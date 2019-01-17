package contract

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
	"strings"
	"sync/atomic"
)

var (
	logExchangeSigHash common.Hash
	logExchangeNFTHash common.Hash
)

const (
	maxWaitingBlock = 50  // come from web3js, if wait more than 50 blocks, transaction time out
	GasLimit            = 3000000
)

type TransactionConfig struct {
	Gaslimit uint64
	GasPrice *big.Int
	key      *keystore.Key
	nonce    uint64
}

type SendOpts struct {
	From common.Address
	To common.Address
	GasLimit uint64
	GasPrice *big.Int
	Value *big.Int
}

type Contract struct {
	Client   *ethclient.Client
	txConfig *TransactionConfig
	Address  common.Address
	ABI      *abi.ABI
}

func init() {

	exchangeSignature := []byte("Exchange(address,uint256)")
	logExchangeSigHash = crypto.Keccak256Hash(exchangeSignature)

	exchangeNFTSignature := []byte("ExchangeNFT(uint256,address,uint256,uint256,bool,bool)")
	logExchangeNFTHash = crypto.Keccak256Hash(exchangeNFTSignature)
}

func NewAuth(key *ecdsa.PrivateKey, nonce uint64, value *big.Int) *bind.TransactOpts {
	auth := bind.NewKeyedTransactor(key)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value
	return auth
}

func (c *Contract) Connect(socket string) {
	if c.Client != nil {
		return
	}
	var err error
	c.Client, err = ethclient.Dial(socket)
	if err != nil {
		panic(err.Error())
	}
}

func (c* Contract) LoadABI(contractAbi string) {
	var err error
	c.ABI= new(abi.ABI)
	*c.ABI, err = abi.JSON(strings.NewReader(contractAbi))
	if err !=nil{
		log.Fatal(err.Error())
	}
}

func (c *Contract) LoadKey(file string, password string) {
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err.Error())
	}
	key, _ := keystore.DecryptKey(jsonBytes, password)
	gasPrice, _ := c.Client.SuggestGasPrice(context.Background())
	gasLimit := uint64(GasLimit)
	initialNonce, _ := c.Client.PendingNonceAt(context.Background(), key.Address)

	c.txConfig = &TransactionConfig{
		Gaslimit: gasLimit,
		GasPrice: gasPrice,
		key:      key,
		nonce:    initialNonce,
	}
}

func (c *Contract) Close() error {
	if c.Client == nil {
		return errors.New("has not connected")
	}
	c.Client.Close()
	return nil
}

func (c *Contract) SendTransaction(tx *types.Transaction) (*types.Transaction, error) {
	err := c.Client.SendTransaction(context.Background(), tx)
	return tx, err
}

func (c *Contract) GetEther(address common.Address) (*big.Int, error) {
	balance, err := c.Client.BalanceAt(context.Background(), address, nil)
	return balance, err
}

func (c *Contract) GetNonce(rawAddress string) (uint64, error) {
	address := common.HexToAddress(rawAddress)
	nonce, err := c.Client.PendingNonceAt(context.Background(), address)
	return nonce, err
}

func (c *Contract) GenerateKeyStore(file string, password string) (common.Address, error) {
	ks := keystore.NewKeyStore(file, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	keyJson, err := ks.Export(account, password, password)
	ioutil.WriteFile(file, keyJson, 0777)
	return account.Address, err
}

func (c *Contract) GetReceiptStatus (txHash common.Hash) (uint64,error) {
	count:= 0
	ch:= make(chan *types.Header)
	sub,err:= c.Client.SubscribeNewHead(context.Background(),ch)
	if err!=nil {
		log.Println(err.Error())
		return 0,err
	}

	for {
		select {
		case err:= <- sub.Err():
			log.Println(err.Error())
			return 0,err
		case <- ch:
			count+=1
			log.Println("getting receipt")
			if count>=maxWaitingBlock {
				return 0, errors.New("transaction time out")
			} else {
				receipt, err:= c.Client.TransactionReceipt(context.Background(),txHash)
				if err == nil {
					log.Println("receipt found")
					if receipt.Status ==0 {
						return receipt.Status, errors.New("transaction revert")
					}
					return receipt.Status, nil
				} else {
				}
			}
		}
	}
}

func (c *Contract) sendWithOpts(funcName string, sendOpts *SendOpts,args ...interface{}) error {
	// assemble transaction
	input,err:= c.ABI.Pack(funcName, args...)
	if err!=nil {
		return err
	}

	nonce:= atomic.AddUint64(&c.txConfig.nonce, 1)
	tx:= types.NewTransaction(nonce-1,sendOpts.To,sendOpts.Value,sendOpts.GasLimit,sendOpts.GasPrice,input)
	signedTx, err:= types.SignTx(tx,types.HomesteadSigner{},c.txConfig.key.PrivateKey)
	if err!=nil {
		return err
	}
	err = c.Client.SendTransaction(context.TODO(), signedTx)
	if err!=nil {
		return err
	}
	_,err = c.GetReceiptStatus(signedTx.Hash())
	return err
}

func (c *Contract) sendWithValue(funcName string, value *big.Int ,args ...interface{}) error {
	return c.sendWithOpts(funcName, &SendOpts{
		To: c.Address,
		GasPrice:c.txConfig.GasPrice,
		GasLimit: c.txConfig.Gaslimit,
		Value: value,
	},args...);
}

func (c *Contract) send(funcName string, args ...interface{}) error {
	return c.sendWithValue(funcName,big.NewInt(0),args...)
}

func (c *Contract) sendWithTargetFeeAndValue (funcName string, TargetGasFee string,value *big.Int ,args ...interface{}) error {
	// Assemble transaction
	input,err:= c.ABI.Pack(funcName, args...)
	if err!=nil {
		return err
	}

	// Estimate gas price
	estimatedGas, err := c.EstimateGas(c.txConfig.key.Address, c.Address,input)
	targetGasFee,_,err:= new(big.Float).Parse(TargetGasFee,10)
	if err!=nil {
		return err
	}
	gasPrice,_:=new(big.Float).Quo(targetGasFee, big.NewFloat(float64(estimatedGas))).Int(nil)

	//Send transaction
	nonce:= atomic.AddUint64(&c.txConfig.nonce, 1)
	tx:= types.NewTransaction(nonce, c.Address,value,c.txConfig.Gaslimit,gasPrice,input)
	chainID, err := c.Client.NetworkID(context.Background())
	if err!=nil {
		return err
	}
	signedTx, err:= types.SignTx(tx,types.NewEIP155Signer(chainID),c.txConfig.key.PrivateKey)
	err = c.Client.SendTransaction(context.Background(), signedTx)
	if err!=nil {
		return err
	}
	_,err = c.GetReceiptStatus(signedTx.Hash())
	return err
}

func (c *Contract) sendWithTargetFee (funcName string, TargetGasFee string,args ...interface{}) error {
	return c.sendWithTargetFeeAndValue(funcName,TargetGasFee,big.NewInt(0),args...)
}

func (c *Contract) call(funcName string, args ...interface{}) ( []byte, error) {
	input,err:= c.ABI.Pack(funcName, args...)
	if err!=nil {
		return nil,err
	}
	rVal, err:= c.Client.CallContract(context.Background(),ethereum.CallMsg{
		To: &c.Address,
		Data:input,
	},nil)
	return rVal,err
}

func (c *Contract) EstimateGas(from common.Address,to common.Address,data []byte) (uint64,error) {
	estimatedGas,err:=c.Client.EstimateGas(context.Background(),ethereum.CallMsg{
		From: from,
		To: &to,
		Data:data,
	})

	return estimatedGas,err
}
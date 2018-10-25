package contract

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"io/ioutil"
	"log"
	"math/big"
	"time"
)

const (
	GasLimit = 3000000
)

var (
	logExchangeSigHash common.Hash
	logExchangeNFTHash common.Hash
)

const (
	privateChainTime = 1*time.Second
	publicChainTime = 1*time.Second
	privateChainTimeOut = 10*time.Second
	publicChainTimeOut = 10*time.Second
)

type Avatar struct {
	TokenId *big.Int `json:"tokenId"`
	Gene        *big.Int `json:"gene"`
	AvatarLevel *big.Int `json:"level"`
	Weaponed    bool     `json:"weaponed"`
	Armored bool  `json:"armored"`
}

type LogExchange struct{
	User common.Address
	Amount *big.Int

	TxHash common.Hash
}

type LogExchangeNFT struct {
	TokenID *big.Int
	Owner common.Address
	Gene *big.Int
	AvatarLevel *big.Int
	Weaponed bool
}

type TransactionConfig struct {
	Gaslimit uint64
	GasPrice *big.Int
	key *keystore.Key
	nonce uint64
}

type Contract struct {
	Client *ethclient.Client
	txConfig *TransactionConfig
	Address common.Address
}

func init(){
	exchangeSignature :=[]byte("Exchange(address,uint256)")
	logExchangeSigHash = crypto.Keccak256Hash(exchangeSignature)

	exchangeNFTSignature:= []byte("ExchangeNFT(uint256,address,uint256,uint256,bool)")
	logExchangeNFTHash = crypto.Keccak256Hash(exchangeNFTSignature)
}

func NewAuth(key *ecdsa.PrivateKey, nonce uint64, value * big.Int) *bind.TransactOpts{
	auth:= bind.NewKeyedTransactor(key)
	auth.Nonce =  big.NewInt(int64(nonce))
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

func (c *Contract) LoadKey(file string, password string) {
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err.Error())
	}
	key, _ := keystore.DecryptKey(jsonBytes, password)
	gasPrice, _ := c.Client.SuggestGasPrice(context.Background())
	gasLimit := uint64(GasLimit)
	initialNonce, _ := c.Client.PendingNonceAt(context.Background(), key.Address)

	c.txConfig = & TransactionConfig{
		Gaslimit:gasLimit,
		GasPrice:gasPrice,
		key: key,
		nonce:initialNonce,
	}
}

func (c *Contract) Close() error {
	if c.Client == nil {
		return errors.New("has not connected")
	}
	c.Client.Close()
	return nil
}

func (c *Contract) SendTransaction(rawTx string) (*types.Transaction,error) {
	rawTxBytes, err := hex.DecodeString(rawTx)
	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)
	err = c.Client.SendTransaction(context.Background(), tx)
	return tx,err
}

func (c *Contract) GetEther(rawAccount string) (*big.Int, error) {
	account := common.HexToAddress(rawAccount)
	balance, err := c.Client.BalanceAt(context.Background(), account, nil)
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

func (c *Contract) EventWatcher() (chan types.Log, <-chan error) {
	log.Println("start to watching")
	query:=ethereum.FilterQuery{
		Addresses:[]common.Address{c.Address},
	}
	logs:=make(chan types.Log)
	sub,err:= c.Client.SubscribeFilterLogs(context.Background(), query, logs)

	if err != nil {
		log.Fatal(err)
	}

	return logs, sub.Err()
}
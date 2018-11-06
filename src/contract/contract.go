package contract

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"math/big"
	"time"
)

var (
	logExchangeSigHash common.Hash
	logExchangeNFTHash common.Hash
)

const (
	privateChainTime    = 1 * time.Second
	publicChainTime     = 1 * time.Second
	privateChainTimeOut = 10 * time.Second
	publicChainTimeOut  = 10 * time.Second
	GasLimit            = 3000000
)

type TransactionConfig struct {
	Gaslimit uint64
	GasPrice *big.Int
	key      *keystore.Key
	nonce    uint64
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
package contract

import (
	"context"
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"io/ioutil"
	"math/big"
	"token/privateSlot"
	"token/publicSlot"
)

type Contract struct {
	Client *ethclient.Client
	Auth   *bind.TransactOpts
}

type Pbc struct {
	*Contract
	Instance *publicSlot.PublicSlot
}

type Pvc struct {
	*Contract
	Instance *privateSlot.PrivateSlot
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
	key, err := keystore.DecryptKey(jsonBytes, password)
	gasPrice, err := c.Client.SuggestGasPrice(context.Background())
	auth := bind.NewKeyedTransactor(key.PrivateKey)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	c.Auth = auth
}

func (c *Contract) Close() error {
	if c.Client == nil {
		return errors.New("has not connected")
	}
	c.Client.Close()
	return nil
}

func (c *Contract) SendTransaction(rawTx string) error {
	rawTxBytes, err := hex.DecodeString(rawTx)
	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)
	err = c.Client.SendTransaction(context.Background(), tx)
	return err
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

func (p *Pbc) LoadContract(rawAddress string) {
	address := common.HexToAddress(rawAddress)
	p.Instance, _ = publicSlot.NewPublicSlot(address, p.Client)
}

func (p *Pbc) GetToken(rawAddress string) (*big.Int, error) {
	address := common.HexToAddress(rawAddress)
	token, err := p.Instance.BalanceOf(&bind.CallOpts{}, address)
	return token, err
}

func (p *Pbc) Deploy(initialSupply *big.Int, requiredSignatures *big.Int, authorities []common.Address) (common.Address, error) {
	p.Auth.Value = big.NewInt(0)
	nonce, err := p.Client.PendingNonceAt(context.Background(), p.Auth.From)
	p.Auth.Nonce = big.NewInt(int64(nonce))
	address, _, _, err := publicSlot.DeployPublicSlot(p.Auth, p.Client, initialSupply, requiredSignatures, authorities)
	return address, err
}

func (p *Pvc) LoadContract(rawAddress string) {
	address := common.HexToAddress(rawAddress)
	p.Instance, _ = privateSlot.NewPrivateSlot(address, p.Client)
}

func (p *Pvc) GetToken(rawAddress string) (*big.Int, error) {
	address := common.HexToAddress(rawAddress)
	token, err := p.Instance.BalanceOf(&bind.CallOpts{}, address)
	return token, err
}

func (p *Pvc) Consume(rawAddress string, amount *big.Int) error {
	p.Auth.Value = big.NewInt(0)
	nonce, err := p.Client.PendingNonceAt(context.Background(), p.Auth.From)
	p.Auth.Nonce = big.NewInt(int64(nonce))
	address := common.HexToAddress(rawAddress)
	_, err = p.Instance.Consume(p.Auth, address, amount)
	return err
}

func (p *Pvc) Pay(rawAddress string, amount *big.Int, transactionHash [32]byte) error {
	p.Auth.Value = big.NewInt(0)
	nonce, err := p.Client.PendingNonceAt(context.Background(), p.Auth.From)
	p.Auth.Nonce = big.NewInt(int64(nonce))
	address := common.HexToAddress(rawAddress)
	_, err = p.Instance.Pay(p.Auth, address, amount, transactionHash)
	return err
}

func (p *Pvc) Award(rawAddress string, amount *big.Int) error {
	p.Auth.Value = big.NewInt(0)
	nonce, err := p.Client.PendingNonceAt(context.Background(), p.Auth.From)
	p.Auth.Nonce = big.NewInt(int64(nonce))
	address := common.HexToAddress(rawAddress)
	_, err = p.Instance.Award(p.Auth, address, amount)
	return err
}

func (p *Pvc) Deploy(initialSupply *big.Int, requiredSignatures *big.Int, authorities []common.Address) (common.Address, error) {
	p.Auth.Value = big.NewInt(0)
	nonce, err := p.Client.PendingNonceAt(context.Background(), p.Auth.From)
	p.Auth.Nonce = big.NewInt(int64(nonce))
	address, _, _, err := privateSlot.DeployPrivateSlot(p.Auth, p.Client, initialSupply, requiredSignatures, authorities)
	return address, err
}

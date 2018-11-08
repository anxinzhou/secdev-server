package contract

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"strings"
	"sync/atomic"
	"time"
	"token/publicSlot"
)

type Pbc struct {
	Contract
	Instance *publicSlot.PublicSlot
}

func NewPbc(port string, keystore string, password string, contractAdd string) *Pbc {
	pbc := new(Pbc)
	pbc.Connect(port)
	pbc.LoadKey(keystore, password)
	pbc.LoadContract(contractAdd)
	pbc.LoadABI()
	return pbc
}

func (p* Pbc) LoadABI() {
	var err error
	p.ABI = new(abi.ABI)
	*p.ABI, err = abi.JSON(strings.NewReader(string(publicSlot.PublicSlotABI)))
	if err !=nil{
		log.Fatal(err.Error())
	}
}

func (p *Pbc) LoadContract(rawAddress string) {
	address := common.HexToAddress(rawAddress)
	p.Address = address
	p.Instance, _ = publicSlot.NewPublicSlot(address, p.Client)
}

func (p *Pbc) GetToken(address common.Address) (*big.Int, error) {
	token, err := p.Instance.BalanceOf(nil, address)
	return token, err
}

func (p *Pbc) GetExchangeRate() (string, error) {
	exchangeRate,err:=p.Instance.ExchangeRate(nil)
	if err!=nil {
		return "", err
	}
	exchangeBase,err:=p.Instance.ExchangeBase(nil)
	if err!=nil {
		return "",err
	}
	rate:=new(big.Float).Quo(new(big.Float).SetInt(exchangeRate),new(big.Float).SetInt(exchangeBase)).String()
	return rate,nil
}

func (p *Pbc) ExchangeForToken(address common.Address, amount *big.Int) (*types.Transaction,error) {
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	auth.Value = amount
	auth.GasLimit = p.txConfig.Gaslimit
	tx,err:= p.Instance.ExchangeForToken(auth,address)
	return tx,err
}

func (p *Pbc) ExchangeForEther(address common.Address, amount *big.Int) (*types.Transaction,error){
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	auth.GasLimit = p.txConfig.Gaslimit
	tx,err:=p.Instance.ExchangeForEther(auth,address,amount)
	return tx,err
}

func (p *Pbc) GetReceiptStatus (txHash common.Hash) (uint64,error) {
	count := time.Second
	for {
		time.Sleep(publicChainTime)
		receipt, err :=p.Client.TransactionReceipt(context.Background(),txHash)
		if err==nil {
			if receipt.Status==0{
				return receipt.Status, errors.New("transaction revert")
			}
			return receipt.Status, nil
		}
		count +=time.Second
		if count == publicChainTimeOut {
			break
		}
	}
	return 0, errors.New("Time out, can not get transaction status")
}

func (p *Pbc) Mint(address common.Address, amount *big.Int) (error) {
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	auth.GasLimit = p.txConfig.Gaslimit
	tx, err := p.Instance.Mint(auth, address, amount)
	_,err = p.GetReceiptStatus(tx.Hash())
	return err
}

func (p *Pbc) Burn(address common.Address, amount *big.Int) (error) {
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	auth.GasLimit = p.txConfig.Gaslimit
	tx, err := p.Instance.Burn(auth, address, amount)
	_,err = p.GetReceiptStatus(tx.Hash())
	return err
}

func (p *Pbc) Reward(address common.Address, amount *big.Int) (error) {
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	auth.GasLimit = p.txConfig.Gaslimit
	tx, err := p.Instance.Reward(auth, address, amount)
	_,err = p.GetReceiptStatus(tx.Hash())
	return err
}
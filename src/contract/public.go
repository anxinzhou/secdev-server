package contract

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"strings"
	"sync/atomic"
	"token/publicSlot"
)

const (
	TargetGasFee = "50000000000000"
	TxFee = "0.00005" //ether
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

	data,err:= p.ABI.Pack("exchangeForToken",address)
	if err!=nil {
		return nil,err
	}
	gasPrice,err := p.EstimateGasPrice(address,p.Contract.Address,data)
	if err!=nil {
		return nil,err
	}
	auth.GasPrice = gasPrice

	tx,err:= p.Instance.ExchangeForToken(auth,address)
	return tx,err
}

func (p* Pbc) Create(address common.Address, tokenID *big.Int) (error) {
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	auth.GasLimit = p.txConfig.Gaslimit

	data,err:= p.ABI.Pack("create",address,tokenID)
	if err!=nil {
		return err
	}

	gasPrice, err:= p.EstimateGasPrice(address,p.Contract.Address,data)
	if err!=nil {
		return err
	}
	auth.GasPrice = gasPrice

	tx,err:= p.Instance.Create(auth,address,tokenID)
	_,err = p.GetReceiptStatus(tx.Hash())
	return err
}

func (p *Pbc) ExchangeForEther(address common.Address, amount *big.Int) (*types.Transaction,error){
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	auth.GasLimit = p.txConfig.Gaslimit

	data,err:= p.ABI.Pack("exchangeForEther",address,amount)
	if err!=nil {
		return nil,err
	}
	gasPrice,err := p.EstimateGasPrice(address,p.Contract.Address,data)
	if err!=nil {
		return nil,err
	}
	auth.GasPrice = gasPrice

	tx,err:=p.Instance.ExchangeForEther(auth,address,amount)
	return tx,err
}

//func (p *Pbc) GetReceiptStatus (txHash common.Hash) (uint64,error) {
//	time.Sleep(publicChainTime)
//	time.Sleep(publicChainTime)
//	return 1,nil
//}

func (p *Pbc) GetReceiptStatus (txHash common.Hash) (uint64,error) {
	log.Println("geting transaction receipt for "+txHash.String())
	count:= 0
	ch:= make(chan *types.Header)
	sub,err:= p.Client.SubscribeNewHead(context.Background(),ch)
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
				if count>=maxWaitingBlock {
					return 0, errors.New("transaction time out")
				} else {
					receipt, err:= p.Client.TransactionReceipt(context.Background(),txHash)
					if err == nil {
						if receipt.Status ==0 {
							return receipt.Status, errors.New("transaction revert")
						}
						return receipt.Status, nil
					} else {
						log.Println(err.Error())
					}
				}
		}
	}
}

func (p *Pbc) Mint(address common.Address, amount *big.Int) (error) {
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	auth.GasLimit = p.txConfig.Gaslimit

	data,err:= p.ABI.Pack("mint",address,amount)
	if err!=nil {
		return err
	}
	gasPrice,err := p.EstimateGasPrice(address,p.Contract.Address,data)
	if err!=nil {
		return err
	}
	auth.GasPrice = gasPrice

	tx, err := p.Instance.Mint(auth, address, amount)
	_,err = p.GetReceiptStatus(tx.Hash())
	return err
}

func (p *Pbc) Burn(address common.Address, amount *big.Int) (error) {
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	auth.GasLimit = p.txConfig.Gaslimit

	data,err:= p.ABI.Pack("burn",address,amount)
	if err!=nil {
		return err
	}
	gasPrice,err := p.EstimateGasPrice(address,p.Contract.Address,data)
	if err!=nil {
		return err
	}
	auth.GasPrice = gasPrice

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

func (p *Pbc) EstimateGasPrice(from common.Address,to common.Address,data []byte) (* big.Int,error) {
	esimatedGas,err:=p.Client.EstimateGas(context.Background(),ethereum.CallMsg{
		From: from,
		To: &to,
		Data:data,
	})

	targetGasFee,_,err:= new(big.Float).Parse(TargetGasFee,10)
	if err!=nil {
		return big.NewInt(0),err
	}

	gasPrice,_:=new(big.Float).Quo(targetGasFee, big.NewFloat(float64(esimatedGas))).Int(nil)
	log.Println(gasPrice,esimatedGas)
	return gasPrice,nil
}
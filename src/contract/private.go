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
	"token/privateSlot"
)

type Pvc struct {
	Contract
	Instance *privateSlot.PrivateSlot
}

func NewPvc(port string, keystore string, password string, contractAdd string) *Pvc {
	pvc := new(Pvc)
	pvc.Connect(port)
	pvc.LoadKey(keystore, password)
	pvc.LoadContract(contractAdd)
	pvc.LoadABI()
	return pvc
}

func (p* Pvc) LoadABI() {
	var err error
	p.ABI= new(abi.ABI)
	*p.ABI, err = abi.JSON(strings.NewReader(string(privateSlot.PrivateSlotABI)))
	if err !=nil{
		log.Fatal(err.Error())
	}
}

func (p *Pvc) LoadContract(rawAddress string) {
	address := common.HexToAddress(rawAddress)
	p.Address = address
	p.Instance, _ = privateSlot.NewPrivateSlot(address, p.Client)
}

func (p *Pvc) GetToken(address common.Address) (*big.Int, error) {
	token, err := p.Instance.BalanceOf(nil, address)
	return token, err
}

func (p *Pvc) Consume(address common.Address, amount *big.Int) (error) {
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	auth.GasLimit = p.txConfig.Gaslimit
	tx, err := p.Instance.Consume(auth, address, amount)
	_,err = p.GetReceiptStatus(tx.Hash())
	return err
}

func (p *Pvc) Reward(address common.Address, amount *big.Int) (error) {
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	auth.GasLimit = p.txConfig.Gaslimit
	tx, err := p.Instance.Reward(auth, address, amount)
	_,err = p.GetReceiptStatus(tx.Hash())
	return err
}

func (p *Pvc) Mint(address common.Address, amount *big.Int) (error) {
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	auth.GasLimit = p.txConfig.Gaslimit
	tx, err := p.Instance.Mint(auth, address, amount)
	_,err = p.GetReceiptStatus(tx.Hash())
	return err
}

func (p *Pvc) Burn(address common.Address, amount *big.Int) (error) {
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	auth.GasLimit = p.txConfig.Gaslimit
	tx, err := p.Instance.Burn(auth, address, amount)
	_,err = p.GetReceiptStatus(tx.Hash())
	return err
}

func (p *Pvc) GetReceiptStatus (txHash common.Hash) (uint64,error) {
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
package contract

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	return pbc
}

func (p *Pbc) LoadContract(rawAddress string) {
	address := common.HexToAddress(rawAddress)
	p.Address = address
	p.Instance, _ = publicSlot.NewPublicSlot(address, p.Client)
}

func (p *Pbc) GetToken(rawAddress string) (*big.Int, error) {
	address := common.HexToAddress(rawAddress)
	token, err := p.Instance.BalanceOf(&bind.CallOpts{}, address)
	return token, err
}

func (p *Pbc) AddGameMachine(rawMachine string) error {
	var err error
	machine := common.HexToAddress(rawMachine)
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	tx, err := p.Instance.AddGameMachine(auth, machine)
	_,err =p.GetReceiptStatus(tx.Hash())

	return err
}

func (p *Pbc) Pay(vs [] uint8, rs [][32]byte, ss[][32]byte, message []byte) error {
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	log.Println("public nonce:", nonce)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	auth.GasLimit = p.txConfig.Gaslimit
	tx, err := p.Instance.Pay(auth, vs, rs, ss, message)
	_,err =p.GetReceiptStatus(tx.Hash())

	return err
}

func (p *Pbc) PayNFT(tokenID *big.Int, address common.Address) error {
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	auth.GasLimit = p.txConfig.Gaslimit

	tx, err:= p.Instance.PayNFT(auth, tokenID, address)

	_, err = p.GetReceiptStatus(tx.Hash())

	return err
}

func (p *Pbc) SendTransaction(rawTx string) error {
	tx,err:= p.Contract.SendTransaction(rawTx)
	_, err = p.GetReceiptStatus(tx.Hash())
	return err
}

func (p *Pbc) ExchangeNFTHandler(pvc *Pvc, nft *LogExchangeNFT){
	log.Println("receive an exchangeNFT")
	err:= pvc.PayNFT(nft.TokenID,nft.Owner)
	log.Println(nft.TokenID.Uint64(), nft.Owner.String())
	if err!=nil {
		log.Println("pay nft fail in privatechain", err.Error())
		p.PayNFT(nft.TokenID,nft.Owner)
	}

}

//func (p *Pbc) Deploy(initialSupply *big.Int, requiredSignatures *big.Int, authorities []common.Address) (common.Address, error) {
//	var err error
//
//	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
//	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce, big.NewInt(0))
//
//	address, _, _, err := publicSlot.DeployPublicSlot(auth, p.Client, initialSupply, requiredSignatures, authorities)
//	return address, err
//}

func (p *Pbc) ExchangeHandler(pvc *Pvc, exchangeEvent *LogExchange){
	log.Println("receive an exchange event from public chain")
	log.Println(exchangeEvent.User.String())
	log.Println(exchangeEvent.Amount.Uint64())
	err:=pvc.Pay(exchangeEvent.User, exchangeEvent.Amount, exchangeEvent.TxHash)
	if err!=nil {
		log.Println(err.Error())
	}
}

func (p *Pbc) EventReceiver(pvc *Pvc){
	logs,eventError:=p.EventWatcher()

	contractAbi, err:= abi.JSON(strings.NewReader(string(publicSlot.PublicSlotABI)))
	if err !=nil{
		log.Fatal(err.Error())
	}


	for {
		select {
		case err:=<-eventError:
			log.Println(err.Error())
		case vLog := <-logs:
			switch vLog.Topics[0].Hex() {
			case logExchangeSigHash.Hex():
				var exchangeEvent LogExchange
				contractAbi.Unpack(&exchangeEvent, "Exchange", vLog.Data)
				log.Println(len(vLog.Data))
				log.Println(vLog.Data)
				log.Println(vLog.TxHash.String())
				exchangeEvent.TxHash = vLog.TxHash
				go p.ExchangeHandler(pvc, &exchangeEvent)
			case logExchangeNFTHash.Hex():
				var exchangeNFTEvent LogExchangeNFT
				contractAbi.Unpack(&exchangeNFTEvent, "ExchangeNFT",vLog.Data)
				go p.ExchangeNFTHandler(pvc,&exchangeNFTEvent)
			}
		}
	}
}

func (p *Pbc) GetReceiptStatus (txHash common.Hash) (uint64,error) {
	count := time.Second
	for {
		time.Sleep(privateChainTime)
		receipt, err :=p.Client.TransactionReceipt(context.Background(),txHash)
		if err==nil {
			if receipt.Status==0{
				return receipt.Status, errors.New("transaction revert")
			}
			return receipt.Status, nil
		}
		count +=time.Second
		if count == privateChainTimeOut {
			break
		}
	}
	return 0, errors.New("Time out, can not get transaction status")
}
package contract

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/goware/disque"
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
	if err!=nil {
		log.Fatal(err.Error() ," pay in public fail")
	}
	_,err =p.GetReceiptStatus(tx.Hash())

	return err
}

//func (p *Pbc) PayNFT(tokenID *big.Int, address common.Address) error {
//	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
//	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
//	auth.GasLimit = p.txConfig.Gaslimit
//
//	tx, err:= p.Instance.PayNFT(auth, tokenID, address)
//
//	_, err = p.GetReceiptStatus(tx.Hash())
//
//	return err
//}

func (p *Pbc) SendTransaction(tx *types.Transaction) error {
	tx,err:= p.Contract.SendTransaction(tx)
	_, err = p.GetReceiptStatus(tx.Hash())
	return err
}

func (p *Pbc) GetAvatar(tokenId *big.Int) (*Avatar, error){

	avatar, err :=p.Instance.Avatar(nil,tokenId)
	return &Avatar{
		TokenId:tokenId,
		Gene:avatar.Gene,
		AvatarLevel:avatar.AvatarLevel,
		Weaponed: avatar.Weaponed,
		Armored: avatar.Armored,
	},err
}

func (p *Pbc) ExchangeNFTHandler(pvc *Pvc,jobs *disque.Pool, nft *LogExchangeNFT){
	input, _ := pvc.ABI.Pack("payNFT",nft.TokenID,nft.Owner,nft.Gene,nft.AvatarLevel,nft.Weaponed,nft.Armored)

	tx:=types.NewTransaction( 0, pvc.Address, big.NewInt(0), pvc.txConfig.Gaslimit, pvc.txConfig.GasPrice, input)
	txWrapper, _:= tx.MarshalJSON()
	jobs.Add(nft.Owner.String()+nft.TxHash.String()+string(txWrapper),PvcPayNFTQueue)
}

//func (p *Pbc) Deploy(initialSupply *big.Int, requiredSignatures *big.Int, authorities []common.Address) (common.Address, error) {
//	var err error
//
//	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
//	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
//
//	address, _, _, err := publicSlot.DeployPublicSlot(auth, p.Client, initialSupply, requiredSignatures, authorities)
//	return address, err
//}

func (p *Pbc) ExchangeHandler(pvc *Pvc, jobs *disque.Pool,exchangeEvent *LogExchange){
	input, _:= pvc.ABI.Pack("pay",exchangeEvent.User, exchangeEvent.Amount, exchangeEvent.TxHash)

	//nonce:= atomic.AddUint64(&pbc.txConfig.nonce, 1)
	tx:=types.NewTransaction( 0, pvc.Address, big.NewInt(0), pvc.txConfig.Gaslimit, pvc.txConfig.GasPrice, input)
	txWrapper, _:= tx.MarshalJSON()
	jobs.Add(exchangeEvent.User.String() +exchangeEvent.TxHash.String()+ string(txWrapper),PvcPayQueue)
}

func (p *Pbc) EventReceiver(pvc *Pvc, jobs *disque.Pool){
	logs,eventError:=p.EventWatcher()

	for {
		select {
		case err:=<-eventError:
			log.Println(err.Error())
		case vLog := <-logs:
			switch vLog.Topics[0].Hex() {
			case logExchangeSigHash.Hex():
				var exchangeEvent LogExchange
				p.ABI.Unpack(&exchangeEvent, "Exchange", vLog.Data)
				exchangeEvent.TxHash = vLog.TxHash
				go p.ExchangeHandler(pvc, jobs,&exchangeEvent)
			case logExchangeNFTHash.Hex():
				var exchangeNFTEvent LogExchangeNFT
				p.ABI.Unpack(&exchangeNFTEvent, "ExchangeNFT",vLog.Data)
				exchangeNFTEvent.TxHash = vLog.TxHash
				go p.ExchangeNFTHandler(pvc,jobs,&exchangeNFTEvent)
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

func (p *Pbc) ProcessJob(job *disque.Job) (*types.Transaction,error) {
	tx,_ := p.Contract.ProcessJob(job)
	_,err := p.GetReceiptStatus(tx.Hash())
	return tx,err
}
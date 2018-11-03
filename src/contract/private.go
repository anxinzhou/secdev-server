package contract

import (
	"bytes"
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/goware/disque"
	"log"
	"math/big"
	"strings"
	"sync/atomic"
	"time"
	"token/privateSlot"
	"utils"
)

type Pvc struct {
	Contract
	Instance *privateSlot.PrivateSlot
}

type LogCollectedSignatures struct {
	AuthorityMachineResponsibleForRelay common.Address
	MessageHash common.Hash
}

var (
	logCollectedSignaturesHash common.Hash
)

func init(){
	collectedSignaturesSig := []byte("CollectedSignatures(address,bytes32)")
	logCollectedSignaturesHash = crypto.Keccak256Hash(collectedSignaturesSig)
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

func (p *Pvc) GetToken(rawAddress string) (*big.Int, error) {
	address := common.HexToAddress(rawAddress)
	token, err := p.Instance.BalanceOf(nil, address)
	return token, err
}

func (p *Pvc) Consume(rawAddress string, amount *big.Int) error {
	var err error

	address := common.HexToAddress(rawAddress)
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))

	tx, err := p.Instance.Consume(auth, address, amount)
	_,err =p.GetReceiptStatus(tx.Hash())

	return err
}

func (p *Pvc) Pay(address common.Address, amount *big.Int, transactionHash common.Hash) error {
	var err error

	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	auth.GasLimit = p.txConfig.Gaslimit

	log.Println(amount.Uint64())
	log.Println(address.String())
	log.Println(transactionHash.String())

	tx, err := p.Instance.Pay(auth, address, amount, transactionHash)

	_,err =p.GetReceiptStatus(tx.Hash())

	return err
}

func (p *Pvc) Reward(rawAddress string, amount *big.Int) error {
	var err error

	address := common.HexToAddress(rawAddress)
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	tx, err := p.Instance.Reward(auth, address, amount)
	if err!=nil {
		log.Fatal(err.Error(),"reward fail")
	}
	_,err =p.GetReceiptStatus(tx.Hash())

	return err
}

func (p *Pvc) AddGameMachine(rawMachine string) error {
	var err error
	machine := common.HexToAddress(rawMachine)
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	tx, err := p.Instance.AddGameMachine(auth, machine)
	_,err =p.GetReceiptStatus(tx.Hash())

	return err
}

func (p *Pvc) SubmitSignature(message []byte, signature []byte) error {
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	auth.GasLimit = p.txConfig.Gaslimit

	tx, err:= p.Instance.SubmitSignature(auth,message,signature)

	if err!=nil {
		log.Println(err.Error())
	}
	_,err =p.GetReceiptStatus(tx.Hash())

	return err
}

func (p *Pvc) OwnedTokens(rawAddress string) (*big.Int,error) {
	address := common.HexToAddress(rawAddress)
	ownedAvatar,err:=p.Instance.OwnedAvatars(nil,address)
	if err!=nil {
		log.Println("get owned tokens fail",err.Error())
	}
	return ownedAvatar, err
}

func (p *Pvc) Mint(rawAddress string,tokenId *big.Int) error {
	address := common.HexToAddress(rawAddress)

	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	auth.GasLimit = p.txConfig.Gaslimit

	tx , err:= p.Instance.Mint(auth, address, tokenId)

	_,err =p.GetReceiptStatus(tx.Hash())

	return err

}

func (p *Pvc) Upgrade (tokenId *big.Int) error{

	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	auth.GasLimit = p.txConfig.Gaslimit

	tx,err :=p.Instance.Upgrade(auth,tokenId)

	_,err = p.GetReceiptStatus(tx.Hash())

	return err
}

func (p *Pvc) EquipWeapon (rawAddress string,tokenId *big.Int) error {
	address := common.HexToAddress(rawAddress)
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	auth.GasLimit = p.txConfig.Gaslimit

	tx,err :=p.Instance.EquipWeapon(auth,tokenId,address)

	_,err = p.GetReceiptStatus(tx.Hash())

	return err
}

func (p*Pvc) OwnerOf(tokenId *big.Int) common.Address{
	owner, _:=p.Instance.OwnerOf(nil, tokenId)
	return owner
}

func (p *Pvc) EquipArmor(rawAddress string, tokenId *big.Int) error {
	address := common.HexToAddress(rawAddress)

	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	auth.GasLimit = p.txConfig.Gaslimit

	tx,err :=p.Instance.EquipArmor(auth,tokenId,address)

	_,err = p.GetReceiptStatus(tx.Hash())

	return err
}

func (p *Pvc) GetAvatar(tokenId *big.Int) (*Avatar, error){

	avatar, err :=p.Instance.Avatar(nil,tokenId)
	return &Avatar{
		TokenId:tokenId,
		Gene:avatar.Gene,
		AvatarLevel:avatar.AvatarLevel,
		Weaponed: avatar.Weaponed,
		Armored: avatar.Armored,
	},err
}

//func (p *Pvc) PayNFT(tokenID *big.Int, address common.Address) error {
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

func (p *Pvc) ExchangeHandler(jobs * disque.Pool,exchangeEvent *LogExchange) {
	log.Println("receive an exchange event from private chain")

	message := bytes.Join([][]byte{ exchangeEvent.TxHash[:],
		exchangeEvent.User[:],
		utils.FillZero(exchangeEvent.Amount.Bytes(), 32)}, []byte(""))
	//hash:= crypto.Keccak256Hash([]byte("\x19Ethereum Signed Message:\n"),[]byte(strconv.Itoa(len(message))),message)
	hash:= crypto.Keccak256Hash(message)
	signature, _:=crypto.Sign(hash.Bytes(), p.txConfig.key.PrivateKey)

	input, _:= p.ABI.Pack("submitSignature",message, signature)
	tx:=types.NewTransaction( 0, p.Address, big.NewInt(0), p.txConfig.Gaslimit, p.txConfig.GasPrice, input)
	txWrapper, _:= tx.MarshalJSON()
	log.Println(exchangeEvent.User.String())
	log.Println(exchangeEvent.TxHash.String())
	jobs.Add(exchangeEvent.User.String()+exchangeEvent.TxHash.String()+string(txWrapper), SubmitSignatureQueue)
}

func (p *Pvc) CollectedSignaturesHandler(pbc *Pbc,jobs * disque.Pool, collectedSignaturesEvent *LogCollectedSignatures) {
	log.Println("receive an collectedSignatures event from private chain")

	if collectedSignaturesEvent.AuthorityMachineResponsibleForRelay != p.txConfig.key.Address {
		log.Println("Collected signature should be submited by", collectedSignaturesEvent.AuthorityMachineResponsibleForRelay)
		return
	}
	message, _ := p.Instance.Message(nil, collectedSignaturesEvent.MessageHash)
	rqSigs, _:= p.Instance.RequiredSignatures(nil)

	vs:=make([]uint8,0,rqSigs.Int64())
	rs:=make([][32]byte,0,rqSigs.Int64())
	ss:=make([][32]byte,0,rqSigs.Int64())

	for i:=0;i<int(rqSigs.Int64());i++ {
		sig, err := p.Instance.Signature(nil,collectedSignaturesEvent.MessageHash,big.NewInt(int64(i)))
		if err!=nil {
			log.Println("test", err.Error())
		}

		var (
			r [32]byte
			v uint8
			s [32]byte
		)

		copy(r[:], sig[:32])
		v = sig[64]
		copy(s[:], sig[32:64])

		vs=append(vs,v)
		rs=append(rs,r)
		ss=append(ss,s)
	}

	input, _:= pbc.ABI.Pack("pay",vs, rs, ss, message)
	tx:=types.NewTransaction( 0, pbc.Address, big.NewInt(0), pbc.txConfig.Gaslimit, pbc.txConfig.GasPrice, input)
	txWrapper, _:= tx.MarshalJSON()
	var address common.Address
	var txHash common.Hash
	copy(txHash[:],message[:32])
	copy(address[:],message[32:52])
	jobs.Add(address.String()+txHash.String()+string(txWrapper), PbcPayQueue)
}

func (p *Pvc) ExchangeNFTHandler(pbc *Pbc,jobs *disque.Pool,nft *LogExchangeNFT){

	// pack method to raw byte
	input, _ := pbc.ABI.Pack("payNFT",nft.TokenID,nft.Owner,nft.Gene,nft.AvatarLevel,nft.Weaponed,nft.Armored)

	//nonce:= atomic.AddUint64(&pbc.txConfig.nonce, 1)
	tx:=types.NewTransaction( 0, pbc.Address, big.NewInt(0), pbc.txConfig.Gaslimit, pbc.txConfig.GasPrice, input)
	txWrapper, _:= tx.MarshalJSON()
	log.Println(nft.Owner.String())
	log.Println(nft.TxHash.String())
	log.Println(len(nft.Owner.String())+len(nft.TxHash.String()))
	jobs.Add(nft.Owner.String()+nft.TxHash.String()+string(txWrapper),PbcPayNFTQueue)

}


func (p *Pvc) SendTransaction(tx *types.Transaction) error {
	tx,err:= p.Contract.SendTransaction(tx)
	_, err = p.GetReceiptStatus(tx.Hash())
	return err
}

func (p *Pvc) GetReceiptStatus (txHash common.Hash) (uint64,error) {
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

func (p *Pvc) ProcessJob(job *disque.Job) (*types.Transaction,error) {
	tx,_ := p.Contract.ProcessJob(job)
	_,err := p.GetReceiptStatus(tx.Hash())
	return tx, err
}
//func (p *Pvc) Deploy(initialSupply *big.Int, requiredSignatures *big.Int, authorities []common.Address) (common.Address, error) {
//	var err error
//
//	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
//	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
//
//	address, _, _, err := privateSlot.DeployPrivateSlot(auth, p.Client, initialSupply, requiredSignatures, authorities)
//	return address, err
//}

func (p *Pvc) EventReceiver(pbc *Pbc, jobs *disque.Pool){
	log.Println("start private watcher")
	logs,eventError:=p.EventWatcher()

	log.Println("Waiting for an event")
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
				go p.ExchangeHandler(jobs,&exchangeEvent)

			case logCollectedSignaturesHash.Hex():

				var collectedSignaturesEvent LogCollectedSignatures
				p.ABI.Unpack(&collectedSignaturesEvent, "CollectedSignatures", vLog.Data)
				go p.CollectedSignaturesHandler(pbc, jobs,&collectedSignaturesEvent)

			case logExchangeNFTHash.Hex():
				var exchangeNFTEvent LogExchangeNFT
				p.ABI.Unpack(&exchangeNFTEvent, "ExchangeNFT",vLog.Data)
				exchangeNFTEvent.TxHash = vLog.TxHash
				go p.ExchangeNFTHandler(pbc, jobs, &exchangeNFTEvent)
			}
		}
	}
}
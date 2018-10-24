package contract

import (
	"bytes"
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"strings"
	"sync/atomic"
	"time"
	"token/privateSlot"
)

type Pvc struct {
	Contract
	Instance *privateSlot.PrivateSlot
}

type LogCollectedSignatures struct {
	AuthorityMachineResponsibleForRelay common.Address
	MessageHash common.Hash
}

//event DepositConfirmation(address recipient, uint256 value, bytes32 transactoinHash);
//event Login(address indexed privateAdd, address indexed publicAdd, uint time);
//event Sync(address []  _from, address [] _to, uint [] _amount, uint [] time);
//event SyncConfirmation(address indexed _from, address indexed _to, uint time);
//event Pay(address indexed _user, uint _amount , bytes32 transactoinHash);
//event Award(address indexed _user, uint _amount );
//event Consume(address indexed _user, uint _amount);
//event ExchangeToken(address user, uint amount);

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
	return pvc
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

func (p *Pvc) EquipWeapon (tokenId *big.Int) error {
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
	auth.GasLimit = p.txConfig.Gaslimit

	tx,err :=p.Instance.EquipWeapon(auth,tokenId)

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
	},err
}

func (p *Pvc) ExchangeHandler(exchangeEvent *LogExchange) error {
	log.Println("receive an exchange event from private chain")
	log.Println(exchangeEvent.User.String())
	log.Println(exchangeEvent.Amount.Uint64())
	message := bytes.Join([][]byte{ exchangeEvent.TxHash[:],
		exchangeEvent.User[:],
		fillZero(exchangeEvent.Amount.Bytes(), 32)}, []byte(""))
	//hash:= crypto.Keccak256Hash([]byte("\x19Ethereum Signed Message:\n"),[]byte(strconv.Itoa(len(message))),message)
	hash:= crypto.Keccak256Hash(message)
	signature, err:=crypto.Sign(hash.Bytes(), p.txConfig.key.PrivateKey)
	log.Println("signature:",signature)

	p.SubmitSignature(message, signature)

	return err
}

func (p *Pvc) CollectedSignaturesHandler(pbc *Pbc, collectedSignaturesEvent *LogCollectedSignatures) error{
	log.Println("receive an collectedSignatures event from private chain")

	if collectedSignaturesEvent.AuthorityMachineResponsibleForRelay != p.txConfig.key.Address {
		log.Println("Collected signature should be submited by", collectedSignaturesEvent.AuthorityMachineResponsibleForRelay)
		return nil
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

	err := pbc.Pay(vs, rs, ss, message)
	if err!=nil {
		log.Println(err.Error())
	}
	return err
}

func (p *Pvc) EventReceiver(pbc *Pbc){
	log.Println("start private watcher")
	logs,eventError:=p.EventWatcher()

	contractAbi, err:= abi.JSON(strings.NewReader(string(privateSlot.PrivateSlotABI)))
	if err !=nil{
		log.Fatal(err.Error())
	}

	log.Println("Waiting for an event")
	for {
		select {
		case err:=<-eventError:
			log.Println(err.Error())
		case vLog := <-logs:
			switch vLog.Topics[0].Hex() {
			case logExchangeSigHash.Hex():
				var exchangeEvent LogExchange
				log.Println(len(vLog.Data))
				log.Println(vLog.Data)
				contractAbi.Unpack(&exchangeEvent, "Exchange", vLog.Data)
				exchangeEvent.TxHash = vLog.TxHash
				go p.ExchangeHandler(&exchangeEvent)
			case logCollectedSignaturesHash.Hex():
				var collectedSignaturesEvent LogCollectedSignatures
				contractAbi.Unpack(&collectedSignaturesEvent, "CollectedSignatures", vLog.Data)
				go p.CollectedSignaturesHandler(pbc, &collectedSignaturesEvent)
			}
		}
	}
}

func fillZero(src []byte, length int) []byte{
	prefix:= make([]byte,length,length+len(src))
	var buffer bytes.Buffer
	buffer.Write(prefix)
	buffer.Write(src)
	dst:=buffer.Bytes()
	return dst[len(dst)-length:]
}

func (p *Pvc) GetReceiptStatus (txHash common.Hash) (uint64,error) {
	count := time.Second
	for {
		time.Sleep(privateChainTime)
		receipt, err :=p.Client.TransactionReceipt(context.Background(),txHash)
		if err==nil {
			if receipt.Status==0{
				return receipt.Status, errors.New("transaction time out")
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

//func (p *Pvc) Deploy(initialSupply *big.Int, requiredSignatures *big.Int, authorities []common.Address) (common.Address, error) {
//	var err error
//
//	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
//	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce-1, big.NewInt(0))
//
//	address, _, _, err := privateSlot.DeployPrivateSlot(auth, p.Client, initialSupply, requiredSignatures, authorities)
//	return address, err
//}




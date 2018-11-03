package contract

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis"
	"github.com/goware/disque"
	"io/ioutil"
	"log"
	"math/big"
	"strings"
	"sync/atomic"
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

const (
	HashStringLen = 66
	AddressStringLen = 42
)

//Message for message queue
const (
	//message for queue of public chain
	NFTArriveOnPbc = "NFTonPbc"
	NFTArriveOnPvc = "NFTonPvc"
	SigsCollected  = "pvcSigsCollected"
	//message for queue of private chain
	ArriveOnPbc    = "onPbc"
	ArriveOnPvc    = "onPvc"
)

// Message Queue
const (
	//queue for public chain
	PbcPayNFTQueue = "pbcNFTPay"
	PbcPayQueue    = "pbcPay"
	// queue for private chain
	PvcPayNFTQueue       = "pvcNFTPay"
	PvcPayQueue          = "pvcPay"
	SubmitSignatureQueue = "submitSignature"
)

// transaction status
const (
	//nft tokens
	NFTPvcWFSig = "NFTPvcWFSig"
	NFTPvcWFPay = "NFTPvcWFPay"
	NFTPbcWFPay = "NFTPbcWFPay"
	NFTPbcPayed = "NFTPbcPayed"
	NFTPvcPayed = "NFTPvcPayed"
	// ERC20 tokens
	PvcWFSig = "PvcWFSig"
	PvcWFPay = "PvcWFPay"
	PbcWFPay = "PbcWFPay"
	PbcPayed = "PbcPayed"
	PvcPayed = "PvcPayed"
)

type txdata struct {
	AccountNonce hexutil.Uint64  `json:"nonce"    gencodec:"required"`
	Price        *hexutil.Big    `json:"gasPrice" gencodec:"required"`
	GasLimit     hexutil.Uint64  `json:"gas"      gencodec:"required"`
	Recipient    *common.Address `json:"to"       rlp:"nil"`
	Amount       *hexutil.Big    `json:"value"    gencodec:"required"`
	Payload      hexutil.Bytes   `json:"input"    gencodec:"required"`
	V            *hexutil.Big    `json:"v" gencodec:"required"`
	R            *hexutil.Big    `json:"r" gencodec:"required"`
	S            *hexutil.Big    `json:"s" gencodec:"required"`
	Hash         *common.Hash    `json:"hash" rlp:"-"`
}

type Avatar struct {
	TokenId     *big.Int `json:"tokenId"`
	Gene        *big.Int `json:"gene"`
	AvatarLevel *big.Int `json:"level"`
	Weaponed    bool     `json:"weaponed"`
	Armored     bool     `json:"armored"`
}

type LogExchange struct {
	User   common.Address
	Amount *big.Int

	TxHash common.Hash
}

type LogExchangeNFT struct {
	TokenID     *big.Int
	Owner       common.Address
	Gene        *big.Int
	AvatarLevel *big.Int
	Weaponed    bool
	Armored     bool

	TxHash common.Hash
}

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
	query := ethereum.FilterQuery{
		Addresses: []common.Address{c.Address},
	}
	logs := make(chan types.Log)
	sub, err := c.Client.SubscribeFilterLogs(context.Background(), query, logs)

	if err != nil {
		log.Fatal(err)
	}

	return logs, sub.Err()
}

func (c *Contract) ProcessJob(job *disque.Job) (*types.Transaction, error) {

	log.Println("process ", job.Queue)

	txWrapper := []byte(job.Data)
	var data txdata
	json.Unmarshal(txWrapper, &data)

	var nonce uint64
	nonce = atomic.AddUint64(&c.txConfig.nonce, 1)
	data.AccountNonce = hexutil.Uint64(nonce - 1)
	dataByte, err := json.Marshal(&data)

	tx := new(types.Transaction)
	tx.UnmarshalJSON(dataByte)

	chainID, err := c.Client.NetworkID(context.Background())
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), c.txConfig.key.PrivateKey)
	err = c.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Println(err.Error(), "send transaction fail")
	}
	return signedTx, err
}

func Consumer(jobs *disque.Pool, pvc *Pvc, pbc *Pbc, db *redis.Client) {
	for {
		job, _ := jobs.Get(PbcPayQueue, PvcPayQueue, PbcPayNFTQueue, PvcPayNFTQueue, SubmitSignatureQueue)
		switch job.Queue {
		case PbcPayQueue:
			user,txHash := parseJobHeader(job)
			_ , err := pbc.ProcessJob(job)
			if err != nil {
				jobs.Nack(job)
			} else {
				jobs.Ack(job)
				db.HSet(txHash, "status", PbcPayed)
				jobs.Add(ArriveOnPbc, PbcPayQueue+user)
			}
		case PbcPayNFTQueue:
			user,txHash := parseJobHeader(job)
			_ , err := pbc.ProcessJob(job)
			if err != nil {
				log.Println(err.Error())
				jobs.Nack(job)
			} else {
				jobs.Ack(job)
				db.HSet(txHash, "status", NFTPbcPayed)
				jobs.Add(NFTArriveOnPbc, PbcPayNFTQueue+user)
			}
		case SubmitSignatureQueue:
			user,txHash := parseJobHeader(job)
			_ , err := pvc.ProcessJob(job)
			if err != nil {
				log.Println(err.Error())
				jobs.Nack(job)
			} else {
				jobs.Ack(job)
				log.Println("ack submit Signature")
				db.HSet(txHash, "status", PvcWFPay)
				jobs.Add(SigsCollected, SubmitSignatureQueue+user)
			}
		case PvcPayQueue:
			user,txHash := parseJobHeader(job)
			_ , err := pvc.ProcessJob(job)
			if err != nil {
				jobs.Nack(job)
			} else {
				jobs.Ack(job)
				db.HSet(txHash, "status", PvcPayed)
				jobs.Add(ArriveOnPvc, PvcPayQueue+user)
			}
		case PvcPayNFTQueue:
			user,txHash := parseJobHeader(job)
			_ , err := pvc.ProcessJob(job)
			if err != nil {
				log.Println(err.Error())
				jobs.Nack(job)
			} else {
				jobs.Ack(job)
				db.HSet(txHash, "status", NFTPvcPayed)
				jobs.Add(NFTArriveOnPvc, PbcPayNFTQueue+user)
			}
		}
	}
}

// Truncate user+txHash from Front of job.Data
//
func parseJobHeader(job *disque.Job) (string,string) {
	user := strings.ToLower(job.Data[:AddressStringLen])
	txHash:= job.Data[AddressStringLen:AddressStringLen+HashStringLen]
	log.Println(txHash)
	job.Data = job.Data[AddressStringLen+HashStringLen:]
	return user,txHash
}
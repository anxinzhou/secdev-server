// Package main provides ...
package main

import (
	app "appClient"
	"contract"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-redis/redis"
	h "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"
	"utils"
)

const (
	chainConfigJson = "./etc/chainConfig.json"
	redisConfigJson = "./etc/redisConfig.json"
	disqueConfigJson = "./etc/disqueConfig.json"
)

var (
	pvc *contract.Pvc
	pbc *contract.Pbc
	machine *app.MachineState
	db *redis.Client

	mutex  *sync.Mutex
)

type BasicChainConfig struct {
	Port     string `json:"port"`
	Keystore string `json:"keystore"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

type RedisConfig struct {
	Port string `json:"port"`
	Password string `json:"password"`
	DB int `json:"DB"`
}

type PrivateConfig struct {
	BasicChainConfig
}

type PublicConfig struct {
	BasicChainConfig
}

type ChainConfig struct {
	Pri *PrivateConfig `json:"private"`
	Pub *PublicConfig `json:"public"`
}

func wrapperAndSend(gcuid int64,res interface{},c *websocket.Conn) {
	resWrapper, err:=json.Marshal(res)
	if err!=nil {
		log.Println(err.Error())
		sendError(gcuid, app.ServerErrorCode, app.ServerJsonError, c)
		return
	}
	mutex.Lock()
	c.WriteMessage(websocket.TextMessage, resWrapper)
	mutex.Unlock()
}

func postRes (gcuid int64, c *websocket.Conn) {
	res:= &app.PostTokenUserOrRewardRes {
		PostRes:&app.PostRes{
			Status: app.OkStatus,
			Gcuid: gcuid,
		},
	}
	wrapperAndSend(gcuid,res,c)
}

func sendError(gcuid int64,errorCode int64, reason string, c *websocket.Conn) {
	reqError:= &app.Error{
		Status: app.FailedStatus,
		Code: errorCode,
		Gcuid: gcuid,
		Reason: reason,
	}
	reqErrorWrapper, err:=json.Marshal(reqError)
	if err!=nil {
		log.Println(err.Error())
		sendError(gcuid, app.ServerErrorCode, app.ServerJsonError, c)
		return
	}
	c.WriteMessage(websocket.TextMessage, reqErrorWrapper)
}

func SigninHandler (data []byte,c *websocket.Conn) {
	log.Println("sign in handler")
	var signin app.SigninReq
	err:=json.Unmarshal(data, &signin)
	if err!=nil{
		log.Println(err.Error())
		sendError(app.Signin,app.ClientErrorCode, app.ClientFormatError,c)
		return
	}
	if signin.LoginCode!=app.LoginCode {
		sendError(app.Signin,app.ClientErrorCode,app.WrongUser,c)
		return
	}

	res:= &app.SigninRes{
		Status:app.OkStatus,
		Gcuid:app.Signin,
		Guuid:app.Guuid,
	}

	wrapperAndSend(app.Signin,res,c)
}

func SignoutHandler (data []byte,c *websocket.Conn) {
	log.Println("signout handler")
	postRes(app.Signout,c)
}

func getWallet() ([]*app.Wallet, *app.TransferInfo,error) {
	ether,err:= pbc.GetEther(app.UserAddr)
	if err!=nil {
		return nil,nil,err
	}
	token,err:=pbc.GetToken(app.UserAddr)
	if err!=nil {
		log.Println(err.Error())
		return nil,nil,err
	}

	tokenWallet :=&app.Wallet{
		Name: app.TokenWalletName,
		Amount: new(big.Float).Quo(new(big.Float).SetInt(token),app.TokenBase).String(),
		Id: app.User,
	}
	etherWallet:=&app.Wallet{
		Name: app.EthWalletName,
		Amount: new(big.Float).Quo(new(big.Float).SetInt(ether),app.EtherBase).String(),
		Id: app.User,
	}

	wallets:=make([]*app.Wallet,0,app.WalletCount)
	wallets = append(wallets, etherWallet, tokenWallet)

	rawTransferAmount,err :=pvc.GetToken(app.UserAddr)
	if err!=nil {
		log.Println(err.Error())
		return nil,nil,err
	}
	transferAmount:=new(big.Float).Quo(new(big.Float).SetInt(rawTransferAmount),app.TokenBase).String()
	toTransfer:=&app.TransferInfo{
		Amount:transferAmount,
		Name:app.TokenWalletName,
	}
	return wallets,toTransfer,nil
}

func GetWalletsAndMachineHandler(data []byte, c *websocket.Conn) {

	wallets,toTransfer,err:=getWallet()
	if err!=nil {
		log.Println(err.Error())
		sendError(app.GetWalletsAndMachine,app.ServerErrorCode,app.ServerChainError,c)
		return
	}
	res:= &app.GetWalletsAndMachineRes{
		Gcuid: app.GetWalletsAndMachine,
		Machine: &app.MachineState {
			State: machine.State,
			MachineId: machine.MachineId,
		},
		WalletsCount: app.WalletCount,
		Wallets: wallets,
		Transfer: toTransfer,
	}


	wrapperAndSend(app.GetWalletsAndMachine,res,c)
}

func GetWalletsHandler(data []byte, c *websocket.Conn) {
	wallets,toTransfer,err:=getWallet()
	if err!=nil {
		log.Println(err.Error())
		sendError(app.GetWallets,app.ServerErrorCode,app.ServerChainError,c)
		return
	}
	res:= &app.GetWalletRes{
		Gcuid: app.GetWallets,
		Count: app.WalletCount,
		Wallets: wallets,
		Transfer: toTransfer,
	}

	wrapperAndSend(app.GetWallets,res,c)
}


func GetTransactionsHandler(data []byte, c *websocket.Conn) {
	var req app.GetTransactionsReq
	err:=json.Unmarshal(data,&req)
	if err!=nil {
		sendError(app.GetTransactions,app.ClientErrorCode,app.ClientFormatError,c)
		return
	}

	start:= req.After
	stop:=req.After + req.Amount
	var vals []string
	var collection string
	switch req.From {
	case app.ETH:
		collection =strings.ToLower(app.User)+":"+string(app.ETH)
	case app.Slot:
		collection =strings.ToLower(app.User)+":"+string(app.Slot)
	}
	totalLen,err:= db.LLen(collection).Result()
	log.Println("transaction len:", totalLen)
	if err!=nil{
		log.Println(err.Error())
		sendError(app.GetTransactions,app.ClientErrorCode,err.Error(),c)
		return
	}

	count:=totalLen-start
	if count>req.Amount {
		count = req.Amount
	}

	if count<0 {
		count = 0
	}

	vals,err = db.LRange(collection,start,stop).Result()
	log.Println(vals)
	if err!=nil{
		log.Println(err.Error())
		sendError(app.GetTransactions,app.ClientErrorCode,err.Error(),c)
		return
	}
	txs:= make([]*app.Transaction,len(vals))
	for i,val:=range vals {
		var tx app.Transaction
		err =json.Unmarshal([]byte(val), &tx)
		if err!=nil{
			sendError(app.GetTransactions,app.ServerErrorCode,app.ServerJsonError,c)
			return
		}
		txs[i]=&tx
	}

	res:= &app.GetTransactionRes{
		Gcuid: app.GetTransactions,
		Count: count,
		After: req.After,
		From: req.From,
		Transactions:txs,
	}

	wrapperAndSend(app.GetTransactions,res,c)
}

func GetExchangeRateHandler(data []byte, c *websocket.Conn) {
	var req app.GetExchangeRateReq
	err := json.Unmarshal(data, &req)
	if err!=nil {
		sendError(app.GetExchangeRate,app.ClientErrorCode,app.ClientFormatError,c)
		return
	}

	rate,err:=pbc.GetExchangeRate()
	if err!=nil {
		log.Println(err.Error())
		sendError(app.GetExchangeRate,app.ServerErrorCode,app.ServerChainError,c)
		return
	}
	switch req.ExchangeType {
	case app.SlotToEth:
		rateWrapper,_,_:=new(big.Float).Parse(rate,10)
		rate = new(big.Float).Quo(big.NewFloat(1),rateWrapper).String()
	}

	res:= &app.GetExchangeRateRes{
		Gcuid: app.GetExchangeRate,
		ExchangeType: req.ExchangeType,
		Rate: rate,
	}

	wrapperAndSend(app.GetExchangeRate,res,c)
}

func exchangeResultHandler(req *app.PostExchangeReq,transaction *types.Transaction, exchangeError error, c *websocket.Conn) {
	if exchangeError!=nil {

		log.Println(exchangeError.Error())
		sendError(app.PostExchange,app.ClientErrorCode,app.ClientFormatError,c)
		return
	}
	txHash:=transaction.Hash()
	postRes(app.PostExchange,c)

	_,err :=pbc.GetReceiptStatus(txHash)
	var status string
	if err!=nil {
		log.Println(err.Error())
		status = app.FailedStatus
		return
	} else {
		status = app.SuccessStatus
		amountWrapper,_,_ := new(big.Float).Parse(req.Amount,10)
		rawExchangeRate,_:=pbc.GetExchangeRate()
		exchangeRate,_,_ := new(big.Float).Parse(rawExchangeRate,10)
		// "2006-01-02 15:04:05 is the birth day of golang, fixed format"
	    date:= utils.GetCurrentTime()
		switch req.ExchangeType {
		case app.EthToSLot:
			withdrawTokenAmount := new(big.Float).Mul(amountWrapper,exchangeRate)
			ethTx:= app.Transaction{
				Type:app.Withdraw,
				Amount:req.Amount,
				CreatedDate: date,
			}
			slotTx:= app.Transaction{
				Type: app.Deposit,
				Amount: withdrawTokenAmount.String(),
				CreatedDate: date,
			}
			slotTxWrapper,err:=json.Marshal(&slotTx)
			ethTxWrapper,err:=json.Marshal(&ethTx)
			_,err =db.LPush(strings.ToLower(app.User)+":"+string(app.ETH), ethTxWrapper).Result()
			if err!=nil {
				log.Println(err.Error())
				return
			}
			_,err = db.LPush(strings.ToLower(app.User)+":"+string(app.Slot),slotTxWrapper).Result()
			if err!=nil {
				log.Println(err.Error())
				return
			}
		case app.SlotToEth:
			withdrawEthAmount:= new(big.Float).Quo(amountWrapper,exchangeRate)
			ethTx:=app.Transaction{
				Type:app.Deposit,
				Amount:withdrawEthAmount.String(),
				CreatedDate: date,
			}
			slotTx:= app.Transaction{
				Type:app.Withdraw,
				Amount: req.Amount,
				CreatedDate: date,
		    }
			slotTxWrapper,_:=json.Marshal(&slotTx)
			ethTxWrapper,_:=json.Marshal(&ethTx)
			_,err =db.LPush(strings.ToLower(app.User)+":"+string(app.ETH), ethTxWrapper).Result()
			if err!=nil {
				log.Println(err.Error())
				return
			}
			_,err =db.LPush(strings.ToLower(app.User)+":"+string(app.Slot),slotTxWrapper).Result()
			if err!=nil {
				log.Println(err.Error())
				return
			}
		}
	}

	resultRes := & app.ExchangeResultRes {
		Gcuid: app.NotifyExchangeResult,
		Status:status,
		ExchangeType: req.ExchangeType,
		Amount: req.Amount,
		CreatedDate: utils.GetCurrentTime(),
	}

	wrapperAndSend(app.NotifyTokenChange,resultRes,c)
}

func PostExchangeHandler(data []byte, c *websocket.Conn){
	var req app.PostExchangeReq
	err:=json.Unmarshal(data,&req)
	if err!=nil {
		log.Println(err.Error())
		sendError(app.PostExchange,app.ClientErrorCode,app.ClientFormatError,c)
		return
	}

	amount,_,err:=  new(big.Float).Parse(req.Amount,10)
	if err!=nil {
		log.Println(err.Error())
		sendError(app.PostExchange,app.ClientErrorCode,app.ClientFormatError,c)
		return
	}

	switch req.ExchangeType {
	case app.EthToSLot:
		amountWrapper,_:=new(big.Float).Mul(amount,app.EtherBase).Int(nil)
		tx,err:=pbc.ExchangeForToken(app.UserAddr,amountWrapper)
		exchangeResultHandler(&req,tx,err,c)

	case app.SlotToEth:
		amountWrapper,_:=new(big.Float).Mul(amount,app.TokenBase).Int(nil)
		tx,err:=pbc.ExchangeForEther(app.UserAddr,amountWrapper)
		exchangeResultHandler(&req,tx,err,c)
	}
}

func machineStatusChange(state int64, c *websocket.Conn){
	res:= &app.MachineStatusChangeRes{
		Gcuid: app.NotifyMachineStatusChange,
		State:state,
		MachineId:app.MachineId,
		Act:app.MachineStatusChange,
		CreatedDate: utils.GetCurrentTime(),
	}

	wrapperAndSend(app.NotifyMachineStatusChange,res,c)
}

func PostQRCodeHandler(data []byte, c *websocket.Conn) {
	postRes(app.PostQRCode,c)
	time.Sleep(app.SleepTime)
	machineStatusChange(app.Login,c)
}


func MachineLogoutHandler(data []byte, c *websocket.Conn) {
	postRes(app.MachineLogout,c)
	machineStatusChange(app.Logout,c)
}

func PostTokenUserOrRewardHandler(data []byte, c *websocket.Conn) {
	var req app.PostTokenUseOrRewardReq
	err :=json.Unmarshal(data,&req)
	if err!=nil {
		log.Println(err.Error())
		sendError(app.PostTokenUseOrReward,app.ClientErrorCode,app.ClientFormatError,c)
	}
	amount,_,err:=  new(big.Float).Parse(req.Amount,10)
	if err!=nil {
		log.Println(err.Error())
		sendError(app.PostTokenUseOrReward,app.ClientErrorCode,app.ClientFormatError,c)
		return
	}

	amountWrapper,_:=new(big.Float).Mul(amount,app.TokenBase).Int(nil)
	switch req.Type {
	case app.Used:
		err =pvc.Consume(app.UserAddr,amountWrapper)
	case app.Reward:
		err =pvc.Reward(app.UserAddr,amountWrapper)
	}
	if err!=nil {
		log.Println(err.Error())
		sendError(app.PostTokenUseOrReward, app.ServerErrorCode, err.Error(),c)
		return
	}

	postRes(app.PostTokenUseOrReward,c)
}

func PostGameStartOrEndHandler(data []byte, c *websocket.Conn) {
	var req app.PostGameStartOrEndReq
	err:= json.Unmarshal(data,&req)
	if err!=nil {
		log.Println(err.Error())
		sendError(app.PostGameStartOrEnd,app.ClientErrorCode,app.ClientFormatError,c)
	}
	amount,_,err:=  new(big.Float).Parse(req.Amount,10)
	if err!=nil {
		log.Println(err.Error())
		sendError(app.PostTokenUseOrReward,app.ClientErrorCode,app.ClientFormatError,c)
		return
	}

	amountWrapper,_:=new(big.Float).Mul(amount,app.TokenBase).Int(nil)
	switch req.Type {
	case app.GameStart:
		err =pvc.Mint(app.UserAddr,amountWrapper)
		if err!=nil{
			log.Println(err.Error())
			sendError(app.PostGameStartOrEnd,app.ServerErrorCode,err.Error(),c)
			return
		}
		postRes(app.PostGameStartOrEnd,c)

		err:=pbc.Burn(app.UserAddr,amountWrapper)
		if err!=nil {
			log.Println(err.Error())
			pvc.Burn(app.UserAddr,amountWrapper)
		}

		tokenUpdate := amount.String()
		tx:=app.Transaction{
			Type: app.Spend,
			Amount: tokenUpdate,
			CreatedDate: utils.GetCurrentTime(),
		}
		txWrapper,_:=json.Marshal(tx)
		_,err =db.LPush(strings.ToLower(app.User)+":"+string(app.Slot),txWrapper).Result()
		if err!=nil {
			log.Println(err.Error())
			return
		}

		rawTokenTotal,err :=pbc.GetToken(app.UserAddr)
		if err!=nil {
			log.Println(err.Error())
			return
		}
		tokenTotal :=new(big.Float).Quo(new(big.Float).SetInt(rawTokenTotal),app.TokenBase).String()


		resultRes := & app.TokenCountChangeRes {
			Gcuid: app.NotifyTokenChange,
			State: app.Decrease,
			WalletName: app.TokenWalletName,
			WalletId: app.User,
			TokenUpdate: tokenUpdate,
			TokenTotal: tokenTotal,
			Act: app.TokenChange,
			CreatedDate: utils.GetCurrentTime(),
		}

		wrapperAndSend(app.NotifyTokenChange,resultRes,c)
	case app.GameEnd:

		postRes(app.PostGameStartOrEnd,c)
	}
}

func PostTransferHandler(data []byte, c *websocket.Conn){

	rawToken, err:= pvc.GetToken(app.UserAddr)
	if err!=nil {
		log.Println(err.Error())
		sendError(app.PostTransfer, app.ServerErrorCode, err.Error(),c)
		return
	}

	err = pvc.Burn(app.UserAddr,rawToken)
	if err!=nil {
		log.Println(err.Error())
		sendError(app.PostTransfer, app.ServerErrorCode, err.Error(),c)
		return
	}
	postRes(app.PostTransfer,c)

	err = pbc.Mint(app.UserAddr,rawToken)
	if err!=nil {
		log.Println(err.Error())
		pvc.Mint(app.UserAddr,rawToken)
		return
	}

	outGameSlot:=new(big.Float).Quo(new(big.Float).SetInt(rawToken),app.TokenBase)

	tokenUpdate := outGameSlot.String()
	tx:=app.Transaction{
		Type: app.Gain,
		Amount: tokenUpdate,
		CreatedDate: utils.GetCurrentTime(),
	}
	txWrapper,_:=json.Marshal(tx)
	_,err =db.LPush(strings.ToLower(app.User)+":"+string(app.Slot),txWrapper).Result()
	if err!=nil {
		log.Println(err.Error())
		return
	}

	rawTokenTotal,err :=pbc.GetToken(app.UserAddr)
	if err!=nil {
		log.Println(err.Error())
		return
	}
	tokenTotal :=new(big.Float).Quo(new(big.Float).SetInt(rawTokenTotal),app.TokenBase).String()


	resultRes := & app.TokenCountChangeRes {
		Gcuid: app.NotifyTokenChange,
		State: app.Increase,
		WalletName: app.TokenWalletName,
		WalletId: app.User,
		TokenUpdate: tokenUpdate,
		TokenTotal: tokenTotal,
		Act: app.TokenChange,
		CreatedDate: utils.GetCurrentTime(),
	}

	wrapperAndSend(app.NotifyTokenChange,resultRes,c)
}

func disConnect(c *websocket.Conn) {
	res:=&app.Disconnect{
		Status:app.FailedStatus,
		Code: app.ServerErrorCode,
		Reason: app.ServerDisconnect,
	}

	wrapperAndSend(app.DisConnect,res,c)
}

func requestHandler(c *websocket.Conn) {
	for {
		_,data,err := c.ReadMessage()
		var kvs map[string]interface{}
		if err!=nil {
			log.Println(err.Error())
			return
		}
		json.Unmarshal(data, &kvs)
		gid,ok := kvs["gcuid"]
		if !ok {
			log.Println(errors.New("gcuid not exist"))
			return
		}

		gcuid:= int64(gid.(float64))
		log.Println("gcuid:",gcuid)
		switch gcuid {
		case app.Signin:
			go SigninHandler(data,c)
		case app.Signout:
			go SignoutHandler(data,c)
		case app.GetWalletsAndMachine:
			go GetWalletsAndMachineHandler(data,c)
		case app.GetWallets:
			go GetWalletsHandler(data,c)
		case app.GetTransactions:
			go GetTransactionsHandler(data,c)
		case app.GetExchangeRate:
			go GetExchangeRateHandler(data,c)
		case app.PostExchange:
			go PostExchangeHandler(data,c)
		case app.PostQRCode:
			go PostQRCodeHandler(data,c)
		case app.NotifyMachineStatusChange:
			go PostQRCodeHandler(data,c)
		case app.MachineLogout:
			go MachineLogoutHandler(data,c)
		case app.PostTokenUseOrReward:
			go PostTokenUserOrRewardHandler(data,c)
		case app.PostGameStartOrEnd:
			go PostGameStartOrEndHandler(data,c)
		case app.PostTransfer:
			go PostTransferHandler(data,c)
		}
	}
}

func requestParser(w http.ResponseWriter, r *http.Request) {
	log.Println("receive a reqeust")
	var upgrader = websocket.Upgrader{}
	upgrader.CheckOrigin = func(rq *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	requestHandler(c)
	defer func(){
		disConnect(c)
		c.Close()
	}()
}

func init() {
	// initial contract
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var cc ChainConfig
	data, err := ioutil.ReadFile(chainConfigJson)
	if err != nil {
		panic("can not read chain config file")
	}
	json.Unmarshal(data, &cc)
	pbc = contract.NewPbc(cc.Pub.Port, cc.Pub.Keystore, cc.Pub.Password, cc.Pub.Address)
	pvc = contract.NewPvc(cc.Pri.Port, cc.Pri.Keystore, cc.Pri.Password, cc.Pri.Address)
	log.Println("hello server")

	//inital game machine
	machine = &app.MachineState{
		State: app.Logout,
		MachineId: app.MachineId,
	}

	//inital redis
	var rc RedisConfig
	data, err = ioutil.ReadFile(redisConfigJson)
	if err!=nil {
		panic("can not read chain config file")
	}
	json.Unmarshal(data,&rc)

	db = redis.NewClient(&redis.Options{
		Addr: rc.Port,
		Password: rc.Password,
		DB: rc.DB,
	})

	//lock
	mutex = &sync.Mutex{}
}

//test
func init(){
	amount,_:=new(big.Float).Mul(big.NewFloat(1000),app.TokenBase).Int(nil)
	go pbc.Reward(app.UserAddr,amount)
}


func main() {
	//runtime.GOMAXPROCS(8)
	r := mux.NewRouter()

	r.HandleFunc("/", requestParser).Methods("GET")

	fmt.Println("Running http server")
	http.ListenAndServe(
		":4000",
		h.CORS(
			h.AllowedMethods([]string{"get", "options", "post", "put", "head"}),
			h.AllowedOrigins([]string{"*"}),
			h.AllowedHeaders([]string{"Content-Type"}),
			)(r),
			)
}

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
	"time"
	"utils"
)

const (
	chainConfigJson = "./etc/chainConfig.json"
	redisConfigJson = "./etc/redisConfig.json"
)

var (
	pvc *contract.Pvc
	machine *app.MachineState
	db *redis.Client
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

type ChainConfig struct {
	Pri *PrivateConfig `json:"private"`
}

func sendError(gcuid int64,errorCode int64, reason string, c *websocket.Conn) {
	reqError:= &app.Error{
		Status: app.FailedStatus,
		Code: errorCode,
		Gcuid: gcuid,
		Reason: reason,
	}
	reqErrorWrapper,err:=json.Marshal(reqError)
	if err!=nil {
		log.Println(err.Error())
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
	resWrapper, err:=json.Marshal(res)
	if err!=nil {
		log.Println(err.Error())
		sendError(app.Signin,app.ServerErrorCode,app.ServerJsonError,c)
		return
	}
	c.WriteMessage(websocket.TextMessage, resWrapper)
}

func SignoutHandler (data []byte,c *websocket.Conn) {
	log.Println("signout handler")
	res:= &app.SignoutRes{
		PostRes:&app.PostRes{
			Status: app.OkStatus,
			Gcuid: app.Signout,
		},
	}
	resWrapper, err:=json.Marshal(res)
	if err!=nil {
		log.Println(err.Error())
		sendError(app.Signout,app.ServerErrorCode,app.ServerJsonError,c)
		return
	}
	c.WriteMessage(websocket.TextMessage, resWrapper)
}

func getWallet() ([]*app.Wallet, error) {
	ether,err:= pvc.GetEther(app.UserAddr)
	if err!=nil {
		return nil,err
	}
	token,err:=pvc.GetToken(app.UserAddr)
	if err!=nil {
		log.Println(err.Error())
		return nil,err
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
	return wallets,nil
}

func GetWalletsAndMachineHandler(data []byte, c *websocket.Conn) {

	wallets,err:=getWallet()
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
	}
	resWrapper, err:=json.Marshal(res)
	if err!=nil {
		log.Println(err.Error())
		sendError(app.GetWalletsAndMachine,app.ServerErrorCode,app.ServerJsonError,c)
		return
	}
	c.WriteMessage(websocket.TextMessage, resWrapper)
}

func GetWalletsHandler(data []byte, c *websocket.Conn) {
	wallets,err:=getWallet()
	if err!=nil {
		log.Println(err.Error())
		sendError(app.GetWallets,app.ServerErrorCode,app.ServerChainError,c)
		return
	}
	res:= &app.GetWalletRes{
		Gcuid: app.GetWallets,
		Count: app.WalletCount,
		Wallets: wallets,
	}
	resWrapper, err:=json.Marshal(res)
	if err!=nil {
		log.Println(err.Error())
		sendError(app.GetWallets,app.ServerErrorCode,app.ServerJsonError,c)
		return
	}
	c.WriteMessage(websocket.TextMessage, resWrapper)
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
		collection =app.User+":"+string(app.ETH)
	case app.Slot:
		collection =app.User+":"+string(app.Slot)
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
	resWrapper, err:=json.Marshal(res)
	if err!=nil {
		log.Println(err.Error())
		sendError(app.GetTransactions,app.ServerErrorCode,app.ServerJsonError,c)
		return
	}
	c.WriteMessage(websocket.TextMessage, resWrapper)
}

func GetExchangeRateHandler(data []byte, c *websocket.Conn) {
	var req app.GetExchangeRateReq
	err := json.Unmarshal(data, &req)
	if err!=nil {
		sendError(app.GetExchangeRate,app.ClientErrorCode,app.ClientFormatError,c)
		return
	}

	rate,err:=pvc.GetExchangeRate()
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
	resWrapper, err:=json.Marshal(res)
	if err!=nil {
		log.Println(err.Error())
		sendError(app.GetExchangeRate,app.ServerErrorCode,app.ServerJsonError,c)
	} else {
		c.WriteMessage(websocket.TextMessage, resWrapper)
	}

}

func exchangeResultHandler(req *app.PostExchangeReq,transaction *types.Transaction, exchangeError error, c *websocket.Conn) {
	if exchangeError!=nil {

		log.Println(exchangeError.Error())
		sendError(app.PostExchange,app.ClientErrorCode,app.ClientFormatError,c)
		return
	}
	txHash:=transaction.Hash()
	res:= &app.PostExchangeRes {
		&app.PostRes{
			Status: app.OkStatus,
			Gcuid: app.PostExchange,
		},
	}
	resWrapper, err:=json.Marshal(res)
	if err!=nil {
		log.Println(err.Error())
		sendError(app.PostExchange, app.ServerErrorCode, app.ServerJsonError, c)
		return
	}
	c.WriteMessage(websocket.TextMessage, resWrapper)

	_,err=pvc.GetReceiptStatus(txHash)
	var status string
	if err!=nil {
		log.Println(err.Error())
		status = app.FailedStatus
		return
	} else {
		status = app.SuccessStatus
		amountWrapper,_,_ := new(big.Float).Parse(req.Amount,10)
		rawExchangeRate,_:=pvc.GetExchangeRate()
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
			_,err =db.LPush(app.User+":"+string(app.ETH), ethTxWrapper).Result()
			if err!=nil {
				log.Println(err.Error())
				return
			}
			_,err = db.LPush(app.User+":"+string(app.Slot),slotTxWrapper).Result()
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
			_,err =db.LPush(app.User+":"+string(app.ETH), ethTxWrapper).Result()
			if err!=nil {
				log.Println(err.Error())
				return
			}
			_,err =db.LPush(app.User+":"+string(app.Slot),slotTxWrapper).Result()
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
	resultResWrapper, err:=json.Marshal(resultRes)
	if err!=nil {
		log.Fatalln(err.Error())
		return
	}
	c.WriteMessage(websocket.TextMessage, resultResWrapper)
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
		tx,err:=pvc.ExchangeForToken(app.UserAddr,amountWrapper)
		exchangeResultHandler(&req,tx,err,c)

	case app.SlotToEth:
		amountWrapper,_:=new(big.Float).Mul(amount,app.TokenBase).Int(nil)
		tx,err:=pvc.ExchangeForEther(app.UserAddr,amountWrapper)
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
	resWrapper, err:=json.Marshal(res)
	if err!=nil {
		log.Println(err.Error())
		sendError(app.NotifyMachineStatusChange, app.ServerErrorCode, app.ServerJsonError, c)
		return
	}
	c.WriteMessage(websocket.TextMessage, resWrapper)
}

func PostQRCodeHandler(data []byte, c *websocket.Conn) {
	res:= &app.PostQRCodeRes{
		PostRes:&app.PostRes{
			Status:app.OkStatus,
			Gcuid: app.PostQRCode,
		},
	}
	resWrapper, err:=json.Marshal(res)
	if err!=nil {
		log.Println(err.Error())
		sendError(app.PostQRCode, app.ServerErrorCode, app.ServerJsonError, c)
		return
	}
	c.WriteMessage(websocket.TextMessage, resWrapper)
	time.Sleep(app.SleepTime)
	machineStatusChange(app.Login,c)
}


func MachineLogoutHandler(data []byte, c *websocket.Conn) {
	res:= &app.PostMachineLogoutRes{
		PostRes:&app.PostRes{
			Status: app.OkStatus,
			Gcuid: app.MachineLogout,
		},
	}
	resWrapper, err:=json.Marshal(res)
	if err!=nil {
		log.Println(err.Error())
		sendError(app.MachineLogout, app.ServerErrorCode, app.ServerJsonError, c)
		return
	}
	c.WriteMessage(websocket.TextMessage, resWrapper)
	machineStatusChange(app.Logout,c)
}

func updateResultHandler(req *app.PostTokenUseOrRewardReq, transaction *types.Transaction, updateError error, c *websocket.Conn){
	if updateError!=nil {
		log.Println(updateError.Error())
		sendError(app.PostTokenUseOrReward,app.ClientErrorCode,app.ClientFormatError,c)
		return
	}
	txHash:=transaction.Hash()
	res:= &app.PostTokenUserOrRewardRes {
		PostRes:&app.PostRes{
			Status: app.OkStatus,
			Gcuid: app.PostTokenUseOrReward,
		},
	}
	resWrapper, err:=json.Marshal(res)
	if err!=nil {
		log.Println(err.Error())
		sendError(app.PostTokenUseOrReward, app.ServerErrorCode, app.ServerJsonError, c)
		return
	}
	c.WriteMessage(websocket.TextMessage, resWrapper)

	_,err=pvc.GetReceiptStatus(txHash)
	var tokenUpdate string
	var state int64
	var token string
	if err!=nil {
		log.Println(err.Error())
		tokenUpdate = "0"
		state = 0
		return
	} else {
		var txType int64
		tokenUpdate = req.Amount
		switch req.Type {
		case app.Used:
			txType = app.Spend
			state = app.Decrease
		case app.Reward:
			state = app.Increase
			txType = app.Reward
		}
		tx:=app.Transaction{
			Type: txType,
			Amount: req.Amount,
			CreatedDate: utils.GetCurrentTime(),
		}
		txWrapper,_:=json.Marshal(tx)
		_,err:=db.LPush(app.User+":"+string(app.Slot),txWrapper).Result()
		if err!=nil {
			log.Println(err.Error())
			return
		}
	}

	rawToken,err:=pvc.GetToken(app.UserAddr)
	if err!=nil {
		log.Println(err.Error())
		token = "0"
		return
	} else {
		token =new(big.Float).Quo(new(big.Float).SetInt(rawToken),app.TokenBase).String()
	}

	resultRes := & app.TokenCountChangeRes {
		Gcuid: app.PostTokenUseOrReward,
		State: state,
		WalletName: app.TokenWalletName,
		WalletId: app.User,
		TokenUpdate: tokenUpdate,
		TokenTotal: token,
		Act: app.TokenChange,
		CreatedDate: utils.GetCurrentTime(),
	}
	resultResWrapper, err:=json.Marshal(resultRes)
	if err!=nil {
		log.Fatalln(err.Error())
		return
	}
	c.WriteMessage(websocket.TextMessage, resultResWrapper)
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
		tx,err:=pvc.Consume(app.UserAddr,amountWrapper)
		updateResultHandler(&req,tx,err,c)
	case app.Reward:
		tx,err:=pvc.Reward(app.UserAddr,amountWrapper)
		updateResultHandler(&req,tx,err,c)
	}
}

func disConnect(c *websocket.Conn) {
	res:=&app.Disconnect{
		Status:app.FailedStatus,
		Code: app.ServerErrorCode,
		Reason: app.ServerDisconnect,
	}
	resWrapper, err:=json.Marshal(res)
	if err!=nil {
		log.Println(err.Error())
	}
	c.WriteMessage(websocket.TextMessage, resWrapper)
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
			SigninHandler(data,c)
		case app.Signout:
			SignoutHandler(data,c)
		case app.GetWalletsAndMachine:
			GetWalletsAndMachineHandler(data,c)
		case app.GetWallets:
			GetWalletsHandler(data,c)
		case app.GetTransactions:
			GetTransactionsHandler(data,c)
		case app.GetExchangeRate:
			GetExchangeRateHandler(data,c)
		case app.PostExchange:
			PostExchangeHandler(data,c)
		case app.PostQRCode:
			PostQRCodeHandler(data,c)
		case app.NotifyMachineStatusChange:
			PostQRCodeHandler(data,c)
		case app.MachineLogout:
			MachineLogoutHandler(data,c)
		case app.PostTokenUseOrReward:
			PostTokenUserOrRewardHandler(data,c)
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

	pvc = contract.NewPvc(cc.Pri.Port, cc.Pri.Keystore, cc.Pri.Password, cc.Pri.Address)
	log.Println("hello server")
}

func init(){
	//inital game machine
	machine = &app.MachineState{
		State: app.Logout,
		MachineId: app.MachineId,
	}
}

func init(){
	//inital redis
	var rc RedisConfig
	data, err:= ioutil.ReadFile(redisConfigJson)
	if err!=nil {
		panic("can not read chain config file")
	}
	json.Unmarshal(data,&rc)

	db = redis.NewClient(&redis.Options{
		Addr: rc.Port,
		Password: rc.Password,
		DB: rc.DB,
	})
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

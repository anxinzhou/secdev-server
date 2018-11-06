// Package main provides ...
package main

import (
	app "appClient"
	"contract"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis"
	h "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"time"
)

const (
	chainConfigJson = "src/etc/chainConfig.json"
	redisConfigJson = "src/etc/redisConfig.json"
)

var (
	pvc *contract.Pvc
	machine *app.MachineState
	QRCode string
	db *redis.Client
)

type DisqueConfig struct {
	Port string `json:"port"`
}

type ExPayLoad struct {
	Tx string `json:"tx"`
	User string `json:"user"`
	Kind string `json:"kind"`
}

type ExTokensPayLoad struct {
	*ExPayLoad
	Amount int `json:"amount"`
}

type ExNFTPayLoad struct {
	*ExPayLoad
	TokenId *big.Int `json:"tokenId"`
}

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

func sendError(gcuid int,errorCode int, reason string, c *websocket.Conn) {
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
	res:= &app.SignoutRes{
		&app.PostRes{
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
	res:= &app.GetWalletsAndMachineRes{
		Gcuid: app.GetWallets,
		WalletsCount: app.WalletCount,
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

	switch req.From {
	case app.ETH:
	case app.Slot:
	}
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

func exchangeResultHandler(req *app.PostExchangeReq,txHash common.Hash, exchangeError error, c *websocket.Conn) {
	if exchangeError!=nil {
		log.Println(exchangeError.Error())
		sendError(app.PostExchange,app.ClientErrorCode,app.ClientFormatError,c)
		return
	}
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
	} else {
		status = app.SuccessStatus
		amountWrapper,_,_ := new(big.Float).Parse(req.Amount,10)
		rawExchangeRate,_:=pvc.GetExchangeRate()
		exchangeRate,_,_ := new(big.Float).Parse(rawExchangeRate,10)
		// "2006-01-02 15:04:05 is the birth day of golang, fixed format"
	    date:= time.Now().Format("2006-01-02 15:04:05")
		switch req.ExchangeType {
		case app.EthToSLot:
			withdrawTokenAmount := new(big.Float).Mul(amountWrapper,exchangeRate)
			ethTx:= app.Transaction{
				Type:app.Deposit,
				Amount:req.Amount,
				CreatedDate: date,
			}
			slotTx:= app.Transaction{
				Type: app.Withdraw,
				Amount: new(big.Float).Quo(withdrawTokenAmount,app.TokenBase).String(),
				CreatedDate: date,
			}
			db.LPush(app.User+":"+string(app.ETH), ethTx)
			db.LPush(app.User+":"+string(app.Slot),slotTx)
		case app.SlotToEth:
			withdrawEthAmount:= new(big.Float).Quo(amountWrapper,exchangeRate)
			ethTx:=app.Transaction{
				Type:app.Withdraw,
				Amount:new(big.Float).Quo(withdrawEthAmount,app.EtherBase).String(),
				CreatedDate: date,
			}
			slotTx:= app.Transaction{
				Type:app.Deposit,
				Amount: req.Amount,
				CreatedDate: date,
		    }
			db.LPush(app.User+":"+string(app.ETH), ethTx)
			db.LPush(app.User+":"+string(app.Slot),slotTx)
		}
	}

	resultRes := & app.ExchangeResultRes {
		Gcuid: app.NotifyExchangeResult,
		Status:status,
		ExchangeType: req.ExchangeType,
		Amount: req.Amount,
		// "2006-01-02 15:04:05 is the birth day of golang, fixed format"
		CreatedDate: time.Now().Format("2006-01-02 15:04:05"),
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
		exchangeResultHandler(&req,tx.Hash(),err,c)

	case app.SlotToEth:
		amountWrapper,_:=new(big.Float).Mul(amount,app.TokenBase).Int(nil)
		tx,err:=pvc.ExchangeForEther(app.UserAddr,amountWrapper)
		exchangeResultHandler(&req,tx.Hash(),err,c)
	}
}

func machineStatusChange(state int, c *websocket.Conn){
	res:= &app.MachineStatusChangeRes{
		Gcuid: app.NotifyMachineStatusChange,
		State:state,
		MachineId:app.MachineId,
		Act:app.MachineStatusChange,
		CreatedDate: time.Now().Format("2006-01-02 15:04:05"),
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
		&app.PostRes{
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
		&app.PostRes{
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

func updateResultHandler(req *app.PostTokenUseOrRewardReq, txHash common.Hash, updateError error, c *websocket.Conn){
	if updateError!=nil {
		log.Println(updateError.Error())
		sendError(app.PostTokenUseOrReward,app.ClientErrorCode,app.ClientFormatError,c)
		return
	}
	res:= &app.PostTokenUserOrRewardRes {
		&app.PostRes{
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
	var state int
	var token string
	if err!=nil {
		tokenUpdate = "0"
		state = 0
	} else {
		var txType int
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
			CreatedDate:time.Now().Format("2006-01-02 15:04:05"),
		}
		db.LPush(app.User+":"+string(app.Slot),tx)
	}

	rawToken,err:=pvc.GetToken(app.UserAddr)
	if err!=nil {
		log.Println(err.Error())
		token = "0"
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
		// "2006-01-02 15:04:05 is the birth day of golang, fixed format"
		CreatedDate: time.Now().Format("2006-01-02 15:04:05"),
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
		updateResultHandler(&req,tx.Hash(),err,c)
	case app.Reward:
		tx,err:=pvc.Reward(app.UserAddr,amountWrapper)
		updateResultHandler(&req,tx.Hash(),err,c)
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
		}
		json.Unmarshal(data, &kvs)
		gid,ok := kvs["gcuid"]
		if !ok {
			log.Println(errors.New("gcuid not exist"))
		}

		gcuid:= gid.(int)

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
			PostExchangeHandler(data,c)
		//case app.NotifyMachineStatusChange:
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

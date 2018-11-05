// Package main provides ...
package main

import (
	"contract"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/go-redis/redis"
	h "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/goware/disque"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	app "appClient"
)


// websocket gcuid


const (
	chainConfigJson = "src/etc/chainConfig.json"
	disqueConfigJson = "src/etc/disqueConfig.json"
	redisConfigJson = "src/etc/redisConfig.json"
)

const (
	ToPbc = "ToPbc"
	ToPvc = "ToPvc"
	AvatarToPbc ="AvatarToPbc"
	AvatarToPvc ="AvatarToPvc"
	User = "0x20A40B83a495DD2fbbE33E0b6ad119B09F09151f"
)

var (
	pbc *contract.Pbc
	pvc *contract.Pvc
	jobs *disque.Pool
	db *redis.Client

)

type RedisConfig struct {
	Port string `json:"port"`
	Password string `json:"password"`
	DB int `json:"DB"`
}

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

type PublicConfig struct {
	BasicChainConfig
}

type PrivateConfig struct {
	BasicChainConfig
}

type ChainConfig struct {
	Pub *PublicConfig  `json:"public"`
	Pri *PrivateConfig `json:"private"`
}

func getToken(w http.ResponseWriter, r *http.Request) {
	var (
		err         error
		tokenNumber *big.Int
	)
	vars := mux.Vars(r)
	user := vars["user"]
	kind := vars["chain"]

	if kind == "public" {
		tokenNumber, err = pbc.GetToken(user)
	} else if kind == "private" {
		tokenNumber, err = pvc.GetToken(user)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	tokenNumberWrapper, err := json.Marshal(tokenNumber)

	w.Write(tokenNumberWrapper)
}

func getNonce(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		nonce uint64
	)

	vars := mux.Vars(r)
	user := vars["user"]
	kind := vars["chain"]

	if kind == "public" {
		nonce, err = pbc.GetNonce(user)
	} else if kind == "private" {
		nonce, err = pvc.GetNonce(user)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	nonceWrapper, err := json.Marshal(nonce)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(nonceWrapper)
}

func getEther(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		ether *big.Int
	)
	vars := mux.Vars(r)
	user := vars["user"]
	kind := vars["chain"]
	if kind == "public" {
		ether, err = pbc.GetEther(user)
	} else if kind == "private" {
		ether, err = pvc.GetEther(user)
	}
	//log.Println("kind:",kind,"Deal with get Ether, ether amount: ", ether, "user: ", user)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	etherWrapper, err := json.Marshal(ether)

	w.Write(etherWrapper)
}

func updateToken(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]
	var amount int
	data, err := ioutil.ReadAll(r.Body)
	log.Println(string(data))
	err = json.Unmarshal(data,&amount)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if amount < 0 {
		log.Println("consume")
		err = pvc.Consume(user, big.NewInt(int64(-amount)))
	} else if amount > 0 {
		err = pvc.Reward(user, big.NewInt(int64(amount)))
	}
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Deal with consume/award amount:", amount, "user: ", user)
}

func nftTransfer(w http.ResponseWriter, r *http.Request) {
	log.Println("receive a nft transfer")
	var payload ExNFTPayLoad
	data, err := ioutil.ReadAll(r.Body)
	if err!=nil {
		log.Println("can not parse transfer payload")
	}
	json.Unmarshal(data,&payload)

	rawTxBytes, err := hex.DecodeString(payload.Tx)
	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

	vars := mux.Vars(r)
	kind := vars["chain"]
	if kind == "public" {
		err = pbc.SendTransaction(tx)
	} else if kind == "private" {
		err = pvc.SendTransaction(tx)
	}
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var status string
	switch payload.Kind {
	case AvatarToPbc:
		status = contract.NFTPvcWFSig
	case AvatarToPvc:
		status = contract.NFTPvcWFPay
	}
	kvs := make(map[string]interface{})
	kvs["kind"]=payload.Kind
	kvs["status"]=status
	kvs["tokenId"]=payload.TokenId.Uint64()
	_, err = db.HMSet(tx.Hash().String(), kvs).Result()
	if err!=nil {
		log.Println(err.Error())
	}
	db.SAdd(payload.User, tx.Hash().String())
}

func transfer(w http.ResponseWriter, r *http.Request) {
	log.Println("receive a transfer")
	var payload ExTokensPayLoad
	data, err := ioutil.ReadAll(r.Body)
	log.Println(string(data))
	if err!=nil {
		log.Println("can not parse transfer payload")
	}
	json.Unmarshal(data,&payload)

	rawTxBytes, err := hex.DecodeString(payload.Tx)
	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

	vars := mux.Vars(r)
	kind := vars["chain"]
	if kind == "public" {
		err = pbc.SendTransaction(tx)
	} else if kind == "private" {
		err = pvc.SendTransaction(tx)
	}
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var status string
	switch payload.Kind {
	case ToPbc:
		status = contract.PvcWFSig
	case ToPvc:
		status = contract.PvcWFPay
	}

	kvs:=make(map[string]interface{})
	kvs["kind"]=payload.Kind
	kvs["status"]=status
	kvs["amount"]=payload.Amount
	db.HMSet(tx.Hash().String(),kvs)
	db.SAdd(payload.User, tx.Hash().String())
}

func mint(w http.ResponseWriter, r *http.Request) {
	log.Println("receive a mint")
	tx, _ := ioutil.ReadAll(r.Body)

	var err error
	vars := mux.Vars(r)
	tokenId, err := strconv.ParseInt(vars["tokenId"], 10, 64)

	if err != nil {
		log.Println("can not parse tokenId", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = pvc.Mint(string(tx), big.NewInt(tokenId))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	avatar, err := pvc.GetAvatar(big.NewInt(tokenId))
	if err!=nil {
		log.Println("can not get avatar")
	}

	avatarWrapper, err:= json.Marshal(avatar)

	if err!=nil {
		log.Println("can not unmarshal json")
	}

	w.Write(avatarWrapper)

}

func getNFT(w http.ResponseWriter, r *http.Request) {

	var err error
	vars := mux.Vars(r)
	user := vars["user"]
	kind := vars["chain"]

	ownedAvatar, err := pvc.OwnedTokens(user)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if ownedAvatar.Int64()==0 {
		http.Error(w,errors.New("you don not have an avatar").Error(), http.StatusBadRequest)
		return
	}

	var avatar *contract.Avatar
	if kind=="private" {
		avatar, err = pvc.GetAvatar(ownedAvatar)
	} else {
		avatar, err = pbc.GetAvatar(ownedAvatar)
	}

	if err!=nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	avatarWrapper, _:= json.Marshal(avatar)
	w.Write(avatarWrapper)
}

func upgradeNFT(w http.ResponseWriter, r *http.Request){
	log.Println("receive a upgradeNFT")

	var err error
	vars := mux.Vars(r)
	tokenId, err := strconv.ParseInt(vars["tokenId"], 10, 64)

	err = pvc.Upgrade(big.NewInt(tokenId))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func equipWeapon(w http.ResponseWriter, r *http.Request){
	log.Println("receive buy Weapon")
	user, _ := ioutil.ReadAll(r.Body)

	var err error
	var amount int64 = 2
	vars := mux.Vars(r)
	tokenId, err := strconv.ParseInt(vars["tokenId"], 10, 64)

	//owner := hex.EncodeToString(pvc.OwnerOf(big.NewInt(tokenId)).Bytes())
	//if owner != string(user) {
	//	log.Println("wrong owner")
	//	http.Error(w, errors.New("wrong  owner").Error(), http.StatusInternalServerError)
	//	return
	//}

	err = pvc.Consume(string(user),big.NewInt(amount))

	if err!=nil {
		log.Println("buy weapon fail maybe not enough money")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}



	err = pvc.EquipWeapon(string(user),big.NewInt(tokenId))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		pvc.Reward(string(user),big.NewInt(amount))
	}
}

func equipArmor(w http.ResponseWriter, r *http.Request) {
	log.Println("receive buy Armor")
	user, _ := ioutil.ReadAll(r.Body)

	var err error
	var amount int64 = 2
	vars := mux.Vars(r)
	tokenId, err := strconv.ParseInt(vars["tokenId"], 10, 64)

	//owner := hex.EncodeToString(pvc.OwnerOf(big.NewInt(tokenId)).Bytes())
	//if owner != string(user) {
	//	log.Println("wrong owner")
	//	http.Error(w, errors.New("wrong  owner").Error(), http.StatusInternalServerError)
	//	return
	//}

	err = pvc.Consume(string(user),big.NewInt(amount))
	if err!=nil {
		log.Println("buy armor fail maybe not enough money")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = pvc.EquipArmor(string(user),big.NewInt(tokenId))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		pvc.Reward(string(user),big.NewInt(amount))
	}

}

func messagePush(w http.ResponseWriter, r *http.Request) {
	log.Println("websocket connected")
	var upgrader = websocket.Upgrader{}
	upgrader.CheckOrigin = func(rq *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	_, message, err:= c.ReadMessage()
	user:= strings.ToLower(string(message))
	log.Println(user, "connect websocket")
	defer c.Close()
	for {
		job ,err:=jobs.Get(contract.PbcPayNFTQueue+user,
			contract.PvcPayNFTQueue+user,
			contract.PbcPayQueue+user,
			contract.PvcPayQueue+user,
			)
		if err!=nil {
			log.Println(err.Error())
			continue
		}
		log.Println(job.Data)
		jobs.Ack(job)
		err = c.WriteMessage(websocket.TextMessage, []byte(job.Data))
		if err!=nil {
			log.Println(err.Error())
		}
	}
}

func SigninHandler (data []byte,c *websocket.Conn) {
	var signin app.SigninReq
	json.Unmarshal(data, &signin)
	if signin.LoginCode!=app.LoginCode {
		reqError:= app.Error{
			Status: app.FailedStatus,
			Code: app.ErrorCode,
			Gcuid: app.Signin,
			Reason: app.InvalidRequestInfo,
		}
		reqErrorWrapper,err:=json.Marshal(reqError)
		if err!=nil {
			log.Println(err.Error())
		}
		c.WriteMessage(websocket.TextMessage, reqErrorWrapper)
	} else {
		res:= app.SigninRes{
			Status:app.OkStatus,
			Gcuid:app.Signin,
			Guuid:app.Guuid,
		}
		resWrapper, err:=json.Marshal(res)
		if err!=nil {
			log.Println(err.Error())
		}
		c.WriteMessage(websocket.TextMessage, resWrapper)
	}
}

func SignoutHandler (data []byte,c *websocket.Conn) {
	res:= app.SignoutRes{
		&app.PostRes{
			Status: app.OkStatus,
			Gcuid: app.Signout,
		},
	}
	resWrapper, err:=json.Marshal(res)
	if err!=nil {
		log.Println(err.Error())
	}
	c.WriteMessage(websocket.TextMessage, resWrapper)
}

func

func GetWalletsAndMachineHandler(data []byte, c *websocket.Conn) {

}

func GetWalletsHandler(data []byte, c *websocket.Conn) {

}

func GetTransactionsHandler(data []byte, c *websocket.Conn) {

}

func GetExchangeRateHandler(data []byte, c *websocket.Conn) {

}

func PostExchangeHandler(data []byte, c *websocket.Conn){

}

func PostQRCodeHandler(data []byte, c *websocket.Conn) {

}

func MachineLogoutHandler(data []byte, c *websocket.Conn) {

}

func PostTokenUserOrRewardHandler(data []byte, c *websocket.Conn) {

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

		gcuid,err:= gid.(int)
		if err!=nil {
			log.Println(err.Error())
		}

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
		//case app.NotifyTokenChange:
		//case app.NotifyExchangeResult:
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

	go requestHandler(c)
	defer c.Close()
	for {
		job ,err:=jobs.Get(contract.PbcPayNFTQueue+User,
			contract.PvcPayNFTQueue+User,
			contract.PbcPayQueue+User,
			contract.PvcPayQueue+User,
		)
		if err!=nil {
			log.Println(err.Error())
			continue
		}
		log.Println(job.Data)
		jobs.Ack(job)
		err = c.WriteMessage(websocket.TextMessage, []byte(job.Data))
		if err!=nil {
			log.Println(err.Error())
		}
	}
}

// initial contract
func init() {
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

	//for test use
	pvc.AddGameMachine("0x20A40B83a495DD2fbbE33E0b6ad119B09F09151f")
	pbc.AddGameMachine("0x20A40B83a495DD2fbbE33E0b6ad119B09F09151f")
	//
}

// initial disque
func init() {
	var dc DisqueConfig
	data, err:= ioutil.ReadFile(disqueConfigJson)
	if err!=nil {
		panic("can not read chain config file")
	}
	json.Unmarshal(data,&dc)

	jobs, err =disque.New(dc.Port)
	if err!=nil {
		panic(err.Error())
	}
}

func init() {
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
	defer jobs.Close()
	go pbc.EventReceiver(pvc, jobs)
	go pvc.EventReceiver(pbc, jobs)
	go contract.Consumer(jobs,pvc,pbc,db)
	r := mux.NewRouter()

	//r.HandleFunc("/api/v1/{chain:public|private}/tokens/{user}", getToken).Methods("GET")
	//r.HandleFunc("/api/v1/{chain:public|private}/nonce/{user}", getNonce).Methods("GET")
	//r.HandleFunc("/api/v1/{chain:public|private}/ether/{user}", getEther).Methods("GET")
	//r.HandleFunc("/api/v1/private/tokens/{user}", updateToken).Methods("PUT")
	//r.HandleFunc("/api/v1/{chain:public|private}/transfer/tokens", transfer).Methods("PUT")
	//r.HandleFunc("/api/v1/{chain:public|private}/transfer/nft", nftTransfer).Methods("PUT")
	//r.HandleFunc("/api/v1/nft/{tokenId}", mint).Methods("POST")
	//r.HandleFunc("/api/v1/{chain:public|private}/nft/{user}", getNFT).Methods("GET")
	//r.HandleFunc("/api/v1/nft/{tokenId}/level", upgradeNFT).Methods("PUT")
	//r.HandleFunc("/api/v1/nft/{tokenId}/weapon", equipWeapon).Methods("PUT")
	//r.HandleFunc("/api/v1/nft/{tokenId}/armor", equipArmor).Methods("PUT")
	//r.HandleFunc("/ws",messagePush)
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

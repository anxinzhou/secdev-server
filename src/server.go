// Package main provides ...
package main

import (
	"contract"
	"encoding/json"
	"errors"
	"fmt"
	h "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/goware/disque"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

const (
	chainConfigJson = "src/etc/chainConfig.json"
	disqueConfigJson = "src/etc/disqueConfig.json"
)

var (
	pbc *contract.Pbc
	pvc *contract.Pvc
	jobs *disque.Pool
)

type DisqueConfig struct {
	Port string `json:"port"`
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
	err := json.NewDecoder(r.Body).Decode(&amount)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("updating private token! amount:", amount)

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

func transfer(w http.ResponseWriter, r *http.Request) {
	log.Println("receive a transfer")
	tx, _ := ioutil.ReadAll(r.Body)

	var err error
	vars := mux.Vars(r)
	kind := vars["chain"]
	if kind == "public" {
		err = pbc.SendTransaction(string(tx))
	} else if kind == "private" {
		err = pvc.SendTransaction(string(tx))
	}
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
		log.Println("can not get avatar")
	}

	avatarWrapper, err:= json.Marshal(avatar)

	if err!=nil {
		log.Println("can not unmarshal json")
	}

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
	defer c.Close()
	for {
		job,_:=jobs.Get("pbcNFT"+user)
		time.Sleep(1*time.Second)
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		//err = c.WriteMessage(websocket.TextMessage, []byte("dsad"))
		//if err != nil {
		//	log.Println("write:", err)
		//	break
		//}
	}
}

func init() {
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

func main() {
	//runtime.GOMAXPROCS(8)
	defer jobs.Close()
	go pbc.EventReceiver(pvc, jobs)
	go pvc.EventReceiver(pbc, jobs)
	go contract.Consumer(jobs,pvc,pbc)
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/{chain:public|private}/tokens/{user}", getToken).Methods("GET")
	r.HandleFunc("/api/v1/{chain:public|private}/nonce/{user}", getNonce).Methods("GET")
	r.HandleFunc("/api/v1/{chain:public|private}/ether/{user}", getEther).Methods("GET")
	r.HandleFunc("/api/v1/private/tokens/{user}", updateToken).Methods("PUT")
	r.HandleFunc("/api/v1/{chain:public|private}/transfer", transfer).Methods("PUT")
	r.HandleFunc("/api/v1/nft/{tokenId}", mint).Methods("POST")
	r.HandleFunc("/api/v1/{chain:public|private}/nft/{user}", getNFT).Methods("GET")
	r.HandleFunc("/api/v1/nft/{tokenId}/level", upgradeNFT).Methods("PUT")
	r.HandleFunc("/api/v1/nft/{tokenId}/weapon", equipWeapon).Methods("PUT")
	r.HandleFunc("/api/v1/nft/{tokenId}/armor", equipArmor).Methods("PUT")
	r.HandleFunc("/ws",messagePush)

	fmt.Println("Running http server")
	http.ListenAndServe(":4000",
		h.CORS(h.AllowedMethods([]string{"get", "options", "post", "put", "head"}),
			h.AllowedOrigins([]string{"*"}))(r))
}

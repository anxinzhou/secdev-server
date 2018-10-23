// Package main provides ...
package main

import (
	"contract"
	"encoding/json"
	"fmt"
	h "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
)

const (
	chainConfigJson = "src/etc/config.json"
)

var (
	pbc *contract.Pbc
	pvc *contract.Pvc
)

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
	tx, _ :=ioutil.ReadAll(r.Body)

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

func main() {
	//runtime.GOMAXPROCS(8)
	go pbc.EventReceiver(pvc)
	go pvc.EventReceiver(pbc)
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/{chain:public|private}/tokens/{user}", getToken).Methods("GET")
	r.HandleFunc("/api/v1/{chain:public|private}/nonce/{user}", getNonce).Methods("GET")
	r.HandleFunc("/api/v1/{chain:public|private}/ether/{user}", getEther).Methods("GET")
	r.HandleFunc("/api/v1/private/tokens/{user}", updateToken).Methods("PUT")
	r.HandleFunc("/api/v1/{chain:public|private}/transfer", transfer).Methods("PUT")

	fmt.Println("Running http server")
	http.ListenAndServe(":4000",
		h.CORS(h.AllowedMethods([]string{"get", "options", "post", "put", "head"}),
			h.AllowedOrigins([]string{"*"}))(r))
}

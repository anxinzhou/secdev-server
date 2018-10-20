// Package main provides ...
package main

import (
	"contract"
	"encoding/json"
	"fmt"
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
	Pub * PublicConfig   `json:"public"`
	Pri * PrivateConfig  `json:"private"`
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
		return
	}
	tokenNumberWrapper, err := json.Marshal(tokenNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Deal with get Token, token amount: ", tokenNumber.String())
	w.Header().Set("Content-Type", "application/json")
	w.Write(tokenNumberWrapper)
}

func getUserNonce(w http.ResponseWriter, r *http.Request) {
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
	} else {
		log.Fatal("this should not happend")
		return
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
	w.Header().Set("Content-Type", "application/json")
	w.Write(nonceWrapper)
}

func getUserEther(w http.ResponseWriter, r *http.Request) {
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
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	etherWrapper, err := json.Marshal(ether)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(etherWrapper)
}

func updateToken(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]
	var amount int
	err := json.NewDecoder(r.Body).Decode(&amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if amount < 0 {
		err = pvc.Consume(user, big.NewInt(int64(-amount)))
	} else if amount > 0 {
		err = pvc.Award(user, big.NewInt(int64(amount)))
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func transfer(w http.ResponseWriter, r *http.Request) {
	var tx string
	err := json.NewDecoder(r.Body).Decode(&tx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)
	kind := vars["user"]
	if kind == "public" {
		err = pbc.SendTransaction(tx)
	} else if kind == "private" {
		err = pvc.SendTransaction(tx)
	}
	if err != nil {
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
	json.Unmarshal(data,&cc)

	pbc = contract.NewPbc(cc.Pub.Port,cc.Pub.Keystore,cc.Pub.Password,cc.Pri.Address)
	pvc = contract.NewPvc(cc.Pri.Port, cc.Pri.Keystore, cc.Pri.Password,cc.Pri.Address)
}

func main() {
	//runtime.GOMAXPROCS(8)
	r:=mux.NewRouter()
	r.HandleFunc("/{chain:public|private}/tokens/{user}",getToken).Methods("GET")
	r.HandleFunc("/nonce/{user}", getUserNonce).Methods("GET")
	r.HandleFunc("/{chain:public|private}/ether/{user}", getUserEther).Methods("GET")
	r.HandleFunc("/private/tokens/{user}",updateToken).Methods("PUT")
	r.HandleFunc("/{chain:public|private}/transfer",transfer).Methods("PUT")

	fmt.Println("Running http server")
	log.Fatal(http.ListenAndServe(":4000", r))
}

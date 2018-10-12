// Package main provides ...
package main

import (
	"contract"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"runtime"
)

const (
	chainConfigJson = "./etc/config.json"
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
	Pub PublicConfig
	Pri PrivateConfig
}

func getToken(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		err         error
		tokenNumber *big.Int
	)
	user := ps.ByName("user")
	kind := ps.ByName("chain")
	if kind == "public" {
		tokenNumber, err = pbc.GetToken(user)
	} else if kind == "private" {
		tokenNumber, err = pvc.GetToken(user)
	} else {
		log.Fatal("this should not happen")
		return
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
	w.Header().Set("Content-Type", "application/json")
	w.Write(tokenNumberWrapper)
}

func getUserNonce(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		err   error
		nonce uint64
	)
	user := ps.ByName("user")
	kind := ps.ByName("chain")
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

func getUserEther(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		err   error
		ether *big.Int
	)
	user := ps.ByName("user")
	kind := ps.ByName("chain")
	if kind == "public" {
		ether, err = pbc.GetEther(user)
	} else if kind == "private" {
		ether, err = pvc.GetEther(user)
	} else {
		log.Fatal("this should not happen")
		return
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

func updateToken(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user := ps.ByName("user")
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

func transfer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var tx string
	err := json.NewDecoder(r.Body).Decode(&tx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	kind := ps.ByName("chain")
	if kind == "public" {
		err = pbc.SendTransaction(tx)
	} else if kind == "private" {
		err = pvc.SendTransaction(tx)
	} else {
		log.Fatal("this should not happend")
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func init() {
	cc := new(ChainConfig)
	data, err := ioutil.ReadFile(chainConfigJson)
	if err != nil {
		panic("can not read chain config file")
	}
	err = json.Unmarshal(data, cc)
	if err != nil {
		panic("can not unmarsh json")
	}
	pbc = new(contract.Pbc)
	pbc.Connect(cc.Pub.Port)
	pbc.LoadKey(cc.Pub.Keystore, cc.Pub.Password)
	pbc.LoadContract(cc.Pub.Address)

	pvc = new(contract.Pvc)
	pvc.Connect(cc.Pri.Port)
	pvc.LoadKey(cc.Pri.Keystore, cc.Pri.Password)
	pvc.LoadContract(cc.Pri.Address)
}

func main() {
	runtime.GOMAXPROCS = 8
	router := httprouter.New()
	router.GET("/:chain(public|private)/tokens/:user", getToken)
	router.GET("/:chain(public|private)/nonce/:user", getUserNonce)
	router.GET("/:chain(public|private)/ether/:user", getUserEther)
	router.PUT("/private/tokens/:user", updateToken)
	router.PUT("/:chain(public|private)/transfer", transfer)
	log.Fatal(http.ListenAndServe(":4000", router))
}

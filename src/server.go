// Package main provides ...
package main

import (
	"contract"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"net/rpc"
)

func getToken(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user := ps.ByName("user")
	kind := ps.ByName("chain")
	if kind == "public" {
		tokenNumber, err := pbc.GetToken(user)
	} else if kind == "private" {
		tokenNumber, err := pvc.GetToken(user)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	w.Header().set("Content-Type", "application/json")
	w.Write(tokenNumberWrapper)
}

func getUserNonce(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user := ps.ByName("user")
	kind := ps.ByName("chain")
	if kind == "public" {
		nonce, err := pbc.GetNonce(user)
	} else if kind == "private" {
		nonce, err := pvc.GetNonce(user)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	w.Header().set("Content-Type", "application/json")
	w.Write(nonceWrapper)
}

func getUserEther(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user := ps.ByName("user")
	kind := ps.ByName("chain")
	if kind == "public" {
		nonce, err := pbc.GetEther(user)
	} else if kind == "private" {
		nonce, err := pvc.GetEther(user)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	w.Header().set("Content-Type", "application/json")
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
		err = pvc.consume(user, -amount)
	} else if amount > 0 {
		err = pvc.reward(user, amount)
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
		nonce, err := pbc.sendSignedTransaction(user)
	} else if kind == "private" {
		nonce, err := pvc.sendSignedTransaction(user)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

var c *Client

func main() {
	c, err := rpc.Dial("tcp", "127.0.0.1:8000")
	router := httprouter.New()
	router.GET("/:chain(public|private)/tokens/:user", getToken)
	router.GET("/:chain(public|private)/nonce/:user", getUserNonce)
	router.GET("/:chain(public|private)/ether/:user", getUserEther)
	router.PUT("/private/tokens/:user", updateToken)
	router.PUT("/:chain(public|private)/transfer", transfer)

	log.Fatal(http.ListenAndServe(":4000", router))
}

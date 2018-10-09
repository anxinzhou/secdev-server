// Package main provides ...
package main

import (
	"contract"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func getToken(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user := ps.ByName("user")
	tokenNumber, err := pvc.GetToken(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().set("Content-Type", "application/json")
	w.Write(tokenNumber)
}

func getUserNonce(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func getUserNonce(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func getUserNonce(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func getUserNonce(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func getUserNonce(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func getUserNonce(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func getUserNonce(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func getUserNonce(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func main() {
	router := httprouter.New()
	router.GET("/tokens/:user", getToken)
	router.GET("/nonce/:user", getUserNonce)
	router.GET("/ether/:user", getUserEther)
	router.PUT("/tokens/:user", updateToken)
	router.POST("/validation", validate)
	router.GET("/publicTokens/:user", getPubTOken)
	router.PUT("/publicTOkens/:user", updatePubTOken)
	router.GET("/publicNOnce/:user", getPubUserNonce)
	router.GET("/publicEther/:user", getPubUserEther)
	router.POST("/faucet", requireEther)

	log.Fatal(http.ListenAndServe(":4000", router))
}

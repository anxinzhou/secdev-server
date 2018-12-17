package eos

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
)

const (
	eosServerUrl = "http://127.0.0.1:3000/api/v1"
)

type GetBalanceRes struct {
	Balance *big.Float `json:"balance"`
}

func GetBalance(account string) (*big.Float,error) {
	resp, err := http.Get(eosServerUrl+"/balance/"+account)
	if(err!=nil) {
		return nil,err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err!=nil {
		return nil,err
	}

	var balancePayload GetBalanceRes
	json.Unmarshal(data,&balancePayload)
	return balancePayload.Balance,nil
}

func Update(account string, amount *big.Float) (error) {
	_, err := http.PostForm(eosServerUrl+"/balance"+account,
		url.Values{"key": {"Value"}, "id": {"123"}})
	if err!=nil {
		return err
	}
	return nil
}
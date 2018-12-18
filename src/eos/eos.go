package eos

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
)

const (
	eosServerUrl = "http://127.0.0.1:8000/api/v1"
)

type UpdateTokenReq struct {
	Account string `json:"account"`
	Amount *big.Float `json:"amount"`
}

type GetTokenBalanceRes struct {
	Balance *big.Float `json:"balance"`
	Symbol  string `json:"symbol"`
}

func GetTokenBalance(account string) (*big.Float,error) {
	resp, err := http.Get(eosServerUrl+"/balance/"+account)
	defer resp.Body.Close()
	if(err!=nil) {
		return nil,err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err!=nil {
		return nil,err
	}

	var balancePayload GetTokenBalanceRes
	json.Unmarshal(data,&balancePayload)
	return balancePayload.Balance,nil
}

func updateToken(account string, amount *big.Float) (error) {
	updateTokenPayload := &UpdateTokenReq{
		Account:account,
		Amount: amount,
	}
	data, err:= json.Marshal(updateTokenPayload)
	if err!=nil {
		return nil
	}


	url:= eosServerUrl+"/balance/"+account
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func Consume(account string, amount *big.Float) (error) {
	return updateToken(account,amount)
}

func Reward(account string, amount *big.Float) (error) {
	return updateToken(account,big.NewFloat(0).Neg(amount))
}

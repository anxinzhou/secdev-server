package contract

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"log"
	"math/big"
	"token/publicSlot"
)

const (
	TargetGasFee = "50000000000000"
	TxFee = "0.00005" //ether
)

type Pbc struct {
	ERC20
	Instance *publicSlot.PublicSlot
}

func NewPbc(port string, keystore string, password string, contractAdd common.Address) *Pbc {
	pbc := new(Pbc)
	pbc.Connect(port)
	pbc.LoadKey(keystore, password)
	pbc.Address = contractAdd
	pbc.Instance,_ = publicSlot.NewPublicSlot(contractAdd,pbc.Client)
	pbc.LoadABI(string(publicSlot.PublicSlotABI))
	return pbc
}

func NewPbcFromJSON(configFile string) *Pbc {
	type basicChainConfig struct {
		Port     string `json:"port"`
		Keystore string `json:"keystore"`
		Password string `json:"password"`
		Address  string `json:"address"`
	}
	// initial contract
	var c basicChainConfig
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Println(err.Error())
		panic("can not read chain config file")
	}
	json.Unmarshal(data, &c)
	return NewPbc(c.Port, c.Keystore, c.Password, common.HexToAddress(c.Address))
}

func (p *Pbc) GetExchangeRate() (string, error) {
	exchangeRate,err:=p.Instance.ExchangeRate(nil)
	if err!=nil {
		return "", err
	}
	exchangeBase,err:=p.Instance.ExchangeBase(nil)
	if err!=nil {
		return "",err
	}
	rate:=new(big.Float).Quo(new(big.Float).SetInt(exchangeRate),new(big.Float).SetInt(exchangeBase)).String()
	return rate,nil
}

func (p *Pbc) ExchangeForToken(address common.Address, amount *big.Int) (error) {
	return p.sendWithTargetFeeAndValue("exchangeForToken",TargetGasFee,amount,address)
}

func (p* Pbc) Create(address common.Address, tokenID *big.Int) (error) {
	return p.sendWithTargetFee("create",TargetGasFee,address, tokenID)
}

func (p *Pbc) ExchangeForEther(address common.Address, amount *big.Int) (error){
	return p.sendWithTargetFee("exchangeForEther",TargetGasFee,address,amount)
}
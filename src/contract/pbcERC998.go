package contract

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"log"
	publicERC998 "token/PbcERC998"
)

type PbcERC998 struct {
	ERC998
	Instance *publicERC998.PbcERC998
}

func NewPbcERC998(port string, keystore string, password string, contractAdd common.Address) *PbcERC998 {
	pbcERC998 := new(PbcERC998)
	pbcERC998.Connect(port)
	pbcERC998.LoadKey(keystore, password)
	pbcERC998.Address = contractAdd
	pbcERC998.Instance,_ = publicERC998.NewPbcERC998(contractAdd,pbcERC998.Client)
	pbcERC998.LoadABI(string(publicERC998.PbcERC998ABI))
	return pbcERC998
}

func NewPbcERC998FromJSON(configFile string) *PbcERC998 {
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
	return NewPbcERC998(c.Port, c.Keystore, c.Password, common.HexToAddress(c.Address))
}


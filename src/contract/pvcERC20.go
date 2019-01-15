package contract

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"log"
	"token/privateSlot"
)

type Pvc struct {
	ERC20
	Instance *privateSlot.PrivateSlot
}

func NewPvc(port string, keystore string, password string, contractAdd common.Address) *Pvc {
	pvc := new(Pvc)
	pvc.Connect(port)
	pvc.LoadKey(keystore, password)
	pvc.Address = contractAdd
	pvc.Instance,_ = privateSlot.NewPrivateSlot(contractAdd,pvc.Client)
	pvc.LoadABI(string(privateSlot.PrivateSlotABI))
	return pvc
}

func NewPvcFromJSON(configFile string) *Pvc {
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
	return NewPvc(c.Port, c.Keystore, c.Password, common.HexToAddress(c.Address))
}
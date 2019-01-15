package contract

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"log"
	"math/big"
	privateERC998 "token/PvcERC998"
)

type PvcERC998 struct {
	ERC998
	Instance  *privateERC998.PvcERC998
}

func NewPvcERC998(port string, keystore string, password string, contractAdd common.Address) *PvcERC998 {
	pvcERC998 := new(PvcERC998)
	pvcERC998.Connect(port)
	pvcERC998.LoadKey(keystore, password)
	pvcERC998.Address = contractAdd
	pvcERC998.Instance,_ = privateERC998.NewPvcERC998(contractAdd,pvcERC998.Client)
	pvcERC998.LoadABI(string(privateERC998.PvcERC998ABI))
	return pvcERC998
}

func NewPvcERC998FromJSON(configFile string) *PvcERC998 {
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
	return NewPvcERC998(c.Port, c.Keystore, c.Password, common.HexToAddress(c.Address))
}

func (p *PvcERC998) ChangeOwner(owner common.Address, tokenId *big.Int) error {
	return p.send("changeOwner", owner, tokenId)
}

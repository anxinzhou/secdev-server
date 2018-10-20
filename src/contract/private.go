package contract

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"sync/atomic"
	"token/privateSlot"
)

type Pvc struct {
	Contract
	Instance *privateSlot.PrivateSlot
}

func NewPvc(port string, keystore string, password string, contractAdd string) *Pvc {
	pvc := new(Pvc)
	pvc.Connect(port)
	pvc.LoadKey(keystore, password)
	pvc.LoadContract(contractAdd)
	return pvc
}

func (p *Pvc) LoadContract(rawAddress string) {
	address := common.HexToAddress(rawAddress)
	p.Instance, _ = privateSlot.NewPrivateSlot(address, p.Client)
}

func (p *Pvc) GetToken(rawAddress string) (*big.Int, error) {
	address := common.HexToAddress(rawAddress)
	token, err := p.Instance.BalanceOf(nil, address)
	return token, err
}

func (p *Pvc) Consume(rawAddress string, amount *big.Int) error {
	var err error

	address := common.HexToAddress(rawAddress)
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce, big.NewInt(0))

	_, err = p.Instance.Consume(auth, address, amount)
	return err
}

func (p *Pvc) Pay(rawAddress string, amount *big.Int, transactionHash [32]byte) error {
	var err error

	address := common.HexToAddress(rawAddress)
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce, big.NewInt(0))

	_, err = p.Instance.Pay(auth, address, amount, transactionHash)
	return err
}

func (p *Pvc) Award(rawAddress string, amount *big.Int) error {
	var err error

	address := common.HexToAddress(rawAddress)
	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce, big.NewInt(0))

	_, err = p.Instance.Award(auth, address, amount)
	return err
}

func (p *Pvc) Deploy(initialSupply *big.Int, requiredSignatures *big.Int, authorities []common.Address) (common.Address, error) {
	var err error

	nonce:= atomic.AddUint64(&p.txConfig.nonce, 1)
	auth:= NewAuth(p.txConfig.key.PrivateKey, nonce, big.NewInt(0))

	address, _, _, err := privateSlot.DeployPrivateSlot(auth, p.Client, initialSupply, requiredSignatures, authorities)
	return address, err
}
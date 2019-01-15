package contract

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type ERC20 struct {
	Contract
}

func (c *ERC20) Consume(address common.Address, amount *big.Int) (error) {
	return c.send("consume", address,amount)
}

func (c *ERC20) Reward(address common.Address, amount *big.Int) (error) {
	return c.send("reward", address,amount)
}

func (c *ERC20) Mint(address common.Address, amount *big.Int) (error) {
	return c.send("mint", address, amount)
}

func (c *ERC20) Burn(address common.Address, amount *big.Int) (error) {
	return c.send("burn", address, amount)
}

func (c *ERC20) GetToken(address common.Address) (*big.Int, error) {
	data, err:= c.call("balanceOf",address)
	if err!=nil {
		return nil,err
	}
	return new(big.Int).SetBytes(data),nil
}
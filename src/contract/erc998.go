package contract

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type ERC998 struct{
	Contract
}

func (c *ERC998) Lock(tokenId * big.Int) error {
	return c.send("lock", tokenId)
}

func (c *ERC998) Unlock(tokenId *big.Int) error {
	return c.send("unlock", tokenId)
}

func (c *ERC998) Mint(to common.Address, tokenId *big.Int, uri string) error {
	return c.send("mint", to, tokenId, uri)
}

func (c *ERC998) MintERC721Child(parentTokenId *big.Int, childContract common.Address, childTokenId *big.Int, uri string) error {
	return c.send("mintERC721Child", parentTokenId, childContract, childTokenId, uri)
}

func (c *ERC998) BurnERC721Child(parentTokenId *big.Int, childContract common.Address, childTokenId *big.Int) error {
	return c.send("burnERC721Child", parentTokenId, childContract, childTokenId)
}

func (c *ERC998) MintERC20Child(parentTokenId *big.Int,childContract common.Address,value *big.Int) error {
	return c.send("mintERC20Child", parentTokenId, childContract, value)
}

func (c *ERC998) BurnERC20Child(parentTokenId *big.Int,childContract common.Address,value *big.Int) error {
	return c.send("burnERC20Child", parentTokenId, childContract, value)
}

func (c *ERC998) TokenURI(tokenId *big.Int) (string,error) {
	data,err:= c.call("tokenURI",tokenId)
	if err!=nil {
		return "",nil
	}
	return string(data),nil
}

func (c *ERC998) SetTokenURI(tokenId *big.Int,uri string) error {
	return c.send("setTokenURI", tokenId, uri)
}

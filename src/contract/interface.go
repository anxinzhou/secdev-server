package contract

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type IERC20 interface {
	Consume(common.Address, *big.Int) (error)
	Reward(common.Address, *big.Int) (error)
	Mint(common.Address, *big.Int) (error)
	Burn(common.Address, *big.Int) (error)
	GetToken(common.Address) (*big.Int, error)
}

type IERC721 interface {
	Lock(* big.Int) error
	Unlock(*big.Int) error
	Mint(common.Address, *big.Int, string) error
	MintERC721Child(*big.Int, common.Address, *big.Int, string) error
	BurnERC721Child(*big.Int, common.Address, *big.Int) error
	MintERC20Child(*big.Int, common.Address, *big.Int) error
	BurnERC20Child(*big.Int, common.Address, *big.Int) error
	TokenURI(*big.Int) (string,error)
	SetTokenURI(*big.Int, string) error
}


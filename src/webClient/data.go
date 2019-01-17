package webClient

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

//
const (
	success = 1
	fail = 0
	public = "public"
	private = "private"
)

// only for test
//
var (
	PvcERC721 common.Address
	PbcERC721 common.Address
)

//error
const (
	ServerJsonError = "can not parse json data in server"
	ClientFormatError = "dataformat is not correct"
)

//Gcuid
const (
	createERC998 = 1+iota
	createERC721
	transferERC998
	transferERC721
)

type createERC998Req struct {
	Gcuid int64 `json:"gcuid"`
	Account string `json:"account"`
	TokenId *big.Int `json:"tokenId"`
	Uri string `json:"uri"`
}

type createERC998Res struct {
	Status int64 `json:"status"`
	Gcuid int64 `json:"gcuid"`
	TokenId *big.Int `json:"tokenId"`
}

type createERC721Req struct {
	Gcuid int64 `json:"gcuid"`
	Object string `json:"object"`
	ParentTokenId *big.Int `json:"parentTokenId"`
	ChildContract string `json:"childContract"`
	ChildTokenId *big.Int `json:"childTokenId"`
	Uri string `json:"uri"`
}

type createERC721Res struct {
	Gcuid int64 `json:"gcuid"`
	Status int64 `json:"status"`
	Object string `json:"object"`
	ChildTokenId *big.Int `json:"childTokenId"`
}

type transferERC998Req struct {
	Gcuid int64 `json:"gcuid"`
	Chain string `json:"chain"`
	TokenId *big.Int `json:"tokenId"`
}

type transferERC998Res struct {
	Gcuid int64 `json:"gcuid"`
	Status int64 `json:"status"`
	Chain string `json:"chain"`
	TokenId *big.Int `json:"tokenId"`
}

type transferERC721Req struct {
	Gcuid int64 `json:"gcuid"`
	Chain string `json:"chain"`
	Object string `json:"object"`
	ParentTokenId *big.Int `json:"parentTokenId"`
	ChildContract string `json:"childContract"`
	ChildTokenId *big.Int `json:"childTokenId"`
	Uri string `json:"uri"`
}

type transferERC721Res struct {
	Gcuid int64 `json:"gcuid"`
	Object string `json:"object"`
	Chain string `json:"chain"`
	Status int64 `json:"status"`
	ChildTokenId *big.Int `json:"childTokenId"`
}

type reqSuccess struct {
	Status int64 `json:"status"`
	Gcuid int64 `json:"gcuid"`
}

type reqError struct {
	Status int64 `json:"status"`
	Gcuid int64 `json:"gcuid"`
	Reason string `json:"reason"`
}
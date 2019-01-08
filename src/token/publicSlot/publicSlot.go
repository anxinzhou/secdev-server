// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package publicSlot

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PublicSlotABI is the input ABI used to generate the binding from.
const PublicSlotABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_tokenID\",\"type\":\"uint256\"}],\"name\":\"create\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"reward\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"by\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"consume\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchangeRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"exchangeForEther\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"exchangeForToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchangeBase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NFTName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NFTSymbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// PublicSlotBin is the compiled bytecode used for deploying new contracts.
const PublicSlotBin = `60806040526101f460075560016008556040805190810160405280600681526020017f6176617461720000000000000000000000000000000000000000000000000000815250600c90805190602001906200005c92919062000233565b506040805190810160405280600381526020017f4156540000000000000000000000000000000000000000000000000000000000815250600d9080519060200190620000aa92919062000233565b506040805190810160405280601481526020017f68747470733a2f2f7365636465762e736974652f000000000000000000000000815250600e9080519060200190620000f892919062000233565b5060405162001e4638038062001e468339810180604052810190808051906020019092919080518201929190602001805182019291905050508282826012600260006101000a81548160ff021916908360ff160217905550600260009054906101000a900460ff1660ff16600a0a8302600081905550600054600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508160019080519060200190620001cc92919062000233565b508060039080519060200190620001e592919062000233565b5050505033600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050620002e2565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200027657805160ff1916838001178555620002a7565b82800160010185558215620002a7579182015b82811115620002a657825182559160200191906001019062000289565b5b509050620002b69190620002ba565b5090565b620002df91905b80821115620002db576000816000905550600101620002c1565b5090565b90565b611b5480620002f26000396000f30060806040526004361061011d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde0314610122578063095ea7b3146101b25780630ecaea731461021757806318160ddd1461027c57806321670f22146102a7578063224b5c721461030c57806323b872dd14610371578063313ce567146103f65780633ba0b9a91461042757806340c10f19146104525780635f1e6941146104b757806370a082311461051c5780637317d44a146105735780637e8b3f89146105c157806394f4e679146105ec57806395d89b411461067c5780639dc29fac1461070c578063a42d20ae14610771578063a9059cbb14610801578063c87b56dd14610866578063dd62ed3e1461090c575b600080fd5b34801561012e57600080fd5b50610137610983565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561017757808201518184015260208101905061015c565b50505050905090810190601f1680156101a45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101be57600080fd5b506101fd600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a21565b604051808215151515815260200191505060405180910390f35b34801561022357600080fd5b50610262600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610b4e565b604051808215151515815260200191505060405180910390f35b34801561028857600080fd5b50610291610ce1565b6040518082815260200191505060405180910390f35b3480156102b357600080fd5b506102f2600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610cea565b604051808215151515815260200191505060405180910390f35b34801561031857600080fd5b50610357600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610d43565b604051808215151515815260200191505060405180910390f35b34801561037d57600080fd5b506103dc600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610dea565b604051808215151515815260200191505060405180910390f35b34801561040257600080fd5b5061040b611001565b604051808260ff1660ff16815260200191505060405180910390f35b34801561043357600080fd5b5061043c611014565b6040518082815260200191505060405180910390f35b34801561045e57600080fd5b5061049d600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061101a565b604051808215151515815260200191505060405180910390f35b3480156104c357600080fd5b50610502600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061106f565b604051808215151515815260200191505060405180910390f35b34801561052857600080fd5b5061055d600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061116f565b6040518082815260200191505060405180910390f35b6105a7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506111b8565b604051808215151515815260200191505060405180910390f35b3480156105cd57600080fd5b506105d6611222565b6040518082815260200191505060405180910390f35b3480156105f857600080fd5b50610601611228565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610641578082015181840152602081019050610626565b50505050905090810190601f16801561066e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561068857600080fd5b506106916112ca565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106d15780820151818401526020810190506106b6565b50505050905090810190601f1680156106fe5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561071857600080fd5b50610757600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611368565b604051808215151515815260200191505060405180910390f35b34801561077d57600080fd5b5061078661140b565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156107c65780820151818401526020810190506107ab565b50505050905090810190601f1680156107f35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561080d57600080fd5b5061084c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506114ad565b604051808215151515815260200191505060405180910390f35b34801561087257600080fd5b5061089160048036038101908080359060200190929190505050611529565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156108d15780820151818401526020810190506108b6565b50505050905090810190601f1680156108fe5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561091857600080fd5b5061096d600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506115de565b6040518082815260200191505060405180910390f35b60018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a195780601f106109ee57610100808354040283529160200191610a19565b820191906000526020600020905b8154815290600101906020018083116109fc57829003601f168201915b505050505081565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614151515610a5e57600080fd5b81600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b600082600b600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600a6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610cb4600e8054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ca15780601f10610c7657610100808354040283529160200191610ca1565b820191906000526020600020905b815481529060010190602001808311610c8457829003601f168201915b5050505050610caf84611665565b6117bc565b600960008481526020019081526020016000209080519060200190610cda929190611a83565b5092915050565b60008054905090565b600081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055506001905092915050565b600081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410151515610d9357600080fd5b81600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055506001905092915050565b6000600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548211151515610e7757600080fd5b610f0682600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461188890919063ffffffff16565b600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610f918484846118a9565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3600190509392505050565b600260009054906101000a900460ff1681565b60075481565b600081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555092915050565b600080600754600854840281151561108357fe5b04905082600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101515156110d457600080fd5b82600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055508373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015611167573d6000803e3d6000fd5b505092915050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60008060085460075434028115156111cc57fe5b04905080600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555050919050565b60085481565b6060600c8054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156112c05780601f10611295576101008083540402835291602001916112c0565b820191906000526020600020905b8154815290600101906020018083116112a357829003601f168201915b5050505050905090565b60038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156113605780601f1061133557610100808354040283529160200191611360565b820191906000526020600020905b81548152906001019060200180831161134357829003601f168201915b505050505081565b600081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101515156113b857600080fd5b81600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555092915050565b6060600d8054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156114a35780601f10611478576101008083540402835291602001916114a3565b820191906000526020600020905b81548152906001019060200180831161148657829003601f168201915b5050505050905090565b60006114ba3384846118a9565b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a36001905092915050565b6060600960008381526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156115d25780601f106115a7576101008083540402835291602001916115d2565b820191906000526020600020905b8154815290600101906020018083116115b557829003601f168201915b50505050509050919050565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b606060008060606000808614156116b3576040805190810160405280600181526020017f300000000000000000000000000000000000000000000000000000000000000081525094506117b3565b8593505b6000841415156116dd578280600101935050600a848115156116d557fe5b0493506116b7565b826040519080825280601f01601f1916602001820160405280156117105781602001602082028038833980820191505090505b5091506001830390505b6000861415156117af57600a8681151561173057fe5b066030017f01000000000000000000000000000000000000000000000000000000000000000282828060019003935081518110151561176b57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a868115156117a757fe5b04955061171a565b8194505b50505050919050565b606082826040516020018083805190602001908083835b6020831015156117f857805182526020820191506020810190506020830392506117d3565b6001836020036101000a03801982511681845116808217855250505050505090500182805190602001908083835b60208310151561184b5780518252602082019150602081019050602083039250611826565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052905092915050565b60008083831115151561189a57600080fd5b82840390508091505092915050565b600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481111515156118f757600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151561193357600080fd5b61198581600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461188890919063ffffffff16565b600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611a1a81600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611a6290919063ffffffff16565b600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505050565b6000808284019050838110151515611a7957600080fd5b8091505092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611ac457805160ff1916838001178555611af2565b82800160010185558215611af2579182015b82811115611af1578251825591602001919060010190611ad6565b5b509050611aff9190611b03565b5090565b611b2591905b80821115611b21576000816000905550600101611b09565b5090565b905600a165627a7a7230582094e8d7bf05f40984dd4a981b70804df1891e0f54d86b8baaf28cf1980bfd79190029`

// DeployPublicSlot deploys a new Ethereum contract, binding an instance of PublicSlot to it.
func DeployPublicSlot(auth *bind.TransactOpts, backend bind.ContractBackend, totalSupply *big.Int, tokenName string, tokenSymbol string) (common.Address, *types.Transaction, *PublicSlot, error) {
	parsed, err := abi.JSON(strings.NewReader(PublicSlotABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PublicSlotBin), backend, totalSupply, tokenName, tokenSymbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PublicSlot{PublicSlotCaller: PublicSlotCaller{contract: contract}, PublicSlotTransactor: PublicSlotTransactor{contract: contract}, PublicSlotFilterer: PublicSlotFilterer{contract: contract}}, nil
}

// PublicSlot is an auto generated Go binding around an Ethereum contract.
type PublicSlot struct {
	PublicSlotCaller     // Read-only binding to the contract
	PublicSlotTransactor // Write-only binding to the contract
	PublicSlotFilterer   // Log filterer for contract events
}

// PublicSlotCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicSlotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicSlotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicSlotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicSlotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicSlotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicSlotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicSlotSession struct {
	Contract     *PublicSlot       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PublicSlotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicSlotCallerSession struct {
	Contract *PublicSlotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PublicSlotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicSlotTransactorSession struct {
	Contract     *PublicSlotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PublicSlotRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicSlotRaw struct {
	Contract *PublicSlot // Generic contract binding to access the raw methods on
}

// PublicSlotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicSlotCallerRaw struct {
	Contract *PublicSlotCaller // Generic read-only contract binding to access the raw methods on
}

// PublicSlotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicSlotTransactorRaw struct {
	Contract *PublicSlotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicSlot creates a new instance of PublicSlot, bound to a specific deployed contract.
func NewPublicSlot(address common.Address, backend bind.ContractBackend) (*PublicSlot, error) {
	contract, err := bindPublicSlot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicSlot{PublicSlotCaller: PublicSlotCaller{contract: contract}, PublicSlotTransactor: PublicSlotTransactor{contract: contract}, PublicSlotFilterer: PublicSlotFilterer{contract: contract}}, nil
}

// NewPublicSlotCaller creates a new read-only instance of PublicSlot, bound to a specific deployed contract.
func NewPublicSlotCaller(address common.Address, caller bind.ContractCaller) (*PublicSlotCaller, error) {
	contract, err := bindPublicSlot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicSlotCaller{contract: contract}, nil
}

// NewPublicSlotTransactor creates a new write-only instance of PublicSlot, bound to a specific deployed contract.
func NewPublicSlotTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicSlotTransactor, error) {
	contract, err := bindPublicSlot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicSlotTransactor{contract: contract}, nil
}

// NewPublicSlotFilterer creates a new log filterer instance of PublicSlot, bound to a specific deployed contract.
func NewPublicSlotFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicSlotFilterer, error) {
	contract, err := bindPublicSlot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicSlotFilterer{contract: contract}, nil
}

// bindPublicSlot binds a generic wrapper to an already deployed contract.
func bindPublicSlot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PublicSlotABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicSlot *PublicSlotRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PublicSlot.Contract.PublicSlotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicSlot *PublicSlotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicSlot.Contract.PublicSlotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicSlot *PublicSlotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicSlot.Contract.PublicSlotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicSlot *PublicSlotCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PublicSlot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicSlot *PublicSlotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicSlot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicSlot *PublicSlotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicSlot.Contract.contract.Transact(opts, method, params...)
}

// NFTName is a free data retrieval call binding the contract method 0x94f4e679.
//
// Solidity: function NFTName() constant returns(string)
func (_PublicSlot *PublicSlotCaller) NFTName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PublicSlot.contract.Call(opts, out, "NFTName")
	return *ret0, err
}

// NFTName is a free data retrieval call binding the contract method 0x94f4e679.
//
// Solidity: function NFTName() constant returns(string)
func (_PublicSlot *PublicSlotSession) NFTName() (string, error) {
	return _PublicSlot.Contract.NFTName(&_PublicSlot.CallOpts)
}

// NFTName is a free data retrieval call binding the contract method 0x94f4e679.
//
// Solidity: function NFTName() constant returns(string)
func (_PublicSlot *PublicSlotCallerSession) NFTName() (string, error) {
	return _PublicSlot.Contract.NFTName(&_PublicSlot.CallOpts)
}

// NFTSymbol is a free data retrieval call binding the contract method 0xa42d20ae.
//
// Solidity: function NFTSymbol() constant returns(string)
func (_PublicSlot *PublicSlotCaller) NFTSymbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PublicSlot.contract.Call(opts, out, "NFTSymbol")
	return *ret0, err
}

// NFTSymbol is a free data retrieval call binding the contract method 0xa42d20ae.
//
// Solidity: function NFTSymbol() constant returns(string)
func (_PublicSlot *PublicSlotSession) NFTSymbol() (string, error) {
	return _PublicSlot.Contract.NFTSymbol(&_PublicSlot.CallOpts)
}

// NFTSymbol is a free data retrieval call binding the contract method 0xa42d20ae.
//
// Solidity: function NFTSymbol() constant returns(string)
func (_PublicSlot *PublicSlotCallerSession) NFTSymbol() (string, error) {
	return _PublicSlot.Contract.NFTSymbol(&_PublicSlot.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_PublicSlot *PublicSlotCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PublicSlot.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_PublicSlot *PublicSlotSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PublicSlot.Contract.Allowance(&_PublicSlot.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_PublicSlot *PublicSlotCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PublicSlot.Contract.Allowance(&_PublicSlot.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_PublicSlot *PublicSlotCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PublicSlot.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_PublicSlot *PublicSlotSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _PublicSlot.Contract.BalanceOf(&_PublicSlot.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_PublicSlot *PublicSlotCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _PublicSlot.Contract.BalanceOf(&_PublicSlot.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_PublicSlot *PublicSlotCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _PublicSlot.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_PublicSlot *PublicSlotSession) Decimals() (uint8, error) {
	return _PublicSlot.Contract.Decimals(&_PublicSlot.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_PublicSlot *PublicSlotCallerSession) Decimals() (uint8, error) {
	return _PublicSlot.Contract.Decimals(&_PublicSlot.CallOpts)
}

// ExchangeBase is a free data retrieval call binding the contract method 0x7e8b3f89.
//
// Solidity: function exchangeBase() constant returns(uint256)
func (_PublicSlot *PublicSlotCaller) ExchangeBase(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PublicSlot.contract.Call(opts, out, "exchangeBase")
	return *ret0, err
}

// ExchangeBase is a free data retrieval call binding the contract method 0x7e8b3f89.
//
// Solidity: function exchangeBase() constant returns(uint256)
func (_PublicSlot *PublicSlotSession) ExchangeBase() (*big.Int, error) {
	return _PublicSlot.Contract.ExchangeBase(&_PublicSlot.CallOpts)
}

// ExchangeBase is a free data retrieval call binding the contract method 0x7e8b3f89.
//
// Solidity: function exchangeBase() constant returns(uint256)
func (_PublicSlot *PublicSlotCallerSession) ExchangeBase() (*big.Int, error) {
	return _PublicSlot.Contract.ExchangeBase(&_PublicSlot.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() constant returns(uint256)
func (_PublicSlot *PublicSlotCaller) ExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PublicSlot.contract.Call(opts, out, "exchangeRate")
	return *ret0, err
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() constant returns(uint256)
func (_PublicSlot *PublicSlotSession) ExchangeRate() (*big.Int, error) {
	return _PublicSlot.Contract.ExchangeRate(&_PublicSlot.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() constant returns(uint256)
func (_PublicSlot *PublicSlotCallerSession) ExchangeRate() (*big.Int, error) {
	return _PublicSlot.Contract.ExchangeRate(&_PublicSlot.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_PublicSlot *PublicSlotCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PublicSlot.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_PublicSlot *PublicSlotSession) Name() (string, error) {
	return _PublicSlot.Contract.Name(&_PublicSlot.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_PublicSlot *PublicSlotCallerSession) Name() (string, error) {
	return _PublicSlot.Contract.Name(&_PublicSlot.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_PublicSlot *PublicSlotCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PublicSlot.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_PublicSlot *PublicSlotSession) Symbol() (string, error) {
	return _PublicSlot.Contract.Symbol(&_PublicSlot.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_PublicSlot *PublicSlotCallerSession) Symbol() (string, error) {
	return _PublicSlot.Contract.Symbol(&_PublicSlot.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) constant returns(string)
func (_PublicSlot *PublicSlotCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PublicSlot.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) constant returns(string)
func (_PublicSlot *PublicSlotSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _PublicSlot.Contract.TokenURI(&_PublicSlot.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) constant returns(string)
func (_PublicSlot *PublicSlotCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _PublicSlot.Contract.TokenURI(&_PublicSlot.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PublicSlot *PublicSlotCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PublicSlot.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PublicSlot *PublicSlotSession) TotalSupply() (*big.Int, error) {
	return _PublicSlot.Contract.TotalSupply(&_PublicSlot.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PublicSlot *PublicSlotCallerSession) TotalSupply() (*big.Int, error) {
	return _PublicSlot.Contract.TotalSupply(&_PublicSlot.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_PublicSlot *PublicSlotTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_PublicSlot *PublicSlotSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Approve(&_PublicSlot.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_PublicSlot *PublicSlotTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Approve(&_PublicSlot.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 amount) returns(bool)
func (_PublicSlot *PublicSlotTransactor) Burn(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "burn", user, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 amount) returns(bool)
func (_PublicSlot *PublicSlotSession) Burn(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Burn(&_PublicSlot.TransactOpts, user, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 amount) returns(bool)
func (_PublicSlot *PublicSlotTransactorSession) Burn(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Burn(&_PublicSlot.TransactOpts, user, amount)
}

// Consume is a paid mutator transaction binding the contract method 0x224b5c72.
//
// Solidity: function consume(address by, uint256 value) returns(bool)
func (_PublicSlot *PublicSlotTransactor) Consume(opts *bind.TransactOpts, by common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "consume", by, value)
}

// Consume is a paid mutator transaction binding the contract method 0x224b5c72.
//
// Solidity: function consume(address by, uint256 value) returns(bool)
func (_PublicSlot *PublicSlotSession) Consume(by common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Consume(&_PublicSlot.TransactOpts, by, value)
}

// Consume is a paid mutator transaction binding the contract method 0x224b5c72.
//
// Solidity: function consume(address by, uint256 value) returns(bool)
func (_PublicSlot *PublicSlotTransactorSession) Consume(by common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Consume(&_PublicSlot.TransactOpts, by, value)
}

// Create is a paid mutator transaction binding the contract method 0x0ecaea73.
//
// Solidity: function create(address _user, uint256 _tokenID) returns(bool)
func (_PublicSlot *PublicSlotTransactor) Create(opts *bind.TransactOpts, _user common.Address, _tokenID *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "create", _user, _tokenID)
}

// Create is a paid mutator transaction binding the contract method 0x0ecaea73.
//
// Solidity: function create(address _user, uint256 _tokenID) returns(bool)
func (_PublicSlot *PublicSlotSession) Create(_user common.Address, _tokenID *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Create(&_PublicSlot.TransactOpts, _user, _tokenID)
}

// Create is a paid mutator transaction binding the contract method 0x0ecaea73.
//
// Solidity: function create(address _user, uint256 _tokenID) returns(bool)
func (_PublicSlot *PublicSlotTransactorSession) Create(_user common.Address, _tokenID *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Create(&_PublicSlot.TransactOpts, _user, _tokenID)
}

// ExchangeForEther is a paid mutator transaction binding the contract method 0x5f1e6941.
//
// Solidity: function exchangeForEther(address user, uint256 amount) returns(bool)
func (_PublicSlot *PublicSlotTransactor) ExchangeForEther(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "exchangeForEther", user, amount)
}

// ExchangeForEther is a paid mutator transaction binding the contract method 0x5f1e6941.
//
// Solidity: function exchangeForEther(address user, uint256 amount) returns(bool)
func (_PublicSlot *PublicSlotSession) ExchangeForEther(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.ExchangeForEther(&_PublicSlot.TransactOpts, user, amount)
}

// ExchangeForEther is a paid mutator transaction binding the contract method 0x5f1e6941.
//
// Solidity: function exchangeForEther(address user, uint256 amount) returns(bool)
func (_PublicSlot *PublicSlotTransactorSession) ExchangeForEther(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.ExchangeForEther(&_PublicSlot.TransactOpts, user, amount)
}

// ExchangeForToken is a paid mutator transaction binding the contract method 0x7317d44a.
//
// Solidity: function exchangeForToken(address user) returns(bool)
func (_PublicSlot *PublicSlotTransactor) ExchangeForToken(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "exchangeForToken", user)
}

// ExchangeForToken is a paid mutator transaction binding the contract method 0x7317d44a.
//
// Solidity: function exchangeForToken(address user) returns(bool)
func (_PublicSlot *PublicSlotSession) ExchangeForToken(user common.Address) (*types.Transaction, error) {
	return _PublicSlot.Contract.ExchangeForToken(&_PublicSlot.TransactOpts, user)
}

// ExchangeForToken is a paid mutator transaction binding the contract method 0x7317d44a.
//
// Solidity: function exchangeForToken(address user) returns(bool)
func (_PublicSlot *PublicSlotTransactorSession) ExchangeForToken(user common.Address) (*types.Transaction, error) {
	return _PublicSlot.Contract.ExchangeForToken(&_PublicSlot.TransactOpts, user)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address user, uint256 amount) returns(bool)
func (_PublicSlot *PublicSlotTransactor) Mint(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "mint", user, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address user, uint256 amount) returns(bool)
func (_PublicSlot *PublicSlotSession) Mint(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Mint(&_PublicSlot.TransactOpts, user, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address user, uint256 amount) returns(bool)
func (_PublicSlot *PublicSlotTransactorSession) Mint(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Mint(&_PublicSlot.TransactOpts, user, amount)
}

// Reward is a paid mutator transaction binding the contract method 0x21670f22.
//
// Solidity: function reward(address to, uint256 value) returns(bool)
func (_PublicSlot *PublicSlotTransactor) Reward(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "reward", to, value)
}

// Reward is a paid mutator transaction binding the contract method 0x21670f22.
//
// Solidity: function reward(address to, uint256 value) returns(bool)
func (_PublicSlot *PublicSlotSession) Reward(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Reward(&_PublicSlot.TransactOpts, to, value)
}

// Reward is a paid mutator transaction binding the contract method 0x21670f22.
//
// Solidity: function reward(address to, uint256 value) returns(bool)
func (_PublicSlot *PublicSlotTransactorSession) Reward(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Reward(&_PublicSlot.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_PublicSlot *PublicSlotTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_PublicSlot *PublicSlotSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Transfer(&_PublicSlot.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_PublicSlot *PublicSlotTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Transfer(&_PublicSlot.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_PublicSlot *PublicSlotTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_PublicSlot *PublicSlotSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.TransferFrom(&_PublicSlot.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_PublicSlot *PublicSlotTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.TransferFrom(&_PublicSlot.TransactOpts, from, to, value)
}

// PublicSlotApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PublicSlot contract.
type PublicSlotApprovalIterator struct {
	Event *PublicSlotApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PublicSlotApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicSlotApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PublicSlotApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PublicSlotApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicSlotApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicSlotApproval represents a Approval event raised by the PublicSlot contract.
type PublicSlotApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_PublicSlot *PublicSlotFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*PublicSlotApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _PublicSlot.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &PublicSlotApprovalIterator{contract: _PublicSlot.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_PublicSlot *PublicSlotFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PublicSlotApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _PublicSlot.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicSlotApproval)
				if err := _PublicSlot.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PublicSlotBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the PublicSlot contract.
type PublicSlotBurnIterator struct {
	Event *PublicSlotBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PublicSlotBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicSlotBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PublicSlotBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PublicSlotBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicSlotBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicSlotBurn represents a Burn event raised by the PublicSlot contract.
type PublicSlotBurn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address user, uint256 amount)
func (_PublicSlot *PublicSlotFilterer) FilterBurn(opts *bind.FilterOpts) (*PublicSlotBurnIterator, error) {

	logs, sub, err := _PublicSlot.contract.FilterLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return &PublicSlotBurnIterator{contract: _PublicSlot.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address user, uint256 amount)
func (_PublicSlot *PublicSlotFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *PublicSlotBurn) (event.Subscription, error) {

	logs, sub, err := _PublicSlot.contract.WatchLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicSlotBurn)
				if err := _PublicSlot.contract.UnpackLog(event, "Burn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PublicSlotMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the PublicSlot contract.
type PublicSlotMintIterator struct {
	Event *PublicSlotMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PublicSlotMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicSlotMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PublicSlotMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PublicSlotMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicSlotMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicSlotMint represents a Mint event raised by the PublicSlot contract.
type PublicSlotMint struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address user, uint256 amount)
func (_PublicSlot *PublicSlotFilterer) FilterMint(opts *bind.FilterOpts) (*PublicSlotMintIterator, error) {

	logs, sub, err := _PublicSlot.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &PublicSlotMintIterator{contract: _PublicSlot.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address user, uint256 amount)
func (_PublicSlot *PublicSlotFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *PublicSlotMint) (event.Subscription, error) {

	logs, sub, err := _PublicSlot.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicSlotMint)
				if err := _PublicSlot.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PublicSlotTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PublicSlot contract.
type PublicSlotTransferIterator struct {
	Event *PublicSlotTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PublicSlotTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicSlotTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PublicSlotTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PublicSlotTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicSlotTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicSlotTransfer represents a Transfer event raised by the PublicSlot contract.
type PublicSlotTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_PublicSlot *PublicSlotFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*PublicSlotTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _PublicSlot.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &PublicSlotTransferIterator{contract: _PublicSlot.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_PublicSlot *PublicSlotFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PublicSlotTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _PublicSlot.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicSlotTransfer)
				if err := _PublicSlot.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

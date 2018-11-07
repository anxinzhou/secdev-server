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
const PublicSlotABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"reward\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"by\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"consume\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchangeRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"exchangeForEther\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"exchangeForToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchangeBase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// PublicSlotBin is the compiled bytecode used for deploying new contracts.
const PublicSlotBin = `60806040526101f46007556001600855604051620014b1380380620014b18339810180604052810190808051906020019092919080518201929190602001805182019291905050508282826012600260006101000a81548160ff021916908360ff160217905550600260009054906101000a900460ff1660ff16600a0a8302600081905550600054600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508160019080519060200190620000e292919062000149565b508060039080519060200190620000fb92919062000149565b5050505033600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050620001f8565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200018c57805160ff1916838001178555620001bd565b82800160010185558215620001bd579182015b82811115620001bc5782518255916020019190600101906200019f565b5b509050620001cc9190620001d0565b5090565b620001f591905b80821115620001f1576000816000905550600101620001d7565b5090565b90565b6112a980620002086000396000f3006080604052600436106100f1576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde03146100f6578063095ea7b31461018657806318160ddd146101eb57806321670f2214610216578063224b5c721461027b57806323b872dd146102e0578063313ce567146103655780633ba0b9a91461039657806340c10f19146103c15780635f1e69411461042657806370a082311461048b5780637317d44a146104e25780637e8b3f891461053057806395d89b411461055b5780639dc29fac146105eb578063a9059cbb14610650578063dd62ed3e146106b5575b600080fd5b34801561010257600080fd5b5061010b61072c565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561014b578082015181840152602081019050610130565b50505050905090810190601f1680156101785780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561019257600080fd5b506101d1600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506107ca565b604051808215151515815260200191505060405180910390f35b3480156101f757600080fd5b506102006108f7565b6040518082815260200191505060405180910390f35b34801561022257600080fd5b50610261600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610900565b604051808215151515815260200191505060405180910390f35b34801561028757600080fd5b506102c6600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610959565b604051808215151515815260200191505060405180910390f35b3480156102ec57600080fd5b5061034b600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a00565b604051808215151515815260200191505060405180910390f35b34801561037157600080fd5b5061037a610c17565b604051808260ff1660ff16815260200191505060405180910390f35b3480156103a257600080fd5b506103ab610c2a565b6040518082815260200191505060405180910390f35b3480156103cd57600080fd5b5061040c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610c30565b604051808215151515815260200191505060405180910390f35b34801561043257600080fd5b50610471600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610c85565b604051808215151515815260200191505060405180910390f35b34801561049757600080fd5b506104cc600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d85565b6040518082815260200191505060405180910390f35b610516600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610dce565b604051808215151515815260200191505060405180910390f35b34801561053c57600080fd5b50610545610e38565b6040518082815260200191505060405180910390f35b34801561056757600080fd5b50610570610e3e565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156105b0578082015181840152602081019050610595565b50505050905090810190601f1680156105dd5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156105f757600080fd5b50610636600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610edc565b604051808215151515815260200191505060405180910390f35b34801561065c57600080fd5b5061069b600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610f7f565b604051808215151515815260200191505060405180910390f35b3480156106c157600080fd5b50610716600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610ffb565b6040518082815260200191505060405180910390f35b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107c25780601f10610797576101008083540402835291602001916107c2565b820191906000526020600020905b8154815290600101906020018083116107a557829003601f168201915b505050505081565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415151561080757600080fd5b81600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b60008054905090565b600081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055506001905092915050565b600081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101515156109a957600080fd5b81600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055506001905092915050565b6000600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548211151515610a8d57600080fd5b610b1c82600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461108290919063ffffffff16565b600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610ba78484846110a3565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3600190509392505050565b600260009054906101000a900460ff1681565b60075481565b600081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555092915050565b6000806007546008548402811515610c9957fe5b04905082600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410151515610cea57600080fd5b82600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055508373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610d7d573d6000803e3d6000fd5b505092915050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000806008546007543402811515610de257fe5b04905080600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555050919050565b60085481565b60038054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ed45780601f10610ea957610100808354040283529160200191610ed4565b820191906000526020600020905b815481529060010190602001808311610eb757829003601f168201915b505050505081565b600081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410151515610f2c57600080fd5b81600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555092915050565b6000610f8c3384846110a3565b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a36001905092915050565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b60008083831115151561109457600080fd5b82840390508091505092915050565b600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481111515156110f157600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151561112d57600080fd5b61117f81600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461108290919063ffffffff16565b600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061121481600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461125c90919063ffffffff16565b600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505050565b600080828401905083811015151561127357600080fd5b80915050929150505600a165627a7a72305820288ef6835f976f76b37ca302b3ba4eb4b1be317227410ceaaec31e49d84844da0029`

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

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
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
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_PublicSlot *PublicSlotSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PublicSlot.Contract.Allowance(&_PublicSlot.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_PublicSlot *PublicSlotCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PublicSlot.Contract.Allowance(&_PublicSlot.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
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
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_PublicSlot *PublicSlotSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _PublicSlot.Contract.BalanceOf(&_PublicSlot.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
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
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_PublicSlot *PublicSlotTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_PublicSlot *PublicSlotSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Approve(&_PublicSlot.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_PublicSlot *PublicSlotTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Approve(&_PublicSlot.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(user address, amount uint256) returns(bool)
func (_PublicSlot *PublicSlotTransactor) Burn(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "burn", user, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(user address, amount uint256) returns(bool)
func (_PublicSlot *PublicSlotSession) Burn(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Burn(&_PublicSlot.TransactOpts, user, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(user address, amount uint256) returns(bool)
func (_PublicSlot *PublicSlotTransactorSession) Burn(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Burn(&_PublicSlot.TransactOpts, user, amount)
}

// Consume is a paid mutator transaction binding the contract method 0x224b5c72.
//
// Solidity: function consume(by address, value uint256) returns(bool)
func (_PublicSlot *PublicSlotTransactor) Consume(opts *bind.TransactOpts, by common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "consume", by, value)
}

// Consume is a paid mutator transaction binding the contract method 0x224b5c72.
//
// Solidity: function consume(by address, value uint256) returns(bool)
func (_PublicSlot *PublicSlotSession) Consume(by common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Consume(&_PublicSlot.TransactOpts, by, value)
}

// Consume is a paid mutator transaction binding the contract method 0x224b5c72.
//
// Solidity: function consume(by address, value uint256) returns(bool)
func (_PublicSlot *PublicSlotTransactorSession) Consume(by common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Consume(&_PublicSlot.TransactOpts, by, value)
}

// ExchangeForEther is a paid mutator transaction binding the contract method 0x5f1e6941.
//
// Solidity: function exchangeForEther(user address, amount uint256) returns(bool)
func (_PublicSlot *PublicSlotTransactor) ExchangeForEther(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "exchangeForEther", user, amount)
}

// ExchangeForEther is a paid mutator transaction binding the contract method 0x5f1e6941.
//
// Solidity: function exchangeForEther(user address, amount uint256) returns(bool)
func (_PublicSlot *PublicSlotSession) ExchangeForEther(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.ExchangeForEther(&_PublicSlot.TransactOpts, user, amount)
}

// ExchangeForEther is a paid mutator transaction binding the contract method 0x5f1e6941.
//
// Solidity: function exchangeForEther(user address, amount uint256) returns(bool)
func (_PublicSlot *PublicSlotTransactorSession) ExchangeForEther(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.ExchangeForEther(&_PublicSlot.TransactOpts, user, amount)
}

// ExchangeForToken is a paid mutator transaction binding the contract method 0x7317d44a.
//
// Solidity: function exchangeForToken(user address) returns(bool)
func (_PublicSlot *PublicSlotTransactor) ExchangeForToken(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "exchangeForToken", user)
}

// ExchangeForToken is a paid mutator transaction binding the contract method 0x7317d44a.
//
// Solidity: function exchangeForToken(user address) returns(bool)
func (_PublicSlot *PublicSlotSession) ExchangeForToken(user common.Address) (*types.Transaction, error) {
	return _PublicSlot.Contract.ExchangeForToken(&_PublicSlot.TransactOpts, user)
}

// ExchangeForToken is a paid mutator transaction binding the contract method 0x7317d44a.
//
// Solidity: function exchangeForToken(user address) returns(bool)
func (_PublicSlot *PublicSlotTransactorSession) ExchangeForToken(user common.Address) (*types.Transaction, error) {
	return _PublicSlot.Contract.ExchangeForToken(&_PublicSlot.TransactOpts, user)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(user address, amount uint256) returns(bool)
func (_PublicSlot *PublicSlotTransactor) Mint(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "mint", user, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(user address, amount uint256) returns(bool)
func (_PublicSlot *PublicSlotSession) Mint(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Mint(&_PublicSlot.TransactOpts, user, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(user address, amount uint256) returns(bool)
func (_PublicSlot *PublicSlotTransactorSession) Mint(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Mint(&_PublicSlot.TransactOpts, user, amount)
}

// Reward is a paid mutator transaction binding the contract method 0x21670f22.
//
// Solidity: function reward(to address, value uint256) returns(bool)
func (_PublicSlot *PublicSlotTransactor) Reward(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "reward", to, value)
}

// Reward is a paid mutator transaction binding the contract method 0x21670f22.
//
// Solidity: function reward(to address, value uint256) returns(bool)
func (_PublicSlot *PublicSlotSession) Reward(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Reward(&_PublicSlot.TransactOpts, to, value)
}

// Reward is a paid mutator transaction binding the contract method 0x21670f22.
//
// Solidity: function reward(to address, value uint256) returns(bool)
func (_PublicSlot *PublicSlotTransactorSession) Reward(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Reward(&_PublicSlot.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_PublicSlot *PublicSlotTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_PublicSlot *PublicSlotSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Transfer(&_PublicSlot.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_PublicSlot *PublicSlotTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Transfer(&_PublicSlot.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_PublicSlot *PublicSlotTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_PublicSlot *PublicSlotSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.TransferFrom(&_PublicSlot.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
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
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
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
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
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
// Solidity: e Burn(user address, amount uint256)
func (_PublicSlot *PublicSlotFilterer) FilterBurn(opts *bind.FilterOpts) (*PublicSlotBurnIterator, error) {

	logs, sub, err := _PublicSlot.contract.FilterLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return &PublicSlotBurnIterator{contract: _PublicSlot.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(user address, amount uint256)
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
// Solidity: e Mint(user address, amount uint256)
func (_PublicSlot *PublicSlotFilterer) FilterMint(opts *bind.FilterOpts) (*PublicSlotMintIterator, error) {

	logs, sub, err := _PublicSlot.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &PublicSlotMintIterator{contract: _PublicSlot.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(user address, amount uint256)
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
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
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
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
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

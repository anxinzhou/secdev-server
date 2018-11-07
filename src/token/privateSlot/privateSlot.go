// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package privateSlot

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

// PrivateSlotABI is the input ABI used to generate the binding from.
const PrivateSlotABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"reward\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"by\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"consume\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchangeRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"exchangeForEther\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"exchangeForToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchangeBase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// PrivateSlotBin is the compiled bytecode used for deploying new contracts.
const PrivateSlotBin = `60806040526101f46007556001600855604051620014b1380380620014b18339810180604052810190808051906020019092919080518201929190602001805182019291905050508282826012600260006101000a81548160ff021916908360ff160217905550600260009054906101000a900460ff1660ff16600a0a8302600081905550600054600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508160019080519060200190620000e292919062000149565b508060039080519060200190620000fb92919062000149565b5050505033600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050620001f8565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200018c57805160ff1916838001178555620001bd565b82800160010185558215620001bd579182015b82811115620001bc5782518255916020019190600101906200019f565b5b509050620001cc9190620001d0565b5090565b620001f591905b80821115620001f1576000816000905550600101620001d7565b5090565b90565b6112a980620002086000396000f3006080604052600436106100f1576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde03146100f6578063095ea7b31461018657806318160ddd146101eb57806321670f2214610216578063224b5c721461027b57806323b872dd146102e0578063313ce567146103655780633ba0b9a91461039657806340c10f19146103c15780635f1e69411461042657806370a082311461048b5780637317d44a146104e25780637e8b3f891461053057806395d89b411461055b5780639dc29fac146105eb578063a9059cbb14610650578063dd62ed3e146106b5575b600080fd5b34801561010257600080fd5b5061010b61072c565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561014b578082015181840152602081019050610130565b50505050905090810190601f1680156101785780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561019257600080fd5b506101d1600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506107ca565b604051808215151515815260200191505060405180910390f35b3480156101f757600080fd5b506102006108f7565b6040518082815260200191505060405180910390f35b34801561022257600080fd5b50610261600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610900565b604051808215151515815260200191505060405180910390f35b34801561028757600080fd5b506102c6600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610959565b604051808215151515815260200191505060405180910390f35b3480156102ec57600080fd5b5061034b600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a00565b604051808215151515815260200191505060405180910390f35b34801561037157600080fd5b5061037a610c17565b604051808260ff1660ff16815260200191505060405180910390f35b3480156103a257600080fd5b506103ab610c2a565b6040518082815260200191505060405180910390f35b3480156103cd57600080fd5b5061040c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610c30565b604051808215151515815260200191505060405180910390f35b34801561043257600080fd5b50610471600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610c85565b604051808215151515815260200191505060405180910390f35b34801561049757600080fd5b506104cc600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d85565b6040518082815260200191505060405180910390f35b610516600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610dce565b604051808215151515815260200191505060405180910390f35b34801561053c57600080fd5b50610545610e38565b6040518082815260200191505060405180910390f35b34801561056757600080fd5b50610570610e3e565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156105b0578082015181840152602081019050610595565b50505050905090810190601f1680156105dd5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156105f757600080fd5b50610636600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610edc565b604051808215151515815260200191505060405180910390f35b34801561065c57600080fd5b5061069b600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610f7f565b604051808215151515815260200191505060405180910390f35b3480156106c157600080fd5b50610716600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610ffb565b6040518082815260200191505060405180910390f35b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107c25780601f10610797576101008083540402835291602001916107c2565b820191906000526020600020905b8154815290600101906020018083116107a557829003601f168201915b505050505081565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415151561080757600080fd5b81600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b60008054905090565b600081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055506001905092915050565b600081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101515156109a957600080fd5b81600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055506001905092915050565b6000600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548211151515610a8d57600080fd5b610b1c82600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461108290919063ffffffff16565b600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610ba78484846110a3565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3600190509392505050565b600260009054906101000a900460ff1681565b60075481565b600081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555092915050565b6000806007546008548402811515610c9957fe5b04905082600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410151515610cea57600080fd5b82600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055508373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610d7d573d6000803e3d6000fd5b505092915050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000806008546007543402811515610de257fe5b04905080600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555050919050565b60085481565b60038054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ed45780601f10610ea957610100808354040283529160200191610ed4565b820191906000526020600020905b815481529060010190602001808311610eb757829003601f168201915b505050505081565b600081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410151515610f2c57600080fd5b81600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555092915050565b6000610f8c3384846110a3565b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a36001905092915050565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b60008083831115151561109457600080fd5b82840390508091505092915050565b600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481111515156110f157600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151561112d57600080fd5b61117f81600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461108290919063ffffffff16565b600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061121481600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461125c90919063ffffffff16565b600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505050565b600080828401905083811015151561127357600080fd5b80915050929150505600a165627a7a72305820ea1a5fe58ff78082c052032258fc4485d89394c7783233365ac75b04d71ca21f0029`

// DeployPrivateSlot deploys a new Ethereum contract, binding an instance of PrivateSlot to it.
func DeployPrivateSlot(auth *bind.TransactOpts, backend bind.ContractBackend, totalSupply *big.Int, tokenName string, tokenSymbol string) (common.Address, *types.Transaction, *PrivateSlot, error) {
	parsed, err := abi.JSON(strings.NewReader(PrivateSlotABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PrivateSlotBin), backend, totalSupply, tokenName, tokenSymbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PrivateSlot{PrivateSlotCaller: PrivateSlotCaller{contract: contract}, PrivateSlotTransactor: PrivateSlotTransactor{contract: contract}, PrivateSlotFilterer: PrivateSlotFilterer{contract: contract}}, nil
}

// PrivateSlot is an auto generated Go binding around an Ethereum contract.
type PrivateSlot struct {
	PrivateSlotCaller     // Read-only binding to the contract
	PrivateSlotTransactor // Write-only binding to the contract
	PrivateSlotFilterer   // Log filterer for contract events
}

// PrivateSlotCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrivateSlotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivateSlotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrivateSlotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivateSlotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrivateSlotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivateSlotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrivateSlotSession struct {
	Contract     *PrivateSlot      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PrivateSlotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrivateSlotCallerSession struct {
	Contract *PrivateSlotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PrivateSlotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrivateSlotTransactorSession struct {
	Contract     *PrivateSlotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PrivateSlotRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrivateSlotRaw struct {
	Contract *PrivateSlot // Generic contract binding to access the raw methods on
}

// PrivateSlotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrivateSlotCallerRaw struct {
	Contract *PrivateSlotCaller // Generic read-only contract binding to access the raw methods on
}

// PrivateSlotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrivateSlotTransactorRaw struct {
	Contract *PrivateSlotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrivateSlot creates a new instance of PrivateSlot, bound to a specific deployed contract.
func NewPrivateSlot(address common.Address, backend bind.ContractBackend) (*PrivateSlot, error) {
	contract, err := bindPrivateSlot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PrivateSlot{PrivateSlotCaller: PrivateSlotCaller{contract: contract}, PrivateSlotTransactor: PrivateSlotTransactor{contract: contract}, PrivateSlotFilterer: PrivateSlotFilterer{contract: contract}}, nil
}

// NewPrivateSlotCaller creates a new read-only instance of PrivateSlot, bound to a specific deployed contract.
func NewPrivateSlotCaller(address common.Address, caller bind.ContractCaller) (*PrivateSlotCaller, error) {
	contract, err := bindPrivateSlot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrivateSlotCaller{contract: contract}, nil
}

// NewPrivateSlotTransactor creates a new write-only instance of PrivateSlot, bound to a specific deployed contract.
func NewPrivateSlotTransactor(address common.Address, transactor bind.ContractTransactor) (*PrivateSlotTransactor, error) {
	contract, err := bindPrivateSlot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrivateSlotTransactor{contract: contract}, nil
}

// NewPrivateSlotFilterer creates a new log filterer instance of PrivateSlot, bound to a specific deployed contract.
func NewPrivateSlotFilterer(address common.Address, filterer bind.ContractFilterer) (*PrivateSlotFilterer, error) {
	contract, err := bindPrivateSlot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrivateSlotFilterer{contract: contract}, nil
}

// bindPrivateSlot binds a generic wrapper to an already deployed contract.
func bindPrivateSlot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PrivateSlotABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrivateSlot *PrivateSlotRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PrivateSlot.Contract.PrivateSlotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrivateSlot *PrivateSlotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivateSlot.Contract.PrivateSlotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrivateSlot *PrivateSlotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrivateSlot.Contract.PrivateSlotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrivateSlot *PrivateSlotCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PrivateSlot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrivateSlot *PrivateSlotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivateSlot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrivateSlot *PrivateSlotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrivateSlot.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_PrivateSlot *PrivateSlotCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivateSlot.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_PrivateSlot *PrivateSlotSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PrivateSlot.Contract.Allowance(&_PrivateSlot.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_PrivateSlot *PrivateSlotCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PrivateSlot.Contract.Allowance(&_PrivateSlot.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_PrivateSlot *PrivateSlotCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivateSlot.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_PrivateSlot *PrivateSlotSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _PrivateSlot.Contract.BalanceOf(&_PrivateSlot.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_PrivateSlot *PrivateSlotCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _PrivateSlot.Contract.BalanceOf(&_PrivateSlot.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_PrivateSlot *PrivateSlotCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _PrivateSlot.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_PrivateSlot *PrivateSlotSession) Decimals() (uint8, error) {
	return _PrivateSlot.Contract.Decimals(&_PrivateSlot.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_PrivateSlot *PrivateSlotCallerSession) Decimals() (uint8, error) {
	return _PrivateSlot.Contract.Decimals(&_PrivateSlot.CallOpts)
}

// ExchangeBase is a free data retrieval call binding the contract method 0x7e8b3f89.
//
// Solidity: function exchangeBase() constant returns(uint256)
func (_PrivateSlot *PrivateSlotCaller) ExchangeBase(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivateSlot.contract.Call(opts, out, "exchangeBase")
	return *ret0, err
}

// ExchangeBase is a free data retrieval call binding the contract method 0x7e8b3f89.
//
// Solidity: function exchangeBase() constant returns(uint256)
func (_PrivateSlot *PrivateSlotSession) ExchangeBase() (*big.Int, error) {
	return _PrivateSlot.Contract.ExchangeBase(&_PrivateSlot.CallOpts)
}

// ExchangeBase is a free data retrieval call binding the contract method 0x7e8b3f89.
//
// Solidity: function exchangeBase() constant returns(uint256)
func (_PrivateSlot *PrivateSlotCallerSession) ExchangeBase() (*big.Int, error) {
	return _PrivateSlot.Contract.ExchangeBase(&_PrivateSlot.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() constant returns(uint256)
func (_PrivateSlot *PrivateSlotCaller) ExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivateSlot.contract.Call(opts, out, "exchangeRate")
	return *ret0, err
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() constant returns(uint256)
func (_PrivateSlot *PrivateSlotSession) ExchangeRate() (*big.Int, error) {
	return _PrivateSlot.Contract.ExchangeRate(&_PrivateSlot.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() constant returns(uint256)
func (_PrivateSlot *PrivateSlotCallerSession) ExchangeRate() (*big.Int, error) {
	return _PrivateSlot.Contract.ExchangeRate(&_PrivateSlot.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_PrivateSlot *PrivateSlotCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PrivateSlot.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_PrivateSlot *PrivateSlotSession) Name() (string, error) {
	return _PrivateSlot.Contract.Name(&_PrivateSlot.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_PrivateSlot *PrivateSlotCallerSession) Name() (string, error) {
	return _PrivateSlot.Contract.Name(&_PrivateSlot.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_PrivateSlot *PrivateSlotCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PrivateSlot.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_PrivateSlot *PrivateSlotSession) Symbol() (string, error) {
	return _PrivateSlot.Contract.Symbol(&_PrivateSlot.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_PrivateSlot *PrivateSlotCallerSession) Symbol() (string, error) {
	return _PrivateSlot.Contract.Symbol(&_PrivateSlot.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PrivateSlot *PrivateSlotCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivateSlot.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PrivateSlot *PrivateSlotSession) TotalSupply() (*big.Int, error) {
	return _PrivateSlot.Contract.TotalSupply(&_PrivateSlot.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PrivateSlot *PrivateSlotCallerSession) TotalSupply() (*big.Int, error) {
	return _PrivateSlot.Contract.TotalSupply(&_PrivateSlot.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_PrivateSlot *PrivateSlotTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_PrivateSlot *PrivateSlotSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Approve(&_PrivateSlot.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_PrivateSlot *PrivateSlotTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Approve(&_PrivateSlot.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(user address, amount uint256) returns(bool)
func (_PrivateSlot *PrivateSlotTransactor) Burn(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "burn", user, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(user address, amount uint256) returns(bool)
func (_PrivateSlot *PrivateSlotSession) Burn(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Burn(&_PrivateSlot.TransactOpts, user, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(user address, amount uint256) returns(bool)
func (_PrivateSlot *PrivateSlotTransactorSession) Burn(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Burn(&_PrivateSlot.TransactOpts, user, amount)
}

// Consume is a paid mutator transaction binding the contract method 0x224b5c72.
//
// Solidity: function consume(by address, value uint256) returns(bool)
func (_PrivateSlot *PrivateSlotTransactor) Consume(opts *bind.TransactOpts, by common.Address, value *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "consume", by, value)
}

// Consume is a paid mutator transaction binding the contract method 0x224b5c72.
//
// Solidity: function consume(by address, value uint256) returns(bool)
func (_PrivateSlot *PrivateSlotSession) Consume(by common.Address, value *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Consume(&_PrivateSlot.TransactOpts, by, value)
}

// Consume is a paid mutator transaction binding the contract method 0x224b5c72.
//
// Solidity: function consume(by address, value uint256) returns(bool)
func (_PrivateSlot *PrivateSlotTransactorSession) Consume(by common.Address, value *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Consume(&_PrivateSlot.TransactOpts, by, value)
}

// ExchangeForEther is a paid mutator transaction binding the contract method 0x5f1e6941.
//
// Solidity: function exchangeForEther(user address, amount uint256) returns(bool)
func (_PrivateSlot *PrivateSlotTransactor) ExchangeForEther(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "exchangeForEther", user, amount)
}

// ExchangeForEther is a paid mutator transaction binding the contract method 0x5f1e6941.
//
// Solidity: function exchangeForEther(user address, amount uint256) returns(bool)
func (_PrivateSlot *PrivateSlotSession) ExchangeForEther(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.ExchangeForEther(&_PrivateSlot.TransactOpts, user, amount)
}

// ExchangeForEther is a paid mutator transaction binding the contract method 0x5f1e6941.
//
// Solidity: function exchangeForEther(user address, amount uint256) returns(bool)
func (_PrivateSlot *PrivateSlotTransactorSession) ExchangeForEther(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.ExchangeForEther(&_PrivateSlot.TransactOpts, user, amount)
}

// ExchangeForToken is a paid mutator transaction binding the contract method 0x7317d44a.
//
// Solidity: function exchangeForToken(user address) returns(bool)
func (_PrivateSlot *PrivateSlotTransactor) ExchangeForToken(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "exchangeForToken", user)
}

// ExchangeForToken is a paid mutator transaction binding the contract method 0x7317d44a.
//
// Solidity: function exchangeForToken(user address) returns(bool)
func (_PrivateSlot *PrivateSlotSession) ExchangeForToken(user common.Address) (*types.Transaction, error) {
	return _PrivateSlot.Contract.ExchangeForToken(&_PrivateSlot.TransactOpts, user)
}

// ExchangeForToken is a paid mutator transaction binding the contract method 0x7317d44a.
//
// Solidity: function exchangeForToken(user address) returns(bool)
func (_PrivateSlot *PrivateSlotTransactorSession) ExchangeForToken(user common.Address) (*types.Transaction, error) {
	return _PrivateSlot.Contract.ExchangeForToken(&_PrivateSlot.TransactOpts, user)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(user address, amount uint256) returns(bool)
func (_PrivateSlot *PrivateSlotTransactor) Mint(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "mint", user, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(user address, amount uint256) returns(bool)
func (_PrivateSlot *PrivateSlotSession) Mint(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Mint(&_PrivateSlot.TransactOpts, user, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(user address, amount uint256) returns(bool)
func (_PrivateSlot *PrivateSlotTransactorSession) Mint(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Mint(&_PrivateSlot.TransactOpts, user, amount)
}

// Reward is a paid mutator transaction binding the contract method 0x21670f22.
//
// Solidity: function reward(to address, value uint256) returns(bool)
func (_PrivateSlot *PrivateSlotTransactor) Reward(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "reward", to, value)
}

// Reward is a paid mutator transaction binding the contract method 0x21670f22.
//
// Solidity: function reward(to address, value uint256) returns(bool)
func (_PrivateSlot *PrivateSlotSession) Reward(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Reward(&_PrivateSlot.TransactOpts, to, value)
}

// Reward is a paid mutator transaction binding the contract method 0x21670f22.
//
// Solidity: function reward(to address, value uint256) returns(bool)
func (_PrivateSlot *PrivateSlotTransactorSession) Reward(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Reward(&_PrivateSlot.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_PrivateSlot *PrivateSlotTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_PrivateSlot *PrivateSlotSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Transfer(&_PrivateSlot.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_PrivateSlot *PrivateSlotTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Transfer(&_PrivateSlot.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_PrivateSlot *PrivateSlotTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_PrivateSlot *PrivateSlotSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.TransferFrom(&_PrivateSlot.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_PrivateSlot *PrivateSlotTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.TransferFrom(&_PrivateSlot.TransactOpts, from, to, value)
}

// PrivateSlotApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PrivateSlot contract.
type PrivateSlotApprovalIterator struct {
	Event *PrivateSlotApproval // Event containing the contract specifics and raw log

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
func (it *PrivateSlotApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateSlotApproval)
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
		it.Event = new(PrivateSlotApproval)
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
func (it *PrivateSlotApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateSlotApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateSlotApproval represents a Approval event raised by the PrivateSlot contract.
type PrivateSlotApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_PrivateSlot *PrivateSlotFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*PrivateSlotApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _PrivateSlot.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &PrivateSlotApprovalIterator{contract: _PrivateSlot.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_PrivateSlot *PrivateSlotFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PrivateSlotApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _PrivateSlot.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateSlotApproval)
				if err := _PrivateSlot.contract.UnpackLog(event, "Approval", log); err != nil {
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

// PrivateSlotBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the PrivateSlot contract.
type PrivateSlotBurnIterator struct {
	Event *PrivateSlotBurn // Event containing the contract specifics and raw log

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
func (it *PrivateSlotBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateSlotBurn)
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
		it.Event = new(PrivateSlotBurn)
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
func (it *PrivateSlotBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateSlotBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateSlotBurn represents a Burn event raised by the PrivateSlot contract.
type PrivateSlotBurn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(user address, amount uint256)
func (_PrivateSlot *PrivateSlotFilterer) FilterBurn(opts *bind.FilterOpts) (*PrivateSlotBurnIterator, error) {

	logs, sub, err := _PrivateSlot.contract.FilterLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return &PrivateSlotBurnIterator{contract: _PrivateSlot.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(user address, amount uint256)
func (_PrivateSlot *PrivateSlotFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *PrivateSlotBurn) (event.Subscription, error) {

	logs, sub, err := _PrivateSlot.contract.WatchLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateSlotBurn)
				if err := _PrivateSlot.contract.UnpackLog(event, "Burn", log); err != nil {
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

// PrivateSlotMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the PrivateSlot contract.
type PrivateSlotMintIterator struct {
	Event *PrivateSlotMint // Event containing the contract specifics and raw log

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
func (it *PrivateSlotMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateSlotMint)
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
		it.Event = new(PrivateSlotMint)
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
func (it *PrivateSlotMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateSlotMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateSlotMint represents a Mint event raised by the PrivateSlot contract.
type PrivateSlotMint struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(user address, amount uint256)
func (_PrivateSlot *PrivateSlotFilterer) FilterMint(opts *bind.FilterOpts) (*PrivateSlotMintIterator, error) {

	logs, sub, err := _PrivateSlot.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &PrivateSlotMintIterator{contract: _PrivateSlot.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(user address, amount uint256)
func (_PrivateSlot *PrivateSlotFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *PrivateSlotMint) (event.Subscription, error) {

	logs, sub, err := _PrivateSlot.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateSlotMint)
				if err := _PrivateSlot.contract.UnpackLog(event, "Mint", log); err != nil {
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

// PrivateSlotTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PrivateSlot contract.
type PrivateSlotTransferIterator struct {
	Event *PrivateSlotTransfer // Event containing the contract specifics and raw log

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
func (it *PrivateSlotTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateSlotTransfer)
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
		it.Event = new(PrivateSlotTransfer)
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
func (it *PrivateSlotTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateSlotTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateSlotTransfer represents a Transfer event raised by the PrivateSlot contract.
type PrivateSlotTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_PrivateSlot *PrivateSlotFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*PrivateSlotTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _PrivateSlot.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &PrivateSlotTransferIterator{contract: _PrivateSlot.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_PrivateSlot *PrivateSlotFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PrivateSlotTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _PrivateSlot.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateSlotTransfer)
				if err := _PrivateSlot.contract.UnpackLog(event, "Transfer", log); err != nil {
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

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
const PublicSlotABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"_success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pays\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"authorities\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"_balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requiredSignatures\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vs\",\"type\":\"uint8[]\"},{\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"pay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"_remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"name\":\"_requiredSignatures\",\"type\":\"uint256\"},{\"name\":\"_authorities\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Exchange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"Pay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"DepositConfirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// PublicSlotBin is the compiled bytecode used for deploying new contracts.
const PublicSlotBin = `60806040526040805190810160405280600481526020017f736c6f7400000000000000000000000000000000000000000000000000000000815250600090805190602001906200005192919062000168565b506040805190810160405280600381526020017f736c740000000000000000000000000000000000000000000000000000000000815250600190805190602001906200009f92919062000168565b5060006002556064600455348015620000b757600080fd5b50604051620016b9380380620016b98339810180604052810190808051906020019092919080519060200190929190805182019291905050508280600381905550600354600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550508160088190555080600990805190602001906200015e929190620001ef565b50505050620002ec565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001ab57805160ff1916838001178555620001dc565b82800160010185558215620001dc579182015b82811115620001db578251825591602001919060010190620001be565b5b509050620001eb91906200027e565b5090565b8280548282559060005260206000209081019282156200026b579160200282015b828111156200026a5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509160200191906001019062000210565b5b5090506200027a9190620002a6565b5090565b620002a391905b808211156200029f57600081600090555060010162000285565b5090565b90565b620002e991905b80821115620002e557600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101620002ad565b5090565b90565b6113bd80620002fc6000396000f3006080604052600436106100a4576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063045d0389146100a9578063095ea7b3146100f657806323b872dd1461015b5780633a1dce51146101e0578063494503d41461022957806370a08231146102965780638d068043146102ed578063a835ba0b14610318578063a9059cbb1461044a578063dd62ed3e146104af575b600080fd5b3480156100b557600080fd5b506100f4600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610526565b005b34801561010257600080fd5b50610141600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061062a565b604051808215151515815260200191505060405180910390f35b34801561016757600080fd5b506101c6600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061071c565b604051808215151515815260200191505060405180910390f35b3480156101ec57600080fd5b5061020f60048036038101908080356000191690602001909291905050506107bd565b604051808215151515815260200191505060405180910390f35b34801561023557600080fd5b50610254600480360381019080803590602001909291905050506107dd565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102a257600080fd5b506102d7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061081b565b6040518082815260200191505060405180910390f35b3480156102f957600080fd5b50610302610864565b6040518082815260200191505060405180910390f35b34801561032457600080fd5b50610448600480360381019080803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505091929192908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919291929080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061086a565b005b34801561045657600080fd5b50610495600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a8f565b604051808215151515815260200191505060405180910390f35b3480156104bb57600080fd5b50610510600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610aa4565b6040518082815260200191505060405180910390f35b80600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015151561057457600080fd5b80600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f477b65cf658c5207a9d60bb5ebe4f60a504af024949bdffa6efc396d01ced3f6836040518082815260200191505060405180910390a35050565b600081600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b6000600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482111515156107a957600080fd5b6107b4848484610b2b565b90509392505050565b60076020528060005260406000206000915054906101000a900460ff1681565b6009818154811015156107ec57fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60085481565b60006060806000806034602087510381151561088257fe5b04945061091b868a8a8a600980548060200260200160405190810160405280929190818152602001828054801561090e57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116108c4575b5050505050600854610d14565b151561092657600080fd5b6109308686610edd565b935061093c8686610f97565b925061094786611023565b915060076000836000191660001916815260200190815260200160002060009054906101000a900460ff1615151561097e57600080fd5b600160076000846000191660001916815260200190815260200160002060006101000a81548160ff021916908315150217905550600090505b84811015610a455782818151811015156109cd57fe5b906020019060200201516005600086848151811015156109e957fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508060010190506109b7565b7f4d1ed1c4c1a67f5295b85ef96fba12c16430d883bb2ebb1614bf898511ebb4418260405180826000191660001916815260200191505060405180910390a1505050505050505050565b6000610a9c338484610b2b565b905092915050565b6000600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6000600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548211151515610b7b57600080fd5b600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020540110151515610c0a57600080fd5b81600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555081600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3600190509392505050565b6000806060600080858a511015610d2e5760009450610ecf565b610d378b611036565b93508651604051908082528060200260200182016040528015610d695781602001602082028038833980820191505090505b509250600091505b85821015610eca576001848b84815181101515610d8a57fe5b906020019060200201518b85815181101515610da257fe5b906020019060200201518b86815181101515610dba57fe5b90602001906020020151604051600081526020016040526040518085600019166000191681526020018460ff1660ff1681526020018360001916600019168152602001826000191660001916815260200194505050505060206040516020810390808403906000865af1158015610e35573d6000803e3d6000fd5b505050602060405103519050610e4b87826111fc565b1515610e5a5760009450610ecf565b610e6483826111fc565b15610e725760009450610ecf565b808383815181101515610e8157fe5b9060200190602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508180600101925050610d71565b600194505b505050509695505050505050565b606080600080600085604051908082528060200260200182016040528015610f145781602001602082028038833980820191505090505b50935060349150600090505b85811015610f8a57818701519250828482815181101515610f3d57fe5b9060200190602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050603482019150806001019050610f20565b8394505050505092915050565b606080600080600085604051908082528060200260200182016040528015610fce5781602001602082028038833980820191505090505b50935060549150600090505b8581101561101657818701519250828482815181101515610ff757fe5b9060200190602002018181525050603482019150806001019050610fda565b8394505050505092915050565b6000806020830151905080915050919050565b600060606040805190810160405280601a81526020017f19457468657265756d205369676e6564204d6573736167653a0a00000000000081525090508061107d845161127a565b846040516020018084805190602001908083835b6020831015156110b65780518252602082019150602081019050602083039250611091565b6001836020036101000a03801982511681845116808217855250505050505090500183805190602001908083835b60208310151561110957805182526020820191506020810190506020830392506110e4565b6001836020036101000a03801982511681845116808217855250505050505090500182805190602001908083835b60208310151561115c5780518252602082019150602081019050602083039250611137565b6001836020036101000a03801982511681845116808217855250505050505090500193505050506040516020818303038152906040526040518082805190602001908083835b6020831015156111c757805182526020820191506020810190506020830392506111a2565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020915050919050565b600080600090505b835181101561126e578273ffffffffffffffffffffffffffffffffffffffff16848281518110151561123257fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1614156112615760019150611273565b8080600101915050611204565b600091505b5092915050565b6060600080606060008093508592505b8380600101945050600a8381151561129e57fe5b0492506000831415156112b05761128a565b836040519080825280601f01601f1916602001820160405280156112e35781602001602082028038833980820191505090505b5091506001840390508592505b600a838115156112fc57fe5b066030017f01000000000000000000000000000000000000000000000000000000000000000282828060019003935081518110151561133757fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a8381151561137357fe5b049250600083141515611385576112f0565b819450505050509190505600a165627a7a72305820a9102516bf11483563d74618668c798776567e943a1372857babbe8438d30fe30029`

// DeployPublicSlot deploys a new Ethereum contract, binding an instance of PublicSlot to it.
func DeployPublicSlot(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, _requiredSignatures *big.Int, _authorities []common.Address) (common.Address, *types.Transaction, *PublicSlot, error) {
	parsed, err := abi.JSON(strings.NewReader(PublicSlotABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PublicSlotBin), backend, initialSupply, _requiredSignatures, _authorities)
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
// Solidity: function allowance(_owner address, _spender address) constant returns(_remaining uint256)
func (_PublicSlot *PublicSlotCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PublicSlot.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(_remaining uint256)
func (_PublicSlot *PublicSlotSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _PublicSlot.Contract.Allowance(&_PublicSlot.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(_remaining uint256)
func (_PublicSlot *PublicSlotCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _PublicSlot.Contract.Allowance(&_PublicSlot.CallOpts, _owner, _spender)
}

// Authorities is a free data retrieval call binding the contract method 0x494503d4.
//
// Solidity: function authorities( uint256) constant returns(address)
func (_PublicSlot *PublicSlotCaller) Authorities(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PublicSlot.contract.Call(opts, out, "authorities", arg0)
	return *ret0, err
}

// Authorities is a free data retrieval call binding the contract method 0x494503d4.
//
// Solidity: function authorities( uint256) constant returns(address)
func (_PublicSlot *PublicSlotSession) Authorities(arg0 *big.Int) (common.Address, error) {
	return _PublicSlot.Contract.Authorities(&_PublicSlot.CallOpts, arg0)
}

// Authorities is a free data retrieval call binding the contract method 0x494503d4.
//
// Solidity: function authorities( uint256) constant returns(address)
func (_PublicSlot *PublicSlotCallerSession) Authorities(arg0 *big.Int) (common.Address, error) {
	return _PublicSlot.Contract.Authorities(&_PublicSlot.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_PublicSlot *PublicSlotCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PublicSlot.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_PublicSlot *PublicSlotSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _PublicSlot.Contract.BalanceOf(&_PublicSlot.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_PublicSlot *PublicSlotCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _PublicSlot.Contract.BalanceOf(&_PublicSlot.CallOpts, _owner)
}

// Pays is a free data retrieval call binding the contract method 0x3a1dce51.
//
// Solidity: function pays( bytes32) constant returns(bool)
func (_PublicSlot *PublicSlotCaller) Pays(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PublicSlot.contract.Call(opts, out, "pays", arg0)
	return *ret0, err
}

// Pays is a free data retrieval call binding the contract method 0x3a1dce51.
//
// Solidity: function pays( bytes32) constant returns(bool)
func (_PublicSlot *PublicSlotSession) Pays(arg0 [32]byte) (bool, error) {
	return _PublicSlot.Contract.Pays(&_PublicSlot.CallOpts, arg0)
}

// Pays is a free data retrieval call binding the contract method 0x3a1dce51.
//
// Solidity: function pays( bytes32) constant returns(bool)
func (_PublicSlot *PublicSlotCallerSession) Pays(arg0 [32]byte) (bool, error) {
	return _PublicSlot.Contract.Pays(&_PublicSlot.CallOpts, arg0)
}

// RequiredSignatures is a free data retrieval call binding the contract method 0x8d068043.
//
// Solidity: function requiredSignatures() constant returns(uint256)
func (_PublicSlot *PublicSlotCaller) RequiredSignatures(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PublicSlot.contract.Call(opts, out, "requiredSignatures")
	return *ret0, err
}

// RequiredSignatures is a free data retrieval call binding the contract method 0x8d068043.
//
// Solidity: function requiredSignatures() constant returns(uint256)
func (_PublicSlot *PublicSlotSession) RequiredSignatures() (*big.Int, error) {
	return _PublicSlot.Contract.RequiredSignatures(&_PublicSlot.CallOpts)
}

// RequiredSignatures is a free data retrieval call binding the contract method 0x8d068043.
//
// Solidity: function requiredSignatures() constant returns(uint256)
func (_PublicSlot *PublicSlotCallerSession) RequiredSignatures() (*big.Int, error) {
	return _PublicSlot.Contract.RequiredSignatures(&_PublicSlot.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(_success bool)
func (_PublicSlot *PublicSlotTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(_success bool)
func (_PublicSlot *PublicSlotSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Approve(&_PublicSlot.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(_success bool)
func (_PublicSlot *PublicSlotTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Approve(&_PublicSlot.TransactOpts, _spender, _value)
}

// Exchange is a paid mutator transaction binding the contract method 0x045d0389.
//
// Solidity: function exchange(_to address, _amount uint256) returns()
func (_PublicSlot *PublicSlotTransactor) Exchange(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "exchange", _to, _amount)
}

// Exchange is a paid mutator transaction binding the contract method 0x045d0389.
//
// Solidity: function exchange(_to address, _amount uint256) returns()
func (_PublicSlot *PublicSlotSession) Exchange(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Exchange(&_PublicSlot.TransactOpts, _to, _amount)
}

// Exchange is a paid mutator transaction binding the contract method 0x045d0389.
//
// Solidity: function exchange(_to address, _amount uint256) returns()
func (_PublicSlot *PublicSlotTransactorSession) Exchange(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Exchange(&_PublicSlot.TransactOpts, _to, _amount)
}

// Pay is a paid mutator transaction binding the contract method 0xa835ba0b.
//
// Solidity: function pay(vs uint8[], rs bytes32[], ss bytes32[], message bytes) returns()
func (_PublicSlot *PublicSlotTransactor) Pay(opts *bind.TransactOpts, vs []uint8, rs [][32]byte, ss [][32]byte, message []byte) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "pay", vs, rs, ss, message)
}

// Pay is a paid mutator transaction binding the contract method 0xa835ba0b.
//
// Solidity: function pay(vs uint8[], rs bytes32[], ss bytes32[], message bytes) returns()
func (_PublicSlot *PublicSlotSession) Pay(vs []uint8, rs [][32]byte, ss [][32]byte, message []byte) (*types.Transaction, error) {
	return _PublicSlot.Contract.Pay(&_PublicSlot.TransactOpts, vs, rs, ss, message)
}

// Pay is a paid mutator transaction binding the contract method 0xa835ba0b.
//
// Solidity: function pay(vs uint8[], rs bytes32[], ss bytes32[], message bytes) returns()
func (_PublicSlot *PublicSlotTransactorSession) Pay(vs []uint8, rs [][32]byte, ss [][32]byte, message []byte) (*types.Transaction, error) {
	return _PublicSlot.Contract.Pay(&_PublicSlot.TransactOpts, vs, rs, ss, message)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PublicSlot *PublicSlotTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PublicSlot *PublicSlotSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Transfer(&_PublicSlot.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PublicSlot *PublicSlotTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Transfer(&_PublicSlot.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PublicSlot *PublicSlotTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PublicSlot *PublicSlotSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.TransferFrom(&_PublicSlot.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PublicSlot *PublicSlotTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.TransferFrom(&_PublicSlot.TransactOpts, _from, _to, _value)
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

// PublicSlotDepositConfirmationIterator is returned from FilterDepositConfirmation and is used to iterate over the raw logs and unpacked data for DepositConfirmation events raised by the PublicSlot contract.
type PublicSlotDepositConfirmationIterator struct {
	Event *PublicSlotDepositConfirmation // Event containing the contract specifics and raw log

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
func (it *PublicSlotDepositConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicSlotDepositConfirmation)
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
		it.Event = new(PublicSlotDepositConfirmation)
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
func (it *PublicSlotDepositConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicSlotDepositConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicSlotDepositConfirmation represents a DepositConfirmation event raised by the PublicSlot contract.
type PublicSlotDepositConfirmation struct {
	Recipient       common.Address
	Value           *big.Int
	TransactionHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDepositConfirmation is a free log retrieval operation binding the contract event 0x82e9885b59946d922664d6e9c439efafebc983ce46aede53b1916dbb897c137a.
//
// Solidity: e DepositConfirmation(recipient address, value uint256, transactionHash bytes32)
func (_PublicSlot *PublicSlotFilterer) FilterDepositConfirmation(opts *bind.FilterOpts) (*PublicSlotDepositConfirmationIterator, error) {

	logs, sub, err := _PublicSlot.contract.FilterLogs(opts, "DepositConfirmation")
	if err != nil {
		return nil, err
	}
	return &PublicSlotDepositConfirmationIterator{contract: _PublicSlot.contract, event: "DepositConfirmation", logs: logs, sub: sub}, nil
}

// WatchDepositConfirmation is a free log subscription operation binding the contract event 0x82e9885b59946d922664d6e9c439efafebc983ce46aede53b1916dbb897c137a.
//
// Solidity: e DepositConfirmation(recipient address, value uint256, transactionHash bytes32)
func (_PublicSlot *PublicSlotFilterer) WatchDepositConfirmation(opts *bind.WatchOpts, sink chan<- *PublicSlotDepositConfirmation) (event.Subscription, error) {

	logs, sub, err := _PublicSlot.contract.WatchLogs(opts, "DepositConfirmation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicSlotDepositConfirmation)
				if err := _PublicSlot.contract.UnpackLog(event, "DepositConfirmation", log); err != nil {
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

// PublicSlotExchangeIterator is returned from FilterExchange and is used to iterate over the raw logs and unpacked data for Exchange events raised by the PublicSlot contract.
type PublicSlotExchangeIterator struct {
	Event *PublicSlotExchange // Event containing the contract specifics and raw log

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
func (it *PublicSlotExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicSlotExchange)
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
		it.Event = new(PublicSlotExchange)
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
func (it *PublicSlotExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicSlotExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicSlotExchange represents a Exchange event raised by the PublicSlot contract.
type PublicSlotExchange struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExchange is a free log retrieval operation binding the contract event 0x477b65cf658c5207a9d60bb5ebe4f60a504af024949bdffa6efc396d01ced3f6.
//
// Solidity: e Exchange(_from indexed address, _to indexed address, _amount uint256)
func (_PublicSlot *PublicSlotFilterer) FilterExchange(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*PublicSlotExchangeIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _PublicSlot.contract.FilterLogs(opts, "Exchange", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &PublicSlotExchangeIterator{contract: _PublicSlot.contract, event: "Exchange", logs: logs, sub: sub}, nil
}

// WatchExchange is a free log subscription operation binding the contract event 0x477b65cf658c5207a9d60bb5ebe4f60a504af024949bdffa6efc396d01ced3f6.
//
// Solidity: e Exchange(_from indexed address, _to indexed address, _amount uint256)
func (_PublicSlot *PublicSlotFilterer) WatchExchange(opts *bind.WatchOpts, sink chan<- *PublicSlotExchange, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _PublicSlot.contract.WatchLogs(opts, "Exchange", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicSlotExchange)
				if err := _PublicSlot.contract.UnpackLog(event, "Exchange", log); err != nil {
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

// PublicSlotPayIterator is returned from FilterPay and is used to iterate over the raw logs and unpacked data for Pay events raised by the PublicSlot contract.
type PublicSlotPayIterator struct {
	Event *PublicSlotPay // Event containing the contract specifics and raw log

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
func (it *PublicSlotPayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicSlotPay)
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
		it.Event = new(PublicSlotPay)
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
func (it *PublicSlotPayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicSlotPayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicSlotPay represents a Pay event raised by the PublicSlot contract.
type PublicSlotPay struct {
	TransactionHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPay is a free log retrieval operation binding the contract event 0x4d1ed1c4c1a67f5295b85ef96fba12c16430d883bb2ebb1614bf898511ebb441.
//
// Solidity: e Pay(transactionHash bytes32)
func (_PublicSlot *PublicSlotFilterer) FilterPay(opts *bind.FilterOpts) (*PublicSlotPayIterator, error) {

	logs, sub, err := _PublicSlot.contract.FilterLogs(opts, "Pay")
	if err != nil {
		return nil, err
	}
	return &PublicSlotPayIterator{contract: _PublicSlot.contract, event: "Pay", logs: logs, sub: sub}, nil
}

// WatchPay is a free log subscription operation binding the contract event 0x4d1ed1c4c1a67f5295b85ef96fba12c16430d883bb2ebb1614bf898511ebb441.
//
// Solidity: e Pay(transactionHash bytes32)
func (_PublicSlot *PublicSlotFilterer) WatchPay(opts *bind.WatchOpts, sink chan<- *PublicSlotPay) (event.Subscription, error) {

	logs, sub, err := _PublicSlot.contract.WatchLogs(opts, "Pay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicSlotPay)
				if err := _PublicSlot.contract.UnpackLog(event, "Pay", log); err != nil {
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

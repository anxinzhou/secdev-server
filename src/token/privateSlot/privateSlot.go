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
const PrivateSlotABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"signature\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"reward\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"by\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"consume\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"machine\",\"type\":\"address\"}],\"name\":\"removeGameMachine\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ownedAvatars\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"message_hash\",\"type\":\"bytes32\"}],\"name\":\"message\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"machine\",\"type\":\"address\"}],\"name\":\"addGameMachine\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"equipArmor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"pay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"avatar\",\"outputs\":[{\"name\":\"gene\",\"type\":\"uint256\"},{\"name\":\"avatarLevel\",\"type\":\"uint256\"},{\"name\":\"weaponed\",\"type\":\"bool\"},{\"name\":\"armored\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"submitSignature\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenID\",\"type\":\"uint256\"},{\"name\":\"avatarOwner\",\"type\":\"address\"},{\"name\":\"gene\",\"type\":\"uint256\"},{\"name\":\"avatarLevel\",\"type\":\"uint256\"},{\"name\":\"weaponed\",\"type\":\"bool\"},{\"name\":\"armored\",\"type\":\"bool\"}],\"name\":\"payNFT\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newRequiredSignatures\",\"type\":\"uint256\"}],\"name\":\"setRequiredSignatures\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requiredSignatures\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"_authorizdedMachines\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"equipWeapon\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"exchangeNFT\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"name\":\"decimalUnits\",\"type\":\"uint8\"},{\"name\":\"_requiredSignatures\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"WithdrawSignatureSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"authorityMachineResponsibleForRelay\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"CollectedSignatures\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"transactoinHash\",\"type\":\"bytes32\"}],\"name\":\"DepositConfirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Exchange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"transactoinHash\",\"type\":\"bytes32\"}],\"name\":\"Pay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"gene\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"avatarLevel\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"weaponed\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"armored\",\"type\":\"bool\"}],\"name\":\"ExchangeNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"machine\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Reward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"machine\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Consume\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// PrivateSlotBin is the compiled bytecode used for deploying new contracts.
const PrivateSlotBin = `60806040523480156200001157600080fd5b5060405162002d0e38038062002d0e833981018060405281019080805190602001909291908051820192919060200180518201929190602001805190602001909291908051906020019092919050505084848484838383838360008190555083600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508260019080519060200190620000cc9291906200019d565b5080600260006101000a81548160ff021916908360ff1602179055508160039080519060200190620001009291906200019d565b505050505033600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505080600b8190555033600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050506200024c565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001e057805160ff191683800117855562000211565b8280016001018555821562000211579182015b8281111562000210578251825591602001919060010190620001f3565b5b50905062000220919062000224565b5090565b6200024991905b80821115620002455760008160009055506001016200022b565b5090565b90565b612ab2806200025c6000396000f300608060405260043610610180576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063045d03891461018557806306fdde03146101d2578063095ea7b3146102625780631812d996146102c757806318160ddd1461037b57806321670f22146103a6578063224b5c721461040b57806323b872dd146104705780633016e108146104f5578063313ce56714610538578063399bd4f01461056957806340c10f19146105c057806345977d031461060d578063490a32c61461063a57806350dc43a1146106e4578063541486b6146107275780635e5571ac146107745780635e5741f5146107cf578063630cea8e1461082d5780636352211e146108dc57806366ff660f1461094957806370a08231146109c25780637d2b9cc014610a195780638d06804314610a4657806390883c0814610a7157806395d89b4114610acc578063a9059cbb14610b5c578063bf6feb4014610bc1578063dd62ed3e14610c0e578063e9565b8314610c85575b600080fd5b34801561019157600080fd5b506101d0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610cb2565b005b3480156101de57600080fd5b506101e7610d4e565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561022757808201518184015260208101905061020c565b50505050905090810190601f1680156102545780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561026e57600080fd5b506102ad600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610dec565b604051808215151515815260200191505060405180910390f35b3480156102d357600080fd5b50610300600480360381019080803560001916906020019092919080359060200190929190505050610f19565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610340578082015181840152602081019050610325565b50505050905090810190601f16801561036d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561038757600080fd5b50610390610ff1565b6040518082815260200191505060405180910390f35b3480156103b257600080fd5b506103f1600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610ffa565b604051808215151515815260200191505060405180910390f35b34801561041757600080fd5b50610456600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506110ce565b604051808215151515815260200191505060405180910390f35b34801561047c57600080fd5b506104db600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506111a2565b604051808215151515815260200191505060405180910390f35b34801561050157600080fd5b50610536600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506113b9565b005b34801561054457600080fd5b5061054d611470565b604051808260ff1660ff16815260200191505060405180910390f35b34801561057557600080fd5b506105aa600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611483565b6040518082815260200191505060405180910390f35b3480156105cc57600080fd5b5061060b600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506114cc565b005b34801561061957600080fd5b50610638600480360381019080803590602001909291905050506115c9565b005b34801561064657600080fd5b506106696004803603810190808035600019169060200190929190505050611616565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106a957808201518184015260208101905061068e565b50505050905090810190601f1680156106d65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156106f057600080fd5b50610725600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506116d6565b005b34801561073357600080fd5b5061077260048036038101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061178d565b005b34801561078057600080fd5b506107cd600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803560001916906020019092919050505061185d565b005b3480156107db57600080fd5b506107fa60048036038101908080359060200190929190505050611bfc565b60405180858152602001848152602001831515151581526020018215151515815260200194505050505060405180910390f35b34801561083957600080fd5b506108da600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611c46565b005b3480156108e857600080fd5b5061090760048036038101908080359060200190929190505050611fb5565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561095557600080fd5b506109c060048036038101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190803515159060200190929190803515159060200190929190505050612033565b005b3480156109ce57600080fd5b50610a03600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612163565b6040518082815260200191505060405180910390f35b348015610a2557600080fd5b50610a44600480360381019080803590602001909291905050506121ac565b005b348015610a5257600080fd5b50610a5b612212565b6040518082815260200191505060405180910390f35b348015610a7d57600080fd5b50610ab2600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612218565b604051808215151515815260200191505060405180910390f35b348015610ad857600080fd5b50610ae1612238565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610b21578082015181840152602081019050610b06565b50505050905090810190601f168015610b4e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b348015610b6857600080fd5b50610ba7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506122d6565b604051808215151515815260200191505060405180910390f35b348015610bcd57600080fd5b50610c0c60048036038101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612352565b005b348015610c1a57600080fd5b50610c6f600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612422565b6040518082815260200191505060405180910390f35b348015610c9157600080fd5b50610cb0600480360381019080803590602001909291905050506124a9565b005b610cdf82600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683612768565b7f5988e4c12f4844b895de0739f562558435dca9602fd8b970720ee3cf8dff39be8282604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a15050565b60018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610de45780601f10610db957610100808354040283529160200191610de4565b820191906000526020600020905b815481529060010190602001808311610dc757829003601f168201915b505050505081565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614151515610e2957600080fd5b81600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b6060600d6000846000191660001916815260200190815260200160002060020182815481101515610f4657fe5b906000526020600020018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610fe45780601f10610fb957610100808354040283529160200191610fe4565b820191906000526020600020905b815481529060010190602001808311610fc757829003601f168201915b5050505050905092915050565b60008054905090565b6000600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151561105457600080fd5b61105f338484612768565b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f6b053894d8fdbdcc936dd753e21291f0c48e68ef12306eb39a63a374147ba4bd846040518082815260200191505060405180910390a36001905092915050565b6000600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151561112857600080fd5b611133833384612768565b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f073d7524ea37f2bed5915b0eca0c75d97c80044d12f5d5b2ecc1bc88a685780f846040518082815260200191505060405180910390a36001905092915050565b6000600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054821115151561122f57600080fd5b6112be82600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461292190919063ffffffff16565b600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611349848484612768565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3600190509392505050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561141557600080fd5b6000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b600260009054906101000a900460ff1681565b6000600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151561150857600080fd5b816008600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506002428115156115aa57fe5b0660096000838152602001908152602001600020600001819055505050565b600260096000838152602001908152602001600020600101541015156115ee57600080fd5b6001600960008381526020019081526020016000206001016000828254019250508190555050565b6060600d600083600019166000191681526020019081526020016000206000018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156116ca5780601f1061169f576101008083540402835291602001916116ca565b820191906000526020600020905b8154815290600101906020018083116116ad57829003601f168201915b50505050509050919050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561173257600080fd5b6001600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b8073ffffffffffffffffffffffffffffffffffffffff166008600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415156117fa57600080fd5b6009600083815260200190815260200160002060020160019054906101000a900460ff1615151561182a57600080fd5b60016009600084815260200190815260200160002060020160016101000a81548160ff0219169083151502179055505050565b6000600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156118b757600080fd5b838383604051602001808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401838152602001826000191660001916815260200193505050506040516020818303038152906040526040518082805190602001908083835b60208310151561195f578051825260208201915060208101905060208303925061193a565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209050611a37600c60008360001916600019168152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015611a2c57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116119e2575b505050505033612942565b151515611a4357600080fd5b600c600082600019166000191681526020019081526020016000203390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600b54600c60008360001916600019168152602001908152602001600020805490501415611b7a57611b17600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168585612768565b8373ffffffffffffffffffffffffffffffffffffffff167f1b174056799bea141540e324bb093eb297a02b564c15e75840a30cf0d0f4837784846040518083815260200182600019166000191681526020019250505060405180910390a2611bf6565b7f82e9885b59946d922664d6e9c439efafebc983ce46aede53b1916dbb897c137a848484604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018260001916600019168152602001935050505060405180910390a15b50505050565b60096020528060005260406000206000915090508060000154908060010154908060020160009054906101000a900460ff16908060020160019054906101000a900460ff16905084565b6000600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515611ca057600080fd5b826040518082805190602001908083835b602083101515611cd65780518252602082019150602081019050602083039250611cb1565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209050611db1600d60008360001916600019168152602001908152602001600020600101805480602002602001604051908101604052809291908181526020018280548015611da657602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611d5c575b505050505033612942565b151515611dbd57600080fd5b82600d600083600019166000191681526020019081526020016000206000019080519060200190611def9291906129e1565b50600d600082600019166000191681526020019081526020016000206001013390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600d60008260001916600019168152602001908152602001600020600201829080600181540180825580915050906001820390600052602060002001600090919290919091509080519060200190611ecb9291906129e1565b5050600b54600d60008360001916600019168152602001908152602001600020600101805490501415611f70577feb043d149eedb81369bec43d4c3a3a53087debc88d2525f13bfaa3eecda28b5c3382604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182600019166000191681526020019250505060405180910390a1611fb0565b7ffc962d5f0de9bd4b90181a7ec60354bcefe1e162eaf760919b07496a184d665a8160405180826000191660001916815260200191505060405180910390a15b505050565b6000806008600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561202a57600080fd5b80915050919050565b85600a60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550846008600088815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550836009600088815260200190815260200160002060000181905550826009600088815260200190815260200160002060010181905550816009600088815260200190815260200160002060020160006101000a81548160ff021916908315150217905550806009600088815260200190815260200160002060020160016101000a81548160ff021916908315150217905550505050505050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561220857600080fd5b80600b8190555050565b600b5481565b60076020528060005260406000206000915054906101000a900460ff1681565b60038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156122ce5780601f106122a3576101008083540402835291602001916122ce565b820191906000526020600020905b8154815290600101906020018083116122b157829003601f168201915b505050505081565b60006122e3338484612768565b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a36001905092915050565b8073ffffffffffffffffffffffffffffffffffffffff166008600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415156123bf57600080fd5b6009600083815260200190815260200160002060020160009054906101000a900460ff161515156123ef57600080fd5b60016009600084815260200190815260200160002060020160006101000a81548160ff0219169083151502179055505050565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b60008060008060006008600087815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1694508473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561252157600080fd5b6000600a60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060006008600088815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060096000878152602001908152602001600020600001549350600960008781526020019081526020016000206001015492506009600087815260200190815260200160002060020160009054906101000a900460ff1691506009600087815260200190815260200160002060020160019054906101000a900460ff169050600060096000888152602001908152602001600020600001819055506000600960008881526020019081526020016000206001018190555060006009600088815260200190815260200160002060020160006101000a81548160ff02191690831515021790555060006009600088815260200190815260200160002060020160016101000a81548160ff0219169083151502179055507f2a797bc85caee23a3e29a6e8fb0614b4ffa925d434e247936693cc55389d7f87868686868686604051808781526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020018315151515815260200182151515158152602001965050505050505060405180910390a1505050505050565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481111515156127b657600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141515156127f257600080fd5b61284481600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461292190919063ffffffff16565b600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506128d981600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546129c090919063ffffffff16565b600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505050565b60008083831115151561293357600080fd5b82840390508091505092915050565b600080600090505b83518110156129b4578273ffffffffffffffffffffffffffffffffffffffff16848281518110151561297857fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1614156129a757600191506129b9565b808060010191505061294a565b600091505b5092915050565b60008082840190508381101515156129d757600080fd5b8091505092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10612a2257805160ff1916838001178555612a50565b82800160010185558215612a50579182015b82811115612a4f578251825591602001919060010190612a34565b5b509050612a5d9190612a61565b5090565b612a8391905b80821115612a7f576000816000905550600101612a67565b5090565b905600a165627a7a72305820e1481d8f1b84fd2dbd3308c9ec6f11573b8d7cab9bab536c8dce7f61e3b5979e0029`

// DeployPrivateSlot deploys a new Ethereum contract, binding an instance of PrivateSlot to it.
func DeployPrivateSlot(auth *bind.TransactOpts, backend bind.ContractBackend, totalSupply *big.Int, tokenName string, tokenSymbol string, decimalUnits uint8, _requiredSignatures *big.Int) (common.Address, *types.Transaction, *PrivateSlot, error) {
	parsed, err := abi.JSON(strings.NewReader(PrivateSlotABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PrivateSlotBin), backend, totalSupply, tokenName, tokenSymbol, decimalUnits, _requiredSignatures)
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

// AuthorizdedMachines is a free data retrieval call binding the contract method 0x90883c08.
//
// Solidity: function _authorizdedMachines( address) constant returns(bool)
func (_PrivateSlot *PrivateSlotCaller) AuthorizdedMachines(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PrivateSlot.contract.Call(opts, out, "_authorizdedMachines", arg0)
	return *ret0, err
}

// AuthorizdedMachines is a free data retrieval call binding the contract method 0x90883c08.
//
// Solidity: function _authorizdedMachines( address) constant returns(bool)
func (_PrivateSlot *PrivateSlotSession) AuthorizdedMachines(arg0 common.Address) (bool, error) {
	return _PrivateSlot.Contract.AuthorizdedMachines(&_PrivateSlot.CallOpts, arg0)
}

// AuthorizdedMachines is a free data retrieval call binding the contract method 0x90883c08.
//
// Solidity: function _authorizdedMachines( address) constant returns(bool)
func (_PrivateSlot *PrivateSlotCallerSession) AuthorizdedMachines(arg0 common.Address) (bool, error) {
	return _PrivateSlot.Contract.AuthorizdedMachines(&_PrivateSlot.CallOpts, arg0)
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

// Avatar is a free data retrieval call binding the contract method 0x5e5741f5.
//
// Solidity: function avatar( uint256) constant returns(gene uint256, avatarLevel uint256, weaponed bool, armored bool)
func (_PrivateSlot *PrivateSlotCaller) Avatar(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Gene        *big.Int
	AvatarLevel *big.Int
	Weaponed    bool
	Armored     bool
}, error) {
	ret := new(struct {
		Gene        *big.Int
		AvatarLevel *big.Int
		Weaponed    bool
		Armored     bool
	})
	out := ret
	err := _PrivateSlot.contract.Call(opts, out, "avatar", arg0)
	return *ret, err
}

// Avatar is a free data retrieval call binding the contract method 0x5e5741f5.
//
// Solidity: function avatar( uint256) constant returns(gene uint256, avatarLevel uint256, weaponed bool, armored bool)
func (_PrivateSlot *PrivateSlotSession) Avatar(arg0 *big.Int) (struct {
	Gene        *big.Int
	AvatarLevel *big.Int
	Weaponed    bool
	Armored     bool
}, error) {
	return _PrivateSlot.Contract.Avatar(&_PrivateSlot.CallOpts, arg0)
}

// Avatar is a free data retrieval call binding the contract method 0x5e5741f5.
//
// Solidity: function avatar( uint256) constant returns(gene uint256, avatarLevel uint256, weaponed bool, armored bool)
func (_PrivateSlot *PrivateSlotCallerSession) Avatar(arg0 *big.Int) (struct {
	Gene        *big.Int
	AvatarLevel *big.Int
	Weaponed    bool
	Armored     bool
}, error) {
	return _PrivateSlot.Contract.Avatar(&_PrivateSlot.CallOpts, arg0)
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

// Message is a free data retrieval call binding the contract method 0x490a32c6.
//
// Solidity: function message(message_hash bytes32) constant returns(bytes)
func (_PrivateSlot *PrivateSlotCaller) Message(opts *bind.CallOpts, message_hash [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _PrivateSlot.contract.Call(opts, out, "message", message_hash)
	return *ret0, err
}

// Message is a free data retrieval call binding the contract method 0x490a32c6.
//
// Solidity: function message(message_hash bytes32) constant returns(bytes)
func (_PrivateSlot *PrivateSlotSession) Message(message_hash [32]byte) ([]byte, error) {
	return _PrivateSlot.Contract.Message(&_PrivateSlot.CallOpts, message_hash)
}

// Message is a free data retrieval call binding the contract method 0x490a32c6.
//
// Solidity: function message(message_hash bytes32) constant returns(bytes)
func (_PrivateSlot *PrivateSlotCallerSession) Message(message_hash [32]byte) ([]byte, error) {
	return _PrivateSlot.Contract.Message(&_PrivateSlot.CallOpts, message_hash)
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

// OwnedAvatars is a free data retrieval call binding the contract method 0x399bd4f0.
//
// Solidity: function ownedAvatars(owner address) constant returns(uint256)
func (_PrivateSlot *PrivateSlotCaller) OwnedAvatars(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivateSlot.contract.Call(opts, out, "ownedAvatars", owner)
	return *ret0, err
}

// OwnedAvatars is a free data retrieval call binding the contract method 0x399bd4f0.
//
// Solidity: function ownedAvatars(owner address) constant returns(uint256)
func (_PrivateSlot *PrivateSlotSession) OwnedAvatars(owner common.Address) (*big.Int, error) {
	return _PrivateSlot.Contract.OwnedAvatars(&_PrivateSlot.CallOpts, owner)
}

// OwnedAvatars is a free data retrieval call binding the contract method 0x399bd4f0.
//
// Solidity: function ownedAvatars(owner address) constant returns(uint256)
func (_PrivateSlot *PrivateSlotCallerSession) OwnedAvatars(owner common.Address) (*big.Int, error) {
	return _PrivateSlot.Contract.OwnedAvatars(&_PrivateSlot.CallOpts, owner)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(tokenId uint256) constant returns(address)
func (_PrivateSlot *PrivateSlotCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PrivateSlot.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(tokenId uint256) constant returns(address)
func (_PrivateSlot *PrivateSlotSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _PrivateSlot.Contract.OwnerOf(&_PrivateSlot.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(tokenId uint256) constant returns(address)
func (_PrivateSlot *PrivateSlotCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _PrivateSlot.Contract.OwnerOf(&_PrivateSlot.CallOpts, tokenId)
}

// RequiredSignatures is a free data retrieval call binding the contract method 0x8d068043.
//
// Solidity: function requiredSignatures() constant returns(uint256)
func (_PrivateSlot *PrivateSlotCaller) RequiredSignatures(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivateSlot.contract.Call(opts, out, "requiredSignatures")
	return *ret0, err
}

// RequiredSignatures is a free data retrieval call binding the contract method 0x8d068043.
//
// Solidity: function requiredSignatures() constant returns(uint256)
func (_PrivateSlot *PrivateSlotSession) RequiredSignatures() (*big.Int, error) {
	return _PrivateSlot.Contract.RequiredSignatures(&_PrivateSlot.CallOpts)
}

// RequiredSignatures is a free data retrieval call binding the contract method 0x8d068043.
//
// Solidity: function requiredSignatures() constant returns(uint256)
func (_PrivateSlot *PrivateSlotCallerSession) RequiredSignatures() (*big.Int, error) {
	return _PrivateSlot.Contract.RequiredSignatures(&_PrivateSlot.CallOpts)
}

// Signature is a free data retrieval call binding the contract method 0x1812d996.
//
// Solidity: function signature(messageHash bytes32, index uint256) constant returns(bytes)
func (_PrivateSlot *PrivateSlotCaller) Signature(opts *bind.CallOpts, messageHash [32]byte, index *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _PrivateSlot.contract.Call(opts, out, "signature", messageHash, index)
	return *ret0, err
}

// Signature is a free data retrieval call binding the contract method 0x1812d996.
//
// Solidity: function signature(messageHash bytes32, index uint256) constant returns(bytes)
func (_PrivateSlot *PrivateSlotSession) Signature(messageHash [32]byte, index *big.Int) ([]byte, error) {
	return _PrivateSlot.Contract.Signature(&_PrivateSlot.CallOpts, messageHash, index)
}

// Signature is a free data retrieval call binding the contract method 0x1812d996.
//
// Solidity: function signature(messageHash bytes32, index uint256) constant returns(bytes)
func (_PrivateSlot *PrivateSlotCallerSession) Signature(messageHash [32]byte, index *big.Int) ([]byte, error) {
	return _PrivateSlot.Contract.Signature(&_PrivateSlot.CallOpts, messageHash, index)
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

// AddGameMachine is a paid mutator transaction binding the contract method 0x50dc43a1.
//
// Solidity: function addGameMachine(machine address) returns()
func (_PrivateSlot *PrivateSlotTransactor) AddGameMachine(opts *bind.TransactOpts, machine common.Address) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "addGameMachine", machine)
}

// AddGameMachine is a paid mutator transaction binding the contract method 0x50dc43a1.
//
// Solidity: function addGameMachine(machine address) returns()
func (_PrivateSlot *PrivateSlotSession) AddGameMachine(machine common.Address) (*types.Transaction, error) {
	return _PrivateSlot.Contract.AddGameMachine(&_PrivateSlot.TransactOpts, machine)
}

// AddGameMachine is a paid mutator transaction binding the contract method 0x50dc43a1.
//
// Solidity: function addGameMachine(machine address) returns()
func (_PrivateSlot *PrivateSlotTransactorSession) AddGameMachine(machine common.Address) (*types.Transaction, error) {
	return _PrivateSlot.Contract.AddGameMachine(&_PrivateSlot.TransactOpts, machine)
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

// EquipArmor is a paid mutator transaction binding the contract method 0x541486b6.
//
// Solidity: function equipArmor(tokenId uint256, user address) returns()
func (_PrivateSlot *PrivateSlotTransactor) EquipArmor(opts *bind.TransactOpts, tokenId *big.Int, user common.Address) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "equipArmor", tokenId, user)
}

// EquipArmor is a paid mutator transaction binding the contract method 0x541486b6.
//
// Solidity: function equipArmor(tokenId uint256, user address) returns()
func (_PrivateSlot *PrivateSlotSession) EquipArmor(tokenId *big.Int, user common.Address) (*types.Transaction, error) {
	return _PrivateSlot.Contract.EquipArmor(&_PrivateSlot.TransactOpts, tokenId, user)
}

// EquipArmor is a paid mutator transaction binding the contract method 0x541486b6.
//
// Solidity: function equipArmor(tokenId uint256, user address) returns()
func (_PrivateSlot *PrivateSlotTransactorSession) EquipArmor(tokenId *big.Int, user common.Address) (*types.Transaction, error) {
	return _PrivateSlot.Contract.EquipArmor(&_PrivateSlot.TransactOpts, tokenId, user)
}

// EquipWeapon is a paid mutator transaction binding the contract method 0xbf6feb40.
//
// Solidity: function equipWeapon(tokenId uint256, user address) returns()
func (_PrivateSlot *PrivateSlotTransactor) EquipWeapon(opts *bind.TransactOpts, tokenId *big.Int, user common.Address) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "equipWeapon", tokenId, user)
}

// EquipWeapon is a paid mutator transaction binding the contract method 0xbf6feb40.
//
// Solidity: function equipWeapon(tokenId uint256, user address) returns()
func (_PrivateSlot *PrivateSlotSession) EquipWeapon(tokenId *big.Int, user common.Address) (*types.Transaction, error) {
	return _PrivateSlot.Contract.EquipWeapon(&_PrivateSlot.TransactOpts, tokenId, user)
}

// EquipWeapon is a paid mutator transaction binding the contract method 0xbf6feb40.
//
// Solidity: function equipWeapon(tokenId uint256, user address) returns()
func (_PrivateSlot *PrivateSlotTransactorSession) EquipWeapon(tokenId *big.Int, user common.Address) (*types.Transaction, error) {
	return _PrivateSlot.Contract.EquipWeapon(&_PrivateSlot.TransactOpts, tokenId, user)
}

// Exchange is a paid mutator transaction binding the contract method 0x045d0389.
//
// Solidity: function exchange(user address, amount uint256) returns()
func (_PrivateSlot *PrivateSlotTransactor) Exchange(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "exchange", user, amount)
}

// Exchange is a paid mutator transaction binding the contract method 0x045d0389.
//
// Solidity: function exchange(user address, amount uint256) returns()
func (_PrivateSlot *PrivateSlotSession) Exchange(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Exchange(&_PrivateSlot.TransactOpts, user, amount)
}

// Exchange is a paid mutator transaction binding the contract method 0x045d0389.
//
// Solidity: function exchange(user address, amount uint256) returns()
func (_PrivateSlot *PrivateSlotTransactorSession) Exchange(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Exchange(&_PrivateSlot.TransactOpts, user, amount)
}

// ExchangeNFT is a paid mutator transaction binding the contract method 0xe9565b83.
//
// Solidity: function exchangeNFT(tokenID uint256) returns()
func (_PrivateSlot *PrivateSlotTransactor) ExchangeNFT(opts *bind.TransactOpts, tokenID *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "exchangeNFT", tokenID)
}

// ExchangeNFT is a paid mutator transaction binding the contract method 0xe9565b83.
//
// Solidity: function exchangeNFT(tokenID uint256) returns()
func (_PrivateSlot *PrivateSlotSession) ExchangeNFT(tokenID *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.ExchangeNFT(&_PrivateSlot.TransactOpts, tokenID)
}

// ExchangeNFT is a paid mutator transaction binding the contract method 0xe9565b83.
//
// Solidity: function exchangeNFT(tokenID uint256) returns()
func (_PrivateSlot *PrivateSlotTransactorSession) ExchangeNFT(tokenID *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.ExchangeNFT(&_PrivateSlot.TransactOpts, tokenID)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(to address, tokenId uint256) returns()
func (_PrivateSlot *PrivateSlotTransactor) Mint(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "mint", to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(to address, tokenId uint256) returns()
func (_PrivateSlot *PrivateSlotSession) Mint(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Mint(&_PrivateSlot.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(to address, tokenId uint256) returns()
func (_PrivateSlot *PrivateSlotTransactorSession) Mint(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Mint(&_PrivateSlot.TransactOpts, to, tokenId)
}

// Pay is a paid mutator transaction binding the contract method 0x5e5571ac.
//
// Solidity: function pay(user address, amount uint256, transactionHash bytes32) returns()
func (_PrivateSlot *PrivateSlotTransactor) Pay(opts *bind.TransactOpts, user common.Address, amount *big.Int, transactionHash [32]byte) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "pay", user, amount, transactionHash)
}

// Pay is a paid mutator transaction binding the contract method 0x5e5571ac.
//
// Solidity: function pay(user address, amount uint256, transactionHash bytes32) returns()
func (_PrivateSlot *PrivateSlotSession) Pay(user common.Address, amount *big.Int, transactionHash [32]byte) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Pay(&_PrivateSlot.TransactOpts, user, amount, transactionHash)
}

// Pay is a paid mutator transaction binding the contract method 0x5e5571ac.
//
// Solidity: function pay(user address, amount uint256, transactionHash bytes32) returns()
func (_PrivateSlot *PrivateSlotTransactorSession) Pay(user common.Address, amount *big.Int, transactionHash [32]byte) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Pay(&_PrivateSlot.TransactOpts, user, amount, transactionHash)
}

// PayNFT is a paid mutator transaction binding the contract method 0x66ff660f.
//
// Solidity: function payNFT(tokenID uint256, avatarOwner address, gene uint256, avatarLevel uint256, weaponed bool, armored bool) returns()
func (_PrivateSlot *PrivateSlotTransactor) PayNFT(opts *bind.TransactOpts, tokenID *big.Int, avatarOwner common.Address, gene *big.Int, avatarLevel *big.Int, weaponed bool, armored bool) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "payNFT", tokenID, avatarOwner, gene, avatarLevel, weaponed, armored)
}

// PayNFT is a paid mutator transaction binding the contract method 0x66ff660f.
//
// Solidity: function payNFT(tokenID uint256, avatarOwner address, gene uint256, avatarLevel uint256, weaponed bool, armored bool) returns()
func (_PrivateSlot *PrivateSlotSession) PayNFT(tokenID *big.Int, avatarOwner common.Address, gene *big.Int, avatarLevel *big.Int, weaponed bool, armored bool) (*types.Transaction, error) {
	return _PrivateSlot.Contract.PayNFT(&_PrivateSlot.TransactOpts, tokenID, avatarOwner, gene, avatarLevel, weaponed, armored)
}

// PayNFT is a paid mutator transaction binding the contract method 0x66ff660f.
//
// Solidity: function payNFT(tokenID uint256, avatarOwner address, gene uint256, avatarLevel uint256, weaponed bool, armored bool) returns()
func (_PrivateSlot *PrivateSlotTransactorSession) PayNFT(tokenID *big.Int, avatarOwner common.Address, gene *big.Int, avatarLevel *big.Int, weaponed bool, armored bool) (*types.Transaction, error) {
	return _PrivateSlot.Contract.PayNFT(&_PrivateSlot.TransactOpts, tokenID, avatarOwner, gene, avatarLevel, weaponed, armored)
}

// RemoveGameMachine is a paid mutator transaction binding the contract method 0x3016e108.
//
// Solidity: function removeGameMachine(machine address) returns()
func (_PrivateSlot *PrivateSlotTransactor) RemoveGameMachine(opts *bind.TransactOpts, machine common.Address) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "removeGameMachine", machine)
}

// RemoveGameMachine is a paid mutator transaction binding the contract method 0x3016e108.
//
// Solidity: function removeGameMachine(machine address) returns()
func (_PrivateSlot *PrivateSlotSession) RemoveGameMachine(machine common.Address) (*types.Transaction, error) {
	return _PrivateSlot.Contract.RemoveGameMachine(&_PrivateSlot.TransactOpts, machine)
}

// RemoveGameMachine is a paid mutator transaction binding the contract method 0x3016e108.
//
// Solidity: function removeGameMachine(machine address) returns()
func (_PrivateSlot *PrivateSlotTransactorSession) RemoveGameMachine(machine common.Address) (*types.Transaction, error) {
	return _PrivateSlot.Contract.RemoveGameMachine(&_PrivateSlot.TransactOpts, machine)
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

// SetRequiredSignatures is a paid mutator transaction binding the contract method 0x7d2b9cc0.
//
// Solidity: function setRequiredSignatures(newRequiredSignatures uint256) returns()
func (_PrivateSlot *PrivateSlotTransactor) SetRequiredSignatures(opts *bind.TransactOpts, newRequiredSignatures *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "setRequiredSignatures", newRequiredSignatures)
}

// SetRequiredSignatures is a paid mutator transaction binding the contract method 0x7d2b9cc0.
//
// Solidity: function setRequiredSignatures(newRequiredSignatures uint256) returns()
func (_PrivateSlot *PrivateSlotSession) SetRequiredSignatures(newRequiredSignatures *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.SetRequiredSignatures(&_PrivateSlot.TransactOpts, newRequiredSignatures)
}

// SetRequiredSignatures is a paid mutator transaction binding the contract method 0x7d2b9cc0.
//
// Solidity: function setRequiredSignatures(newRequiredSignatures uint256) returns()
func (_PrivateSlot *PrivateSlotTransactorSession) SetRequiredSignatures(newRequiredSignatures *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.SetRequiredSignatures(&_PrivateSlot.TransactOpts, newRequiredSignatures)
}

// SubmitSignature is a paid mutator transaction binding the contract method 0x630cea8e.
//
// Solidity: function submitSignature(message bytes, signature bytes) returns()
func (_PrivateSlot *PrivateSlotTransactor) SubmitSignature(opts *bind.TransactOpts, message []byte, signature []byte) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "submitSignature", message, signature)
}

// SubmitSignature is a paid mutator transaction binding the contract method 0x630cea8e.
//
// Solidity: function submitSignature(message bytes, signature bytes) returns()
func (_PrivateSlot *PrivateSlotSession) SubmitSignature(message []byte, signature []byte) (*types.Transaction, error) {
	return _PrivateSlot.Contract.SubmitSignature(&_PrivateSlot.TransactOpts, message, signature)
}

// SubmitSignature is a paid mutator transaction binding the contract method 0x630cea8e.
//
// Solidity: function submitSignature(message bytes, signature bytes) returns()
func (_PrivateSlot *PrivateSlotTransactorSession) SubmitSignature(message []byte, signature []byte) (*types.Transaction, error) {
	return _PrivateSlot.Contract.SubmitSignature(&_PrivateSlot.TransactOpts, message, signature)
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

// Upgrade is a paid mutator transaction binding the contract method 0x45977d03.
//
// Solidity: function upgrade(tokenId uint256) returns()
func (_PrivateSlot *PrivateSlotTransactor) Upgrade(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "upgrade", tokenId)
}

// Upgrade is a paid mutator transaction binding the contract method 0x45977d03.
//
// Solidity: function upgrade(tokenId uint256) returns()
func (_PrivateSlot *PrivateSlotSession) Upgrade(tokenId *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Upgrade(&_PrivateSlot.TransactOpts, tokenId)
}

// Upgrade is a paid mutator transaction binding the contract method 0x45977d03.
//
// Solidity: function upgrade(tokenId uint256) returns()
func (_PrivateSlot *PrivateSlotTransactorSession) Upgrade(tokenId *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Upgrade(&_PrivateSlot.TransactOpts, tokenId)
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

// PrivateSlotCollectedSignaturesIterator is returned from FilterCollectedSignatures and is used to iterate over the raw logs and unpacked data for CollectedSignatures events raised by the PrivateSlot contract.
type PrivateSlotCollectedSignaturesIterator struct {
	Event *PrivateSlotCollectedSignatures // Event containing the contract specifics and raw log

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
func (it *PrivateSlotCollectedSignaturesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateSlotCollectedSignatures)
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
		it.Event = new(PrivateSlotCollectedSignatures)
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
func (it *PrivateSlotCollectedSignaturesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateSlotCollectedSignaturesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateSlotCollectedSignatures represents a CollectedSignatures event raised by the PrivateSlot contract.
type PrivateSlotCollectedSignatures struct {
	AuthorityMachineResponsibleForRelay common.Address
	MessageHash                         [32]byte
	Raw                                 types.Log // Blockchain specific contextual infos
}

// FilterCollectedSignatures is a free log retrieval operation binding the contract event 0xeb043d149eedb81369bec43d4c3a3a53087debc88d2525f13bfaa3eecda28b5c.
//
// Solidity: e CollectedSignatures(authorityMachineResponsibleForRelay address, messageHash bytes32)
func (_PrivateSlot *PrivateSlotFilterer) FilterCollectedSignatures(opts *bind.FilterOpts) (*PrivateSlotCollectedSignaturesIterator, error) {

	logs, sub, err := _PrivateSlot.contract.FilterLogs(opts, "CollectedSignatures")
	if err != nil {
		return nil, err
	}
	return &PrivateSlotCollectedSignaturesIterator{contract: _PrivateSlot.contract, event: "CollectedSignatures", logs: logs, sub: sub}, nil
}

// WatchCollectedSignatures is a free log subscription operation binding the contract event 0xeb043d149eedb81369bec43d4c3a3a53087debc88d2525f13bfaa3eecda28b5c.
//
// Solidity: e CollectedSignatures(authorityMachineResponsibleForRelay address, messageHash bytes32)
func (_PrivateSlot *PrivateSlotFilterer) WatchCollectedSignatures(opts *bind.WatchOpts, sink chan<- *PrivateSlotCollectedSignatures) (event.Subscription, error) {

	logs, sub, err := _PrivateSlot.contract.WatchLogs(opts, "CollectedSignatures")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateSlotCollectedSignatures)
				if err := _PrivateSlot.contract.UnpackLog(event, "CollectedSignatures", log); err != nil {
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

// PrivateSlotConsumeIterator is returned from FilterConsume and is used to iterate over the raw logs and unpacked data for Consume events raised by the PrivateSlot contract.
type PrivateSlotConsumeIterator struct {
	Event *PrivateSlotConsume // Event containing the contract specifics and raw log

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
func (it *PrivateSlotConsumeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateSlotConsume)
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
		it.Event = new(PrivateSlotConsume)
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
func (it *PrivateSlotConsumeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateSlotConsumeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateSlotConsume represents a Consume event raised by the PrivateSlot contract.
type PrivateSlotConsume struct {
	Machine common.Address
	Player  common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterConsume is a free log retrieval operation binding the contract event 0x073d7524ea37f2bed5915b0eca0c75d97c80044d12f5d5b2ecc1bc88a685780f.
//
// Solidity: e Consume(machine indexed address, player indexed address, value uint256)
func (_PrivateSlot *PrivateSlotFilterer) FilterConsume(opts *bind.FilterOpts, machine []common.Address, player []common.Address) (*PrivateSlotConsumeIterator, error) {

	var machineRule []interface{}
	for _, machineItem := range machine {
		machineRule = append(machineRule, machineItem)
	}
	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _PrivateSlot.contract.FilterLogs(opts, "Consume", machineRule, playerRule)
	if err != nil {
		return nil, err
	}
	return &PrivateSlotConsumeIterator{contract: _PrivateSlot.contract, event: "Consume", logs: logs, sub: sub}, nil
}

// WatchConsume is a free log subscription operation binding the contract event 0x073d7524ea37f2bed5915b0eca0c75d97c80044d12f5d5b2ecc1bc88a685780f.
//
// Solidity: e Consume(machine indexed address, player indexed address, value uint256)
func (_PrivateSlot *PrivateSlotFilterer) WatchConsume(opts *bind.WatchOpts, sink chan<- *PrivateSlotConsume, machine []common.Address, player []common.Address) (event.Subscription, error) {

	var machineRule []interface{}
	for _, machineItem := range machine {
		machineRule = append(machineRule, machineItem)
	}
	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _PrivateSlot.contract.WatchLogs(opts, "Consume", machineRule, playerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateSlotConsume)
				if err := _PrivateSlot.contract.UnpackLog(event, "Consume", log); err != nil {
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

// PrivateSlotDepositConfirmationIterator is returned from FilterDepositConfirmation and is used to iterate over the raw logs and unpacked data for DepositConfirmation events raised by the PrivateSlot contract.
type PrivateSlotDepositConfirmationIterator struct {
	Event *PrivateSlotDepositConfirmation // Event containing the contract specifics and raw log

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
func (it *PrivateSlotDepositConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateSlotDepositConfirmation)
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
		it.Event = new(PrivateSlotDepositConfirmation)
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
func (it *PrivateSlotDepositConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateSlotDepositConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateSlotDepositConfirmation represents a DepositConfirmation event raised by the PrivateSlot contract.
type PrivateSlotDepositConfirmation struct {
	Recipient       common.Address
	Value           *big.Int
	TransactoinHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDepositConfirmation is a free log retrieval operation binding the contract event 0x82e9885b59946d922664d6e9c439efafebc983ce46aede53b1916dbb897c137a.
//
// Solidity: e DepositConfirmation(recipient address, value uint256, transactoinHash bytes32)
func (_PrivateSlot *PrivateSlotFilterer) FilterDepositConfirmation(opts *bind.FilterOpts) (*PrivateSlotDepositConfirmationIterator, error) {

	logs, sub, err := _PrivateSlot.contract.FilterLogs(opts, "DepositConfirmation")
	if err != nil {
		return nil, err
	}
	return &PrivateSlotDepositConfirmationIterator{contract: _PrivateSlot.contract, event: "DepositConfirmation", logs: logs, sub: sub}, nil
}

// WatchDepositConfirmation is a free log subscription operation binding the contract event 0x82e9885b59946d922664d6e9c439efafebc983ce46aede53b1916dbb897c137a.
//
// Solidity: e DepositConfirmation(recipient address, value uint256, transactoinHash bytes32)
func (_PrivateSlot *PrivateSlotFilterer) WatchDepositConfirmation(opts *bind.WatchOpts, sink chan<- *PrivateSlotDepositConfirmation) (event.Subscription, error) {

	logs, sub, err := _PrivateSlot.contract.WatchLogs(opts, "DepositConfirmation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateSlotDepositConfirmation)
				if err := _PrivateSlot.contract.UnpackLog(event, "DepositConfirmation", log); err != nil {
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

// PrivateSlotExchangeIterator is returned from FilterExchange and is used to iterate over the raw logs and unpacked data for Exchange events raised by the PrivateSlot contract.
type PrivateSlotExchangeIterator struct {
	Event *PrivateSlotExchange // Event containing the contract specifics and raw log

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
func (it *PrivateSlotExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateSlotExchange)
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
		it.Event = new(PrivateSlotExchange)
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
func (it *PrivateSlotExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateSlotExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateSlotExchange represents a Exchange event raised by the PrivateSlot contract.
type PrivateSlotExchange struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExchange is a free log retrieval operation binding the contract event 0x5988e4c12f4844b895de0739f562558435dca9602fd8b970720ee3cf8dff39be.
//
// Solidity: e Exchange(user address, amount uint256)
func (_PrivateSlot *PrivateSlotFilterer) FilterExchange(opts *bind.FilterOpts) (*PrivateSlotExchangeIterator, error) {

	logs, sub, err := _PrivateSlot.contract.FilterLogs(opts, "Exchange")
	if err != nil {
		return nil, err
	}
	return &PrivateSlotExchangeIterator{contract: _PrivateSlot.contract, event: "Exchange", logs: logs, sub: sub}, nil
}

// WatchExchange is a free log subscription operation binding the contract event 0x5988e4c12f4844b895de0739f562558435dca9602fd8b970720ee3cf8dff39be.
//
// Solidity: e Exchange(user address, amount uint256)
func (_PrivateSlot *PrivateSlotFilterer) WatchExchange(opts *bind.WatchOpts, sink chan<- *PrivateSlotExchange) (event.Subscription, error) {

	logs, sub, err := _PrivateSlot.contract.WatchLogs(opts, "Exchange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateSlotExchange)
				if err := _PrivateSlot.contract.UnpackLog(event, "Exchange", log); err != nil {
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

// PrivateSlotExchangeNFTIterator is returned from FilterExchangeNFT and is used to iterate over the raw logs and unpacked data for ExchangeNFT events raised by the PrivateSlot contract.
type PrivateSlotExchangeNFTIterator struct {
	Event *PrivateSlotExchangeNFT // Event containing the contract specifics and raw log

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
func (it *PrivateSlotExchangeNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateSlotExchangeNFT)
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
		it.Event = new(PrivateSlotExchangeNFT)
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
func (it *PrivateSlotExchangeNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateSlotExchangeNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateSlotExchangeNFT represents a ExchangeNFT event raised by the PrivateSlot contract.
type PrivateSlotExchangeNFT struct {
	TokenID     *big.Int
	Owner       common.Address
	Gene        *big.Int
	AvatarLevel *big.Int
	Weaponed    bool
	Armored     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExchangeNFT is a free log retrieval operation binding the contract event 0x2a797bc85caee23a3e29a6e8fb0614b4ffa925d434e247936693cc55389d7f87.
//
// Solidity: e ExchangeNFT(tokenID uint256, owner address, gene uint256, avatarLevel uint256, weaponed bool, armored bool)
func (_PrivateSlot *PrivateSlotFilterer) FilterExchangeNFT(opts *bind.FilterOpts) (*PrivateSlotExchangeNFTIterator, error) {

	logs, sub, err := _PrivateSlot.contract.FilterLogs(opts, "ExchangeNFT")
	if err != nil {
		return nil, err
	}
	return &PrivateSlotExchangeNFTIterator{contract: _PrivateSlot.contract, event: "ExchangeNFT", logs: logs, sub: sub}, nil
}

// WatchExchangeNFT is a free log subscription operation binding the contract event 0x2a797bc85caee23a3e29a6e8fb0614b4ffa925d434e247936693cc55389d7f87.
//
// Solidity: e ExchangeNFT(tokenID uint256, owner address, gene uint256, avatarLevel uint256, weaponed bool, armored bool)
func (_PrivateSlot *PrivateSlotFilterer) WatchExchangeNFT(opts *bind.WatchOpts, sink chan<- *PrivateSlotExchangeNFT) (event.Subscription, error) {

	logs, sub, err := _PrivateSlot.contract.WatchLogs(opts, "ExchangeNFT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateSlotExchangeNFT)
				if err := _PrivateSlot.contract.UnpackLog(event, "ExchangeNFT", log); err != nil {
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

// PrivateSlotPayIterator is returned from FilterPay and is used to iterate over the raw logs and unpacked data for Pay events raised by the PrivateSlot contract.
type PrivateSlotPayIterator struct {
	Event *PrivateSlotPay // Event containing the contract specifics and raw log

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
func (it *PrivateSlotPayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateSlotPay)
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
		it.Event = new(PrivateSlotPay)
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
func (it *PrivateSlotPayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateSlotPayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateSlotPay represents a Pay event raised by the PrivateSlot contract.
type PrivateSlotPay struct {
	User            common.Address
	Amount          *big.Int
	TransactoinHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPay is a free log retrieval operation binding the contract event 0x1b174056799bea141540e324bb093eb297a02b564c15e75840a30cf0d0f48377.
//
// Solidity: e Pay(user indexed address, amount uint256, transactoinHash bytes32)
func (_PrivateSlot *PrivateSlotFilterer) FilterPay(opts *bind.FilterOpts, user []common.Address) (*PrivateSlotPayIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PrivateSlot.contract.FilterLogs(opts, "Pay", userRule)
	if err != nil {
		return nil, err
	}
	return &PrivateSlotPayIterator{contract: _PrivateSlot.contract, event: "Pay", logs: logs, sub: sub}, nil
}

// WatchPay is a free log subscription operation binding the contract event 0x1b174056799bea141540e324bb093eb297a02b564c15e75840a30cf0d0f48377.
//
// Solidity: e Pay(user indexed address, amount uint256, transactoinHash bytes32)
func (_PrivateSlot *PrivateSlotFilterer) WatchPay(opts *bind.WatchOpts, sink chan<- *PrivateSlotPay, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PrivateSlot.contract.WatchLogs(opts, "Pay", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateSlotPay)
				if err := _PrivateSlot.contract.UnpackLog(event, "Pay", log); err != nil {
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

// PrivateSlotRewardIterator is returned from FilterReward and is used to iterate over the raw logs and unpacked data for Reward events raised by the PrivateSlot contract.
type PrivateSlotRewardIterator struct {
	Event *PrivateSlotReward // Event containing the contract specifics and raw log

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
func (it *PrivateSlotRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateSlotReward)
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
		it.Event = new(PrivateSlotReward)
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
func (it *PrivateSlotRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateSlotRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateSlotReward represents a Reward event raised by the PrivateSlot contract.
type PrivateSlotReward struct {
	Machine common.Address
	Player  common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReward is a free log retrieval operation binding the contract event 0x6b053894d8fdbdcc936dd753e21291f0c48e68ef12306eb39a63a374147ba4bd.
//
// Solidity: e Reward(machine indexed address, player indexed address, value uint256)
func (_PrivateSlot *PrivateSlotFilterer) FilterReward(opts *bind.FilterOpts, machine []common.Address, player []common.Address) (*PrivateSlotRewardIterator, error) {

	var machineRule []interface{}
	for _, machineItem := range machine {
		machineRule = append(machineRule, machineItem)
	}
	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _PrivateSlot.contract.FilterLogs(opts, "Reward", machineRule, playerRule)
	if err != nil {
		return nil, err
	}
	return &PrivateSlotRewardIterator{contract: _PrivateSlot.contract, event: "Reward", logs: logs, sub: sub}, nil
}

// WatchReward is a free log subscription operation binding the contract event 0x6b053894d8fdbdcc936dd753e21291f0c48e68ef12306eb39a63a374147ba4bd.
//
// Solidity: e Reward(machine indexed address, player indexed address, value uint256)
func (_PrivateSlot *PrivateSlotFilterer) WatchReward(opts *bind.WatchOpts, sink chan<- *PrivateSlotReward, machine []common.Address, player []common.Address) (event.Subscription, error) {

	var machineRule []interface{}
	for _, machineItem := range machine {
		machineRule = append(machineRule, machineItem)
	}
	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _PrivateSlot.contract.WatchLogs(opts, "Reward", machineRule, playerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateSlotReward)
				if err := _PrivateSlot.contract.UnpackLog(event, "Reward", log); err != nil {
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

// PrivateSlotWithdrawSignatureSubmittedIterator is returned from FilterWithdrawSignatureSubmitted and is used to iterate over the raw logs and unpacked data for WithdrawSignatureSubmitted events raised by the PrivateSlot contract.
type PrivateSlotWithdrawSignatureSubmittedIterator struct {
	Event *PrivateSlotWithdrawSignatureSubmitted // Event containing the contract specifics and raw log

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
func (it *PrivateSlotWithdrawSignatureSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateSlotWithdrawSignatureSubmitted)
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
		it.Event = new(PrivateSlotWithdrawSignatureSubmitted)
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
func (it *PrivateSlotWithdrawSignatureSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateSlotWithdrawSignatureSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateSlotWithdrawSignatureSubmitted represents a WithdrawSignatureSubmitted event raised by the PrivateSlot contract.
type PrivateSlotWithdrawSignatureSubmitted struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawSignatureSubmitted is a free log retrieval operation binding the contract event 0xfc962d5f0de9bd4b90181a7ec60354bcefe1e162eaf760919b07496a184d665a.
//
// Solidity: e WithdrawSignatureSubmitted(messageHash bytes32)
func (_PrivateSlot *PrivateSlotFilterer) FilterWithdrawSignatureSubmitted(opts *bind.FilterOpts) (*PrivateSlotWithdrawSignatureSubmittedIterator, error) {

	logs, sub, err := _PrivateSlot.contract.FilterLogs(opts, "WithdrawSignatureSubmitted")
	if err != nil {
		return nil, err
	}
	return &PrivateSlotWithdrawSignatureSubmittedIterator{contract: _PrivateSlot.contract, event: "WithdrawSignatureSubmitted", logs: logs, sub: sub}, nil
}

// WatchWithdrawSignatureSubmitted is a free log subscription operation binding the contract event 0xfc962d5f0de9bd4b90181a7ec60354bcefe1e162eaf760919b07496a184d665a.
//
// Solidity: e WithdrawSignatureSubmitted(messageHash bytes32)
func (_PrivateSlot *PrivateSlotFilterer) WatchWithdrawSignatureSubmitted(opts *bind.WatchOpts, sink chan<- *PrivateSlotWithdrawSignatureSubmitted) (event.Subscription, error) {

	logs, sub, err := _PrivateSlot.contract.WatchLogs(opts, "WithdrawSignatureSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateSlotWithdrawSignatureSubmitted)
				if err := _PrivateSlot.contract.UnpackLog(event, "WithdrawSignatureSubmitted", log); err != nil {
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

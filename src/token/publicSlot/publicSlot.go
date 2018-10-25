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
const PublicSlotABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"reward\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"by\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"consume\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"machine\",\"type\":\"address\"}],\"name\":\"removeGameMachine\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ownedAvatars\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenID\",\"type\":\"uint256\"},{\"name\":\"avatarOwner\",\"type\":\"address\"}],\"name\":\"payNFT\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"machine\",\"type\":\"address\"}],\"name\":\"addGameMachine\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"equipArmor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"avatar\",\"outputs\":[{\"name\":\"gene\",\"type\":\"uint256\"},{\"name\":\"avatarLevel\",\"type\":\"uint256\"},{\"name\":\"weaponed\",\"type\":\"bool\"},{\"name\":\"armored\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newRequiredSignatures\",\"type\":\"uint256\"}],\"name\":\"setRequiredSignatures\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requiredSignatures\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"_authorizdedMachines\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vs\",\"type\":\"uint8[]\"},{\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"pay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"equipWeapon\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"exchangeNFT\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"payed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"name\":\"decimalUnits\",\"type\":\"uint8\"},{\"name\":\"_requiredSignatures\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Exchange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"Pay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"gene\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"avatarLevel\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"weaponed\",\"type\":\"bool\"}],\"name\":\"ExchangeNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"machine\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Reward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"machine\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Consume\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// PublicSlotBin is the compiled bytecode used for deploying new contracts.
const PublicSlotBin = `60806040523480156200001157600080fd5b506040516200224e3803806200224e833981018060405281019080805190602001909291908051820192919060200180518201929190602001805190602001909291908051906020019092919050505084848484838383838360008190555083600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508260019080519060200190620000cc9291906200019d565b5080600260006101000a81548160ff021916908360ff1602179055508160039080519060200190620001009291906200019d565b505050505033600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505080600c8190555033600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050506200024c565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001e057805160ff191683800117855562000211565b8280016001018555821562000211579182015b8281111562000210578251825591602001919060010190620001f3565b5b50905062000220919062000224565b5090565b6200024991905b80821115620002455760008160009055506001016200022b565b5090565b90565b611ff2806200025c6000396000f30060806040526004361061016a576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063045d03891461016f57806306fdde03146101bc578063095ea7b31461024c57806318160ddd146102b157806321670f22146102dc578063224b5c721461034157806323b872dd146103a65780633016e1081461042b578063313ce5671461046e578063399bd4f01461049f57806340c10f19146104f657806345977d0314610543578063507023a21461057057806350dc43a1146105bd578063541486b6146106005780635e5741f51461064d5780636352211e146106ab57806370a08231146107185780637d2b9cc01461076f5780638d0680431461079c57806390883c08146107c757806395d89b4114610822578063a835ba0b146108b2578063a9059cbb146109e4578063bf6feb4014610a49578063dd62ed3e14610a96578063e9565b8314610b0d578063ff88aeb914610b3a575b600080fd5b34801561017b57600080fd5b506101ba600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610b83565b005b3480156101c857600080fd5b506101d1610c1f565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102115780820151818401526020810190506101f6565b50505050905090810190601f16801561023e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561025857600080fd5b50610297600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610cbd565b604051808215151515815260200191505060405180910390f35b3480156102bd57600080fd5b506102c6610dea565b6040518082815260200191505060405180910390f35b3480156102e857600080fd5b50610327600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610df3565b604051808215151515815260200191505060405180910390f35b34801561034d57600080fd5b5061038c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610ec7565b604051808215151515815260200191505060405180910390f35b3480156103b257600080fd5b50610411600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610f9b565b604051808215151515815260200191505060405180910390f35b34801561043757600080fd5b5061046c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506111b2565b005b34801561047a57600080fd5b50610483611269565b604051808260ff1660ff16815260200191505060405180910390f35b3480156104ab57600080fd5b506104e0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061127c565b6040518082815260200191505060405180910390f35b34801561050257600080fd5b50610541600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506112c5565b005b34801561054f57600080fd5b5061056e600480360381019080803590602001909291905050506113c2565b005b34801561057c57600080fd5b506105bb60048036038101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061140f565b005b3480156105c957600080fd5b506105fe600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506114a9565b005b34801561060c57600080fd5b5061064b60048036038101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611560565b005b34801561065957600080fd5b5061067860048036038101908080359060200190929190505050611630565b60405180858152602001848152602001831515151581526020018215151515815260200194505050505060405180910390f35b3480156106b757600080fd5b506106d66004803603810190808035906020019092919050505061167a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561072457600080fd5b50610759600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506116f8565b6040518082815260200191505060405180910390f35b34801561077b57600080fd5b5061079a60048036038101908080359060200190929190505050611741565b005b3480156107a857600080fd5b506107b16117a7565b6040518082815260200191505060405180910390f35b3480156107d357600080fd5b50610808600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506117ad565b604051808215151515815260200191505060405180910390f35b34801561082e57600080fd5b506108376117cd565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561087757808201518184015260208101905061085c565b50505050905090810190601f1680156108a45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156108be57600080fd5b506109e2600480360381019080803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505091929192908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919291929080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061186b565b005b3480156109f057600080fd5b50610a2f600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506119bb565b604051808215151515815260200191505060405180910390f35b348015610a5557600080fd5b50610a9460048036038101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611a37565b005b348015610aa257600080fd5b50610af7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611b07565b6040518082815260200191505060405180910390f35b348015610b1957600080fd5b50610b3860048036038101908080359060200190929190505050611b8e565b005b348015610b4657600080fd5b50610b696004803603810190808035600019169060200190929190505050611d72565b604051808215151515815260200191505060405180910390f35b610bb082600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683611d92565b7f5988e4c12f4844b895de0739f562558435dca9602fd8b970720ee3cf8dff39be8282604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a15050565b60018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610cb55780601f10610c8a57610100808354040283529160200191610cb5565b820191906000526020600020905b815481529060010190602001808311610c9857829003601f168201915b505050505081565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614151515610cfa57600080fd5b81600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b60008054905090565b6000600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610e4d57600080fd5b610e58338484611d92565b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f6b053894d8fdbdcc936dd753e21291f0c48e68ef12306eb39a63a374147ba4bd846040518082815260200191505060405180910390a36001905092915050565b6000600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610f2157600080fd5b610f2c833384611d92565b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f073d7524ea37f2bed5915b0eca0c75d97c80044d12f5d5b2ecc1bc88a685780f846040518082815260200191505060405180910390a36001905092915050565b6000600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054821115151561102857600080fd5b6110b782600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611f4b90919063ffffffff16565b600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611142848484611d92565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3600190509392505050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561120e57600080fd5b6000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b600260009054906101000a900460ff1681565b6000600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151561130157600080fd5b816008600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506002428115156113a357fe5b0660096000838152602001908152602001600020600001819055505050565b600260096000838152602001908152602001600020600101541015156113e757600080fd5b6001600960008381526020019081526020016000206001016000828254019250508190555050565b81600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550806008600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561150557600080fd5b6001600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b8073ffffffffffffffffffffffffffffffffffffffff166008600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415156115cd57600080fd5b6009600083815260200190815260200160002060020160019054906101000a900460ff161515156115fd57600080fd5b60016009600084815260200190815260200160002060020160016101000a81548160ff0219169083151502179055505050565b60096020528060005260406000206000915090508060000154908060010154908060020160009054906101000a900460ff16908060020160019054906101000a900460ff16905084565b6000806008600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156116ef57600080fd5b80915050919050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561179d57600080fd5b80600c8190555050565b600c5481565b60076020528060005260406000206000915054906101000a900460ff1681565b60038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156118635780601f1061183857610100808354040283529160200191611863565b820191906000526020600020905b81548152906001019060200180831161184657829003601f168201915b505050505081565b60008060006054845114151561188057600080fd5b61188984611f6c565b925061189484611f7f565b915061189f84611f92565b9050600b6000826000191660001916815260200190815260200160002060009054906101000a900460ff161515156118d657600080fd5b6001600b6000836000191660001916815260200190815260200160002060006101000a81548160ff021916908315150217905550611937600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168484611d92565b7f1b174056799bea141540e324bb093eb297a02b564c15e75840a30cf0d0f48377838383604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018260001916600019168152602001935050505060405180910390a150505050505050565b60006119c8338484611d92565b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a36001905092915050565b8073ffffffffffffffffffffffffffffffffffffffff166008600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515611aa457600080fd5b6009600083815260200190815260200160002060020160009054906101000a900460ff16151515611ad457600080fd5b60016009600084815260200190815260200160002060020160006101000a81548160ff0219169083151502179055505050565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b60006008600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611c0057600080fd5b6000600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060006008600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f610a7660a1a3dfb5b4577a836313c496b3992c167b445d7ef51d178da791e9188282600960008681526020019081526020016000206000015460096000878152602001908152602001600020600101546009600088815260200190815260200160002060020160009054906101000a900460ff16604051808681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001848152602001838152602001821515151581526020019550505050505060405180910390a15050565b600b6020528060005260406000206000915054906101000a900460ff1681565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548111151515611de057600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515611e1c57600080fd5b611e6e81600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611f4b90919063ffffffff16565b600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611f0381600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611fa590919063ffffffff16565b600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505050565b600080838311151515611f5d57600080fd5b82840390508091505092915050565b6000806034830151905080915050919050565b6000806054830151905080915050919050565b6000806020830151905080915050919050565b6000808284019050838110151515611fbc57600080fd5b80915050929150505600a165627a7a72305820d2901357b8aa3cb3db1f038ba3c8f4f5064d787df8fe97a85a8a504fe8614c780029`

// DeployPublicSlot deploys a new Ethereum contract, binding an instance of PublicSlot to it.
func DeployPublicSlot(auth *bind.TransactOpts, backend bind.ContractBackend, totalSupply *big.Int, tokenName string, tokenSymbol string, decimalUnits uint8, _requiredSignatures *big.Int) (common.Address, *types.Transaction, *PublicSlot, error) {
	parsed, err := abi.JSON(strings.NewReader(PublicSlotABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PublicSlotBin), backend, totalSupply, tokenName, tokenSymbol, decimalUnits, _requiredSignatures)
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

// AuthorizdedMachines is a free data retrieval call binding the contract method 0x90883c08.
//
// Solidity: function _authorizdedMachines( address) constant returns(bool)
func (_PublicSlot *PublicSlotCaller) AuthorizdedMachines(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PublicSlot.contract.Call(opts, out, "_authorizdedMachines", arg0)
	return *ret0, err
}

// AuthorizdedMachines is a free data retrieval call binding the contract method 0x90883c08.
//
// Solidity: function _authorizdedMachines( address) constant returns(bool)
func (_PublicSlot *PublicSlotSession) AuthorizdedMachines(arg0 common.Address) (bool, error) {
	return _PublicSlot.Contract.AuthorizdedMachines(&_PublicSlot.CallOpts, arg0)
}

// AuthorizdedMachines is a free data retrieval call binding the contract method 0x90883c08.
//
// Solidity: function _authorizdedMachines( address) constant returns(bool)
func (_PublicSlot *PublicSlotCallerSession) AuthorizdedMachines(arg0 common.Address) (bool, error) {
	return _PublicSlot.Contract.AuthorizdedMachines(&_PublicSlot.CallOpts, arg0)
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

// Avatar is a free data retrieval call binding the contract method 0x5e5741f5.
//
// Solidity: function avatar( uint256) constant returns(gene uint256, avatarLevel uint256, weaponed bool, armored bool)
func (_PublicSlot *PublicSlotCaller) Avatar(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _PublicSlot.contract.Call(opts, out, "avatar", arg0)
	return *ret, err
}

// Avatar is a free data retrieval call binding the contract method 0x5e5741f5.
//
// Solidity: function avatar( uint256) constant returns(gene uint256, avatarLevel uint256, weaponed bool, armored bool)
func (_PublicSlot *PublicSlotSession) Avatar(arg0 *big.Int) (struct {
	Gene        *big.Int
	AvatarLevel *big.Int
	Weaponed    bool
	Armored     bool
}, error) {
	return _PublicSlot.Contract.Avatar(&_PublicSlot.CallOpts, arg0)
}

// Avatar is a free data retrieval call binding the contract method 0x5e5741f5.
//
// Solidity: function avatar( uint256) constant returns(gene uint256, avatarLevel uint256, weaponed bool, armored bool)
func (_PublicSlot *PublicSlotCallerSession) Avatar(arg0 *big.Int) (struct {
	Gene        *big.Int
	AvatarLevel *big.Int
	Weaponed    bool
	Armored     bool
}, error) {
	return _PublicSlot.Contract.Avatar(&_PublicSlot.CallOpts, arg0)
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

// OwnedAvatars is a free data retrieval call binding the contract method 0x399bd4f0.
//
// Solidity: function ownedAvatars(owner address) constant returns(uint256)
func (_PublicSlot *PublicSlotCaller) OwnedAvatars(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PublicSlot.contract.Call(opts, out, "ownedAvatars", owner)
	return *ret0, err
}

// OwnedAvatars is a free data retrieval call binding the contract method 0x399bd4f0.
//
// Solidity: function ownedAvatars(owner address) constant returns(uint256)
func (_PublicSlot *PublicSlotSession) OwnedAvatars(owner common.Address) (*big.Int, error) {
	return _PublicSlot.Contract.OwnedAvatars(&_PublicSlot.CallOpts, owner)
}

// OwnedAvatars is a free data retrieval call binding the contract method 0x399bd4f0.
//
// Solidity: function ownedAvatars(owner address) constant returns(uint256)
func (_PublicSlot *PublicSlotCallerSession) OwnedAvatars(owner common.Address) (*big.Int, error) {
	return _PublicSlot.Contract.OwnedAvatars(&_PublicSlot.CallOpts, owner)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(tokenId uint256) constant returns(address)
func (_PublicSlot *PublicSlotCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PublicSlot.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(tokenId uint256) constant returns(address)
func (_PublicSlot *PublicSlotSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _PublicSlot.Contract.OwnerOf(&_PublicSlot.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(tokenId uint256) constant returns(address)
func (_PublicSlot *PublicSlotCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _PublicSlot.Contract.OwnerOf(&_PublicSlot.CallOpts, tokenId)
}

// Payed is a free data retrieval call binding the contract method 0xff88aeb9.
//
// Solidity: function payed( bytes32) constant returns(bool)
func (_PublicSlot *PublicSlotCaller) Payed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PublicSlot.contract.Call(opts, out, "payed", arg0)
	return *ret0, err
}

// Payed is a free data retrieval call binding the contract method 0xff88aeb9.
//
// Solidity: function payed( bytes32) constant returns(bool)
func (_PublicSlot *PublicSlotSession) Payed(arg0 [32]byte) (bool, error) {
	return _PublicSlot.Contract.Payed(&_PublicSlot.CallOpts, arg0)
}

// Payed is a free data retrieval call binding the contract method 0xff88aeb9.
//
// Solidity: function payed( bytes32) constant returns(bool)
func (_PublicSlot *PublicSlotCallerSession) Payed(arg0 [32]byte) (bool, error) {
	return _PublicSlot.Contract.Payed(&_PublicSlot.CallOpts, arg0)
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

// AddGameMachine is a paid mutator transaction binding the contract method 0x50dc43a1.
//
// Solidity: function addGameMachine(machine address) returns()
func (_PublicSlot *PublicSlotTransactor) AddGameMachine(opts *bind.TransactOpts, machine common.Address) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "addGameMachine", machine)
}

// AddGameMachine is a paid mutator transaction binding the contract method 0x50dc43a1.
//
// Solidity: function addGameMachine(machine address) returns()
func (_PublicSlot *PublicSlotSession) AddGameMachine(machine common.Address) (*types.Transaction, error) {
	return _PublicSlot.Contract.AddGameMachine(&_PublicSlot.TransactOpts, machine)
}

// AddGameMachine is a paid mutator transaction binding the contract method 0x50dc43a1.
//
// Solidity: function addGameMachine(machine address) returns()
func (_PublicSlot *PublicSlotTransactorSession) AddGameMachine(machine common.Address) (*types.Transaction, error) {
	return _PublicSlot.Contract.AddGameMachine(&_PublicSlot.TransactOpts, machine)
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

// EquipArmor is a paid mutator transaction binding the contract method 0x541486b6.
//
// Solidity: function equipArmor(tokenId uint256, user address) returns()
func (_PublicSlot *PublicSlotTransactor) EquipArmor(opts *bind.TransactOpts, tokenId *big.Int, user common.Address) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "equipArmor", tokenId, user)
}

// EquipArmor is a paid mutator transaction binding the contract method 0x541486b6.
//
// Solidity: function equipArmor(tokenId uint256, user address) returns()
func (_PublicSlot *PublicSlotSession) EquipArmor(tokenId *big.Int, user common.Address) (*types.Transaction, error) {
	return _PublicSlot.Contract.EquipArmor(&_PublicSlot.TransactOpts, tokenId, user)
}

// EquipArmor is a paid mutator transaction binding the contract method 0x541486b6.
//
// Solidity: function equipArmor(tokenId uint256, user address) returns()
func (_PublicSlot *PublicSlotTransactorSession) EquipArmor(tokenId *big.Int, user common.Address) (*types.Transaction, error) {
	return _PublicSlot.Contract.EquipArmor(&_PublicSlot.TransactOpts, tokenId, user)
}

// EquipWeapon is a paid mutator transaction binding the contract method 0xbf6feb40.
//
// Solidity: function equipWeapon(tokenId uint256, user address) returns()
func (_PublicSlot *PublicSlotTransactor) EquipWeapon(opts *bind.TransactOpts, tokenId *big.Int, user common.Address) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "equipWeapon", tokenId, user)
}

// EquipWeapon is a paid mutator transaction binding the contract method 0xbf6feb40.
//
// Solidity: function equipWeapon(tokenId uint256, user address) returns()
func (_PublicSlot *PublicSlotSession) EquipWeapon(tokenId *big.Int, user common.Address) (*types.Transaction, error) {
	return _PublicSlot.Contract.EquipWeapon(&_PublicSlot.TransactOpts, tokenId, user)
}

// EquipWeapon is a paid mutator transaction binding the contract method 0xbf6feb40.
//
// Solidity: function equipWeapon(tokenId uint256, user address) returns()
func (_PublicSlot *PublicSlotTransactorSession) EquipWeapon(tokenId *big.Int, user common.Address) (*types.Transaction, error) {
	return _PublicSlot.Contract.EquipWeapon(&_PublicSlot.TransactOpts, tokenId, user)
}

// Exchange is a paid mutator transaction binding the contract method 0x045d0389.
//
// Solidity: function exchange(user address, amount uint256) returns()
func (_PublicSlot *PublicSlotTransactor) Exchange(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "exchange", user, amount)
}

// Exchange is a paid mutator transaction binding the contract method 0x045d0389.
//
// Solidity: function exchange(user address, amount uint256) returns()
func (_PublicSlot *PublicSlotSession) Exchange(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Exchange(&_PublicSlot.TransactOpts, user, amount)
}

// Exchange is a paid mutator transaction binding the contract method 0x045d0389.
//
// Solidity: function exchange(user address, amount uint256) returns()
func (_PublicSlot *PublicSlotTransactorSession) Exchange(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Exchange(&_PublicSlot.TransactOpts, user, amount)
}

// ExchangeNFT is a paid mutator transaction binding the contract method 0xe9565b83.
//
// Solidity: function exchangeNFT(tokenID uint256) returns()
func (_PublicSlot *PublicSlotTransactor) ExchangeNFT(opts *bind.TransactOpts, tokenID *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "exchangeNFT", tokenID)
}

// ExchangeNFT is a paid mutator transaction binding the contract method 0xe9565b83.
//
// Solidity: function exchangeNFT(tokenID uint256) returns()
func (_PublicSlot *PublicSlotSession) ExchangeNFT(tokenID *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.ExchangeNFT(&_PublicSlot.TransactOpts, tokenID)
}

// ExchangeNFT is a paid mutator transaction binding the contract method 0xe9565b83.
//
// Solidity: function exchangeNFT(tokenID uint256) returns()
func (_PublicSlot *PublicSlotTransactorSession) ExchangeNFT(tokenID *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.ExchangeNFT(&_PublicSlot.TransactOpts, tokenID)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(to address, tokenId uint256) returns()
func (_PublicSlot *PublicSlotTransactor) Mint(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "mint", to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(to address, tokenId uint256) returns()
func (_PublicSlot *PublicSlotSession) Mint(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Mint(&_PublicSlot.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(to address, tokenId uint256) returns()
func (_PublicSlot *PublicSlotTransactorSession) Mint(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Mint(&_PublicSlot.TransactOpts, to, tokenId)
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

// PayNFT is a paid mutator transaction binding the contract method 0x507023a2.
//
// Solidity: function payNFT(tokenID uint256, avatarOwner address) returns()
func (_PublicSlot *PublicSlotTransactor) PayNFT(opts *bind.TransactOpts, tokenID *big.Int, avatarOwner common.Address) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "payNFT", tokenID, avatarOwner)
}

// PayNFT is a paid mutator transaction binding the contract method 0x507023a2.
//
// Solidity: function payNFT(tokenID uint256, avatarOwner address) returns()
func (_PublicSlot *PublicSlotSession) PayNFT(tokenID *big.Int, avatarOwner common.Address) (*types.Transaction, error) {
	return _PublicSlot.Contract.PayNFT(&_PublicSlot.TransactOpts, tokenID, avatarOwner)
}

// PayNFT is a paid mutator transaction binding the contract method 0x507023a2.
//
// Solidity: function payNFT(tokenID uint256, avatarOwner address) returns()
func (_PublicSlot *PublicSlotTransactorSession) PayNFT(tokenID *big.Int, avatarOwner common.Address) (*types.Transaction, error) {
	return _PublicSlot.Contract.PayNFT(&_PublicSlot.TransactOpts, tokenID, avatarOwner)
}

// RemoveGameMachine is a paid mutator transaction binding the contract method 0x3016e108.
//
// Solidity: function removeGameMachine(machine address) returns()
func (_PublicSlot *PublicSlotTransactor) RemoveGameMachine(opts *bind.TransactOpts, machine common.Address) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "removeGameMachine", machine)
}

// RemoveGameMachine is a paid mutator transaction binding the contract method 0x3016e108.
//
// Solidity: function removeGameMachine(machine address) returns()
func (_PublicSlot *PublicSlotSession) RemoveGameMachine(machine common.Address) (*types.Transaction, error) {
	return _PublicSlot.Contract.RemoveGameMachine(&_PublicSlot.TransactOpts, machine)
}

// RemoveGameMachine is a paid mutator transaction binding the contract method 0x3016e108.
//
// Solidity: function removeGameMachine(machine address) returns()
func (_PublicSlot *PublicSlotTransactorSession) RemoveGameMachine(machine common.Address) (*types.Transaction, error) {
	return _PublicSlot.Contract.RemoveGameMachine(&_PublicSlot.TransactOpts, machine)
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

// SetRequiredSignatures is a paid mutator transaction binding the contract method 0x7d2b9cc0.
//
// Solidity: function setRequiredSignatures(newRequiredSignatures uint256) returns()
func (_PublicSlot *PublicSlotTransactor) SetRequiredSignatures(opts *bind.TransactOpts, newRequiredSignatures *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "setRequiredSignatures", newRequiredSignatures)
}

// SetRequiredSignatures is a paid mutator transaction binding the contract method 0x7d2b9cc0.
//
// Solidity: function setRequiredSignatures(newRequiredSignatures uint256) returns()
func (_PublicSlot *PublicSlotSession) SetRequiredSignatures(newRequiredSignatures *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.SetRequiredSignatures(&_PublicSlot.TransactOpts, newRequiredSignatures)
}

// SetRequiredSignatures is a paid mutator transaction binding the contract method 0x7d2b9cc0.
//
// Solidity: function setRequiredSignatures(newRequiredSignatures uint256) returns()
func (_PublicSlot *PublicSlotTransactorSession) SetRequiredSignatures(newRequiredSignatures *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.SetRequiredSignatures(&_PublicSlot.TransactOpts, newRequiredSignatures)
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

// Upgrade is a paid mutator transaction binding the contract method 0x45977d03.
//
// Solidity: function upgrade(tokenId uint256) returns()
func (_PublicSlot *PublicSlotTransactor) Upgrade(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _PublicSlot.contract.Transact(opts, "upgrade", tokenId)
}

// Upgrade is a paid mutator transaction binding the contract method 0x45977d03.
//
// Solidity: function upgrade(tokenId uint256) returns()
func (_PublicSlot *PublicSlotSession) Upgrade(tokenId *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Upgrade(&_PublicSlot.TransactOpts, tokenId)
}

// Upgrade is a paid mutator transaction binding the contract method 0x45977d03.
//
// Solidity: function upgrade(tokenId uint256) returns()
func (_PublicSlot *PublicSlotTransactorSession) Upgrade(tokenId *big.Int) (*types.Transaction, error) {
	return _PublicSlot.Contract.Upgrade(&_PublicSlot.TransactOpts, tokenId)
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

// PublicSlotConsumeIterator is returned from FilterConsume and is used to iterate over the raw logs and unpacked data for Consume events raised by the PublicSlot contract.
type PublicSlotConsumeIterator struct {
	Event *PublicSlotConsume // Event containing the contract specifics and raw log

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
func (it *PublicSlotConsumeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicSlotConsume)
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
		it.Event = new(PublicSlotConsume)
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
func (it *PublicSlotConsumeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicSlotConsumeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicSlotConsume represents a Consume event raised by the PublicSlot contract.
type PublicSlotConsume struct {
	Machine common.Address
	Player  common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterConsume is a free log retrieval operation binding the contract event 0x073d7524ea37f2bed5915b0eca0c75d97c80044d12f5d5b2ecc1bc88a685780f.
//
// Solidity: e Consume(machine indexed address, player indexed address, value uint256)
func (_PublicSlot *PublicSlotFilterer) FilterConsume(opts *bind.FilterOpts, machine []common.Address, player []common.Address) (*PublicSlotConsumeIterator, error) {

	var machineRule []interface{}
	for _, machineItem := range machine {
		machineRule = append(machineRule, machineItem)
	}
	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _PublicSlot.contract.FilterLogs(opts, "Consume", machineRule, playerRule)
	if err != nil {
		return nil, err
	}
	return &PublicSlotConsumeIterator{contract: _PublicSlot.contract, event: "Consume", logs: logs, sub: sub}, nil
}

// WatchConsume is a free log subscription operation binding the contract event 0x073d7524ea37f2bed5915b0eca0c75d97c80044d12f5d5b2ecc1bc88a685780f.
//
// Solidity: e Consume(machine indexed address, player indexed address, value uint256)
func (_PublicSlot *PublicSlotFilterer) WatchConsume(opts *bind.WatchOpts, sink chan<- *PublicSlotConsume, machine []common.Address, player []common.Address) (event.Subscription, error) {

	var machineRule []interface{}
	for _, machineItem := range machine {
		machineRule = append(machineRule, machineItem)
	}
	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _PublicSlot.contract.WatchLogs(opts, "Consume", machineRule, playerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicSlotConsume)
				if err := _PublicSlot.contract.UnpackLog(event, "Consume", log); err != nil {
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
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExchange is a free log retrieval operation binding the contract event 0x5988e4c12f4844b895de0739f562558435dca9602fd8b970720ee3cf8dff39be.
//
// Solidity: e Exchange(user address, amount uint256)
func (_PublicSlot *PublicSlotFilterer) FilterExchange(opts *bind.FilterOpts) (*PublicSlotExchangeIterator, error) {

	logs, sub, err := _PublicSlot.contract.FilterLogs(opts, "Exchange")
	if err != nil {
		return nil, err
	}
	return &PublicSlotExchangeIterator{contract: _PublicSlot.contract, event: "Exchange", logs: logs, sub: sub}, nil
}

// WatchExchange is a free log subscription operation binding the contract event 0x5988e4c12f4844b895de0739f562558435dca9602fd8b970720ee3cf8dff39be.
//
// Solidity: e Exchange(user address, amount uint256)
func (_PublicSlot *PublicSlotFilterer) WatchExchange(opts *bind.WatchOpts, sink chan<- *PublicSlotExchange) (event.Subscription, error) {

	logs, sub, err := _PublicSlot.contract.WatchLogs(opts, "Exchange")
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

// PublicSlotExchangeNFTIterator is returned from FilterExchangeNFT and is used to iterate over the raw logs and unpacked data for ExchangeNFT events raised by the PublicSlot contract.
type PublicSlotExchangeNFTIterator struct {
	Event *PublicSlotExchangeNFT // Event containing the contract specifics and raw log

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
func (it *PublicSlotExchangeNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicSlotExchangeNFT)
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
		it.Event = new(PublicSlotExchangeNFT)
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
func (it *PublicSlotExchangeNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicSlotExchangeNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicSlotExchangeNFT represents a ExchangeNFT event raised by the PublicSlot contract.
type PublicSlotExchangeNFT struct {
	TokenID     *big.Int
	Owner       common.Address
	Gene        *big.Int
	AvatarLevel *big.Int
	Weaponed    bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExchangeNFT is a free log retrieval operation binding the contract event 0x610a7660a1a3dfb5b4577a836313c496b3992c167b445d7ef51d178da791e918.
//
// Solidity: e ExchangeNFT(tokenID uint256, owner address, gene uint256, avatarLevel uint256, weaponed bool)
func (_PublicSlot *PublicSlotFilterer) FilterExchangeNFT(opts *bind.FilterOpts) (*PublicSlotExchangeNFTIterator, error) {

	logs, sub, err := _PublicSlot.contract.FilterLogs(opts, "ExchangeNFT")
	if err != nil {
		return nil, err
	}
	return &PublicSlotExchangeNFTIterator{contract: _PublicSlot.contract, event: "ExchangeNFT", logs: logs, sub: sub}, nil
}

// WatchExchangeNFT is a free log subscription operation binding the contract event 0x610a7660a1a3dfb5b4577a836313c496b3992c167b445d7ef51d178da791e918.
//
// Solidity: e ExchangeNFT(tokenID uint256, owner address, gene uint256, avatarLevel uint256, weaponed bool)
func (_PublicSlot *PublicSlotFilterer) WatchExchangeNFT(opts *bind.WatchOpts, sink chan<- *PublicSlotExchangeNFT) (event.Subscription, error) {

	logs, sub, err := _PublicSlot.contract.WatchLogs(opts, "ExchangeNFT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicSlotExchangeNFT)
				if err := _PublicSlot.contract.UnpackLog(event, "ExchangeNFT", log); err != nil {
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
	User            common.Address
	Amount          *big.Int
	TransactionHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPay is a free log retrieval operation binding the contract event 0x1b174056799bea141540e324bb093eb297a02b564c15e75840a30cf0d0f48377.
//
// Solidity: e Pay(user address, amount uint256, transactionHash bytes32)
func (_PublicSlot *PublicSlotFilterer) FilterPay(opts *bind.FilterOpts) (*PublicSlotPayIterator, error) {

	logs, sub, err := _PublicSlot.contract.FilterLogs(opts, "Pay")
	if err != nil {
		return nil, err
	}
	return &PublicSlotPayIterator{contract: _PublicSlot.contract, event: "Pay", logs: logs, sub: sub}, nil
}

// WatchPay is a free log subscription operation binding the contract event 0x1b174056799bea141540e324bb093eb297a02b564c15e75840a30cf0d0f48377.
//
// Solidity: e Pay(user address, amount uint256, transactionHash bytes32)
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

// PublicSlotRewardIterator is returned from FilterReward and is used to iterate over the raw logs and unpacked data for Reward events raised by the PublicSlot contract.
type PublicSlotRewardIterator struct {
	Event *PublicSlotReward // Event containing the contract specifics and raw log

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
func (it *PublicSlotRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicSlotReward)
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
		it.Event = new(PublicSlotReward)
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
func (it *PublicSlotRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicSlotRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicSlotReward represents a Reward event raised by the PublicSlot contract.
type PublicSlotReward struct {
	Machine common.Address
	Player  common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReward is a free log retrieval operation binding the contract event 0x6b053894d8fdbdcc936dd753e21291f0c48e68ef12306eb39a63a374147ba4bd.
//
// Solidity: e Reward(machine indexed address, player indexed address, value uint256)
func (_PublicSlot *PublicSlotFilterer) FilterReward(opts *bind.FilterOpts, machine []common.Address, player []common.Address) (*PublicSlotRewardIterator, error) {

	var machineRule []interface{}
	for _, machineItem := range machine {
		machineRule = append(machineRule, machineItem)
	}
	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _PublicSlot.contract.FilterLogs(opts, "Reward", machineRule, playerRule)
	if err != nil {
		return nil, err
	}
	return &PublicSlotRewardIterator{contract: _PublicSlot.contract, event: "Reward", logs: logs, sub: sub}, nil
}

// WatchReward is a free log subscription operation binding the contract event 0x6b053894d8fdbdcc936dd753e21291f0c48e68ef12306eb39a63a374147ba4bd.
//
// Solidity: e Reward(machine indexed address, player indexed address, value uint256)
func (_PublicSlot *PublicSlotFilterer) WatchReward(opts *bind.WatchOpts, sink chan<- *PublicSlotReward, machine []common.Address, player []common.Address) (event.Subscription, error) {

	var machineRule []interface{}
	for _, machineItem := range machine {
		machineRule = append(machineRule, machineItem)
	}
	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _PublicSlot.contract.WatchLogs(opts, "Reward", machineRule, playerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicSlotReward)
				if err := _PublicSlot.contract.UnpackLog(event, "Reward", log); err != nil {
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

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
const PrivateSlotABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_user\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256[]\"},{\"name\":\"_time\",\"type\":\"uint256[]\"}],\"name\":\"sync\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"_success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"signature\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"consume\",\"outputs\":[{\"name\":\"_success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAuthority\",\"type\":\"address\"}],\"name\":\"addAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"message_hash\",\"type\":\"bytes32\"}],\"name\":\"message\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"authorities\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"award\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_transactionHash\",\"type\":\"bytes32\"}],\"name\":\"pay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"submitSignature\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"_balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newRequiredSignatures\",\"type\":\"uint256\"}],\"name\":\"setRequiredSignatures\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requiredSignatures\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRequiredSignatures\",\"outputs\":[{\"name\":\"_requiredSignatures\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"_remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"name\":\"_requiredSignatures\",\"type\":\"uint256\"},{\"name\":\"_authorities\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"WithdrawSignatureSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"authorityResponsibleForRelay\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"CollectedSignatures\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"transactoinHash\",\"type\":\"bytes32\"}],\"name\":\"DepositConfirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"privateAdd\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"publicAdd\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"Login\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address[]\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address[]\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"time\",\"type\":\"uint256[]\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"SyncConfirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"transactoinHash\",\"type\":\"bytes32\"}],\"name\":\"Pay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Award\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Consume\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ExchangeToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// PrivateSlotBin is the compiled bytecode used for deploying new contracts.
const PrivateSlotBin = `60806040526040805190810160405280600481526020017f736c6f7400000000000000000000000000000000000000000000000000000000815250600090805190602001906200005192919062000168565b506040805190810160405280600381526020017f736c740000000000000000000000000000000000000000000000000000000000815250600190805190602001906200009f92919062000168565b5060006002556064600455348015620000b757600080fd5b506040516200234d3803806200234d8339810180604052810190808051906020019092919080519060200190929190805182019291905050508280600381905550600354600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550508160078190555080600890805190602001906200015e929190620001ef565b50505050620002ec565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001ab57805160ff1916838001178555620001dc565b82800160010185558215620001dc579182015b82811115620001db578251825591602001919060010190620001be565b5b509050620001eb91906200027e565b5090565b8280548282559060005260206000209081019282156200026b579160200282015b828111156200026a5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509160200191906001019062000210565b5b5090506200027a9190620002a6565b5090565b620002a391905b808211156200029f57600081600090555060010162000285565b5090565b90565b620002e991905b80821115620002e557600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101620002ad565b5090565b90565b61205180620002fc6000396000f3006080604052600436106100f1576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063059d4135146100f6578063095ea7b3146101e25780631812d99614610247578063224b5c72146102fb57806323b872dd1461036057806326defa73146103e5578063490a32c614610428578063494503d4146104d25780635d8a776e1461053f5780635e5571ac1461058c578063630cea8e146105e757806370a08231146106965780637d2b9cc0146106ed5780638d0680431461071a578063a9059cbb14610745578063ccd93998146107aa578063dd62ed3e146107d5575b600080fd5b34801561010257600080fd5b506101e060048036038101908080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505091929192908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919291929050505061084c565b005b3480156101ee57600080fd5b5061022d600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610abc565b604051808215151515815260200191505060405180910390f35b34801561025357600080fd5b50610280600480360381019080803560001916906020019092919080359060200190929190505050610bae565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102c05780820151818401526020810190506102a5565b50505050905090810190601f1680156102ed5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561030757600080fd5b50610346600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610c86565b604051808215151515815260200191505060405180910390f35b34801561036c57600080fd5b506103cb600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610d7b565b604051808215151515815260200191505060405180910390f35b3480156103f157600080fd5b50610426600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e1c565b005b34801561043457600080fd5b506104576004803603810190808035600019169060200190929190505050610f20565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561049757808201518184015260208101905061047c565b50505050905090810190601f1680156104c45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156104de57600080fd5b506104fd60048036038101908080359060200190929190505050610fe0565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561054b57600080fd5b5061058a600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061101e565b005b34801561059857600080fd5b506105e5600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080356000191690602001909291905050506110fe565b005b3480156105f357600080fd5b50610694600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506113cb565b005b3480156106a257600080fd5b506106d7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506117c0565b6040518082815260200191505060405180910390f35b3480156106f957600080fd5b5061071860048036038101908080359060200190929190505050611809565b005b34801561072657600080fd5b5061072f6118ae565b6040518082815260200191505060405180910390f35b34801561075157600080fd5b50610790600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506118b4565b604051808215151515815260200191505060405180910390f35b3480156107b657600080fd5b506107bf6118c9565b6040518082815260200191505060405180910390f35b3480156107e157600080fd5b50610836600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506118d3565b6040518082815260200191505060405180910390f35b60008084519150600090505b8181101561096257838181518110151561086e57fe5b9060200190602002015160056000878481518110151561088a57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101515156108dc57600080fd5b83818151811015156108ea57fe5b9060200190602002015160056000878481518110151561090657fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540392505081905550806001019050610858565b7f97958b0c39735ab08db9a8f6ceab791cd74b6842a2dd8cb67a7fff23cc72b679858686866040518080602001806020018060200180602001858103855289818151815260200191508051906020019060200280838360005b838110156109d65780820151818401526020810190506109bb565b50505050905001858103845288818151815260200191508051906020019060200280838360005b83811015610a185780820151818401526020810190506109fd565b50505050905001858103835287818151815260200191508051906020019060200280838360005b83811015610a5a578082015181840152602081019050610a3f565b50505050905001858103825286818151815260200191508051906020019060200280838360005b83811015610a9c578082015181840152602081019050610a81565b505050509050019850505050505050505060405180910390a15050505050565b600081600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b6060600a6000846000191660001916815260200190815260200160002060020182815481101515610bdb57fe5b906000526020600020018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c795780601f10610c4e57610100808354040283529160200191610c79565b820191906000526020600020905b815481529060010190602001808311610c5c57829003601f168201915b5050505050905092915050565b600081600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410151515610cd657600080fd5b81600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055508273ffffffffffffffffffffffffffffffffffffffff167fb3762e93ec66871dd27c421b64edc79636345ff0a949cd04f7f8efce5bd4240e836040518082815260200191505060405180910390a26001905092915050565b6000600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548211151515610e0857600080fd5b610e1384848461195a565b90509392505050565b610eac6008805480602002602001604051908101604052809291908181526020018280548015610ea157602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610e57575b505050505033611b43565b1515610eb757600080fd5b60088190806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6060600a600083600019166000191681526020019081526020016000206000018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610fd45780601f10610fa957610100808354040283529160200191610fd4565b820191906000526020600020905b815481529060010190602001808311610fb757829003601f168201915b50505050509050919050565b600881815481101515610fef57fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205401101515156110ad57600080fd5b80600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055505050565b6000611190600880548060200260200160405190810160405280929190818152602001828054801561118557602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161113b575b505050505033611b43565b151561119b57600080fd5b838383604051602001808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401838152602001826000191660001916815260200193505050506040516020818303038152906040526040518082805190602001908083835b602083101515611243578051825260208201915060208101905060208303925061121e565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902090506009600082600019166000191681526020019081526020016000203390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506007546009600083600019166000191681526020019081526020016000208054905014156113c55782600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508373ffffffffffffffffffffffffffffffffffffffff167f1b174056799bea141540e324bb093eb297a02b564c15e75840a30cf0d0f4837784846040518083815260200182600019166000191681526020019250505060405180910390a25b50505050565b600061145d600880548060200260200160405190810160405280929190818152602001828054801561145257602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611408575b505050505033611b43565b151561146857600080fd5b6114728284611bc1565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156114ab57600080fd5b826040518082805190602001908083835b6020831015156114e157805182526020820191506020810190506020830392506114bc565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902090506115bc600a600083600019166000191681526020019081526020016000206001018054806020026020016040519081016040528092919081815260200182805480156115b157602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611567575b505050505033611b43565b1515156115c857600080fd5b82600a6000836000191660001916815260200190815260200160002060000190805190602001906115fa929190611f80565b50600a600082600019166000191681526020019081526020016000206001013390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600a600082600019166000191681526020019081526020016000206002018290806001815401808255809150509060018203906000526020600020016000909192909190915090805190602001906116d6929190611f80565b5050600754600a6000836000191660001916815260200190815260200160002060010180549050141561177b577feb043d149eedb81369bec43d4c3a3a53087debc88d2525f13bfaa3eecda28b5c3382604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182600019166000191681526020019250505060405180910390a16117bb565b7ffc962d5f0de9bd4b90181a7ec60354bcefe1e162eaf760919b07496a184d665a8160405180826000191660001916815260200191505060405180910390a15b505050565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b611899600880548060200260200160405190810160405280929190818152602001828054801561188e57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611844575b505050505033611b43565b15156118a457600080fd5b8060078190555050565b60075481565b60006118c133848461195a565b905092915050565b6000600754905090565b6000600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6000600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482111515156119aa57600080fd5b600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020540110151515611a3957600080fd5b81600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555081600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3600190509392505050565b600080600090505b8351811015611bb5578273ffffffffffffffffffffffffffffffffffffffff168482815181101515611b7957fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff161415611ba85760019150611bba565b8080600101915050611b4b565b600091505b5092915050565b60008060008060418651141515611bd757600080fd5b6020860151925060408601519150606086015190506001611bf786611ca3565b827f010000000000000000000000000000000000000000000000000000000000000090048585604051600081526020016040526040518085600019166000191681526020018460ff1660ff1681526020018360001916600019168152602001826000191660001916815260200194505050505060206040516020810390808403906000865af1158015611c8e573d6000803e3d6000fd5b50505060206040510351935050505092915050565b600060606040805190810160405280601a81526020017f19457468657265756d205369676e6564204d6573736167653a0a000000000000815250905080611cea8451611e69565b846040516020018084805190602001908083835b602083101515611d235780518252602082019150602081019050602083039250611cfe565b6001836020036101000a03801982511681845116808217855250505050505090500183805190602001908083835b602083101515611d765780518252602082019150602081019050602083039250611d51565b6001836020036101000a03801982511681845116808217855250505050505090500182805190602001908083835b602083101515611dc95780518252602082019150602081019050602083039250611da4565b6001836020036101000a03801982511681845116808217855250505050505090500193505050506040516020818303038152906040526040518082805190602001908083835b602083101515611e345780518252602082019150602081019050602083039250611e0f565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020915050919050565b6060600080606060008093508592505b8380600101945050600a83811515611e8d57fe5b049250600083141515611e9f57611e79565b836040519080825280601f01601f191660200182016040528015611ed25781602001602082028038833980820191505090505b5091506001840390508592505b600a83811515611eeb57fe5b066030017f010000000000000000000000000000000000000000000000000000000000000002828280600190039350815181101515611f2657fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a83811515611f6257fe5b049250600083141515611f7457611edf565b81945050505050919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611fc157805160ff1916838001178555611fef565b82800160010185558215611fef579182015b82811115611fee578251825591602001919060010190611fd3565b5b509050611ffc9190612000565b5090565b61202291905b8082111561201e576000816000905550600101612006565b5090565b905600a165627a7a723058205630aca62d2bf3bad0d4ec9e98f95225d788c8e5162516d5094d2add0b3463fd0029`

// DeployPrivateSlot deploys a new Ethereum contract, binding an instance of PrivateSlot to it.
func DeployPrivateSlot(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, _requiredSignatures *big.Int, _authorities []common.Address) (common.Address, *types.Transaction, *PrivateSlot, error) {
	parsed, err := abi.JSON(strings.NewReader(PrivateSlotABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PrivateSlotBin), backend, initialSupply, _requiredSignatures, _authorities)
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
// Solidity: function allowance(_owner address, _spender address) constant returns(_remaining uint256)
func (_PrivateSlot *PrivateSlotCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivateSlot.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(_remaining uint256)
func (_PrivateSlot *PrivateSlotSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _PrivateSlot.Contract.Allowance(&_PrivateSlot.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(_remaining uint256)
func (_PrivateSlot *PrivateSlotCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _PrivateSlot.Contract.Allowance(&_PrivateSlot.CallOpts, _owner, _spender)
}

// Authorities is a free data retrieval call binding the contract method 0x494503d4.
//
// Solidity: function authorities( uint256) constant returns(address)
func (_PrivateSlot *PrivateSlotCaller) Authorities(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PrivateSlot.contract.Call(opts, out, "authorities", arg0)
	return *ret0, err
}

// Authorities is a free data retrieval call binding the contract method 0x494503d4.
//
// Solidity: function authorities( uint256) constant returns(address)
func (_PrivateSlot *PrivateSlotSession) Authorities(arg0 *big.Int) (common.Address, error) {
	return _PrivateSlot.Contract.Authorities(&_PrivateSlot.CallOpts, arg0)
}

// Authorities is a free data retrieval call binding the contract method 0x494503d4.
//
// Solidity: function authorities( uint256) constant returns(address)
func (_PrivateSlot *PrivateSlotCallerSession) Authorities(arg0 *big.Int) (common.Address, error) {
	return _PrivateSlot.Contract.Authorities(&_PrivateSlot.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_PrivateSlot *PrivateSlotCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivateSlot.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_PrivateSlot *PrivateSlotSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _PrivateSlot.Contract.BalanceOf(&_PrivateSlot.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_PrivateSlot *PrivateSlotCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _PrivateSlot.Contract.BalanceOf(&_PrivateSlot.CallOpts, _owner)
}

// GetRequiredSignatures is a free data retrieval call binding the contract method 0xccd93998.
//
// Solidity: function getRequiredSignatures() constant returns(_requiredSignatures uint256)
func (_PrivateSlot *PrivateSlotCaller) GetRequiredSignatures(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivateSlot.contract.Call(opts, out, "getRequiredSignatures")
	return *ret0, err
}

// GetRequiredSignatures is a free data retrieval call binding the contract method 0xccd93998.
//
// Solidity: function getRequiredSignatures() constant returns(_requiredSignatures uint256)
func (_PrivateSlot *PrivateSlotSession) GetRequiredSignatures() (*big.Int, error) {
	return _PrivateSlot.Contract.GetRequiredSignatures(&_PrivateSlot.CallOpts)
}

// GetRequiredSignatures is a free data retrieval call binding the contract method 0xccd93998.
//
// Solidity: function getRequiredSignatures() constant returns(_requiredSignatures uint256)
func (_PrivateSlot *PrivateSlotCallerSession) GetRequiredSignatures() (*big.Int, error) {
	return _PrivateSlot.Contract.GetRequiredSignatures(&_PrivateSlot.CallOpts)
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

// AddAuthority is a paid mutator transaction binding the contract method 0x26defa73.
//
// Solidity: function addAuthority(newAuthority address) returns()
func (_PrivateSlot *PrivateSlotTransactor) AddAuthority(opts *bind.TransactOpts, newAuthority common.Address) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "addAuthority", newAuthority)
}

// AddAuthority is a paid mutator transaction binding the contract method 0x26defa73.
//
// Solidity: function addAuthority(newAuthority address) returns()
func (_PrivateSlot *PrivateSlotSession) AddAuthority(newAuthority common.Address) (*types.Transaction, error) {
	return _PrivateSlot.Contract.AddAuthority(&_PrivateSlot.TransactOpts, newAuthority)
}

// AddAuthority is a paid mutator transaction binding the contract method 0x26defa73.
//
// Solidity: function addAuthority(newAuthority address) returns()
func (_PrivateSlot *PrivateSlotTransactorSession) AddAuthority(newAuthority common.Address) (*types.Transaction, error) {
	return _PrivateSlot.Contract.AddAuthority(&_PrivateSlot.TransactOpts, newAuthority)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(_success bool)
func (_PrivateSlot *PrivateSlotTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(_success bool)
func (_PrivateSlot *PrivateSlotSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Approve(&_PrivateSlot.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(_success bool)
func (_PrivateSlot *PrivateSlotTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Approve(&_PrivateSlot.TransactOpts, _spender, _value)
}

// Award is a paid mutator transaction binding the contract method 0x5d8a776e.
//
// Solidity: function award(_user address, _amount uint256) returns()
func (_PrivateSlot *PrivateSlotTransactor) Award(opts *bind.TransactOpts, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "award", _user, _amount)
}

// Award is a paid mutator transaction binding the contract method 0x5d8a776e.
//
// Solidity: function award(_user address, _amount uint256) returns()
func (_PrivateSlot *PrivateSlotSession) Award(_user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Award(&_PrivateSlot.TransactOpts, _user, _amount)
}

// Award is a paid mutator transaction binding the contract method 0x5d8a776e.
//
// Solidity: function award(_user address, _amount uint256) returns()
func (_PrivateSlot *PrivateSlotTransactorSession) Award(_user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Award(&_PrivateSlot.TransactOpts, _user, _amount)
}

// Consume is a paid mutator transaction binding the contract method 0x224b5c72.
//
// Solidity: function consume(_user address, _amount uint256) returns(_success bool)
func (_PrivateSlot *PrivateSlotTransactor) Consume(opts *bind.TransactOpts, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "consume", _user, _amount)
}

// Consume is a paid mutator transaction binding the contract method 0x224b5c72.
//
// Solidity: function consume(_user address, _amount uint256) returns(_success bool)
func (_PrivateSlot *PrivateSlotSession) Consume(_user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Consume(&_PrivateSlot.TransactOpts, _user, _amount)
}

// Consume is a paid mutator transaction binding the contract method 0x224b5c72.
//
// Solidity: function consume(_user address, _amount uint256) returns(_success bool)
func (_PrivateSlot *PrivateSlotTransactorSession) Consume(_user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Consume(&_PrivateSlot.TransactOpts, _user, _amount)
}

// Pay is a paid mutator transaction binding the contract method 0x5e5571ac.
//
// Solidity: function pay(_user address, _amount uint256, _transactionHash bytes32) returns()
func (_PrivateSlot *PrivateSlotTransactor) Pay(opts *bind.TransactOpts, _user common.Address, _amount *big.Int, _transactionHash [32]byte) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "pay", _user, _amount, _transactionHash)
}

// Pay is a paid mutator transaction binding the contract method 0x5e5571ac.
//
// Solidity: function pay(_user address, _amount uint256, _transactionHash bytes32) returns()
func (_PrivateSlot *PrivateSlotSession) Pay(_user common.Address, _amount *big.Int, _transactionHash [32]byte) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Pay(&_PrivateSlot.TransactOpts, _user, _amount, _transactionHash)
}

// Pay is a paid mutator transaction binding the contract method 0x5e5571ac.
//
// Solidity: function pay(_user address, _amount uint256, _transactionHash bytes32) returns()
func (_PrivateSlot *PrivateSlotTransactorSession) Pay(_user common.Address, _amount *big.Int, _transactionHash [32]byte) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Pay(&_PrivateSlot.TransactOpts, _user, _amount, _transactionHash)
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

// Sync is a paid mutator transaction binding the contract method 0x059d4135.
//
// Solidity: function sync(_user address[], _amount uint256[], _time uint256[]) returns()
func (_PrivateSlot *PrivateSlotTransactor) Sync(opts *bind.TransactOpts, _user []common.Address, _amount []*big.Int, _time []*big.Int) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "sync", _user, _amount, _time)
}

// Sync is a paid mutator transaction binding the contract method 0x059d4135.
//
// Solidity: function sync(_user address[], _amount uint256[], _time uint256[]) returns()
func (_PrivateSlot *PrivateSlotSession) Sync(_user []common.Address, _amount []*big.Int, _time []*big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Sync(&_PrivateSlot.TransactOpts, _user, _amount, _time)
}

// Sync is a paid mutator transaction binding the contract method 0x059d4135.
//
// Solidity: function sync(_user address[], _amount uint256[], _time uint256[]) returns()
func (_PrivateSlot *PrivateSlotTransactorSession) Sync(_user []common.Address, _amount []*big.Int, _time []*big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Sync(&_PrivateSlot.TransactOpts, _user, _amount, _time)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PrivateSlot *PrivateSlotTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PrivateSlot *PrivateSlotSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Transfer(&_PrivateSlot.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PrivateSlot *PrivateSlotTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.Transfer(&_PrivateSlot.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PrivateSlot *PrivateSlotTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PrivateSlot *PrivateSlotSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.TransferFrom(&_PrivateSlot.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PrivateSlot *PrivateSlotTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PrivateSlot.Contract.TransferFrom(&_PrivateSlot.TransactOpts, _from, _to, _value)
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

// PrivateSlotAwardIterator is returned from FilterAward and is used to iterate over the raw logs and unpacked data for Award events raised by the PrivateSlot contract.
type PrivateSlotAwardIterator struct {
	Event *PrivateSlotAward // Event containing the contract specifics and raw log

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
func (it *PrivateSlotAwardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateSlotAward)
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
		it.Event = new(PrivateSlotAward)
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
func (it *PrivateSlotAwardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateSlotAwardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateSlotAward represents a Award event raised by the PrivateSlot contract.
type PrivateSlotAward struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAward is a free log retrieval operation binding the contract event 0x473edf73b107bf5d270ea55a7ea4ce98a1b5618dd196e00d5a48e101299b26d4.
//
// Solidity: e Award(_user indexed address, _amount uint256)
func (_PrivateSlot *PrivateSlotFilterer) FilterAward(opts *bind.FilterOpts, _user []common.Address) (*PrivateSlotAwardIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _PrivateSlot.contract.FilterLogs(opts, "Award", _userRule)
	if err != nil {
		return nil, err
	}
	return &PrivateSlotAwardIterator{contract: _PrivateSlot.contract, event: "Award", logs: logs, sub: sub}, nil
}

// WatchAward is a free log subscription operation binding the contract event 0x473edf73b107bf5d270ea55a7ea4ce98a1b5618dd196e00d5a48e101299b26d4.
//
// Solidity: e Award(_user indexed address, _amount uint256)
func (_PrivateSlot *PrivateSlotFilterer) WatchAward(opts *bind.WatchOpts, sink chan<- *PrivateSlotAward, _user []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _PrivateSlot.contract.WatchLogs(opts, "Award", _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateSlotAward)
				if err := _PrivateSlot.contract.UnpackLog(event, "Award", log); err != nil {
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
	AuthorityResponsibleForRelay common.Address
	MessageHash                  [32]byte
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterCollectedSignatures is a free log retrieval operation binding the contract event 0xeb043d149eedb81369bec43d4c3a3a53087debc88d2525f13bfaa3eecda28b5c.
//
// Solidity: e CollectedSignatures(authorityResponsibleForRelay address, messageHash bytes32)
func (_PrivateSlot *PrivateSlotFilterer) FilterCollectedSignatures(opts *bind.FilterOpts) (*PrivateSlotCollectedSignaturesIterator, error) {

	logs, sub, err := _PrivateSlot.contract.FilterLogs(opts, "CollectedSignatures")
	if err != nil {
		return nil, err
	}
	return &PrivateSlotCollectedSignaturesIterator{contract: _PrivateSlot.contract, event: "CollectedSignatures", logs: logs, sub: sub}, nil
}

// WatchCollectedSignatures is a free log subscription operation binding the contract event 0xeb043d149eedb81369bec43d4c3a3a53087debc88d2525f13bfaa3eecda28b5c.
//
// Solidity: e CollectedSignatures(authorityResponsibleForRelay address, messageHash bytes32)
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
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterConsume is a free log retrieval operation binding the contract event 0xb3762e93ec66871dd27c421b64edc79636345ff0a949cd04f7f8efce5bd4240e.
//
// Solidity: e Consume(_user indexed address, _amount uint256)
func (_PrivateSlot *PrivateSlotFilterer) FilterConsume(opts *bind.FilterOpts, _user []common.Address) (*PrivateSlotConsumeIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _PrivateSlot.contract.FilterLogs(opts, "Consume", _userRule)
	if err != nil {
		return nil, err
	}
	return &PrivateSlotConsumeIterator{contract: _PrivateSlot.contract, event: "Consume", logs: logs, sub: sub}, nil
}

// WatchConsume is a free log subscription operation binding the contract event 0xb3762e93ec66871dd27c421b64edc79636345ff0a949cd04f7f8efce5bd4240e.
//
// Solidity: e Consume(_user indexed address, _amount uint256)
func (_PrivateSlot *PrivateSlotFilterer) WatchConsume(opts *bind.WatchOpts, sink chan<- *PrivateSlotConsume, _user []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _PrivateSlot.contract.WatchLogs(opts, "Consume", _userRule)
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

// PrivateSlotExchangeTokenIterator is returned from FilterExchangeToken and is used to iterate over the raw logs and unpacked data for ExchangeToken events raised by the PrivateSlot contract.
type PrivateSlotExchangeTokenIterator struct {
	Event *PrivateSlotExchangeToken // Event containing the contract specifics and raw log

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
func (it *PrivateSlotExchangeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateSlotExchangeToken)
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
		it.Event = new(PrivateSlotExchangeToken)
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
func (it *PrivateSlotExchangeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateSlotExchangeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateSlotExchangeToken represents a ExchangeToken event raised by the PrivateSlot contract.
type PrivateSlotExchangeToken struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExchangeToken is a free log retrieval operation binding the contract event 0xf50dd38e54ff60d756c111498c058ab11d7f0de3ea8a045e08313ff844b136f6.
//
// Solidity: e ExchangeToken(user address, amount uint256)
func (_PrivateSlot *PrivateSlotFilterer) FilterExchangeToken(opts *bind.FilterOpts) (*PrivateSlotExchangeTokenIterator, error) {

	logs, sub, err := _PrivateSlot.contract.FilterLogs(opts, "ExchangeToken")
	if err != nil {
		return nil, err
	}
	return &PrivateSlotExchangeTokenIterator{contract: _PrivateSlot.contract, event: "ExchangeToken", logs: logs, sub: sub}, nil
}

// WatchExchangeToken is a free log subscription operation binding the contract event 0xf50dd38e54ff60d756c111498c058ab11d7f0de3ea8a045e08313ff844b136f6.
//
// Solidity: e ExchangeToken(user address, amount uint256)
func (_PrivateSlot *PrivateSlotFilterer) WatchExchangeToken(opts *bind.WatchOpts, sink chan<- *PrivateSlotExchangeToken) (event.Subscription, error) {

	logs, sub, err := _PrivateSlot.contract.WatchLogs(opts, "ExchangeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateSlotExchangeToken)
				if err := _PrivateSlot.contract.UnpackLog(event, "ExchangeToken", log); err != nil {
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

// PrivateSlotLoginIterator is returned from FilterLogin and is used to iterate over the raw logs and unpacked data for Login events raised by the PrivateSlot contract.
type PrivateSlotLoginIterator struct {
	Event *PrivateSlotLogin // Event containing the contract specifics and raw log

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
func (it *PrivateSlotLoginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateSlotLogin)
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
		it.Event = new(PrivateSlotLogin)
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
func (it *PrivateSlotLoginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateSlotLoginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateSlotLogin represents a Login event raised by the PrivateSlot contract.
type PrivateSlotLogin struct {
	PrivateAdd common.Address
	PublicAdd  common.Address
	Time       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogin is a free log retrieval operation binding the contract event 0xf5b7370f7cf13ae233764c6a70ddf85ca7a721e8d09bba8ab4287bd0388e11f6.
//
// Solidity: e Login(privateAdd indexed address, publicAdd indexed address, time uint256)
func (_PrivateSlot *PrivateSlotFilterer) FilterLogin(opts *bind.FilterOpts, privateAdd []common.Address, publicAdd []common.Address) (*PrivateSlotLoginIterator, error) {

	var privateAddRule []interface{}
	for _, privateAddItem := range privateAdd {
		privateAddRule = append(privateAddRule, privateAddItem)
	}
	var publicAddRule []interface{}
	for _, publicAddItem := range publicAdd {
		publicAddRule = append(publicAddRule, publicAddItem)
	}

	logs, sub, err := _PrivateSlot.contract.FilterLogs(opts, "Login", privateAddRule, publicAddRule)
	if err != nil {
		return nil, err
	}
	return &PrivateSlotLoginIterator{contract: _PrivateSlot.contract, event: "Login", logs: logs, sub: sub}, nil
}

// WatchLogin is a free log subscription operation binding the contract event 0xf5b7370f7cf13ae233764c6a70ddf85ca7a721e8d09bba8ab4287bd0388e11f6.
//
// Solidity: e Login(privateAdd indexed address, publicAdd indexed address, time uint256)
func (_PrivateSlot *PrivateSlotFilterer) WatchLogin(opts *bind.WatchOpts, sink chan<- *PrivateSlotLogin, privateAdd []common.Address, publicAdd []common.Address) (event.Subscription, error) {

	var privateAddRule []interface{}
	for _, privateAddItem := range privateAdd {
		privateAddRule = append(privateAddRule, privateAddItem)
	}
	var publicAddRule []interface{}
	for _, publicAddItem := range publicAdd {
		publicAddRule = append(publicAddRule, publicAddItem)
	}

	logs, sub, err := _PrivateSlot.contract.WatchLogs(opts, "Login", privateAddRule, publicAddRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateSlotLogin)
				if err := _PrivateSlot.contract.UnpackLog(event, "Login", log); err != nil {
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
// Solidity: e Pay(_user indexed address, _amount uint256, transactoinHash bytes32)
func (_PrivateSlot *PrivateSlotFilterer) FilterPay(opts *bind.FilterOpts, _user []common.Address) (*PrivateSlotPayIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _PrivateSlot.contract.FilterLogs(opts, "Pay", _userRule)
	if err != nil {
		return nil, err
	}
	return &PrivateSlotPayIterator{contract: _PrivateSlot.contract, event: "Pay", logs: logs, sub: sub}, nil
}

// WatchPay is a free log subscription operation binding the contract event 0x1b174056799bea141540e324bb093eb297a02b564c15e75840a30cf0d0f48377.
//
// Solidity: e Pay(_user indexed address, _amount uint256, transactoinHash bytes32)
func (_PrivateSlot *PrivateSlotFilterer) WatchPay(opts *bind.WatchOpts, sink chan<- *PrivateSlotPay, _user []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _PrivateSlot.contract.WatchLogs(opts, "Pay", _userRule)
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

// PrivateSlotSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the PrivateSlot contract.
type PrivateSlotSyncIterator struct {
	Event *PrivateSlotSync // Event containing the contract specifics and raw log

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
func (it *PrivateSlotSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateSlotSync)
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
		it.Event = new(PrivateSlotSync)
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
func (it *PrivateSlotSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateSlotSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateSlotSync represents a Sync event raised by the PrivateSlot contract.
type PrivateSlotSync struct {
	From   []common.Address
	To     []common.Address
	Amount []*big.Int
	Time   []*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0x97958b0c39735ab08db9a8f6ceab791cd74b6842a2dd8cb67a7fff23cc72b679.
//
// Solidity: e Sync(_from address[], _to address[], _amount uint256[], time uint256[])
func (_PrivateSlot *PrivateSlotFilterer) FilterSync(opts *bind.FilterOpts) (*PrivateSlotSyncIterator, error) {

	logs, sub, err := _PrivateSlot.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &PrivateSlotSyncIterator{contract: _PrivateSlot.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0x97958b0c39735ab08db9a8f6ceab791cd74b6842a2dd8cb67a7fff23cc72b679.
//
// Solidity: e Sync(_from address[], _to address[], _amount uint256[], time uint256[])
func (_PrivateSlot *PrivateSlotFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *PrivateSlotSync) (event.Subscription, error) {

	logs, sub, err := _PrivateSlot.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateSlotSync)
				if err := _PrivateSlot.contract.UnpackLog(event, "Sync", log); err != nil {
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

// PrivateSlotSyncConfirmationIterator is returned from FilterSyncConfirmation and is used to iterate over the raw logs and unpacked data for SyncConfirmation events raised by the PrivateSlot contract.
type PrivateSlotSyncConfirmationIterator struct {
	Event *PrivateSlotSyncConfirmation // Event containing the contract specifics and raw log

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
func (it *PrivateSlotSyncConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateSlotSyncConfirmation)
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
		it.Event = new(PrivateSlotSyncConfirmation)
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
func (it *PrivateSlotSyncConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateSlotSyncConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateSlotSyncConfirmation represents a SyncConfirmation event raised by the PrivateSlot contract.
type PrivateSlotSyncConfirmation struct {
	From common.Address
	To   common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSyncConfirmation is a free log retrieval operation binding the contract event 0x65caa97aee6cdc0218d86358a64a76db2568ce77774417e7f97ec394f35b0c69.
//
// Solidity: e SyncConfirmation(_from indexed address, _to indexed address, time uint256)
func (_PrivateSlot *PrivateSlotFilterer) FilterSyncConfirmation(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*PrivateSlotSyncConfirmationIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _PrivateSlot.contract.FilterLogs(opts, "SyncConfirmation", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &PrivateSlotSyncConfirmationIterator{contract: _PrivateSlot.contract, event: "SyncConfirmation", logs: logs, sub: sub}, nil
}

// WatchSyncConfirmation is a free log subscription operation binding the contract event 0x65caa97aee6cdc0218d86358a64a76db2568ce77774417e7f97ec394f35b0c69.
//
// Solidity: e SyncConfirmation(_from indexed address, _to indexed address, time uint256)
func (_PrivateSlot *PrivateSlotFilterer) WatchSyncConfirmation(opts *bind.WatchOpts, sink chan<- *PrivateSlotSyncConfirmation, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _PrivateSlot.contract.WatchLogs(opts, "SyncConfirmation", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateSlotSyncConfirmation)
				if err := _PrivateSlot.contract.UnpackLog(event, "SyncConfirmation", log); err != nil {
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

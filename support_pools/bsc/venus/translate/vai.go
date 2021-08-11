// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package translate

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// VaiABI is the input ABI used to generate the binding from.
const VaiABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"move\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"pull\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"push\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VaiBin is the compiled bytecode used for deploying new contracts.
var VaiBin = "0x608060405234801561001057600080fd5b506040516112923803806112928339818101604052602081101561003357600080fd5b5051336000908152602081905260409081902060019055518060526112408239604080519182900360520182208282018252600e83526d2b20a49029ba30b13632b1b7b4b760911b6020938401528151808301835260018152603160f81b908401528151808401919091527fdd18a44c5a898a5e53e119eaa2f28613c9f2172cf68c0fb769cbcdd3e3296b05818301527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6606082015260808101949094523060a0808601919091528151808603909101815260c09094019052825192019190912060055550611119806101276000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c80637ecebe00116100b8578063a9059cbb1161007c578063a9059cbb146103e0578063b753a98c1461040c578063bb35783b14610438578063bf353dbb1461046e578063dd62ed3e14610494578063f2d5d56b146104c257610142565b80637ecebe00146103065780638fcbaf0c1461032c57806395d89b41146103865780639c52a7f11461038e5780639dc29fac146103b457610142565b8063313ce5671161010a578063313ce5671461025e5780633644e5151461027c57806340c10f191461028457806354fd4d50146102b257806365fae35e146102ba57806370a08231146102e057610142565b806306fdde0314610147578063095ea7b3146101c657806318160ddd1461020657806323b872dd1461022057806330adf81f14610256575b600080fd5b61014f6104ee565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561018b578082015183820152602001610173565b50505050905090810190601f1680156101b85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101f2600480360360408110156101dc57600080fd5b506001600160a01b038135169060200135610516565b604051901515815260200160405180910390f35b61020e610588565b60405190815260200160405180910390f35b6101f26004803603606081101561023657600080fd5b506001600160a01b0381358116916020810135909116906040013561058e565b61020e6107db565b6102666107ff565b60405160ff909116815260200160405180910390f35b61020e610804565b6102b06004803603604081101561029a57600080fd5b506001600160a01b03813516906020013561080a565b005b61014f6108f1565b6102b0600480360360208110156102d057600080fd5b50356001600160a01b031661090c565b61020e600480360360208110156102f657600080fd5b50356001600160a01b03166109ba565b61020e6004803603602081101561031c57600080fd5b50356001600160a01b03166109ce565b6102b0600480360361010081101561034357600080fd5b506001600160a01b038135811691602081013590911690604081013590606081013590608081013515159060ff60a0820135169060c08101359060e001356109e2565b61014f610cfc565b6102b0600480360360208110156103a457600080fd5b50356001600160a01b0316610d19565b6102b0600480360360408110156103ca57600080fd5b506001600160a01b038135169060200135610dc4565b6101f2600480360360408110156103f657600080fd5b506001600160a01b038135169060200135610fde565b6102b06004803603604081101561042257600080fd5b506001600160a01b038135169060200135610ff2565b6102b06004803603606081101561044e57600080fd5b506001600160a01b03813581169160208101359091169060400135611002565b61020e6004803603602081101561048457600080fd5b50356001600160a01b0316611013565b61020e600480360360408110156104aa57600080fd5b506001600160a01b0381358116916020013516611027565b6102b0600480360360408110156104d857600080fd5b506001600160a01b038135169060200135611049565b60405160408082019052600e81526d2b20a49029ba30b13632b1b7b4b760911b602082015281565b336000908152600360205281604082206001600160a01b038516600090815260209190915260409020556001600160a01b038316337f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405190815260200160405180910390a35060015b92915050565b60015481565b6001600160a01b0383166000908152600260205281604082205410156105f55760405162461bcd60e51b81526020600482015260186024820152775641492f696e73756666696369656e742d62616c616e636560401b604482015260640160405180910390fd5b6001600160a01b038416331480159061063957506001600160a01b038416600090815260036020526000199060409020336000908152602091909152604090205414155b15610718576001600160a01b03841660009081526003602052829060409020336000908152602091909152604090205410156106bb5760405162461bcd60e51b815260206004820152601a60248201527f5641492f696e73756666696369656e742d616c6c6f77616e6365000000000000604482015260640160405180910390fd5b6001600160a01b038416600090815260036020526106ee9060409020336000908152602091909152604090205483611054565b6001600160a01b038516600090815260036020526040902033600090815260209190915260409020555b6001600160a01b0384166000908152600260205261073b90604090205483611054565b6001600160a01b0385166000908152600260205260409020556001600160a01b038316600090815260026020526107779060409020548361109c565b6001600160a01b0384166000908152600260205260409020556001600160a01b038084169085167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a35060019392505050565b7fea2aa0a1be11a07ed86d755c93467f4f82362b452371d1ba94d1715123511acb81565b601281565b60055481565b336000908152602081905260409020546001146108625760405162461bcd60e51b81526020600482015260126024820152711590524bdb9bdd0b585d5d1a1bdc9a5e995960721b604482015260640160405180910390fd5b6001600160a01b038216600090815260026020526108859060409020548261109c565b6001600160a01b0383166000908152600260205260409020556001546108ab908261109c565b6001556001600160a01b03821660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405190815260200160405180910390a35050565b6040516040808201905260018152603160f81b602082015281565b336000908152602081905260409020546001146109645760405162461bcd60e51b81526020600482015260126024820152711590524bdb9bdd0b585d5d1a1bdc9a5e995960721b604482015260640160405180910390fd5b6001600160a01b0381166000908152602081905260019060409020555961012081016040526020815260e0602082015260e060006040830137602435600435336001600160e01b03196000351661012085a45050565b600260205280600052604060002054905081565b600460205280600052604060002054905081565b6005546000907fea2aa0a1be11a07ed86d755c93467f4f82362b452371d1ba94d1715123511acb8a8a8a8a8a60405160208101969096526001600160a01b03948516604080880191909152939094166060860152608085019190915260a084015290151560c083015260e090910190516020818303038152906040528051906020012060405161190160f01b6020820152602281019290925260428201526062016040516020818303038152906040528051906020012090506001600160a01b038916610aed5760405162461bcd60e51b815260206004820152601560248201527405641492f696e76616c69642d616464726573732d3605c1b604482015260640160405180910390fd5b60018185858560405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015610b4a573d6000803e3d6000fd5b505050602060405103516001600160a01b0316896001600160a01b031614610bad5760405162461bcd60e51b81526020600482015260126024820152711590524bda5b9d985b1a590b5c195c9b5a5d60721b604482015260640160405180910390fd5b851580610bba5750854211155b610bff5760405162461bcd60e51b81526020600482015260126024820152711590524bdc195c9b5a5d0b595e1c1a5c995960721b604482015260640160405180910390fd5b6001600160a01b03891660009081526004602052604090208054600181019091558714610c665760405162461bcd60e51b81526020600482015260116024820152705641492f696e76616c69642d6e6f6e636560781b604482015260640160405180910390fd5b600085610c74576000610c78565b6000195b6001600160a01b038b16600090815260036020529091508190604090206001600160a01b038b16600090815260209190915260409020556001600160a01b03808a16908b167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258360405190815260200160405180910390a350505050505050505050565b60405160408082019052600381526256414960e81b602082015281565b33600090815260208190526040902054600114610d715760405162461bcd60e51b81526020600482015260126024820152711590524bdb9bdd0b585d5d1a1bdc9a5e995960721b604482015260640160405180910390fd5b6001600160a01b0381166000908152602081905260408120555961012081016040526020815260e0602082015260e060006040830137602435600435336001600160e01b03196000351661012085a45050565b6001600160a01b03821660009081526002602052819060409020541015610e2c5760405162461bcd60e51b81526020600482015260186024820152775641492f696e73756666696369656e742d62616c616e636560401b604482015260640160405180910390fd5b6001600160a01b0382163314801590610e7057506001600160a01b038216600090815260036020526000199060409020336000908152602091909152604090205414155b15610f4f576001600160a01b0382166000908152600360205281906040902033600090815260209190915260409020541015610ef25760405162461bcd60e51b815260206004820152601a60248201527f5641492f696e73756666696369656e742d616c6c6f77616e6365000000000000604482015260640160405180910390fd5b6001600160a01b03821660009081526003602052610f259060409020336000908152602091909152604090205482611054565b6001600160a01b038316600090815260036020526040902033600090815260209190915260409020555b6001600160a01b03821660009081526002602052610f7290604090205482611054565b6001600160a01b038316600090815260026020526040902055600154610f989082611054565b60015560006001600160a01b0383167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405190815260200160405180910390a35050565b6000610feb33848461058e565b9392505050565b610ffd33838361058e565b505050565b61100d83838361058e565b50505050565b600060205280600052604060002054905081565b6003602052816000526040600020602052806000526040600020549150829050565b610ffd82338361058e565b808203828111156105825760405162461bcd60e51b815260206004820152600e60248201526d2b20a49036b0ba341032b93937b960911b604482015260640160405180910390fd5b808201828110156105825760405162461bcd60e51b815260206004820152600e60248201526d2b20a49036b0ba341032b93937b960911b604482015260640160405180910390fdfea265627a7a723158207032c3145907d0d892823eb54be4d00d91e9bd88d47a2a55b2e4936e569e9b5d64736f6c63430005110032454950373132446f6d61696e28737472696e67206e616d652c737472696e672076657273696f6e2c75696e7432353620636861696e49642c6164647265737320766572696679696e67436f6e7472616374290000000000000000000000000000000000000000000000000000000000000038"

// DeployVai deploys a new Ethereum contract, binding an instance of Vai to it.
func DeployVai(auth *bind.TransactOpts, backend bind.ContractBackend, chainId_ *big.Int) (common.Address, *types.Transaction, *Vai, error) {
	parsed, err := abi.JSON(strings.NewReader(VaiABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaiBin), backend, chainId_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vai{VaiCaller: VaiCaller{contract: contract}, VaiTransactor: VaiTransactor{contract: contract}, VaiFilterer: VaiFilterer{contract: contract}}, nil
}

// Vai is an auto generated Go binding around an Ethereum contract.
type Vai struct {
	VaiCaller     // Read-only binding to the contract
	VaiTransactor // Write-only binding to the contract
	VaiFilterer   // Log filterer for contract events
}

// VaiCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaiSession struct {
	Contract     *Vai              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaiCallerSession struct {
	Contract *VaiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VaiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaiTransactorSession struct {
	Contract     *VaiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaiRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaiRaw struct {
	Contract *Vai // Generic contract binding to access the raw methods on
}

// VaiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaiCallerRaw struct {
	Contract *VaiCaller // Generic read-only contract binding to access the raw methods on
}

// VaiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaiTransactorRaw struct {
	Contract *VaiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVai creates a new instance of Vai, bound to a specific deployed contract.
func NewVai(address common.Address, backend bind.ContractBackend) (*Vai, error) {
	contract, err := bindVai(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vai{VaiCaller: VaiCaller{contract: contract}, VaiTransactor: VaiTransactor{contract: contract}, VaiFilterer: VaiFilterer{contract: contract}}, nil
}

// NewVaiCaller creates a new read-only instance of Vai, bound to a specific deployed contract.
func NewVaiCaller(address common.Address, caller bind.ContractCaller) (*VaiCaller, error) {
	contract, err := bindVai(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaiCaller{contract: contract}, nil
}

// NewVaiTransactor creates a new write-only instance of Vai, bound to a specific deployed contract.
func NewVaiTransactor(address common.Address, transactor bind.ContractTransactor) (*VaiTransactor, error) {
	contract, err := bindVai(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaiTransactor{contract: contract}, nil
}

// NewVaiFilterer creates a new log filterer instance of Vai, bound to a specific deployed contract.
func NewVaiFilterer(address common.Address, filterer bind.ContractFilterer) (*VaiFilterer, error) {
	contract, err := bindVai(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaiFilterer{contract: contract}, nil
}

// bindVai binds a generic wrapper to an already deployed contract.
func bindVai(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vai *VaiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vai.Contract.VaiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vai *VaiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vai.Contract.VaiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vai *VaiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vai.Contract.VaiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vai *VaiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vai.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vai *VaiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vai.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vai *VaiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vai.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Vai *VaiCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vai.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Vai *VaiSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Vai.Contract.DOMAINSEPARATOR(&_Vai.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Vai *VaiCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Vai.Contract.DOMAINSEPARATOR(&_Vai.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Vai *VaiCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vai.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Vai *VaiSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Vai.Contract.PERMITTYPEHASH(&_Vai.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Vai *VaiCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Vai.Contract.PERMITTYPEHASH(&_Vai.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Vai *VaiCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vai.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Vai *VaiSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Vai.Contract.Allowance(&_Vai.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Vai *VaiCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Vai.Contract.Allowance(&_Vai.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Vai *VaiCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vai.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Vai *VaiSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Vai.Contract.BalanceOf(&_Vai.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Vai *VaiCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Vai.Contract.BalanceOf(&_Vai.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Vai *VaiCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Vai.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Vai *VaiSession) Decimals() (uint8, error) {
	return _Vai.Contract.Decimals(&_Vai.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Vai *VaiCallerSession) Decimals() (uint8, error) {
	return _Vai.Contract.Decimals(&_Vai.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vai *VaiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vai.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vai *VaiSession) Name() (string, error) {
	return _Vai.Contract.Name(&_Vai.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vai *VaiCallerSession) Name() (string, error) {
	return _Vai.Contract.Name(&_Vai.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Vai *VaiCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vai.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Vai *VaiSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Vai.Contract.Nonces(&_Vai.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Vai *VaiCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Vai.Contract.Nonces(&_Vai.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vai *VaiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vai.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vai *VaiSession) Symbol() (string, error) {
	return _Vai.Contract.Symbol(&_Vai.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vai *VaiCallerSession) Symbol() (string, error) {
	return _Vai.Contract.Symbol(&_Vai.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vai *VaiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vai.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vai *VaiSession) TotalSupply() (*big.Int, error) {
	return _Vai.Contract.TotalSupply(&_Vai.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vai *VaiCallerSession) TotalSupply() (*big.Int, error) {
	return _Vai.Contract.TotalSupply(&_Vai.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Vai *VaiCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vai.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Vai *VaiSession) Version() (string, error) {
	return _Vai.Contract.Version(&_Vai.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Vai *VaiCallerSession) Version() (string, error) {
	return _Vai.Contract.Version(&_Vai.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Vai *VaiCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vai.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Vai *VaiSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Vai.Contract.Wards(&_Vai.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Vai *VaiCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Vai.Contract.Wards(&_Vai.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address usr, uint256 wad) returns(bool)
func (_Vai *VaiTransactor) Approve(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vai.contract.Transact(opts, "approve", usr, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address usr, uint256 wad) returns(bool)
func (_Vai *VaiSession) Approve(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vai.Contract.Approve(&_Vai.TransactOpts, usr, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address usr, uint256 wad) returns(bool)
func (_Vai *VaiTransactorSession) Approve(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vai.Contract.Approve(&_Vai.TransactOpts, usr, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address usr, uint256 wad) returns()
func (_Vai *VaiTransactor) Burn(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vai.contract.Transact(opts, "burn", usr, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address usr, uint256 wad) returns()
func (_Vai *VaiSession) Burn(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vai.Contract.Burn(&_Vai.TransactOpts, usr, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address usr, uint256 wad) returns()
func (_Vai *VaiTransactorSession) Burn(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vai.Contract.Burn(&_Vai.TransactOpts, usr, wad)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_Vai *VaiTransactor) Deny(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _Vai.contract.Transact(opts, "deny", guy)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_Vai *VaiSession) Deny(guy common.Address) (*types.Transaction, error) {
	return _Vai.Contract.Deny(&_Vai.TransactOpts, guy)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_Vai *VaiTransactorSession) Deny(guy common.Address) (*types.Transaction, error) {
	return _Vai.Contract.Deny(&_Vai.TransactOpts, guy)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address usr, uint256 wad) returns()
func (_Vai *VaiTransactor) Mint(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vai.contract.Transact(opts, "mint", usr, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address usr, uint256 wad) returns()
func (_Vai *VaiSession) Mint(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vai.Contract.Mint(&_Vai.TransactOpts, usr, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address usr, uint256 wad) returns()
func (_Vai *VaiTransactorSession) Mint(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vai.Contract.Mint(&_Vai.TransactOpts, usr, wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 wad) returns()
func (_Vai *VaiTransactor) Move(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vai.contract.Transact(opts, "move", src, dst, wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 wad) returns()
func (_Vai *VaiSession) Move(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vai.Contract.Move(&_Vai.TransactOpts, src, dst, wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 wad) returns()
func (_Vai *VaiTransactorSession) Move(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vai.Contract.Move(&_Vai.TransactOpts, src, dst, wad)
}

// Permit is a paid mutator transaction binding the contract method 0x8fcbaf0c.
//
// Solidity: function permit(address holder, address spender, uint256 nonce, uint256 expiry, bool allowed, uint8 v, bytes32 r, bytes32 s) returns()
func (_Vai *VaiTransactor) Permit(opts *bind.TransactOpts, holder common.Address, spender common.Address, nonce *big.Int, expiry *big.Int, allowed bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Vai.contract.Transact(opts, "permit", holder, spender, nonce, expiry, allowed, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0x8fcbaf0c.
//
// Solidity: function permit(address holder, address spender, uint256 nonce, uint256 expiry, bool allowed, uint8 v, bytes32 r, bytes32 s) returns()
func (_Vai *VaiSession) Permit(holder common.Address, spender common.Address, nonce *big.Int, expiry *big.Int, allowed bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Vai.Contract.Permit(&_Vai.TransactOpts, holder, spender, nonce, expiry, allowed, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0x8fcbaf0c.
//
// Solidity: function permit(address holder, address spender, uint256 nonce, uint256 expiry, bool allowed, uint8 v, bytes32 r, bytes32 s) returns()
func (_Vai *VaiTransactorSession) Permit(holder common.Address, spender common.Address, nonce *big.Int, expiry *big.Int, allowed bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Vai.Contract.Permit(&_Vai.TransactOpts, holder, spender, nonce, expiry, allowed, v, r, s)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address usr, uint256 wad) returns()
func (_Vai *VaiTransactor) Pull(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vai.contract.Transact(opts, "pull", usr, wad)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address usr, uint256 wad) returns()
func (_Vai *VaiSession) Pull(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vai.Contract.Pull(&_Vai.TransactOpts, usr, wad)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address usr, uint256 wad) returns()
func (_Vai *VaiTransactorSession) Pull(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vai.Contract.Pull(&_Vai.TransactOpts, usr, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(address usr, uint256 wad) returns()
func (_Vai *VaiTransactor) Push(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vai.contract.Transact(opts, "push", usr, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(address usr, uint256 wad) returns()
func (_Vai *VaiSession) Push(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vai.Contract.Push(&_Vai.TransactOpts, usr, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(address usr, uint256 wad) returns()
func (_Vai *VaiTransactorSession) Push(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vai.Contract.Push(&_Vai.TransactOpts, usr, wad)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_Vai *VaiTransactor) Rely(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _Vai.contract.Transact(opts, "rely", guy)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_Vai *VaiSession) Rely(guy common.Address) (*types.Transaction, error) {
	return _Vai.Contract.Rely(&_Vai.TransactOpts, guy)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_Vai *VaiTransactorSession) Rely(guy common.Address) (*types.Transaction, error) {
	return _Vai.Contract.Rely(&_Vai.TransactOpts, guy)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Vai *VaiTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vai.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Vai *VaiSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vai.Contract.Transfer(&_Vai.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Vai *VaiTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vai.Contract.Transfer(&_Vai.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Vai *VaiTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vai.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Vai *VaiSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vai.Contract.TransferFrom(&_Vai.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Vai *VaiTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vai.Contract.TransferFrom(&_Vai.TransactOpts, src, dst, wad)
}

// VaiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Vai contract.
type VaiApprovalIterator struct {
	Event *VaiApproval // Event containing the contract specifics and raw log

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
func (it *VaiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaiApproval)
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
		it.Event = new(VaiApproval)
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
func (it *VaiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaiApproval represents a Approval event raised by the Vai contract.
type VaiApproval struct {
	Src common.Address
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_Vai *VaiFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*VaiApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Vai.contract.FilterLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return &VaiApprovalIterator{contract: _Vai.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_Vai *VaiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VaiApproval, src []common.Address, guy []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Vai.contract.WatchLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaiApproval)
				if err := _Vai.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_Vai *VaiFilterer) ParseApproval(log types.Log) (*VaiApproval, error) {
	event := new(VaiApproval)
	if err := _Vai.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Vai contract.
type VaiTransferIterator struct {
	Event *VaiTransfer // Event containing the contract specifics and raw log

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
func (it *VaiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaiTransfer)
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
		it.Event = new(VaiTransfer)
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
func (it *VaiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaiTransfer represents a Transfer event raised by the Vai contract.
type VaiTransfer struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_Vai *VaiFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*VaiTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Vai.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &VaiTransferIterator{contract: _Vai.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_Vai *VaiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VaiTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Vai.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaiTransfer)
				if err := _Vai.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_Vai *VaiFilterer) ParseTransfer(log types.Log) (*VaiTransfer, error) {
	event := new(VaiTransfer)
	if err := _Vai.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

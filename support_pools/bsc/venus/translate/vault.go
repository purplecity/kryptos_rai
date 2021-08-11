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

// VaultABI is the input ABI used to generate the binding from.
const VaultABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminTransfered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVAIVaultProxy\",\"name\":\"vaiVaultProxy\",\"type\":\"address\"}],\"name\":\"_become\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accXVSPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"burnAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingVAIVaultImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingXVS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setNewAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_xvs\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vai\",\"type\":\"address\"}],\"name\":\"setVenusInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updatePendingRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vai\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaiVaultImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"xvs\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"xvsBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VaultBin is the compiled bytecode used for deploying new contracts.
var VaultBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b031916331790556105c3806100326000396000f3fe60806040526004361061007b5760003560e01c8063e992a0411161004e578063e992a0411461019e578063e9c714f2146101d1578063f661bb86146101e6578063f851a440146101fb5761007b565b8063211de6b6146100fe578063267822471461012f578063b71d1a0c14610144578063c1e8033414610189575b6002546040516000916001600160a01b031690829036908083838082843760405192019450600093509091505080830381855af49150503d80600081146100de576040519150601f19603f3d011682016040523d82523d6000602084013e6100e3565b606091505b505090506040513d6000823e8180156100fa573d82f35b3d82fd5b34801561010a57600080fd5b50610113610210565b604080516001600160a01b039092168252519081900360200190f35b34801561013b57600080fd5b5061011361021f565b34801561015057600080fd5b506101776004803603602081101561016757600080fd5b50356001600160a01b031661022e565b60408051918252519081900360200190f35b34801561019557600080fd5b506101776102bf565b3480156101aa57600080fd5b50610177600480360360208110156101c157600080fd5b50356001600160a01b03166103a4565b3480156101dd57600080fd5b50610177610428565b3480156101f257600080fd5b50610113610503565b34801561020757600080fd5b50610113610512565b6003546001600160a01b031681565b6001546001600160a01b031681565b600080546001600160a01b031633146102545761024d60016002610521565b90506102ba565b600180546001600160a01b038481166001600160a01b0319831681179093556040805191909216808252602082019390935281517fca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9929181900390910190a160005b9150505b919050565b6003546000906001600160a01b031633146102e6576102df600180610521565b90506103a1565b60028054600380546001600160a01b038082166001600160a01b031980861682179687905590921690925560408051938316808552949092166020840152815190927fd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a92908290030190a1600354604080516001600160a01b038085168252909216602083015280517fe945ccee5d701fc83f9b8aa8ca94ea4219ec1fcbd4f4cab4f0ea57c5c3e1d8159281900390910190a160005b925050505b90565b600080546001600160a01b031633146103c35761024d60016003610521565b600380546001600160a01b038481166001600160a01b0319831617928390556040805192821680845293909116602083015280517fe945ccee5d701fc83f9b8aa8ca94ea4219ec1fcbd4f4cab4f0ea57c5c3e1d8159281900390910190a160006102b6565b6001546000906001600160a01b03163314610449576102df60016000610521565b60008054600180546001600160a01b038082166001600160a01b031980861682179687905590921690925560408051938316808552949092166020840152815190927ff9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc92908290030190a1600154604080516001600160a01b038085168252909216602083015280517fca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a99281900390910190a1600061039c565b6002546001600160a01b031681565b6000546001600160a01b031681565b60007f45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa083600181111561055057fe5b83600381111561055c57fe5b604080519283526020830191909152600082820152519081900360600190a182600181111561058757fe5b939250505056fea265627a7a723158204f9eb826cb4e2e59e6913f7469288cf96aec04a96f7856fca321efc4934cbfb564736f6c63430005110032"

// DeployVault deploys a new Ethereum contract, binding an instance of Vault to it.
func DeployVault(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Vault, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaultBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// Vault is an auto generated Go binding around an Ethereum contract.
type Vault struct {
	VaultCaller     // Read-only binding to the contract
	VaultTransactor // Write-only binding to the contract
	VaultFilterer   // Log filterer for contract events
}

// VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultSession struct {
	Contract     *Vault            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultCallerSession struct {
	Contract *VaultCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultTransactorSession struct {
	Contract     *VaultTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultRaw struct {
	Contract *Vault // Generic contract binding to access the raw methods on
}

// VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultCallerRaw struct {
	Contract *VaultCaller // Generic read-only contract binding to access the raw methods on
}

// VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultTransactorRaw struct {
	Contract *VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVault creates a new instance of Vault, bound to a specific deployed contract.
func NewVault(address common.Address, backend bind.ContractBackend) (*Vault, error) {
	contract, err := bindVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// NewVaultCaller creates a new read-only instance of Vault, bound to a specific deployed contract.
func NewVaultCaller(address common.Address, caller bind.ContractCaller) (*VaultCaller, error) {
	contract, err := bindVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultCaller{contract: contract}, nil
}

// NewVaultTransactor creates a new write-only instance of Vault, bound to a specific deployed contract.
func NewVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultTransactor, error) {
	contract, err := bindVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTransactor{contract: contract}, nil
}

// NewVaultFilterer creates a new log filterer instance of Vault, bound to a specific deployed contract.
func NewVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultFilterer, error) {
	contract, err := bindVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultFilterer{contract: contract}, nil
}

// bindVault binds a generic wrapper to an already deployed contract.
func bindVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transact(opts, method, params...)
}

// AccXVSPerShare is a free data retrieval call binding the contract method 0x91cc3da1.
//
// Solidity: function accXVSPerShare() view returns(uint256)
func (_Vault *VaultCaller) AccXVSPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "accXVSPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccXVSPerShare is a free data retrieval call binding the contract method 0x91cc3da1.
//
// Solidity: function accXVSPerShare() view returns(uint256)
func (_Vault *VaultSession) AccXVSPerShare() (*big.Int, error) {
	return _Vault.Contract.AccXVSPerShare(&_Vault.CallOpts)
}

// AccXVSPerShare is a free data retrieval call binding the contract method 0x91cc3da1.
//
// Solidity: function accXVSPerShare() view returns(uint256)
func (_Vault *VaultCallerSession) AccXVSPerShare() (*big.Int, error) {
	return _Vault.Contract.AccXVSPerShare(&_Vault.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Vault *VaultCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Vault *VaultSession) Admin() (common.Address, error) {
	return _Vault.Contract.Admin(&_Vault.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Vault *VaultCallerSession) Admin() (common.Address, error) {
	return _Vault.Contract.Admin(&_Vault.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_Vault *VaultCaller) GetAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_Vault *VaultSession) GetAdmin() (common.Address, error) {
	return _Vault.Contract.GetAdmin(&_Vault.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_Vault *VaultCallerSession) GetAdmin() (common.Address, error) {
	return _Vault.Contract.GetAdmin(&_Vault.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Vault *VaultCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Vault *VaultSession) PendingAdmin() (common.Address, error) {
	return _Vault.Contract.PendingAdmin(&_Vault.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Vault *VaultCallerSession) PendingAdmin() (common.Address, error) {
	return _Vault.Contract.PendingAdmin(&_Vault.CallOpts)
}

// PendingRewards is a free data retrieval call binding the contract method 0xeded3fda.
//
// Solidity: function pendingRewards() view returns(uint256)
func (_Vault *VaultCaller) PendingRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "pendingRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingRewards is a free data retrieval call binding the contract method 0xeded3fda.
//
// Solidity: function pendingRewards() view returns(uint256)
func (_Vault *VaultSession) PendingRewards() (*big.Int, error) {
	return _Vault.Contract.PendingRewards(&_Vault.CallOpts)
}

// PendingRewards is a free data retrieval call binding the contract method 0xeded3fda.
//
// Solidity: function pendingRewards() view returns(uint256)
func (_Vault *VaultCallerSession) PendingRewards() (*big.Int, error) {
	return _Vault.Contract.PendingRewards(&_Vault.CallOpts)
}

// PendingVAIVaultImplementation is a free data retrieval call binding the contract method 0x211de6b6.
//
// Solidity: function pendingVAIVaultImplementation() view returns(address)
func (_Vault *VaultCaller) PendingVAIVaultImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "pendingVAIVaultImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingVAIVaultImplementation is a free data retrieval call binding the contract method 0x211de6b6.
//
// Solidity: function pendingVAIVaultImplementation() view returns(address)
func (_Vault *VaultSession) PendingVAIVaultImplementation() (common.Address, error) {
	return _Vault.Contract.PendingVAIVaultImplementation(&_Vault.CallOpts)
}

// PendingVAIVaultImplementation is a free data retrieval call binding the contract method 0x211de6b6.
//
// Solidity: function pendingVAIVaultImplementation() view returns(address)
func (_Vault *VaultCallerSession) PendingVAIVaultImplementation() (common.Address, error) {
	return _Vault.Contract.PendingVAIVaultImplementation(&_Vault.CallOpts)
}

// PendingXVS is a free data retrieval call binding the contract method 0x79c3b944.
//
// Solidity: function pendingXVS(address _user) view returns(uint256)
func (_Vault *VaultCaller) PendingXVS(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "pendingXVS", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingXVS is a free data retrieval call binding the contract method 0x79c3b944.
//
// Solidity: function pendingXVS(address _user) view returns(uint256)
func (_Vault *VaultSession) PendingXVS(_user common.Address) (*big.Int, error) {
	return _Vault.Contract.PendingXVS(&_Vault.CallOpts, _user)
}

// PendingXVS is a free data retrieval call binding the contract method 0x79c3b944.
//
// Solidity: function pendingXVS(address _user) view returns(uint256)
func (_Vault *VaultCallerSession) PendingXVS(_user common.Address) (*big.Int, error) {
	return _Vault.Contract.PendingXVS(&_Vault.CallOpts, _user)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Vault *VaultCaller) UserInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "userInfo", arg0)

	outstruct := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Vault *VaultSession) UserInfo(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Vault.Contract.UserInfo(&_Vault.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Vault *VaultCallerSession) UserInfo(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Vault.Contract.UserInfo(&_Vault.CallOpts, arg0)
}

// Vai is a free data retrieval call binding the contract method 0xb62e4c92.
//
// Solidity: function vai() view returns(address)
func (_Vault *VaultCaller) Vai(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "vai")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vai is a free data retrieval call binding the contract method 0xb62e4c92.
//
// Solidity: function vai() view returns(address)
func (_Vault *VaultSession) Vai() (common.Address, error) {
	return _Vault.Contract.Vai(&_Vault.CallOpts)
}

// Vai is a free data retrieval call binding the contract method 0xb62e4c92.
//
// Solidity: function vai() view returns(address)
func (_Vault *VaultCallerSession) Vai() (common.Address, error) {
	return _Vault.Contract.Vai(&_Vault.CallOpts)
}

// VaiVaultImplementation is a free data retrieval call binding the contract method 0xf661bb86.
//
// Solidity: function vaiVaultImplementation() view returns(address)
func (_Vault *VaultCaller) VaiVaultImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "vaiVaultImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaiVaultImplementation is a free data retrieval call binding the contract method 0xf661bb86.
//
// Solidity: function vaiVaultImplementation() view returns(address)
func (_Vault *VaultSession) VaiVaultImplementation() (common.Address, error) {
	return _Vault.Contract.VaiVaultImplementation(&_Vault.CallOpts)
}

// VaiVaultImplementation is a free data retrieval call binding the contract method 0xf661bb86.
//
// Solidity: function vaiVaultImplementation() view returns(address)
func (_Vault *VaultCallerSession) VaiVaultImplementation() (common.Address, error) {
	return _Vault.Contract.VaiVaultImplementation(&_Vault.CallOpts)
}

// Xvs is a free data retrieval call binding the contract method 0x4e79ed3c.
//
// Solidity: function xvs() view returns(address)
func (_Vault *VaultCaller) Xvs(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "xvs")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Xvs is a free data retrieval call binding the contract method 0x4e79ed3c.
//
// Solidity: function xvs() view returns(address)
func (_Vault *VaultSession) Xvs() (common.Address, error) {
	return _Vault.Contract.Xvs(&_Vault.CallOpts)
}

// Xvs is a free data retrieval call binding the contract method 0x4e79ed3c.
//
// Solidity: function xvs() view returns(address)
func (_Vault *VaultCallerSession) Xvs() (common.Address, error) {
	return _Vault.Contract.Xvs(&_Vault.CallOpts)
}

// XvsBalance is a free data retrieval call binding the contract method 0x761692ba.
//
// Solidity: function xvsBalance() view returns(uint256)
func (_Vault *VaultCaller) XvsBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "xvsBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XvsBalance is a free data retrieval call binding the contract method 0x761692ba.
//
// Solidity: function xvsBalance() view returns(uint256)
func (_Vault *VaultSession) XvsBalance() (*big.Int, error) {
	return _Vault.Contract.XvsBalance(&_Vault.CallOpts)
}

// XvsBalance is a free data retrieval call binding the contract method 0x761692ba.
//
// Solidity: function xvsBalance() view returns(uint256)
func (_Vault *VaultCallerSession) XvsBalance() (*big.Int, error) {
	return _Vault.Contract.XvsBalance(&_Vault.CallOpts)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address vaiVaultProxy) returns()
func (_Vault *VaultTransactor) Become(opts *bind.TransactOpts, vaiVaultProxy common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "_become", vaiVaultProxy)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address vaiVaultProxy) returns()
func (_Vault *VaultSession) Become(vaiVaultProxy common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Become(&_Vault.TransactOpts, vaiVaultProxy)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address vaiVaultProxy) returns()
func (_Vault *VaultTransactorSession) Become(vaiVaultProxy common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Become(&_Vault.TransactOpts, vaiVaultProxy)
}

// BurnAdmin is a paid mutator transaction binding the contract method 0x81bdf98c.
//
// Solidity: function burnAdmin() returns()
func (_Vault *VaultTransactor) BurnAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "burnAdmin")
}

// BurnAdmin is a paid mutator transaction binding the contract method 0x81bdf98c.
//
// Solidity: function burnAdmin() returns()
func (_Vault *VaultSession) BurnAdmin() (*types.Transaction, error) {
	return _Vault.Contract.BurnAdmin(&_Vault.TransactOpts)
}

// BurnAdmin is a paid mutator transaction binding the contract method 0x81bdf98c.
//
// Solidity: function burnAdmin() returns()
func (_Vault *VaultTransactorSession) BurnAdmin() (*types.Transaction, error) {
	return _Vault.Contract.BurnAdmin(&_Vault.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Vault *VaultTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Vault *VaultSession) Claim() (*types.Transaction, error) {
	return _Vault.Contract.Claim(&_Vault.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Vault *VaultTransactorSession) Claim() (*types.Transaction, error) {
	return _Vault.Contract.Claim(&_Vault.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Vault *VaultTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Vault *VaultSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Vault *VaultTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, _amount)
}

// SetNewAdmin is a paid mutator transaction binding the contract method 0x8eec99c8.
//
// Solidity: function setNewAdmin(address newAdmin) returns()
func (_Vault *VaultTransactor) SetNewAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setNewAdmin", newAdmin)
}

// SetNewAdmin is a paid mutator transaction binding the contract method 0x8eec99c8.
//
// Solidity: function setNewAdmin(address newAdmin) returns()
func (_Vault *VaultSession) SetNewAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetNewAdmin(&_Vault.TransactOpts, newAdmin)
}

// SetNewAdmin is a paid mutator transaction binding the contract method 0x8eec99c8.
//
// Solidity: function setNewAdmin(address newAdmin) returns()
func (_Vault *VaultTransactorSession) SetNewAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetNewAdmin(&_Vault.TransactOpts, newAdmin)
}

// SetVenusInfo is a paid mutator transaction binding the contract method 0x44c0e8ee.
//
// Solidity: function setVenusInfo(address _xvs, address _vai) returns()
func (_Vault *VaultTransactor) SetVenusInfo(opts *bind.TransactOpts, _xvs common.Address, _vai common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setVenusInfo", _xvs, _vai)
}

// SetVenusInfo is a paid mutator transaction binding the contract method 0x44c0e8ee.
//
// Solidity: function setVenusInfo(address _xvs, address _vai) returns()
func (_Vault *VaultSession) SetVenusInfo(_xvs common.Address, _vai common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetVenusInfo(&_Vault.TransactOpts, _xvs, _vai)
}

// SetVenusInfo is a paid mutator transaction binding the contract method 0x44c0e8ee.
//
// Solidity: function setVenusInfo(address _xvs, address _vai) returns()
func (_Vault *VaultTransactorSession) SetVenusInfo(_xvs common.Address, _vai common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetVenusInfo(&_Vault.TransactOpts, _xvs, _vai)
}

// UpdatePendingRewards is a paid mutator transaction binding the contract method 0xfaa1809e.
//
// Solidity: function updatePendingRewards() returns()
func (_Vault *VaultTransactor) UpdatePendingRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "updatePendingRewards")
}

// UpdatePendingRewards is a paid mutator transaction binding the contract method 0xfaa1809e.
//
// Solidity: function updatePendingRewards() returns()
func (_Vault *VaultSession) UpdatePendingRewards() (*types.Transaction, error) {
	return _Vault.Contract.UpdatePendingRewards(&_Vault.TransactOpts)
}

// UpdatePendingRewards is a paid mutator transaction binding the contract method 0xfaa1809e.
//
// Solidity: function updatePendingRewards() returns()
func (_Vault *VaultTransactorSession) UpdatePendingRewards() (*types.Transaction, error) {
	return _Vault.Contract.UpdatePendingRewards(&_Vault.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Vault *VaultTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Vault *VaultSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Vault *VaultTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, _amount)
}

// VaultAdminTransferedIterator is returned from FilterAdminTransfered and is used to iterate over the raw logs and unpacked data for AdminTransfered events raised by the Vault contract.
type VaultAdminTransferedIterator struct {
	Event *VaultAdminTransfered // Event containing the contract specifics and raw log

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
func (it *VaultAdminTransferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultAdminTransfered)
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
		it.Event = new(VaultAdminTransfered)
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
func (it *VaultAdminTransferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultAdminTransferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultAdminTransfered represents a AdminTransfered event raised by the Vault contract.
type VaultAdminTransfered struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminTransfered is a free log retrieval operation binding the contract event 0x173de3514d8508f36ce8c81d509adcd8c8c76098400f685d3042b36f9a4160c3.
//
// Solidity: event AdminTransfered(address indexed oldAdmin, address indexed newAdmin)
func (_Vault *VaultFilterer) FilterAdminTransfered(opts *bind.FilterOpts, oldAdmin []common.Address, newAdmin []common.Address) (*VaultAdminTransferedIterator, error) {

	var oldAdminRule []interface{}
	for _, oldAdminItem := range oldAdmin {
		oldAdminRule = append(oldAdminRule, oldAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "AdminTransfered", oldAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &VaultAdminTransferedIterator{contract: _Vault.contract, event: "AdminTransfered", logs: logs, sub: sub}, nil
}

// WatchAdminTransfered is a free log subscription operation binding the contract event 0x173de3514d8508f36ce8c81d509adcd8c8c76098400f685d3042b36f9a4160c3.
//
// Solidity: event AdminTransfered(address indexed oldAdmin, address indexed newAdmin)
func (_Vault *VaultFilterer) WatchAdminTransfered(opts *bind.WatchOpts, sink chan<- *VaultAdminTransfered, oldAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var oldAdminRule []interface{}
	for _, oldAdminItem := range oldAdmin {
		oldAdminRule = append(oldAdminRule, oldAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "AdminTransfered", oldAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultAdminTransfered)
				if err := _Vault.contract.UnpackLog(event, "AdminTransfered", log); err != nil {
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

// ParseAdminTransfered is a log parse operation binding the contract event 0x173de3514d8508f36ce8c81d509adcd8c8c76098400f685d3042b36f9a4160c3.
//
// Solidity: event AdminTransfered(address indexed oldAdmin, address indexed newAdmin)
func (_Vault *VaultFilterer) ParseAdminTransfered(log types.Log) (*VaultAdminTransfered, error) {
	event := new(VaultAdminTransfered)
	if err := _Vault.contract.UnpackLog(event, "AdminTransfered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Vault contract.
type VaultDepositIterator struct {
	Event *VaultDeposit // Event containing the contract specifics and raw log

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
func (it *VaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultDeposit)
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
		it.Event = new(VaultDeposit)
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
func (it *VaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultDeposit represents a Deposit event raised by the Vault contract.
type VaultDeposit struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Vault *VaultFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*VaultDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &VaultDepositIterator{contract: _Vault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Vault *VaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *VaultDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultDeposit)
				if err := _Vault.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Vault *VaultFilterer) ParseDeposit(log types.Log) (*VaultDeposit, error) {
	event := new(VaultDeposit)
	if err := _Vault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Vault contract.
type VaultWithdrawIterator struct {
	Event *VaultWithdraw // Event containing the contract specifics and raw log

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
func (it *VaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultWithdraw)
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
		it.Event = new(VaultWithdraw)
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
func (it *VaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultWithdraw represents a Withdraw event raised by the Vault contract.
type VaultWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Vault *VaultFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*VaultWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &VaultWithdrawIterator{contract: _Vault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Vault *VaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *VaultWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultWithdraw)
				if err := _Vault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Vault *VaultFilterer) ParseWithdraw(log types.Log) (*VaultWithdraw, error) {
	event := new(VaultWithdraw)
	if err := _Vault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

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

// VaiComtrollerABI is the input ABI used to generate the binding from.
const VaiComtrollerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"LiquidateVAI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"MintFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintVAIAmount\",\"type\":\"uint256\"}],\"name\":\"MintVAI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"oldComptroller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"NewComptroller\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTreasuryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTreasuryAddress\",\"type\":\"address\"}],\"name\":\"NewTreasuryAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTreasuryGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTreasuryGuardian\",\"type\":\"address\"}],\"name\":\"NewTreasuryGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTreasuryPercent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTreasuryPercent\",\"type\":\"uint256\"}],\"name\":\"NewTreasuryPercent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayVAIAmount\",\"type\":\"uint256\"}],\"name\":\"RepayVAI\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVAIUnitroller\",\"name\":\"unitroller\",\"type\":\"address\"}],\"name\":\"_become\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"_initializeVenusVAIState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"comptroller_\",\"type\":\"address\"}],\"name\":\"_setComptroller\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTreasuryGuardian\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newTreasuryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newTreasuryPercent\",\"type\":\"uint256\"}],\"name\":\"_setTreasuryData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vaiMinter\",\"type\":\"address\"}],\"name\":\"calcDistributeVAIMinterVenus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptroller\",\"outputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"getMintableVAI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVAIAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isVenusVAIInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractVTokenInterface\",\"name\":\"vTokenCollateral\",\"type\":\"address\"}],\"name\":\"liquidateVAI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintVAIAmount\",\"type\":\"uint256\"}],\"name\":\"mintVAI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingVAIControllerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"repayVAIAmount\",\"type\":\"uint256\"}],\"name\":\"repayVAI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateVenusVAIMintIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaiControllerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"venusInitialIndex\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"\",\"type\":\"uint224\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusVAIMinterIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"venusVAIState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VaiComtrollerBin is the compiled bytecode used for deploying new contracts.
var VaiComtrollerBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b031916331790556105e3806100326000396000f3fe60806040526004361061007a5760003560e01c8063c1e803341161004e578063c1e803341461019d578063e992a041146101b2578063e9c714f2146101e5578063f851a440146101fa5761007a565b80623b5884146100fd578063267822471461012e578063b06bb42614610143578063b71d1a0c14610158575b6002546040516000916001600160a01b031690829036908083838082843760405192019450600093509091505080830381855af49150503d80600081146100dd576040519150601f19603f3d011682016040523d82523d6000602084013e6100e2565b606091505b505090506040513d6000823e8180156100f9573d82f35b3d82fd5b34801561010957600080fd5b5061011261020f565b604080516001600160a01b039092168252519081900360200190f35b34801561013a57600080fd5b5061011261021e565b34801561014f57600080fd5b5061011261022d565b34801561016457600080fd5b5061018b6004803603602081101561017b57600080fd5b50356001600160a01b031661023c565b60408051918252519081900360200190f35b3480156101a957600080fd5b5061018b6102cd565b3480156101be57600080fd5b5061018b600480360360208110156101d557600080fd5b50356001600160a01b03166103c9565b3480156101f157600080fd5b5061018b61044c565b34801561020657600080fd5b50610112610532565b6002546001600160a01b031681565b6001546001600160a01b031681565b6003546001600160a01b031681565b600080546001600160a01b031633146102625761025b60016000610541565b90506102c8565b600180546001600160a01b038481166001600160a01b0319831681179093556040805191909216808252602082019390935281517fca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9929181900390910190a160005b9150505b919050565b6003546000906001600160a01b0316331415806102f357506003546001600160a01b0316155b1561030b5761030460016004610541565b90506103c6565b60028054600380546001600160a01b038082166001600160a01b031980861682179687905590921690925560408051938316808552949092166020840152815190927fd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a92908290030190a1600354604080516001600160a01b038085168252909216602083015280517fe945ccee5d701fc83f9b8aa8ca94ea4219ec1fcbd4f4cab4f0ea57c5c3e1d8159281900390910190a160005b925050505b90565b600080546001600160a01b031633146103e75761025b600180610541565b600380546001600160a01b038481166001600160a01b0319831617928390556040805192821680845293909116602083015280517fe945ccee5d701fc83f9b8aa8ca94ea4219ec1fcbd4f4cab4f0ea57c5c3e1d8159281900390910190a160006102c4565b6001546000906001600160a01b031633141580610467575033155b156104785761030460016003610541565b60008054600180546001600160a01b038082166001600160a01b031980861682179687905590921690925560408051938316808552949092166020840152815190927ff9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc92908290030190a1600154604080516001600160a01b038085168252909216602083015280517fca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a99281900390910190a160006103c1565b6000546001600160a01b031681565b60007f45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa083600681111561057057fe5b83600681111561057c57fe5b604080519283526020830191909152600082820152519081900360600190a18260068111156105a757fe5b939250505056fea265627a7a72315820f80a26062e16402e92475a12e2c595652e33c933793fb67be676c9603eec453764736f6c63430005110032"

// DeployVaiComtroller deploys a new Ethereum contract, binding an instance of VaiComtroller to it.
func DeployVaiComtroller(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VaiComtroller, error) {
	parsed, err := abi.JSON(strings.NewReader(VaiComtrollerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaiComtrollerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VaiComtroller{VaiComtrollerCaller: VaiComtrollerCaller{contract: contract}, VaiComtrollerTransactor: VaiComtrollerTransactor{contract: contract}, VaiComtrollerFilterer: VaiComtrollerFilterer{contract: contract}}, nil
}

// VaiComtroller is an auto generated Go binding around an Ethereum contract.
type VaiComtroller struct {
	VaiComtrollerCaller     // Read-only binding to the contract
	VaiComtrollerTransactor // Write-only binding to the contract
	VaiComtrollerFilterer   // Log filterer for contract events
}

// VaiComtrollerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaiComtrollerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaiComtrollerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaiComtrollerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaiComtrollerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaiComtrollerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaiComtrollerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaiComtrollerSession struct {
	Contract     *VaiComtroller    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaiComtrollerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaiComtrollerCallerSession struct {
	Contract *VaiComtrollerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// VaiComtrollerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaiComtrollerTransactorSession struct {
	Contract     *VaiComtrollerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VaiComtrollerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaiComtrollerRaw struct {
	Contract *VaiComtroller // Generic contract binding to access the raw methods on
}

// VaiComtrollerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaiComtrollerCallerRaw struct {
	Contract *VaiComtrollerCaller // Generic read-only contract binding to access the raw methods on
}

// VaiComtrollerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaiComtrollerTransactorRaw struct {
	Contract *VaiComtrollerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaiComtroller creates a new instance of VaiComtroller, bound to a specific deployed contract.
func NewVaiComtroller(address common.Address, backend bind.ContractBackend) (*VaiComtroller, error) {
	contract, err := bindVaiComtroller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaiComtroller{VaiComtrollerCaller: VaiComtrollerCaller{contract: contract}, VaiComtrollerTransactor: VaiComtrollerTransactor{contract: contract}, VaiComtrollerFilterer: VaiComtrollerFilterer{contract: contract}}, nil
}

// NewVaiComtrollerCaller creates a new read-only instance of VaiComtroller, bound to a specific deployed contract.
func NewVaiComtrollerCaller(address common.Address, caller bind.ContractCaller) (*VaiComtrollerCaller, error) {
	contract, err := bindVaiComtroller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaiComtrollerCaller{contract: contract}, nil
}

// NewVaiComtrollerTransactor creates a new write-only instance of VaiComtroller, bound to a specific deployed contract.
func NewVaiComtrollerTransactor(address common.Address, transactor bind.ContractTransactor) (*VaiComtrollerTransactor, error) {
	contract, err := bindVaiComtroller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaiComtrollerTransactor{contract: contract}, nil
}

// NewVaiComtrollerFilterer creates a new log filterer instance of VaiComtroller, bound to a specific deployed contract.
func NewVaiComtrollerFilterer(address common.Address, filterer bind.ContractFilterer) (*VaiComtrollerFilterer, error) {
	contract, err := bindVaiComtroller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaiComtrollerFilterer{contract: contract}, nil
}

// bindVaiComtroller binds a generic wrapper to an already deployed contract.
func bindVaiComtroller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaiComtrollerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaiComtroller *VaiComtrollerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaiComtroller.Contract.VaiComtrollerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaiComtroller *VaiComtrollerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaiComtroller.Contract.VaiComtrollerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaiComtroller *VaiComtrollerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaiComtroller.Contract.VaiComtrollerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaiComtroller *VaiComtrollerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaiComtroller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaiComtroller *VaiComtrollerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaiComtroller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaiComtroller *VaiComtrollerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaiComtroller.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_VaiComtroller *VaiComtrollerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaiComtroller.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_VaiComtroller *VaiComtrollerSession) Admin() (common.Address, error) {
	return _VaiComtroller.Contract.Admin(&_VaiComtroller.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_VaiComtroller *VaiComtrollerCallerSession) Admin() (common.Address, error) {
	return _VaiComtroller.Contract.Admin(&_VaiComtroller.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_VaiComtroller *VaiComtrollerCaller) Comptroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaiComtroller.contract.Call(opts, &out, "comptroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_VaiComtroller *VaiComtrollerSession) Comptroller() (common.Address, error) {
	return _VaiComtroller.Contract.Comptroller(&_VaiComtroller.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_VaiComtroller *VaiComtrollerCallerSession) Comptroller() (common.Address, error) {
	return _VaiComtroller.Contract.Comptroller(&_VaiComtroller.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_VaiComtroller *VaiComtrollerCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaiComtroller.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_VaiComtroller *VaiComtrollerSession) GetBlockNumber() (*big.Int, error) {
	return _VaiComtroller.Contract.GetBlockNumber(&_VaiComtroller.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_VaiComtroller *VaiComtrollerCallerSession) GetBlockNumber() (*big.Int, error) {
	return _VaiComtroller.Contract.GetBlockNumber(&_VaiComtroller.CallOpts)
}

// GetMintableVAI is a free data retrieval call binding the contract method 0x3785d1d6.
//
// Solidity: function getMintableVAI(address minter) view returns(uint256, uint256)
func (_VaiComtroller *VaiComtrollerCaller) GetMintableVAI(opts *bind.CallOpts, minter common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _VaiComtroller.contract.Call(opts, &out, "getMintableVAI", minter)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetMintableVAI is a free data retrieval call binding the contract method 0x3785d1d6.
//
// Solidity: function getMintableVAI(address minter) view returns(uint256, uint256)
func (_VaiComtroller *VaiComtrollerSession) GetMintableVAI(minter common.Address) (*big.Int, *big.Int, error) {
	return _VaiComtroller.Contract.GetMintableVAI(&_VaiComtroller.CallOpts, minter)
}

// GetMintableVAI is a free data retrieval call binding the contract method 0x3785d1d6.
//
// Solidity: function getMintableVAI(address minter) view returns(uint256, uint256)
func (_VaiComtroller *VaiComtrollerCallerSession) GetMintableVAI(minter common.Address) (*big.Int, *big.Int, error) {
	return _VaiComtroller.Contract.GetMintableVAI(&_VaiComtroller.CallOpts, minter)
}

// GetVAIAddress is a free data retrieval call binding the contract method 0xcbeb2b28.
//
// Solidity: function getVAIAddress() view returns(address)
func (_VaiComtroller *VaiComtrollerCaller) GetVAIAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaiComtroller.contract.Call(opts, &out, "getVAIAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVAIAddress is a free data retrieval call binding the contract method 0xcbeb2b28.
//
// Solidity: function getVAIAddress() view returns(address)
func (_VaiComtroller *VaiComtrollerSession) GetVAIAddress() (common.Address, error) {
	return _VaiComtroller.Contract.GetVAIAddress(&_VaiComtroller.CallOpts)
}

// GetVAIAddress is a free data retrieval call binding the contract method 0xcbeb2b28.
//
// Solidity: function getVAIAddress() view returns(address)
func (_VaiComtroller *VaiComtrollerCallerSession) GetVAIAddress() (common.Address, error) {
	return _VaiComtroller.Contract.GetVAIAddress(&_VaiComtroller.CallOpts)
}

// IsVenusVAIInitialized is a free data retrieval call binding the contract method 0x60c954ef.
//
// Solidity: function isVenusVAIInitialized() view returns(bool)
func (_VaiComtroller *VaiComtrollerCaller) IsVenusVAIInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaiComtroller.contract.Call(opts, &out, "isVenusVAIInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVenusVAIInitialized is a free data retrieval call binding the contract method 0x60c954ef.
//
// Solidity: function isVenusVAIInitialized() view returns(bool)
func (_VaiComtroller *VaiComtrollerSession) IsVenusVAIInitialized() (bool, error) {
	return _VaiComtroller.Contract.IsVenusVAIInitialized(&_VaiComtroller.CallOpts)
}

// IsVenusVAIInitialized is a free data retrieval call binding the contract method 0x60c954ef.
//
// Solidity: function isVenusVAIInitialized() view returns(bool)
func (_VaiComtroller *VaiComtrollerCallerSession) IsVenusVAIInitialized() (bool, error) {
	return _VaiComtroller.Contract.IsVenusVAIInitialized(&_VaiComtroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_VaiComtroller *VaiComtrollerCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaiComtroller.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_VaiComtroller *VaiComtrollerSession) PendingAdmin() (common.Address, error) {
	return _VaiComtroller.Contract.PendingAdmin(&_VaiComtroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_VaiComtroller *VaiComtrollerCallerSession) PendingAdmin() (common.Address, error) {
	return _VaiComtroller.Contract.PendingAdmin(&_VaiComtroller.CallOpts)
}

// PendingVAIControllerImplementation is a free data retrieval call binding the contract method 0xb06bb426.
//
// Solidity: function pendingVAIControllerImplementation() view returns(address)
func (_VaiComtroller *VaiComtrollerCaller) PendingVAIControllerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaiComtroller.contract.Call(opts, &out, "pendingVAIControllerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingVAIControllerImplementation is a free data retrieval call binding the contract method 0xb06bb426.
//
// Solidity: function pendingVAIControllerImplementation() view returns(address)
func (_VaiComtroller *VaiComtrollerSession) PendingVAIControllerImplementation() (common.Address, error) {
	return _VaiComtroller.Contract.PendingVAIControllerImplementation(&_VaiComtroller.CallOpts)
}

// PendingVAIControllerImplementation is a free data retrieval call binding the contract method 0xb06bb426.
//
// Solidity: function pendingVAIControllerImplementation() view returns(address)
func (_VaiComtroller *VaiComtrollerCallerSession) PendingVAIControllerImplementation() (common.Address, error) {
	return _VaiComtroller.Contract.PendingVAIControllerImplementation(&_VaiComtroller.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_VaiComtroller *VaiComtrollerCaller) TreasuryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaiComtroller.contract.Call(opts, &out, "treasuryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_VaiComtroller *VaiComtrollerSession) TreasuryAddress() (common.Address, error) {
	return _VaiComtroller.Contract.TreasuryAddress(&_VaiComtroller.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_VaiComtroller *VaiComtrollerCallerSession) TreasuryAddress() (common.Address, error) {
	return _VaiComtroller.Contract.TreasuryAddress(&_VaiComtroller.CallOpts)
}

// TreasuryGuardian is a free data retrieval call binding the contract method 0xb2eafc39.
//
// Solidity: function treasuryGuardian() view returns(address)
func (_VaiComtroller *VaiComtrollerCaller) TreasuryGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaiComtroller.contract.Call(opts, &out, "treasuryGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryGuardian is a free data retrieval call binding the contract method 0xb2eafc39.
//
// Solidity: function treasuryGuardian() view returns(address)
func (_VaiComtroller *VaiComtrollerSession) TreasuryGuardian() (common.Address, error) {
	return _VaiComtroller.Contract.TreasuryGuardian(&_VaiComtroller.CallOpts)
}

// TreasuryGuardian is a free data retrieval call binding the contract method 0xb2eafc39.
//
// Solidity: function treasuryGuardian() view returns(address)
func (_VaiComtroller *VaiComtrollerCallerSession) TreasuryGuardian() (common.Address, error) {
	return _VaiComtroller.Contract.TreasuryGuardian(&_VaiComtroller.CallOpts)
}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_VaiComtroller *VaiComtrollerCaller) TreasuryPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaiComtroller.contract.Call(opts, &out, "treasuryPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_VaiComtroller *VaiComtrollerSession) TreasuryPercent() (*big.Int, error) {
	return _VaiComtroller.Contract.TreasuryPercent(&_VaiComtroller.CallOpts)
}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_VaiComtroller *VaiComtrollerCallerSession) TreasuryPercent() (*big.Int, error) {
	return _VaiComtroller.Contract.TreasuryPercent(&_VaiComtroller.CallOpts)
}

// VaiControllerImplementation is a free data retrieval call binding the contract method 0x003b5884.
//
// Solidity: function vaiControllerImplementation() view returns(address)
func (_VaiComtroller *VaiComtrollerCaller) VaiControllerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaiComtroller.contract.Call(opts, &out, "vaiControllerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaiControllerImplementation is a free data retrieval call binding the contract method 0x003b5884.
//
// Solidity: function vaiControllerImplementation() view returns(address)
func (_VaiComtroller *VaiComtrollerSession) VaiControllerImplementation() (common.Address, error) {
	return _VaiComtroller.Contract.VaiControllerImplementation(&_VaiComtroller.CallOpts)
}

// VaiControllerImplementation is a free data retrieval call binding the contract method 0x003b5884.
//
// Solidity: function vaiControllerImplementation() view returns(address)
func (_VaiComtroller *VaiComtrollerCallerSession) VaiControllerImplementation() (common.Address, error) {
	return _VaiComtroller.Contract.VaiControllerImplementation(&_VaiComtroller.CallOpts)
}

// VenusInitialIndex is a free data retrieval call binding the contract method 0xc5b4db55.
//
// Solidity: function venusInitialIndex() view returns(uint224)
func (_VaiComtroller *VaiComtrollerCaller) VenusInitialIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaiComtroller.contract.Call(opts, &out, "venusInitialIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusInitialIndex is a free data retrieval call binding the contract method 0xc5b4db55.
//
// Solidity: function venusInitialIndex() view returns(uint224)
func (_VaiComtroller *VaiComtrollerSession) VenusInitialIndex() (*big.Int, error) {
	return _VaiComtroller.Contract.VenusInitialIndex(&_VaiComtroller.CallOpts)
}

// VenusInitialIndex is a free data retrieval call binding the contract method 0xc5b4db55.
//
// Solidity: function venusInitialIndex() view returns(uint224)
func (_VaiComtroller *VaiComtrollerCallerSession) VenusInitialIndex() (*big.Int, error) {
	return _VaiComtroller.Contract.VenusInitialIndex(&_VaiComtroller.CallOpts)
}

// VenusVAIMinterIndex is a free data retrieval call binding the contract method 0x24650602.
//
// Solidity: function venusVAIMinterIndex(address ) view returns(uint256)
func (_VaiComtroller *VaiComtrollerCaller) VenusVAIMinterIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaiComtroller.contract.Call(opts, &out, "venusVAIMinterIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusVAIMinterIndex is a free data retrieval call binding the contract method 0x24650602.
//
// Solidity: function venusVAIMinterIndex(address ) view returns(uint256)
func (_VaiComtroller *VaiComtrollerSession) VenusVAIMinterIndex(arg0 common.Address) (*big.Int, error) {
	return _VaiComtroller.Contract.VenusVAIMinterIndex(&_VaiComtroller.CallOpts, arg0)
}

// VenusVAIMinterIndex is a free data retrieval call binding the contract method 0x24650602.
//
// Solidity: function venusVAIMinterIndex(address ) view returns(uint256)
func (_VaiComtroller *VaiComtrollerCallerSession) VenusVAIMinterIndex(arg0 common.Address) (*big.Int, error) {
	return _VaiComtroller.Contract.VenusVAIMinterIndex(&_VaiComtroller.CallOpts, arg0)
}

// VenusVAIState is a free data retrieval call binding the contract method 0xe44e6168.
//
// Solidity: function venusVAIState() view returns(uint224 index, uint32 block)
func (_VaiComtroller *VaiComtrollerCaller) VenusVAIState(opts *bind.CallOpts) (struct {
	Index *big.Int
	Block uint32
}, error) {
	var out []interface{}
	err := _VaiComtroller.contract.Call(opts, &out, "venusVAIState")

	outstruct := new(struct {
		Index *big.Int
		Block uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Block = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// VenusVAIState is a free data retrieval call binding the contract method 0xe44e6168.
//
// Solidity: function venusVAIState() view returns(uint224 index, uint32 block)
func (_VaiComtroller *VaiComtrollerSession) VenusVAIState() (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _VaiComtroller.Contract.VenusVAIState(&_VaiComtroller.CallOpts)
}

// VenusVAIState is a free data retrieval call binding the contract method 0xe44e6168.
//
// Solidity: function venusVAIState() view returns(uint224 index, uint32 block)
func (_VaiComtroller *VaiComtrollerCallerSession) VenusVAIState() (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _VaiComtroller.Contract.VenusVAIState(&_VaiComtroller.CallOpts)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_VaiComtroller *VaiComtrollerTransactor) Become(opts *bind.TransactOpts, unitroller common.Address) (*types.Transaction, error) {
	return _VaiComtroller.contract.Transact(opts, "_become", unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_VaiComtroller *VaiComtrollerSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _VaiComtroller.Contract.Become(&_VaiComtroller.TransactOpts, unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_VaiComtroller *VaiComtrollerTransactorSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _VaiComtroller.Contract.Become(&_VaiComtroller.TransactOpts, unitroller)
}

// InitializeVenusVAIState is a paid mutator transaction binding the contract method 0x5035920a.
//
// Solidity: function _initializeVenusVAIState(uint256 blockNumber) returns(uint256)
func (_VaiComtroller *VaiComtrollerTransactor) InitializeVenusVAIState(opts *bind.TransactOpts, blockNumber *big.Int) (*types.Transaction, error) {
	return _VaiComtroller.contract.Transact(opts, "_initializeVenusVAIState", blockNumber)
}

// InitializeVenusVAIState is a paid mutator transaction binding the contract method 0x5035920a.
//
// Solidity: function _initializeVenusVAIState(uint256 blockNumber) returns(uint256)
func (_VaiComtroller *VaiComtrollerSession) InitializeVenusVAIState(blockNumber *big.Int) (*types.Transaction, error) {
	return _VaiComtroller.Contract.InitializeVenusVAIState(&_VaiComtroller.TransactOpts, blockNumber)
}

// InitializeVenusVAIState is a paid mutator transaction binding the contract method 0x5035920a.
//
// Solidity: function _initializeVenusVAIState(uint256 blockNumber) returns(uint256)
func (_VaiComtroller *VaiComtrollerTransactorSession) InitializeVenusVAIState(blockNumber *big.Int) (*types.Transaction, error) {
	return _VaiComtroller.Contract.InitializeVenusVAIState(&_VaiComtroller.TransactOpts, blockNumber)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address comptroller_) returns(uint256)
func (_VaiComtroller *VaiComtrollerTransactor) SetComptroller(opts *bind.TransactOpts, comptroller_ common.Address) (*types.Transaction, error) {
	return _VaiComtroller.contract.Transact(opts, "_setComptroller", comptroller_)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address comptroller_) returns(uint256)
func (_VaiComtroller *VaiComtrollerSession) SetComptroller(comptroller_ common.Address) (*types.Transaction, error) {
	return _VaiComtroller.Contract.SetComptroller(&_VaiComtroller.TransactOpts, comptroller_)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address comptroller_) returns(uint256)
func (_VaiComtroller *VaiComtrollerTransactorSession) SetComptroller(comptroller_ common.Address) (*types.Transaction, error) {
	return _VaiComtroller.Contract.SetComptroller(&_VaiComtroller.TransactOpts, comptroller_)
}

// SetTreasuryData is a paid mutator transaction binding the contract method 0xd24febad.
//
// Solidity: function _setTreasuryData(address newTreasuryGuardian, address newTreasuryAddress, uint256 newTreasuryPercent) returns(uint256)
func (_VaiComtroller *VaiComtrollerTransactor) SetTreasuryData(opts *bind.TransactOpts, newTreasuryGuardian common.Address, newTreasuryAddress common.Address, newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _VaiComtroller.contract.Transact(opts, "_setTreasuryData", newTreasuryGuardian, newTreasuryAddress, newTreasuryPercent)
}

// SetTreasuryData is a paid mutator transaction binding the contract method 0xd24febad.
//
// Solidity: function _setTreasuryData(address newTreasuryGuardian, address newTreasuryAddress, uint256 newTreasuryPercent) returns(uint256)
func (_VaiComtroller *VaiComtrollerSession) SetTreasuryData(newTreasuryGuardian common.Address, newTreasuryAddress common.Address, newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _VaiComtroller.Contract.SetTreasuryData(&_VaiComtroller.TransactOpts, newTreasuryGuardian, newTreasuryAddress, newTreasuryPercent)
}

// SetTreasuryData is a paid mutator transaction binding the contract method 0xd24febad.
//
// Solidity: function _setTreasuryData(address newTreasuryGuardian, address newTreasuryAddress, uint256 newTreasuryPercent) returns(uint256)
func (_VaiComtroller *VaiComtrollerTransactorSession) SetTreasuryData(newTreasuryGuardian common.Address, newTreasuryAddress common.Address, newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _VaiComtroller.Contract.SetTreasuryData(&_VaiComtroller.TransactOpts, newTreasuryGuardian, newTreasuryAddress, newTreasuryPercent)
}

// CalcDistributeVAIMinterVenus is a paid mutator transaction binding the contract method 0xd178273b.
//
// Solidity: function calcDistributeVAIMinterVenus(address vaiMinter) returns(uint256, uint256, uint256, uint256)
func (_VaiComtroller *VaiComtrollerTransactor) CalcDistributeVAIMinterVenus(opts *bind.TransactOpts, vaiMinter common.Address) (*types.Transaction, error) {
	return _VaiComtroller.contract.Transact(opts, "calcDistributeVAIMinterVenus", vaiMinter)
}

// CalcDistributeVAIMinterVenus is a paid mutator transaction binding the contract method 0xd178273b.
//
// Solidity: function calcDistributeVAIMinterVenus(address vaiMinter) returns(uint256, uint256, uint256, uint256)
func (_VaiComtroller *VaiComtrollerSession) CalcDistributeVAIMinterVenus(vaiMinter common.Address) (*types.Transaction, error) {
	return _VaiComtroller.Contract.CalcDistributeVAIMinterVenus(&_VaiComtroller.TransactOpts, vaiMinter)
}

// CalcDistributeVAIMinterVenus is a paid mutator transaction binding the contract method 0xd178273b.
//
// Solidity: function calcDistributeVAIMinterVenus(address vaiMinter) returns(uint256, uint256, uint256, uint256)
func (_VaiComtroller *VaiComtrollerTransactorSession) CalcDistributeVAIMinterVenus(vaiMinter common.Address) (*types.Transaction, error) {
	return _VaiComtroller.Contract.CalcDistributeVAIMinterVenus(&_VaiComtroller.TransactOpts, vaiMinter)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_VaiComtroller *VaiComtrollerTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaiComtroller.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_VaiComtroller *VaiComtrollerSession) Initialize() (*types.Transaction, error) {
	return _VaiComtroller.Contract.Initialize(&_VaiComtroller.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_VaiComtroller *VaiComtrollerTransactorSession) Initialize() (*types.Transaction, error) {
	return _VaiComtroller.Contract.Initialize(&_VaiComtroller.TransactOpts)
}

// LiquidateVAI is a paid mutator transaction binding the contract method 0x11b3d5e7.
//
// Solidity: function liquidateVAI(address borrower, uint256 repayAmount, address vTokenCollateral) returns(uint256, uint256)
func (_VaiComtroller *VaiComtrollerTransactor) LiquidateVAI(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _VaiComtroller.contract.Transact(opts, "liquidateVAI", borrower, repayAmount, vTokenCollateral)
}

// LiquidateVAI is a paid mutator transaction binding the contract method 0x11b3d5e7.
//
// Solidity: function liquidateVAI(address borrower, uint256 repayAmount, address vTokenCollateral) returns(uint256, uint256)
func (_VaiComtroller *VaiComtrollerSession) LiquidateVAI(borrower common.Address, repayAmount *big.Int, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _VaiComtroller.Contract.LiquidateVAI(&_VaiComtroller.TransactOpts, borrower, repayAmount, vTokenCollateral)
}

// LiquidateVAI is a paid mutator transaction binding the contract method 0x11b3d5e7.
//
// Solidity: function liquidateVAI(address borrower, uint256 repayAmount, address vTokenCollateral) returns(uint256, uint256)
func (_VaiComtroller *VaiComtrollerTransactorSession) LiquidateVAI(borrower common.Address, repayAmount *big.Int, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _VaiComtroller.Contract.LiquidateVAI(&_VaiComtroller.TransactOpts, borrower, repayAmount, vTokenCollateral)
}

// MintVAI is a paid mutator transaction binding the contract method 0x4712ee7d.
//
// Solidity: function mintVAI(uint256 mintVAIAmount) returns(uint256)
func (_VaiComtroller *VaiComtrollerTransactor) MintVAI(opts *bind.TransactOpts, mintVAIAmount *big.Int) (*types.Transaction, error) {
	return _VaiComtroller.contract.Transact(opts, "mintVAI", mintVAIAmount)
}

// MintVAI is a paid mutator transaction binding the contract method 0x4712ee7d.
//
// Solidity: function mintVAI(uint256 mintVAIAmount) returns(uint256)
func (_VaiComtroller *VaiComtrollerSession) MintVAI(mintVAIAmount *big.Int) (*types.Transaction, error) {
	return _VaiComtroller.Contract.MintVAI(&_VaiComtroller.TransactOpts, mintVAIAmount)
}

// MintVAI is a paid mutator transaction binding the contract method 0x4712ee7d.
//
// Solidity: function mintVAI(uint256 mintVAIAmount) returns(uint256)
func (_VaiComtroller *VaiComtrollerTransactorSession) MintVAI(mintVAIAmount *big.Int) (*types.Transaction, error) {
	return _VaiComtroller.Contract.MintVAI(&_VaiComtroller.TransactOpts, mintVAIAmount)
}

// RepayVAI is a paid mutator transaction binding the contract method 0x6fe74a21.
//
// Solidity: function repayVAI(uint256 repayVAIAmount) returns(uint256, uint256)
func (_VaiComtroller *VaiComtrollerTransactor) RepayVAI(opts *bind.TransactOpts, repayVAIAmount *big.Int) (*types.Transaction, error) {
	return _VaiComtroller.contract.Transact(opts, "repayVAI", repayVAIAmount)
}

// RepayVAI is a paid mutator transaction binding the contract method 0x6fe74a21.
//
// Solidity: function repayVAI(uint256 repayVAIAmount) returns(uint256, uint256)
func (_VaiComtroller *VaiComtrollerSession) RepayVAI(repayVAIAmount *big.Int) (*types.Transaction, error) {
	return _VaiComtroller.Contract.RepayVAI(&_VaiComtroller.TransactOpts, repayVAIAmount)
}

// RepayVAI is a paid mutator transaction binding the contract method 0x6fe74a21.
//
// Solidity: function repayVAI(uint256 repayVAIAmount) returns(uint256, uint256)
func (_VaiComtroller *VaiComtrollerTransactorSession) RepayVAI(repayVAIAmount *big.Int) (*types.Transaction, error) {
	return _VaiComtroller.Contract.RepayVAI(&_VaiComtroller.TransactOpts, repayVAIAmount)
}

// UpdateVenusVAIMintIndex is a paid mutator transaction binding the contract method 0xd244b081.
//
// Solidity: function updateVenusVAIMintIndex() returns(uint256)
func (_VaiComtroller *VaiComtrollerTransactor) UpdateVenusVAIMintIndex(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaiComtroller.contract.Transact(opts, "updateVenusVAIMintIndex")
}

// UpdateVenusVAIMintIndex is a paid mutator transaction binding the contract method 0xd244b081.
//
// Solidity: function updateVenusVAIMintIndex() returns(uint256)
func (_VaiComtroller *VaiComtrollerSession) UpdateVenusVAIMintIndex() (*types.Transaction, error) {
	return _VaiComtroller.Contract.UpdateVenusVAIMintIndex(&_VaiComtroller.TransactOpts)
}

// UpdateVenusVAIMintIndex is a paid mutator transaction binding the contract method 0xd244b081.
//
// Solidity: function updateVenusVAIMintIndex() returns(uint256)
func (_VaiComtroller *VaiComtrollerTransactorSession) UpdateVenusVAIMintIndex() (*types.Transaction, error) {
	return _VaiComtroller.Contract.UpdateVenusVAIMintIndex(&_VaiComtroller.TransactOpts)
}

// VaiComtrollerFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the VaiComtroller contract.
type VaiComtrollerFailureIterator struct {
	Event *VaiComtrollerFailure // Event containing the contract specifics and raw log

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
func (it *VaiComtrollerFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaiComtrollerFailure)
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
		it.Event = new(VaiComtrollerFailure)
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
func (it *VaiComtrollerFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaiComtrollerFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaiComtrollerFailure represents a Failure event raised by the VaiComtroller contract.
type VaiComtrollerFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_VaiComtroller *VaiComtrollerFilterer) FilterFailure(opts *bind.FilterOpts) (*VaiComtrollerFailureIterator, error) {

	logs, sub, err := _VaiComtroller.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &VaiComtrollerFailureIterator{contract: _VaiComtroller.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_VaiComtroller *VaiComtrollerFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *VaiComtrollerFailure) (event.Subscription, error) {

	logs, sub, err := _VaiComtroller.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaiComtrollerFailure)
				if err := _VaiComtroller.contract.UnpackLog(event, "Failure", log); err != nil {
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

// ParseFailure is a log parse operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_VaiComtroller *VaiComtrollerFilterer) ParseFailure(log types.Log) (*VaiComtrollerFailure, error) {
	event := new(VaiComtrollerFailure)
	if err := _VaiComtroller.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaiComtrollerLiquidateVAIIterator is returned from FilterLiquidateVAI and is used to iterate over the raw logs and unpacked data for LiquidateVAI events raised by the VaiComtroller contract.
type VaiComtrollerLiquidateVAIIterator struct {
	Event *VaiComtrollerLiquidateVAI // Event containing the contract specifics and raw log

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
func (it *VaiComtrollerLiquidateVAIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaiComtrollerLiquidateVAI)
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
		it.Event = new(VaiComtrollerLiquidateVAI)
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
func (it *VaiComtrollerLiquidateVAIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaiComtrollerLiquidateVAIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaiComtrollerLiquidateVAI represents a LiquidateVAI event raised by the VaiComtroller contract.
type VaiComtrollerLiquidateVAI struct {
	Liquidator       common.Address
	Borrower         common.Address
	RepayAmount      *big.Int
	VTokenCollateral common.Address
	SeizeTokens      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLiquidateVAI is a free log retrieval operation binding the contract event 0x42d401f96718a0c42e5cea8108973f0022677b7e2e5f4ee19851b2de7a0394e7.
//
// Solidity: event LiquidateVAI(address liquidator, address borrower, uint256 repayAmount, address vTokenCollateral, uint256 seizeTokens)
func (_VaiComtroller *VaiComtrollerFilterer) FilterLiquidateVAI(opts *bind.FilterOpts) (*VaiComtrollerLiquidateVAIIterator, error) {

	logs, sub, err := _VaiComtroller.contract.FilterLogs(opts, "LiquidateVAI")
	if err != nil {
		return nil, err
	}
	return &VaiComtrollerLiquidateVAIIterator{contract: _VaiComtroller.contract, event: "LiquidateVAI", logs: logs, sub: sub}, nil
}

// WatchLiquidateVAI is a free log subscription operation binding the contract event 0x42d401f96718a0c42e5cea8108973f0022677b7e2e5f4ee19851b2de7a0394e7.
//
// Solidity: event LiquidateVAI(address liquidator, address borrower, uint256 repayAmount, address vTokenCollateral, uint256 seizeTokens)
func (_VaiComtroller *VaiComtrollerFilterer) WatchLiquidateVAI(opts *bind.WatchOpts, sink chan<- *VaiComtrollerLiquidateVAI) (event.Subscription, error) {

	logs, sub, err := _VaiComtroller.contract.WatchLogs(opts, "LiquidateVAI")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaiComtrollerLiquidateVAI)
				if err := _VaiComtroller.contract.UnpackLog(event, "LiquidateVAI", log); err != nil {
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

// ParseLiquidateVAI is a log parse operation binding the contract event 0x42d401f96718a0c42e5cea8108973f0022677b7e2e5f4ee19851b2de7a0394e7.
//
// Solidity: event LiquidateVAI(address liquidator, address borrower, uint256 repayAmount, address vTokenCollateral, uint256 seizeTokens)
func (_VaiComtroller *VaiComtrollerFilterer) ParseLiquidateVAI(log types.Log) (*VaiComtrollerLiquidateVAI, error) {
	event := new(VaiComtrollerLiquidateVAI)
	if err := _VaiComtroller.contract.UnpackLog(event, "LiquidateVAI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaiComtrollerMintFeeIterator is returned from FilterMintFee and is used to iterate over the raw logs and unpacked data for MintFee events raised by the VaiComtroller contract.
type VaiComtrollerMintFeeIterator struct {
	Event *VaiComtrollerMintFee // Event containing the contract specifics and raw log

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
func (it *VaiComtrollerMintFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaiComtrollerMintFee)
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
		it.Event = new(VaiComtrollerMintFee)
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
func (it *VaiComtrollerMintFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaiComtrollerMintFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaiComtrollerMintFee represents a MintFee event raised by the VaiComtroller contract.
type VaiComtrollerMintFee struct {
	Minter    common.Address
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMintFee is a free log retrieval operation binding the contract event 0xb0715a6d41a37c1b0672c22c09a31a0642c1fb3f9efa2d5fd5c6d2d891ee78c6.
//
// Solidity: event MintFee(address minter, uint256 feeAmount)
func (_VaiComtroller *VaiComtrollerFilterer) FilterMintFee(opts *bind.FilterOpts) (*VaiComtrollerMintFeeIterator, error) {

	logs, sub, err := _VaiComtroller.contract.FilterLogs(opts, "MintFee")
	if err != nil {
		return nil, err
	}
	return &VaiComtrollerMintFeeIterator{contract: _VaiComtroller.contract, event: "MintFee", logs: logs, sub: sub}, nil
}

// WatchMintFee is a free log subscription operation binding the contract event 0xb0715a6d41a37c1b0672c22c09a31a0642c1fb3f9efa2d5fd5c6d2d891ee78c6.
//
// Solidity: event MintFee(address minter, uint256 feeAmount)
func (_VaiComtroller *VaiComtrollerFilterer) WatchMintFee(opts *bind.WatchOpts, sink chan<- *VaiComtrollerMintFee) (event.Subscription, error) {

	logs, sub, err := _VaiComtroller.contract.WatchLogs(opts, "MintFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaiComtrollerMintFee)
				if err := _VaiComtroller.contract.UnpackLog(event, "MintFee", log); err != nil {
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

// ParseMintFee is a log parse operation binding the contract event 0xb0715a6d41a37c1b0672c22c09a31a0642c1fb3f9efa2d5fd5c6d2d891ee78c6.
//
// Solidity: event MintFee(address minter, uint256 feeAmount)
func (_VaiComtroller *VaiComtrollerFilterer) ParseMintFee(log types.Log) (*VaiComtrollerMintFee, error) {
	event := new(VaiComtrollerMintFee)
	if err := _VaiComtroller.contract.UnpackLog(event, "MintFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaiComtrollerMintVAIIterator is returned from FilterMintVAI and is used to iterate over the raw logs and unpacked data for MintVAI events raised by the VaiComtroller contract.
type VaiComtrollerMintVAIIterator struct {
	Event *VaiComtrollerMintVAI // Event containing the contract specifics and raw log

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
func (it *VaiComtrollerMintVAIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaiComtrollerMintVAI)
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
		it.Event = new(VaiComtrollerMintVAI)
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
func (it *VaiComtrollerMintVAIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaiComtrollerMintVAIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaiComtrollerMintVAI represents a MintVAI event raised by the VaiComtroller contract.
type VaiComtrollerMintVAI struct {
	Minter        common.Address
	MintVAIAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMintVAI is a free log retrieval operation binding the contract event 0x002e68ab1600fc5e7290e2ceaa79e2f86b4dbaca84a48421e167e0b40409218a.
//
// Solidity: event MintVAI(address minter, uint256 mintVAIAmount)
func (_VaiComtroller *VaiComtrollerFilterer) FilterMintVAI(opts *bind.FilterOpts) (*VaiComtrollerMintVAIIterator, error) {

	logs, sub, err := _VaiComtroller.contract.FilterLogs(opts, "MintVAI")
	if err != nil {
		return nil, err
	}
	return &VaiComtrollerMintVAIIterator{contract: _VaiComtroller.contract, event: "MintVAI", logs: logs, sub: sub}, nil
}

// WatchMintVAI is a free log subscription operation binding the contract event 0x002e68ab1600fc5e7290e2ceaa79e2f86b4dbaca84a48421e167e0b40409218a.
//
// Solidity: event MintVAI(address minter, uint256 mintVAIAmount)
func (_VaiComtroller *VaiComtrollerFilterer) WatchMintVAI(opts *bind.WatchOpts, sink chan<- *VaiComtrollerMintVAI) (event.Subscription, error) {

	logs, sub, err := _VaiComtroller.contract.WatchLogs(opts, "MintVAI")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaiComtrollerMintVAI)
				if err := _VaiComtroller.contract.UnpackLog(event, "MintVAI", log); err != nil {
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

// ParseMintVAI is a log parse operation binding the contract event 0x002e68ab1600fc5e7290e2ceaa79e2f86b4dbaca84a48421e167e0b40409218a.
//
// Solidity: event MintVAI(address minter, uint256 mintVAIAmount)
func (_VaiComtroller *VaiComtrollerFilterer) ParseMintVAI(log types.Log) (*VaiComtrollerMintVAI, error) {
	event := new(VaiComtrollerMintVAI)
	if err := _VaiComtroller.contract.UnpackLog(event, "MintVAI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaiComtrollerNewComptrollerIterator is returned from FilterNewComptroller and is used to iterate over the raw logs and unpacked data for NewComptroller events raised by the VaiComtroller contract.
type VaiComtrollerNewComptrollerIterator struct {
	Event *VaiComtrollerNewComptroller // Event containing the contract specifics and raw log

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
func (it *VaiComtrollerNewComptrollerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaiComtrollerNewComptroller)
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
		it.Event = new(VaiComtrollerNewComptroller)
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
func (it *VaiComtrollerNewComptrollerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaiComtrollerNewComptrollerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaiComtrollerNewComptroller represents a NewComptroller event raised by the VaiComtroller contract.
type VaiComtrollerNewComptroller struct {
	OldComptroller common.Address
	NewComptroller common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewComptroller is a free log retrieval operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_VaiComtroller *VaiComtrollerFilterer) FilterNewComptroller(opts *bind.FilterOpts) (*VaiComtrollerNewComptrollerIterator, error) {

	logs, sub, err := _VaiComtroller.contract.FilterLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return &VaiComtrollerNewComptrollerIterator{contract: _VaiComtroller.contract, event: "NewComptroller", logs: logs, sub: sub}, nil
}

// WatchNewComptroller is a free log subscription operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_VaiComtroller *VaiComtrollerFilterer) WatchNewComptroller(opts *bind.WatchOpts, sink chan<- *VaiComtrollerNewComptroller) (event.Subscription, error) {

	logs, sub, err := _VaiComtroller.contract.WatchLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaiComtrollerNewComptroller)
				if err := _VaiComtroller.contract.UnpackLog(event, "NewComptroller", log); err != nil {
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

// ParseNewComptroller is a log parse operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_VaiComtroller *VaiComtrollerFilterer) ParseNewComptroller(log types.Log) (*VaiComtrollerNewComptroller, error) {
	event := new(VaiComtrollerNewComptroller)
	if err := _VaiComtroller.contract.UnpackLog(event, "NewComptroller", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaiComtrollerNewTreasuryAddressIterator is returned from FilterNewTreasuryAddress and is used to iterate over the raw logs and unpacked data for NewTreasuryAddress events raised by the VaiComtroller contract.
type VaiComtrollerNewTreasuryAddressIterator struct {
	Event *VaiComtrollerNewTreasuryAddress // Event containing the contract specifics and raw log

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
func (it *VaiComtrollerNewTreasuryAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaiComtrollerNewTreasuryAddress)
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
		it.Event = new(VaiComtrollerNewTreasuryAddress)
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
func (it *VaiComtrollerNewTreasuryAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaiComtrollerNewTreasuryAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaiComtrollerNewTreasuryAddress represents a NewTreasuryAddress event raised by the VaiComtroller contract.
type VaiComtrollerNewTreasuryAddress struct {
	OldTreasuryAddress common.Address
	NewTreasuryAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewTreasuryAddress is a free log retrieval operation binding the contract event 0x8de763046d7b8f08b6c3d03543de1d615309417842bb5d2d62f110f65809ddac.
//
// Solidity: event NewTreasuryAddress(address oldTreasuryAddress, address newTreasuryAddress)
func (_VaiComtroller *VaiComtrollerFilterer) FilterNewTreasuryAddress(opts *bind.FilterOpts) (*VaiComtrollerNewTreasuryAddressIterator, error) {

	logs, sub, err := _VaiComtroller.contract.FilterLogs(opts, "NewTreasuryAddress")
	if err != nil {
		return nil, err
	}
	return &VaiComtrollerNewTreasuryAddressIterator{contract: _VaiComtroller.contract, event: "NewTreasuryAddress", logs: logs, sub: sub}, nil
}

// WatchNewTreasuryAddress is a free log subscription operation binding the contract event 0x8de763046d7b8f08b6c3d03543de1d615309417842bb5d2d62f110f65809ddac.
//
// Solidity: event NewTreasuryAddress(address oldTreasuryAddress, address newTreasuryAddress)
func (_VaiComtroller *VaiComtrollerFilterer) WatchNewTreasuryAddress(opts *bind.WatchOpts, sink chan<- *VaiComtrollerNewTreasuryAddress) (event.Subscription, error) {

	logs, sub, err := _VaiComtroller.contract.WatchLogs(opts, "NewTreasuryAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaiComtrollerNewTreasuryAddress)
				if err := _VaiComtroller.contract.UnpackLog(event, "NewTreasuryAddress", log); err != nil {
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

// ParseNewTreasuryAddress is a log parse operation binding the contract event 0x8de763046d7b8f08b6c3d03543de1d615309417842bb5d2d62f110f65809ddac.
//
// Solidity: event NewTreasuryAddress(address oldTreasuryAddress, address newTreasuryAddress)
func (_VaiComtroller *VaiComtrollerFilterer) ParseNewTreasuryAddress(log types.Log) (*VaiComtrollerNewTreasuryAddress, error) {
	event := new(VaiComtrollerNewTreasuryAddress)
	if err := _VaiComtroller.contract.UnpackLog(event, "NewTreasuryAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaiComtrollerNewTreasuryGuardianIterator is returned from FilterNewTreasuryGuardian and is used to iterate over the raw logs and unpacked data for NewTreasuryGuardian events raised by the VaiComtroller contract.
type VaiComtrollerNewTreasuryGuardianIterator struct {
	Event *VaiComtrollerNewTreasuryGuardian // Event containing the contract specifics and raw log

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
func (it *VaiComtrollerNewTreasuryGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaiComtrollerNewTreasuryGuardian)
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
		it.Event = new(VaiComtrollerNewTreasuryGuardian)
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
func (it *VaiComtrollerNewTreasuryGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaiComtrollerNewTreasuryGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaiComtrollerNewTreasuryGuardian represents a NewTreasuryGuardian event raised by the VaiComtroller contract.
type VaiComtrollerNewTreasuryGuardian struct {
	OldTreasuryGuardian common.Address
	NewTreasuryGuardian common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewTreasuryGuardian is a free log retrieval operation binding the contract event 0x29f06ea15931797ebaed313d81d100963dc22cb213cb4ce2737b5a62b1a8b1e8.
//
// Solidity: event NewTreasuryGuardian(address oldTreasuryGuardian, address newTreasuryGuardian)
func (_VaiComtroller *VaiComtrollerFilterer) FilterNewTreasuryGuardian(opts *bind.FilterOpts) (*VaiComtrollerNewTreasuryGuardianIterator, error) {

	logs, sub, err := _VaiComtroller.contract.FilterLogs(opts, "NewTreasuryGuardian")
	if err != nil {
		return nil, err
	}
	return &VaiComtrollerNewTreasuryGuardianIterator{contract: _VaiComtroller.contract, event: "NewTreasuryGuardian", logs: logs, sub: sub}, nil
}

// WatchNewTreasuryGuardian is a free log subscription operation binding the contract event 0x29f06ea15931797ebaed313d81d100963dc22cb213cb4ce2737b5a62b1a8b1e8.
//
// Solidity: event NewTreasuryGuardian(address oldTreasuryGuardian, address newTreasuryGuardian)
func (_VaiComtroller *VaiComtrollerFilterer) WatchNewTreasuryGuardian(opts *bind.WatchOpts, sink chan<- *VaiComtrollerNewTreasuryGuardian) (event.Subscription, error) {

	logs, sub, err := _VaiComtroller.contract.WatchLogs(opts, "NewTreasuryGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaiComtrollerNewTreasuryGuardian)
				if err := _VaiComtroller.contract.UnpackLog(event, "NewTreasuryGuardian", log); err != nil {
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

// ParseNewTreasuryGuardian is a log parse operation binding the contract event 0x29f06ea15931797ebaed313d81d100963dc22cb213cb4ce2737b5a62b1a8b1e8.
//
// Solidity: event NewTreasuryGuardian(address oldTreasuryGuardian, address newTreasuryGuardian)
func (_VaiComtroller *VaiComtrollerFilterer) ParseNewTreasuryGuardian(log types.Log) (*VaiComtrollerNewTreasuryGuardian, error) {
	event := new(VaiComtrollerNewTreasuryGuardian)
	if err := _VaiComtroller.contract.UnpackLog(event, "NewTreasuryGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaiComtrollerNewTreasuryPercentIterator is returned from FilterNewTreasuryPercent and is used to iterate over the raw logs and unpacked data for NewTreasuryPercent events raised by the VaiComtroller contract.
type VaiComtrollerNewTreasuryPercentIterator struct {
	Event *VaiComtrollerNewTreasuryPercent // Event containing the contract specifics and raw log

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
func (it *VaiComtrollerNewTreasuryPercentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaiComtrollerNewTreasuryPercent)
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
		it.Event = new(VaiComtrollerNewTreasuryPercent)
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
func (it *VaiComtrollerNewTreasuryPercentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaiComtrollerNewTreasuryPercentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaiComtrollerNewTreasuryPercent represents a NewTreasuryPercent event raised by the VaiComtroller contract.
type VaiComtrollerNewTreasuryPercent struct {
	OldTreasuryPercent *big.Int
	NewTreasuryPercent *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewTreasuryPercent is a free log retrieval operation binding the contract event 0x0893f8f4101baaabbeb513f96761e7a36eb837403c82cc651c292a4abdc94ed7.
//
// Solidity: event NewTreasuryPercent(uint256 oldTreasuryPercent, uint256 newTreasuryPercent)
func (_VaiComtroller *VaiComtrollerFilterer) FilterNewTreasuryPercent(opts *bind.FilterOpts) (*VaiComtrollerNewTreasuryPercentIterator, error) {

	logs, sub, err := _VaiComtroller.contract.FilterLogs(opts, "NewTreasuryPercent")
	if err != nil {
		return nil, err
	}
	return &VaiComtrollerNewTreasuryPercentIterator{contract: _VaiComtroller.contract, event: "NewTreasuryPercent", logs: logs, sub: sub}, nil
}

// WatchNewTreasuryPercent is a free log subscription operation binding the contract event 0x0893f8f4101baaabbeb513f96761e7a36eb837403c82cc651c292a4abdc94ed7.
//
// Solidity: event NewTreasuryPercent(uint256 oldTreasuryPercent, uint256 newTreasuryPercent)
func (_VaiComtroller *VaiComtrollerFilterer) WatchNewTreasuryPercent(opts *bind.WatchOpts, sink chan<- *VaiComtrollerNewTreasuryPercent) (event.Subscription, error) {

	logs, sub, err := _VaiComtroller.contract.WatchLogs(opts, "NewTreasuryPercent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaiComtrollerNewTreasuryPercent)
				if err := _VaiComtroller.contract.UnpackLog(event, "NewTreasuryPercent", log); err != nil {
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

// ParseNewTreasuryPercent is a log parse operation binding the contract event 0x0893f8f4101baaabbeb513f96761e7a36eb837403c82cc651c292a4abdc94ed7.
//
// Solidity: event NewTreasuryPercent(uint256 oldTreasuryPercent, uint256 newTreasuryPercent)
func (_VaiComtroller *VaiComtrollerFilterer) ParseNewTreasuryPercent(log types.Log) (*VaiComtrollerNewTreasuryPercent, error) {
	event := new(VaiComtrollerNewTreasuryPercent)
	if err := _VaiComtroller.contract.UnpackLog(event, "NewTreasuryPercent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaiComtrollerRepayVAIIterator is returned from FilterRepayVAI and is used to iterate over the raw logs and unpacked data for RepayVAI events raised by the VaiComtroller contract.
type VaiComtrollerRepayVAIIterator struct {
	Event *VaiComtrollerRepayVAI // Event containing the contract specifics and raw log

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
func (it *VaiComtrollerRepayVAIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaiComtrollerRepayVAI)
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
		it.Event = new(VaiComtrollerRepayVAI)
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
func (it *VaiComtrollerRepayVAIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaiComtrollerRepayVAIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaiComtrollerRepayVAI represents a RepayVAI event raised by the VaiComtroller contract.
type VaiComtrollerRepayVAI struct {
	Payer          common.Address
	Borrower       common.Address
	RepayVAIAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRepayVAI is a free log retrieval operation binding the contract event 0x1db858e6f7e1a0d5e92c10c6507d42b3dabfe0a4867fe90c5a14d9963662ef7e.
//
// Solidity: event RepayVAI(address payer, address borrower, uint256 repayVAIAmount)
func (_VaiComtroller *VaiComtrollerFilterer) FilterRepayVAI(opts *bind.FilterOpts) (*VaiComtrollerRepayVAIIterator, error) {

	logs, sub, err := _VaiComtroller.contract.FilterLogs(opts, "RepayVAI")
	if err != nil {
		return nil, err
	}
	return &VaiComtrollerRepayVAIIterator{contract: _VaiComtroller.contract, event: "RepayVAI", logs: logs, sub: sub}, nil
}

// WatchRepayVAI is a free log subscription operation binding the contract event 0x1db858e6f7e1a0d5e92c10c6507d42b3dabfe0a4867fe90c5a14d9963662ef7e.
//
// Solidity: event RepayVAI(address payer, address borrower, uint256 repayVAIAmount)
func (_VaiComtroller *VaiComtrollerFilterer) WatchRepayVAI(opts *bind.WatchOpts, sink chan<- *VaiComtrollerRepayVAI) (event.Subscription, error) {

	logs, sub, err := _VaiComtroller.contract.WatchLogs(opts, "RepayVAI")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaiComtrollerRepayVAI)
				if err := _VaiComtroller.contract.UnpackLog(event, "RepayVAI", log); err != nil {
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

// ParseRepayVAI is a log parse operation binding the contract event 0x1db858e6f7e1a0d5e92c10c6507d42b3dabfe0a4867fe90c5a14d9963662ef7e.
//
// Solidity: event RepayVAI(address payer, address borrower, uint256 repayVAIAmount)
func (_VaiComtroller *VaiComtrollerFilterer) ParseRepayVAI(log types.Log) (*VaiComtrollerRepayVAI, error) {
	event := new(VaiComtrollerRepayVAI)
	if err := _VaiComtroller.contract.UnpackLog(event, "RepayVAI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

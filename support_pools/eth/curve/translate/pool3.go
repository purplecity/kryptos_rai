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

// Pool3ABI is the input ABI used to generate the binding from.
const Pool3ABI = "[{\"name\":\"balances\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Pool3 is an auto generated Go binding around an Ethereum contract.
type Pool3 struct {
	Pool3Caller     // Read-only binding to the contract
	Pool3Transactor // Write-only binding to the contract
	Pool3Filterer   // Log filterer for contract events
}

// Pool3Caller is an auto generated read-only Go binding around an Ethereum contract.
type Pool3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Pool3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Pool3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Pool3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Pool3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Pool3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Pool3Session struct {
	Contract     *Pool3            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Pool3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Pool3CallerSession struct {
	Contract *Pool3Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Pool3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Pool3TransactorSession struct {
	Contract     *Pool3Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Pool3Raw is an auto generated low-level Go binding around an Ethereum contract.
type Pool3Raw struct {
	Contract *Pool3 // Generic contract binding to access the raw methods on
}

// Pool3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Pool3CallerRaw struct {
	Contract *Pool3Caller // Generic read-only contract binding to access the raw methods on
}

// Pool3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Pool3TransactorRaw struct {
	Contract *Pool3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPool3 creates a new instance of Pool3, bound to a specific deployed contract.
func NewPool3(address common.Address, backend bind.ContractBackend) (*Pool3, error) {
	contract, err := bindPool3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pool3{Pool3Caller: Pool3Caller{contract: contract}, Pool3Transactor: Pool3Transactor{contract: contract}, Pool3Filterer: Pool3Filterer{contract: contract}}, nil
}

// NewPool3Caller creates a new read-only instance of Pool3, bound to a specific deployed contract.
func NewPool3Caller(address common.Address, caller bind.ContractCaller) (*Pool3Caller, error) {
	contract, err := bindPool3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Pool3Caller{contract: contract}, nil
}

// NewPool3Transactor creates a new write-only instance of Pool3, bound to a specific deployed contract.
func NewPool3Transactor(address common.Address, transactor bind.ContractTransactor) (*Pool3Transactor, error) {
	contract, err := bindPool3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Pool3Transactor{contract: contract}, nil
}

// NewPool3Filterer creates a new log filterer instance of Pool3, bound to a specific deployed contract.
func NewPool3Filterer(address common.Address, filterer bind.ContractFilterer) (*Pool3Filterer, error) {
	contract, err := bindPool3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Pool3Filterer{contract: contract}, nil
}

// bindPool3 binds a generic wrapper to an already deployed contract.
func bindPool3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Pool3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool3 *Pool3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool3.Contract.Pool3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool3 *Pool3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool3.Contract.Pool3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool3 *Pool3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool3.Contract.Pool3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool3 *Pool3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool3 *Pool3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool3 *Pool3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool3.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_Pool3 *Pool3Caller) Balances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool3.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_Pool3 *Pool3Session) Balances(arg0 *big.Int) (*big.Int, error) {
	return _Pool3.Contract.Balances(&_Pool3.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_Pool3 *Pool3CallerSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _Pool3.Contract.Balances(&_Pool3.CallOpts, arg0)
}

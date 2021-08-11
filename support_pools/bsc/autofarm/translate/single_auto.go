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

// SingleAutoABI is the input ABI used to generate the binding from.
const SingleAutoABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"_want\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_strat\",\"type\":\"address\"}],\"name\":\"Add\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_want\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_strat\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_wantAmt\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"inCaseTokensGetStuck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"want\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accAUTOPerShare\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"strat\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"stakedWantTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_wantAmt\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SingleAutoBin is the compiled bytecode used for deploying new contracts.
var SingleAutoBin = "0x6080604052600160045534801561001557600080fd5b506000610020610073565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35060018055610077565b3390565b6115d8806100866000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80637b84daec116100715780637b84daec146101a75780638da5cb5b146101d357806393f1a40b146101f7578063c6d758cb14610223578063e2bbb1581461024f578063f2fde38b14610272576100b4565b8063081e3eda146100b95780631526fe27146100d357806317caf6f114610134578063441a3e701461013c5780635f70edcf14610161578063715018a61461019f575b600080fd5b6100c1610298565b60408051918252519081900360200190f35b6100f0600480360360208110156100e957600080fd5b503561029e565b60405180866001600160a01b03168152602001858152602001848152602001838152602001826001600160a01b031681526020019550505050505060405180910390f35b6100c16102e9565b61015f6004803603604081101561015257600080fd5b50803590602001356102ef565b005b61015f6004803603608081101561017757600080fd5b508035906001600160a01b036020820135811691604081013515159160609091013516610721565b61015f610903565b6100c1600480360360408110156101bd57600080fd5b50803590602001356001600160a01b03166109a5565b6101db610b22565b604080516001600160a01b039092168252519081900360200190f35b6100c16004803603604081101561020d57600080fd5b50803590602001356001600160a01b0316610b31565b61015f6004803603604081101561023957600080fd5b506001600160a01b038135169060200135610b4e565b61015f6004803603604081101561026557600080fd5b5080359060200135610bbe565b61015f6004803603602081101561028857600080fd5b50356001600160a01b0316610d83565b60025490565b600281815481106102ab57fe5b6000918252602090912060059091020180546001820154600283015460038401546004909401546001600160a01b0393841695509193909290911685565b60045481565b60026001541415610347576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b600260018190555060006002838154811061035e57fe5b60009182526020808320868452600382526040808520338652909252908320600280546005909402909201945092918690811061039757fe5b906000526020600020906005020160040160009054906101000a90046001600160a01b03166001600160a01b03166342da4eb36040518163ffffffff1660e01b815260040160206040518083038186803b1580156103f457600080fd5b505afa158015610408573d6000803e3d6000fd5b505050506040513d602081101561041e57600080fd5b5051600280549192506000918790811061043457fe5b906000526020600020906005020160040160009054906101000a90046001600160a01b03166001600160a01b03166344a3955e6040518163ffffffff1660e01b815260040160206040518083038186803b15801561049157600080fd5b505afa1580156104a5573d6000803e3d6000fd5b505050506040513d60208110156104bb57600080fd5b50518354909150610506576040805162461bcd60e51b815260206004820152601060248201526f0757365722e73686172657320697320360841b604482015290519081900360640190fd5b6000811161054e576040805162461bcd60e51b815260206004820152601060248201526f0736861726573546f74616c20697320360841b604482015290519081900360640190fd5b82546000906105699083906105639086610e7b565b90610edb565b905080861115610577578095505b85156106dd5760006002888154811061058c57fe5b6000918252602080832060046005909302018201546040805163f3fef3a360e01b81523394810194909452602484018c9052516001600160a01b039091169363f3fef3a3936044808201949392918390030190829087803b1580156105f057600080fd5b505af1158015610604573d6000803e3d6000fd5b505050506040513d602081101561061a57600080fd5b505185549091508111156106315760008555610640565b845461063d9082610f1d565b85555b8554604080516370a0823160e01b815230600482015290516000926001600160a01b0316916370a08231916024808301926020929190829003018186803b15801561068a57600080fd5b505afa15801561069e573d6000803e3d6000fd5b505050506040513d60208110156106b457600080fd5b50519050878110156106c4578097505b86546106da906001600160a01b0316338a610f5f565b50505b604080518781529051889133917ff279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b5689181900360200190a35050600180555050505050565b610729610fb6565b6000546001600160a01b03908116911614610779576040805162461bcd60e51b81526020600482018190526024820152600080516020611559833981519152604482015290519081900360640190fd5b6004546107869085610fba565b6004556040805160a0810182526001600160a01b038086168083526020808401898152600085870181815260608088018381528a88166080808b018281526002805460018101825597529a5160059096027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace81018054978c166001600160a01b031998891617905596517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5acf88015593517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ad087015590517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ad186015597517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ad2909401805494909716939092169290921790945585518a815291820192909252861515818601529182019290925291517f7ac3b3d636ea03a286c25edba9e75801c02c65437a5fc04d2160a9693f7d31b39281900390910190a150505050565b61090b610fb6565b6000546001600160a01b0390811691161461095b576040805162461bcd60e51b81526020600482018190526024820152600080516020611559833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b600080600284815481106109b557fe5b600091825260208083208784526003825260408085206001600160a01b03808a16875290845281862060046005909602909301858101548351632251caaf60e11b8152935191985093969593909116936344a3955e9383810193919291829003018186803b158015610a2657600080fd5b505afa158015610a3a573d6000803e3d6000fd5b505050506040513d6020811015610a5057600080fd5b50516002805491925060009188908110610a6657fe5b906000526020600020906005020160040160009054906101000a90046001600160a01b03166001600160a01b03166342da4eb36040518163ffffffff1660e01b815260040160206040518083038186803b158015610ac357600080fd5b505afa158015610ad7573d6000803e3d6000fd5b505050506040513d6020811015610aed57600080fd5b5051905081610b03576000945050505050610b1c565b8254610b159083906105639084610e7b565b9450505050505b92915050565b6000546001600160a01b031690565b600360209081526000928352604080842090915290825290205481565b610b56610fb6565b6000546001600160a01b03908116911614610ba6576040805162461bcd60e51b81526020600482018190526024820152600080516020611559833981519152604482015290519081900360640190fd5b610bba6001600160a01b0383163383610f5f565b5050565b60026001541415610c16576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6002600181905550600060028381548110610c2d57fe5b600091825260208083208684526003825260408085203386529092529220600590910290910191508215610d42578154610c72906001600160a01b0316333086611014565b60048201548254610c90916001600160a01b03918216911685611074565b600060028581548110610c9f57fe5b600091825260208083206004600590930201820154604080516311f9fbc960e21b8152339481019490945260248401899052516001600160a01b03909116936347e7ef24936044808201949392918390030190829087803b158015610d0357600080fd5b505af1158015610d17573d6000803e3d6000fd5b505050506040513d6020811015610d2d57600080fd5b50518254909150610d3e9082610fba565b8255505b604080518481529051859133917f90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a159181900360200190a35050600180555050565b610d8b610fb6565b6000546001600160a01b03908116911614610ddb576040805162461bcd60e51b81526020600482018190526024820152600080516020611559833981519152604482015290519081900360640190fd5b6001600160a01b038116610e205760405162461bcd60e51b81526004018080602001828103825260268152602001806114ec6026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b600082610e8a57506000610b1c565b82820282848281610e9757fe5b0414610ed45760405162461bcd60e51b81526004018080602001828103825260218152602001806115386021913960400191505060405180910390fd5b9392505050565b6000610ed483836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f00000000000081525061115f565b6000610ed483836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611201565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052610fb190849061125b565b505050565b3390565b600082820183811015610ed4576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b604080516001600160a01b0380861660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b17905261106e90859061125b565b50505050565b600061110a82856001600160a01b031663dd62ed3e30876040518363ffffffff1660e01b815260040180836001600160a01b03168152602001826001600160a01b031681526020019250505060206040518083038186803b1580156110d857600080fd5b505afa1580156110ec573d6000803e3d6000fd5b505050506040513d602081101561110257600080fd5b505190610fba565b604080516001600160a01b038616602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b17905290915061106e90859061125b565b600081836111eb5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156111b0578181015183820152602001611198565b50505050905090810190601f1680156111dd5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385816111f757fe5b0495945050505050565b600081848411156112535760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156111b0578181015183820152602001611198565b505050900390565b60606112b0826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661130c9092919063ffffffff16565b805190915015610fb1578080602001905160208110156112cf57600080fd5b5051610fb15760405162461bcd60e51b815260040180806020018281038252602a815260200180611579602a913960400191505060405180910390fd5b606061131b8484600085611323565b949350505050565b6060824710156113645760405162461bcd60e51b81526004018080602001828103825260268152602001806115126026913960400191505060405180910390fd5b61136d8561147f565b6113be576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b60006060866001600160a01b031685876040518082805190602001908083835b602083106113fd5780518252601f1990920191602091820191016113de565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d806000811461145f576040519150601f19603f3d011682016040523d82523d6000602084013e611464565b606091505b5091509150611474828286611485565b979650505050505050565b3b151590565b60608315611494575081610ed4565b8251156114a45782518084602001fd5b60405162461bcd60e51b81526020600482018181528451602484015284518593919283926044019190850190808383600083156111b057818101518382015260200161119856fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f774f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a2646970667358221220f450f9cad1dafc309e491bf63e2b1109c795b4d73ab3e5fc3eed13ea718959e764736f6c634300060c0033"

// DeploySingleAuto deploys a new Ethereum contract, binding an instance of SingleAuto to it.
func DeploySingleAuto(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SingleAuto, error) {
	parsed, err := abi.JSON(strings.NewReader(SingleAutoABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SingleAutoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SingleAuto{SingleAutoCaller: SingleAutoCaller{contract: contract}, SingleAutoTransactor: SingleAutoTransactor{contract: contract}, SingleAutoFilterer: SingleAutoFilterer{contract: contract}}, nil
}

// SingleAuto is an auto generated Go binding around an Ethereum contract.
type SingleAuto struct {
	SingleAutoCaller     // Read-only binding to the contract
	SingleAutoTransactor // Write-only binding to the contract
	SingleAutoFilterer   // Log filterer for contract events
}

// SingleAutoCaller is an auto generated read-only Go binding around an Ethereum contract.
type SingleAutoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SingleAutoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SingleAutoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SingleAutoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SingleAutoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SingleAutoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SingleAutoSession struct {
	Contract     *SingleAuto       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SingleAutoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SingleAutoCallerSession struct {
	Contract *SingleAutoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SingleAutoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SingleAutoTransactorSession struct {
	Contract     *SingleAutoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SingleAutoRaw is an auto generated low-level Go binding around an Ethereum contract.
type SingleAutoRaw struct {
	Contract *SingleAuto // Generic contract binding to access the raw methods on
}

// SingleAutoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SingleAutoCallerRaw struct {
	Contract *SingleAutoCaller // Generic read-only contract binding to access the raw methods on
}

// SingleAutoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SingleAutoTransactorRaw struct {
	Contract *SingleAutoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSingleAuto creates a new instance of SingleAuto, bound to a specific deployed contract.
func NewSingleAuto(address common.Address, backend bind.ContractBackend) (*SingleAuto, error) {
	contract, err := bindSingleAuto(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SingleAuto{SingleAutoCaller: SingleAutoCaller{contract: contract}, SingleAutoTransactor: SingleAutoTransactor{contract: contract}, SingleAutoFilterer: SingleAutoFilterer{contract: contract}}, nil
}

// NewSingleAutoCaller creates a new read-only instance of SingleAuto, bound to a specific deployed contract.
func NewSingleAutoCaller(address common.Address, caller bind.ContractCaller) (*SingleAutoCaller, error) {
	contract, err := bindSingleAuto(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SingleAutoCaller{contract: contract}, nil
}

// NewSingleAutoTransactor creates a new write-only instance of SingleAuto, bound to a specific deployed contract.
func NewSingleAutoTransactor(address common.Address, transactor bind.ContractTransactor) (*SingleAutoTransactor, error) {
	contract, err := bindSingleAuto(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SingleAutoTransactor{contract: contract}, nil
}

// NewSingleAutoFilterer creates a new log filterer instance of SingleAuto, bound to a specific deployed contract.
func NewSingleAutoFilterer(address common.Address, filterer bind.ContractFilterer) (*SingleAutoFilterer, error) {
	contract, err := bindSingleAuto(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SingleAutoFilterer{contract: contract}, nil
}

// bindSingleAuto binds a generic wrapper to an already deployed contract.
func bindSingleAuto(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SingleAutoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SingleAuto *SingleAutoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SingleAuto.Contract.SingleAutoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SingleAuto *SingleAutoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SingleAuto.Contract.SingleAutoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SingleAuto *SingleAutoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SingleAuto.Contract.SingleAutoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SingleAuto *SingleAutoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SingleAuto.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SingleAuto *SingleAutoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SingleAuto.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SingleAuto *SingleAutoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SingleAuto.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SingleAuto *SingleAutoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SingleAuto.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SingleAuto *SingleAutoSession) Owner() (common.Address, error) {
	return _SingleAuto.Contract.Owner(&_SingleAuto.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SingleAuto *SingleAutoCallerSession) Owner() (common.Address, error) {
	return _SingleAuto.Contract.Owner(&_SingleAuto.CallOpts)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address want, uint256 allocPoint, uint256 lastRewardBlock, uint256 accAUTOPerShare, address strat)
func (_SingleAuto *SingleAutoCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Want            common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccAUTOPerShare *big.Int
	Strat           common.Address
}, error) {
	var out []interface{}
	err := _SingleAuto.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		Want            common.Address
		AllocPoint      *big.Int
		LastRewardBlock *big.Int
		AccAUTOPerShare *big.Int
		Strat           common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Want = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AllocPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastRewardBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccAUTOPerShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Strat = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address want, uint256 allocPoint, uint256 lastRewardBlock, uint256 accAUTOPerShare, address strat)
func (_SingleAuto *SingleAutoSession) PoolInfo(arg0 *big.Int) (struct {
	Want            common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccAUTOPerShare *big.Int
	Strat           common.Address
}, error) {
	return _SingleAuto.Contract.PoolInfo(&_SingleAuto.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address want, uint256 allocPoint, uint256 lastRewardBlock, uint256 accAUTOPerShare, address strat)
func (_SingleAuto *SingleAutoCallerSession) PoolInfo(arg0 *big.Int) (struct {
	Want            common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccAUTOPerShare *big.Int
	Strat           common.Address
}, error) {
	return _SingleAuto.Contract.PoolInfo(&_SingleAuto.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_SingleAuto *SingleAutoCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SingleAuto.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_SingleAuto *SingleAutoSession) PoolLength() (*big.Int, error) {
	return _SingleAuto.Contract.PoolLength(&_SingleAuto.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_SingleAuto *SingleAutoCallerSession) PoolLength() (*big.Int, error) {
	return _SingleAuto.Contract.PoolLength(&_SingleAuto.CallOpts)
}

// StakedWantTokens is a free data retrieval call binding the contract method 0x7b84daec.
//
// Solidity: function stakedWantTokens(uint256 _pid, address _user) view returns(uint256)
func (_SingleAuto *SingleAutoCaller) StakedWantTokens(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SingleAuto.contract.Call(opts, &out, "stakedWantTokens", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedWantTokens is a free data retrieval call binding the contract method 0x7b84daec.
//
// Solidity: function stakedWantTokens(uint256 _pid, address _user) view returns(uint256)
func (_SingleAuto *SingleAutoSession) StakedWantTokens(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _SingleAuto.Contract.StakedWantTokens(&_SingleAuto.CallOpts, _pid, _user)
}

// StakedWantTokens is a free data retrieval call binding the contract method 0x7b84daec.
//
// Solidity: function stakedWantTokens(uint256 _pid, address _user) view returns(uint256)
func (_SingleAuto *SingleAutoCallerSession) StakedWantTokens(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _SingleAuto.Contract.StakedWantTokens(&_SingleAuto.CallOpts, _pid, _user)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_SingleAuto *SingleAutoCaller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SingleAuto.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_SingleAuto *SingleAutoSession) TotalAllocPoint() (*big.Int, error) {
	return _SingleAuto.Contract.TotalAllocPoint(&_SingleAuto.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_SingleAuto *SingleAutoCallerSession) TotalAllocPoint() (*big.Int, error) {
	return _SingleAuto.Contract.TotalAllocPoint(&_SingleAuto.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 shares)
func (_SingleAuto *SingleAutoCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SingleAuto.contract.Call(opts, &out, "userInfo", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 shares)
func (_SingleAuto *SingleAutoSession) UserInfo(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _SingleAuto.Contract.UserInfo(&_SingleAuto.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 shares)
func (_SingleAuto *SingleAutoCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _SingleAuto.Contract.UserInfo(&_SingleAuto.CallOpts, arg0, arg1)
}

// Add is a paid mutator transaction binding the contract method 0x5f70edcf.
//
// Solidity: function add(uint256 _allocPoint, address _want, bool _withUpdate, address _strat) returns()
func (_SingleAuto *SingleAutoTransactor) Add(opts *bind.TransactOpts, _allocPoint *big.Int, _want common.Address, _withUpdate bool, _strat common.Address) (*types.Transaction, error) {
	return _SingleAuto.contract.Transact(opts, "add", _allocPoint, _want, _withUpdate, _strat)
}

// Add is a paid mutator transaction binding the contract method 0x5f70edcf.
//
// Solidity: function add(uint256 _allocPoint, address _want, bool _withUpdate, address _strat) returns()
func (_SingleAuto *SingleAutoSession) Add(_allocPoint *big.Int, _want common.Address, _withUpdate bool, _strat common.Address) (*types.Transaction, error) {
	return _SingleAuto.Contract.Add(&_SingleAuto.TransactOpts, _allocPoint, _want, _withUpdate, _strat)
}

// Add is a paid mutator transaction binding the contract method 0x5f70edcf.
//
// Solidity: function add(uint256 _allocPoint, address _want, bool _withUpdate, address _strat) returns()
func (_SingleAuto *SingleAutoTransactorSession) Add(_allocPoint *big.Int, _want common.Address, _withUpdate bool, _strat common.Address) (*types.Transaction, error) {
	return _SingleAuto.Contract.Add(&_SingleAuto.TransactOpts, _allocPoint, _want, _withUpdate, _strat)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _wantAmt) returns()
func (_SingleAuto *SingleAutoTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _wantAmt *big.Int) (*types.Transaction, error) {
	return _SingleAuto.contract.Transact(opts, "deposit", _pid, _wantAmt)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _wantAmt) returns()
func (_SingleAuto *SingleAutoSession) Deposit(_pid *big.Int, _wantAmt *big.Int) (*types.Transaction, error) {
	return _SingleAuto.Contract.Deposit(&_SingleAuto.TransactOpts, _pid, _wantAmt)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _wantAmt) returns()
func (_SingleAuto *SingleAutoTransactorSession) Deposit(_pid *big.Int, _wantAmt *big.Int) (*types.Transaction, error) {
	return _SingleAuto.Contract.Deposit(&_SingleAuto.TransactOpts, _pid, _wantAmt)
}

// InCaseTokensGetStuck is a paid mutator transaction binding the contract method 0xc6d758cb.
//
// Solidity: function inCaseTokensGetStuck(address _token, uint256 _amount) returns()
func (_SingleAuto *SingleAutoTransactor) InCaseTokensGetStuck(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SingleAuto.contract.Transact(opts, "inCaseTokensGetStuck", _token, _amount)
}

// InCaseTokensGetStuck is a paid mutator transaction binding the contract method 0xc6d758cb.
//
// Solidity: function inCaseTokensGetStuck(address _token, uint256 _amount) returns()
func (_SingleAuto *SingleAutoSession) InCaseTokensGetStuck(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SingleAuto.Contract.InCaseTokensGetStuck(&_SingleAuto.TransactOpts, _token, _amount)
}

// InCaseTokensGetStuck is a paid mutator transaction binding the contract method 0xc6d758cb.
//
// Solidity: function inCaseTokensGetStuck(address _token, uint256 _amount) returns()
func (_SingleAuto *SingleAutoTransactorSession) InCaseTokensGetStuck(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SingleAuto.Contract.InCaseTokensGetStuck(&_SingleAuto.TransactOpts, _token, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SingleAuto *SingleAutoTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SingleAuto.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SingleAuto *SingleAutoSession) RenounceOwnership() (*types.Transaction, error) {
	return _SingleAuto.Contract.RenounceOwnership(&_SingleAuto.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SingleAuto *SingleAutoTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SingleAuto.Contract.RenounceOwnership(&_SingleAuto.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SingleAuto *SingleAutoTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SingleAuto.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SingleAuto *SingleAutoSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SingleAuto.Contract.TransferOwnership(&_SingleAuto.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SingleAuto *SingleAutoTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SingleAuto.Contract.TransferOwnership(&_SingleAuto.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _wantAmt) returns()
func (_SingleAuto *SingleAutoTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _wantAmt *big.Int) (*types.Transaction, error) {
	return _SingleAuto.contract.Transact(opts, "withdraw", _pid, _wantAmt)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _wantAmt) returns()
func (_SingleAuto *SingleAutoSession) Withdraw(_pid *big.Int, _wantAmt *big.Int) (*types.Transaction, error) {
	return _SingleAuto.Contract.Withdraw(&_SingleAuto.TransactOpts, _pid, _wantAmt)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _wantAmt) returns()
func (_SingleAuto *SingleAutoTransactorSession) Withdraw(_pid *big.Int, _wantAmt *big.Int) (*types.Transaction, error) {
	return _SingleAuto.Contract.Withdraw(&_SingleAuto.TransactOpts, _pid, _wantAmt)
}

// SingleAutoAddIterator is returned from FilterAdd and is used to iterate over the raw logs and unpacked data for Add events raised by the SingleAuto contract.
type SingleAutoAddIterator struct {
	Event *SingleAutoAdd // Event containing the contract specifics and raw log

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
func (it *SingleAutoAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleAutoAdd)
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
		it.Event = new(SingleAutoAdd)
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
func (it *SingleAutoAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SingleAutoAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SingleAutoAdd represents a Add event raised by the SingleAuto contract.
type SingleAutoAdd struct {
	AllocPoint *big.Int
	Want       common.Address
	WithUpdate bool
	Strat      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAdd is a free log retrieval operation binding the contract event 0x7ac3b3d636ea03a286c25edba9e75801c02c65437a5fc04d2160a9693f7d31b3.
//
// Solidity: event Add(uint256 _allocPoint, address _want, bool _withUpdate, address _strat)
func (_SingleAuto *SingleAutoFilterer) FilterAdd(opts *bind.FilterOpts) (*SingleAutoAddIterator, error) {

	logs, sub, err := _SingleAuto.contract.FilterLogs(opts, "Add")
	if err != nil {
		return nil, err
	}
	return &SingleAutoAddIterator{contract: _SingleAuto.contract, event: "Add", logs: logs, sub: sub}, nil
}

// WatchAdd is a free log subscription operation binding the contract event 0x7ac3b3d636ea03a286c25edba9e75801c02c65437a5fc04d2160a9693f7d31b3.
//
// Solidity: event Add(uint256 _allocPoint, address _want, bool _withUpdate, address _strat)
func (_SingleAuto *SingleAutoFilterer) WatchAdd(opts *bind.WatchOpts, sink chan<- *SingleAutoAdd) (event.Subscription, error) {

	logs, sub, err := _SingleAuto.contract.WatchLogs(opts, "Add")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SingleAutoAdd)
				if err := _SingleAuto.contract.UnpackLog(event, "Add", log); err != nil {
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

// ParseAdd is a log parse operation binding the contract event 0x7ac3b3d636ea03a286c25edba9e75801c02c65437a5fc04d2160a9693f7d31b3.
//
// Solidity: event Add(uint256 _allocPoint, address _want, bool _withUpdate, address _strat)
func (_SingleAuto *SingleAutoFilterer) ParseAdd(log types.Log) (*SingleAutoAdd, error) {
	event := new(SingleAutoAdd)
	if err := _SingleAuto.contract.UnpackLog(event, "Add", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SingleAutoDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the SingleAuto contract.
type SingleAutoDepositIterator struct {
	Event *SingleAutoDeposit // Event containing the contract specifics and raw log

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
func (it *SingleAutoDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleAutoDeposit)
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
		it.Event = new(SingleAutoDeposit)
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
func (it *SingleAutoDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SingleAutoDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SingleAutoDeposit represents a Deposit event raised by the SingleAuto contract.
type SingleAutoDeposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_SingleAuto *SingleAutoFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*SingleAutoDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _SingleAuto.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &SingleAutoDepositIterator{contract: _SingleAuto.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_SingleAuto *SingleAutoFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SingleAutoDeposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _SingleAuto.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SingleAutoDeposit)
				if err := _SingleAuto.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_SingleAuto *SingleAutoFilterer) ParseDeposit(log types.Log) (*SingleAutoDeposit, error) {
	event := new(SingleAutoDeposit)
	if err := _SingleAuto.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SingleAutoOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SingleAuto contract.
type SingleAutoOwnershipTransferredIterator struct {
	Event *SingleAutoOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SingleAutoOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleAutoOwnershipTransferred)
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
		it.Event = new(SingleAutoOwnershipTransferred)
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
func (it *SingleAutoOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SingleAutoOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SingleAutoOwnershipTransferred represents a OwnershipTransferred event raised by the SingleAuto contract.
type SingleAutoOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SingleAuto *SingleAutoFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SingleAutoOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SingleAuto.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SingleAutoOwnershipTransferredIterator{contract: _SingleAuto.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SingleAuto *SingleAutoFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SingleAutoOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SingleAuto.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SingleAutoOwnershipTransferred)
				if err := _SingleAuto.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SingleAuto *SingleAutoFilterer) ParseOwnershipTransferred(log types.Log) (*SingleAutoOwnershipTransferred, error) {
	event := new(SingleAutoOwnershipTransferred)
	if err := _SingleAuto.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SingleAutoWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the SingleAuto contract.
type SingleAutoWithdrawIterator struct {
	Event *SingleAutoWithdraw // Event containing the contract specifics and raw log

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
func (it *SingleAutoWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleAutoWithdraw)
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
		it.Event = new(SingleAutoWithdraw)
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
func (it *SingleAutoWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SingleAutoWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SingleAutoWithdraw represents a Withdraw event raised by the SingleAuto contract.
type SingleAutoWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_SingleAuto *SingleAutoFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*SingleAutoWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _SingleAuto.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &SingleAutoWithdrawIterator{contract: _SingleAuto.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_SingleAuto *SingleAutoFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SingleAutoWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _SingleAuto.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SingleAutoWithdraw)
				if err := _SingleAuto.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_SingleAuto *SingleAutoFilterer) ParseWithdraw(log types.Log) (*SingleAutoWithdraw, error) {
	event := new(SingleAutoWithdraw)
	if err := _SingleAuto.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

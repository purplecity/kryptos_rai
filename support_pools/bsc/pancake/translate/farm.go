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

// FarmABI is the input ABI used to generate the binding from.
const FarmABI = "[{\"inputs\":[{\"internalType\":\"contractCakeToken\",\"name\":\"_cake\",\"type\":\"address\"},{\"internalType\":\"contractSyrupBar\",\"name\":\"_syrup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_devaddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_cakePerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BONUS_MULTIPLIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIBEP20\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cake\",\"outputs\":[{\"internalType\":\"contractCakeToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cakePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_devaddr\",\"type\":\"address\"}],\"name\":\"dev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"devaddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"enterStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_to\",\"type\":\"uint256\"}],\"name\":\"getMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"leaveStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrator\",\"outputs\":[{\"internalType\":\"contractIMigratorChef\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingCake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accCakePerShare\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMigratorChef\",\"name\":\"_migrator\",\"type\":\"address\"}],\"name\":\"setMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"syrup\",\"outputs\":[{\"internalType\":\"contractSyrupBar\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"multiplierNumber\",\"type\":\"uint256\"}],\"name\":\"updateMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FarmBin is the compiled bytecode used for deploying new contracts.
var FarmBin = "0x60806040526001600555600060095534801561001a57600080fd5b506040516125b73803806125b7833981810160405260a081101561003d57600080fd5b5080516020820151604083015160608401516080909401519293919290919060006100666101e0565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600180546001600160a01b03199081166001600160a01b03978816908117835560028054831697891697909717909655600380548216958816959095179094556004928355600a829055604080516080810182529586526103e86020870181815291870193845260006060880181815260078054958601815590915296517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688939095029283018054909616949097169390931790935590517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c689830155517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68a82015590517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68b909101556009556101e4565b3390565b6123c4806101f36000396000f3fe608060405234801561001057600080fd5b50600436106101cf5760003560e01c80635ffe6146116101045780638d88a90e116100a2578063d49e77cd11610071578063d49e77cd146104b1578063dce17484146104b9578063e2bbb158146104c1578063f2fde38b146104e4576101cf565b80638d88a90e1461041b5780638da5cb5b146104415780638dbb1e3a1461044957806393f1a40b1461046c576101cf565b8063715018a6116100de578063715018a6146103df5780637cd07e47146103e757806386a952c41461040b5780638aa2855014610413576101cf565b80635ffe61461461038f578063630b5ba1146103ac57806364482f79146103b4576101cf565b806323cf311811610171578063454b06081161014b578063454b06081461033057806348cd4cb11461034d57806351eb05a6146103555780635312ea8e14610372576101cf565b806323cf3118146102ca57806341441d3b146102f0578063441a3e701461030d576101cf565b80631175a1dd116101ad5780631175a1dd146102155780631526fe271461024157806317caf6f11461028e5780631eaaa04514610296576101cf565b80630755e0b6146101d4578063081e3eda146101ee5780631058d281146101f6575b600080fd5b6101dc61050a565b60408051918252519081900360200190f35b6101dc610510565b6102136004803603602081101561020c57600080fd5b5035610516565b005b6101dc6004803603604081101561022b57600080fd5b50803590602001356001600160a01b031661071e565b61025e6004803603602081101561025757600080fd5b50356108a1565b604080516001600160a01b0390951685526020850193909352838301919091526060830152519081900360800190f35b6101dc6108e2565b610213600480360360608110156102ac57600080fd5b508035906001600160a01b03602082013516906040013515156108e8565b610213600480360360208110156102e057600080fd5b50356001600160a01b0316610a98565b6102136004803603602081101561030657600080fd5b5035610b3c565b6102136004803603604081101561032357600080fd5b5080359060200135610ce9565b6102136004803603602081101561034657600080fd5b5035610e99565b6101dc611172565b6102136004803603602081101561036b57600080fd5b5035611178565b6102136004803603602081101561038857600080fd5b50356113d0565b610213600480360360208110156103a557600080fd5b503561146b565b6102136114da565b610213600480360360608110156103ca57600080fd5b508035906020810135906040013515156114fd565b61021361160d565b6103ef6116d9565b604080516001600160a01b039092168252519081900360200190f35b6103ef6116e8565b6101dc6116f7565b6102136004803603602081101561043157600080fd5b50356001600160a01b03166116fd565b6103ef611796565b6101dc6004803603604081101561045f57600080fd5b50803590602001356117a5565b6104986004803603604081101561048257600080fd5b50803590602001356001600160a01b03166117c0565b6040805192835260208301919091528051918290030190f35b6103ef6117e4565b6103ef6117f3565b610213600480360360408110156104d757600080fd5b5080359060200135611802565b610213600480360360208110156104fa57600080fd5b50356001600160a01b0316611966565b60045481565b60075490565b6000600760008154811061052657fe5b600091825260208083203384527f5eff886ea0ce6ca488a3d6e336d6c0f75f46d19b42c06ce5ee98e42c96d256c790915260409092208054600490920290920192508311156105bc576040805162461bcd60e51b815260206004820152601260248201527f77697468647261773a206e6f7420676f6f640000000000000000000000000000604482015290519081900360640190fd5b6105c66000611178565b600061060082600101546105fa64e8d4a510006105f4876003015487600001546119d990919063ffffffff16565b90611a32565b90611a74565b90508015610612576106123382611ab6565b831561063c5781546106249085611a74565b8255825461063c906001600160a01b03163386611b40565b600383015482546106579164e8d4a51000916105f4916119d9565b6001830155600254604080517f9dc29fac0000000000000000000000000000000000000000000000000000000081523360048201526024810187905290516001600160a01b0390921691639dc29fac9160448082019260009290919082900301818387803b1580156106c857600080fd5b505af11580156106dc573d6000803e3d6000fd5b5050604080518781529051600093503392507ff279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b5689181900360200190a350505050565b6000806007848154811061072e57fe5b600091825260208083208784526008825260408085206001600160a01b03898116875290845281862060049586029093016003810154815484517f70a082310000000000000000000000000000000000000000000000000000000081523098810198909852935191985093969395939492909116926370a08231926024808301939192829003018186803b1580156107c557600080fd5b505afa1580156107d9573d6000803e3d6000fd5b505050506040513d60208110156107ef57600080fd5b505160028501549091504311801561080657508015155b1561086c57600061081b8560020154436117a5565b905060006108486009546105f48860010154610842600454876119d990919063ffffffff16565b906119d9565b9050610867610860846105f48464e8d4a510006119d9565b8590611bc5565b935050505b61089483600101546105fa64e8d4a510006105f48688600001546119d990919063ffffffff16565b9450505050505b92915050565b600781815481106108ae57fe5b600091825260209091206004909102018054600182015460028301546003909301546001600160a01b039092169350919084565b60095481565b6108f0611c1f565b6000546001600160a01b03908116911614610952576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b8015610960576109606114da565b6000600a54431161097357600a54610975565b435b6009549091506109859085611bc5565b600955604080516080810182526001600160a01b0385811682526020820187815292820184815260006060840181815260078054600181018255925293517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688600490920291820180547fffffffffffffffffffffffff000000000000000000000000000000000000000016919094161790925592517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68982015591517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68a830155517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68b90910155610a92611c23565b50505050565b610aa0611c1f565b6000546001600160a01b03908116911614610b02576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600680547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60006007600081548110610b4c57fe5b600091825260208083203384527f5eff886ea0ce6ca488a3d6e336d6c0f75f46d19b42c06ce5ee98e42c96d256c790915260408320600490920201925090610b9390611178565b805415610bdc576000610bc882600101546105fa64e8d4a510006105f4876003015487600001546119d990919063ffffffff16565b90508015610bda57610bda3382611ab6565b505b8215610c08578154610bf9906001600160a01b0316333086611cc5565b8054610c059084611bc5565b81555b60038201548154610c239164e8d4a51000916105f4916119d9565b6001820155600254604080517f40c10f190000000000000000000000000000000000000000000000000000000081523360048201526024810186905290516001600160a01b03909216916340c10f199160448082019260009290919082900301818387803b158015610c9457600080fd5b505af1158015610ca8573d6000803e3d6000fd5b5050604080518681529051600093503392507f90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a159181900360200190a3505050565b81610d3b576040805162461bcd60e51b815260206004820152601a60248201527f77697468647261772043414b4520627920756e7374616b696e67000000000000604482015290519081900360640190fd5b600060078381548110610d4a57fe5b600091825260208083208684526008825260408085203386529092529220805460049092029092019250831115610dc8576040805162461bcd60e51b815260206004820152601260248201527f77697468647261773a206e6f7420676f6f640000000000000000000000000000604482015290519081900360640190fd5b610dd184611178565b6000610dff82600101546105fa64e8d4a510006105f4876003015487600001546119d990919063ffffffff16565b90508015610e1157610e113382611ab6565b8315610e3b578154610e239085611a74565b82558254610e3b906001600160a01b03163386611b40565b60038301548254610e569164e8d4a51000916105f4916119d9565b6001830155604080518581529051869133917ff279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b5689181900360200190a35050505050565b6006546001600160a01b0316610ef6576040805162461bcd60e51b815260206004820152601460248201527f6d6967726174653a206e6f206d69677261746f72000000000000000000000000604482015290519081900360640190fd5b600060078281548110610f0557fe5b600091825260208083206004928302018054604080517f70a082310000000000000000000000000000000000000000000000000000000081523095810195909552519195506001600160a01b0316939284926370a0823192602480840193829003018186803b158015610f7757600080fd5b505afa158015610f8b573d6000803e3d6000fd5b505050506040513d6020811015610fa157600080fd5b5051600654909150610fc0906001600160a01b03848116911683611d4d565b600654604080517fce5494bb0000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301529151600093929092169163ce5494bb9160248082019260209290919082900301818787803b15801561102b57600080fd5b505af115801561103f573d6000803e3d6000fd5b505050506040513d602081101561105557600080fd5b5051604080517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290519192506001600160a01b038316916370a0823191602480820192602092909190829003018186803b1580156110ba57600080fd5b505afa1580156110ce573d6000803e3d6000fd5b505050506040513d60208110156110e457600080fd5b50518214611139576040805162461bcd60e51b815260206004820152600c60248201527f6d6967726174653a206261640000000000000000000000000000000000000000604482015290519081900360640190fd5b83547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039190911617909255505050565b600a5481565b60006007828154811061118757fe5b90600052602060002090600402019050806002015443116111a857506113cd565b8054604080517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290516000926001600160a01b0316916370a08231916024808301926020929190829003018186803b15801561120b57600080fd5b505afa15801561121f573d6000803e3d6000fd5b505050506040513d602081101561123557600080fd5b505190508061124b5750436002909101556113cd565b600061125b8360020154436117a5565b905060006112826009546105f48660010154610842600454876119d990919063ffffffff16565b6001546003549192506001600160a01b03908116916340c10f1991166112a984600a611a32565b6040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050600060405180830381600087803b1580156112ef57600080fd5b505af1158015611303573d6000803e3d6000fd5b5050600154600254604080517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810187905290519190921693506340c10f199250604480830192600092919082900301818387803b15801561137857600080fd5b505af115801561138c573d6000803e3d6000fd5b505050506113ba6113af846105f464e8d4a51000856119d990919063ffffffff16565b600386015490611bc5565b6003850155505043600290920191909155505b50565b6000600782815481106113df57fe5b60009182526020808320858452600882526040808520338087529352909320805460049093029093018054909450611424926001600160a01b03919091169190611b40565b80546040805191825251849133917fbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae05959181900360200190a360008082556001909101555050565b611473611c1f565b6000546001600160a01b039081169116146114d5576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600555565b60075460005b818110156114f9576114f181611178565b6001016114e0565b5050565b611505611c1f565b6000546001600160a01b03908116911614611567576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b8015611575576115756114da565b6115b2826115ac6007868154811061158957fe5b906000526020600020906004020160010154600954611a7490919063ffffffff16565b90611bc5565b6009819055506000600784815481106115c757fe5b906000526020600020906004020160010154905082600785815481106115e957fe5b906000526020600020906004020160010181905550828114610a9257610a92611c23565b611615611c1f565b6000546001600160a01b03908116911614611677576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055565b6006546001600160a01b031681565b6002546001600160a01b031681565b60055481565b6003546001600160a01b0316331461175c576040805162461bcd60e51b815260206004820152600960248201527f6465763a207775743f0000000000000000000000000000000000000000000000604482015290519081900360640190fd5b600380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6000546001600160a01b031690565b6005546000906117b9906108428486611a74565b9392505050565b60086020908152600092835260408084209091529082529020805460019091015482565b6003546001600160a01b031681565b6001546001600160a01b031681565b81611854576040805162461bcd60e51b815260206004820152601760248201527f6465706f7369742043414b45206279207374616b696e67000000000000000000604482015290519081900360640190fd5b60006007838154811061186357fe5b6000918252602080832086845260088252604080852033865290925292206004909102909101915061189484611178565b8054156118dd5760006118c982600101546105fa64e8d4a510006105f4876003015487600001546119d990919063ffffffff16565b905080156118db576118db3382611ab6565b505b82156119095781546118fa906001600160a01b0316333086611cc5565b80546119069084611bc5565b81555b600382015481546119249164e8d4a51000916105f4916119d9565b6001820155604080518481529051859133917f90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a159181900360200190a350505050565b61196e611c1f565b6000546001600160a01b039081169116146119d0576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6113cd81611ea7565b6000826119e85750600061089b565b828202828482816119f557fe5b04146117b95760405162461bcd60e51b81526004018080602001828103825260218152602001806123386021913960400191505060405180910390fd5b60006117b983836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250611f5f565b60006117b983836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250612001565b600254604080517fa2e6ddcc0000000000000000000000000000000000000000000000000000000081526001600160a01b038581166004830152602482018590529151919092169163a2e6ddcc91604480830192600092919082900301818387803b158015611b2457600080fd5b505af1158015611b38573d6000803e3d6000fd5b505050505050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052611bc090849061205b565b505050565b6000828201838110156117b9576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b3390565b600754600060015b82811015611c6e57611c6460078281548110611c4357fe5b90600052602060002090600402016001015483611bc590919063ffffffff16565b9150600101611c2b565b5080156114f957611c80816003611a32565b9050611c97816115ac600760008154811061158957fe5b600981905550806007600081548110611cac57fe5b9060005260206000209060040201600101819055505050565b604080516001600160a01b0380861660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd00000000000000000000000000000000000000000000000000000000179052610a9290859061205b565b801580611dec5750604080517fdd62ed3e0000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b03848116602483015291519185169163dd62ed3e91604480820192602092909190829003018186803b158015611dbe57600080fd5b505afa158015611dd2573d6000803e3d6000fd5b505050506040513d6020811015611de857600080fd5b5051155b611e275760405162461bcd60e51b81526004018080602001828103825260368152602001806123596036913960400191505060405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f095ea7b300000000000000000000000000000000000000000000000000000000179052611bc090849061205b565b6001600160a01b038116611eec5760405162461bcd60e51b81526004018080602001828103825260268152602001806123126026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60008183611feb5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611fb0578181015183820152602001611f98565b50505050905090810190601f168015611fdd5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506000838581611ff757fe5b0495945050505050565b600081848411156120535760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315611fb0578181015183820152602001611f98565b505050900390565b60606120b0826040518060400160405280602081526020017f5361666542455032303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661210c9092919063ffffffff16565b805190915015611bc0578080602001905160208110156120cf57600080fd5b5051611bc05760405162461bcd60e51b815260040180806020018281038252602a8152602001806122e8602a913960400191505060405180910390fd5b606061211b8484600085612123565b949350505050565b606061212e856122ae565b61217f576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b60006060866001600160a01b031685876040518082805190602001908083835b602083106121dc57805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0909201916020918201910161219f565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d806000811461223e576040519150601f19603f3d011682016040523d82523d6000602084013e612243565b606091505b5091509150811561225757915061211b9050565b8051156122675780518082602001fd5b60405162461bcd60e51b8152602060048201818152865160248401528651879391928392604401919085019080838360008315611fb0578181015183820152602001611f98565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47081811480159061211b57505015159291505056fe5361666542455032303a204245503230206f7065726174696f6e20646964206e6f7420737563636565644f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f775361666542455032303a20617070726f76652066726f6d206e6f6e2d7a65726f20746f206e6f6e2d7a65726f20616c6c6f77616e6365a2646970667358221220077495f6efd1597cf54586cb058707bf146388300d17c943ece392463f56203964736f6c634300060c00330000000000000000000000000e09fabb73bd3ade0a17ecc321fd13a19e81ce82000000000000000000000000009cf7bc57584b7998236eff51b98a168dcea9b0000000000000000000000000b9fa21a62fc96cb2ac635a051061e2e50d9640510000000000000000000000000000000000000000000000022b1c8c1227a0000000000000000000000000000000000000000000000000000000000000000abd4c"

// DeployFarm deploys a new Ethereum contract, binding an instance of Farm to it.
func DeployFarm(auth *bind.TransactOpts, backend bind.ContractBackend, _cake common.Address, _syrup common.Address, _devaddr common.Address, _cakePerBlock *big.Int, _startBlock *big.Int) (common.Address, *types.Transaction, *Farm, error) {
	parsed, err := abi.JSON(strings.NewReader(FarmABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FarmBin), backend, _cake, _syrup, _devaddr, _cakePerBlock, _startBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Farm{FarmCaller: FarmCaller{contract: contract}, FarmTransactor: FarmTransactor{contract: contract}, FarmFilterer: FarmFilterer{contract: contract}}, nil
}

// Farm is an auto generated Go binding around an Ethereum contract.
type Farm struct {
	FarmCaller     // Read-only binding to the contract
	FarmTransactor // Write-only binding to the contract
	FarmFilterer   // Log filterer for contract events
}

// FarmCaller is an auto generated read-only Go binding around an Ethereum contract.
type FarmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FarmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FarmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FarmSession struct {
	Contract     *Farm             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FarmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FarmCallerSession struct {
	Contract *FarmCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FarmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FarmTransactorSession struct {
	Contract     *FarmTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FarmRaw is an auto generated low-level Go binding around an Ethereum contract.
type FarmRaw struct {
	Contract *Farm // Generic contract binding to access the raw methods on
}

// FarmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FarmCallerRaw struct {
	Contract *FarmCaller // Generic read-only contract binding to access the raw methods on
}

// FarmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FarmTransactorRaw struct {
	Contract *FarmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFarm creates a new instance of Farm, bound to a specific deployed contract.
func NewFarm(address common.Address, backend bind.ContractBackend) (*Farm, error) {
	contract, err := bindFarm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Farm{FarmCaller: FarmCaller{contract: contract}, FarmTransactor: FarmTransactor{contract: contract}, FarmFilterer: FarmFilterer{contract: contract}}, nil
}

// NewFarmCaller creates a new read-only instance of Farm, bound to a specific deployed contract.
func NewFarmCaller(address common.Address, caller bind.ContractCaller) (*FarmCaller, error) {
	contract, err := bindFarm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FarmCaller{contract: contract}, nil
}

// NewFarmTransactor creates a new write-only instance of Farm, bound to a specific deployed contract.
func NewFarmTransactor(address common.Address, transactor bind.ContractTransactor) (*FarmTransactor, error) {
	contract, err := bindFarm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FarmTransactor{contract: contract}, nil
}

// NewFarmFilterer creates a new log filterer instance of Farm, bound to a specific deployed contract.
func NewFarmFilterer(address common.Address, filterer bind.ContractFilterer) (*FarmFilterer, error) {
	contract, err := bindFarm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FarmFilterer{contract: contract}, nil
}

// bindFarm binds a generic wrapper to an already deployed contract.
func bindFarm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FarmABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Farm *FarmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Farm.Contract.FarmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Farm *FarmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Farm.Contract.FarmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Farm *FarmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Farm.Contract.FarmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Farm *FarmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Farm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Farm *FarmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Farm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Farm *FarmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Farm.Contract.contract.Transact(opts, method, params...)
}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() view returns(uint256)
func (_Farm *FarmCaller) BONUSMULTIPLIER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "BONUS_MULTIPLIER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() view returns(uint256)
func (_Farm *FarmSession) BONUSMULTIPLIER() (*big.Int, error) {
	return _Farm.Contract.BONUSMULTIPLIER(&_Farm.CallOpts)
}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() view returns(uint256)
func (_Farm *FarmCallerSession) BONUSMULTIPLIER() (*big.Int, error) {
	return _Farm.Contract.BONUSMULTIPLIER(&_Farm.CallOpts)
}

// Cake is a free data retrieval call binding the contract method 0xdce17484.
//
// Solidity: function cake() view returns(address)
func (_Farm *FarmCaller) Cake(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "cake")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cake is a free data retrieval call binding the contract method 0xdce17484.
//
// Solidity: function cake() view returns(address)
func (_Farm *FarmSession) Cake() (common.Address, error) {
	return _Farm.Contract.Cake(&_Farm.CallOpts)
}

// Cake is a free data retrieval call binding the contract method 0xdce17484.
//
// Solidity: function cake() view returns(address)
func (_Farm *FarmCallerSession) Cake() (common.Address, error) {
	return _Farm.Contract.Cake(&_Farm.CallOpts)
}

// CakePerBlock is a free data retrieval call binding the contract method 0x0755e0b6.
//
// Solidity: function cakePerBlock() view returns(uint256)
func (_Farm *FarmCaller) CakePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "cakePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CakePerBlock is a free data retrieval call binding the contract method 0x0755e0b6.
//
// Solidity: function cakePerBlock() view returns(uint256)
func (_Farm *FarmSession) CakePerBlock() (*big.Int, error) {
	return _Farm.Contract.CakePerBlock(&_Farm.CallOpts)
}

// CakePerBlock is a free data retrieval call binding the contract method 0x0755e0b6.
//
// Solidity: function cakePerBlock() view returns(uint256)
func (_Farm *FarmCallerSession) CakePerBlock() (*big.Int, error) {
	return _Farm.Contract.CakePerBlock(&_Farm.CallOpts)
}

// Devaddr is a free data retrieval call binding the contract method 0xd49e77cd.
//
// Solidity: function devaddr() view returns(address)
func (_Farm *FarmCaller) Devaddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "devaddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Devaddr is a free data retrieval call binding the contract method 0xd49e77cd.
//
// Solidity: function devaddr() view returns(address)
func (_Farm *FarmSession) Devaddr() (common.Address, error) {
	return _Farm.Contract.Devaddr(&_Farm.CallOpts)
}

// Devaddr is a free data retrieval call binding the contract method 0xd49e77cd.
//
// Solidity: function devaddr() view returns(address)
func (_Farm *FarmCallerSession) Devaddr() (common.Address, error) {
	return _Farm.Contract.Devaddr(&_Farm.CallOpts)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_Farm *FarmCaller) GetMultiplier(opts *bind.CallOpts, _from *big.Int, _to *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "getMultiplier", _from, _to)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_Farm *FarmSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _Farm.Contract.GetMultiplier(&_Farm.CallOpts, _from, _to)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_Farm *FarmCallerSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _Farm.Contract.GetMultiplier(&_Farm.CallOpts, _from, _to)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_Farm *FarmCaller) Migrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "migrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_Farm *FarmSession) Migrator() (common.Address, error) {
	return _Farm.Contract.Migrator(&_Farm.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_Farm *FarmCallerSession) Migrator() (common.Address, error) {
	return _Farm.Contract.Migrator(&_Farm.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Farm *FarmCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Farm *FarmSession) Owner() (common.Address, error) {
	return _Farm.Contract.Owner(&_Farm.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Farm *FarmCallerSession) Owner() (common.Address, error) {
	return _Farm.Contract.Owner(&_Farm.CallOpts)
}

// PendingCake is a free data retrieval call binding the contract method 0x1175a1dd.
//
// Solidity: function pendingCake(uint256 _pid, address _user) view returns(uint256)
func (_Farm *FarmCaller) PendingCake(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "pendingCake", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingCake is a free data retrieval call binding the contract method 0x1175a1dd.
//
// Solidity: function pendingCake(uint256 _pid, address _user) view returns(uint256)
func (_Farm *FarmSession) PendingCake(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _Farm.Contract.PendingCake(&_Farm.CallOpts, _pid, _user)
}

// PendingCake is a free data retrieval call binding the contract method 0x1175a1dd.
//
// Solidity: function pendingCake(uint256 _pid, address _user) view returns(uint256)
func (_Farm *FarmCallerSession) PendingCake(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _Farm.Contract.PendingCake(&_Farm.CallOpts, _pid, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accCakePerShare)
func (_Farm *FarmCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccCakePerShare *big.Int
}, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		LpToken         common.Address
		AllocPoint      *big.Int
		LastRewardBlock *big.Int
		AccCakePerShare *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AllocPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastRewardBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccCakePerShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accCakePerShare)
func (_Farm *FarmSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccCakePerShare *big.Int
}, error) {
	return _Farm.Contract.PoolInfo(&_Farm.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accCakePerShare)
func (_Farm *FarmCallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccCakePerShare *big.Int
}, error) {
	return _Farm.Contract.PoolInfo(&_Farm.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Farm *FarmCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Farm *FarmSession) PoolLength() (*big.Int, error) {
	return _Farm.Contract.PoolLength(&_Farm.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Farm *FarmCallerSession) PoolLength() (*big.Int, error) {
	return _Farm.Contract.PoolLength(&_Farm.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Farm *FarmCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Farm *FarmSession) StartBlock() (*big.Int, error) {
	return _Farm.Contract.StartBlock(&_Farm.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Farm *FarmCallerSession) StartBlock() (*big.Int, error) {
	return _Farm.Contract.StartBlock(&_Farm.CallOpts)
}

// Syrup is a free data retrieval call binding the contract method 0x86a952c4.
//
// Solidity: function syrup() view returns(address)
func (_Farm *FarmCaller) Syrup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "syrup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Syrup is a free data retrieval call binding the contract method 0x86a952c4.
//
// Solidity: function syrup() view returns(address)
func (_Farm *FarmSession) Syrup() (common.Address, error) {
	return _Farm.Contract.Syrup(&_Farm.CallOpts)
}

// Syrup is a free data retrieval call binding the contract method 0x86a952c4.
//
// Solidity: function syrup() view returns(address)
func (_Farm *FarmCallerSession) Syrup() (common.Address, error) {
	return _Farm.Contract.Syrup(&_Farm.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Farm *FarmCaller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Farm *FarmSession) TotalAllocPoint() (*big.Int, error) {
	return _Farm.Contract.TotalAllocPoint(&_Farm.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Farm *FarmCallerSession) TotalAllocPoint() (*big.Int, error) {
	return _Farm.Contract.TotalAllocPoint(&_Farm.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Farm *FarmCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "userInfo", arg0, arg1)

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

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Farm *FarmSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Farm.Contract.UserInfo(&_Farm.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Farm *FarmCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Farm.Contract.UserInfo(&_Farm.CallOpts, arg0, arg1)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_Farm *FarmTransactor) Add(opts *bind.TransactOpts, _allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "add", _allocPoint, _lpToken, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_Farm *FarmSession) Add(_allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _Farm.Contract.Add(&_Farm.TransactOpts, _allocPoint, _lpToken, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_Farm *FarmTransactorSession) Add(_allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _Farm.Contract.Add(&_Farm.TransactOpts, _allocPoint, _lpToken, _withUpdate)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Farm *FarmTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Farm *FarmSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Farm.Contract.Deposit(&_Farm.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Farm *FarmTransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Farm.Contract.Deposit(&_Farm.TransactOpts, _pid, _amount)
}

// Dev is a paid mutator transaction binding the contract method 0x8d88a90e.
//
// Solidity: function dev(address _devaddr) returns()
func (_Farm *FarmTransactor) Dev(opts *bind.TransactOpts, _devaddr common.Address) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "dev", _devaddr)
}

// Dev is a paid mutator transaction binding the contract method 0x8d88a90e.
//
// Solidity: function dev(address _devaddr) returns()
func (_Farm *FarmSession) Dev(_devaddr common.Address) (*types.Transaction, error) {
	return _Farm.Contract.Dev(&_Farm.TransactOpts, _devaddr)
}

// Dev is a paid mutator transaction binding the contract method 0x8d88a90e.
//
// Solidity: function dev(address _devaddr) returns()
func (_Farm *FarmTransactorSession) Dev(_devaddr common.Address) (*types.Transaction, error) {
	return _Farm.Contract.Dev(&_Farm.TransactOpts, _devaddr)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Farm *FarmTransactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Farm *FarmSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _Farm.Contract.EmergencyWithdraw(&_Farm.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Farm *FarmTransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _Farm.Contract.EmergencyWithdraw(&_Farm.TransactOpts, _pid)
}

// EnterStaking is a paid mutator transaction binding the contract method 0x41441d3b.
//
// Solidity: function enterStaking(uint256 _amount) returns()
func (_Farm *FarmTransactor) EnterStaking(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "enterStaking", _amount)
}

// EnterStaking is a paid mutator transaction binding the contract method 0x41441d3b.
//
// Solidity: function enterStaking(uint256 _amount) returns()
func (_Farm *FarmSession) EnterStaking(_amount *big.Int) (*types.Transaction, error) {
	return _Farm.Contract.EnterStaking(&_Farm.TransactOpts, _amount)
}

// EnterStaking is a paid mutator transaction binding the contract method 0x41441d3b.
//
// Solidity: function enterStaking(uint256 _amount) returns()
func (_Farm *FarmTransactorSession) EnterStaking(_amount *big.Int) (*types.Transaction, error) {
	return _Farm.Contract.EnterStaking(&_Farm.TransactOpts, _amount)
}

// LeaveStaking is a paid mutator transaction binding the contract method 0x1058d281.
//
// Solidity: function leaveStaking(uint256 _amount) returns()
func (_Farm *FarmTransactor) LeaveStaking(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "leaveStaking", _amount)
}

// LeaveStaking is a paid mutator transaction binding the contract method 0x1058d281.
//
// Solidity: function leaveStaking(uint256 _amount) returns()
func (_Farm *FarmSession) LeaveStaking(_amount *big.Int) (*types.Transaction, error) {
	return _Farm.Contract.LeaveStaking(&_Farm.TransactOpts, _amount)
}

// LeaveStaking is a paid mutator transaction binding the contract method 0x1058d281.
//
// Solidity: function leaveStaking(uint256 _amount) returns()
func (_Farm *FarmTransactorSession) LeaveStaking(_amount *big.Int) (*types.Transaction, error) {
	return _Farm.Contract.LeaveStaking(&_Farm.TransactOpts, _amount)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Farm *FarmTransactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Farm *FarmSession) MassUpdatePools() (*types.Transaction, error) {
	return _Farm.Contract.MassUpdatePools(&_Farm.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Farm *FarmTransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _Farm.Contract.MassUpdatePools(&_Farm.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_Farm *FarmTransactor) Migrate(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "migrate", _pid)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_Farm *FarmSession) Migrate(_pid *big.Int) (*types.Transaction, error) {
	return _Farm.Contract.Migrate(&_Farm.TransactOpts, _pid)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_Farm *FarmTransactorSession) Migrate(_pid *big.Int) (*types.Transaction, error) {
	return _Farm.Contract.Migrate(&_Farm.TransactOpts, _pid)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Farm *FarmTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Farm *FarmSession) RenounceOwnership() (*types.Transaction, error) {
	return _Farm.Contract.RenounceOwnership(&_Farm.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Farm *FarmTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Farm.Contract.RenounceOwnership(&_Farm.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_Farm *FarmTransactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "set", _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_Farm *FarmSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _Farm.Contract.Set(&_Farm.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_Farm *FarmTransactorSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _Farm.Contract.Set(&_Farm.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_Farm *FarmTransactor) SetMigrator(opts *bind.TransactOpts, _migrator common.Address) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "setMigrator", _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_Farm *FarmSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _Farm.Contract.SetMigrator(&_Farm.TransactOpts, _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_Farm *FarmTransactorSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _Farm.Contract.SetMigrator(&_Farm.TransactOpts, _migrator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Farm *FarmTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Farm *FarmSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Farm.Contract.TransferOwnership(&_Farm.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Farm *FarmTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Farm.Contract.TransferOwnership(&_Farm.TransactOpts, newOwner)
}

// UpdateMultiplier is a paid mutator transaction binding the contract method 0x5ffe6146.
//
// Solidity: function updateMultiplier(uint256 multiplierNumber) returns()
func (_Farm *FarmTransactor) UpdateMultiplier(opts *bind.TransactOpts, multiplierNumber *big.Int) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "updateMultiplier", multiplierNumber)
}

// UpdateMultiplier is a paid mutator transaction binding the contract method 0x5ffe6146.
//
// Solidity: function updateMultiplier(uint256 multiplierNumber) returns()
func (_Farm *FarmSession) UpdateMultiplier(multiplierNumber *big.Int) (*types.Transaction, error) {
	return _Farm.Contract.UpdateMultiplier(&_Farm.TransactOpts, multiplierNumber)
}

// UpdateMultiplier is a paid mutator transaction binding the contract method 0x5ffe6146.
//
// Solidity: function updateMultiplier(uint256 multiplierNumber) returns()
func (_Farm *FarmTransactorSession) UpdateMultiplier(multiplierNumber *big.Int) (*types.Transaction, error) {
	return _Farm.Contract.UpdateMultiplier(&_Farm.TransactOpts, multiplierNumber)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Farm *FarmTransactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Farm *FarmSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _Farm.Contract.UpdatePool(&_Farm.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Farm *FarmTransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _Farm.Contract.UpdatePool(&_Farm.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Farm *FarmTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Farm *FarmSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Farm.Contract.Withdraw(&_Farm.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Farm *FarmTransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Farm.Contract.Withdraw(&_Farm.TransactOpts, _pid, _amount)
}

// FarmDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Farm contract.
type FarmDepositIterator struct {
	Event *FarmDeposit // Event containing the contract specifics and raw log

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
func (it *FarmDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmDeposit)
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
		it.Event = new(FarmDeposit)
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
func (it *FarmDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmDeposit represents a Deposit event raised by the Farm contract.
type FarmDeposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Farm *FarmFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*FarmDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Farm.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &FarmDepositIterator{contract: _Farm.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Farm *FarmFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *FarmDeposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Farm.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmDeposit)
				if err := _Farm.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_Farm *FarmFilterer) ParseDeposit(log types.Log) (*FarmDeposit, error) {
	event := new(FarmDeposit)
	if err := _Farm.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the Farm contract.
type FarmEmergencyWithdrawIterator struct {
	Event *FarmEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *FarmEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmEmergencyWithdraw)
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
		it.Event = new(FarmEmergencyWithdraw)
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
func (it *FarmEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmEmergencyWithdraw represents a EmergencyWithdraw event raised by the Farm contract.
type FarmEmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Farm *FarmFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*FarmEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Farm.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &FarmEmergencyWithdrawIterator{contract: _Farm.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Farm *FarmFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *FarmEmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Farm.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmEmergencyWithdraw)
				if err := _Farm.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Farm *FarmFilterer) ParseEmergencyWithdraw(log types.Log) (*FarmEmergencyWithdraw, error) {
	event := new(FarmEmergencyWithdraw)
	if err := _Farm.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Farm contract.
type FarmOwnershipTransferredIterator struct {
	Event *FarmOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FarmOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmOwnershipTransferred)
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
		it.Event = new(FarmOwnershipTransferred)
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
func (it *FarmOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmOwnershipTransferred represents a OwnershipTransferred event raised by the Farm contract.
type FarmOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Farm *FarmFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FarmOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Farm.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FarmOwnershipTransferredIterator{contract: _Farm.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Farm *FarmFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FarmOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Farm.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmOwnershipTransferred)
				if err := _Farm.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Farm *FarmFilterer) ParseOwnershipTransferred(log types.Log) (*FarmOwnershipTransferred, error) {
	event := new(FarmOwnershipTransferred)
	if err := _Farm.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Farm contract.
type FarmWithdrawIterator struct {
	Event *FarmWithdraw // Event containing the contract specifics and raw log

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
func (it *FarmWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmWithdraw)
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
		it.Event = new(FarmWithdraw)
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
func (it *FarmWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmWithdraw represents a Withdraw event raised by the Farm contract.
type FarmWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Farm *FarmFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*FarmWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Farm.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &FarmWithdrawIterator{contract: _Farm.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Farm *FarmFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *FarmWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Farm.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmWithdraw)
				if err := _Farm.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_Farm *FarmFilterer) ParseWithdraw(log types.Log) (*FarmWithdraw, error) {
	event := new(FarmWithdraw)
	if err := _Farm.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

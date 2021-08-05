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

// PoolABI is the input ABI used to generate the binding from.
const PoolABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenRecovered\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AdminTokenRecovery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolLimitPerUser\",\"type\":\"uint256\"}],\"name\":\"NewPoolLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPerBlock\",\"type\":\"uint256\"}],\"name\":\"NewRewardPerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"name\":\"NewStartAndEndBlocks\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"RewardsStop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PRECISION_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SMART_CHEF_FACTORY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accTokenPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusEndBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"emergencyRewardWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasUserLimit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"_stakedToken\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bonusEndBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_poolLimitPerUser\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRewardBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLimitPerUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"}],\"name\":\"recoverWrongTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedToken\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_hasUserLimit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_poolLimitPerUser\",\"type\":\"uint256\"}],\"name\":\"updatePoolLimitPerUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardPerBlock\",\"type\":\"uint256\"}],\"name\":\"updateRewardPerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bonusEndBlock\",\"type\":\"uint256\"}],\"name\":\"updateStartAndEndBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PoolBin is the compiled bytecode used for deploying new contracts.
var PoolBin = "0x608060405234801561001057600080fd5b50600061001b610080565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35060018055600280546001600160a01b03191633179055610084565b3390565b61244a806100936000396000f3fe608060405234801561001057600080fd5b50600436106101b95760003560e01c80638da5cb5b116100f9578063bd61719111610097578063db2e21bc11610071578063db2e21bc1461041d578063f2fde38b14610425578063f40f0f5214610458578063f7c618c11461048b576101b9565b8063bd61719114610405578063cc7a262e1461040d578063ccd34cd514610415576101b9565b80639513997f116100d35780639513997f14610398578063a0b40905146103bb578063a9f8d181146103e0578063b6b55f25146103e8576101b9565b80638da5cb5b146103575780638f6629151461038857806392e8990e14610390576101b9565b80633f138d4b11610166578063715018a611610140578063715018a6146102e257806380dc0672146102ea57806383a5041c146102f25780638ae39cac1461034f576101b9565b80633f138d4b1461029957806348cd4cb1146102d257806366fe9f8a146102da576101b9565b80632e1a7d4d116101975780632e1a7d4d146102435780633279beab14610260578063392e53cd1461027d576101b9565b806301f8a976146101be5780631959a002146101dd5780631aed655314610229575b600080fd5b6101db600480360360208110156101d457600080fd5b5035610493565b005b610210600480360360208110156101f357600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166105e6565b6040805192835260208301919091528051918290030190f35b6102316105ff565b60408051918252519081900360200190f35b6101db6004803603602081101561025957600080fd5b5035610605565b6101db6004803603602081101561027657600080fd5b50356107f7565b6102856108c6565b604080519115158252519081900360200190f35b6101db600480360360408110156102af57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356108e8565b610231610b1a565b610231610b20565b6101db610b26565b6101db610c3d565b6101db600480360360e081101561030857600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135811691602081013582169160408201359160608101359160808201359160a08101359160c09091013516610ceb565b610231611020565b61035f611026565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b610231611042565b610285611048565b6101db600480360360408110156103ae57600080fd5b5080359060200135611069565b6101db600480360360408110156103d157600080fd5b5080351515906020013561127f565b6102316114b2565b6101db600480360360208110156103fe57600080fd5b50356114b8565b61035f6116d8565b61035f6116f4565b610231611710565b6101db611716565b6101db6004803603602081101561043b57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16611814565b6102316004803603602081101561046e57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166119b5565b61035f611b32565b61049b611b4e565b73ffffffffffffffffffffffffffffffffffffffff166104b9611026565b73ffffffffffffffffffffffffffffffffffffffff161461053b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b60055443106105ab57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f506f6f6c20686173207374617274656400000000000000000000000000000000604482015290519081900360640190fd5b60088190556040805182815290517f0c4d677eef92893ac7ec52faf8140fc6c851ab4736302b4f3a89dfb20696a0df9181900360200190a150565b600c602052600090815260409020805460019091015482565b60045481565b6002600154141561067757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6002600155336000908152600c6020526040902080548211156106fb57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f416d6f756e7420746f20776974686472617720746f6f20686967680000000000604482015290519081900360640190fd5b610703611b52565b6000610738826001015461073260095461072c6003548760000154611c6090919063ffffffff16565b90611cdc565b90611d5d565b9050821561077257815461074c9084611d5d565b8255600b546107729073ffffffffffffffffffffffffffffffffffffffff163385611dd4565b801561079c57600a5461079c9073ffffffffffffffffffffffffffffffffffffffff163383611dd4565b60095460035483546107b3929161072c9190611c60565b600183015560408051848152905133917f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364919081900360200190a250506001805550565b6107ff611b4e565b73ffffffffffffffffffffffffffffffffffffffff1661081d611026565b73ffffffffffffffffffffffffffffffffffffffff161461089f57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600a546108c39073ffffffffffffffffffffffffffffffffffffffff163383611dd4565b50565b6002547501000000000000000000000000000000000000000000900460ff1681565b6108f0611b4e565b73ffffffffffffffffffffffffffffffffffffffff1661090e611026565b73ffffffffffffffffffffffffffffffffffffffff161461099057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600b5473ffffffffffffffffffffffffffffffffffffffff83811691161415610a1a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f43616e6e6f74206265207374616b656420746f6b656e00000000000000000000604482015290519081900360640190fd5b600a5473ffffffffffffffffffffffffffffffffffffffff83811691161415610aa457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f43616e6e6f742062652072657761726420746f6b656e00000000000000000000604482015290519081900360640190fd5b610ac573ffffffffffffffffffffffffffffffffffffffff83163383611dd4565b6040805173ffffffffffffffffffffffffffffffffffffffff841681526020810183905281517f74545154aac348a3eac92596bd1971957ca94795f4e954ec5f613b55fab78129929181900390910190a15050565b60055481565b60075481565b610b2e611b4e565b73ffffffffffffffffffffffffffffffffffffffff16610b4c611026565b73ffffffffffffffffffffffffffffffffffffffff1614610bce57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6000805460405173ffffffffffffffffffffffffffffffffffffffff909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055565b610c45611b4e565b73ffffffffffffffffffffffffffffffffffffffff16610c63611026565b73ffffffffffffffffffffffffffffffffffffffff1614610ce557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b43600455565b6002547501000000000000000000000000000000000000000000900460ff1615610d7657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f416c726561647920696e697469616c697a656400000000000000000000000000604482015290519081900360640190fd5b60025473ffffffffffffffffffffffffffffffffffffffff163314610dfc57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4e6f7420666163746f7279000000000000000000000000000000000000000000604482015290519081900360640190fd5b600280547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff167501000000000000000000000000000000000000000000179055600b805473ffffffffffffffffffffffffffffffffffffffff808a167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255600a8054928916929091169190911790556008859055600584905560048390558115610eeb57600280547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000017905560078290555b600a54604080517f313ce567000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163313ce567916004808301926020929190829003018186803b158015610f5657600080fd5b505afa158015610f6a573d6000803e3d6000fd5b505050506040513d6020811015610f8057600080fd5b505160ff169050601e8110610ff657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d75737420626520696e666572696f7220746f20333000000000000000000000604482015290519081900360640190fd5b611001601e82611d5d565b600a0a60095560055460065561101682611814565b5050505050505050565b60085481565b60005473ffffffffffffffffffffffffffffffffffffffff1690565b60035481565b60025474010000000000000000000000000000000000000000900460ff1681565b611071611b4e565b73ffffffffffffffffffffffffffffffffffffffff1661108f611026565b73ffffffffffffffffffffffffffffffffffffffff161461111157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600554431061118157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f506f6f6c20686173207374617274656400000000000000000000000000000000604482015290519081900360640190fd5b8082106111d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e815260200180612370602e913960400191505060405180910390fd5b814310611231576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806123c46030913960400191505060405180910390fd5b600582905560048190556006829055604080518381526020810183905281517f7cd0ab87d19036f3dfadadb232c78aa4879dda3f0c994a9d637532410ee2ce06929181900390910190a15050565b611287611b4e565b73ffffffffffffffffffffffffffffffffffffffff166112a5611026565b73ffffffffffffffffffffffffffffffffffffffff161461132757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b60025474010000000000000000000000000000000000000000900460ff166113b057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4d75737420626520736574000000000000000000000000000000000000000000604482015290519081900360640190fd5b811561143057600754811161142657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4e6577206c696d6974206d757374206265206869676865720000000000000000604482015290519081900360640190fd5b6007819055611479565b600280547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000008415150217905560006007555b60075460408051918252517f241f67ee5f41b7a5cabf911367329be7215900f602ebfc47f89dce2a6bcd847c9181900360200190a15050565b60065481565b6002600154141561152a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b60026001819055336000908152600c60205260409020905474010000000000000000000000000000000000000000900460ff16156115df576007548154611572908490611e66565b11156115df57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f5573657220616d6f756e742061626f7665206c696d6974000000000000000000604482015290519081900360640190fd5b6115e7611b52565b805415611645576000611617826001015461073260095461072c6003548760000154611c6090919063ffffffff16565b9050801561164357600a546116439073ffffffffffffffffffffffffffffffffffffffff163383611dd4565b505b811561167e5780546116579083611e66565b8155600b5461167e9073ffffffffffffffffffffffffffffffffffffffff16333085611eda565b6009546003548254611695929161072c9190611c60565b600182015560408051838152905133917fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c919081900360200190a2505060018055565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b600b5473ffffffffffffffffffffffffffffffffffffffff1681565b60095481565b6002600154141561178857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b60026001908155336000908152600c602052604081208054828255928101919091559080156117d557600b546117d59073ffffffffffffffffffffffffffffffffffffffff163383611dd4565b8154604080519182525133917f5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695919081900360200190a2505060018055565b61181c611b4e565b73ffffffffffffffffffffffffffffffffffffffff1661183a611026565b73ffffffffffffffffffffffffffffffffffffffff16146118bc57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff8116611928576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602681526020018061234a6026913960400191505060405180910390fd5b6000805460405173ffffffffffffffffffffffffffffffffffffffff808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b73ffffffffffffffffffffffffffffffffffffffff8082166000908152600c60209081526040808320600b5482517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152925194959194869491909216926370a082319260248083019392829003018186803b158015611a3a57600080fd5b505afa158015611a4e573d6000803e3d6000fd5b505050506040513d6020811015611a6457600080fd5b505160065490915043118015611a7957508015155b15611b01576000611a8c60065443611f75565b90506000611aa560085483611c6090919063ffffffff16565b90506000611ace611ac58561072c60095486611c6090919063ffffffff16565b60035490611e66565b9050611af5856001015461073260095461072c858a60000154611c6090919063ffffffff16565b95505050505050611b2d565b611b28826001015461073260095461072c6003548760000154611c6090919063ffffffff16565b925050505b919050565b600a5473ffffffffffffffffffffffffffffffffffffffff1681565b3390565b6006544311611b6057611c5e565b600b54604080517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152905160009273ffffffffffffffffffffffffffffffffffffffff16916370a08231916024808301926020929190829003018186803b158015611bd157600080fd5b505afa158015611be5573d6000803e3d6000fd5b505050506040513d6020811015611bfb57600080fd5b5051905080611c0e575043600655611c5e565b6000611c1c60065443611f75565b90506000611c3560085483611c6090919063ffffffff16565b9050611c53611ac58461072c60095485611c6090919063ffffffff16565b600355505043600655505b565b600082611c6f57506000611cd6565b82820282848281611c7c57fe5b0414611cd3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806123f46021913960400191505060405180910390fd5b90505b92915050565b6000808211611d4c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b818381611d5557fe5b049392505050565b600082821115611dce57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052611e61908490611faf565b505050565b600082820183811015611cd357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6040805173ffffffffffffffffffffffffffffffffffffffff80861660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd00000000000000000000000000000000000000000000000000000000179052611f6f908590611faf565b50505050565b60006004548211611f9157611f8a8284611d5d565b9050611cd6565b6004548310611fa257506000611cd6565b600454611f8a9084611d5d565b6060612011826040518060400160405280602081526020017f5361666542455032303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166120879092919063ffffffff16565b805190915015611e615780806020019051602081101561203057600080fd5b5051611e61576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180612320602a913960400191505060405180910390fd5b606061209684846000856120a0565b90505b9392505050565b6060824710156120fb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602681526020018061239e6026913960400191505060405180910390fd5b6121048561225b565b61216f57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b600060608673ffffffffffffffffffffffffffffffffffffffff1685876040518082805190602001908083835b602083106121d957805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0909201916020918201910161219c565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d806000811461223b576040519150601f19603f3d011682016040523d82523d6000602084013e612240565b606091505b5091509150612250828286612261565b979650505050505050565b3b151590565b60608315612270575081612099565b8251156122805782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156122e45781810151838201526020016122cc565b50505050905090810190601f1680156123115780820380516001836020036101000a031916815260200191505b509250505060405180910390fdfe5361666542455032303a204245503230206f7065726174696f6e20646964206e6f7420737563636565644f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734e6577207374617274426c6f636b206d757374206265206c6f776572207468616e206e657720656e64426c6f636b416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c4e6577207374617274426c6f636b206d75737420626520686967686572207468616e2063757272656e7420626c6f636b536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a264697066735822122073a95bccadd3bc470e49a218813498f4abaf5fad7c08581560c7aa88c60f27c564736f6c634300060c0033"

// DeployPool deploys a new Ethereum contract, binding an instance of Pool to it.
func DeployPool(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pool, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PoolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// Pool is an auto generated Go binding around an Ethereum contract.
type Pool struct {
	PoolCaller     // Read-only binding to the contract
	PoolTransactor // Write-only binding to the contract
	PoolFilterer   // Log filterer for contract events
}

// PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolSession struct {
	Contract     *Pool             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolCallerSession struct {
	Contract *PoolCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolTransactorSession struct {
	Contract     *PoolTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolRaw struct {
	Contract *Pool // Generic contract binding to access the raw methods on
}

// PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolCallerRaw struct {
	Contract *PoolCaller // Generic read-only contract binding to access the raw methods on
}

// PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolTransactorRaw struct {
	Contract *PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPool creates a new instance of Pool, bound to a specific deployed contract.
func NewPool(address common.Address, backend bind.ContractBackend) (*Pool, error) {
	contract, err := bindPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// NewPoolCaller creates a new read-only instance of Pool, bound to a specific deployed contract.
func NewPoolCaller(address common.Address, caller bind.ContractCaller) (*PoolCaller, error) {
	contract, err := bindPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCaller{contract: contract}, nil
}

// NewPoolTransactor creates a new write-only instance of Pool, bound to a specific deployed contract.
func NewPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolTransactor, error) {
	contract, err := bindPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTransactor{contract: contract}, nil
}

// NewPoolFilterer creates a new log filterer instance of Pool, bound to a specific deployed contract.
func NewPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolFilterer, error) {
	contract, err := bindPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolFilterer{contract: contract}, nil
}

// bindPool binds a generic wrapper to an already deployed contract.
func bindPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transact(opts, method, params...)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_Pool *PoolCaller) PRECISIONFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "PRECISION_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_Pool *PoolSession) PRECISIONFACTOR() (*big.Int, error) {
	return _Pool.Contract.PRECISIONFACTOR(&_Pool.CallOpts)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_Pool *PoolCallerSession) PRECISIONFACTOR() (*big.Int, error) {
	return _Pool.Contract.PRECISIONFACTOR(&_Pool.CallOpts)
}

// SMARTCHEFFACTORY is a free data retrieval call binding the contract method 0xbd617191.
//
// Solidity: function SMART_CHEF_FACTORY() view returns(address)
func (_Pool *PoolCaller) SMARTCHEFFACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "SMART_CHEF_FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SMARTCHEFFACTORY is a free data retrieval call binding the contract method 0xbd617191.
//
// Solidity: function SMART_CHEF_FACTORY() view returns(address)
func (_Pool *PoolSession) SMARTCHEFFACTORY() (common.Address, error) {
	return _Pool.Contract.SMARTCHEFFACTORY(&_Pool.CallOpts)
}

// SMARTCHEFFACTORY is a free data retrieval call binding the contract method 0xbd617191.
//
// Solidity: function SMART_CHEF_FACTORY() view returns(address)
func (_Pool *PoolCallerSession) SMARTCHEFFACTORY() (common.Address, error) {
	return _Pool.Contract.SMARTCHEFFACTORY(&_Pool.CallOpts)
}

// AccTokenPerShare is a free data retrieval call binding the contract method 0x8f662915.
//
// Solidity: function accTokenPerShare() view returns(uint256)
func (_Pool *PoolCaller) AccTokenPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "accTokenPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccTokenPerShare is a free data retrieval call binding the contract method 0x8f662915.
//
// Solidity: function accTokenPerShare() view returns(uint256)
func (_Pool *PoolSession) AccTokenPerShare() (*big.Int, error) {
	return _Pool.Contract.AccTokenPerShare(&_Pool.CallOpts)
}

// AccTokenPerShare is a free data retrieval call binding the contract method 0x8f662915.
//
// Solidity: function accTokenPerShare() view returns(uint256)
func (_Pool *PoolCallerSession) AccTokenPerShare() (*big.Int, error) {
	return _Pool.Contract.AccTokenPerShare(&_Pool.CallOpts)
}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_Pool *PoolCaller) BonusEndBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "bonusEndBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_Pool *PoolSession) BonusEndBlock() (*big.Int, error) {
	return _Pool.Contract.BonusEndBlock(&_Pool.CallOpts)
}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_Pool *PoolCallerSession) BonusEndBlock() (*big.Int, error) {
	return _Pool.Contract.BonusEndBlock(&_Pool.CallOpts)
}

// HasUserLimit is a free data retrieval call binding the contract method 0x92e8990e.
//
// Solidity: function hasUserLimit() view returns(bool)
func (_Pool *PoolCaller) HasUserLimit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "hasUserLimit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasUserLimit is a free data retrieval call binding the contract method 0x92e8990e.
//
// Solidity: function hasUserLimit() view returns(bool)
func (_Pool *PoolSession) HasUserLimit() (bool, error) {
	return _Pool.Contract.HasUserLimit(&_Pool.CallOpts)
}

// HasUserLimit is a free data retrieval call binding the contract method 0x92e8990e.
//
// Solidity: function hasUserLimit() view returns(bool)
func (_Pool *PoolCallerSession) HasUserLimit() (bool, error) {
	return _Pool.Contract.HasUserLimit(&_Pool.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Pool *PoolCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Pool *PoolSession) IsInitialized() (bool, error) {
	return _Pool.Contract.IsInitialized(&_Pool.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Pool *PoolCallerSession) IsInitialized() (bool, error) {
	return _Pool.Contract.IsInitialized(&_Pool.CallOpts)
}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_Pool *PoolCaller) LastRewardBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "lastRewardBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_Pool *PoolSession) LastRewardBlock() (*big.Int, error) {
	return _Pool.Contract.LastRewardBlock(&_Pool.CallOpts)
}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_Pool *PoolCallerSession) LastRewardBlock() (*big.Int, error) {
	return _Pool.Contract.LastRewardBlock(&_Pool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pool *PoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pool *PoolSession) Owner() (common.Address, error) {
	return _Pool.Contract.Owner(&_Pool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pool *PoolCallerSession) Owner() (common.Address, error) {
	return _Pool.Contract.Owner(&_Pool.CallOpts)
}

// PendingReward is a free data retrieval call binding the contract method 0xf40f0f52.
//
// Solidity: function pendingReward(address _user) view returns(uint256)
func (_Pool *PoolCaller) PendingReward(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "pendingReward", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingReward is a free data retrieval call binding the contract method 0xf40f0f52.
//
// Solidity: function pendingReward(address _user) view returns(uint256)
func (_Pool *PoolSession) PendingReward(_user common.Address) (*big.Int, error) {
	return _Pool.Contract.PendingReward(&_Pool.CallOpts, _user)
}

// PendingReward is a free data retrieval call binding the contract method 0xf40f0f52.
//
// Solidity: function pendingReward(address _user) view returns(uint256)
func (_Pool *PoolCallerSession) PendingReward(_user common.Address) (*big.Int, error) {
	return _Pool.Contract.PendingReward(&_Pool.CallOpts, _user)
}

// PoolLimitPerUser is a free data retrieval call binding the contract method 0x66fe9f8a.
//
// Solidity: function poolLimitPerUser() view returns(uint256)
func (_Pool *PoolCaller) PoolLimitPerUser(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "poolLimitPerUser")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLimitPerUser is a free data retrieval call binding the contract method 0x66fe9f8a.
//
// Solidity: function poolLimitPerUser() view returns(uint256)
func (_Pool *PoolSession) PoolLimitPerUser() (*big.Int, error) {
	return _Pool.Contract.PoolLimitPerUser(&_Pool.CallOpts)
}

// PoolLimitPerUser is a free data retrieval call binding the contract method 0x66fe9f8a.
//
// Solidity: function poolLimitPerUser() view returns(uint256)
func (_Pool *PoolCallerSession) PoolLimitPerUser() (*big.Int, error) {
	return _Pool.Contract.PoolLimitPerUser(&_Pool.CallOpts)
}

// RewardPerBlock is a free data retrieval call binding the contract method 0x8ae39cac.
//
// Solidity: function rewardPerBlock() view returns(uint256)
func (_Pool *PoolCaller) RewardPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "rewardPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerBlock is a free data retrieval call binding the contract method 0x8ae39cac.
//
// Solidity: function rewardPerBlock() view returns(uint256)
func (_Pool *PoolSession) RewardPerBlock() (*big.Int, error) {
	return _Pool.Contract.RewardPerBlock(&_Pool.CallOpts)
}

// RewardPerBlock is a free data retrieval call binding the contract method 0x8ae39cac.
//
// Solidity: function rewardPerBlock() view returns(uint256)
func (_Pool *PoolCallerSession) RewardPerBlock() (*big.Int, error) {
	return _Pool.Contract.RewardPerBlock(&_Pool.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Pool *PoolCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Pool *PoolSession) RewardToken() (common.Address, error) {
	return _Pool.Contract.RewardToken(&_Pool.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Pool *PoolCallerSession) RewardToken() (common.Address, error) {
	return _Pool.Contract.RewardToken(&_Pool.CallOpts)
}

// StakedToken is a free data retrieval call binding the contract method 0xcc7a262e.
//
// Solidity: function stakedToken() view returns(address)
func (_Pool *PoolCaller) StakedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "stakedToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedToken is a free data retrieval call binding the contract method 0xcc7a262e.
//
// Solidity: function stakedToken() view returns(address)
func (_Pool *PoolSession) StakedToken() (common.Address, error) {
	return _Pool.Contract.StakedToken(&_Pool.CallOpts)
}

// StakedToken is a free data retrieval call binding the contract method 0xcc7a262e.
//
// Solidity: function stakedToken() view returns(address)
func (_Pool *PoolCallerSession) StakedToken() (common.Address, error) {
	return _Pool.Contract.StakedToken(&_Pool.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Pool *PoolCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Pool *PoolSession) StartBlock() (*big.Int, error) {
	return _Pool.Contract.StartBlock(&_Pool.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Pool *PoolCallerSession) StartBlock() (*big.Int, error) {
	return _Pool.Contract.StartBlock(&_Pool.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Pool *PoolCaller) UserInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "userInfo", arg0)

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
func (_Pool *PoolSession) UserInfo(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Pool.Contract.UserInfo(&_Pool.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Pool *PoolCallerSession) UserInfo(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Pool.Contract.UserInfo(&_Pool.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Pool *PoolTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Pool *PoolSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Deposit(&_Pool.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Pool *PoolTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Deposit(&_Pool.TransactOpts, _amount)
}

// EmergencyRewardWithdraw is a paid mutator transaction binding the contract method 0x3279beab.
//
// Solidity: function emergencyRewardWithdraw(uint256 _amount) returns()
func (_Pool *PoolTransactor) EmergencyRewardWithdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "emergencyRewardWithdraw", _amount)
}

// EmergencyRewardWithdraw is a paid mutator transaction binding the contract method 0x3279beab.
//
// Solidity: function emergencyRewardWithdraw(uint256 _amount) returns()
func (_Pool *PoolSession) EmergencyRewardWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.EmergencyRewardWithdraw(&_Pool.TransactOpts, _amount)
}

// EmergencyRewardWithdraw is a paid mutator transaction binding the contract method 0x3279beab.
//
// Solidity: function emergencyRewardWithdraw(uint256 _amount) returns()
func (_Pool *PoolTransactorSession) EmergencyRewardWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.EmergencyRewardWithdraw(&_Pool.TransactOpts, _amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_Pool *PoolTransactor) EmergencyWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "emergencyWithdraw")
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_Pool *PoolSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _Pool.Contract.EmergencyWithdraw(&_Pool.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_Pool *PoolTransactorSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _Pool.Contract.EmergencyWithdraw(&_Pool.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x83a5041c.
//
// Solidity: function initialize(address _stakedToken, address _rewardToken, uint256 _rewardPerBlock, uint256 _startBlock, uint256 _bonusEndBlock, uint256 _poolLimitPerUser, address _admin) returns()
func (_Pool *PoolTransactor) Initialize(opts *bind.TransactOpts, _stakedToken common.Address, _rewardToken common.Address, _rewardPerBlock *big.Int, _startBlock *big.Int, _bonusEndBlock *big.Int, _poolLimitPerUser *big.Int, _admin common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "initialize", _stakedToken, _rewardToken, _rewardPerBlock, _startBlock, _bonusEndBlock, _poolLimitPerUser, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x83a5041c.
//
// Solidity: function initialize(address _stakedToken, address _rewardToken, uint256 _rewardPerBlock, uint256 _startBlock, uint256 _bonusEndBlock, uint256 _poolLimitPerUser, address _admin) returns()
func (_Pool *PoolSession) Initialize(_stakedToken common.Address, _rewardToken common.Address, _rewardPerBlock *big.Int, _startBlock *big.Int, _bonusEndBlock *big.Int, _poolLimitPerUser *big.Int, _admin common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Initialize(&_Pool.TransactOpts, _stakedToken, _rewardToken, _rewardPerBlock, _startBlock, _bonusEndBlock, _poolLimitPerUser, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x83a5041c.
//
// Solidity: function initialize(address _stakedToken, address _rewardToken, uint256 _rewardPerBlock, uint256 _startBlock, uint256 _bonusEndBlock, uint256 _poolLimitPerUser, address _admin) returns()
func (_Pool *PoolTransactorSession) Initialize(_stakedToken common.Address, _rewardToken common.Address, _rewardPerBlock *big.Int, _startBlock *big.Int, _bonusEndBlock *big.Int, _poolLimitPerUser *big.Int, _admin common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Initialize(&_Pool.TransactOpts, _stakedToken, _rewardToken, _rewardPerBlock, _startBlock, _bonusEndBlock, _poolLimitPerUser, _admin)
}

// RecoverWrongTokens is a paid mutator transaction binding the contract method 0x3f138d4b.
//
// Solidity: function recoverWrongTokens(address _tokenAddress, uint256 _tokenAmount) returns()
func (_Pool *PoolTransactor) RecoverWrongTokens(opts *bind.TransactOpts, _tokenAddress common.Address, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "recoverWrongTokens", _tokenAddress, _tokenAmount)
}

// RecoverWrongTokens is a paid mutator transaction binding the contract method 0x3f138d4b.
//
// Solidity: function recoverWrongTokens(address _tokenAddress, uint256 _tokenAmount) returns()
func (_Pool *PoolSession) RecoverWrongTokens(_tokenAddress common.Address, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RecoverWrongTokens(&_Pool.TransactOpts, _tokenAddress, _tokenAmount)
}

// RecoverWrongTokens is a paid mutator transaction binding the contract method 0x3f138d4b.
//
// Solidity: function recoverWrongTokens(address _tokenAddress, uint256 _tokenAmount) returns()
func (_Pool *PoolTransactorSession) RecoverWrongTokens(_tokenAddress common.Address, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RecoverWrongTokens(&_Pool.TransactOpts, _tokenAddress, _tokenAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pool *PoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pool *PoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pool.Contract.RenounceOwnership(&_Pool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pool *PoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pool.Contract.RenounceOwnership(&_Pool.TransactOpts)
}

// StopReward is a paid mutator transaction binding the contract method 0x80dc0672.
//
// Solidity: function stopReward() returns()
func (_Pool *PoolTransactor) StopReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "stopReward")
}

// StopReward is a paid mutator transaction binding the contract method 0x80dc0672.
//
// Solidity: function stopReward() returns()
func (_Pool *PoolSession) StopReward() (*types.Transaction, error) {
	return _Pool.Contract.StopReward(&_Pool.TransactOpts)
}

// StopReward is a paid mutator transaction binding the contract method 0x80dc0672.
//
// Solidity: function stopReward() returns()
func (_Pool *PoolTransactorSession) StopReward() (*types.Transaction, error) {
	return _Pool.Contract.StopReward(&_Pool.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pool *PoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pool *PoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pool.Contract.TransferOwnership(&_Pool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pool *PoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pool.Contract.TransferOwnership(&_Pool.TransactOpts, newOwner)
}

// UpdatePoolLimitPerUser is a paid mutator transaction binding the contract method 0xa0b40905.
//
// Solidity: function updatePoolLimitPerUser(bool _hasUserLimit, uint256 _poolLimitPerUser) returns()
func (_Pool *PoolTransactor) UpdatePoolLimitPerUser(opts *bind.TransactOpts, _hasUserLimit bool, _poolLimitPerUser *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "updatePoolLimitPerUser", _hasUserLimit, _poolLimitPerUser)
}

// UpdatePoolLimitPerUser is a paid mutator transaction binding the contract method 0xa0b40905.
//
// Solidity: function updatePoolLimitPerUser(bool _hasUserLimit, uint256 _poolLimitPerUser) returns()
func (_Pool *PoolSession) UpdatePoolLimitPerUser(_hasUserLimit bool, _poolLimitPerUser *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.UpdatePoolLimitPerUser(&_Pool.TransactOpts, _hasUserLimit, _poolLimitPerUser)
}

// UpdatePoolLimitPerUser is a paid mutator transaction binding the contract method 0xa0b40905.
//
// Solidity: function updatePoolLimitPerUser(bool _hasUserLimit, uint256 _poolLimitPerUser) returns()
func (_Pool *PoolTransactorSession) UpdatePoolLimitPerUser(_hasUserLimit bool, _poolLimitPerUser *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.UpdatePoolLimitPerUser(&_Pool.TransactOpts, _hasUserLimit, _poolLimitPerUser)
}

// UpdateRewardPerBlock is a paid mutator transaction binding the contract method 0x01f8a976.
//
// Solidity: function updateRewardPerBlock(uint256 _rewardPerBlock) returns()
func (_Pool *PoolTransactor) UpdateRewardPerBlock(opts *bind.TransactOpts, _rewardPerBlock *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "updateRewardPerBlock", _rewardPerBlock)
}

// UpdateRewardPerBlock is a paid mutator transaction binding the contract method 0x01f8a976.
//
// Solidity: function updateRewardPerBlock(uint256 _rewardPerBlock) returns()
func (_Pool *PoolSession) UpdateRewardPerBlock(_rewardPerBlock *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.UpdateRewardPerBlock(&_Pool.TransactOpts, _rewardPerBlock)
}

// UpdateRewardPerBlock is a paid mutator transaction binding the contract method 0x01f8a976.
//
// Solidity: function updateRewardPerBlock(uint256 _rewardPerBlock) returns()
func (_Pool *PoolTransactorSession) UpdateRewardPerBlock(_rewardPerBlock *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.UpdateRewardPerBlock(&_Pool.TransactOpts, _rewardPerBlock)
}

// UpdateStartAndEndBlocks is a paid mutator transaction binding the contract method 0x9513997f.
//
// Solidity: function updateStartAndEndBlocks(uint256 _startBlock, uint256 _bonusEndBlock) returns()
func (_Pool *PoolTransactor) UpdateStartAndEndBlocks(opts *bind.TransactOpts, _startBlock *big.Int, _bonusEndBlock *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "updateStartAndEndBlocks", _startBlock, _bonusEndBlock)
}

// UpdateStartAndEndBlocks is a paid mutator transaction binding the contract method 0x9513997f.
//
// Solidity: function updateStartAndEndBlocks(uint256 _startBlock, uint256 _bonusEndBlock) returns()
func (_Pool *PoolSession) UpdateStartAndEndBlocks(_startBlock *big.Int, _bonusEndBlock *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.UpdateStartAndEndBlocks(&_Pool.TransactOpts, _startBlock, _bonusEndBlock)
}

// UpdateStartAndEndBlocks is a paid mutator transaction binding the contract method 0x9513997f.
//
// Solidity: function updateStartAndEndBlocks(uint256 _startBlock, uint256 _bonusEndBlock) returns()
func (_Pool *PoolTransactorSession) UpdateStartAndEndBlocks(_startBlock *big.Int, _bonusEndBlock *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.UpdateStartAndEndBlocks(&_Pool.TransactOpts, _startBlock, _bonusEndBlock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Pool *PoolTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Pool *PoolSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Withdraw(&_Pool.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Pool *PoolTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Withdraw(&_Pool.TransactOpts, _amount)
}

// PoolAdminTokenRecoveryIterator is returned from FilterAdminTokenRecovery and is used to iterate over the raw logs and unpacked data for AdminTokenRecovery events raised by the Pool contract.
type PoolAdminTokenRecoveryIterator struct {
	Event *PoolAdminTokenRecovery // Event containing the contract specifics and raw log

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
func (it *PoolAdminTokenRecoveryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAdminTokenRecovery)
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
		it.Event = new(PoolAdminTokenRecovery)
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
func (it *PoolAdminTokenRecoveryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAdminTokenRecoveryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAdminTokenRecovery represents a AdminTokenRecovery event raised by the Pool contract.
type PoolAdminTokenRecovery struct {
	TokenRecovered common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAdminTokenRecovery is a free log retrieval operation binding the contract event 0x74545154aac348a3eac92596bd1971957ca94795f4e954ec5f613b55fab78129.
//
// Solidity: event AdminTokenRecovery(address tokenRecovered, uint256 amount)
func (_Pool *PoolFilterer) FilterAdminTokenRecovery(opts *bind.FilterOpts) (*PoolAdminTokenRecoveryIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "AdminTokenRecovery")
	if err != nil {
		return nil, err
	}
	return &PoolAdminTokenRecoveryIterator{contract: _Pool.contract, event: "AdminTokenRecovery", logs: logs, sub: sub}, nil
}

// WatchAdminTokenRecovery is a free log subscription operation binding the contract event 0x74545154aac348a3eac92596bd1971957ca94795f4e954ec5f613b55fab78129.
//
// Solidity: event AdminTokenRecovery(address tokenRecovered, uint256 amount)
func (_Pool *PoolFilterer) WatchAdminTokenRecovery(opts *bind.WatchOpts, sink chan<- *PoolAdminTokenRecovery) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "AdminTokenRecovery")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAdminTokenRecovery)
				if err := _Pool.contract.UnpackLog(event, "AdminTokenRecovery", log); err != nil {
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

// ParseAdminTokenRecovery is a log parse operation binding the contract event 0x74545154aac348a3eac92596bd1971957ca94795f4e954ec5f613b55fab78129.
//
// Solidity: event AdminTokenRecovery(address tokenRecovered, uint256 amount)
func (_Pool *PoolFilterer) ParseAdminTokenRecovery(log types.Log) (*PoolAdminTokenRecovery, error) {
	event := new(PoolAdminTokenRecovery)
	if err := _Pool.contract.UnpackLog(event, "AdminTokenRecovery", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Pool contract.
type PoolDepositIterator struct {
	Event *PoolDeposit // Event containing the contract specifics and raw log

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
func (it *PoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolDeposit)
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
		it.Event = new(PoolDeposit)
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
func (it *PoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolDeposit represents a Deposit event raised by the Pool contract.
type PoolDeposit struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Pool *PoolFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*PoolDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &PoolDepositIterator{contract: _Pool.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Pool *PoolFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *PoolDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolDeposit)
				if err := _Pool.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_Pool *PoolFilterer) ParseDeposit(log types.Log) (*PoolDeposit, error) {
	event := new(PoolDeposit)
	if err := _Pool.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the Pool contract.
type PoolEmergencyWithdrawIterator struct {
	Event *PoolEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *PoolEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolEmergencyWithdraw)
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
		it.Event = new(PoolEmergencyWithdraw)
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
func (it *PoolEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolEmergencyWithdraw represents a EmergencyWithdraw event raised by the Pool contract.
type PoolEmergencyWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 amount)
func (_Pool *PoolFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address) (*PoolEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "EmergencyWithdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &PoolEmergencyWithdrawIterator{contract: _Pool.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 amount)
func (_Pool *PoolFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *PoolEmergencyWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "EmergencyWithdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolEmergencyWithdraw)
				if err := _Pool.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 amount)
func (_Pool *PoolFilterer) ParseEmergencyWithdraw(log types.Log) (*PoolEmergencyWithdraw, error) {
	event := new(PoolEmergencyWithdraw)
	if err := _Pool.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolNewPoolLimitIterator is returned from FilterNewPoolLimit and is used to iterate over the raw logs and unpacked data for NewPoolLimit events raised by the Pool contract.
type PoolNewPoolLimitIterator struct {
	Event *PoolNewPoolLimit // Event containing the contract specifics and raw log

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
func (it *PoolNewPoolLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolNewPoolLimit)
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
		it.Event = new(PoolNewPoolLimit)
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
func (it *PoolNewPoolLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolNewPoolLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolNewPoolLimit represents a NewPoolLimit event raised by the Pool contract.
type PoolNewPoolLimit struct {
	PoolLimitPerUser *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewPoolLimit is a free log retrieval operation binding the contract event 0x241f67ee5f41b7a5cabf911367329be7215900f602ebfc47f89dce2a6bcd847c.
//
// Solidity: event NewPoolLimit(uint256 poolLimitPerUser)
func (_Pool *PoolFilterer) FilterNewPoolLimit(opts *bind.FilterOpts) (*PoolNewPoolLimitIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "NewPoolLimit")
	if err != nil {
		return nil, err
	}
	return &PoolNewPoolLimitIterator{contract: _Pool.contract, event: "NewPoolLimit", logs: logs, sub: sub}, nil
}

// WatchNewPoolLimit is a free log subscription operation binding the contract event 0x241f67ee5f41b7a5cabf911367329be7215900f602ebfc47f89dce2a6bcd847c.
//
// Solidity: event NewPoolLimit(uint256 poolLimitPerUser)
func (_Pool *PoolFilterer) WatchNewPoolLimit(opts *bind.WatchOpts, sink chan<- *PoolNewPoolLimit) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "NewPoolLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolNewPoolLimit)
				if err := _Pool.contract.UnpackLog(event, "NewPoolLimit", log); err != nil {
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

// ParseNewPoolLimit is a log parse operation binding the contract event 0x241f67ee5f41b7a5cabf911367329be7215900f602ebfc47f89dce2a6bcd847c.
//
// Solidity: event NewPoolLimit(uint256 poolLimitPerUser)
func (_Pool *PoolFilterer) ParseNewPoolLimit(log types.Log) (*PoolNewPoolLimit, error) {
	event := new(PoolNewPoolLimit)
	if err := _Pool.contract.UnpackLog(event, "NewPoolLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolNewRewardPerBlockIterator is returned from FilterNewRewardPerBlock and is used to iterate over the raw logs and unpacked data for NewRewardPerBlock events raised by the Pool contract.
type PoolNewRewardPerBlockIterator struct {
	Event *PoolNewRewardPerBlock // Event containing the contract specifics and raw log

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
func (it *PoolNewRewardPerBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolNewRewardPerBlock)
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
		it.Event = new(PoolNewRewardPerBlock)
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
func (it *PoolNewRewardPerBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolNewRewardPerBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolNewRewardPerBlock represents a NewRewardPerBlock event raised by the Pool contract.
type PoolNewRewardPerBlock struct {
	RewardPerBlock *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewRewardPerBlock is a free log retrieval operation binding the contract event 0x0c4d677eef92893ac7ec52faf8140fc6c851ab4736302b4f3a89dfb20696a0df.
//
// Solidity: event NewRewardPerBlock(uint256 rewardPerBlock)
func (_Pool *PoolFilterer) FilterNewRewardPerBlock(opts *bind.FilterOpts) (*PoolNewRewardPerBlockIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "NewRewardPerBlock")
	if err != nil {
		return nil, err
	}
	return &PoolNewRewardPerBlockIterator{contract: _Pool.contract, event: "NewRewardPerBlock", logs: logs, sub: sub}, nil
}

// WatchNewRewardPerBlock is a free log subscription operation binding the contract event 0x0c4d677eef92893ac7ec52faf8140fc6c851ab4736302b4f3a89dfb20696a0df.
//
// Solidity: event NewRewardPerBlock(uint256 rewardPerBlock)
func (_Pool *PoolFilterer) WatchNewRewardPerBlock(opts *bind.WatchOpts, sink chan<- *PoolNewRewardPerBlock) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "NewRewardPerBlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolNewRewardPerBlock)
				if err := _Pool.contract.UnpackLog(event, "NewRewardPerBlock", log); err != nil {
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

// ParseNewRewardPerBlock is a log parse operation binding the contract event 0x0c4d677eef92893ac7ec52faf8140fc6c851ab4736302b4f3a89dfb20696a0df.
//
// Solidity: event NewRewardPerBlock(uint256 rewardPerBlock)
func (_Pool *PoolFilterer) ParseNewRewardPerBlock(log types.Log) (*PoolNewRewardPerBlock, error) {
	event := new(PoolNewRewardPerBlock)
	if err := _Pool.contract.UnpackLog(event, "NewRewardPerBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolNewStartAndEndBlocksIterator is returned from FilterNewStartAndEndBlocks and is used to iterate over the raw logs and unpacked data for NewStartAndEndBlocks events raised by the Pool contract.
type PoolNewStartAndEndBlocksIterator struct {
	Event *PoolNewStartAndEndBlocks // Event containing the contract specifics and raw log

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
func (it *PoolNewStartAndEndBlocksIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolNewStartAndEndBlocks)
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
		it.Event = new(PoolNewStartAndEndBlocks)
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
func (it *PoolNewStartAndEndBlocksIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolNewStartAndEndBlocksIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolNewStartAndEndBlocks represents a NewStartAndEndBlocks event raised by the Pool contract.
type PoolNewStartAndEndBlocks struct {
	StartBlock *big.Int
	EndBlock   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewStartAndEndBlocks is a free log retrieval operation binding the contract event 0x7cd0ab87d19036f3dfadadb232c78aa4879dda3f0c994a9d637532410ee2ce06.
//
// Solidity: event NewStartAndEndBlocks(uint256 startBlock, uint256 endBlock)
func (_Pool *PoolFilterer) FilterNewStartAndEndBlocks(opts *bind.FilterOpts) (*PoolNewStartAndEndBlocksIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "NewStartAndEndBlocks")
	if err != nil {
		return nil, err
	}
	return &PoolNewStartAndEndBlocksIterator{contract: _Pool.contract, event: "NewStartAndEndBlocks", logs: logs, sub: sub}, nil
}

// WatchNewStartAndEndBlocks is a free log subscription operation binding the contract event 0x7cd0ab87d19036f3dfadadb232c78aa4879dda3f0c994a9d637532410ee2ce06.
//
// Solidity: event NewStartAndEndBlocks(uint256 startBlock, uint256 endBlock)
func (_Pool *PoolFilterer) WatchNewStartAndEndBlocks(opts *bind.WatchOpts, sink chan<- *PoolNewStartAndEndBlocks) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "NewStartAndEndBlocks")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolNewStartAndEndBlocks)
				if err := _Pool.contract.UnpackLog(event, "NewStartAndEndBlocks", log); err != nil {
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

// ParseNewStartAndEndBlocks is a log parse operation binding the contract event 0x7cd0ab87d19036f3dfadadb232c78aa4879dda3f0c994a9d637532410ee2ce06.
//
// Solidity: event NewStartAndEndBlocks(uint256 startBlock, uint256 endBlock)
func (_Pool *PoolFilterer) ParseNewStartAndEndBlocks(log types.Log) (*PoolNewStartAndEndBlocks, error) {
	event := new(PoolNewStartAndEndBlocks)
	if err := _Pool.contract.UnpackLog(event, "NewStartAndEndBlocks", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pool contract.
type PoolOwnershipTransferredIterator struct {
	Event *PoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolOwnershipTransferred)
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
		it.Event = new(PoolOwnershipTransferred)
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
func (it *PoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolOwnershipTransferred represents a OwnershipTransferred event raised by the Pool contract.
type PoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pool *PoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PoolOwnershipTransferredIterator{contract: _Pool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pool *PoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolOwnershipTransferred)
				if err := _Pool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Pool *PoolFilterer) ParseOwnershipTransferred(log types.Log) (*PoolOwnershipTransferred, error) {
	event := new(PoolOwnershipTransferred)
	if err := _Pool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolRewardsStopIterator is returned from FilterRewardsStop and is used to iterate over the raw logs and unpacked data for RewardsStop events raised by the Pool contract.
type PoolRewardsStopIterator struct {
	Event *PoolRewardsStop // Event containing the contract specifics and raw log

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
func (it *PoolRewardsStopIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolRewardsStop)
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
		it.Event = new(PoolRewardsStop)
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
func (it *PoolRewardsStopIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolRewardsStopIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolRewardsStop represents a RewardsStop event raised by the Pool contract.
type PoolRewardsStop struct {
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardsStop is a free log retrieval operation binding the contract event 0xfed9fcb0ca3d1e761a4b929792bb24082fba92dca81252646ad306d306806566.
//
// Solidity: event RewardsStop(uint256 blockNumber)
func (_Pool *PoolFilterer) FilterRewardsStop(opts *bind.FilterOpts) (*PoolRewardsStopIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "RewardsStop")
	if err != nil {
		return nil, err
	}
	return &PoolRewardsStopIterator{contract: _Pool.contract, event: "RewardsStop", logs: logs, sub: sub}, nil
}

// WatchRewardsStop is a free log subscription operation binding the contract event 0xfed9fcb0ca3d1e761a4b929792bb24082fba92dca81252646ad306d306806566.
//
// Solidity: event RewardsStop(uint256 blockNumber)
func (_Pool *PoolFilterer) WatchRewardsStop(opts *bind.WatchOpts, sink chan<- *PoolRewardsStop) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "RewardsStop")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolRewardsStop)
				if err := _Pool.contract.UnpackLog(event, "RewardsStop", log); err != nil {
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

// ParseRewardsStop is a log parse operation binding the contract event 0xfed9fcb0ca3d1e761a4b929792bb24082fba92dca81252646ad306d306806566.
//
// Solidity: event RewardsStop(uint256 blockNumber)
func (_Pool *PoolFilterer) ParseRewardsStop(log types.Log) (*PoolRewardsStop, error) {
	event := new(PoolRewardsStop)
	if err := _Pool.contract.UnpackLog(event, "RewardsStop", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Pool contract.
type PoolWithdrawIterator struct {
	Event *PoolWithdraw // Event containing the contract specifics and raw log

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
func (it *PoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolWithdraw)
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
		it.Event = new(PoolWithdraw)
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
func (it *PoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolWithdraw represents a Withdraw event raised by the Pool contract.
type PoolWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Pool *PoolFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*PoolWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &PoolWithdrawIterator{contract: _Pool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Pool *PoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PoolWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolWithdraw)
				if err := _Pool.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_Pool *PoolFilterer) ParseWithdraw(log types.Log) (*PoolWithdraw, error) {
	event := new(PoolWithdraw)
	if err := _Pool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

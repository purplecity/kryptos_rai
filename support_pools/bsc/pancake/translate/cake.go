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

// CakeABI is the input ABI used to generate the binding from.
const CakeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromDelegate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toDelegate\",\"type\":\"address\"}],\"name\":\"DelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"DelegateVotesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DELEGATION_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"checkpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"fromBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getCurrentVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getPriorVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"numCheckpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CakeBin is the compiled bytecode used for deploying new contracts.
var CakeBin = "0x60806040523480156200001157600080fd5b50604051806040016040528060118152602001702830b731b0b5b2a9bbb0b8102a37b5b2b760791b8152506040518060400160405280600481526020016343616b6560e01b81525060006200006b620000f660201b60201c565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3508151620000ca906004906020850190620000fa565b508051620000e0906005906020840190620000fa565b50506006805460ff191660121790555062000196565b3390565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200013d57805160ff19168380011785556200016d565b828001600101855582156200016d579182015b828111156200016d57825182559160200191906001019062000150565b506200017b9291506200017f565b5090565b5b808211156200017b576000815560010162000180565b611c7580620001a66000396000f3fe608060405234801561001057600080fd5b50600436106101b95760003560e01c8063782d6fe1116100f9578063a9059cbb11610097578063dd62ed3e11610071578063dd62ed3e1461056c578063e7a324dc1461059a578063f1127ed8146105a2578063f2fde38b146105f4576101b9565b8063a9059cbb146104d3578063b4b5ea57146104ff578063c3cda52014610525576101b9565b80638da5cb5b116100d35780638da5cb5b1461047a57806395d89b4114610482578063a0712d681461048a578063a457c2d7146104a7576101b9565b8063782d6fe1146104205780637ecebe001461044c578063893d20e814610472576101b9565b806339509351116101665780635c19a95c116101405780635c19a95c1461038d5780636fcfff45146103b357806370a08231146103f2578063715018a614610418576101b9565b806339509351146102f157806340c10f191461031d578063587cde1e1461034b576101b9565b806320606b701161019757806320606b701461029557806323b872dd1461029d578063313ce567146102d3576101b9565b806306fdde03146101be578063095ea7b31461023b57806318160ddd1461027b575b600080fd5b6101c661061a565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102005781810151838201526020016101e8565b50505050905090810190601f16801561022d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102676004803603604081101561025157600080fd5b506001600160a01b0381351690602001356106b0565b604080519115158252519081900360200190f35b6102836106ce565b60408051918252519081900360200190f35b6102836106d4565b610267600480360360608110156102b357600080fd5b506001600160a01b038135811691602081013590911690604001356106f8565b6102db61077f565b6040805160ff9092168252519081900360200190f35b6102676004803603604081101561030757600080fd5b506001600160a01b038135169060200135610788565b6103496004803603604081101561033357600080fd5b506001600160a01b0381351690602001356107d6565b005b6103716004803603602081101561036157600080fd5b50356001600160a01b0316610873565b604080516001600160a01b039092168252519081900360200190f35b610349600480360360208110156103a357600080fd5b50356001600160a01b0316610891565b6103d9600480360360208110156103c957600080fd5b50356001600160a01b031661089e565b6040805163ffffffff9092168252519081900360200190f35b6102836004803603602081101561040857600080fd5b50356001600160a01b03166108b6565b6103496108d1565b6102836004803603604081101561043657600080fd5b506001600160a01b03813516906020013561099d565b6102836004803603602081101561046257600080fd5b50356001600160a01b0316610ba5565b610371610bb7565b610371610bc6565b6101c6610bd5565b610267600480360360208110156104a057600080fd5b5035610c36565b610267600480360360408110156104bd57600080fd5b506001600160a01b038135169060200135610cbb565b610267600480360360408110156104e957600080fd5b506001600160a01b038135169060200135610d23565b6102836004803603602081101561051557600080fd5b50356001600160a01b0316610d37565b610349600480360360c081101561053b57600080fd5b506001600160a01b038135169060208101359060408101359060ff6060820135169060808101359060a00135610d9b565b6102836004803603604081101561058257600080fd5b506001600160a01b0381358116916020013516611065565b610283611090565b6105d4600480360360408110156105b857600080fd5b5080356001600160a01b0316906020013563ffffffff166110b4565b6040805163ffffffff909316835260208301919091528051918290030190f35b6103496004803603602081101561060a57600080fd5b50356001600160a01b03166110e1565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156106a65780601f1061067b576101008083540402835291602001916106a6565b820191906000526020600020905b81548152906001019060200180831161068957829003601f168201915b5050505050905090565b60006106c46106bd611154565b8484611158565b5060015b92915050565b60035490565b7f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a86681565b6000610705848484611244565b61077584610711611154565b61077085604051806060016040528060288152602001611ae5602891396001600160a01b038a1660009081526002602052604081209061074f611154565b6001600160a01b031681526020810191909152604001600020549190611396565b611158565b5060019392505050565b60065460ff1690565b60006106c4610795611154565b8461077085600260006107a6611154565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549061142d565b6107de611154565b6000546001600160a01b03908116911614610840576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b61084a8282611487565b6001600160a01b0380831660009081526007602052604081205461086f92168361156d565b5050565b6001600160a01b039081166000908152600760205260409020541690565b61089b33826116af565b50565b60096020526000908152604090205463ffffffff1681565b6001600160a01b031660009081526001602052604090205490565b6108d9611154565b6000546001600160a01b0390811691161461093b576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055565b60004382106109dd5760405162461bcd60e51b8152600401808060200182810382526027815260200180611b676027913960400191505060405180910390fd5b6001600160a01b03831660009081526009602052604090205463ffffffff1680610a0b5760009150506106c8565b6001600160a01b038416600090815260086020908152604080832063ffffffff600019860181168552925290912054168310610a7a576001600160a01b03841660009081526008602090815260408083206000199490940163ffffffff168352929052206001015490506106c8565b6001600160a01b038416600090815260086020908152604080832083805290915290205463ffffffff16831015610ab55760009150506106c8565b600060001982015b8163ffffffff168163ffffffff161115610b6e57600282820363ffffffff16048103610ae7611a38565b506001600160a01b038716600090815260086020908152604080832063ffffffff808616855290835292819020815180830190925280549093168082526001909301549181019190915290871415610b49576020015194506106c89350505050565b805163ffffffff16871115610b6057819350610b67565b6001820392505b5050610abd565b506001600160a01b038516600090815260086020908152604080832063ffffffff9094168352929052206001015491505092915050565b600a6020526000908152604090205481565b6000610bc1610bc6565b905090565b6000546001600160a01b031690565b60058054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156106a65780601f1061067b576101008083540402835291602001916106a6565b6000610c40611154565b6000546001600160a01b03908116911614610ca2576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b610cb3610cad611154565b83611487565b506001919050565b60006106c4610cc8611154565b8461077085604051806060016040528060258152602001611bf96025913960026000610cf2611154565b6001600160a01b03908116825260208083019390935260409182016000908120918d16815292529020549190611396565b60006106c4610d30611154565b8484611244565b6001600160a01b03811660009081526009602052604081205463ffffffff1680610d62576000610d94565b6001600160a01b038316600090815260086020908152604080832063ffffffff60001986011684529091529020600101545b9392505050565b60007f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866610dc661061a565b80519060200120610dd561175c565b60408051602080820195909552808201939093526060830191909152306080808401919091528151808403909101815260a0830182528051908401207fe48329057bfd03d55e49b547132e39cffd9c1820ad7b9d4c5307691425d15adf60c08401526001600160a01b038b1660e084015261010083018a90526101208084018a905282518085039091018152610140840183528051908501207f19010000000000000000000000000000000000000000000000000000000000006101608501526101628401829052610182808501829052835180860390910181526101a285018085528151918701919091206000918290526101c2860180865281905260ff8b166101e287015261020286018a905261022286018990529351929650909492939092600192610242808401937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08301929081900390910190855afa158015610f41573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001519150506001600160a01b038116610fb15760405162461bcd60e51b8152600401808060200182810382526026815260200180611b0d6026913960400191505060405180910390fd5b6001600160a01b0381166000908152600a60205260409020805460018101909155891461100f5760405162461bcd60e51b8152600401808060200182810382526022815260200180611bb46022913960400191505060405180910390fd5b8742111561104e5760405162461bcd60e51b8152600401808060200182810382526026815260200180611abf6026913960400191505060405180910390fd5b611058818b6116af565b505050505b505050505050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b7fe48329057bfd03d55e49b547132e39cffd9c1820ad7b9d4c5307691425d15adf81565b60086020908152600092835260408084209091529082529020805460019091015463ffffffff9091169082565b6110e9611154565b6000546001600160a01b0390811691161461114b576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b61089b81611760565b3390565b6001600160a01b03831661119d5760405162461bcd60e51b8152600401808060200182810382526024815260200180611a756024913960400191505060405180910390fd5b6001600160a01b0382166111e25760405162461bcd60e51b8152600401808060200182810382526022815260200180611c1e6022913960400191505060405180910390fd5b6001600160a01b03808416600081815260026020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166112895760405162461bcd60e51b8152600401808060200182810382526025815260200180611a506025913960400191505060405180910390fd5b6001600160a01b0382166112ce5760405162461bcd60e51b8152600401808060200182810382526023815260200180611bd66023913960400191505060405180910390fd5b61130b81604051806060016040528060268152602001611b8e602691396001600160a01b0386166000908152600160205260409020549190611396565b6001600160a01b03808516600090815260016020526040808220939093559084168152205461133a908261142d565b6001600160a01b0380841660008181526001602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600081848411156114255760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156113ea5781810151838201526020016113d2565b50505050905090810190601f1680156114175780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600082820183811015610d94576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6001600160a01b0382166114e2576040805162461bcd60e51b815260206004820152601f60248201527f42455032303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b6003546114ef908261142d565b6003556001600160a01b038216600090815260016020526040902054611515908261142d565b6001600160a01b03831660008181526001602090815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b816001600160a01b0316836001600160a01b03161415801561158f5750600081115b156116aa576001600160a01b03831615611621576001600160a01b03831660009081526009602052604081205463ffffffff1690816115cf576000611601565b6001600160a01b038516600090815260086020908152604080832063ffffffff60001987011684529091529020600101545b9050600061160f8285611818565b905061161d8684848461185a565b5050505b6001600160a01b038216156116aa576001600160a01b03821660009081526009602052604081205463ffffffff16908161165c57600061168e565b6001600160a01b038416600090815260086020908152604080832063ffffffff60001987011684529091529020600101545b9050600061169c828561142d565b905061105d8584848461185a565b505050565b6001600160a01b03808316600090815260076020526040812054909116906116d6846108b6565b6001600160a01b0385811660008181526007602052604080822080547fffffffffffffffffffffffff000000000000000000000000000000000000000016898616908117909155905194955093928616927f3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f9190a461175682848361156d565b50505050565b4690565b6001600160a01b0381166117a55760405162461bcd60e51b8152600401808060200182810382526026815260200180611a996026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6000610d9483836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611396565b600061187e43604051806060016040528060348152602001611b33603491396119da565b905060008463ffffffff161180156118c757506001600160a01b038516600090815260086020908152604080832063ffffffff6000198901811685529252909120548282169116145b15611904576001600160a01b038516600090815260086020908152604080832063ffffffff60001989011684529091529020600101829055611990565b60408051808201825263ffffffff808416825260208083018681526001600160a01b038a166000818152600884528681208b861682528452868120955186549086167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000918216178755925160019687015590815260099092529390208054928801909116919092161790555b604080518481526020810184905281516001600160a01b038816927fdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724928290030190a25050505050565b6000816401000000008410611a305760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156113ea5781810151838201526020016113d2565b509192915050565b60408051808201909152600080825260208201529056fe42455032303a207472616e736665722066726f6d20746865207a65726f206164647265737342455032303a20617070726f76652066726f6d20746865207a65726f20616464726573734f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737343414b453a3a64656c656761746542795369673a207369676e6174757265206578706972656442455032303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636543414b453a3a64656c656761746542795369673a20696e76616c6964207369676e617475726543414b453a3a5f7772697465436865636b706f696e743a20626c6f636b206e756d6265722065786365656473203332206269747343414b453a3a6765745072696f72566f7465733a206e6f74207965742064657465726d696e656442455032303a207472616e7366657220616d6f756e7420657863656564732062616c616e636543414b453a3a64656c656761746542795369673a20696e76616c6964206e6f6e636542455032303a207472616e7366657220746f20746865207a65726f206164647265737342455032303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726f42455032303a20617070726f766520746f20746865207a65726f2061646472657373a26469706673582212209755240809e31aec9fa5790314682233ca431b7c4f252d7e4bba347e2e74208664736f6c634300060c0033"

// DeployCake deploys a new Ethereum contract, binding an instance of Cake to it.
func DeployCake(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Cake, error) {
	parsed, err := abi.JSON(strings.NewReader(CakeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CakeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cake{CakeCaller: CakeCaller{contract: contract}, CakeTransactor: CakeTransactor{contract: contract}, CakeFilterer: CakeFilterer{contract: contract}}, nil
}

// Cake is an auto generated Go binding around an Ethereum contract.
type Cake struct {
	CakeCaller     // Read-only binding to the contract
	CakeTransactor // Write-only binding to the contract
	CakeFilterer   // Log filterer for contract events
}

// CakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CakeSession struct {
	Contract     *Cake             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CakeCallerSession struct {
	Contract *CakeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CakeTransactorSession struct {
	Contract     *CakeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CakeRaw struct {
	Contract *Cake // Generic contract binding to access the raw methods on
}

// CakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CakeCallerRaw struct {
	Contract *CakeCaller // Generic read-only contract binding to access the raw methods on
}

// CakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CakeTransactorRaw struct {
	Contract *CakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCake creates a new instance of Cake, bound to a specific deployed contract.
func NewCake(address common.Address, backend bind.ContractBackend) (*Cake, error) {
	contract, err := bindCake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cake{CakeCaller: CakeCaller{contract: contract}, CakeTransactor: CakeTransactor{contract: contract}, CakeFilterer: CakeFilterer{contract: contract}}, nil
}

// NewCakeCaller creates a new read-only instance of Cake, bound to a specific deployed contract.
func NewCakeCaller(address common.Address, caller bind.ContractCaller) (*CakeCaller, error) {
	contract, err := bindCake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CakeCaller{contract: contract}, nil
}

// NewCakeTransactor creates a new write-only instance of Cake, bound to a specific deployed contract.
func NewCakeTransactor(address common.Address, transactor bind.ContractTransactor) (*CakeTransactor, error) {
	contract, err := bindCake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CakeTransactor{contract: contract}, nil
}

// NewCakeFilterer creates a new log filterer instance of Cake, bound to a specific deployed contract.
func NewCakeFilterer(address common.Address, filterer bind.ContractFilterer) (*CakeFilterer, error) {
	contract, err := bindCake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CakeFilterer{contract: contract}, nil
}

// bindCake binds a generic wrapper to an already deployed contract.
func bindCake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CakeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cake *CakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cake.Contract.CakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cake *CakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cake.Contract.CakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cake *CakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cake.Contract.CakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cake *CakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cake *CakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cake *CakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cake.Contract.contract.Transact(opts, method, params...)
}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_Cake *CakeCaller) DELEGATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "DELEGATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_Cake *CakeSession) DELEGATIONTYPEHASH() ([32]byte, error) {
	return _Cake.Contract.DELEGATIONTYPEHASH(&_Cake.CallOpts)
}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_Cake *CakeCallerSession) DELEGATIONTYPEHASH() ([32]byte, error) {
	return _Cake.Contract.DELEGATIONTYPEHASH(&_Cake.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Cake *CakeCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Cake *CakeSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Cake.Contract.DOMAINTYPEHASH(&_Cake.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Cake *CakeCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Cake.Contract.DOMAINTYPEHASH(&_Cake.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Cake *CakeCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Cake *CakeSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Cake.Contract.Allowance(&_Cake.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Cake *CakeCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Cake.Contract.Allowance(&_Cake.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Cake *CakeCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Cake *CakeSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Cake.Contract.BalanceOf(&_Cake.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Cake *CakeCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Cake.Contract.BalanceOf(&_Cake.CallOpts, account)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint256 votes)
func (_Cake *CakeCaller) Checkpoints(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "checkpoints", arg0, arg1)

	outstruct := new(struct {
		FromBlock uint32
		Votes     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FromBlock = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Votes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint256 votes)
func (_Cake *CakeSession) Checkpoints(arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	return _Cake.Contract.Checkpoints(&_Cake.CallOpts, arg0, arg1)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint256 votes)
func (_Cake *CakeCallerSession) Checkpoints(arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	return _Cake.Contract.Checkpoints(&_Cake.CallOpts, arg0, arg1)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Cake *CakeCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Cake *CakeSession) Decimals() (uint8, error) {
	return _Cake.Contract.Decimals(&_Cake.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Cake *CakeCallerSession) Decimals() (uint8, error) {
	return _Cake.Contract.Decimals(&_Cake.CallOpts)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address delegator) view returns(address)
func (_Cake *CakeCaller) Delegates(opts *bind.CallOpts, delegator common.Address) (common.Address, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "delegates", delegator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address delegator) view returns(address)
func (_Cake *CakeSession) Delegates(delegator common.Address) (common.Address, error) {
	return _Cake.Contract.Delegates(&_Cake.CallOpts, delegator)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address delegator) view returns(address)
func (_Cake *CakeCallerSession) Delegates(delegator common.Address) (common.Address, error) {
	return _Cake.Contract.Delegates(&_Cake.CallOpts, delegator)
}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint256)
func (_Cake *CakeCaller) GetCurrentVotes(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "getCurrentVotes", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint256)
func (_Cake *CakeSession) GetCurrentVotes(account common.Address) (*big.Int, error) {
	return _Cake.Contract.GetCurrentVotes(&_Cake.CallOpts, account)
}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint256)
func (_Cake *CakeCallerSession) GetCurrentVotes(account common.Address) (*big.Int, error) {
	return _Cake.Contract.GetCurrentVotes(&_Cake.CallOpts, account)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Cake *CakeCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Cake *CakeSession) GetOwner() (common.Address, error) {
	return _Cake.Contract.GetOwner(&_Cake.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Cake *CakeCallerSession) GetOwner() (common.Address, error) {
	return _Cake.Contract.GetOwner(&_Cake.CallOpts)
}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint256)
func (_Cake *CakeCaller) GetPriorVotes(opts *bind.CallOpts, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "getPriorVotes", account, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint256)
func (_Cake *CakeSession) GetPriorVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _Cake.Contract.GetPriorVotes(&_Cake.CallOpts, account, blockNumber)
}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint256)
func (_Cake *CakeCallerSession) GetPriorVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _Cake.Contract.GetPriorVotes(&_Cake.CallOpts, account, blockNumber)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cake *CakeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cake *CakeSession) Name() (string, error) {
	return _Cake.Contract.Name(&_Cake.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cake *CakeCallerSession) Name() (string, error) {
	return _Cake.Contract.Name(&_Cake.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Cake *CakeCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Cake *CakeSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Cake.Contract.Nonces(&_Cake.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Cake *CakeCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Cake.Contract.Nonces(&_Cake.CallOpts, arg0)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_Cake *CakeCaller) NumCheckpoints(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "numCheckpoints", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_Cake *CakeSession) NumCheckpoints(arg0 common.Address) (uint32, error) {
	return _Cake.Contract.NumCheckpoints(&_Cake.CallOpts, arg0)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_Cake *CakeCallerSession) NumCheckpoints(arg0 common.Address) (uint32, error) {
	return _Cake.Contract.NumCheckpoints(&_Cake.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cake *CakeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cake *CakeSession) Owner() (common.Address, error) {
	return _Cake.Contract.Owner(&_Cake.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cake *CakeCallerSession) Owner() (common.Address, error) {
	return _Cake.Contract.Owner(&_Cake.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cake *CakeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cake *CakeSession) Symbol() (string, error) {
	return _Cake.Contract.Symbol(&_Cake.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cake *CakeCallerSession) Symbol() (string, error) {
	return _Cake.Contract.Symbol(&_Cake.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cake *CakeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cake *CakeSession) TotalSupply() (*big.Int, error) {
	return _Cake.Contract.TotalSupply(&_Cake.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cake *CakeCallerSession) TotalSupply() (*big.Int, error) {
	return _Cake.Contract.TotalSupply(&_Cake.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Cake *CakeTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cake.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Cake *CakeSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.Approve(&_Cake.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Cake *CakeTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.Approve(&_Cake.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Cake *CakeTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Cake.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Cake *CakeSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.DecreaseAllowance(&_Cake.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Cake *CakeTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.DecreaseAllowance(&_Cake.TransactOpts, spender, subtractedValue)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Cake *CakeTransactor) Delegate(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _Cake.contract.Transact(opts, "delegate", delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Cake *CakeSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _Cake.Contract.Delegate(&_Cake.TransactOpts, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Cake *CakeTransactorSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _Cake.Contract.Delegate(&_Cake.TransactOpts, delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Cake *CakeTransactor) DelegateBySig(opts *bind.TransactOpts, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Cake.contract.Transact(opts, "delegateBySig", delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Cake *CakeSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Cake.Contract.DelegateBySig(&_Cake.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Cake *CakeTransactorSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Cake.Contract.DelegateBySig(&_Cake.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Cake *CakeTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Cake.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Cake *CakeSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.IncreaseAllowance(&_Cake.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Cake *CakeTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.IncreaseAllowance(&_Cake.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Cake *CakeTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Cake.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Cake *CakeSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.Mint(&_Cake.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Cake *CakeTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.Mint(&_Cake.TransactOpts, _to, _amount)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_Cake *CakeTransactor) Mint0(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Cake.contract.Transact(opts, "mint0", amount)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_Cake *CakeSession) Mint0(amount *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.Mint0(&_Cake.TransactOpts, amount)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_Cake *CakeTransactorSession) Mint0(amount *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.Mint0(&_Cake.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cake *CakeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cake.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cake *CakeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cake.Contract.RenounceOwnership(&_Cake.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cake *CakeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cake.Contract.RenounceOwnership(&_Cake.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Cake *CakeTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cake.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Cake *CakeSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.Transfer(&_Cake.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Cake *CakeTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.Transfer(&_Cake.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Cake *CakeTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cake.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Cake *CakeSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.TransferFrom(&_Cake.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Cake *CakeTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.TransferFrom(&_Cake.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cake *CakeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Cake.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cake *CakeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cake.Contract.TransferOwnership(&_Cake.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cake *CakeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cake.Contract.TransferOwnership(&_Cake.TransactOpts, newOwner)
}

// CakeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Cake contract.
type CakeApprovalIterator struct {
	Event *CakeApproval // Event containing the contract specifics and raw log

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
func (it *CakeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakeApproval)
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
		it.Event = new(CakeApproval)
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
func (it *CakeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakeApproval represents a Approval event raised by the Cake contract.
type CakeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Cake *CakeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CakeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Cake.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CakeApprovalIterator{contract: _Cake.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Cake *CakeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CakeApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Cake.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakeApproval)
				if err := _Cake.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Cake *CakeFilterer) ParseApproval(log types.Log) (*CakeApproval, error) {
	event := new(CakeApproval)
	if err := _Cake.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakeDelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the Cake contract.
type CakeDelegateChangedIterator struct {
	Event *CakeDelegateChanged // Event containing the contract specifics and raw log

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
func (it *CakeDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakeDelegateChanged)
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
		it.Event = new(CakeDelegateChanged)
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
func (it *CakeDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakeDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakeDelegateChanged represents a DelegateChanged event raised by the Cake contract.
type CakeDelegateChanged struct {
	Delegator    common.Address
	FromDelegate common.Address
	ToDelegate   common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_Cake *CakeFilterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (*CakeDelegateChangedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _Cake.contract.FilterLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return &CakeDelegateChangedIterator{contract: _Cake.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_Cake *CakeFilterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *CakeDelegateChanged, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _Cake.contract.WatchLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakeDelegateChanged)
				if err := _Cake.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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

// ParseDelegateChanged is a log parse operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_Cake *CakeFilterer) ParseDelegateChanged(log types.Log) (*CakeDelegateChanged, error) {
	event := new(CakeDelegateChanged)
	if err := _Cake.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakeDelegateVotesChangedIterator is returned from FilterDelegateVotesChanged and is used to iterate over the raw logs and unpacked data for DelegateVotesChanged events raised by the Cake contract.
type CakeDelegateVotesChangedIterator struct {
	Event *CakeDelegateVotesChanged // Event containing the contract specifics and raw log

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
func (it *CakeDelegateVotesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakeDelegateVotesChanged)
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
		it.Event = new(CakeDelegateVotesChanged)
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
func (it *CakeDelegateVotesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakeDelegateVotesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakeDelegateVotesChanged represents a DelegateVotesChanged event raised by the Cake contract.
type CakeDelegateVotesChanged struct {
	Delegate        common.Address
	PreviousBalance *big.Int
	NewBalance      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegateVotesChanged is a free log retrieval operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_Cake *CakeFilterer) FilterDelegateVotesChanged(opts *bind.FilterOpts, delegate []common.Address) (*CakeDelegateVotesChangedIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _Cake.contract.FilterLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return &CakeDelegateVotesChangedIterator{contract: _Cake.contract, event: "DelegateVotesChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateVotesChanged is a free log subscription operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_Cake *CakeFilterer) WatchDelegateVotesChanged(opts *bind.WatchOpts, sink chan<- *CakeDelegateVotesChanged, delegate []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _Cake.contract.WatchLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakeDelegateVotesChanged)
				if err := _Cake.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
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

// ParseDelegateVotesChanged is a log parse operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_Cake *CakeFilterer) ParseDelegateVotesChanged(log types.Log) (*CakeDelegateVotesChanged, error) {
	event := new(CakeDelegateVotesChanged)
	if err := _Cake.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Cake contract.
type CakeOwnershipTransferredIterator struct {
	Event *CakeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CakeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakeOwnershipTransferred)
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
		it.Event = new(CakeOwnershipTransferred)
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
func (it *CakeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakeOwnershipTransferred represents a OwnershipTransferred event raised by the Cake contract.
type CakeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cake *CakeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CakeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cake.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CakeOwnershipTransferredIterator{contract: _Cake.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cake *CakeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CakeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cake.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakeOwnershipTransferred)
				if err := _Cake.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Cake *CakeFilterer) ParseOwnershipTransferred(log types.Log) (*CakeOwnershipTransferred, error) {
	event := new(CakeOwnershipTransferred)
	if err := _Cake.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Cake contract.
type CakeTransferIterator struct {
	Event *CakeTransfer // Event containing the contract specifics and raw log

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
func (it *CakeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakeTransfer)
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
		it.Event = new(CakeTransfer)
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
func (it *CakeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakeTransfer represents a Transfer event raised by the Cake contract.
type CakeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Cake *CakeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CakeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Cake.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CakeTransferIterator{contract: _Cake.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Cake *CakeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CakeTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Cake.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakeTransfer)
				if err := _Cake.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Cake *CakeFilterer) ParseTransfer(log types.Log) (*CakeTransfer, error) {
	event := new(CakeTransfer)
	if err := _Cake.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

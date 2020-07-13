// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eccd_abi

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

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// EthCrossChainDataABI is the input ABI used to generate the binding from.
const EthCrossChainDataABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EthTxHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EthTxHashIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ExtraData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"LatestBookKeeperHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"LatestHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"MCHeaderBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MCKeeperHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"MCKeeperPubBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_crossChainTx\",\"type\":\"bytes32\"}],\"name\":\"getCrossChainTxExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethTxHashIndex\",\"type\":\"uint256\"}],\"name\":\"getEthTxHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEthTxHashIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key2\",\"type\":\"bytes32\"}],\"name\":\"getExtraData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLatestBookKeeperHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLatestHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_height\",\"type\":\"uint64\"}],\"name\":\"getMCHeaderBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMCKeeperHeight\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"getMCKeeperPubBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_height\",\"type\":\"uint64\"}],\"name\":\"getMCKeeperPubKeybytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initGenesisBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isGenesisBlockInited\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_crossChainTx\",\"type\":\"bytes32\"}],\"name\":\"putCrossChainTxExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ethTxHash\",\"type\":\"bytes32\"}],\"name\":\"putEthTxHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"putExtraData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_latestBookKeeperHeight\",\"type\":\"uint64\"}],\"name\":\"putLatestBookKeeperHeight\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_latestHeight\",\"type\":\"uint64\"}],\"name\":\"putLatestHeight\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_rawHeader\",\"type\":\"bytes\"}],\"name\":\"putMCHeaderBytes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"putMCKeeperHeight\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_MCKeeperPubBytes\",\"type\":\"bytes\"}],\"name\":\"putMCKeeperPubBytes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_keepersBytes\",\"type\":\"bytes\"}],\"name\":\"putMCKeeperPubKeybytes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EthCrossChainDataFuncSigs maps the 4-byte function signature to its string representation.
var EthCrossChainDataFuncSigs = map[string]string{
	"28e46109": "EthTxHash(uint256)",
	"8a463d3c": "EthTxHashIndex()",
	"20bbde38": "ExtraData(bytes32,bytes32)",
	"e244e4b1": "LatestBookKeeperHeight()",
	"e6aefed8": "LatestHeight()",
	"228a7cc5": "MCHeaderBytes(uint64)",
	"c6521e0f": "MCKeeperHeight(uint256)",
	"82cf210e": "MCKeeperPubBytes(uint64)",
	"cceb218d": "getCrossChainTxExist(bytes32)",
	"29927875": "getEthTxHash(uint256)",
	"ff3d24a1": "getEthTxHashIndex()",
	"40602bb5": "getExtraData(bytes32,bytes32)",
	"71a77750": "getLatestBookKeeperHeight()",
	"4ed1d8cc": "getLatestHeight()",
	"9743e9d6": "getMCHeaderBytes(uint64)",
	"f0b1fcfe": "getMCKeeperHeight()",
	"7d737b8b": "getMCKeeperPubBytes(uint64)",
	"09705412": "getMCKeeperPubKeybytes(uint64)",
	"8b39aa25": "initGenesisBlock()",
	"64f1acb0": "isGenesisBlockInited()",
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"794b0300": "putCrossChainTxExist(bytes32)",
	"4c3ccf64": "putEthTxHash(bytes32)",
	"1afe374e": "putExtraData(bytes32,bytes32,bytes)",
	"a6c6fe7c": "putLatestBookKeeperHeight(uint64)",
	"6d044440": "putLatestHeight(uint64)",
	"021d70ab": "putMCHeaderBytes(uint64,bytes)",
	"c6dbd0b7": "putMCKeeperHeight(uint64)",
	"191b0ab1": "putMCKeeperPubBytes(uint64,bytes)",
	"73c2f5f6": "putMCKeeperPubKeybytes(uint64,bytes)",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
	"3f4ba83a": "unpause()",
}

// EthCrossChainDataBin is the compiled bytecode used for deploying new contracts.
var EthCrossChainDataBin = "0x608060405234801561001057600080fd5b5060006100246001600160e01b0361008016565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506000805460ff60a01b19169055610084565b3390565b6118c3806100936000396000f3fe608060405234801561001057600080fd5b506004361061021c5760003560e01c8063794b030011610125578063a6c6fe7c116100ad578063e244e4b11161007c578063e244e4b11461076a578063e6aefed814610772578063f0b1fcfe1461077a578063f2fde38b146107d2578063ff3d24a1146107f85761021c565b8063a6c6fe7c146106e4578063c6521e0f1461070a578063c6dbd0b714610727578063cceb218d1461074d5761021c565b80638a463d3c116100f45780638a463d3c146106825780638b39aa251461068a5780638da5cb5b146106925780638f32d59b146106b65780639743e9d6146106be5761021c565b8063794b0300146106375780637d737b8b146102eb57806382cf210e146106545780638456cb591461067a5761021c565b806340602bb5116101a857806364f1acb01161017757806364f1acb0146105f75780636d044440146105ff578063715018a61461062557806371a777501461062f57806373c2f5f6146103865761021c565b806340602bb51461058b5780634c3ccf64146105ae5780634ed1d8cc146105cb5780635c975abb146105ef5761021c565b806320bbde38116101ef57806320bbde38146104ee578063228a7cc51461051157806328e461091461053757806329927875146105665780633f4ba83a146105835761021c565b8063021d70ab1461022157806309705412146102eb578063191b0ab1146103865780631afe374e1461043c575b600080fd5b6102d76004803603604081101561023757600080fd5b6001600160401b03823516919081019060408101602082013564010000000081111561026257600080fd5b82018360208201111561027457600080fd5b8035906020019184600183028401116401000000008311171561029657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610800945050505050565b604080519115158252519081900360200190f35b6103116004803603602081101561030157600080fd5b50356001600160401b03166108cd565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561034b578181015183820152602001610333565b50505050905090810190601f1680156103785780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102d76004803603604081101561039c57600080fd5b6001600160401b0382351691908101906040810160208201356401000000008111156103c757600080fd5b8201836020820111156103d957600080fd5b803590602001918460018302840111640100000000831117156103fb57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610978945050505050565b6102d76004803603606081101561045257600080fd5b81359160208101359181019060608101604082013564010000000081111561047957600080fd5b82018360208201111561048b57600080fd5b803590602001918460018302840111640100000000831117156104ad57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a3b945050505050565b6103116004803603604081101561050457600080fd5b5080359060200135610b07565b6103116004803603602081101561052757600080fd5b50356001600160401b0316610bab565b6105546004803603602081101561054d57600080fd5b5035610c13565b60408051918252519081900360200190f35b6105546004803603602081101561057c57600080fd5b5035610c25565b6102d7610c37565b610311600480360360408110156105a157600080fd5b5080359060200135610ce4565b6102d7600480360360208110156105c457600080fd5b5035610d8e565b6105d3610e4b565b604080516001600160401b039092168252519081900360200190f35b6102d7610e61565b6102d7610e71565b6102d76004803603602081101561061557600080fd5b50356001600160401b0316610e81565b61062d610f4f565b005b6105d3610fe0565b6102d76004803603602081101561064d57600080fd5b5035610fef565b6103116004803603602081101561066a57600080fd5b50356001600160401b03166110a9565b6102d7611111565b6105546111b4565b6102d76111ba565b61069a61121c565b604080516001600160a01b039092168252519081900360200190f35b6102d761122b565b610311600480360360208110156106d457600080fd5b50356001600160401b031661124f565b6102d7600480360360208110156106fa57600080fd5b50356001600160401b03166112c3565b6105d36004803603602081101561072057600080fd5b5035611383565b6102d76004803603602081101561073d57600080fd5b50356001600160401b03166113bd565b6102d76004803603602081101561076357600080fd5b50356114bd565b6105d36114d2565b6105d36114e1565b6107826114f7565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156107be5781810151838201526020016107a6565b505050509050019250505060405180910390f35b61062d600480360360208110156107e857600080fd5b50356001600160a01b0316611581565b6105546115d4565b60008054600160a01b900460ff1615610853576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b61085b61122b565b61089a576040805162461bcd60e51b8152602060048201819052602482015260008051602061186f833981519152604482015290519081900360640190fd5b6001600160401b038316600090815260036020908152604090912083516108c3928501906117b0565b5060019392505050565b6001600160401b03811660009081526004602090815260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084526060939283018282801561096c5780601f106109415761010080835404028352916020019161096c565b820191906000526020600020905b81548152906001019060200180831161094f57829003601f168201915b50505050509050919050565b60008054600160a01b900460ff16156109cb576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b6109d361122b565b610a12576040805162461bcd60e51b8152602060048201819052602482015260008051602061186f833981519152604482015290519081900360640190fd5b6001600160401b038316600090815260046020908152604090912083516108c3928501906117b0565b60008054600160a01b900460ff1615610a8e576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b610a9661122b565b610ad5576040805162461bcd60e51b8152602060048201819052602482015260008051602061186f833981519152604482015290519081900360640190fd5b600084815260086020908152604080832086845282529091208351610afc928501906117b0565b506001949350505050565b60086020908152600092835260408084208252918352918190208054825160026001831615610100026000190190921691909104601f810185900485028201850190935282815292909190830182828015610ba35780601f10610b7857610100808354040283529160200191610ba3565b820191906000526020600020905b815481529060010190602001808311610b8657829003601f168201915b505050505081565b60036020908152600091825260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015610ba35780601f10610b7857610100808354040283529160200191610ba3565b60016020526000908152604090205481565b60009081526001602052604090205490565b6000610c4161122b565b610c80576040805162461bcd60e51b8152602060048201819052602482015260008051602061186f833981519152604482015290519081900360640190fd5b600054600160a01b900460ff16610cd5576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b610cdd6115da565b5060015b90565b600082815260086020908152604080832084845282529182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845260609392830182828015610d815780601f10610d5657610100808354040283529160200191610d81565b820191906000526020600020905b815481529060010190602001808311610d6457829003601f168201915b5050505050905092915050565b60008054600160a01b900460ff1615610de1576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b610de961122b565b610e28576040805162461bcd60e51b8152602060048201819052602482015260008051602061186f833981519152604482015290519081900360640190fd5b506002805460009081526001602081905260409091209290925580548201905590565b600754600160401b90046001600160401b031690565b600054600160a01b900460ff1690565b600754600160801b900460ff1690565b60008054600160a01b900460ff1615610ed4576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b610edc61122b565b610f1b576040805162461bcd60e51b8152602060048201819052602482015260008051602061186f833981519152604482015290519081900360640190fd5b50600780546001600160401b038316600160401b026fffffffffffffffff0000000000000000199091161790556001919050565b610f5761122b565b610f96576040805162461bcd60e51b8152602060048201819052602482015260008051602061186f833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6007546001600160401b031690565b60008054600160a01b900460ff1615611042576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b61104a61122b565b611089576040805162461bcd60e51b8152602060048201819052602482015260008051602061186f833981519152604482015290519081900360640190fd5b506000908152600560205260409020805460ff1916600190811790915590565b60046020908152600091825260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015610ba35780601f10610b7857610100808354040283529160200191610ba3565b600061111b61122b565b61115a576040805162461bcd60e51b8152602060048201819052602482015260008051602061186f833981519152604482015290519081900360640190fd5b600054600160a01b900460ff16156111ac576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b610cdd611682565b60025481565b60006111c461122b565b611203576040805162461bcd60e51b8152602060048201819052602482015260008051602061186f833981519152604482015290519081900360640190fd5b506007805460ff60801b1916600160801b179055600190565b6000546001600160a01b031690565b600080546001600160a01b031661124061170c565b6001600160a01b031614905090565b6001600160401b03811660009081526003602090815260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084526060939283018282801561096c5780601f106109415761010080835404028352916020019161096c565b60008054600160a01b900460ff1615611316576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b61131e61122b565b61135d576040805162461bcd60e51b8152602060048201819052602482015260008051602061186f833981519152604482015290519081900360640190fd5b50600780546001600160401b03831667ffffffffffffffff199091161790556001919050565b6006818154811061139057fe5b9060005260206000209060049182820401919006600802915054906101000a90046001600160401b031681565b60008054600160a01b900460ff1615611410576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b61141861122b565b611457576040805162461bcd60e51b8152602060048201819052602482015260008051602061186f833981519152604482015290519081900360640190fd5b506006805460018181018355600092909252600481047ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f0180546001600160401b0380861660086003909516949094026101000a93840293021916919091179055919050565b60009081526005602052604090205460ff1690565b6007546001600160401b031681565b600754600160401b90046001600160401b031681565b6060600680548060200260200160405190810160405280929190818152602001828054801561157757602002820191906000526020600020906000905b82829054906101000a90046001600160401b03166001600160401b0316815260200190600801906020826007010492830192600103820291508084116115345790505b5050505050905090565b61158961122b565b6115c8576040805162461bcd60e51b8152602060048201819052602482015260008051602061186f833981519152604482015290519081900360640190fd5b6115d181611710565b50565b60025490565b600054600160a01b900460ff1661162f576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa61166561170c565b604080516001600160a01b039092168252519081900360200190a1565b600054600160a01b900460ff16156116d4576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586116655b3390565b6001600160a01b0381166117555760405162461bcd60e51b81526004018080602001828103825260268152602001806118496026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106117f157805160ff191683800117855561181e565b8280016001018555821561181e579182015b8281111561181e578251825591602001919060010190611803565b5061182a92915061182e565b5090565b610ce191905b8082111561182a576000815560010161183456fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a265627a7a72315820ae935078ffe9cfec37b19556fef2009e99e5a6cd279135bf4a619c6e2ea45f3764736f6c634300050f0032"

// DeployEthCrossChainData deploys a new Ethereum contract, binding an instance of EthCrossChainData to it.
func DeployEthCrossChainData(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthCrossChainData, error) {
	parsed, err := abi.JSON(strings.NewReader(EthCrossChainDataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthCrossChainDataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthCrossChainData{EthCrossChainDataCaller: EthCrossChainDataCaller{contract: contract}, EthCrossChainDataTransactor: EthCrossChainDataTransactor{contract: contract}, EthCrossChainDataFilterer: EthCrossChainDataFilterer{contract: contract}}, nil
}

// EthCrossChainData is an auto generated Go binding around an Ethereum contract.
type EthCrossChainData struct {
	EthCrossChainDataCaller     // Read-only binding to the contract
	EthCrossChainDataTransactor // Write-only binding to the contract
	EthCrossChainDataFilterer   // Log filterer for contract events
}

// EthCrossChainDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthCrossChainDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthCrossChainDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthCrossChainDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthCrossChainDataSession struct {
	Contract     *EthCrossChainData // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EthCrossChainDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthCrossChainDataCallerSession struct {
	Contract *EthCrossChainDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// EthCrossChainDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthCrossChainDataTransactorSession struct {
	Contract     *EthCrossChainDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// EthCrossChainDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthCrossChainDataRaw struct {
	Contract *EthCrossChainData // Generic contract binding to access the raw methods on
}

// EthCrossChainDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthCrossChainDataCallerRaw struct {
	Contract *EthCrossChainDataCaller // Generic read-only contract binding to access the raw methods on
}

// EthCrossChainDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthCrossChainDataTransactorRaw struct {
	Contract *EthCrossChainDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthCrossChainData creates a new instance of EthCrossChainData, bound to a specific deployed contract.
func NewEthCrossChainData(address common.Address, backend bind.ContractBackend) (*EthCrossChainData, error) {
	contract, err := bindEthCrossChainData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainData{EthCrossChainDataCaller: EthCrossChainDataCaller{contract: contract}, EthCrossChainDataTransactor: EthCrossChainDataTransactor{contract: contract}, EthCrossChainDataFilterer: EthCrossChainDataFilterer{contract: contract}}, nil
}

// NewEthCrossChainDataCaller creates a new read-only instance of EthCrossChainData, bound to a specific deployed contract.
func NewEthCrossChainDataCaller(address common.Address, caller bind.ContractCaller) (*EthCrossChainDataCaller, error) {
	contract, err := bindEthCrossChainData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainDataCaller{contract: contract}, nil
}

// NewEthCrossChainDataTransactor creates a new write-only instance of EthCrossChainData, bound to a specific deployed contract.
func NewEthCrossChainDataTransactor(address common.Address, transactor bind.ContractTransactor) (*EthCrossChainDataTransactor, error) {
	contract, err := bindEthCrossChainData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainDataTransactor{contract: contract}, nil
}

// NewEthCrossChainDataFilterer creates a new log filterer instance of EthCrossChainData, bound to a specific deployed contract.
func NewEthCrossChainDataFilterer(address common.Address, filterer bind.ContractFilterer) (*EthCrossChainDataFilterer, error) {
	contract, err := bindEthCrossChainData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainDataFilterer{contract: contract}, nil
}

// bindEthCrossChainData binds a generic wrapper to an already deployed contract.
func bindEthCrossChainData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthCrossChainDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthCrossChainData *EthCrossChainDataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthCrossChainData.Contract.EthCrossChainDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthCrossChainData *EthCrossChainDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.EthCrossChainDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthCrossChainData *EthCrossChainDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.EthCrossChainDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthCrossChainData *EthCrossChainDataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthCrossChainData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthCrossChainData *EthCrossChainDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthCrossChainData *EthCrossChainDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.contract.Transact(opts, method, params...)
}

// EthTxHash is a free data retrieval call binding the contract method 0x28e46109.
//
// Solidity: function EthTxHash(uint256 ) view returns(bytes32)
func (_EthCrossChainData *EthCrossChainDataCaller) EthTxHash(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "EthTxHash", arg0)
	return *ret0, err
}

// EthTxHash is a free data retrieval call binding the contract method 0x28e46109.
//
// Solidity: function EthTxHash(uint256 ) view returns(bytes32)
func (_EthCrossChainData *EthCrossChainDataSession) EthTxHash(arg0 *big.Int) ([32]byte, error) {
	return _EthCrossChainData.Contract.EthTxHash(&_EthCrossChainData.CallOpts, arg0)
}

// EthTxHash is a free data retrieval call binding the contract method 0x28e46109.
//
// Solidity: function EthTxHash(uint256 ) view returns(bytes32)
func (_EthCrossChainData *EthCrossChainDataCallerSession) EthTxHash(arg0 *big.Int) ([32]byte, error) {
	return _EthCrossChainData.Contract.EthTxHash(&_EthCrossChainData.CallOpts, arg0)
}

// EthTxHashIndex is a free data retrieval call binding the contract method 0x8a463d3c.
//
// Solidity: function EthTxHashIndex() view returns(uint256)
func (_EthCrossChainData *EthCrossChainDataCaller) EthTxHashIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "EthTxHashIndex")
	return *ret0, err
}

// EthTxHashIndex is a free data retrieval call binding the contract method 0x8a463d3c.
//
// Solidity: function EthTxHashIndex() view returns(uint256)
func (_EthCrossChainData *EthCrossChainDataSession) EthTxHashIndex() (*big.Int, error) {
	return _EthCrossChainData.Contract.EthTxHashIndex(&_EthCrossChainData.CallOpts)
}

// EthTxHashIndex is a free data retrieval call binding the contract method 0x8a463d3c.
//
// Solidity: function EthTxHashIndex() view returns(uint256)
func (_EthCrossChainData *EthCrossChainDataCallerSession) EthTxHashIndex() (*big.Int, error) {
	return _EthCrossChainData.Contract.EthTxHashIndex(&_EthCrossChainData.CallOpts)
}

// ExtraData is a free data retrieval call binding the contract method 0x20bbde38.
//
// Solidity: function ExtraData(bytes32 , bytes32 ) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCaller) ExtraData(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "ExtraData", arg0, arg1)
	return *ret0, err
}

// ExtraData is a free data retrieval call binding the contract method 0x20bbde38.
//
// Solidity: function ExtraData(bytes32 , bytes32 ) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataSession) ExtraData(arg0 [32]byte, arg1 [32]byte) ([]byte, error) {
	return _EthCrossChainData.Contract.ExtraData(&_EthCrossChainData.CallOpts, arg0, arg1)
}

// ExtraData is a free data retrieval call binding the contract method 0x20bbde38.
//
// Solidity: function ExtraData(bytes32 , bytes32 ) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCallerSession) ExtraData(arg0 [32]byte, arg1 [32]byte) ([]byte, error) {
	return _EthCrossChainData.Contract.ExtraData(&_EthCrossChainData.CallOpts, arg0, arg1)
}

// LatestBookKeeperHeight is a free data retrieval call binding the contract method 0xe244e4b1.
//
// Solidity: function LatestBookKeeperHeight() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataCaller) LatestBookKeeperHeight(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "LatestBookKeeperHeight")
	return *ret0, err
}

// LatestBookKeeperHeight is a free data retrieval call binding the contract method 0xe244e4b1.
//
// Solidity: function LatestBookKeeperHeight() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataSession) LatestBookKeeperHeight() (uint64, error) {
	return _EthCrossChainData.Contract.LatestBookKeeperHeight(&_EthCrossChainData.CallOpts)
}

// LatestBookKeeperHeight is a free data retrieval call binding the contract method 0xe244e4b1.
//
// Solidity: function LatestBookKeeperHeight() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataCallerSession) LatestBookKeeperHeight() (uint64, error) {
	return _EthCrossChainData.Contract.LatestBookKeeperHeight(&_EthCrossChainData.CallOpts)
}

// LatestHeight is a free data retrieval call binding the contract method 0xe6aefed8.
//
// Solidity: function LatestHeight() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataCaller) LatestHeight(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "LatestHeight")
	return *ret0, err
}

// LatestHeight is a free data retrieval call binding the contract method 0xe6aefed8.
//
// Solidity: function LatestHeight() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataSession) LatestHeight() (uint64, error) {
	return _EthCrossChainData.Contract.LatestHeight(&_EthCrossChainData.CallOpts)
}

// LatestHeight is a free data retrieval call binding the contract method 0xe6aefed8.
//
// Solidity: function LatestHeight() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataCallerSession) LatestHeight() (uint64, error) {
	return _EthCrossChainData.Contract.LatestHeight(&_EthCrossChainData.CallOpts)
}

// MCHeaderBytes is a free data retrieval call binding the contract method 0x228a7cc5.
//
// Solidity: function MCHeaderBytes(uint64 ) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCaller) MCHeaderBytes(opts *bind.CallOpts, arg0 uint64) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "MCHeaderBytes", arg0)
	return *ret0, err
}

// MCHeaderBytes is a free data retrieval call binding the contract method 0x228a7cc5.
//
// Solidity: function MCHeaderBytes(uint64 ) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataSession) MCHeaderBytes(arg0 uint64) ([]byte, error) {
	return _EthCrossChainData.Contract.MCHeaderBytes(&_EthCrossChainData.CallOpts, arg0)
}

// MCHeaderBytes is a free data retrieval call binding the contract method 0x228a7cc5.
//
// Solidity: function MCHeaderBytes(uint64 ) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCallerSession) MCHeaderBytes(arg0 uint64) ([]byte, error) {
	return _EthCrossChainData.Contract.MCHeaderBytes(&_EthCrossChainData.CallOpts, arg0)
}

// MCKeeperHeight is a free data retrieval call binding the contract method 0xc6521e0f.
//
// Solidity: function MCKeeperHeight(uint256 ) view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataCaller) MCKeeperHeight(opts *bind.CallOpts, arg0 *big.Int) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "MCKeeperHeight", arg0)
	return *ret0, err
}

// MCKeeperHeight is a free data retrieval call binding the contract method 0xc6521e0f.
//
// Solidity: function MCKeeperHeight(uint256 ) view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataSession) MCKeeperHeight(arg0 *big.Int) (uint64, error) {
	return _EthCrossChainData.Contract.MCKeeperHeight(&_EthCrossChainData.CallOpts, arg0)
}

// MCKeeperHeight is a free data retrieval call binding the contract method 0xc6521e0f.
//
// Solidity: function MCKeeperHeight(uint256 ) view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataCallerSession) MCKeeperHeight(arg0 *big.Int) (uint64, error) {
	return _EthCrossChainData.Contract.MCKeeperHeight(&_EthCrossChainData.CallOpts, arg0)
}

// MCKeeperPubBytes is a free data retrieval call binding the contract method 0x82cf210e.
//
// Solidity: function MCKeeperPubBytes(uint64 ) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCaller) MCKeeperPubBytes(opts *bind.CallOpts, arg0 uint64) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "MCKeeperPubBytes", arg0)
	return *ret0, err
}

// MCKeeperPubBytes is a free data retrieval call binding the contract method 0x82cf210e.
//
// Solidity: function MCKeeperPubBytes(uint64 ) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataSession) MCKeeperPubBytes(arg0 uint64) ([]byte, error) {
	return _EthCrossChainData.Contract.MCKeeperPubBytes(&_EthCrossChainData.CallOpts, arg0)
}

// MCKeeperPubBytes is a free data retrieval call binding the contract method 0x82cf210e.
//
// Solidity: function MCKeeperPubBytes(uint64 ) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCallerSession) MCKeeperPubBytes(arg0 uint64) ([]byte, error) {
	return _EthCrossChainData.Contract.MCKeeperPubBytes(&_EthCrossChainData.CallOpts, arg0)
}

// GetCrossChainTxExist is a free data retrieval call binding the contract method 0xcceb218d.
//
// Solidity: function getCrossChainTxExist(bytes32 _crossChainTx) view returns(bool)
func (_EthCrossChainData *EthCrossChainDataCaller) GetCrossChainTxExist(opts *bind.CallOpts, _crossChainTx [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "getCrossChainTxExist", _crossChainTx)
	return *ret0, err
}

// GetCrossChainTxExist is a free data retrieval call binding the contract method 0xcceb218d.
//
// Solidity: function getCrossChainTxExist(bytes32 _crossChainTx) view returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) GetCrossChainTxExist(_crossChainTx [32]byte) (bool, error) {
	return _EthCrossChainData.Contract.GetCrossChainTxExist(&_EthCrossChainData.CallOpts, _crossChainTx)
}

// GetCrossChainTxExist is a free data retrieval call binding the contract method 0xcceb218d.
//
// Solidity: function getCrossChainTxExist(bytes32 _crossChainTx) view returns(bool)
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetCrossChainTxExist(_crossChainTx [32]byte) (bool, error) {
	return _EthCrossChainData.Contract.GetCrossChainTxExist(&_EthCrossChainData.CallOpts, _crossChainTx)
}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 _ethTxHashIndex) view returns(bytes32)
func (_EthCrossChainData *EthCrossChainDataCaller) GetEthTxHash(opts *bind.CallOpts, _ethTxHashIndex *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "getEthTxHash", _ethTxHashIndex)
	return *ret0, err
}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 _ethTxHashIndex) view returns(bytes32)
func (_EthCrossChainData *EthCrossChainDataSession) GetEthTxHash(_ethTxHashIndex *big.Int) ([32]byte, error) {
	return _EthCrossChainData.Contract.GetEthTxHash(&_EthCrossChainData.CallOpts, _ethTxHashIndex)
}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 _ethTxHashIndex) view returns(bytes32)
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetEthTxHash(_ethTxHashIndex *big.Int) ([32]byte, error) {
	return _EthCrossChainData.Contract.GetEthTxHash(&_EthCrossChainData.CallOpts, _ethTxHashIndex)
}

// GetEthTxHashIndex is a free data retrieval call binding the contract method 0xff3d24a1.
//
// Solidity: function getEthTxHashIndex() view returns(uint256)
func (_EthCrossChainData *EthCrossChainDataCaller) GetEthTxHashIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "getEthTxHashIndex")
	return *ret0, err
}

// GetEthTxHashIndex is a free data retrieval call binding the contract method 0xff3d24a1.
//
// Solidity: function getEthTxHashIndex() view returns(uint256)
func (_EthCrossChainData *EthCrossChainDataSession) GetEthTxHashIndex() (*big.Int, error) {
	return _EthCrossChainData.Contract.GetEthTxHashIndex(&_EthCrossChainData.CallOpts)
}

// GetEthTxHashIndex is a free data retrieval call binding the contract method 0xff3d24a1.
//
// Solidity: function getEthTxHashIndex() view returns(uint256)
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetEthTxHashIndex() (*big.Int, error) {
	return _EthCrossChainData.Contract.GetEthTxHashIndex(&_EthCrossChainData.CallOpts)
}

// GetExtraData is a free data retrieval call binding the contract method 0x40602bb5.
//
// Solidity: function getExtraData(bytes32 key1, bytes32 key2) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCaller) GetExtraData(opts *bind.CallOpts, key1 [32]byte, key2 [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "getExtraData", key1, key2)
	return *ret0, err
}

// GetExtraData is a free data retrieval call binding the contract method 0x40602bb5.
//
// Solidity: function getExtraData(bytes32 key1, bytes32 key2) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataSession) GetExtraData(key1 [32]byte, key2 [32]byte) ([]byte, error) {
	return _EthCrossChainData.Contract.GetExtraData(&_EthCrossChainData.CallOpts, key1, key2)
}

// GetExtraData is a free data retrieval call binding the contract method 0x40602bb5.
//
// Solidity: function getExtraData(bytes32 key1, bytes32 key2) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetExtraData(key1 [32]byte, key2 [32]byte) ([]byte, error) {
	return _EthCrossChainData.Contract.GetExtraData(&_EthCrossChainData.CallOpts, key1, key2)
}

// GetLatestBookKeeperHeight is a free data retrieval call binding the contract method 0x71a77750.
//
// Solidity: function getLatestBookKeeperHeight() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataCaller) GetLatestBookKeeperHeight(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "getLatestBookKeeperHeight")
	return *ret0, err
}

// GetLatestBookKeeperHeight is a free data retrieval call binding the contract method 0x71a77750.
//
// Solidity: function getLatestBookKeeperHeight() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataSession) GetLatestBookKeeperHeight() (uint64, error) {
	return _EthCrossChainData.Contract.GetLatestBookKeeperHeight(&_EthCrossChainData.CallOpts)
}

// GetLatestBookKeeperHeight is a free data retrieval call binding the contract method 0x71a77750.
//
// Solidity: function getLatestBookKeeperHeight() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetLatestBookKeeperHeight() (uint64, error) {
	return _EthCrossChainData.Contract.GetLatestBookKeeperHeight(&_EthCrossChainData.CallOpts)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x4ed1d8cc.
//
// Solidity: function getLatestHeight() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataCaller) GetLatestHeight(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "getLatestHeight")
	return *ret0, err
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x4ed1d8cc.
//
// Solidity: function getLatestHeight() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataSession) GetLatestHeight() (uint64, error) {
	return _EthCrossChainData.Contract.GetLatestHeight(&_EthCrossChainData.CallOpts)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x4ed1d8cc.
//
// Solidity: function getLatestHeight() view returns(uint64)
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetLatestHeight() (uint64, error) {
	return _EthCrossChainData.Contract.GetLatestHeight(&_EthCrossChainData.CallOpts)
}

// GetMCHeaderBytes is a free data retrieval call binding the contract method 0x9743e9d6.
//
// Solidity: function getMCHeaderBytes(uint64 _height) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCaller) GetMCHeaderBytes(opts *bind.CallOpts, _height uint64) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "getMCHeaderBytes", _height)
	return *ret0, err
}

// GetMCHeaderBytes is a free data retrieval call binding the contract method 0x9743e9d6.
//
// Solidity: function getMCHeaderBytes(uint64 _height) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataSession) GetMCHeaderBytes(_height uint64) ([]byte, error) {
	return _EthCrossChainData.Contract.GetMCHeaderBytes(&_EthCrossChainData.CallOpts, _height)
}

// GetMCHeaderBytes is a free data retrieval call binding the contract method 0x9743e9d6.
//
// Solidity: function getMCHeaderBytes(uint64 _height) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetMCHeaderBytes(_height uint64) ([]byte, error) {
	return _EthCrossChainData.Contract.GetMCHeaderBytes(&_EthCrossChainData.CallOpts, _height)
}

// GetMCKeeperHeight is a free data retrieval call binding the contract method 0xf0b1fcfe.
//
// Solidity: function getMCKeeperHeight() view returns(uint64[])
func (_EthCrossChainData *EthCrossChainDataCaller) GetMCKeeperHeight(opts *bind.CallOpts) ([]uint64, error) {
	var (
		ret0 = new([]uint64)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "getMCKeeperHeight")
	return *ret0, err
}

// GetMCKeeperHeight is a free data retrieval call binding the contract method 0xf0b1fcfe.
//
// Solidity: function getMCKeeperHeight() view returns(uint64[])
func (_EthCrossChainData *EthCrossChainDataSession) GetMCKeeperHeight() ([]uint64, error) {
	return _EthCrossChainData.Contract.GetMCKeeperHeight(&_EthCrossChainData.CallOpts)
}

// GetMCKeeperHeight is a free data retrieval call binding the contract method 0xf0b1fcfe.
//
// Solidity: function getMCKeeperHeight() view returns(uint64[])
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetMCKeeperHeight() ([]uint64, error) {
	return _EthCrossChainData.Contract.GetMCKeeperHeight(&_EthCrossChainData.CallOpts)
}

// GetMCKeeperPubBytes is a free data retrieval call binding the contract method 0x7d737b8b.
//
// Solidity: function getMCKeeperPubBytes(uint64 height) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCaller) GetMCKeeperPubBytes(opts *bind.CallOpts, height uint64) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "getMCKeeperPubBytes", height)
	return *ret0, err
}

// GetMCKeeperPubBytes is a free data retrieval call binding the contract method 0x7d737b8b.
//
// Solidity: function getMCKeeperPubBytes(uint64 height) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataSession) GetMCKeeperPubBytes(height uint64) ([]byte, error) {
	return _EthCrossChainData.Contract.GetMCKeeperPubBytes(&_EthCrossChainData.CallOpts, height)
}

// GetMCKeeperPubBytes is a free data retrieval call binding the contract method 0x7d737b8b.
//
// Solidity: function getMCKeeperPubBytes(uint64 height) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetMCKeeperPubBytes(height uint64) ([]byte, error) {
	return _EthCrossChainData.Contract.GetMCKeeperPubBytes(&_EthCrossChainData.CallOpts, height)
}

// GetMCKeeperPubKeybytes is a free data retrieval call binding the contract method 0x09705412.
//
// Solidity: function getMCKeeperPubKeybytes(uint64 _height) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCaller) GetMCKeeperPubKeybytes(opts *bind.CallOpts, _height uint64) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "getMCKeeperPubKeybytes", _height)
	return *ret0, err
}

// GetMCKeeperPubKeybytes is a free data retrieval call binding the contract method 0x09705412.
//
// Solidity: function getMCKeeperPubKeybytes(uint64 _height) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataSession) GetMCKeeperPubKeybytes(_height uint64) ([]byte, error) {
	return _EthCrossChainData.Contract.GetMCKeeperPubKeybytes(&_EthCrossChainData.CallOpts, _height)
}

// GetMCKeeperPubKeybytes is a free data retrieval call binding the contract method 0x09705412.
//
// Solidity: function getMCKeeperPubKeybytes(uint64 _height) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetMCKeeperPubKeybytes(_height uint64) ([]byte, error) {
	return _EthCrossChainData.Contract.GetMCKeeperPubKeybytes(&_EthCrossChainData.CallOpts, _height)
}

// IsGenesisBlockInited is a free data retrieval call binding the contract method 0x64f1acb0.
//
// Solidity: function isGenesisBlockInited() view returns(bool)
func (_EthCrossChainData *EthCrossChainDataCaller) IsGenesisBlockInited(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "isGenesisBlockInited")
	return *ret0, err
}

// IsGenesisBlockInited is a free data retrieval call binding the contract method 0x64f1acb0.
//
// Solidity: function isGenesisBlockInited() view returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) IsGenesisBlockInited() (bool, error) {
	return _EthCrossChainData.Contract.IsGenesisBlockInited(&_EthCrossChainData.CallOpts)
}

// IsGenesisBlockInited is a free data retrieval call binding the contract method 0x64f1acb0.
//
// Solidity: function isGenesisBlockInited() view returns(bool)
func (_EthCrossChainData *EthCrossChainDataCallerSession) IsGenesisBlockInited() (bool, error) {
	return _EthCrossChainData.Contract.IsGenesisBlockInited(&_EthCrossChainData.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_EthCrossChainData *EthCrossChainDataCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) IsOwner() (bool, error) {
	return _EthCrossChainData.Contract.IsOwner(&_EthCrossChainData.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_EthCrossChainData *EthCrossChainDataCallerSession) IsOwner() (bool, error) {
	return _EthCrossChainData.Contract.IsOwner(&_EthCrossChainData.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthCrossChainData *EthCrossChainDataCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthCrossChainData *EthCrossChainDataSession) Owner() (common.Address, error) {
	return _EthCrossChainData.Contract.Owner(&_EthCrossChainData.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthCrossChainData *EthCrossChainDataCallerSession) Owner() (common.Address, error) {
	return _EthCrossChainData.Contract.Owner(&_EthCrossChainData.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EthCrossChainData *EthCrossChainDataCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) Paused() (bool, error) {
	return _EthCrossChainData.Contract.Paused(&_EthCrossChainData.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EthCrossChainData *EthCrossChainDataCallerSession) Paused() (bool, error) {
	return _EthCrossChainData.Contract.Paused(&_EthCrossChainData.CallOpts)
}

// InitGenesisBlock is a paid mutator transaction binding the contract method 0x8b39aa25.
//
// Solidity: function initGenesisBlock() returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) InitGenesisBlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "initGenesisBlock")
}

// InitGenesisBlock is a paid mutator transaction binding the contract method 0x8b39aa25.
//
// Solidity: function initGenesisBlock() returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) InitGenesisBlock() (*types.Transaction, error) {
	return _EthCrossChainData.Contract.InitGenesisBlock(&_EthCrossChainData.TransactOpts)
}

// InitGenesisBlock is a paid mutator transaction binding the contract method 0x8b39aa25.
//
// Solidity: function initGenesisBlock() returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) InitGenesisBlock() (*types.Transaction, error) {
	return _EthCrossChainData.Contract.InitGenesisBlock(&_EthCrossChainData.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) Pause() (*types.Transaction, error) {
	return _EthCrossChainData.Contract.Pause(&_EthCrossChainData.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) Pause() (*types.Transaction, error) {
	return _EthCrossChainData.Contract.Pause(&_EthCrossChainData.TransactOpts)
}

// PutCrossChainTxExist is a paid mutator transaction binding the contract method 0x794b0300.
//
// Solidity: function putCrossChainTxExist(bytes32 _crossChainTx) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) PutCrossChainTxExist(opts *bind.TransactOpts, _crossChainTx [32]byte) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "putCrossChainTxExist", _crossChainTx)
}

// PutCrossChainTxExist is a paid mutator transaction binding the contract method 0x794b0300.
//
// Solidity: function putCrossChainTxExist(bytes32 _crossChainTx) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) PutCrossChainTxExist(_crossChainTx [32]byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutCrossChainTxExist(&_EthCrossChainData.TransactOpts, _crossChainTx)
}

// PutCrossChainTxExist is a paid mutator transaction binding the contract method 0x794b0300.
//
// Solidity: function putCrossChainTxExist(bytes32 _crossChainTx) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) PutCrossChainTxExist(_crossChainTx [32]byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutCrossChainTxExist(&_EthCrossChainData.TransactOpts, _crossChainTx)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x4c3ccf64.
//
// Solidity: function putEthTxHash(bytes32 _ethTxHash) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) PutEthTxHash(opts *bind.TransactOpts, _ethTxHash [32]byte) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "putEthTxHash", _ethTxHash)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x4c3ccf64.
//
// Solidity: function putEthTxHash(bytes32 _ethTxHash) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) PutEthTxHash(_ethTxHash [32]byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutEthTxHash(&_EthCrossChainData.TransactOpts, _ethTxHash)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x4c3ccf64.
//
// Solidity: function putEthTxHash(bytes32 _ethTxHash) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) PutEthTxHash(_ethTxHash [32]byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutEthTxHash(&_EthCrossChainData.TransactOpts, _ethTxHash)
}

// PutExtraData is a paid mutator transaction binding the contract method 0x1afe374e.
//
// Solidity: function putExtraData(bytes32 key1, bytes32 key2, bytes value) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) PutExtraData(opts *bind.TransactOpts, key1 [32]byte, key2 [32]byte, value []byte) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "putExtraData", key1, key2, value)
}

// PutExtraData is a paid mutator transaction binding the contract method 0x1afe374e.
//
// Solidity: function putExtraData(bytes32 key1, bytes32 key2, bytes value) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) PutExtraData(key1 [32]byte, key2 [32]byte, value []byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutExtraData(&_EthCrossChainData.TransactOpts, key1, key2, value)
}

// PutExtraData is a paid mutator transaction binding the contract method 0x1afe374e.
//
// Solidity: function putExtraData(bytes32 key1, bytes32 key2, bytes value) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) PutExtraData(key1 [32]byte, key2 [32]byte, value []byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutExtraData(&_EthCrossChainData.TransactOpts, key1, key2, value)
}

// PutLatestBookKeeperHeight is a paid mutator transaction binding the contract method 0xa6c6fe7c.
//
// Solidity: function putLatestBookKeeperHeight(uint64 _latestBookKeeperHeight) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) PutLatestBookKeeperHeight(opts *bind.TransactOpts, _latestBookKeeperHeight uint64) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "putLatestBookKeeperHeight", _latestBookKeeperHeight)
}

// PutLatestBookKeeperHeight is a paid mutator transaction binding the contract method 0xa6c6fe7c.
//
// Solidity: function putLatestBookKeeperHeight(uint64 _latestBookKeeperHeight) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) PutLatestBookKeeperHeight(_latestBookKeeperHeight uint64) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutLatestBookKeeperHeight(&_EthCrossChainData.TransactOpts, _latestBookKeeperHeight)
}

// PutLatestBookKeeperHeight is a paid mutator transaction binding the contract method 0xa6c6fe7c.
//
// Solidity: function putLatestBookKeeperHeight(uint64 _latestBookKeeperHeight) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) PutLatestBookKeeperHeight(_latestBookKeeperHeight uint64) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutLatestBookKeeperHeight(&_EthCrossChainData.TransactOpts, _latestBookKeeperHeight)
}

// PutLatestHeight is a paid mutator transaction binding the contract method 0x6d044440.
//
// Solidity: function putLatestHeight(uint64 _latestHeight) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) PutLatestHeight(opts *bind.TransactOpts, _latestHeight uint64) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "putLatestHeight", _latestHeight)
}

// PutLatestHeight is a paid mutator transaction binding the contract method 0x6d044440.
//
// Solidity: function putLatestHeight(uint64 _latestHeight) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) PutLatestHeight(_latestHeight uint64) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutLatestHeight(&_EthCrossChainData.TransactOpts, _latestHeight)
}

// PutLatestHeight is a paid mutator transaction binding the contract method 0x6d044440.
//
// Solidity: function putLatestHeight(uint64 _latestHeight) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) PutLatestHeight(_latestHeight uint64) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutLatestHeight(&_EthCrossChainData.TransactOpts, _latestHeight)
}

// PutMCHeaderBytes is a paid mutator transaction binding the contract method 0x021d70ab.
//
// Solidity: function putMCHeaderBytes(uint64 _height, bytes _rawHeader) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) PutMCHeaderBytes(opts *bind.TransactOpts, _height uint64, _rawHeader []byte) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "putMCHeaderBytes", _height, _rawHeader)
}

// PutMCHeaderBytes is a paid mutator transaction binding the contract method 0x021d70ab.
//
// Solidity: function putMCHeaderBytes(uint64 _height, bytes _rawHeader) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) PutMCHeaderBytes(_height uint64, _rawHeader []byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutMCHeaderBytes(&_EthCrossChainData.TransactOpts, _height, _rawHeader)
}

// PutMCHeaderBytes is a paid mutator transaction binding the contract method 0x021d70ab.
//
// Solidity: function putMCHeaderBytes(uint64 _height, bytes _rawHeader) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) PutMCHeaderBytes(_height uint64, _rawHeader []byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutMCHeaderBytes(&_EthCrossChainData.TransactOpts, _height, _rawHeader)
}

// PutMCKeeperHeight is a paid mutator transaction binding the contract method 0xc6dbd0b7.
//
// Solidity: function putMCKeeperHeight(uint64 height) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) PutMCKeeperHeight(opts *bind.TransactOpts, height uint64) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "putMCKeeperHeight", height)
}

// PutMCKeeperHeight is a paid mutator transaction binding the contract method 0xc6dbd0b7.
//
// Solidity: function putMCKeeperHeight(uint64 height) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) PutMCKeeperHeight(height uint64) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutMCKeeperHeight(&_EthCrossChainData.TransactOpts, height)
}

// PutMCKeeperHeight is a paid mutator transaction binding the contract method 0xc6dbd0b7.
//
// Solidity: function putMCKeeperHeight(uint64 height) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) PutMCKeeperHeight(height uint64) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutMCKeeperHeight(&_EthCrossChainData.TransactOpts, height)
}

// PutMCKeeperPubBytes is a paid mutator transaction binding the contract method 0x191b0ab1.
//
// Solidity: function putMCKeeperPubBytes(uint64 height, bytes _MCKeeperPubBytes) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) PutMCKeeperPubBytes(opts *bind.TransactOpts, height uint64, _MCKeeperPubBytes []byte) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "putMCKeeperPubBytes", height, _MCKeeperPubBytes)
}

// PutMCKeeperPubBytes is a paid mutator transaction binding the contract method 0x191b0ab1.
//
// Solidity: function putMCKeeperPubBytes(uint64 height, bytes _MCKeeperPubBytes) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) PutMCKeeperPubBytes(height uint64, _MCKeeperPubBytes []byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutMCKeeperPubBytes(&_EthCrossChainData.TransactOpts, height, _MCKeeperPubBytes)
}

// PutMCKeeperPubBytes is a paid mutator transaction binding the contract method 0x191b0ab1.
//
// Solidity: function putMCKeeperPubBytes(uint64 height, bytes _MCKeeperPubBytes) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) PutMCKeeperPubBytes(height uint64, _MCKeeperPubBytes []byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutMCKeeperPubBytes(&_EthCrossChainData.TransactOpts, height, _MCKeeperPubBytes)
}

// PutMCKeeperPubKeybytes is a paid mutator transaction binding the contract method 0x73c2f5f6.
//
// Solidity: function putMCKeeperPubKeybytes(uint64 _height, bytes _keepersBytes) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) PutMCKeeperPubKeybytes(opts *bind.TransactOpts, _height uint64, _keepersBytes []byte) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "putMCKeeperPubKeybytes", _height, _keepersBytes)
}

// PutMCKeeperPubKeybytes is a paid mutator transaction binding the contract method 0x73c2f5f6.
//
// Solidity: function putMCKeeperPubKeybytes(uint64 _height, bytes _keepersBytes) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) PutMCKeeperPubKeybytes(_height uint64, _keepersBytes []byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutMCKeeperPubKeybytes(&_EthCrossChainData.TransactOpts, _height, _keepersBytes)
}

// PutMCKeeperPubKeybytes is a paid mutator transaction binding the contract method 0x73c2f5f6.
//
// Solidity: function putMCKeeperPubKeybytes(uint64 _height, bytes _keepersBytes) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) PutMCKeeperPubKeybytes(_height uint64, _keepersBytes []byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutMCKeeperPubKeybytes(&_EthCrossChainData.TransactOpts, _height, _keepersBytes)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthCrossChainData *EthCrossChainDataTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthCrossChainData *EthCrossChainDataSession) RenounceOwnership() (*types.Transaction, error) {
	return _EthCrossChainData.Contract.RenounceOwnership(&_EthCrossChainData.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthCrossChainData *EthCrossChainDataTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EthCrossChainData.Contract.RenounceOwnership(&_EthCrossChainData.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthCrossChainData *EthCrossChainDataTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthCrossChainData *EthCrossChainDataSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.TransferOwnership(&_EthCrossChainData.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthCrossChainData *EthCrossChainDataTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.TransferOwnership(&_EthCrossChainData.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) Unpause() (*types.Transaction, error) {
	return _EthCrossChainData.Contract.Unpause(&_EthCrossChainData.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) Unpause() (*types.Transaction, error) {
	return _EthCrossChainData.Contract.Unpause(&_EthCrossChainData.TransactOpts)
}

// EthCrossChainDataOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EthCrossChainData contract.
type EthCrossChainDataOwnershipTransferredIterator struct {
	Event *EthCrossChainDataOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EthCrossChainDataOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainDataOwnershipTransferred)
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
		it.Event = new(EthCrossChainDataOwnershipTransferred)
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
func (it *EthCrossChainDataOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainDataOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainDataOwnershipTransferred represents a OwnershipTransferred event raised by the EthCrossChainData contract.
type EthCrossChainDataOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EthCrossChainData *EthCrossChainDataFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EthCrossChainDataOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EthCrossChainData.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainDataOwnershipTransferredIterator{contract: _EthCrossChainData.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EthCrossChainData *EthCrossChainDataFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EthCrossChainDataOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EthCrossChainData.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainDataOwnershipTransferred)
				if err := _EthCrossChainData.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EthCrossChainData *EthCrossChainDataFilterer) ParseOwnershipTransferred(log types.Log) (*EthCrossChainDataOwnershipTransferred, error) {
	event := new(EthCrossChainDataOwnershipTransferred)
	if err := _EthCrossChainData.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainDataPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EthCrossChainData contract.
type EthCrossChainDataPausedIterator struct {
	Event *EthCrossChainDataPaused // Event containing the contract specifics and raw log

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
func (it *EthCrossChainDataPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainDataPaused)
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
		it.Event = new(EthCrossChainDataPaused)
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
func (it *EthCrossChainDataPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainDataPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainDataPaused represents a Paused event raised by the EthCrossChainData contract.
type EthCrossChainDataPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EthCrossChainData *EthCrossChainDataFilterer) FilterPaused(opts *bind.FilterOpts) (*EthCrossChainDataPausedIterator, error) {

	logs, sub, err := _EthCrossChainData.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainDataPausedIterator{contract: _EthCrossChainData.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EthCrossChainData *EthCrossChainDataFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EthCrossChainDataPaused) (event.Subscription, error) {

	logs, sub, err := _EthCrossChainData.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainDataPaused)
				if err := _EthCrossChainData.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EthCrossChainData *EthCrossChainDataFilterer) ParsePaused(log types.Log) (*EthCrossChainDataPaused, error) {
	event := new(EthCrossChainDataPaused)
	if err := _EthCrossChainData.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainDataUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EthCrossChainData contract.
type EthCrossChainDataUnpausedIterator struct {
	Event *EthCrossChainDataUnpaused // Event containing the contract specifics and raw log

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
func (it *EthCrossChainDataUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainDataUnpaused)
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
		it.Event = new(EthCrossChainDataUnpaused)
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
func (it *EthCrossChainDataUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainDataUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainDataUnpaused represents a Unpaused event raised by the EthCrossChainData contract.
type EthCrossChainDataUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EthCrossChainData *EthCrossChainDataFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EthCrossChainDataUnpausedIterator, error) {

	logs, sub, err := _EthCrossChainData.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainDataUnpausedIterator{contract: _EthCrossChainData.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EthCrossChainData *EthCrossChainDataFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EthCrossChainDataUnpaused) (event.Subscription, error) {

	logs, sub, err := _EthCrossChainData.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainDataUnpaused)
				if err := _EthCrossChainData.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EthCrossChainData *EthCrossChainDataFilterer) ParseUnpaused(log types.Log) (*EthCrossChainDataUnpaused, error) {
	event := new(EthCrossChainDataUnpaused)
	if err := _EthCrossChainData.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IEthCrossChainDataABI is the input ABI used to generate the binding from.
const IEthCrossChainDataABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_crossChainTx\",\"type\":\"bytes32\"}],\"name\":\"getCrossChainTxExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethTxHashIndex\",\"type\":\"uint256\"}],\"name\":\"getEthTxHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEthTxHashIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key2\",\"type\":\"bytes32\"}],\"name\":\"getExtraData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLatestBookKeeperHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLatestHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_height\",\"type\":\"uint64\"}],\"name\":\"getMCHeaderBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMCKeeperHeight\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"getMCKeeperPubBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_height\",\"type\":\"uint64\"}],\"name\":\"getMCKeeperPubKeybytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initGenesisBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isGenesisBlockInited\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_crossChainTx\",\"type\":\"bytes32\"}],\"name\":\"putCrossChainTxExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ethTxHash\",\"type\":\"bytes32\"}],\"name\":\"putEthTxHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"putExtraData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_latestBookKeeperHeight\",\"type\":\"uint64\"}],\"name\":\"putLatestBookKeeperHeight\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_latestHeight\",\"type\":\"uint64\"}],\"name\":\"putLatestHeight\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_rawHeader\",\"type\":\"bytes\"}],\"name\":\"putMCHeaderBytes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"putMCKeeperHeight\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_MCKeeperPubBytes\",\"type\":\"bytes\"}],\"name\":\"putMCKeeperPubBytes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_keepersBytes\",\"type\":\"bytes\"}],\"name\":\"putMCKeeperPubKeybytes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IEthCrossChainDataFuncSigs maps the 4-byte function signature to its string representation.
var IEthCrossChainDataFuncSigs = map[string]string{
	"cceb218d": "getCrossChainTxExist(bytes32)",
	"29927875": "getEthTxHash(uint256)",
	"ff3d24a1": "getEthTxHashIndex()",
	"40602bb5": "getExtraData(bytes32,bytes32)",
	"71a77750": "getLatestBookKeeperHeight()",
	"4ed1d8cc": "getLatestHeight()",
	"9743e9d6": "getMCHeaderBytes(uint64)",
	"f0b1fcfe": "getMCKeeperHeight()",
	"7d737b8b": "getMCKeeperPubBytes(uint64)",
	"09705412": "getMCKeeperPubKeybytes(uint64)",
	"8b39aa25": "initGenesisBlock()",
	"64f1acb0": "isGenesisBlockInited()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"794b0300": "putCrossChainTxExist(bytes32)",
	"4c3ccf64": "putEthTxHash(bytes32)",
	"1afe374e": "putExtraData(bytes32,bytes32,bytes)",
	"a6c6fe7c": "putLatestBookKeeperHeight(uint64)",
	"6d044440": "putLatestHeight(uint64)",
	"021d70ab": "putMCHeaderBytes(uint64,bytes)",
	"c6dbd0b7": "putMCKeeperHeight(uint64)",
	"191b0ab1": "putMCKeeperPubBytes(uint64,bytes)",
	"73c2f5f6": "putMCKeeperPubKeybytes(uint64,bytes)",
	"f2fde38b": "transferOwnership(address)",
	"3f4ba83a": "unpause()",
}

// IEthCrossChainData is an auto generated Go binding around an Ethereum contract.
type IEthCrossChainData struct {
	IEthCrossChainDataCaller     // Read-only binding to the contract
	IEthCrossChainDataTransactor // Write-only binding to the contract
	IEthCrossChainDataFilterer   // Log filterer for contract events
}

// IEthCrossChainDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEthCrossChainDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEthCrossChainDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEthCrossChainDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEthCrossChainDataSession struct {
	Contract     *IEthCrossChainData // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IEthCrossChainDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEthCrossChainDataCallerSession struct {
	Contract *IEthCrossChainDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IEthCrossChainDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEthCrossChainDataTransactorSession struct {
	Contract     *IEthCrossChainDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IEthCrossChainDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEthCrossChainDataRaw struct {
	Contract *IEthCrossChainData // Generic contract binding to access the raw methods on
}

// IEthCrossChainDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEthCrossChainDataCallerRaw struct {
	Contract *IEthCrossChainDataCaller // Generic read-only contract binding to access the raw methods on
}

// IEthCrossChainDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEthCrossChainDataTransactorRaw struct {
	Contract *IEthCrossChainDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEthCrossChainData creates a new instance of IEthCrossChainData, bound to a specific deployed contract.
func NewIEthCrossChainData(address common.Address, backend bind.ContractBackend) (*IEthCrossChainData, error) {
	contract, err := bindIEthCrossChainData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainData{IEthCrossChainDataCaller: IEthCrossChainDataCaller{contract: contract}, IEthCrossChainDataTransactor: IEthCrossChainDataTransactor{contract: contract}, IEthCrossChainDataFilterer: IEthCrossChainDataFilterer{contract: contract}}, nil
}

// NewIEthCrossChainDataCaller creates a new read-only instance of IEthCrossChainData, bound to a specific deployed contract.
func NewIEthCrossChainDataCaller(address common.Address, caller bind.ContractCaller) (*IEthCrossChainDataCaller, error) {
	contract, err := bindIEthCrossChainData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainDataCaller{contract: contract}, nil
}

// NewIEthCrossChainDataTransactor creates a new write-only instance of IEthCrossChainData, bound to a specific deployed contract.
func NewIEthCrossChainDataTransactor(address common.Address, transactor bind.ContractTransactor) (*IEthCrossChainDataTransactor, error) {
	contract, err := bindIEthCrossChainData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainDataTransactor{contract: contract}, nil
}

// NewIEthCrossChainDataFilterer creates a new log filterer instance of IEthCrossChainData, bound to a specific deployed contract.
func NewIEthCrossChainDataFilterer(address common.Address, filterer bind.ContractFilterer) (*IEthCrossChainDataFilterer, error) {
	contract, err := bindIEthCrossChainData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainDataFilterer{contract: contract}, nil
}

// bindIEthCrossChainData binds a generic wrapper to an already deployed contract.
func bindIEthCrossChainData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IEthCrossChainDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainData *IEthCrossChainDataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainData.Contract.IEthCrossChainDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainData *IEthCrossChainDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.IEthCrossChainDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainData *IEthCrossChainDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.IEthCrossChainDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainData *IEthCrossChainDataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainData *IEthCrossChainDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainData *IEthCrossChainDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.contract.Transact(opts, method, params...)
}

// GetCrossChainTxExist is a free data retrieval call binding the contract method 0xcceb218d.
//
// Solidity: function getCrossChainTxExist(bytes32 _crossChainTx) view returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetCrossChainTxExist(opts *bind.CallOpts, _crossChainTx [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getCrossChainTxExist", _crossChainTx)
	return *ret0, err
}

// GetCrossChainTxExist is a free data retrieval call binding the contract method 0xcceb218d.
//
// Solidity: function getCrossChainTxExist(bytes32 _crossChainTx) view returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetCrossChainTxExist(_crossChainTx [32]byte) (bool, error) {
	return _IEthCrossChainData.Contract.GetCrossChainTxExist(&_IEthCrossChainData.CallOpts, _crossChainTx)
}

// GetCrossChainTxExist is a free data retrieval call binding the contract method 0xcceb218d.
//
// Solidity: function getCrossChainTxExist(bytes32 _crossChainTx) view returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetCrossChainTxExist(_crossChainTx [32]byte) (bool, error) {
	return _IEthCrossChainData.Contract.GetCrossChainTxExist(&_IEthCrossChainData.CallOpts, _crossChainTx)
}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 _ethTxHashIndex) view returns(bytes32)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetEthTxHash(opts *bind.CallOpts, _ethTxHashIndex *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getEthTxHash", _ethTxHashIndex)
	return *ret0, err
}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 _ethTxHashIndex) view returns(bytes32)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetEthTxHash(_ethTxHashIndex *big.Int) ([32]byte, error) {
	return _IEthCrossChainData.Contract.GetEthTxHash(&_IEthCrossChainData.CallOpts, _ethTxHashIndex)
}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 _ethTxHashIndex) view returns(bytes32)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetEthTxHash(_ethTxHashIndex *big.Int) ([32]byte, error) {
	return _IEthCrossChainData.Contract.GetEthTxHash(&_IEthCrossChainData.CallOpts, _ethTxHashIndex)
}

// GetEthTxHashIndex is a free data retrieval call binding the contract method 0xff3d24a1.
//
// Solidity: function getEthTxHashIndex() view returns(uint256)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetEthTxHashIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getEthTxHashIndex")
	return *ret0, err
}

// GetEthTxHashIndex is a free data retrieval call binding the contract method 0xff3d24a1.
//
// Solidity: function getEthTxHashIndex() view returns(uint256)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetEthTxHashIndex() (*big.Int, error) {
	return _IEthCrossChainData.Contract.GetEthTxHashIndex(&_IEthCrossChainData.CallOpts)
}

// GetEthTxHashIndex is a free data retrieval call binding the contract method 0xff3d24a1.
//
// Solidity: function getEthTxHashIndex() view returns(uint256)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetEthTxHashIndex() (*big.Int, error) {
	return _IEthCrossChainData.Contract.GetEthTxHashIndex(&_IEthCrossChainData.CallOpts)
}

// GetExtraData is a free data retrieval call binding the contract method 0x40602bb5.
//
// Solidity: function getExtraData(bytes32 key1, bytes32 key2) view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetExtraData(opts *bind.CallOpts, key1 [32]byte, key2 [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getExtraData", key1, key2)
	return *ret0, err
}

// GetExtraData is a free data retrieval call binding the contract method 0x40602bb5.
//
// Solidity: function getExtraData(bytes32 key1, bytes32 key2) view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetExtraData(key1 [32]byte, key2 [32]byte) ([]byte, error) {
	return _IEthCrossChainData.Contract.GetExtraData(&_IEthCrossChainData.CallOpts, key1, key2)
}

// GetExtraData is a free data retrieval call binding the contract method 0x40602bb5.
//
// Solidity: function getExtraData(bytes32 key1, bytes32 key2) view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetExtraData(key1 [32]byte, key2 [32]byte) ([]byte, error) {
	return _IEthCrossChainData.Contract.GetExtraData(&_IEthCrossChainData.CallOpts, key1, key2)
}

// GetLatestBookKeeperHeight is a free data retrieval call binding the contract method 0x71a77750.
//
// Solidity: function getLatestBookKeeperHeight() view returns(uint64)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetLatestBookKeeperHeight(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getLatestBookKeeperHeight")
	return *ret0, err
}

// GetLatestBookKeeperHeight is a free data retrieval call binding the contract method 0x71a77750.
//
// Solidity: function getLatestBookKeeperHeight() view returns(uint64)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetLatestBookKeeperHeight() (uint64, error) {
	return _IEthCrossChainData.Contract.GetLatestBookKeeperHeight(&_IEthCrossChainData.CallOpts)
}

// GetLatestBookKeeperHeight is a free data retrieval call binding the contract method 0x71a77750.
//
// Solidity: function getLatestBookKeeperHeight() view returns(uint64)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetLatestBookKeeperHeight() (uint64, error) {
	return _IEthCrossChainData.Contract.GetLatestBookKeeperHeight(&_IEthCrossChainData.CallOpts)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x4ed1d8cc.
//
// Solidity: function getLatestHeight() view returns(uint64)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetLatestHeight(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getLatestHeight")
	return *ret0, err
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x4ed1d8cc.
//
// Solidity: function getLatestHeight() view returns(uint64)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetLatestHeight() (uint64, error) {
	return _IEthCrossChainData.Contract.GetLatestHeight(&_IEthCrossChainData.CallOpts)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x4ed1d8cc.
//
// Solidity: function getLatestHeight() view returns(uint64)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetLatestHeight() (uint64, error) {
	return _IEthCrossChainData.Contract.GetLatestHeight(&_IEthCrossChainData.CallOpts)
}

// GetMCHeaderBytes is a free data retrieval call binding the contract method 0x9743e9d6.
//
// Solidity: function getMCHeaderBytes(uint64 _height) view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetMCHeaderBytes(opts *bind.CallOpts, _height uint64) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getMCHeaderBytes", _height)
	return *ret0, err
}

// GetMCHeaderBytes is a free data retrieval call binding the contract method 0x9743e9d6.
//
// Solidity: function getMCHeaderBytes(uint64 _height) view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetMCHeaderBytes(_height uint64) ([]byte, error) {
	return _IEthCrossChainData.Contract.GetMCHeaderBytes(&_IEthCrossChainData.CallOpts, _height)
}

// GetMCHeaderBytes is a free data retrieval call binding the contract method 0x9743e9d6.
//
// Solidity: function getMCHeaderBytes(uint64 _height) view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetMCHeaderBytes(_height uint64) ([]byte, error) {
	return _IEthCrossChainData.Contract.GetMCHeaderBytes(&_IEthCrossChainData.CallOpts, _height)
}

// GetMCKeeperHeight is a free data retrieval call binding the contract method 0xf0b1fcfe.
//
// Solidity: function getMCKeeperHeight() view returns(uint64[])
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetMCKeeperHeight(opts *bind.CallOpts) ([]uint64, error) {
	var (
		ret0 = new([]uint64)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getMCKeeperHeight")
	return *ret0, err
}

// GetMCKeeperHeight is a free data retrieval call binding the contract method 0xf0b1fcfe.
//
// Solidity: function getMCKeeperHeight() view returns(uint64[])
func (_IEthCrossChainData *IEthCrossChainDataSession) GetMCKeeperHeight() ([]uint64, error) {
	return _IEthCrossChainData.Contract.GetMCKeeperHeight(&_IEthCrossChainData.CallOpts)
}

// GetMCKeeperHeight is a free data retrieval call binding the contract method 0xf0b1fcfe.
//
// Solidity: function getMCKeeperHeight() view returns(uint64[])
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetMCKeeperHeight() ([]uint64, error) {
	return _IEthCrossChainData.Contract.GetMCKeeperHeight(&_IEthCrossChainData.CallOpts)
}

// GetMCKeeperPubBytes is a free data retrieval call binding the contract method 0x7d737b8b.
//
// Solidity: function getMCKeeperPubBytes(uint64 height) view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetMCKeeperPubBytes(opts *bind.CallOpts, height uint64) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getMCKeeperPubBytes", height)
	return *ret0, err
}

// GetMCKeeperPubBytes is a free data retrieval call binding the contract method 0x7d737b8b.
//
// Solidity: function getMCKeeperPubBytes(uint64 height) view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetMCKeeperPubBytes(height uint64) ([]byte, error) {
	return _IEthCrossChainData.Contract.GetMCKeeperPubBytes(&_IEthCrossChainData.CallOpts, height)
}

// GetMCKeeperPubBytes is a free data retrieval call binding the contract method 0x7d737b8b.
//
// Solidity: function getMCKeeperPubBytes(uint64 height) view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetMCKeeperPubBytes(height uint64) ([]byte, error) {
	return _IEthCrossChainData.Contract.GetMCKeeperPubBytes(&_IEthCrossChainData.CallOpts, height)
}

// GetMCKeeperPubKeybytes is a free data retrieval call binding the contract method 0x09705412.
//
// Solidity: function getMCKeeperPubKeybytes(uint64 _height) view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetMCKeeperPubKeybytes(opts *bind.CallOpts, _height uint64) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getMCKeeperPubKeybytes", _height)
	return *ret0, err
}

// GetMCKeeperPubKeybytes is a free data retrieval call binding the contract method 0x09705412.
//
// Solidity: function getMCKeeperPubKeybytes(uint64 _height) view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetMCKeeperPubKeybytes(_height uint64) ([]byte, error) {
	return _IEthCrossChainData.Contract.GetMCKeeperPubKeybytes(&_IEthCrossChainData.CallOpts, _height)
}

// GetMCKeeperPubKeybytes is a free data retrieval call binding the contract method 0x09705412.
//
// Solidity: function getMCKeeperPubKeybytes(uint64 _height) view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetMCKeeperPubKeybytes(_height uint64) ([]byte, error) {
	return _IEthCrossChainData.Contract.GetMCKeeperPubKeybytes(&_IEthCrossChainData.CallOpts, _height)
}

// IsGenesisBlockInited is a free data retrieval call binding the contract method 0x64f1acb0.
//
// Solidity: function isGenesisBlockInited() view returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataCaller) IsGenesisBlockInited(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "isGenesisBlockInited")
	return *ret0, err
}

// IsGenesisBlockInited is a free data retrieval call binding the contract method 0x64f1acb0.
//
// Solidity: function isGenesisBlockInited() view returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) IsGenesisBlockInited() (bool, error) {
	return _IEthCrossChainData.Contract.IsGenesisBlockInited(&_IEthCrossChainData.CallOpts)
}

// IsGenesisBlockInited is a free data retrieval call binding the contract method 0x64f1acb0.
//
// Solidity: function isGenesisBlockInited() view returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) IsGenesisBlockInited() (bool, error) {
	return _IEthCrossChainData.Contract.IsGenesisBlockInited(&_IEthCrossChainData.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) Paused() (bool, error) {
	return _IEthCrossChainData.Contract.Paused(&_IEthCrossChainData.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) Paused() (bool, error) {
	return _IEthCrossChainData.Contract.Paused(&_IEthCrossChainData.CallOpts)
}

// InitGenesisBlock is a paid mutator transaction binding the contract method 0x8b39aa25.
//
// Solidity: function initGenesisBlock() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) InitGenesisBlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "initGenesisBlock")
}

// InitGenesisBlock is a paid mutator transaction binding the contract method 0x8b39aa25.
//
// Solidity: function initGenesisBlock() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) InitGenesisBlock() (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.InitGenesisBlock(&_IEthCrossChainData.TransactOpts)
}

// InitGenesisBlock is a paid mutator transaction binding the contract method 0x8b39aa25.
//
// Solidity: function initGenesisBlock() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) InitGenesisBlock() (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.InitGenesisBlock(&_IEthCrossChainData.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) Pause() (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.Pause(&_IEthCrossChainData.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) Pause() (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.Pause(&_IEthCrossChainData.TransactOpts)
}

// PutCrossChainTxExist is a paid mutator transaction binding the contract method 0x794b0300.
//
// Solidity: function putCrossChainTxExist(bytes32 _crossChainTx) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutCrossChainTxExist(opts *bind.TransactOpts, _crossChainTx [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putCrossChainTxExist", _crossChainTx)
}

// PutCrossChainTxExist is a paid mutator transaction binding the contract method 0x794b0300.
//
// Solidity: function putCrossChainTxExist(bytes32 _crossChainTx) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutCrossChainTxExist(_crossChainTx [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutCrossChainTxExist(&_IEthCrossChainData.TransactOpts, _crossChainTx)
}

// PutCrossChainTxExist is a paid mutator transaction binding the contract method 0x794b0300.
//
// Solidity: function putCrossChainTxExist(bytes32 _crossChainTx) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutCrossChainTxExist(_crossChainTx [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutCrossChainTxExist(&_IEthCrossChainData.TransactOpts, _crossChainTx)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x4c3ccf64.
//
// Solidity: function putEthTxHash(bytes32 _ethTxHash) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutEthTxHash(opts *bind.TransactOpts, _ethTxHash [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putEthTxHash", _ethTxHash)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x4c3ccf64.
//
// Solidity: function putEthTxHash(bytes32 _ethTxHash) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutEthTxHash(_ethTxHash [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutEthTxHash(&_IEthCrossChainData.TransactOpts, _ethTxHash)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x4c3ccf64.
//
// Solidity: function putEthTxHash(bytes32 _ethTxHash) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutEthTxHash(_ethTxHash [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutEthTxHash(&_IEthCrossChainData.TransactOpts, _ethTxHash)
}

// PutExtraData is a paid mutator transaction binding the contract method 0x1afe374e.
//
// Solidity: function putExtraData(bytes32 key1, bytes32 key2, bytes value) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutExtraData(opts *bind.TransactOpts, key1 [32]byte, key2 [32]byte, value []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putExtraData", key1, key2, value)
}

// PutExtraData is a paid mutator transaction binding the contract method 0x1afe374e.
//
// Solidity: function putExtraData(bytes32 key1, bytes32 key2, bytes value) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutExtraData(key1 [32]byte, key2 [32]byte, value []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutExtraData(&_IEthCrossChainData.TransactOpts, key1, key2, value)
}

// PutExtraData is a paid mutator transaction binding the contract method 0x1afe374e.
//
// Solidity: function putExtraData(bytes32 key1, bytes32 key2, bytes value) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutExtraData(key1 [32]byte, key2 [32]byte, value []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutExtraData(&_IEthCrossChainData.TransactOpts, key1, key2, value)
}

// PutLatestBookKeeperHeight is a paid mutator transaction binding the contract method 0xa6c6fe7c.
//
// Solidity: function putLatestBookKeeperHeight(uint64 _latestBookKeeperHeight) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutLatestBookKeeperHeight(opts *bind.TransactOpts, _latestBookKeeperHeight uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putLatestBookKeeperHeight", _latestBookKeeperHeight)
}

// PutLatestBookKeeperHeight is a paid mutator transaction binding the contract method 0xa6c6fe7c.
//
// Solidity: function putLatestBookKeeperHeight(uint64 _latestBookKeeperHeight) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutLatestBookKeeperHeight(_latestBookKeeperHeight uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutLatestBookKeeperHeight(&_IEthCrossChainData.TransactOpts, _latestBookKeeperHeight)
}

// PutLatestBookKeeperHeight is a paid mutator transaction binding the contract method 0xa6c6fe7c.
//
// Solidity: function putLatestBookKeeperHeight(uint64 _latestBookKeeperHeight) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutLatestBookKeeperHeight(_latestBookKeeperHeight uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutLatestBookKeeperHeight(&_IEthCrossChainData.TransactOpts, _latestBookKeeperHeight)
}

// PutLatestHeight is a paid mutator transaction binding the contract method 0x6d044440.
//
// Solidity: function putLatestHeight(uint64 _latestHeight) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutLatestHeight(opts *bind.TransactOpts, _latestHeight uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putLatestHeight", _latestHeight)
}

// PutLatestHeight is a paid mutator transaction binding the contract method 0x6d044440.
//
// Solidity: function putLatestHeight(uint64 _latestHeight) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutLatestHeight(_latestHeight uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutLatestHeight(&_IEthCrossChainData.TransactOpts, _latestHeight)
}

// PutLatestHeight is a paid mutator transaction binding the contract method 0x6d044440.
//
// Solidity: function putLatestHeight(uint64 _latestHeight) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutLatestHeight(_latestHeight uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutLatestHeight(&_IEthCrossChainData.TransactOpts, _latestHeight)
}

// PutMCHeaderBytes is a paid mutator transaction binding the contract method 0x021d70ab.
//
// Solidity: function putMCHeaderBytes(uint64 _height, bytes _rawHeader) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutMCHeaderBytes(opts *bind.TransactOpts, _height uint64, _rawHeader []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putMCHeaderBytes", _height, _rawHeader)
}

// PutMCHeaderBytes is a paid mutator transaction binding the contract method 0x021d70ab.
//
// Solidity: function putMCHeaderBytes(uint64 _height, bytes _rawHeader) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutMCHeaderBytes(_height uint64, _rawHeader []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutMCHeaderBytes(&_IEthCrossChainData.TransactOpts, _height, _rawHeader)
}

// PutMCHeaderBytes is a paid mutator transaction binding the contract method 0x021d70ab.
//
// Solidity: function putMCHeaderBytes(uint64 _height, bytes _rawHeader) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutMCHeaderBytes(_height uint64, _rawHeader []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutMCHeaderBytes(&_IEthCrossChainData.TransactOpts, _height, _rawHeader)
}

// PutMCKeeperHeight is a paid mutator transaction binding the contract method 0xc6dbd0b7.
//
// Solidity: function putMCKeeperHeight(uint64 height) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutMCKeeperHeight(opts *bind.TransactOpts, height uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putMCKeeperHeight", height)
}

// PutMCKeeperHeight is a paid mutator transaction binding the contract method 0xc6dbd0b7.
//
// Solidity: function putMCKeeperHeight(uint64 height) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutMCKeeperHeight(height uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutMCKeeperHeight(&_IEthCrossChainData.TransactOpts, height)
}

// PutMCKeeperHeight is a paid mutator transaction binding the contract method 0xc6dbd0b7.
//
// Solidity: function putMCKeeperHeight(uint64 height) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutMCKeeperHeight(height uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutMCKeeperHeight(&_IEthCrossChainData.TransactOpts, height)
}

// PutMCKeeperPubBytes is a paid mutator transaction binding the contract method 0x191b0ab1.
//
// Solidity: function putMCKeeperPubBytes(uint64 height, bytes _MCKeeperPubBytes) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutMCKeeperPubBytes(opts *bind.TransactOpts, height uint64, _MCKeeperPubBytes []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putMCKeeperPubBytes", height, _MCKeeperPubBytes)
}

// PutMCKeeperPubBytes is a paid mutator transaction binding the contract method 0x191b0ab1.
//
// Solidity: function putMCKeeperPubBytes(uint64 height, bytes _MCKeeperPubBytes) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutMCKeeperPubBytes(height uint64, _MCKeeperPubBytes []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutMCKeeperPubBytes(&_IEthCrossChainData.TransactOpts, height, _MCKeeperPubBytes)
}

// PutMCKeeperPubBytes is a paid mutator transaction binding the contract method 0x191b0ab1.
//
// Solidity: function putMCKeeperPubBytes(uint64 height, bytes _MCKeeperPubBytes) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutMCKeeperPubBytes(height uint64, _MCKeeperPubBytes []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutMCKeeperPubBytes(&_IEthCrossChainData.TransactOpts, height, _MCKeeperPubBytes)
}

// PutMCKeeperPubKeybytes is a paid mutator transaction binding the contract method 0x73c2f5f6.
//
// Solidity: function putMCKeeperPubKeybytes(uint64 _height, bytes _keepersBytes) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutMCKeeperPubKeybytes(opts *bind.TransactOpts, _height uint64, _keepersBytes []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putMCKeeperPubKeybytes", _height, _keepersBytes)
}

// PutMCKeeperPubKeybytes is a paid mutator transaction binding the contract method 0x73c2f5f6.
//
// Solidity: function putMCKeeperPubKeybytes(uint64 _height, bytes _keepersBytes) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutMCKeeperPubKeybytes(_height uint64, _keepersBytes []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutMCKeeperPubKeybytes(&_IEthCrossChainData.TransactOpts, _height, _keepersBytes)
}

// PutMCKeeperPubKeybytes is a paid mutator transaction binding the contract method 0x73c2f5f6.
//
// Solidity: function putMCKeeperPubKeybytes(uint64 _height, bytes _keepersBytes) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutMCKeeperPubKeybytes(_height uint64, _keepersBytes []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutMCKeeperPubKeybytes(&_IEthCrossChainData.TransactOpts, _height, _keepersBytes)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IEthCrossChainData *IEthCrossChainDataTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IEthCrossChainData *IEthCrossChainDataSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.TransferOwnership(&_IEthCrossChainData.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.TransferOwnership(&_IEthCrossChainData.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) Unpause() (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.Unpause(&_IEthCrossChainData.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) Unpause() (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.Unpause(&_IEthCrossChainData.TransactOpts)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Ownable *OwnableCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Ownable *OwnableSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Ownable *OwnableCallerSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PausableABI is the input ABI used to generate the binding from.
const PausableABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PausableFuncSigs maps the 4-byte function signature to its string representation.
var PausableFuncSigs = map[string]string{
	"5c975abb": "paused()",
}

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// PausablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Pausable contract.
type PausablePausedIterator struct {
	Event *PausablePaused // Event containing the contract specifics and raw log

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
func (it *PausablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePaused)
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
		it.Event = new(PausablePaused)
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
func (it *PausablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePaused represents a Paused event raised by the Pausable contract.
type PausablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) FilterPaused(opts *bind.FilterOpts) (*PausablePausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PausablePausedIterator{contract: _Pausable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PausablePaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePaused)
				if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) ParsePaused(log types.Log) (*PausablePaused, error) {
	event := new(PausablePaused)
	if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PausableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Pausable contract.
type PausableUnpausedIterator struct {
	Event *PausableUnpaused // Event containing the contract specifics and raw log

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
func (it *PausableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpaused)
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
		it.Event = new(PausableUnpaused)
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
func (it *PausableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpaused represents a Unpaused event raised by the Pausable contract.
type PausableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PausableUnpausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PausableUnpausedIterator{contract: _Pausable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PausableUnpaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpaused)
				if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) ParseUnpaused(log types.Log) (*PausableUnpaused, error) {
	event := new(PausableUnpaused)
	if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}


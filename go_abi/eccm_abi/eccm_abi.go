// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eccm_abi

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

// ECCUtilsABI is the input ABI used to generate the binding from.
const ECCUtilsABI = "[]"

// ECCUtilsBin is the compiled bytecode used for deploying new contracts.
var ECCUtilsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820d5215014dd289390811fa8fa3a975f5906f395be805770cfcc6faaa7f999e21064736f6c63430005110032"

// DeployECCUtils deploys a new Ethereum contract, binding an instance of ECCUtils to it.
func DeployECCUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECCUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(ECCUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ECCUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECCUtils{ECCUtilsCaller: ECCUtilsCaller{contract: contract}, ECCUtilsTransactor: ECCUtilsTransactor{contract: contract}, ECCUtilsFilterer: ECCUtilsFilterer{contract: contract}}, nil
}

// ECCUtils is an auto generated Go binding around an Ethereum contract.
type ECCUtils struct {
	ECCUtilsCaller     // Read-only binding to the contract
	ECCUtilsTransactor // Write-only binding to the contract
	ECCUtilsFilterer   // Log filterer for contract events
}

// ECCUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECCUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECCUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECCUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECCUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECCUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECCUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECCUtilsSession struct {
	Contract     *ECCUtils         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECCUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECCUtilsCallerSession struct {
	Contract *ECCUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ECCUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECCUtilsTransactorSession struct {
	Contract     *ECCUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ECCUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECCUtilsRaw struct {
	Contract *ECCUtils // Generic contract binding to access the raw methods on
}

// ECCUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECCUtilsCallerRaw struct {
	Contract *ECCUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// ECCUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECCUtilsTransactorRaw struct {
	Contract *ECCUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECCUtils creates a new instance of ECCUtils, bound to a specific deployed contract.
func NewECCUtils(address common.Address, backend bind.ContractBackend) (*ECCUtils, error) {
	contract, err := bindECCUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECCUtils{ECCUtilsCaller: ECCUtilsCaller{contract: contract}, ECCUtilsTransactor: ECCUtilsTransactor{contract: contract}, ECCUtilsFilterer: ECCUtilsFilterer{contract: contract}}, nil
}

// NewECCUtilsCaller creates a new read-only instance of ECCUtils, bound to a specific deployed contract.
func NewECCUtilsCaller(address common.Address, caller bind.ContractCaller) (*ECCUtilsCaller, error) {
	contract, err := bindECCUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECCUtilsCaller{contract: contract}, nil
}

// NewECCUtilsTransactor creates a new write-only instance of ECCUtils, bound to a specific deployed contract.
func NewECCUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*ECCUtilsTransactor, error) {
	contract, err := bindECCUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECCUtilsTransactor{contract: contract}, nil
}

// NewECCUtilsFilterer creates a new log filterer instance of ECCUtils, bound to a specific deployed contract.
func NewECCUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*ECCUtilsFilterer, error) {
	contract, err := bindECCUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECCUtilsFilterer{contract: contract}, nil
}

// bindECCUtils binds a generic wrapper to an already deployed contract.
func bindECCUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECCUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECCUtils *ECCUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECCUtils.Contract.ECCUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECCUtils *ECCUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECCUtils.Contract.ECCUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECCUtils *ECCUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECCUtils.Contract.ECCUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECCUtils *ECCUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECCUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECCUtils *ECCUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECCUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECCUtils *ECCUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECCUtils.Contract.contract.Transact(opts, method, params...)
}

// EthCrossChainManagerABI is the input ABI used to generate the binding from.
const EthCrossChainManagerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_eccd\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_chainId\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"fromContractWhiteList\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"toContractWhiteList\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"methodWhiteList\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"curEpochPkBytes\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rawHeader\",\"type\":\"bytes\"}],\"name\":\"ChangeBookKeeperEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"txId\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxyOrAssetContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toContract\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rawdata\",\"type\":\"bytes\"}],\"name\":\"CrossChainEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rawHeader\",\"type\":\"bytes\"}],\"name\":\"InitGenesisBlockEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fromChainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toContract\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"crossChainTxHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fromChainTxHash\",\"type\":\"bytes\"}],\"name\":\"VerifyHeaderAndExecuteTxEvent\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"EthCrossChainDataAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rawHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyList\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sigList\",\"type\":\"bytes\"}],\"name\":\"changeBookKeeper\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"toContract\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"method\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"txData\",\"type\":\"bytes\"}],\"name\":\"crossChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rawHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyList\",\"type\":\"bytes\"}],\"name\":\"initGenesisBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"EpochPkBytes\",\"type\":\"bytes\"}],\"name\":\"recoverEpochPk\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_newChainId\",\"type\":\"uint64\"}],\"name\":\"setChainId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"unsetEpochPkBytes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newEthCrossChainManagerAddress\",\"type\":\"address\"}],\"name\":\"upgradeToNew\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rawHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"headerProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"curRawHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"headerSig\",\"type\":\"bytes\"}],\"name\":\"verifyHeaderAndExecuteTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteListFromContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"whiteListMethod\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteListToContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EthCrossChainManagerFuncSigs maps the 4-byte function signature to its string representation.
var EthCrossChainManagerFuncSigs = map[string]string{
	"00ba1694": "EthCrossChainDataAddress()",
	"9a8a0592": "chainId()",
	"29dcf4ab": "changeBookKeeper(bytes,bytes,bytes)",
	"bd5cf625": "crossChain(uint64,bytes,bytes,bytes)",
	"34a773eb": "initGenesisBlock(bytes,bytes)",
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"7bcf7dc1": "recoverEpochPk(bytes)",
	"715018a6": "renounceOwnership()",
	"6f31031d": "setChainId(uint64)",
	"f2fde38b": "transferOwnership(address)",
	"3f4ba83a": "unpause()",
	"8c5aea41": "unsetEpochPkBytes(bytes)",
	"7e724ff3": "upgradeToNew(address)",
	"d450e04c": "verifyHeaderAndExecuteTx(bytes,bytes,bytes,bytes,bytes)",
	"73f53ba4": "whiteListFromContract(address)",
	"7b935854": "whiteListMethod(bytes)",
	"b91ca106": "whiteListToContract(address)",
}

// EthCrossChainManagerBin is the compiled bytecode used for deploying new contracts.
var EthCrossChainManagerBin = "0x60806040523480156200001157600080fd5b5060405162004ec938038062004ec98339810160408190526200003491620003a7565b858560006200004b6001600160e01b036200023f16565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506000805460ff60a01b19168155600180546001600160401b03909316600160a01b02600160a01b600160e01b03196001600160a01b039095166001600160a01b03199094169390931793909316919091179091555b845181101562000142576001600260008784815181106200010757fe5b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916911515919091179055600101620000ea565b5060005b83518110156200019e576001600360008684815181106200016357fe5b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff191691151591909117905560010162000146565b5060005b8251811015620001fd5760016004848381518110620001bd57fe5b6020026020010151604051620001d49190620004d0565b908152604051908190036020019020805491151560ff19909216919091179055600101620001a2565b506001600582604051620002129190620004d0565b908152604051908190036020019020805491151560ff1990921691909117905550620005d4945050505050565b3390565b80516200025081620005af565b92915050565b600082601f8301126200026857600080fd5b81516200027f62000279826200050c565b620004e5565b91508181835260208401935060208101905083856020840282011115620002a557600080fd5b60005b83811015620002d55781620002be888262000243565b8452506020928301929190910190600101620002a8565b5050505092915050565b600082601f830112620002f157600080fd5b81516200030262000279826200050c565b81815260209384019390925082018360005b83811015620002d557815186016200032d888262000344565b845250602092830192919091019060010162000314565b600082601f8301126200035657600080fd5b81516200036762000279826200052d565b915080825260208301602083018583830111156200038457600080fd5b620003918382846200057c565b50505092915050565b80516200025081620005c9565b60008060008060008060c08789031215620003c157600080fd5b6000620003cf898962000243565b9650506020620003e289828a016200039a565b95505060408701516001600160401b03811115620003ff57600080fd5b6200040d89828a0162000256565b94505060608701516001600160401b038111156200042a57600080fd5b6200043889828a0162000256565b93505060808701516001600160401b038111156200045557600080fd5b6200046389828a01620002df565b92505060a08701516001600160401b038111156200048057600080fd5b6200048e89828a0162000344565b9150509295509295509295565b6000620004a88262000555565b620004b4818562000559565b9350620004c68185602086016200057c565b9290920192915050565b6000620004de82846200049b565b9392505050565b6040518181016001600160401b03811182821017156200050457600080fd5b604052919050565b60006001600160401b038211156200052357600080fd5b5060209081020190565b60006001600160401b038211156200054457600080fd5b506020601f91909101601f19160190565b5190565b919050565b60006001600160a01b03821662000250565b6001600160401b031690565b60005b83811015620005995781810151838201526020016200057f565b83811115620005a9576000848401525b50505050565b620005ba816200055e565b8114620005c657600080fd5b50565b620005ba8162000570565b6148e580620005e46000396000f3fe608060405234801561001057600080fd5b506004361061012b5760003560e01c80637e724ff3116100ad5780639a8a0592116100715780639a8a059214610225578063b91ca1061461023a578063bd5cf6251461024d578063d450e04c14610260578063f2fde38b146102735761012b565b80637e724ff3146101e75780638456cb59146101fa5780638c5aea41146102025780638da5cb5b146102155780638f32d59b1461021d5761012b565b80636f31031d116100f45780636f31031d14610191578063715018a6146101a457806373f53ba4146101ae5780637b935854146101c15780637bcf7dc1146101d45761012b565b8062ba16941461013057806329dcf4ab1461014e57806334a773eb1461016e5780633f4ba83a146101815780635c975abb14610189575b600080fd5b610138610286565b6040516101459190614268565b60405180910390f35b61016161015c36600461305a565b610295565b6040516101459190614284565b61016161017c366004612ff3565b610663565b610161610916565b610161610a6b565b61016161019f3660046131f0565b610a7b565b6101ac610af5565b005b6101616101bc366004612f31565b610b63565b6101616101cf366004612f8b565b610b78565b6101ac6101e2366004612f8b565b610b98565b6101616101f5366004612f31565b610cb4565b610161610d6f565b610161610210366004612f8b565b610eba565b610138610eda565b610161610ee9565b61022d610f0d565b60405161014591906146c8565b610161610248366004612f31565b610f23565b61016161025b36600461320e565b610f38565b61016161026e3660046130eb565b6112c1565b6101ac610281366004612f31565b611799565b6001546001600160a01b031681565b60008054600160a01b900460ff16156102c95760405162461bcd60e51b81526004016102c09061446a565b60405180910390fd5b6102d1612d5a565b6102da856117c9565b90506000600160009054906101000a90046001600160a01b031690506000816001600160a01b0316635ac407906040518163ffffffff1660e01b815260040160206040518083038186803b15801561033157600080fd5b505afa158015610345573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061036991908101906131d2565b63ffffffff169050806001600160401b0316836060015163ffffffff16116103a35760405162461bcd60e51b81526004016102c0906144aa565b6101408301516bffffffffffffffffffffffff19166103d45760405162461bcd60e51b81526004016102c0906143ba565b6060610453836001600160a01b03166369d480746040518163ffffffff1660e01b815260040160006040518083038186803b15801561041257600080fd5b505afa158015610426573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261044e9190810190612fbf565b6118e8565b805190915061046e898884600360001986015b04850361199b565b61048a5760405162461bcd60e51b81526004016102c09061444a565b600060606104978a611b4f565b91509150816001600160601b0319168761014001516001600160601b031916146104d35760405162461bcd60e51b81526004016102c09061448a565b6060870151604051638a8bd17f60e01b81526001600160a01b03881691638a8bd17f91610503919060040161469a565b602060405180830381600087803b15801561051d57600080fd5b505af1158015610531573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506105559190810190612f4f565b6105715760405162461bcd60e51b81526004016102c09061455a565b856001600160a01b03166341973cd961058983611bcb565b6040518263ffffffff1660e01b81526004016105a591906142d5565b602060405180830381600087803b1580156105bf57600080fd5b505af11580156105d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506105f79190810190612f4f565b6106135760405162461bcd60e51b81526004016102c0906144ca565b7fe60d33488cba3977bf65766cd2f8ac9617f64bf3b3198aff6240ce5c7d43b69087606001518c6040516106489291906146a8565b60405180910390a160019750505050505050505b9392505050565b60008054600160a01b900460ff161561068e5760405162461bcd60e51b81526004016102c09061446a565b60015460408051631a75201d60e21b815290516001600160a01b039092169182916369d48074916004808301926000929190829003018186803b1580156106d457600080fd5b505afa1580156106e8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526107109190810190612fbf565b511561072e5760405162461bcd60e51b81526004016102c0906145ea565b610736612d5a565b61073f856117c9565b90506000606061074e86611b4f565b91509150816001600160601b0319168361014001516001600160601b0319161461078a5760405162461bcd60e51b81526004016102c09061448a565b6060830151604051638a8bd17f60e01b81526001600160a01b03861691638a8bd17f916107ba919060040161469a565b602060405180830381600087803b1580156107d457600080fd5b505af11580156107e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061080c9190810190612f4f565b6108285760405162461bcd60e51b81526004016102c09061441a565b836001600160a01b03166341973cd961084083611bcb565b6040518263ffffffff1660e01b815260040161085c91906142d5565b602060405180830381600087803b15801561087657600080fd5b505af115801561088a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506108ae9190810190612f4f565b6108ca5760405162461bcd60e51b81526004016102c09061460a565b7ff01968fc3a2655cf1b5144cb32de6dc898f91b9239c103744e8457152ab2fbde8360600151886040516108ff9291906146a8565b60405180910390a160019450505050505b92915050565b6000610920610ee9565b61093c5760405162461bcd60e51b81526004016102c0906144fa565b610944610a6b565b1561095157610951611c3a565b60015460408051635c975abb60e01b815290516001600160a01b03909216918291635c975abb916004808301926020929190829003018186803b15801561099757600080fd5b505afa1580156109ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506109cf9190810190612f4f565b15610a6357806001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610a0f57600080fd5b505af1158015610a23573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610a479190810190612f4f565b610a635760405162461bcd60e51b81526004016102c0906143ca565b600191505090565b600054600160a01b900460ff1690565b60008054600160a01b900460ff16610aa55760405162461bcd60e51b81526004016102c09061437a565b610aad610ee9565b610ac95760405162461bcd60e51b81526004016102c0906144fa565b506001805467ffffffffffffffff60a01b1916600160a01b6001600160401b038416021781555b919050565b610afd610ee9565b610b195760405162461bcd60e51b81526004016102c0906144fa565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60026020526000908152604090205460ff1681565b805160208183018101805160048252928201919093012091525460ff1681565b600054600160a01b900460ff16610bc15760405162461bcd60e51b81526004016102c09061437a565b600581604051610bd191906141cc565b9081526040519081900360200190205460ff16610c005760405162461bcd60e51b81526004016102c0906144ea565b6000600582604051610c1291906141cc565b908152604051908190036020018120805492151560ff19909316929092179091556001546341973cd960e01b82526001600160a01b0316906341973cd990610c5e9084906004016142d5565b602060405180830381600087803b158015610c7857600080fd5b505af1158015610c8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610cb09190810190612f4f565b5050565b60008054600160a01b900460ff16610cde5760405162461bcd60e51b81526004016102c09061437a565b610ce6610ee9565b610d025760405162461bcd60e51b81526004016102c0906144fa565b60015460405163f2fde38b60e01b81526001600160a01b0390911690819063f2fde38b90610d34908690600401614268565b600060405180830381600087803b158015610d4e57600080fd5b505af1158015610d62573d6000803e3d6000fd5b5060019695505050505050565b6000610d79610ee9565b610d955760405162461bcd60e51b81526004016102c0906144fa565b610d9d610a6b565b610da957610da9611cb0565b60015460408051635c975abb60e01b815290516001600160a01b03909216918291635c975abb916004808301926020929190829003018186803b158015610def57600080fd5b505afa158015610e03573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610e279190810190612f4f565b610a6357806001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610e6657600080fd5b505af1158015610e7a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610e9e9190810190612f4f565b610a635760405162461bcd60e51b81526004016102c09061439a565b805160208183018101805160058252928201919093012091525460ff1681565b6000546001600160a01b031690565b600080546001600160a01b0316610efe611d12565b6001600160a01b031614905090565b600154600160a01b90046001600160401b031681565b60036020526000908152604090205460ff1681565b60008054600160a01b900460ff1615610f635760405162461bcd60e51b81526004016102c09061446a565b3360009081526002602052604090205460ff16610f925760405162461bcd60e51b81526004016102c09061445a565b60015460408051600162c2db5f60e01b0319815290516001600160a01b0390921691600091839163ff3d24a191600480820192602092909190829003018186803b158015610fdf57600080fd5b505afa158015610ff3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506110179190810190612f6d565b9050606061102482611d16565b9050606061103182611d5a565b6110c76002308560405160200161104992919061412c565b60408051601f1981840301815290829052611063916141cc565b602060405180830381855afa158015611080573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052506110a39190810190612f6d565b6040516020016110b3919061419b565b604051602081830303815290604052611d5a565b6110d86110d333611d91565b611d5a565b6110e18f611dac565b6111208f8f8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611d5a92505050565b61115f8e8e8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611d5a92505050565b61119e8d8d8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611d5a92505050565b6040516020016111b497969594939291906141f0565b60408051601f19818403018152908290528051602082012063130f33d960e21b83529092506001600160a01b03861691634c3ccf64916111f691600401614292565b602060405180830381600087803b15801561121057600080fd5b505af1158015611224573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506112489190810190612f4f565b6112645760405162461bcd60e51b81526004016102c09061465a565b326001600160a01b03167f6ad3bf15c1988bc04bc153490cab16db8efb9a3990215bf1c64ea6e28be8848383338f8f8f876040516112a7969594939291906142e6565b60405180910390a25060019b9a5050505050505050505050565b60008054600160a01b900460ff16156112ec5760405162461bcd60e51b81526004016102c09061446a565b6112f4612d5a565b6112fd866117c9565b90506000600160009054906101000a90046001600160a01b031690506060611357826001600160a01b03166369d480746040518163ffffffff1660e01b815260040160006040518083038186803b15801561041257600080fd5b90506000826001600160a01b0316635ac407906040518163ffffffff1660e01b815260040160206040518083038186803b15801561139457600080fd5b505afa1580156113a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506113cc91908101906131d2565b63ffffffff16905060008251905081856060015163ffffffff161061141e576113fd8a888560036000198601610466565b6114195760405162461bcd60e51b81526004016102c09061468a565b6114a4565b61143088888560036000198601610466565b61144c5760405162461bcd60e51b81526004016102c09061457a565b611454612d5a565b61145d896117c9565b905060606114708b836101000151611def565b905061147b81611ed7565b6114848d611f02565b146114a15760405162461bcd60e51b81526004016102c0906143ea565b50505b60606114b48c8760e00151611def565b90506114be612db5565b6114c782611fbf565b9050856001600160a01b0316630586763c82602001516114ea8460000151611ed7565b6040518363ffffffff1660e01b81526004016115079291906146d6565b60206040518083038186803b15801561151f57600080fd5b505afa158015611533573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506115579190810190612f4f565b156115745760405162461bcd60e51b81526004016102c09061464a565b856001600160a01b031663e90bfdcf82602001516115958460000151611ed7565b6040518363ffffffff1660e01b81526004016115b29291906146d6565b602060405180830381600087803b1580156115cc57600080fd5b505af11580156115e0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506116049190810190612f4f565b6116205760405162461bcd60e51b81526004016102c09061440a565b6001546040820151606001516001600160401b03908116600160a01b909204161461165d5760405162461bcd60e51b81526004016102c09061442a565b600061167082604001516080015161209b565b6001600160a01b03811660009081526003602052604090205490915060ff166116ab5760405162461bcd60e51b81526004016102c09061449a565b6004826040015160a001516040516116c391906141cc565b9081526040519081900360200190205460ff166116f25760405162461bcd60e51b81526004016102c09061459a565b61171b81836040015160a00151846040015160c0015185604001516040015186602001516120c6565b6117375760405162461bcd60e51b81526004016102c09061438a565b602082015160408084015160808101518551915192517f8a4a2663ce60ce4955c595da2894de0415240f1ace024cfbff85f513b656bdae9461177b949093916146f1565b60405180910390a16001985050505050505050505b95945050505050565b6117a1610ee9565b6117bd5760405162461bcd60e51b81526004016102c0906144fa565b6117c681612237565b50565b6117d1612d5a565b6117d9612d5a565b60006117e584826122b8565b63ffffffff909116835290506117fb8482612341565b6001600160401b039091166020840152905061181784826123c8565b60a0840191909152905061182b84826123c8565b60c0840191909152905061183f84826123c8565b60e0840191909152905061185384826123c8565b610100840191909152905061186884826122b8565b63ffffffff9091166040840152905061188184826122b8565b63ffffffff9091166060840152905061189a8482612341565b6001600160401b03909116608084015290506118b6848261240d565b61012084019190915290506118cb84826124c1565b506bffffffffffffffffffffffff19166101408301525092915050565b60606000806118f78482612341565b80935081925050506060816001600160401b0316604051908082528060200260200182016040528015611934578160200160208202803883390190505b509050606060005b836001600160401b031681101561199057611957878661240d565b955091506119648261209b565b83828151811061197057fe5b6001600160a01b039092166020928302919091019091015260010161193c565b509095945050505050565b6000806119a786611f02565b9050600060418651816119b657fe5b0490506060816040519080825280602002602001820160405280156119e5578160200160208202803883390190505b50905060008080805b85811015611b3257611a0d611a088c60418402602061250a565b611ed7565b9350611a24611a088c60418402602001602061250a565b92508a6041820260400181518110611a3857fe5b602001015160f81c60f81b60f81c601b0191506001600288604051602001611a60919061419b565b60408051601f1981840301815290829052611a7a916141cc565b602060405180830381855afa158015611a97573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250611aba9190810190612f6d565b83868660405160008152602001604052604051611ada94939291906142a0565b6020604051602081039080840390855afa158015611afc573d6000803e3d6000fd5b50505060206040510351858281518110611b1257fe5b6001600160a01b03909216602092830291909101909101526001016119ee565b50611b3e89858a61258a565b96505050505050505b949350505050565b600060606043835181611b5e57fe5b0615611b7c5760405162461bcd60e51b81526004016102c09061453a565b60006043845181611b8957fe5b0490506001811015611bad5760405162461bcd60e51b81526004016102c0906145ba565b611bc181600360001982010483038661262b565b9250925050915091565b805160609081611bda82611dac565b905060005b82811015611c325781611c076110d3878481518110611bfa57fe5b6020026020010151611d91565b604051602001611c189291906141d8565b60408051601f198184030181529190529150600101611bdf565b509392505050565b600054600160a01b900460ff16611c635760405162461bcd60e51b81526004016102c09061437a565b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa611c99611d12565b604051611ca69190614276565b60405180910390a1565b600054600160a01b900460ff1615611cda5760405162461bcd60e51b81526004016102c09061446a565b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258611c995b3390565b60606001600160ff1b03821115611d3f5760405162461bcd60e51b81526004016102c0906144da565b60405190506020815281602082015260408101604052919050565b8051606090611d68816127e0565b83604051602001611d7a9291906141d8565b604051602081830303815290604052915050919050565b604080516014815260609290921b6020830152818101905290565b6040516008808252606091906000601f5b82821015611ddf5785811a826020860101536001919091019060001901611dbd565b5050506028810160405292915050565b6060600081611dfe858361240d565b925090506000611e0d82612897565b9050600060218488510381611e1e57fe5b049050600080805b83811015611eaa57611e388a886128b2565b97509150611e468a886123c8565b975092506001600160f81b03198216611e6a57611e6383866128fb565b9450611ea2565b600160f81b6001600160f81b031983161415611e8a57611e6385846128fb565b60405162461bcd60e51b81526004016102c09061451a565b600101611e26565b50878414611eca5760405162461bcd60e51b81526004016102c09061447a565b5092979650505050505050565b60008151602014611efa5760405162461bcd60e51b81526004016102c0906145aa565b506020015190565b600060028083604051611f1591906141cc565b602060405180830381855afa158015611f32573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250611f559190810190612f6d565b604051602001611f65919061419b565b60408051601f1981840301815290829052611f7f916141cc565b602060405180830381855afa158015611f9c573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052506109109190810190612f6d565b611fc7612db5565b611fcf612db5565b6000611fdb848261240d565b9083529050611fea8482612341565b6001600160401b0390911660208401529050612004612dd9565b61200e858361240d565b908252915061201d858361240d565b60208301919091529150612031858361240d565b604083019190915291506120458583612341565b6001600160401b0390911660608301529150612061858361240d565b60808301919091529150612075858361240d565b60a08301919091529150612089858361240d565b5060c082015260408301525092915050565b600081516014146120be5760405162461bcd60e51b81526004016102c0906143da565b506014015190565b60006120d186612972565b6120ed5760405162461bcd60e51b81526004016102c09061454a565b60606000876001600160a01b03168760405160200161210c9190614251565b6040516020818303038152906040528051906020012087878760405160200161213793929190614346565b60408051601f198184030181529082905261215592916020016141b0565b60408051601f198184030181529082905261216f916141cc565b6000604051808303816000865af19150503d80600081146121ac576040519150601f19603f3d011682016040523d82523d6000602084013e6121b1565b606091505b50925090506001811515146121d85760405162461bcd60e51b81526004016102c0906144ba565b81516121f65760405162461bcd60e51b81526004016102c09061458a565b600061220383601f6129a9565b5090506001811515146122285760405162461bcd60e51b81526004016102c0906145ca565b50600198975050505050505050565b6001600160a01b03811661225d5760405162461bcd60e51b81526004016102c0906143fa565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b600080835183600401111580156122d157508260040183105b6122ed5760405162461bcd60e51b81526004016102c09061466a565b600060405160046000600182038760208a0101515b838310156123225780821a83860153600183019250600182039150612302565b50505080820160405260200390035192505050600482015b9250929050565b6000808351836008011115801561235a57508260080183105b6123765760405162461bcd60e51b81526004016102c09061462a565b600060405160086000600182038760208a0101515b838310156123ab5780821a8386015360018301925060018203915061238b565b505050808201604052602003900351956008949094019450505050565b600080835183602001111580156123e157508260200183105b6123fd5760405162461bcd60e51b81526004016102c09061461a565b5050602091810182015192910190565b606060008061241c8585612a42565b865190955090915081850111801590612436575080840184105b6124525760405162461bcd60e51b81526004016102c09061467a565b60608115801561246d576040519150602082016040526124b7565b6040519150601f8316801560200281840101848101888315602002848c0101015b818310156124a657805183526020928301920161248e565b5050848452601f01601f1916604052505b5095930193505050565b600080835183601401111580156124da57508260140183105b6124f65760405162461bcd60e51b81526004016102c09061443a565b505081810160200151601482019250929050565b60608183018451101561251c57600080fd5b60608215801561253757604051915060208201604052612581565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015612570578051835260209283019201612558565b5050858452601f01601f1916604052505b50949350505050565b600080805b845181101561261f5760005b8651811015612616578681815181106125b057fe5b60200260200101516001600160a01b03168683815181106125cd57fe5b60200260200101516001600160a01b0316141561260e5782806001019350508681815181106125f857fe5b6020026020010160006001600160a01b03168152505b60010161259b565b5060010161258f565b50909111159392505050565b600060608061263986612b87565b9050606086604051908082528060200260200182016040528015612667578160200160208202803883390190505b50905060006060815b898110156126fe576126878860438302604361250a565b9150846126966110d384612bca565b6040516020016126a79291906141d8565b60405160208183030381529060405294506126c5826003604061250a565b8051906020012092508260001c8482815181106126de57fe5b6001600160a01b0390921660209283029190910190910152600101612670565b508361270989612b87565b60405160200161271a9291906141d8565b60405160208183030381529060405293506000600360028660405161273f91906141cc565b602060405180830381855afa15801561275c573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525061277f9190810190612f6d565b60405160200161278f919061419b565b60408051601f19818403018152908290526127a9916141cc565b602060405180830381855afa1580156127c6573d6000803e3d6000fd5b50506040515160601b9b949a509398505050505050505050565b606060fd826001600160401b03161015612804576127fd82612c83565b9050610af0565b61ffff826001600160401b0316116128535761282360fd60f81b612c9f565b61282c83612b87565b60405160200161283d9291906141d8565b6040516020818303038152906040529050610af0565b63ffffffff826001600160401b03161161287d57612874607f60f91b612c9f565b61282c83612cad565b61288e6001600160f81b0319612c9f565b61282c83611dac565b60006002600060f81b83604051602001611f6592919061417f565b600080835183600101111580156128cb57508260010183105b6128e75760405162461bcd60e51b81526004016102c09061452a565b505081810160200151600182019250929050565b60006002600160f81b848460405160200161291893929190614148565b60408051601f1981840301815290829052612932916141cc565b602060405180830381855afa15801561294f573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525061065c9190810190612f6d565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4708115801590611b475750141592915050565b600080835183600101111580156129c257508260010183105b6129de5760405162461bcd60e51b81526004016102c0906145fa565b838301602001516000600160f81b6001600160f81b031983161415612a0557506001612a34565b6001600160f81b03198216612a1c57506000612a34565b60405162461bcd60e51b81526004016102c09061450a565b956001949094019450505050565b6000806000612a5185856128b2565b94509050600060fd60f81b6001600160f81b031983161415612aba57612a778686612cf0565b955061ffff16905060fd8110801590612a92575061ffff8111155b612aae5760405162461bcd60e51b81526004016102c0906143aa565b925083915061233a9050565b607f60f91b6001600160f81b031983161415612b1557612ada86866122b8565b955063ffffffff16905061ffff81118015612af9575063ffffffff8111155b612aae5760405162461bcd60e51b81526004016102c09061456a565b6001600160f81b03198083161415612b6157612b318686612341565b95506001600160401b0316905063ffffffff8111612aae5760405162461bcd60e51b81526004016102c09061456a565b5060f881901c60fd8110612aae5760405162461bcd60e51b81526004016102c09061456a565b6040516002808252606091906000601f5b82821015612bba5785811a826020860101536001919091019060001901612b98565b5050506022810160405292915050565b6060604382511015612bee5760405162461bcd60e51b81526004016102c09061463a565b612bfb826000602361250a565b9050600282604281518110612c0c57fe5b016020015160f81c81612c1b57fe5b0660ff1660001415612c5557600260f81b81600281518110612c3957fe5b60200101906001600160f81b031916908160001a905350610af0565b600360f81b81600281518110612c6757fe5b60200101906001600160f81b031916908160001a905350919050565b604080516001815260f89290921b602083015260218201905290565b60606109108260f81c612c83565b6040516004808252606091906000601f5b82821015612ce05785811a826020860101536001919091019060001901612cbe565b5050506024810160405292915050565b60008083518360020111158015612d0957508260020183105b612d255760405162461bcd60e51b81526004016102c0906145da565b6000604051846020870101518060011a82538060001a6001830153506002818101604052601d19909101519694019450505050565b604080516101608101825260008082526020820181905291810182905260608082018390526080820183905260a0820183905260c0820183905260e08201839052610100820183905261012082015261014081019190915290565b6040805160608082018352815260006020820152908101612dd4612dd9565b905290565b6040518060e0016040528060608152602001606081526020016060815260200160006001600160401b031681526020016060815260200160608152602001606081525090565b80356109108161486a565b80516109108161487e565b805161091081614887565b60008083601f840112612e5257600080fd5b5081356001600160401b03811115612e6957600080fd5b60208301915083600182028301111561233a57600080fd5b600082601f830112612e9257600080fd5b8135612ea5612ea082614769565b614743565b91508082526020830160208301858383011115612ec157600080fd5b612ecc83828461480d565b50505092915050565b600082601f830112612ee657600080fd5b8151612ef4612ea082614769565b91508082526020830160208301858383011115612f1057600080fd5b612ecc838284614819565b805161091081614890565b803561091081614899565b600060208284031215612f4357600080fd5b6000611b478484612e1f565b600060208284031215612f6157600080fd5b6000611b478484612e2a565b600060208284031215612f7f57600080fd5b6000611b478484612e35565b600060208284031215612f9d57600080fd5b81356001600160401b03811115612fb357600080fd5b611b4784828501612e81565b600060208284031215612fd157600080fd5b81516001600160401b03811115612fe757600080fd5b611b4784828501612ed5565b6000806040838503121561300657600080fd5b82356001600160401b0381111561301c57600080fd5b61302885828601612e81565b92505060208301356001600160401b0381111561304457600080fd5b61305085828601612e81565b9150509250929050565b60008060006060848603121561306f57600080fd5b83356001600160401b0381111561308557600080fd5b61309186828701612e81565b93505060208401356001600160401b038111156130ad57600080fd5b6130b986828701612e81565b92505060408401356001600160401b038111156130d557600080fd5b6130e186828701612e81565b9150509250925092565b600080600080600060a0868803121561310357600080fd5b85356001600160401b0381111561311957600080fd5b61312588828901612e81565b95505060208601356001600160401b0381111561314157600080fd5b61314d88828901612e81565b94505060408601356001600160401b0381111561316957600080fd5b61317588828901612e81565b93505060608601356001600160401b0381111561319157600080fd5b61319d88828901612e81565b92505060808601356001600160401b038111156131b957600080fd5b6131c588828901612e81565b9150509295509295909350565b6000602082840312156131e457600080fd5b6000611b478484612f1b565b60006020828403121561320257600080fd5b6000611b478484612f26565b60008060008060008060006080888a03121561322957600080fd5b60006132358a8a612f26565b97505060208801356001600160401b0381111561325157600080fd5b61325d8a828b01612e40565b965096505060408801356001600160401b0381111561327b57600080fd5b6132878a828b01612e40565b945094505060608801356001600160401b038111156132a557600080fd5b6132b18a828b01612e40565b925092505092959891949750929550565b6132cb816147f1565b82525050565b6132cb8161479d565b6132cb6132e68261479d565b614849565b6132cb816147a8565b6132cb613300826147ad565b6147ba565b6132cb816147ba565b6132cb613300826147ba565b6132cb613300826147bd565b60006133328385614794565b935061333f83858461480d565b6133488361485a565b9093019392505050565b600061335d82614790565b6133678185614794565b9350613377818560208601614819565b6133488161485a565b600061338b82614790565b6133958185610af0565b93506133a5818560208601614819565b9290920192915050565b60006133bc601483614794565b7314185d5cd8589b194e881b9bdd081c185d5cd95960621b815260200192915050565b60006133ec601d83614794565b7f457865637574652043726f7373436861696e205478206661696c656421000000815260200192915050565b6000613425602783614794565b7f70617573652045746843726f7373436861696e4461746120636f6e74726163748152660819985a5b195960ca1b602082015260400192915050565b600061346e601f83614794565b7f4e65787455696e7431362c2076616c7565206f7574736964652072616e676500815260200192915050565b60006134a7602583614794565b7f546865206e657874426f6f6b4b6565706572206f662068656164657220697320815264656d70747960d81b602082015260400192915050565b60006134ee602983614794565b7f756e70617573652045746843726f7373436861696e4461746120636f6e74726181526818dd0819985a5b195960ba1b602082015260400192915050565b6000613539602383614794565b7f6279746573206c656e67746820646f6573206e6f74206d61746368206164647281526265737360e81b602082015260400192915050565b600061357e601b83614794565b7f766572696679206865616465722070726f6f66206661696c6564210000000000815260200192915050565b60006135b7602683614794565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206181526564647265737360d01b602082015260400192915050565b60006135ff602083614794565b7f536176652063726f7373636861696e207478206578697374206661696c656421815260200192915050565b6000613638604383614794565b7f5361766520506f6c7920636861696e2063757272656e742065706f636820737481527f6172742068656967687420746f204461746120636f6e7472616374206661696c60208201526265642160e81b604082015260600192915050565b60006136a3602683614794565b7f54686973205478206973206e6f742061696d696e672061742074686973206e6581526574776f726b2160d01b602082015260400192915050565b60006136eb602383614794565b7f4e657874427974657332302c206f66667365742065786365656473206d6178698152626d756d60e81b602082015260400192915050565b6000613730601883614794565b7f566572696679207369676e6174757265206661696c6564210000000000000000815260200192915050565b6000613769601583614794565b74125b9d985b1a5908199c9bdb4818dbdb9d1c9858dd605a1b815260200192915050565b600061379a601083614794565b6f14185d5cd8589b194e881c185d5cd95960821b815260200192915050565b60006137c6603183614794565b7f6d65726b6c6550726f76652c2065787065637420726f6f74206973206e6f7420815270195c5d585b081858dd1d585b081c9bdbdd607a1b602082015260400192915050565b6000613819601383614794565b7213995e1d109bdbdad95c9cc81a5b1b1959d85b606a1b815260200192915050565b6000613848601383614794565b72125b9d985b1a59081d1bc818dbdb9d1c9858dd606a1b815260200192915050565b6000613877603e83614794565b7f54686520686569676874206f6620686561646572206973206c6f77657220746881527f616e2063757272656e742065706f636820737461727420686569676874210000602082015260400192915050565b60006138d6602b83614794565b7f45746843726f7373436861696e2063616c6c20627573696e65737320636f6e7481526a1c9858dd0819985a5b195960aa1b602082015260400192915050565b6000613923603b83614794565b7f5361766520506f6c7920636861696e20626f6f6b206b6565706572732062797481527f657320746f204461746120636f6e7472616374206661696c6564210000000000602082015260400192915050565b6000613982601783614794565b7f56616c75652065786365656473207468652072616e6765000000000000000000815260200192915050565b60006139bb601583614794565b74111bdb89dd08185c989a5d1c985c9a5b1e481cd95d605a1b815260200192915050565b60006139ec602083614794565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572815260200192915050565b6000613a25601483614794565b732732bc3a2137b7b6103b30b63ab29032b93937b960611b815260200192915050565b6000613a55602e83614794565b7f6d65726b6c6550726f76652c204e6578744279746520666f7220706f7369746981526d1bdb881a5b999bc819985a5b195960921b602082015260400192915050565b6000613aa5602083614794565b7f4e657874427974652c204f66667365742065786365656473206d6178696d756d815260200192915050565b6000613ade601b83614794565b7f5f7075624b65794c697374206c656e67746820696c6c6567616c210000000000815260200192915050565b6000613b17602883614794565b7f5468652070617373656420696e2061646472657373206973206e6f74206120638152676f6e74726163742160c01b602082015260400192915050565b6000613b61602d83614794565b7f53617665204d43204c617465737448656967687420746f204461746120636f6e81526c7472616374206661696c65642160981b602082015260400192915050565b6000613bb0602083614794565b7f4e65787456617255696e742c2076616c7565206f7574736964652072616e6765815260200192915050565b6000613be9603883614794565b7f56657269667920706f6c7920636861696e2063757272656e742065706f63682081527f686561646572207369676e6174757265206661696c6564210000000000000000602082015260400192915050565b6000613c48601483610af0565b732862797465732c62797465732c75696e7436342960601b815260140192915050565b6000613c78602783614794565b7f4e6f2072657475726e2076616c75652066726f6d20627573696e65737320636f8152666e74726163742160c81b602082015260400192915050565b6000613cc1600e83614794565b6d125b9d985b1a59081b595d1a1bd960921b815260200192915050565b6000613ceb601783614794565b7f6279746573206c656e677468206973206e6f742033322e000000000000000000815260200192915050565b6000613d24601683614794565b75746f6f2073686f7274205f7075624b65794c6973742160501b815260200192915050565b6000613d56603783614794565b7f45746843726f7373436861696e2063616c6c20627573696e65737320636f6e7481527f726163742072657475726e206973206e6f742074727565000000000000000000602082015260400192915050565b6000613db5602283614794565b7f4e65787455696e7431362c206f66667365742065786365656473206d6178696d815261756d60f01b602082015260400192915050565b6000613df9603883614794565b7f45746843726f7373436861696e4461746120636f6e747261637420686173206181527f6c7265616479206265656e20696e697469616c697a6564210000000000000000602082015260400192915050565b6000613e58601483614794565b7313d9999cd95d08195e18d959591cc81b1a5b5a5d60621b815260200192915050565b6000613e88604383614794565b7f5361766520506f6c7920636861696e2063757272656e742065706f636820626f81527f6f6b206b65657065727320746f204461746120636f6e7472616374206661696c60208201526265642160e81b604082015260600192915050565b6000613ef3602083614794565b7f4e657874486173682c206f66667365742065786365656473206d6178696d756d815260200192915050565b6000613f2c602283614794565b7f4e65787455696e7436342c206f66667365742065786365656473206d6178696d815261756d60f01b602082015260400192915050565b6000613f70601783614794565b7f6b6579206c656e67676820697320746f6f2073686f7274000000000000000000815260200192915050565b6000613fa9602283614794565b7f746865207472616e73616374696f6e20686173206265656e2065786563757465815261642160f01b602082015260400192915050565b6000613fed603083614794565b7f536176652065746854784861736820627920696e64657820746f20446174612081526f636f6e7472616374206661696c65642160801b602082015260400192915050565b600061403f602283614794565b7f4e65787455696e7433322c206f66667365742065786365656473206d6178696d815261756d60f01b602082015260400192915050565b6000614083602483614794565b7f4e65787456617242797465732c206f66667365742065786365656473206d6178815263696d756d60e01b602082015260400192915050565b60006140c9602a83614794565b7f56657269667920706f6c7920636861696e20686561646572207369676e6174758152697265206661696c65642160b01b602082015260400192915050565b6132cb81614802565b6132cb816147d6565b6132cb816147df565b6132cb816147eb565b600061413882856132da565b601482019150611b478284613380565b600061415482866132f4565b600182019150614164828561330e565b602082019150614174828461330e565b506020019392505050565b600061418b82856132f4565b600182019150611b478284613380565b60006141a7828461330e565b50602001919050565b60006141bc828561331a565b600482019150611b478284613380565b600061065c8284613380565b60006141e48285613380565b9150611b478284613380565b60006141fc828a613380565b91506142088289613380565b91506142148288613380565b91506142208287613380565b915061422c8286613380565b91506142388285613380565b91506142448284613380565b9998505050505050505050565b600061425d8284613380565b915061065c82613c3b565b6020810161091082846132d1565b6020810161091082846132c2565b6020810161091082846132eb565b602081016109108284613305565b608081016142ae8287613305565b6142bb6020830186614123565b6142c86040830185613305565b6117906060830184613305565b6020808252810161065c8184613352565b60a080825281016142f78189613352565b905061430660208301886132c2565b614313604083018761411a565b8181036060830152614326818587613326565b9050818103608083015261433a8184613352565b98975050505050505050565b606080825281016143578186613352565b9050818103602083015261436b8185613352565b9050611b47604083018461411a565b60208082528101610910816133af565b60208082528101610910816133df565b6020808252810161091081613418565b6020808252810161091081613461565b602080825281016109108161349a565b60208082528101610910816134e1565b602080825281016109108161352c565b6020808252810161091081613571565b60208082528101610910816135aa565b60208082528101610910816135f2565b602080825281016109108161362b565b6020808252810161091081613696565b60208082528101610910816136de565b6020808252810161091081613723565b602080825281016109108161375c565b602080825281016109108161378d565b60208082528101610910816137b9565b602080825281016109108161380c565b602080825281016109108161383b565b602080825281016109108161386a565b60208082528101610910816138c9565b6020808252810161091081613916565b6020808252810161091081613975565b60208082528101610910816139ae565b60208082528101610910816139df565b6020808252810161091081613a18565b6020808252810161091081613a48565b6020808252810161091081613a98565b6020808252810161091081613ad1565b6020808252810161091081613b0a565b6020808252810161091081613b54565b6020808252810161091081613ba3565b6020808252810161091081613bdc565b6020808252810161091081613c6b565b6020808252810161091081613cb4565b6020808252810161091081613cde565b6020808252810161091081613d17565b6020808252810161091081613d49565b6020808252810161091081613da8565b6020808252810161091081613dec565b6020808252810161091081613e4b565b6020808252810161091081613e7b565b6020808252810161091081613ee6565b6020808252810161091081613f1f565b6020808252810161091081613f63565b6020808252810161091081613f9c565b6020808252810161091081613fe0565b6020808252810161091081614032565b6020808252810161091081614076565b60208082528101610910816140bc565b602081016109108284614111565b604081016146b68285614108565b8181036020830152611b478184613352565b60208101610910828461411a565b604081016146e4828561411a565b61065c6020830184613305565b608081016146ff828761411a565b81810360208301526147118186613352565b905081810360408301526147258185613352565b905081810360608301526147398184613352565b9695505050505050565b6040518181016001600160401b038111828210171561476157600080fd5b604052919050565b60006001600160401b0382111561477f57600080fd5b506020601f91909101601f19160190565b5190565b90815260200190565b6000610910826147ca565b151590565b6001600160f81b03191690565b90565b6001600160e01b03191690565b6001600160a01b031690565b63ffffffff1690565b6001600160401b031690565b60ff1690565b60006109108260006109108261479d565b6000610910826147d6565b82818337506000910152565b60005b8381101561483457818101518382015260200161481c565b83811115614843576000848401525b50505050565b600061091082600061091082614864565b601f01601f191690565b60601b90565b6148738161479d565b81146117c657600080fd5b614873816147a8565b614873816147ba565b614873816147d6565b614873816147df56fea365627a7a72315820182231cfbbbb0d9399442a89c652c7861f0aea19a0556423814ccd531753ff886c6578706572696d656e74616cf564736f6c63430005110040"

// DeployEthCrossChainManager deploys a new Ethereum contract, binding an instance of EthCrossChainManager to it.
func DeployEthCrossChainManager(auth *bind.TransactOpts, backend bind.ContractBackend, _eccd common.Address, _chainId uint64, fromContractWhiteList []common.Address, toContractWhiteList []common.Address, methodWhiteList [][]byte, curEpochPkBytes []byte) (common.Address, *types.Transaction, *EthCrossChainManager, error) {
	parsed, err := abi.JSON(strings.NewReader(EthCrossChainManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthCrossChainManagerBin), backend, _eccd, _chainId, fromContractWhiteList, toContractWhiteList, methodWhiteList, curEpochPkBytes)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthCrossChainManager{EthCrossChainManagerCaller: EthCrossChainManagerCaller{contract: contract}, EthCrossChainManagerTransactor: EthCrossChainManagerTransactor{contract: contract}, EthCrossChainManagerFilterer: EthCrossChainManagerFilterer{contract: contract}}, nil
}

// EthCrossChainManager is an auto generated Go binding around an Ethereum contract.
type EthCrossChainManager struct {
	EthCrossChainManagerCaller     // Read-only binding to the contract
	EthCrossChainManagerTransactor // Write-only binding to the contract
	EthCrossChainManagerFilterer   // Log filterer for contract events
}

// EthCrossChainManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthCrossChainManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthCrossChainManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthCrossChainManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthCrossChainManagerSession struct {
	Contract     *EthCrossChainManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EthCrossChainManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthCrossChainManagerCallerSession struct {
	Contract *EthCrossChainManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// EthCrossChainManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthCrossChainManagerTransactorSession struct {
	Contract     *EthCrossChainManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// EthCrossChainManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthCrossChainManagerRaw struct {
	Contract *EthCrossChainManager // Generic contract binding to access the raw methods on
}

// EthCrossChainManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthCrossChainManagerCallerRaw struct {
	Contract *EthCrossChainManagerCaller // Generic read-only contract binding to access the raw methods on
}

// EthCrossChainManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthCrossChainManagerTransactorRaw struct {
	Contract *EthCrossChainManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthCrossChainManager creates a new instance of EthCrossChainManager, bound to a specific deployed contract.
func NewEthCrossChainManager(address common.Address, backend bind.ContractBackend) (*EthCrossChainManager, error) {
	contract, err := bindEthCrossChainManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManager{EthCrossChainManagerCaller: EthCrossChainManagerCaller{contract: contract}, EthCrossChainManagerTransactor: EthCrossChainManagerTransactor{contract: contract}, EthCrossChainManagerFilterer: EthCrossChainManagerFilterer{contract: contract}}, nil
}

// NewEthCrossChainManagerCaller creates a new read-only instance of EthCrossChainManager, bound to a specific deployed contract.
func NewEthCrossChainManagerCaller(address common.Address, caller bind.ContractCaller) (*EthCrossChainManagerCaller, error) {
	contract, err := bindEthCrossChainManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerCaller{contract: contract}, nil
}

// NewEthCrossChainManagerTransactor creates a new write-only instance of EthCrossChainManager, bound to a specific deployed contract.
func NewEthCrossChainManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*EthCrossChainManagerTransactor, error) {
	contract, err := bindEthCrossChainManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerTransactor{contract: contract}, nil
}

// NewEthCrossChainManagerFilterer creates a new log filterer instance of EthCrossChainManager, bound to a specific deployed contract.
func NewEthCrossChainManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*EthCrossChainManagerFilterer, error) {
	contract, err := bindEthCrossChainManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerFilterer{contract: contract}, nil
}

// bindEthCrossChainManager binds a generic wrapper to an already deployed contract.
func bindEthCrossChainManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthCrossChainManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthCrossChainManager *EthCrossChainManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthCrossChainManager.Contract.EthCrossChainManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthCrossChainManager *EthCrossChainManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.EthCrossChainManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthCrossChainManager *EthCrossChainManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.EthCrossChainManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthCrossChainManager *EthCrossChainManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthCrossChainManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthCrossChainManager *EthCrossChainManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthCrossChainManager *EthCrossChainManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.contract.Transact(opts, method, params...)
}

// EthCrossChainDataAddress is a free data retrieval call binding the contract method 0x00ba1694.
//
// Solidity: function EthCrossChainDataAddress() view returns(address)
func (_EthCrossChainManager *EthCrossChainManagerCaller) EthCrossChainDataAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EthCrossChainManager.contract.Call(opts, out, "EthCrossChainDataAddress")
	return *ret0, err
}

// EthCrossChainDataAddress is a free data retrieval call binding the contract method 0x00ba1694.
//
// Solidity: function EthCrossChainDataAddress() view returns(address)
func (_EthCrossChainManager *EthCrossChainManagerSession) EthCrossChainDataAddress() (common.Address, error) {
	return _EthCrossChainManager.Contract.EthCrossChainDataAddress(&_EthCrossChainManager.CallOpts)
}

// EthCrossChainDataAddress is a free data retrieval call binding the contract method 0x00ba1694.
//
// Solidity: function EthCrossChainDataAddress() view returns(address)
func (_EthCrossChainManager *EthCrossChainManagerCallerSession) EthCrossChainDataAddress() (common.Address, error) {
	return _EthCrossChainManager.Contract.EthCrossChainDataAddress(&_EthCrossChainManager.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_EthCrossChainManager *EthCrossChainManagerCaller) ChainId(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _EthCrossChainManager.contract.Call(opts, out, "chainId")
	return *ret0, err
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_EthCrossChainManager *EthCrossChainManagerSession) ChainId() (uint64, error) {
	return _EthCrossChainManager.Contract.ChainId(&_EthCrossChainManager.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_EthCrossChainManager *EthCrossChainManagerCallerSession) ChainId() (uint64, error) {
	return _EthCrossChainManager.Contract.ChainId(&_EthCrossChainManager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EthCrossChainManager.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) IsOwner() (bool, error) {
	return _EthCrossChainManager.Contract.IsOwner(&_EthCrossChainManager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerCallerSession) IsOwner() (bool, error) {
	return _EthCrossChainManager.Contract.IsOwner(&_EthCrossChainManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthCrossChainManager *EthCrossChainManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EthCrossChainManager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthCrossChainManager *EthCrossChainManagerSession) Owner() (common.Address, error) {
	return _EthCrossChainManager.Contract.Owner(&_EthCrossChainManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthCrossChainManager *EthCrossChainManagerCallerSession) Owner() (common.Address, error) {
	return _EthCrossChainManager.Contract.Owner(&_EthCrossChainManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EthCrossChainManager.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) Paused() (bool, error) {
	return _EthCrossChainManager.Contract.Paused(&_EthCrossChainManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerCallerSession) Paused() (bool, error) {
	return _EthCrossChainManager.Contract.Paused(&_EthCrossChainManager.CallOpts)
}

// UnsetEpochPkBytes is a free data retrieval call binding the contract method 0x8c5aea41.
//
// Solidity: function unsetEpochPkBytes(bytes ) view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerCaller) UnsetEpochPkBytes(opts *bind.CallOpts, arg0 []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EthCrossChainManager.contract.Call(opts, out, "unsetEpochPkBytes", arg0)
	return *ret0, err
}

// UnsetEpochPkBytes is a free data retrieval call binding the contract method 0x8c5aea41.
//
// Solidity: function unsetEpochPkBytes(bytes ) view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) UnsetEpochPkBytes(arg0 []byte) (bool, error) {
	return _EthCrossChainManager.Contract.UnsetEpochPkBytes(&_EthCrossChainManager.CallOpts, arg0)
}

// UnsetEpochPkBytes is a free data retrieval call binding the contract method 0x8c5aea41.
//
// Solidity: function unsetEpochPkBytes(bytes ) view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerCallerSession) UnsetEpochPkBytes(arg0 []byte) (bool, error) {
	return _EthCrossChainManager.Contract.UnsetEpochPkBytes(&_EthCrossChainManager.CallOpts, arg0)
}

// WhiteListFromContract is a free data retrieval call binding the contract method 0x73f53ba4.
//
// Solidity: function whiteListFromContract(address ) view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerCaller) WhiteListFromContract(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EthCrossChainManager.contract.Call(opts, out, "whiteListFromContract", arg0)
	return *ret0, err
}

// WhiteListFromContract is a free data retrieval call binding the contract method 0x73f53ba4.
//
// Solidity: function whiteListFromContract(address ) view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) WhiteListFromContract(arg0 common.Address) (bool, error) {
	return _EthCrossChainManager.Contract.WhiteListFromContract(&_EthCrossChainManager.CallOpts, arg0)
}

// WhiteListFromContract is a free data retrieval call binding the contract method 0x73f53ba4.
//
// Solidity: function whiteListFromContract(address ) view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerCallerSession) WhiteListFromContract(arg0 common.Address) (bool, error) {
	return _EthCrossChainManager.Contract.WhiteListFromContract(&_EthCrossChainManager.CallOpts, arg0)
}

// WhiteListMethod is a free data retrieval call binding the contract method 0x7b935854.
//
// Solidity: function whiteListMethod(bytes ) view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerCaller) WhiteListMethod(opts *bind.CallOpts, arg0 []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EthCrossChainManager.contract.Call(opts, out, "whiteListMethod", arg0)
	return *ret0, err
}

// WhiteListMethod is a free data retrieval call binding the contract method 0x7b935854.
//
// Solidity: function whiteListMethod(bytes ) view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) WhiteListMethod(arg0 []byte) (bool, error) {
	return _EthCrossChainManager.Contract.WhiteListMethod(&_EthCrossChainManager.CallOpts, arg0)
}

// WhiteListMethod is a free data retrieval call binding the contract method 0x7b935854.
//
// Solidity: function whiteListMethod(bytes ) view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerCallerSession) WhiteListMethod(arg0 []byte) (bool, error) {
	return _EthCrossChainManager.Contract.WhiteListMethod(&_EthCrossChainManager.CallOpts, arg0)
}

// WhiteListToContract is a free data retrieval call binding the contract method 0xb91ca106.
//
// Solidity: function whiteListToContract(address ) view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerCaller) WhiteListToContract(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EthCrossChainManager.contract.Call(opts, out, "whiteListToContract", arg0)
	return *ret0, err
}

// WhiteListToContract is a free data retrieval call binding the contract method 0xb91ca106.
//
// Solidity: function whiteListToContract(address ) view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) WhiteListToContract(arg0 common.Address) (bool, error) {
	return _EthCrossChainManager.Contract.WhiteListToContract(&_EthCrossChainManager.CallOpts, arg0)
}

// WhiteListToContract is a free data retrieval call binding the contract method 0xb91ca106.
//
// Solidity: function whiteListToContract(address ) view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerCallerSession) WhiteListToContract(arg0 common.Address) (bool, error) {
	return _EthCrossChainManager.Contract.WhiteListToContract(&_EthCrossChainManager.CallOpts, arg0)
}

// ChangeBookKeeper is a paid mutator transaction binding the contract method 0x29dcf4ab.
//
// Solidity: function changeBookKeeper(bytes rawHeader, bytes pubKeyList, bytes sigList) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) ChangeBookKeeper(opts *bind.TransactOpts, rawHeader []byte, pubKeyList []byte, sigList []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "changeBookKeeper", rawHeader, pubKeyList, sigList)
}

// ChangeBookKeeper is a paid mutator transaction binding the contract method 0x29dcf4ab.
//
// Solidity: function changeBookKeeper(bytes rawHeader, bytes pubKeyList, bytes sigList) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) ChangeBookKeeper(rawHeader []byte, pubKeyList []byte, sigList []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.ChangeBookKeeper(&_EthCrossChainManager.TransactOpts, rawHeader, pubKeyList, sigList)
}

// ChangeBookKeeper is a paid mutator transaction binding the contract method 0x29dcf4ab.
//
// Solidity: function changeBookKeeper(bytes rawHeader, bytes pubKeyList, bytes sigList) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) ChangeBookKeeper(rawHeader []byte, pubKeyList []byte, sigList []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.ChangeBookKeeper(&_EthCrossChainManager.TransactOpts, rawHeader, pubKeyList, sigList)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 toChainId, bytes toContract, bytes method, bytes txData) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) CrossChain(opts *bind.TransactOpts, toChainId uint64, toContract []byte, method []byte, txData []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "crossChain", toChainId, toContract, method, txData)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 toChainId, bytes toContract, bytes method, bytes txData) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) CrossChain(toChainId uint64, toContract []byte, method []byte, txData []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.CrossChain(&_EthCrossChainManager.TransactOpts, toChainId, toContract, method, txData)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 toChainId, bytes toContract, bytes method, bytes txData) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) CrossChain(toChainId uint64, toContract []byte, method []byte, txData []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.CrossChain(&_EthCrossChainManager.TransactOpts, toChainId, toContract, method, txData)
}

// InitGenesisBlock is a paid mutator transaction binding the contract method 0x34a773eb.
//
// Solidity: function initGenesisBlock(bytes rawHeader, bytes pubKeyList) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) InitGenesisBlock(opts *bind.TransactOpts, rawHeader []byte, pubKeyList []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "initGenesisBlock", rawHeader, pubKeyList)
}

// InitGenesisBlock is a paid mutator transaction binding the contract method 0x34a773eb.
//
// Solidity: function initGenesisBlock(bytes rawHeader, bytes pubKeyList) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) InitGenesisBlock(rawHeader []byte, pubKeyList []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.InitGenesisBlock(&_EthCrossChainManager.TransactOpts, rawHeader, pubKeyList)
}

// InitGenesisBlock is a paid mutator transaction binding the contract method 0x34a773eb.
//
// Solidity: function initGenesisBlock(bytes rawHeader, bytes pubKeyList) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) InitGenesisBlock(rawHeader []byte, pubKeyList []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.InitGenesisBlock(&_EthCrossChainManager.TransactOpts, rawHeader, pubKeyList)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) Pause() (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.Pause(&_EthCrossChainManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) Pause() (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.Pause(&_EthCrossChainManager.TransactOpts)
}

// RecoverEpochPk is a paid mutator transaction binding the contract method 0x7bcf7dc1.
//
// Solidity: function recoverEpochPk(bytes EpochPkBytes) returns()
func (_EthCrossChainManager *EthCrossChainManagerTransactor) RecoverEpochPk(opts *bind.TransactOpts, EpochPkBytes []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "recoverEpochPk", EpochPkBytes)
}

// RecoverEpochPk is a paid mutator transaction binding the contract method 0x7bcf7dc1.
//
// Solidity: function recoverEpochPk(bytes EpochPkBytes) returns()
func (_EthCrossChainManager *EthCrossChainManagerSession) RecoverEpochPk(EpochPkBytes []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.RecoverEpochPk(&_EthCrossChainManager.TransactOpts, EpochPkBytes)
}

// RecoverEpochPk is a paid mutator transaction binding the contract method 0x7bcf7dc1.
//
// Solidity: function recoverEpochPk(bytes EpochPkBytes) returns()
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) RecoverEpochPk(EpochPkBytes []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.RecoverEpochPk(&_EthCrossChainManager.TransactOpts, EpochPkBytes)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthCrossChainManager *EthCrossChainManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthCrossChainManager *EthCrossChainManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.RenounceOwnership(&_EthCrossChainManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.RenounceOwnership(&_EthCrossChainManager.TransactOpts)
}

// SetChainId is a paid mutator transaction binding the contract method 0x6f31031d.
//
// Solidity: function setChainId(uint64 _newChainId) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) SetChainId(opts *bind.TransactOpts, _newChainId uint64) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "setChainId", _newChainId)
}

// SetChainId is a paid mutator transaction binding the contract method 0x6f31031d.
//
// Solidity: function setChainId(uint64 _newChainId) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) SetChainId(_newChainId uint64) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.SetChainId(&_EthCrossChainManager.TransactOpts, _newChainId)
}

// SetChainId is a paid mutator transaction binding the contract method 0x6f31031d.
//
// Solidity: function setChainId(uint64 _newChainId) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) SetChainId(_newChainId uint64) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.SetChainId(&_EthCrossChainManager.TransactOpts, _newChainId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthCrossChainManager *EthCrossChainManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthCrossChainManager *EthCrossChainManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.TransferOwnership(&_EthCrossChainManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.TransferOwnership(&_EthCrossChainManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) Unpause() (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.Unpause(&_EthCrossChainManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.Unpause(&_EthCrossChainManager.TransactOpts)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address newEthCrossChainManagerAddress) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) UpgradeToNew(opts *bind.TransactOpts, newEthCrossChainManagerAddress common.Address) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "upgradeToNew", newEthCrossChainManagerAddress)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address newEthCrossChainManagerAddress) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) UpgradeToNew(newEthCrossChainManagerAddress common.Address) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.UpgradeToNew(&_EthCrossChainManager.TransactOpts, newEthCrossChainManagerAddress)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address newEthCrossChainManagerAddress) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) UpgradeToNew(newEthCrossChainManagerAddress common.Address) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.UpgradeToNew(&_EthCrossChainManager.TransactOpts, newEthCrossChainManagerAddress)
}

// VerifyHeaderAndExecuteTx is a paid mutator transaction binding the contract method 0xd450e04c.
//
// Solidity: function verifyHeaderAndExecuteTx(bytes proof, bytes rawHeader, bytes headerProof, bytes curRawHeader, bytes headerSig) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) VerifyHeaderAndExecuteTx(opts *bind.TransactOpts, proof []byte, rawHeader []byte, headerProof []byte, curRawHeader []byte, headerSig []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "verifyHeaderAndExecuteTx", proof, rawHeader, headerProof, curRawHeader, headerSig)
}

// VerifyHeaderAndExecuteTx is a paid mutator transaction binding the contract method 0xd450e04c.
//
// Solidity: function verifyHeaderAndExecuteTx(bytes proof, bytes rawHeader, bytes headerProof, bytes curRawHeader, bytes headerSig) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) VerifyHeaderAndExecuteTx(proof []byte, rawHeader []byte, headerProof []byte, curRawHeader []byte, headerSig []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.VerifyHeaderAndExecuteTx(&_EthCrossChainManager.TransactOpts, proof, rawHeader, headerProof, curRawHeader, headerSig)
}

// VerifyHeaderAndExecuteTx is a paid mutator transaction binding the contract method 0xd450e04c.
//
// Solidity: function verifyHeaderAndExecuteTx(bytes proof, bytes rawHeader, bytes headerProof, bytes curRawHeader, bytes headerSig) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) VerifyHeaderAndExecuteTx(proof []byte, rawHeader []byte, headerProof []byte, curRawHeader []byte, headerSig []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.VerifyHeaderAndExecuteTx(&_EthCrossChainManager.TransactOpts, proof, rawHeader, headerProof, curRawHeader, headerSig)
}

// EthCrossChainManagerChangeBookKeeperEventIterator is returned from FilterChangeBookKeeperEvent and is used to iterate over the raw logs and unpacked data for ChangeBookKeeperEvent events raised by the EthCrossChainManager contract.
type EthCrossChainManagerChangeBookKeeperEventIterator struct {
	Event *EthCrossChainManagerChangeBookKeeperEvent // Event containing the contract specifics and raw log

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
func (it *EthCrossChainManagerChangeBookKeeperEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerChangeBookKeeperEvent)
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
		it.Event = new(EthCrossChainManagerChangeBookKeeperEvent)
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
func (it *EthCrossChainManagerChangeBookKeeperEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerChangeBookKeeperEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerChangeBookKeeperEvent represents a ChangeBookKeeperEvent event raised by the EthCrossChainManager contract.
type EthCrossChainManagerChangeBookKeeperEvent struct {
	Height    *big.Int
	RawHeader []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChangeBookKeeperEvent is a free log retrieval operation binding the contract event 0xe60d33488cba3977bf65766cd2f8ac9617f64bf3b3198aff6240ce5c7d43b690.
//
// Solidity: event ChangeBookKeeperEvent(uint256 height, bytes rawHeader)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) FilterChangeBookKeeperEvent(opts *bind.FilterOpts) (*EthCrossChainManagerChangeBookKeeperEventIterator, error) {

	logs, sub, err := _EthCrossChainManager.contract.FilterLogs(opts, "ChangeBookKeeperEvent")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerChangeBookKeeperEventIterator{contract: _EthCrossChainManager.contract, event: "ChangeBookKeeperEvent", logs: logs, sub: sub}, nil
}

// WatchChangeBookKeeperEvent is a free log subscription operation binding the contract event 0xe60d33488cba3977bf65766cd2f8ac9617f64bf3b3198aff6240ce5c7d43b690.
//
// Solidity: event ChangeBookKeeperEvent(uint256 height, bytes rawHeader)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) WatchChangeBookKeeperEvent(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerChangeBookKeeperEvent) (event.Subscription, error) {

	logs, sub, err := _EthCrossChainManager.contract.WatchLogs(opts, "ChangeBookKeeperEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerChangeBookKeeperEvent)
				if err := _EthCrossChainManager.contract.UnpackLog(event, "ChangeBookKeeperEvent", log); err != nil {
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

// ParseChangeBookKeeperEvent is a log parse operation binding the contract event 0xe60d33488cba3977bf65766cd2f8ac9617f64bf3b3198aff6240ce5c7d43b690.
//
// Solidity: event ChangeBookKeeperEvent(uint256 height, bytes rawHeader)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) ParseChangeBookKeeperEvent(log types.Log) (*EthCrossChainManagerChangeBookKeeperEvent, error) {
	event := new(EthCrossChainManagerChangeBookKeeperEvent)
	if err := _EthCrossChainManager.contract.UnpackLog(event, "ChangeBookKeeperEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainManagerCrossChainEventIterator is returned from FilterCrossChainEvent and is used to iterate over the raw logs and unpacked data for CrossChainEvent events raised by the EthCrossChainManager contract.
type EthCrossChainManagerCrossChainEventIterator struct {
	Event *EthCrossChainManagerCrossChainEvent // Event containing the contract specifics and raw log

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
func (it *EthCrossChainManagerCrossChainEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerCrossChainEvent)
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
		it.Event = new(EthCrossChainManagerCrossChainEvent)
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
func (it *EthCrossChainManagerCrossChainEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerCrossChainEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerCrossChainEvent represents a CrossChainEvent event raised by the EthCrossChainManager contract.
type EthCrossChainManagerCrossChainEvent struct {
	Sender               common.Address
	TxId                 []byte
	ProxyOrAssetContract common.Address
	ToChainId            uint64
	ToContract           []byte
	Rawdata              []byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterCrossChainEvent is a free log retrieval operation binding the contract event 0x6ad3bf15c1988bc04bc153490cab16db8efb9a3990215bf1c64ea6e28be88483.
//
// Solidity: event CrossChainEvent(address indexed sender, bytes txId, address proxyOrAssetContract, uint64 toChainId, bytes toContract, bytes rawdata)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) FilterCrossChainEvent(opts *bind.FilterOpts, sender []common.Address) (*EthCrossChainManagerCrossChainEventIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EthCrossChainManager.contract.FilterLogs(opts, "CrossChainEvent", senderRule)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerCrossChainEventIterator{contract: _EthCrossChainManager.contract, event: "CrossChainEvent", logs: logs, sub: sub}, nil
}

// WatchCrossChainEvent is a free log subscription operation binding the contract event 0x6ad3bf15c1988bc04bc153490cab16db8efb9a3990215bf1c64ea6e28be88483.
//
// Solidity: event CrossChainEvent(address indexed sender, bytes txId, address proxyOrAssetContract, uint64 toChainId, bytes toContract, bytes rawdata)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) WatchCrossChainEvent(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerCrossChainEvent, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EthCrossChainManager.contract.WatchLogs(opts, "CrossChainEvent", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerCrossChainEvent)
				if err := _EthCrossChainManager.contract.UnpackLog(event, "CrossChainEvent", log); err != nil {
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

// ParseCrossChainEvent is a log parse operation binding the contract event 0x6ad3bf15c1988bc04bc153490cab16db8efb9a3990215bf1c64ea6e28be88483.
//
// Solidity: event CrossChainEvent(address indexed sender, bytes txId, address proxyOrAssetContract, uint64 toChainId, bytes toContract, bytes rawdata)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) ParseCrossChainEvent(log types.Log) (*EthCrossChainManagerCrossChainEvent, error) {
	event := new(EthCrossChainManagerCrossChainEvent)
	if err := _EthCrossChainManager.contract.UnpackLog(event, "CrossChainEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainManagerInitGenesisBlockEventIterator is returned from FilterInitGenesisBlockEvent and is used to iterate over the raw logs and unpacked data for InitGenesisBlockEvent events raised by the EthCrossChainManager contract.
type EthCrossChainManagerInitGenesisBlockEventIterator struct {
	Event *EthCrossChainManagerInitGenesisBlockEvent // Event containing the contract specifics and raw log

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
func (it *EthCrossChainManagerInitGenesisBlockEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerInitGenesisBlockEvent)
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
		it.Event = new(EthCrossChainManagerInitGenesisBlockEvent)
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
func (it *EthCrossChainManagerInitGenesisBlockEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerInitGenesisBlockEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerInitGenesisBlockEvent represents a InitGenesisBlockEvent event raised by the EthCrossChainManager contract.
type EthCrossChainManagerInitGenesisBlockEvent struct {
	Height    *big.Int
	RawHeader []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterInitGenesisBlockEvent is a free log retrieval operation binding the contract event 0xf01968fc3a2655cf1b5144cb32de6dc898f91b9239c103744e8457152ab2fbde.
//
// Solidity: event InitGenesisBlockEvent(uint256 height, bytes rawHeader)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) FilterInitGenesisBlockEvent(opts *bind.FilterOpts) (*EthCrossChainManagerInitGenesisBlockEventIterator, error) {

	logs, sub, err := _EthCrossChainManager.contract.FilterLogs(opts, "InitGenesisBlockEvent")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerInitGenesisBlockEventIterator{contract: _EthCrossChainManager.contract, event: "InitGenesisBlockEvent", logs: logs, sub: sub}, nil
}

// WatchInitGenesisBlockEvent is a free log subscription operation binding the contract event 0xf01968fc3a2655cf1b5144cb32de6dc898f91b9239c103744e8457152ab2fbde.
//
// Solidity: event InitGenesisBlockEvent(uint256 height, bytes rawHeader)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) WatchInitGenesisBlockEvent(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerInitGenesisBlockEvent) (event.Subscription, error) {

	logs, sub, err := _EthCrossChainManager.contract.WatchLogs(opts, "InitGenesisBlockEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerInitGenesisBlockEvent)
				if err := _EthCrossChainManager.contract.UnpackLog(event, "InitGenesisBlockEvent", log); err != nil {
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

// ParseInitGenesisBlockEvent is a log parse operation binding the contract event 0xf01968fc3a2655cf1b5144cb32de6dc898f91b9239c103744e8457152ab2fbde.
//
// Solidity: event InitGenesisBlockEvent(uint256 height, bytes rawHeader)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) ParseInitGenesisBlockEvent(log types.Log) (*EthCrossChainManagerInitGenesisBlockEvent, error) {
	event := new(EthCrossChainManagerInitGenesisBlockEvent)
	if err := _EthCrossChainManager.contract.UnpackLog(event, "InitGenesisBlockEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EthCrossChainManager contract.
type EthCrossChainManagerOwnershipTransferredIterator struct {
	Event *EthCrossChainManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EthCrossChainManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerOwnershipTransferred)
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
		it.Event = new(EthCrossChainManagerOwnershipTransferred)
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
func (it *EthCrossChainManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerOwnershipTransferred represents a OwnershipTransferred event raised by the EthCrossChainManager contract.
type EthCrossChainManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EthCrossChainManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EthCrossChainManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerOwnershipTransferredIterator{contract: _EthCrossChainManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EthCrossChainManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerOwnershipTransferred)
				if err := _EthCrossChainManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EthCrossChainManager *EthCrossChainManagerFilterer) ParseOwnershipTransferred(log types.Log) (*EthCrossChainManagerOwnershipTransferred, error) {
	event := new(EthCrossChainManagerOwnershipTransferred)
	if err := _EthCrossChainManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EthCrossChainManager contract.
type EthCrossChainManagerPausedIterator struct {
	Event *EthCrossChainManagerPaused // Event containing the contract specifics and raw log

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
func (it *EthCrossChainManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerPaused)
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
		it.Event = new(EthCrossChainManagerPaused)
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
func (it *EthCrossChainManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerPaused represents a Paused event raised by the EthCrossChainManager contract.
type EthCrossChainManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*EthCrossChainManagerPausedIterator, error) {

	logs, sub, err := _EthCrossChainManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerPausedIterator{contract: _EthCrossChainManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerPaused) (event.Subscription, error) {

	logs, sub, err := _EthCrossChainManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerPaused)
				if err := _EthCrossChainManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_EthCrossChainManager *EthCrossChainManagerFilterer) ParsePaused(log types.Log) (*EthCrossChainManagerPaused, error) {
	event := new(EthCrossChainManagerPaused)
	if err := _EthCrossChainManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EthCrossChainManager contract.
type EthCrossChainManagerUnpausedIterator struct {
	Event *EthCrossChainManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *EthCrossChainManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerUnpaused)
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
		it.Event = new(EthCrossChainManagerUnpaused)
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
func (it *EthCrossChainManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerUnpaused represents a Unpaused event raised by the EthCrossChainManager contract.
type EthCrossChainManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EthCrossChainManagerUnpausedIterator, error) {

	logs, sub, err := _EthCrossChainManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerUnpausedIterator{contract: _EthCrossChainManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _EthCrossChainManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerUnpaused)
				if err := _EthCrossChainManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_EthCrossChainManager *EthCrossChainManagerFilterer) ParseUnpaused(log types.Log) (*EthCrossChainManagerUnpaused, error) {
	event := new(EthCrossChainManagerUnpaused)
	if err := _EthCrossChainManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainManagerVerifyHeaderAndExecuteTxEventIterator is returned from FilterVerifyHeaderAndExecuteTxEvent and is used to iterate over the raw logs and unpacked data for VerifyHeaderAndExecuteTxEvent events raised by the EthCrossChainManager contract.
type EthCrossChainManagerVerifyHeaderAndExecuteTxEventIterator struct {
	Event *EthCrossChainManagerVerifyHeaderAndExecuteTxEvent // Event containing the contract specifics and raw log

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
func (it *EthCrossChainManagerVerifyHeaderAndExecuteTxEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerVerifyHeaderAndExecuteTxEvent)
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
		it.Event = new(EthCrossChainManagerVerifyHeaderAndExecuteTxEvent)
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
func (it *EthCrossChainManagerVerifyHeaderAndExecuteTxEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerVerifyHeaderAndExecuteTxEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerVerifyHeaderAndExecuteTxEvent represents a VerifyHeaderAndExecuteTxEvent event raised by the EthCrossChainManager contract.
type EthCrossChainManagerVerifyHeaderAndExecuteTxEvent struct {
	FromChainID      uint64
	ToContract       []byte
	CrossChainTxHash []byte
	FromChainTxHash  []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterVerifyHeaderAndExecuteTxEvent is a free log retrieval operation binding the contract event 0x8a4a2663ce60ce4955c595da2894de0415240f1ace024cfbff85f513b656bdae.
//
// Solidity: event VerifyHeaderAndExecuteTxEvent(uint64 fromChainID, bytes toContract, bytes crossChainTxHash, bytes fromChainTxHash)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) FilterVerifyHeaderAndExecuteTxEvent(opts *bind.FilterOpts) (*EthCrossChainManagerVerifyHeaderAndExecuteTxEventIterator, error) {

	logs, sub, err := _EthCrossChainManager.contract.FilterLogs(opts, "VerifyHeaderAndExecuteTxEvent")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerVerifyHeaderAndExecuteTxEventIterator{contract: _EthCrossChainManager.contract, event: "VerifyHeaderAndExecuteTxEvent", logs: logs, sub: sub}, nil
}

// WatchVerifyHeaderAndExecuteTxEvent is a free log subscription operation binding the contract event 0x8a4a2663ce60ce4955c595da2894de0415240f1ace024cfbff85f513b656bdae.
//
// Solidity: event VerifyHeaderAndExecuteTxEvent(uint64 fromChainID, bytes toContract, bytes crossChainTxHash, bytes fromChainTxHash)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) WatchVerifyHeaderAndExecuteTxEvent(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerVerifyHeaderAndExecuteTxEvent) (event.Subscription, error) {

	logs, sub, err := _EthCrossChainManager.contract.WatchLogs(opts, "VerifyHeaderAndExecuteTxEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerVerifyHeaderAndExecuteTxEvent)
				if err := _EthCrossChainManager.contract.UnpackLog(event, "VerifyHeaderAndExecuteTxEvent", log); err != nil {
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

// ParseVerifyHeaderAndExecuteTxEvent is a log parse operation binding the contract event 0x8a4a2663ce60ce4955c595da2894de0415240f1ace024cfbff85f513b656bdae.
//
// Solidity: event VerifyHeaderAndExecuteTxEvent(uint64 fromChainID, bytes toContract, bytes crossChainTxHash, bytes fromChainTxHash)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) ParseVerifyHeaderAndExecuteTxEvent(log types.Log) (*EthCrossChainManagerVerifyHeaderAndExecuteTxEvent, error) {
	event := new(EthCrossChainManagerVerifyHeaderAndExecuteTxEvent)
	if err := _EthCrossChainManager.contract.UnpackLog(event, "VerifyHeaderAndExecuteTxEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IEthCrossChainDataABI is the input ABI used to generate the binding from.
const IEthCrossChainDataABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"fromChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"fromChainTx\",\"type\":\"bytes32\"}],\"name\":\"checkIfFromChainTxExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurEpochConPubKeyBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurEpochStartHeight\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ethTxHashIndex\",\"type\":\"uint256\"}],\"name\":\"getEthTxHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEthTxHashIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key2\",\"type\":\"bytes32\"}],\"name\":\"getExtraData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"fromChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"fromChainTx\",\"type\":\"bytes32\"}],\"name\":\"markFromChainTxExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"curEpochPkBytes\",\"type\":\"bytes\"}],\"name\":\"putCurEpochConPubKeyBytes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"curEpochStartHeight\",\"type\":\"uint32\"}],\"name\":\"putCurEpochStartHeight\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ethTxHash\",\"type\":\"bytes32\"}],\"name\":\"putEthTxHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"putExtraData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IEthCrossChainDataFuncSigs maps the 4-byte function signature to its string representation.
var IEthCrossChainDataFuncSigs = map[string]string{
	"0586763c": "checkIfFromChainTxExist(uint64,bytes32)",
	"69d48074": "getCurEpochConPubKeyBytes()",
	"5ac40790": "getCurEpochStartHeight()",
	"29927875": "getEthTxHash(uint256)",
	"ff3d24a1": "getEthTxHashIndex()",
	"40602bb5": "getExtraData(bytes32,bytes32)",
	"e90bfdcf": "markFromChainTxExist(uint64,bytes32)",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"41973cd9": "putCurEpochConPubKeyBytes(bytes)",
	"8a8bd17f": "putCurEpochStartHeight(uint32)",
	"4c3ccf64": "putEthTxHash(bytes32)",
	"1afe374e": "putExtraData(bytes32,bytes32,bytes)",
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

// CheckIfFromChainTxExist is a free data retrieval call binding the contract method 0x0586763c.
//
// Solidity: function checkIfFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) view returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataCaller) CheckIfFromChainTxExist(opts *bind.CallOpts, fromChainId uint64, fromChainTx [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "checkIfFromChainTxExist", fromChainId, fromChainTx)
	return *ret0, err
}

// CheckIfFromChainTxExist is a free data retrieval call binding the contract method 0x0586763c.
//
// Solidity: function checkIfFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) view returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) CheckIfFromChainTxExist(fromChainId uint64, fromChainTx [32]byte) (bool, error) {
	return _IEthCrossChainData.Contract.CheckIfFromChainTxExist(&_IEthCrossChainData.CallOpts, fromChainId, fromChainTx)
}

// CheckIfFromChainTxExist is a free data retrieval call binding the contract method 0x0586763c.
//
// Solidity: function checkIfFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) view returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) CheckIfFromChainTxExist(fromChainId uint64, fromChainTx [32]byte) (bool, error) {
	return _IEthCrossChainData.Contract.CheckIfFromChainTxExist(&_IEthCrossChainData.CallOpts, fromChainId, fromChainTx)
}

// GetCurEpochConPubKeyBytes is a free data retrieval call binding the contract method 0x69d48074.
//
// Solidity: function getCurEpochConPubKeyBytes() view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetCurEpochConPubKeyBytes(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getCurEpochConPubKeyBytes")
	return *ret0, err
}

// GetCurEpochConPubKeyBytes is a free data retrieval call binding the contract method 0x69d48074.
//
// Solidity: function getCurEpochConPubKeyBytes() view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetCurEpochConPubKeyBytes() ([]byte, error) {
	return _IEthCrossChainData.Contract.GetCurEpochConPubKeyBytes(&_IEthCrossChainData.CallOpts)
}

// GetCurEpochConPubKeyBytes is a free data retrieval call binding the contract method 0x69d48074.
//
// Solidity: function getCurEpochConPubKeyBytes() view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetCurEpochConPubKeyBytes() ([]byte, error) {
	return _IEthCrossChainData.Contract.GetCurEpochConPubKeyBytes(&_IEthCrossChainData.CallOpts)
}

// GetCurEpochStartHeight is a free data retrieval call binding the contract method 0x5ac40790.
//
// Solidity: function getCurEpochStartHeight() view returns(uint32)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetCurEpochStartHeight(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getCurEpochStartHeight")
	return *ret0, err
}

// GetCurEpochStartHeight is a free data retrieval call binding the contract method 0x5ac40790.
//
// Solidity: function getCurEpochStartHeight() view returns(uint32)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetCurEpochStartHeight() (uint32, error) {
	return _IEthCrossChainData.Contract.GetCurEpochStartHeight(&_IEthCrossChainData.CallOpts)
}

// GetCurEpochStartHeight is a free data retrieval call binding the contract method 0x5ac40790.
//
// Solidity: function getCurEpochStartHeight() view returns(uint32)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetCurEpochStartHeight() (uint32, error) {
	return _IEthCrossChainData.Contract.GetCurEpochStartHeight(&_IEthCrossChainData.CallOpts)
}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 ethTxHashIndex) view returns(bytes32)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetEthTxHash(opts *bind.CallOpts, ethTxHashIndex *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getEthTxHash", ethTxHashIndex)
	return *ret0, err
}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 ethTxHashIndex) view returns(bytes32)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetEthTxHash(ethTxHashIndex *big.Int) ([32]byte, error) {
	return _IEthCrossChainData.Contract.GetEthTxHash(&_IEthCrossChainData.CallOpts, ethTxHashIndex)
}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 ethTxHashIndex) view returns(bytes32)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetEthTxHash(ethTxHashIndex *big.Int) ([32]byte, error) {
	return _IEthCrossChainData.Contract.GetEthTxHash(&_IEthCrossChainData.CallOpts, ethTxHashIndex)
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

// MarkFromChainTxExist is a paid mutator transaction binding the contract method 0xe90bfdcf.
//
// Solidity: function markFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) MarkFromChainTxExist(opts *bind.TransactOpts, fromChainId uint64, fromChainTx [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "markFromChainTxExist", fromChainId, fromChainTx)
}

// MarkFromChainTxExist is a paid mutator transaction binding the contract method 0xe90bfdcf.
//
// Solidity: function markFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) MarkFromChainTxExist(fromChainId uint64, fromChainTx [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.MarkFromChainTxExist(&_IEthCrossChainData.TransactOpts, fromChainId, fromChainTx)
}

// MarkFromChainTxExist is a paid mutator transaction binding the contract method 0xe90bfdcf.
//
// Solidity: function markFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) MarkFromChainTxExist(fromChainId uint64, fromChainTx [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.MarkFromChainTxExist(&_IEthCrossChainData.TransactOpts, fromChainId, fromChainTx)
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

// PutCurEpochConPubKeyBytes is a paid mutator transaction binding the contract method 0x41973cd9.
//
// Solidity: function putCurEpochConPubKeyBytes(bytes curEpochPkBytes) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutCurEpochConPubKeyBytes(opts *bind.TransactOpts, curEpochPkBytes []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putCurEpochConPubKeyBytes", curEpochPkBytes)
}

// PutCurEpochConPubKeyBytes is a paid mutator transaction binding the contract method 0x41973cd9.
//
// Solidity: function putCurEpochConPubKeyBytes(bytes curEpochPkBytes) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutCurEpochConPubKeyBytes(curEpochPkBytes []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutCurEpochConPubKeyBytes(&_IEthCrossChainData.TransactOpts, curEpochPkBytes)
}

// PutCurEpochConPubKeyBytes is a paid mutator transaction binding the contract method 0x41973cd9.
//
// Solidity: function putCurEpochConPubKeyBytes(bytes curEpochPkBytes) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutCurEpochConPubKeyBytes(curEpochPkBytes []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutCurEpochConPubKeyBytes(&_IEthCrossChainData.TransactOpts, curEpochPkBytes)
}

// PutCurEpochStartHeight is a paid mutator transaction binding the contract method 0x8a8bd17f.
//
// Solidity: function putCurEpochStartHeight(uint32 curEpochStartHeight) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutCurEpochStartHeight(opts *bind.TransactOpts, curEpochStartHeight uint32) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putCurEpochStartHeight", curEpochStartHeight)
}

// PutCurEpochStartHeight is a paid mutator transaction binding the contract method 0x8a8bd17f.
//
// Solidity: function putCurEpochStartHeight(uint32 curEpochStartHeight) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutCurEpochStartHeight(curEpochStartHeight uint32) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutCurEpochStartHeight(&_IEthCrossChainData.TransactOpts, curEpochStartHeight)
}

// PutCurEpochStartHeight is a paid mutator transaction binding the contract method 0x8a8bd17f.
//
// Solidity: function putCurEpochStartHeight(uint32 curEpochStartHeight) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutCurEpochStartHeight(curEpochStartHeight uint32) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutCurEpochStartHeight(&_IEthCrossChainData.TransactOpts, curEpochStartHeight)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x4c3ccf64.
//
// Solidity: function putEthTxHash(bytes32 ethTxHash) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutEthTxHash(opts *bind.TransactOpts, ethTxHash [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putEthTxHash", ethTxHash)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x4c3ccf64.
//
// Solidity: function putEthTxHash(bytes32 ethTxHash) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutEthTxHash(ethTxHash [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutEthTxHash(&_IEthCrossChainData.TransactOpts, ethTxHash)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x4c3ccf64.
//
// Solidity: function putEthTxHash(bytes32 ethTxHash) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutEthTxHash(ethTxHash [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutEthTxHash(&_IEthCrossChainData.TransactOpts, ethTxHash)
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

// IEthCrossChainManagerABI is the input ABI used to generate the binding from.
const IEthCrossChainManagerABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_toContract\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_method\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_txData\",\"type\":\"bytes\"}],\"name\":\"crossChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IEthCrossChainManagerFuncSigs maps the 4-byte function signature to its string representation.
var IEthCrossChainManagerFuncSigs = map[string]string{
	"bd5cf625": "crossChain(uint64,bytes,bytes,bytes)",
}

// IEthCrossChainManager is an auto generated Go binding around an Ethereum contract.
type IEthCrossChainManager struct {
	IEthCrossChainManagerCaller     // Read-only binding to the contract
	IEthCrossChainManagerTransactor // Write-only binding to the contract
	IEthCrossChainManagerFilterer   // Log filterer for contract events
}

// IEthCrossChainManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEthCrossChainManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEthCrossChainManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEthCrossChainManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEthCrossChainManagerSession struct {
	Contract     *IEthCrossChainManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IEthCrossChainManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEthCrossChainManagerCallerSession struct {
	Contract *IEthCrossChainManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IEthCrossChainManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEthCrossChainManagerTransactorSession struct {
	Contract     *IEthCrossChainManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IEthCrossChainManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEthCrossChainManagerRaw struct {
	Contract *IEthCrossChainManager // Generic contract binding to access the raw methods on
}

// IEthCrossChainManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEthCrossChainManagerCallerRaw struct {
	Contract *IEthCrossChainManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IEthCrossChainManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEthCrossChainManagerTransactorRaw struct {
	Contract *IEthCrossChainManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEthCrossChainManager creates a new instance of IEthCrossChainManager, bound to a specific deployed contract.
func NewIEthCrossChainManager(address common.Address, backend bind.ContractBackend) (*IEthCrossChainManager, error) {
	contract, err := bindIEthCrossChainManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManager{IEthCrossChainManagerCaller: IEthCrossChainManagerCaller{contract: contract}, IEthCrossChainManagerTransactor: IEthCrossChainManagerTransactor{contract: contract}, IEthCrossChainManagerFilterer: IEthCrossChainManagerFilterer{contract: contract}}, nil
}

// NewIEthCrossChainManagerCaller creates a new read-only instance of IEthCrossChainManager, bound to a specific deployed contract.
func NewIEthCrossChainManagerCaller(address common.Address, caller bind.ContractCaller) (*IEthCrossChainManagerCaller, error) {
	contract, err := bindIEthCrossChainManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerCaller{contract: contract}, nil
}

// NewIEthCrossChainManagerTransactor creates a new write-only instance of IEthCrossChainManager, bound to a specific deployed contract.
func NewIEthCrossChainManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IEthCrossChainManagerTransactor, error) {
	contract, err := bindIEthCrossChainManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerTransactor{contract: contract}, nil
}

// NewIEthCrossChainManagerFilterer creates a new log filterer instance of IEthCrossChainManager, bound to a specific deployed contract.
func NewIEthCrossChainManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IEthCrossChainManagerFilterer, error) {
	contract, err := bindIEthCrossChainManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerFilterer{contract: contract}, nil
}

// bindIEthCrossChainManager binds a generic wrapper to an already deployed contract.
func bindIEthCrossChainManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IEthCrossChainManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainManager *IEthCrossChainManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainManager.Contract.IEthCrossChainManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainManager *IEthCrossChainManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.IEthCrossChainManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainManager *IEthCrossChainManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.IEthCrossChainManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainManager *IEthCrossChainManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainManager *IEthCrossChainManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainManager *IEthCrossChainManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.contract.Transact(opts, method, params...)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) returns(bool)
func (_IEthCrossChainManager *IEthCrossChainManagerTransactor) CrossChain(opts *bind.TransactOpts, _toChainId uint64, _toContract []byte, _method []byte, _txData []byte) (*types.Transaction, error) {
	return _IEthCrossChainManager.contract.Transact(opts, "crossChain", _toChainId, _toContract, _method, _txData)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) returns(bool)
func (_IEthCrossChainManager *IEthCrossChainManagerSession) CrossChain(_toChainId uint64, _toContract []byte, _method []byte, _txData []byte) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.CrossChain(&_IEthCrossChainManager.TransactOpts, _toChainId, _toContract, _method, _txData)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) returns(bool)
func (_IEthCrossChainManager *IEthCrossChainManagerTransactorSession) CrossChain(_toChainId uint64, _toContract []byte, _method []byte, _txData []byte) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.CrossChain(&_IEthCrossChainManager.TransactOpts, _toChainId, _toContract, _method, _txData)
}

// IUpgradableECCMABI is the input ABI used to generate the binding from.
const IUpgradableECCMABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_newChainId\",\"type\":\"uint64\"}],\"name\":\"setChainId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"upgradeToNew\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IUpgradableECCMFuncSigs maps the 4-byte function signature to its string representation.
var IUpgradableECCMFuncSigs = map[string]string{
	"8f32d59b": "isOwner()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"6f31031d": "setChainId(uint64)",
	"3f4ba83a": "unpause()",
	"7e724ff3": "upgradeToNew(address)",
}

// IUpgradableECCM is an auto generated Go binding around an Ethereum contract.
type IUpgradableECCM struct {
	IUpgradableECCMCaller     // Read-only binding to the contract
	IUpgradableECCMTransactor // Write-only binding to the contract
	IUpgradableECCMFilterer   // Log filterer for contract events
}

// IUpgradableECCMCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUpgradableECCMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUpgradableECCMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUpgradableECCMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUpgradableECCMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUpgradableECCMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUpgradableECCMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUpgradableECCMSession struct {
	Contract     *IUpgradableECCM  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IUpgradableECCMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUpgradableECCMCallerSession struct {
	Contract *IUpgradableECCMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IUpgradableECCMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUpgradableECCMTransactorSession struct {
	Contract     *IUpgradableECCMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IUpgradableECCMRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUpgradableECCMRaw struct {
	Contract *IUpgradableECCM // Generic contract binding to access the raw methods on
}

// IUpgradableECCMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUpgradableECCMCallerRaw struct {
	Contract *IUpgradableECCMCaller // Generic read-only contract binding to access the raw methods on
}

// IUpgradableECCMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUpgradableECCMTransactorRaw struct {
	Contract *IUpgradableECCMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUpgradableECCM creates a new instance of IUpgradableECCM, bound to a specific deployed contract.
func NewIUpgradableECCM(address common.Address, backend bind.ContractBackend) (*IUpgradableECCM, error) {
	contract, err := bindIUpgradableECCM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUpgradableECCM{IUpgradableECCMCaller: IUpgradableECCMCaller{contract: contract}, IUpgradableECCMTransactor: IUpgradableECCMTransactor{contract: contract}, IUpgradableECCMFilterer: IUpgradableECCMFilterer{contract: contract}}, nil
}

// NewIUpgradableECCMCaller creates a new read-only instance of IUpgradableECCM, bound to a specific deployed contract.
func NewIUpgradableECCMCaller(address common.Address, caller bind.ContractCaller) (*IUpgradableECCMCaller, error) {
	contract, err := bindIUpgradableECCM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUpgradableECCMCaller{contract: contract}, nil
}

// NewIUpgradableECCMTransactor creates a new write-only instance of IUpgradableECCM, bound to a specific deployed contract.
func NewIUpgradableECCMTransactor(address common.Address, transactor bind.ContractTransactor) (*IUpgradableECCMTransactor, error) {
	contract, err := bindIUpgradableECCM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUpgradableECCMTransactor{contract: contract}, nil
}

// NewIUpgradableECCMFilterer creates a new log filterer instance of IUpgradableECCM, bound to a specific deployed contract.
func NewIUpgradableECCMFilterer(address common.Address, filterer bind.ContractFilterer) (*IUpgradableECCMFilterer, error) {
	contract, err := bindIUpgradableECCM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUpgradableECCMFilterer{contract: contract}, nil
}

// bindIUpgradableECCM binds a generic wrapper to an already deployed contract.
func bindIUpgradableECCM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUpgradableECCMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUpgradableECCM *IUpgradableECCMRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IUpgradableECCM.Contract.IUpgradableECCMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUpgradableECCM *IUpgradableECCMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.IUpgradableECCMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUpgradableECCM *IUpgradableECCMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.IUpgradableECCMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUpgradableECCM *IUpgradableECCMCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IUpgradableECCM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUpgradableECCM *IUpgradableECCMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUpgradableECCM *IUpgradableECCMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_IUpgradableECCM *IUpgradableECCMCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IUpgradableECCM.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_IUpgradableECCM *IUpgradableECCMSession) IsOwner() (bool, error) {
	return _IUpgradableECCM.Contract.IsOwner(&_IUpgradableECCM.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_IUpgradableECCM *IUpgradableECCMCallerSession) IsOwner() (bool, error) {
	return _IUpgradableECCM.Contract.IsOwner(&_IUpgradableECCM.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IUpgradableECCM *IUpgradableECCMCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IUpgradableECCM.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IUpgradableECCM *IUpgradableECCMSession) Paused() (bool, error) {
	return _IUpgradableECCM.Contract.Paused(&_IUpgradableECCM.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IUpgradableECCM *IUpgradableECCMCallerSession) Paused() (bool, error) {
	return _IUpgradableECCM.Contract.Paused(&_IUpgradableECCM.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_IUpgradableECCM *IUpgradableECCMTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUpgradableECCM.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_IUpgradableECCM *IUpgradableECCMSession) Pause() (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.Pause(&_IUpgradableECCM.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_IUpgradableECCM *IUpgradableECCMTransactorSession) Pause() (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.Pause(&_IUpgradableECCM.TransactOpts)
}

// SetChainId is a paid mutator transaction binding the contract method 0x6f31031d.
//
// Solidity: function setChainId(uint64 _newChainId) returns(bool)
func (_IUpgradableECCM *IUpgradableECCMTransactor) SetChainId(opts *bind.TransactOpts, _newChainId uint64) (*types.Transaction, error) {
	return _IUpgradableECCM.contract.Transact(opts, "setChainId", _newChainId)
}

// SetChainId is a paid mutator transaction binding the contract method 0x6f31031d.
//
// Solidity: function setChainId(uint64 _newChainId) returns(bool)
func (_IUpgradableECCM *IUpgradableECCMSession) SetChainId(_newChainId uint64) (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.SetChainId(&_IUpgradableECCM.TransactOpts, _newChainId)
}

// SetChainId is a paid mutator transaction binding the contract method 0x6f31031d.
//
// Solidity: function setChainId(uint64 _newChainId) returns(bool)
func (_IUpgradableECCM *IUpgradableECCMTransactorSession) SetChainId(_newChainId uint64) (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.SetChainId(&_IUpgradableECCM.TransactOpts, _newChainId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_IUpgradableECCM *IUpgradableECCMTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUpgradableECCM.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_IUpgradableECCM *IUpgradableECCMSession) Unpause() (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.Unpause(&_IUpgradableECCM.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_IUpgradableECCM *IUpgradableECCMTransactorSession) Unpause() (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.Unpause(&_IUpgradableECCM.TransactOpts)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address ) returns(bool)
func (_IUpgradableECCM *IUpgradableECCMTransactor) UpgradeToNew(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IUpgradableECCM.contract.Transact(opts, "upgradeToNew", arg0)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address ) returns(bool)
func (_IUpgradableECCM *IUpgradableECCMSession) UpgradeToNew(arg0 common.Address) (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.UpgradeToNew(&_IUpgradableECCM.TransactOpts, arg0)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address ) returns(bool)
func (_IUpgradableECCM *IUpgradableECCMTransactorSession) UpgradeToNew(arg0 common.Address) (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.UpgradeToNew(&_IUpgradableECCM.TransactOpts, arg0)
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158208dcea9a9a54063ebe733acda3d5b2886fa847c60bc371f7f71b46ec7553bdd5e64736f6c63430005110032"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// UpgradableECCMABI is the input ABI used to generate the binding from.
const UpgradableECCMABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethCrossChainDataAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_chainId\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"EthCrossChainDataAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_newChainId\",\"type\":\"uint64\"}],\"name\":\"setChainId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newEthCrossChainManagerAddress\",\"type\":\"address\"}],\"name\":\"upgradeToNew\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// UpgradableECCMFuncSigs maps the 4-byte function signature to its string representation.
var UpgradableECCMFuncSigs = map[string]string{
	"00ba1694": "EthCrossChainDataAddress()",
	"9a8a0592": "chainId()",
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"715018a6": "renounceOwnership()",
	"6f31031d": "setChainId(uint64)",
	"f2fde38b": "transferOwnership(address)",
	"3f4ba83a": "unpause()",
	"7e724ff3": "upgradeToNew(address)",
}

// UpgradableECCMBin is the compiled bytecode used for deploying new contracts.
var UpgradableECCMBin = "0x608060405234801561001057600080fd5b50604051610b6c380380610b6c8339818101604052604081101561003357600080fd5b508051602090910151600061004f6001600160e01b036100ee16565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506000805460ff60a01b19169055600180546001600160401b03909216600160a01b02600160a01b600160e01b03196001600160a01b039094166001600160a01b031990931692909217929092161790556100f2565b3390565b610a6b806101016000396000f3fe608060405234801561001057600080fd5b50600436106100a85760003560e01c80637e724ff3116100715780637e724ff3146101265780638456cb591461014c5780638da5cb5b146101545780638f32d59b1461015c5780639a8a059214610164578063f2fde38b14610189576100a8565b8062ba1694146100ad5780633f4ba83a146100d15780635c975abb146100ed5780636f31031d146100f5578063715018a61461011c575b600080fd5b6100b56101af565b604080516001600160a01b039092168252519081900360200190f35b6100d96101be565b604080519115158252519081900360200190f35b6100d961033d565b6100d96004803603602081101561010b57600080fd5b503567ffffffffffffffff1661034d565b610124610418565b005b6100d96004803603602081101561013c57600080fd5b50356001600160a01b03166104a9565b6100d96105b8565b6100b561072d565b6100d961073c565b61016c610760565b6040805167ffffffffffffffff9092168252519081900360200190f35b6101246004803603602081101561019f57600080fd5b50356001600160a01b0316610777565b6001546001600160a01b031681565b60006101c861073c565b610207576040805162461bcd60e51b81526020600482018190526024820152600080516020610a17833981519152604482015290519081900360640190fd5b61020f61033d565b1561021c5761021c6107ca565b60015460408051635c975abb60e01b815290516001600160a01b03909216918291635c975abb916004808301926020929190829003018186803b15801561026257600080fd5b505afa158015610276573d6000803e3d6000fd5b505050506040513d602081101561028c57600080fd5b50511561033557806001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156102ce57600080fd5b505af11580156102e2573d6000803e3d6000fd5b505050506040513d60208110156102f857600080fd5b50516103355760405162461bcd60e51b81526004018080602001828103825260298152602001806109c86029913960400191505060405180910390fd5b600191505090565b600054600160a01b900460ff1690565b60008054600160a01b900460ff166103a3576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b6103ab61073c565b6103ea576040805162461bcd60e51b81526020600482018190526024820152600080516020610a17833981519152604482015290519081900360640190fd5b506001805467ffffffffffffffff8316600160a01b0267ffffffffffffffff60a01b19909116178155919050565b61042061073c565b61045f576040805162461bcd60e51b81526020600482018190526024820152600080516020610a17833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60008054600160a01b900460ff166104ff576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b61050761073c565b610546576040805162461bcd60e51b81526020600482018190526024820152600080516020610a17833981519152604482015290519081900360640190fd5b6001546040805163f2fde38b60e01b81526001600160a01b03858116600483015291519190921691829163f2fde38b9160248082019260009290919082900301818387803b15801561059757600080fd5b505af11580156105ab573d6000803e3d6000fd5b5060019695505050505050565b60006105c261073c565b610601576040805162461bcd60e51b81526020600482018190526024820152600080516020610a17833981519152604482015290519081900360640190fd5b61060961033d565b61061557610615610872565b60015460408051635c975abb60e01b815290516001600160a01b03909216918291635c975abb916004808301926020929190829003018186803b15801561065b57600080fd5b505afa15801561066f573d6000803e3d6000fd5b505050506040513d602081101561068557600080fd5b505161033557806001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156106c657600080fd5b505af11580156106da573d6000803e3d6000fd5b505050506040513d60208110156106f057600080fd5b50516103355760405162461bcd60e51b81526004018080602001828103825260278152602001806109a16027913960400191505060405180910390fd5b6000546001600160a01b031690565b600080546001600160a01b03166107516108fc565b6001600160a01b031614905090565b600154600160a01b900467ffffffffffffffff1681565b61077f61073c565b6107be576040805162461bcd60e51b81526020600482018190526024820152600080516020610a17833981519152604482015290519081900360640190fd5b6107c781610900565b50565b600054600160a01b900460ff1661081f576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6108556108fc565b604080516001600160a01b039092168252519081900360200190a1565b600054600160a01b900460ff16156108c4576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586108555b3390565b6001600160a01b0381166109455760405162461bcd60e51b81526004018080602001828103825260268152602001806109f16026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fe70617573652045746843726f7373436861696e4461746120636f6e7472616374206661696c6564756e70617573652045746843726f7373436861696e4461746120636f6e7472616374206661696c65644f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a265627a7a723158201582793e17f9bf88c24fcafae155e42e6cd8395013de73a81041a01550bf2e2864736f6c63430005110032"

// DeployUpgradableECCM deploys a new Ethereum contract, binding an instance of UpgradableECCM to it.
func DeployUpgradableECCM(auth *bind.TransactOpts, backend bind.ContractBackend, ethCrossChainDataAddr common.Address, _chainId uint64) (common.Address, *types.Transaction, *UpgradableECCM, error) {
	parsed, err := abi.JSON(strings.NewReader(UpgradableECCMABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UpgradableECCMBin), backend, ethCrossChainDataAddr, _chainId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UpgradableECCM{UpgradableECCMCaller: UpgradableECCMCaller{contract: contract}, UpgradableECCMTransactor: UpgradableECCMTransactor{contract: contract}, UpgradableECCMFilterer: UpgradableECCMFilterer{contract: contract}}, nil
}

// UpgradableECCM is an auto generated Go binding around an Ethereum contract.
type UpgradableECCM struct {
	UpgradableECCMCaller     // Read-only binding to the contract
	UpgradableECCMTransactor // Write-only binding to the contract
	UpgradableECCMFilterer   // Log filterer for contract events
}

// UpgradableECCMCaller is an auto generated read-only Go binding around an Ethereum contract.
type UpgradableECCMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradableECCMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UpgradableECCMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradableECCMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UpgradableECCMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradableECCMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UpgradableECCMSession struct {
	Contract     *UpgradableECCM   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UpgradableECCMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UpgradableECCMCallerSession struct {
	Contract *UpgradableECCMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// UpgradableECCMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UpgradableECCMTransactorSession struct {
	Contract     *UpgradableECCMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// UpgradableECCMRaw is an auto generated low-level Go binding around an Ethereum contract.
type UpgradableECCMRaw struct {
	Contract *UpgradableECCM // Generic contract binding to access the raw methods on
}

// UpgradableECCMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UpgradableECCMCallerRaw struct {
	Contract *UpgradableECCMCaller // Generic read-only contract binding to access the raw methods on
}

// UpgradableECCMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UpgradableECCMTransactorRaw struct {
	Contract *UpgradableECCMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpgradableECCM creates a new instance of UpgradableECCM, bound to a specific deployed contract.
func NewUpgradableECCM(address common.Address, backend bind.ContractBackend) (*UpgradableECCM, error) {
	contract, err := bindUpgradableECCM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UpgradableECCM{UpgradableECCMCaller: UpgradableECCMCaller{contract: contract}, UpgradableECCMTransactor: UpgradableECCMTransactor{contract: contract}, UpgradableECCMFilterer: UpgradableECCMFilterer{contract: contract}}, nil
}

// NewUpgradableECCMCaller creates a new read-only instance of UpgradableECCM, bound to a specific deployed contract.
func NewUpgradableECCMCaller(address common.Address, caller bind.ContractCaller) (*UpgradableECCMCaller, error) {
	contract, err := bindUpgradableECCM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradableECCMCaller{contract: contract}, nil
}

// NewUpgradableECCMTransactor creates a new write-only instance of UpgradableECCM, bound to a specific deployed contract.
func NewUpgradableECCMTransactor(address common.Address, transactor bind.ContractTransactor) (*UpgradableECCMTransactor, error) {
	contract, err := bindUpgradableECCM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradableECCMTransactor{contract: contract}, nil
}

// NewUpgradableECCMFilterer creates a new log filterer instance of UpgradableECCM, bound to a specific deployed contract.
func NewUpgradableECCMFilterer(address common.Address, filterer bind.ContractFilterer) (*UpgradableECCMFilterer, error) {
	contract, err := bindUpgradableECCM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpgradableECCMFilterer{contract: contract}, nil
}

// bindUpgradableECCM binds a generic wrapper to an already deployed contract.
func bindUpgradableECCM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UpgradableECCMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradableECCM *UpgradableECCMRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UpgradableECCM.Contract.UpgradableECCMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradableECCM *UpgradableECCMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.UpgradableECCMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradableECCM *UpgradableECCMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.UpgradableECCMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradableECCM *UpgradableECCMCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UpgradableECCM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradableECCM *UpgradableECCMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradableECCM *UpgradableECCMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.contract.Transact(opts, method, params...)
}

// EthCrossChainDataAddress is a free data retrieval call binding the contract method 0x00ba1694.
//
// Solidity: function EthCrossChainDataAddress() view returns(address)
func (_UpgradableECCM *UpgradableECCMCaller) EthCrossChainDataAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UpgradableECCM.contract.Call(opts, out, "EthCrossChainDataAddress")
	return *ret0, err
}

// EthCrossChainDataAddress is a free data retrieval call binding the contract method 0x00ba1694.
//
// Solidity: function EthCrossChainDataAddress() view returns(address)
func (_UpgradableECCM *UpgradableECCMSession) EthCrossChainDataAddress() (common.Address, error) {
	return _UpgradableECCM.Contract.EthCrossChainDataAddress(&_UpgradableECCM.CallOpts)
}

// EthCrossChainDataAddress is a free data retrieval call binding the contract method 0x00ba1694.
//
// Solidity: function EthCrossChainDataAddress() view returns(address)
func (_UpgradableECCM *UpgradableECCMCallerSession) EthCrossChainDataAddress() (common.Address, error) {
	return _UpgradableECCM.Contract.EthCrossChainDataAddress(&_UpgradableECCM.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_UpgradableECCM *UpgradableECCMCaller) ChainId(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _UpgradableECCM.contract.Call(opts, out, "chainId")
	return *ret0, err
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_UpgradableECCM *UpgradableECCMSession) ChainId() (uint64, error) {
	return _UpgradableECCM.Contract.ChainId(&_UpgradableECCM.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_UpgradableECCM *UpgradableECCMCallerSession) ChainId() (uint64, error) {
	return _UpgradableECCM.Contract.ChainId(&_UpgradableECCM.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_UpgradableECCM *UpgradableECCMCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _UpgradableECCM.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_UpgradableECCM *UpgradableECCMSession) IsOwner() (bool, error) {
	return _UpgradableECCM.Contract.IsOwner(&_UpgradableECCM.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_UpgradableECCM *UpgradableECCMCallerSession) IsOwner() (bool, error) {
	return _UpgradableECCM.Contract.IsOwner(&_UpgradableECCM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UpgradableECCM *UpgradableECCMCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UpgradableECCM.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UpgradableECCM *UpgradableECCMSession) Owner() (common.Address, error) {
	return _UpgradableECCM.Contract.Owner(&_UpgradableECCM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UpgradableECCM *UpgradableECCMCallerSession) Owner() (common.Address, error) {
	return _UpgradableECCM.Contract.Owner(&_UpgradableECCM.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_UpgradableECCM *UpgradableECCMCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _UpgradableECCM.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_UpgradableECCM *UpgradableECCMSession) Paused() (bool, error) {
	return _UpgradableECCM.Contract.Paused(&_UpgradableECCM.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_UpgradableECCM *UpgradableECCMCallerSession) Paused() (bool, error) {
	return _UpgradableECCM.Contract.Paused(&_UpgradableECCM.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_UpgradableECCM *UpgradableECCMTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradableECCM.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_UpgradableECCM *UpgradableECCMSession) Pause() (*types.Transaction, error) {
	return _UpgradableECCM.Contract.Pause(&_UpgradableECCM.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_UpgradableECCM *UpgradableECCMTransactorSession) Pause() (*types.Transaction, error) {
	return _UpgradableECCM.Contract.Pause(&_UpgradableECCM.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UpgradableECCM *UpgradableECCMTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradableECCM.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UpgradableECCM *UpgradableECCMSession) RenounceOwnership() (*types.Transaction, error) {
	return _UpgradableECCM.Contract.RenounceOwnership(&_UpgradableECCM.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UpgradableECCM *UpgradableECCMTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UpgradableECCM.Contract.RenounceOwnership(&_UpgradableECCM.TransactOpts)
}

// SetChainId is a paid mutator transaction binding the contract method 0x6f31031d.
//
// Solidity: function setChainId(uint64 _newChainId) returns(bool)
func (_UpgradableECCM *UpgradableECCMTransactor) SetChainId(opts *bind.TransactOpts, _newChainId uint64) (*types.Transaction, error) {
	return _UpgradableECCM.contract.Transact(opts, "setChainId", _newChainId)
}

// SetChainId is a paid mutator transaction binding the contract method 0x6f31031d.
//
// Solidity: function setChainId(uint64 _newChainId) returns(bool)
func (_UpgradableECCM *UpgradableECCMSession) SetChainId(_newChainId uint64) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.SetChainId(&_UpgradableECCM.TransactOpts, _newChainId)
}

// SetChainId is a paid mutator transaction binding the contract method 0x6f31031d.
//
// Solidity: function setChainId(uint64 _newChainId) returns(bool)
func (_UpgradableECCM *UpgradableECCMTransactorSession) SetChainId(_newChainId uint64) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.SetChainId(&_UpgradableECCM.TransactOpts, _newChainId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpgradableECCM *UpgradableECCMTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UpgradableECCM.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpgradableECCM *UpgradableECCMSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.TransferOwnership(&_UpgradableECCM.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpgradableECCM *UpgradableECCMTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.TransferOwnership(&_UpgradableECCM.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_UpgradableECCM *UpgradableECCMTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradableECCM.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_UpgradableECCM *UpgradableECCMSession) Unpause() (*types.Transaction, error) {
	return _UpgradableECCM.Contract.Unpause(&_UpgradableECCM.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_UpgradableECCM *UpgradableECCMTransactorSession) Unpause() (*types.Transaction, error) {
	return _UpgradableECCM.Contract.Unpause(&_UpgradableECCM.TransactOpts)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address newEthCrossChainManagerAddress) returns(bool)
func (_UpgradableECCM *UpgradableECCMTransactor) UpgradeToNew(opts *bind.TransactOpts, newEthCrossChainManagerAddress common.Address) (*types.Transaction, error) {
	return _UpgradableECCM.contract.Transact(opts, "upgradeToNew", newEthCrossChainManagerAddress)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address newEthCrossChainManagerAddress) returns(bool)
func (_UpgradableECCM *UpgradableECCMSession) UpgradeToNew(newEthCrossChainManagerAddress common.Address) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.UpgradeToNew(&_UpgradableECCM.TransactOpts, newEthCrossChainManagerAddress)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address newEthCrossChainManagerAddress) returns(bool)
func (_UpgradableECCM *UpgradableECCMTransactorSession) UpgradeToNew(newEthCrossChainManagerAddress common.Address) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.UpgradeToNew(&_UpgradableECCM.TransactOpts, newEthCrossChainManagerAddress)
}

// UpgradableECCMOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UpgradableECCM contract.
type UpgradableECCMOwnershipTransferredIterator struct {
	Event *UpgradableECCMOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UpgradableECCMOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradableECCMOwnershipTransferred)
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
		it.Event = new(UpgradableECCMOwnershipTransferred)
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
func (it *UpgradableECCMOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradableECCMOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradableECCMOwnershipTransferred represents a OwnershipTransferred event raised by the UpgradableECCM contract.
type UpgradableECCMOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UpgradableECCM *UpgradableECCMFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UpgradableECCMOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UpgradableECCM.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UpgradableECCMOwnershipTransferredIterator{contract: _UpgradableECCM.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UpgradableECCM *UpgradableECCMFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UpgradableECCMOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UpgradableECCM.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradableECCMOwnershipTransferred)
				if err := _UpgradableECCM.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_UpgradableECCM *UpgradableECCMFilterer) ParseOwnershipTransferred(log types.Log) (*UpgradableECCMOwnershipTransferred, error) {
	event := new(UpgradableECCMOwnershipTransferred)
	if err := _UpgradableECCM.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UpgradableECCMPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the UpgradableECCM contract.
type UpgradableECCMPausedIterator struct {
	Event *UpgradableECCMPaused // Event containing the contract specifics and raw log

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
func (it *UpgradableECCMPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradableECCMPaused)
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
		it.Event = new(UpgradableECCMPaused)
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
func (it *UpgradableECCMPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradableECCMPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradableECCMPaused represents a Paused event raised by the UpgradableECCM contract.
type UpgradableECCMPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_UpgradableECCM *UpgradableECCMFilterer) FilterPaused(opts *bind.FilterOpts) (*UpgradableECCMPausedIterator, error) {

	logs, sub, err := _UpgradableECCM.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &UpgradableECCMPausedIterator{contract: _UpgradableECCM.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_UpgradableECCM *UpgradableECCMFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *UpgradableECCMPaused) (event.Subscription, error) {

	logs, sub, err := _UpgradableECCM.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradableECCMPaused)
				if err := _UpgradableECCM.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_UpgradableECCM *UpgradableECCMFilterer) ParsePaused(log types.Log) (*UpgradableECCMPaused, error) {
	event := new(UpgradableECCMPaused)
	if err := _UpgradableECCM.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UpgradableECCMUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the UpgradableECCM contract.
type UpgradableECCMUnpausedIterator struct {
	Event *UpgradableECCMUnpaused // Event containing the contract specifics and raw log

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
func (it *UpgradableECCMUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradableECCMUnpaused)
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
		it.Event = new(UpgradableECCMUnpaused)
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
func (it *UpgradableECCMUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradableECCMUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradableECCMUnpaused represents a Unpaused event raised by the UpgradableECCM contract.
type UpgradableECCMUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_UpgradableECCM *UpgradableECCMFilterer) FilterUnpaused(opts *bind.FilterOpts) (*UpgradableECCMUnpausedIterator, error) {

	logs, sub, err := _UpgradableECCM.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &UpgradableECCMUnpausedIterator{contract: _UpgradableECCM.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_UpgradableECCM *UpgradableECCMFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *UpgradableECCMUnpaused) (event.Subscription, error) {

	logs, sub, err := _UpgradableECCM.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradableECCMUnpaused)
				if err := _UpgradableECCM.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_UpgradableECCM *UpgradableECCMFilterer) ParseUnpaused(log types.Log) (*UpgradableECCMUnpaused, error) {
	event := new(UpgradableECCMUnpaused)
	if err := _UpgradableECCM.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UtilsABI is the input ABI used to generate the binding from.
const UtilsABI = "[]"

// UtilsBin is the compiled bytecode used for deploying new contracts.
var UtilsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158205f9120778c140a57a599e3f429150b29f8ef5702675093c8d587f2c2d5c8994a64736f6c63430005110032"

// DeployUtils deploys a new Ethereum contract, binding an instance of Utils to it.
func DeployUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Utils, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// Utils is an auto generated Go binding around an Ethereum contract.
type Utils struct {
	UtilsCaller     // Read-only binding to the contract
	UtilsTransactor // Write-only binding to the contract
	UtilsFilterer   // Log filterer for contract events
}

// UtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilsSession struct {
	Contract     *Utils            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilsCallerSession struct {
	Contract *UtilsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilsTransactorSession struct {
	Contract     *UtilsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilsRaw struct {
	Contract *Utils // Generic contract binding to access the raw methods on
}

// UtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilsCallerRaw struct {
	Contract *UtilsCaller // Generic read-only contract binding to access the raw methods on
}

// UtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilsTransactorRaw struct {
	Contract *UtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtils creates a new instance of Utils, bound to a specific deployed contract.
func NewUtils(address common.Address, backend bind.ContractBackend) (*Utils, error) {
	contract, err := bindUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// NewUtilsCaller creates a new read-only instance of Utils, bound to a specific deployed contract.
func NewUtilsCaller(address common.Address, caller bind.ContractCaller) (*UtilsCaller, error) {
	contract, err := bindUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsCaller{contract: contract}, nil
}

// NewUtilsTransactor creates a new write-only instance of Utils, bound to a specific deployed contract.
func NewUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilsTransactor, error) {
	contract, err := bindUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsTransactor{contract: contract}, nil
}

// NewUtilsFilterer creates a new log filterer instance of Utils, bound to a specific deployed contract.
func NewUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilsFilterer, error) {
	contract, err := bindUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilsFilterer{contract: contract}, nil
}

// bindUtils binds a generic wrapper to an already deployed contract.
func bindUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.UtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transact(opts, method, params...)
}

// ZeroCopySinkABI is the input ABI used to generate the binding from.
const ZeroCopySinkABI = "[]"

// ZeroCopySinkBin is the compiled bytecode used for deploying new contracts.
var ZeroCopySinkBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820283c467f4d3857355a17e6c56f5021c71ae4de199154b010340642aefecc31bc64736f6c63430005110032"

// DeployZeroCopySink deploys a new Ethereum contract, binding an instance of ZeroCopySink to it.
func DeployZeroCopySink(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZeroCopySink, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroCopySinkABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZeroCopySinkBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZeroCopySink{ZeroCopySinkCaller: ZeroCopySinkCaller{contract: contract}, ZeroCopySinkTransactor: ZeroCopySinkTransactor{contract: contract}, ZeroCopySinkFilterer: ZeroCopySinkFilterer{contract: contract}}, nil
}

// ZeroCopySink is an auto generated Go binding around an Ethereum contract.
type ZeroCopySink struct {
	ZeroCopySinkCaller     // Read-only binding to the contract
	ZeroCopySinkTransactor // Write-only binding to the contract
	ZeroCopySinkFilterer   // Log filterer for contract events
}

// ZeroCopySinkCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZeroCopySinkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySinkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZeroCopySinkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySinkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZeroCopySinkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySinkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZeroCopySinkSession struct {
	Contract     *ZeroCopySink     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZeroCopySinkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZeroCopySinkCallerSession struct {
	Contract *ZeroCopySinkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ZeroCopySinkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZeroCopySinkTransactorSession struct {
	Contract     *ZeroCopySinkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ZeroCopySinkRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZeroCopySinkRaw struct {
	Contract *ZeroCopySink // Generic contract binding to access the raw methods on
}

// ZeroCopySinkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZeroCopySinkCallerRaw struct {
	Contract *ZeroCopySinkCaller // Generic read-only contract binding to access the raw methods on
}

// ZeroCopySinkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZeroCopySinkTransactorRaw struct {
	Contract *ZeroCopySinkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZeroCopySink creates a new instance of ZeroCopySink, bound to a specific deployed contract.
func NewZeroCopySink(address common.Address, backend bind.ContractBackend) (*ZeroCopySink, error) {
	contract, err := bindZeroCopySink(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySink{ZeroCopySinkCaller: ZeroCopySinkCaller{contract: contract}, ZeroCopySinkTransactor: ZeroCopySinkTransactor{contract: contract}, ZeroCopySinkFilterer: ZeroCopySinkFilterer{contract: contract}}, nil
}

// NewZeroCopySinkCaller creates a new read-only instance of ZeroCopySink, bound to a specific deployed contract.
func NewZeroCopySinkCaller(address common.Address, caller bind.ContractCaller) (*ZeroCopySinkCaller, error) {
	contract, err := bindZeroCopySink(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySinkCaller{contract: contract}, nil
}

// NewZeroCopySinkTransactor creates a new write-only instance of ZeroCopySink, bound to a specific deployed contract.
func NewZeroCopySinkTransactor(address common.Address, transactor bind.ContractTransactor) (*ZeroCopySinkTransactor, error) {
	contract, err := bindZeroCopySink(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySinkTransactor{contract: contract}, nil
}

// NewZeroCopySinkFilterer creates a new log filterer instance of ZeroCopySink, bound to a specific deployed contract.
func NewZeroCopySinkFilterer(address common.Address, filterer bind.ContractFilterer) (*ZeroCopySinkFilterer, error) {
	contract, err := bindZeroCopySink(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySinkFilterer{contract: contract}, nil
}

// bindZeroCopySink binds a generic wrapper to an already deployed contract.
func bindZeroCopySink(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroCopySinkABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroCopySink *ZeroCopySinkRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZeroCopySink.Contract.ZeroCopySinkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroCopySink *ZeroCopySinkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroCopySink.Contract.ZeroCopySinkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroCopySink *ZeroCopySinkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroCopySink.Contract.ZeroCopySinkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroCopySink *ZeroCopySinkCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZeroCopySink.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroCopySink *ZeroCopySinkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroCopySink.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroCopySink *ZeroCopySinkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroCopySink.Contract.contract.Transact(opts, method, params...)
}

// ZeroCopySourceABI is the input ABI used to generate the binding from.
const ZeroCopySourceABI = "[]"

// ZeroCopySourceBin is the compiled bytecode used for deploying new contracts.
var ZeroCopySourceBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158209555b9835947880b0164932574379af90a2d92e58ded2092cca4c290a4151eb964736f6c63430005110032"

// DeployZeroCopySource deploys a new Ethereum contract, binding an instance of ZeroCopySource to it.
func DeployZeroCopySource(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZeroCopySource, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroCopySourceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZeroCopySourceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZeroCopySource{ZeroCopySourceCaller: ZeroCopySourceCaller{contract: contract}, ZeroCopySourceTransactor: ZeroCopySourceTransactor{contract: contract}, ZeroCopySourceFilterer: ZeroCopySourceFilterer{contract: contract}}, nil
}

// ZeroCopySource is an auto generated Go binding around an Ethereum contract.
type ZeroCopySource struct {
	ZeroCopySourceCaller     // Read-only binding to the contract
	ZeroCopySourceTransactor // Write-only binding to the contract
	ZeroCopySourceFilterer   // Log filterer for contract events
}

// ZeroCopySourceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZeroCopySourceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySourceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZeroCopySourceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySourceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZeroCopySourceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySourceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZeroCopySourceSession struct {
	Contract     *ZeroCopySource   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZeroCopySourceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZeroCopySourceCallerSession struct {
	Contract *ZeroCopySourceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ZeroCopySourceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZeroCopySourceTransactorSession struct {
	Contract     *ZeroCopySourceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ZeroCopySourceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZeroCopySourceRaw struct {
	Contract *ZeroCopySource // Generic contract binding to access the raw methods on
}

// ZeroCopySourceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZeroCopySourceCallerRaw struct {
	Contract *ZeroCopySourceCaller // Generic read-only contract binding to access the raw methods on
}

// ZeroCopySourceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZeroCopySourceTransactorRaw struct {
	Contract *ZeroCopySourceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZeroCopySource creates a new instance of ZeroCopySource, bound to a specific deployed contract.
func NewZeroCopySource(address common.Address, backend bind.ContractBackend) (*ZeroCopySource, error) {
	contract, err := bindZeroCopySource(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySource{ZeroCopySourceCaller: ZeroCopySourceCaller{contract: contract}, ZeroCopySourceTransactor: ZeroCopySourceTransactor{contract: contract}, ZeroCopySourceFilterer: ZeroCopySourceFilterer{contract: contract}}, nil
}

// NewZeroCopySourceCaller creates a new read-only instance of ZeroCopySource, bound to a specific deployed contract.
func NewZeroCopySourceCaller(address common.Address, caller bind.ContractCaller) (*ZeroCopySourceCaller, error) {
	contract, err := bindZeroCopySource(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySourceCaller{contract: contract}, nil
}

// NewZeroCopySourceTransactor creates a new write-only instance of ZeroCopySource, bound to a specific deployed contract.
func NewZeroCopySourceTransactor(address common.Address, transactor bind.ContractTransactor) (*ZeroCopySourceTransactor, error) {
	contract, err := bindZeroCopySource(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySourceTransactor{contract: contract}, nil
}

// NewZeroCopySourceFilterer creates a new log filterer instance of ZeroCopySource, bound to a specific deployed contract.
func NewZeroCopySourceFilterer(address common.Address, filterer bind.ContractFilterer) (*ZeroCopySourceFilterer, error) {
	contract, err := bindZeroCopySource(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySourceFilterer{contract: contract}, nil
}

// bindZeroCopySource binds a generic wrapper to an already deployed contract.
func bindZeroCopySource(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroCopySourceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroCopySource *ZeroCopySourceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZeroCopySource.Contract.ZeroCopySourceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroCopySource *ZeroCopySourceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroCopySource.Contract.ZeroCopySourceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroCopySource *ZeroCopySourceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroCopySource.Contract.ZeroCopySourceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroCopySource *ZeroCopySourceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZeroCopySource.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroCopySource *ZeroCopySourceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroCopySource.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroCopySource *ZeroCopySourceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroCopySource.Contract.contract.Transact(opts, method, params...)
}


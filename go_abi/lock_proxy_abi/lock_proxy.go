// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lock_proxy_abi

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

// ERC20InterfaceABI is the input ABI used to generate the binding from.
const ERC20InterfaceABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20InterfaceFuncSigs maps the 4-byte function signature to its string representation.
var ERC20InterfaceFuncSigs = map[string]string{
	"70a08231": "balanceOf(address)",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20Interface is an auto generated Go binding around an Ethereum contract.
type ERC20Interface struct {
	ERC20InterfaceCaller     // Read-only binding to the contract
	ERC20InterfaceTransactor // Write-only binding to the contract
	ERC20InterfaceFilterer   // Log filterer for contract events
}

// ERC20InterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20InterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20InterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20InterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20InterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20InterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20InterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20InterfaceSession struct {
	Contract     *ERC20Interface   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20InterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20InterfaceCallerSession struct {
	Contract *ERC20InterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC20InterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20InterfaceTransactorSession struct {
	Contract     *ERC20InterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC20InterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20InterfaceRaw struct {
	Contract *ERC20Interface // Generic contract binding to access the raw methods on
}

// ERC20InterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20InterfaceCallerRaw struct {
	Contract *ERC20InterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20InterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20InterfaceTransactorRaw struct {
	Contract *ERC20InterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Interface creates a new instance of ERC20Interface, bound to a specific deployed contract.
func NewERC20Interface(address common.Address, backend bind.ContractBackend) (*ERC20Interface, error) {
	contract, err := bindERC20Interface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Interface{ERC20InterfaceCaller: ERC20InterfaceCaller{contract: contract}, ERC20InterfaceTransactor: ERC20InterfaceTransactor{contract: contract}, ERC20InterfaceFilterer: ERC20InterfaceFilterer{contract: contract}}, nil
}

// NewERC20InterfaceCaller creates a new read-only instance of ERC20Interface, bound to a specific deployed contract.
func NewERC20InterfaceCaller(address common.Address, caller bind.ContractCaller) (*ERC20InterfaceCaller, error) {
	contract, err := bindERC20Interface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20InterfaceCaller{contract: contract}, nil
}

// NewERC20InterfaceTransactor creates a new write-only instance of ERC20Interface, bound to a specific deployed contract.
func NewERC20InterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20InterfaceTransactor, error) {
	contract, err := bindERC20Interface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20InterfaceTransactor{contract: contract}, nil
}

// NewERC20InterfaceFilterer creates a new log filterer instance of ERC20Interface, bound to a specific deployed contract.
func NewERC20InterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20InterfaceFilterer, error) {
	contract, err := bindERC20Interface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20InterfaceFilterer{contract: contract}, nil
}

// bindERC20Interface binds a generic wrapper to an already deployed contract.
func bindERC20Interface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20InterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Interface *ERC20InterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Interface.Contract.ERC20InterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Interface *ERC20InterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Interface.Contract.ERC20InterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Interface *ERC20InterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Interface.Contract.ERC20InterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Interface *ERC20InterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Interface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Interface *ERC20InterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Interface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Interface *ERC20InterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Interface.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Interface *ERC20InterfaceCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Interface.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Interface *ERC20InterfaceSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Interface.Contract.BalanceOf(&_ERC20Interface.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Interface *ERC20InterfaceCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Interface.Contract.BalanceOf(&_ERC20Interface.CallOpts, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20Interface *ERC20InterfaceTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20Interface *ERC20InterfaceSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.Transfer(&_ERC20Interface.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20Interface *ERC20InterfaceTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.Transfer(&_ERC20Interface.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_ERC20Interface *ERC20InterfaceTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_ERC20Interface *ERC20InterfaceSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.TransferFrom(&_ERC20Interface.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_ERC20Interface *ERC20InterfaceTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.TransferFrom(&_ERC20Interface.TransactOpts, _from, _to, _value)
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

// IEthCrossChainManagerProxyABI is the input ABI used to generate the binding from.
const IEthCrossChainManagerProxyABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getEthCrossChainManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IEthCrossChainManagerProxyFuncSigs maps the 4-byte function signature to its string representation.
var IEthCrossChainManagerProxyFuncSigs = map[string]string{
	"87939a7f": "getEthCrossChainManager()",
}

// IEthCrossChainManagerProxy is an auto generated Go binding around an Ethereum contract.
type IEthCrossChainManagerProxy struct {
	IEthCrossChainManagerProxyCaller     // Read-only binding to the contract
	IEthCrossChainManagerProxyTransactor // Write-only binding to the contract
	IEthCrossChainManagerProxyFilterer   // Log filterer for contract events
}

// IEthCrossChainManagerProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEthCrossChainManagerProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEthCrossChainManagerProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEthCrossChainManagerProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEthCrossChainManagerProxySession struct {
	Contract     *IEthCrossChainManagerProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IEthCrossChainManagerProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEthCrossChainManagerProxyCallerSession struct {
	Contract *IEthCrossChainManagerProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// IEthCrossChainManagerProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEthCrossChainManagerProxyTransactorSession struct {
	Contract     *IEthCrossChainManagerProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// IEthCrossChainManagerProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEthCrossChainManagerProxyRaw struct {
	Contract *IEthCrossChainManagerProxy // Generic contract binding to access the raw methods on
}

// IEthCrossChainManagerProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEthCrossChainManagerProxyCallerRaw struct {
	Contract *IEthCrossChainManagerProxyCaller // Generic read-only contract binding to access the raw methods on
}

// IEthCrossChainManagerProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEthCrossChainManagerProxyTransactorRaw struct {
	Contract *IEthCrossChainManagerProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEthCrossChainManagerProxy creates a new instance of IEthCrossChainManagerProxy, bound to a specific deployed contract.
func NewIEthCrossChainManagerProxy(address common.Address, backend bind.ContractBackend) (*IEthCrossChainManagerProxy, error) {
	contract, err := bindIEthCrossChainManagerProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerProxy{IEthCrossChainManagerProxyCaller: IEthCrossChainManagerProxyCaller{contract: contract}, IEthCrossChainManagerProxyTransactor: IEthCrossChainManagerProxyTransactor{contract: contract}, IEthCrossChainManagerProxyFilterer: IEthCrossChainManagerProxyFilterer{contract: contract}}, nil
}

// NewIEthCrossChainManagerProxyCaller creates a new read-only instance of IEthCrossChainManagerProxy, bound to a specific deployed contract.
func NewIEthCrossChainManagerProxyCaller(address common.Address, caller bind.ContractCaller) (*IEthCrossChainManagerProxyCaller, error) {
	contract, err := bindIEthCrossChainManagerProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerProxyCaller{contract: contract}, nil
}

// NewIEthCrossChainManagerProxyTransactor creates a new write-only instance of IEthCrossChainManagerProxy, bound to a specific deployed contract.
func NewIEthCrossChainManagerProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*IEthCrossChainManagerProxyTransactor, error) {
	contract, err := bindIEthCrossChainManagerProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerProxyTransactor{contract: contract}, nil
}

// NewIEthCrossChainManagerProxyFilterer creates a new log filterer instance of IEthCrossChainManagerProxy, bound to a specific deployed contract.
func NewIEthCrossChainManagerProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*IEthCrossChainManagerProxyFilterer, error) {
	contract, err := bindIEthCrossChainManagerProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerProxyFilterer{contract: contract}, nil
}

// bindIEthCrossChainManagerProxy binds a generic wrapper to an already deployed contract.
func bindIEthCrossChainManagerProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IEthCrossChainManagerProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainManagerProxy.Contract.IEthCrossChainManagerProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainManagerProxy.Contract.IEthCrossChainManagerProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainManagerProxy.Contract.IEthCrossChainManagerProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainManagerProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainManagerProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainManagerProxy.Contract.contract.Transact(opts, method, params...)
}

// GetEthCrossChainManager is a free data retrieval call binding the contract method 0x87939a7f.
//
// Solidity: function getEthCrossChainManager() view returns(address)
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyCaller) GetEthCrossChainManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IEthCrossChainManagerProxy.contract.Call(opts, out, "getEthCrossChainManager")
	return *ret0, err
}

// GetEthCrossChainManager is a free data retrieval call binding the contract method 0x87939a7f.
//
// Solidity: function getEthCrossChainManager() view returns(address)
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxySession) GetEthCrossChainManager() (common.Address, error) {
	return _IEthCrossChainManagerProxy.Contract.GetEthCrossChainManager(&_IEthCrossChainManagerProxy.CallOpts)
}

// GetEthCrossChainManager is a free data retrieval call binding the contract method 0x87939a7f.
//
// Solidity: function getEthCrossChainManager() view returns(address)
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyCallerSession) GetEthCrossChainManager() (common.Address, error) {
	return _IEthCrossChainManagerProxy.Contract.GetEthCrossChainManager(&_IEthCrossChainManagerProxy.CallOpts)
}

// LockProxyABI is the input ABI used to generate the binding from.
const LockProxyABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"targetProxyHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialAmount\",\"type\":\"uint256\"}],\"name\":\"BindAssetEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"targetProxyHash\",\"type\":\"bytes\"}],\"name\":\"BindProxyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toAssetHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LockEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"SetManagerProxyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toAssetHash\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnlockEvent\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"assetHashMap\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"toAssetHash\",\"type\":\"bytes\"}],\"name\":\"bindAssetHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"targetProxyHash\",\"type\":\"bytes\"}],\"name\":\"bindProxyHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"}],\"name\":\"getBalanceFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"managerProxyContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"proxyHashMap\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethCCMProxyAddr\",\"type\":\"address\"}],\"name\":\"setManagerProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"argsBs\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"fromContractAddr\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"fromChainId\",\"type\":\"uint64\"}],\"name\":\"unlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// LockProxyFuncSigs maps the 4-byte function signature to its string representation.
var LockProxyFuncSigs = map[string]string{
	"4f7d9808": "assetHashMap(address,uint64)",
	"3348f63b": "bindAssetHash(address,uint64,bytes)",
	"379b98f6": "bindProxyHash(uint64,bytes)",
	"59c589a1": "getBalanceFor(address)",
	"84a6d055": "lock(address,uint64,bytes,uint256)",
	"d798f881": "managerProxyContract()",
	"570ca735": "operator()",
	"9e5767aa": "proxyHashMap(uint64)",
	"af9980f0": "setManagerProxy(address)",
	"06af4b9f": "unlock(bytes,bytes,uint64)",
}

// LockProxyBin is the compiled bytecode used for deploying new contracts.
var LockProxyBin = "0x608060405234801561001057600080fd5b506100226001600160e01b0361004716565b600080546001600160a01b0319166001600160a01b039290921691909117905561004b565b3390565b6127678061005a6000396000f3fe6080604052600436106100915760003560e01c806359c589a11161005957806359c589a11461046357806384a6d055146104a85780639e5767aa1461056c578063af9980f01461059f578063d798f881146105d457610091565b806306af4b9f146100965780633348f63b146101eb578063379b98f6146102ba5780634f7d98081461037b578063570ca73514610432575b600080fd5b3480156100a257600080fd5b506101d7600480360360608110156100b957600080fd5b810190602081018135600160201b8111156100d357600080fd5b8201836020820111156100e557600080fd5b803590602001918460018302840111600160201b8311171561010657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561015857600080fd5b82018360208201111561016a57600080fd5b803590602001918460018302840111600160201b8311171561018b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160401b031691506105e99050565b604080519115158252519081900360200190f35b3480156101f757600080fd5b506101d76004803603606081101561020e57600080fd5b6001600160a01b03823516916001600160401b0360208201351691810190606081016040820135600160201b81111561024657600080fd5b82018360208201111561025857600080fd5b803590602001918460018302840111600160201b8311171561027957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506108dc945050505050565b3480156102c657600080fd5b506101d7600480360360408110156102dd57600080fd5b6001600160401b038235169190810190604081016020820135600160201b81111561030757600080fd5b82018360208201111561031957600080fd5b803590602001918460018302840111600160201b8311171561033a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a26945050505050565b34801561038757600080fd5b506103bd6004803603604081101561039e57600080fd5b5080356001600160a01b031690602001356001600160401b0316610b34565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103f75781810151838201526020016103df565b50505050905090810190601f1680156104245780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561043e57600080fd5b50610447610bd8565b604080516001600160a01b039092168252519081900360200190f35b34801561046f57600080fd5b506104966004803603602081101561048657600080fd5b50356001600160a01b0316610be7565b60408051918252519081900360200190f35b6101d7600480360360808110156104be57600080fd5b6001600160a01b03823516916001600160401b0360208201351691810190606081016040820135600160201b8111156104f657600080fd5b82018360208201111561050857600080fd5b803590602001918460018302840111600160201b8311171561052957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610c82915050565b34801561057857600080fd5b506103bd6004803603602081101561058f57600080fd5b50356001600160401b0316611300565b3480156105ab57600080fd5b506105d2600480360360208110156105c257600080fd5b50356001600160a01b0316611366565b005b3480156105e057600080fd5b506104476113e7565b600154604080516387939a7f60e01b815290516000926001600160a01b03169182916387939a7f91600480820192602092909190829003018186803b15801561063157600080fd5b505afa158015610645573d6000803e3d6000fd5b505050506040513d602081101561065b57600080fd5b50516001600160a01b031661066e6113f6565b6001600160a01b0316146106b35760405162461bcd60e51b815260040180806020018281038252602d815260200180612555602d913960400191505060405180910390fd5b6106bb6123e6565b6106c4866113fb565b90508451600014156107075760405162461bcd60e51b815260040180806020018281038252602b815260200180612582602b913960400191505060405180910390fd5b6001600160401b03841660009081526002602052604090206107299086611447565b6107645760405162461bcd60e51b81526004018080602001828103825260228152602001806126f06022913960400191505060405180910390fd5b8051516107b8576040805162461bcd60e51b815260206004820152601b60248201527f746f4173736574486173682063616e6e6f7420626520656d7074790000000000604482015290519081900360640190fd5b60006107c782600001516114fb565b905081602001515160001415610824576040805162461bcd60e51b815260206004820152601960248201527f746f416464726573732063616e6e6f7420626520656d70747900000000000000604482015290519081900360640190fd5b600061083383602001516114fb565b905061084482828560400151611545565b61087f5760405162461bcd60e51b815260040180806020018281038252603c815260200180612519603c913960400191505060405180910390fd5b60408084015181516001600160a01b038086168252841660208201528083019190915290517fd90288730b87c2b8e0c45bd82260fd22478aba30ae1c4d578b8daba9261604df9181900360600190a1506001979650505050505050565b600080546001600160a01b03166108f16113f6565b6001600160a01b03161461090457600080fd5b6001600160a01b03841660009081526003602090815260408083206001600160401b03871684528252909120835161093e92850190612407565b507f1628c8374c1bdfeb2275fd9f4c90562fd3fae974783dc522c8234e36abcfc58e84848461096c88610be7565b60405180856001600160a01b03166001600160a01b03168152602001846001600160401b03166001600160401b0316815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b838110156109df5781810151838201526020016109c7565b50505050905090810190601f168015610a0c5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a15060019392505050565b600080546001600160a01b0316610a3b6113f6565b6001600160a01b031614610a4e57600080fd5b6001600160401b03831660009081526002602090815260409091208351610a7792850190612407565b507fdacd7d303272a3b58aec6620d6d1fb588f4996a5b46858ed437f1c34348f2d0f838360405180836001600160401b03166001600160401b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610af0578181015183820152602001610ad8565b50505050905090810190601f168015610b1d5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a150600192915050565b60036020908152600092835260408084208252918352918190208054825160026001831615610100026000190190921691909104601f810185900485028201850190935282815292909190830182828015610bd05780601f10610ba557610100808354040283529160200191610bd0565b820191906000526020600020905b815481529060010190602001808311610bb357829003601f168201915b505050505081565b6000546001600160a01b031681565b60006001600160a01b038216610bff57503031610c7d565b604080516370a0823160e01b8152306004820152905183916001600160a01b038316916370a0823191602480820192602092909190829003018186803b158015610c4857600080fd5b505afa158015610c5c573d6000803e3d6000fd5b505050506040513d6020811015610c7257600080fd5b50519150610c7d9050565b919050565b600081610cd6576040805162461bcd60e51b815260206004820152601960248201527f616d6f756e74206973206c657373207468616e207a65726f2100000000000000604482015290519081900360640190fd5b610ce085836115e1565b610d1b5760405162461bcd60e51b815260040180806020018281038252603f8152602001806125fe603f913960400191505060405180910390fd5b6001600160a01b03851660009081526003602090815260408083206001600160401b038816845282529182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845260609392830182828015610dcb5780601f10610da057610100808354040283529160200191610dcb565b820191906000526020600020905b815481529060010190602001808311610dae57829003601f168201915b50505050509050805160001415610e29576040805162461bcd60e51b815260206004820152601960248201527f656d70747920696c6c6567616c20746f41737365744861736800000000000000604482015290519081900360640190fd5b610e316123e6565b60405180606001604052808381526020018681526020018581525090506060610e598261170f565b90506000600160009054906101000a90046001600160a01b031690506000816001600160a01b03166387939a7f6040518163ffffffff1660e01b815260040160206040518083038186803b158015610eb057600080fd5b505afa158015610ec4573d6000803e3d6000fd5b505050506040513d6020811015610eda57600080fd5b50516001600160401b038a16600090815260026020818152604092839020805484516001821615610100026000190190911693909304601f810183900483028401830190945283835293945084936060939091830182828015610f7e5780601f10610f5357610100808354040283529160200191610f7e565b820191906000526020600020905b815481529060010190602001808311610f6157829003601f168201915b50505050509050805160001415610fdc576040805162461bcd60e51b815260206004820152601960248201527f656d70747920696c6c6567616c20746f50726f78794861736800000000000000604482015290519081900360640190fd5b816001600160a01b031663bd5cf6258c83886040518463ffffffff1660e01b815260040180846001600160401b03166001600160401b03168152602001806020018060200180602001848103845286818151815260200191508051906020019080838360005b8381101561105a578181015183820152602001611042565b50505050905090810190601f1680156110875780820380516001836020036101000a031916815260200191505b508481038352600681526020018065756e6c6f636b60d01b815250602001848103825285818151815260200191508051906020019080838360005b838110156110da5781810151838201526020016110c2565b50505050905090810190601f1680156111075780820380516001836020036101000a031916815260200191505b509650505050505050602060405180830381600087803b15801561112a57600080fd5b505af115801561113e573d6000803e3d6000fd5b505050506040513d602081101561115457600080fd5b50516111915760405162461bcd60e51b815260040180806020018281038252602f8152602001806125cf602f913960400191505060405180910390fd5b7f8636abd6d0e464fe725a13346c7ac779b73561c705506044a2e6b2cdb1295ea58c6111bb6113f6565b8d8a8e8e60405180876001600160a01b03166001600160a01b03168152602001866001600160a01b03166001600160a01b03168152602001856001600160401b03166001600160401b031681526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360005b8381101561124e578181015183820152602001611236565b50505050905090810190601f16801561127b5780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b838110156112ae578181015183820152602001611296565b50505050905090810190601f1680156112db5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390a15060019b9a5050505050505050505050565b600260208181526000928352604092839020805484516001821615610100026000190190911693909304601f8101839004830284018301909452838352919290830182828015610bd05780601f10610ba557610100808354040283529160200191610bd0565b6000546001600160a01b031661137a6113f6565b6001600160a01b03161461138d57600080fd5b600180546001600160a01b0319166001600160a01b03838116919091179182905560408051929091168252517f43b1a8ec337adb61e8311ed025d99c80db65c02fe5c5027c1b6a93b40970cec4916020908290030190a150565b6001546001600160a01b031681565b335b90565b6114036123e6565b61140b6123e6565b6000611417848261183c565b9083529050611426848261183c565b6020840191909152905061143a8482611914565b5060408301525092915050565b6000806001905083546002600180831615610100020382160484518082146001811461147657600094506114ef565b82156114ef5760208310600181146114d457600189600052602060002060208a018581015b6002848284100114156114cb5781518354146114ba5760009950600093505b60018301925060208201915061149b565b505050506114ed565b610100808604029450602088015185146114ed57600095505b505b50929695505050505050565b6000815160141461153d5760405162461bcd60e51b81526004018080602001828103825260238152602001806124c36023913960400191505060405180910390fd5b506014015190565b60006001600160a01b038416611591576040516001600160a01b0384169083156108fc029084906000818181858888f1935050505015801561158b573d6000803e3d6000fd5b506115d7565b61159c848484611a11565b6115d75760405162461bcd60e51b81526004018080602001828103825260338152602001806124e66033913960400191505060405180910390fd5b5060019392505050565b60006001600160a01b03831661167b573461162d5760405162461bcd60e51b81526004018080602001828103825260218152602001806127126021913960400191505060405180910390fd5b813414801561163b57503415155b6116765760405162461bcd60e51b815260040180806020018281038252602981526020018061263d6029913960400191505060405180910390fd5b611706565b34156116b85760405162461bcd60e51b81526004018080602001828103825260228152602001806125ad6022913960400191505060405180910390fd5b6116cb836116c46113f6565b3085611af0565b6117065760405162461bcd60e51b81526004018080602001828103825260338152602001806124e66033913960400191505060405180910390fd5b50600192915050565b60608061171f8360000151611bd8565b61172c8460200151611bd8565b6117398560400151611c9e565b6040516020018084805190602001908083835b6020831061176b5780518252601f19909201916020918201910161174c565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106117b35780518252601f199092019160209182019101611794565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106117fb5780518252601f1990920191602091820191016117dc565b6001836020036101000a0380198251168184511680821785525050505050509050019350505050604051602081830303815290604052905080915050919050565b606060008061184b8585611d3b565b865190955090915081850111801590611865575080840184105b6118a05760405162461bcd60e51b81526004018080602001828103825260248152602001806126cc6024913960400191505060405180910390fd5b6060811580156118bb57604051915060208201604052611905565b6040519150601f8316801560200281840101848101888315602002848c0101015b818310156118f45780518352602092830192016118dc565b5050848452601f01601f1916604052505b509250830190505b9250929050565b6000808351836020011115801561192d57508260200183105b6119685760405162461bcd60e51b81526004018080602001828103825260238152602001806124a06023913960400191505060405180910390fd5b600060405160206000600182038760208a0101515b8383101561199d5780821a8386015360018301925060018203915061197d565b50505081016040525190506001600160ff1b03811115611a04576040805162461bcd60e51b815260206004820152601760248201527f56616c75652065786365656473207468652072616e6765000000000000000000604482015290519081900360640190fd5b9460209390930193505050565b6040805163a9059cbb60e01b81526001600160a01b03848116600483015260248201849052915160009286929083169163a9059cbb9160448082019260209290919082900301818887803b158015611a6857600080fd5b505af1158015611a7c573d6000803e3d6000fd5b505050506040513d6020811015611a9257600080fd5b5051611ae5576040805162461bcd60e51b815260206004820152601c60248201527f747261736e66657220455243323020546f6b656e206661696c65642100000000604482015290519081900360640190fd5b506001949350505050565b604080516323b872dd60e01b81526001600160a01b03858116600483015284811660248301526044820184905291516000928792908316916323b872dd9160648082019260209290919082900301818887803b158015611b4f57600080fd5b505af1158015611b63573d6000803e3d6000fd5b505050506040513d6020811015611b7957600080fd5b5051611bcc576040805162461bcd60e51b815260206004820152601c60248201527f747261736e66657220455243323020546f6b656e206661696c65642100000000604482015290519081900360640190fd5b50600195945050505050565b8051606090611be681611f54565b836040516020018083805190602001908083835b60208310611c195780518252601f199092019160209182019101611bfa565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b60208310611c615780518252601f199092019160209182019101611c42565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052915050919050565b60606001600160ff1b03821115611cfc576040805162461bcd60e51b815260206004820152601b60248201527f56616c756520657863656564732075696e743235352072616e67650000000000604482015290519081900360640190fd5b60405160208082526000601f5b82821015611d2b5785811a826020860101536001919091019060001901611d09565b5050506040818101905292915050565b6000806000611d4a858561209a565b94509050600060fd60f81b6001600160f81b031983161415611de857611d708686612118565b955061ffff16905060fd8110801590611d8b575061ffff8111155b611ddc576040805162461bcd60e51b815260206004820152601f60248201527f4e65787455696e7431362c2076616c7565206f7574736964652072616e676500604482015290519081900360640190fd5b925083915061190d9050565b607f60f91b6001600160f81b031983161415611e7857611e0886866121a1565b955063ffffffff16905061ffff81118015611e27575063ffffffff8111155b611ddc576040805162461bcd60e51b815260206004820181905260248201527f4e65787456617255696e742c2076616c7565206f7574736964652072616e6765604482015290519081900360640190fd5b6001600160f81b03198083161415611ef957611e948686612247565b95506001600160401b0316905063ffffffff8111611ddc576040805162461bcd60e51b815260206004820181905260248201527f4e65787456617255696e742c2076616c7565206f7574736964652072616e6765604482015290519081900360640190fd5b5060f881901c60fd8110611ddc576040805162461bcd60e51b815260206004820181905260248201527f4e65787456617255696e742c2076616c7565206f7574736964652072616e6765604482015290519081900360640190fd5b606060fd826001600160401b03161015611f7857611f71826122ed565b9050610c7d565b61ffff826001600160401b03161161205657611f9760fd60f81b612309565b611fa08361231d565b6040516020018083805190602001908083835b60208310611fd25780518252601f199092019160209182019101611fb3565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b6020831061201a5780518252601f199092019160209182019101611ffb565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529050610c7d565b63ffffffff826001600160401b03161161208057612077607f60f91b612309565b611fa083612360565b6120916001600160f81b0319612309565b611fa0836123a3565b600080835183600101111580156120b357508260010183105b612104576040805162461bcd60e51b815260206004820181905260248201527f4e657874427974652c204f66667365742065786365656473206d6178696d756d604482015290519081900360640190fd5b505081810160200151600182019250929050565b6000808351836002011115801561213157508260020183105b61216c5760405162461bcd60e51b81526004018080602001828103825260228152602001806126666022913960400191505060405180910390fd5b6000604051846020870101518060011a82538060001a6001830153506002818101604052601d19909101519694019450505050565b600080835183600401111580156121ba57508260040183105b6121f55760405162461bcd60e51b81526004018080602001828103825260228152602001806126aa6022913960400191505060405180910390fd5b600060405160046000600182038760208a0101515b8383101561222a5780821a8386015360018301925060018203915061220a565b505050808201604052602003900351956004949094019450505050565b6000808351836008011115801561226057508260080183105b61229b5760405162461bcd60e51b81526004018080602001828103825260228152602001806126886022913960400191505060405180910390fd5b600060405160086000600182038760208a0101515b838310156122d05780821a838601536001830192506001820391506122b0565b505050808201604052602003900351956008949094019450505050565b604080516001815260f89290921b602083015260218201905290565b60606123178260f81c6122ed565b92915050565b6040516002808252606091906000601f5b828210156123505785811a82602086010153600191909101906000190161232e565b5050506022810160405292915050565b6040516004808252606091906000601f5b828210156123935785811a826020860101536001919091019060001901612371565b5050506024810160405292915050565b6040516008808252606091906000601f5b828210156123d65785811a8260208601015360019190910190600019016123b4565b5050506028810160405292915050565b60405180606001604052806060815260200160608152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061244857805160ff1916838001178555612475565b82800160010185558215612475579182015b8281111561247557825182559160200191906001019061245a565b50612481929150612485565b5090565b6113f891905b80821115612481576000815560010161248b56fe4e65787455696e743235352c206f66667365742065786365656473206d6178696d756d6279746573206c656e67746820646f6573206e6f74206d6174636820616464726573737472616e7366657220657263323020617373657420746f206c6f636b5f70726f787920636f6e7472616374206661696c6564217472616e736665722061737365742066726f6d206c6f636b5f70726f787920636f6e747261637420746f20746f41646472657373206661696c6564216d736753656e646572206973206e6f742045746843726f7373436861696e4d616e61676572436f6e747261637466726f6d2070726f787920636f6e747261637420616464726573732063616e6e6f7420626520656d70747974686572652073686f756c64206265206e6f206574686572207472616e736665722145746843726f7373436861696e4d616e616765722063726f7373436861696e206578656375746564206572726f72217472616e736665722061737365742066726f6d2066726f6d4164647265737320746f206c6f636b5f70726f787920636f6e747261637420206661696c6564217472616e73666572726564206574686572206973206e6f7420657175616c20746f20616d6f756e74214e65787455696e7431362c206f66667365742065786365656473206d6178696d756d4e65787455696e7436342c206f66667365742065786365656473206d6178696d756d4e65787455696e7433322c206f66667365742065786365656473206d6178696d756d4e65787456617242797465732c206f66667365742065786365656473206d6178696d756d46726f6d2050726f787920636f6e74726163742061646472657373206572726f72217472616e736665727265642065746865722063616e6e6f74206265207a65726f21a265627a7a723158202dd993160c5bb84a3514a7304addecd913bc3abeecbb749dcd43ab4d31d96a5964736f6c634300050f0032"

// DeployLockProxy deploys a new Ethereum contract, binding an instance of LockProxy to it.
func DeployLockProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LockProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(LockProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LockProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LockProxy{LockProxyCaller: LockProxyCaller{contract: contract}, LockProxyTransactor: LockProxyTransactor{contract: contract}, LockProxyFilterer: LockProxyFilterer{contract: contract}}, nil
}

// LockProxy is an auto generated Go binding around an Ethereum contract.
type LockProxy struct {
	LockProxyCaller     // Read-only binding to the contract
	LockProxyTransactor // Write-only binding to the contract
	LockProxyFilterer   // Log filterer for contract events
}

// LockProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type LockProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LockProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LockProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LockProxySession struct {
	Contract     *LockProxy        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LockProxyCallerSession struct {
	Contract *LockProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// LockProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LockProxyTransactorSession struct {
	Contract     *LockProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// LockProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type LockProxyRaw struct {
	Contract *LockProxy // Generic contract binding to access the raw methods on
}

// LockProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LockProxyCallerRaw struct {
	Contract *LockProxyCaller // Generic read-only contract binding to access the raw methods on
}

// LockProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LockProxyTransactorRaw struct {
	Contract *LockProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLockProxy creates a new instance of LockProxy, bound to a specific deployed contract.
func NewLockProxy(address common.Address, backend bind.ContractBackend) (*LockProxy, error) {
	contract, err := bindLockProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LockProxy{LockProxyCaller: LockProxyCaller{contract: contract}, LockProxyTransactor: LockProxyTransactor{contract: contract}, LockProxyFilterer: LockProxyFilterer{contract: contract}}, nil
}

// NewLockProxyCaller creates a new read-only instance of LockProxy, bound to a specific deployed contract.
func NewLockProxyCaller(address common.Address, caller bind.ContractCaller) (*LockProxyCaller, error) {
	contract, err := bindLockProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockProxyCaller{contract: contract}, nil
}

// NewLockProxyTransactor creates a new write-only instance of LockProxy, bound to a specific deployed contract.
func NewLockProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*LockProxyTransactor, error) {
	contract, err := bindLockProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockProxyTransactor{contract: contract}, nil
}

// NewLockProxyFilterer creates a new log filterer instance of LockProxy, bound to a specific deployed contract.
func NewLockProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*LockProxyFilterer, error) {
	contract, err := bindLockProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockProxyFilterer{contract: contract}, nil
}

// bindLockProxy binds a generic wrapper to an already deployed contract.
func bindLockProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LockProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LockProxy *LockProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LockProxy.Contract.LockProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LockProxy *LockProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockProxy.Contract.LockProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LockProxy *LockProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockProxy.Contract.LockProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LockProxy *LockProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LockProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LockProxy *LockProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LockProxy *LockProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockProxy.Contract.contract.Transact(opts, method, params...)
}

// AssetHashMap is a free data retrieval call binding the contract method 0x4f7d9808.
//
// Solidity: function assetHashMap(address , uint64 ) view returns(bytes)
func (_LockProxy *LockProxyCaller) AssetHashMap(opts *bind.CallOpts, arg0 common.Address, arg1 uint64) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _LockProxy.contract.Call(opts, out, "assetHashMap", arg0, arg1)
	return *ret0, err
}

// AssetHashMap is a free data retrieval call binding the contract method 0x4f7d9808.
//
// Solidity: function assetHashMap(address , uint64 ) view returns(bytes)
func (_LockProxy *LockProxySession) AssetHashMap(arg0 common.Address, arg1 uint64) ([]byte, error) {
	return _LockProxy.Contract.AssetHashMap(&_LockProxy.CallOpts, arg0, arg1)
}

// AssetHashMap is a free data retrieval call binding the contract method 0x4f7d9808.
//
// Solidity: function assetHashMap(address , uint64 ) view returns(bytes)
func (_LockProxy *LockProxyCallerSession) AssetHashMap(arg0 common.Address, arg1 uint64) ([]byte, error) {
	return _LockProxy.Contract.AssetHashMap(&_LockProxy.CallOpts, arg0, arg1)
}

// GetBalanceFor is a free data retrieval call binding the contract method 0x59c589a1.
//
// Solidity: function getBalanceFor(address fromAssetHash) view returns(uint256)
func (_LockProxy *LockProxyCaller) GetBalanceFor(opts *bind.CallOpts, fromAssetHash common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LockProxy.contract.Call(opts, out, "getBalanceFor", fromAssetHash)
	return *ret0, err
}

// GetBalanceFor is a free data retrieval call binding the contract method 0x59c589a1.
//
// Solidity: function getBalanceFor(address fromAssetHash) view returns(uint256)
func (_LockProxy *LockProxySession) GetBalanceFor(fromAssetHash common.Address) (*big.Int, error) {
	return _LockProxy.Contract.GetBalanceFor(&_LockProxy.CallOpts, fromAssetHash)
}

// GetBalanceFor is a free data retrieval call binding the contract method 0x59c589a1.
//
// Solidity: function getBalanceFor(address fromAssetHash) view returns(uint256)
func (_LockProxy *LockProxyCallerSession) GetBalanceFor(fromAssetHash common.Address) (*big.Int, error) {
	return _LockProxy.Contract.GetBalanceFor(&_LockProxy.CallOpts, fromAssetHash)
}

// ManagerProxyContract is a free data retrieval call binding the contract method 0xd798f881.
//
// Solidity: function managerProxyContract() view returns(address)
func (_LockProxy *LockProxyCaller) ManagerProxyContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LockProxy.contract.Call(opts, out, "managerProxyContract")
	return *ret0, err
}

// ManagerProxyContract is a free data retrieval call binding the contract method 0xd798f881.
//
// Solidity: function managerProxyContract() view returns(address)
func (_LockProxy *LockProxySession) ManagerProxyContract() (common.Address, error) {
	return _LockProxy.Contract.ManagerProxyContract(&_LockProxy.CallOpts)
}

// ManagerProxyContract is a free data retrieval call binding the contract method 0xd798f881.
//
// Solidity: function managerProxyContract() view returns(address)
func (_LockProxy *LockProxyCallerSession) ManagerProxyContract() (common.Address, error) {
	return _LockProxy.Contract.ManagerProxyContract(&_LockProxy.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_LockProxy *LockProxyCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LockProxy.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_LockProxy *LockProxySession) Operator() (common.Address, error) {
	return _LockProxy.Contract.Operator(&_LockProxy.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_LockProxy *LockProxyCallerSession) Operator() (common.Address, error) {
	return _LockProxy.Contract.Operator(&_LockProxy.CallOpts)
}

// ProxyHashMap is a free data retrieval call binding the contract method 0x9e5767aa.
//
// Solidity: function proxyHashMap(uint64 ) view returns(bytes)
func (_LockProxy *LockProxyCaller) ProxyHashMap(opts *bind.CallOpts, arg0 uint64) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _LockProxy.contract.Call(opts, out, "proxyHashMap", arg0)
	return *ret0, err
}

// ProxyHashMap is a free data retrieval call binding the contract method 0x9e5767aa.
//
// Solidity: function proxyHashMap(uint64 ) view returns(bytes)
func (_LockProxy *LockProxySession) ProxyHashMap(arg0 uint64) ([]byte, error) {
	return _LockProxy.Contract.ProxyHashMap(&_LockProxy.CallOpts, arg0)
}

// ProxyHashMap is a free data retrieval call binding the contract method 0x9e5767aa.
//
// Solidity: function proxyHashMap(uint64 ) view returns(bytes)
func (_LockProxy *LockProxyCallerSession) ProxyHashMap(arg0 uint64) ([]byte, error) {
	return _LockProxy.Contract.ProxyHashMap(&_LockProxy.CallOpts, arg0)
}

// BindAssetHash is a paid mutator transaction binding the contract method 0x3348f63b.
//
// Solidity: function bindAssetHash(address fromAssetHash, uint64 toChainId, bytes toAssetHash) returns(bool)
func (_LockProxy *LockProxyTransactor) BindAssetHash(opts *bind.TransactOpts, fromAssetHash common.Address, toChainId uint64, toAssetHash []byte) (*types.Transaction, error) {
	return _LockProxy.contract.Transact(opts, "bindAssetHash", fromAssetHash, toChainId, toAssetHash)
}

// BindAssetHash is a paid mutator transaction binding the contract method 0x3348f63b.
//
// Solidity: function bindAssetHash(address fromAssetHash, uint64 toChainId, bytes toAssetHash) returns(bool)
func (_LockProxy *LockProxySession) BindAssetHash(fromAssetHash common.Address, toChainId uint64, toAssetHash []byte) (*types.Transaction, error) {
	return _LockProxy.Contract.BindAssetHash(&_LockProxy.TransactOpts, fromAssetHash, toChainId, toAssetHash)
}

// BindAssetHash is a paid mutator transaction binding the contract method 0x3348f63b.
//
// Solidity: function bindAssetHash(address fromAssetHash, uint64 toChainId, bytes toAssetHash) returns(bool)
func (_LockProxy *LockProxyTransactorSession) BindAssetHash(fromAssetHash common.Address, toChainId uint64, toAssetHash []byte) (*types.Transaction, error) {
	return _LockProxy.Contract.BindAssetHash(&_LockProxy.TransactOpts, fromAssetHash, toChainId, toAssetHash)
}

// BindProxyHash is a paid mutator transaction binding the contract method 0x379b98f6.
//
// Solidity: function bindProxyHash(uint64 toChainId, bytes targetProxyHash) returns(bool)
func (_LockProxy *LockProxyTransactor) BindProxyHash(opts *bind.TransactOpts, toChainId uint64, targetProxyHash []byte) (*types.Transaction, error) {
	return _LockProxy.contract.Transact(opts, "bindProxyHash", toChainId, targetProxyHash)
}

// BindProxyHash is a paid mutator transaction binding the contract method 0x379b98f6.
//
// Solidity: function bindProxyHash(uint64 toChainId, bytes targetProxyHash) returns(bool)
func (_LockProxy *LockProxySession) BindProxyHash(toChainId uint64, targetProxyHash []byte) (*types.Transaction, error) {
	return _LockProxy.Contract.BindProxyHash(&_LockProxy.TransactOpts, toChainId, targetProxyHash)
}

// BindProxyHash is a paid mutator transaction binding the contract method 0x379b98f6.
//
// Solidity: function bindProxyHash(uint64 toChainId, bytes targetProxyHash) returns(bool)
func (_LockProxy *LockProxyTransactorSession) BindProxyHash(toChainId uint64, targetProxyHash []byte) (*types.Transaction, error) {
	return _LockProxy.Contract.BindProxyHash(&_LockProxy.TransactOpts, toChainId, targetProxyHash)
}

// Lock is a paid mutator transaction binding the contract method 0x84a6d055.
//
// Solidity: function lock(address fromAssetHash, uint64 toChainId, bytes toAddress, uint256 amount) payable returns(bool)
func (_LockProxy *LockProxyTransactor) Lock(opts *bind.TransactOpts, fromAssetHash common.Address, toChainId uint64, toAddress []byte, amount *big.Int) (*types.Transaction, error) {
	return _LockProxy.contract.Transact(opts, "lock", fromAssetHash, toChainId, toAddress, amount)
}

// Lock is a paid mutator transaction binding the contract method 0x84a6d055.
//
// Solidity: function lock(address fromAssetHash, uint64 toChainId, bytes toAddress, uint256 amount) payable returns(bool)
func (_LockProxy *LockProxySession) Lock(fromAssetHash common.Address, toChainId uint64, toAddress []byte, amount *big.Int) (*types.Transaction, error) {
	return _LockProxy.Contract.Lock(&_LockProxy.TransactOpts, fromAssetHash, toChainId, toAddress, amount)
}

// Lock is a paid mutator transaction binding the contract method 0x84a6d055.
//
// Solidity: function lock(address fromAssetHash, uint64 toChainId, bytes toAddress, uint256 amount) payable returns(bool)
func (_LockProxy *LockProxyTransactorSession) Lock(fromAssetHash common.Address, toChainId uint64, toAddress []byte, amount *big.Int) (*types.Transaction, error) {
	return _LockProxy.Contract.Lock(&_LockProxy.TransactOpts, fromAssetHash, toChainId, toAddress, amount)
}

// SetManagerProxy is a paid mutator transaction binding the contract method 0xaf9980f0.
//
// Solidity: function setManagerProxy(address ethCCMProxyAddr) returns()
func (_LockProxy *LockProxyTransactor) SetManagerProxy(opts *bind.TransactOpts, ethCCMProxyAddr common.Address) (*types.Transaction, error) {
	return _LockProxy.contract.Transact(opts, "setManagerProxy", ethCCMProxyAddr)
}

// SetManagerProxy is a paid mutator transaction binding the contract method 0xaf9980f0.
//
// Solidity: function setManagerProxy(address ethCCMProxyAddr) returns()
func (_LockProxy *LockProxySession) SetManagerProxy(ethCCMProxyAddr common.Address) (*types.Transaction, error) {
	return _LockProxy.Contract.SetManagerProxy(&_LockProxy.TransactOpts, ethCCMProxyAddr)
}

// SetManagerProxy is a paid mutator transaction binding the contract method 0xaf9980f0.
//
// Solidity: function setManagerProxy(address ethCCMProxyAddr) returns()
func (_LockProxy *LockProxyTransactorSession) SetManagerProxy(ethCCMProxyAddr common.Address) (*types.Transaction, error) {
	return _LockProxy.Contract.SetManagerProxy(&_LockProxy.TransactOpts, ethCCMProxyAddr)
}

// Unlock is a paid mutator transaction binding the contract method 0x06af4b9f.
//
// Solidity: function unlock(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_LockProxy *LockProxyTransactor) Unlock(opts *bind.TransactOpts, argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _LockProxy.contract.Transact(opts, "unlock", argsBs, fromContractAddr, fromChainId)
}

// Unlock is a paid mutator transaction binding the contract method 0x06af4b9f.
//
// Solidity: function unlock(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_LockProxy *LockProxySession) Unlock(argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _LockProxy.Contract.Unlock(&_LockProxy.TransactOpts, argsBs, fromContractAddr, fromChainId)
}

// Unlock is a paid mutator transaction binding the contract method 0x06af4b9f.
//
// Solidity: function unlock(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_LockProxy *LockProxyTransactorSession) Unlock(argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _LockProxy.Contract.Unlock(&_LockProxy.TransactOpts, argsBs, fromContractAddr, fromChainId)
}

// LockProxyBindAssetEventIterator is returned from FilterBindAssetEvent and is used to iterate over the raw logs and unpacked data for BindAssetEvent events raised by the LockProxy contract.
type LockProxyBindAssetEventIterator struct {
	Event *LockProxyBindAssetEvent // Event containing the contract specifics and raw log

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
func (it *LockProxyBindAssetEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockProxyBindAssetEvent)
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
		it.Event = new(LockProxyBindAssetEvent)
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
func (it *LockProxyBindAssetEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockProxyBindAssetEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockProxyBindAssetEvent represents a BindAssetEvent event raised by the LockProxy contract.
type LockProxyBindAssetEvent struct {
	FromAssetHash   common.Address
	ToChainId       uint64
	TargetProxyHash []byte
	InitialAmount   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBindAssetEvent is a free log retrieval operation binding the contract event 0x1628c8374c1bdfeb2275fd9f4c90562fd3fae974783dc522c8234e36abcfc58e.
//
// Solidity: event BindAssetEvent(address fromAssetHash, uint64 toChainId, bytes targetProxyHash, uint256 initialAmount)
func (_LockProxy *LockProxyFilterer) FilterBindAssetEvent(opts *bind.FilterOpts) (*LockProxyBindAssetEventIterator, error) {

	logs, sub, err := _LockProxy.contract.FilterLogs(opts, "BindAssetEvent")
	if err != nil {
		return nil, err
	}
	return &LockProxyBindAssetEventIterator{contract: _LockProxy.contract, event: "BindAssetEvent", logs: logs, sub: sub}, nil
}

// WatchBindAssetEvent is a free log subscription operation binding the contract event 0x1628c8374c1bdfeb2275fd9f4c90562fd3fae974783dc522c8234e36abcfc58e.
//
// Solidity: event BindAssetEvent(address fromAssetHash, uint64 toChainId, bytes targetProxyHash, uint256 initialAmount)
func (_LockProxy *LockProxyFilterer) WatchBindAssetEvent(opts *bind.WatchOpts, sink chan<- *LockProxyBindAssetEvent) (event.Subscription, error) {

	logs, sub, err := _LockProxy.contract.WatchLogs(opts, "BindAssetEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockProxyBindAssetEvent)
				if err := _LockProxy.contract.UnpackLog(event, "BindAssetEvent", log); err != nil {
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

// ParseBindAssetEvent is a log parse operation binding the contract event 0x1628c8374c1bdfeb2275fd9f4c90562fd3fae974783dc522c8234e36abcfc58e.
//
// Solidity: event BindAssetEvent(address fromAssetHash, uint64 toChainId, bytes targetProxyHash, uint256 initialAmount)
func (_LockProxy *LockProxyFilterer) ParseBindAssetEvent(log types.Log) (*LockProxyBindAssetEvent, error) {
	event := new(LockProxyBindAssetEvent)
	if err := _LockProxy.contract.UnpackLog(event, "BindAssetEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LockProxyBindProxyEventIterator is returned from FilterBindProxyEvent and is used to iterate over the raw logs and unpacked data for BindProxyEvent events raised by the LockProxy contract.
type LockProxyBindProxyEventIterator struct {
	Event *LockProxyBindProxyEvent // Event containing the contract specifics and raw log

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
func (it *LockProxyBindProxyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockProxyBindProxyEvent)
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
		it.Event = new(LockProxyBindProxyEvent)
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
func (it *LockProxyBindProxyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockProxyBindProxyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockProxyBindProxyEvent represents a BindProxyEvent event raised by the LockProxy contract.
type LockProxyBindProxyEvent struct {
	ToChainId       uint64
	TargetProxyHash []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBindProxyEvent is a free log retrieval operation binding the contract event 0xdacd7d303272a3b58aec6620d6d1fb588f4996a5b46858ed437f1c34348f2d0f.
//
// Solidity: event BindProxyEvent(uint64 toChainId, bytes targetProxyHash)
func (_LockProxy *LockProxyFilterer) FilterBindProxyEvent(opts *bind.FilterOpts) (*LockProxyBindProxyEventIterator, error) {

	logs, sub, err := _LockProxy.contract.FilterLogs(opts, "BindProxyEvent")
	if err != nil {
		return nil, err
	}
	return &LockProxyBindProxyEventIterator{contract: _LockProxy.contract, event: "BindProxyEvent", logs: logs, sub: sub}, nil
}

// WatchBindProxyEvent is a free log subscription operation binding the contract event 0xdacd7d303272a3b58aec6620d6d1fb588f4996a5b46858ed437f1c34348f2d0f.
//
// Solidity: event BindProxyEvent(uint64 toChainId, bytes targetProxyHash)
func (_LockProxy *LockProxyFilterer) WatchBindProxyEvent(opts *bind.WatchOpts, sink chan<- *LockProxyBindProxyEvent) (event.Subscription, error) {

	logs, sub, err := _LockProxy.contract.WatchLogs(opts, "BindProxyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockProxyBindProxyEvent)
				if err := _LockProxy.contract.UnpackLog(event, "BindProxyEvent", log); err != nil {
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

// ParseBindProxyEvent is a log parse operation binding the contract event 0xdacd7d303272a3b58aec6620d6d1fb588f4996a5b46858ed437f1c34348f2d0f.
//
// Solidity: event BindProxyEvent(uint64 toChainId, bytes targetProxyHash)
func (_LockProxy *LockProxyFilterer) ParseBindProxyEvent(log types.Log) (*LockProxyBindProxyEvent, error) {
	event := new(LockProxyBindProxyEvent)
	if err := _LockProxy.contract.UnpackLog(event, "BindProxyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LockProxyLockEventIterator is returned from FilterLockEvent and is used to iterate over the raw logs and unpacked data for LockEvent events raised by the LockProxy contract.
type LockProxyLockEventIterator struct {
	Event *LockProxyLockEvent // Event containing the contract specifics and raw log

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
func (it *LockProxyLockEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockProxyLockEvent)
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
		it.Event = new(LockProxyLockEvent)
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
func (it *LockProxyLockEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockProxyLockEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockProxyLockEvent represents a LockEvent event raised by the LockProxy contract.
type LockProxyLockEvent struct {
	FromAssetHash common.Address
	FromAddress   common.Address
	ToChainId     uint64
	ToAssetHash   []byte
	ToAddress     []byte
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLockEvent is a free log retrieval operation binding the contract event 0x8636abd6d0e464fe725a13346c7ac779b73561c705506044a2e6b2cdb1295ea5.
//
// Solidity: event LockEvent(address fromAssetHash, address fromAddress, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount)
func (_LockProxy *LockProxyFilterer) FilterLockEvent(opts *bind.FilterOpts) (*LockProxyLockEventIterator, error) {

	logs, sub, err := _LockProxy.contract.FilterLogs(opts, "LockEvent")
	if err != nil {
		return nil, err
	}
	return &LockProxyLockEventIterator{contract: _LockProxy.contract, event: "LockEvent", logs: logs, sub: sub}, nil
}

// WatchLockEvent is a free log subscription operation binding the contract event 0x8636abd6d0e464fe725a13346c7ac779b73561c705506044a2e6b2cdb1295ea5.
//
// Solidity: event LockEvent(address fromAssetHash, address fromAddress, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount)
func (_LockProxy *LockProxyFilterer) WatchLockEvent(opts *bind.WatchOpts, sink chan<- *LockProxyLockEvent) (event.Subscription, error) {

	logs, sub, err := _LockProxy.contract.WatchLogs(opts, "LockEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockProxyLockEvent)
				if err := _LockProxy.contract.UnpackLog(event, "LockEvent", log); err != nil {
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

// ParseLockEvent is a log parse operation binding the contract event 0x8636abd6d0e464fe725a13346c7ac779b73561c705506044a2e6b2cdb1295ea5.
//
// Solidity: event LockEvent(address fromAssetHash, address fromAddress, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount)
func (_LockProxy *LockProxyFilterer) ParseLockEvent(log types.Log) (*LockProxyLockEvent, error) {
	event := new(LockProxyLockEvent)
	if err := _LockProxy.contract.UnpackLog(event, "LockEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LockProxySetManagerProxyEventIterator is returned from FilterSetManagerProxyEvent and is used to iterate over the raw logs and unpacked data for SetManagerProxyEvent events raised by the LockProxy contract.
type LockProxySetManagerProxyEventIterator struct {
	Event *LockProxySetManagerProxyEvent // Event containing the contract specifics and raw log

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
func (it *LockProxySetManagerProxyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockProxySetManagerProxyEvent)
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
		it.Event = new(LockProxySetManagerProxyEvent)
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
func (it *LockProxySetManagerProxyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockProxySetManagerProxyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockProxySetManagerProxyEvent represents a SetManagerProxyEvent event raised by the LockProxy contract.
type LockProxySetManagerProxyEvent struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetManagerProxyEvent is a free log retrieval operation binding the contract event 0x43b1a8ec337adb61e8311ed025d99c80db65c02fe5c5027c1b6a93b40970cec4.
//
// Solidity: event SetManagerProxyEvent(address manager)
func (_LockProxy *LockProxyFilterer) FilterSetManagerProxyEvent(opts *bind.FilterOpts) (*LockProxySetManagerProxyEventIterator, error) {

	logs, sub, err := _LockProxy.contract.FilterLogs(opts, "SetManagerProxyEvent")
	if err != nil {
		return nil, err
	}
	return &LockProxySetManagerProxyEventIterator{contract: _LockProxy.contract, event: "SetManagerProxyEvent", logs: logs, sub: sub}, nil
}

// WatchSetManagerProxyEvent is a free log subscription operation binding the contract event 0x43b1a8ec337adb61e8311ed025d99c80db65c02fe5c5027c1b6a93b40970cec4.
//
// Solidity: event SetManagerProxyEvent(address manager)
func (_LockProxy *LockProxyFilterer) WatchSetManagerProxyEvent(opts *bind.WatchOpts, sink chan<- *LockProxySetManagerProxyEvent) (event.Subscription, error) {

	logs, sub, err := _LockProxy.contract.WatchLogs(opts, "SetManagerProxyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockProxySetManagerProxyEvent)
				if err := _LockProxy.contract.UnpackLog(event, "SetManagerProxyEvent", log); err != nil {
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

// ParseSetManagerProxyEvent is a log parse operation binding the contract event 0x43b1a8ec337adb61e8311ed025d99c80db65c02fe5c5027c1b6a93b40970cec4.
//
// Solidity: event SetManagerProxyEvent(address manager)
func (_LockProxy *LockProxyFilterer) ParseSetManagerProxyEvent(log types.Log) (*LockProxySetManagerProxyEvent, error) {
	event := new(LockProxySetManagerProxyEvent)
	if err := _LockProxy.contract.UnpackLog(event, "SetManagerProxyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LockProxyUnlockEventIterator is returned from FilterUnlockEvent and is used to iterate over the raw logs and unpacked data for UnlockEvent events raised by the LockProxy contract.
type LockProxyUnlockEventIterator struct {
	Event *LockProxyUnlockEvent // Event containing the contract specifics and raw log

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
func (it *LockProxyUnlockEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockProxyUnlockEvent)
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
		it.Event = new(LockProxyUnlockEvent)
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
func (it *LockProxyUnlockEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockProxyUnlockEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockProxyUnlockEvent represents a UnlockEvent event raised by the LockProxy contract.
type LockProxyUnlockEvent struct {
	ToAssetHash common.Address
	ToAddress   common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnlockEvent is a free log retrieval operation binding the contract event 0xd90288730b87c2b8e0c45bd82260fd22478aba30ae1c4d578b8daba9261604df.
//
// Solidity: event UnlockEvent(address toAssetHash, address toAddress, uint256 amount)
func (_LockProxy *LockProxyFilterer) FilterUnlockEvent(opts *bind.FilterOpts) (*LockProxyUnlockEventIterator, error) {

	logs, sub, err := _LockProxy.contract.FilterLogs(opts, "UnlockEvent")
	if err != nil {
		return nil, err
	}
	return &LockProxyUnlockEventIterator{contract: _LockProxy.contract, event: "UnlockEvent", logs: logs, sub: sub}, nil
}

// WatchUnlockEvent is a free log subscription operation binding the contract event 0xd90288730b87c2b8e0c45bd82260fd22478aba30ae1c4d578b8daba9261604df.
//
// Solidity: event UnlockEvent(address toAssetHash, address toAddress, uint256 amount)
func (_LockProxy *LockProxyFilterer) WatchUnlockEvent(opts *bind.WatchOpts, sink chan<- *LockProxyUnlockEvent) (event.Subscription, error) {

	logs, sub, err := _LockProxy.contract.WatchLogs(opts, "UnlockEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockProxyUnlockEvent)
				if err := _LockProxy.contract.UnpackLog(event, "UnlockEvent", log); err != nil {
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

// ParseUnlockEvent is a log parse operation binding the contract event 0xd90288730b87c2b8e0c45bd82260fd22478aba30ae1c4d578b8daba9261604df.
//
// Solidity: event UnlockEvent(address toAssetHash, address toAddress, uint256 amount)
func (_LockProxy *LockProxyFilterer) ParseUnlockEvent(log types.Log) (*LockProxyUnlockEvent, error) {
	event := new(LockProxyUnlockEvent)
	if err := _LockProxy.contract.UnpackLog(event, "UnlockEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158203d1637be9b56ceb98635efac8ed78a3613aacd2af1a906d51f09e708c8864eee64736f6c634300050f0032"

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

// UtilsABI is the input ABI used to generate the binding from.
const UtilsABI = "[]"

// UtilsBin is the compiled bytecode used for deploying new contracts.
var UtilsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820050dca012050b404de47b755368d860587e735c8b8ad9500f58b365fbb406df664736f6c634300050f0032"

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
var ZeroCopySinkBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158200d7209163aef4edca6e943dc5ba5502dc1fbcfa0829e293d0819c8c05dd6f69e64736f6c634300050f0032"

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
var ZeroCopySourceBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820a657b3a40ce034eca00b3cb7dad6596556e4a01575d19c83524645676fc4efaf64736f6c634300050f0032"

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


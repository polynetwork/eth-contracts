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
const EthCrossChainManagerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_eccd\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_chainId\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"contracts\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"curEpochPkBytes\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rawHeader\",\"type\":\"bytes\"}],\"name\":\"ChangeBookKeeperEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"txId\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxyOrAssetContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toContract\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rawdata\",\"type\":\"bytes\"}],\"name\":\"CrossChainEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rawHeader\",\"type\":\"bytes\"}],\"name\":\"InitGenesisBlockEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fromChainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toContract\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"crossChainTxHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fromChainTxHash\",\"type\":\"bytes\"}],\"name\":\"VerifyHeaderAndExecuteTxEvent\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"EthCrossChainDataAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rawHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyList\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sigList\",\"type\":\"bytes\"}],\"name\":\"changeBookKeeper\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"toContract\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"method\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"txData\",\"type\":\"bytes\"}],\"name\":\"crossChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rawHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyList\",\"type\":\"bytes\"}],\"name\":\"initGenesisBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"EpochPkBytes\",\"type\":\"bytes\"}],\"name\":\"recoverEpochPk\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_newChainId\",\"type\":\"uint64\"}],\"name\":\"setChainId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"unsetEpochPkBytes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newEthCrossChainManagerAddress\",\"type\":\"address\"}],\"name\":\"upgradeToNew\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rawHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"headerProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"curRawHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"headerSig\",\"type\":\"bytes\"}],\"name\":\"verifyHeaderAndExecuteTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteListContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

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
	"b44ec921": "whiteListContract(address)",
}

// EthCrossChainManagerBin is the compiled bytecode used for deploying new contracts.
var EthCrossChainManagerBin = "0x60806040523480156200001157600080fd5b506040516200562e3803806200562e833981810160405260808110156200003757600080fd5b815160208301516040808501805191519395929483019291846401000000008211156200006357600080fd5b9083019060208201858111156200007957600080fd5b82518660208202830111640100000000821117156200009757600080fd5b82525081516020918201928201910280838360005b83811015620000c6578181015183820152602001620000ac565b5050505090500160405260200180516040519392919084640100000000821115620000f057600080fd5b9083019060208201858111156200010657600080fd5b82516401000000008111828201881017156200012157600080fd5b82525081516020918201929091019080838360005b838110156200015057818101518382015260200162000136565b50505050905090810190601f1680156200017e5780820380516001836020036101000a031916815260200191505b5060405250505083836000620001996200030f60201b60201c565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506000805460ff60a01b19168155600180546001600160401b03909316600160a01b02600160a01b600160e01b03196001600160a01b039095166001600160a01b03199094169390931793909316919091179091555b825181101562000290576001600260008584815181106200025557fe5b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff191691151591909117905560010162000238565b5060016003826040518082805190602001908083835b60208310620002c75780518252601f199092019160209182019101620002a6565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220805460ff1916931515939093179092555062000313945050505050565b3390565b61530b80620003236000396000f3fe608060405234801561001057600080fd5b50600436106101155760003560e01c80638456cb59116100a25780639a8a0592116100715780639a8a0592146105ef578063b44ec92114610613578063bd5cf62514610639578063d450e04c14610757578063f2fde38b14610a0f57610115565b80638456cb59146105335780638c5aea411461053b5780638da5cb5b146105df5780638f32d59b146105e757610115565b80635c975abb116100e95780635c975abb146104315780636f31031d14610439578063715018a61461045f5780637bcf7dc1146104695780637e724ff31461050d57610115565b8062ba16941461011a57806329dcf4ab1461013e57806334a773eb146103005780633f4ba83a14610429575b600080fd5b610122610a35565b604080516001600160a01b039092168252519081900360200190f35b6102ec6004803603606081101561015457600080fd5b810190602081018135600160201b81111561016e57600080fd5b82018360208201111561018057600080fd5b803590602001918460018302840111600160201b831117156101a157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156101f357600080fd5b82018360208201111561020557600080fd5b803590602001918460018302840111600160201b8311171561022657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561027857600080fd5b82018360208201111561028a57600080fd5b803590602001918460018302840111600160201b831117156102ab57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a44945050505050565b604080519115158252519081900360200190f35b6102ec6004803603604081101561031657600080fd5b810190602081018135600160201b81111561033057600080fd5b82018360208201111561034257600080fd5b803590602001918460018302840111600160201b8311171561036357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156103b557600080fd5b8201836020820111156103c757600080fd5b803590602001918460018302840111600160201b831117156103e857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611066945050505050565b6102ec61153c565b6102ec6116bb565b6102ec6004803603602081101561044f57600080fd5b50356001600160401b03166116cb565b610467611794565b005b6104676004803603602081101561047f57600080fd5b810190602081018135600160201b81111561049957600080fd5b8201836020820111156104ab57600080fd5b803590602001918460018302840111600160201b831117156104cc57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611825945050505050565b6102ec6004803603602081101561052357600080fd5b50356001600160a01b0316611a6b565b6102ec611b7a565b6102ec6004803603602081101561055157600080fd5b810190602081018135600160201b81111561056b57600080fd5b82018360208201111561057d57600080fd5b803590602001918460018302840111600160201b8311171561059e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611cef945050505050565b610122611d0f565b6102ec611d1e565b6105f7611d42565b604080516001600160401b039092168252519081900360200190f35b6102ec6004803603602081101561062957600080fd5b50356001600160a01b0316611d58565b6102ec6004803603608081101561064f57600080fd5b6001600160401b038235169190810190604081016020820135600160201b81111561067957600080fd5b82018360208201111561068b57600080fd5b803590602001918460018302840111600160201b831117156106ac57600080fd5b919390929091602081019035600160201b8111156106c957600080fd5b8201836020820111156106db57600080fd5b803590602001918460018302840111600160201b831117156106fc57600080fd5b919390929091602081019035600160201b81111561071957600080fd5b82018360208201111561072b57600080fd5b803590602001918460018302840111600160201b8311171561074c57600080fd5b509092509050611d6d565b6102ec600480360360a081101561076d57600080fd5b810190602081018135600160201b81111561078757600080fd5b82018360208201111561079957600080fd5b803590602001918460018302840111600160201b831117156107ba57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561080c57600080fd5b82018360208201111561081e57600080fd5b803590602001918460018302840111600160201b8311171561083f57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561089157600080fd5b8201836020820111156108a357600080fd5b803590602001918460018302840111600160201b831117156108c457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561091657600080fd5b82018360208201111561092857600080fd5b803590602001918460018302840111600160201b8311171561094957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561099b57600080fd5b8201836020820111156109ad57600080fd5b803590602001918460018302840111600160201b831117156109ce57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506124ac945050505050565b61046760048036036020811015610a2557600080fd5b50356001600160a01b0316612bfc565b6001546001600160a01b031681565b60008054600160a01b900460ff1615610a97576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b610a9f614d38565b610aa885612c4f565b90506000600160009054906101000a90046001600160a01b031690506000816001600160a01b0316635ac407906040518163ffffffff1660e01b815260040160206040518083038186803b158015610aff57600080fd5b505afa158015610b13573d6000803e3d6000fd5b505050506040513d6020811015610b2957600080fd5b5051606084015163ffffffff9182169250168110610b785760405162461bcd60e51b815260040180806020018281038252603e815260200180614f79603e913960400191505060405180910390fd5b6101408301516bffffffffffffffffffffffff1916610bc85760405162461bcd60e51b8152600401808060200182810382526025815260200180614e256025913960400191505060405180910390fd5b6060610cf6836001600160a01b03166369d480746040518163ffffffff1660e01b815260040160006040518083038186803b158015610c0657600080fd5b505afa158015610c1a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610c4357600080fd5b8101908080516040519392919084600160201b821115610c6257600080fd5b908301906020820185811115610c7757600080fd5b8251600160201b811182820188101715610c9057600080fd5b82525081516020918201929091019080838360005b83811015610cbd578181015183820152602001610ca5565b50505050905090810190601f168015610cea5780820380516001836020036101000a031916815260200191505b50604052505050612d6e565b8051909150610d11898884600360001986015b048503612e21565b610d62576040805162461bcd60e51b815260206004820152601860248201527f566572696679207369676e6174757265206661696c6564210000000000000000604482015290519081900360640190fd5b60006060610d6f8a613020565b91509150816001600160601b0319168761014001516001600160601b03191614610dd6576040805162461bcd60e51b815260206004820152601360248201527213995e1d109bdbdad95c9cc81a5b1b1959d85b606a1b604482015290519081900360640190fd5b856001600160a01b0316638a8bd17f88606001516040518263ffffffff1660e01b8152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b158015610e2c57600080fd5b505af1158015610e40573d6000803e3d6000fd5b505050506040513d6020811015610e5657600080fd5b5051610e935760405162461bcd60e51b815260040180806020018281038252602d815260200180615093602d913960400191505060405180910390fd5b856001600160a01b03166341973cd9610eab836130ff565b6040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610ef7578181015183820152602001610edf565b50505050905090810190601f168015610f245780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b158015610f4357600080fd5b505af1158015610f57573d6000803e3d6000fd5b505050506040513d6020811015610f6d57600080fd5b5051610faa5760405162461bcd60e51b815260040180806020018281038252603b815260200180614fe2603b913960400191505060405180910390fd5b7fe60d33488cba3977bf65766cd2f8ac9617f64bf3b3198aff6240ce5c7d43b69087606001518c604051808363ffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561101a578181015183820152602001611002565b50505050905090810190601f1680156110475780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15060019a9950505050505050505050565b60008054600160a01b900460ff16156110b9576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b60015460408051631a75201d60e21b815290516001600160a01b039092169182916369d48074916004808301926000929190829003018186803b1580156110ff57600080fd5b505afa158015611113573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561113c57600080fd5b8101908080516040519392919084600160201b82111561115b57600080fd5b90830190602082018581111561117057600080fd5b8251600160201b81118282018810171561118957600080fd5b82525081516020918201929091019080838360005b838110156111b657818101518382015260200161119e565b50505050905090810190601f1680156111e35780820380516001836020036101000a031916815260200191505b50604052505050516000146112295760405162461bcd60e51b81526004018080602001828103825260388152602001806151786038913960400191505060405180910390fd5b611231614d38565b61123a85612c4f565b90506000606061124986613020565b91509150816001600160601b0319168361014001516001600160601b031916146112b0576040805162461bcd60e51b815260206004820152601360248201527213995e1d109bdbdad95c9cc81a5b1b1959d85b606a1b604482015290519081900360640190fd5b836001600160a01b0316638a8bd17f84606001516040518263ffffffff1660e01b8152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561130657600080fd5b505af115801561131a573d6000803e3d6000fd5b505050506040513d602081101561133057600080fd5b505161136d5760405162461bcd60e51b8152600401808060200182810382526043815260200180614ebc6043913960600191505060405180910390fd5b836001600160a01b03166341973cd9611385836130ff565b6040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156113d15781810151838201526020016113b9565b50505050905090810190601f1680156113fe5780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b15801561141d57600080fd5b505af1158015611431573d6000803e3d6000fd5b505050506040513d602081101561144757600080fd5b50516114845760405162461bcd60e51b81526004018080602001828103825260438152602001806151b06043913960600191505060405180910390fd5b7ff01968fc3a2655cf1b5144cb32de6dc898f91b9239c103744e8457152ab2fbde836060015188604051808363ffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b838110156114f45781810151838201526020016114dc565b50505050905090810190601f1680156115215780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15060019695505050505050565b6000611546611d1e565b611585576040805162461bcd60e51b8152602060048201819052602482015260008051602061501d833981519152604482015290519081900360640190fd5b61158d6116bb565b1561159a5761159a613201565b60015460408051635c975abb60e01b815290516001600160a01b03909216918291635c975abb916004808301926020929190829003018186803b1580156115e057600080fd5b505afa1580156115f4573d6000803e3d6000fd5b505050506040513d602081101561160a57600080fd5b5051156116b357806001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561164c57600080fd5b505af1158015611660573d6000803e3d6000fd5b505050506040513d602081101561167657600080fd5b50516116b35760405162461bcd60e51b8152600401808060200182810382526029815260200180614e4a6029913960400191505060405180910390fd5b600191505090565b600054600160a01b900460ff1690565b60008054600160a01b900460ff16611721576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b611729611d1e565b611768576040805162461bcd60e51b8152602060048201819052602482015260008051602061501d833981519152604482015290519081900360640190fd5b506001805467ffffffffffffffff60a01b1916600160a01b6001600160401b038416021781555b919050565b61179c611d1e565b6117db576040805162461bcd60e51b8152602060048201819052602482015260008051602061501d833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b600054600160a01b900460ff1615611877576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b6003816040518082805190602001908083835b602083106118a95780518252601f19909201916020918201910161188a565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff1691506119259050576040805162461bcd60e51b8152602060048201526015602482015274111bdb89dd08185c989a5d1c985c9a5b1e481cd95d605a1b604482015290519081900360640190fd5b60006003826040518082805190602001908083835b602083106119595780518252601f19909201916020918201910161193a565b51815160209384036101000a60001901801990921691161790529201948552506040519384900381018420805460ff19169515159590951790945550506001546341973cd960e01b8252600482018381528451602484015284516001600160a01b03909216936341973cd9938693839260449092019185019080838360005b838110156119f05781810151838201526020016119d8565b50505050905090810190601f168015611a1d5780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b158015611a3c57600080fd5b505af1158015611a50573d6000803e3d6000fd5b505050506040513d6020811015611a6657600080fd5b505050565b60008054600160a01b900460ff16611ac1576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b611ac9611d1e565b611b08576040805162461bcd60e51b8152602060048201819052602482015260008051602061501d833981519152604482015290519081900360640190fd5b6001546040805163f2fde38b60e01b81526001600160a01b03858116600483015291519190921691829163f2fde38b9160248082019260009290919082900301818387803b158015611b5957600080fd5b505af1158015611b6d573d6000803e3d6000fd5b5060019695505050505050565b6000611b84611d1e565b611bc3576040805162461bcd60e51b8152602060048201819052602482015260008051602061501d833981519152604482015290519081900360640190fd5b611bcb6116bb565b611bd757611bd76132a9565b60015460408051635c975abb60e01b815290516001600160a01b03909216918291635c975abb916004808301926020929190829003018186803b158015611c1d57600080fd5b505afa158015611c31573d6000803e3d6000fd5b505050506040513d6020811015611c4757600080fd5b50516116b357806001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401602060405180830381600087803b158015611c8857600080fd5b505af1158015611c9c573d6000803e3d6000fd5b505050506040513d6020811015611cb257600080fd5b50516116b35760405162461bcd60e51b8152600401808060200182810382526027815260200180614dfe6027913960400191505060405180910390fd5b805160208183018101805160038252928201919093012091525460ff1681565b6000546001600160a01b031690565b600080546001600160a01b0316611d33613333565b6001600160a01b031614905090565b600154600160a01b90046001600160401b031681565b60026020526000908152604090205460ff1681565b60008054600160a01b900460ff1615611dc0576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b60015460408051600162c2db5f60e01b0319815290516001600160a01b0390921691600091839163ff3d24a191600480820192602092909190829003018186803b158015611e0d57600080fd5b505afa158015611e21573d6000803e3d6000fd5b505050506040513d6020811015611e3757600080fd5b505190506060611e4682613337565b90506060611e53826133b0565b611f816002308560405160200180836001600160a01b03166001600160a01b031660601b815260140182805190602001908083835b60208310611ea75780518252601f199092019160209182019101611e88565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040526040518082805190602001908083835b60208310611f0b5780518252601f199092019160209182019101611eec565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015611f4a573d6000803e3d6000fd5b5050506040513d6020811015611f5f57600080fd5b50516040805160208181019390935281518082039093018352810190526133b0565b611f92611f8d33613476565b6133b0565b611f9b8f613491565b611fda8f8f8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506133b092505050565b6120198e8e8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506133b092505050565b6120588d8d8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506133b092505050565b6040516020018088805190602001908083835b6020831061208a5780518252601f19909201916020918201910161206b565b51815160209384036101000a60001901801990921691161790528a5191909301928a0191508083835b602083106120d25780518252601f1990920191602091820191016120b3565b51815160209384036101000a600019018019909216911617905289519190930192890191508083835b6020831061211a5780518252601f1990920191602091820191016120fb565b51815160209384036101000a600019018019909216911617905288519190930192880191508083835b602083106121625780518252601f199092019160209182019101612143565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b602083106121aa5780518252601f19909201916020918201910161218b565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106121f25780518252601f1990920191602091820191016121d3565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b6020831061223a5780518252601f19909201916020918201910161221b565b6001836020036101000a0380198251168184511680821785525050505050509050019750505050505050506040516020818303038152906040529050836001600160a01b0316634c3ccf6482805190602001206040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b1580156122c357600080fd5b505af11580156122d7573d6000803e3d6000fd5b505050506040513d60208110156122ed57600080fd5b505161232a5760405162461bcd60e51b81526004018080602001828103825260308152602001806152376030913960400191505060405180910390fd5b326001600160a01b03167f6ad3bf15c1988bc04bc153490cab16db8efb9a3990215bf1c64ea6e28be8848383338f8f8f876040518080602001876001600160a01b03166001600160a01b03168152602001866001600160401b03166001600160401b03168152602001806020018060200184810384528a818151815260200191508051906020019080838360005b838110156123d05781810151838201526020016123b8565b50505050905090810190601f1680156123fd5780820380516001836020036101000a031916815260200191505b5084810383528681526020018787808284376000838201819052601f909101601f191690920186810384528751815287516020918201939189019250908190849084905b83811015612459578181015183820152602001612441565b50505050905090810190601f1680156124865780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390a25060019b9a5050505050505050505050565b60008054600160a01b900460ff16156124ff576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b612507614d38565b61251086612c4f565b90506000600160009054906101000a90046001600160a01b03169050606061256a826001600160a01b03166369d480746040518163ffffffff1660e01b815260040160006040518083038186803b158015610c0657600080fd5b90506000826001600160a01b0316635ac407906040518163ffffffff1660e01b815260040160206040518083038186803b1580156125a757600080fd5b505afa1580156125bb573d6000803e3d6000fd5b505050506040513d60208110156125d157600080fd5b50518251606086015163ffffffff9283169350909116821161263f576125ff8a888560036000198601610d09565b61263a5760405162461bcd60e51b815260040180806020018281038252602a8152602001806152ad602a913960400191505060405180910390fd5b612719565b61265188888560036000198601610d09565b61268c5760405162461bcd60e51b81526004018080602001828103825260388152602001806150c06038913960400191505060405180910390fd5b612694614d38565b61269d89612c4f565b905060606126b08b8361010001516134d4565b90506126bb816135fa565b6126c48d61365a565b14612716576040805162461bcd60e51b815260206004820152601b60248201527f766572696679206865616465722070726f6f66206661696c6564210000000000604482015290519081900360640190fd5b50505b60606127298c8760e001516134d4565b9050612733614d93565b61273c8261378b565b9050856001600160a01b0316630586763c826020015161275f84600001516135fa565b6040518363ffffffff1660e01b815260040180836001600160401b03166001600160401b031681526020018281526020019250505060206040518083038186803b1580156127ac57600080fd5b505afa1580156127c0573d6000803e3d6000fd5b505050506040513d60208110156127d657600080fd5b5051156128145760405162461bcd60e51b81526004018080602001828103825260228152602001806152156022913960400191505060405180910390fd5b856001600160a01b031663e90bfdcf826020015161283584600001516135fa565b6040518363ffffffff1660e01b815260040180836001600160401b03166001600160401b0316815260200182815260200192505050602060405180830381600087803b15801561288457600080fd5b505af1158015612898573d6000803e3d6000fd5b505050506040513d60208110156128ae57600080fd5b5051612901576040805162461bcd60e51b815260206004820181905260248201527f536176652063726f7373636861696e207478206578697374206661696c656421604482015290519081900360640190fd5b6001546040820151606001516001600160401b03908116600160a01b909204161461295d5760405162461bcd60e51b8152600401808060200182810382526026815260200180614eff6026913960400191505060405180910390fd5b6000612970826040015160800151613867565b6001600160a01b03811660009081526002602052604090205490915060ff166129d5576040805162461bcd60e51b81526020600482015260126024820152711a5b9d985b1a59081d1bd0dbdb9d1c9858dd60721b604482015290519081900360640190fd5b6129fe81836040015160a00151846040015160c0015185604001516040015186602001516138b1565b612a4f576040805162461bcd60e51b815260206004820152601d60248201527f457865637574652043726f7373436861696e205478206661696c656421000000604482015290519081900360640190fd5b7f8a4a2663ce60ce4955c595da2894de0415240f1ace024cfbff85f513b656bdae8260200151836040015160800151846000015185604001516000015160405180856001600160401b03166001600160401b03168152602001806020018060200180602001848103845287818151815260200191508051906020019080838360005b83811015612ae9578181015183820152602001612ad1565b50505050905090810190601f168015612b165780820380516001836020036101000a031916815260200191505b50848103835286518152865160209182019188019080838360005b83811015612b49578181015183820152602001612b31565b50505050905090810190601f168015612b765780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838360005b83811015612ba9578181015183820152602001612b91565b50505050905090810190601f168015612bd65780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a15060019d9c50505050505050505050505050565b612c04611d1e565b612c43576040805162461bcd60e51b8152602060048201819052602482015260008051602061501d833981519152604482015290519081900360640190fd5b612c4c81613c85565b50565b612c57614d38565b612c5f614d38565b6000612c6b8482613d25565b63ffffffff90911683529050612c818482613dcd565b6001600160401b0390911660208401529050612c9d8482613e73565b60a08401919091529050612cb18482613e73565b60c08401919091529050612cc58482613e73565b60e08401919091529050612cd98482613e73565b6101008401919091529050612cee8482613d25565b63ffffffff90911660408401529050612d078482613d25565b63ffffffff90911660608401529050612d208482613dcd565b6001600160401b0390911660808401529050612d3c8482613eed565b6101208401919091529050612d518482613fc0565b506bffffffffffffffffffffffff19166101408301525092915050565b6060600080612d7d8482613dcd565b80935081925050506060816001600160401b0316604051908082528060200260200182016040528015612dba578160200160208202803883390190505b509050606060005b836001600160401b0316811015612e1657612ddd8786613eed565b95509150612dea82613867565b838281518110612df657fe5b6001600160a01b0390921660209283029190910190910152600101612dc2565b509095945050505050565b600080612e2d8661365a565b905060006041865181612e3c57fe5b049050606081604051908082528060200260200182016040528015612e6b578160200160208202803883390190505b50905060008080805b8581101561300557612e93612e8e8c604184026020614028565b6135fa565b9350612eaa612e8e8c604184026020016020614028565b92508a6041820260400181518110612ebe57fe5b602001015160f81c60f81b60f81c601b0191506001600288604051602001808281526020019150506040516020818303038152906040526040518082805190602001908083835b60208310612f245780518252601f199092019160209182019101612f05565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015612f63573d6000803e3d6000fd5b5050506040513d6020811015612f7857600080fd5b5051604080516000815260208181018084529390935260ff8616818301526060810188905260808101879052905160a08083019392601f198301929081900390910190855afa158015612fcf573d6000803e3d6000fd5b50505060206040510351858281518110612fe557fe5b6001600160a01b0390921660209283029190910190910152600101612e74565b5061301189858a6140a8565b9b9a5050505050505050505050565b60006060604383518161302f57fe5b0615613082576040805162461bcd60e51b815260206004820152601b60248201527f5f7075624b65794c697374206c656e67746820696c6c6567616c210000000000604482015290519081900360640190fd5b6000604384518161308f57fe5b04905060018110156130e1576040805162461bcd60e51b8152602060048201526016602482015275746f6f2073686f7274205f7075624b65794c6973742160501b604482015290519081900360640190fd5b6130f5816003600019820104830386614149565b9250925050915091565b80516060908161310e82613491565b905060005b828110156131f9578161313b611f8d87848151811061312e57fe5b6020026020010151613476565b6040516020018083805190602001908083835b6020831061316d5780518252601f19909201916020918201910161314e565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106131b55780518252601f199092019160209182019101613196565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405160208183030381529060405291508080600101915050613113565b509392505050565b600054600160a01b900460ff16613256576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa61328c613333565b604080516001600160a01b039092168252519081900360200190a1565b600054600160a01b900460ff16156132fb576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861328c5b3390565b60606001600160ff1b03821115613395576040805162461bcd60e51b815260206004820152601760248201527f56616c75652065786365656473207468652072616e6765000000000000000000604482015290519081900360640190fd5b60405190506020815281602082015260408101604052919050565b80516060906133be81614496565b836040516020018083805190602001908083835b602083106133f15780518252601f1990920191602091820191016133d2565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106134395780518252601f19909201916020918201910161341a565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052915050919050565b604080516014815260609290921b6020830152818101905290565b6040516008808252606091906000601f5b828210156134c45785811a8260208601015360019190910190600019016134a2565b5050506028810160405292915050565b60606000816134e38583613eed565b9250905060006134f2826145dc565b905060006021848851038161350357fe5b049050600080805b838110156135ae5761351d8a88614683565b9750915061352b8a88613e73565b975092506001600160f81b0319821661354f576135488386614701565b94506135a6565b600160f81b6001600160f81b03198316141561356f576135488584614701565b60405162461bcd60e51b815260040180806020018281038252602e81526020018061503d602e913960400191505060405180910390fd5b60010161350b565b508784146135ed5760405162461bcd60e51b8152600401808060200182810382526031815260200180614f486031913960400191505060405180910390fd5b5092979650505050505050565b60008151602014613652576040805162461bcd60e51b815260206004820152601760248201527f6279746573206c656e677468206973206e6f742033322e000000000000000000604482015290519081900360640190fd5b506020015190565b6000600280836040518082805190602001908083835b6020831061368f5780518252601f199092019160209182019101613670565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa1580156136ce573d6000803e3d6000fd5b5050506040513d60208110156136e357600080fd5b50516040805160208181019390935281518082038401815290820191829052805190928291908401908083835b6020831061372f5780518252601f199092019160209182019101613710565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa15801561376e573d6000803e3d6000fd5b5050506040513d602081101561378357600080fd5b505192915050565b613793614d93565b61379b614d93565b60006137a78482613eed565b90835290506137b68482613dcd565b6001600160401b03909116602084015290506137d0614db7565b6137da8583613eed565b90825291506137e98583613eed565b602083019190915291506137fd8583613eed565b604083019190915291506138118583613dcd565b6001600160401b039091166060830152915061382d8583613eed565b608083019190915291506138418583613eed565b60a083019190915291506138558583613eed565b5060c082015260408301525092915050565b600081516014146138a95760405162461bcd60e51b8152600401808060200182810382526023815260200180614e736023913960400191505060405180910390fd5b506014015190565b60006138bc866147c5565b6138f75760405162461bcd60e51b815260040180806020018281038252602881526020018061506b6028913960400191505060405180910390fd5b60606000876001600160a01b0316876040516020018082805190602001908083835b602083106139385780518252601f199092019160209182019101613919565b51815160001960209485036101000a01908116901991909116179052732862797465732c62797465732c75696e7436342960601b9390910192835260408051600b19818603018152601485019091528051908201206001600160401b038b1660748501526060603485019081528d5160948601528d519196508d95508c948c945090928392605483019260b4019188019080838360005b838110156139e75781810151838201526020016139cf565b50505050905090810190601f168015613a145780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b83811015613a47578181015183820152602001613a2f565b50505050905090810190601f168015613a745780820380516001836020036101000a031916815260200191505b509550505050505060405160208183030381529060405260405160200180836001600160e01b0319166001600160e01b031916815260040182805190602001908083835b60208310613ad75780518252601f199092019160209182019101613ab8565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040526040518082805190602001908083835b60208310613b3b5780518252601f199092019160209182019101613b1c565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114613b9d576040519150601f19603f3d011682016040523d82523d6000602084013e613ba2565b606091505b5092509050600181151514613be85760405162461bcd60e51b815260040180806020018281038252602b815260200180614fb7602b913960400191505060405180910390fd5b8151613c255760405162461bcd60e51b81526004018080602001828103825260278152602001806150f86027913960400191505060405180910390fd5b6000613c3283601f614801565b509050600181151514613c765760405162461bcd60e51b815260040180806020018281038252603781526020018061511f6037913960400191505060405180910390fd5b50600198975050505050505050565b6001600160a01b038116613cca5760405162461bcd60e51b8152600401808060200182810382526026815260200180614e966026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b60008083518360040111158015613d3e57508260040183105b613d795760405162461bcd60e51b81526004018080602001828103825260228152602001806152676022913960400191505060405180910390fd5b600060405160046000600182038760208a0101515b83831015613dae5780821a83860153600183019250600182039150613d8e565b50505080820160405260200390035192505050600482015b9250929050565b60008083518360080111158015613de657508260080183105b613e215760405162461bcd60e51b81526004018080602001828103825260228152602001806151f36022913960400191505060405180910390fd5b600060405160086000600182038760208a0101515b83831015613e565780821a83860153600183019250600182039150613e36565b505050808201604052602003900351956008949094019450505050565b60008083518360200111158015613e8c57508260200183105b613edd576040805162461bcd60e51b815260206004820181905260248201527f4e657874486173682c206f66667365742065786365656473206d6178696d756d604482015290519081900360640190fd5b5050602091810182015192910190565b6060600080613efc85856148f2565b865190955090915081850111801590613f16575080840184105b613f515760405162461bcd60e51b81526004018080602001828103825260248152602001806152896024913960400191505060405180910390fd5b606081158015613f6c57604051915060208201604052613fb6565b6040519150601f8316801560200281840101848101888315602002848c0101015b81831015613fa5578051835260209283019201613f8d565b5050848452601f01601f1916604052505b5095930193505050565b60008083518360140111158015613fd957508260140183105b6140145760405162461bcd60e51b8152600401808060200182810382526023815260200180614f256023913960400191505060405180910390fd5b505081810160200151601482019250929050565b60608183018451101561403a57600080fd5b6060821580156140555760405191506020820160405261409f565b6040519150601f8416801560200281840101858101878315602002848b0101015b8183101561408e578051835260209283019201614076565b5050858452601f01601f1916604052505b50949350505050565b600080805b845181101561413d5760005b8651811015614134578681815181106140ce57fe5b60200260200101516001600160a01b03168683815181106140eb57fe5b60200260200101516001600160a01b0316141561412c57828060010193505086818151811061411657fe5b6020026020010160006001600160a01b03168152505b6001016140b9565b506001016140ad565b50909111159392505050565b600060608061415786614b0b565b9050606086604051908082528060200260200182016040528015614185578160200160208202803883390190505b50905060006060815b898110156142ab576141a588604383026043614028565b9150846141b4611f8d84614b4e565b6040516020018083805190602001908083835b602083106141e65780518252601f1990920191602091820191016141c7565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b6020831061422e5780518252601f19909201916020918201910161420f565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405160208183030381529060405294506142728260036040614028565b8051906020012092508260001c84828151811061428b57fe5b6001600160a01b039092166020928302919091019091015260010161418e565b50836142b689614b0b565b6040516020018083805190602001908083835b602083106142e85780518252601f1990920191602091820191016142c9565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106143305780518252601f199092019160209182019101614311565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529350600060036002866040518082805190602001908083835b6020831061439d5780518252601f19909201916020918201910161437e565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa1580156143dc573d6000803e3d6000fd5b5050506040513d60208110156143f157600080fd5b50516040805160208181019390935281518082038401815290820191829052805190928291908401908083835b6020831061443d5780518252601f19909201916020918201910161441e565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa15801561447c573d6000803e3d6000fd5b50506040515160601b9b949a509398505050505050505050565b606060fd826001600160401b031610156144ba576144b382614c3c565b905061178f565b61ffff826001600160401b031611614598576144d960fd60f81b614c58565b6144e283614b0b565b6040516020018083805190602001908083835b602083106145145780518252601f1990920191602091820191016144f5565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b6020831061455c5780518252601f19909201916020918201910161453d565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052905061178f565b63ffffffff826001600160401b0316116145c2576145b9607f60f91b614c58565b6144e283614c6c565b6145d36001600160f81b0319614c58565b6144e283613491565b6040516000602080830182815284519293600293859387939260210191908401908083835b602083106146205780518252601f199092019160209182019101614601565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040526040518082805190602001908083836020831061372f5780518252601f199092019160209182019101613710565b6000808351836001011115801561469c57508260010183105b6146ed576040805162461bcd60e51b815260206004820181905260248201527f4e657874427974652c204f66667365742065786365656473206d6178696d756d604482015290519081900360640190fd5b505081810160200151600182019250929050565b60408051600160f81b60208083019190915260218201859052604180830185905283518084039091018152606190920192839052815160009360029392909182918401908083835b602083106147685780518252601f199092019160209182019101614749565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa1580156147a7573d6000803e3d6000fd5b5050506040513d60208110156147bc57600080fd5b50519392505050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47081158015906147f95750808214155b949350505050565b6000808351836001011115801561481a57508260010183105b614862576040805162461bcd60e51b815260206004820152601460248201527313d9999cd95d08195e18d959591cc81b1a5b5a5d60621b604482015290519081900360640190fd5b838301602001516000600160f81b6001600160f81b031983161415614889575060016148e4565b6001600160f81b031982166148a0575060006148e4565b6040805162461bcd60e51b81526020600482015260146024820152732732bc3a2137b7b6103b30b63ab29032b93937b960611b604482015290519081900360640190fd5b956001949094019450505050565b60008060006149018585614683565b94509050600060fd60f81b6001600160f81b03198316141561499f576149278686614caf565b955061ffff16905060fd8110801590614942575061ffff8111155b614993576040805162461bcd60e51b815260206004820152601f60248201527f4e65787455696e7431362c2076616c7565206f7574736964652072616e676500604482015290519081900360640190fd5b9250839150613dc69050565b607f60f91b6001600160f81b031983161415614a2f576149bf8686613d25565b955063ffffffff16905061ffff811180156149de575063ffffffff8111155b614993576040805162461bcd60e51b815260206004820181905260248201527f4e65787456617255696e742c2076616c7565206f7574736964652072616e6765604482015290519081900360640190fd5b6001600160f81b03198083161415614ab057614a4b8686613dcd565b95506001600160401b0316905063ffffffff8111614993576040805162461bcd60e51b815260206004820181905260248201527f4e65787456617255696e742c2076616c7565206f7574736964652072616e6765604482015290519081900360640190fd5b5060f881901c60fd8110614993576040805162461bcd60e51b815260206004820181905260248201527f4e65787456617255696e742c2076616c7565206f7574736964652072616e6765604482015290519081900360640190fd5b6040516002808252606091906000601f5b82821015614b3e5785811a826020860101536001919091019060001901614b1c565b5050506022810160405292915050565b6060604382511015614ba7576040805162461bcd60e51b815260206004820152601760248201527f6b6579206c656e67676820697320746f6f2073686f7274000000000000000000604482015290519081900360640190fd5b614bb48260006023614028565b9050600282604281518110614bc557fe5b016020015160f81c81614bd457fe5b0660ff1660001415614c0e57600260f81b81600281518110614bf257fe5b60200101906001600160f81b031916908160001a90535061178f565b600360f81b81600281518110614c2057fe5b60200101906001600160f81b031916908160001a905350919050565b604080516001815260f89290921b602083015260218201905290565b6060614c668260f81c614c3c565b92915050565b6040516004808252606091906000601f5b82821015614c9f5785811a826020860101536001919091019060001901614c7d565b5050506024810160405292915050565b60008083518360020111158015614cc857508260020183105b614d035760405162461bcd60e51b81526004018080602001828103825260228152602001806151566022913960400191505060405180910390fd5b6000604051846020870101518060011a82538060001a6001830153506002818101604052601d19909101519694019450505050565b604080516101608101825260008082526020820181905291810182905260608082018390526080820183905260a0820183905260c0820183905260e08201839052610100820183905261012082015261014081019190915290565b6040805160608082018352815260006020820152908101614db2614db7565b905290565b6040518060e0016040528060608152602001606081526020016060815260200160006001600160401b03168152602001606081526020016060815260200160608152509056fe70617573652045746843726f7373436861696e4461746120636f6e7472616374206661696c6564546865206e657874426f6f6b4b6565706572206f662068656164657220697320656d707479756e70617573652045746843726f7373436861696e4461746120636f6e7472616374206661696c65646279746573206c656e67746820646f6573206e6f74206d6174636820616464726573734f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573735361766520506f6c7920636861696e2063757272656e742065706f63682073746172742068656967687420746f204461746120636f6e7472616374206661696c65642154686973205478206973206e6f742061696d696e672061742074686973206e6574776f726b214e657874427974657332302c206f66667365742065786365656473206d6178696d756d6d65726b6c6550726f76652c2065787065637420726f6f74206973206e6f7420657175616c2061637475616c20726f6f7454686520686569676874206f6620686561646572206973206c6f776572207468616e2063757272656e742065706f6368207374617274206865696768742145746843726f7373436861696e2063616c6c20627573696e65737320636f6e7472616374206661696c65645361766520506f6c7920636861696e20626f6f6b206b65657065727320627974657320746f204461746120636f6e7472616374206661696c6564214f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726d65726b6c6550726f76652c204e6578744279746520666f7220706f736974696f6e20696e666f206661696c65645468652070617373656420696e2061646472657373206973206e6f74206120636f6e74726163742153617665204d43204c617465737448656967687420746f204461746120636f6e7472616374206661696c65642156657269667920706f6c7920636861696e2063757272656e742065706f636820686561646572207369676e6174757265206661696c6564214e6f2072657475726e2076616c75652066726f6d20627573696e65737320636f6e74726163742145746843726f7373436861696e2063616c6c20627573696e65737320636f6e74726163742072657475726e206973206e6f7420747275654e65787455696e7431362c206f66667365742065786365656473206d6178696d756d45746843726f7373436861696e4461746120636f6e74726163742068617320616c7265616479206265656e20696e697469616c697a6564215361766520506f6c7920636861696e2063757272656e742065706f636820626f6f6b206b65657065727320746f204461746120636f6e7472616374206661696c6564214e65787455696e7436342c206f66667365742065786365656473206d6178696d756d746865207472616e73616374696f6e20686173206265656e20657865637574656421536176652065746854784861736820627920696e64657820746f204461746120636f6e7472616374206661696c6564214e65787455696e7433322c206f66667365742065786365656473206d6178696d756d4e65787456617242797465732c206f66667365742065786365656473206d6178696d756d56657269667920706f6c7920636861696e20686561646572207369676e6174757265206661696c656421a265627a7a72315820a5299a442205fc2f912e24a730ca2f3e8ab4627a6fc3f710535c8db4d5efa61f64736f6c63430005110032"

// DeployEthCrossChainManager deploys a new Ethereum contract, binding an instance of EthCrossChainManager to it.
func DeployEthCrossChainManager(auth *bind.TransactOpts, backend bind.ContractBackend, _eccd common.Address, _chainId uint64, contracts []common.Address, curEpochPkBytes []byte) (common.Address, *types.Transaction, *EthCrossChainManager, error) {
	parsed, err := abi.JSON(strings.NewReader(EthCrossChainManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthCrossChainManagerBin), backend, _eccd, _chainId, contracts, curEpochPkBytes)
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

// WhiteListContract is a free data retrieval call binding the contract method 0xb44ec921.
//
// Solidity: function whiteListContract(address ) view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerCaller) WhiteListContract(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EthCrossChainManager.contract.Call(opts, out, "whiteListContract", arg0)
	return *ret0, err
}

// WhiteListContract is a free data retrieval call binding the contract method 0xb44ec921.
//
// Solidity: function whiteListContract(address ) view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) WhiteListContract(arg0 common.Address) (bool, error) {
	return _EthCrossChainManager.Contract.WhiteListContract(&_EthCrossChainManager.CallOpts, arg0)
}

// WhiteListContract is a free data retrieval call binding the contract method 0xb44ec921.
//
// Solidity: function whiteListContract(address ) view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerCallerSession) WhiteListContract(arg0 common.Address) (bool, error) {
	return _EthCrossChainManager.Contract.WhiteListContract(&_EthCrossChainManager.CallOpts, arg0)
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


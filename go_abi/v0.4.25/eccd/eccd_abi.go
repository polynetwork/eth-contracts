// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eccd

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
const EthCrossChainDataABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"EthToPolyTxHashIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"fromChainId\",\"type\":\"uint64\"},{\"name\":\"fromChainTx\",\"type\":\"bytes32\"}],\"name\":\"checkIfFromChainTxExist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key1\",\"type\":\"bytes32\"},{\"name\":\"key2\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"putExtraData\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ExtraData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ethTxHashIndex\",\"type\":\"uint256\"}],\"name\":\"getEthTxHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key1\",\"type\":\"bytes32\"},{\"name\":\"key2\",\"type\":\"bytes32\"}],\"name\":\"getExtraData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"curEpochPkBytes\",\"type\":\"bytes\"}],\"name\":\"putCurEpochConPubKeyBytes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ethTxHash\",\"type\":\"bytes32\"}],\"name\":\"putEthTxHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EthToPolyTxHashMap\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurEpochStartHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CurEpochStartHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurEpochConPubKeyBytes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"curEpochStartHeight\",\"type\":\"uint32\"}],\"name\":\"putCurEpochStartHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"fromChainId\",\"type\":\"uint64\"},{\"name\":\"fromChainTx\",\"type\":\"bytes32\"}],\"name\":\"markFromChainTxExist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ConKeepersPkBytes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEthTxHashIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// EthCrossChainDataFuncSigs maps the 4-byte function signature to its string representation.
var EthCrossChainDataFuncSigs = map[string]string{
	"fcbdc1e0": "ConKeepersPkBytes()",
	"67e31a74": "CurEpochStartHeight()",
	"00c5fff8": "EthToPolyTxHashIndex()",
	"529caad8": "EthToPolyTxHashMap(uint256)",
	"20bbde38": "ExtraData(bytes32,bytes32)",
	"0586763c": "checkIfFromChainTxExist(uint64,bytes32)",
	"69d48074": "getCurEpochConPubKeyBytes()",
	"5ac40790": "getCurEpochStartHeight()",
	"29927875": "getEthTxHash(uint256)",
	"ff3d24a1": "getEthTxHashIndex()",
	"40602bb5": "getExtraData(bytes32,bytes32)",
	"8f32d59b": "isOwner()",
	"e90bfdcf": "markFromChainTxExist(uint64,bytes32)",
	"8da5cb5b": "owner()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"41973cd9": "putCurEpochConPubKeyBytes(bytes)",
	"8a8bd17f": "putCurEpochStartHeight(uint32)",
	"4c3ccf64": "putEthTxHash(bytes32)",
	"1afe374e": "putExtraData(bytes32,bytes32,bytes)",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
	"3f4ba83a": "unpause()",
}

// EthCrossChainDataBin is the compiled bytecode used for deploying new contracts.
var EthCrossChainDataBin = "0x60806040526000610017640100000000610076810204565b60008054600160a060020a031916600160a060020a0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506000805460a060020a60ff021916905561007a565b3390565b6110d0806100896000396000f3006080604052600436106101315763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041662c5fff881146101365780630586763c1461015d5780631afe374e1461019657806320bbde38146101be578063299278751461024e5780633f4ba83a1461026657806340602bb51461027b57806341973cd9146102965780634c3ccf64146102b6578063529caad8146102ce5780635ac40790146102e65780635c975abb1461031457806367e31a741461032957806369d480741461033e578063715018a6146103535780638456cb591461036a5780638a8bd17f1461037f5780638da5cb5b1461039d5780638f32d59b146103ce578063e90bfdcf146103e3578063f2fde38b14610408578063fcbdc1e014610429578063ff3d24a11461043e575b600080fd5b34801561014257600080fd5b5061014b610453565b60408051918252519081900360200190f35b34801561016957600080fd5b5061018267ffffffffffffffff60043516602435610459565b604080519115158252519081900360200190f35b3480156101a257600080fd5b5061018260048035906024803591604435918201910135610485565b3480156101ca57600080fd5b506101d9600435602435610552565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102135781810151838201526020016101fb565b50505050905090810190601f1680156102405780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561025a57600080fd5b5061014b6004356105f6565b34801561027257600080fd5b50610182610608565b34801561028757600080fd5b506101d96004356024356106c8565b3480156102a257600080fd5b506101826004803560248101910135610772565b3480156102c257600080fd5b50610182600435610825565b3480156102da57600080fd5b5061014b6004356108e5565b3480156102f257600080fd5b506102fb6108f7565b6040805163ffffffff9092168252519081900360200190f35b34801561032057600080fd5b50610182610903565b34801561033557600080fd5b506102fb610913565b34801561034a57600080fd5b506101d961091f565b34801561035f57600080fd5b506103686109b5565b005b34801561037657600080fd5b50610182610a58565b34801561038b57600080fd5b5061018263ffffffff60043516610afe565b3480156103a957600080fd5b506103b2610bba565b60408051600160a060020a039092168252519081900360200190f35b3480156103da57600080fd5b50610182610bc9565b3480156103ef57600080fd5b5061018267ffffffffffffffff60043516602435610bed565b34801561041457600080fd5b50610368600160a060020a0360043516610cc1565b34801561043557600080fd5b506101d9610d19565b34801561044a57600080fd5b5061014b610d74565b60025481565b67ffffffffffffffff919091166000908152600560209081526040808320938352929052205460ff1690565b6000805460a060020a900460ff16156104d6576040805160e560020a62461bcd0281526020600482015260106024820152600080516020611085833981519152604482015290519081900360640190fd5b6104de610bc9565b1515610522576040805160e560020a62461bcd0281526020600482018190526024820152600080516020611065833981519152604482015290519081900360640190fd5b60008581526006602090815260408083208784529091529020610546908484610fcc565b50600195945050505050565b60066020908152600092835260408084208252918352918190208054825160026001831615610100026000190190921691909104601f8101859004850282018501909352828152929091908301828280156105ee5780601f106105c3576101008083540402835291602001916105ee565b820191906000526020600020905b8154815290600101906020018083116105d157829003601f168201915b505050505081565b60009081526001602052604090205490565b6000610612610bc9565b1515610656576040805160e560020a62461bcd0281526020600482018190526024820152600080516020611065833981519152604482015290519081900360640190fd5b60005460a060020a900460ff1615156106b9576040805160e560020a62461bcd02815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015290519081900360640190fd5b6106c1610d7a565b5060015b90565b600082815260066020908152604080832084845282529182902080548351601f60026101006001851615026000190190931692909204918201849004840281018401909452808452606093928301828280156107655780601f1061073a57610100808354040283529160200191610765565b820191906000526020600020905b81548152906001019060200180831161074857829003601f168201915b5050505050905092915050565b6000805460a060020a900460ff16156107c3576040805160e560020a62461bcd0281526020600482015260106024820152600080516020611085833981519152604482015290519081900360640190fd5b6107cb610bc9565b151561080f576040805160e560020a62461bcd0281526020600482018190526024820152600080516020611065833981519152604482015290519081900360640190fd5b61081b60038484610fcc565b5060019392505050565b6000805460a060020a900460ff1615610876576040805160e560020a62461bcd0281526020600482015260106024820152600080516020611085833981519152604482015290519081900360640190fd5b61087e610bc9565b15156108c2576040805160e560020a62461bcd0281526020600482018190526024820152600080516020611065833981519152604482015290519081900360640190fd5b506002805460009081526001602081905260409091209290925580548201905590565b60016020526000908152604090205481565b60045463ffffffff1690565b60005460a060020a900460ff1690565b60045463ffffffff1681565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156109ab5780601f10610980576101008083540402835291602001916109ab565b820191906000526020600020905b81548152906001019060200180831161098e57829003601f168201915b5050505050905090565b6109bd610bc9565b1515610a01576040805160e560020a62461bcd0281526020600482018190526024820152600080516020611065833981519152604482015290519081900360640190fd5b60008054604051600160a060020a03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a36000805473ffffffffffffffffffffffffffffffffffffffff19169055565b6000610a62610bc9565b1515610aa6576040805160e560020a62461bcd0281526020600482018190526024820152600080516020611065833981519152604482015290519081900360640190fd5b60005460a060020a900460ff1615610af6576040805160e560020a62461bcd0281526020600482015260106024820152600080516020611085833981519152604482015290519081900360640190fd5b6106c1610e41565b6000805460a060020a900460ff1615610b4f576040805160e560020a62461bcd0281526020600482015260106024820152600080516020611085833981519152604482015290519081900360640190fd5b610b57610bc9565b1515610b9b576040805160e560020a62461bcd0281526020600482018190526024820152600080516020611065833981519152604482015290519081900360640190fd5b506004805463ffffffff831663ffffffff199091161790556001919050565b600054600160a060020a031690565b60008054600160a060020a0316610bde610eda565b600160a060020a031614905090565b6000805460a060020a900460ff1615610c3e576040805160e560020a62461bcd0281526020600482015260106024820152600080516020611085833981519152604482015290519081900360640190fd5b610c46610bc9565b1515610c8a576040805160e560020a62461bcd0281526020600482018190526024820152600080516020611065833981519152604482015290519081900360640190fd5b5067ffffffffffffffff91909116600090815260056020908152604080832093835292905220805460ff1916600190811790915590565b610cc9610bc9565b1515610d0d576040805160e560020a62461bcd0281526020600482018190526024820152600080516020611065833981519152604482015290519081900360640190fd5b610d1681610ede565b50565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156105ee5780601f106105c3576101008083540402835291602001916105ee565b60025490565b60005460a060020a900460ff161515610ddd576040805160e560020a62461bcd02815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015290519081900360640190fd5b6000805474ff0000000000000000000000000000000000000000191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa610e24610eda565b60408051600160a060020a039092168252519081900360200190a1565b60005460a060020a900460ff1615610e91576040805160e560020a62461bcd0281526020600482015260106024820152600080516020611085833981519152604482015290519081900360640190fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610e245b3390565b600160a060020a0381161515610f64576040805160e560020a62461bcd02815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061100d5782800160ff1982351617855561103a565b8280016001018555821561103a579182015b8281111561103a57823582559160200191906001019061101f565b5061104692915061104a565b5090565b6106c591905b80821115611046576000815560010161105056004f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725061757361626c653a2070617573656400000000000000000000000000000000a165627a7a723058209aad2591f6a75eac5e63d2f2d6feff1f41a72841feeb7e5b77177e15f907c30c0029"

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

// ConKeepersPkBytes is a free data retrieval call binding the contract method 0xfcbdc1e0.
//
// Solidity: function ConKeepersPkBytes() view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCaller) ConKeepersPkBytes(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "ConKeepersPkBytes")
	return *ret0, err
}

// ConKeepersPkBytes is a free data retrieval call binding the contract method 0xfcbdc1e0.
//
// Solidity: function ConKeepersPkBytes() view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataSession) ConKeepersPkBytes() ([]byte, error) {
	return _EthCrossChainData.Contract.ConKeepersPkBytes(&_EthCrossChainData.CallOpts)
}

// ConKeepersPkBytes is a free data retrieval call binding the contract method 0xfcbdc1e0.
//
// Solidity: function ConKeepersPkBytes() view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCallerSession) ConKeepersPkBytes() ([]byte, error) {
	return _EthCrossChainData.Contract.ConKeepersPkBytes(&_EthCrossChainData.CallOpts)
}

// CurEpochStartHeight is a free data retrieval call binding the contract method 0x67e31a74.
//
// Solidity: function CurEpochStartHeight() view returns(uint32)
func (_EthCrossChainData *EthCrossChainDataCaller) CurEpochStartHeight(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "CurEpochStartHeight")
	return *ret0, err
}

// CurEpochStartHeight is a free data retrieval call binding the contract method 0x67e31a74.
//
// Solidity: function CurEpochStartHeight() view returns(uint32)
func (_EthCrossChainData *EthCrossChainDataSession) CurEpochStartHeight() (uint32, error) {
	return _EthCrossChainData.Contract.CurEpochStartHeight(&_EthCrossChainData.CallOpts)
}

// CurEpochStartHeight is a free data retrieval call binding the contract method 0x67e31a74.
//
// Solidity: function CurEpochStartHeight() view returns(uint32)
func (_EthCrossChainData *EthCrossChainDataCallerSession) CurEpochStartHeight() (uint32, error) {
	return _EthCrossChainData.Contract.CurEpochStartHeight(&_EthCrossChainData.CallOpts)
}

// EthToPolyTxHashIndex is a free data retrieval call binding the contract method 0x00c5fff8.
//
// Solidity: function EthToPolyTxHashIndex() view returns(uint256)
func (_EthCrossChainData *EthCrossChainDataCaller) EthToPolyTxHashIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "EthToPolyTxHashIndex")
	return *ret0, err
}

// EthToPolyTxHashIndex is a free data retrieval call binding the contract method 0x00c5fff8.
//
// Solidity: function EthToPolyTxHashIndex() view returns(uint256)
func (_EthCrossChainData *EthCrossChainDataSession) EthToPolyTxHashIndex() (*big.Int, error) {
	return _EthCrossChainData.Contract.EthToPolyTxHashIndex(&_EthCrossChainData.CallOpts)
}

// EthToPolyTxHashIndex is a free data retrieval call binding the contract method 0x00c5fff8.
//
// Solidity: function EthToPolyTxHashIndex() view returns(uint256)
func (_EthCrossChainData *EthCrossChainDataCallerSession) EthToPolyTxHashIndex() (*big.Int, error) {
	return _EthCrossChainData.Contract.EthToPolyTxHashIndex(&_EthCrossChainData.CallOpts)
}

// EthToPolyTxHashMap is a free data retrieval call binding the contract method 0x529caad8.
//
// Solidity: function EthToPolyTxHashMap(uint256 ) view returns(bytes32)
func (_EthCrossChainData *EthCrossChainDataCaller) EthToPolyTxHashMap(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "EthToPolyTxHashMap", arg0)
	return *ret0, err
}

// EthToPolyTxHashMap is a free data retrieval call binding the contract method 0x529caad8.
//
// Solidity: function EthToPolyTxHashMap(uint256 ) view returns(bytes32)
func (_EthCrossChainData *EthCrossChainDataSession) EthToPolyTxHashMap(arg0 *big.Int) ([32]byte, error) {
	return _EthCrossChainData.Contract.EthToPolyTxHashMap(&_EthCrossChainData.CallOpts, arg0)
}

// EthToPolyTxHashMap is a free data retrieval call binding the contract method 0x529caad8.
//
// Solidity: function EthToPolyTxHashMap(uint256 ) view returns(bytes32)
func (_EthCrossChainData *EthCrossChainDataCallerSession) EthToPolyTxHashMap(arg0 *big.Int) ([32]byte, error) {
	return _EthCrossChainData.Contract.EthToPolyTxHashMap(&_EthCrossChainData.CallOpts, arg0)
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

// CheckIfFromChainTxExist is a free data retrieval call binding the contract method 0x0586763c.
//
// Solidity: function checkIfFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) view returns(bool)
func (_EthCrossChainData *EthCrossChainDataCaller) CheckIfFromChainTxExist(opts *bind.CallOpts, fromChainId uint64, fromChainTx [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "checkIfFromChainTxExist", fromChainId, fromChainTx)
	return *ret0, err
}

// CheckIfFromChainTxExist is a free data retrieval call binding the contract method 0x0586763c.
//
// Solidity: function checkIfFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) view returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) CheckIfFromChainTxExist(fromChainId uint64, fromChainTx [32]byte) (bool, error) {
	return _EthCrossChainData.Contract.CheckIfFromChainTxExist(&_EthCrossChainData.CallOpts, fromChainId, fromChainTx)
}

// CheckIfFromChainTxExist is a free data retrieval call binding the contract method 0x0586763c.
//
// Solidity: function checkIfFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) view returns(bool)
func (_EthCrossChainData *EthCrossChainDataCallerSession) CheckIfFromChainTxExist(fromChainId uint64, fromChainTx [32]byte) (bool, error) {
	return _EthCrossChainData.Contract.CheckIfFromChainTxExist(&_EthCrossChainData.CallOpts, fromChainId, fromChainTx)
}

// GetCurEpochConPubKeyBytes is a free data retrieval call binding the contract method 0x69d48074.
//
// Solidity: function getCurEpochConPubKeyBytes() view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCaller) GetCurEpochConPubKeyBytes(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "getCurEpochConPubKeyBytes")
	return *ret0, err
}

// GetCurEpochConPubKeyBytes is a free data retrieval call binding the contract method 0x69d48074.
//
// Solidity: function getCurEpochConPubKeyBytes() view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataSession) GetCurEpochConPubKeyBytes() ([]byte, error) {
	return _EthCrossChainData.Contract.GetCurEpochConPubKeyBytes(&_EthCrossChainData.CallOpts)
}

// GetCurEpochConPubKeyBytes is a free data retrieval call binding the contract method 0x69d48074.
//
// Solidity: function getCurEpochConPubKeyBytes() view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetCurEpochConPubKeyBytes() ([]byte, error) {
	return _EthCrossChainData.Contract.GetCurEpochConPubKeyBytes(&_EthCrossChainData.CallOpts)
}

// GetCurEpochStartHeight is a free data retrieval call binding the contract method 0x5ac40790.
//
// Solidity: function getCurEpochStartHeight() view returns(uint32)
func (_EthCrossChainData *EthCrossChainDataCaller) GetCurEpochStartHeight(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "getCurEpochStartHeight")
	return *ret0, err
}

// GetCurEpochStartHeight is a free data retrieval call binding the contract method 0x5ac40790.
//
// Solidity: function getCurEpochStartHeight() view returns(uint32)
func (_EthCrossChainData *EthCrossChainDataSession) GetCurEpochStartHeight() (uint32, error) {
	return _EthCrossChainData.Contract.GetCurEpochStartHeight(&_EthCrossChainData.CallOpts)
}

// GetCurEpochStartHeight is a free data retrieval call binding the contract method 0x5ac40790.
//
// Solidity: function getCurEpochStartHeight() view returns(uint32)
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetCurEpochStartHeight() (uint32, error) {
	return _EthCrossChainData.Contract.GetCurEpochStartHeight(&_EthCrossChainData.CallOpts)
}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 ethTxHashIndex) view returns(bytes32)
func (_EthCrossChainData *EthCrossChainDataCaller) GetEthTxHash(opts *bind.CallOpts, ethTxHashIndex *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _EthCrossChainData.contract.Call(opts, out, "getEthTxHash", ethTxHashIndex)
	return *ret0, err
}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 ethTxHashIndex) view returns(bytes32)
func (_EthCrossChainData *EthCrossChainDataSession) GetEthTxHash(ethTxHashIndex *big.Int) ([32]byte, error) {
	return _EthCrossChainData.Contract.GetEthTxHash(&_EthCrossChainData.CallOpts, ethTxHashIndex)
}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 ethTxHashIndex) view returns(bytes32)
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetEthTxHash(ethTxHashIndex *big.Int) ([32]byte, error) {
	return _EthCrossChainData.Contract.GetEthTxHash(&_EthCrossChainData.CallOpts, ethTxHashIndex)
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

// MarkFromChainTxExist is a paid mutator transaction binding the contract method 0xe90bfdcf.
//
// Solidity: function markFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) MarkFromChainTxExist(opts *bind.TransactOpts, fromChainId uint64, fromChainTx [32]byte) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "markFromChainTxExist", fromChainId, fromChainTx)
}

// MarkFromChainTxExist is a paid mutator transaction binding the contract method 0xe90bfdcf.
//
// Solidity: function markFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) MarkFromChainTxExist(fromChainId uint64, fromChainTx [32]byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.MarkFromChainTxExist(&_EthCrossChainData.TransactOpts, fromChainId, fromChainTx)
}

// MarkFromChainTxExist is a paid mutator transaction binding the contract method 0xe90bfdcf.
//
// Solidity: function markFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) MarkFromChainTxExist(fromChainId uint64, fromChainTx [32]byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.MarkFromChainTxExist(&_EthCrossChainData.TransactOpts, fromChainId, fromChainTx)
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

// PutCurEpochConPubKeyBytes is a paid mutator transaction binding the contract method 0x41973cd9.
//
// Solidity: function putCurEpochConPubKeyBytes(bytes curEpochPkBytes) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) PutCurEpochConPubKeyBytes(opts *bind.TransactOpts, curEpochPkBytes []byte) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "putCurEpochConPubKeyBytes", curEpochPkBytes)
}

// PutCurEpochConPubKeyBytes is a paid mutator transaction binding the contract method 0x41973cd9.
//
// Solidity: function putCurEpochConPubKeyBytes(bytes curEpochPkBytes) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) PutCurEpochConPubKeyBytes(curEpochPkBytes []byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutCurEpochConPubKeyBytes(&_EthCrossChainData.TransactOpts, curEpochPkBytes)
}

// PutCurEpochConPubKeyBytes is a paid mutator transaction binding the contract method 0x41973cd9.
//
// Solidity: function putCurEpochConPubKeyBytes(bytes curEpochPkBytes) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) PutCurEpochConPubKeyBytes(curEpochPkBytes []byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutCurEpochConPubKeyBytes(&_EthCrossChainData.TransactOpts, curEpochPkBytes)
}

// PutCurEpochStartHeight is a paid mutator transaction binding the contract method 0x8a8bd17f.
//
// Solidity: function putCurEpochStartHeight(uint32 curEpochStartHeight) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) PutCurEpochStartHeight(opts *bind.TransactOpts, curEpochStartHeight uint32) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "putCurEpochStartHeight", curEpochStartHeight)
}

// PutCurEpochStartHeight is a paid mutator transaction binding the contract method 0x8a8bd17f.
//
// Solidity: function putCurEpochStartHeight(uint32 curEpochStartHeight) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) PutCurEpochStartHeight(curEpochStartHeight uint32) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutCurEpochStartHeight(&_EthCrossChainData.TransactOpts, curEpochStartHeight)
}

// PutCurEpochStartHeight is a paid mutator transaction binding the contract method 0x8a8bd17f.
//
// Solidity: function putCurEpochStartHeight(uint32 curEpochStartHeight) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) PutCurEpochStartHeight(curEpochStartHeight uint32) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutCurEpochStartHeight(&_EthCrossChainData.TransactOpts, curEpochStartHeight)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x4c3ccf64.
//
// Solidity: function putEthTxHash(bytes32 ethTxHash) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) PutEthTxHash(opts *bind.TransactOpts, ethTxHash [32]byte) (*types.Transaction, error) {
	return _EthCrossChainData.contract.Transact(opts, "putEthTxHash", ethTxHash)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x4c3ccf64.
//
// Solidity: function putEthTxHash(bytes32 ethTxHash) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) PutEthTxHash(ethTxHash [32]byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutEthTxHash(&_EthCrossChainData.TransactOpts, ethTxHash)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x4c3ccf64.
//
// Solidity: function putEthTxHash(bytes32 ethTxHash) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) PutEthTxHash(ethTxHash [32]byte) (*types.Transaction, error) {
	return _EthCrossChainData.Contract.PutEthTxHash(&_EthCrossChainData.TransactOpts, ethTxHash)
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
const IEthCrossChainDataABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"fromChainId\",\"type\":\"uint64\"},{\"name\":\"fromChainTx\",\"type\":\"bytes32\"}],\"name\":\"checkIfFromChainTxExist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key1\",\"type\":\"bytes32\"},{\"name\":\"key2\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"putExtraData\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ethTxHashIndex\",\"type\":\"uint256\"}],\"name\":\"getEthTxHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key1\",\"type\":\"bytes32\"},{\"name\":\"key2\",\"type\":\"bytes32\"}],\"name\":\"getExtraData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"curEpochPkBytes\",\"type\":\"bytes\"}],\"name\":\"putCurEpochConPubKeyBytes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ethTxHash\",\"type\":\"bytes32\"}],\"name\":\"putEthTxHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurEpochStartHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurEpochConPubKeyBytes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"curEpochStartHeight\",\"type\":\"uint32\"}],\"name\":\"putCurEpochStartHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"fromChainId\",\"type\":\"uint64\"},{\"name\":\"fromChainTx\",\"type\":\"bytes32\"}],\"name\":\"markFromChainTxExist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEthTxHashIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

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

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

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
const PausableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"}]"

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

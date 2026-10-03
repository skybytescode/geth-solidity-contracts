// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
	_ = time.Tick
	_ = context.Background
)

// EthSenderMetaData contains all meta data concerning the EthSender contract.
var EthSenderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountPerClaim_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cooldown_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nextClaimAt\",\"type\":\"uint256\"}],\"name\":\"TooSoon\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"amountPerClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cooldown\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getTotalSent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sendEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalSent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60e060405234801561000f575f5ffd5b506040516106c03803806106c083398101604081905261002e91610040565b3360805260a09190915260c052610062565b5f5f60408385031215610051575f5ffd5b505080516020909101519092909150565b60805160a05160c0516105ef6100d15f395f818161016901528181610228015261025a01525f81816098015281816102a6015281816102e10152818161032b0152818161036101526103be01525f81816101d001528181610451015281816104bb015261051f01526105ef5ff3fe60806040526004361061007c575f3560e01c80635c16e15e1161004c5780635c16e15e1461012d578063787a08a614610158578063806c0b2b1461018b5780638da5cb5b146101bf575f5ffd5b8063013d88471461008757806306e99fef146100cd57806308f81a93146100e35780632e1a7d4d1461010e575f5ffd5b3661008357005b5f5ffd5b348015610092575f5ffd5b506100ba7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020015b60405180910390f35b3480156100d8575f5ffd5b506100e161020a565b005b3480156100ee575f5ffd5b506100ba6100fd366004610550565b5f6020819052908152604090205481565b348015610119575f5ffd5b506100e161012836600461057d565b610446565b348015610138575f5ffd5b506100ba610147366004610550565b60016020525f908152604090205481565b348015610163575f5ffd5b506100ba7f000000000000000000000000000000000000000000000000000000000000000081565b348015610196575f5ffd5b506100ba6101a5366004610550565b6001600160a01b03165f9081526020819052604090205490565b3480156101ca575f5ffd5b506101f27f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100c4565b335f908152600160205260409020548015801590610250575061024d7f000000000000000000000000000000000000000000000000000000000000000082610594565b42105b156102a45761027f7f000000000000000000000000000000000000000000000000000000000000000082610594565b604051637437acf560e11b815260040161029b91815260200190565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000047101561030d5760405163cf47918160e01b81524760048201527f0000000000000000000000000000000000000000000000000000000000000000602482015260440161029b565b335f90815260016020908152604080832042905590829052812080547f00000000000000000000000000000000000000000000000000000000000000009290610357908490610594565b90915550506040517f0000000000000000000000000000000000000000000000000000000000000000815233907f78f5cdad99320ec2ba57132d7dffb1d125775c823239e60ff5e9300fd4ac898c9060200160405180910390a25f336001600160a01b03167f00000000000000000000000000000000000000000000000000000000000000006040515b5f6040518083038185875af1925050503d805f811461041b576040519150601f19603f3d011682016040523d82523d5f602084013e610420565b606091505b5050905080610442576040516312171d8360e31b815260040160405180910390fd5b5050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461048f576040516330cd747160e01b815260040160405180910390fd5b804710156104b95760405163cf47918160e01b81524760048201526024810182905260440161029b565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03167f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d58260405161051491815260200190565b60405180910390a25f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826040516103e1565b5f60208284031215610560575f5ffd5b81356001600160a01b0381168114610576575f5ffd5b9392505050565b5f6020828403121561058d575f5ffd5b5035919050565b808201808211156105b357634e487b7160e01b5f52601160045260245ffd5b9291505056fea26469706673582212203d2a5b2b7629243d51d5c0fc7726a722c0fec254a14ca4dd97cd1429b7b2be9c64736f6c634300081e0033",
}

// EthSenderABI is the input ABI used to generate the binding from.
// Deprecated: Use EthSenderMetaData.ABI instead.
var EthSenderABI = EthSenderMetaData.ABI

// EthSenderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthSenderMetaData.Bin instead.
var EthSenderBin = EthSenderMetaData.Bin

// DeployEthSender deploys a new Ethereum contract, binding an instance of EthSender to it.
func DeployEthSender(auth *bind.TransactOpts, backend bind.ContractBackend, amountPerClaim_ *big.Int, cooldown_ *big.Int) (common.Address, *types.Transaction, *EthSender, error) {
	parsed, err := EthSenderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthSenderBin), backend, amountPerClaim_, cooldown_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthSender{EthSenderCaller: EthSenderCaller{contract: contract}, EthSenderTransactor: EthSenderTransactor{contract: contract}, EthSenderFilterer: EthSenderFilterer{contract: contract}}, nil
}

// EthSender is an auto generated Go binding around an Ethereum contract.
type EthSender struct {
	EthSenderCaller     // Read-only binding to the contract
	EthSenderTransactor // Write-only binding to the contract
	EthSenderFilterer   // Log filterer for contract events
}

// EthSenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthSenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthSenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthSenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthSenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthSenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthSenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthSenderSession struct {
	Contract     *EthSender        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthSenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthSenderCallerSession struct {
	Contract *EthSenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// EthSenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthSenderTransactorSession struct {
	Contract     *EthSenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EthSenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthSenderRaw struct {
	Contract *EthSender // Generic contract binding to access the raw methods on
}

// EthSenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthSenderCallerRaw struct {
	Contract *EthSenderCaller // Generic read-only contract binding to access the raw methods on
}

// EthSenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthSenderTransactorRaw struct {
	Contract *EthSenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthSender creates a new instance of EthSender, bound to a specific deployed contract.
func NewEthSender(address common.Address, backend bind.ContractBackend) (*EthSender, error) {
	contract, err := bindEthSender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthSender{EthSenderCaller: EthSenderCaller{contract: contract}, EthSenderTransactor: EthSenderTransactor{contract: contract}, EthSenderFilterer: EthSenderFilterer{contract: contract}}, nil
}

// NewEthSenderCaller creates a new read-only instance of EthSender, bound to a specific deployed contract.
func NewEthSenderCaller(address common.Address, caller bind.ContractCaller) (*EthSenderCaller, error) {
	contract, err := bindEthSender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthSenderCaller{contract: contract}, nil
}

// NewEthSenderTransactor creates a new write-only instance of EthSender, bound to a specific deployed contract.
func NewEthSenderTransactor(address common.Address, transactor bind.ContractTransactor) (*EthSenderTransactor, error) {
	contract, err := bindEthSender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthSenderTransactor{contract: contract}, nil
}

// NewEthSenderFilterer creates a new log filterer instance of EthSender, bound to a specific deployed contract.
func NewEthSenderFilterer(address common.Address, filterer bind.ContractFilterer) (*EthSenderFilterer, error) {
	contract, err := bindEthSender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthSenderFilterer{contract: contract}, nil
}

// bindEthSender binds a generic wrapper to an already deployed contract.
func bindEthSender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EthSenderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthSender *EthSenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthSender.Contract.EthSenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthSender *EthSenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthSender.Contract.EthSenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthSender *EthSenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthSender.Contract.EthSenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthSender *EthSenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthSender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthSender *EthSenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthSender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthSender *EthSenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthSender.Contract.contract.Transact(opts, method, params...)
}

// AmountPerClaim is a free data retrieval call binding the contract method 0x013d8847.
//
// Solidity: function amountPerClaim() view returns(uint256)
func (_EthSender *EthSenderCaller) AmountPerClaim(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthSender.contract.Call(opts, &out, "amountPerClaim")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountPerClaim is a free data retrieval call binding the contract method 0x013d8847.
//
// Solidity: function amountPerClaim() view returns(uint256)
func (_EthSender *EthSenderSession) AmountPerClaim() (*big.Int, error) {
	return _EthSender.Contract.AmountPerClaim(&_EthSender.CallOpts)
}

// AmountPerClaim is a free data retrieval call binding the contract method 0x013d8847.
//
// Solidity: function amountPerClaim() view returns(uint256)
func (_EthSender *EthSenderCallerSession) AmountPerClaim() (*big.Int, error) {
	return _EthSender.Contract.AmountPerClaim(&_EthSender.CallOpts)
}

// Cooldown is a free data retrieval call binding the contract method 0x787a08a6.
//
// Solidity: function cooldown() view returns(uint256)
func (_EthSender *EthSenderCaller) Cooldown(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthSender.contract.Call(opts, &out, "cooldown")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cooldown is a free data retrieval call binding the contract method 0x787a08a6.
//
// Solidity: function cooldown() view returns(uint256)
func (_EthSender *EthSenderSession) Cooldown() (*big.Int, error) {
	return _EthSender.Contract.Cooldown(&_EthSender.CallOpts)
}

// Cooldown is a free data retrieval call binding the contract method 0x787a08a6.
//
// Solidity: function cooldown() view returns(uint256)
func (_EthSender *EthSenderCallerSession) Cooldown() (*big.Int, error) {
	return _EthSender.Contract.Cooldown(&_EthSender.CallOpts)
}

// GetTotalSent is a free data retrieval call binding the contract method 0x806c0b2b.
//
// Solidity: function getTotalSent(address account) view returns(uint256)
func (_EthSender *EthSenderCaller) GetTotalSent(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EthSender.contract.Call(opts, &out, "getTotalSent", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSent is a free data retrieval call binding the contract method 0x806c0b2b.
//
// Solidity: function getTotalSent(address account) view returns(uint256)
func (_EthSender *EthSenderSession) GetTotalSent(account common.Address) (*big.Int, error) {
	return _EthSender.Contract.GetTotalSent(&_EthSender.CallOpts, account)
}

// GetTotalSent is a free data retrieval call binding the contract method 0x806c0b2b.
//
// Solidity: function getTotalSent(address account) view returns(uint256)
func (_EthSender *EthSenderCallerSession) GetTotalSent(account common.Address) (*big.Int, error) {
	return _EthSender.Contract.GetTotalSent(&_EthSender.CallOpts, account)
}

// LastClaim is a free data retrieval call binding the contract method 0x5c16e15e.
//
// Solidity: function lastClaim(address ) view returns(uint256)
func (_EthSender *EthSenderCaller) LastClaim(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EthSender.contract.Call(opts, &out, "lastClaim", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastClaim is a free data retrieval call binding the contract method 0x5c16e15e.
//
// Solidity: function lastClaim(address ) view returns(uint256)
func (_EthSender *EthSenderSession) LastClaim(arg0 common.Address) (*big.Int, error) {
	return _EthSender.Contract.LastClaim(&_EthSender.CallOpts, arg0)
}

// LastClaim is a free data retrieval call binding the contract method 0x5c16e15e.
//
// Solidity: function lastClaim(address ) view returns(uint256)
func (_EthSender *EthSenderCallerSession) LastClaim(arg0 common.Address) (*big.Int, error) {
	return _EthSender.Contract.LastClaim(&_EthSender.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthSender *EthSenderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthSender.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthSender *EthSenderSession) Owner() (common.Address, error) {
	return _EthSender.Contract.Owner(&_EthSender.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthSender *EthSenderCallerSession) Owner() (common.Address, error) {
	return _EthSender.Contract.Owner(&_EthSender.CallOpts)
}

// TotalSent is a free data retrieval call binding the contract method 0x08f81a93.
//
// Solidity: function totalSent(address ) view returns(uint256)
func (_EthSender *EthSenderCaller) TotalSent(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EthSender.contract.Call(opts, &out, "totalSent", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSent is a free data retrieval call binding the contract method 0x08f81a93.
//
// Solidity: function totalSent(address ) view returns(uint256)
func (_EthSender *EthSenderSession) TotalSent(arg0 common.Address) (*big.Int, error) {
	return _EthSender.Contract.TotalSent(&_EthSender.CallOpts, arg0)
}

// TotalSent is a free data retrieval call binding the contract method 0x08f81a93.
//
// Solidity: function totalSent(address ) view returns(uint256)
func (_EthSender *EthSenderCallerSession) TotalSent(arg0 common.Address) (*big.Int, error) {
	return _EthSender.Contract.TotalSent(&_EthSender.CallOpts, arg0)
}

// SendEth is a paid mutator transaction binding the contract method 0x06e99fef.
//
// Solidity: function sendEth() returns()
func (_EthSender *EthSenderTransactor) SendEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthSender.contract.Transact(opts, "sendEth")
}

// SendEth is a paid mutator transaction binding the contract method 0x06e99fef.
//
// Solidity: function sendEth() returns()
func (_EthSender *EthSenderSession) SendEth() (*types.Transaction, error) {
	return _EthSender.Contract.SendEth(&_EthSender.TransactOpts)
}

// SendEth is a paid mutator transaction binding the contract method 0x06e99fef.
//
// Solidity: function sendEth() returns()
func (_EthSender *EthSenderTransactorSession) SendEth() (*types.Transaction, error) {
	return _EthSender.Contract.SendEth(&_EthSender.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_EthSender *EthSenderTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _EthSender.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_EthSender *EthSenderSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _EthSender.Contract.Withdraw(&_EthSender.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_EthSender *EthSenderTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _EthSender.Contract.Withdraw(&_EthSender.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EthSender *EthSenderTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthSender.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EthSender *EthSenderSession) Receive() (*types.Transaction, error) {
	return _EthSender.Contract.Receive(&_EthSender.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EthSender *EthSenderTransactorSession) Receive() (*types.Transaction, error) {
	return _EthSender.Contract.Receive(&_EthSender.TransactOpts)
}

// EthSenderEthSentIterator is returned from FilterEthSent and is used to iterate over the raw logs and unpacked data for EthSent events raised by the EthSender contract.
type EthSenderEthSentIterator struct {
	Event *EthSenderEthSent // Event containing the contract specifics and raw log

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
func (it *EthSenderEthSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthSenderEthSent)
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
		it.Event = new(EthSenderEthSent)
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
func (it *EthSenderEthSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthSenderEthSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthSenderEthSent represents a EthSent event raised by the EthSender contract.
type EthSenderEthSent struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEthSent is a free log retrieval operation binding the contract event 0x78f5cdad99320ec2ba57132d7dffb1d125775c823239e60ff5e9300fd4ac898c.
//
// Solidity: event EthSent(address indexed recipient, uint256 amount)
func (_EthSender *EthSenderFilterer) FilterEthSent(opts *bind.FilterOpts, recipient []common.Address) (*EthSenderEthSentIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _EthSender.contract.FilterLogs(opts, "EthSent", recipientRule)
	if err != nil {
		return nil, err
	}
	return &EthSenderEthSentIterator{contract: _EthSender.contract, event: "EthSent", logs: logs, sub: sub}, nil
}

// WatchEthSent is a free log subscription operation binding the contract event 0x78f5cdad99320ec2ba57132d7dffb1d125775c823239e60ff5e9300fd4ac898c.
//
// Solidity: event EthSent(address indexed recipient, uint256 amount)
func (_EthSender *EthSenderFilterer) WatchEthSent(opts *bind.WatchOpts, sink chan<- *EthSenderEthSent, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _EthSender.contract.WatchLogs(opts, "EthSent", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthSenderEthSent)
				if err := _EthSender.contract.UnpackLog(event, "EthSent", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
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

// ParseEthSent is a log parse operation binding the contract event 0x78f5cdad99320ec2ba57132d7dffb1d125775c823239e60ff5e9300fd4ac898c.
//
// Solidity: event EthSent(address indexed recipient, uint256 amount)
func (_EthSender *EthSenderFilterer) ParseEthSent(log types.Log) (*EthSenderEthSent, error) {
	event := new(EthSenderEthSent)
	if err := _EthSender.contract.UnpackLog(event, "EthSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthSenderWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the EthSender contract.
type EthSenderWithdrawnIterator struct {
	Event *EthSenderWithdrawn // Event containing the contract specifics and raw log

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
func (it *EthSenderWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthSenderWithdrawn)
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
		it.Event = new(EthSenderWithdrawn)
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
func (it *EthSenderWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthSenderWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthSenderWithdrawn represents a Withdrawn event raised by the EthSender contract.
type EthSenderWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_EthSender *EthSenderFilterer) FilterWithdrawn(opts *bind.FilterOpts, to []common.Address) (*EthSenderWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EthSender.contract.FilterLogs(opts, "Withdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &EthSenderWithdrawnIterator{contract: _EthSender.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_EthSender *EthSenderFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *EthSenderWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EthSender.contract.WatchLogs(opts, "Withdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthSenderWithdrawn)
				if err := _EthSender.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_EthSender *EthSenderFilterer) ParseWithdrawn(log types.Log) (*EthSenderWithdrawn, error) {
	event := new(EthSenderWithdrawn)
	if err := _EthSender.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gasestimator

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/0xsequence/ethkit/go-ethereum"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi/bind"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/ethkit/go-ethereum/event"
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
)

// GasEstimatorMetaData contains all meta data concerning the GasEstimator contract.
var GasEstimatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"estimate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610208806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80630eb34cd314610030575b600080fd5b6100bd6004803603604081101561004657600080fd5b73ffffffffffffffffffffffffffffffffffffffff823516919081019060408101602082013564010000000081111561007e57600080fd5b82018360208201111561009057600080fd5b803590602001918460018302840111640100000000831117156100b257600080fd5b509092509050610145565b60405180841515815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b838110156101085781810151838201526020016100f0565b50505050905090810190601f1680156101355780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b600060606000805a90508673ffffffffffffffffffffffffffffffffffffffff168686604051808383808284376040519201945060009350909150508083038183865af19150503d80600081146101b8576040519150601f19603f3d011682016040523d82523d6000602084013e6101bd565b606091505b5090945092505a81039150509350935093905056fea26469706673582212202294382602e928b6c962b89f8f706713a73b32b446a191ebcfa9b77f7004d19064736f6c63430007060033",
}

// GasEstimatorABI is the input ABI used to generate the binding from.
// Deprecated: Use GasEstimatorMetaData.ABI instead.
var GasEstimatorABI = GasEstimatorMetaData.ABI

// GasEstimatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GasEstimatorMetaData.Bin instead.
const GasEstimatorBin = GasEstimatorMetaData.Bin

// DeployGasEstimator deploys a new Ethereum contract, binding an instance of GasEstimator to it.
func DeployGasEstimator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GasEstimator, error) {
	parsed, err := GasEstimatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GasEstimatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GasEstimator{GasEstimatorCaller: GasEstimatorCaller{contract: contract}, GasEstimatorTransactor: GasEstimatorTransactor{contract: contract}, GasEstimatorFilterer: GasEstimatorFilterer{contract: contract}}, nil
}

// GasEstimatorDeployedBin is the resulting bytecode of the created contract
const GasEstimatorDeployedBin = "0x608060405234801561001057600080fd5b506004361061002b5760003560e01c80630eb34cd314610030575b600080fd5b6100bd6004803603604081101561004657600080fd5b73ffffffffffffffffffffffffffffffffffffffff823516919081019060408101602082013564010000000081111561007e57600080fd5b82018360208201111561009057600080fd5b803590602001918460018302840111640100000000831117156100b257600080fd5b509092509050610145565b60405180841515815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b838110156101085781810151838201526020016100f0565b50505050905090810190601f1680156101355780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b600060606000805a90508673ffffffffffffffffffffffffffffffffffffffff168686604051808383808284376040519201945060009350909150508083038183865af19150503d80600081146101b8576040519150601f19603f3d011682016040523d82523d6000602084013e6101bd565b606091505b5090945092505a81039150509350935093905056fea26469706673582212202294382602e928b6c962b89f8f706713a73b32b446a191ebcfa9b77f7004d19064736f6c63430007060033"

// GasEstimator is an auto generated Go binding around an Ethereum contract.
type GasEstimator struct {
	GasEstimatorCaller     // Read-only binding to the contract
	GasEstimatorTransactor // Write-only binding to the contract
	GasEstimatorFilterer   // Log filterer for contract events
}

// GasEstimatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type GasEstimatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasEstimatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GasEstimatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasEstimatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GasEstimatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasEstimatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GasEstimatorSession struct {
	Contract     *GasEstimator     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GasEstimatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GasEstimatorCallerSession struct {
	Contract *GasEstimatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// GasEstimatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GasEstimatorTransactorSession struct {
	Contract     *GasEstimatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GasEstimatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type GasEstimatorRaw struct {
	Contract *GasEstimator // Generic contract binding to access the raw methods on
}

// GasEstimatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GasEstimatorCallerRaw struct {
	Contract *GasEstimatorCaller // Generic read-only contract binding to access the raw methods on
}

// GasEstimatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GasEstimatorTransactorRaw struct {
	Contract *GasEstimatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGasEstimator creates a new instance of GasEstimator, bound to a specific deployed contract.
func NewGasEstimator(address common.Address, backend bind.ContractBackend) (*GasEstimator, error) {
	contract, err := bindGasEstimator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GasEstimator{GasEstimatorCaller: GasEstimatorCaller{contract: contract}, GasEstimatorTransactor: GasEstimatorTransactor{contract: contract}, GasEstimatorFilterer: GasEstimatorFilterer{contract: contract}}, nil
}

// NewGasEstimatorCaller creates a new read-only instance of GasEstimator, bound to a specific deployed contract.
func NewGasEstimatorCaller(address common.Address, caller bind.ContractCaller) (*GasEstimatorCaller, error) {
	contract, err := bindGasEstimator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GasEstimatorCaller{contract: contract}, nil
}

// NewGasEstimatorTransactor creates a new write-only instance of GasEstimator, bound to a specific deployed contract.
func NewGasEstimatorTransactor(address common.Address, transactor bind.ContractTransactor) (*GasEstimatorTransactor, error) {
	contract, err := bindGasEstimator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GasEstimatorTransactor{contract: contract}, nil
}

// NewGasEstimatorFilterer creates a new log filterer instance of GasEstimator, bound to a specific deployed contract.
func NewGasEstimatorFilterer(address common.Address, filterer bind.ContractFilterer) (*GasEstimatorFilterer, error) {
	contract, err := bindGasEstimator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GasEstimatorFilterer{contract: contract}, nil
}

// bindGasEstimator binds a generic wrapper to an already deployed contract.
func bindGasEstimator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GasEstimatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasEstimator *GasEstimatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasEstimator.Contract.GasEstimatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasEstimator *GasEstimatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasEstimator.Contract.GasEstimatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasEstimator *GasEstimatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasEstimator.Contract.GasEstimatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasEstimator *GasEstimatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasEstimator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasEstimator *GasEstimatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasEstimator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasEstimator *GasEstimatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasEstimator.Contract.contract.Transact(opts, method, params...)
}

// Estimate is a paid mutator transaction binding the contract method 0x0eb34cd3.
//
// Solidity: function estimate(address _to, bytes _data) returns(bool success, bytes result, uint256 gas)
func (_GasEstimator *GasEstimatorTransactor) Estimate(opts *bind.TransactOpts, _to common.Address, _data []byte) (*types.Transaction, error) {
	return _GasEstimator.contract.Transact(opts, "estimate", _to, _data)
}

// Estimate is a paid mutator transaction binding the contract method 0x0eb34cd3.
//
// Solidity: function estimate(address _to, bytes _data) returns(bool success, bytes result, uint256 gas)
func (_GasEstimator *GasEstimatorSession) Estimate(_to common.Address, _data []byte) (*types.Transaction, error) {
	return _GasEstimator.Contract.Estimate(&_GasEstimator.TransactOpts, _to, _data)
}

// Estimate is a paid mutator transaction binding the contract method 0x0eb34cd3.
//
// Solidity: function estimate(address _to, bytes _data) returns(bool success, bytes result, uint256 gas)
func (_GasEstimator *GasEstimatorTransactorSession) Estimate(_to common.Address, _data []byte) (*types.Transaction, error) {
	return _GasEstimator.Contract.Estimate(&_GasEstimator.TransactOpts, _to, _data)
}

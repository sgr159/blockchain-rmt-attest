// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OperContract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// OperContABI is the input ABI used to generate the binding from.
const OperContABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"usr\",\"type\":\"address\"},{\"name\":\"username\",\"type\":\"string\"},{\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"addAuthorizedUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"binName\",\"type\":\"string\"},{\"name\":\"measurement\",\"type\":\"uint256\"}],\"name\":\"verifyBinary\",\"outputs\":[{\"name\":\"res\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"binName\",\"type\":\"string\"},{\"name\":\"measurement\",\"type\":\"uint256\"}],\"name\":\"addBinary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"username\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OperCont is an auto generated Go binding around an Ethereum contract.
type OperCont struct {
	OperContCaller     // Read-only binding to the contract
	OperContTransactor // Write-only binding to the contract
}

// OperContCaller is an auto generated read-only Go binding around an Ethereum contract.
type OperContCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperContTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OperContTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperContSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OperContSession struct {
	Contract     *OperCont         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OperContCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OperContCallerSession struct {
	Contract *OperContCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// OperContTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OperContTransactorSession struct {
	Contract     *OperContTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OperContRaw is an auto generated low-level Go binding around an Ethereum contract.
type OperContRaw struct {
	Contract *OperCont // Generic contract binding to access the raw methods on
}

// OperContCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OperContCallerRaw struct {
	Contract *OperContCaller // Generic read-only contract binding to access the raw methods on
}

// OperContTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OperContTransactorRaw struct {
	Contract *OperContTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOperCont creates a new instance of OperCont, bound to a specific deployed contract.
func NewOperCont(address common.Address, backend bind.ContractBackend) (*OperCont, error) {
	contract, err := bindOperCont(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OperCont{OperContCaller: OperContCaller{contract: contract}, OperContTransactor: OperContTransactor{contract: contract}}, nil
}

// NewOperContCaller creates a new read-only instance of OperCont, bound to a specific deployed contract.
func NewOperContCaller(address common.Address, caller bind.ContractCaller) (*OperContCaller, error) {
	contract, err := bindOperCont(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &OperContCaller{contract: contract}, nil
}

// NewOperContTransactor creates a new write-only instance of OperCont, bound to a specific deployed contract.
func NewOperContTransactor(address common.Address, transactor bind.ContractTransactor) (*OperContTransactor, error) {
	contract, err := bindOperCont(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &OperContTransactor{contract: contract}, nil
}

// bindOperCont binds a generic wrapper to an already deployed contract.
func bindOperCont(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OperContABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OperCont *OperContRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OperCont.Contract.OperContCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OperCont *OperContRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperCont.Contract.OperContTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OperCont *OperContRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OperCont.Contract.OperContTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OperCont *OperContCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OperCont.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OperCont *OperContTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperCont.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OperCont *OperContTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OperCont.Contract.contract.Transact(opts, method, params...)
}

// VerifyBinary is a free data retrieval call binding the contract method 0xaf44c218.
//
// Solidity: function verifyBinary(binName string, measurement uint256) constant returns(res bool)
func (_OperCont *OperContCaller) VerifyBinary(opts *bind.CallOpts, binName string, measurement *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OperCont.contract.Call(opts, out, "verifyBinary", binName, measurement)
	return *ret0, err
}

// VerifyBinary is a free data retrieval call binding the contract method 0xaf44c218.
//
// Solidity: function verifyBinary(binName string, measurement uint256) constant returns(res bool)
func (_OperCont *OperContSession) VerifyBinary(binName string, measurement *big.Int) (bool, error) {
	return _OperCont.Contract.VerifyBinary(&_OperCont.CallOpts, binName, measurement)
}

// VerifyBinary is a free data retrieval call binding the contract method 0xaf44c218.
//
// Solidity: function verifyBinary(binName string, measurement uint256) constant returns(res bool)
func (_OperCont *OperContCallerSession) VerifyBinary(binName string, measurement *big.Int) (bool, error) {
	return _OperCont.Contract.VerifyBinary(&_OperCont.CallOpts, binName, measurement)
}

// AddAuthorizedUser is a paid mutator transaction binding the contract method 0x90fbdd78.
//
// Solidity: function addAuthorizedUser(usr address, username string, role uint8) returns()
func (_OperCont *OperContTransactor) AddAuthorizedUser(opts *bind.TransactOpts, usr common.Address, username string, role uint8) (*types.Transaction, error) {
	return _OperCont.contract.Transact(opts, "addAuthorizedUser", usr, username, role)
}

// AddAuthorizedUser is a paid mutator transaction binding the contract method 0x90fbdd78.
//
// Solidity: function addAuthorizedUser(usr address, username string, role uint8) returns()
func (_OperCont *OperContSession) AddAuthorizedUser(usr common.Address, username string, role uint8) (*types.Transaction, error) {
	return _OperCont.Contract.AddAuthorizedUser(&_OperCont.TransactOpts, usr, username, role)
}

// AddAuthorizedUser is a paid mutator transaction binding the contract method 0x90fbdd78.
//
// Solidity: function addAuthorizedUser(usr address, username string, role uint8) returns()
func (_OperCont *OperContTransactorSession) AddAuthorizedUser(usr common.Address, username string, role uint8) (*types.Transaction, error) {
	return _OperCont.Contract.AddAuthorizedUser(&_OperCont.TransactOpts, usr, username, role)
}

// AddBinary is a paid mutator transaction binding the contract method 0xe42422af.
//
// Solidity: function addBinary(binName string, measurement uint256) returns()
func (_OperCont *OperContTransactor) AddBinary(opts *bind.TransactOpts, binName string, measurement *big.Int) (*types.Transaction, error) {
	return _OperCont.contract.Transact(opts, "addBinary", binName, measurement)
}

// AddBinary is a paid mutator transaction binding the contract method 0xe42422af.
//
// Solidity: function addBinary(binName string, measurement uint256) returns()
func (_OperCont *OperContSession) AddBinary(binName string, measurement *big.Int) (*types.Transaction, error) {
	return _OperCont.Contract.AddBinary(&_OperCont.TransactOpts, binName, measurement)
}

// AddBinary is a paid mutator transaction binding the contract method 0xe42422af.
//
// Solidity: function addBinary(binName string, measurement uint256) returns()
func (_OperCont *OperContTransactorSession) AddBinary(binName string, measurement *big.Int) (*types.Transaction, error) {
	return _OperCont.Contract.AddBinary(&_OperCont.TransactOpts, binName, measurement)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_OperCont *OperContTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperCont.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_OperCont *OperContSession) Kill() (*types.Transaction, error) {
	return _OperCont.Contract.Kill(&_OperCont.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_OperCont *OperContTransactorSession) Kill() (*types.Transaction, error) {
	return _OperCont.Contract.Kill(&_OperCont.TransactOpts)
}

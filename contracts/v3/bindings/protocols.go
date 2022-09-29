// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package swivel

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ProtocolsMetaData contains all meta data concerning the Protocols contract.
var ProtocolsMetaData = &bind.MetaData{
	ABI: "[]",
}

// ProtocolsABI is the input ABI used to generate the binding from.
// Deprecated: Use ProtocolsMetaData.ABI instead.
var ProtocolsABI = ProtocolsMetaData.ABI

// Protocols is an auto generated Go binding around an Ethereum contract.
type Protocols struct {
	ProtocolsCaller     // Read-only binding to the contract
	ProtocolsTransactor // Write-only binding to the contract
	ProtocolsFilterer   // Log filterer for contract events
}

// ProtocolsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProtocolsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProtocolsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProtocolsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProtocolsSession struct {
	Contract     *Protocols        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProtocolsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProtocolsCallerSession struct {
	Contract *ProtocolsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ProtocolsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProtocolsTransactorSession struct {
	Contract     *ProtocolsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ProtocolsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProtocolsRaw struct {
	Contract *Protocols // Generic contract binding to access the raw methods on
}

// ProtocolsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProtocolsCallerRaw struct {
	Contract *ProtocolsCaller // Generic read-only contract binding to access the raw methods on
}

// ProtocolsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProtocolsTransactorRaw struct {
	Contract *ProtocolsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProtocols creates a new instance of Protocols, bound to a specific deployed contract.
func NewProtocols(address common.Address, backend bind.ContractBackend) (*Protocols, error) {
	contract, err := bindProtocols(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Protocols{ProtocolsCaller: ProtocolsCaller{contract: contract}, ProtocolsTransactor: ProtocolsTransactor{contract: contract}, ProtocolsFilterer: ProtocolsFilterer{contract: contract}}, nil
}

// NewProtocolsCaller creates a new read-only instance of Protocols, bound to a specific deployed contract.
func NewProtocolsCaller(address common.Address, caller bind.ContractCaller) (*ProtocolsCaller, error) {
	contract, err := bindProtocols(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolsCaller{contract: contract}, nil
}

// NewProtocolsTransactor creates a new write-only instance of Protocols, bound to a specific deployed contract.
func NewProtocolsTransactor(address common.Address, transactor bind.ContractTransactor) (*ProtocolsTransactor, error) {
	contract, err := bindProtocols(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolsTransactor{contract: contract}, nil
}

// NewProtocolsFilterer creates a new log filterer instance of Protocols, bound to a specific deployed contract.
func NewProtocolsFilterer(address common.Address, filterer bind.ContractFilterer) (*ProtocolsFilterer, error) {
	contract, err := bindProtocols(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProtocolsFilterer{contract: contract}, nil
}

// bindProtocols binds a generic wrapper to an already deployed contract.
func bindProtocols(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProtocolsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Protocols *ProtocolsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Protocols.Contract.ProtocolsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Protocols *ProtocolsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Protocols.Contract.ProtocolsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Protocols *ProtocolsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Protocols.Contract.ProtocolsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Protocols *ProtocolsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Protocols.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Protocols *ProtocolsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Protocols.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Protocols *ProtocolsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Protocols.Contract.contract.Transact(opts, method, params...)
}

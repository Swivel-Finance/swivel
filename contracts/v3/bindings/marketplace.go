// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package marketplace

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

// MarketPlaceMetaData contains all meta data concerning the MarketPlace contract.
var MarketPlaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Exception\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"protocol\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"zcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vaultTracker\",\"type\":\"address\"}],\"name\":\"Create\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"protocol\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"zcTarget\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nTarget\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CustodialExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"protocol\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"zcTarget\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nTarget\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CustodialInitiate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"protocol\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maturityRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"matured\",\"type\":\"uint256\"}],\"name\":\"Mature\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"protocol\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"P2pVaultExchange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"protocol\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"P2pZcTokenExchange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"protocol\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RedeemVaultInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"protocol\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RedeemZcToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"SetAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"protocol\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferVaultNotional\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"authRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"burnZcTokenRemovingNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"cTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"}],\"name\":\"createMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"z\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"n\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"custodialExit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"z\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"n\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"custodialInitiate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"}],\"name\":\"exchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"cTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zcToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vaultTracker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturityRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"matureMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"mintZcTokenAddingNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"p2pVaultExchange\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"p2pZcTokenExchange\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"pause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"rates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"redeemVaultInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"redeemZcToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"}],\"name\":\"setSwivel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swivel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferVaultNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferVaultNotionalFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MarketPlaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketPlaceMetaData.ABI instead.
var MarketPlaceABI = MarketPlaceMetaData.ABI

// MarketPlace is an auto generated Go binding around an Ethereum contract.
type MarketPlace struct {
	MarketPlaceCaller     // Read-only binding to the contract
	MarketPlaceTransactor // Write-only binding to the contract
	MarketPlaceFilterer   // Log filterer for contract events
}

// MarketPlaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketPlaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketPlaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketPlaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketPlaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketPlaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketPlaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketPlaceSession struct {
	Contract     *MarketPlace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketPlaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketPlaceCallerSession struct {
	Contract *MarketPlaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketPlaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketPlaceTransactorSession struct {
	Contract     *MarketPlaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketPlaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketPlaceRaw struct {
	Contract *MarketPlace // Generic contract binding to access the raw methods on
}

// MarketPlaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketPlaceCallerRaw struct {
	Contract *MarketPlaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketPlaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketPlaceTransactorRaw struct {
	Contract *MarketPlaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketPlace creates a new instance of MarketPlace, bound to a specific deployed contract.
func NewMarketPlace(address common.Address, backend bind.ContractBackend) (*MarketPlace, error) {
	contract, err := bindMarketPlace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarketPlace{MarketPlaceCaller: MarketPlaceCaller{contract: contract}, MarketPlaceTransactor: MarketPlaceTransactor{contract: contract}, MarketPlaceFilterer: MarketPlaceFilterer{contract: contract}}, nil
}

// NewMarketPlaceCaller creates a new read-only instance of MarketPlace, bound to a specific deployed contract.
func NewMarketPlaceCaller(address common.Address, caller bind.ContractCaller) (*MarketPlaceCaller, error) {
	contract, err := bindMarketPlace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceCaller{contract: contract}, nil
}

// NewMarketPlaceTransactor creates a new write-only instance of MarketPlace, bound to a specific deployed contract.
func NewMarketPlaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketPlaceTransactor, error) {
	contract, err := bindMarketPlace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceTransactor{contract: contract}, nil
}

// NewMarketPlaceFilterer creates a new log filterer instance of MarketPlace, bound to a specific deployed contract.
func NewMarketPlaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketPlaceFilterer, error) {
	contract, err := bindMarketPlace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceFilterer{contract: contract}, nil
}

// bindMarketPlace binds a generic wrapper to an already deployed contract.
func bindMarketPlace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketPlaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketPlace *MarketPlaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketPlace.Contract.MarketPlaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketPlace *MarketPlaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketPlace.Contract.MarketPlaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketPlace *MarketPlaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketPlace.Contract.MarketPlaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketPlace *MarketPlaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketPlace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketPlace *MarketPlaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketPlace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketPlace *MarketPlaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketPlace.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_MarketPlace *MarketPlaceCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_MarketPlace *MarketPlaceSession) Admin() (common.Address, error) {
	return _MarketPlace.Contract.Admin(&_MarketPlace.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_MarketPlace *MarketPlaceCallerSession) Admin() (common.Address, error) {
	return _MarketPlace.Contract.Admin(&_MarketPlace.CallOpts)
}

// CTokenAddress is a free data retrieval call binding the contract method 0x35bdafab.
//
// Solidity: function cTokenAddress(uint8 p, address u, uint256 m) view returns(address)
func (_MarketPlace *MarketPlaceCaller) CTokenAddress(opts *bind.CallOpts, p uint8, u common.Address, m *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "cTokenAddress", p, u, m)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CTokenAddress is a free data retrieval call binding the contract method 0x35bdafab.
//
// Solidity: function cTokenAddress(uint8 p, address u, uint256 m) view returns(address)
func (_MarketPlace *MarketPlaceSession) CTokenAddress(p uint8, u common.Address, m *big.Int) (common.Address, error) {
	return _MarketPlace.Contract.CTokenAddress(&_MarketPlace.CallOpts, p, u, m)
}

// CTokenAddress is a free data retrieval call binding the contract method 0x35bdafab.
//
// Solidity: function cTokenAddress(uint8 p, address u, uint256 m) view returns(address)
func (_MarketPlace *MarketPlaceCallerSession) CTokenAddress(p uint8, u common.Address, m *big.Int) (common.Address, error) {
	return _MarketPlace.Contract.CTokenAddress(&_MarketPlace.CallOpts, p, u, m)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_MarketPlace *MarketPlaceCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_MarketPlace *MarketPlaceSession) Creator() (common.Address, error) {
	return _MarketPlace.Contract.Creator(&_MarketPlace.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_MarketPlace *MarketPlaceCallerSession) Creator() (common.Address, error) {
	return _MarketPlace.Contract.Creator(&_MarketPlace.CallOpts)
}

// Markets is a free data retrieval call binding the contract method 0x305a21bf.
//
// Solidity: function markets(uint8 , address , uint256 ) view returns(address cTokenAddr, address zcToken, address vaultTracker, uint256 maturityRate)
func (_MarketPlace *MarketPlaceCaller) Markets(opts *bind.CallOpts, arg0 uint8, arg1 common.Address, arg2 *big.Int) (struct {
	CTokenAddr   common.Address
	ZcToken      common.Address
	VaultTracker common.Address
	MaturityRate *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "markets", arg0, arg1, arg2)

	outstruct := new(struct {
		CTokenAddr   common.Address
		ZcToken      common.Address
		VaultTracker common.Address
		MaturityRate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CTokenAddr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ZcToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.VaultTracker = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.MaturityRate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Markets is a free data retrieval call binding the contract method 0x305a21bf.
//
// Solidity: function markets(uint8 , address , uint256 ) view returns(address cTokenAddr, address zcToken, address vaultTracker, uint256 maturityRate)
func (_MarketPlace *MarketPlaceSession) Markets(arg0 uint8, arg1 common.Address, arg2 *big.Int) (struct {
	CTokenAddr   common.Address
	ZcToken      common.Address
	VaultTracker common.Address
	MaturityRate *big.Int
}, error) {
	return _MarketPlace.Contract.Markets(&_MarketPlace.CallOpts, arg0, arg1, arg2)
}

// Markets is a free data retrieval call binding the contract method 0x305a21bf.
//
// Solidity: function markets(uint8 , address , uint256 ) view returns(address cTokenAddr, address zcToken, address vaultTracker, uint256 maturityRate)
func (_MarketPlace *MarketPlaceCallerSession) Markets(arg0 uint8, arg1 common.Address, arg2 *big.Int) (struct {
	CTokenAddr   common.Address
	ZcToken      common.Address
	VaultTracker common.Address
	MaturityRate *big.Int
}, error) {
	return _MarketPlace.Contract.Markets(&_MarketPlace.CallOpts, arg0, arg1, arg2)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 ) view returns(bool)
func (_MarketPlace *MarketPlaceCaller) Paused(opts *bind.CallOpts, arg0 uint8) (bool, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "paused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 ) view returns(bool)
func (_MarketPlace *MarketPlaceSession) Paused(arg0 uint8) (bool, error) {
	return _MarketPlace.Contract.Paused(&_MarketPlace.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 ) view returns(bool)
func (_MarketPlace *MarketPlaceCallerSession) Paused(arg0 uint8) (bool, error) {
	return _MarketPlace.Contract.Paused(&_MarketPlace.CallOpts, arg0)
}

// Swivel is a free data retrieval call binding the contract method 0x012b264a.
//
// Solidity: function swivel() view returns(address)
func (_MarketPlace *MarketPlaceCaller) Swivel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "swivel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Swivel is a free data retrieval call binding the contract method 0x012b264a.
//
// Solidity: function swivel() view returns(address)
func (_MarketPlace *MarketPlaceSession) Swivel() (common.Address, error) {
	return _MarketPlace.Contract.Swivel(&_MarketPlace.CallOpts)
}

// Swivel is a free data retrieval call binding the contract method 0x012b264a.
//
// Solidity: function swivel() view returns(address)
func (_MarketPlace *MarketPlaceCallerSession) Swivel() (common.Address, error) {
	return _MarketPlace.Contract.Swivel(&_MarketPlace.CallOpts)
}

// AuthRedeem is a paid mutator transaction binding the contract method 0x52bc9430.
//
// Solidity: function authRedeem(uint8 p, address u, uint256 m, address f, address t, uint256 a) returns(uint256)
func (_MarketPlace *MarketPlaceTransactor) AuthRedeem(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "authRedeem", p, u, m, f, t, a)
}

// AuthRedeem is a paid mutator transaction binding the contract method 0x52bc9430.
//
// Solidity: function authRedeem(uint8 p, address u, uint256 m, address f, address t, uint256 a) returns(uint256)
func (_MarketPlace *MarketPlaceSession) AuthRedeem(p uint8, u common.Address, m *big.Int, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.AuthRedeem(&_MarketPlace.TransactOpts, p, u, m, f, t, a)
}

// AuthRedeem is a paid mutator transaction binding the contract method 0x52bc9430.
//
// Solidity: function authRedeem(uint8 p, address u, uint256 m, address f, address t, uint256 a) returns(uint256)
func (_MarketPlace *MarketPlaceTransactorSession) AuthRedeem(p uint8, u common.Address, m *big.Int, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.AuthRedeem(&_MarketPlace.TransactOpts, p, u, m, f, t, a)
}

// BurnZcTokenRemovingNotional is a paid mutator transaction binding the contract method 0x87e157c1.
//
// Solidity: function burnZcTokenRemovingNotional(uint8 p, address u, uint256 m, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) BurnZcTokenRemovingNotional(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "burnZcTokenRemovingNotional", p, u, m, t, a)
}

// BurnZcTokenRemovingNotional is a paid mutator transaction binding the contract method 0x87e157c1.
//
// Solidity: function burnZcTokenRemovingNotional(uint8 p, address u, uint256 m, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) BurnZcTokenRemovingNotional(p uint8, u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.BurnZcTokenRemovingNotional(&_MarketPlace.TransactOpts, p, u, m, t, a)
}

// BurnZcTokenRemovingNotional is a paid mutator transaction binding the contract method 0x87e157c1.
//
// Solidity: function burnZcTokenRemovingNotional(uint8 p, address u, uint256 m, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) BurnZcTokenRemovingNotional(p uint8, u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.BurnZcTokenRemovingNotional(&_MarketPlace.TransactOpts, p, u, m, t, a)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x174d2548.
//
// Solidity: function createMarket(uint8 p, uint256 m, address c, string n, string s) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) CreateMarket(opts *bind.TransactOpts, p uint8, m *big.Int, c common.Address, n string, s string) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "createMarket", p, m, c, n, s)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x174d2548.
//
// Solidity: function createMarket(uint8 p, uint256 m, address c, string n, string s) returns(bool)
func (_MarketPlace *MarketPlaceSession) CreateMarket(p uint8, m *big.Int, c common.Address, n string, s string) (*types.Transaction, error) {
	return _MarketPlace.Contract.CreateMarket(&_MarketPlace.TransactOpts, p, m, c, n, s)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x174d2548.
//
// Solidity: function createMarket(uint8 p, uint256 m, address c, string n, string s) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) CreateMarket(p uint8, m *big.Int, c common.Address, n string, s string) (*types.Transaction, error) {
	return _MarketPlace.Contract.CreateMarket(&_MarketPlace.TransactOpts, p, m, c, n, s)
}

// CustodialExit is a paid mutator transaction binding the contract method 0x0f0016b6.
//
// Solidity: function custodialExit(uint8 p, address u, uint256 m, address z, address n, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) CustodialExit(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, z common.Address, n common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "custodialExit", p, u, m, z, n, a)
}

// CustodialExit is a paid mutator transaction binding the contract method 0x0f0016b6.
//
// Solidity: function custodialExit(uint8 p, address u, uint256 m, address z, address n, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) CustodialExit(p uint8, u common.Address, m *big.Int, z common.Address, n common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CustodialExit(&_MarketPlace.TransactOpts, p, u, m, z, n, a)
}

// CustodialExit is a paid mutator transaction binding the contract method 0x0f0016b6.
//
// Solidity: function custodialExit(uint8 p, address u, uint256 m, address z, address n, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) CustodialExit(p uint8, u common.Address, m *big.Int, z common.Address, n common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CustodialExit(&_MarketPlace.TransactOpts, p, u, m, z, n, a)
}

// CustodialInitiate is a paid mutator transaction binding the contract method 0xc06760c7.
//
// Solidity: function custodialInitiate(uint8 p, address u, uint256 m, address z, address n, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) CustodialInitiate(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, z common.Address, n common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "custodialInitiate", p, u, m, z, n, a)
}

// CustodialInitiate is a paid mutator transaction binding the contract method 0xc06760c7.
//
// Solidity: function custodialInitiate(uint8 p, address u, uint256 m, address z, address n, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) CustodialInitiate(p uint8, u common.Address, m *big.Int, z common.Address, n common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CustodialInitiate(&_MarketPlace.TransactOpts, p, u, m, z, n, a)
}

// CustodialInitiate is a paid mutator transaction binding the contract method 0xc06760c7.
//
// Solidity: function custodialInitiate(uint8 p, address u, uint256 m, address z, address n, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) CustodialInitiate(p uint8, u common.Address, m *big.Int, z common.Address, n common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CustodialInitiate(&_MarketPlace.TransactOpts, p, u, m, z, n, a)
}

// ExchangeRate is a paid mutator transaction binding the contract method 0x5755d763.
//
// Solidity: function exchangeRate(uint8 p, address c) returns(uint256)
func (_MarketPlace *MarketPlaceTransactor) ExchangeRate(opts *bind.TransactOpts, p uint8, c common.Address) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "exchangeRate", p, c)
}

// ExchangeRate is a paid mutator transaction binding the contract method 0x5755d763.
//
// Solidity: function exchangeRate(uint8 p, address c) returns(uint256)
func (_MarketPlace *MarketPlaceSession) ExchangeRate(p uint8, c common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.ExchangeRate(&_MarketPlace.TransactOpts, p, c)
}

// ExchangeRate is a paid mutator transaction binding the contract method 0x5755d763.
//
// Solidity: function exchangeRate(uint8 p, address c) returns(uint256)
func (_MarketPlace *MarketPlaceTransactorSession) ExchangeRate(p uint8, c common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.ExchangeRate(&_MarketPlace.TransactOpts, p, c)
}

// MatureMarket is a paid mutator transaction binding the contract method 0x872e9f6c.
//
// Solidity: function matureMarket(uint8 p, address u, uint256 m) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) MatureMarket(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "matureMarket", p, u, m)
}

// MatureMarket is a paid mutator transaction binding the contract method 0x872e9f6c.
//
// Solidity: function matureMarket(uint8 p, address u, uint256 m) returns(bool)
func (_MarketPlace *MarketPlaceSession) MatureMarket(p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.MatureMarket(&_MarketPlace.TransactOpts, p, u, m)
}

// MatureMarket is a paid mutator transaction binding the contract method 0x872e9f6c.
//
// Solidity: function matureMarket(uint8 p, address u, uint256 m) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) MatureMarket(p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.MatureMarket(&_MarketPlace.TransactOpts, p, u, m)
}

// MintZcTokenAddingNotional is a paid mutator transaction binding the contract method 0x01cc6448.
//
// Solidity: function mintZcTokenAddingNotional(uint8 p, address u, uint256 m, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) MintZcTokenAddingNotional(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "mintZcTokenAddingNotional", p, u, m, t, a)
}

// MintZcTokenAddingNotional is a paid mutator transaction binding the contract method 0x01cc6448.
//
// Solidity: function mintZcTokenAddingNotional(uint8 p, address u, uint256 m, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) MintZcTokenAddingNotional(p uint8, u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.MintZcTokenAddingNotional(&_MarketPlace.TransactOpts, p, u, m, t, a)
}

// MintZcTokenAddingNotional is a paid mutator transaction binding the contract method 0x01cc6448.
//
// Solidity: function mintZcTokenAddingNotional(uint8 p, address u, uint256 m, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) MintZcTokenAddingNotional(p uint8, u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.MintZcTokenAddingNotional(&_MarketPlace.TransactOpts, p, u, m, t, a)
}

// P2pVaultExchange is a paid mutator transaction binding the contract method 0x15042ddf.
//
// Solidity: function p2pVaultExchange(uint8 p, address u, uint256 m, address f, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) P2pVaultExchange(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "p2pVaultExchange", p, u, m, f, t, a)
}

// P2pVaultExchange is a paid mutator transaction binding the contract method 0x15042ddf.
//
// Solidity: function p2pVaultExchange(uint8 p, address u, uint256 m, address f, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) P2pVaultExchange(p uint8, u common.Address, m *big.Int, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.P2pVaultExchange(&_MarketPlace.TransactOpts, p, u, m, f, t, a)
}

// P2pVaultExchange is a paid mutator transaction binding the contract method 0x15042ddf.
//
// Solidity: function p2pVaultExchange(uint8 p, address u, uint256 m, address f, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) P2pVaultExchange(p uint8, u common.Address, m *big.Int, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.P2pVaultExchange(&_MarketPlace.TransactOpts, p, u, m, f, t, a)
}

// P2pZcTokenExchange is a paid mutator transaction binding the contract method 0xfcbaab2e.
//
// Solidity: function p2pZcTokenExchange(uint8 p, address u, uint256 m, address f, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) P2pZcTokenExchange(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "p2pZcTokenExchange", p, u, m, f, t, a)
}

// P2pZcTokenExchange is a paid mutator transaction binding the contract method 0xfcbaab2e.
//
// Solidity: function p2pZcTokenExchange(uint8 p, address u, uint256 m, address f, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) P2pZcTokenExchange(p uint8, u common.Address, m *big.Int, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.P2pZcTokenExchange(&_MarketPlace.TransactOpts, p, u, m, f, t, a)
}

// P2pZcTokenExchange is a paid mutator transaction binding the contract method 0xfcbaab2e.
//
// Solidity: function p2pZcTokenExchange(uint8 p, address u, uint256 m, address f, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) P2pZcTokenExchange(p uint8, u common.Address, m *big.Int, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.P2pZcTokenExchange(&_MarketPlace.TransactOpts, p, u, m, f, t, a)
}

// Pause is a paid mutator transaction binding the contract method 0xfe3ee169.
//
// Solidity: function pause(uint8 p, bool b) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) Pause(opts *bind.TransactOpts, p uint8, b bool) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "pause", p, b)
}

// Pause is a paid mutator transaction binding the contract method 0xfe3ee169.
//
// Solidity: function pause(uint8 p, bool b) returns(bool)
func (_MarketPlace *MarketPlaceSession) Pause(p uint8, b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.Pause(&_MarketPlace.TransactOpts, p, b)
}

// Pause is a paid mutator transaction binding the contract method 0xfe3ee169.
//
// Solidity: function pause(uint8 p, bool b) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) Pause(p uint8, b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.Pause(&_MarketPlace.TransactOpts, p, b)
}

// Rates is a paid mutator transaction binding the contract method 0xf7de8b1f.
//
// Solidity: function rates(uint8 p, address u, uint256 m) returns(uint256, uint256)
func (_MarketPlace *MarketPlaceTransactor) Rates(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "rates", p, u, m)
}

// Rates is a paid mutator transaction binding the contract method 0xf7de8b1f.
//
// Solidity: function rates(uint8 p, address u, uint256 m) returns(uint256, uint256)
func (_MarketPlace *MarketPlaceSession) Rates(p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.Rates(&_MarketPlace.TransactOpts, p, u, m)
}

// Rates is a paid mutator transaction binding the contract method 0xf7de8b1f.
//
// Solidity: function rates(uint8 p, address u, uint256 m) returns(uint256, uint256)
func (_MarketPlace *MarketPlaceTransactorSession) Rates(p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.Rates(&_MarketPlace.TransactOpts, p, u, m)
}

// RedeemVaultInterest is a paid mutator transaction binding the contract method 0x3a660bd8.
//
// Solidity: function redeemVaultInterest(uint8 p, address u, uint256 m, address t) returns(uint256)
func (_MarketPlace *MarketPlaceTransactor) RedeemVaultInterest(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, t common.Address) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "redeemVaultInterest", p, u, m, t)
}

// RedeemVaultInterest is a paid mutator transaction binding the contract method 0x3a660bd8.
//
// Solidity: function redeemVaultInterest(uint8 p, address u, uint256 m, address t) returns(uint256)
func (_MarketPlace *MarketPlaceSession) RedeemVaultInterest(p uint8, u common.Address, m *big.Int, t common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.RedeemVaultInterest(&_MarketPlace.TransactOpts, p, u, m, t)
}

// RedeemVaultInterest is a paid mutator transaction binding the contract method 0x3a660bd8.
//
// Solidity: function redeemVaultInterest(uint8 p, address u, uint256 m, address t) returns(uint256)
func (_MarketPlace *MarketPlaceTransactorSession) RedeemVaultInterest(p uint8, u common.Address, m *big.Int, t common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.RedeemVaultInterest(&_MarketPlace.TransactOpts, p, u, m, t)
}

// RedeemZcToken is a paid mutator transaction binding the contract method 0x9f6eddc4.
//
// Solidity: function redeemZcToken(uint8 p, address u, uint256 m, address t, uint256 a) returns(uint256)
func (_MarketPlace *MarketPlaceTransactor) RedeemZcToken(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "redeemZcToken", p, u, m, t, a)
}

// RedeemZcToken is a paid mutator transaction binding the contract method 0x9f6eddc4.
//
// Solidity: function redeemZcToken(uint8 p, address u, uint256 m, address t, uint256 a) returns(uint256)
func (_MarketPlace *MarketPlaceSession) RedeemZcToken(p uint8, u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.RedeemZcToken(&_MarketPlace.TransactOpts, p, u, m, t, a)
}

// RedeemZcToken is a paid mutator transaction binding the contract method 0x9f6eddc4.
//
// Solidity: function redeemZcToken(uint8 p, address u, uint256 m, address t, uint256 a) returns(uint256)
func (_MarketPlace *MarketPlaceTransactorSession) RedeemZcToken(p uint8, u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.RedeemZcToken(&_MarketPlace.TransactOpts, p, u, m, t, a)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) SetAdmin(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "setAdmin", a)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address a) returns(bool)
func (_MarketPlace *MarketPlaceSession) SetAdmin(a common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.SetAdmin(&_MarketPlace.TransactOpts, a)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) SetAdmin(a common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.SetAdmin(&_MarketPlace.TransactOpts, a)
}

// SetSwivel is a paid mutator transaction binding the contract method 0xb79eb926.
//
// Solidity: function setSwivel(address s) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) SetSwivel(opts *bind.TransactOpts, s common.Address) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "setSwivel", s)
}

// SetSwivel is a paid mutator transaction binding the contract method 0xb79eb926.
//
// Solidity: function setSwivel(address s) returns(bool)
func (_MarketPlace *MarketPlaceSession) SetSwivel(s common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.SetSwivel(&_MarketPlace.TransactOpts, s)
}

// SetSwivel is a paid mutator transaction binding the contract method 0xb79eb926.
//
// Solidity: function setSwivel(address s) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) SetSwivel(s common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.SetSwivel(&_MarketPlace.TransactOpts, s)
}

// TransferVaultNotional is a paid mutator transaction binding the contract method 0x7dcad278.
//
// Solidity: function transferVaultNotional(uint8 p, address u, uint256 m, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) TransferVaultNotional(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "transferVaultNotional", p, u, m, t, a)
}

// TransferVaultNotional is a paid mutator transaction binding the contract method 0x7dcad278.
//
// Solidity: function transferVaultNotional(uint8 p, address u, uint256 m, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) TransferVaultNotional(p uint8, u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferVaultNotional(&_MarketPlace.TransactOpts, p, u, m, t, a)
}

// TransferVaultNotional is a paid mutator transaction binding the contract method 0x7dcad278.
//
// Solidity: function transferVaultNotional(uint8 p, address u, uint256 m, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) TransferVaultNotional(p uint8, u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferVaultNotional(&_MarketPlace.TransactOpts, p, u, m, t, a)
}

// TransferVaultNotionalFee is a paid mutator transaction binding the contract method 0xdb850901.
//
// Solidity: function transferVaultNotionalFee(uint8 p, address u, uint256 m, address f, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) TransferVaultNotionalFee(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, f common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "transferVaultNotionalFee", p, u, m, f, a)
}

// TransferVaultNotionalFee is a paid mutator transaction binding the contract method 0xdb850901.
//
// Solidity: function transferVaultNotionalFee(uint8 p, address u, uint256 m, address f, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) TransferVaultNotionalFee(p uint8, u common.Address, m *big.Int, f common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferVaultNotionalFee(&_MarketPlace.TransactOpts, p, u, m, f, a)
}

// TransferVaultNotionalFee is a paid mutator transaction binding the contract method 0xdb850901.
//
// Solidity: function transferVaultNotionalFee(uint8 p, address u, uint256 m, address f, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) TransferVaultNotionalFee(p uint8, u common.Address, m *big.Int, f common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferVaultNotionalFee(&_MarketPlace.TransactOpts, p, u, m, f, a)
}

// MarketPlaceCreateIterator is returned from FilterCreate and is used to iterate over the raw logs and unpacked data for Create events raised by the MarketPlace contract.
type MarketPlaceCreateIterator struct {
	Event *MarketPlaceCreate // Event containing the contract specifics and raw log

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
func (it *MarketPlaceCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketPlaceCreate)
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
		it.Event = new(MarketPlaceCreate)
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
func (it *MarketPlaceCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketPlaceCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketPlaceCreate represents a Create event raised by the MarketPlace contract.
type MarketPlaceCreate struct {
	Protocol     uint8
	Underlying   common.Address
	Maturity     *big.Int
	CToken       common.Address
	ZcToken      common.Address
	VaultTracker common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreate is a free log retrieval operation binding the contract event 0x39fc19388929259a60a8806c6aecb45c336e28e3295fa6fed5813474e6d2b7e8.
//
// Solidity: event Create(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address cToken, address zcToken, address vaultTracker)
func (_MarketPlace *MarketPlaceFilterer) FilterCreate(opts *bind.FilterOpts, protocol []uint8, underlying []common.Address, maturity []*big.Int) (*MarketPlaceCreateIterator, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _MarketPlace.contract.FilterLogs(opts, "Create", protocolRule, underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceCreateIterator{contract: _MarketPlace.contract, event: "Create", logs: logs, sub: sub}, nil
}

// WatchCreate is a free log subscription operation binding the contract event 0x39fc19388929259a60a8806c6aecb45c336e28e3295fa6fed5813474e6d2b7e8.
//
// Solidity: event Create(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address cToken, address zcToken, address vaultTracker)
func (_MarketPlace *MarketPlaceFilterer) WatchCreate(opts *bind.WatchOpts, sink chan<- *MarketPlaceCreate, protocol []uint8, underlying []common.Address, maturity []*big.Int) (event.Subscription, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _MarketPlace.contract.WatchLogs(opts, "Create", protocolRule, underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketPlaceCreate)
				if err := _MarketPlace.contract.UnpackLog(event, "Create", log); err != nil {
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

// ParseCreate is a log parse operation binding the contract event 0x39fc19388929259a60a8806c6aecb45c336e28e3295fa6fed5813474e6d2b7e8.
//
// Solidity: event Create(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address cToken, address zcToken, address vaultTracker)
func (_MarketPlace *MarketPlaceFilterer) ParseCreate(log types.Log) (*MarketPlaceCreate, error) {
	event := new(MarketPlaceCreate)
	if err := _MarketPlace.contract.UnpackLog(event, "Create", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketPlaceCustodialExitIterator is returned from FilterCustodialExit and is used to iterate over the raw logs and unpacked data for CustodialExit events raised by the MarketPlace contract.
type MarketPlaceCustodialExitIterator struct {
	Event *MarketPlaceCustodialExit // Event containing the contract specifics and raw log

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
func (it *MarketPlaceCustodialExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketPlaceCustodialExit)
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
		it.Event = new(MarketPlaceCustodialExit)
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
func (it *MarketPlaceCustodialExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketPlaceCustodialExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketPlaceCustodialExit represents a CustodialExit event raised by the MarketPlace contract.
type MarketPlaceCustodialExit struct {
	Protocol   uint8
	Underlying common.Address
	Maturity   *big.Int
	ZcTarget   common.Address
	NTarget    common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCustodialExit is a free log retrieval operation binding the contract event 0x34fa475e6431f76e8146368a631a5a0ef8ee78b77d3359d0f8e50de4d7a4ff8b.
//
// Solidity: event CustodialExit(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address zcTarget, address nTarget, uint256 amount)
func (_MarketPlace *MarketPlaceFilterer) FilterCustodialExit(opts *bind.FilterOpts, protocol []uint8, underlying []common.Address, maturity []*big.Int) (*MarketPlaceCustodialExitIterator, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _MarketPlace.contract.FilterLogs(opts, "CustodialExit", protocolRule, underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceCustodialExitIterator{contract: _MarketPlace.contract, event: "CustodialExit", logs: logs, sub: sub}, nil
}

// WatchCustodialExit is a free log subscription operation binding the contract event 0x34fa475e6431f76e8146368a631a5a0ef8ee78b77d3359d0f8e50de4d7a4ff8b.
//
// Solidity: event CustodialExit(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address zcTarget, address nTarget, uint256 amount)
func (_MarketPlace *MarketPlaceFilterer) WatchCustodialExit(opts *bind.WatchOpts, sink chan<- *MarketPlaceCustodialExit, protocol []uint8, underlying []common.Address, maturity []*big.Int) (event.Subscription, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _MarketPlace.contract.WatchLogs(opts, "CustodialExit", protocolRule, underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketPlaceCustodialExit)
				if err := _MarketPlace.contract.UnpackLog(event, "CustodialExit", log); err != nil {
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

// ParseCustodialExit is a log parse operation binding the contract event 0x34fa475e6431f76e8146368a631a5a0ef8ee78b77d3359d0f8e50de4d7a4ff8b.
//
// Solidity: event CustodialExit(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address zcTarget, address nTarget, uint256 amount)
func (_MarketPlace *MarketPlaceFilterer) ParseCustodialExit(log types.Log) (*MarketPlaceCustodialExit, error) {
	event := new(MarketPlaceCustodialExit)
	if err := _MarketPlace.contract.UnpackLog(event, "CustodialExit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketPlaceCustodialInitiateIterator is returned from FilterCustodialInitiate and is used to iterate over the raw logs and unpacked data for CustodialInitiate events raised by the MarketPlace contract.
type MarketPlaceCustodialInitiateIterator struct {
	Event *MarketPlaceCustodialInitiate // Event containing the contract specifics and raw log

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
func (it *MarketPlaceCustodialInitiateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketPlaceCustodialInitiate)
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
		it.Event = new(MarketPlaceCustodialInitiate)
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
func (it *MarketPlaceCustodialInitiateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketPlaceCustodialInitiateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketPlaceCustodialInitiate represents a CustodialInitiate event raised by the MarketPlace contract.
type MarketPlaceCustodialInitiate struct {
	Protocol   uint8
	Underlying common.Address
	Maturity   *big.Int
	ZcTarget   common.Address
	NTarget    common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCustodialInitiate is a free log retrieval operation binding the contract event 0x4ccb07dd34f02abf1c514fa611ba589e66ca2f7bcdb3c0cb3b65852354da4398.
//
// Solidity: event CustodialInitiate(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address zcTarget, address nTarget, uint256 amount)
func (_MarketPlace *MarketPlaceFilterer) FilterCustodialInitiate(opts *bind.FilterOpts, protocol []uint8, underlying []common.Address, maturity []*big.Int) (*MarketPlaceCustodialInitiateIterator, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _MarketPlace.contract.FilterLogs(opts, "CustodialInitiate", protocolRule, underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceCustodialInitiateIterator{contract: _MarketPlace.contract, event: "CustodialInitiate", logs: logs, sub: sub}, nil
}

// WatchCustodialInitiate is a free log subscription operation binding the contract event 0x4ccb07dd34f02abf1c514fa611ba589e66ca2f7bcdb3c0cb3b65852354da4398.
//
// Solidity: event CustodialInitiate(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address zcTarget, address nTarget, uint256 amount)
func (_MarketPlace *MarketPlaceFilterer) WatchCustodialInitiate(opts *bind.WatchOpts, sink chan<- *MarketPlaceCustodialInitiate, protocol []uint8, underlying []common.Address, maturity []*big.Int) (event.Subscription, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _MarketPlace.contract.WatchLogs(opts, "CustodialInitiate", protocolRule, underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketPlaceCustodialInitiate)
				if err := _MarketPlace.contract.UnpackLog(event, "CustodialInitiate", log); err != nil {
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

// ParseCustodialInitiate is a log parse operation binding the contract event 0x4ccb07dd34f02abf1c514fa611ba589e66ca2f7bcdb3c0cb3b65852354da4398.
//
// Solidity: event CustodialInitiate(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address zcTarget, address nTarget, uint256 amount)
func (_MarketPlace *MarketPlaceFilterer) ParseCustodialInitiate(log types.Log) (*MarketPlaceCustodialInitiate, error) {
	event := new(MarketPlaceCustodialInitiate)
	if err := _MarketPlace.contract.UnpackLog(event, "CustodialInitiate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketPlaceMatureIterator is returned from FilterMature and is used to iterate over the raw logs and unpacked data for Mature events raised by the MarketPlace contract.
type MarketPlaceMatureIterator struct {
	Event *MarketPlaceMature // Event containing the contract specifics and raw log

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
func (it *MarketPlaceMatureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketPlaceMature)
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
		it.Event = new(MarketPlaceMature)
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
func (it *MarketPlaceMatureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketPlaceMatureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketPlaceMature represents a Mature event raised by the MarketPlace contract.
type MarketPlaceMature struct {
	Protocol     uint8
	Underlying   common.Address
	Maturity     *big.Int
	MaturityRate *big.Int
	Matured      *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMature is a free log retrieval operation binding the contract event 0xa43c0392e4bc23fcadd5a4c4d6d69a1148b6bcec3ac53d7654921bcc33f5addf.
//
// Solidity: event Mature(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, uint256 maturityRate, uint256 matured)
func (_MarketPlace *MarketPlaceFilterer) FilterMature(opts *bind.FilterOpts, protocol []uint8, underlying []common.Address, maturity []*big.Int) (*MarketPlaceMatureIterator, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _MarketPlace.contract.FilterLogs(opts, "Mature", protocolRule, underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceMatureIterator{contract: _MarketPlace.contract, event: "Mature", logs: logs, sub: sub}, nil
}

// WatchMature is a free log subscription operation binding the contract event 0xa43c0392e4bc23fcadd5a4c4d6d69a1148b6bcec3ac53d7654921bcc33f5addf.
//
// Solidity: event Mature(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, uint256 maturityRate, uint256 matured)
func (_MarketPlace *MarketPlaceFilterer) WatchMature(opts *bind.WatchOpts, sink chan<- *MarketPlaceMature, protocol []uint8, underlying []common.Address, maturity []*big.Int) (event.Subscription, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _MarketPlace.contract.WatchLogs(opts, "Mature", protocolRule, underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketPlaceMature)
				if err := _MarketPlace.contract.UnpackLog(event, "Mature", log); err != nil {
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

// ParseMature is a log parse operation binding the contract event 0xa43c0392e4bc23fcadd5a4c4d6d69a1148b6bcec3ac53d7654921bcc33f5addf.
//
// Solidity: event Mature(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, uint256 maturityRate, uint256 matured)
func (_MarketPlace *MarketPlaceFilterer) ParseMature(log types.Log) (*MarketPlaceMature, error) {
	event := new(MarketPlaceMature)
	if err := _MarketPlace.contract.UnpackLog(event, "Mature", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketPlaceP2pVaultExchangeIterator is returned from FilterP2pVaultExchange and is used to iterate over the raw logs and unpacked data for P2pVaultExchange events raised by the MarketPlace contract.
type MarketPlaceP2pVaultExchangeIterator struct {
	Event *MarketPlaceP2pVaultExchange // Event containing the contract specifics and raw log

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
func (it *MarketPlaceP2pVaultExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketPlaceP2pVaultExchange)
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
		it.Event = new(MarketPlaceP2pVaultExchange)
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
func (it *MarketPlaceP2pVaultExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketPlaceP2pVaultExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketPlaceP2pVaultExchange represents a P2pVaultExchange event raised by the MarketPlace contract.
type MarketPlaceP2pVaultExchange struct {
	Protocol   uint8
	Underlying common.Address
	Maturity   *big.Int
	From       common.Address
	To         common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterP2pVaultExchange is a free log retrieval operation binding the contract event 0x4a50decadd365d7ca023f61a307490ce8e696b1b81e112ae6f743b3366d13b6c.
//
// Solidity: event P2pVaultExchange(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address from, address to, uint256 amount)
func (_MarketPlace *MarketPlaceFilterer) FilterP2pVaultExchange(opts *bind.FilterOpts, protocol []uint8, underlying []common.Address, maturity []*big.Int) (*MarketPlaceP2pVaultExchangeIterator, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _MarketPlace.contract.FilterLogs(opts, "P2pVaultExchange", protocolRule, underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceP2pVaultExchangeIterator{contract: _MarketPlace.contract, event: "P2pVaultExchange", logs: logs, sub: sub}, nil
}

// WatchP2pVaultExchange is a free log subscription operation binding the contract event 0x4a50decadd365d7ca023f61a307490ce8e696b1b81e112ae6f743b3366d13b6c.
//
// Solidity: event P2pVaultExchange(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address from, address to, uint256 amount)
func (_MarketPlace *MarketPlaceFilterer) WatchP2pVaultExchange(opts *bind.WatchOpts, sink chan<- *MarketPlaceP2pVaultExchange, protocol []uint8, underlying []common.Address, maturity []*big.Int) (event.Subscription, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _MarketPlace.contract.WatchLogs(opts, "P2pVaultExchange", protocolRule, underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketPlaceP2pVaultExchange)
				if err := _MarketPlace.contract.UnpackLog(event, "P2pVaultExchange", log); err != nil {
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

// ParseP2pVaultExchange is a log parse operation binding the contract event 0x4a50decadd365d7ca023f61a307490ce8e696b1b81e112ae6f743b3366d13b6c.
//
// Solidity: event P2pVaultExchange(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address from, address to, uint256 amount)
func (_MarketPlace *MarketPlaceFilterer) ParseP2pVaultExchange(log types.Log) (*MarketPlaceP2pVaultExchange, error) {
	event := new(MarketPlaceP2pVaultExchange)
	if err := _MarketPlace.contract.UnpackLog(event, "P2pVaultExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketPlaceP2pZcTokenExchangeIterator is returned from FilterP2pZcTokenExchange and is used to iterate over the raw logs and unpacked data for P2pZcTokenExchange events raised by the MarketPlace contract.
type MarketPlaceP2pZcTokenExchangeIterator struct {
	Event *MarketPlaceP2pZcTokenExchange // Event containing the contract specifics and raw log

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
func (it *MarketPlaceP2pZcTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketPlaceP2pZcTokenExchange)
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
		it.Event = new(MarketPlaceP2pZcTokenExchange)
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
func (it *MarketPlaceP2pZcTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketPlaceP2pZcTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketPlaceP2pZcTokenExchange represents a P2pZcTokenExchange event raised by the MarketPlace contract.
type MarketPlaceP2pZcTokenExchange struct {
	Protocol   uint8
	Underlying common.Address
	Maturity   *big.Int
	From       common.Address
	To         common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterP2pZcTokenExchange is a free log retrieval operation binding the contract event 0x0c5d0fa58187faeb475ea625004aed68162ededc2acf69844eed15a090a02b32.
//
// Solidity: event P2pZcTokenExchange(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address from, address to, uint256 amount)
func (_MarketPlace *MarketPlaceFilterer) FilterP2pZcTokenExchange(opts *bind.FilterOpts, protocol []uint8, underlying []common.Address, maturity []*big.Int) (*MarketPlaceP2pZcTokenExchangeIterator, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _MarketPlace.contract.FilterLogs(opts, "P2pZcTokenExchange", protocolRule, underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceP2pZcTokenExchangeIterator{contract: _MarketPlace.contract, event: "P2pZcTokenExchange", logs: logs, sub: sub}, nil
}

// WatchP2pZcTokenExchange is a free log subscription operation binding the contract event 0x0c5d0fa58187faeb475ea625004aed68162ededc2acf69844eed15a090a02b32.
//
// Solidity: event P2pZcTokenExchange(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address from, address to, uint256 amount)
func (_MarketPlace *MarketPlaceFilterer) WatchP2pZcTokenExchange(opts *bind.WatchOpts, sink chan<- *MarketPlaceP2pZcTokenExchange, protocol []uint8, underlying []common.Address, maturity []*big.Int) (event.Subscription, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _MarketPlace.contract.WatchLogs(opts, "P2pZcTokenExchange", protocolRule, underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketPlaceP2pZcTokenExchange)
				if err := _MarketPlace.contract.UnpackLog(event, "P2pZcTokenExchange", log); err != nil {
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

// ParseP2pZcTokenExchange is a log parse operation binding the contract event 0x0c5d0fa58187faeb475ea625004aed68162ededc2acf69844eed15a090a02b32.
//
// Solidity: event P2pZcTokenExchange(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address from, address to, uint256 amount)
func (_MarketPlace *MarketPlaceFilterer) ParseP2pZcTokenExchange(log types.Log) (*MarketPlaceP2pZcTokenExchange, error) {
	event := new(MarketPlaceP2pZcTokenExchange)
	if err := _MarketPlace.contract.UnpackLog(event, "P2pZcTokenExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketPlaceRedeemVaultInterestIterator is returned from FilterRedeemVaultInterest and is used to iterate over the raw logs and unpacked data for RedeemVaultInterest events raised by the MarketPlace contract.
type MarketPlaceRedeemVaultInterestIterator struct {
	Event *MarketPlaceRedeemVaultInterest // Event containing the contract specifics and raw log

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
func (it *MarketPlaceRedeemVaultInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketPlaceRedeemVaultInterest)
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
		it.Event = new(MarketPlaceRedeemVaultInterest)
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
func (it *MarketPlaceRedeemVaultInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketPlaceRedeemVaultInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketPlaceRedeemVaultInterest represents a RedeemVaultInterest event raised by the MarketPlace contract.
type MarketPlaceRedeemVaultInterest struct {
	Protocol   uint8
	Underlying common.Address
	Maturity   *big.Int
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRedeemVaultInterest is a free log retrieval operation binding the contract event 0x602f2da12d1008cffacf50314af0e5f78e8759bba815d4221b390e0b9e73639a.
//
// Solidity: event RedeemVaultInterest(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address sender)
func (_MarketPlace *MarketPlaceFilterer) FilterRedeemVaultInterest(opts *bind.FilterOpts, protocol []uint8, underlying []common.Address, maturity []*big.Int) (*MarketPlaceRedeemVaultInterestIterator, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _MarketPlace.contract.FilterLogs(opts, "RedeemVaultInterest", protocolRule, underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceRedeemVaultInterestIterator{contract: _MarketPlace.contract, event: "RedeemVaultInterest", logs: logs, sub: sub}, nil
}

// WatchRedeemVaultInterest is a free log subscription operation binding the contract event 0x602f2da12d1008cffacf50314af0e5f78e8759bba815d4221b390e0b9e73639a.
//
// Solidity: event RedeemVaultInterest(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address sender)
func (_MarketPlace *MarketPlaceFilterer) WatchRedeemVaultInterest(opts *bind.WatchOpts, sink chan<- *MarketPlaceRedeemVaultInterest, protocol []uint8, underlying []common.Address, maturity []*big.Int) (event.Subscription, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _MarketPlace.contract.WatchLogs(opts, "RedeemVaultInterest", protocolRule, underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketPlaceRedeemVaultInterest)
				if err := _MarketPlace.contract.UnpackLog(event, "RedeemVaultInterest", log); err != nil {
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

// ParseRedeemVaultInterest is a log parse operation binding the contract event 0x602f2da12d1008cffacf50314af0e5f78e8759bba815d4221b390e0b9e73639a.
//
// Solidity: event RedeemVaultInterest(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address sender)
func (_MarketPlace *MarketPlaceFilterer) ParseRedeemVaultInterest(log types.Log) (*MarketPlaceRedeemVaultInterest, error) {
	event := new(MarketPlaceRedeemVaultInterest)
	if err := _MarketPlace.contract.UnpackLog(event, "RedeemVaultInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketPlaceRedeemZcTokenIterator is returned from FilterRedeemZcToken and is used to iterate over the raw logs and unpacked data for RedeemZcToken events raised by the MarketPlace contract.
type MarketPlaceRedeemZcTokenIterator struct {
	Event *MarketPlaceRedeemZcToken // Event containing the contract specifics and raw log

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
func (it *MarketPlaceRedeemZcTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketPlaceRedeemZcToken)
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
		it.Event = new(MarketPlaceRedeemZcToken)
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
func (it *MarketPlaceRedeemZcTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketPlaceRedeemZcTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketPlaceRedeemZcToken represents a RedeemZcToken event raised by the MarketPlace contract.
type MarketPlaceRedeemZcToken struct {
	Protocol   uint8
	Underlying common.Address
	Maturity   *big.Int
	Sender     common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRedeemZcToken is a free log retrieval operation binding the contract event 0x1d3b3ead9f6b17c584914d99c0019883ab43e6e354df48d46185e166f43c68b4.
//
// Solidity: event RedeemZcToken(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address sender, uint256 amount)
func (_MarketPlace *MarketPlaceFilterer) FilterRedeemZcToken(opts *bind.FilterOpts, protocol []uint8, underlying []common.Address, maturity []*big.Int) (*MarketPlaceRedeemZcTokenIterator, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _MarketPlace.contract.FilterLogs(opts, "RedeemZcToken", protocolRule, underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceRedeemZcTokenIterator{contract: _MarketPlace.contract, event: "RedeemZcToken", logs: logs, sub: sub}, nil
}

// WatchRedeemZcToken is a free log subscription operation binding the contract event 0x1d3b3ead9f6b17c584914d99c0019883ab43e6e354df48d46185e166f43c68b4.
//
// Solidity: event RedeemZcToken(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address sender, uint256 amount)
func (_MarketPlace *MarketPlaceFilterer) WatchRedeemZcToken(opts *bind.WatchOpts, sink chan<- *MarketPlaceRedeemZcToken, protocol []uint8, underlying []common.Address, maturity []*big.Int) (event.Subscription, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _MarketPlace.contract.WatchLogs(opts, "RedeemZcToken", protocolRule, underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketPlaceRedeemZcToken)
				if err := _MarketPlace.contract.UnpackLog(event, "RedeemZcToken", log); err != nil {
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

// ParseRedeemZcToken is a log parse operation binding the contract event 0x1d3b3ead9f6b17c584914d99c0019883ab43e6e354df48d46185e166f43c68b4.
//
// Solidity: event RedeemZcToken(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address sender, uint256 amount)
func (_MarketPlace *MarketPlaceFilterer) ParseRedeemZcToken(log types.Log) (*MarketPlaceRedeemZcToken, error) {
	event := new(MarketPlaceRedeemZcToken)
	if err := _MarketPlace.contract.UnpackLog(event, "RedeemZcToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketPlaceSetAdminIterator is returned from FilterSetAdmin and is used to iterate over the raw logs and unpacked data for SetAdmin events raised by the MarketPlace contract.
type MarketPlaceSetAdminIterator struct {
	Event *MarketPlaceSetAdmin // Event containing the contract specifics and raw log

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
func (it *MarketPlaceSetAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketPlaceSetAdmin)
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
		it.Event = new(MarketPlaceSetAdmin)
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
func (it *MarketPlaceSetAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketPlaceSetAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketPlaceSetAdmin represents a SetAdmin event raised by the MarketPlace contract.
type MarketPlaceSetAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetAdmin is a free log retrieval operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address indexed admin)
func (_MarketPlace *MarketPlaceFilterer) FilterSetAdmin(opts *bind.FilterOpts, admin []common.Address) (*MarketPlaceSetAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _MarketPlace.contract.FilterLogs(opts, "SetAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceSetAdminIterator{contract: _MarketPlace.contract, event: "SetAdmin", logs: logs, sub: sub}, nil
}

// WatchSetAdmin is a free log subscription operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address indexed admin)
func (_MarketPlace *MarketPlaceFilterer) WatchSetAdmin(opts *bind.WatchOpts, sink chan<- *MarketPlaceSetAdmin, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _MarketPlace.contract.WatchLogs(opts, "SetAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketPlaceSetAdmin)
				if err := _MarketPlace.contract.UnpackLog(event, "SetAdmin", log); err != nil {
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

// ParseSetAdmin is a log parse operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address indexed admin)
func (_MarketPlace *MarketPlaceFilterer) ParseSetAdmin(log types.Log) (*MarketPlaceSetAdmin, error) {
	event := new(MarketPlaceSetAdmin)
	if err := _MarketPlace.contract.UnpackLog(event, "SetAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketPlaceTransferVaultNotionalIterator is returned from FilterTransferVaultNotional and is used to iterate over the raw logs and unpacked data for TransferVaultNotional events raised by the MarketPlace contract.
type MarketPlaceTransferVaultNotionalIterator struct {
	Event *MarketPlaceTransferVaultNotional // Event containing the contract specifics and raw log

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
func (it *MarketPlaceTransferVaultNotionalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketPlaceTransferVaultNotional)
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
		it.Event = new(MarketPlaceTransferVaultNotional)
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
func (it *MarketPlaceTransferVaultNotionalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketPlaceTransferVaultNotionalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketPlaceTransferVaultNotional represents a TransferVaultNotional event raised by the MarketPlace contract.
type MarketPlaceTransferVaultNotional struct {
	Protocol   uint8
	Underlying common.Address
	Maturity   *big.Int
	From       common.Address
	To         common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransferVaultNotional is a free log retrieval operation binding the contract event 0xe401e2d61a180e97aba5ebeb66d643bbc7e6516b91281ff0e480dac7e206c88f.
//
// Solidity: event TransferVaultNotional(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address from, address to, uint256 amount)
func (_MarketPlace *MarketPlaceFilterer) FilterTransferVaultNotional(opts *bind.FilterOpts, protocol []uint8, underlying []common.Address, maturity []*big.Int) (*MarketPlaceTransferVaultNotionalIterator, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _MarketPlace.contract.FilterLogs(opts, "TransferVaultNotional", protocolRule, underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceTransferVaultNotionalIterator{contract: _MarketPlace.contract, event: "TransferVaultNotional", logs: logs, sub: sub}, nil
}

// WatchTransferVaultNotional is a free log subscription operation binding the contract event 0xe401e2d61a180e97aba5ebeb66d643bbc7e6516b91281ff0e480dac7e206c88f.
//
// Solidity: event TransferVaultNotional(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address from, address to, uint256 amount)
func (_MarketPlace *MarketPlaceFilterer) WatchTransferVaultNotional(opts *bind.WatchOpts, sink chan<- *MarketPlaceTransferVaultNotional, protocol []uint8, underlying []common.Address, maturity []*big.Int) (event.Subscription, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _MarketPlace.contract.WatchLogs(opts, "TransferVaultNotional", protocolRule, underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketPlaceTransferVaultNotional)
				if err := _MarketPlace.contract.UnpackLog(event, "TransferVaultNotional", log); err != nil {
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

// ParseTransferVaultNotional is a log parse operation binding the contract event 0xe401e2d61a180e97aba5ebeb66d643bbc7e6516b91281ff0e480dac7e206c88f.
//
// Solidity: event TransferVaultNotional(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address from, address to, uint256 amount)
func (_MarketPlace *MarketPlaceFilterer) ParseTransferVaultNotional(log types.Log) (*MarketPlaceTransferVaultNotional, error) {
	event := new(MarketPlaceTransferVaultNotional)
	if err := _MarketPlace.contract.UnpackLog(event, "TransferVaultNotional", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

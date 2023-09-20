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

// HashOrder is an auto generated low-level Go binding around an user-defined struct.
type HashOrder struct {
	Key        [32]byte
	Protocol   uint8
	Maker      common.Address
	Underlying common.Address
	Vault      bool
	Exit       bool
	Principal  *big.Int
	Premium    *big.Int
	Maturity   *big.Int
	Expiry     *big.Int
}

// SigComponents is an auto generated low-level Go binding around an user-defined struct.
type SigComponents struct {
	V uint8
	R [32]byte
	S [32]byte
}

// SwivelMetaData contains all meta data concerning the Swivel contract.
var SwivelMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"m\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ApproveFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Exception\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"S\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"blocked\",\"type\":\"address\"}],\"name\":\"BlockApproval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BlockFeeChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BlockWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"Cancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ChangeFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"vault\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"filled\",\"type\":\"uint256\"}],\"name\":\"Exit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"vault\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"filled\",\"type\":\"uint256\"}],\"name\":\"Initiate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hold\",\"type\":\"uint256\"}],\"name\":\"ScheduleApproval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16[4]\",\"name\":\"proposal\",\"type\":\"uint16[4]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hold\",\"type\":\"uint256\"}],\"name\":\"ScheduleFeeChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hold\",\"type\":\"uint256\"}],\"name\":\"ScheduleWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"SetAdmin\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"HOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_FEENOMINATOR\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aaveAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"u\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"c\",\"type\":\"address[]\"}],\"name\":\"approveUnderlying\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"authRedeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"e\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"blockApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockFeeChange\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"e\",\"type\":\"address\"}],\"name\":\"blockWithdrawal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"protocol\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"vault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structHash.Order[]\",\"name\":\"o\",\"type\":\"tuple[]\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cancelled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16[4]\",\"name\":\"f\",\"type\":\"uint16[4]\"}],\"name\":\"changeFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"combineTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domain\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"protocol\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"vault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structHash.Order[]\",\"name\":\"o\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"a\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSig.Components[]\",\"name\":\"c\",\"type\":\"tuple[]\"}],\"name\":\"exit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeChange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"feenominators\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"filled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"protocol\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"vault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structHash.Order[]\",\"name\":\"o\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"a\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSig.Components[]\",\"name\":\"c\",\"type\":\"tuple[]\"}],\"name\":\"initiate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketPlace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"redeemSwivelVaultInterest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"redeemVaultInterest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"redeemZcToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"e\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"scheduleApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16[4]\",\"name\":\"f\",\"type\":\"uint16[4]\"}],\"name\":\"scheduleFeeChange\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"e\",\"type\":\"address\"}],\"name\":\"scheduleWithdrawal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"splitUnderlying\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"e\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SwivelABI is the input ABI used to generate the binding from.
// Deprecated: Use SwivelMetaData.ABI instead.
var SwivelABI = SwivelMetaData.ABI

// Swivel is an auto generated Go binding around an Ethereum contract.
type Swivel struct {
	SwivelCaller     // Read-only binding to the contract
	SwivelTransactor // Write-only binding to the contract
	SwivelFilterer   // Log filterer for contract events
}

// SwivelCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwivelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwivelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwivelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwivelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwivelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwivelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwivelSession struct {
	Contract     *Swivel           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwivelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwivelCallerSession struct {
	Contract *SwivelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SwivelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwivelTransactorSession struct {
	Contract     *SwivelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwivelRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwivelRaw struct {
	Contract *Swivel // Generic contract binding to access the raw methods on
}

// SwivelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwivelCallerRaw struct {
	Contract *SwivelCaller // Generic read-only contract binding to access the raw methods on
}

// SwivelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwivelTransactorRaw struct {
	Contract *SwivelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwivel creates a new instance of Swivel, bound to a specific deployed contract.
func NewSwivel(address common.Address, backend bind.ContractBackend) (*Swivel, error) {
	contract, err := bindSwivel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swivel{SwivelCaller: SwivelCaller{contract: contract}, SwivelTransactor: SwivelTransactor{contract: contract}, SwivelFilterer: SwivelFilterer{contract: contract}}, nil
}

// NewSwivelCaller creates a new read-only instance of Swivel, bound to a specific deployed contract.
func NewSwivelCaller(address common.Address, caller bind.ContractCaller) (*SwivelCaller, error) {
	contract, err := bindSwivel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwivelCaller{contract: contract}, nil
}

// NewSwivelTransactor creates a new write-only instance of Swivel, bound to a specific deployed contract.
func NewSwivelTransactor(address common.Address, transactor bind.ContractTransactor) (*SwivelTransactor, error) {
	contract, err := bindSwivel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwivelTransactor{contract: contract}, nil
}

// NewSwivelFilterer creates a new log filterer instance of Swivel, bound to a specific deployed contract.
func NewSwivelFilterer(address common.Address, filterer bind.ContractFilterer) (*SwivelFilterer, error) {
	contract, err := bindSwivel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwivelFilterer{contract: contract}, nil
}

// bindSwivel binds a generic wrapper to an already deployed contract.
func bindSwivel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwivelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swivel *SwivelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swivel.Contract.SwivelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swivel *SwivelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swivel.Contract.SwivelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swivel *SwivelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swivel.Contract.SwivelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swivel *SwivelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swivel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swivel *SwivelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swivel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swivel *SwivelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swivel.Contract.contract.Transact(opts, method, params...)
}

// HOLD is a free data retrieval call binding the contract method 0xd0886f97.
//
// Solidity: function HOLD() view returns(uint256)
func (_Swivel *SwivelCaller) HOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "HOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HOLD is a free data retrieval call binding the contract method 0xd0886f97.
//
// Solidity: function HOLD() view returns(uint256)
func (_Swivel *SwivelSession) HOLD() (*big.Int, error) {
	return _Swivel.Contract.HOLD(&_Swivel.CallOpts)
}

// HOLD is a free data retrieval call binding the contract method 0xd0886f97.
//
// Solidity: function HOLD() view returns(uint256)
func (_Swivel *SwivelCallerSession) HOLD() (*big.Int, error) {
	return _Swivel.Contract.HOLD(&_Swivel.CallOpts)
}

// MINFEENOMINATOR is a free data retrieval call binding the contract method 0x0d3f5352.
//
// Solidity: function MIN_FEENOMINATOR() view returns(uint16)
func (_Swivel *SwivelCaller) MINFEENOMINATOR(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "MIN_FEENOMINATOR")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MINFEENOMINATOR is a free data retrieval call binding the contract method 0x0d3f5352.
//
// Solidity: function MIN_FEENOMINATOR() view returns(uint16)
func (_Swivel *SwivelSession) MINFEENOMINATOR() (uint16, error) {
	return _Swivel.Contract.MINFEENOMINATOR(&_Swivel.CallOpts)
}

// MINFEENOMINATOR is a free data retrieval call binding the contract method 0x0d3f5352.
//
// Solidity: function MIN_FEENOMINATOR() view returns(uint16)
func (_Swivel *SwivelCallerSession) MINFEENOMINATOR() (uint16, error) {
	return _Swivel.Contract.MINFEENOMINATOR(&_Swivel.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Swivel *SwivelCaller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Swivel *SwivelSession) NAME() (string, error) {
	return _Swivel.Contract.NAME(&_Swivel.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Swivel *SwivelCallerSession) NAME() (string, error) {
	return _Swivel.Contract.NAME(&_Swivel.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Swivel *SwivelCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Swivel *SwivelSession) VERSION() (string, error) {
	return _Swivel.Contract.VERSION(&_Swivel.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Swivel *SwivelCallerSession) VERSION() (string, error) {
	return _Swivel.Contract.VERSION(&_Swivel.CallOpts)
}

// AaveAddr is a free data retrieval call binding the contract method 0x81bb8d63.
//
// Solidity: function aaveAddr() view returns(address)
func (_Swivel *SwivelCaller) AaveAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "aaveAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AaveAddr is a free data retrieval call binding the contract method 0x81bb8d63.
//
// Solidity: function aaveAddr() view returns(address)
func (_Swivel *SwivelSession) AaveAddr() (common.Address, error) {
	return _Swivel.Contract.AaveAddr(&_Swivel.CallOpts)
}

// AaveAddr is a free data retrieval call binding the contract method 0x81bb8d63.
//
// Solidity: function aaveAddr() view returns(address)
func (_Swivel *SwivelCallerSession) AaveAddr() (common.Address, error) {
	return _Swivel.Contract.AaveAddr(&_Swivel.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Swivel *SwivelCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Swivel *SwivelSession) Admin() (common.Address, error) {
	return _Swivel.Contract.Admin(&_Swivel.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Swivel *SwivelCallerSession) Admin() (common.Address, error) {
	return _Swivel.Contract.Admin(&_Swivel.CallOpts)
}

// Approvals is a free data retrieval call binding the contract method 0xa32ce11e.
//
// Solidity: function approvals(address , address ) view returns(uint256)
func (_Swivel *SwivelCaller) Approvals(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "approvals", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Approvals is a free data retrieval call binding the contract method 0xa32ce11e.
//
// Solidity: function approvals(address , address ) view returns(uint256)
func (_Swivel *SwivelSession) Approvals(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Swivel.Contract.Approvals(&_Swivel.CallOpts, arg0, arg1)
}

// Approvals is a free data retrieval call binding the contract method 0xa32ce11e.
//
// Solidity: function approvals(address , address ) view returns(uint256)
func (_Swivel *SwivelCallerSession) Approvals(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Swivel.Contract.Approvals(&_Swivel.CallOpts, arg0, arg1)
}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled(bytes32 ) view returns(bool)
func (_Swivel *SwivelCaller) Cancelled(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "cancelled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled(bytes32 ) view returns(bool)
func (_Swivel *SwivelSession) Cancelled(arg0 [32]byte) (bool, error) {
	return _Swivel.Contract.Cancelled(&_Swivel.CallOpts, arg0)
}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled(bytes32 ) view returns(bool)
func (_Swivel *SwivelCallerSession) Cancelled(arg0 [32]byte) (bool, error) {
	return _Swivel.Contract.Cancelled(&_Swivel.CallOpts, arg0)
}

// Domain is a free data retrieval call binding the contract method 0xc2fb26a6.
//
// Solidity: function domain() view returns(bytes32)
func (_Swivel *SwivelCaller) Domain(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "domain")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Domain is a free data retrieval call binding the contract method 0xc2fb26a6.
//
// Solidity: function domain() view returns(bytes32)
func (_Swivel *SwivelSession) Domain() ([32]byte, error) {
	return _Swivel.Contract.Domain(&_Swivel.CallOpts)
}

// Domain is a free data retrieval call binding the contract method 0xc2fb26a6.
//
// Solidity: function domain() view returns(bytes32)
func (_Swivel *SwivelCallerSession) Domain() ([32]byte, error) {
	return _Swivel.Contract.Domain(&_Swivel.CallOpts)
}

// FeeChange is a free data retrieval call binding the contract method 0x35197f9e.
//
// Solidity: function feeChange() view returns(uint256)
func (_Swivel *SwivelCaller) FeeChange(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "feeChange")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeChange is a free data retrieval call binding the contract method 0x35197f9e.
//
// Solidity: function feeChange() view returns(uint256)
func (_Swivel *SwivelSession) FeeChange() (*big.Int, error) {
	return _Swivel.Contract.FeeChange(&_Swivel.CallOpts)
}

// FeeChange is a free data retrieval call binding the contract method 0x35197f9e.
//
// Solidity: function feeChange() view returns(uint256)
func (_Swivel *SwivelCallerSession) FeeChange() (*big.Int, error) {
	return _Swivel.Contract.FeeChange(&_Swivel.CallOpts)
}

// Feenominators is a free data retrieval call binding the contract method 0x95cb60c4.
//
// Solidity: function feenominators(uint256 ) view returns(uint16)
func (_Swivel *SwivelCaller) Feenominators(opts *bind.CallOpts, arg0 *big.Int) (uint16, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "feenominators", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Feenominators is a free data retrieval call binding the contract method 0x95cb60c4.
//
// Solidity: function feenominators(uint256 ) view returns(uint16)
func (_Swivel *SwivelSession) Feenominators(arg0 *big.Int) (uint16, error) {
	return _Swivel.Contract.Feenominators(&_Swivel.CallOpts, arg0)
}

// Feenominators is a free data retrieval call binding the contract method 0x95cb60c4.
//
// Solidity: function feenominators(uint256 ) view returns(uint16)
func (_Swivel *SwivelCallerSession) Feenominators(arg0 *big.Int) (uint16, error) {
	return _Swivel.Contract.Feenominators(&_Swivel.CallOpts, arg0)
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) view returns(uint256)
func (_Swivel *SwivelCaller) Filled(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "filled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) view returns(uint256)
func (_Swivel *SwivelSession) Filled(arg0 [32]byte) (*big.Int, error) {
	return _Swivel.Contract.Filled(&_Swivel.CallOpts, arg0)
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) view returns(uint256)
func (_Swivel *SwivelCallerSession) Filled(arg0 [32]byte) (*big.Int, error) {
	return _Swivel.Contract.Filled(&_Swivel.CallOpts, arg0)
}

// MarketPlace is a free data retrieval call binding the contract method 0x2e25d2a6.
//
// Solidity: function marketPlace() view returns(address)
func (_Swivel *SwivelCaller) MarketPlace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "marketPlace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketPlace is a free data retrieval call binding the contract method 0x2e25d2a6.
//
// Solidity: function marketPlace() view returns(address)
func (_Swivel *SwivelSession) MarketPlace() (common.Address, error) {
	return _Swivel.Contract.MarketPlace(&_Swivel.CallOpts)
}

// MarketPlace is a free data retrieval call binding the contract method 0x2e25d2a6.
//
// Solidity: function marketPlace() view returns(address)
func (_Swivel *SwivelCallerSession) MarketPlace() (common.Address, error) {
	return _Swivel.Contract.MarketPlace(&_Swivel.CallOpts)
}

// Withdrawals is a free data retrieval call binding the contract method 0x7a9262a2.
//
// Solidity: function withdrawals(address ) view returns(uint256)
func (_Swivel *SwivelCaller) Withdrawals(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "withdrawals", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Withdrawals is a free data retrieval call binding the contract method 0x7a9262a2.
//
// Solidity: function withdrawals(address ) view returns(uint256)
func (_Swivel *SwivelSession) Withdrawals(arg0 common.Address) (*big.Int, error) {
	return _Swivel.Contract.Withdrawals(&_Swivel.CallOpts, arg0)
}

// Withdrawals is a free data retrieval call binding the contract method 0x7a9262a2.
//
// Solidity: function withdrawals(address ) view returns(uint256)
func (_Swivel *SwivelCallerSession) Withdrawals(arg0 common.Address) (*big.Int, error) {
	return _Swivel.Contract.Withdrawals(&_Swivel.CallOpts, arg0)
}

// ApproveUnderlying is a paid mutator transaction binding the contract method 0x25dedb85.
//
// Solidity: function approveUnderlying(address[] u, address[] c) returns(bool)
func (_Swivel *SwivelTransactor) ApproveUnderlying(opts *bind.TransactOpts, u []common.Address, c []common.Address) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "approveUnderlying", u, c)
}

// ApproveUnderlying is a paid mutator transaction binding the contract method 0x25dedb85.
//
// Solidity: function approveUnderlying(address[] u, address[] c) returns(bool)
func (_Swivel *SwivelSession) ApproveUnderlying(u []common.Address, c []common.Address) (*types.Transaction, error) {
	return _Swivel.Contract.ApproveUnderlying(&_Swivel.TransactOpts, u, c)
}

// ApproveUnderlying is a paid mutator transaction binding the contract method 0x25dedb85.
//
// Solidity: function approveUnderlying(address[] u, address[] c) returns(bool)
func (_Swivel *SwivelTransactorSession) ApproveUnderlying(u []common.Address, c []common.Address) (*types.Transaction, error) {
	return _Swivel.Contract.ApproveUnderlying(&_Swivel.TransactOpts, u, c)
}

// AuthRedeem is a paid mutator transaction binding the contract method 0x8bfa66be.
//
// Solidity: function authRedeem(uint8 p, address u, address c, address t, uint256 a) returns(bool)
func (_Swivel *SwivelTransactor) AuthRedeem(opts *bind.TransactOpts, p uint8, u common.Address, c common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "authRedeem", p, u, c, t, a)
}

// AuthRedeem is a paid mutator transaction binding the contract method 0x8bfa66be.
//
// Solidity: function authRedeem(uint8 p, address u, address c, address t, uint256 a) returns(bool)
func (_Swivel *SwivelSession) AuthRedeem(p uint8, u common.Address, c common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Swivel.Contract.AuthRedeem(&_Swivel.TransactOpts, p, u, c, t, a)
}

// AuthRedeem is a paid mutator transaction binding the contract method 0x8bfa66be.
//
// Solidity: function authRedeem(uint8 p, address u, address c, address t, uint256 a) returns(bool)
func (_Swivel *SwivelTransactorSession) AuthRedeem(p uint8, u common.Address, c common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Swivel.Contract.AuthRedeem(&_Swivel.TransactOpts, p, u, c, t, a)
}

// BlockApproval is a paid mutator transaction binding the contract method 0x2c4957b2.
//
// Solidity: function blockApproval(address e, address a) returns(bool)
func (_Swivel *SwivelTransactor) BlockApproval(opts *bind.TransactOpts, e common.Address, a common.Address) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "blockApproval", e, a)
}

// BlockApproval is a paid mutator transaction binding the contract method 0x2c4957b2.
//
// Solidity: function blockApproval(address e, address a) returns(bool)
func (_Swivel *SwivelSession) BlockApproval(e common.Address, a common.Address) (*types.Transaction, error) {
	return _Swivel.Contract.BlockApproval(&_Swivel.TransactOpts, e, a)
}

// BlockApproval is a paid mutator transaction binding the contract method 0x2c4957b2.
//
// Solidity: function blockApproval(address e, address a) returns(bool)
func (_Swivel *SwivelTransactorSession) BlockApproval(e common.Address, a common.Address) (*types.Transaction, error) {
	return _Swivel.Contract.BlockApproval(&_Swivel.TransactOpts, e, a)
}

// BlockFeeChange is a paid mutator transaction binding the contract method 0xf9ad473d.
//
// Solidity: function blockFeeChange() returns(bool)
func (_Swivel *SwivelTransactor) BlockFeeChange(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "blockFeeChange")
}

// BlockFeeChange is a paid mutator transaction binding the contract method 0xf9ad473d.
//
// Solidity: function blockFeeChange() returns(bool)
func (_Swivel *SwivelSession) BlockFeeChange() (*types.Transaction, error) {
	return _Swivel.Contract.BlockFeeChange(&_Swivel.TransactOpts)
}

// BlockFeeChange is a paid mutator transaction binding the contract method 0xf9ad473d.
//
// Solidity: function blockFeeChange() returns(bool)
func (_Swivel *SwivelTransactorSession) BlockFeeChange() (*types.Transaction, error) {
	return _Swivel.Contract.BlockFeeChange(&_Swivel.TransactOpts)
}

// BlockWithdrawal is a paid mutator transaction binding the contract method 0xa102e384.
//
// Solidity: function blockWithdrawal(address e) returns(bool)
func (_Swivel *SwivelTransactor) BlockWithdrawal(opts *bind.TransactOpts, e common.Address) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "blockWithdrawal", e)
}

// BlockWithdrawal is a paid mutator transaction binding the contract method 0xa102e384.
//
// Solidity: function blockWithdrawal(address e) returns(bool)
func (_Swivel *SwivelSession) BlockWithdrawal(e common.Address) (*types.Transaction, error) {
	return _Swivel.Contract.BlockWithdrawal(&_Swivel.TransactOpts, e)
}

// BlockWithdrawal is a paid mutator transaction binding the contract method 0xa102e384.
//
// Solidity: function blockWithdrawal(address e) returns(bool)
func (_Swivel *SwivelTransactorSession) BlockWithdrawal(e common.Address) (*types.Transaction, error) {
	return _Swivel.Contract.BlockWithdrawal(&_Swivel.TransactOpts, e)
}

// Cancel is a paid mutator transaction binding the contract method 0x12a444fa.
//
// Solidity: function cancel((bytes32,uint8,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o) returns(bool)
func (_Swivel *SwivelTransactor) Cancel(opts *bind.TransactOpts, o []HashOrder) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "cancel", o)
}

// Cancel is a paid mutator transaction binding the contract method 0x12a444fa.
//
// Solidity: function cancel((bytes32,uint8,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o) returns(bool)
func (_Swivel *SwivelSession) Cancel(o []HashOrder) (*types.Transaction, error) {
	return _Swivel.Contract.Cancel(&_Swivel.TransactOpts, o)
}

// Cancel is a paid mutator transaction binding the contract method 0x12a444fa.
//
// Solidity: function cancel((bytes32,uint8,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o) returns(bool)
func (_Swivel *SwivelTransactorSession) Cancel(o []HashOrder) (*types.Transaction, error) {
	return _Swivel.Contract.Cancel(&_Swivel.TransactOpts, o)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x172bfa94.
//
// Solidity: function changeFee(uint16[4] f) returns(bool)
func (_Swivel *SwivelTransactor) ChangeFee(opts *bind.TransactOpts, f [4]uint16) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "changeFee", f)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x172bfa94.
//
// Solidity: function changeFee(uint16[4] f) returns(bool)
func (_Swivel *SwivelSession) ChangeFee(f [4]uint16) (*types.Transaction, error) {
	return _Swivel.Contract.ChangeFee(&_Swivel.TransactOpts, f)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x172bfa94.
//
// Solidity: function changeFee(uint16[4] f) returns(bool)
func (_Swivel *SwivelTransactorSession) ChangeFee(f [4]uint16) (*types.Transaction, error) {
	return _Swivel.Contract.ChangeFee(&_Swivel.TransactOpts, f)
}

// CombineTokens is a paid mutator transaction binding the contract method 0x2fd9ee8a.
//
// Solidity: function combineTokens(uint8 p, address u, uint256 m, uint256 a) returns(bool)
func (_Swivel *SwivelTransactor) CombineTokens(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "combineTokens", p, u, m, a)
}

// CombineTokens is a paid mutator transaction binding the contract method 0x2fd9ee8a.
//
// Solidity: function combineTokens(uint8 p, address u, uint256 m, uint256 a) returns(bool)
func (_Swivel *SwivelSession) CombineTokens(p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _Swivel.Contract.CombineTokens(&_Swivel.TransactOpts, p, u, m, a)
}

// CombineTokens is a paid mutator transaction binding the contract method 0x2fd9ee8a.
//
// Solidity: function combineTokens(uint8 p, address u, uint256 m, uint256 a) returns(bool)
func (_Swivel *SwivelTransactorSession) CombineTokens(p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _Swivel.Contract.CombineTokens(&_Swivel.TransactOpts, p, u, m, a)
}

// Exit is a paid mutator transaction binding the contract method 0x4af60607.
//
// Solidity: function exit((bytes32,uint8,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] c) returns(bool)
func (_Swivel *SwivelTransactor) Exit(opts *bind.TransactOpts, o []HashOrder, a []*big.Int, c []SigComponents) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "exit", o, a, c)
}

// Exit is a paid mutator transaction binding the contract method 0x4af60607.
//
// Solidity: function exit((bytes32,uint8,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] c) returns(bool)
func (_Swivel *SwivelSession) Exit(o []HashOrder, a []*big.Int, c []SigComponents) (*types.Transaction, error) {
	return _Swivel.Contract.Exit(&_Swivel.TransactOpts, o, a, c)
}

// Exit is a paid mutator transaction binding the contract method 0x4af60607.
//
// Solidity: function exit((bytes32,uint8,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] c) returns(bool)
func (_Swivel *SwivelTransactorSession) Exit(o []HashOrder, a []*big.Int, c []SigComponents) (*types.Transaction, error) {
	return _Swivel.Contract.Exit(&_Swivel.TransactOpts, o, a, c)
}

// Initiate is a paid mutator transaction binding the contract method 0x10510f11.
//
// Solidity: function initiate((bytes32,uint8,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] c) returns(bool)
func (_Swivel *SwivelTransactor) Initiate(opts *bind.TransactOpts, o []HashOrder, a []*big.Int, c []SigComponents) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "initiate", o, a, c)
}

// Initiate is a paid mutator transaction binding the contract method 0x10510f11.
//
// Solidity: function initiate((bytes32,uint8,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] c) returns(bool)
func (_Swivel *SwivelSession) Initiate(o []HashOrder, a []*big.Int, c []SigComponents) (*types.Transaction, error) {
	return _Swivel.Contract.Initiate(&_Swivel.TransactOpts, o, a, c)
}

// Initiate is a paid mutator transaction binding the contract method 0x10510f11.
//
// Solidity: function initiate((bytes32,uint8,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] c) returns(bool)
func (_Swivel *SwivelTransactorSession) Initiate(o []HashOrder, a []*big.Int, c []SigComponents) (*types.Transaction, error) {
	return _Swivel.Contract.Initiate(&_Swivel.TransactOpts, o, a, c)
}

// RedeemSwivelVaultInterest is a paid mutator transaction binding the contract method 0xb6907677.
//
// Solidity: function redeemSwivelVaultInterest(uint8 p, address u, uint256 m) returns(bool)
func (_Swivel *SwivelTransactor) RedeemSwivelVaultInterest(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "redeemSwivelVaultInterest", p, u, m)
}

// RedeemSwivelVaultInterest is a paid mutator transaction binding the contract method 0xb6907677.
//
// Solidity: function redeemSwivelVaultInterest(uint8 p, address u, uint256 m) returns(bool)
func (_Swivel *SwivelSession) RedeemSwivelVaultInterest(p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _Swivel.Contract.RedeemSwivelVaultInterest(&_Swivel.TransactOpts, p, u, m)
}

// RedeemSwivelVaultInterest is a paid mutator transaction binding the contract method 0xb6907677.
//
// Solidity: function redeemSwivelVaultInterest(uint8 p, address u, uint256 m) returns(bool)
func (_Swivel *SwivelTransactorSession) RedeemSwivelVaultInterest(p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _Swivel.Contract.RedeemSwivelVaultInterest(&_Swivel.TransactOpts, p, u, m)
}

// RedeemVaultInterest is a paid mutator transaction binding the contract method 0xa0b06b50.
//
// Solidity: function redeemVaultInterest(uint8 p, address u, uint256 m) returns(bool)
func (_Swivel *SwivelTransactor) RedeemVaultInterest(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "redeemVaultInterest", p, u, m)
}

// RedeemVaultInterest is a paid mutator transaction binding the contract method 0xa0b06b50.
//
// Solidity: function redeemVaultInterest(uint8 p, address u, uint256 m) returns(bool)
func (_Swivel *SwivelSession) RedeemVaultInterest(p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _Swivel.Contract.RedeemVaultInterest(&_Swivel.TransactOpts, p, u, m)
}

// RedeemVaultInterest is a paid mutator transaction binding the contract method 0xa0b06b50.
//
// Solidity: function redeemVaultInterest(uint8 p, address u, uint256 m) returns(bool)
func (_Swivel *SwivelTransactorSession) RedeemVaultInterest(p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _Swivel.Contract.RedeemVaultInterest(&_Swivel.TransactOpts, p, u, m)
}

// RedeemZcToken is a paid mutator transaction binding the contract method 0xb16a33cc.
//
// Solidity: function redeemZcToken(uint8 p, address u, uint256 m, uint256 a) returns(bool)
func (_Swivel *SwivelTransactor) RedeemZcToken(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "redeemZcToken", p, u, m, a)
}

// RedeemZcToken is a paid mutator transaction binding the contract method 0xb16a33cc.
//
// Solidity: function redeemZcToken(uint8 p, address u, uint256 m, uint256 a) returns(bool)
func (_Swivel *SwivelSession) RedeemZcToken(p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _Swivel.Contract.RedeemZcToken(&_Swivel.TransactOpts, p, u, m, a)
}

// RedeemZcToken is a paid mutator transaction binding the contract method 0xb16a33cc.
//
// Solidity: function redeemZcToken(uint8 p, address u, uint256 m, uint256 a) returns(bool)
func (_Swivel *SwivelTransactorSession) RedeemZcToken(p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _Swivel.Contract.RedeemZcToken(&_Swivel.TransactOpts, p, u, m, a)
}

// ScheduleApproval is a paid mutator transaction binding the contract method 0x830e4d64.
//
// Solidity: function scheduleApproval(address e, address a) returns(bool)
func (_Swivel *SwivelTransactor) ScheduleApproval(opts *bind.TransactOpts, e common.Address, a common.Address) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "scheduleApproval", e, a)
}

// ScheduleApproval is a paid mutator transaction binding the contract method 0x830e4d64.
//
// Solidity: function scheduleApproval(address e, address a) returns(bool)
func (_Swivel *SwivelSession) ScheduleApproval(e common.Address, a common.Address) (*types.Transaction, error) {
	return _Swivel.Contract.ScheduleApproval(&_Swivel.TransactOpts, e, a)
}

// ScheduleApproval is a paid mutator transaction binding the contract method 0x830e4d64.
//
// Solidity: function scheduleApproval(address e, address a) returns(bool)
func (_Swivel *SwivelTransactorSession) ScheduleApproval(e common.Address, a common.Address) (*types.Transaction, error) {
	return _Swivel.Contract.ScheduleApproval(&_Swivel.TransactOpts, e, a)
}

// ScheduleFeeChange is a paid mutator transaction binding the contract method 0x761a7e70.
//
// Solidity: function scheduleFeeChange(uint16[4] f) returns(bool)
func (_Swivel *SwivelTransactor) ScheduleFeeChange(opts *bind.TransactOpts, f [4]uint16) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "scheduleFeeChange", f)
}

// ScheduleFeeChange is a paid mutator transaction binding the contract method 0x761a7e70.
//
// Solidity: function scheduleFeeChange(uint16[4] f) returns(bool)
func (_Swivel *SwivelSession) ScheduleFeeChange(f [4]uint16) (*types.Transaction, error) {
	return _Swivel.Contract.ScheduleFeeChange(&_Swivel.TransactOpts, f)
}

// ScheduleFeeChange is a paid mutator transaction binding the contract method 0x761a7e70.
//
// Solidity: function scheduleFeeChange(uint16[4] f) returns(bool)
func (_Swivel *SwivelTransactorSession) ScheduleFeeChange(f [4]uint16) (*types.Transaction, error) {
	return _Swivel.Contract.ScheduleFeeChange(&_Swivel.TransactOpts, f)
}

// ScheduleWithdrawal is a paid mutator transaction binding the contract method 0xf8eaad35.
//
// Solidity: function scheduleWithdrawal(address e) returns(bool)
func (_Swivel *SwivelTransactor) ScheduleWithdrawal(opts *bind.TransactOpts, e common.Address) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "scheduleWithdrawal", e)
}

// ScheduleWithdrawal is a paid mutator transaction binding the contract method 0xf8eaad35.
//
// Solidity: function scheduleWithdrawal(address e) returns(bool)
func (_Swivel *SwivelSession) ScheduleWithdrawal(e common.Address) (*types.Transaction, error) {
	return _Swivel.Contract.ScheduleWithdrawal(&_Swivel.TransactOpts, e)
}

// ScheduleWithdrawal is a paid mutator transaction binding the contract method 0xf8eaad35.
//
// Solidity: function scheduleWithdrawal(address e) returns(bool)
func (_Swivel *SwivelTransactorSession) ScheduleWithdrawal(e common.Address) (*types.Transaction, error) {
	return _Swivel.Contract.ScheduleWithdrawal(&_Swivel.TransactOpts, e)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address a) returns(bool)
func (_Swivel *SwivelTransactor) SetAdmin(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "setAdmin", a)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address a) returns(bool)
func (_Swivel *SwivelSession) SetAdmin(a common.Address) (*types.Transaction, error) {
	return _Swivel.Contract.SetAdmin(&_Swivel.TransactOpts, a)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address a) returns(bool)
func (_Swivel *SwivelTransactorSession) SetAdmin(a common.Address) (*types.Transaction, error) {
	return _Swivel.Contract.SetAdmin(&_Swivel.TransactOpts, a)
}

// SplitUnderlying is a paid mutator transaction binding the contract method 0x52a4eec6.
//
// Solidity: function splitUnderlying(uint8 p, address u, uint256 m, uint256 a) returns(bool)
func (_Swivel *SwivelTransactor) SplitUnderlying(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "splitUnderlying", p, u, m, a)
}

// SplitUnderlying is a paid mutator transaction binding the contract method 0x52a4eec6.
//
// Solidity: function splitUnderlying(uint8 p, address u, uint256 m, uint256 a) returns(bool)
func (_Swivel *SwivelSession) SplitUnderlying(p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _Swivel.Contract.SplitUnderlying(&_Swivel.TransactOpts, p, u, m, a)
}

// SplitUnderlying is a paid mutator transaction binding the contract method 0x52a4eec6.
//
// Solidity: function splitUnderlying(uint8 p, address u, uint256 m, uint256 a) returns(bool)
func (_Swivel *SwivelTransactorSession) SplitUnderlying(p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _Swivel.Contract.SplitUnderlying(&_Swivel.TransactOpts, p, u, m, a)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address e) returns(bool)
func (_Swivel *SwivelTransactor) Withdraw(opts *bind.TransactOpts, e common.Address) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "withdraw", e)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address e) returns(bool)
func (_Swivel *SwivelSession) Withdraw(e common.Address) (*types.Transaction, error) {
	return _Swivel.Contract.Withdraw(&_Swivel.TransactOpts, e)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address e) returns(bool)
func (_Swivel *SwivelTransactorSession) Withdraw(e common.Address) (*types.Transaction, error) {
	return _Swivel.Contract.Withdraw(&_Swivel.TransactOpts, e)
}

// SwivelBlockApprovalIterator is returned from FilterBlockApproval and is used to iterate over the raw logs and unpacked data for BlockApproval events raised by the Swivel contract.
type SwivelBlockApprovalIterator struct {
	Event *SwivelBlockApproval // Event containing the contract specifics and raw log

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
func (it *SwivelBlockApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwivelBlockApproval)
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
		it.Event = new(SwivelBlockApproval)
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
func (it *SwivelBlockApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwivelBlockApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwivelBlockApproval represents a BlockApproval event raised by the Swivel contract.
type SwivelBlockApproval struct {
	Token   common.Address
	Blocked common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBlockApproval is a free log retrieval operation binding the contract event 0xe9ef359316c4387875730b408d25b96f623198f8e17f3340089b491f24a16319.
//
// Solidity: event BlockApproval(address indexed token, address indexed blocked)
func (_Swivel *SwivelFilterer) FilterBlockApproval(opts *bind.FilterOpts, token []common.Address, blocked []common.Address) (*SwivelBlockApprovalIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var blockedRule []interface{}
	for _, blockedItem := range blocked {
		blockedRule = append(blockedRule, blockedItem)
	}

	logs, sub, err := _Swivel.contract.FilterLogs(opts, "BlockApproval", tokenRule, blockedRule)
	if err != nil {
		return nil, err
	}
	return &SwivelBlockApprovalIterator{contract: _Swivel.contract, event: "BlockApproval", logs: logs, sub: sub}, nil
}

// WatchBlockApproval is a free log subscription operation binding the contract event 0xe9ef359316c4387875730b408d25b96f623198f8e17f3340089b491f24a16319.
//
// Solidity: event BlockApproval(address indexed token, address indexed blocked)
func (_Swivel *SwivelFilterer) WatchBlockApproval(opts *bind.WatchOpts, sink chan<- *SwivelBlockApproval, token []common.Address, blocked []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var blockedRule []interface{}
	for _, blockedItem := range blocked {
		blockedRule = append(blockedRule, blockedItem)
	}

	logs, sub, err := _Swivel.contract.WatchLogs(opts, "BlockApproval", tokenRule, blockedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwivelBlockApproval)
				if err := _Swivel.contract.UnpackLog(event, "BlockApproval", log); err != nil {
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

// ParseBlockApproval is a log parse operation binding the contract event 0xe9ef359316c4387875730b408d25b96f623198f8e17f3340089b491f24a16319.
//
// Solidity: event BlockApproval(address indexed token, address indexed blocked)
func (_Swivel *SwivelFilterer) ParseBlockApproval(log types.Log) (*SwivelBlockApproval, error) {
	event := new(SwivelBlockApproval)
	if err := _Swivel.contract.UnpackLog(event, "BlockApproval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwivelBlockFeeChangeIterator is returned from FilterBlockFeeChange and is used to iterate over the raw logs and unpacked data for BlockFeeChange events raised by the Swivel contract.
type SwivelBlockFeeChangeIterator struct {
	Event *SwivelBlockFeeChange // Event containing the contract specifics and raw log

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
func (it *SwivelBlockFeeChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwivelBlockFeeChange)
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
		it.Event = new(SwivelBlockFeeChange)
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
func (it *SwivelBlockFeeChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwivelBlockFeeChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwivelBlockFeeChange represents a BlockFeeChange event raised by the Swivel contract.
type SwivelBlockFeeChange struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBlockFeeChange is a free log retrieval operation binding the contract event 0x6875685eb5dbc8e2796d75d2dc9e9cb607b610d0558ee7336df418a26d4846e8.
//
// Solidity: event BlockFeeChange()
func (_Swivel *SwivelFilterer) FilterBlockFeeChange(opts *bind.FilterOpts) (*SwivelBlockFeeChangeIterator, error) {

	logs, sub, err := _Swivel.contract.FilterLogs(opts, "BlockFeeChange")
	if err != nil {
		return nil, err
	}
	return &SwivelBlockFeeChangeIterator{contract: _Swivel.contract, event: "BlockFeeChange", logs: logs, sub: sub}, nil
}

// WatchBlockFeeChange is a free log subscription operation binding the contract event 0x6875685eb5dbc8e2796d75d2dc9e9cb607b610d0558ee7336df418a26d4846e8.
//
// Solidity: event BlockFeeChange()
func (_Swivel *SwivelFilterer) WatchBlockFeeChange(opts *bind.WatchOpts, sink chan<- *SwivelBlockFeeChange) (event.Subscription, error) {

	logs, sub, err := _Swivel.contract.WatchLogs(opts, "BlockFeeChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwivelBlockFeeChange)
				if err := _Swivel.contract.UnpackLog(event, "BlockFeeChange", log); err != nil {
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

// ParseBlockFeeChange is a log parse operation binding the contract event 0x6875685eb5dbc8e2796d75d2dc9e9cb607b610d0558ee7336df418a26d4846e8.
//
// Solidity: event BlockFeeChange()
func (_Swivel *SwivelFilterer) ParseBlockFeeChange(log types.Log) (*SwivelBlockFeeChange, error) {
	event := new(SwivelBlockFeeChange)
	if err := _Swivel.contract.UnpackLog(event, "BlockFeeChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwivelBlockWithdrawalIterator is returned from FilterBlockWithdrawal and is used to iterate over the raw logs and unpacked data for BlockWithdrawal events raised by the Swivel contract.
type SwivelBlockWithdrawalIterator struct {
	Event *SwivelBlockWithdrawal // Event containing the contract specifics and raw log

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
func (it *SwivelBlockWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwivelBlockWithdrawal)
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
		it.Event = new(SwivelBlockWithdrawal)
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
func (it *SwivelBlockWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwivelBlockWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwivelBlockWithdrawal represents a BlockWithdrawal event raised by the Swivel contract.
type SwivelBlockWithdrawal struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBlockWithdrawal is a free log retrieval operation binding the contract event 0xb1c1232c5dd039bb1c46cc05eaf25828e4f8596b7f68bdb23073ba78b9ca382d.
//
// Solidity: event BlockWithdrawal(address indexed token)
func (_Swivel *SwivelFilterer) FilterBlockWithdrawal(opts *bind.FilterOpts, token []common.Address) (*SwivelBlockWithdrawalIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Swivel.contract.FilterLogs(opts, "BlockWithdrawal", tokenRule)
	if err != nil {
		return nil, err
	}
	return &SwivelBlockWithdrawalIterator{contract: _Swivel.contract, event: "BlockWithdrawal", logs: logs, sub: sub}, nil
}

// WatchBlockWithdrawal is a free log subscription operation binding the contract event 0xb1c1232c5dd039bb1c46cc05eaf25828e4f8596b7f68bdb23073ba78b9ca382d.
//
// Solidity: event BlockWithdrawal(address indexed token)
func (_Swivel *SwivelFilterer) WatchBlockWithdrawal(opts *bind.WatchOpts, sink chan<- *SwivelBlockWithdrawal, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Swivel.contract.WatchLogs(opts, "BlockWithdrawal", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwivelBlockWithdrawal)
				if err := _Swivel.contract.UnpackLog(event, "BlockWithdrawal", log); err != nil {
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

// ParseBlockWithdrawal is a log parse operation binding the contract event 0xb1c1232c5dd039bb1c46cc05eaf25828e4f8596b7f68bdb23073ba78b9ca382d.
//
// Solidity: event BlockWithdrawal(address indexed token)
func (_Swivel *SwivelFilterer) ParseBlockWithdrawal(log types.Log) (*SwivelBlockWithdrawal, error) {
	event := new(SwivelBlockWithdrawal)
	if err := _Swivel.contract.UnpackLog(event, "BlockWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwivelCancelIterator is returned from FilterCancel and is used to iterate over the raw logs and unpacked data for Cancel events raised by the Swivel contract.
type SwivelCancelIterator struct {
	Event *SwivelCancel // Event containing the contract specifics and raw log

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
func (it *SwivelCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwivelCancel)
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
		it.Event = new(SwivelCancel)
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
func (it *SwivelCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwivelCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwivelCancel represents a Cancel event raised by the Swivel contract.
type SwivelCancel struct {
	Key  [32]byte
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCancel is a free log retrieval operation binding the contract event 0x9e5d8891dc1b047de610617bc9bc2d8ccffebbc3d63363431a546831245858a6.
//
// Solidity: event Cancel(bytes32 indexed key, bytes32 hash)
func (_Swivel *SwivelFilterer) FilterCancel(opts *bind.FilterOpts, key [][32]byte) (*SwivelCancelIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Swivel.contract.FilterLogs(opts, "Cancel", keyRule)
	if err != nil {
		return nil, err
	}
	return &SwivelCancelIterator{contract: _Swivel.contract, event: "Cancel", logs: logs, sub: sub}, nil
}

// WatchCancel is a free log subscription operation binding the contract event 0x9e5d8891dc1b047de610617bc9bc2d8ccffebbc3d63363431a546831245858a6.
//
// Solidity: event Cancel(bytes32 indexed key, bytes32 hash)
func (_Swivel *SwivelFilterer) WatchCancel(opts *bind.WatchOpts, sink chan<- *SwivelCancel, key [][32]byte) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Swivel.contract.WatchLogs(opts, "Cancel", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwivelCancel)
				if err := _Swivel.contract.UnpackLog(event, "Cancel", log); err != nil {
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

// ParseCancel is a log parse operation binding the contract event 0x9e5d8891dc1b047de610617bc9bc2d8ccffebbc3d63363431a546831245858a6.
//
// Solidity: event Cancel(bytes32 indexed key, bytes32 hash)
func (_Swivel *SwivelFilterer) ParseCancel(log types.Log) (*SwivelCancel, error) {
	event := new(SwivelCancel)
	if err := _Swivel.contract.UnpackLog(event, "Cancel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwivelChangeFeeIterator is returned from FilterChangeFee and is used to iterate over the raw logs and unpacked data for ChangeFee events raised by the Swivel contract.
type SwivelChangeFeeIterator struct {
	Event *SwivelChangeFee // Event containing the contract specifics and raw log

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
func (it *SwivelChangeFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwivelChangeFee)
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
		it.Event = new(SwivelChangeFee)
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
func (it *SwivelChangeFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwivelChangeFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwivelChangeFee represents a ChangeFee event raised by the Swivel contract.
type SwivelChangeFee struct {
	Index *big.Int
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterChangeFee is a free log retrieval operation binding the contract event 0x91e72fa36e0202be93e86c97a3d3d3497cf0a06cf859b14b616a304367835a8e.
//
// Solidity: event ChangeFee(uint256 indexed index, uint256 indexed value)
func (_Swivel *SwivelFilterer) FilterChangeFee(opts *bind.FilterOpts, index []*big.Int, value []*big.Int) (*SwivelChangeFeeIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Swivel.contract.FilterLogs(opts, "ChangeFee", indexRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &SwivelChangeFeeIterator{contract: _Swivel.contract, event: "ChangeFee", logs: logs, sub: sub}, nil
}

// WatchChangeFee is a free log subscription operation binding the contract event 0x91e72fa36e0202be93e86c97a3d3d3497cf0a06cf859b14b616a304367835a8e.
//
// Solidity: event ChangeFee(uint256 indexed index, uint256 indexed value)
func (_Swivel *SwivelFilterer) WatchChangeFee(opts *bind.WatchOpts, sink chan<- *SwivelChangeFee, index []*big.Int, value []*big.Int) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Swivel.contract.WatchLogs(opts, "ChangeFee", indexRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwivelChangeFee)
				if err := _Swivel.contract.UnpackLog(event, "ChangeFee", log); err != nil {
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

// ParseChangeFee is a log parse operation binding the contract event 0x91e72fa36e0202be93e86c97a3d3d3497cf0a06cf859b14b616a304367835a8e.
//
// Solidity: event ChangeFee(uint256 indexed index, uint256 indexed value)
func (_Swivel *SwivelFilterer) ParseChangeFee(log types.Log) (*SwivelChangeFee, error) {
	event := new(SwivelChangeFee)
	if err := _Swivel.contract.UnpackLog(event, "ChangeFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwivelExitIterator is returned from FilterExit and is used to iterate over the raw logs and unpacked data for Exit events raised by the Swivel contract.
type SwivelExitIterator struct {
	Event *SwivelExit // Event containing the contract specifics and raw log

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
func (it *SwivelExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwivelExit)
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
		it.Event = new(SwivelExit)
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
func (it *SwivelExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwivelExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwivelExit represents a Exit event raised by the Swivel contract.
type SwivelExit struct {
	Key    [32]byte
	Hash   [32]byte
	Maker  common.Address
	Vault  bool
	Exit   bool
	Sender common.Address
	Amount *big.Int
	Filled *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExit is a free log retrieval operation binding the contract event 0x51cad9177cf46d59109ae978bb3cf5ffed2bb3d53fb3682fa56fbd9266712834.
//
// Solidity: event Exit(bytes32 indexed key, bytes32 hash, address indexed maker, bool vault, bool exit, address indexed sender, uint256 amount, uint256 filled)
func (_Swivel *SwivelFilterer) FilterExit(opts *bind.FilterOpts, key [][32]byte, maker []common.Address, sender []common.Address) (*SwivelExitIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Swivel.contract.FilterLogs(opts, "Exit", keyRule, makerRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SwivelExitIterator{contract: _Swivel.contract, event: "Exit", logs: logs, sub: sub}, nil
}

// WatchExit is a free log subscription operation binding the contract event 0x51cad9177cf46d59109ae978bb3cf5ffed2bb3d53fb3682fa56fbd9266712834.
//
// Solidity: event Exit(bytes32 indexed key, bytes32 hash, address indexed maker, bool vault, bool exit, address indexed sender, uint256 amount, uint256 filled)
func (_Swivel *SwivelFilterer) WatchExit(opts *bind.WatchOpts, sink chan<- *SwivelExit, key [][32]byte, maker []common.Address, sender []common.Address) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Swivel.contract.WatchLogs(opts, "Exit", keyRule, makerRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwivelExit)
				if err := _Swivel.contract.UnpackLog(event, "Exit", log); err != nil {
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

// ParseExit is a log parse operation binding the contract event 0x51cad9177cf46d59109ae978bb3cf5ffed2bb3d53fb3682fa56fbd9266712834.
//
// Solidity: event Exit(bytes32 indexed key, bytes32 hash, address indexed maker, bool vault, bool exit, address indexed sender, uint256 amount, uint256 filled)
func (_Swivel *SwivelFilterer) ParseExit(log types.Log) (*SwivelExit, error) {
	event := new(SwivelExit)
	if err := _Swivel.contract.UnpackLog(event, "Exit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwivelInitiateIterator is returned from FilterInitiate and is used to iterate over the raw logs and unpacked data for Initiate events raised by the Swivel contract.
type SwivelInitiateIterator struct {
	Event *SwivelInitiate // Event containing the contract specifics and raw log

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
func (it *SwivelInitiateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwivelInitiate)
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
		it.Event = new(SwivelInitiate)
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
func (it *SwivelInitiateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwivelInitiateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwivelInitiate represents a Initiate event raised by the Swivel contract.
type SwivelInitiate struct {
	Key    [32]byte
	Hash   [32]byte
	Maker  common.Address
	Vault  bool
	Exit   bool
	Sender common.Address
	Amount *big.Int
	Filled *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInitiate is a free log retrieval operation binding the contract event 0x32bc401d77ffde781b234d480866e0c360e724770a30ea3299309f9171e400ef.
//
// Solidity: event Initiate(bytes32 indexed key, bytes32 hash, address indexed maker, bool vault, bool exit, address indexed sender, uint256 amount, uint256 filled)
func (_Swivel *SwivelFilterer) FilterInitiate(opts *bind.FilterOpts, key [][32]byte, maker []common.Address, sender []common.Address) (*SwivelInitiateIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Swivel.contract.FilterLogs(opts, "Initiate", keyRule, makerRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SwivelInitiateIterator{contract: _Swivel.contract, event: "Initiate", logs: logs, sub: sub}, nil
}

// WatchInitiate is a free log subscription operation binding the contract event 0x32bc401d77ffde781b234d480866e0c360e724770a30ea3299309f9171e400ef.
//
// Solidity: event Initiate(bytes32 indexed key, bytes32 hash, address indexed maker, bool vault, bool exit, address indexed sender, uint256 amount, uint256 filled)
func (_Swivel *SwivelFilterer) WatchInitiate(opts *bind.WatchOpts, sink chan<- *SwivelInitiate, key [][32]byte, maker []common.Address, sender []common.Address) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Swivel.contract.WatchLogs(opts, "Initiate", keyRule, makerRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwivelInitiate)
				if err := _Swivel.contract.UnpackLog(event, "Initiate", log); err != nil {
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

// ParseInitiate is a log parse operation binding the contract event 0x32bc401d77ffde781b234d480866e0c360e724770a30ea3299309f9171e400ef.
//
// Solidity: event Initiate(bytes32 indexed key, bytes32 hash, address indexed maker, bool vault, bool exit, address indexed sender, uint256 amount, uint256 filled)
func (_Swivel *SwivelFilterer) ParseInitiate(log types.Log) (*SwivelInitiate, error) {
	event := new(SwivelInitiate)
	if err := _Swivel.contract.UnpackLog(event, "Initiate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwivelScheduleApprovalIterator is returned from FilterScheduleApproval and is used to iterate over the raw logs and unpacked data for ScheduleApproval events raised by the Swivel contract.
type SwivelScheduleApprovalIterator struct {
	Event *SwivelScheduleApproval // Event containing the contract specifics and raw log

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
func (it *SwivelScheduleApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwivelScheduleApproval)
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
		it.Event = new(SwivelScheduleApproval)
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
func (it *SwivelScheduleApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwivelScheduleApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwivelScheduleApproval represents a ScheduleApproval event raised by the Swivel contract.
type SwivelScheduleApproval struct {
	Token    common.Address
	Approved common.Address
	Hold     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterScheduleApproval is a free log retrieval operation binding the contract event 0x0ebba6ff355b3d0ba132f77cc4b6e7e312fa368ab818ea06d6f4cc788b4b2ceb.
//
// Solidity: event ScheduleApproval(address indexed token, address indexed approved, uint256 hold)
func (_Swivel *SwivelFilterer) FilterScheduleApproval(opts *bind.FilterOpts, token []common.Address, approved []common.Address) (*SwivelScheduleApprovalIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}

	logs, sub, err := _Swivel.contract.FilterLogs(opts, "ScheduleApproval", tokenRule, approvedRule)
	if err != nil {
		return nil, err
	}
	return &SwivelScheduleApprovalIterator{contract: _Swivel.contract, event: "ScheduleApproval", logs: logs, sub: sub}, nil
}

// WatchScheduleApproval is a free log subscription operation binding the contract event 0x0ebba6ff355b3d0ba132f77cc4b6e7e312fa368ab818ea06d6f4cc788b4b2ceb.
//
// Solidity: event ScheduleApproval(address indexed token, address indexed approved, uint256 hold)
func (_Swivel *SwivelFilterer) WatchScheduleApproval(opts *bind.WatchOpts, sink chan<- *SwivelScheduleApproval, token []common.Address, approved []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}

	logs, sub, err := _Swivel.contract.WatchLogs(opts, "ScheduleApproval", tokenRule, approvedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwivelScheduleApproval)
				if err := _Swivel.contract.UnpackLog(event, "ScheduleApproval", log); err != nil {
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

// ParseScheduleApproval is a log parse operation binding the contract event 0x0ebba6ff355b3d0ba132f77cc4b6e7e312fa368ab818ea06d6f4cc788b4b2ceb.
//
// Solidity: event ScheduleApproval(address indexed token, address indexed approved, uint256 hold)
func (_Swivel *SwivelFilterer) ParseScheduleApproval(log types.Log) (*SwivelScheduleApproval, error) {
	event := new(SwivelScheduleApproval)
	if err := _Swivel.contract.UnpackLog(event, "ScheduleApproval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwivelScheduleFeeChangeIterator is returned from FilterScheduleFeeChange and is used to iterate over the raw logs and unpacked data for ScheduleFeeChange events raised by the Swivel contract.
type SwivelScheduleFeeChangeIterator struct {
	Event *SwivelScheduleFeeChange // Event containing the contract specifics and raw log

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
func (it *SwivelScheduleFeeChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwivelScheduleFeeChange)
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
		it.Event = new(SwivelScheduleFeeChange)
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
func (it *SwivelScheduleFeeChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwivelScheduleFeeChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwivelScheduleFeeChange represents a ScheduleFeeChange event raised by the Swivel contract.
type SwivelScheduleFeeChange struct {
	Proposal [4]uint16
	Hold     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterScheduleFeeChange is a free log retrieval operation binding the contract event 0x996dcde4d8c5a2baa561812ac7ceca074aea97d25dcd0d75106a654433c5544a.
//
// Solidity: event ScheduleFeeChange(uint16[4] proposal, uint256 hold)
func (_Swivel *SwivelFilterer) FilterScheduleFeeChange(opts *bind.FilterOpts) (*SwivelScheduleFeeChangeIterator, error) {

	logs, sub, err := _Swivel.contract.FilterLogs(opts, "ScheduleFeeChange")
	if err != nil {
		return nil, err
	}
	return &SwivelScheduleFeeChangeIterator{contract: _Swivel.contract, event: "ScheduleFeeChange", logs: logs, sub: sub}, nil
}

// WatchScheduleFeeChange is a free log subscription operation binding the contract event 0x996dcde4d8c5a2baa561812ac7ceca074aea97d25dcd0d75106a654433c5544a.
//
// Solidity: event ScheduleFeeChange(uint16[4] proposal, uint256 hold)
func (_Swivel *SwivelFilterer) WatchScheduleFeeChange(opts *bind.WatchOpts, sink chan<- *SwivelScheduleFeeChange) (event.Subscription, error) {

	logs, sub, err := _Swivel.contract.WatchLogs(opts, "ScheduleFeeChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwivelScheduleFeeChange)
				if err := _Swivel.contract.UnpackLog(event, "ScheduleFeeChange", log); err != nil {
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

// ParseScheduleFeeChange is a log parse operation binding the contract event 0x996dcde4d8c5a2baa561812ac7ceca074aea97d25dcd0d75106a654433c5544a.
//
// Solidity: event ScheduleFeeChange(uint16[4] proposal, uint256 hold)
func (_Swivel *SwivelFilterer) ParseScheduleFeeChange(log types.Log) (*SwivelScheduleFeeChange, error) {
	event := new(SwivelScheduleFeeChange)
	if err := _Swivel.contract.UnpackLog(event, "ScheduleFeeChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwivelScheduleWithdrawalIterator is returned from FilterScheduleWithdrawal and is used to iterate over the raw logs and unpacked data for ScheduleWithdrawal events raised by the Swivel contract.
type SwivelScheduleWithdrawalIterator struct {
	Event *SwivelScheduleWithdrawal // Event containing the contract specifics and raw log

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
func (it *SwivelScheduleWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwivelScheduleWithdrawal)
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
		it.Event = new(SwivelScheduleWithdrawal)
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
func (it *SwivelScheduleWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwivelScheduleWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwivelScheduleWithdrawal represents a ScheduleWithdrawal event raised by the Swivel contract.
type SwivelScheduleWithdrawal struct {
	Token common.Address
	Hold  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterScheduleWithdrawal is a free log retrieval operation binding the contract event 0xe4b67652e856f57a7747dd2473850ce987087f4b1744a870504f1c047cb56f4f.
//
// Solidity: event ScheduleWithdrawal(address indexed token, uint256 hold)
func (_Swivel *SwivelFilterer) FilterScheduleWithdrawal(opts *bind.FilterOpts, token []common.Address) (*SwivelScheduleWithdrawalIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Swivel.contract.FilterLogs(opts, "ScheduleWithdrawal", tokenRule)
	if err != nil {
		return nil, err
	}
	return &SwivelScheduleWithdrawalIterator{contract: _Swivel.contract, event: "ScheduleWithdrawal", logs: logs, sub: sub}, nil
}

// WatchScheduleWithdrawal is a free log subscription operation binding the contract event 0xe4b67652e856f57a7747dd2473850ce987087f4b1744a870504f1c047cb56f4f.
//
// Solidity: event ScheduleWithdrawal(address indexed token, uint256 hold)
func (_Swivel *SwivelFilterer) WatchScheduleWithdrawal(opts *bind.WatchOpts, sink chan<- *SwivelScheduleWithdrawal, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Swivel.contract.WatchLogs(opts, "ScheduleWithdrawal", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwivelScheduleWithdrawal)
				if err := _Swivel.contract.UnpackLog(event, "ScheduleWithdrawal", log); err != nil {
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

// ParseScheduleWithdrawal is a log parse operation binding the contract event 0xe4b67652e856f57a7747dd2473850ce987087f4b1744a870504f1c047cb56f4f.
//
// Solidity: event ScheduleWithdrawal(address indexed token, uint256 hold)
func (_Swivel *SwivelFilterer) ParseScheduleWithdrawal(log types.Log) (*SwivelScheduleWithdrawal, error) {
	event := new(SwivelScheduleWithdrawal)
	if err := _Swivel.contract.UnpackLog(event, "ScheduleWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwivelSetAdminIterator is returned from FilterSetAdmin and is used to iterate over the raw logs and unpacked data for SetAdmin events raised by the Swivel contract.
type SwivelSetAdminIterator struct {
	Event *SwivelSetAdmin // Event containing the contract specifics and raw log

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
func (it *SwivelSetAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwivelSetAdmin)
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
		it.Event = new(SwivelSetAdmin)
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
func (it *SwivelSetAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwivelSetAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwivelSetAdmin represents a SetAdmin event raised by the Swivel contract.
type SwivelSetAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetAdmin is a free log retrieval operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address indexed admin)
func (_Swivel *SwivelFilterer) FilterSetAdmin(opts *bind.FilterOpts, admin []common.Address) (*SwivelSetAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Swivel.contract.FilterLogs(opts, "SetAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return &SwivelSetAdminIterator{contract: _Swivel.contract, event: "SetAdmin", logs: logs, sub: sub}, nil
}

// WatchSetAdmin is a free log subscription operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address indexed admin)
func (_Swivel *SwivelFilterer) WatchSetAdmin(opts *bind.WatchOpts, sink chan<- *SwivelSetAdmin, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Swivel.contract.WatchLogs(opts, "SetAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwivelSetAdmin)
				if err := _Swivel.contract.UnpackLog(event, "SetAdmin", log); err != nil {
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
func (_Swivel *SwivelFilterer) ParseSetAdmin(log types.Log) (*SwivelSetAdmin, error) {
	event := new(SwivelSetAdmin)
	if err := _Swivel.contract.UnpackLog(event, "SetAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

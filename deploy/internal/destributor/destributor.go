// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package destributor

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

// DestributorABI is the input ABI used to generate the binding from.
const DestributorABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"distribution\",\"type\":\"uint256\"}],\"name\":\"Distribute\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"p\",\"type\":\"bytes32[]\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"claimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"}],\"name\":\"distribute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distribution\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"pause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DestributorBin is the compiled bytecode used for deploying new contracts.
var DestributorBin = "0x60a060405234801561001057600080fd5b5060405161107a38038061107a83398101604081905261002f91610085565b60048054336001600160a01b031990911617905560609190911b6001600160601b03191660805260008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5556100bd565b60008060408385031215610097578182fd5b82516001600160a01b03811681146100ad578283fd5b6020939093015192949293505050565b60805160601c610f916100e9600039600081816101dd015281816105f301526109740152610f916000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c806375829def11610076578063eb0edcfb1161005b578063eb0edcfb14610180578063f851a44014610193578063fc0c546a146101d857600080fd5b806375829def1461015a578063a54ab4571461016d57600080fd5b80633c70b357116100a75780633c70b357146100fe5780635c975abb1461012c5780635ee58efc1461015157600080fd5b806302329a29146100c35780632e7ba6ef146100eb575b600080fd5b6100d66100d1366004610d7a565b6101ff565b60405190151581526020015b60405180910390f35b6100d66100f9366004610de2565b6102d8565b61011e61010c366004610db2565b60006020819052908152604090205481565b6040519081526020016100e2565b6004546100d69074010000000000000000000000000000000000000000900460ff1681565b61011e60035481565b6100d6610168366004610d18565b61073c565b6100d661017b366004610e74565b61080d565b6100d661018e366004610d39565b6108c1565b6004546101b39073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100e2565b6101b37f000000000000000000000000000000000000000000000000000000000000000081565b60045460009073ffffffffffffffffffffffffffffffffffffffff16338114610289576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f73656e646572206d75737420626520617574686f72697a65640000000000000060448201526064015b60405180910390fd5b6004805484151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff9091161790556001915050919050565b60045460009074010000000000000000000000000000000000000000900460ff1615610360576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f636c61696d696e672069732070617573656400000000000000000000000000006044820152606401610280565b61036c8660035461080d565b156103d3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f646973747269627574696f6e20636c61696d65640000000000000000000000006044820152606401610280565b60035460009081526002602052604090205460ff161561044f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f646973747269627574696f6e2063616e63656c6c6564000000000000000000006044820152606401610280565b60408051602081018890527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606088901b1691810191909152605481018590526000906074016040516020818303038152906040528051906020012090506104fa84848080602002602001604051908101604052809392919081815260200183836020028082843760009201829052506003548152602081905260409020549250859150610c169050565b610560576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f696e76616c69642070726f6f66000000000000000000000000000000000000006044820152606401610280565b600061056e61010089610e95565b9050600061057e6101008a610f07565b600354600090815260016020818152604080842087855290915291829020805491841b9091179055517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a81166004830152602482018a90529192507f00000000000000000000000000000000000000000000000000000000000000009091169063a9059cbb90604401602060405180830381600087803b15801561063957600080fd5b505af115801561064d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106719190610d96565b6106d7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f7472616e73666572206661696c65642e000000000000000000000000000000006044820152606401610280565b604080518a815273ffffffffffffffffffffffffffffffffffffffff8a1660208201529081018890527f3ed1528b0fdc7c5207c1bf935e34a667e13656b9ed165260c522be0bc544f3039060600160405180910390a150600198975050505050505050565b60045460009073ffffffffffffffffffffffffffffffffffffffff163381146107c1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f73656e646572206d75737420626520617574686f72697a6564000000000000006044820152606401610280565b6004805473ffffffffffffffffffffffffffffffffffffffff85167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790556001915050919050565b6000808311610878576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f70617373656420696e646578206d757374206265203e203000000000000000006044820152606401610280565b600061088661010085610e95565b9050600061089661010086610f07565b600094855260016020818152604080882095885294905292909420549190931b908116149392505050565b60045460009073ffffffffffffffffffffffffffffffffffffffff16338114610946576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f73656e646572206d75737420626520617574686f72697a6564000000000000006044820152606401610280565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201527f00000000000000000000000000000000000000000000000000000000000000009060009073ffffffffffffffffffffffffffffffffffffffff8316906370a0823190602401602060405180830381600087803b1580156109d257600080fd5b505af11580156109e6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a0a9190610dca565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8981166004830152602482018390529192509083169063a9059cbb90604401602060405180830381600087803b158015610a7e57600080fd5b505af1158015610a92573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ab69190610d96565b506040517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8981166004830152306024830152604482018890528316906323b872dd90606401602060405180830381600087803b158015610b2d57600080fd5b505af1158015610b41573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b659190610d96565b50600354600081815260026020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580610baa81610ea9565b60008181526020819052604081208990556003829055909250610bcd91506101ff565b5060408051878152602081018390527f9d30f704de5f2d6155b872bae8dcfb406d03fd551cbf6dce019b9fc8f56c2033910160405180910390a150600198975050505050505050565b825160009082825b82811015610ce3576000878281518110610c61577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101519050808311610ca3576040805160208101859052908101829052606001604051602081830303815290604052805190602001209250610cd0565b60408051602081018390529081018490526060016040516020818303038152906040528051906020012092505b5080610cdb81610ea9565b915050610c1e565b50909314949350505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610d1357600080fd5b919050565b600060208284031215610d29578081fd5b610d3282610cef565b9392505050565b60008060008060808587031215610d4e578283fd5b610d5785610cef565b9350610d6560208601610cef565b93969395505050506040820135916060013590565b600060208284031215610d8b578081fd5b8135610d3281610f4a565b600060208284031215610da7578081fd5b8151610d3281610f4a565b600060208284031215610dc3578081fd5b5035919050565b600060208284031215610ddb578081fd5b5051919050565b600080600080600060808688031215610df9578081fd5b85359450610e0960208701610cef565b935060408601359250606086013567ffffffffffffffff80821115610e2c578283fd5b818801915088601f830112610e3f578283fd5b813581811115610e4d578384fd5b8960208260051b8501011115610e61578384fd5b9699959850939650602001949392505050565b60008060408385031215610e86578182fd5b50508035926020909101359150565b600082610ea457610ea4610f1b565b500490565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415610f00577f4e487b710000000000000000000000000000000000000000000000000000000081526011600452602481fd5b5060010190565b600082610f1657610f16610f1b565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b8015158114610f5857600080fd5b5056fea26469706673582212204843107ce43bf623ed52096277bc12fadef635253e39b1c07bb3042948ed5a8164736f6c63430008040033"

// DeployDestributor deploys a new Ethereum contract, binding an instance of Destributor to it.
func DeployDestributor(auth *bind.TransactOpts, backend bind.ContractBackend, t common.Address, r [32]byte) (common.Address, *types.Transaction, *Destributor, error) {
	parsed, err := abi.JSON(strings.NewReader(DestributorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DestributorBin), backend, t, r)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Destributor{DestributorCaller: DestributorCaller{contract: contract}, DestributorTransactor: DestributorTransactor{contract: contract}, DestributorFilterer: DestributorFilterer{contract: contract}}, nil
}

// Destributor is an auto generated Go binding around an Ethereum contract.
type Destributor struct {
	DestributorCaller     // Read-only binding to the contract
	DestributorTransactor // Write-only binding to the contract
	DestributorFilterer   // Log filterer for contract events
}

// DestributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type DestributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DestributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DestributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DestributorSession struct {
	Contract     *Destributor      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DestributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DestributorCallerSession struct {
	Contract *DestributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DestributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DestributorTransactorSession struct {
	Contract     *DestributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DestributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type DestributorRaw struct {
	Contract *Destributor // Generic contract binding to access the raw methods on
}

// DestributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DestributorCallerRaw struct {
	Contract *DestributorCaller // Generic read-only contract binding to access the raw methods on
}

// DestributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DestributorTransactorRaw struct {
	Contract *DestributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDestributor creates a new instance of Destributor, bound to a specific deployed contract.
func NewDestributor(address common.Address, backend bind.ContractBackend) (*Destributor, error) {
	contract, err := bindDestributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Destributor{DestributorCaller: DestributorCaller{contract: contract}, DestributorTransactor: DestributorTransactor{contract: contract}, DestributorFilterer: DestributorFilterer{contract: contract}}, nil
}

// NewDestributorCaller creates a new read-only instance of Destributor, bound to a specific deployed contract.
func NewDestributorCaller(address common.Address, caller bind.ContractCaller) (*DestributorCaller, error) {
	contract, err := bindDestributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DestributorCaller{contract: contract}, nil
}

// NewDestributorTransactor creates a new write-only instance of Destributor, bound to a specific deployed contract.
func NewDestributorTransactor(address common.Address, transactor bind.ContractTransactor) (*DestributorTransactor, error) {
	contract, err := bindDestributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DestributorTransactor{contract: contract}, nil
}

// NewDestributorFilterer creates a new log filterer instance of Destributor, bound to a specific deployed contract.
func NewDestributorFilterer(address common.Address, filterer bind.ContractFilterer) (*DestributorFilterer, error) {
	contract, err := bindDestributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DestributorFilterer{contract: contract}, nil
}

// bindDestributor binds a generic wrapper to an already deployed contract.
func bindDestributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DestributorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Destributor *DestributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Destributor.Contract.DestributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Destributor *DestributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destributor.Contract.DestributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Destributor *DestributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Destributor.Contract.DestributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Destributor *DestributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Destributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Destributor *DestributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Destributor *DestributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Destributor.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Destributor *DestributorCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Destributor.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Destributor *DestributorSession) Admin() (common.Address, error) {
	return _Destributor.Contract.Admin(&_Destributor.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Destributor *DestributorCallerSession) Admin() (common.Address, error) {
	return _Destributor.Contract.Admin(&_Destributor.CallOpts)
}

// Claimed is a free data retrieval call binding the contract method 0xa54ab457.
//
// Solidity: function claimed(uint256 i, uint256 d) view returns(bool)
func (_Destributor *DestributorCaller) Claimed(opts *bind.CallOpts, i *big.Int, d *big.Int) (bool, error) {
	var out []interface{}
	err := _Destributor.contract.Call(opts, &out, "claimed", i, d)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Claimed is a free data retrieval call binding the contract method 0xa54ab457.
//
// Solidity: function claimed(uint256 i, uint256 d) view returns(bool)
func (_Destributor *DestributorSession) Claimed(i *big.Int, d *big.Int) (bool, error) {
	return _Destributor.Contract.Claimed(&_Destributor.CallOpts, i, d)
}

// Claimed is a free data retrieval call binding the contract method 0xa54ab457.
//
// Solidity: function claimed(uint256 i, uint256 d) view returns(bool)
func (_Destributor *DestributorCallerSession) Claimed(i *big.Int, d *big.Int) (bool, error) {
	return _Destributor.Contract.Claimed(&_Destributor.CallOpts, i, d)
}

// Distribution is a free data retrieval call binding the contract method 0x5ee58efc.
//
// Solidity: function distribution() view returns(uint256)
func (_Destributor *DestributorCaller) Distribution(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Destributor.contract.Call(opts, &out, "distribution")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Distribution is a free data retrieval call binding the contract method 0x5ee58efc.
//
// Solidity: function distribution() view returns(uint256)
func (_Destributor *DestributorSession) Distribution() (*big.Int, error) {
	return _Destributor.Contract.Distribution(&_Destributor.CallOpts)
}

// Distribution is a free data retrieval call binding the contract method 0x5ee58efc.
//
// Solidity: function distribution() view returns(uint256)
func (_Destributor *DestributorCallerSession) Distribution() (*big.Int, error) {
	return _Destributor.Contract.Distribution(&_Destributor.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x3c70b357.
//
// Solidity: function merkleRoot(uint256 ) view returns(bytes32)
func (_Destributor *DestributorCaller) MerkleRoot(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Destributor.contract.Call(opts, &out, "merkleRoot", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x3c70b357.
//
// Solidity: function merkleRoot(uint256 ) view returns(bytes32)
func (_Destributor *DestributorSession) MerkleRoot(arg0 *big.Int) ([32]byte, error) {
	return _Destributor.Contract.MerkleRoot(&_Destributor.CallOpts, arg0)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x3c70b357.
//
// Solidity: function merkleRoot(uint256 ) view returns(bytes32)
func (_Destributor *DestributorCallerSession) MerkleRoot(arg0 *big.Int) ([32]byte, error) {
	return _Destributor.Contract.MerkleRoot(&_Destributor.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Destributor *DestributorCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Destributor.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Destributor *DestributorSession) Paused() (bool, error) {
	return _Destributor.Contract.Paused(&_Destributor.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Destributor *DestributorCallerSession) Paused() (bool, error) {
	return _Destributor.Contract.Paused(&_Destributor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Destributor *DestributorCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Destributor.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Destributor *DestributorSession) Token() (common.Address, error) {
	return _Destributor.Contract.Token(&_Destributor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Destributor *DestributorCallerSession) Token() (common.Address, error) {
	return _Destributor.Contract.Token(&_Destributor.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x2e7ba6ef.
//
// Solidity: function claim(uint256 i, address o, uint256 a, bytes32[] p) returns(bool)
func (_Destributor *DestributorTransactor) Claim(opts *bind.TransactOpts, i *big.Int, o common.Address, a *big.Int, p [][32]byte) (*types.Transaction, error) {
	return _Destributor.contract.Transact(opts, "claim", i, o, a, p)
}

// Claim is a paid mutator transaction binding the contract method 0x2e7ba6ef.
//
// Solidity: function claim(uint256 i, address o, uint256 a, bytes32[] p) returns(bool)
func (_Destributor *DestributorSession) Claim(i *big.Int, o common.Address, a *big.Int, p [][32]byte) (*types.Transaction, error) {
	return _Destributor.Contract.Claim(&_Destributor.TransactOpts, i, o, a, p)
}

// Claim is a paid mutator transaction binding the contract method 0x2e7ba6ef.
//
// Solidity: function claim(uint256 i, address o, uint256 a, bytes32[] p) returns(bool)
func (_Destributor *DestributorTransactorSession) Claim(i *big.Int, o common.Address, a *big.Int, p [][32]byte) (*types.Transaction, error) {
	return _Destributor.Contract.Claim(&_Destributor.TransactOpts, i, o, a, p)
}

// Distribute is a paid mutator transaction binding the contract method 0xeb0edcfb.
//
// Solidity: function distribute(address f, address t, uint256 a, bytes32 r) returns(bool)
func (_Destributor *DestributorTransactor) Distribute(opts *bind.TransactOpts, f common.Address, t common.Address, a *big.Int, r [32]byte) (*types.Transaction, error) {
	return _Destributor.contract.Transact(opts, "distribute", f, t, a, r)
}

// Distribute is a paid mutator transaction binding the contract method 0xeb0edcfb.
//
// Solidity: function distribute(address f, address t, uint256 a, bytes32 r) returns(bool)
func (_Destributor *DestributorSession) Distribute(f common.Address, t common.Address, a *big.Int, r [32]byte) (*types.Transaction, error) {
	return _Destributor.Contract.Distribute(&_Destributor.TransactOpts, f, t, a, r)
}

// Distribute is a paid mutator transaction binding the contract method 0xeb0edcfb.
//
// Solidity: function distribute(address f, address t, uint256 a, bytes32 r) returns(bool)
func (_Destributor *DestributorTransactorSession) Distribute(f common.Address, t common.Address, a *big.Int, r [32]byte) (*types.Transaction, error) {
	return _Destributor.Contract.Distribute(&_Destributor.TransactOpts, f, t, a, r)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool b) returns(bool)
func (_Destributor *DestributorTransactor) Pause(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _Destributor.contract.Transact(opts, "pause", b)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool b) returns(bool)
func (_Destributor *DestributorSession) Pause(b bool) (*types.Transaction, error) {
	return _Destributor.Contract.Pause(&_Destributor.TransactOpts, b)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool b) returns(bool)
func (_Destributor *DestributorTransactorSession) Pause(b bool) (*types.Transaction, error) {
	return _Destributor.Contract.Pause(&_Destributor.TransactOpts, b)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address a) returns(bool)
func (_Destributor *DestributorTransactor) TransferAdmin(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Destributor.contract.Transact(opts, "transferAdmin", a)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address a) returns(bool)
func (_Destributor *DestributorSession) TransferAdmin(a common.Address) (*types.Transaction, error) {
	return _Destributor.Contract.TransferAdmin(&_Destributor.TransactOpts, a)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address a) returns(bool)
func (_Destributor *DestributorTransactorSession) TransferAdmin(a common.Address) (*types.Transaction, error) {
	return _Destributor.Contract.TransferAdmin(&_Destributor.TransactOpts, a)
}

// DestributorClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Destributor contract.
type DestributorClaimIterator struct {
	Event *DestributorClaim // Event containing the contract specifics and raw log

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
func (it *DestributorClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestributorClaim)
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
		it.Event = new(DestributorClaim)
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
func (it *DestributorClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestributorClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestributorClaim represents a Claim event raised by the Destributor contract.
type DestributorClaim struct {
	Index  *big.Int
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x3ed1528b0fdc7c5207c1bf935e34a667e13656b9ed165260c522be0bc544f303.
//
// Solidity: event Claim(uint256 index, address owner, uint256 amount)
func (_Destributor *DestributorFilterer) FilterClaim(opts *bind.FilterOpts) (*DestributorClaimIterator, error) {

	logs, sub, err := _Destributor.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &DestributorClaimIterator{contract: _Destributor.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x3ed1528b0fdc7c5207c1bf935e34a667e13656b9ed165260c522be0bc544f303.
//
// Solidity: event Claim(uint256 index, address owner, uint256 amount)
func (_Destributor *DestributorFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *DestributorClaim) (event.Subscription, error) {

	logs, sub, err := _Destributor.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestributorClaim)
				if err := _Destributor.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x3ed1528b0fdc7c5207c1bf935e34a667e13656b9ed165260c522be0bc544f303.
//
// Solidity: event Claim(uint256 index, address owner, uint256 amount)
func (_Destributor *DestributorFilterer) ParseClaim(log types.Log) (*DestributorClaim, error) {
	event := new(DestributorClaim)
	if err := _Destributor.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestributorDistributeIterator is returned from FilterDistribute and is used to iterate over the raw logs and unpacked data for Distribute events raised by the Destributor contract.
type DestributorDistributeIterator struct {
	Event *DestributorDistribute // Event containing the contract specifics and raw log

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
func (it *DestributorDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestributorDistribute)
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
		it.Event = new(DestributorDistribute)
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
func (it *DestributorDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestributorDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestributorDistribute represents a Distribute event raised by the Destributor contract.
type DestributorDistribute struct {
	MerkleRoot   [32]byte
	Distribution *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDistribute is a free log retrieval operation binding the contract event 0x9d30f704de5f2d6155b872bae8dcfb406d03fd551cbf6dce019b9fc8f56c2033.
//
// Solidity: event Distribute(bytes32 merkleRoot, uint256 distribution)
func (_Destributor *DestributorFilterer) FilterDistribute(opts *bind.FilterOpts) (*DestributorDistributeIterator, error) {

	logs, sub, err := _Destributor.contract.FilterLogs(opts, "Distribute")
	if err != nil {
		return nil, err
	}
	return &DestributorDistributeIterator{contract: _Destributor.contract, event: "Distribute", logs: logs, sub: sub}, nil
}

// WatchDistribute is a free log subscription operation binding the contract event 0x9d30f704de5f2d6155b872bae8dcfb406d03fd551cbf6dce019b9fc8f56c2033.
//
// Solidity: event Distribute(bytes32 merkleRoot, uint256 distribution)
func (_Destributor *DestributorFilterer) WatchDistribute(opts *bind.WatchOpts, sink chan<- *DestributorDistribute) (event.Subscription, error) {

	logs, sub, err := _Destributor.contract.WatchLogs(opts, "Distribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestributorDistribute)
				if err := _Destributor.contract.UnpackLog(event, "Distribute", log); err != nil {
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

// ParseDistribute is a log parse operation binding the contract event 0x9d30f704de5f2d6155b872bae8dcfb406d03fd551cbf6dce019b9fc8f56c2033.
//
// Solidity: event Distribute(bytes32 merkleRoot, uint256 distribution)
func (_Destributor *DestributorFilterer) ParseDistribute(log types.Log) (*DestributorDistribute, error) {
	event := new(DestributorDistribute)
	if err := _Destributor.contract.UnpackLog(event, "Distribute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"strings"

	"github.com/wukongcheng/databank/accounts/abi"
	"github.com/wukongcheng/databank/accounts/abi/bind"
	"github.com/wukongcheng/databank/common"
	"github.com/wukongcheng/databank/core/types"
)

// WhitelistABI is the input ABI used to generate the binding from.
const WhitelistABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"enode\",\"type\":\"string\"},{\"name\":\"DIDJson\",\"type\":\"string\"}],\"name\":\"addNewNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"enode\",\"type\":\"string\"}],\"name\":\"inWhiteList\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"enode\",\"type\":\"string\"}],\"name\":\"getDID\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// WhitelistBin is the compiled bytecode used for deploying new contracts.
const WhitelistBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a0319909116179055600180556104a18061003f6000396000f3006060604052600436106100565763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663760ca6fc811461005b578063af657f3314610087578063e0783898146100ec575b600080fd5b341561006657600080fd5b61008560246004803582810192908201359181359182019101356101b4565b005b341561009257600080fd5b6100d860046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061023195505050505050565b604051901515815260200160405180910390f35b34156100f757600080fd5b61013d60046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506102a295505050505050565b60405160208082528190810183818151815260200191508051906020019080838360005b83811015610179578082015183820152602001610161565b50505050905090810190601f1680156101a65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116146101dc57600080fd5b60015460028585604051808383808284378201915050925050509081526020016040519081900390205560015460009081526003602052604090206102229083836103c8565b50506001805481019055505050565b6000806002836040518082805190602001908083835b602083106102665780518252601f199092019160209182019101610247565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902054119050919050565b6102aa610446565b60006002836040518082805190602001908083835b602083106102de5780518252601f1990920191602091820191016102bf565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020549050600360008281526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103bb5780601f10610390576101008083540402835291602001916103bb565b820191906000526020600020905b81548152906001019060200180831161039e57829003601f168201915b5050505050915050919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106104095782800160ff19823516178555610436565b82800160010185558215610436579182015b8281111561043657823582559160200191906001019061041b565b50610442929150610458565b5090565b60206040519081016040526000815290565b61047291905b80821115610442576000815560010161045e565b905600a165627a7a72305820519fdd565d1780972f8a5fa83be185295bca4c9d13077f530d56cf64f81212150029`

// DeployWhitelist deploys a new Ethereum contract, binding an instance of Whitelist to it.
func DeployWhitelist(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Whitelist, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WhitelistBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Whitelist{WhitelistCaller: WhitelistCaller{contract: contract}, WhitelistTransactor: WhitelistTransactor{contract: contract}, WhitelistFilterer: WhitelistFilterer{contract: contract}}, nil
}

// Whitelist is an auto generated Go binding around an Ethereum contract.
type Whitelist struct {
	WhitelistCaller     // Read-only binding to the contract
	WhitelistTransactor // Write-only binding to the contract
	WhitelistFilterer   // Log filterer for contract events
}

// WhitelistCaller is an auto generated read-only Go binding around an Ethereum contract.
type WhitelistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WhitelistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WhitelistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WhitelistSession struct {
	Contract     *Whitelist        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WhitelistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WhitelistCallerSession struct {
	Contract *WhitelistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// WhitelistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WhitelistTransactorSession struct {
	Contract     *WhitelistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// WhitelistRaw is an auto generated low-level Go binding around an Ethereum contract.
type WhitelistRaw struct {
	Contract *Whitelist // Generic contract binding to access the raw methods on
}

// WhitelistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WhitelistCallerRaw struct {
	Contract *WhitelistCaller // Generic read-only contract binding to access the raw methods on
}

// WhitelistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WhitelistTransactorRaw struct {
	Contract *WhitelistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWhitelist creates a new instance of Whitelist, bound to a specific deployed contract.
func NewWhitelist(address common.Address, backend bind.ContractBackend) (*Whitelist, error) {
	contract, err := bindWhitelist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Whitelist{WhitelistCaller: WhitelistCaller{contract: contract}, WhitelistTransactor: WhitelistTransactor{contract: contract}, WhitelistFilterer: WhitelistFilterer{contract: contract}}, nil
}

// NewWhitelistCaller creates a new read-only instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistCaller(address common.Address, caller bind.ContractCaller) (*WhitelistCaller, error) {
	contract, err := bindWhitelist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistCaller{contract: contract}, nil
}

// NewWhitelistTransactor creates a new write-only instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistTransactor(address common.Address, transactor bind.ContractTransactor) (*WhitelistTransactor, error) {
	contract, err := bindWhitelist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistTransactor{contract: contract}, nil
}

// NewWhitelistFilterer creates a new log filterer instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistFilterer(address common.Address, filterer bind.ContractFilterer) (*WhitelistFilterer, error) {
	contract, err := bindWhitelist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WhitelistFilterer{contract: contract}, nil
}

// bindWhitelist binds a generic wrapper to an already deployed contract.
func bindWhitelist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Whitelist *WhitelistRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Whitelist.Contract.WhitelistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Whitelist *WhitelistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Whitelist.Contract.WhitelistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Whitelist *WhitelistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Whitelist.Contract.WhitelistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Whitelist *WhitelistCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Whitelist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Whitelist *WhitelistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Whitelist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Whitelist *WhitelistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Whitelist.Contract.contract.Transact(opts, method, params...)
}

// GetDID is a free data retrieval call binding the contract method 0xe0783898.
//
// Solidity: function getDID(enode string) constant returns(string)
func (_Whitelist *WhitelistCaller) GetDID(opts *bind.CallOpts, enode string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Whitelist.contract.Call(opts, out, "getDID", enode)
	return *ret0, err
}

// GetDID is a free data retrieval call binding the contract method 0xe0783898.
//
// Solidity: function getDID(enode string) constant returns(string)
func (_Whitelist *WhitelistSession) GetDID(enode string) (string, error) {
	return _Whitelist.Contract.GetDID(&_Whitelist.CallOpts, enode)
}

// GetDID is a free data retrieval call binding the contract method 0xe0783898.
//
// Solidity: function getDID(enode string) constant returns(string)
func (_Whitelist *WhitelistCallerSession) GetDID(enode string) (string, error) {
	return _Whitelist.Contract.GetDID(&_Whitelist.CallOpts, enode)
}

// InWhiteList is a free data retrieval call binding the contract method 0xaf657f33.
//
// Solidity: function inWhiteList(enode string) constant returns(bool)
func (_Whitelist *WhitelistCaller) InWhiteList(opts *bind.CallOpts, enode string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Whitelist.contract.Call(opts, out, "inWhiteList", enode)
	return *ret0, err
}

// InWhiteList is a free data retrieval call binding the contract method 0xaf657f33.
//
// Solidity: function inWhiteList(enode string) constant returns(bool)
func (_Whitelist *WhitelistSession) InWhiteList(enode string) (bool, error) {
	return _Whitelist.Contract.InWhiteList(&_Whitelist.CallOpts, enode)
}

// InWhiteList is a free data retrieval call binding the contract method 0xaf657f33.
//
// Solidity: function inWhiteList(enode string) constant returns(bool)
func (_Whitelist *WhitelistCallerSession) InWhiteList(enode string) (bool, error) {
	return _Whitelist.Contract.InWhiteList(&_Whitelist.CallOpts, enode)
}

// AddNewNode is a paid mutator transaction binding the contract method 0x760ca6fc.
//
// Solidity: function addNewNode(enode string, DIDJson string) returns()
func (_Whitelist *WhitelistTransactor) AddNewNode(opts *bind.TransactOpts, enode string, DIDJson string) (*types.Transaction, error) {
	return _Whitelist.contract.Transact(opts, "addNewNode", enode, DIDJson)
}

// AddNewNode is a paid mutator transaction binding the contract method 0x760ca6fc.
//
// Solidity: function addNewNode(enode string, DIDJson string) returns()
func (_Whitelist *WhitelistSession) AddNewNode(enode string, DIDJson string) (*types.Transaction, error) {
	return _Whitelist.Contract.AddNewNode(&_Whitelist.TransactOpts, enode, DIDJson)
}

// AddNewNode is a paid mutator transaction binding the contract method 0x760ca6fc.
//
// Solidity: function addNewNode(enode string, DIDJson string) returns()
func (_Whitelist *WhitelistTransactorSession) AddNewNode(enode string, DIDJson string) (*types.Transaction, error) {
	return _Whitelist.Contract.AddNewNode(&_Whitelist.TransactOpts, enode, DIDJson)
}

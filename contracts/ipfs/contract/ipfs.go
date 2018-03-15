// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"strings"

	"github.com/xcareteam/xci/accounts/abi"
	"github.com/xcareteam/xci/accounts/abi/bind"
	"github.com/xcareteam/xci/common"
	"github.com/xcareteam/xci/core/types"
)

// IpfsABI is the input ABI used to generate the binding from.
const IpfsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"date\",\"type\":\"string\"},{\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"getIpfsUrl\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"date\",\"type\":\"string\"}],\"name\":\"getFileListByDate\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"date\",\"type\":\"string\"},{\"name\":\"fileName\",\"type\":\"string\"},{\"name\":\"url\",\"type\":\"string\"}],\"name\":\"addNewIpfsUrl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// IpfsBin is the compiled bytecode used for deploying new contracts.
const IpfsBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a0319909116179055610c1f8061003b6000396000f3006060604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166385be4c098114610066578063893d20e814610170578063b85cbc121461019f578063e19c9a04146101f0575b600080fd5b341561007157600080fd5b6100f960046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061022895505050505050565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561013557808201518382015260200161011d565b50505050905090810190601f1680156101625780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561017b57600080fd5b610183610397565b604051600160a060020a03909116815260200160405180910390f35b34156101aa57600080fd5b6100f960046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506103a795505050505050565b34156101fb57600080fd5b61022660246004803582810192908201359181358083019290820135916044359182019101356104cc565b005b610230610adb565b610238610adb565b6102768460408051908101604052600181527f2d00000000000000000000000000000000000000000000000000000000000000602082015285610945565b600160a060020a033316600090815260016020526040908190209192508290518082805190602001908083835b602083106102c25780518252601f1990920191602091820191016102a3565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103895780601f1061035e57610100808354040283529160200191610389565b820191906000526020600020905b81548152906001019060200180831161036c57829003601f168201915b505050505091505092915050565b600054600160a060020a03165b90565b6103af610adb565b600160a060020a03331660009081526002602052604090819020908390518082805190602001908083835b602083106103f95780518252601f1990920191602091820191016103da565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104c05780601f10610495576101008083540402835291602001916104c0565b820191906000526020600020905b8154815290600101906020018083116104a357829003601f168201915b50505050509050919050565b6104d4610adb565b6104dc610adb565b6104e4610adb565b6104ec610adb565b6104f4610adb565b6104fc610adb565b61059b8c8c8080601f016020809104026020016040519081016040528181529291906020840183838082843750604094508493505050505190810160405280600181526020017f2d000000000000000000000000000000000000000000000000000000000000008152508c8c8080601f016020809104026020016040519081016040528181529291906020840183838082843750610945945050505050565b600160a060020a033316600090815260016020526040908190209197508790518082805190602001908083835b602083106105e75780518252601f1990920191602091820191016105c8565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106ae5780601f10610683576101008083540402835291602001916106ae565b820191906000526020600020905b81548152906001019060200180831161069157829003601f168201915b505050505094508493508351156106c157fe5b87876001600033600160a060020a0316600160a060020a03168152602001908152602001600020886040518082805190602001908083835b602083106107185780518252601f1990920191602091820191016106f9565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020610757929091610aed565b50600160a060020a03331660009081526002602052604090819020908d908d90518083838082843782019150509250505090815260200160405180910390208054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561082a5780601f106107ff5761010080835404028352916020019161082a565b820191906000526020600020905b81548152906001019060200180831161080d57829003601f168201915b50505050509250829150815115156108745789898080601f0160208091040260200160405190810160405281815292919060208401838380828437509495506108e6945050505050565b6108e3836040805190810160405280600181526020017f2c000000000000000000000000000000000000000000000000000000000000008152508c8c8080601f016020809104026020016040519081016040528181529291906020840183838082843750610945945050505050565b90505b600160a060020a0333166000908152600260205260409081902082918e908e9051808383808284378201915050925050509081526020016040518091039020908051610936929160200190610b6b565b50505050505050505050505050565b61094d610adb565b610955610adb565b61095d610adb565b610965610adb565b61096d610adb565b610975610adb565b6000808a965089955088945084518651885101016040518059106109965750595b818152601f19601f83011681016020016040529050935083925060009150600090505b8651811015610a12578681815181106109ce57fe5b016020015160f860020a900460f860020a028383806001019450815181106109f257fe5b906020010190600160f860020a031916908160001a9053506001016109b9565b5060005b8551811015610a6f57858181518110610a2b57fe5b016020015160f860020a900460f860020a02838380600101945081518110610a4f57fe5b906020010190600160f860020a031916908160001a905350600101610a16565b5060005b8451811015610acc57848181518110610a8857fe5b016020015160f860020a900460f860020a02838380600101945081518110610aac57fe5b906020010190600160f860020a031916908160001a905350600101610a73565b50909998505050505050505050565b60206040519081016040526000815290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b2e5782800160ff19823516178555610b5b565b82800160010185558215610b5b579182015b82811115610b5b578235825591602001919060010190610b40565b50610b67929150610bd9565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610bac57805160ff1916838001178555610b5b565b82800160010185558215610b5b579182015b82811115610b5b578251825591602001919060010190610bbe565b6103a491905b80821115610b675760008155600101610bdf5600a165627a7a723058207203cb6b15f4356c765f6839e826ac05def2d99b7722d69f6b693e2777a666bf0029`

// DeployIpfs deploys a new Ethereum contract, binding an instance of Ipfs to it.
func DeployIpfs(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ipfs, error) {
	parsed, err := abi.JSON(strings.NewReader(IpfsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IpfsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ipfs{IpfsCaller: IpfsCaller{contract: contract}, IpfsTransactor: IpfsTransactor{contract: contract}, IpfsFilterer: IpfsFilterer{contract: contract}}, nil
}

// Ipfs is an auto generated Go binding around an Ethereum contract.
type Ipfs struct {
	IpfsCaller     // Read-only binding to the contract
	IpfsTransactor // Write-only binding to the contract
	IpfsFilterer   // Log filterer for contract events
}

// IpfsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IpfsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpfsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IpfsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpfsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IpfsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpfsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IpfsSession struct {
	Contract     *Ipfs             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IpfsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IpfsCallerSession struct {
	Contract *IpfsCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IpfsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IpfsTransactorSession struct {
	Contract     *IpfsTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IpfsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IpfsRaw struct {
	Contract *Ipfs // Generic contract binding to access the raw methods on
}

// IpfsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IpfsCallerRaw struct {
	Contract *IpfsCaller // Generic read-only contract binding to access the raw methods on
}

// IpfsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IpfsTransactorRaw struct {
	Contract *IpfsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIpfs creates a new instance of Ipfs, bound to a specific deployed contract.
func NewIpfs(address common.Address, backend bind.ContractBackend) (*Ipfs, error) {
	contract, err := bindIpfs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ipfs{IpfsCaller: IpfsCaller{contract: contract}, IpfsTransactor: IpfsTransactor{contract: contract}, IpfsFilterer: IpfsFilterer{contract: contract}}, nil
}

// NewIpfsCaller creates a new read-only instance of Ipfs, bound to a specific deployed contract.
func NewIpfsCaller(address common.Address, caller bind.ContractCaller) (*IpfsCaller, error) {
	contract, err := bindIpfs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IpfsCaller{contract: contract}, nil
}

// NewIpfsTransactor creates a new write-only instance of Ipfs, bound to a specific deployed contract.
func NewIpfsTransactor(address common.Address, transactor bind.ContractTransactor) (*IpfsTransactor, error) {
	contract, err := bindIpfs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IpfsTransactor{contract: contract}, nil
}

// NewIpfsFilterer creates a new log filterer instance of Ipfs, bound to a specific deployed contract.
func NewIpfsFilterer(address common.Address, filterer bind.ContractFilterer) (*IpfsFilterer, error) {
	contract, err := bindIpfs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IpfsFilterer{contract: contract}, nil
}

// bindIpfs binds a generic wrapper to an already deployed contract.
func bindIpfs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IpfsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipfs *IpfsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ipfs.Contract.IpfsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipfs *IpfsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipfs.Contract.IpfsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipfs *IpfsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipfs.Contract.IpfsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipfs *IpfsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ipfs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipfs *IpfsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipfs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipfs *IpfsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipfs.Contract.contract.Transact(opts, method, params...)
}

// GetFileListByDate is a free data retrieval call binding the contract method 0xb85cbc12.
//
// Solidity: function getFileListByDate(date string) constant returns(string)
func (_Ipfs *IpfsCaller) GetFileListByDate(opts *bind.CallOpts, date string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ipfs.contract.Call(opts, out, "getFileListByDate", date)
	return *ret0, err
}

// GetFileListByDate is a free data retrieval call binding the contract method 0xb85cbc12.
//
// Solidity: function getFileListByDate(date string) constant returns(string)
func (_Ipfs *IpfsSession) GetFileListByDate(date string) (string, error) {
	return _Ipfs.Contract.GetFileListByDate(&_Ipfs.CallOpts, date)
}

// GetFileListByDate is a free data retrieval call binding the contract method 0xb85cbc12.
//
// Solidity: function getFileListByDate(date string) constant returns(string)
func (_Ipfs *IpfsCallerSession) GetFileListByDate(date string) (string, error) {
	return _Ipfs.Contract.GetFileListByDate(&_Ipfs.CallOpts, date)
}

// GetIpfsUrl is a free data retrieval call binding the contract method 0x85be4c09.
//
// Solidity: function getIpfsUrl(date string, fileName string) constant returns(string)
func (_Ipfs *IpfsCaller) GetIpfsUrl(opts *bind.CallOpts, date string, fileName string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ipfs.contract.Call(opts, out, "getIpfsUrl", date, fileName)
	return *ret0, err
}

// GetIpfsUrl is a free data retrieval call binding the contract method 0x85be4c09.
//
// Solidity: function getIpfsUrl(date string, fileName string) constant returns(string)
func (_Ipfs *IpfsSession) GetIpfsUrl(date string, fileName string) (string, error) {
	return _Ipfs.Contract.GetIpfsUrl(&_Ipfs.CallOpts, date, fileName)
}

// GetIpfsUrl is a free data retrieval call binding the contract method 0x85be4c09.
//
// Solidity: function getIpfsUrl(date string, fileName string) constant returns(string)
func (_Ipfs *IpfsCallerSession) GetIpfsUrl(date string, fileName string) (string, error) {
	return _Ipfs.Contract.GetIpfsUrl(&_Ipfs.CallOpts, date, fileName)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Ipfs *IpfsCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ipfs.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Ipfs *IpfsSession) GetOwner() (common.Address, error) {
	return _Ipfs.Contract.GetOwner(&_Ipfs.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Ipfs *IpfsCallerSession) GetOwner() (common.Address, error) {
	return _Ipfs.Contract.GetOwner(&_Ipfs.CallOpts)
}

// AddNewIpfsUrl is a paid mutator transaction binding the contract method 0xe19c9a04.
//
// Solidity: function addNewIpfsUrl(date string, fileName string, url string) returns()
func (_Ipfs *IpfsTransactor) AddNewIpfsUrl(opts *bind.TransactOpts, date string, fileName string, url string) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "addNewIpfsUrl", date, fileName, url)
}

// AddNewIpfsUrl is a paid mutator transaction binding the contract method 0xe19c9a04.
//
// Solidity: function addNewIpfsUrl(date string, fileName string, url string) returns()
func (_Ipfs *IpfsSession) AddNewIpfsUrl(date string, fileName string, url string) (*types.Transaction, error) {
	return _Ipfs.Contract.AddNewIpfsUrl(&_Ipfs.TransactOpts, date, fileName, url)
}

// AddNewIpfsUrl is a paid mutator transaction binding the contract method 0xe19c9a04.
//
// Solidity: function addNewIpfsUrl(date string, fileName string, url string) returns()
func (_Ipfs *IpfsTransactorSession) AddNewIpfsUrl(date string, fileName string, url string) (*types.Transaction, error) {
	return _Ipfs.Contract.AddNewIpfsUrl(&_Ipfs.TransactOpts, date, fileName, url)
}

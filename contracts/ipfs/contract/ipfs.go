// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	"github.com/wukongcheng/databank/accounts/abi"
	"github.com/wukongcheng/databank/accounts/abi/bind"
	"github.com/wukongcheng/databank/common"
	"github.com/wukongcheng/databank/core/types"
)

// IpfsABI is the input ABI used to generate the binding from.
const IpfsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"getIpfsUrl\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getFileQuantity\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getFileNameByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"fileName\",\"type\":\"string\"},{\"name\":\"ipfsUrl\",\"type\":\"string\"}],\"name\":\"addNewIpfsUrl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// IpfsBin is the compiled bytecode used for deploying new contracts.
const IpfsBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a03199091161790556106758061003b6000396000f30060606040526004361061006c5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166315eaebdc81146100715780631b192716146101475780634662959e146101785780635ac072b71461019a578063893d20e8146101c6575b600080fd5b341561007c57600080fd5b6100d060048035600160a060020a03169060446024803590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506101f595505050505050565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561010c5780820151838201526020016100f4565b50505050905090810190601f1680156101395780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561015257600080fd5b610166600160a060020a036004351661031b565b60405190815260200160405180910390f35b341561018357600080fd5b6100d0600160a060020a0360043516602435610336565b34156101a557600080fd5b6101c460246004803582810192908201359181359182019101356103dd565b005b34156101d157600080fd5b6101d961058f565b604051600160a060020a03909116815260200160405180910390f35b6101fd61059f565b600160a060020a03831660009081526001602052604090819020908390518082805190602001908083835b602083106102475780518252601f199092019160209182019101610228565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561030e5780601f106102e35761010080835404028352916020019161030e565b820191906000526020600020905b8154815290600101906020018083116102f157829003601f168201915b5050505050905092915050565b600160a060020a031660009081526003602052604090205490565b61033e61059f565b6002600084600160a060020a0316600160a060020a0316815260200190815260200160002060008381526020019081526020016000208054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561030e5780601f106102e35761010080835404028352916020019161030e565b60006103e761059f565b6103ef61059f565b600160a060020a0333166000908152600360209081526040808320546001909252918290209094509088908890518083838082843782019150509250505090815260200160405180910390208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104cf5780601f106104a4576101008083540402835291602001916104cf565b820191906000526020600020905b8154815290600101906020018083116104b257829003601f168201915b505050505091508190508051156104e257fe5b84846001600033600160a060020a0316600160a060020a031681526020019081526020016000208989604051808383808284378201915050925050509081526020016040519081900390206105389290916105b1565b50600160a060020a033316600090815260026020908152604080832086845290915290206105679088886105b1565b505050600160a060020a03331660009081526003602052604090206001909101905550505050565b600054600160a060020a03165b90565b60206040519081016040526000815290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106105f25782800160ff1982351617855561061f565b8280016001018555821561061f579182015b8281111561061f578235825591602001919060010190610604565b5061062b92915061062f565b5090565b61059c91905b8082111561062b57600081556001016106355600a165627a7a72305820c93eff2ca23455b70eda6c2f33cd5bc4f1708a9db0bec63700eb52490855d27c0029`

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

// GetFileNameByIndex is a free data retrieval call binding the contract method 0x4662959e.
//
// Solidity: function getFileNameByIndex(account address, index uint256) constant returns(string)
func (_Ipfs *IpfsCaller) GetFileNameByIndex(opts *bind.CallOpts, account common.Address, index *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ipfs.contract.Call(opts, out, "getFileNameByIndex", account, index)
	return *ret0, err
}

// GetFileNameByIndex is a free data retrieval call binding the contract method 0x4662959e.
//
// Solidity: function getFileNameByIndex(account address, index uint256) constant returns(string)
func (_Ipfs *IpfsSession) GetFileNameByIndex(account common.Address, index *big.Int) (string, error) {
	return _Ipfs.Contract.GetFileNameByIndex(&_Ipfs.CallOpts, account, index)
}

// GetFileNameByIndex is a free data retrieval call binding the contract method 0x4662959e.
//
// Solidity: function getFileNameByIndex(account address, index uint256) constant returns(string)
func (_Ipfs *IpfsCallerSession) GetFileNameByIndex(account common.Address, index *big.Int) (string, error) {
	return _Ipfs.Contract.GetFileNameByIndex(&_Ipfs.CallOpts, account, index)
}

// GetFileQuantity is a free data retrieval call binding the contract method 0x1b192716.
//
// Solidity: function getFileQuantity(account address) constant returns(uint256)
func (_Ipfs *IpfsCaller) GetFileQuantity(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ipfs.contract.Call(opts, out, "getFileQuantity", account)
	return *ret0, err
}

// GetFileQuantity is a free data retrieval call binding the contract method 0x1b192716.
//
// Solidity: function getFileQuantity(account address) constant returns(uint256)
func (_Ipfs *IpfsSession) GetFileQuantity(account common.Address) (*big.Int, error) {
	return _Ipfs.Contract.GetFileQuantity(&_Ipfs.CallOpts, account)
}

// GetFileQuantity is a free data retrieval call binding the contract method 0x1b192716.
//
// Solidity: function getFileQuantity(account address) constant returns(uint256)
func (_Ipfs *IpfsCallerSession) GetFileQuantity(account common.Address) (*big.Int, error) {
	return _Ipfs.Contract.GetFileQuantity(&_Ipfs.CallOpts, account)
}

// GetIpfsUrl is a free data retrieval call binding the contract method 0x15eaebdc.
//
// Solidity: function getIpfsUrl(account address, fileName string) constant returns(string)
func (_Ipfs *IpfsCaller) GetIpfsUrl(opts *bind.CallOpts, account common.Address, fileName string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ipfs.contract.Call(opts, out, "getIpfsUrl", account, fileName)
	return *ret0, err
}

// GetIpfsUrl is a free data retrieval call binding the contract method 0x15eaebdc.
//
// Solidity: function getIpfsUrl(account address, fileName string) constant returns(string)
func (_Ipfs *IpfsSession) GetIpfsUrl(account common.Address, fileName string) (string, error) {
	return _Ipfs.Contract.GetIpfsUrl(&_Ipfs.CallOpts, account, fileName)
}

// GetIpfsUrl is a free data retrieval call binding the contract method 0x15eaebdc.
//
// Solidity: function getIpfsUrl(account address, fileName string) constant returns(string)
func (_Ipfs *IpfsCallerSession) GetIpfsUrl(account common.Address, fileName string) (string, error) {
	return _Ipfs.Contract.GetIpfsUrl(&_Ipfs.CallOpts, account, fileName)
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

// AddNewIpfsUrl is a paid mutator transaction binding the contract method 0x5ac072b7.
//
// Solidity: function addNewIpfsUrl(fileName string, ipfsUrl string) returns()
func (_Ipfs *IpfsTransactor) AddNewIpfsUrl(opts *bind.TransactOpts, fileName string, ipfsUrl string) (*types.Transaction, error) {
	return _Ipfs.contract.Transact(opts, "addNewIpfsUrl", fileName, ipfsUrl)
}

// AddNewIpfsUrl is a paid mutator transaction binding the contract method 0x5ac072b7.
//
// Solidity: function addNewIpfsUrl(fileName string, ipfsUrl string) returns()
func (_Ipfs *IpfsSession) AddNewIpfsUrl(fileName string, ipfsUrl string) (*types.Transaction, error) {
	return _Ipfs.Contract.AddNewIpfsUrl(&_Ipfs.TransactOpts, fileName, ipfsUrl)
}

// AddNewIpfsUrl is a paid mutator transaction binding the contract method 0x5ac072b7.
//
// Solidity: function addNewIpfsUrl(fileName string, ipfsUrl string) returns()
func (_Ipfs *IpfsTransactorSession) AddNewIpfsUrl(fileName string, ipfsUrl string) (*types.Transaction, error) {
	return _Ipfs.Contract.AddNewIpfsUrl(&_Ipfs.TransactOpts, fileName, ipfsUrl)
}

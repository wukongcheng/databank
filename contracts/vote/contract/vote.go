// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	"github.com/xcareteam/xci/common"
	"github.com/xcareteam/xci/accounts/abi"
	ethereum "github.com/xcareteam/xci"
	"github.com/xcareteam/xci/accounts/abi/bind"
	"github.com/xcareteam/xci/core/types"
	"github.com/xcareteam/xci/event"
)

// BallotManagerABI is the input ABI used to generate the binding from.
const BallotManagerABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"ballotId\",\"type\":\"uint256\"}],\"name\":\"winnerName\",\"outputs\":[{\"name\":\"winnerName_\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_erc20Address\",\"type\":\"address\"}],\"name\":\"setERC20Address\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballotId\",\"type\":\"uint256\"},{\"name\":\"proposal\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"xcareTokenContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ballotId\",\"type\":\"uint256\"}],\"name\":\"winningProposal\",\"outputs\":[{\"name\":\"winningProposal_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ballotname\",\"type\":\"string\"},{\"name\":\"proposalNames\",\"type\":\"bytes32[]\"}],\"name\":\"NewBallot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"ballotId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"ballotName\",\"type\":\"string\"}],\"name\":\"CreateBallot\",\"type\":\"event\"}]"

// BallotManagerBin is the compiled bytecode used for deploying new contracts.
const BallotManagerBin = `0x6060604052341561000f57600080fd5b60028054600160a060020a03191633600160a060020a0316179055610770806100396000396000f3006060604052600436106100775763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663205758ba811461007c57806341bec0d2146100a4578063b384abef146100c5578063bab2cefe146100de578063bcd60f6c1461010d578063db94bd4714610123575b600080fd5b341561008757600080fd5b6100926004356101b4565b60405190815260200160405180910390f35b34156100af57600080fd5b6100c3600160a060020a0360043516610204565b005b34156100d057600080fd5b6100c360043560243561024e565b34156100e957600080fd5b6100f16103b8565b604051600160a060020a03909116815260200160405180910390f35b341561011857600080fd5b6100926004356103c7565b341561012e57600080fd5b6100c360046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437820191505050505050919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284375094965061045d95505050505050565b6000805482106101c357600080fd5b60008054839081106101d157fe5b906000526020600020906004020160020160006101ed846103c7565b815260208101919091526040016000205492915050565b60025433600160a060020a0390811691161461021f57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60008054831061025d57600080fd5b600080548490811061026b57fe5b9060005260206000209060040201600101548210151561028a57600080fd5b600080548490811061029857fe5b60009182526020808320600160a060020a0333168452600360049093020191909101905260409020600181015490915060ff16156102d557600080fd5b6001818101805460ff1916821790556002820183905554600160a060020a03166370a08231336000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561035b57600080fd5b6102c65a03f1151561036c57600080fd5b505050604051805180835560008054919250908590811061038957fe5b600091825260208083209583526002600490920290950101909352604090922060010180549092019091555050565b600154600160a060020a031681565b60008054819081908190819086106103de57600080fd5b600093506000868154811015156103f157fe5b9060005260206000209060040201925082600101549150600090505b818110156104545760008181526002840160205260409020600101548490111561044c5760008181526002840160205260409020600101549094509250835b60010161040d565b50505050919050565b600254600090819033600160a060020a0390811691161461047d57600080fd5b600080546001810161048f8382610607565b916000526020600020906004020160006040805190810160405280888152602001875190529190508151819080516104cb929160200190610638565b50602082015160019091015550506000805460001901925090505b825181101561056157604080519081016040528084838151811061050657fe5b9060200190602002015181526000602090910181905280548490811061052857fe5b600091825260208083208584526002600490930201919091019052604090208151815560208201516001918201559190910190506104e6565b7fbd6b7eff2074354e9b4d638f4aea8af1536168cf8b8f03b6725ff26d8e722fde828560405182815260406020820181815290820183818151815260200191508051906020019080838360005b838110156105c65780820151838201526020016105ae565b50505050905090810190601f1680156105f35780820380516001836020036101000a031916815260200191505b50935050505060405180910390a150505050565b8154818355818115116106335760040281600402836000526020600020918201910161063391906106b6565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061067957805160ff19168380011785556106a6565b828001600101855582156106a6579182015b828111156106a657825182559160200191906001019061068b565b506106b29291506106e3565b5090565b6106e091905b808211156106b25760006106d082826106fd565b50600060018201556004016106bc565b90565b6106e091905b808211156106b257600081556001016106e9565b50805460018160011615610100020316600290046000825580601f106107235750610741565b601f01602090049060005260206000209081019061074191906106e3565b505600a165627a7a7230582001425d2571c12d75e3d2765e8eecc1bad370ca1f4ed302f3438acedfa5326fce0029`

// DeployBallotManager deploys a new Ethereum contract, binding an instance of BallotManager to it.
func DeployBallotManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BallotManager, error) {
	parsed, err := abi.JSON(strings.NewReader(BallotManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BallotManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BallotManager{BallotManagerCaller: BallotManagerCaller{contract: contract}, BallotManagerTransactor: BallotManagerTransactor{contract: contract}, BallotManagerFilterer: BallotManagerFilterer{contract: contract}}, nil
}

// BallotManager is an auto generated Go binding around an Ethereum contract.
type BallotManager struct {
	BallotManagerCaller     // Read-only binding to the contract
	BallotManagerTransactor // Write-only binding to the contract
	BallotManagerFilterer   // Log filterer for contract events
}

// BallotManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BallotManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BallotManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BallotManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BallotManagerSession struct {
	Contract     *BallotManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BallotManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BallotManagerCallerSession struct {
	Contract *BallotManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BallotManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BallotManagerTransactorSession struct {
	Contract     *BallotManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BallotManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BallotManagerRaw struct {
	Contract *BallotManager // Generic contract binding to access the raw methods on
}

// BallotManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BallotManagerCallerRaw struct {
	Contract *BallotManagerCaller // Generic read-only contract binding to access the raw methods on
}

// BallotManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BallotManagerTransactorRaw struct {
	Contract *BallotManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBallotManager creates a new instance of BallotManager, bound to a specific deployed contract.
func NewBallotManager(address common.Address, backend bind.ContractBackend) (*BallotManager, error) {
	contract, err := bindBallotManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BallotManager{BallotManagerCaller: BallotManagerCaller{contract: contract}, BallotManagerTransactor: BallotManagerTransactor{contract: contract}, BallotManagerFilterer: BallotManagerFilterer{contract: contract}}, nil
}

// NewBallotManagerCaller creates a new read-only instance of BallotManager, bound to a specific deployed contract.
func NewBallotManagerCaller(address common.Address, caller bind.ContractCaller) (*BallotManagerCaller, error) {
	contract, err := bindBallotManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BallotManagerCaller{contract: contract}, nil
}

// NewBallotManagerTransactor creates a new write-only instance of BallotManager, bound to a specific deployed contract.
func NewBallotManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*BallotManagerTransactor, error) {
	contract, err := bindBallotManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BallotManagerTransactor{contract: contract}, nil
}

// NewBallotManagerFilterer creates a new log filterer instance of BallotManager, bound to a specific deployed contract.
func NewBallotManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*BallotManagerFilterer, error) {
	contract, err := bindBallotManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BallotManagerFilterer{contract: contract}, nil
}

// bindBallotManager binds a generic wrapper to an already deployed contract.
func bindBallotManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BallotManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BallotManager *BallotManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BallotManager.Contract.BallotManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BallotManager *BallotManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BallotManager.Contract.BallotManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BallotManager *BallotManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BallotManager.Contract.BallotManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BallotManager *BallotManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BallotManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BallotManager *BallotManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BallotManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BallotManager *BallotManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BallotManager.Contract.contract.Transact(opts, method, params...)
}

// WinnerName is a free data retrieval call binding the contract method 0x205758ba.
//
// Solidity: function winnerName(ballotId uint256) constant returns(winnerName_ bytes32)
func (_BallotManager *BallotManagerCaller) WinnerName(opts *bind.CallOpts, ballotId *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BallotManager.contract.Call(opts, out, "winnerName", ballotId)
	return *ret0, err
}

// WinnerName is a free data retrieval call binding the contract method 0x205758ba.
//
// Solidity: function winnerName(ballotId uint256) constant returns(winnerName_ bytes32)
func (_BallotManager *BallotManagerSession) WinnerName(ballotId *big.Int) ([32]byte, error) {
	return _BallotManager.Contract.WinnerName(&_BallotManager.CallOpts, ballotId)
}

// WinnerName is a free data retrieval call binding the contract method 0x205758ba.
//
// Solidity: function winnerName(ballotId uint256) constant returns(winnerName_ bytes32)
func (_BallotManager *BallotManagerCallerSession) WinnerName(ballotId *big.Int) ([32]byte, error) {
	return _BallotManager.Contract.WinnerName(&_BallotManager.CallOpts, ballotId)
}

// WinningProposal is a free data retrieval call binding the contract method 0xbcd60f6c.
//
// Solidity: function winningProposal(ballotId uint256) constant returns(winningProposal_ uint256)
func (_BallotManager *BallotManagerCaller) WinningProposal(opts *bind.CallOpts, ballotId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BallotManager.contract.Call(opts, out, "winningProposal", ballotId)
	return *ret0, err
}

// WinningProposal is a free data retrieval call binding the contract method 0xbcd60f6c.
//
// Solidity: function winningProposal(ballotId uint256) constant returns(winningProposal_ uint256)
func (_BallotManager *BallotManagerSession) WinningProposal(ballotId *big.Int) (*big.Int, error) {
	return _BallotManager.Contract.WinningProposal(&_BallotManager.CallOpts, ballotId)
}

// WinningProposal is a free data retrieval call binding the contract method 0xbcd60f6c.
//
// Solidity: function winningProposal(ballotId uint256) constant returns(winningProposal_ uint256)
func (_BallotManager *BallotManagerCallerSession) WinningProposal(ballotId *big.Int) (*big.Int, error) {
	return _BallotManager.Contract.WinningProposal(&_BallotManager.CallOpts, ballotId)
}

// XcareTokenContract is a free data retrieval call binding the contract method 0xbab2cefe.
//
// Solidity: function xcareTokenContract() constant returns(address)
func (_BallotManager *BallotManagerCaller) XcareTokenContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BallotManager.contract.Call(opts, out, "xcareTokenContract")
	return *ret0, err
}

// XcareTokenContract is a free data retrieval call binding the contract method 0xbab2cefe.
//
// Solidity: function xcareTokenContract() constant returns(address)
func (_BallotManager *BallotManagerSession) XcareTokenContract() (common.Address, error) {
	return _BallotManager.Contract.XcareTokenContract(&_BallotManager.CallOpts)
}

// XcareTokenContract is a free data retrieval call binding the contract method 0xbab2cefe.
//
// Solidity: function xcareTokenContract() constant returns(address)
func (_BallotManager *BallotManagerCallerSession) XcareTokenContract() (common.Address, error) {
	return _BallotManager.Contract.XcareTokenContract(&_BallotManager.CallOpts)
}

// NewBallot is a paid mutator transaction binding the contract method 0xdb94bd47.
//
// Solidity: function NewBallot(ballotname string, proposalNames bytes32[]) returns()
func (_BallotManager *BallotManagerTransactor) NewBallot(opts *bind.TransactOpts, ballotname string, proposalNames [][32]byte) (*types.Transaction, error) {
	return _BallotManager.contract.Transact(opts, "NewBallot", ballotname, proposalNames)
}

// NewBallot is a paid mutator transaction binding the contract method 0xdb94bd47.
//
// Solidity: function NewBallot(ballotname string, proposalNames bytes32[]) returns()
func (_BallotManager *BallotManagerSession) NewBallot(ballotname string, proposalNames [][32]byte) (*types.Transaction, error) {
	return _BallotManager.Contract.NewBallot(&_BallotManager.TransactOpts, ballotname, proposalNames)
}

// NewBallot is a paid mutator transaction binding the contract method 0xdb94bd47.
//
// Solidity: function NewBallot(ballotname string, proposalNames bytes32[]) returns()
func (_BallotManager *BallotManagerTransactorSession) NewBallot(ballotname string, proposalNames [][32]byte) (*types.Transaction, error) {
	return _BallotManager.Contract.NewBallot(&_BallotManager.TransactOpts, ballotname, proposalNames)
}

// SetERC20Address is a paid mutator transaction binding the contract method 0x41bec0d2.
//
// Solidity: function setERC20Address(_erc20Address address) returns()
func (_BallotManager *BallotManagerTransactor) SetERC20Address(opts *bind.TransactOpts, _erc20Address common.Address) (*types.Transaction, error) {
	return _BallotManager.contract.Transact(opts, "setERC20Address", _erc20Address)
}

// SetERC20Address is a paid mutator transaction binding the contract method 0x41bec0d2.
//
// Solidity: function setERC20Address(_erc20Address address) returns()
func (_BallotManager *BallotManagerSession) SetERC20Address(_erc20Address common.Address) (*types.Transaction, error) {
	return _BallotManager.Contract.SetERC20Address(&_BallotManager.TransactOpts, _erc20Address)
}

// SetERC20Address is a paid mutator transaction binding the contract method 0x41bec0d2.
//
// Solidity: function setERC20Address(_erc20Address address) returns()
func (_BallotManager *BallotManagerTransactorSession) SetERC20Address(_erc20Address common.Address) (*types.Transaction, error) {
	return _BallotManager.Contract.SetERC20Address(&_BallotManager.TransactOpts, _erc20Address)
}

// Vote is a paid mutator transaction binding the contract method 0xb384abef.
//
// Solidity: function vote(ballotId uint256, proposal uint256) returns()
func (_BallotManager *BallotManagerTransactor) Vote(opts *bind.TransactOpts, ballotId *big.Int, proposal *big.Int) (*types.Transaction, error) {
	return _BallotManager.contract.Transact(opts, "vote", ballotId, proposal)
}

// Vote is a paid mutator transaction binding the contract method 0xb384abef.
//
// Solidity: function vote(ballotId uint256, proposal uint256) returns()
func (_BallotManager *BallotManagerSession) Vote(ballotId *big.Int, proposal *big.Int) (*types.Transaction, error) {
	return _BallotManager.Contract.Vote(&_BallotManager.TransactOpts, ballotId, proposal)
}

// Vote is a paid mutator transaction binding the contract method 0xb384abef.
//
// Solidity: function vote(ballotId uint256, proposal uint256) returns()
func (_BallotManager *BallotManagerTransactorSession) Vote(ballotId *big.Int, proposal *big.Int) (*types.Transaction, error) {
	return _BallotManager.Contract.Vote(&_BallotManager.TransactOpts, ballotId, proposal)
}

// BallotManagerCreateBallotIterator is returned from FilterCreateBallot and is used to iterate over the raw logs and unpacked data for CreateBallot events raised by the BallotManager contract.
type BallotManagerCreateBallotIterator struct {
	Event *BallotManagerCreateBallot // Event containing the contract specifics and raw log

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
func (it *BallotManagerCreateBallotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BallotManagerCreateBallot)
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
		it.Event = new(BallotManagerCreateBallot)
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
func (it *BallotManagerCreateBallotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BallotManagerCreateBallotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BallotManagerCreateBallot represents a CreateBallot event raised by the BallotManager contract.
type BallotManagerCreateBallot struct {
	BallotId   *big.Int
	BallotName string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreateBallot is a free log retrieval operation binding the contract event 0xbd6b7eff2074354e9b4d638f4aea8af1536168cf8b8f03b6725ff26d8e722fde.
//
// Solidity: event CreateBallot(ballotId uint256, ballotName string)
func (_BallotManager *BallotManagerFilterer) FilterCreateBallot(opts *bind.FilterOpts) (*BallotManagerCreateBallotIterator, error) {

	logs, sub, err := _BallotManager.contract.FilterLogs(opts, "CreateBallot")
	if err != nil {
		return nil, err
	}
	return &BallotManagerCreateBallotIterator{contract: _BallotManager.contract, event: "CreateBallot", logs: logs, sub: sub}, nil
}

// WatchCreateBallot is a free log subscription operation binding the contract event 0xbd6b7eff2074354e9b4d638f4aea8af1536168cf8b8f03b6725ff26d8e722fde.
//
// Solidity: event CreateBallot(ballotId uint256, ballotName string)
func (_BallotManager *BallotManagerFilterer) WatchCreateBallot(opts *bind.WatchOpts, sink chan<- *BallotManagerCreateBallot) (event.Subscription, error) {

	logs, sub, err := _BallotManager.contract.WatchLogs(opts, "CreateBallot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BallotManagerCreateBallot)
				if err := _BallotManager.contract.UnpackLog(event, "CreateBallot", log); err != nil {
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

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"keys\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"key\",\"type\":\"string\"}],\"name\":\"Register\",\"type\":\"event\"}]"

// ERC20Bin is the compiled bytecode used for deploying new contracts.
const ERC20Bin = `0x`

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20 *ERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20 *ERC20Session) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20 *ERC20CallerSession) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Keys is a free data retrieval call binding the contract method 0x670d14b2.
//
// Solidity: function keys( address) constant returns(string)
func (_ERC20 *ERC20Caller) Keys(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "keys", arg0)
	return *ret0, err
}

// Keys is a free data retrieval call binding the contract method 0x670d14b2.
//
// Solidity: function keys( address) constant returns(string)
func (_ERC20 *ERC20Session) Keys(arg0 common.Address) (string, error) {
	return _ERC20.Contract.Keys(&_ERC20.CallOpts, arg0)
}

// Keys is a free data retrieval call binding the contract method 0x670d14b2.
//
// Solidity: function keys( address) constant returns(string)
func (_ERC20 *ERC20CallerSession) Keys(arg0 common.Address) (string, error) {
	return _ERC20.Contract.Keys(&_ERC20.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC20 *ERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC20 *ERC20Session) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC20 *ERC20CallerSession) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC20 *ERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC20 *ERC20Session) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC20 *ERC20CallerSession) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_ERC20 *ERC20Session) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_ERC20 *ERC20TransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, _spender, _value)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(key string) returns()
func (_ERC20 *ERC20Transactor) Register(opts *bind.TransactOpts, key string) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "register", key)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(key string) returns()
func (_ERC20 *ERC20Session) Register(key string) (*types.Transaction, error) {
	return _ERC20.Contract.Register(&_ERC20.TransactOpts, key)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(key string) returns()
func (_ERC20 *ERC20TransactorSession) Register(key string) (*types.Transaction, error) {
	return _ERC20.Contract.Register(&_ERC20.TransactOpts, key)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns()
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns()
func (_ERC20 *ERC20Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns()
func (_ERC20 *ERC20TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_ERC20 *ERC20Session) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, _from, _to, _value)
}

// ERC20RegisterIterator is returned from FilterRegister and is used to iterate over the raw logs and unpacked data for Register events raised by the ERC20 contract.
type ERC20RegisterIterator struct {
	Event *ERC20Register // Event containing the contract specifics and raw log

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
func (it *ERC20RegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Register)
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
		it.Event = new(ERC20Register)
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
func (it *ERC20RegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20RegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Register represents a Register event raised by the ERC20 contract.
type ERC20Register struct {
	User common.Address
	Key  string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRegister is a free log retrieval operation binding the contract event 0x6ba0831d2f62ae5cbf7214bcc1d79c5da1d705f12811efda0beaa840006f874e.
//
// Solidity: event Register(user address, key string)
func (_ERC20 *ERC20Filterer) FilterRegister(opts *bind.FilterOpts) (*ERC20RegisterIterator, error) {

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Register")
	if err != nil {
		return nil, err
	}
	return &ERC20RegisterIterator{contract: _ERC20.contract, event: "Register", logs: logs, sub: sub}, nil
}

// WatchRegister is a free log subscription operation binding the contract event 0x6ba0831d2f62ae5cbf7214bcc1d79c5da1d705f12811efda0beaa840006f874e.
//
// Solidity: event Register(user address, key string)
func (_ERC20 *ERC20Filterer) WatchRegister(opts *bind.WatchOpts, sink chan<- *ERC20Register) (event.Subscription, error) {

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Register")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Register)
				if err := _ERC20.contract.UnpackLog(event, "Register", log); err != nil {
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

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

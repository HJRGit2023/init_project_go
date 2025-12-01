// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aWork

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
	_ = abi.ConvertType
)

// AWorkMetaData contains all meta data concerning the AWork contract.
var AWorkMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"by\",\"type\":\"uint256\"}],\"name\":\"Increment\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counterAdd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506101eb8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c806306661abd146100385780633f31a3c214610056575b5f5ffd5b610040610060565b60405161004d91906100cd565b60405180910390f35b61005e610065565b005b5f5481565b5f5f81548092919061007690610113565b91905055507f51af157c2eee40f68107a47a49c32fbbeb0a3c9e5cd37aa56e88e6be92368a8160016040516100ab919061019c565b60405180910390a1565b5f819050919050565b6100c7816100b5565b82525050565b5f6020820190506100e05f8301846100be565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61011d826100b5565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361014f5761014e6100e6565b5b600182019050919050565b5f819050919050565b5f819050919050565b5f61018661018161017c8461015a565b610163565b6100b5565b9050919050565b6101968161016c565b82525050565b5f6020820190506101af5f83018461018d565b9291505056fea26469706673582212206c7c724a55f6f8fcc79eeed62f1acd8d26fb1062200403e6f3fc1fca8b9d228d64736f6c634300081e0033",
}

// AWorkABI is the input ABI used to generate the binding from.
// Deprecated: Use AWorkMetaData.ABI instead.
var AWorkABI = AWorkMetaData.ABI

// AWorkBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AWorkMetaData.Bin instead.
var AWorkBin = AWorkMetaData.Bin

// DeployAWork deploys a new Ethereum contract, binding an instance of AWork to it.
func DeployAWork(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AWork, error) {
	parsed, err := AWorkMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AWorkBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AWork{AWorkCaller: AWorkCaller{contract: contract}, AWorkTransactor: AWorkTransactor{contract: contract}, AWorkFilterer: AWorkFilterer{contract: contract}}, nil
}

// AWork is an auto generated Go binding around an Ethereum contract.
type AWork struct {
	AWorkCaller     // Read-only binding to the contract
	AWorkTransactor // Write-only binding to the contract
	AWorkFilterer   // Log filterer for contract events
}

// AWorkCaller is an auto generated read-only Go binding around an Ethereum contract.
type AWorkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AWorkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AWorkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AWorkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AWorkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AWorkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AWorkSession struct {
	Contract     *AWork            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AWorkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AWorkCallerSession struct {
	Contract *AWorkCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AWorkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AWorkTransactorSession struct {
	Contract     *AWorkTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AWorkRaw is an auto generated low-level Go binding around an Ethereum contract.
type AWorkRaw struct {
	Contract *AWork // Generic contract binding to access the raw methods on
}

// AWorkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AWorkCallerRaw struct {
	Contract *AWorkCaller // Generic read-only contract binding to access the raw methods on
}

// AWorkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AWorkTransactorRaw struct {
	Contract *AWorkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAWork creates a new instance of AWork, bound to a specific deployed contract.
func NewAWork(address common.Address, backend bind.ContractBackend) (*AWork, error) {
	contract, err := bindAWork(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AWork{AWorkCaller: AWorkCaller{contract: contract}, AWorkTransactor: AWorkTransactor{contract: contract}, AWorkFilterer: AWorkFilterer{contract: contract}}, nil
}

// NewAWorkCaller creates a new read-only instance of AWork, bound to a specific deployed contract.
func NewAWorkCaller(address common.Address, caller bind.ContractCaller) (*AWorkCaller, error) {
	contract, err := bindAWork(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AWorkCaller{contract: contract}, nil
}

// NewAWorkTransactor creates a new write-only instance of AWork, bound to a specific deployed contract.
func NewAWorkTransactor(address common.Address, transactor bind.ContractTransactor) (*AWorkTransactor, error) {
	contract, err := bindAWork(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AWorkTransactor{contract: contract}, nil
}

// NewAWorkFilterer creates a new log filterer instance of AWork, bound to a specific deployed contract.
func NewAWorkFilterer(address common.Address, filterer bind.ContractFilterer) (*AWorkFilterer, error) {
	contract, err := bindAWork(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AWorkFilterer{contract: contract}, nil
}

// bindAWork binds a generic wrapper to an already deployed contract.
func bindAWork(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AWorkMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AWork *AWorkRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AWork.Contract.AWorkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AWork *AWorkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AWork.Contract.AWorkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AWork *AWorkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AWork.Contract.AWorkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AWork *AWorkCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AWork.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AWork *AWorkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AWork.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AWork *AWorkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AWork.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_AWork *AWorkCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AWork.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_AWork *AWorkSession) Count() (*big.Int, error) {
	return _AWork.Contract.Count(&_AWork.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_AWork *AWorkCallerSession) Count() (*big.Int, error) {
	return _AWork.Contract.Count(&_AWork.CallOpts)
}

// CounterAdd is a paid mutator transaction binding the contract method 0x3f31a3c2.
//
// Solidity: function counterAdd() returns()
func (_AWork *AWorkTransactor) CounterAdd(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AWork.contract.Transact(opts, "counterAdd")
}

// CounterAdd is a paid mutator transaction binding the contract method 0x3f31a3c2.
//
// Solidity: function counterAdd() returns()
func (_AWork *AWorkSession) CounterAdd() (*types.Transaction, error) {
	return _AWork.Contract.CounterAdd(&_AWork.TransactOpts)
}

// CounterAdd is a paid mutator transaction binding the contract method 0x3f31a3c2.
//
// Solidity: function counterAdd() returns()
func (_AWork *AWorkTransactorSession) CounterAdd() (*types.Transaction, error) {
	return _AWork.Contract.CounterAdd(&_AWork.TransactOpts)
}

// AWorkIncrementIterator is returned from FilterIncrement and is used to iterate over the raw logs and unpacked data for Increment events raised by the AWork contract.
type AWorkIncrementIterator struct {
	Event *AWorkIncrement // Event containing the contract specifics and raw log

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
func (it *AWorkIncrementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AWorkIncrement)
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
		it.Event = new(AWorkIncrement)
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
func (it *AWorkIncrementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AWorkIncrementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AWorkIncrement represents a Increment event raised by the AWork contract.
type AWorkIncrement struct {
	By  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterIncrement is a free log retrieval operation binding the contract event 0x51af157c2eee40f68107a47a49c32fbbeb0a3c9e5cd37aa56e88e6be92368a81.
//
// Solidity: event Increment(uint256 by)
func (_AWork *AWorkFilterer) FilterIncrement(opts *bind.FilterOpts) (*AWorkIncrementIterator, error) {

	logs, sub, err := _AWork.contract.FilterLogs(opts, "Increment")
	if err != nil {
		return nil, err
	}
	return &AWorkIncrementIterator{contract: _AWork.contract, event: "Increment", logs: logs, sub: sub}, nil
}

// WatchIncrement is a free log subscription operation binding the contract event 0x51af157c2eee40f68107a47a49c32fbbeb0a3c9e5cd37aa56e88e6be92368a81.
//
// Solidity: event Increment(uint256 by)
func (_AWork *AWorkFilterer) WatchIncrement(opts *bind.WatchOpts, sink chan<- *AWorkIncrement) (event.Subscription, error) {

	logs, sub, err := _AWork.contract.WatchLogs(opts, "Increment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AWorkIncrement)
				if err := _AWork.contract.UnpackLog(event, "Increment", log); err != nil {
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

// ParseIncrement is a log parse operation binding the contract event 0x51af157c2eee40f68107a47a49c32fbbeb0a3c9e5cd37aa56e88e6be92368a81.
//
// Solidity: event Increment(uint256 by)
func (_AWork *AWorkFilterer) ParseIncrement(log types.Log) (*AWorkIncrement, error) {
	event := new(AWorkIncrement)
	if err := _AWork.contract.UnpackLog(event, "Increment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

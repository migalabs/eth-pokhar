// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cmv2

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

// IMetaRegistryExternalOperator is an auto generated low-level Go binding around an user-defined struct.
type IMetaRegistryExternalOperator struct {
	Data []byte
}

// IMetaRegistryOperatorGroup is an auto generated low-level Go binding around an user-defined struct.
type IMetaRegistryOperatorGroup struct {
	SubNodeOperators  []IMetaRegistrySubNodeOperator
	ExternalOperators []IMetaRegistryExternalOperator
}

// IMetaRegistrySubNodeOperator is an auto generated low-level Go binding around an user-defined struct.
type IMetaRegistrySubNodeOperator struct {
	NodeOperatorId uint64
	Share          uint16
}

// OperatorMetadata is an auto generated low-level Go binding around an user-defined struct.
type OperatorMetadata struct {
	Name                 string
	Description          string
	OwnerEditsRestricted bool
}

// MetaRegistryMetaData contains all meta data concerning the MetaRegistry contract.
var MetaRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"module\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ACCOUNTING\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAccounting\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MANAGE_OPERATOR_GROUPS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MODULE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractICuratedModule\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NO_GROUP_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SET_BOND_CURVE_WEIGHT_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SET_OPERATOR_INFO_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_ROUTER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakingRouter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createOrUpdateOperatorGroup\",\"inputs\":[{\"name\":\"groupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"groupInfo\",\"type\":\"tuple\",\"internalType\":\"structIMetaRegistry.OperatorGroup\",\"components\":[{\"name\":\"subNodeOperators\",\"type\":\"tuple[]\",\"internalType\":\"structIMetaRegistry.SubNodeOperator[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"share\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"externalOperators\",\"type\":\"tuple[]\",\"internalType\":\"structIMetaRegistry.ExternalOperator[]\",\"components\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBondCurveWeight\",\"inputs\":[{\"name\":\"curveId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExternalOperatorGroupId\",\"inputs\":[{\"name\":\"op\",\"type\":\"tuple\",\"internalType\":\"structIMetaRegistry.ExternalOperator\",\"components\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"operatorGroupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInitializedVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOperatorGroupId\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"operatorGroupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOperatorWeight\",\"inputs\":[{\"name\":\"noId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOperatorWeightAndExternalStake\",\"inputs\":[{\"name\":\"noId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"externalStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorGroup\",\"inputs\":[{\"name\":\"groupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"groupInfo\",\"type\":\"tuple\",\"internalType\":\"structIMetaRegistry.OperatorGroup\",\"components\":[{\"name\":\"subNodeOperators\",\"type\":\"tuple[]\",\"internalType\":\"structIMetaRegistry.SubNodeOperator[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"share\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"externalOperators\",\"type\":\"tuple[]\",\"internalType\":\"structIMetaRegistry.ExternalOperator[]\",\"components\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorGroupsCount\",\"inputs\":[],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorMetadata\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"metadata\",\"type\":\"tuple\",\"internalType\":\"structOperatorMetadata\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"ownerEditsRestricted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorWeights\",\"inputs\":[{\"name\":\"nodeOperatorIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"operatorWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMember\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMemberCount\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMembers\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"refreshOperatorWeight\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBondCurveWeight\",\"inputs\":[{\"name\":\"curveId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorMetadataAsAdmin\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"metadata\",\"type\":\"tuple\",\"internalType\":\"structOperatorMetadata\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"ownerEditsRestricted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorMetadataAsOwner\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BondCurveWeightSet\",\"inputs\":[{\"name\":\"curveId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeOperatorEffectiveWeightChanged\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"oldWeight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newWeight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorGroupCleared\",\"inputs\":[{\"name\":\"groupId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorGroupCreated\",\"inputs\":[{\"name\":\"groupId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"groupInfo\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIMetaRegistry.OperatorGroup\",\"components\":[{\"name\":\"subNodeOperators\",\"type\":\"tuple[]\",\"internalType\":\"structIMetaRegistry.SubNodeOperator[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"share\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"externalOperators\",\"type\":\"tuple[]\",\"internalType\":\"structIMetaRegistry.ExternalOperator[]\",\"components\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorGroupUpdated\",\"inputs\":[{\"name\":\"groupId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"groupInfo\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIMetaRegistry.OperatorGroup\",\"components\":[{\"name\":\"subNodeOperators\",\"type\":\"tuple[]\",\"internalType\":\"structIMetaRegistry.SubNodeOperator[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"share\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"externalOperators\",\"type\":\"tuple[]\",\"internalType\":\"structIMetaRegistry.ExternalOperator[]\",\"components\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorMetadataSet\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"metadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorMetadata\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"ownerEditsRestricted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AlreadyUsedAsExternalOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBondCurveWeight\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExternalOperatorDataEntry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorGroup\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorGroupId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSubNodeOperatorShares\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ModuleAddressNotCached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NodeOperatorAlreadyInGroup\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NodeOperatorDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorDescriptionTooLong\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNameTooLong\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerEditsRestricted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SameBondCurveWeight\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderIsNotEligible\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAdminAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroModuleAddress\",\"inputs\":[]}]",
}

// MetaRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use MetaRegistryMetaData.ABI instead.
var MetaRegistryABI = MetaRegistryMetaData.ABI

// MetaRegistry is an auto generated Go binding around an Ethereum contract.
type MetaRegistry struct {
	MetaRegistryCaller     // Read-only binding to the contract
	MetaRegistryTransactor // Write-only binding to the contract
	MetaRegistryFilterer   // Log filterer for contract events
}

// MetaRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetaRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetaRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MetaRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetaRegistrySession struct {
	Contract     *MetaRegistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetaRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetaRegistryCallerSession struct {
	Contract *MetaRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MetaRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetaRegistryTransactorSession struct {
	Contract     *MetaRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MetaRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetaRegistryRaw struct {
	Contract *MetaRegistry // Generic contract binding to access the raw methods on
}

// MetaRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetaRegistryCallerRaw struct {
	Contract *MetaRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// MetaRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetaRegistryTransactorRaw struct {
	Contract *MetaRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetaRegistry creates a new instance of MetaRegistry, bound to a specific deployed contract.
func NewMetaRegistry(address common.Address, backend bind.ContractBackend) (*MetaRegistry, error) {
	contract, err := bindMetaRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MetaRegistry{MetaRegistryCaller: MetaRegistryCaller{contract: contract}, MetaRegistryTransactor: MetaRegistryTransactor{contract: contract}, MetaRegistryFilterer: MetaRegistryFilterer{contract: contract}}, nil
}

// NewMetaRegistryCaller creates a new read-only instance of MetaRegistry, bound to a specific deployed contract.
func NewMetaRegistryCaller(address common.Address, caller bind.ContractCaller) (*MetaRegistryCaller, error) {
	contract, err := bindMetaRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetaRegistryCaller{contract: contract}, nil
}

// NewMetaRegistryTransactor creates a new write-only instance of MetaRegistry, bound to a specific deployed contract.
func NewMetaRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*MetaRegistryTransactor, error) {
	contract, err := bindMetaRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetaRegistryTransactor{contract: contract}, nil
}

// NewMetaRegistryFilterer creates a new log filterer instance of MetaRegistry, bound to a specific deployed contract.
func NewMetaRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*MetaRegistryFilterer, error) {
	contract, err := bindMetaRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetaRegistryFilterer{contract: contract}, nil
}

// bindMetaRegistry binds a generic wrapper to an already deployed contract.
func bindMetaRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MetaRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaRegistry *MetaRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetaRegistry.Contract.MetaRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaRegistry *MetaRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaRegistry.Contract.MetaRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaRegistry *MetaRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaRegistry.Contract.MetaRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaRegistry *MetaRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetaRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaRegistry *MetaRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaRegistry *MetaRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaRegistry.Contract.contract.Transact(opts, method, params...)
}

// ACCOUNTING is a free data retrieval call binding the contract method 0x6dc3f2bd.
//
// Solidity: function ACCOUNTING() view returns(address)
func (_MetaRegistry *MetaRegistryCaller) ACCOUNTING(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaRegistry.contract.Call(opts, &out, "ACCOUNTING")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ACCOUNTING is a free data retrieval call binding the contract method 0x6dc3f2bd.
//
// Solidity: function ACCOUNTING() view returns(address)
func (_MetaRegistry *MetaRegistrySession) ACCOUNTING() (common.Address, error) {
	return _MetaRegistry.Contract.ACCOUNTING(&_MetaRegistry.CallOpts)
}

// ACCOUNTING is a free data retrieval call binding the contract method 0x6dc3f2bd.
//
// Solidity: function ACCOUNTING() view returns(address)
func (_MetaRegistry *MetaRegistryCallerSession) ACCOUNTING() (common.Address, error) {
	return _MetaRegistry.Contract.ACCOUNTING(&_MetaRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MetaRegistry *MetaRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MetaRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MetaRegistry *MetaRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MetaRegistry.Contract.DEFAULTADMINROLE(&_MetaRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MetaRegistry *MetaRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MetaRegistry.Contract.DEFAULTADMINROLE(&_MetaRegistry.CallOpts)
}

// MANAGEOPERATORGROUPSROLE is a free data retrieval call binding the contract method 0x443ef3fe.
//
// Solidity: function MANAGE_OPERATOR_GROUPS_ROLE() view returns(bytes32)
func (_MetaRegistry *MetaRegistryCaller) MANAGEOPERATORGROUPSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MetaRegistry.contract.Call(opts, &out, "MANAGE_OPERATOR_GROUPS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGEOPERATORGROUPSROLE is a free data retrieval call binding the contract method 0x443ef3fe.
//
// Solidity: function MANAGE_OPERATOR_GROUPS_ROLE() view returns(bytes32)
func (_MetaRegistry *MetaRegistrySession) MANAGEOPERATORGROUPSROLE() ([32]byte, error) {
	return _MetaRegistry.Contract.MANAGEOPERATORGROUPSROLE(&_MetaRegistry.CallOpts)
}

// MANAGEOPERATORGROUPSROLE is a free data retrieval call binding the contract method 0x443ef3fe.
//
// Solidity: function MANAGE_OPERATOR_GROUPS_ROLE() view returns(bytes32)
func (_MetaRegistry *MetaRegistryCallerSession) MANAGEOPERATORGROUPSROLE() ([32]byte, error) {
	return _MetaRegistry.Contract.MANAGEOPERATORGROUPSROLE(&_MetaRegistry.CallOpts)
}

// MODULE is a free data retrieval call binding the contract method 0x094d3a34.
//
// Solidity: function MODULE() view returns(address)
func (_MetaRegistry *MetaRegistryCaller) MODULE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaRegistry.contract.Call(opts, &out, "MODULE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULE is a free data retrieval call binding the contract method 0x094d3a34.
//
// Solidity: function MODULE() view returns(address)
func (_MetaRegistry *MetaRegistrySession) MODULE() (common.Address, error) {
	return _MetaRegistry.Contract.MODULE(&_MetaRegistry.CallOpts)
}

// MODULE is a free data retrieval call binding the contract method 0x094d3a34.
//
// Solidity: function MODULE() view returns(address)
func (_MetaRegistry *MetaRegistryCallerSession) MODULE() (common.Address, error) {
	return _MetaRegistry.Contract.MODULE(&_MetaRegistry.CallOpts)
}

// NOGROUPID is a free data retrieval call binding the contract method 0x566fbdab.
//
// Solidity: function NO_GROUP_ID() view returns(uint256)
func (_MetaRegistry *MetaRegistryCaller) NOGROUPID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetaRegistry.contract.Call(opts, &out, "NO_GROUP_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NOGROUPID is a free data retrieval call binding the contract method 0x566fbdab.
//
// Solidity: function NO_GROUP_ID() view returns(uint256)
func (_MetaRegistry *MetaRegistrySession) NOGROUPID() (*big.Int, error) {
	return _MetaRegistry.Contract.NOGROUPID(&_MetaRegistry.CallOpts)
}

// NOGROUPID is a free data retrieval call binding the contract method 0x566fbdab.
//
// Solidity: function NO_GROUP_ID() view returns(uint256)
func (_MetaRegistry *MetaRegistryCallerSession) NOGROUPID() (*big.Int, error) {
	return _MetaRegistry.Contract.NOGROUPID(&_MetaRegistry.CallOpts)
}

// SETBONDCURVEWEIGHTROLE is a free data retrieval call binding the contract method 0x83de9c6c.
//
// Solidity: function SET_BOND_CURVE_WEIGHT_ROLE() view returns(bytes32)
func (_MetaRegistry *MetaRegistryCaller) SETBONDCURVEWEIGHTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MetaRegistry.contract.Call(opts, &out, "SET_BOND_CURVE_WEIGHT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETBONDCURVEWEIGHTROLE is a free data retrieval call binding the contract method 0x83de9c6c.
//
// Solidity: function SET_BOND_CURVE_WEIGHT_ROLE() view returns(bytes32)
func (_MetaRegistry *MetaRegistrySession) SETBONDCURVEWEIGHTROLE() ([32]byte, error) {
	return _MetaRegistry.Contract.SETBONDCURVEWEIGHTROLE(&_MetaRegistry.CallOpts)
}

// SETBONDCURVEWEIGHTROLE is a free data retrieval call binding the contract method 0x83de9c6c.
//
// Solidity: function SET_BOND_CURVE_WEIGHT_ROLE() view returns(bytes32)
func (_MetaRegistry *MetaRegistryCallerSession) SETBONDCURVEWEIGHTROLE() ([32]byte, error) {
	return _MetaRegistry.Contract.SETBONDCURVEWEIGHTROLE(&_MetaRegistry.CallOpts)
}

// SETOPERATORINFOROLE is a free data retrieval call binding the contract method 0x88b81774.
//
// Solidity: function SET_OPERATOR_INFO_ROLE() view returns(bytes32)
func (_MetaRegistry *MetaRegistryCaller) SETOPERATORINFOROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MetaRegistry.contract.Call(opts, &out, "SET_OPERATOR_INFO_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETOPERATORINFOROLE is a free data retrieval call binding the contract method 0x88b81774.
//
// Solidity: function SET_OPERATOR_INFO_ROLE() view returns(bytes32)
func (_MetaRegistry *MetaRegistrySession) SETOPERATORINFOROLE() ([32]byte, error) {
	return _MetaRegistry.Contract.SETOPERATORINFOROLE(&_MetaRegistry.CallOpts)
}

// SETOPERATORINFOROLE is a free data retrieval call binding the contract method 0x88b81774.
//
// Solidity: function SET_OPERATOR_INFO_ROLE() view returns(bytes32)
func (_MetaRegistry *MetaRegistryCallerSession) SETOPERATORINFOROLE() ([32]byte, error) {
	return _MetaRegistry.Contract.SETOPERATORINFOROLE(&_MetaRegistry.CallOpts)
}

// STAKINGROUTER is a free data retrieval call binding the contract method 0xc8a5f8e6.
//
// Solidity: function STAKING_ROUTER() view returns(address)
func (_MetaRegistry *MetaRegistryCaller) STAKINGROUTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaRegistry.contract.Call(opts, &out, "STAKING_ROUTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKINGROUTER is a free data retrieval call binding the contract method 0xc8a5f8e6.
//
// Solidity: function STAKING_ROUTER() view returns(address)
func (_MetaRegistry *MetaRegistrySession) STAKINGROUTER() (common.Address, error) {
	return _MetaRegistry.Contract.STAKINGROUTER(&_MetaRegistry.CallOpts)
}

// STAKINGROUTER is a free data retrieval call binding the contract method 0xc8a5f8e6.
//
// Solidity: function STAKING_ROUTER() view returns(address)
func (_MetaRegistry *MetaRegistryCallerSession) STAKINGROUTER() (common.Address, error) {
	return _MetaRegistry.Contract.STAKINGROUTER(&_MetaRegistry.CallOpts)
}

// GetBondCurveWeight is a free data retrieval call binding the contract method 0x29bbfcb9.
//
// Solidity: function getBondCurveWeight(uint256 curveId) view returns(uint256 weight)
func (_MetaRegistry *MetaRegistryCaller) GetBondCurveWeight(opts *bind.CallOpts, curveId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MetaRegistry.contract.Call(opts, &out, "getBondCurveWeight", curveId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBondCurveWeight is a free data retrieval call binding the contract method 0x29bbfcb9.
//
// Solidity: function getBondCurveWeight(uint256 curveId) view returns(uint256 weight)
func (_MetaRegistry *MetaRegistrySession) GetBondCurveWeight(curveId *big.Int) (*big.Int, error) {
	return _MetaRegistry.Contract.GetBondCurveWeight(&_MetaRegistry.CallOpts, curveId)
}

// GetBondCurveWeight is a free data retrieval call binding the contract method 0x29bbfcb9.
//
// Solidity: function getBondCurveWeight(uint256 curveId) view returns(uint256 weight)
func (_MetaRegistry *MetaRegistryCallerSession) GetBondCurveWeight(curveId *big.Int) (*big.Int, error) {
	return _MetaRegistry.Contract.GetBondCurveWeight(&_MetaRegistry.CallOpts, curveId)
}

// GetExternalOperatorGroupId is a free data retrieval call binding the contract method 0x7a900bf9.
//
// Solidity: function getExternalOperatorGroupId((bytes) op) view returns(uint256 operatorGroupId)
func (_MetaRegistry *MetaRegistryCaller) GetExternalOperatorGroupId(opts *bind.CallOpts, op IMetaRegistryExternalOperator) (*big.Int, error) {
	var out []interface{}
	err := _MetaRegistry.contract.Call(opts, &out, "getExternalOperatorGroupId", op)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExternalOperatorGroupId is a free data retrieval call binding the contract method 0x7a900bf9.
//
// Solidity: function getExternalOperatorGroupId((bytes) op) view returns(uint256 operatorGroupId)
func (_MetaRegistry *MetaRegistrySession) GetExternalOperatorGroupId(op IMetaRegistryExternalOperator) (*big.Int, error) {
	return _MetaRegistry.Contract.GetExternalOperatorGroupId(&_MetaRegistry.CallOpts, op)
}

// GetExternalOperatorGroupId is a free data retrieval call binding the contract method 0x7a900bf9.
//
// Solidity: function getExternalOperatorGroupId((bytes) op) view returns(uint256 operatorGroupId)
func (_MetaRegistry *MetaRegistryCallerSession) GetExternalOperatorGroupId(op IMetaRegistryExternalOperator) (*big.Int, error) {
	return _MetaRegistry.Contract.GetExternalOperatorGroupId(&_MetaRegistry.CallOpts, op)
}

// GetInitializedVersion is a free data retrieval call binding the contract method 0xb3c65015.
//
// Solidity: function getInitializedVersion() view returns(uint64)
func (_MetaRegistry *MetaRegistryCaller) GetInitializedVersion(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MetaRegistry.contract.Call(opts, &out, "getInitializedVersion")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetInitializedVersion is a free data retrieval call binding the contract method 0xb3c65015.
//
// Solidity: function getInitializedVersion() view returns(uint64)
func (_MetaRegistry *MetaRegistrySession) GetInitializedVersion() (uint64, error) {
	return _MetaRegistry.Contract.GetInitializedVersion(&_MetaRegistry.CallOpts)
}

// GetInitializedVersion is a free data retrieval call binding the contract method 0xb3c65015.
//
// Solidity: function getInitializedVersion() view returns(uint64)
func (_MetaRegistry *MetaRegistryCallerSession) GetInitializedVersion() (uint64, error) {
	return _MetaRegistry.Contract.GetInitializedVersion(&_MetaRegistry.CallOpts)
}

// GetNodeOperatorGroupId is a free data retrieval call binding the contract method 0x70dc92e2.
//
// Solidity: function getNodeOperatorGroupId(uint256 nodeOperatorId) view returns(uint256 operatorGroupId)
func (_MetaRegistry *MetaRegistryCaller) GetNodeOperatorGroupId(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MetaRegistry.contract.Call(opts, &out, "getNodeOperatorGroupId", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeOperatorGroupId is a free data retrieval call binding the contract method 0x70dc92e2.
//
// Solidity: function getNodeOperatorGroupId(uint256 nodeOperatorId) view returns(uint256 operatorGroupId)
func (_MetaRegistry *MetaRegistrySession) GetNodeOperatorGroupId(nodeOperatorId *big.Int) (*big.Int, error) {
	return _MetaRegistry.Contract.GetNodeOperatorGroupId(&_MetaRegistry.CallOpts, nodeOperatorId)
}

// GetNodeOperatorGroupId is a free data retrieval call binding the contract method 0x70dc92e2.
//
// Solidity: function getNodeOperatorGroupId(uint256 nodeOperatorId) view returns(uint256 operatorGroupId)
func (_MetaRegistry *MetaRegistryCallerSession) GetNodeOperatorGroupId(nodeOperatorId *big.Int) (*big.Int, error) {
	return _MetaRegistry.Contract.GetNodeOperatorGroupId(&_MetaRegistry.CallOpts, nodeOperatorId)
}

// GetNodeOperatorWeight is a free data retrieval call binding the contract method 0xa940e268.
//
// Solidity: function getNodeOperatorWeight(uint256 noId) view returns(uint256 weight)
func (_MetaRegistry *MetaRegistryCaller) GetNodeOperatorWeight(opts *bind.CallOpts, noId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MetaRegistry.contract.Call(opts, &out, "getNodeOperatorWeight", noId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeOperatorWeight is a free data retrieval call binding the contract method 0xa940e268.
//
// Solidity: function getNodeOperatorWeight(uint256 noId) view returns(uint256 weight)
func (_MetaRegistry *MetaRegistrySession) GetNodeOperatorWeight(noId *big.Int) (*big.Int, error) {
	return _MetaRegistry.Contract.GetNodeOperatorWeight(&_MetaRegistry.CallOpts, noId)
}

// GetNodeOperatorWeight is a free data retrieval call binding the contract method 0xa940e268.
//
// Solidity: function getNodeOperatorWeight(uint256 noId) view returns(uint256 weight)
func (_MetaRegistry *MetaRegistryCallerSession) GetNodeOperatorWeight(noId *big.Int) (*big.Int, error) {
	return _MetaRegistry.Contract.GetNodeOperatorWeight(&_MetaRegistry.CallOpts, noId)
}

// GetNodeOperatorWeightAndExternalStake is a free data retrieval call binding the contract method 0x9703e7bc.
//
// Solidity: function getNodeOperatorWeightAndExternalStake(uint256 noId) view returns(uint256 weight, uint256 externalStake)
func (_MetaRegistry *MetaRegistryCaller) GetNodeOperatorWeightAndExternalStake(opts *bind.CallOpts, noId *big.Int) (struct {
	Weight        *big.Int
	ExternalStake *big.Int
}, error) {
	var out []interface{}
	err := _MetaRegistry.contract.Call(opts, &out, "getNodeOperatorWeightAndExternalStake", noId)

	outstruct := new(struct {
		Weight        *big.Int
		ExternalStake *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Weight = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ExternalStake = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNodeOperatorWeightAndExternalStake is a free data retrieval call binding the contract method 0x9703e7bc.
//
// Solidity: function getNodeOperatorWeightAndExternalStake(uint256 noId) view returns(uint256 weight, uint256 externalStake)
func (_MetaRegistry *MetaRegistrySession) GetNodeOperatorWeightAndExternalStake(noId *big.Int) (struct {
	Weight        *big.Int
	ExternalStake *big.Int
}, error) {
	return _MetaRegistry.Contract.GetNodeOperatorWeightAndExternalStake(&_MetaRegistry.CallOpts, noId)
}

// GetNodeOperatorWeightAndExternalStake is a free data retrieval call binding the contract method 0x9703e7bc.
//
// Solidity: function getNodeOperatorWeightAndExternalStake(uint256 noId) view returns(uint256 weight, uint256 externalStake)
func (_MetaRegistry *MetaRegistryCallerSession) GetNodeOperatorWeightAndExternalStake(noId *big.Int) (struct {
	Weight        *big.Int
	ExternalStake *big.Int
}, error) {
	return _MetaRegistry.Contract.GetNodeOperatorWeightAndExternalStake(&_MetaRegistry.CallOpts, noId)
}

// GetOperatorGroup is a free data retrieval call binding the contract method 0x0c665463.
//
// Solidity: function getOperatorGroup(uint256 groupId) view returns(((uint64,uint16)[],(bytes)[]) groupInfo)
func (_MetaRegistry *MetaRegistryCaller) GetOperatorGroup(opts *bind.CallOpts, groupId *big.Int) (IMetaRegistryOperatorGroup, error) {
	var out []interface{}
	err := _MetaRegistry.contract.Call(opts, &out, "getOperatorGroup", groupId)

	if err != nil {
		return *new(IMetaRegistryOperatorGroup), err
	}

	out0 := *abi.ConvertType(out[0], new(IMetaRegistryOperatorGroup)).(*IMetaRegistryOperatorGroup)

	return out0, err

}

// GetOperatorGroup is a free data retrieval call binding the contract method 0x0c665463.
//
// Solidity: function getOperatorGroup(uint256 groupId) view returns(((uint64,uint16)[],(bytes)[]) groupInfo)
func (_MetaRegistry *MetaRegistrySession) GetOperatorGroup(groupId *big.Int) (IMetaRegistryOperatorGroup, error) {
	return _MetaRegistry.Contract.GetOperatorGroup(&_MetaRegistry.CallOpts, groupId)
}

// GetOperatorGroup is a free data retrieval call binding the contract method 0x0c665463.
//
// Solidity: function getOperatorGroup(uint256 groupId) view returns(((uint64,uint16)[],(bytes)[]) groupInfo)
func (_MetaRegistry *MetaRegistryCallerSession) GetOperatorGroup(groupId *big.Int) (IMetaRegistryOperatorGroup, error) {
	return _MetaRegistry.Contract.GetOperatorGroup(&_MetaRegistry.CallOpts, groupId)
}

// GetOperatorGroupsCount is a free data retrieval call binding the contract method 0x5badee3b.
//
// Solidity: function getOperatorGroupsCount() view returns(uint256 count)
func (_MetaRegistry *MetaRegistryCaller) GetOperatorGroupsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetaRegistry.contract.Call(opts, &out, "getOperatorGroupsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorGroupsCount is a free data retrieval call binding the contract method 0x5badee3b.
//
// Solidity: function getOperatorGroupsCount() view returns(uint256 count)
func (_MetaRegistry *MetaRegistrySession) GetOperatorGroupsCount() (*big.Int, error) {
	return _MetaRegistry.Contract.GetOperatorGroupsCount(&_MetaRegistry.CallOpts)
}

// GetOperatorGroupsCount is a free data retrieval call binding the contract method 0x5badee3b.
//
// Solidity: function getOperatorGroupsCount() view returns(uint256 count)
func (_MetaRegistry *MetaRegistryCallerSession) GetOperatorGroupsCount() (*big.Int, error) {
	return _MetaRegistry.Contract.GetOperatorGroupsCount(&_MetaRegistry.CallOpts)
}

// GetOperatorMetadata is a free data retrieval call binding the contract method 0xaca6d2da.
//
// Solidity: function getOperatorMetadata(uint256 nodeOperatorId) view returns((string,string,bool) metadata)
func (_MetaRegistry *MetaRegistryCaller) GetOperatorMetadata(opts *bind.CallOpts, nodeOperatorId *big.Int) (OperatorMetadata, error) {
	var out []interface{}
	err := _MetaRegistry.contract.Call(opts, &out, "getOperatorMetadata", nodeOperatorId)

	if err != nil {
		return *new(OperatorMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorMetadata)).(*OperatorMetadata)

	return out0, err

}

// GetOperatorMetadata is a free data retrieval call binding the contract method 0xaca6d2da.
//
// Solidity: function getOperatorMetadata(uint256 nodeOperatorId) view returns((string,string,bool) metadata)
func (_MetaRegistry *MetaRegistrySession) GetOperatorMetadata(nodeOperatorId *big.Int) (OperatorMetadata, error) {
	return _MetaRegistry.Contract.GetOperatorMetadata(&_MetaRegistry.CallOpts, nodeOperatorId)
}

// GetOperatorMetadata is a free data retrieval call binding the contract method 0xaca6d2da.
//
// Solidity: function getOperatorMetadata(uint256 nodeOperatorId) view returns((string,string,bool) metadata)
func (_MetaRegistry *MetaRegistryCallerSession) GetOperatorMetadata(nodeOperatorId *big.Int) (OperatorMetadata, error) {
	return _MetaRegistry.Contract.GetOperatorMetadata(&_MetaRegistry.CallOpts, nodeOperatorId)
}

// GetOperatorWeights is a free data retrieval call binding the contract method 0x08ca8c6d.
//
// Solidity: function getOperatorWeights(uint256[] nodeOperatorIds) view returns(uint256[] operatorWeights)
func (_MetaRegistry *MetaRegistryCaller) GetOperatorWeights(opts *bind.CallOpts, nodeOperatorIds []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _MetaRegistry.contract.Call(opts, &out, "getOperatorWeights", nodeOperatorIds)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOperatorWeights is a free data retrieval call binding the contract method 0x08ca8c6d.
//
// Solidity: function getOperatorWeights(uint256[] nodeOperatorIds) view returns(uint256[] operatorWeights)
func (_MetaRegistry *MetaRegistrySession) GetOperatorWeights(nodeOperatorIds []*big.Int) ([]*big.Int, error) {
	return _MetaRegistry.Contract.GetOperatorWeights(&_MetaRegistry.CallOpts, nodeOperatorIds)
}

// GetOperatorWeights is a free data retrieval call binding the contract method 0x08ca8c6d.
//
// Solidity: function getOperatorWeights(uint256[] nodeOperatorIds) view returns(uint256[] operatorWeights)
func (_MetaRegistry *MetaRegistryCallerSession) GetOperatorWeights(nodeOperatorIds []*big.Int) ([]*big.Int, error) {
	return _MetaRegistry.Contract.GetOperatorWeights(&_MetaRegistry.CallOpts, nodeOperatorIds)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MetaRegistry *MetaRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MetaRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MetaRegistry *MetaRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MetaRegistry.Contract.GetRoleAdmin(&_MetaRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MetaRegistry *MetaRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MetaRegistry.Contract.GetRoleAdmin(&_MetaRegistry.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MetaRegistry *MetaRegistryCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MetaRegistry.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MetaRegistry *MetaRegistrySession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _MetaRegistry.Contract.GetRoleMember(&_MetaRegistry.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MetaRegistry *MetaRegistryCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _MetaRegistry.Contract.GetRoleMember(&_MetaRegistry.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MetaRegistry *MetaRegistryCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MetaRegistry.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MetaRegistry *MetaRegistrySession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _MetaRegistry.Contract.GetRoleMemberCount(&_MetaRegistry.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MetaRegistry *MetaRegistryCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _MetaRegistry.Contract.GetRoleMemberCount(&_MetaRegistry.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_MetaRegistry *MetaRegistryCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _MetaRegistry.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_MetaRegistry *MetaRegistrySession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _MetaRegistry.Contract.GetRoleMembers(&_MetaRegistry.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_MetaRegistry *MetaRegistryCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _MetaRegistry.Contract.GetRoleMembers(&_MetaRegistry.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MetaRegistry *MetaRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _MetaRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MetaRegistry *MetaRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MetaRegistry.Contract.HasRole(&_MetaRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MetaRegistry *MetaRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MetaRegistry.Contract.HasRole(&_MetaRegistry.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MetaRegistry *MetaRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MetaRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MetaRegistry *MetaRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MetaRegistry.Contract.SupportsInterface(&_MetaRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MetaRegistry *MetaRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MetaRegistry.Contract.SupportsInterface(&_MetaRegistry.CallOpts, interfaceId)
}

// CreateOrUpdateOperatorGroup is a paid mutator transaction binding the contract method 0x63fe89d7.
//
// Solidity: function createOrUpdateOperatorGroup(uint256 groupId, ((uint64,uint16)[],(bytes)[]) groupInfo) returns()
func (_MetaRegistry *MetaRegistryTransactor) CreateOrUpdateOperatorGroup(opts *bind.TransactOpts, groupId *big.Int, groupInfo IMetaRegistryOperatorGroup) (*types.Transaction, error) {
	return _MetaRegistry.contract.Transact(opts, "createOrUpdateOperatorGroup", groupId, groupInfo)
}

// CreateOrUpdateOperatorGroup is a paid mutator transaction binding the contract method 0x63fe89d7.
//
// Solidity: function createOrUpdateOperatorGroup(uint256 groupId, ((uint64,uint16)[],(bytes)[]) groupInfo) returns()
func (_MetaRegistry *MetaRegistrySession) CreateOrUpdateOperatorGroup(groupId *big.Int, groupInfo IMetaRegistryOperatorGroup) (*types.Transaction, error) {
	return _MetaRegistry.Contract.CreateOrUpdateOperatorGroup(&_MetaRegistry.TransactOpts, groupId, groupInfo)
}

// CreateOrUpdateOperatorGroup is a paid mutator transaction binding the contract method 0x63fe89d7.
//
// Solidity: function createOrUpdateOperatorGroup(uint256 groupId, ((uint64,uint16)[],(bytes)[]) groupInfo) returns()
func (_MetaRegistry *MetaRegistryTransactorSession) CreateOrUpdateOperatorGroup(groupId *big.Int, groupInfo IMetaRegistryOperatorGroup) (*types.Transaction, error) {
	return _MetaRegistry.Contract.CreateOrUpdateOperatorGroup(&_MetaRegistry.TransactOpts, groupId, groupInfo)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MetaRegistry *MetaRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MetaRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MetaRegistry *MetaRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MetaRegistry.Contract.GrantRole(&_MetaRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MetaRegistry *MetaRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MetaRegistry.Contract.GrantRole(&_MetaRegistry.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_MetaRegistry *MetaRegistryTransactor) Initialize(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _MetaRegistry.contract.Transact(opts, "initialize", admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_MetaRegistry *MetaRegistrySession) Initialize(admin common.Address) (*types.Transaction, error) {
	return _MetaRegistry.Contract.Initialize(&_MetaRegistry.TransactOpts, admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_MetaRegistry *MetaRegistryTransactorSession) Initialize(admin common.Address) (*types.Transaction, error) {
	return _MetaRegistry.Contract.Initialize(&_MetaRegistry.TransactOpts, admin)
}

// RefreshOperatorWeight is a paid mutator transaction binding the contract method 0x2e09e97f.
//
// Solidity: function refreshOperatorWeight(uint256 nodeOperatorId) returns()
func (_MetaRegistry *MetaRegistryTransactor) RefreshOperatorWeight(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _MetaRegistry.contract.Transact(opts, "refreshOperatorWeight", nodeOperatorId)
}

// RefreshOperatorWeight is a paid mutator transaction binding the contract method 0x2e09e97f.
//
// Solidity: function refreshOperatorWeight(uint256 nodeOperatorId) returns()
func (_MetaRegistry *MetaRegistrySession) RefreshOperatorWeight(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _MetaRegistry.Contract.RefreshOperatorWeight(&_MetaRegistry.TransactOpts, nodeOperatorId)
}

// RefreshOperatorWeight is a paid mutator transaction binding the contract method 0x2e09e97f.
//
// Solidity: function refreshOperatorWeight(uint256 nodeOperatorId) returns()
func (_MetaRegistry *MetaRegistryTransactorSession) RefreshOperatorWeight(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _MetaRegistry.Contract.RefreshOperatorWeight(&_MetaRegistry.TransactOpts, nodeOperatorId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MetaRegistry *MetaRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MetaRegistry.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MetaRegistry *MetaRegistrySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MetaRegistry.Contract.RenounceRole(&_MetaRegistry.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MetaRegistry *MetaRegistryTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MetaRegistry.Contract.RenounceRole(&_MetaRegistry.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MetaRegistry *MetaRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MetaRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MetaRegistry *MetaRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MetaRegistry.Contract.RevokeRole(&_MetaRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MetaRegistry *MetaRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MetaRegistry.Contract.RevokeRole(&_MetaRegistry.TransactOpts, role, account)
}

// SetBondCurveWeight is a paid mutator transaction binding the contract method 0x4c5353a3.
//
// Solidity: function setBondCurveWeight(uint256 curveId, uint256 weight) returns()
func (_MetaRegistry *MetaRegistryTransactor) SetBondCurveWeight(opts *bind.TransactOpts, curveId *big.Int, weight *big.Int) (*types.Transaction, error) {
	return _MetaRegistry.contract.Transact(opts, "setBondCurveWeight", curveId, weight)
}

// SetBondCurveWeight is a paid mutator transaction binding the contract method 0x4c5353a3.
//
// Solidity: function setBondCurveWeight(uint256 curveId, uint256 weight) returns()
func (_MetaRegistry *MetaRegistrySession) SetBondCurveWeight(curveId *big.Int, weight *big.Int) (*types.Transaction, error) {
	return _MetaRegistry.Contract.SetBondCurveWeight(&_MetaRegistry.TransactOpts, curveId, weight)
}

// SetBondCurveWeight is a paid mutator transaction binding the contract method 0x4c5353a3.
//
// Solidity: function setBondCurveWeight(uint256 curveId, uint256 weight) returns()
func (_MetaRegistry *MetaRegistryTransactorSession) SetBondCurveWeight(curveId *big.Int, weight *big.Int) (*types.Transaction, error) {
	return _MetaRegistry.Contract.SetBondCurveWeight(&_MetaRegistry.TransactOpts, curveId, weight)
}

// SetOperatorMetadataAsAdmin is a paid mutator transaction binding the contract method 0x586df724.
//
// Solidity: function setOperatorMetadataAsAdmin(uint256 nodeOperatorId, (string,string,bool) metadata) returns()
func (_MetaRegistry *MetaRegistryTransactor) SetOperatorMetadataAsAdmin(opts *bind.TransactOpts, nodeOperatorId *big.Int, metadata OperatorMetadata) (*types.Transaction, error) {
	return _MetaRegistry.contract.Transact(opts, "setOperatorMetadataAsAdmin", nodeOperatorId, metadata)
}

// SetOperatorMetadataAsAdmin is a paid mutator transaction binding the contract method 0x586df724.
//
// Solidity: function setOperatorMetadataAsAdmin(uint256 nodeOperatorId, (string,string,bool) metadata) returns()
func (_MetaRegistry *MetaRegistrySession) SetOperatorMetadataAsAdmin(nodeOperatorId *big.Int, metadata OperatorMetadata) (*types.Transaction, error) {
	return _MetaRegistry.Contract.SetOperatorMetadataAsAdmin(&_MetaRegistry.TransactOpts, nodeOperatorId, metadata)
}

// SetOperatorMetadataAsAdmin is a paid mutator transaction binding the contract method 0x586df724.
//
// Solidity: function setOperatorMetadataAsAdmin(uint256 nodeOperatorId, (string,string,bool) metadata) returns()
func (_MetaRegistry *MetaRegistryTransactorSession) SetOperatorMetadataAsAdmin(nodeOperatorId *big.Int, metadata OperatorMetadata) (*types.Transaction, error) {
	return _MetaRegistry.Contract.SetOperatorMetadataAsAdmin(&_MetaRegistry.TransactOpts, nodeOperatorId, metadata)
}

// SetOperatorMetadataAsOwner is a paid mutator transaction binding the contract method 0xc7fcad43.
//
// Solidity: function setOperatorMetadataAsOwner(uint256 nodeOperatorId, string name, string description) returns()
func (_MetaRegistry *MetaRegistryTransactor) SetOperatorMetadataAsOwner(opts *bind.TransactOpts, nodeOperatorId *big.Int, name string, description string) (*types.Transaction, error) {
	return _MetaRegistry.contract.Transact(opts, "setOperatorMetadataAsOwner", nodeOperatorId, name, description)
}

// SetOperatorMetadataAsOwner is a paid mutator transaction binding the contract method 0xc7fcad43.
//
// Solidity: function setOperatorMetadataAsOwner(uint256 nodeOperatorId, string name, string description) returns()
func (_MetaRegistry *MetaRegistrySession) SetOperatorMetadataAsOwner(nodeOperatorId *big.Int, name string, description string) (*types.Transaction, error) {
	return _MetaRegistry.Contract.SetOperatorMetadataAsOwner(&_MetaRegistry.TransactOpts, nodeOperatorId, name, description)
}

// SetOperatorMetadataAsOwner is a paid mutator transaction binding the contract method 0xc7fcad43.
//
// Solidity: function setOperatorMetadataAsOwner(uint256 nodeOperatorId, string name, string description) returns()
func (_MetaRegistry *MetaRegistryTransactorSession) SetOperatorMetadataAsOwner(nodeOperatorId *big.Int, name string, description string) (*types.Transaction, error) {
	return _MetaRegistry.Contract.SetOperatorMetadataAsOwner(&_MetaRegistry.TransactOpts, nodeOperatorId, name, description)
}

// MetaRegistryBondCurveWeightSetIterator is returned from FilterBondCurveWeightSet and is used to iterate over the raw logs and unpacked data for BondCurveWeightSet events raised by the MetaRegistry contract.
type MetaRegistryBondCurveWeightSetIterator struct {
	Event *MetaRegistryBondCurveWeightSet // Event containing the contract specifics and raw log

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
func (it *MetaRegistryBondCurveWeightSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaRegistryBondCurveWeightSet)
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
		it.Event = new(MetaRegistryBondCurveWeightSet)
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
func (it *MetaRegistryBondCurveWeightSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaRegistryBondCurveWeightSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaRegistryBondCurveWeightSet represents a BondCurveWeightSet event raised by the MetaRegistry contract.
type MetaRegistryBondCurveWeightSet struct {
	CurveId *big.Int
	Weight  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBondCurveWeightSet is a free log retrieval operation binding the contract event 0x6f7d3cc32dbb04cf62121ac85aa0a4ad6ea80a7ee589075cef0df07f9b646142.
//
// Solidity: event BondCurveWeightSet(uint256 indexed curveId, uint256 weight)
func (_MetaRegistry *MetaRegistryFilterer) FilterBondCurveWeightSet(opts *bind.FilterOpts, curveId []*big.Int) (*MetaRegistryBondCurveWeightSetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _MetaRegistry.contract.FilterLogs(opts, "BondCurveWeightSet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &MetaRegistryBondCurveWeightSetIterator{contract: _MetaRegistry.contract, event: "BondCurveWeightSet", logs: logs, sub: sub}, nil
}

// WatchBondCurveWeightSet is a free log subscription operation binding the contract event 0x6f7d3cc32dbb04cf62121ac85aa0a4ad6ea80a7ee589075cef0df07f9b646142.
//
// Solidity: event BondCurveWeightSet(uint256 indexed curveId, uint256 weight)
func (_MetaRegistry *MetaRegistryFilterer) WatchBondCurveWeightSet(opts *bind.WatchOpts, sink chan<- *MetaRegistryBondCurveWeightSet, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _MetaRegistry.contract.WatchLogs(opts, "BondCurveWeightSet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaRegistryBondCurveWeightSet)
				if err := _MetaRegistry.contract.UnpackLog(event, "BondCurveWeightSet", log); err != nil {
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

// ParseBondCurveWeightSet is a log parse operation binding the contract event 0x6f7d3cc32dbb04cf62121ac85aa0a4ad6ea80a7ee589075cef0df07f9b646142.
//
// Solidity: event BondCurveWeightSet(uint256 indexed curveId, uint256 weight)
func (_MetaRegistry *MetaRegistryFilterer) ParseBondCurveWeightSet(log types.Log) (*MetaRegistryBondCurveWeightSet, error) {
	event := new(MetaRegistryBondCurveWeightSet)
	if err := _MetaRegistry.contract.UnpackLog(event, "BondCurveWeightSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MetaRegistry contract.
type MetaRegistryInitializedIterator struct {
	Event *MetaRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *MetaRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaRegistryInitialized)
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
		it.Event = new(MetaRegistryInitialized)
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
func (it *MetaRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaRegistryInitialized represents a Initialized event raised by the MetaRegistry contract.
type MetaRegistryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MetaRegistry *MetaRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*MetaRegistryInitializedIterator, error) {

	logs, sub, err := _MetaRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MetaRegistryInitializedIterator{contract: _MetaRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MetaRegistry *MetaRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MetaRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _MetaRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaRegistryInitialized)
				if err := _MetaRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MetaRegistry *MetaRegistryFilterer) ParseInitialized(log types.Log) (*MetaRegistryInitialized, error) {
	event := new(MetaRegistryInitialized)
	if err := _MetaRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaRegistryNodeOperatorEffectiveWeightChangedIterator is returned from FilterNodeOperatorEffectiveWeightChanged and is used to iterate over the raw logs and unpacked data for NodeOperatorEffectiveWeightChanged events raised by the MetaRegistry contract.
type MetaRegistryNodeOperatorEffectiveWeightChangedIterator struct {
	Event *MetaRegistryNodeOperatorEffectiveWeightChanged // Event containing the contract specifics and raw log

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
func (it *MetaRegistryNodeOperatorEffectiveWeightChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaRegistryNodeOperatorEffectiveWeightChanged)
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
		it.Event = new(MetaRegistryNodeOperatorEffectiveWeightChanged)
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
func (it *MetaRegistryNodeOperatorEffectiveWeightChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaRegistryNodeOperatorEffectiveWeightChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaRegistryNodeOperatorEffectiveWeightChanged represents a NodeOperatorEffectiveWeightChanged event raised by the MetaRegistry contract.
type MetaRegistryNodeOperatorEffectiveWeightChanged struct {
	NodeOperatorId *big.Int
	OldWeight      *big.Int
	NewWeight      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorEffectiveWeightChanged is a free log retrieval operation binding the contract event 0x24f422ac3e13fecf5ab70d1db8427a6ed91b508ef302eacc12ead0a91f4c541b.
//
// Solidity: event NodeOperatorEffectiveWeightChanged(uint256 indexed nodeOperatorId, uint256 oldWeight, uint256 newWeight)
func (_MetaRegistry *MetaRegistryFilterer) FilterNodeOperatorEffectiveWeightChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*MetaRegistryNodeOperatorEffectiveWeightChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _MetaRegistry.contract.FilterLogs(opts, "NodeOperatorEffectiveWeightChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &MetaRegistryNodeOperatorEffectiveWeightChangedIterator{contract: _MetaRegistry.contract, event: "NodeOperatorEffectiveWeightChanged", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorEffectiveWeightChanged is a free log subscription operation binding the contract event 0x24f422ac3e13fecf5ab70d1db8427a6ed91b508ef302eacc12ead0a91f4c541b.
//
// Solidity: event NodeOperatorEffectiveWeightChanged(uint256 indexed nodeOperatorId, uint256 oldWeight, uint256 newWeight)
func (_MetaRegistry *MetaRegistryFilterer) WatchNodeOperatorEffectiveWeightChanged(opts *bind.WatchOpts, sink chan<- *MetaRegistryNodeOperatorEffectiveWeightChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _MetaRegistry.contract.WatchLogs(opts, "NodeOperatorEffectiveWeightChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaRegistryNodeOperatorEffectiveWeightChanged)
				if err := _MetaRegistry.contract.UnpackLog(event, "NodeOperatorEffectiveWeightChanged", log); err != nil {
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

// ParseNodeOperatorEffectiveWeightChanged is a log parse operation binding the contract event 0x24f422ac3e13fecf5ab70d1db8427a6ed91b508ef302eacc12ead0a91f4c541b.
//
// Solidity: event NodeOperatorEffectiveWeightChanged(uint256 indexed nodeOperatorId, uint256 oldWeight, uint256 newWeight)
func (_MetaRegistry *MetaRegistryFilterer) ParseNodeOperatorEffectiveWeightChanged(log types.Log) (*MetaRegistryNodeOperatorEffectiveWeightChanged, error) {
	event := new(MetaRegistryNodeOperatorEffectiveWeightChanged)
	if err := _MetaRegistry.contract.UnpackLog(event, "NodeOperatorEffectiveWeightChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaRegistryOperatorGroupClearedIterator is returned from FilterOperatorGroupCleared and is used to iterate over the raw logs and unpacked data for OperatorGroupCleared events raised by the MetaRegistry contract.
type MetaRegistryOperatorGroupClearedIterator struct {
	Event *MetaRegistryOperatorGroupCleared // Event containing the contract specifics and raw log

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
func (it *MetaRegistryOperatorGroupClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaRegistryOperatorGroupCleared)
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
		it.Event = new(MetaRegistryOperatorGroupCleared)
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
func (it *MetaRegistryOperatorGroupClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaRegistryOperatorGroupClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaRegistryOperatorGroupCleared represents a OperatorGroupCleared event raised by the MetaRegistry contract.
type MetaRegistryOperatorGroupCleared struct {
	GroupId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOperatorGroupCleared is a free log retrieval operation binding the contract event 0x9ca6c49aaad163ac75ea3efd0b66ef3b2f164a30ce0a0eb32db352798d223081.
//
// Solidity: event OperatorGroupCleared(uint256 indexed groupId)
func (_MetaRegistry *MetaRegistryFilterer) FilterOperatorGroupCleared(opts *bind.FilterOpts, groupId []*big.Int) (*MetaRegistryOperatorGroupClearedIterator, error) {

	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _MetaRegistry.contract.FilterLogs(opts, "OperatorGroupCleared", groupIdRule)
	if err != nil {
		return nil, err
	}
	return &MetaRegistryOperatorGroupClearedIterator{contract: _MetaRegistry.contract, event: "OperatorGroupCleared", logs: logs, sub: sub}, nil
}

// WatchOperatorGroupCleared is a free log subscription operation binding the contract event 0x9ca6c49aaad163ac75ea3efd0b66ef3b2f164a30ce0a0eb32db352798d223081.
//
// Solidity: event OperatorGroupCleared(uint256 indexed groupId)
func (_MetaRegistry *MetaRegistryFilterer) WatchOperatorGroupCleared(opts *bind.WatchOpts, sink chan<- *MetaRegistryOperatorGroupCleared, groupId []*big.Int) (event.Subscription, error) {

	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _MetaRegistry.contract.WatchLogs(opts, "OperatorGroupCleared", groupIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaRegistryOperatorGroupCleared)
				if err := _MetaRegistry.contract.UnpackLog(event, "OperatorGroupCleared", log); err != nil {
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

// ParseOperatorGroupCleared is a log parse operation binding the contract event 0x9ca6c49aaad163ac75ea3efd0b66ef3b2f164a30ce0a0eb32db352798d223081.
//
// Solidity: event OperatorGroupCleared(uint256 indexed groupId)
func (_MetaRegistry *MetaRegistryFilterer) ParseOperatorGroupCleared(log types.Log) (*MetaRegistryOperatorGroupCleared, error) {
	event := new(MetaRegistryOperatorGroupCleared)
	if err := _MetaRegistry.contract.UnpackLog(event, "OperatorGroupCleared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaRegistryOperatorGroupCreatedIterator is returned from FilterOperatorGroupCreated and is used to iterate over the raw logs and unpacked data for OperatorGroupCreated events raised by the MetaRegistry contract.
type MetaRegistryOperatorGroupCreatedIterator struct {
	Event *MetaRegistryOperatorGroupCreated // Event containing the contract specifics and raw log

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
func (it *MetaRegistryOperatorGroupCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaRegistryOperatorGroupCreated)
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
		it.Event = new(MetaRegistryOperatorGroupCreated)
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
func (it *MetaRegistryOperatorGroupCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaRegistryOperatorGroupCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaRegistryOperatorGroupCreated represents a OperatorGroupCreated event raised by the MetaRegistry contract.
type MetaRegistryOperatorGroupCreated struct {
	GroupId   *big.Int
	GroupInfo IMetaRegistryOperatorGroup
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOperatorGroupCreated is a free log retrieval operation binding the contract event 0x4ce64360c4e6921537ecff38795cd28d67eaf7bad79eb0098d23a29b7d0988c5.
//
// Solidity: event OperatorGroupCreated(uint256 indexed groupId, ((uint64,uint16)[],(bytes)[]) groupInfo)
func (_MetaRegistry *MetaRegistryFilterer) FilterOperatorGroupCreated(opts *bind.FilterOpts, groupId []*big.Int) (*MetaRegistryOperatorGroupCreatedIterator, error) {

	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _MetaRegistry.contract.FilterLogs(opts, "OperatorGroupCreated", groupIdRule)
	if err != nil {
		return nil, err
	}
	return &MetaRegistryOperatorGroupCreatedIterator{contract: _MetaRegistry.contract, event: "OperatorGroupCreated", logs: logs, sub: sub}, nil
}

// WatchOperatorGroupCreated is a free log subscription operation binding the contract event 0x4ce64360c4e6921537ecff38795cd28d67eaf7bad79eb0098d23a29b7d0988c5.
//
// Solidity: event OperatorGroupCreated(uint256 indexed groupId, ((uint64,uint16)[],(bytes)[]) groupInfo)
func (_MetaRegistry *MetaRegistryFilterer) WatchOperatorGroupCreated(opts *bind.WatchOpts, sink chan<- *MetaRegistryOperatorGroupCreated, groupId []*big.Int) (event.Subscription, error) {

	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _MetaRegistry.contract.WatchLogs(opts, "OperatorGroupCreated", groupIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaRegistryOperatorGroupCreated)
				if err := _MetaRegistry.contract.UnpackLog(event, "OperatorGroupCreated", log); err != nil {
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

// ParseOperatorGroupCreated is a log parse operation binding the contract event 0x4ce64360c4e6921537ecff38795cd28d67eaf7bad79eb0098d23a29b7d0988c5.
//
// Solidity: event OperatorGroupCreated(uint256 indexed groupId, ((uint64,uint16)[],(bytes)[]) groupInfo)
func (_MetaRegistry *MetaRegistryFilterer) ParseOperatorGroupCreated(log types.Log) (*MetaRegistryOperatorGroupCreated, error) {
	event := new(MetaRegistryOperatorGroupCreated)
	if err := _MetaRegistry.contract.UnpackLog(event, "OperatorGroupCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaRegistryOperatorGroupUpdatedIterator is returned from FilterOperatorGroupUpdated and is used to iterate over the raw logs and unpacked data for OperatorGroupUpdated events raised by the MetaRegistry contract.
type MetaRegistryOperatorGroupUpdatedIterator struct {
	Event *MetaRegistryOperatorGroupUpdated // Event containing the contract specifics and raw log

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
func (it *MetaRegistryOperatorGroupUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaRegistryOperatorGroupUpdated)
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
		it.Event = new(MetaRegistryOperatorGroupUpdated)
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
func (it *MetaRegistryOperatorGroupUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaRegistryOperatorGroupUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaRegistryOperatorGroupUpdated represents a OperatorGroupUpdated event raised by the MetaRegistry contract.
type MetaRegistryOperatorGroupUpdated struct {
	GroupId   *big.Int
	GroupInfo IMetaRegistryOperatorGroup
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOperatorGroupUpdated is a free log retrieval operation binding the contract event 0xfe2fcc8be7d25f7502080d4872f1e2a4263a8ccbe6a414f8715bca670ee56ea4.
//
// Solidity: event OperatorGroupUpdated(uint256 indexed groupId, ((uint64,uint16)[],(bytes)[]) groupInfo)
func (_MetaRegistry *MetaRegistryFilterer) FilterOperatorGroupUpdated(opts *bind.FilterOpts, groupId []*big.Int) (*MetaRegistryOperatorGroupUpdatedIterator, error) {

	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _MetaRegistry.contract.FilterLogs(opts, "OperatorGroupUpdated", groupIdRule)
	if err != nil {
		return nil, err
	}
	return &MetaRegistryOperatorGroupUpdatedIterator{contract: _MetaRegistry.contract, event: "OperatorGroupUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorGroupUpdated is a free log subscription operation binding the contract event 0xfe2fcc8be7d25f7502080d4872f1e2a4263a8ccbe6a414f8715bca670ee56ea4.
//
// Solidity: event OperatorGroupUpdated(uint256 indexed groupId, ((uint64,uint16)[],(bytes)[]) groupInfo)
func (_MetaRegistry *MetaRegistryFilterer) WatchOperatorGroupUpdated(opts *bind.WatchOpts, sink chan<- *MetaRegistryOperatorGroupUpdated, groupId []*big.Int) (event.Subscription, error) {

	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _MetaRegistry.contract.WatchLogs(opts, "OperatorGroupUpdated", groupIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaRegistryOperatorGroupUpdated)
				if err := _MetaRegistry.contract.UnpackLog(event, "OperatorGroupUpdated", log); err != nil {
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

// ParseOperatorGroupUpdated is a log parse operation binding the contract event 0xfe2fcc8be7d25f7502080d4872f1e2a4263a8ccbe6a414f8715bca670ee56ea4.
//
// Solidity: event OperatorGroupUpdated(uint256 indexed groupId, ((uint64,uint16)[],(bytes)[]) groupInfo)
func (_MetaRegistry *MetaRegistryFilterer) ParseOperatorGroupUpdated(log types.Log) (*MetaRegistryOperatorGroupUpdated, error) {
	event := new(MetaRegistryOperatorGroupUpdated)
	if err := _MetaRegistry.contract.UnpackLog(event, "OperatorGroupUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaRegistryOperatorMetadataSetIterator is returned from FilterOperatorMetadataSet and is used to iterate over the raw logs and unpacked data for OperatorMetadataSet events raised by the MetaRegistry contract.
type MetaRegistryOperatorMetadataSetIterator struct {
	Event *MetaRegistryOperatorMetadataSet // Event containing the contract specifics and raw log

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
func (it *MetaRegistryOperatorMetadataSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaRegistryOperatorMetadataSet)
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
		it.Event = new(MetaRegistryOperatorMetadataSet)
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
func (it *MetaRegistryOperatorMetadataSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaRegistryOperatorMetadataSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaRegistryOperatorMetadataSet represents a OperatorMetadataSet event raised by the MetaRegistry contract.
type MetaRegistryOperatorMetadataSet struct {
	NodeOperatorId *big.Int
	Metadata       OperatorMetadata
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorMetadataSet is a free log retrieval operation binding the contract event 0x27c4b5559e9eb92ed9a6cba3ab99a51ef6b41103e9eaae8877132a3d680dc8c5.
//
// Solidity: event OperatorMetadataSet(uint256 indexed nodeOperatorId, (string,string,bool) metadata)
func (_MetaRegistry *MetaRegistryFilterer) FilterOperatorMetadataSet(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*MetaRegistryOperatorMetadataSetIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _MetaRegistry.contract.FilterLogs(opts, "OperatorMetadataSet", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &MetaRegistryOperatorMetadataSetIterator{contract: _MetaRegistry.contract, event: "OperatorMetadataSet", logs: logs, sub: sub}, nil
}

// WatchOperatorMetadataSet is a free log subscription operation binding the contract event 0x27c4b5559e9eb92ed9a6cba3ab99a51ef6b41103e9eaae8877132a3d680dc8c5.
//
// Solidity: event OperatorMetadataSet(uint256 indexed nodeOperatorId, (string,string,bool) metadata)
func (_MetaRegistry *MetaRegistryFilterer) WatchOperatorMetadataSet(opts *bind.WatchOpts, sink chan<- *MetaRegistryOperatorMetadataSet, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _MetaRegistry.contract.WatchLogs(opts, "OperatorMetadataSet", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaRegistryOperatorMetadataSet)
				if err := _MetaRegistry.contract.UnpackLog(event, "OperatorMetadataSet", log); err != nil {
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

// ParseOperatorMetadataSet is a log parse operation binding the contract event 0x27c4b5559e9eb92ed9a6cba3ab99a51ef6b41103e9eaae8877132a3d680dc8c5.
//
// Solidity: event OperatorMetadataSet(uint256 indexed nodeOperatorId, (string,string,bool) metadata)
func (_MetaRegistry *MetaRegistryFilterer) ParseOperatorMetadataSet(log types.Log) (*MetaRegistryOperatorMetadataSet, error) {
	event := new(MetaRegistryOperatorMetadataSet)
	if err := _MetaRegistry.contract.UnpackLog(event, "OperatorMetadataSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the MetaRegistry contract.
type MetaRegistryRoleAdminChangedIterator struct {
	Event *MetaRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MetaRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaRegistryRoleAdminChanged)
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
		it.Event = new(MetaRegistryRoleAdminChanged)
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
func (it *MetaRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the MetaRegistry contract.
type MetaRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MetaRegistry *MetaRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MetaRegistryRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MetaRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MetaRegistryRoleAdminChangedIterator{contract: _MetaRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MetaRegistry *MetaRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MetaRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MetaRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaRegistryRoleAdminChanged)
				if err := _MetaRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MetaRegistry *MetaRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*MetaRegistryRoleAdminChanged, error) {
	event := new(MetaRegistryRoleAdminChanged)
	if err := _MetaRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the MetaRegistry contract.
type MetaRegistryRoleGrantedIterator struct {
	Event *MetaRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *MetaRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaRegistryRoleGranted)
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
		it.Event = new(MetaRegistryRoleGranted)
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
func (it *MetaRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaRegistryRoleGranted represents a RoleGranted event raised by the MetaRegistry contract.
type MetaRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MetaRegistry *MetaRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MetaRegistryRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MetaRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MetaRegistryRoleGrantedIterator{contract: _MetaRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MetaRegistry *MetaRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MetaRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MetaRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaRegistryRoleGranted)
				if err := _MetaRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MetaRegistry *MetaRegistryFilterer) ParseRoleGranted(log types.Log) (*MetaRegistryRoleGranted, error) {
	event := new(MetaRegistryRoleGranted)
	if err := _MetaRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the MetaRegistry contract.
type MetaRegistryRoleRevokedIterator struct {
	Event *MetaRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *MetaRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaRegistryRoleRevoked)
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
		it.Event = new(MetaRegistryRoleRevoked)
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
func (it *MetaRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaRegistryRoleRevoked represents a RoleRevoked event raised by the MetaRegistry contract.
type MetaRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MetaRegistry *MetaRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MetaRegistryRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MetaRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MetaRegistryRoleRevokedIterator{contract: _MetaRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MetaRegistry *MetaRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MetaRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MetaRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaRegistryRoleRevoked)
				if err := _MetaRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MetaRegistry *MetaRegistryFilterer) ParseRoleRevoked(log types.Log) (*MetaRegistryRoleRevoked, error) {
	event := new(MetaRegistryRoleRevoked)
	if err := _MetaRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

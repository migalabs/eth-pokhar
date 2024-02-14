// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lido

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

// NodeOperatorsRegistryMetaData contains all meta data concerning the NodeOperatorsRegistry contract.
var NodeOperatorsRegistryMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"hasInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"},{\"name\":\"_keysCount\",\"type\":\"uint256\"},{\"name\":\"_publicKeys\",\"type\":\"bytes\"},{\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"addSigningKeys\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getType\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_script\",\"type\":\"bytes\"}],\"name\":\"getEVMScriptExecutor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"clearNodeOperatorPenalty\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRecoveryVault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_offset\",\"type\":\"uint256\"},{\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getNodeOperatorIds\",\"outputs\":[{\"name\":\"nodeOperatorIds\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"},{\"name\":\"_offset\",\"type\":\"uint256\"},{\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getSigningKeys\",\"outputs\":[{\"name\":\"pubkeys\",\"type\":\"bytes\"},{\"name\":\"signatures\",\"type\":\"bytes\"},{\"name\":\"used\",\"type\":\"bool[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"},{\"name\":\"_fromIndex\",\"type\":\"uint256\"},{\"name\":\"_keysCount\",\"type\":\"uint256\"}],\"name\":\"removeSigningKeysOperatorBH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getNodeOperatorIsActive\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setNodeOperatorName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_totalRewardShares\",\"type\":\"uint256\"}],\"name\":\"getRewardsDistribution\",\"outputs\":[{\"name\":\"recipients\",\"type\":\"address[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\"},{\"name\":\"penalized\",\"type\":\"bool[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_indexFrom\",\"type\":\"uint256\"},{\"name\":\"_indexTo\",\"type\":\"uint256\"}],\"name\":\"invalidateReadyToDepositKeysRange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_locator\",\"type\":\"address\"},{\"name\":\"_type\",\"type\":\"bytes32\"},{\"name\":\"_stuckPenaltyDelay\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_delay\",\"type\":\"uint256\"}],\"name\":\"setStuckPenaltyDelay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStuckPenaltyDelay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"removeSigningKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"},{\"name\":\"_fromIndex\",\"type\":\"uint256\"},{\"name\":\"_keysCount\",\"type\":\"uint256\"}],\"name\":\"removeSigningKeys\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"isOperatorPenalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"deactivateNodeOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"allowRecoverability\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STAKING_ROUTER_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"},{\"name\":\"_keysCount\",\"type\":\"uint256\"},{\"name\":\"_publicKeys\",\"type\":\"bytes\"},{\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"addSigningKeysOperatorBH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"appId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getActiveNodeOperatorsCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_rewardAddress\",\"type\":\"address\"}],\"name\":\"addNodeOperator\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getContractVersion\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitializationBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getUnusedSigningKeyCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"onRewardsMinted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MANAGE_NODE_OPERATOR_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"onWithdrawalCredentialsChanged\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"activateNodeOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"},{\"name\":\"_rewardAddress\",\"type\":\"address\"}],\"name\":\"setNodeOperatorRewardAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"},{\"name\":\"_fullInfo\",\"type\":\"bool\"}],\"name\":\"getNodeOperator\",\"outputs\":[{\"name\":\"active\",\"type\":\"bool\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"rewardAddress\",\"type\":\"address\"},{\"name\":\"stakingLimit\",\"type\":\"uint64\"},{\"name\":\"stoppedValidators\",\"type\":\"uint64\"},{\"name\":\"totalSigningKeys\",\"type\":\"uint64\"},{\"name\":\"usedSigningKeys\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_locator\",\"type\":\"address\"},{\"name\":\"_type\",\"type\":\"bytes32\"},{\"name\":\"_stuckPenaltyDelay\",\"type\":\"uint256\"}],\"name\":\"finalizeUpgrade_v2\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStakingModuleSummary\",\"outputs\":[{\"name\":\"totalExitedValidators\",\"type\":\"uint256\"},{\"name\":\"totalDepositedValidators\",\"type\":\"uint256\"},{\"name\":\"depositableValidatorsCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeOperatorIds\",\"type\":\"bytes\"},{\"name\":\"_exitedValidatorsCounts\",\"type\":\"bytes\"}],\"name\":\"updateExitedValidatorsCount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeOperatorIds\",\"type\":\"bytes\"},{\"name\":\"_stuckValidatorsCounts\",\"type\":\"bytes\"}],\"name\":\"updateStuckValidatorsCount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"transferToVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"bytes32\"},{\"name\":\"_params\",\"type\":\"uint256[]\"}],\"name\":\"canPerform\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"},{\"name\":\"_refundedValidatorsCount\",\"type\":\"uint256\"}],\"name\":\"updateRefundedValidatorsCount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEVMScriptRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNodeOperatorsCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"},{\"name\":\"_isTargetLimitActive\",\"type\":\"bool\"},{\"name\":\"_targetLimit\",\"type\":\"uint256\"}],\"name\":\"updateTargetValidatorsLimits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"},{\"name\":\"_vettedSigningKeysCount\",\"type\":\"uint64\"}],\"name\":\"setNodeOperatorStakingLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getNodeOperatorSummary\",\"outputs\":[{\"name\":\"isTargetLimitActive\",\"type\":\"bool\"},{\"name\":\"targetValidatorsCount\",\"type\":\"uint256\"},{\"name\":\"stuckValidatorsCount\",\"type\":\"uint256\"},{\"name\":\"refundedValidatorsCount\",\"type\":\"uint256\"},{\"name\":\"stuckPenaltyEndTimestamp\",\"type\":\"uint256\"},{\"name\":\"totalExitedValidators\",\"type\":\"uint256\"},{\"name\":\"totalDepositedValidators\",\"type\":\"uint256\"},{\"name\":\"depositableValidatorsCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getSigningKey\",\"outputs\":[{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"depositSignature\",\"type\":\"bytes\"},{\"name\":\"used\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_NODE_OPERATOR_NAME_LENGTH\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_depositsCount\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"obtainDepositData\",\"outputs\":[{\"name\":\"publicKeys\",\"type\":\"bytes\"},{\"name\":\"signatures\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getKeysOpIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kernel\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLocator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SET_NODE_OPERATOR_LIMIT_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getTotalSigningKeyCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPetrified\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_STUCK_PENALTY_DELAY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"onExitedAndStuckValidatorsCountsUpdated\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_NODE_OPERATORS_COUNT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"removeSigningKeyOperatorBH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"},{\"name\":\"_exitedValidatorsCount\",\"type\":\"uint256\"},{\"name\":\"_stuckValidatorsCount\",\"type\":\"uint256\"}],\"name\":\"unsafeUpdateValidatorsCount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MANAGE_SIGNING_KEYS\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"isOperatorPenaltyCleared\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"rewardAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"stakingLimit\",\"type\":\"uint64\"}],\"name\":\"NodeOperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"NodeOperatorActiveSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NodeOperatorNameSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rewardAddress\",\"type\":\"address\"}],\"name\":\"NodeOperatorRewardAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalKeysTrimmed\",\"type\":\"uint64\"}],\"name\":\"NodeOperatorTotalKeysTrimmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"keysOpIndex\",\"type\":\"uint256\"}],\"name\":\"KeysOpIndexSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"moduleType\",\"type\":\"bytes32\"}],\"name\":\"StakingModuleTypeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"rewardAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sharesAmount\",\"type\":\"uint256\"}],\"name\":\"RewardsDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"locatorAddress\",\"type\":\"address\"}],\"name\":\"LocatorContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"approvedValidatorsCount\",\"type\":\"uint256\"}],\"name\":\"VettedSigningKeysCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"depositedValidatorsCount\",\"type\":\"uint256\"}],\"name\":\"DepositedSigningKeysCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"exitedValidatorsCount\",\"type\":\"uint256\"}],\"name\":\"ExitedSigningKeysCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalValidatorsCount\",\"type\":\"uint256\"}],\"name\":\"TotalSigningKeysCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"NonceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"stuckPenaltyDelay\",\"type\":\"uint256\"}],\"name\":\"StuckPenaltyDelayChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"stuckValidatorsCount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"refundedValidatorsCount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"stuckPenaltyEndTimestamp\",\"type\":\"uint256\"}],\"name\":\"StuckPenaltyStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"targetValidatorsCount\",\"type\":\"uint256\"}],\"name\":\"TargetValidatorsCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"recipientAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sharesPenalizedAmount\",\"type\":\"uint256\"}],\"name\":\"NodeOperatorPenalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"ContractVersionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"script\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"input\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ScriptResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RecoverToVault\",\"type\":\"event\"}]",
}

// NodeOperatorsRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use NodeOperatorsRegistryMetaData.ABI instead.
var NodeOperatorsRegistryABI = NodeOperatorsRegistryMetaData.ABI

// NodeOperatorsRegistry is an auto generated Go binding around an Ethereum contract.
type NodeOperatorsRegistry struct {
	NodeOperatorsRegistryCaller     // Read-only binding to the contract
	NodeOperatorsRegistryTransactor // Write-only binding to the contract
	NodeOperatorsRegistryFilterer   // Log filterer for contract events
}

// NodeOperatorsRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeOperatorsRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeOperatorsRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeOperatorsRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeOperatorsRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeOperatorsRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeOperatorsRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeOperatorsRegistrySession struct {
	Contract     *NodeOperatorsRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NodeOperatorsRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeOperatorsRegistryCallerSession struct {
	Contract *NodeOperatorsRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// NodeOperatorsRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeOperatorsRegistryTransactorSession struct {
	Contract     *NodeOperatorsRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// NodeOperatorsRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeOperatorsRegistryRaw struct {
	Contract *NodeOperatorsRegistry // Generic contract binding to access the raw methods on
}

// NodeOperatorsRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeOperatorsRegistryCallerRaw struct {
	Contract *NodeOperatorsRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// NodeOperatorsRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeOperatorsRegistryTransactorRaw struct {
	Contract *NodeOperatorsRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeOperatorsRegistry creates a new instance of NodeOperatorsRegistry, bound to a specific deployed contract.
func NewNodeOperatorsRegistry(address common.Address, backend bind.ContractBackend) (*NodeOperatorsRegistry, error) {
	contract, err := bindNodeOperatorsRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistry{NodeOperatorsRegistryCaller: NodeOperatorsRegistryCaller{contract: contract}, NodeOperatorsRegistryTransactor: NodeOperatorsRegistryTransactor{contract: contract}, NodeOperatorsRegistryFilterer: NodeOperatorsRegistryFilterer{contract: contract}}, nil
}

// NewNodeOperatorsRegistryCaller creates a new read-only instance of NodeOperatorsRegistry, bound to a specific deployed contract.
func NewNodeOperatorsRegistryCaller(address common.Address, caller bind.ContractCaller) (*NodeOperatorsRegistryCaller, error) {
	contract, err := bindNodeOperatorsRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryCaller{contract: contract}, nil
}

// NewNodeOperatorsRegistryTransactor creates a new write-only instance of NodeOperatorsRegistry, bound to a specific deployed contract.
func NewNodeOperatorsRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeOperatorsRegistryTransactor, error) {
	contract, err := bindNodeOperatorsRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryTransactor{contract: contract}, nil
}

// NewNodeOperatorsRegistryFilterer creates a new log filterer instance of NodeOperatorsRegistry, bound to a specific deployed contract.
func NewNodeOperatorsRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeOperatorsRegistryFilterer, error) {
	contract, err := bindNodeOperatorsRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryFilterer{contract: contract}, nil
}

// bindNodeOperatorsRegistry binds a generic wrapper to an already deployed contract.
func bindNodeOperatorsRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeOperatorsRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeOperatorsRegistry *NodeOperatorsRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeOperatorsRegistry.Contract.NodeOperatorsRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeOperatorsRegistry *NodeOperatorsRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.NodeOperatorsRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeOperatorsRegistry *NodeOperatorsRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.NodeOperatorsRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeOperatorsRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.contract.Transact(opts, method, params...)
}

// MANAGENODEOPERATORROLE is a free data retrieval call binding the contract method 0x8ece9995.
//
// Solidity: function MANAGE_NODE_OPERATOR_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) MANAGENODEOPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "MANAGE_NODE_OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGENODEOPERATORROLE is a free data retrieval call binding the contract method 0x8ece9995.
//
// Solidity: function MANAGE_NODE_OPERATOR_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) MANAGENODEOPERATORROLE() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.MANAGENODEOPERATORROLE(&_NodeOperatorsRegistry.CallOpts)
}

// MANAGENODEOPERATORROLE is a free data retrieval call binding the contract method 0x8ece9995.
//
// Solidity: function MANAGE_NODE_OPERATOR_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) MANAGENODEOPERATORROLE() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.MANAGENODEOPERATORROLE(&_NodeOperatorsRegistry.CallOpts)
}

// MANAGESIGNINGKEYS is a free data retrieval call binding the contract method 0xf31bd9c1.
//
// Solidity: function MANAGE_SIGNING_KEYS() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) MANAGESIGNINGKEYS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "MANAGE_SIGNING_KEYS")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGESIGNINGKEYS is a free data retrieval call binding the contract method 0xf31bd9c1.
//
// Solidity: function MANAGE_SIGNING_KEYS() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) MANAGESIGNINGKEYS() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.MANAGESIGNINGKEYS(&_NodeOperatorsRegistry.CallOpts)
}

// MANAGESIGNINGKEYS is a free data retrieval call binding the contract method 0xf31bd9c1.
//
// Solidity: function MANAGE_SIGNING_KEYS() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) MANAGESIGNINGKEYS() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.MANAGESIGNINGKEYS(&_NodeOperatorsRegistry.CallOpts)
}

// MAXNODEOPERATORSCOUNT is a free data retrieval call binding the contract method 0xec5af3a4.
//
// Solidity: function MAX_NODE_OPERATORS_COUNT() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) MAXNODEOPERATORSCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "MAX_NODE_OPERATORS_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXNODEOPERATORSCOUNT is a free data retrieval call binding the contract method 0xec5af3a4.
//
// Solidity: function MAX_NODE_OPERATORS_COUNT() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) MAXNODEOPERATORSCOUNT() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.MAXNODEOPERATORSCOUNT(&_NodeOperatorsRegistry.CallOpts)
}

// MAXNODEOPERATORSCOUNT is a free data retrieval call binding the contract method 0xec5af3a4.
//
// Solidity: function MAX_NODE_OPERATORS_COUNT() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) MAXNODEOPERATORSCOUNT() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.MAXNODEOPERATORSCOUNT(&_NodeOperatorsRegistry.CallOpts)
}

// MAXNODEOPERATORNAMELENGTH is a free data retrieval call binding the contract method 0xb4971833.
//
// Solidity: function MAX_NODE_OPERATOR_NAME_LENGTH() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) MAXNODEOPERATORNAMELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "MAX_NODE_OPERATOR_NAME_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXNODEOPERATORNAMELENGTH is a free data retrieval call binding the contract method 0xb4971833.
//
// Solidity: function MAX_NODE_OPERATOR_NAME_LENGTH() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) MAXNODEOPERATORNAMELENGTH() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.MAXNODEOPERATORNAMELENGTH(&_NodeOperatorsRegistry.CallOpts)
}

// MAXNODEOPERATORNAMELENGTH is a free data retrieval call binding the contract method 0xb4971833.
//
// Solidity: function MAX_NODE_OPERATOR_NAME_LENGTH() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) MAXNODEOPERATORNAMELENGTH() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.MAXNODEOPERATORNAMELENGTH(&_NodeOperatorsRegistry.CallOpts)
}

// MAXSTUCKPENALTYDELAY is a free data retrieval call binding the contract method 0xe204d09b.
//
// Solidity: function MAX_STUCK_PENALTY_DELAY() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) MAXSTUCKPENALTYDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "MAX_STUCK_PENALTY_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSTUCKPENALTYDELAY is a free data retrieval call binding the contract method 0xe204d09b.
//
// Solidity: function MAX_STUCK_PENALTY_DELAY() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) MAXSTUCKPENALTYDELAY() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.MAXSTUCKPENALTYDELAY(&_NodeOperatorsRegistry.CallOpts)
}

// MAXSTUCKPENALTYDELAY is a free data retrieval call binding the contract method 0xe204d09b.
//
// Solidity: function MAX_STUCK_PENALTY_DELAY() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) MAXSTUCKPENALTYDELAY() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.MAXSTUCKPENALTYDELAY(&_NodeOperatorsRegistry.CallOpts)
}

// SETNODEOPERATORLIMITROLE is a free data retrieval call binding the contract method 0xd8e71cd1.
//
// Solidity: function SET_NODE_OPERATOR_LIMIT_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) SETNODEOPERATORLIMITROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "SET_NODE_OPERATOR_LIMIT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETNODEOPERATORLIMITROLE is a free data retrieval call binding the contract method 0xd8e71cd1.
//
// Solidity: function SET_NODE_OPERATOR_LIMIT_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) SETNODEOPERATORLIMITROLE() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.SETNODEOPERATORLIMITROLE(&_NodeOperatorsRegistry.CallOpts)
}

// SETNODEOPERATORLIMITROLE is a free data retrieval call binding the contract method 0xd8e71cd1.
//
// Solidity: function SET_NODE_OPERATOR_LIMIT_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) SETNODEOPERATORLIMITROLE() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.SETNODEOPERATORLIMITROLE(&_NodeOperatorsRegistry.CallOpts)
}

// STAKINGROUTERROLE is a free data retrieval call binding the contract method 0x80231f15.
//
// Solidity: function STAKING_ROUTER_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) STAKINGROUTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "STAKING_ROUTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGROUTERROLE is a free data retrieval call binding the contract method 0x80231f15.
//
// Solidity: function STAKING_ROUTER_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) STAKINGROUTERROLE() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.STAKINGROUTERROLE(&_NodeOperatorsRegistry.CallOpts)
}

// STAKINGROUTERROLE is a free data retrieval call binding the contract method 0x80231f15.
//
// Solidity: function STAKING_ROUTER_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) STAKINGROUTERROLE() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.STAKINGROUTERROLE(&_NodeOperatorsRegistry.CallOpts)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) AllowRecoverability(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "allowRecoverability", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) AllowRecoverability(token common.Address) (bool, error) {
	return _NodeOperatorsRegistry.Contract.AllowRecoverability(&_NodeOperatorsRegistry.CallOpts, token)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) AllowRecoverability(token common.Address) (bool, error) {
	return _NodeOperatorsRegistry.Contract.AllowRecoverability(&_NodeOperatorsRegistry.CallOpts, token)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) AppId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "appId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) AppId() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.AppId(&_NodeOperatorsRegistry.CallOpts)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) AppId() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.AppId(&_NodeOperatorsRegistry.CallOpts)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) CanPerform(opts *bind.CallOpts, _sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "canPerform", _sender, _role, _params)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _NodeOperatorsRegistry.Contract.CanPerform(&_NodeOperatorsRegistry.CallOpts, _sender, _role, _params)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _NodeOperatorsRegistry.Contract.CanPerform(&_NodeOperatorsRegistry.CallOpts, _sender, _role, _params)
}

// GetActiveNodeOperatorsCount is a free data retrieval call binding the contract method 0x8469cbd3.
//
// Solidity: function getActiveNodeOperatorsCount() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetActiveNodeOperatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getActiveNodeOperatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveNodeOperatorsCount is a free data retrieval call binding the contract method 0x8469cbd3.
//
// Solidity: function getActiveNodeOperatorsCount() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetActiveNodeOperatorsCount() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetActiveNodeOperatorsCount(&_NodeOperatorsRegistry.CallOpts)
}

// GetActiveNodeOperatorsCount is a free data retrieval call binding the contract method 0x8469cbd3.
//
// Solidity: function getActiveNodeOperatorsCount() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetActiveNodeOperatorsCount() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetActiveNodeOperatorsCount(&_NodeOperatorsRegistry.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetContractVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getContractVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetContractVersion() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetContractVersion(&_NodeOperatorsRegistry.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetContractVersion() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetContractVersion(&_NodeOperatorsRegistry.CallOpts)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetEVMScriptExecutor(opts *bind.CallOpts, _script []byte) (common.Address, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getEVMScriptExecutor", _script)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _NodeOperatorsRegistry.Contract.GetEVMScriptExecutor(&_NodeOperatorsRegistry.CallOpts, _script)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _NodeOperatorsRegistry.Contract.GetEVMScriptExecutor(&_NodeOperatorsRegistry.CallOpts, _script)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetEVMScriptRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getEVMScriptRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetEVMScriptRegistry() (common.Address, error) {
	return _NodeOperatorsRegistry.Contract.GetEVMScriptRegistry(&_NodeOperatorsRegistry.CallOpts)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetEVMScriptRegistry() (common.Address, error) {
	return _NodeOperatorsRegistry.Contract.GetEVMScriptRegistry(&_NodeOperatorsRegistry.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetInitializationBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getInitializationBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetInitializationBlock() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetInitializationBlock(&_NodeOperatorsRegistry.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetInitializationBlock() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetInitializationBlock(&_NodeOperatorsRegistry.CallOpts)
}

// GetKeysOpIndex is a free data retrieval call binding the contract method 0xd07442f1.
//
// Solidity: function getKeysOpIndex() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetKeysOpIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getKeysOpIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetKeysOpIndex is a free data retrieval call binding the contract method 0xd07442f1.
//
// Solidity: function getKeysOpIndex() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetKeysOpIndex() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetKeysOpIndex(&_NodeOperatorsRegistry.CallOpts)
}

// GetKeysOpIndex is a free data retrieval call binding the contract method 0xd07442f1.
//
// Solidity: function getKeysOpIndex() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetKeysOpIndex() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetKeysOpIndex(&_NodeOperatorsRegistry.CallOpts)
}

// GetLocator is a free data retrieval call binding the contract method 0xd8343dcb.
//
// Solidity: function getLocator() view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetLocator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getLocator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLocator is a free data retrieval call binding the contract method 0xd8343dcb.
//
// Solidity: function getLocator() view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetLocator() (common.Address, error) {
	return _NodeOperatorsRegistry.Contract.GetLocator(&_NodeOperatorsRegistry.CallOpts)
}

// GetLocator is a free data retrieval call binding the contract method 0xd8343dcb.
//
// Solidity: function getLocator() view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetLocator() (common.Address, error) {
	return _NodeOperatorsRegistry.Contract.GetLocator(&_NodeOperatorsRegistry.CallOpts)
}

// GetNodeOperator is a free data retrieval call binding the contract method 0x9a56983c.
//
// Solidity: function getNodeOperator(uint256 _nodeOperatorId, bool _fullInfo) view returns(bool active, string name, address rewardAddress, uint64 stakingLimit, uint64 stoppedValidators, uint64 totalSigningKeys, uint64 usedSigningKeys)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetNodeOperator(opts *bind.CallOpts, _nodeOperatorId *big.Int, _fullInfo bool) (struct {
	Active            bool
	Name              string
	RewardAddress     common.Address
	StakingLimit      uint64
	StoppedValidators uint64
	TotalSigningKeys  uint64
	UsedSigningKeys   uint64
}, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getNodeOperator", _nodeOperatorId, _fullInfo)

	outstruct := new(struct {
		Active            bool
		Name              string
		RewardAddress     common.Address
		StakingLimit      uint64
		StoppedValidators uint64
		TotalSigningKeys  uint64
		UsedSigningKeys   uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Active = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.RewardAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.StakingLimit = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.StoppedValidators = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.TotalSigningKeys = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.UsedSigningKeys = *abi.ConvertType(out[6], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetNodeOperator is a free data retrieval call binding the contract method 0x9a56983c.
//
// Solidity: function getNodeOperator(uint256 _nodeOperatorId, bool _fullInfo) view returns(bool active, string name, address rewardAddress, uint64 stakingLimit, uint64 stoppedValidators, uint64 totalSigningKeys, uint64 usedSigningKeys)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetNodeOperator(_nodeOperatorId *big.Int, _fullInfo bool) (struct {
	Active            bool
	Name              string
	RewardAddress     common.Address
	StakingLimit      uint64
	StoppedValidators uint64
	TotalSigningKeys  uint64
	UsedSigningKeys   uint64
}, error) {
	return _NodeOperatorsRegistry.Contract.GetNodeOperator(&_NodeOperatorsRegistry.CallOpts, _nodeOperatorId, _fullInfo)
}

// GetNodeOperator is a free data retrieval call binding the contract method 0x9a56983c.
//
// Solidity: function getNodeOperator(uint256 _nodeOperatorId, bool _fullInfo) view returns(bool active, string name, address rewardAddress, uint64 stakingLimit, uint64 stoppedValidators, uint64 totalSigningKeys, uint64 usedSigningKeys)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetNodeOperator(_nodeOperatorId *big.Int, _fullInfo bool) (struct {
	Active            bool
	Name              string
	RewardAddress     common.Address
	StakingLimit      uint64
	StoppedValidators uint64
	TotalSigningKeys  uint64
	UsedSigningKeys   uint64
}, error) {
	return _NodeOperatorsRegistry.Contract.GetNodeOperator(&_NodeOperatorsRegistry.CallOpts, _nodeOperatorId, _fullInfo)
}

// GetNodeOperatorIds is a free data retrieval call binding the contract method 0x4febc81b.
//
// Solidity: function getNodeOperatorIds(uint256 _offset, uint256 _limit) view returns(uint256[] nodeOperatorIds)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetNodeOperatorIds(opts *bind.CallOpts, _offset *big.Int, _limit *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getNodeOperatorIds", _offset, _limit)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetNodeOperatorIds is a free data retrieval call binding the contract method 0x4febc81b.
//
// Solidity: function getNodeOperatorIds(uint256 _offset, uint256 _limit) view returns(uint256[] nodeOperatorIds)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetNodeOperatorIds(_offset *big.Int, _limit *big.Int) ([]*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetNodeOperatorIds(&_NodeOperatorsRegistry.CallOpts, _offset, _limit)
}

// GetNodeOperatorIds is a free data retrieval call binding the contract method 0x4febc81b.
//
// Solidity: function getNodeOperatorIds(uint256 _offset, uint256 _limit) view returns(uint256[] nodeOperatorIds)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetNodeOperatorIds(_offset *big.Int, _limit *big.Int) ([]*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetNodeOperatorIds(&_NodeOperatorsRegistry.CallOpts, _offset, _limit)
}

// GetNodeOperatorIsActive is a free data retrieval call binding the contract method 0x5e2fb908.
//
// Solidity: function getNodeOperatorIsActive(uint256 _nodeOperatorId) view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetNodeOperatorIsActive(opts *bind.CallOpts, _nodeOperatorId *big.Int) (bool, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getNodeOperatorIsActive", _nodeOperatorId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetNodeOperatorIsActive is a free data retrieval call binding the contract method 0x5e2fb908.
//
// Solidity: function getNodeOperatorIsActive(uint256 _nodeOperatorId) view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetNodeOperatorIsActive(_nodeOperatorId *big.Int) (bool, error) {
	return _NodeOperatorsRegistry.Contract.GetNodeOperatorIsActive(&_NodeOperatorsRegistry.CallOpts, _nodeOperatorId)
}

// GetNodeOperatorIsActive is a free data retrieval call binding the contract method 0x5e2fb908.
//
// Solidity: function getNodeOperatorIsActive(uint256 _nodeOperatorId) view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetNodeOperatorIsActive(_nodeOperatorId *big.Int) (bool, error) {
	return _NodeOperatorsRegistry.Contract.GetNodeOperatorIsActive(&_NodeOperatorsRegistry.CallOpts, _nodeOperatorId)
}

// GetNodeOperatorSummary is a free data retrieval call binding the contract method 0xb3076c3c.
//
// Solidity: function getNodeOperatorSummary(uint256 _nodeOperatorId) view returns(bool isTargetLimitActive, uint256 targetValidatorsCount, uint256 stuckValidatorsCount, uint256 refundedValidatorsCount, uint256 stuckPenaltyEndTimestamp, uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetNodeOperatorSummary(opts *bind.CallOpts, _nodeOperatorId *big.Int) (struct {
	IsTargetLimitActive        bool
	TargetValidatorsCount      *big.Int
	StuckValidatorsCount       *big.Int
	RefundedValidatorsCount    *big.Int
	StuckPenaltyEndTimestamp   *big.Int
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getNodeOperatorSummary", _nodeOperatorId)

	outstruct := new(struct {
		IsTargetLimitActive        bool
		TargetValidatorsCount      *big.Int
		StuckValidatorsCount       *big.Int
		RefundedValidatorsCount    *big.Int
		StuckPenaltyEndTimestamp   *big.Int
		TotalExitedValidators      *big.Int
		TotalDepositedValidators   *big.Int
		DepositableValidatorsCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsTargetLimitActive = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.TargetValidatorsCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StuckValidatorsCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RefundedValidatorsCount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StuckPenaltyEndTimestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TotalExitedValidators = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TotalDepositedValidators = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.DepositableValidatorsCount = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNodeOperatorSummary is a free data retrieval call binding the contract method 0xb3076c3c.
//
// Solidity: function getNodeOperatorSummary(uint256 _nodeOperatorId) view returns(bool isTargetLimitActive, uint256 targetValidatorsCount, uint256 stuckValidatorsCount, uint256 refundedValidatorsCount, uint256 stuckPenaltyEndTimestamp, uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetNodeOperatorSummary(_nodeOperatorId *big.Int) (struct {
	IsTargetLimitActive        bool
	TargetValidatorsCount      *big.Int
	StuckValidatorsCount       *big.Int
	RefundedValidatorsCount    *big.Int
	StuckPenaltyEndTimestamp   *big.Int
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	return _NodeOperatorsRegistry.Contract.GetNodeOperatorSummary(&_NodeOperatorsRegistry.CallOpts, _nodeOperatorId)
}

// GetNodeOperatorSummary is a free data retrieval call binding the contract method 0xb3076c3c.
//
// Solidity: function getNodeOperatorSummary(uint256 _nodeOperatorId) view returns(bool isTargetLimitActive, uint256 targetValidatorsCount, uint256 stuckValidatorsCount, uint256 refundedValidatorsCount, uint256 stuckPenaltyEndTimestamp, uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetNodeOperatorSummary(_nodeOperatorId *big.Int) (struct {
	IsTargetLimitActive        bool
	TargetValidatorsCount      *big.Int
	StuckValidatorsCount       *big.Int
	RefundedValidatorsCount    *big.Int
	StuckPenaltyEndTimestamp   *big.Int
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	return _NodeOperatorsRegistry.Contract.GetNodeOperatorSummary(&_NodeOperatorsRegistry.CallOpts, _nodeOperatorId)
}

// GetNodeOperatorsCount is a free data retrieval call binding the contract method 0xa70c70e4.
//
// Solidity: function getNodeOperatorsCount() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetNodeOperatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getNodeOperatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeOperatorsCount is a free data retrieval call binding the contract method 0xa70c70e4.
//
// Solidity: function getNodeOperatorsCount() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetNodeOperatorsCount() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetNodeOperatorsCount(&_NodeOperatorsRegistry.CallOpts)
}

// GetNodeOperatorsCount is a free data retrieval call binding the contract method 0xa70c70e4.
//
// Solidity: function getNodeOperatorsCount() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetNodeOperatorsCount() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetNodeOperatorsCount(&_NodeOperatorsRegistry.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetNonce() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetNonce(&_NodeOperatorsRegistry.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetNonce() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetNonce(&_NodeOperatorsRegistry.CallOpts)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetRecoveryVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getRecoveryVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetRecoveryVault() (common.Address, error) {
	return _NodeOperatorsRegistry.Contract.GetRecoveryVault(&_NodeOperatorsRegistry.CallOpts)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetRecoveryVault() (common.Address, error) {
	return _NodeOperatorsRegistry.Contract.GetRecoveryVault(&_NodeOperatorsRegistry.CallOpts)
}

// GetRewardsDistribution is a free data retrieval call binding the contract method 0x62dcfda1.
//
// Solidity: function getRewardsDistribution(uint256 _totalRewardShares) view returns(address[] recipients, uint256[] shares, bool[] penalized)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetRewardsDistribution(opts *bind.CallOpts, _totalRewardShares *big.Int) (struct {
	Recipients []common.Address
	Shares     []*big.Int
	Penalized  []bool
}, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getRewardsDistribution", _totalRewardShares)

	outstruct := new(struct {
		Recipients []common.Address
		Shares     []*big.Int
		Penalized  []bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Recipients = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Shares = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Penalized = *abi.ConvertType(out[2], new([]bool)).(*[]bool)

	return *outstruct, err

}

// GetRewardsDistribution is a free data retrieval call binding the contract method 0x62dcfda1.
//
// Solidity: function getRewardsDistribution(uint256 _totalRewardShares) view returns(address[] recipients, uint256[] shares, bool[] penalized)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetRewardsDistribution(_totalRewardShares *big.Int) (struct {
	Recipients []common.Address
	Shares     []*big.Int
	Penalized  []bool
}, error) {
	return _NodeOperatorsRegistry.Contract.GetRewardsDistribution(&_NodeOperatorsRegistry.CallOpts, _totalRewardShares)
}

// GetRewardsDistribution is a free data retrieval call binding the contract method 0x62dcfda1.
//
// Solidity: function getRewardsDistribution(uint256 _totalRewardShares) view returns(address[] recipients, uint256[] shares, bool[] penalized)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetRewardsDistribution(_totalRewardShares *big.Int) (struct {
	Recipients []common.Address
	Shares     []*big.Int
	Penalized  []bool
}, error) {
	return _NodeOperatorsRegistry.Contract.GetRewardsDistribution(&_NodeOperatorsRegistry.CallOpts, _totalRewardShares)
}

// GetSigningKey is a free data retrieval call binding the contract method 0xb449402a.
//
// Solidity: function getSigningKey(uint256 _nodeOperatorId, uint256 _index) view returns(bytes key, bytes depositSignature, bool used)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetSigningKey(opts *bind.CallOpts, _nodeOperatorId *big.Int, _index *big.Int) (struct {
	Key              []byte
	DepositSignature []byte
	Used             bool
}, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getSigningKey", _nodeOperatorId, _index)

	outstruct := new(struct {
		Key              []byte
		DepositSignature []byte
		Used             bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Key = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.DepositSignature = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.Used = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// GetSigningKey is a free data retrieval call binding the contract method 0xb449402a.
//
// Solidity: function getSigningKey(uint256 _nodeOperatorId, uint256 _index) view returns(bytes key, bytes depositSignature, bool used)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetSigningKey(_nodeOperatorId *big.Int, _index *big.Int) (struct {
	Key              []byte
	DepositSignature []byte
	Used             bool
}, error) {
	return _NodeOperatorsRegistry.Contract.GetSigningKey(&_NodeOperatorsRegistry.CallOpts, _nodeOperatorId, _index)
}

// GetSigningKey is a free data retrieval call binding the contract method 0xb449402a.
//
// Solidity: function getSigningKey(uint256 _nodeOperatorId, uint256 _index) view returns(bytes key, bytes depositSignature, bool used)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetSigningKey(_nodeOperatorId *big.Int, _index *big.Int) (struct {
	Key              []byte
	DepositSignature []byte
	Used             bool
}, error) {
	return _NodeOperatorsRegistry.Contract.GetSigningKey(&_NodeOperatorsRegistry.CallOpts, _nodeOperatorId, _index)
}

// GetSigningKeys is a free data retrieval call binding the contract method 0x59e25c12.
//
// Solidity: function getSigningKeys(uint256 _nodeOperatorId, uint256 _offset, uint256 _limit) view returns(bytes pubkeys, bytes signatures, bool[] used)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetSigningKeys(opts *bind.CallOpts, _nodeOperatorId *big.Int, _offset *big.Int, _limit *big.Int) (struct {
	Pubkeys    []byte
	Signatures []byte
	Used       []bool
}, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getSigningKeys", _nodeOperatorId, _offset, _limit)

	outstruct := new(struct {
		Pubkeys    []byte
		Signatures []byte
		Used       []bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pubkeys = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Signatures = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.Used = *abi.ConvertType(out[2], new([]bool)).(*[]bool)

	return *outstruct, err

}

// GetSigningKeys is a free data retrieval call binding the contract method 0x59e25c12.
//
// Solidity: function getSigningKeys(uint256 _nodeOperatorId, uint256 _offset, uint256 _limit) view returns(bytes pubkeys, bytes signatures, bool[] used)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetSigningKeys(_nodeOperatorId *big.Int, _offset *big.Int, _limit *big.Int) (struct {
	Pubkeys    []byte
	Signatures []byte
	Used       []bool
}, error) {
	return _NodeOperatorsRegistry.Contract.GetSigningKeys(&_NodeOperatorsRegistry.CallOpts, _nodeOperatorId, _offset, _limit)
}

// GetSigningKeys is a free data retrieval call binding the contract method 0x59e25c12.
//
// Solidity: function getSigningKeys(uint256 _nodeOperatorId, uint256 _offset, uint256 _limit) view returns(bytes pubkeys, bytes signatures, bool[] used)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetSigningKeys(_nodeOperatorId *big.Int, _offset *big.Int, _limit *big.Int) (struct {
	Pubkeys    []byte
	Signatures []byte
	Used       []bool
}, error) {
	return _NodeOperatorsRegistry.Contract.GetSigningKeys(&_NodeOperatorsRegistry.CallOpts, _nodeOperatorId, _offset, _limit)
}

// GetStakingModuleSummary is a free data retrieval call binding the contract method 0x9abddf09.
//
// Solidity: function getStakingModuleSummary() view returns(uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetStakingModuleSummary(opts *bind.CallOpts) (struct {
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getStakingModuleSummary")

	outstruct := new(struct {
		TotalExitedValidators      *big.Int
		TotalDepositedValidators   *big.Int
		DepositableValidatorsCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalExitedValidators = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalDepositedValidators = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DepositableValidatorsCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStakingModuleSummary is a free data retrieval call binding the contract method 0x9abddf09.
//
// Solidity: function getStakingModuleSummary() view returns(uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetStakingModuleSummary() (struct {
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	return _NodeOperatorsRegistry.Contract.GetStakingModuleSummary(&_NodeOperatorsRegistry.CallOpts)
}

// GetStakingModuleSummary is a free data retrieval call binding the contract method 0x9abddf09.
//
// Solidity: function getStakingModuleSummary() view returns(uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetStakingModuleSummary() (struct {
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	return _NodeOperatorsRegistry.Contract.GetStakingModuleSummary(&_NodeOperatorsRegistry.CallOpts)
}

// GetStuckPenaltyDelay is a free data retrieval call binding the contract method 0x6da7d0a7.
//
// Solidity: function getStuckPenaltyDelay() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetStuckPenaltyDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getStuckPenaltyDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStuckPenaltyDelay is a free data retrieval call binding the contract method 0x6da7d0a7.
//
// Solidity: function getStuckPenaltyDelay() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetStuckPenaltyDelay() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetStuckPenaltyDelay(&_NodeOperatorsRegistry.CallOpts)
}

// GetStuckPenaltyDelay is a free data retrieval call binding the contract method 0x6da7d0a7.
//
// Solidity: function getStuckPenaltyDelay() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetStuckPenaltyDelay() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetStuckPenaltyDelay(&_NodeOperatorsRegistry.CallOpts)
}

// GetTotalSigningKeyCount is a free data retrieval call binding the contract method 0xdb9887ea.
//
// Solidity: function getTotalSigningKeyCount(uint256 _nodeOperatorId) view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetTotalSigningKeyCount(opts *bind.CallOpts, _nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getTotalSigningKeyCount", _nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSigningKeyCount is a free data retrieval call binding the contract method 0xdb9887ea.
//
// Solidity: function getTotalSigningKeyCount(uint256 _nodeOperatorId) view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetTotalSigningKeyCount(_nodeOperatorId *big.Int) (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetTotalSigningKeyCount(&_NodeOperatorsRegistry.CallOpts, _nodeOperatorId)
}

// GetTotalSigningKeyCount is a free data retrieval call binding the contract method 0xdb9887ea.
//
// Solidity: function getTotalSigningKeyCount(uint256 _nodeOperatorId) view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetTotalSigningKeyCount(_nodeOperatorId *big.Int) (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetTotalSigningKeyCount(&_NodeOperatorsRegistry.CallOpts, _nodeOperatorId)
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetType() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.GetType(&_NodeOperatorsRegistry.CallOpts)
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetType() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.GetType(&_NodeOperatorsRegistry.CallOpts)
}

// GetUnusedSigningKeyCount is a free data retrieval call binding the contract method 0x8ca7c052.
//
// Solidity: function getUnusedSigningKeyCount(uint256 _nodeOperatorId) view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetUnusedSigningKeyCount(opts *bind.CallOpts, _nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getUnusedSigningKeyCount", _nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnusedSigningKeyCount is a free data retrieval call binding the contract method 0x8ca7c052.
//
// Solidity: function getUnusedSigningKeyCount(uint256 _nodeOperatorId) view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetUnusedSigningKeyCount(_nodeOperatorId *big.Int) (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetUnusedSigningKeyCount(&_NodeOperatorsRegistry.CallOpts, _nodeOperatorId)
}

// GetUnusedSigningKeyCount is a free data retrieval call binding the contract method 0x8ca7c052.
//
// Solidity: function getUnusedSigningKeyCount(uint256 _nodeOperatorId) view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetUnusedSigningKeyCount(_nodeOperatorId *big.Int) (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetUnusedSigningKeyCount(&_NodeOperatorsRegistry.CallOpts, _nodeOperatorId)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) HasInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "hasInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) HasInitialized() (bool, error) {
	return _NodeOperatorsRegistry.Contract.HasInitialized(&_NodeOperatorsRegistry.CallOpts)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) HasInitialized() (bool, error) {
	return _NodeOperatorsRegistry.Contract.HasInitialized(&_NodeOperatorsRegistry.CallOpts)
}

// IsOperatorPenalized is a free data retrieval call binding the contract method 0x75049ad8.
//
// Solidity: function isOperatorPenalized(uint256 _nodeOperatorId) view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) IsOperatorPenalized(opts *bind.CallOpts, _nodeOperatorId *big.Int) (bool, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "isOperatorPenalized", _nodeOperatorId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorPenalized is a free data retrieval call binding the contract method 0x75049ad8.
//
// Solidity: function isOperatorPenalized(uint256 _nodeOperatorId) view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) IsOperatorPenalized(_nodeOperatorId *big.Int) (bool, error) {
	return _NodeOperatorsRegistry.Contract.IsOperatorPenalized(&_NodeOperatorsRegistry.CallOpts, _nodeOperatorId)
}

// IsOperatorPenalized is a free data retrieval call binding the contract method 0x75049ad8.
//
// Solidity: function isOperatorPenalized(uint256 _nodeOperatorId) view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) IsOperatorPenalized(_nodeOperatorId *big.Int) (bool, error) {
	return _NodeOperatorsRegistry.Contract.IsOperatorPenalized(&_NodeOperatorsRegistry.CallOpts, _nodeOperatorId)
}

// IsOperatorPenaltyCleared is a free data retrieval call binding the contract method 0xfbc77ef1.
//
// Solidity: function isOperatorPenaltyCleared(uint256 _nodeOperatorId) view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) IsOperatorPenaltyCleared(opts *bind.CallOpts, _nodeOperatorId *big.Int) (bool, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "isOperatorPenaltyCleared", _nodeOperatorId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorPenaltyCleared is a free data retrieval call binding the contract method 0xfbc77ef1.
//
// Solidity: function isOperatorPenaltyCleared(uint256 _nodeOperatorId) view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) IsOperatorPenaltyCleared(_nodeOperatorId *big.Int) (bool, error) {
	return _NodeOperatorsRegistry.Contract.IsOperatorPenaltyCleared(&_NodeOperatorsRegistry.CallOpts, _nodeOperatorId)
}

// IsOperatorPenaltyCleared is a free data retrieval call binding the contract method 0xfbc77ef1.
//
// Solidity: function isOperatorPenaltyCleared(uint256 _nodeOperatorId) view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) IsOperatorPenaltyCleared(_nodeOperatorId *big.Int) (bool, error) {
	return _NodeOperatorsRegistry.Contract.IsOperatorPenaltyCleared(&_NodeOperatorsRegistry.CallOpts, _nodeOperatorId)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) IsPetrified(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "isPetrified")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) IsPetrified() (bool, error) {
	return _NodeOperatorsRegistry.Contract.IsPetrified(&_NodeOperatorsRegistry.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) IsPetrified() (bool, error) {
	return _NodeOperatorsRegistry.Contract.IsPetrified(&_NodeOperatorsRegistry.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) Kernel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "kernel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) Kernel() (common.Address, error) {
	return _NodeOperatorsRegistry.Contract.Kernel(&_NodeOperatorsRegistry.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) Kernel() (common.Address, error) {
	return _NodeOperatorsRegistry.Contract.Kernel(&_NodeOperatorsRegistry.CallOpts)
}

// OnRewardsMinted is a free data retrieval call binding the contract method 0x8d7e4017.
//
// Solidity: function onRewardsMinted(uint256 ) view returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) OnRewardsMinted(opts *bind.CallOpts, arg0 *big.Int) error {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "onRewardsMinted", arg0)

	if err != nil {
		return err
	}

	return err

}

// OnRewardsMinted is a free data retrieval call binding the contract method 0x8d7e4017.
//
// Solidity: function onRewardsMinted(uint256 ) view returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) OnRewardsMinted(arg0 *big.Int) error {
	return _NodeOperatorsRegistry.Contract.OnRewardsMinted(&_NodeOperatorsRegistry.CallOpts, arg0)
}

// OnRewardsMinted is a free data retrieval call binding the contract method 0x8d7e4017.
//
// Solidity: function onRewardsMinted(uint256 ) view returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) OnRewardsMinted(arg0 *big.Int) error {
	return _NodeOperatorsRegistry.Contract.OnRewardsMinted(&_NodeOperatorsRegistry.CallOpts, arg0)
}

// ActivateNodeOperator is a paid mutator transaction binding the contract method 0x91dcd6b2.
//
// Solidity: function activateNodeOperator(uint256 _nodeOperatorId) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) ActivateNodeOperator(opts *bind.TransactOpts, _nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "activateNodeOperator", _nodeOperatorId)
}

// ActivateNodeOperator is a paid mutator transaction binding the contract method 0x91dcd6b2.
//
// Solidity: function activateNodeOperator(uint256 _nodeOperatorId) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) ActivateNodeOperator(_nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.ActivateNodeOperator(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId)
}

// ActivateNodeOperator is a paid mutator transaction binding the contract method 0x91dcd6b2.
//
// Solidity: function activateNodeOperator(uint256 _nodeOperatorId) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) ActivateNodeOperator(_nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.ActivateNodeOperator(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId)
}

// AddNodeOperator is a paid mutator transaction binding the contract method 0x85fa63d7.
//
// Solidity: function addNodeOperator(string _name, address _rewardAddress) returns(uint256 id)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) AddNodeOperator(opts *bind.TransactOpts, _name string, _rewardAddress common.Address) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "addNodeOperator", _name, _rewardAddress)
}

// AddNodeOperator is a paid mutator transaction binding the contract method 0x85fa63d7.
//
// Solidity: function addNodeOperator(string _name, address _rewardAddress) returns(uint256 id)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) AddNodeOperator(_name string, _rewardAddress common.Address) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.AddNodeOperator(&_NodeOperatorsRegistry.TransactOpts, _name, _rewardAddress)
}

// AddNodeOperator is a paid mutator transaction binding the contract method 0x85fa63d7.
//
// Solidity: function addNodeOperator(string _name, address _rewardAddress) returns(uint256 id)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) AddNodeOperator(_name string, _rewardAddress common.Address) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.AddNodeOperator(&_NodeOperatorsRegistry.TransactOpts, _name, _rewardAddress)
}

// AddSigningKeys is a paid mutator transaction binding the contract method 0x096b7b35.
//
// Solidity: function addSigningKeys(uint256 _nodeOperatorId, uint256 _keysCount, bytes _publicKeys, bytes _signatures) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) AddSigningKeys(opts *bind.TransactOpts, _nodeOperatorId *big.Int, _keysCount *big.Int, _publicKeys []byte, _signatures []byte) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "addSigningKeys", _nodeOperatorId, _keysCount, _publicKeys, _signatures)
}

// AddSigningKeys is a paid mutator transaction binding the contract method 0x096b7b35.
//
// Solidity: function addSigningKeys(uint256 _nodeOperatorId, uint256 _keysCount, bytes _publicKeys, bytes _signatures) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) AddSigningKeys(_nodeOperatorId *big.Int, _keysCount *big.Int, _publicKeys []byte, _signatures []byte) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.AddSigningKeys(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId, _keysCount, _publicKeys, _signatures)
}

// AddSigningKeys is a paid mutator transaction binding the contract method 0x096b7b35.
//
// Solidity: function addSigningKeys(uint256 _nodeOperatorId, uint256 _keysCount, bytes _publicKeys, bytes _signatures) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) AddSigningKeys(_nodeOperatorId *big.Int, _keysCount *big.Int, _publicKeys []byte, _signatures []byte) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.AddSigningKeys(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId, _keysCount, _publicKeys, _signatures)
}

// AddSigningKeysOperatorBH is a paid mutator transaction binding the contract method 0x805911ae.
//
// Solidity: function addSigningKeysOperatorBH(uint256 _nodeOperatorId, uint256 _keysCount, bytes _publicKeys, bytes _signatures) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) AddSigningKeysOperatorBH(opts *bind.TransactOpts, _nodeOperatorId *big.Int, _keysCount *big.Int, _publicKeys []byte, _signatures []byte) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "addSigningKeysOperatorBH", _nodeOperatorId, _keysCount, _publicKeys, _signatures)
}

// AddSigningKeysOperatorBH is a paid mutator transaction binding the contract method 0x805911ae.
//
// Solidity: function addSigningKeysOperatorBH(uint256 _nodeOperatorId, uint256 _keysCount, bytes _publicKeys, bytes _signatures) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) AddSigningKeysOperatorBH(_nodeOperatorId *big.Int, _keysCount *big.Int, _publicKeys []byte, _signatures []byte) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.AddSigningKeysOperatorBH(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId, _keysCount, _publicKeys, _signatures)
}

// AddSigningKeysOperatorBH is a paid mutator transaction binding the contract method 0x805911ae.
//
// Solidity: function addSigningKeysOperatorBH(uint256 _nodeOperatorId, uint256 _keysCount, bytes _publicKeys, bytes _signatures) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) AddSigningKeysOperatorBH(_nodeOperatorId *big.Int, _keysCount *big.Int, _publicKeys []byte, _signatures []byte) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.AddSigningKeysOperatorBH(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId, _keysCount, _publicKeys, _signatures)
}

// ClearNodeOperatorPenalty is a paid mutator transaction binding the contract method 0x30a90f01.
//
// Solidity: function clearNodeOperatorPenalty(uint256 _nodeOperatorId) returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) ClearNodeOperatorPenalty(opts *bind.TransactOpts, _nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "clearNodeOperatorPenalty", _nodeOperatorId)
}

// ClearNodeOperatorPenalty is a paid mutator transaction binding the contract method 0x30a90f01.
//
// Solidity: function clearNodeOperatorPenalty(uint256 _nodeOperatorId) returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) ClearNodeOperatorPenalty(_nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.ClearNodeOperatorPenalty(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId)
}

// ClearNodeOperatorPenalty is a paid mutator transaction binding the contract method 0x30a90f01.
//
// Solidity: function clearNodeOperatorPenalty(uint256 _nodeOperatorId) returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) ClearNodeOperatorPenalty(_nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.ClearNodeOperatorPenalty(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId)
}

// DeactivateNodeOperator is a paid mutator transaction binding the contract method 0x75a080d5.
//
// Solidity: function deactivateNodeOperator(uint256 _nodeOperatorId) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) DeactivateNodeOperator(opts *bind.TransactOpts, _nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "deactivateNodeOperator", _nodeOperatorId)
}

// DeactivateNodeOperator is a paid mutator transaction binding the contract method 0x75a080d5.
//
// Solidity: function deactivateNodeOperator(uint256 _nodeOperatorId) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) DeactivateNodeOperator(_nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.DeactivateNodeOperator(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId)
}

// DeactivateNodeOperator is a paid mutator transaction binding the contract method 0x75a080d5.
//
// Solidity: function deactivateNodeOperator(uint256 _nodeOperatorId) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) DeactivateNodeOperator(_nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.DeactivateNodeOperator(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0x9a7c2ade.
//
// Solidity: function finalizeUpgrade_v2(address _locator, bytes32 _type, uint256 _stuckPenaltyDelay) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) FinalizeUpgradeV2(opts *bind.TransactOpts, _locator common.Address, _type [32]byte, _stuckPenaltyDelay *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "finalizeUpgrade_v2", _locator, _type, _stuckPenaltyDelay)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0x9a7c2ade.
//
// Solidity: function finalizeUpgrade_v2(address _locator, bytes32 _type, uint256 _stuckPenaltyDelay) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) FinalizeUpgradeV2(_locator common.Address, _type [32]byte, _stuckPenaltyDelay *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.FinalizeUpgradeV2(&_NodeOperatorsRegistry.TransactOpts, _locator, _type, _stuckPenaltyDelay)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0x9a7c2ade.
//
// Solidity: function finalizeUpgrade_v2(address _locator, bytes32 _type, uint256 _stuckPenaltyDelay) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) FinalizeUpgradeV2(_locator common.Address, _type [32]byte, _stuckPenaltyDelay *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.FinalizeUpgradeV2(&_NodeOperatorsRegistry.TransactOpts, _locator, _type, _stuckPenaltyDelay)
}

// Initialize is a paid mutator transaction binding the contract method 0x684560a2.
//
// Solidity: function initialize(address _locator, bytes32 _type, uint256 _stuckPenaltyDelay) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) Initialize(opts *bind.TransactOpts, _locator common.Address, _type [32]byte, _stuckPenaltyDelay *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "initialize", _locator, _type, _stuckPenaltyDelay)
}

// Initialize is a paid mutator transaction binding the contract method 0x684560a2.
//
// Solidity: function initialize(address _locator, bytes32 _type, uint256 _stuckPenaltyDelay) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) Initialize(_locator common.Address, _type [32]byte, _stuckPenaltyDelay *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.Initialize(&_NodeOperatorsRegistry.TransactOpts, _locator, _type, _stuckPenaltyDelay)
}

// Initialize is a paid mutator transaction binding the contract method 0x684560a2.
//
// Solidity: function initialize(address _locator, bytes32 _type, uint256 _stuckPenaltyDelay) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) Initialize(_locator common.Address, _type [32]byte, _stuckPenaltyDelay *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.Initialize(&_NodeOperatorsRegistry.TransactOpts, _locator, _type, _stuckPenaltyDelay)
}

// InvalidateReadyToDepositKeysRange is a paid mutator transaction binding the contract method 0x65cc369a.
//
// Solidity: function invalidateReadyToDepositKeysRange(uint256 _indexFrom, uint256 _indexTo) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) InvalidateReadyToDepositKeysRange(opts *bind.TransactOpts, _indexFrom *big.Int, _indexTo *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "invalidateReadyToDepositKeysRange", _indexFrom, _indexTo)
}

// InvalidateReadyToDepositKeysRange is a paid mutator transaction binding the contract method 0x65cc369a.
//
// Solidity: function invalidateReadyToDepositKeysRange(uint256 _indexFrom, uint256 _indexTo) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) InvalidateReadyToDepositKeysRange(_indexFrom *big.Int, _indexTo *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.InvalidateReadyToDepositKeysRange(&_NodeOperatorsRegistry.TransactOpts, _indexFrom, _indexTo)
}

// InvalidateReadyToDepositKeysRange is a paid mutator transaction binding the contract method 0x65cc369a.
//
// Solidity: function invalidateReadyToDepositKeysRange(uint256 _indexFrom, uint256 _indexTo) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) InvalidateReadyToDepositKeysRange(_indexFrom *big.Int, _indexTo *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.InvalidateReadyToDepositKeysRange(&_NodeOperatorsRegistry.TransactOpts, _indexFrom, _indexTo)
}

// ObtainDepositData is a paid mutator transaction binding the contract method 0xbee41b58.
//
// Solidity: function obtainDepositData(uint256 _depositsCount, bytes ) returns(bytes publicKeys, bytes signatures)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) ObtainDepositData(opts *bind.TransactOpts, _depositsCount *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "obtainDepositData", _depositsCount, arg1)
}

// ObtainDepositData is a paid mutator transaction binding the contract method 0xbee41b58.
//
// Solidity: function obtainDepositData(uint256 _depositsCount, bytes ) returns(bytes publicKeys, bytes signatures)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) ObtainDepositData(_depositsCount *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.ObtainDepositData(&_NodeOperatorsRegistry.TransactOpts, _depositsCount, arg1)
}

// ObtainDepositData is a paid mutator transaction binding the contract method 0xbee41b58.
//
// Solidity: function obtainDepositData(uint256 _depositsCount, bytes ) returns(bytes publicKeys, bytes signatures)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) ObtainDepositData(_depositsCount *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.ObtainDepositData(&_NodeOperatorsRegistry.TransactOpts, _depositsCount, arg1)
}

// OnExitedAndStuckValidatorsCountsUpdated is a paid mutator transaction binding the contract method 0xe864299e.
//
// Solidity: function onExitedAndStuckValidatorsCountsUpdated() returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) OnExitedAndStuckValidatorsCountsUpdated(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "onExitedAndStuckValidatorsCountsUpdated")
}

// OnExitedAndStuckValidatorsCountsUpdated is a paid mutator transaction binding the contract method 0xe864299e.
//
// Solidity: function onExitedAndStuckValidatorsCountsUpdated() returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) OnExitedAndStuckValidatorsCountsUpdated() (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.OnExitedAndStuckValidatorsCountsUpdated(&_NodeOperatorsRegistry.TransactOpts)
}

// OnExitedAndStuckValidatorsCountsUpdated is a paid mutator transaction binding the contract method 0xe864299e.
//
// Solidity: function onExitedAndStuckValidatorsCountsUpdated() returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) OnExitedAndStuckValidatorsCountsUpdated() (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.OnExitedAndStuckValidatorsCountsUpdated(&_NodeOperatorsRegistry.TransactOpts)
}

// OnWithdrawalCredentialsChanged is a paid mutator transaction binding the contract method 0x90c09bdb.
//
// Solidity: function onWithdrawalCredentialsChanged() returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) OnWithdrawalCredentialsChanged(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "onWithdrawalCredentialsChanged")
}

// OnWithdrawalCredentialsChanged is a paid mutator transaction binding the contract method 0x90c09bdb.
//
// Solidity: function onWithdrawalCredentialsChanged() returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) OnWithdrawalCredentialsChanged() (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.OnWithdrawalCredentialsChanged(&_NodeOperatorsRegistry.TransactOpts)
}

// OnWithdrawalCredentialsChanged is a paid mutator transaction binding the contract method 0x90c09bdb.
//
// Solidity: function onWithdrawalCredentialsChanged() returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) OnWithdrawalCredentialsChanged() (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.OnWithdrawalCredentialsChanged(&_NodeOperatorsRegistry.TransactOpts)
}

// RemoveSigningKey is a paid mutator transaction binding the contract method 0x6ef355f1.
//
// Solidity: function removeSigningKey(uint256 _nodeOperatorId, uint256 _index) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) RemoveSigningKey(opts *bind.TransactOpts, _nodeOperatorId *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "removeSigningKey", _nodeOperatorId, _index)
}

// RemoveSigningKey is a paid mutator transaction binding the contract method 0x6ef355f1.
//
// Solidity: function removeSigningKey(uint256 _nodeOperatorId, uint256 _index) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) RemoveSigningKey(_nodeOperatorId *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.RemoveSigningKey(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId, _index)
}

// RemoveSigningKey is a paid mutator transaction binding the contract method 0x6ef355f1.
//
// Solidity: function removeSigningKey(uint256 _nodeOperatorId, uint256 _index) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) RemoveSigningKey(_nodeOperatorId *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.RemoveSigningKey(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId, _index)
}

// RemoveSigningKeyOperatorBH is a paid mutator transaction binding the contract method 0xed5cfa41.
//
// Solidity: function removeSigningKeyOperatorBH(uint256 _nodeOperatorId, uint256 _index) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) RemoveSigningKeyOperatorBH(opts *bind.TransactOpts, _nodeOperatorId *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "removeSigningKeyOperatorBH", _nodeOperatorId, _index)
}

// RemoveSigningKeyOperatorBH is a paid mutator transaction binding the contract method 0xed5cfa41.
//
// Solidity: function removeSigningKeyOperatorBH(uint256 _nodeOperatorId, uint256 _index) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) RemoveSigningKeyOperatorBH(_nodeOperatorId *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.RemoveSigningKeyOperatorBH(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId, _index)
}

// RemoveSigningKeyOperatorBH is a paid mutator transaction binding the contract method 0xed5cfa41.
//
// Solidity: function removeSigningKeyOperatorBH(uint256 _nodeOperatorId, uint256 _index) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) RemoveSigningKeyOperatorBH(_nodeOperatorId *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.RemoveSigningKeyOperatorBH(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId, _index)
}

// RemoveSigningKeys is a paid mutator transaction binding the contract method 0x7038b141.
//
// Solidity: function removeSigningKeys(uint256 _nodeOperatorId, uint256 _fromIndex, uint256 _keysCount) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) RemoveSigningKeys(opts *bind.TransactOpts, _nodeOperatorId *big.Int, _fromIndex *big.Int, _keysCount *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "removeSigningKeys", _nodeOperatorId, _fromIndex, _keysCount)
}

// RemoveSigningKeys is a paid mutator transaction binding the contract method 0x7038b141.
//
// Solidity: function removeSigningKeys(uint256 _nodeOperatorId, uint256 _fromIndex, uint256 _keysCount) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) RemoveSigningKeys(_nodeOperatorId *big.Int, _fromIndex *big.Int, _keysCount *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.RemoveSigningKeys(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId, _fromIndex, _keysCount)
}

// RemoveSigningKeys is a paid mutator transaction binding the contract method 0x7038b141.
//
// Solidity: function removeSigningKeys(uint256 _nodeOperatorId, uint256 _fromIndex, uint256 _keysCount) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) RemoveSigningKeys(_nodeOperatorId *big.Int, _fromIndex *big.Int, _keysCount *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.RemoveSigningKeys(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId, _fromIndex, _keysCount)
}

// RemoveSigningKeysOperatorBH is a paid mutator transaction binding the contract method 0x5ddde810.
//
// Solidity: function removeSigningKeysOperatorBH(uint256 _nodeOperatorId, uint256 _fromIndex, uint256 _keysCount) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) RemoveSigningKeysOperatorBH(opts *bind.TransactOpts, _nodeOperatorId *big.Int, _fromIndex *big.Int, _keysCount *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "removeSigningKeysOperatorBH", _nodeOperatorId, _fromIndex, _keysCount)
}

// RemoveSigningKeysOperatorBH is a paid mutator transaction binding the contract method 0x5ddde810.
//
// Solidity: function removeSigningKeysOperatorBH(uint256 _nodeOperatorId, uint256 _fromIndex, uint256 _keysCount) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) RemoveSigningKeysOperatorBH(_nodeOperatorId *big.Int, _fromIndex *big.Int, _keysCount *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.RemoveSigningKeysOperatorBH(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId, _fromIndex, _keysCount)
}

// RemoveSigningKeysOperatorBH is a paid mutator transaction binding the contract method 0x5ddde810.
//
// Solidity: function removeSigningKeysOperatorBH(uint256 _nodeOperatorId, uint256 _fromIndex, uint256 _keysCount) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) RemoveSigningKeysOperatorBH(_nodeOperatorId *big.Int, _fromIndex *big.Int, _keysCount *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.RemoveSigningKeysOperatorBH(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId, _fromIndex, _keysCount)
}

// SetNodeOperatorName is a paid mutator transaction binding the contract method 0x5e57d742.
//
// Solidity: function setNodeOperatorName(uint256 _nodeOperatorId, string _name) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) SetNodeOperatorName(opts *bind.TransactOpts, _nodeOperatorId *big.Int, _name string) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "setNodeOperatorName", _nodeOperatorId, _name)
}

// SetNodeOperatorName is a paid mutator transaction binding the contract method 0x5e57d742.
//
// Solidity: function setNodeOperatorName(uint256 _nodeOperatorId, string _name) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) SetNodeOperatorName(_nodeOperatorId *big.Int, _name string) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.SetNodeOperatorName(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId, _name)
}

// SetNodeOperatorName is a paid mutator transaction binding the contract method 0x5e57d742.
//
// Solidity: function setNodeOperatorName(uint256 _nodeOperatorId, string _name) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) SetNodeOperatorName(_nodeOperatorId *big.Int, _name string) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.SetNodeOperatorName(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId, _name)
}

// SetNodeOperatorRewardAddress is a paid mutator transaction binding the contract method 0x973e9328.
//
// Solidity: function setNodeOperatorRewardAddress(uint256 _nodeOperatorId, address _rewardAddress) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) SetNodeOperatorRewardAddress(opts *bind.TransactOpts, _nodeOperatorId *big.Int, _rewardAddress common.Address) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "setNodeOperatorRewardAddress", _nodeOperatorId, _rewardAddress)
}

// SetNodeOperatorRewardAddress is a paid mutator transaction binding the contract method 0x973e9328.
//
// Solidity: function setNodeOperatorRewardAddress(uint256 _nodeOperatorId, address _rewardAddress) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) SetNodeOperatorRewardAddress(_nodeOperatorId *big.Int, _rewardAddress common.Address) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.SetNodeOperatorRewardAddress(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId, _rewardAddress)
}

// SetNodeOperatorRewardAddress is a paid mutator transaction binding the contract method 0x973e9328.
//
// Solidity: function setNodeOperatorRewardAddress(uint256 _nodeOperatorId, address _rewardAddress) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) SetNodeOperatorRewardAddress(_nodeOperatorId *big.Int, _rewardAddress common.Address) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.SetNodeOperatorRewardAddress(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId, _rewardAddress)
}

// SetNodeOperatorStakingLimit is a paid mutator transaction binding the contract method 0xae962acf.
//
// Solidity: function setNodeOperatorStakingLimit(uint256 _nodeOperatorId, uint64 _vettedSigningKeysCount) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) SetNodeOperatorStakingLimit(opts *bind.TransactOpts, _nodeOperatorId *big.Int, _vettedSigningKeysCount uint64) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "setNodeOperatorStakingLimit", _nodeOperatorId, _vettedSigningKeysCount)
}

// SetNodeOperatorStakingLimit is a paid mutator transaction binding the contract method 0xae962acf.
//
// Solidity: function setNodeOperatorStakingLimit(uint256 _nodeOperatorId, uint64 _vettedSigningKeysCount) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) SetNodeOperatorStakingLimit(_nodeOperatorId *big.Int, _vettedSigningKeysCount uint64) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.SetNodeOperatorStakingLimit(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId, _vettedSigningKeysCount)
}

// SetNodeOperatorStakingLimit is a paid mutator transaction binding the contract method 0xae962acf.
//
// Solidity: function setNodeOperatorStakingLimit(uint256 _nodeOperatorId, uint64 _vettedSigningKeysCount) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) SetNodeOperatorStakingLimit(_nodeOperatorId *big.Int, _vettedSigningKeysCount uint64) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.SetNodeOperatorStakingLimit(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId, _vettedSigningKeysCount)
}

// SetStuckPenaltyDelay is a paid mutator transaction binding the contract method 0x6ccc7562.
//
// Solidity: function setStuckPenaltyDelay(uint256 _delay) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) SetStuckPenaltyDelay(opts *bind.TransactOpts, _delay *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "setStuckPenaltyDelay", _delay)
}

// SetStuckPenaltyDelay is a paid mutator transaction binding the contract method 0x6ccc7562.
//
// Solidity: function setStuckPenaltyDelay(uint256 _delay) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) SetStuckPenaltyDelay(_delay *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.SetStuckPenaltyDelay(&_NodeOperatorsRegistry.TransactOpts, _delay)
}

// SetStuckPenaltyDelay is a paid mutator transaction binding the contract method 0x6ccc7562.
//
// Solidity: function setStuckPenaltyDelay(uint256 _delay) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) SetStuckPenaltyDelay(_delay *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.SetStuckPenaltyDelay(&_NodeOperatorsRegistry.TransactOpts, _delay)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) TransferToVault(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "transferToVault", _token)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) TransferToVault(_token common.Address) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.TransferToVault(&_NodeOperatorsRegistry.TransactOpts, _token)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) TransferToVault(_token common.Address) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.TransferToVault(&_NodeOperatorsRegistry.TransactOpts, _token)
}

// UnsafeUpdateValidatorsCount is a paid mutator transaction binding the contract method 0xf2e2ca63.
//
// Solidity: function unsafeUpdateValidatorsCount(uint256 _nodeOperatorId, uint256 _exitedValidatorsCount, uint256 _stuckValidatorsCount) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) UnsafeUpdateValidatorsCount(opts *bind.TransactOpts, _nodeOperatorId *big.Int, _exitedValidatorsCount *big.Int, _stuckValidatorsCount *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "unsafeUpdateValidatorsCount", _nodeOperatorId, _exitedValidatorsCount, _stuckValidatorsCount)
}

// UnsafeUpdateValidatorsCount is a paid mutator transaction binding the contract method 0xf2e2ca63.
//
// Solidity: function unsafeUpdateValidatorsCount(uint256 _nodeOperatorId, uint256 _exitedValidatorsCount, uint256 _stuckValidatorsCount) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) UnsafeUpdateValidatorsCount(_nodeOperatorId *big.Int, _exitedValidatorsCount *big.Int, _stuckValidatorsCount *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.UnsafeUpdateValidatorsCount(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId, _exitedValidatorsCount, _stuckValidatorsCount)
}

// UnsafeUpdateValidatorsCount is a paid mutator transaction binding the contract method 0xf2e2ca63.
//
// Solidity: function unsafeUpdateValidatorsCount(uint256 _nodeOperatorId, uint256 _exitedValidatorsCount, uint256 _stuckValidatorsCount) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) UnsafeUpdateValidatorsCount(_nodeOperatorId *big.Int, _exitedValidatorsCount *big.Int, _stuckValidatorsCount *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.UnsafeUpdateValidatorsCount(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId, _exitedValidatorsCount, _stuckValidatorsCount)
}

// UpdateExitedValidatorsCount is a paid mutator transaction binding the contract method 0x9b00c146.
//
// Solidity: function updateExitedValidatorsCount(bytes _nodeOperatorIds, bytes _exitedValidatorsCounts) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) UpdateExitedValidatorsCount(opts *bind.TransactOpts, _nodeOperatorIds []byte, _exitedValidatorsCounts []byte) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "updateExitedValidatorsCount", _nodeOperatorIds, _exitedValidatorsCounts)
}

// UpdateExitedValidatorsCount is a paid mutator transaction binding the contract method 0x9b00c146.
//
// Solidity: function updateExitedValidatorsCount(bytes _nodeOperatorIds, bytes _exitedValidatorsCounts) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) UpdateExitedValidatorsCount(_nodeOperatorIds []byte, _exitedValidatorsCounts []byte) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.UpdateExitedValidatorsCount(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorIds, _exitedValidatorsCounts)
}

// UpdateExitedValidatorsCount is a paid mutator transaction binding the contract method 0x9b00c146.
//
// Solidity: function updateExitedValidatorsCount(bytes _nodeOperatorIds, bytes _exitedValidatorsCounts) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) UpdateExitedValidatorsCount(_nodeOperatorIds []byte, _exitedValidatorsCounts []byte) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.UpdateExitedValidatorsCount(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorIds, _exitedValidatorsCounts)
}

// UpdateRefundedValidatorsCount is a paid mutator transaction binding the contract method 0xa2e080f1.
//
// Solidity: function updateRefundedValidatorsCount(uint256 _nodeOperatorId, uint256 _refundedValidatorsCount) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) UpdateRefundedValidatorsCount(opts *bind.TransactOpts, _nodeOperatorId *big.Int, _refundedValidatorsCount *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "updateRefundedValidatorsCount", _nodeOperatorId, _refundedValidatorsCount)
}

// UpdateRefundedValidatorsCount is a paid mutator transaction binding the contract method 0xa2e080f1.
//
// Solidity: function updateRefundedValidatorsCount(uint256 _nodeOperatorId, uint256 _refundedValidatorsCount) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) UpdateRefundedValidatorsCount(_nodeOperatorId *big.Int, _refundedValidatorsCount *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.UpdateRefundedValidatorsCount(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId, _refundedValidatorsCount)
}

// UpdateRefundedValidatorsCount is a paid mutator transaction binding the contract method 0xa2e080f1.
//
// Solidity: function updateRefundedValidatorsCount(uint256 _nodeOperatorId, uint256 _refundedValidatorsCount) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) UpdateRefundedValidatorsCount(_nodeOperatorId *big.Int, _refundedValidatorsCount *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.UpdateRefundedValidatorsCount(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId, _refundedValidatorsCount)
}

// UpdateStuckValidatorsCount is a paid mutator transaction binding the contract method 0x9b3d1900.
//
// Solidity: function updateStuckValidatorsCount(bytes _nodeOperatorIds, bytes _stuckValidatorsCounts) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) UpdateStuckValidatorsCount(opts *bind.TransactOpts, _nodeOperatorIds []byte, _stuckValidatorsCounts []byte) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "updateStuckValidatorsCount", _nodeOperatorIds, _stuckValidatorsCounts)
}

// UpdateStuckValidatorsCount is a paid mutator transaction binding the contract method 0x9b3d1900.
//
// Solidity: function updateStuckValidatorsCount(bytes _nodeOperatorIds, bytes _stuckValidatorsCounts) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) UpdateStuckValidatorsCount(_nodeOperatorIds []byte, _stuckValidatorsCounts []byte) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.UpdateStuckValidatorsCount(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorIds, _stuckValidatorsCounts)
}

// UpdateStuckValidatorsCount is a paid mutator transaction binding the contract method 0x9b3d1900.
//
// Solidity: function updateStuckValidatorsCount(bytes _nodeOperatorIds, bytes _stuckValidatorsCounts) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) UpdateStuckValidatorsCount(_nodeOperatorIds []byte, _stuckValidatorsCounts []byte) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.UpdateStuckValidatorsCount(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorIds, _stuckValidatorsCounts)
}

// UpdateTargetValidatorsLimits is a paid mutator transaction binding the contract method 0xa9e7a846.
//
// Solidity: function updateTargetValidatorsLimits(uint256 _nodeOperatorId, bool _isTargetLimitActive, uint256 _targetLimit) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) UpdateTargetValidatorsLimits(opts *bind.TransactOpts, _nodeOperatorId *big.Int, _isTargetLimitActive bool, _targetLimit *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "updateTargetValidatorsLimits", _nodeOperatorId, _isTargetLimitActive, _targetLimit)
}

// UpdateTargetValidatorsLimits is a paid mutator transaction binding the contract method 0xa9e7a846.
//
// Solidity: function updateTargetValidatorsLimits(uint256 _nodeOperatorId, bool _isTargetLimitActive, uint256 _targetLimit) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) UpdateTargetValidatorsLimits(_nodeOperatorId *big.Int, _isTargetLimitActive bool, _targetLimit *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.UpdateTargetValidatorsLimits(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId, _isTargetLimitActive, _targetLimit)
}

// UpdateTargetValidatorsLimits is a paid mutator transaction binding the contract method 0xa9e7a846.
//
// Solidity: function updateTargetValidatorsLimits(uint256 _nodeOperatorId, bool _isTargetLimitActive, uint256 _targetLimit) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) UpdateTargetValidatorsLimits(_nodeOperatorId *big.Int, _isTargetLimitActive bool, _targetLimit *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.UpdateTargetValidatorsLimits(&_NodeOperatorsRegistry.TransactOpts, _nodeOperatorId, _isTargetLimitActive, _targetLimit)
}

// NodeOperatorsRegistryContractVersionSetIterator is returned from FilterContractVersionSet and is used to iterate over the raw logs and unpacked data for ContractVersionSet events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryContractVersionSetIterator struct {
	Event *NodeOperatorsRegistryContractVersionSet // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryContractVersionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryContractVersionSet)
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
		it.Event = new(NodeOperatorsRegistryContractVersionSet)
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
func (it *NodeOperatorsRegistryContractVersionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryContractVersionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryContractVersionSet represents a ContractVersionSet event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryContractVersionSet struct {
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractVersionSet is a free log retrieval operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterContractVersionSet(opts *bind.FilterOpts) (*NodeOperatorsRegistryContractVersionSetIterator, error) {

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryContractVersionSetIterator{contract: _NodeOperatorsRegistry.contract, event: "ContractVersionSet", logs: logs, sub: sub}, nil
}

// WatchContractVersionSet is a free log subscription operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchContractVersionSet(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryContractVersionSet) (event.Subscription, error) {

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryContractVersionSet)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
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

// ParseContractVersionSet is a log parse operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseContractVersionSet(log types.Log) (*NodeOperatorsRegistryContractVersionSet, error) {
	event := new(NodeOperatorsRegistryContractVersionSet)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryDepositedSigningKeysCountChangedIterator is returned from FilterDepositedSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for DepositedSigningKeysCountChanged events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryDepositedSigningKeysCountChangedIterator struct {
	Event *NodeOperatorsRegistryDepositedSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryDepositedSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryDepositedSigningKeysCountChanged)
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
		it.Event = new(NodeOperatorsRegistryDepositedSigningKeysCountChanged)
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
func (it *NodeOperatorsRegistryDepositedSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryDepositedSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryDepositedSigningKeysCountChanged represents a DepositedSigningKeysCountChanged event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryDepositedSigningKeysCountChanged struct {
	NodeOperatorId           *big.Int
	DepositedValidatorsCount *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterDepositedSigningKeysCountChanged is a free log retrieval operation binding the contract event 0x24eb1c9e765ba41accf9437300ea91ece5ed3f897ec3cdee0e9debd7fe309b78.
//
// Solidity: event DepositedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 depositedValidatorsCount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterDepositedSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*NodeOperatorsRegistryDepositedSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "DepositedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryDepositedSigningKeysCountChangedIterator{contract: _NodeOperatorsRegistry.contract, event: "DepositedSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchDepositedSigningKeysCountChanged is a free log subscription operation binding the contract event 0x24eb1c9e765ba41accf9437300ea91ece5ed3f897ec3cdee0e9debd7fe309b78.
//
// Solidity: event DepositedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 depositedValidatorsCount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchDepositedSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryDepositedSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "DepositedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryDepositedSigningKeysCountChanged)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "DepositedSigningKeysCountChanged", log); err != nil {
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

// ParseDepositedSigningKeysCountChanged is a log parse operation binding the contract event 0x24eb1c9e765ba41accf9437300ea91ece5ed3f897ec3cdee0e9debd7fe309b78.
//
// Solidity: event DepositedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 depositedValidatorsCount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseDepositedSigningKeysCountChanged(log types.Log) (*NodeOperatorsRegistryDepositedSigningKeysCountChanged, error) {
	event := new(NodeOperatorsRegistryDepositedSigningKeysCountChanged)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "DepositedSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryExitedSigningKeysCountChangedIterator is returned from FilterExitedSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for ExitedSigningKeysCountChanged events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryExitedSigningKeysCountChangedIterator struct {
	Event *NodeOperatorsRegistryExitedSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryExitedSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryExitedSigningKeysCountChanged)
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
		it.Event = new(NodeOperatorsRegistryExitedSigningKeysCountChanged)
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
func (it *NodeOperatorsRegistryExitedSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryExitedSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryExitedSigningKeysCountChanged represents a ExitedSigningKeysCountChanged event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryExitedSigningKeysCountChanged struct {
	NodeOperatorId        *big.Int
	ExitedValidatorsCount *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterExitedSigningKeysCountChanged is a free log retrieval operation binding the contract event 0x0f67960648751434ae86bf350db61194f387fda387e7f568b0ccd0ae0c220166.
//
// Solidity: event ExitedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 exitedValidatorsCount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterExitedSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*NodeOperatorsRegistryExitedSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "ExitedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryExitedSigningKeysCountChangedIterator{contract: _NodeOperatorsRegistry.contract, event: "ExitedSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchExitedSigningKeysCountChanged is a free log subscription operation binding the contract event 0x0f67960648751434ae86bf350db61194f387fda387e7f568b0ccd0ae0c220166.
//
// Solidity: event ExitedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 exitedValidatorsCount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchExitedSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryExitedSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "ExitedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryExitedSigningKeysCountChanged)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "ExitedSigningKeysCountChanged", log); err != nil {
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

// ParseExitedSigningKeysCountChanged is a log parse operation binding the contract event 0x0f67960648751434ae86bf350db61194f387fda387e7f568b0ccd0ae0c220166.
//
// Solidity: event ExitedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 exitedValidatorsCount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseExitedSigningKeysCountChanged(log types.Log) (*NodeOperatorsRegistryExitedSigningKeysCountChanged, error) {
	event := new(NodeOperatorsRegistryExitedSigningKeysCountChanged)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "ExitedSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryKeysOpIndexSetIterator is returned from FilterKeysOpIndexSet and is used to iterate over the raw logs and unpacked data for KeysOpIndexSet events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryKeysOpIndexSetIterator struct {
	Event *NodeOperatorsRegistryKeysOpIndexSet // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryKeysOpIndexSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryKeysOpIndexSet)
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
		it.Event = new(NodeOperatorsRegistryKeysOpIndexSet)
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
func (it *NodeOperatorsRegistryKeysOpIndexSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryKeysOpIndexSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryKeysOpIndexSet represents a KeysOpIndexSet event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryKeysOpIndexSet struct {
	KeysOpIndex *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterKeysOpIndexSet is a free log retrieval operation binding the contract event 0xfb992daec9d46d64898e3a9336d02811349df6cbea8b95d4deb2fa6c7b454f0d.
//
// Solidity: event KeysOpIndexSet(uint256 keysOpIndex)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterKeysOpIndexSet(opts *bind.FilterOpts) (*NodeOperatorsRegistryKeysOpIndexSetIterator, error) {

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "KeysOpIndexSet")
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryKeysOpIndexSetIterator{contract: _NodeOperatorsRegistry.contract, event: "KeysOpIndexSet", logs: logs, sub: sub}, nil
}

// WatchKeysOpIndexSet is a free log subscription operation binding the contract event 0xfb992daec9d46d64898e3a9336d02811349df6cbea8b95d4deb2fa6c7b454f0d.
//
// Solidity: event KeysOpIndexSet(uint256 keysOpIndex)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchKeysOpIndexSet(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryKeysOpIndexSet) (event.Subscription, error) {

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "KeysOpIndexSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryKeysOpIndexSet)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "KeysOpIndexSet", log); err != nil {
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

// ParseKeysOpIndexSet is a log parse operation binding the contract event 0xfb992daec9d46d64898e3a9336d02811349df6cbea8b95d4deb2fa6c7b454f0d.
//
// Solidity: event KeysOpIndexSet(uint256 keysOpIndex)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseKeysOpIndexSet(log types.Log) (*NodeOperatorsRegistryKeysOpIndexSet, error) {
	event := new(NodeOperatorsRegistryKeysOpIndexSet)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "KeysOpIndexSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryLocatorContractSetIterator is returned from FilterLocatorContractSet and is used to iterate over the raw logs and unpacked data for LocatorContractSet events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryLocatorContractSetIterator struct {
	Event *NodeOperatorsRegistryLocatorContractSet // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryLocatorContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryLocatorContractSet)
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
		it.Event = new(NodeOperatorsRegistryLocatorContractSet)
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
func (it *NodeOperatorsRegistryLocatorContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryLocatorContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryLocatorContractSet represents a LocatorContractSet event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryLocatorContractSet struct {
	LocatorAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLocatorContractSet is a free log retrieval operation binding the contract event 0xa44aa4b7320163340e971b1f22f153bbb8a0151d783bd58377018ea5bc96d0c9.
//
// Solidity: event LocatorContractSet(address locatorAddress)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterLocatorContractSet(opts *bind.FilterOpts) (*NodeOperatorsRegistryLocatorContractSetIterator, error) {

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "LocatorContractSet")
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryLocatorContractSetIterator{contract: _NodeOperatorsRegistry.contract, event: "LocatorContractSet", logs: logs, sub: sub}, nil
}

// WatchLocatorContractSet is a free log subscription operation binding the contract event 0xa44aa4b7320163340e971b1f22f153bbb8a0151d783bd58377018ea5bc96d0c9.
//
// Solidity: event LocatorContractSet(address locatorAddress)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchLocatorContractSet(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryLocatorContractSet) (event.Subscription, error) {

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "LocatorContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryLocatorContractSet)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "LocatorContractSet", log); err != nil {
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

// ParseLocatorContractSet is a log parse operation binding the contract event 0xa44aa4b7320163340e971b1f22f153bbb8a0151d783bd58377018ea5bc96d0c9.
//
// Solidity: event LocatorContractSet(address locatorAddress)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseLocatorContractSet(log types.Log) (*NodeOperatorsRegistryLocatorContractSet, error) {
	event := new(NodeOperatorsRegistryLocatorContractSet)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "LocatorContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryNodeOperatorActiveSetIterator is returned from FilterNodeOperatorActiveSet and is used to iterate over the raw logs and unpacked data for NodeOperatorActiveSet events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNodeOperatorActiveSetIterator struct {
	Event *NodeOperatorsRegistryNodeOperatorActiveSet // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryNodeOperatorActiveSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryNodeOperatorActiveSet)
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
		it.Event = new(NodeOperatorsRegistryNodeOperatorActiveSet)
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
func (it *NodeOperatorsRegistryNodeOperatorActiveSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryNodeOperatorActiveSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryNodeOperatorActiveSet represents a NodeOperatorActiveSet event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNodeOperatorActiveSet struct {
	NodeOperatorId *big.Int
	Active         bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorActiveSet is a free log retrieval operation binding the contract event 0xecdf08e8a6c4493efb460f6abc7d14532074fa339c3a6410623a1d3ee0fb2cac.
//
// Solidity: event NodeOperatorActiveSet(uint256 indexed nodeOperatorId, bool active)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterNodeOperatorActiveSet(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*NodeOperatorsRegistryNodeOperatorActiveSetIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "NodeOperatorActiveSet", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryNodeOperatorActiveSetIterator{contract: _NodeOperatorsRegistry.contract, event: "NodeOperatorActiveSet", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorActiveSet is a free log subscription operation binding the contract event 0xecdf08e8a6c4493efb460f6abc7d14532074fa339c3a6410623a1d3ee0fb2cac.
//
// Solidity: event NodeOperatorActiveSet(uint256 indexed nodeOperatorId, bool active)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchNodeOperatorActiveSet(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryNodeOperatorActiveSet, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "NodeOperatorActiveSet", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryNodeOperatorActiveSet)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NodeOperatorActiveSet", log); err != nil {
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

// ParseNodeOperatorActiveSet is a log parse operation binding the contract event 0xecdf08e8a6c4493efb460f6abc7d14532074fa339c3a6410623a1d3ee0fb2cac.
//
// Solidity: event NodeOperatorActiveSet(uint256 indexed nodeOperatorId, bool active)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseNodeOperatorActiveSet(log types.Log) (*NodeOperatorsRegistryNodeOperatorActiveSet, error) {
	event := new(NodeOperatorsRegistryNodeOperatorActiveSet)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NodeOperatorActiveSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryNodeOperatorAddedIterator is returned from FilterNodeOperatorAdded and is used to iterate over the raw logs and unpacked data for NodeOperatorAdded events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNodeOperatorAddedIterator struct {
	Event *NodeOperatorsRegistryNodeOperatorAdded // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryNodeOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryNodeOperatorAdded)
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
		it.Event = new(NodeOperatorsRegistryNodeOperatorAdded)
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
func (it *NodeOperatorsRegistryNodeOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryNodeOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryNodeOperatorAdded represents a NodeOperatorAdded event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNodeOperatorAdded struct {
	NodeOperatorId *big.Int
	Name           string
	RewardAddress  common.Address
	StakingLimit   uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorAdded is a free log retrieval operation binding the contract event 0xc52ec0ad7872dae440d886040390c13677df7bf3cca136d8d81e5e5e7dd62ff1.
//
// Solidity: event NodeOperatorAdded(uint256 nodeOperatorId, string name, address rewardAddress, uint64 stakingLimit)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterNodeOperatorAdded(opts *bind.FilterOpts) (*NodeOperatorsRegistryNodeOperatorAddedIterator, error) {

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "NodeOperatorAdded")
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryNodeOperatorAddedIterator{contract: _NodeOperatorsRegistry.contract, event: "NodeOperatorAdded", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorAdded is a free log subscription operation binding the contract event 0xc52ec0ad7872dae440d886040390c13677df7bf3cca136d8d81e5e5e7dd62ff1.
//
// Solidity: event NodeOperatorAdded(uint256 nodeOperatorId, string name, address rewardAddress, uint64 stakingLimit)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchNodeOperatorAdded(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryNodeOperatorAdded) (event.Subscription, error) {

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "NodeOperatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryNodeOperatorAdded)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NodeOperatorAdded", log); err != nil {
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

// ParseNodeOperatorAdded is a log parse operation binding the contract event 0xc52ec0ad7872dae440d886040390c13677df7bf3cca136d8d81e5e5e7dd62ff1.
//
// Solidity: event NodeOperatorAdded(uint256 nodeOperatorId, string name, address rewardAddress, uint64 stakingLimit)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseNodeOperatorAdded(log types.Log) (*NodeOperatorsRegistryNodeOperatorAdded, error) {
	event := new(NodeOperatorsRegistryNodeOperatorAdded)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NodeOperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryNodeOperatorNameSetIterator is returned from FilterNodeOperatorNameSet and is used to iterate over the raw logs and unpacked data for NodeOperatorNameSet events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNodeOperatorNameSetIterator struct {
	Event *NodeOperatorsRegistryNodeOperatorNameSet // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryNodeOperatorNameSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryNodeOperatorNameSet)
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
		it.Event = new(NodeOperatorsRegistryNodeOperatorNameSet)
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
func (it *NodeOperatorsRegistryNodeOperatorNameSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryNodeOperatorNameSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryNodeOperatorNameSet represents a NodeOperatorNameSet event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNodeOperatorNameSet struct {
	NodeOperatorId *big.Int
	Name           string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorNameSet is a free log retrieval operation binding the contract event 0xcb16868f4831cc58a28d413f658752a2958bd1f50e94ed6391716b936c48093b.
//
// Solidity: event NodeOperatorNameSet(uint256 indexed nodeOperatorId, string name)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterNodeOperatorNameSet(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*NodeOperatorsRegistryNodeOperatorNameSetIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "NodeOperatorNameSet", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryNodeOperatorNameSetIterator{contract: _NodeOperatorsRegistry.contract, event: "NodeOperatorNameSet", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorNameSet is a free log subscription operation binding the contract event 0xcb16868f4831cc58a28d413f658752a2958bd1f50e94ed6391716b936c48093b.
//
// Solidity: event NodeOperatorNameSet(uint256 indexed nodeOperatorId, string name)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchNodeOperatorNameSet(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryNodeOperatorNameSet, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "NodeOperatorNameSet", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryNodeOperatorNameSet)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NodeOperatorNameSet", log); err != nil {
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

// ParseNodeOperatorNameSet is a log parse operation binding the contract event 0xcb16868f4831cc58a28d413f658752a2958bd1f50e94ed6391716b936c48093b.
//
// Solidity: event NodeOperatorNameSet(uint256 indexed nodeOperatorId, string name)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseNodeOperatorNameSet(log types.Log) (*NodeOperatorsRegistryNodeOperatorNameSet, error) {
	event := new(NodeOperatorsRegistryNodeOperatorNameSet)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NodeOperatorNameSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryNodeOperatorPenalizedIterator is returned from FilterNodeOperatorPenalized and is used to iterate over the raw logs and unpacked data for NodeOperatorPenalized events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNodeOperatorPenalizedIterator struct {
	Event *NodeOperatorsRegistryNodeOperatorPenalized // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryNodeOperatorPenalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryNodeOperatorPenalized)
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
		it.Event = new(NodeOperatorsRegistryNodeOperatorPenalized)
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
func (it *NodeOperatorsRegistryNodeOperatorPenalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryNodeOperatorPenalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryNodeOperatorPenalized represents a NodeOperatorPenalized event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNodeOperatorPenalized struct {
	RecipientAddress      common.Address
	SharesPenalizedAmount *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorPenalized is a free log retrieval operation binding the contract event 0xe915a473fc2ef8e0231da98380f853b2aeea117a4392c67e753c54186bfbbd12.
//
// Solidity: event NodeOperatorPenalized(address indexed recipientAddress, uint256 sharesPenalizedAmount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterNodeOperatorPenalized(opts *bind.FilterOpts, recipientAddress []common.Address) (*NodeOperatorsRegistryNodeOperatorPenalizedIterator, error) {

	var recipientAddressRule []interface{}
	for _, recipientAddressItem := range recipientAddress {
		recipientAddressRule = append(recipientAddressRule, recipientAddressItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "NodeOperatorPenalized", recipientAddressRule)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryNodeOperatorPenalizedIterator{contract: _NodeOperatorsRegistry.contract, event: "NodeOperatorPenalized", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorPenalized is a free log subscription operation binding the contract event 0xe915a473fc2ef8e0231da98380f853b2aeea117a4392c67e753c54186bfbbd12.
//
// Solidity: event NodeOperatorPenalized(address indexed recipientAddress, uint256 sharesPenalizedAmount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchNodeOperatorPenalized(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryNodeOperatorPenalized, recipientAddress []common.Address) (event.Subscription, error) {

	var recipientAddressRule []interface{}
	for _, recipientAddressItem := range recipientAddress {
		recipientAddressRule = append(recipientAddressRule, recipientAddressItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "NodeOperatorPenalized", recipientAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryNodeOperatorPenalized)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NodeOperatorPenalized", log); err != nil {
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

// ParseNodeOperatorPenalized is a log parse operation binding the contract event 0xe915a473fc2ef8e0231da98380f853b2aeea117a4392c67e753c54186bfbbd12.
//
// Solidity: event NodeOperatorPenalized(address indexed recipientAddress, uint256 sharesPenalizedAmount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseNodeOperatorPenalized(log types.Log) (*NodeOperatorsRegistryNodeOperatorPenalized, error) {
	event := new(NodeOperatorsRegistryNodeOperatorPenalized)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NodeOperatorPenalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryNodeOperatorRewardAddressSetIterator is returned from FilterNodeOperatorRewardAddressSet and is used to iterate over the raw logs and unpacked data for NodeOperatorRewardAddressSet events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNodeOperatorRewardAddressSetIterator struct {
	Event *NodeOperatorsRegistryNodeOperatorRewardAddressSet // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryNodeOperatorRewardAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryNodeOperatorRewardAddressSet)
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
		it.Event = new(NodeOperatorsRegistryNodeOperatorRewardAddressSet)
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
func (it *NodeOperatorsRegistryNodeOperatorRewardAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryNodeOperatorRewardAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryNodeOperatorRewardAddressSet represents a NodeOperatorRewardAddressSet event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNodeOperatorRewardAddressSet struct {
	NodeOperatorId *big.Int
	RewardAddress  common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorRewardAddressSet is a free log retrieval operation binding the contract event 0x9a52205165d510fc1e428886d52108725dc01ed544da1702dc7bd3fdb3f243b2.
//
// Solidity: event NodeOperatorRewardAddressSet(uint256 indexed nodeOperatorId, address rewardAddress)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterNodeOperatorRewardAddressSet(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*NodeOperatorsRegistryNodeOperatorRewardAddressSetIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "NodeOperatorRewardAddressSet", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryNodeOperatorRewardAddressSetIterator{contract: _NodeOperatorsRegistry.contract, event: "NodeOperatorRewardAddressSet", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorRewardAddressSet is a free log subscription operation binding the contract event 0x9a52205165d510fc1e428886d52108725dc01ed544da1702dc7bd3fdb3f243b2.
//
// Solidity: event NodeOperatorRewardAddressSet(uint256 indexed nodeOperatorId, address rewardAddress)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchNodeOperatorRewardAddressSet(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryNodeOperatorRewardAddressSet, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "NodeOperatorRewardAddressSet", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryNodeOperatorRewardAddressSet)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NodeOperatorRewardAddressSet", log); err != nil {
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

// ParseNodeOperatorRewardAddressSet is a log parse operation binding the contract event 0x9a52205165d510fc1e428886d52108725dc01ed544da1702dc7bd3fdb3f243b2.
//
// Solidity: event NodeOperatorRewardAddressSet(uint256 indexed nodeOperatorId, address rewardAddress)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseNodeOperatorRewardAddressSet(log types.Log) (*NodeOperatorsRegistryNodeOperatorRewardAddressSet, error) {
	event := new(NodeOperatorsRegistryNodeOperatorRewardAddressSet)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NodeOperatorRewardAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryNodeOperatorTotalKeysTrimmedIterator is returned from FilterNodeOperatorTotalKeysTrimmed and is used to iterate over the raw logs and unpacked data for NodeOperatorTotalKeysTrimmed events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNodeOperatorTotalKeysTrimmedIterator struct {
	Event *NodeOperatorsRegistryNodeOperatorTotalKeysTrimmed // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryNodeOperatorTotalKeysTrimmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryNodeOperatorTotalKeysTrimmed)
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
		it.Event = new(NodeOperatorsRegistryNodeOperatorTotalKeysTrimmed)
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
func (it *NodeOperatorsRegistryNodeOperatorTotalKeysTrimmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryNodeOperatorTotalKeysTrimmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryNodeOperatorTotalKeysTrimmed represents a NodeOperatorTotalKeysTrimmed event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNodeOperatorTotalKeysTrimmed struct {
	NodeOperatorId   *big.Int
	TotalKeysTrimmed uint64
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorTotalKeysTrimmed is a free log retrieval operation binding the contract event 0x9824694569ba758f8872bb150515caaf8f1e2cc27e6805679c4ac8c3b9b83d87.
//
// Solidity: event NodeOperatorTotalKeysTrimmed(uint256 indexed nodeOperatorId, uint64 totalKeysTrimmed)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterNodeOperatorTotalKeysTrimmed(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*NodeOperatorsRegistryNodeOperatorTotalKeysTrimmedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "NodeOperatorTotalKeysTrimmed", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryNodeOperatorTotalKeysTrimmedIterator{contract: _NodeOperatorsRegistry.contract, event: "NodeOperatorTotalKeysTrimmed", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorTotalKeysTrimmed is a free log subscription operation binding the contract event 0x9824694569ba758f8872bb150515caaf8f1e2cc27e6805679c4ac8c3b9b83d87.
//
// Solidity: event NodeOperatorTotalKeysTrimmed(uint256 indexed nodeOperatorId, uint64 totalKeysTrimmed)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchNodeOperatorTotalKeysTrimmed(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryNodeOperatorTotalKeysTrimmed, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "NodeOperatorTotalKeysTrimmed", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryNodeOperatorTotalKeysTrimmed)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NodeOperatorTotalKeysTrimmed", log); err != nil {
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

// ParseNodeOperatorTotalKeysTrimmed is a log parse operation binding the contract event 0x9824694569ba758f8872bb150515caaf8f1e2cc27e6805679c4ac8c3b9b83d87.
//
// Solidity: event NodeOperatorTotalKeysTrimmed(uint256 indexed nodeOperatorId, uint64 totalKeysTrimmed)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseNodeOperatorTotalKeysTrimmed(log types.Log) (*NodeOperatorsRegistryNodeOperatorTotalKeysTrimmed, error) {
	event := new(NodeOperatorsRegistryNodeOperatorTotalKeysTrimmed)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NodeOperatorTotalKeysTrimmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryNonceChangedIterator is returned from FilterNonceChanged and is used to iterate over the raw logs and unpacked data for NonceChanged events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNonceChangedIterator struct {
	Event *NodeOperatorsRegistryNonceChanged // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryNonceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryNonceChanged)
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
		it.Event = new(NodeOperatorsRegistryNonceChanged)
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
func (it *NodeOperatorsRegistryNonceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryNonceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryNonceChanged represents a NonceChanged event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNonceChanged struct {
	Nonce *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNonceChanged is a free log retrieval operation binding the contract event 0x7220970e1f1f12864ecccd8942690a837c7a8dd45d158cb891eb45a8a69134aa.
//
// Solidity: event NonceChanged(uint256 nonce)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterNonceChanged(opts *bind.FilterOpts) (*NodeOperatorsRegistryNonceChangedIterator, error) {

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "NonceChanged")
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryNonceChangedIterator{contract: _NodeOperatorsRegistry.contract, event: "NonceChanged", logs: logs, sub: sub}, nil
}

// WatchNonceChanged is a free log subscription operation binding the contract event 0x7220970e1f1f12864ecccd8942690a837c7a8dd45d158cb891eb45a8a69134aa.
//
// Solidity: event NonceChanged(uint256 nonce)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchNonceChanged(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryNonceChanged) (event.Subscription, error) {

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "NonceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryNonceChanged)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NonceChanged", log); err != nil {
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

// ParseNonceChanged is a log parse operation binding the contract event 0x7220970e1f1f12864ecccd8942690a837c7a8dd45d158cb891eb45a8a69134aa.
//
// Solidity: event NonceChanged(uint256 nonce)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseNonceChanged(log types.Log) (*NodeOperatorsRegistryNonceChanged, error) {
	event := new(NodeOperatorsRegistryNonceChanged)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NonceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryRecoverToVaultIterator is returned from FilterRecoverToVault and is used to iterate over the raw logs and unpacked data for RecoverToVault events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryRecoverToVaultIterator struct {
	Event *NodeOperatorsRegistryRecoverToVault // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryRecoverToVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryRecoverToVault)
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
		it.Event = new(NodeOperatorsRegistryRecoverToVault)
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
func (it *NodeOperatorsRegistryRecoverToVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryRecoverToVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryRecoverToVault represents a RecoverToVault event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryRecoverToVault struct {
	Vault  common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecoverToVault is a free log retrieval operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterRecoverToVault(opts *bind.FilterOpts, vault []common.Address, token []common.Address) (*NodeOperatorsRegistryRecoverToVaultIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryRecoverToVaultIterator{contract: _NodeOperatorsRegistry.contract, event: "RecoverToVault", logs: logs, sub: sub}, nil
}

// WatchRecoverToVault is a free log subscription operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchRecoverToVault(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryRecoverToVault, vault []common.Address, token []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryRecoverToVault)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
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

// ParseRecoverToVault is a log parse operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseRecoverToVault(log types.Log) (*NodeOperatorsRegistryRecoverToVault, error) {
	event := new(NodeOperatorsRegistryRecoverToVault)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryRewardsDistributedIterator is returned from FilterRewardsDistributed and is used to iterate over the raw logs and unpacked data for RewardsDistributed events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryRewardsDistributedIterator struct {
	Event *NodeOperatorsRegistryRewardsDistributed // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryRewardsDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryRewardsDistributed)
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
		it.Event = new(NodeOperatorsRegistryRewardsDistributed)
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
func (it *NodeOperatorsRegistryRewardsDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryRewardsDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryRewardsDistributed represents a RewardsDistributed event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryRewardsDistributed struct {
	RewardAddress common.Address
	SharesAmount  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRewardsDistributed is a free log retrieval operation binding the contract event 0xdf29796aad820e4bb192f3a8d631b76519bcd2cbe77cc85af20e9df53cece086.
//
// Solidity: event RewardsDistributed(address indexed rewardAddress, uint256 sharesAmount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterRewardsDistributed(opts *bind.FilterOpts, rewardAddress []common.Address) (*NodeOperatorsRegistryRewardsDistributedIterator, error) {

	var rewardAddressRule []interface{}
	for _, rewardAddressItem := range rewardAddress {
		rewardAddressRule = append(rewardAddressRule, rewardAddressItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "RewardsDistributed", rewardAddressRule)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryRewardsDistributedIterator{contract: _NodeOperatorsRegistry.contract, event: "RewardsDistributed", logs: logs, sub: sub}, nil
}

// WatchRewardsDistributed is a free log subscription operation binding the contract event 0xdf29796aad820e4bb192f3a8d631b76519bcd2cbe77cc85af20e9df53cece086.
//
// Solidity: event RewardsDistributed(address indexed rewardAddress, uint256 sharesAmount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchRewardsDistributed(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryRewardsDistributed, rewardAddress []common.Address) (event.Subscription, error) {

	var rewardAddressRule []interface{}
	for _, rewardAddressItem := range rewardAddress {
		rewardAddressRule = append(rewardAddressRule, rewardAddressItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "RewardsDistributed", rewardAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryRewardsDistributed)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "RewardsDistributed", log); err != nil {
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

// ParseRewardsDistributed is a log parse operation binding the contract event 0xdf29796aad820e4bb192f3a8d631b76519bcd2cbe77cc85af20e9df53cece086.
//
// Solidity: event RewardsDistributed(address indexed rewardAddress, uint256 sharesAmount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseRewardsDistributed(log types.Log) (*NodeOperatorsRegistryRewardsDistributed, error) {
	event := new(NodeOperatorsRegistryRewardsDistributed)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "RewardsDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryScriptResultIterator is returned from FilterScriptResult and is used to iterate over the raw logs and unpacked data for ScriptResult events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryScriptResultIterator struct {
	Event *NodeOperatorsRegistryScriptResult // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryScriptResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryScriptResult)
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
		it.Event = new(NodeOperatorsRegistryScriptResult)
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
func (it *NodeOperatorsRegistryScriptResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryScriptResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryScriptResult represents a ScriptResult event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryScriptResult struct {
	Executor   common.Address
	Script     []byte
	Input      []byte
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterScriptResult is a free log retrieval operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterScriptResult(opts *bind.FilterOpts, executor []common.Address) (*NodeOperatorsRegistryScriptResultIterator, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryScriptResultIterator{contract: _NodeOperatorsRegistry.contract, event: "ScriptResult", logs: logs, sub: sub}, nil
}

// WatchScriptResult is a free log subscription operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchScriptResult(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryScriptResult, executor []common.Address) (event.Subscription, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryScriptResult)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "ScriptResult", log); err != nil {
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

// ParseScriptResult is a log parse operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseScriptResult(log types.Log) (*NodeOperatorsRegistryScriptResult, error) {
	event := new(NodeOperatorsRegistryScriptResult)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "ScriptResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryStakingModuleTypeSetIterator is returned from FilterStakingModuleTypeSet and is used to iterate over the raw logs and unpacked data for StakingModuleTypeSet events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryStakingModuleTypeSetIterator struct {
	Event *NodeOperatorsRegistryStakingModuleTypeSet // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryStakingModuleTypeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryStakingModuleTypeSet)
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
		it.Event = new(NodeOperatorsRegistryStakingModuleTypeSet)
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
func (it *NodeOperatorsRegistryStakingModuleTypeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryStakingModuleTypeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryStakingModuleTypeSet represents a StakingModuleTypeSet event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryStakingModuleTypeSet struct {
	ModuleType [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStakingModuleTypeSet is a free log retrieval operation binding the contract event 0xdb042010b15d1321c99552200b350bba0a95dfa3d0b43869983ce74b44d644ee.
//
// Solidity: event StakingModuleTypeSet(bytes32 moduleType)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterStakingModuleTypeSet(opts *bind.FilterOpts) (*NodeOperatorsRegistryStakingModuleTypeSetIterator, error) {

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "StakingModuleTypeSet")
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryStakingModuleTypeSetIterator{contract: _NodeOperatorsRegistry.contract, event: "StakingModuleTypeSet", logs: logs, sub: sub}, nil
}

// WatchStakingModuleTypeSet is a free log subscription operation binding the contract event 0xdb042010b15d1321c99552200b350bba0a95dfa3d0b43869983ce74b44d644ee.
//
// Solidity: event StakingModuleTypeSet(bytes32 moduleType)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchStakingModuleTypeSet(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryStakingModuleTypeSet) (event.Subscription, error) {

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "StakingModuleTypeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryStakingModuleTypeSet)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "StakingModuleTypeSet", log); err != nil {
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

// ParseStakingModuleTypeSet is a log parse operation binding the contract event 0xdb042010b15d1321c99552200b350bba0a95dfa3d0b43869983ce74b44d644ee.
//
// Solidity: event StakingModuleTypeSet(bytes32 moduleType)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseStakingModuleTypeSet(log types.Log) (*NodeOperatorsRegistryStakingModuleTypeSet, error) {
	event := new(NodeOperatorsRegistryStakingModuleTypeSet)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "StakingModuleTypeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryStuckPenaltyDelayChangedIterator is returned from FilterStuckPenaltyDelayChanged and is used to iterate over the raw logs and unpacked data for StuckPenaltyDelayChanged events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryStuckPenaltyDelayChangedIterator struct {
	Event *NodeOperatorsRegistryStuckPenaltyDelayChanged // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryStuckPenaltyDelayChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryStuckPenaltyDelayChanged)
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
		it.Event = new(NodeOperatorsRegistryStuckPenaltyDelayChanged)
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
func (it *NodeOperatorsRegistryStuckPenaltyDelayChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryStuckPenaltyDelayChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryStuckPenaltyDelayChanged represents a StuckPenaltyDelayChanged event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryStuckPenaltyDelayChanged struct {
	StuckPenaltyDelay *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterStuckPenaltyDelayChanged is a free log retrieval operation binding the contract event 0x4cccd9748bff0341d9852cc61d82652a3003dcebea088f05388c0be1f26b4c8a.
//
// Solidity: event StuckPenaltyDelayChanged(uint256 stuckPenaltyDelay)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterStuckPenaltyDelayChanged(opts *bind.FilterOpts) (*NodeOperatorsRegistryStuckPenaltyDelayChangedIterator, error) {

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "StuckPenaltyDelayChanged")
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryStuckPenaltyDelayChangedIterator{contract: _NodeOperatorsRegistry.contract, event: "StuckPenaltyDelayChanged", logs: logs, sub: sub}, nil
}

// WatchStuckPenaltyDelayChanged is a free log subscription operation binding the contract event 0x4cccd9748bff0341d9852cc61d82652a3003dcebea088f05388c0be1f26b4c8a.
//
// Solidity: event StuckPenaltyDelayChanged(uint256 stuckPenaltyDelay)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchStuckPenaltyDelayChanged(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryStuckPenaltyDelayChanged) (event.Subscription, error) {

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "StuckPenaltyDelayChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryStuckPenaltyDelayChanged)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "StuckPenaltyDelayChanged", log); err != nil {
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

// ParseStuckPenaltyDelayChanged is a log parse operation binding the contract event 0x4cccd9748bff0341d9852cc61d82652a3003dcebea088f05388c0be1f26b4c8a.
//
// Solidity: event StuckPenaltyDelayChanged(uint256 stuckPenaltyDelay)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseStuckPenaltyDelayChanged(log types.Log) (*NodeOperatorsRegistryStuckPenaltyDelayChanged, error) {
	event := new(NodeOperatorsRegistryStuckPenaltyDelayChanged)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "StuckPenaltyDelayChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryStuckPenaltyStateChangedIterator is returned from FilterStuckPenaltyStateChanged and is used to iterate over the raw logs and unpacked data for StuckPenaltyStateChanged events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryStuckPenaltyStateChangedIterator struct {
	Event *NodeOperatorsRegistryStuckPenaltyStateChanged // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryStuckPenaltyStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryStuckPenaltyStateChanged)
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
		it.Event = new(NodeOperatorsRegistryStuckPenaltyStateChanged)
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
func (it *NodeOperatorsRegistryStuckPenaltyStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryStuckPenaltyStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryStuckPenaltyStateChanged represents a StuckPenaltyStateChanged event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryStuckPenaltyStateChanged struct {
	NodeOperatorId           *big.Int
	StuckValidatorsCount     *big.Int
	RefundedValidatorsCount  *big.Int
	StuckPenaltyEndTimestamp *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterStuckPenaltyStateChanged is a free log retrieval operation binding the contract event 0x0ee42dd52dd2b8feb0fc9cc054a08162a23e022c177319db981cf339e5b8ffdb.
//
// Solidity: event StuckPenaltyStateChanged(uint256 indexed nodeOperatorId, uint256 stuckValidatorsCount, uint256 refundedValidatorsCount, uint256 stuckPenaltyEndTimestamp)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterStuckPenaltyStateChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*NodeOperatorsRegistryStuckPenaltyStateChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "StuckPenaltyStateChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryStuckPenaltyStateChangedIterator{contract: _NodeOperatorsRegistry.contract, event: "StuckPenaltyStateChanged", logs: logs, sub: sub}, nil
}

// WatchStuckPenaltyStateChanged is a free log subscription operation binding the contract event 0x0ee42dd52dd2b8feb0fc9cc054a08162a23e022c177319db981cf339e5b8ffdb.
//
// Solidity: event StuckPenaltyStateChanged(uint256 indexed nodeOperatorId, uint256 stuckValidatorsCount, uint256 refundedValidatorsCount, uint256 stuckPenaltyEndTimestamp)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchStuckPenaltyStateChanged(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryStuckPenaltyStateChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "StuckPenaltyStateChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryStuckPenaltyStateChanged)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "StuckPenaltyStateChanged", log); err != nil {
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

// ParseStuckPenaltyStateChanged is a log parse operation binding the contract event 0x0ee42dd52dd2b8feb0fc9cc054a08162a23e022c177319db981cf339e5b8ffdb.
//
// Solidity: event StuckPenaltyStateChanged(uint256 indexed nodeOperatorId, uint256 stuckValidatorsCount, uint256 refundedValidatorsCount, uint256 stuckPenaltyEndTimestamp)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseStuckPenaltyStateChanged(log types.Log) (*NodeOperatorsRegistryStuckPenaltyStateChanged, error) {
	event := new(NodeOperatorsRegistryStuckPenaltyStateChanged)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "StuckPenaltyStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryTargetValidatorsCountChangedIterator is returned from FilterTargetValidatorsCountChanged and is used to iterate over the raw logs and unpacked data for TargetValidatorsCountChanged events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryTargetValidatorsCountChangedIterator struct {
	Event *NodeOperatorsRegistryTargetValidatorsCountChanged // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryTargetValidatorsCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryTargetValidatorsCountChanged)
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
		it.Event = new(NodeOperatorsRegistryTargetValidatorsCountChanged)
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
func (it *NodeOperatorsRegistryTargetValidatorsCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryTargetValidatorsCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryTargetValidatorsCountChanged represents a TargetValidatorsCountChanged event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryTargetValidatorsCountChanged struct {
	NodeOperatorId        *big.Int
	TargetValidatorsCount *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterTargetValidatorsCountChanged is a free log retrieval operation binding the contract event 0xd50ea115db6f0b433ef9cc4b71110dbd9202364a00488be90718990be5bf16a6.
//
// Solidity: event TargetValidatorsCountChanged(uint256 indexed nodeOperatorId, uint256 targetValidatorsCount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterTargetValidatorsCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*NodeOperatorsRegistryTargetValidatorsCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "TargetValidatorsCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryTargetValidatorsCountChangedIterator{contract: _NodeOperatorsRegistry.contract, event: "TargetValidatorsCountChanged", logs: logs, sub: sub}, nil
}

// WatchTargetValidatorsCountChanged is a free log subscription operation binding the contract event 0xd50ea115db6f0b433ef9cc4b71110dbd9202364a00488be90718990be5bf16a6.
//
// Solidity: event TargetValidatorsCountChanged(uint256 indexed nodeOperatorId, uint256 targetValidatorsCount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchTargetValidatorsCountChanged(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryTargetValidatorsCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "TargetValidatorsCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryTargetValidatorsCountChanged)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "TargetValidatorsCountChanged", log); err != nil {
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

// ParseTargetValidatorsCountChanged is a log parse operation binding the contract event 0xd50ea115db6f0b433ef9cc4b71110dbd9202364a00488be90718990be5bf16a6.
//
// Solidity: event TargetValidatorsCountChanged(uint256 indexed nodeOperatorId, uint256 targetValidatorsCount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseTargetValidatorsCountChanged(log types.Log) (*NodeOperatorsRegistryTargetValidatorsCountChanged, error) {
	event := new(NodeOperatorsRegistryTargetValidatorsCountChanged)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "TargetValidatorsCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryTotalSigningKeysCountChangedIterator is returned from FilterTotalSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for TotalSigningKeysCountChanged events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryTotalSigningKeysCountChangedIterator struct {
	Event *NodeOperatorsRegistryTotalSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryTotalSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryTotalSigningKeysCountChanged)
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
		it.Event = new(NodeOperatorsRegistryTotalSigningKeysCountChanged)
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
func (it *NodeOperatorsRegistryTotalSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryTotalSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryTotalSigningKeysCountChanged represents a TotalSigningKeysCountChanged event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryTotalSigningKeysCountChanged struct {
	NodeOperatorId       *big.Int
	TotalValidatorsCount *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTotalSigningKeysCountChanged is a free log retrieval operation binding the contract event 0xdd01838a366ae4dc9a86e1922512c0716abebc9a440baae0e22d2dec578223f0.
//
// Solidity: event TotalSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 totalValidatorsCount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterTotalSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*NodeOperatorsRegistryTotalSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "TotalSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryTotalSigningKeysCountChangedIterator{contract: _NodeOperatorsRegistry.contract, event: "TotalSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchTotalSigningKeysCountChanged is a free log subscription operation binding the contract event 0xdd01838a366ae4dc9a86e1922512c0716abebc9a440baae0e22d2dec578223f0.
//
// Solidity: event TotalSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 totalValidatorsCount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchTotalSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryTotalSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "TotalSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryTotalSigningKeysCountChanged)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "TotalSigningKeysCountChanged", log); err != nil {
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

// ParseTotalSigningKeysCountChanged is a log parse operation binding the contract event 0xdd01838a366ae4dc9a86e1922512c0716abebc9a440baae0e22d2dec578223f0.
//
// Solidity: event TotalSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 totalValidatorsCount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseTotalSigningKeysCountChanged(log types.Log) (*NodeOperatorsRegistryTotalSigningKeysCountChanged, error) {
	event := new(NodeOperatorsRegistryTotalSigningKeysCountChanged)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "TotalSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryVettedSigningKeysCountChangedIterator is returned from FilterVettedSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for VettedSigningKeysCountChanged events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryVettedSigningKeysCountChangedIterator struct {
	Event *NodeOperatorsRegistryVettedSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryVettedSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryVettedSigningKeysCountChanged)
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
		it.Event = new(NodeOperatorsRegistryVettedSigningKeysCountChanged)
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
func (it *NodeOperatorsRegistryVettedSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryVettedSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryVettedSigningKeysCountChanged represents a VettedSigningKeysCountChanged event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryVettedSigningKeysCountChanged struct {
	NodeOperatorId          *big.Int
	ApprovedValidatorsCount *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterVettedSigningKeysCountChanged is a free log retrieval operation binding the contract event 0x947f955eec7e1f626bee3afd2aa47b5de04ddcdd3fe78dc8838213015ef58dfd.
//
// Solidity: event VettedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 approvedValidatorsCount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterVettedSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*NodeOperatorsRegistryVettedSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "VettedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryVettedSigningKeysCountChangedIterator{contract: _NodeOperatorsRegistry.contract, event: "VettedSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchVettedSigningKeysCountChanged is a free log subscription operation binding the contract event 0x947f955eec7e1f626bee3afd2aa47b5de04ddcdd3fe78dc8838213015ef58dfd.
//
// Solidity: event VettedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 approvedValidatorsCount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchVettedSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryVettedSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "VettedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryVettedSigningKeysCountChanged)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "VettedSigningKeysCountChanged", log); err != nil {
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

// ParseVettedSigningKeysCountChanged is a log parse operation binding the contract event 0x947f955eec7e1f626bee3afd2aa47b5de04ddcdd3fe78dc8838213015ef58dfd.
//
// Solidity: event VettedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 approvedValidatorsCount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseVettedSigningKeysCountChanged(log types.Log) (*NodeOperatorsRegistryVettedSigningKeysCountChanged, error) {
	event := new(NodeOperatorsRegistryVettedSigningKeysCountChanged)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "VettedSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

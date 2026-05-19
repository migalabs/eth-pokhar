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

// IAccountingPermitInput is an auto generated low-level Go binding around an user-defined struct.
type IAccountingPermitInput struct {
	Value    *big.Int
	Deadline *big.Int
	V        uint8
	R        [32]byte
	S        [32]byte
}

// NodeOperator is an auto generated low-level Go binding around an user-defined struct.
type NodeOperator struct {
	TotalAddedKeys             uint32
	TotalWithdrawnKeys         uint32
	TotalDepositedKeys         uint32
	TotalVettedKeys            uint32
	StuckValidatorsCount       uint32
	DepositableValidatorsCount uint32
	TargetLimit                uint32
	TargetLimitMode            uint8
	TotalExitedKeys            uint32
	EnqueuedCount              uint32
	ManagerAddress             common.Address
	ProposedManagerAddress     common.Address
	RewardAddress              common.Address
	ProposedRewardAddress      common.Address
	ExtendedManagerPermissions bool
	UsedPriorityQueue          bool
}

// NodeOperatorManagementProperties is an auto generated low-level Go binding around an user-defined struct.
type NodeOperatorManagementProperties struct {
	ManagerAddress             common.Address
	RewardAddress              common.Address
	ExtendedManagerPermissions bool
}

// WithdrawnValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type WithdrawnValidatorInfo struct {
	NodeOperatorId  *big.Int
	KeyIndex        *big.Int
	ExitBalance     *big.Int
	SlashingPenalty *big.Int
	IsSlashed       bool
}

// CuratedModuleMetaData contains all meta data concerning the CuratedModule contract.
var CuratedModuleMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"moduleType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lidoLocator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"parametersRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"accounting\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"exitPenalties\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metaRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ACCOUNTING\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAccounting\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CREATE_NODE_OPERATOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EXIT_PENALTIES\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIExitPenalties\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FEE_DISTRIBUTOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LIDO_LOCATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractILidoLocator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"META_REGISTRY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIMetaRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_ADDRESSES_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PARAMETERS_REGISTRY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIParametersRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSE_INFINITELY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSE_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RECOVERER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REPORT_GENERAL_DELAYED_PENALTY_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REPORT_REGULAR_WITHDRAWN_VALIDATORS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REPORT_SLASHED_WITHDRAWN_VALIDATORS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RESUME_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SETTLE_GENERAL_DELAYED_PENALTY_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_ROUTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STETH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStETH\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERIFIER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addValidatorKeysETH\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"keysCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publicKeys\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signatures\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"addValidatorKeysStETH\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"keysCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publicKeys\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signatures\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"permit\",\"type\":\"tuple\",\"internalType\":\"structIAccounting.PermitInput\",\"components\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addValidatorKeysWstETH\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"keysCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publicKeys\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signatures\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"permit\",\"type\":\"tuple\",\"internalType\":\"structIAccounting.PermitInput\",\"components\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocateDeposits\",\"inputs\":[{\"name\":\"maxDepositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pubkeys\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"keyIndices\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"operatorIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"topUpLimits\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"allocations\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchDepositInfoUpdate\",\"inputs\":[{\"name\":\"maxCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"operatorsLeft\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelGeneralDelayedPenalty\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeNodeOperatorAddresses\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newRewardAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeNodeOperatorRewardAddress\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"compensateGeneralDelayedPenalty\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"confirmNodeOperatorManagerAddressChange\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"confirmNodeOperatorRewardAddressChange\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createNodeOperator\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"managementProperties\",\"type\":\"tuple\",\"internalType\":\"structNodeOperatorManagementProperties\",\"components\":[{\"name\":\"managerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rewardAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extendedManagerPermissions\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decreaseVettedSigningKeysCount\",\"inputs\":[{\"name\":\"nodeOperatorIds\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"vettedSigningKeysCounts\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"exitDeadlineThreshold\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActiveNodeOperatorsCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDepositAllocationTargets\",\"inputs\":[],\"outputs\":[{\"name\":\"currentValidators\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"targetValidators\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDepositsAllocation\",\"inputs\":[{\"name\":\"maxDepositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"allocated\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"operatorIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"allocations\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInitializedVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getKeyAllocatedBalances\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"keysCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"balances\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getKeyConfirmedBalances\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"keysCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"balances\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOperator\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structNodeOperator\",\"components\":[{\"name\":\"totalAddedKeys\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"totalWithdrawnKeys\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"totalDepositedKeys\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"totalVettedKeys\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stuckValidatorsCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"depositableValidatorsCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"targetLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"targetLimitMode\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"totalExitedKeys\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enqueuedCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"managerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proposedManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rewardAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proposedRewardAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extendedManagerPermissions\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"usedPriorityQueue\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOperatorBalance\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOperatorDepositInfoToUpdateCount\",\"inputs\":[],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOperatorIds\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"nodeOperatorIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOperatorIsActive\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOperatorManagementProperties\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structNodeOperatorManagementProperties\",\"components\":[{\"name\":\"managerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rewardAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extendedManagerPermissions\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOperatorNonWithdrawnKeys\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOperatorOwner\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOperatorSummary\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"targetLimitMode\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetValidatorsCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stuckValidatorsCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"refundedValidatorsCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stuckPenaltyEndTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalExitedValidators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDepositedValidators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"depositableValidatorsCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOperatorWeightAndExternalStake\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"externalStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOperatorsCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorWeights\",\"inputs\":[{\"name\":\"operatorIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"operatorWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getResumeSinceTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMember\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMemberCount\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMembers\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSigningKeys\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"keysCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"keys\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSigningKeysWithSignatures\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"keysCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"keys\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signatures\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakingModuleSummary\",\"inputs\":[],\"outputs\":[{\"name\":\"totalExitedValidators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDepositedValidators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"depositableValidatorsCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTopUpAllocationTargets\",\"inputs\":[],\"outputs\":[{\"name\":\"currentAllocations\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"targetAllocations\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalModuleStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidatorExitDelayPenaltyApplicable\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"eligibleToExitInSec\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidatorSlashed\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"keyIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidatorWithdrawn\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"keyIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"notifyNodeOperatorWeightChange\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"oldWeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newWeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"obtainDepositData\",\"inputs\":[{\"name\":\"depositsCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"publicKeys\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signatures\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onExitedAndStuckValidatorsCountsUpdated\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onRewardsMinted\",\"inputs\":[{\"name\":\"totalShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onValidatorExitTriggered\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"elWithdrawalRequestFeePaid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"exitType\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onWithdrawalCredentialsChanged\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseFor\",\"inputs\":[{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proposeNodeOperatorManagerAddressChange\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proposeNodeOperatorRewardAddressChange\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recoverERC1155\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recoverERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recoverERC721\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recoverEther\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeKeys\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"keysCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reportGeneralDelayedPenalty\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"penaltyType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"details\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reportRegularWithdrawnValidators\",\"inputs\":[{\"name\":\"validatorInfos\",\"type\":\"tuple[]\",\"internalType\":\"structWithdrawnValidatorInfo[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"keyIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"exitBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slashingPenalty\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isSlashed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reportSlashedWithdrawnValidators\",\"inputs\":[{\"name\":\"validatorInfos\",\"type\":\"tuple[]\",\"internalType\":\"structWithdrawnValidatorInfo[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"keyIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"exitBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slashingPenalty\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isSlashed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reportValidatorBalance\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"keyIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentBalanceWei\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reportValidatorExitDelay\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"eligibleToExitInSec\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reportValidatorSlashing\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"keyIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestFullDepositInfoUpdate\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resetNodeOperatorManagerAddress\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resume\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"settleGeneralDelayedPenalty\",\"inputs\":[{\"name\":\"nodeOperatorIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"maxAmounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unsafeUpdateValidatorsCount\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"exitedValidatorsCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateDepositInfo\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateDepositableValidatorsCount\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateExitedValidatorsCount\",\"inputs\":[{\"name\":\"nodeOperatorIds\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"exitedValidatorsCounts\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTargetValidatorsLimits\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetLimitMode\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DepositableSigningKeysCountChanged\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"depositableKeysCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositedSigningKeysCountChanged\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"depositedKeysCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ERC1155Recovered\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ERC20Recovered\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ERC721Recovered\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EtherRecovered\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExitedSigningKeysCountChanged\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"exitedKeysCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FullDepositInfoUpdateRequested\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GeneralDelayedPenaltyCancelled\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GeneralDelayedPenaltyCompensated\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GeneralDelayedPenaltyReported\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"penaltyType\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"additionalFine\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"details\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GeneralDelayedPenaltySettled\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"KeyAllocatedBalanceChanged\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"keyIndex\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newTotal\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"KeyConfirmedBalanceChanged\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"keyIndex\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newBalance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"KeyRemovalChargeApplied\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeOperatorAdded\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"managerAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"rewardAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"extendedManagerPermissions\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeOperatorBalanceUpdated\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"balanceWei\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeOperatorDepositInfoFullyUpdated\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeOperatorManagerAddressChangeProposed\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"oldProposedAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newProposedAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeOperatorManagerAddressChanged\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"oldAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeOperatorRewardAddressChangeProposed\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"oldProposedAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newProposedAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeOperatorRewardAddressChanged\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"oldAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NonceChanged\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"duration\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReferrerSet\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"referrer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Resumed\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SigningKeyAdded\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"pubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SigningKeyRemoved\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"pubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StETHSharesRecovered\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TargetValidatorsCountChanged\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"targetLimitMode\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"targetValidatorsCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TotalSigningKeysCountChanged\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"totalKeysCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TotalWithdrawnValidatorsRebuilt\",\"inputs\":[{\"name\":\"totalWithdrawnValidators\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorSlashingReported\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"keyIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"pubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorWithdrawn\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"keyIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"exitBalance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"slashingPenalty\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"pubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VettedSigningKeysCountChanged\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"vettedKeysCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VettedSigningKeysCountDecreased\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AlreadyProposed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotAddKeys\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositInfoIsNotUpToDate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositableKeysWithUnsupportedWithdrawalCredentials\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedToSendEther\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidManagerAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRewardAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSigningKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVetKeysPointer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWithdrawnValidatorInfo\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"KeysLimitExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MethodCallIsNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NodeOperatorDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToRecover\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughKeys\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PauseUntilMustBeInFuture\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PausedExpected\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PubkeyMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ResumedExpected\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SameAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderIsNotEligible\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderIsNotManagerAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderIsNotMetaRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderIsNotProposedAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderIsNotRewardAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SigningKeysInvalidOffset\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SlashingPenaltyIsNotApplicable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnreportableBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorSlashingAlreadyReported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAccountingAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAdminAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroExitBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroExitPenaltiesAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroLocatorAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroManagerAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroMetaRegistryAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroModuleType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroParametersRegistryAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroPauseDuration\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroPenaltyType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroRewardAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroSenderAddress\",\"inputs\":[]}]",
}

// CuratedModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use CuratedModuleMetaData.ABI instead.
var CuratedModuleABI = CuratedModuleMetaData.ABI

// CuratedModule is an auto generated Go binding around an Ethereum contract.
type CuratedModule struct {
	CuratedModuleCaller     // Read-only binding to the contract
	CuratedModuleTransactor // Write-only binding to the contract
	CuratedModuleFilterer   // Log filterer for contract events
}

// CuratedModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type CuratedModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CuratedModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CuratedModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CuratedModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CuratedModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CuratedModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CuratedModuleSession struct {
	Contract     *CuratedModule    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CuratedModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CuratedModuleCallerSession struct {
	Contract *CuratedModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CuratedModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CuratedModuleTransactorSession struct {
	Contract     *CuratedModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CuratedModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type CuratedModuleRaw struct {
	Contract *CuratedModule // Generic contract binding to access the raw methods on
}

// CuratedModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CuratedModuleCallerRaw struct {
	Contract *CuratedModuleCaller // Generic read-only contract binding to access the raw methods on
}

// CuratedModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CuratedModuleTransactorRaw struct {
	Contract *CuratedModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCuratedModule creates a new instance of CuratedModule, bound to a specific deployed contract.
func NewCuratedModule(address common.Address, backend bind.ContractBackend) (*CuratedModule, error) {
	contract, err := bindCuratedModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CuratedModule{CuratedModuleCaller: CuratedModuleCaller{contract: contract}, CuratedModuleTransactor: CuratedModuleTransactor{contract: contract}, CuratedModuleFilterer: CuratedModuleFilterer{contract: contract}}, nil
}

// NewCuratedModuleCaller creates a new read-only instance of CuratedModule, bound to a specific deployed contract.
func NewCuratedModuleCaller(address common.Address, caller bind.ContractCaller) (*CuratedModuleCaller, error) {
	contract, err := bindCuratedModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleCaller{contract: contract}, nil
}

// NewCuratedModuleTransactor creates a new write-only instance of CuratedModule, bound to a specific deployed contract.
func NewCuratedModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*CuratedModuleTransactor, error) {
	contract, err := bindCuratedModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleTransactor{contract: contract}, nil
}

// NewCuratedModuleFilterer creates a new log filterer instance of CuratedModule, bound to a specific deployed contract.
func NewCuratedModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*CuratedModuleFilterer, error) {
	contract, err := bindCuratedModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleFilterer{contract: contract}, nil
}

// bindCuratedModule binds a generic wrapper to an already deployed contract.
func bindCuratedModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CuratedModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CuratedModule *CuratedModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CuratedModule.Contract.CuratedModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CuratedModule *CuratedModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CuratedModule.Contract.CuratedModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CuratedModule *CuratedModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CuratedModule.Contract.CuratedModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CuratedModule *CuratedModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CuratedModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CuratedModule *CuratedModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CuratedModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CuratedModule *CuratedModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CuratedModule.Contract.contract.Transact(opts, method, params...)
}

// ACCOUNTING is a free data retrieval call binding the contract method 0x6dc3f2bd.
//
// Solidity: function ACCOUNTING() view returns(address)
func (_CuratedModule *CuratedModuleCaller) ACCOUNTING(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "ACCOUNTING")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ACCOUNTING is a free data retrieval call binding the contract method 0x6dc3f2bd.
//
// Solidity: function ACCOUNTING() view returns(address)
func (_CuratedModule *CuratedModuleSession) ACCOUNTING() (common.Address, error) {
	return _CuratedModule.Contract.ACCOUNTING(&_CuratedModule.CallOpts)
}

// ACCOUNTING is a free data retrieval call binding the contract method 0x6dc3f2bd.
//
// Solidity: function ACCOUNTING() view returns(address)
func (_CuratedModule *CuratedModuleCallerSession) ACCOUNTING() (common.Address, error) {
	return _CuratedModule.Contract.ACCOUNTING(&_CuratedModule.CallOpts)
}

// CREATENODEOPERATORROLE is a free data retrieval call binding the contract method 0x743f5105.
//
// Solidity: function CREATE_NODE_OPERATOR_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleCaller) CREATENODEOPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "CREATE_NODE_OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CREATENODEOPERATORROLE is a free data retrieval call binding the contract method 0x743f5105.
//
// Solidity: function CREATE_NODE_OPERATOR_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleSession) CREATENODEOPERATORROLE() ([32]byte, error) {
	return _CuratedModule.Contract.CREATENODEOPERATORROLE(&_CuratedModule.CallOpts)
}

// CREATENODEOPERATORROLE is a free data retrieval call binding the contract method 0x743f5105.
//
// Solidity: function CREATE_NODE_OPERATOR_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleCallerSession) CREATENODEOPERATORROLE() ([32]byte, error) {
	return _CuratedModule.Contract.CREATENODEOPERATORROLE(&_CuratedModule.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CuratedModule.Contract.DEFAULTADMINROLE(&_CuratedModule.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CuratedModule.Contract.DEFAULTADMINROLE(&_CuratedModule.CallOpts)
}

// EXITPENALTIES is a free data retrieval call binding the contract method 0xfa367c9e.
//
// Solidity: function EXIT_PENALTIES() view returns(address)
func (_CuratedModule *CuratedModuleCaller) EXITPENALTIES(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "EXIT_PENALTIES")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EXITPENALTIES is a free data retrieval call binding the contract method 0xfa367c9e.
//
// Solidity: function EXIT_PENALTIES() view returns(address)
func (_CuratedModule *CuratedModuleSession) EXITPENALTIES() (common.Address, error) {
	return _CuratedModule.Contract.EXITPENALTIES(&_CuratedModule.CallOpts)
}

// EXITPENALTIES is a free data retrieval call binding the contract method 0xfa367c9e.
//
// Solidity: function EXIT_PENALTIES() view returns(address)
func (_CuratedModule *CuratedModuleCallerSession) EXITPENALTIES() (common.Address, error) {
	return _CuratedModule.Contract.EXITPENALTIES(&_CuratedModule.CallOpts)
}

// FEEDISTRIBUTOR is a free data retrieval call binding the contract method 0x6910dcce.
//
// Solidity: function FEE_DISTRIBUTOR() view returns(address)
func (_CuratedModule *CuratedModuleCaller) FEEDISTRIBUTOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "FEE_DISTRIBUTOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FEEDISTRIBUTOR is a free data retrieval call binding the contract method 0x6910dcce.
//
// Solidity: function FEE_DISTRIBUTOR() view returns(address)
func (_CuratedModule *CuratedModuleSession) FEEDISTRIBUTOR() (common.Address, error) {
	return _CuratedModule.Contract.FEEDISTRIBUTOR(&_CuratedModule.CallOpts)
}

// FEEDISTRIBUTOR is a free data retrieval call binding the contract method 0x6910dcce.
//
// Solidity: function FEE_DISTRIBUTOR() view returns(address)
func (_CuratedModule *CuratedModuleCallerSession) FEEDISTRIBUTOR() (common.Address, error) {
	return _CuratedModule.Contract.FEEDISTRIBUTOR(&_CuratedModule.CallOpts)
}

// LIDOLOCATOR is a free data retrieval call binding the contract method 0xdbba4b48.
//
// Solidity: function LIDO_LOCATOR() view returns(address)
func (_CuratedModule *CuratedModuleCaller) LIDOLOCATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "LIDO_LOCATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIDOLOCATOR is a free data retrieval call binding the contract method 0xdbba4b48.
//
// Solidity: function LIDO_LOCATOR() view returns(address)
func (_CuratedModule *CuratedModuleSession) LIDOLOCATOR() (common.Address, error) {
	return _CuratedModule.Contract.LIDOLOCATOR(&_CuratedModule.CallOpts)
}

// LIDOLOCATOR is a free data retrieval call binding the contract method 0xdbba4b48.
//
// Solidity: function LIDO_LOCATOR() view returns(address)
func (_CuratedModule *CuratedModuleCallerSession) LIDOLOCATOR() (common.Address, error) {
	return _CuratedModule.Contract.LIDOLOCATOR(&_CuratedModule.CallOpts)
}

// METAREGISTRY is a free data retrieval call binding the contract method 0x03a3279d.
//
// Solidity: function META_REGISTRY() view returns(address)
func (_CuratedModule *CuratedModuleCaller) METAREGISTRY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "META_REGISTRY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// METAREGISTRY is a free data retrieval call binding the contract method 0x03a3279d.
//
// Solidity: function META_REGISTRY() view returns(address)
func (_CuratedModule *CuratedModuleSession) METAREGISTRY() (common.Address, error) {
	return _CuratedModule.Contract.METAREGISTRY(&_CuratedModule.CallOpts)
}

// METAREGISTRY is a free data retrieval call binding the contract method 0x03a3279d.
//
// Solidity: function META_REGISTRY() view returns(address)
func (_CuratedModule *CuratedModuleCallerSession) METAREGISTRY() (common.Address, error) {
	return _CuratedModule.Contract.METAREGISTRY(&_CuratedModule.CallOpts)
}

// OPERATORADDRESSESADMINROLE is a free data retrieval call binding the contract method 0x0fcfcc41.
//
// Solidity: function OPERATOR_ADDRESSES_ADMIN_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleCaller) OPERATORADDRESSESADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "OPERATOR_ADDRESSES_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORADDRESSESADMINROLE is a free data retrieval call binding the contract method 0x0fcfcc41.
//
// Solidity: function OPERATOR_ADDRESSES_ADMIN_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleSession) OPERATORADDRESSESADMINROLE() ([32]byte, error) {
	return _CuratedModule.Contract.OPERATORADDRESSESADMINROLE(&_CuratedModule.CallOpts)
}

// OPERATORADDRESSESADMINROLE is a free data retrieval call binding the contract method 0x0fcfcc41.
//
// Solidity: function OPERATOR_ADDRESSES_ADMIN_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleCallerSession) OPERATORADDRESSESADMINROLE() ([32]byte, error) {
	return _CuratedModule.Contract.OPERATORADDRESSESADMINROLE(&_CuratedModule.CallOpts)
}

// PARAMETERSREGISTRY is a free data retrieval call binding the contract method 0x2fc88741.
//
// Solidity: function PARAMETERS_REGISTRY() view returns(address)
func (_CuratedModule *CuratedModuleCaller) PARAMETERSREGISTRY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "PARAMETERS_REGISTRY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PARAMETERSREGISTRY is a free data retrieval call binding the contract method 0x2fc88741.
//
// Solidity: function PARAMETERS_REGISTRY() view returns(address)
func (_CuratedModule *CuratedModuleSession) PARAMETERSREGISTRY() (common.Address, error) {
	return _CuratedModule.Contract.PARAMETERSREGISTRY(&_CuratedModule.CallOpts)
}

// PARAMETERSREGISTRY is a free data retrieval call binding the contract method 0x2fc88741.
//
// Solidity: function PARAMETERS_REGISTRY() view returns(address)
func (_CuratedModule *CuratedModuleCallerSession) PARAMETERSREGISTRY() (common.Address, error) {
	return _CuratedModule.Contract.PARAMETERSREGISTRY(&_CuratedModule.CallOpts)
}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_CuratedModule *CuratedModuleCaller) PAUSEINFINITELY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "PAUSE_INFINITELY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_CuratedModule *CuratedModuleSession) PAUSEINFINITELY() (*big.Int, error) {
	return _CuratedModule.Contract.PAUSEINFINITELY(&_CuratedModule.CallOpts)
}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_CuratedModule *CuratedModuleCallerSession) PAUSEINFINITELY() (*big.Int, error) {
	return _CuratedModule.Contract.PAUSEINFINITELY(&_CuratedModule.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleCaller) PAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleSession) PAUSEROLE() ([32]byte, error) {
	return _CuratedModule.Contract.PAUSEROLE(&_CuratedModule.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleCallerSession) PAUSEROLE() ([32]byte, error) {
	return _CuratedModule.Contract.PAUSEROLE(&_CuratedModule.CallOpts)
}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleCaller) RECOVERERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "RECOVERER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleSession) RECOVERERROLE() ([32]byte, error) {
	return _CuratedModule.Contract.RECOVERERROLE(&_CuratedModule.CallOpts)
}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleCallerSession) RECOVERERROLE() ([32]byte, error) {
	return _CuratedModule.Contract.RECOVERERROLE(&_CuratedModule.CallOpts)
}

// REPORTGENERALDELAYEDPENALTYROLE is a free data retrieval call binding the contract method 0xf8b92d65.
//
// Solidity: function REPORT_GENERAL_DELAYED_PENALTY_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleCaller) REPORTGENERALDELAYEDPENALTYROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "REPORT_GENERAL_DELAYED_PENALTY_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REPORTGENERALDELAYEDPENALTYROLE is a free data retrieval call binding the contract method 0xf8b92d65.
//
// Solidity: function REPORT_GENERAL_DELAYED_PENALTY_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleSession) REPORTGENERALDELAYEDPENALTYROLE() ([32]byte, error) {
	return _CuratedModule.Contract.REPORTGENERALDELAYEDPENALTYROLE(&_CuratedModule.CallOpts)
}

// REPORTGENERALDELAYEDPENALTYROLE is a free data retrieval call binding the contract method 0xf8b92d65.
//
// Solidity: function REPORT_GENERAL_DELAYED_PENALTY_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleCallerSession) REPORTGENERALDELAYEDPENALTYROLE() ([32]byte, error) {
	return _CuratedModule.Contract.REPORTGENERALDELAYEDPENALTYROLE(&_CuratedModule.CallOpts)
}

// REPORTREGULARWITHDRAWNVALIDATORSROLE is a free data retrieval call binding the contract method 0x87f4caf5.
//
// Solidity: function REPORT_REGULAR_WITHDRAWN_VALIDATORS_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleCaller) REPORTREGULARWITHDRAWNVALIDATORSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "REPORT_REGULAR_WITHDRAWN_VALIDATORS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REPORTREGULARWITHDRAWNVALIDATORSROLE is a free data retrieval call binding the contract method 0x87f4caf5.
//
// Solidity: function REPORT_REGULAR_WITHDRAWN_VALIDATORS_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleSession) REPORTREGULARWITHDRAWNVALIDATORSROLE() ([32]byte, error) {
	return _CuratedModule.Contract.REPORTREGULARWITHDRAWNVALIDATORSROLE(&_CuratedModule.CallOpts)
}

// REPORTREGULARWITHDRAWNVALIDATORSROLE is a free data retrieval call binding the contract method 0x87f4caf5.
//
// Solidity: function REPORT_REGULAR_WITHDRAWN_VALIDATORS_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleCallerSession) REPORTREGULARWITHDRAWNVALIDATORSROLE() ([32]byte, error) {
	return _CuratedModule.Contract.REPORTREGULARWITHDRAWNVALIDATORSROLE(&_CuratedModule.CallOpts)
}

// REPORTSLASHEDWITHDRAWNVALIDATORSROLE is a free data retrieval call binding the contract method 0x39a4ffd6.
//
// Solidity: function REPORT_SLASHED_WITHDRAWN_VALIDATORS_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleCaller) REPORTSLASHEDWITHDRAWNVALIDATORSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "REPORT_SLASHED_WITHDRAWN_VALIDATORS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REPORTSLASHEDWITHDRAWNVALIDATORSROLE is a free data retrieval call binding the contract method 0x39a4ffd6.
//
// Solidity: function REPORT_SLASHED_WITHDRAWN_VALIDATORS_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleSession) REPORTSLASHEDWITHDRAWNVALIDATORSROLE() ([32]byte, error) {
	return _CuratedModule.Contract.REPORTSLASHEDWITHDRAWNVALIDATORSROLE(&_CuratedModule.CallOpts)
}

// REPORTSLASHEDWITHDRAWNVALIDATORSROLE is a free data retrieval call binding the contract method 0x39a4ffd6.
//
// Solidity: function REPORT_SLASHED_WITHDRAWN_VALIDATORS_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleCallerSession) REPORTSLASHEDWITHDRAWNVALIDATORSROLE() ([32]byte, error) {
	return _CuratedModule.Contract.REPORTSLASHEDWITHDRAWNVALIDATORSROLE(&_CuratedModule.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleCaller) RESUMEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "RESUME_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleSession) RESUMEROLE() ([32]byte, error) {
	return _CuratedModule.Contract.RESUMEROLE(&_CuratedModule.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleCallerSession) RESUMEROLE() ([32]byte, error) {
	return _CuratedModule.Contract.RESUMEROLE(&_CuratedModule.CallOpts)
}

// SETTLEGENERALDELAYEDPENALTYROLE is a free data retrieval call binding the contract method 0x7a41d945.
//
// Solidity: function SETTLE_GENERAL_DELAYED_PENALTY_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleCaller) SETTLEGENERALDELAYEDPENALTYROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "SETTLE_GENERAL_DELAYED_PENALTY_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETTLEGENERALDELAYEDPENALTYROLE is a free data retrieval call binding the contract method 0x7a41d945.
//
// Solidity: function SETTLE_GENERAL_DELAYED_PENALTY_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleSession) SETTLEGENERALDELAYEDPENALTYROLE() ([32]byte, error) {
	return _CuratedModule.Contract.SETTLEGENERALDELAYEDPENALTYROLE(&_CuratedModule.CallOpts)
}

// SETTLEGENERALDELAYEDPENALTYROLE is a free data retrieval call binding the contract method 0x7a41d945.
//
// Solidity: function SETTLE_GENERAL_DELAYED_PENALTY_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleCallerSession) SETTLEGENERALDELAYEDPENALTYROLE() ([32]byte, error) {
	return _CuratedModule.Contract.SETTLEGENERALDELAYEDPENALTYROLE(&_CuratedModule.CallOpts)
}

// STAKINGROUTERROLE is a free data retrieval call binding the contract method 0x80231f15.
//
// Solidity: function STAKING_ROUTER_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleCaller) STAKINGROUTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "STAKING_ROUTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGROUTERROLE is a free data retrieval call binding the contract method 0x80231f15.
//
// Solidity: function STAKING_ROUTER_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleSession) STAKINGROUTERROLE() ([32]byte, error) {
	return _CuratedModule.Contract.STAKINGROUTERROLE(&_CuratedModule.CallOpts)
}

// STAKINGROUTERROLE is a free data retrieval call binding the contract method 0x80231f15.
//
// Solidity: function STAKING_ROUTER_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleCallerSession) STAKINGROUTERROLE() ([32]byte, error) {
	return _CuratedModule.Contract.STAKINGROUTERROLE(&_CuratedModule.CallOpts)
}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_CuratedModule *CuratedModuleCaller) STETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "STETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_CuratedModule *CuratedModuleSession) STETH() (common.Address, error) {
	return _CuratedModule.Contract.STETH(&_CuratedModule.CallOpts)
}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_CuratedModule *CuratedModuleCallerSession) STETH() (common.Address, error) {
	return _CuratedModule.Contract.STETH(&_CuratedModule.CallOpts)
}

// VERIFIERROLE is a free data retrieval call binding the contract method 0xe7705db6.
//
// Solidity: function VERIFIER_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleCaller) VERIFIERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "VERIFIER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VERIFIERROLE is a free data retrieval call binding the contract method 0xe7705db6.
//
// Solidity: function VERIFIER_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleSession) VERIFIERROLE() ([32]byte, error) {
	return _CuratedModule.Contract.VERIFIERROLE(&_CuratedModule.CallOpts)
}

// VERIFIERROLE is a free data retrieval call binding the contract method 0xe7705db6.
//
// Solidity: function VERIFIER_ROLE() view returns(bytes32)
func (_CuratedModule *CuratedModuleCallerSession) VERIFIERROLE() ([32]byte, error) {
	return _CuratedModule.Contract.VERIFIERROLE(&_CuratedModule.CallOpts)
}

// ExitDeadlineThreshold is a free data retrieval call binding the contract method 0x28d6d36b.
//
// Solidity: function exitDeadlineThreshold(uint256 nodeOperatorId) view returns(uint256)
func (_CuratedModule *CuratedModuleCaller) ExitDeadlineThreshold(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "exitDeadlineThreshold", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExitDeadlineThreshold is a free data retrieval call binding the contract method 0x28d6d36b.
//
// Solidity: function exitDeadlineThreshold(uint256 nodeOperatorId) view returns(uint256)
func (_CuratedModule *CuratedModuleSession) ExitDeadlineThreshold(nodeOperatorId *big.Int) (*big.Int, error) {
	return _CuratedModule.Contract.ExitDeadlineThreshold(&_CuratedModule.CallOpts, nodeOperatorId)
}

// ExitDeadlineThreshold is a free data retrieval call binding the contract method 0x28d6d36b.
//
// Solidity: function exitDeadlineThreshold(uint256 nodeOperatorId) view returns(uint256)
func (_CuratedModule *CuratedModuleCallerSession) ExitDeadlineThreshold(nodeOperatorId *big.Int) (*big.Int, error) {
	return _CuratedModule.Contract.ExitDeadlineThreshold(&_CuratedModule.CallOpts, nodeOperatorId)
}

// GetActiveNodeOperatorsCount is a free data retrieval call binding the contract method 0x8469cbd3.
//
// Solidity: function getActiveNodeOperatorsCount() view returns(uint256)
func (_CuratedModule *CuratedModuleCaller) GetActiveNodeOperatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getActiveNodeOperatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveNodeOperatorsCount is a free data retrieval call binding the contract method 0x8469cbd3.
//
// Solidity: function getActiveNodeOperatorsCount() view returns(uint256)
func (_CuratedModule *CuratedModuleSession) GetActiveNodeOperatorsCount() (*big.Int, error) {
	return _CuratedModule.Contract.GetActiveNodeOperatorsCount(&_CuratedModule.CallOpts)
}

// GetActiveNodeOperatorsCount is a free data retrieval call binding the contract method 0x8469cbd3.
//
// Solidity: function getActiveNodeOperatorsCount() view returns(uint256)
func (_CuratedModule *CuratedModuleCallerSession) GetActiveNodeOperatorsCount() (*big.Int, error) {
	return _CuratedModule.Contract.GetActiveNodeOperatorsCount(&_CuratedModule.CallOpts)
}

// GetDepositAllocationTargets is a free data retrieval call binding the contract method 0xc3b497f9.
//
// Solidity: function getDepositAllocationTargets() view returns(uint256[] currentValidators, uint256[] targetValidators)
func (_CuratedModule *CuratedModuleCaller) GetDepositAllocationTargets(opts *bind.CallOpts) (struct {
	CurrentValidators []*big.Int
	TargetValidators  []*big.Int
}, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getDepositAllocationTargets")

	outstruct := new(struct {
		CurrentValidators []*big.Int
		TargetValidators  []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentValidators = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.TargetValidators = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetDepositAllocationTargets is a free data retrieval call binding the contract method 0xc3b497f9.
//
// Solidity: function getDepositAllocationTargets() view returns(uint256[] currentValidators, uint256[] targetValidators)
func (_CuratedModule *CuratedModuleSession) GetDepositAllocationTargets() (struct {
	CurrentValidators []*big.Int
	TargetValidators  []*big.Int
}, error) {
	return _CuratedModule.Contract.GetDepositAllocationTargets(&_CuratedModule.CallOpts)
}

// GetDepositAllocationTargets is a free data retrieval call binding the contract method 0xc3b497f9.
//
// Solidity: function getDepositAllocationTargets() view returns(uint256[] currentValidators, uint256[] targetValidators)
func (_CuratedModule *CuratedModuleCallerSession) GetDepositAllocationTargets() (struct {
	CurrentValidators []*big.Int
	TargetValidators  []*big.Int
}, error) {
	return _CuratedModule.Contract.GetDepositAllocationTargets(&_CuratedModule.CallOpts)
}

// GetDepositsAllocation is a free data retrieval call binding the contract method 0xc82b1bb1.
//
// Solidity: function getDepositsAllocation(uint256 maxDepositAmount) view returns(uint256 allocated, uint256[] operatorIds, uint256[] allocations)
func (_CuratedModule *CuratedModuleCaller) GetDepositsAllocation(opts *bind.CallOpts, maxDepositAmount *big.Int) (struct {
	Allocated   *big.Int
	OperatorIds []*big.Int
	Allocations []*big.Int
}, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getDepositsAllocation", maxDepositAmount)

	outstruct := new(struct {
		Allocated   *big.Int
		OperatorIds []*big.Int
		Allocations []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Allocated = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.OperatorIds = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Allocations = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetDepositsAllocation is a free data retrieval call binding the contract method 0xc82b1bb1.
//
// Solidity: function getDepositsAllocation(uint256 maxDepositAmount) view returns(uint256 allocated, uint256[] operatorIds, uint256[] allocations)
func (_CuratedModule *CuratedModuleSession) GetDepositsAllocation(maxDepositAmount *big.Int) (struct {
	Allocated   *big.Int
	OperatorIds []*big.Int
	Allocations []*big.Int
}, error) {
	return _CuratedModule.Contract.GetDepositsAllocation(&_CuratedModule.CallOpts, maxDepositAmount)
}

// GetDepositsAllocation is a free data retrieval call binding the contract method 0xc82b1bb1.
//
// Solidity: function getDepositsAllocation(uint256 maxDepositAmount) view returns(uint256 allocated, uint256[] operatorIds, uint256[] allocations)
func (_CuratedModule *CuratedModuleCallerSession) GetDepositsAllocation(maxDepositAmount *big.Int) (struct {
	Allocated   *big.Int
	OperatorIds []*big.Int
	Allocations []*big.Int
}, error) {
	return _CuratedModule.Contract.GetDepositsAllocation(&_CuratedModule.CallOpts, maxDepositAmount)
}

// GetInitializedVersion is a free data retrieval call binding the contract method 0xb3c65015.
//
// Solidity: function getInitializedVersion() view returns(uint64)
func (_CuratedModule *CuratedModuleCaller) GetInitializedVersion(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getInitializedVersion")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetInitializedVersion is a free data retrieval call binding the contract method 0xb3c65015.
//
// Solidity: function getInitializedVersion() view returns(uint64)
func (_CuratedModule *CuratedModuleSession) GetInitializedVersion() (uint64, error) {
	return _CuratedModule.Contract.GetInitializedVersion(&_CuratedModule.CallOpts)
}

// GetInitializedVersion is a free data retrieval call binding the contract method 0xb3c65015.
//
// Solidity: function getInitializedVersion() view returns(uint64)
func (_CuratedModule *CuratedModuleCallerSession) GetInitializedVersion() (uint64, error) {
	return _CuratedModule.Contract.GetInitializedVersion(&_CuratedModule.CallOpts)
}

// GetKeyAllocatedBalances is a free data retrieval call binding the contract method 0x0547c962.
//
// Solidity: function getKeyAllocatedBalances(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(uint256[] balances)
func (_CuratedModule *CuratedModuleCaller) GetKeyAllocatedBalances(opts *bind.CallOpts, nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getKeyAllocatedBalances", nodeOperatorId, startIndex, keysCount)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetKeyAllocatedBalances is a free data retrieval call binding the contract method 0x0547c962.
//
// Solidity: function getKeyAllocatedBalances(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(uint256[] balances)
func (_CuratedModule *CuratedModuleSession) GetKeyAllocatedBalances(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) ([]*big.Int, error) {
	return _CuratedModule.Contract.GetKeyAllocatedBalances(&_CuratedModule.CallOpts, nodeOperatorId, startIndex, keysCount)
}

// GetKeyAllocatedBalances is a free data retrieval call binding the contract method 0x0547c962.
//
// Solidity: function getKeyAllocatedBalances(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(uint256[] balances)
func (_CuratedModule *CuratedModuleCallerSession) GetKeyAllocatedBalances(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) ([]*big.Int, error) {
	return _CuratedModule.Contract.GetKeyAllocatedBalances(&_CuratedModule.CallOpts, nodeOperatorId, startIndex, keysCount)
}

// GetKeyConfirmedBalances is a free data retrieval call binding the contract method 0xcd0af35c.
//
// Solidity: function getKeyConfirmedBalances(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(uint256[] balances)
func (_CuratedModule *CuratedModuleCaller) GetKeyConfirmedBalances(opts *bind.CallOpts, nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getKeyConfirmedBalances", nodeOperatorId, startIndex, keysCount)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetKeyConfirmedBalances is a free data retrieval call binding the contract method 0xcd0af35c.
//
// Solidity: function getKeyConfirmedBalances(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(uint256[] balances)
func (_CuratedModule *CuratedModuleSession) GetKeyConfirmedBalances(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) ([]*big.Int, error) {
	return _CuratedModule.Contract.GetKeyConfirmedBalances(&_CuratedModule.CallOpts, nodeOperatorId, startIndex, keysCount)
}

// GetKeyConfirmedBalances is a free data retrieval call binding the contract method 0xcd0af35c.
//
// Solidity: function getKeyConfirmedBalances(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(uint256[] balances)
func (_CuratedModule *CuratedModuleCallerSession) GetKeyConfirmedBalances(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) ([]*big.Int, error) {
	return _CuratedModule.Contract.GetKeyConfirmedBalances(&_CuratedModule.CallOpts, nodeOperatorId, startIndex, keysCount)
}

// GetNodeOperator is a free data retrieval call binding the contract method 0x65c14dc7.
//
// Solidity: function getNodeOperator(uint256 nodeOperatorId) view returns((uint32,uint32,uint32,uint32,uint32,uint32,uint32,uint8,uint32,uint32,address,address,address,address,bool,bool))
func (_CuratedModule *CuratedModuleCaller) GetNodeOperator(opts *bind.CallOpts, nodeOperatorId *big.Int) (NodeOperator, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getNodeOperator", nodeOperatorId)

	if err != nil {
		return *new(NodeOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(NodeOperator)).(*NodeOperator)

	return out0, err

}

// GetNodeOperator is a free data retrieval call binding the contract method 0x65c14dc7.
//
// Solidity: function getNodeOperator(uint256 nodeOperatorId) view returns((uint32,uint32,uint32,uint32,uint32,uint32,uint32,uint8,uint32,uint32,address,address,address,address,bool,bool))
func (_CuratedModule *CuratedModuleSession) GetNodeOperator(nodeOperatorId *big.Int) (NodeOperator, error) {
	return _CuratedModule.Contract.GetNodeOperator(&_CuratedModule.CallOpts, nodeOperatorId)
}

// GetNodeOperator is a free data retrieval call binding the contract method 0x65c14dc7.
//
// Solidity: function getNodeOperator(uint256 nodeOperatorId) view returns((uint32,uint32,uint32,uint32,uint32,uint32,uint32,uint8,uint32,uint32,address,address,address,address,bool,bool))
func (_CuratedModule *CuratedModuleCallerSession) GetNodeOperator(nodeOperatorId *big.Int) (NodeOperator, error) {
	return _CuratedModule.Contract.GetNodeOperator(&_CuratedModule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorBalance is a free data retrieval call binding the contract method 0xd2e878a6.
//
// Solidity: function getNodeOperatorBalance(uint256 nodeOperatorId) view returns(uint256)
func (_CuratedModule *CuratedModuleCaller) GetNodeOperatorBalance(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getNodeOperatorBalance", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeOperatorBalance is a free data retrieval call binding the contract method 0xd2e878a6.
//
// Solidity: function getNodeOperatorBalance(uint256 nodeOperatorId) view returns(uint256)
func (_CuratedModule *CuratedModuleSession) GetNodeOperatorBalance(nodeOperatorId *big.Int) (*big.Int, error) {
	return _CuratedModule.Contract.GetNodeOperatorBalance(&_CuratedModule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorBalance is a free data retrieval call binding the contract method 0xd2e878a6.
//
// Solidity: function getNodeOperatorBalance(uint256 nodeOperatorId) view returns(uint256)
func (_CuratedModule *CuratedModuleCallerSession) GetNodeOperatorBalance(nodeOperatorId *big.Int) (*big.Int, error) {
	return _CuratedModule.Contract.GetNodeOperatorBalance(&_CuratedModule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorDepositInfoToUpdateCount is a free data retrieval call binding the contract method 0xa1b51f7b.
//
// Solidity: function getNodeOperatorDepositInfoToUpdateCount() view returns(uint256 count)
func (_CuratedModule *CuratedModuleCaller) GetNodeOperatorDepositInfoToUpdateCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getNodeOperatorDepositInfoToUpdateCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeOperatorDepositInfoToUpdateCount is a free data retrieval call binding the contract method 0xa1b51f7b.
//
// Solidity: function getNodeOperatorDepositInfoToUpdateCount() view returns(uint256 count)
func (_CuratedModule *CuratedModuleSession) GetNodeOperatorDepositInfoToUpdateCount() (*big.Int, error) {
	return _CuratedModule.Contract.GetNodeOperatorDepositInfoToUpdateCount(&_CuratedModule.CallOpts)
}

// GetNodeOperatorDepositInfoToUpdateCount is a free data retrieval call binding the contract method 0xa1b51f7b.
//
// Solidity: function getNodeOperatorDepositInfoToUpdateCount() view returns(uint256 count)
func (_CuratedModule *CuratedModuleCallerSession) GetNodeOperatorDepositInfoToUpdateCount() (*big.Int, error) {
	return _CuratedModule.Contract.GetNodeOperatorDepositInfoToUpdateCount(&_CuratedModule.CallOpts)
}

// GetNodeOperatorIds is a free data retrieval call binding the contract method 0x4febc81b.
//
// Solidity: function getNodeOperatorIds(uint256 offset, uint256 limit) view returns(uint256[] nodeOperatorIds)
func (_CuratedModule *CuratedModuleCaller) GetNodeOperatorIds(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getNodeOperatorIds", offset, limit)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetNodeOperatorIds is a free data retrieval call binding the contract method 0x4febc81b.
//
// Solidity: function getNodeOperatorIds(uint256 offset, uint256 limit) view returns(uint256[] nodeOperatorIds)
func (_CuratedModule *CuratedModuleSession) GetNodeOperatorIds(offset *big.Int, limit *big.Int) ([]*big.Int, error) {
	return _CuratedModule.Contract.GetNodeOperatorIds(&_CuratedModule.CallOpts, offset, limit)
}

// GetNodeOperatorIds is a free data retrieval call binding the contract method 0x4febc81b.
//
// Solidity: function getNodeOperatorIds(uint256 offset, uint256 limit) view returns(uint256[] nodeOperatorIds)
func (_CuratedModule *CuratedModuleCallerSession) GetNodeOperatorIds(offset *big.Int, limit *big.Int) ([]*big.Int, error) {
	return _CuratedModule.Contract.GetNodeOperatorIds(&_CuratedModule.CallOpts, offset, limit)
}

// GetNodeOperatorIsActive is a free data retrieval call binding the contract method 0x5e2fb908.
//
// Solidity: function getNodeOperatorIsActive(uint256 nodeOperatorId) view returns(bool)
func (_CuratedModule *CuratedModuleCaller) GetNodeOperatorIsActive(opts *bind.CallOpts, nodeOperatorId *big.Int) (bool, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getNodeOperatorIsActive", nodeOperatorId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetNodeOperatorIsActive is a free data retrieval call binding the contract method 0x5e2fb908.
//
// Solidity: function getNodeOperatorIsActive(uint256 nodeOperatorId) view returns(bool)
func (_CuratedModule *CuratedModuleSession) GetNodeOperatorIsActive(nodeOperatorId *big.Int) (bool, error) {
	return _CuratedModule.Contract.GetNodeOperatorIsActive(&_CuratedModule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorIsActive is a free data retrieval call binding the contract method 0x5e2fb908.
//
// Solidity: function getNodeOperatorIsActive(uint256 nodeOperatorId) view returns(bool)
func (_CuratedModule *CuratedModuleCallerSession) GetNodeOperatorIsActive(nodeOperatorId *big.Int) (bool, error) {
	return _CuratedModule.Contract.GetNodeOperatorIsActive(&_CuratedModule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorManagementProperties is a free data retrieval call binding the contract method 0x499b8e9a.
//
// Solidity: function getNodeOperatorManagementProperties(uint256 nodeOperatorId) view returns((address,address,bool))
func (_CuratedModule *CuratedModuleCaller) GetNodeOperatorManagementProperties(opts *bind.CallOpts, nodeOperatorId *big.Int) (NodeOperatorManagementProperties, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getNodeOperatorManagementProperties", nodeOperatorId)

	if err != nil {
		return *new(NodeOperatorManagementProperties), err
	}

	out0 := *abi.ConvertType(out[0], new(NodeOperatorManagementProperties)).(*NodeOperatorManagementProperties)

	return out0, err

}

// GetNodeOperatorManagementProperties is a free data retrieval call binding the contract method 0x499b8e9a.
//
// Solidity: function getNodeOperatorManagementProperties(uint256 nodeOperatorId) view returns((address,address,bool))
func (_CuratedModule *CuratedModuleSession) GetNodeOperatorManagementProperties(nodeOperatorId *big.Int) (NodeOperatorManagementProperties, error) {
	return _CuratedModule.Contract.GetNodeOperatorManagementProperties(&_CuratedModule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorManagementProperties is a free data retrieval call binding the contract method 0x499b8e9a.
//
// Solidity: function getNodeOperatorManagementProperties(uint256 nodeOperatorId) view returns((address,address,bool))
func (_CuratedModule *CuratedModuleCallerSession) GetNodeOperatorManagementProperties(nodeOperatorId *big.Int) (NodeOperatorManagementProperties, error) {
	return _CuratedModule.Contract.GetNodeOperatorManagementProperties(&_CuratedModule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorNonWithdrawnKeys is a free data retrieval call binding the contract method 0x8ec69028.
//
// Solidity: function getNodeOperatorNonWithdrawnKeys(uint256 nodeOperatorId) view returns(uint256)
func (_CuratedModule *CuratedModuleCaller) GetNodeOperatorNonWithdrawnKeys(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getNodeOperatorNonWithdrawnKeys", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeOperatorNonWithdrawnKeys is a free data retrieval call binding the contract method 0x8ec69028.
//
// Solidity: function getNodeOperatorNonWithdrawnKeys(uint256 nodeOperatorId) view returns(uint256)
func (_CuratedModule *CuratedModuleSession) GetNodeOperatorNonWithdrawnKeys(nodeOperatorId *big.Int) (*big.Int, error) {
	return _CuratedModule.Contract.GetNodeOperatorNonWithdrawnKeys(&_CuratedModule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorNonWithdrawnKeys is a free data retrieval call binding the contract method 0x8ec69028.
//
// Solidity: function getNodeOperatorNonWithdrawnKeys(uint256 nodeOperatorId) view returns(uint256)
func (_CuratedModule *CuratedModuleCallerSession) GetNodeOperatorNonWithdrawnKeys(nodeOperatorId *big.Int) (*big.Int, error) {
	return _CuratedModule.Contract.GetNodeOperatorNonWithdrawnKeys(&_CuratedModule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorOwner is a free data retrieval call binding the contract method 0xb055e15c.
//
// Solidity: function getNodeOperatorOwner(uint256 nodeOperatorId) view returns(address)
func (_CuratedModule *CuratedModuleCaller) GetNodeOperatorOwner(opts *bind.CallOpts, nodeOperatorId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getNodeOperatorOwner", nodeOperatorId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNodeOperatorOwner is a free data retrieval call binding the contract method 0xb055e15c.
//
// Solidity: function getNodeOperatorOwner(uint256 nodeOperatorId) view returns(address)
func (_CuratedModule *CuratedModuleSession) GetNodeOperatorOwner(nodeOperatorId *big.Int) (common.Address, error) {
	return _CuratedModule.Contract.GetNodeOperatorOwner(&_CuratedModule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorOwner is a free data retrieval call binding the contract method 0xb055e15c.
//
// Solidity: function getNodeOperatorOwner(uint256 nodeOperatorId) view returns(address)
func (_CuratedModule *CuratedModuleCallerSession) GetNodeOperatorOwner(nodeOperatorId *big.Int) (common.Address, error) {
	return _CuratedModule.Contract.GetNodeOperatorOwner(&_CuratedModule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorSummary is a free data retrieval call binding the contract method 0xb3076c3c.
//
// Solidity: function getNodeOperatorSummary(uint256 nodeOperatorId) view returns(uint256 targetLimitMode, uint256 targetValidatorsCount, uint256 stuckValidatorsCount, uint256 refundedValidatorsCount, uint256 stuckPenaltyEndTimestamp, uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_CuratedModule *CuratedModuleCaller) GetNodeOperatorSummary(opts *bind.CallOpts, nodeOperatorId *big.Int) (struct {
	TargetLimitMode            *big.Int
	TargetValidatorsCount      *big.Int
	StuckValidatorsCount       *big.Int
	RefundedValidatorsCount    *big.Int
	StuckPenaltyEndTimestamp   *big.Int
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getNodeOperatorSummary", nodeOperatorId)

	outstruct := new(struct {
		TargetLimitMode            *big.Int
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

	outstruct.TargetLimitMode = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
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
// Solidity: function getNodeOperatorSummary(uint256 nodeOperatorId) view returns(uint256 targetLimitMode, uint256 targetValidatorsCount, uint256 stuckValidatorsCount, uint256 refundedValidatorsCount, uint256 stuckPenaltyEndTimestamp, uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_CuratedModule *CuratedModuleSession) GetNodeOperatorSummary(nodeOperatorId *big.Int) (struct {
	TargetLimitMode            *big.Int
	TargetValidatorsCount      *big.Int
	StuckValidatorsCount       *big.Int
	RefundedValidatorsCount    *big.Int
	StuckPenaltyEndTimestamp   *big.Int
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	return _CuratedModule.Contract.GetNodeOperatorSummary(&_CuratedModule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorSummary is a free data retrieval call binding the contract method 0xb3076c3c.
//
// Solidity: function getNodeOperatorSummary(uint256 nodeOperatorId) view returns(uint256 targetLimitMode, uint256 targetValidatorsCount, uint256 stuckValidatorsCount, uint256 refundedValidatorsCount, uint256 stuckPenaltyEndTimestamp, uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_CuratedModule *CuratedModuleCallerSession) GetNodeOperatorSummary(nodeOperatorId *big.Int) (struct {
	TargetLimitMode            *big.Int
	TargetValidatorsCount      *big.Int
	StuckValidatorsCount       *big.Int
	RefundedValidatorsCount    *big.Int
	StuckPenaltyEndTimestamp   *big.Int
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	return _CuratedModule.Contract.GetNodeOperatorSummary(&_CuratedModule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorWeightAndExternalStake is a free data retrieval call binding the contract method 0x9703e7bc.
//
// Solidity: function getNodeOperatorWeightAndExternalStake(uint256 nodeOperatorId) view returns(uint256 weight, uint256 externalStake)
func (_CuratedModule *CuratedModuleCaller) GetNodeOperatorWeightAndExternalStake(opts *bind.CallOpts, nodeOperatorId *big.Int) (struct {
	Weight        *big.Int
	ExternalStake *big.Int
}, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getNodeOperatorWeightAndExternalStake", nodeOperatorId)

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
// Solidity: function getNodeOperatorWeightAndExternalStake(uint256 nodeOperatorId) view returns(uint256 weight, uint256 externalStake)
func (_CuratedModule *CuratedModuleSession) GetNodeOperatorWeightAndExternalStake(nodeOperatorId *big.Int) (struct {
	Weight        *big.Int
	ExternalStake *big.Int
}, error) {
	return _CuratedModule.Contract.GetNodeOperatorWeightAndExternalStake(&_CuratedModule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorWeightAndExternalStake is a free data retrieval call binding the contract method 0x9703e7bc.
//
// Solidity: function getNodeOperatorWeightAndExternalStake(uint256 nodeOperatorId) view returns(uint256 weight, uint256 externalStake)
func (_CuratedModule *CuratedModuleCallerSession) GetNodeOperatorWeightAndExternalStake(nodeOperatorId *big.Int) (struct {
	Weight        *big.Int
	ExternalStake *big.Int
}, error) {
	return _CuratedModule.Contract.GetNodeOperatorWeightAndExternalStake(&_CuratedModule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorsCount is a free data retrieval call binding the contract method 0xa70c70e4.
//
// Solidity: function getNodeOperatorsCount() view returns(uint256)
func (_CuratedModule *CuratedModuleCaller) GetNodeOperatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getNodeOperatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeOperatorsCount is a free data retrieval call binding the contract method 0xa70c70e4.
//
// Solidity: function getNodeOperatorsCount() view returns(uint256)
func (_CuratedModule *CuratedModuleSession) GetNodeOperatorsCount() (*big.Int, error) {
	return _CuratedModule.Contract.GetNodeOperatorsCount(&_CuratedModule.CallOpts)
}

// GetNodeOperatorsCount is a free data retrieval call binding the contract method 0xa70c70e4.
//
// Solidity: function getNodeOperatorsCount() view returns(uint256)
func (_CuratedModule *CuratedModuleCallerSession) GetNodeOperatorsCount() (*big.Int, error) {
	return _CuratedModule.Contract.GetNodeOperatorsCount(&_CuratedModule.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_CuratedModule *CuratedModuleCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_CuratedModule *CuratedModuleSession) GetNonce() (*big.Int, error) {
	return _CuratedModule.Contract.GetNonce(&_CuratedModule.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_CuratedModule *CuratedModuleCallerSession) GetNonce() (*big.Int, error) {
	return _CuratedModule.Contract.GetNonce(&_CuratedModule.CallOpts)
}

// GetOperatorWeights is a free data retrieval call binding the contract method 0x08ca8c6d.
//
// Solidity: function getOperatorWeights(uint256[] operatorIds) view returns(uint256[] operatorWeights)
func (_CuratedModule *CuratedModuleCaller) GetOperatorWeights(opts *bind.CallOpts, operatorIds []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getOperatorWeights", operatorIds)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOperatorWeights is a free data retrieval call binding the contract method 0x08ca8c6d.
//
// Solidity: function getOperatorWeights(uint256[] operatorIds) view returns(uint256[] operatorWeights)
func (_CuratedModule *CuratedModuleSession) GetOperatorWeights(operatorIds []*big.Int) ([]*big.Int, error) {
	return _CuratedModule.Contract.GetOperatorWeights(&_CuratedModule.CallOpts, operatorIds)
}

// GetOperatorWeights is a free data retrieval call binding the contract method 0x08ca8c6d.
//
// Solidity: function getOperatorWeights(uint256[] operatorIds) view returns(uint256[] operatorWeights)
func (_CuratedModule *CuratedModuleCallerSession) GetOperatorWeights(operatorIds []*big.Int) ([]*big.Int, error) {
	return _CuratedModule.Contract.GetOperatorWeights(&_CuratedModule.CallOpts, operatorIds)
}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_CuratedModule *CuratedModuleCaller) GetResumeSinceTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getResumeSinceTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_CuratedModule *CuratedModuleSession) GetResumeSinceTimestamp() (*big.Int, error) {
	return _CuratedModule.Contract.GetResumeSinceTimestamp(&_CuratedModule.CallOpts)
}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_CuratedModule *CuratedModuleCallerSession) GetResumeSinceTimestamp() (*big.Int, error) {
	return _CuratedModule.Contract.GetResumeSinceTimestamp(&_CuratedModule.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CuratedModule *CuratedModuleCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CuratedModule *CuratedModuleSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CuratedModule.Contract.GetRoleAdmin(&_CuratedModule.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CuratedModule *CuratedModuleCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CuratedModule.Contract.GetRoleAdmin(&_CuratedModule.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_CuratedModule *CuratedModuleCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_CuratedModule *CuratedModuleSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _CuratedModule.Contract.GetRoleMember(&_CuratedModule.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_CuratedModule *CuratedModuleCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _CuratedModule.Contract.GetRoleMember(&_CuratedModule.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_CuratedModule *CuratedModuleCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_CuratedModule *CuratedModuleSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _CuratedModule.Contract.GetRoleMemberCount(&_CuratedModule.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_CuratedModule *CuratedModuleCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _CuratedModule.Contract.GetRoleMemberCount(&_CuratedModule.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CuratedModule *CuratedModuleCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CuratedModule *CuratedModuleSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _CuratedModule.Contract.GetRoleMembers(&_CuratedModule.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CuratedModule *CuratedModuleCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _CuratedModule.Contract.GetRoleMembers(&_CuratedModule.CallOpts, role)
}

// GetSigningKeys is a free data retrieval call binding the contract method 0x59e25c12.
//
// Solidity: function getSigningKeys(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(bytes keys)
func (_CuratedModule *CuratedModuleCaller) GetSigningKeys(opts *bind.CallOpts, nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) ([]byte, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getSigningKeys", nodeOperatorId, startIndex, keysCount)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetSigningKeys is a free data retrieval call binding the contract method 0x59e25c12.
//
// Solidity: function getSigningKeys(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(bytes keys)
func (_CuratedModule *CuratedModuleSession) GetSigningKeys(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) ([]byte, error) {
	return _CuratedModule.Contract.GetSigningKeys(&_CuratedModule.CallOpts, nodeOperatorId, startIndex, keysCount)
}

// GetSigningKeys is a free data retrieval call binding the contract method 0x59e25c12.
//
// Solidity: function getSigningKeys(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(bytes keys)
func (_CuratedModule *CuratedModuleCallerSession) GetSigningKeys(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) ([]byte, error) {
	return _CuratedModule.Contract.GetSigningKeys(&_CuratedModule.CallOpts, nodeOperatorId, startIndex, keysCount)
}

// GetSigningKeysWithSignatures is a free data retrieval call binding the contract method 0x50388cb6.
//
// Solidity: function getSigningKeysWithSignatures(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(bytes keys, bytes signatures)
func (_CuratedModule *CuratedModuleCaller) GetSigningKeysWithSignatures(opts *bind.CallOpts, nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) (struct {
	Keys       []byte
	Signatures []byte
}, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getSigningKeysWithSignatures", nodeOperatorId, startIndex, keysCount)

	outstruct := new(struct {
		Keys       []byte
		Signatures []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Keys = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Signatures = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetSigningKeysWithSignatures is a free data retrieval call binding the contract method 0x50388cb6.
//
// Solidity: function getSigningKeysWithSignatures(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(bytes keys, bytes signatures)
func (_CuratedModule *CuratedModuleSession) GetSigningKeysWithSignatures(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) (struct {
	Keys       []byte
	Signatures []byte
}, error) {
	return _CuratedModule.Contract.GetSigningKeysWithSignatures(&_CuratedModule.CallOpts, nodeOperatorId, startIndex, keysCount)
}

// GetSigningKeysWithSignatures is a free data retrieval call binding the contract method 0x50388cb6.
//
// Solidity: function getSigningKeysWithSignatures(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(bytes keys, bytes signatures)
func (_CuratedModule *CuratedModuleCallerSession) GetSigningKeysWithSignatures(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) (struct {
	Keys       []byte
	Signatures []byte
}, error) {
	return _CuratedModule.Contract.GetSigningKeysWithSignatures(&_CuratedModule.CallOpts, nodeOperatorId, startIndex, keysCount)
}

// GetStakingModuleSummary is a free data retrieval call binding the contract method 0x9abddf09.
//
// Solidity: function getStakingModuleSummary() view returns(uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_CuratedModule *CuratedModuleCaller) GetStakingModuleSummary(opts *bind.CallOpts) (struct {
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getStakingModuleSummary")

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
func (_CuratedModule *CuratedModuleSession) GetStakingModuleSummary() (struct {
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	return _CuratedModule.Contract.GetStakingModuleSummary(&_CuratedModule.CallOpts)
}

// GetStakingModuleSummary is a free data retrieval call binding the contract method 0x9abddf09.
//
// Solidity: function getStakingModuleSummary() view returns(uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_CuratedModule *CuratedModuleCallerSession) GetStakingModuleSummary() (struct {
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	return _CuratedModule.Contract.GetStakingModuleSummary(&_CuratedModule.CallOpts)
}

// GetTopUpAllocationTargets is a free data retrieval call binding the contract method 0x7fd63efb.
//
// Solidity: function getTopUpAllocationTargets() view returns(uint256[] currentAllocations, uint256[] targetAllocations)
func (_CuratedModule *CuratedModuleCaller) GetTopUpAllocationTargets(opts *bind.CallOpts) (struct {
	CurrentAllocations []*big.Int
	TargetAllocations  []*big.Int
}, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getTopUpAllocationTargets")

	outstruct := new(struct {
		CurrentAllocations []*big.Int
		TargetAllocations  []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentAllocations = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.TargetAllocations = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetTopUpAllocationTargets is a free data retrieval call binding the contract method 0x7fd63efb.
//
// Solidity: function getTopUpAllocationTargets() view returns(uint256[] currentAllocations, uint256[] targetAllocations)
func (_CuratedModule *CuratedModuleSession) GetTopUpAllocationTargets() (struct {
	CurrentAllocations []*big.Int
	TargetAllocations  []*big.Int
}, error) {
	return _CuratedModule.Contract.GetTopUpAllocationTargets(&_CuratedModule.CallOpts)
}

// GetTopUpAllocationTargets is a free data retrieval call binding the contract method 0x7fd63efb.
//
// Solidity: function getTopUpAllocationTargets() view returns(uint256[] currentAllocations, uint256[] targetAllocations)
func (_CuratedModule *CuratedModuleCallerSession) GetTopUpAllocationTargets() (struct {
	CurrentAllocations []*big.Int
	TargetAllocations  []*big.Int
}, error) {
	return _CuratedModule.Contract.GetTopUpAllocationTargets(&_CuratedModule.CallOpts)
}

// GetTotalModuleStake is a free data retrieval call binding the contract method 0x0c852f5c.
//
// Solidity: function getTotalModuleStake() view returns(uint256)
func (_CuratedModule *CuratedModuleCaller) GetTotalModuleStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getTotalModuleStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalModuleStake is a free data retrieval call binding the contract method 0x0c852f5c.
//
// Solidity: function getTotalModuleStake() view returns(uint256)
func (_CuratedModule *CuratedModuleSession) GetTotalModuleStake() (*big.Int, error) {
	return _CuratedModule.Contract.GetTotalModuleStake(&_CuratedModule.CallOpts)
}

// GetTotalModuleStake is a free data retrieval call binding the contract method 0x0c852f5c.
//
// Solidity: function getTotalModuleStake() view returns(uint256)
func (_CuratedModule *CuratedModuleCallerSession) GetTotalModuleStake() (*big.Int, error) {
	return _CuratedModule.Contract.GetTotalModuleStake(&_CuratedModule.CallOpts)
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(bytes32)
func (_CuratedModule *CuratedModuleCaller) GetType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "getType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(bytes32)
func (_CuratedModule *CuratedModuleSession) GetType() ([32]byte, error) {
	return _CuratedModule.Contract.GetType(&_CuratedModule.CallOpts)
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(bytes32)
func (_CuratedModule *CuratedModuleCallerSession) GetType() ([32]byte, error) {
	return _CuratedModule.Contract.GetType(&_CuratedModule.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CuratedModule *CuratedModuleCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CuratedModule *CuratedModuleSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CuratedModule.Contract.HasRole(&_CuratedModule.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CuratedModule *CuratedModuleCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CuratedModule.Contract.HasRole(&_CuratedModule.CallOpts, role, account)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_CuratedModule *CuratedModuleCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_CuratedModule *CuratedModuleSession) IsPaused() (bool, error) {
	return _CuratedModule.Contract.IsPaused(&_CuratedModule.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_CuratedModule *CuratedModuleCallerSession) IsPaused() (bool, error) {
	return _CuratedModule.Contract.IsPaused(&_CuratedModule.CallOpts)
}

// IsValidatorExitDelayPenaltyApplicable is a free data retrieval call binding the contract method 0x83b57a4e.
//
// Solidity: function isValidatorExitDelayPenaltyApplicable(uint256 nodeOperatorId, uint256 , bytes publicKey, uint256 eligibleToExitInSec) view returns(bool)
func (_CuratedModule *CuratedModuleCaller) IsValidatorExitDelayPenaltyApplicable(opts *bind.CallOpts, nodeOperatorId *big.Int, arg1 *big.Int, publicKey []byte, eligibleToExitInSec *big.Int) (bool, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "isValidatorExitDelayPenaltyApplicable", nodeOperatorId, arg1, publicKey, eligibleToExitInSec)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorExitDelayPenaltyApplicable is a free data retrieval call binding the contract method 0x83b57a4e.
//
// Solidity: function isValidatorExitDelayPenaltyApplicable(uint256 nodeOperatorId, uint256 , bytes publicKey, uint256 eligibleToExitInSec) view returns(bool)
func (_CuratedModule *CuratedModuleSession) IsValidatorExitDelayPenaltyApplicable(nodeOperatorId *big.Int, arg1 *big.Int, publicKey []byte, eligibleToExitInSec *big.Int) (bool, error) {
	return _CuratedModule.Contract.IsValidatorExitDelayPenaltyApplicable(&_CuratedModule.CallOpts, nodeOperatorId, arg1, publicKey, eligibleToExitInSec)
}

// IsValidatorExitDelayPenaltyApplicable is a free data retrieval call binding the contract method 0x83b57a4e.
//
// Solidity: function isValidatorExitDelayPenaltyApplicable(uint256 nodeOperatorId, uint256 , bytes publicKey, uint256 eligibleToExitInSec) view returns(bool)
func (_CuratedModule *CuratedModuleCallerSession) IsValidatorExitDelayPenaltyApplicable(nodeOperatorId *big.Int, arg1 *big.Int, publicKey []byte, eligibleToExitInSec *big.Int) (bool, error) {
	return _CuratedModule.Contract.IsValidatorExitDelayPenaltyApplicable(&_CuratedModule.CallOpts, nodeOperatorId, arg1, publicKey, eligibleToExitInSec)
}

// IsValidatorSlashed is a free data retrieval call binding the contract method 0x3dbe8b5a.
//
// Solidity: function isValidatorSlashed(uint256 nodeOperatorId, uint256 keyIndex) view returns(bool)
func (_CuratedModule *CuratedModuleCaller) IsValidatorSlashed(opts *bind.CallOpts, nodeOperatorId *big.Int, keyIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "isValidatorSlashed", nodeOperatorId, keyIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorSlashed is a free data retrieval call binding the contract method 0x3dbe8b5a.
//
// Solidity: function isValidatorSlashed(uint256 nodeOperatorId, uint256 keyIndex) view returns(bool)
func (_CuratedModule *CuratedModuleSession) IsValidatorSlashed(nodeOperatorId *big.Int, keyIndex *big.Int) (bool, error) {
	return _CuratedModule.Contract.IsValidatorSlashed(&_CuratedModule.CallOpts, nodeOperatorId, keyIndex)
}

// IsValidatorSlashed is a free data retrieval call binding the contract method 0x3dbe8b5a.
//
// Solidity: function isValidatorSlashed(uint256 nodeOperatorId, uint256 keyIndex) view returns(bool)
func (_CuratedModule *CuratedModuleCallerSession) IsValidatorSlashed(nodeOperatorId *big.Int, keyIndex *big.Int) (bool, error) {
	return _CuratedModule.Contract.IsValidatorSlashed(&_CuratedModule.CallOpts, nodeOperatorId, keyIndex)
}

// IsValidatorWithdrawn is a free data retrieval call binding the contract method 0x53433643.
//
// Solidity: function isValidatorWithdrawn(uint256 nodeOperatorId, uint256 keyIndex) view returns(bool)
func (_CuratedModule *CuratedModuleCaller) IsValidatorWithdrawn(opts *bind.CallOpts, nodeOperatorId *big.Int, keyIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "isValidatorWithdrawn", nodeOperatorId, keyIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorWithdrawn is a free data retrieval call binding the contract method 0x53433643.
//
// Solidity: function isValidatorWithdrawn(uint256 nodeOperatorId, uint256 keyIndex) view returns(bool)
func (_CuratedModule *CuratedModuleSession) IsValidatorWithdrawn(nodeOperatorId *big.Int, keyIndex *big.Int) (bool, error) {
	return _CuratedModule.Contract.IsValidatorWithdrawn(&_CuratedModule.CallOpts, nodeOperatorId, keyIndex)
}

// IsValidatorWithdrawn is a free data retrieval call binding the contract method 0x53433643.
//
// Solidity: function isValidatorWithdrawn(uint256 nodeOperatorId, uint256 keyIndex) view returns(bool)
func (_CuratedModule *CuratedModuleCallerSession) IsValidatorWithdrawn(nodeOperatorId *big.Int, keyIndex *big.Int) (bool, error) {
	return _CuratedModule.Contract.IsValidatorWithdrawn(&_CuratedModule.CallOpts, nodeOperatorId, keyIndex)
}

// OnExitedAndStuckValidatorsCountsUpdated is a free data retrieval call binding the contract method 0xe864299e.
//
// Solidity: function onExitedAndStuckValidatorsCountsUpdated() view returns()
func (_CuratedModule *CuratedModuleCaller) OnExitedAndStuckValidatorsCountsUpdated(opts *bind.CallOpts) error {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "onExitedAndStuckValidatorsCountsUpdated")

	if err != nil {
		return err
	}

	return err

}

// OnExitedAndStuckValidatorsCountsUpdated is a free data retrieval call binding the contract method 0xe864299e.
//
// Solidity: function onExitedAndStuckValidatorsCountsUpdated() view returns()
func (_CuratedModule *CuratedModuleSession) OnExitedAndStuckValidatorsCountsUpdated() error {
	return _CuratedModule.Contract.OnExitedAndStuckValidatorsCountsUpdated(&_CuratedModule.CallOpts)
}

// OnExitedAndStuckValidatorsCountsUpdated is a free data retrieval call binding the contract method 0xe864299e.
//
// Solidity: function onExitedAndStuckValidatorsCountsUpdated() view returns()
func (_CuratedModule *CuratedModuleCallerSession) OnExitedAndStuckValidatorsCountsUpdated() error {
	return _CuratedModule.Contract.OnExitedAndStuckValidatorsCountsUpdated(&_CuratedModule.CallOpts)
}

// OnWithdrawalCredentialsChanged is a free data retrieval call binding the contract method 0x90c09bdb.
//
// Solidity: function onWithdrawalCredentialsChanged() view returns()
func (_CuratedModule *CuratedModuleCaller) OnWithdrawalCredentialsChanged(opts *bind.CallOpts) error {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "onWithdrawalCredentialsChanged")

	if err != nil {
		return err
	}

	return err

}

// OnWithdrawalCredentialsChanged is a free data retrieval call binding the contract method 0x90c09bdb.
//
// Solidity: function onWithdrawalCredentialsChanged() view returns()
func (_CuratedModule *CuratedModuleSession) OnWithdrawalCredentialsChanged() error {
	return _CuratedModule.Contract.OnWithdrawalCredentialsChanged(&_CuratedModule.CallOpts)
}

// OnWithdrawalCredentialsChanged is a free data retrieval call binding the contract method 0x90c09bdb.
//
// Solidity: function onWithdrawalCredentialsChanged() view returns()
func (_CuratedModule *CuratedModuleCallerSession) OnWithdrawalCredentialsChanged() error {
	return _CuratedModule.Contract.OnWithdrawalCredentialsChanged(&_CuratedModule.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CuratedModule *CuratedModuleCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CuratedModule.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CuratedModule *CuratedModuleSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CuratedModule.Contract.SupportsInterface(&_CuratedModule.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CuratedModule *CuratedModuleCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CuratedModule.Contract.SupportsInterface(&_CuratedModule.CallOpts, interfaceId)
}

// AddValidatorKeysETH is a paid mutator transaction binding the contract method 0xa1913f4b.
//
// Solidity: function addValidatorKeysETH(address from, uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures) payable returns()
func (_CuratedModule *CuratedModuleTransactor) AddValidatorKeysETH(opts *bind.TransactOpts, from common.Address, nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "addValidatorKeysETH", from, nodeOperatorId, keysCount, publicKeys, signatures)
}

// AddValidatorKeysETH is a paid mutator transaction binding the contract method 0xa1913f4b.
//
// Solidity: function addValidatorKeysETH(address from, uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures) payable returns()
func (_CuratedModule *CuratedModuleSession) AddValidatorKeysETH(from common.Address, nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte) (*types.Transaction, error) {
	return _CuratedModule.Contract.AddValidatorKeysETH(&_CuratedModule.TransactOpts, from, nodeOperatorId, keysCount, publicKeys, signatures)
}

// AddValidatorKeysETH is a paid mutator transaction binding the contract method 0xa1913f4b.
//
// Solidity: function addValidatorKeysETH(address from, uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures) payable returns()
func (_CuratedModule *CuratedModuleTransactorSession) AddValidatorKeysETH(from common.Address, nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte) (*types.Transaction, error) {
	return _CuratedModule.Contract.AddValidatorKeysETH(&_CuratedModule.TransactOpts, from, nodeOperatorId, keysCount, publicKeys, signatures)
}

// AddValidatorKeysStETH is a paid mutator transaction binding the contract method 0xf696ccb3.
//
// Solidity: function addValidatorKeysStETH(address from, uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_CuratedModule *CuratedModuleTransactor) AddValidatorKeysStETH(opts *bind.TransactOpts, from common.Address, nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte, permit IAccountingPermitInput) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "addValidatorKeysStETH", from, nodeOperatorId, keysCount, publicKeys, signatures, permit)
}

// AddValidatorKeysStETH is a paid mutator transaction binding the contract method 0xf696ccb3.
//
// Solidity: function addValidatorKeysStETH(address from, uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_CuratedModule *CuratedModuleSession) AddValidatorKeysStETH(from common.Address, nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte, permit IAccountingPermitInput) (*types.Transaction, error) {
	return _CuratedModule.Contract.AddValidatorKeysStETH(&_CuratedModule.TransactOpts, from, nodeOperatorId, keysCount, publicKeys, signatures, permit)
}

// AddValidatorKeysStETH is a paid mutator transaction binding the contract method 0xf696ccb3.
//
// Solidity: function addValidatorKeysStETH(address from, uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_CuratedModule *CuratedModuleTransactorSession) AddValidatorKeysStETH(from common.Address, nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte, permit IAccountingPermitInput) (*types.Transaction, error) {
	return _CuratedModule.Contract.AddValidatorKeysStETH(&_CuratedModule.TransactOpts, from, nodeOperatorId, keysCount, publicKeys, signatures, permit)
}

// AddValidatorKeysWstETH is a paid mutator transaction binding the contract method 0xa6ab5b9c.
//
// Solidity: function addValidatorKeysWstETH(address from, uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_CuratedModule *CuratedModuleTransactor) AddValidatorKeysWstETH(opts *bind.TransactOpts, from common.Address, nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte, permit IAccountingPermitInput) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "addValidatorKeysWstETH", from, nodeOperatorId, keysCount, publicKeys, signatures, permit)
}

// AddValidatorKeysWstETH is a paid mutator transaction binding the contract method 0xa6ab5b9c.
//
// Solidity: function addValidatorKeysWstETH(address from, uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_CuratedModule *CuratedModuleSession) AddValidatorKeysWstETH(from common.Address, nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte, permit IAccountingPermitInput) (*types.Transaction, error) {
	return _CuratedModule.Contract.AddValidatorKeysWstETH(&_CuratedModule.TransactOpts, from, nodeOperatorId, keysCount, publicKeys, signatures, permit)
}

// AddValidatorKeysWstETH is a paid mutator transaction binding the contract method 0xa6ab5b9c.
//
// Solidity: function addValidatorKeysWstETH(address from, uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_CuratedModule *CuratedModuleTransactorSession) AddValidatorKeysWstETH(from common.Address, nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte, permit IAccountingPermitInput) (*types.Transaction, error) {
	return _CuratedModule.Contract.AddValidatorKeysWstETH(&_CuratedModule.TransactOpts, from, nodeOperatorId, keysCount, publicKeys, signatures, permit)
}

// AllocateDeposits is a paid mutator transaction binding the contract method 0x783b8a65.
//
// Solidity: function allocateDeposits(uint256 maxDepositAmount, bytes[] pubkeys, uint256[] keyIndices, uint256[] operatorIds, uint256[] topUpLimits) returns(uint256[] allocations)
func (_CuratedModule *CuratedModuleTransactor) AllocateDeposits(opts *bind.TransactOpts, maxDepositAmount *big.Int, pubkeys [][]byte, keyIndices []*big.Int, operatorIds []*big.Int, topUpLimits []*big.Int) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "allocateDeposits", maxDepositAmount, pubkeys, keyIndices, operatorIds, topUpLimits)
}

// AllocateDeposits is a paid mutator transaction binding the contract method 0x783b8a65.
//
// Solidity: function allocateDeposits(uint256 maxDepositAmount, bytes[] pubkeys, uint256[] keyIndices, uint256[] operatorIds, uint256[] topUpLimits) returns(uint256[] allocations)
func (_CuratedModule *CuratedModuleSession) AllocateDeposits(maxDepositAmount *big.Int, pubkeys [][]byte, keyIndices []*big.Int, operatorIds []*big.Int, topUpLimits []*big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.AllocateDeposits(&_CuratedModule.TransactOpts, maxDepositAmount, pubkeys, keyIndices, operatorIds, topUpLimits)
}

// AllocateDeposits is a paid mutator transaction binding the contract method 0x783b8a65.
//
// Solidity: function allocateDeposits(uint256 maxDepositAmount, bytes[] pubkeys, uint256[] keyIndices, uint256[] operatorIds, uint256[] topUpLimits) returns(uint256[] allocations)
func (_CuratedModule *CuratedModuleTransactorSession) AllocateDeposits(maxDepositAmount *big.Int, pubkeys [][]byte, keyIndices []*big.Int, operatorIds []*big.Int, topUpLimits []*big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.AllocateDeposits(&_CuratedModule.TransactOpts, maxDepositAmount, pubkeys, keyIndices, operatorIds, topUpLimits)
}

// BatchDepositInfoUpdate is a paid mutator transaction binding the contract method 0x398de960.
//
// Solidity: function batchDepositInfoUpdate(uint256 maxCount) returns(uint256 operatorsLeft)
func (_CuratedModule *CuratedModuleTransactor) BatchDepositInfoUpdate(opts *bind.TransactOpts, maxCount *big.Int) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "batchDepositInfoUpdate", maxCount)
}

// BatchDepositInfoUpdate is a paid mutator transaction binding the contract method 0x398de960.
//
// Solidity: function batchDepositInfoUpdate(uint256 maxCount) returns(uint256 operatorsLeft)
func (_CuratedModule *CuratedModuleSession) BatchDepositInfoUpdate(maxCount *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.BatchDepositInfoUpdate(&_CuratedModule.TransactOpts, maxCount)
}

// BatchDepositInfoUpdate is a paid mutator transaction binding the contract method 0x398de960.
//
// Solidity: function batchDepositInfoUpdate(uint256 maxCount) returns(uint256 operatorsLeft)
func (_CuratedModule *CuratedModuleTransactorSession) BatchDepositInfoUpdate(maxCount *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.BatchDepositInfoUpdate(&_CuratedModule.TransactOpts, maxCount)
}

// CancelGeneralDelayedPenalty is a paid mutator transaction binding the contract method 0x75e2d63f.
//
// Solidity: function cancelGeneralDelayedPenalty(uint256 nodeOperatorId, uint256 amount) returns()
func (_CuratedModule *CuratedModuleTransactor) CancelGeneralDelayedPenalty(opts *bind.TransactOpts, nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "cancelGeneralDelayedPenalty", nodeOperatorId, amount)
}

// CancelGeneralDelayedPenalty is a paid mutator transaction binding the contract method 0x75e2d63f.
//
// Solidity: function cancelGeneralDelayedPenalty(uint256 nodeOperatorId, uint256 amount) returns()
func (_CuratedModule *CuratedModuleSession) CancelGeneralDelayedPenalty(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.CancelGeneralDelayedPenalty(&_CuratedModule.TransactOpts, nodeOperatorId, amount)
}

// CancelGeneralDelayedPenalty is a paid mutator transaction binding the contract method 0x75e2d63f.
//
// Solidity: function cancelGeneralDelayedPenalty(uint256 nodeOperatorId, uint256 amount) returns()
func (_CuratedModule *CuratedModuleTransactorSession) CancelGeneralDelayedPenalty(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.CancelGeneralDelayedPenalty(&_CuratedModule.TransactOpts, nodeOperatorId, amount)
}

// ChangeNodeOperatorAddresses is a paid mutator transaction binding the contract method 0x5fab48c1.
//
// Solidity: function changeNodeOperatorAddresses(uint256 nodeOperatorId, address newManagerAddress, address newRewardAddress) returns()
func (_CuratedModule *CuratedModuleTransactor) ChangeNodeOperatorAddresses(opts *bind.TransactOpts, nodeOperatorId *big.Int, newManagerAddress common.Address, newRewardAddress common.Address) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "changeNodeOperatorAddresses", nodeOperatorId, newManagerAddress, newRewardAddress)
}

// ChangeNodeOperatorAddresses is a paid mutator transaction binding the contract method 0x5fab48c1.
//
// Solidity: function changeNodeOperatorAddresses(uint256 nodeOperatorId, address newManagerAddress, address newRewardAddress) returns()
func (_CuratedModule *CuratedModuleSession) ChangeNodeOperatorAddresses(nodeOperatorId *big.Int, newManagerAddress common.Address, newRewardAddress common.Address) (*types.Transaction, error) {
	return _CuratedModule.Contract.ChangeNodeOperatorAddresses(&_CuratedModule.TransactOpts, nodeOperatorId, newManagerAddress, newRewardAddress)
}

// ChangeNodeOperatorAddresses is a paid mutator transaction binding the contract method 0x5fab48c1.
//
// Solidity: function changeNodeOperatorAddresses(uint256 nodeOperatorId, address newManagerAddress, address newRewardAddress) returns()
func (_CuratedModule *CuratedModuleTransactorSession) ChangeNodeOperatorAddresses(nodeOperatorId *big.Int, newManagerAddress common.Address, newRewardAddress common.Address) (*types.Transaction, error) {
	return _CuratedModule.Contract.ChangeNodeOperatorAddresses(&_CuratedModule.TransactOpts, nodeOperatorId, newManagerAddress, newRewardAddress)
}

// ChangeNodeOperatorRewardAddress is a paid mutator transaction binding the contract method 0x75a401da.
//
// Solidity: function changeNodeOperatorRewardAddress(uint256 nodeOperatorId, address newAddress) returns()
func (_CuratedModule *CuratedModuleTransactor) ChangeNodeOperatorRewardAddress(opts *bind.TransactOpts, nodeOperatorId *big.Int, newAddress common.Address) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "changeNodeOperatorRewardAddress", nodeOperatorId, newAddress)
}

// ChangeNodeOperatorRewardAddress is a paid mutator transaction binding the contract method 0x75a401da.
//
// Solidity: function changeNodeOperatorRewardAddress(uint256 nodeOperatorId, address newAddress) returns()
func (_CuratedModule *CuratedModuleSession) ChangeNodeOperatorRewardAddress(nodeOperatorId *big.Int, newAddress common.Address) (*types.Transaction, error) {
	return _CuratedModule.Contract.ChangeNodeOperatorRewardAddress(&_CuratedModule.TransactOpts, nodeOperatorId, newAddress)
}

// ChangeNodeOperatorRewardAddress is a paid mutator transaction binding the contract method 0x75a401da.
//
// Solidity: function changeNodeOperatorRewardAddress(uint256 nodeOperatorId, address newAddress) returns()
func (_CuratedModule *CuratedModuleTransactorSession) ChangeNodeOperatorRewardAddress(nodeOperatorId *big.Int, newAddress common.Address) (*types.Transaction, error) {
	return _CuratedModule.Contract.ChangeNodeOperatorRewardAddress(&_CuratedModule.TransactOpts, nodeOperatorId, newAddress)
}

// CompensateGeneralDelayedPenalty is a paid mutator transaction binding the contract method 0x71d62617.
//
// Solidity: function compensateGeneralDelayedPenalty(uint256 nodeOperatorId) returns()
func (_CuratedModule *CuratedModuleTransactor) CompensateGeneralDelayedPenalty(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "compensateGeneralDelayedPenalty", nodeOperatorId)
}

// CompensateGeneralDelayedPenalty is a paid mutator transaction binding the contract method 0x71d62617.
//
// Solidity: function compensateGeneralDelayedPenalty(uint256 nodeOperatorId) returns()
func (_CuratedModule *CuratedModuleSession) CompensateGeneralDelayedPenalty(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.CompensateGeneralDelayedPenalty(&_CuratedModule.TransactOpts, nodeOperatorId)
}

// CompensateGeneralDelayedPenalty is a paid mutator transaction binding the contract method 0x71d62617.
//
// Solidity: function compensateGeneralDelayedPenalty(uint256 nodeOperatorId) returns()
func (_CuratedModule *CuratedModuleTransactorSession) CompensateGeneralDelayedPenalty(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.CompensateGeneralDelayedPenalty(&_CuratedModule.TransactOpts, nodeOperatorId)
}

// ConfirmNodeOperatorManagerAddressChange is a paid mutator transaction binding the contract method 0x6bb1bfdf.
//
// Solidity: function confirmNodeOperatorManagerAddressChange(uint256 nodeOperatorId) returns()
func (_CuratedModule *CuratedModuleTransactor) ConfirmNodeOperatorManagerAddressChange(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "confirmNodeOperatorManagerAddressChange", nodeOperatorId)
}

// ConfirmNodeOperatorManagerAddressChange is a paid mutator transaction binding the contract method 0x6bb1bfdf.
//
// Solidity: function confirmNodeOperatorManagerAddressChange(uint256 nodeOperatorId) returns()
func (_CuratedModule *CuratedModuleSession) ConfirmNodeOperatorManagerAddressChange(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.ConfirmNodeOperatorManagerAddressChange(&_CuratedModule.TransactOpts, nodeOperatorId)
}

// ConfirmNodeOperatorManagerAddressChange is a paid mutator transaction binding the contract method 0x6bb1bfdf.
//
// Solidity: function confirmNodeOperatorManagerAddressChange(uint256 nodeOperatorId) returns()
func (_CuratedModule *CuratedModuleTransactorSession) ConfirmNodeOperatorManagerAddressChange(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.ConfirmNodeOperatorManagerAddressChange(&_CuratedModule.TransactOpts, nodeOperatorId)
}

// ConfirmNodeOperatorRewardAddressChange is a paid mutator transaction binding the contract method 0x5204281c.
//
// Solidity: function confirmNodeOperatorRewardAddressChange(uint256 nodeOperatorId) returns()
func (_CuratedModule *CuratedModuleTransactor) ConfirmNodeOperatorRewardAddressChange(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "confirmNodeOperatorRewardAddressChange", nodeOperatorId)
}

// ConfirmNodeOperatorRewardAddressChange is a paid mutator transaction binding the contract method 0x5204281c.
//
// Solidity: function confirmNodeOperatorRewardAddressChange(uint256 nodeOperatorId) returns()
func (_CuratedModule *CuratedModuleSession) ConfirmNodeOperatorRewardAddressChange(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.ConfirmNodeOperatorRewardAddressChange(&_CuratedModule.TransactOpts, nodeOperatorId)
}

// ConfirmNodeOperatorRewardAddressChange is a paid mutator transaction binding the contract method 0x5204281c.
//
// Solidity: function confirmNodeOperatorRewardAddressChange(uint256 nodeOperatorId) returns()
func (_CuratedModule *CuratedModuleTransactorSession) ConfirmNodeOperatorRewardAddressChange(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.ConfirmNodeOperatorRewardAddressChange(&_CuratedModule.TransactOpts, nodeOperatorId)
}

// CreateNodeOperator is a paid mutator transaction binding the contract method 0xa4516c98.
//
// Solidity: function createNodeOperator(address from, (address,address,bool) managementProperties, address ) returns(uint256 nodeOperatorId)
func (_CuratedModule *CuratedModuleTransactor) CreateNodeOperator(opts *bind.TransactOpts, from common.Address, managementProperties NodeOperatorManagementProperties, arg2 common.Address) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "createNodeOperator", from, managementProperties, arg2)
}

// CreateNodeOperator is a paid mutator transaction binding the contract method 0xa4516c98.
//
// Solidity: function createNodeOperator(address from, (address,address,bool) managementProperties, address ) returns(uint256 nodeOperatorId)
func (_CuratedModule *CuratedModuleSession) CreateNodeOperator(from common.Address, managementProperties NodeOperatorManagementProperties, arg2 common.Address) (*types.Transaction, error) {
	return _CuratedModule.Contract.CreateNodeOperator(&_CuratedModule.TransactOpts, from, managementProperties, arg2)
}

// CreateNodeOperator is a paid mutator transaction binding the contract method 0xa4516c98.
//
// Solidity: function createNodeOperator(address from, (address,address,bool) managementProperties, address ) returns(uint256 nodeOperatorId)
func (_CuratedModule *CuratedModuleTransactorSession) CreateNodeOperator(from common.Address, managementProperties NodeOperatorManagementProperties, arg2 common.Address) (*types.Transaction, error) {
	return _CuratedModule.Contract.CreateNodeOperator(&_CuratedModule.TransactOpts, from, managementProperties, arg2)
}

// DecreaseVettedSigningKeysCount is a paid mutator transaction binding the contract method 0xb643189b.
//
// Solidity: function decreaseVettedSigningKeysCount(bytes nodeOperatorIds, bytes vettedSigningKeysCounts) returns()
func (_CuratedModule *CuratedModuleTransactor) DecreaseVettedSigningKeysCount(opts *bind.TransactOpts, nodeOperatorIds []byte, vettedSigningKeysCounts []byte) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "decreaseVettedSigningKeysCount", nodeOperatorIds, vettedSigningKeysCounts)
}

// DecreaseVettedSigningKeysCount is a paid mutator transaction binding the contract method 0xb643189b.
//
// Solidity: function decreaseVettedSigningKeysCount(bytes nodeOperatorIds, bytes vettedSigningKeysCounts) returns()
func (_CuratedModule *CuratedModuleSession) DecreaseVettedSigningKeysCount(nodeOperatorIds []byte, vettedSigningKeysCounts []byte) (*types.Transaction, error) {
	return _CuratedModule.Contract.DecreaseVettedSigningKeysCount(&_CuratedModule.TransactOpts, nodeOperatorIds, vettedSigningKeysCounts)
}

// DecreaseVettedSigningKeysCount is a paid mutator transaction binding the contract method 0xb643189b.
//
// Solidity: function decreaseVettedSigningKeysCount(bytes nodeOperatorIds, bytes vettedSigningKeysCounts) returns()
func (_CuratedModule *CuratedModuleTransactorSession) DecreaseVettedSigningKeysCount(nodeOperatorIds []byte, vettedSigningKeysCounts []byte) (*types.Transaction, error) {
	return _CuratedModule.Contract.DecreaseVettedSigningKeysCount(&_CuratedModule.TransactOpts, nodeOperatorIds, vettedSigningKeysCounts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CuratedModule *CuratedModuleTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CuratedModule *CuratedModuleSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CuratedModule.Contract.GrantRole(&_CuratedModule.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CuratedModule *CuratedModuleTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CuratedModule.Contract.GrantRole(&_CuratedModule.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_CuratedModule *CuratedModuleTransactor) Initialize(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "initialize", admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_CuratedModule *CuratedModuleSession) Initialize(admin common.Address) (*types.Transaction, error) {
	return _CuratedModule.Contract.Initialize(&_CuratedModule.TransactOpts, admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_CuratedModule *CuratedModuleTransactorSession) Initialize(admin common.Address) (*types.Transaction, error) {
	return _CuratedModule.Contract.Initialize(&_CuratedModule.TransactOpts, admin)
}

// NotifyNodeOperatorWeightChange is a paid mutator transaction binding the contract method 0x20a55b48.
//
// Solidity: function notifyNodeOperatorWeightChange(uint256 nodeOperatorId, uint256 oldWeight, uint256 newWeight) returns()
func (_CuratedModule *CuratedModuleTransactor) NotifyNodeOperatorWeightChange(opts *bind.TransactOpts, nodeOperatorId *big.Int, oldWeight *big.Int, newWeight *big.Int) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "notifyNodeOperatorWeightChange", nodeOperatorId, oldWeight, newWeight)
}

// NotifyNodeOperatorWeightChange is a paid mutator transaction binding the contract method 0x20a55b48.
//
// Solidity: function notifyNodeOperatorWeightChange(uint256 nodeOperatorId, uint256 oldWeight, uint256 newWeight) returns()
func (_CuratedModule *CuratedModuleSession) NotifyNodeOperatorWeightChange(nodeOperatorId *big.Int, oldWeight *big.Int, newWeight *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.NotifyNodeOperatorWeightChange(&_CuratedModule.TransactOpts, nodeOperatorId, oldWeight, newWeight)
}

// NotifyNodeOperatorWeightChange is a paid mutator transaction binding the contract method 0x20a55b48.
//
// Solidity: function notifyNodeOperatorWeightChange(uint256 nodeOperatorId, uint256 oldWeight, uint256 newWeight) returns()
func (_CuratedModule *CuratedModuleTransactorSession) NotifyNodeOperatorWeightChange(nodeOperatorId *big.Int, oldWeight *big.Int, newWeight *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.NotifyNodeOperatorWeightChange(&_CuratedModule.TransactOpts, nodeOperatorId, oldWeight, newWeight)
}

// ObtainDepositData is a paid mutator transaction binding the contract method 0xbee41b58.
//
// Solidity: function obtainDepositData(uint256 depositsCount, bytes ) returns(bytes publicKeys, bytes signatures)
func (_CuratedModule *CuratedModuleTransactor) ObtainDepositData(opts *bind.TransactOpts, depositsCount *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "obtainDepositData", depositsCount, arg1)
}

// ObtainDepositData is a paid mutator transaction binding the contract method 0xbee41b58.
//
// Solidity: function obtainDepositData(uint256 depositsCount, bytes ) returns(bytes publicKeys, bytes signatures)
func (_CuratedModule *CuratedModuleSession) ObtainDepositData(depositsCount *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _CuratedModule.Contract.ObtainDepositData(&_CuratedModule.TransactOpts, depositsCount, arg1)
}

// ObtainDepositData is a paid mutator transaction binding the contract method 0xbee41b58.
//
// Solidity: function obtainDepositData(uint256 depositsCount, bytes ) returns(bytes publicKeys, bytes signatures)
func (_CuratedModule *CuratedModuleTransactorSession) ObtainDepositData(depositsCount *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _CuratedModule.Contract.ObtainDepositData(&_CuratedModule.TransactOpts, depositsCount, arg1)
}

// OnRewardsMinted is a paid mutator transaction binding the contract method 0x8d7e4017.
//
// Solidity: function onRewardsMinted(uint256 totalShares) returns()
func (_CuratedModule *CuratedModuleTransactor) OnRewardsMinted(opts *bind.TransactOpts, totalShares *big.Int) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "onRewardsMinted", totalShares)
}

// OnRewardsMinted is a paid mutator transaction binding the contract method 0x8d7e4017.
//
// Solidity: function onRewardsMinted(uint256 totalShares) returns()
func (_CuratedModule *CuratedModuleSession) OnRewardsMinted(totalShares *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.OnRewardsMinted(&_CuratedModule.TransactOpts, totalShares)
}

// OnRewardsMinted is a paid mutator transaction binding the contract method 0x8d7e4017.
//
// Solidity: function onRewardsMinted(uint256 totalShares) returns()
func (_CuratedModule *CuratedModuleTransactorSession) OnRewardsMinted(totalShares *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.OnRewardsMinted(&_CuratedModule.TransactOpts, totalShares)
}

// OnValidatorExitTriggered is a paid mutator transaction binding the contract method 0x693cc600.
//
// Solidity: function onValidatorExitTriggered(uint256 nodeOperatorId, bytes publicKey, uint256 elWithdrawalRequestFeePaid, uint256 exitType) returns()
func (_CuratedModule *CuratedModuleTransactor) OnValidatorExitTriggered(opts *bind.TransactOpts, nodeOperatorId *big.Int, publicKey []byte, elWithdrawalRequestFeePaid *big.Int, exitType *big.Int) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "onValidatorExitTriggered", nodeOperatorId, publicKey, elWithdrawalRequestFeePaid, exitType)
}

// OnValidatorExitTriggered is a paid mutator transaction binding the contract method 0x693cc600.
//
// Solidity: function onValidatorExitTriggered(uint256 nodeOperatorId, bytes publicKey, uint256 elWithdrawalRequestFeePaid, uint256 exitType) returns()
func (_CuratedModule *CuratedModuleSession) OnValidatorExitTriggered(nodeOperatorId *big.Int, publicKey []byte, elWithdrawalRequestFeePaid *big.Int, exitType *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.OnValidatorExitTriggered(&_CuratedModule.TransactOpts, nodeOperatorId, publicKey, elWithdrawalRequestFeePaid, exitType)
}

// OnValidatorExitTriggered is a paid mutator transaction binding the contract method 0x693cc600.
//
// Solidity: function onValidatorExitTriggered(uint256 nodeOperatorId, bytes publicKey, uint256 elWithdrawalRequestFeePaid, uint256 exitType) returns()
func (_CuratedModule *CuratedModuleTransactorSession) OnValidatorExitTriggered(nodeOperatorId *big.Int, publicKey []byte, elWithdrawalRequestFeePaid *big.Int, exitType *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.OnValidatorExitTriggered(&_CuratedModule.TransactOpts, nodeOperatorId, publicKey, elWithdrawalRequestFeePaid, exitType)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 duration) returns()
func (_CuratedModule *CuratedModuleTransactor) PauseFor(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "pauseFor", duration)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 duration) returns()
func (_CuratedModule *CuratedModuleSession) PauseFor(duration *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.PauseFor(&_CuratedModule.TransactOpts, duration)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 duration) returns()
func (_CuratedModule *CuratedModuleTransactorSession) PauseFor(duration *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.PauseFor(&_CuratedModule.TransactOpts, duration)
}

// ProposeNodeOperatorManagerAddressChange is a paid mutator transaction binding the contract method 0x8cabe959.
//
// Solidity: function proposeNodeOperatorManagerAddressChange(uint256 nodeOperatorId, address proposedAddress) returns()
func (_CuratedModule *CuratedModuleTransactor) ProposeNodeOperatorManagerAddressChange(opts *bind.TransactOpts, nodeOperatorId *big.Int, proposedAddress common.Address) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "proposeNodeOperatorManagerAddressChange", nodeOperatorId, proposedAddress)
}

// ProposeNodeOperatorManagerAddressChange is a paid mutator transaction binding the contract method 0x8cabe959.
//
// Solidity: function proposeNodeOperatorManagerAddressChange(uint256 nodeOperatorId, address proposedAddress) returns()
func (_CuratedModule *CuratedModuleSession) ProposeNodeOperatorManagerAddressChange(nodeOperatorId *big.Int, proposedAddress common.Address) (*types.Transaction, error) {
	return _CuratedModule.Contract.ProposeNodeOperatorManagerAddressChange(&_CuratedModule.TransactOpts, nodeOperatorId, proposedAddress)
}

// ProposeNodeOperatorManagerAddressChange is a paid mutator transaction binding the contract method 0x8cabe959.
//
// Solidity: function proposeNodeOperatorManagerAddressChange(uint256 nodeOperatorId, address proposedAddress) returns()
func (_CuratedModule *CuratedModuleTransactorSession) ProposeNodeOperatorManagerAddressChange(nodeOperatorId *big.Int, proposedAddress common.Address) (*types.Transaction, error) {
	return _CuratedModule.Contract.ProposeNodeOperatorManagerAddressChange(&_CuratedModule.TransactOpts, nodeOperatorId, proposedAddress)
}

// ProposeNodeOperatorRewardAddressChange is a paid mutator transaction binding the contract method 0x1b40b231.
//
// Solidity: function proposeNodeOperatorRewardAddressChange(uint256 nodeOperatorId, address proposedAddress) returns()
func (_CuratedModule *CuratedModuleTransactor) ProposeNodeOperatorRewardAddressChange(opts *bind.TransactOpts, nodeOperatorId *big.Int, proposedAddress common.Address) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "proposeNodeOperatorRewardAddressChange", nodeOperatorId, proposedAddress)
}

// ProposeNodeOperatorRewardAddressChange is a paid mutator transaction binding the contract method 0x1b40b231.
//
// Solidity: function proposeNodeOperatorRewardAddressChange(uint256 nodeOperatorId, address proposedAddress) returns()
func (_CuratedModule *CuratedModuleSession) ProposeNodeOperatorRewardAddressChange(nodeOperatorId *big.Int, proposedAddress common.Address) (*types.Transaction, error) {
	return _CuratedModule.Contract.ProposeNodeOperatorRewardAddressChange(&_CuratedModule.TransactOpts, nodeOperatorId, proposedAddress)
}

// ProposeNodeOperatorRewardAddressChange is a paid mutator transaction binding the contract method 0x1b40b231.
//
// Solidity: function proposeNodeOperatorRewardAddressChange(uint256 nodeOperatorId, address proposedAddress) returns()
func (_CuratedModule *CuratedModuleTransactorSession) ProposeNodeOperatorRewardAddressChange(nodeOperatorId *big.Int, proposedAddress common.Address) (*types.Transaction, error) {
	return _CuratedModule.Contract.ProposeNodeOperatorRewardAddressChange(&_CuratedModule.TransactOpts, nodeOperatorId, proposedAddress)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_CuratedModule *CuratedModuleTransactor) RecoverERC1155(opts *bind.TransactOpts, token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "recoverERC1155", token, tokenId)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_CuratedModule *CuratedModuleSession) RecoverERC1155(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.RecoverERC1155(&_CuratedModule.TransactOpts, token, tokenId)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_CuratedModule *CuratedModuleTransactorSession) RecoverERC1155(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.RecoverERC1155(&_CuratedModule.TransactOpts, token, tokenId)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_CuratedModule *CuratedModuleTransactor) RecoverERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "recoverERC20", token, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_CuratedModule *CuratedModuleSession) RecoverERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.RecoverERC20(&_CuratedModule.TransactOpts, token, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_CuratedModule *CuratedModuleTransactorSession) RecoverERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.RecoverERC20(&_CuratedModule.TransactOpts, token, amount)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_CuratedModule *CuratedModuleTransactor) RecoverERC721(opts *bind.TransactOpts, token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "recoverERC721", token, tokenId)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_CuratedModule *CuratedModuleSession) RecoverERC721(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.RecoverERC721(&_CuratedModule.TransactOpts, token, tokenId)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_CuratedModule *CuratedModuleTransactorSession) RecoverERC721(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.RecoverERC721(&_CuratedModule.TransactOpts, token, tokenId)
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_CuratedModule *CuratedModuleTransactor) RecoverEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "recoverEther")
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_CuratedModule *CuratedModuleSession) RecoverEther() (*types.Transaction, error) {
	return _CuratedModule.Contract.RecoverEther(&_CuratedModule.TransactOpts)
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_CuratedModule *CuratedModuleTransactorSession) RecoverEther() (*types.Transaction, error) {
	return _CuratedModule.Contract.RecoverEther(&_CuratedModule.TransactOpts)
}

// RemoveKeys is a paid mutator transaction binding the contract method 0x8b3ac71d.
//
// Solidity: function removeKeys(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) returns()
func (_CuratedModule *CuratedModuleTransactor) RemoveKeys(opts *bind.TransactOpts, nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "removeKeys", nodeOperatorId, startIndex, keysCount)
}

// RemoveKeys is a paid mutator transaction binding the contract method 0x8b3ac71d.
//
// Solidity: function removeKeys(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) returns()
func (_CuratedModule *CuratedModuleSession) RemoveKeys(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.RemoveKeys(&_CuratedModule.TransactOpts, nodeOperatorId, startIndex, keysCount)
}

// RemoveKeys is a paid mutator transaction binding the contract method 0x8b3ac71d.
//
// Solidity: function removeKeys(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) returns()
func (_CuratedModule *CuratedModuleTransactorSession) RemoveKeys(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.RemoveKeys(&_CuratedModule.TransactOpts, nodeOperatorId, startIndex, keysCount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CuratedModule *CuratedModuleTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CuratedModule *CuratedModuleSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CuratedModule.Contract.RenounceRole(&_CuratedModule.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CuratedModule *CuratedModuleTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CuratedModule.Contract.RenounceRole(&_CuratedModule.TransactOpts, role, callerConfirmation)
}

// ReportGeneralDelayedPenalty is a paid mutator transaction binding the contract method 0x3368fccd.
//
// Solidity: function reportGeneralDelayedPenalty(uint256 nodeOperatorId, bytes32 penaltyType, uint256 amount, string details) returns()
func (_CuratedModule *CuratedModuleTransactor) ReportGeneralDelayedPenalty(opts *bind.TransactOpts, nodeOperatorId *big.Int, penaltyType [32]byte, amount *big.Int, details string) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "reportGeneralDelayedPenalty", nodeOperatorId, penaltyType, amount, details)
}

// ReportGeneralDelayedPenalty is a paid mutator transaction binding the contract method 0x3368fccd.
//
// Solidity: function reportGeneralDelayedPenalty(uint256 nodeOperatorId, bytes32 penaltyType, uint256 amount, string details) returns()
func (_CuratedModule *CuratedModuleSession) ReportGeneralDelayedPenalty(nodeOperatorId *big.Int, penaltyType [32]byte, amount *big.Int, details string) (*types.Transaction, error) {
	return _CuratedModule.Contract.ReportGeneralDelayedPenalty(&_CuratedModule.TransactOpts, nodeOperatorId, penaltyType, amount, details)
}

// ReportGeneralDelayedPenalty is a paid mutator transaction binding the contract method 0x3368fccd.
//
// Solidity: function reportGeneralDelayedPenalty(uint256 nodeOperatorId, bytes32 penaltyType, uint256 amount, string details) returns()
func (_CuratedModule *CuratedModuleTransactorSession) ReportGeneralDelayedPenalty(nodeOperatorId *big.Int, penaltyType [32]byte, amount *big.Int, details string) (*types.Transaction, error) {
	return _CuratedModule.Contract.ReportGeneralDelayedPenalty(&_CuratedModule.TransactOpts, nodeOperatorId, penaltyType, amount, details)
}

// ReportRegularWithdrawnValidators is a paid mutator transaction binding the contract method 0x81ecc658.
//
// Solidity: function reportRegularWithdrawnValidators((uint256,uint256,uint256,uint256,bool)[] validatorInfos) returns()
func (_CuratedModule *CuratedModuleTransactor) ReportRegularWithdrawnValidators(opts *bind.TransactOpts, validatorInfos []WithdrawnValidatorInfo) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "reportRegularWithdrawnValidators", validatorInfos)
}

// ReportRegularWithdrawnValidators is a paid mutator transaction binding the contract method 0x81ecc658.
//
// Solidity: function reportRegularWithdrawnValidators((uint256,uint256,uint256,uint256,bool)[] validatorInfos) returns()
func (_CuratedModule *CuratedModuleSession) ReportRegularWithdrawnValidators(validatorInfos []WithdrawnValidatorInfo) (*types.Transaction, error) {
	return _CuratedModule.Contract.ReportRegularWithdrawnValidators(&_CuratedModule.TransactOpts, validatorInfos)
}

// ReportRegularWithdrawnValidators is a paid mutator transaction binding the contract method 0x81ecc658.
//
// Solidity: function reportRegularWithdrawnValidators((uint256,uint256,uint256,uint256,bool)[] validatorInfos) returns()
func (_CuratedModule *CuratedModuleTransactorSession) ReportRegularWithdrawnValidators(validatorInfos []WithdrawnValidatorInfo) (*types.Transaction, error) {
	return _CuratedModule.Contract.ReportRegularWithdrawnValidators(&_CuratedModule.TransactOpts, validatorInfos)
}

// ReportSlashedWithdrawnValidators is a paid mutator transaction binding the contract method 0x4412f7aa.
//
// Solidity: function reportSlashedWithdrawnValidators((uint256,uint256,uint256,uint256,bool)[] validatorInfos) returns()
func (_CuratedModule *CuratedModuleTransactor) ReportSlashedWithdrawnValidators(opts *bind.TransactOpts, validatorInfos []WithdrawnValidatorInfo) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "reportSlashedWithdrawnValidators", validatorInfos)
}

// ReportSlashedWithdrawnValidators is a paid mutator transaction binding the contract method 0x4412f7aa.
//
// Solidity: function reportSlashedWithdrawnValidators((uint256,uint256,uint256,uint256,bool)[] validatorInfos) returns()
func (_CuratedModule *CuratedModuleSession) ReportSlashedWithdrawnValidators(validatorInfos []WithdrawnValidatorInfo) (*types.Transaction, error) {
	return _CuratedModule.Contract.ReportSlashedWithdrawnValidators(&_CuratedModule.TransactOpts, validatorInfos)
}

// ReportSlashedWithdrawnValidators is a paid mutator transaction binding the contract method 0x4412f7aa.
//
// Solidity: function reportSlashedWithdrawnValidators((uint256,uint256,uint256,uint256,bool)[] validatorInfos) returns()
func (_CuratedModule *CuratedModuleTransactorSession) ReportSlashedWithdrawnValidators(validatorInfos []WithdrawnValidatorInfo) (*types.Transaction, error) {
	return _CuratedModule.Contract.ReportSlashedWithdrawnValidators(&_CuratedModule.TransactOpts, validatorInfos)
}

// ReportValidatorBalance is a paid mutator transaction binding the contract method 0x69629163.
//
// Solidity: function reportValidatorBalance(uint256 nodeOperatorId, uint256 keyIndex, uint256 currentBalanceWei) returns()
func (_CuratedModule *CuratedModuleTransactor) ReportValidatorBalance(opts *bind.TransactOpts, nodeOperatorId *big.Int, keyIndex *big.Int, currentBalanceWei *big.Int) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "reportValidatorBalance", nodeOperatorId, keyIndex, currentBalanceWei)
}

// ReportValidatorBalance is a paid mutator transaction binding the contract method 0x69629163.
//
// Solidity: function reportValidatorBalance(uint256 nodeOperatorId, uint256 keyIndex, uint256 currentBalanceWei) returns()
func (_CuratedModule *CuratedModuleSession) ReportValidatorBalance(nodeOperatorId *big.Int, keyIndex *big.Int, currentBalanceWei *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.ReportValidatorBalance(&_CuratedModule.TransactOpts, nodeOperatorId, keyIndex, currentBalanceWei)
}

// ReportValidatorBalance is a paid mutator transaction binding the contract method 0x69629163.
//
// Solidity: function reportValidatorBalance(uint256 nodeOperatorId, uint256 keyIndex, uint256 currentBalanceWei) returns()
func (_CuratedModule *CuratedModuleTransactorSession) ReportValidatorBalance(nodeOperatorId *big.Int, keyIndex *big.Int, currentBalanceWei *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.ReportValidatorBalance(&_CuratedModule.TransactOpts, nodeOperatorId, keyIndex, currentBalanceWei)
}

// ReportValidatorExitDelay is a paid mutator transaction binding the contract method 0x57f9c341.
//
// Solidity: function reportValidatorExitDelay(uint256 nodeOperatorId, uint256 , bytes publicKey, uint256 eligibleToExitInSec) returns()
func (_CuratedModule *CuratedModuleTransactor) ReportValidatorExitDelay(opts *bind.TransactOpts, nodeOperatorId *big.Int, arg1 *big.Int, publicKey []byte, eligibleToExitInSec *big.Int) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "reportValidatorExitDelay", nodeOperatorId, arg1, publicKey, eligibleToExitInSec)
}

// ReportValidatorExitDelay is a paid mutator transaction binding the contract method 0x57f9c341.
//
// Solidity: function reportValidatorExitDelay(uint256 nodeOperatorId, uint256 , bytes publicKey, uint256 eligibleToExitInSec) returns()
func (_CuratedModule *CuratedModuleSession) ReportValidatorExitDelay(nodeOperatorId *big.Int, arg1 *big.Int, publicKey []byte, eligibleToExitInSec *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.ReportValidatorExitDelay(&_CuratedModule.TransactOpts, nodeOperatorId, arg1, publicKey, eligibleToExitInSec)
}

// ReportValidatorExitDelay is a paid mutator transaction binding the contract method 0x57f9c341.
//
// Solidity: function reportValidatorExitDelay(uint256 nodeOperatorId, uint256 , bytes publicKey, uint256 eligibleToExitInSec) returns()
func (_CuratedModule *CuratedModuleTransactorSession) ReportValidatorExitDelay(nodeOperatorId *big.Int, arg1 *big.Int, publicKey []byte, eligibleToExitInSec *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.ReportValidatorExitDelay(&_CuratedModule.TransactOpts, nodeOperatorId, arg1, publicKey, eligibleToExitInSec)
}

// ReportValidatorSlashing is a paid mutator transaction binding the contract method 0x228b2e63.
//
// Solidity: function reportValidatorSlashing(uint256 nodeOperatorId, uint256 keyIndex) returns()
func (_CuratedModule *CuratedModuleTransactor) ReportValidatorSlashing(opts *bind.TransactOpts, nodeOperatorId *big.Int, keyIndex *big.Int) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "reportValidatorSlashing", nodeOperatorId, keyIndex)
}

// ReportValidatorSlashing is a paid mutator transaction binding the contract method 0x228b2e63.
//
// Solidity: function reportValidatorSlashing(uint256 nodeOperatorId, uint256 keyIndex) returns()
func (_CuratedModule *CuratedModuleSession) ReportValidatorSlashing(nodeOperatorId *big.Int, keyIndex *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.ReportValidatorSlashing(&_CuratedModule.TransactOpts, nodeOperatorId, keyIndex)
}

// ReportValidatorSlashing is a paid mutator transaction binding the contract method 0x228b2e63.
//
// Solidity: function reportValidatorSlashing(uint256 nodeOperatorId, uint256 keyIndex) returns()
func (_CuratedModule *CuratedModuleTransactorSession) ReportValidatorSlashing(nodeOperatorId *big.Int, keyIndex *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.ReportValidatorSlashing(&_CuratedModule.TransactOpts, nodeOperatorId, keyIndex)
}

// RequestFullDepositInfoUpdate is a paid mutator transaction binding the contract method 0x24bae21c.
//
// Solidity: function requestFullDepositInfoUpdate() returns()
func (_CuratedModule *CuratedModuleTransactor) RequestFullDepositInfoUpdate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "requestFullDepositInfoUpdate")
}

// RequestFullDepositInfoUpdate is a paid mutator transaction binding the contract method 0x24bae21c.
//
// Solidity: function requestFullDepositInfoUpdate() returns()
func (_CuratedModule *CuratedModuleSession) RequestFullDepositInfoUpdate() (*types.Transaction, error) {
	return _CuratedModule.Contract.RequestFullDepositInfoUpdate(&_CuratedModule.TransactOpts)
}

// RequestFullDepositInfoUpdate is a paid mutator transaction binding the contract method 0x24bae21c.
//
// Solidity: function requestFullDepositInfoUpdate() returns()
func (_CuratedModule *CuratedModuleTransactorSession) RequestFullDepositInfoUpdate() (*types.Transaction, error) {
	return _CuratedModule.Contract.RequestFullDepositInfoUpdate(&_CuratedModule.TransactOpts)
}

// ResetNodeOperatorManagerAddress is a paid mutator transaction binding the contract method 0x6a6304cc.
//
// Solidity: function resetNodeOperatorManagerAddress(uint256 nodeOperatorId) returns()
func (_CuratedModule *CuratedModuleTransactor) ResetNodeOperatorManagerAddress(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "resetNodeOperatorManagerAddress", nodeOperatorId)
}

// ResetNodeOperatorManagerAddress is a paid mutator transaction binding the contract method 0x6a6304cc.
//
// Solidity: function resetNodeOperatorManagerAddress(uint256 nodeOperatorId) returns()
func (_CuratedModule *CuratedModuleSession) ResetNodeOperatorManagerAddress(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.ResetNodeOperatorManagerAddress(&_CuratedModule.TransactOpts, nodeOperatorId)
}

// ResetNodeOperatorManagerAddress is a paid mutator transaction binding the contract method 0x6a6304cc.
//
// Solidity: function resetNodeOperatorManagerAddress(uint256 nodeOperatorId) returns()
func (_CuratedModule *CuratedModuleTransactorSession) ResetNodeOperatorManagerAddress(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.ResetNodeOperatorManagerAddress(&_CuratedModule.TransactOpts, nodeOperatorId)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_CuratedModule *CuratedModuleTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_CuratedModule *CuratedModuleSession) Resume() (*types.Transaction, error) {
	return _CuratedModule.Contract.Resume(&_CuratedModule.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_CuratedModule *CuratedModuleTransactorSession) Resume() (*types.Transaction, error) {
	return _CuratedModule.Contract.Resume(&_CuratedModule.TransactOpts)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CuratedModule *CuratedModuleTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CuratedModule *CuratedModuleSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CuratedModule.Contract.RevokeRole(&_CuratedModule.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CuratedModule *CuratedModuleTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CuratedModule.Contract.RevokeRole(&_CuratedModule.TransactOpts, role, account)
}

// SettleGeneralDelayedPenalty is a paid mutator transaction binding the contract method 0x187d9f92.
//
// Solidity: function settleGeneralDelayedPenalty(uint256[] nodeOperatorIds, uint256[] maxAmounts) returns()
func (_CuratedModule *CuratedModuleTransactor) SettleGeneralDelayedPenalty(opts *bind.TransactOpts, nodeOperatorIds []*big.Int, maxAmounts []*big.Int) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "settleGeneralDelayedPenalty", nodeOperatorIds, maxAmounts)
}

// SettleGeneralDelayedPenalty is a paid mutator transaction binding the contract method 0x187d9f92.
//
// Solidity: function settleGeneralDelayedPenalty(uint256[] nodeOperatorIds, uint256[] maxAmounts) returns()
func (_CuratedModule *CuratedModuleSession) SettleGeneralDelayedPenalty(nodeOperatorIds []*big.Int, maxAmounts []*big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.SettleGeneralDelayedPenalty(&_CuratedModule.TransactOpts, nodeOperatorIds, maxAmounts)
}

// SettleGeneralDelayedPenalty is a paid mutator transaction binding the contract method 0x187d9f92.
//
// Solidity: function settleGeneralDelayedPenalty(uint256[] nodeOperatorIds, uint256[] maxAmounts) returns()
func (_CuratedModule *CuratedModuleTransactorSession) SettleGeneralDelayedPenalty(nodeOperatorIds []*big.Int, maxAmounts []*big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.SettleGeneralDelayedPenalty(&_CuratedModule.TransactOpts, nodeOperatorIds, maxAmounts)
}

// UnsafeUpdateValidatorsCount is a paid mutator transaction binding the contract method 0x94120368.
//
// Solidity: function unsafeUpdateValidatorsCount(uint256 nodeOperatorId, uint256 exitedValidatorsCount) returns()
func (_CuratedModule *CuratedModuleTransactor) UnsafeUpdateValidatorsCount(opts *bind.TransactOpts, nodeOperatorId *big.Int, exitedValidatorsCount *big.Int) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "unsafeUpdateValidatorsCount", nodeOperatorId, exitedValidatorsCount)
}

// UnsafeUpdateValidatorsCount is a paid mutator transaction binding the contract method 0x94120368.
//
// Solidity: function unsafeUpdateValidatorsCount(uint256 nodeOperatorId, uint256 exitedValidatorsCount) returns()
func (_CuratedModule *CuratedModuleSession) UnsafeUpdateValidatorsCount(nodeOperatorId *big.Int, exitedValidatorsCount *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.UnsafeUpdateValidatorsCount(&_CuratedModule.TransactOpts, nodeOperatorId, exitedValidatorsCount)
}

// UnsafeUpdateValidatorsCount is a paid mutator transaction binding the contract method 0x94120368.
//
// Solidity: function unsafeUpdateValidatorsCount(uint256 nodeOperatorId, uint256 exitedValidatorsCount) returns()
func (_CuratedModule *CuratedModuleTransactorSession) UnsafeUpdateValidatorsCount(nodeOperatorId *big.Int, exitedValidatorsCount *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.UnsafeUpdateValidatorsCount(&_CuratedModule.TransactOpts, nodeOperatorId, exitedValidatorsCount)
}

// UpdateDepositInfo is a paid mutator transaction binding the contract method 0xbf9155fe.
//
// Solidity: function updateDepositInfo(uint256 nodeOperatorId) returns()
func (_CuratedModule *CuratedModuleTransactor) UpdateDepositInfo(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "updateDepositInfo", nodeOperatorId)
}

// UpdateDepositInfo is a paid mutator transaction binding the contract method 0xbf9155fe.
//
// Solidity: function updateDepositInfo(uint256 nodeOperatorId) returns()
func (_CuratedModule *CuratedModuleSession) UpdateDepositInfo(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.UpdateDepositInfo(&_CuratedModule.TransactOpts, nodeOperatorId)
}

// UpdateDepositInfo is a paid mutator transaction binding the contract method 0xbf9155fe.
//
// Solidity: function updateDepositInfo(uint256 nodeOperatorId) returns()
func (_CuratedModule *CuratedModuleTransactorSession) UpdateDepositInfo(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.UpdateDepositInfo(&_CuratedModule.TransactOpts, nodeOperatorId)
}

// UpdateDepositableValidatorsCount is a paid mutator transaction binding the contract method 0x8eab3cd0.
//
// Solidity: function updateDepositableValidatorsCount(uint256 nodeOperatorId) returns()
func (_CuratedModule *CuratedModuleTransactor) UpdateDepositableValidatorsCount(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "updateDepositableValidatorsCount", nodeOperatorId)
}

// UpdateDepositableValidatorsCount is a paid mutator transaction binding the contract method 0x8eab3cd0.
//
// Solidity: function updateDepositableValidatorsCount(uint256 nodeOperatorId) returns()
func (_CuratedModule *CuratedModuleSession) UpdateDepositableValidatorsCount(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.UpdateDepositableValidatorsCount(&_CuratedModule.TransactOpts, nodeOperatorId)
}

// UpdateDepositableValidatorsCount is a paid mutator transaction binding the contract method 0x8eab3cd0.
//
// Solidity: function updateDepositableValidatorsCount(uint256 nodeOperatorId) returns()
func (_CuratedModule *CuratedModuleTransactorSession) UpdateDepositableValidatorsCount(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.UpdateDepositableValidatorsCount(&_CuratedModule.TransactOpts, nodeOperatorId)
}

// UpdateExitedValidatorsCount is a paid mutator transaction binding the contract method 0x9b00c146.
//
// Solidity: function updateExitedValidatorsCount(bytes nodeOperatorIds, bytes exitedValidatorsCounts) returns()
func (_CuratedModule *CuratedModuleTransactor) UpdateExitedValidatorsCount(opts *bind.TransactOpts, nodeOperatorIds []byte, exitedValidatorsCounts []byte) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "updateExitedValidatorsCount", nodeOperatorIds, exitedValidatorsCounts)
}

// UpdateExitedValidatorsCount is a paid mutator transaction binding the contract method 0x9b00c146.
//
// Solidity: function updateExitedValidatorsCount(bytes nodeOperatorIds, bytes exitedValidatorsCounts) returns()
func (_CuratedModule *CuratedModuleSession) UpdateExitedValidatorsCount(nodeOperatorIds []byte, exitedValidatorsCounts []byte) (*types.Transaction, error) {
	return _CuratedModule.Contract.UpdateExitedValidatorsCount(&_CuratedModule.TransactOpts, nodeOperatorIds, exitedValidatorsCounts)
}

// UpdateExitedValidatorsCount is a paid mutator transaction binding the contract method 0x9b00c146.
//
// Solidity: function updateExitedValidatorsCount(bytes nodeOperatorIds, bytes exitedValidatorsCounts) returns()
func (_CuratedModule *CuratedModuleTransactorSession) UpdateExitedValidatorsCount(nodeOperatorIds []byte, exitedValidatorsCounts []byte) (*types.Transaction, error) {
	return _CuratedModule.Contract.UpdateExitedValidatorsCount(&_CuratedModule.TransactOpts, nodeOperatorIds, exitedValidatorsCounts)
}

// UpdateTargetValidatorsLimits is a paid mutator transaction binding the contract method 0x08a679ad.
//
// Solidity: function updateTargetValidatorsLimits(uint256 nodeOperatorId, uint256 targetLimitMode, uint256 targetLimit) returns()
func (_CuratedModule *CuratedModuleTransactor) UpdateTargetValidatorsLimits(opts *bind.TransactOpts, nodeOperatorId *big.Int, targetLimitMode *big.Int, targetLimit *big.Int) (*types.Transaction, error) {
	return _CuratedModule.contract.Transact(opts, "updateTargetValidatorsLimits", nodeOperatorId, targetLimitMode, targetLimit)
}

// UpdateTargetValidatorsLimits is a paid mutator transaction binding the contract method 0x08a679ad.
//
// Solidity: function updateTargetValidatorsLimits(uint256 nodeOperatorId, uint256 targetLimitMode, uint256 targetLimit) returns()
func (_CuratedModule *CuratedModuleSession) UpdateTargetValidatorsLimits(nodeOperatorId *big.Int, targetLimitMode *big.Int, targetLimit *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.UpdateTargetValidatorsLimits(&_CuratedModule.TransactOpts, nodeOperatorId, targetLimitMode, targetLimit)
}

// UpdateTargetValidatorsLimits is a paid mutator transaction binding the contract method 0x08a679ad.
//
// Solidity: function updateTargetValidatorsLimits(uint256 nodeOperatorId, uint256 targetLimitMode, uint256 targetLimit) returns()
func (_CuratedModule *CuratedModuleTransactorSession) UpdateTargetValidatorsLimits(nodeOperatorId *big.Int, targetLimitMode *big.Int, targetLimit *big.Int) (*types.Transaction, error) {
	return _CuratedModule.Contract.UpdateTargetValidatorsLimits(&_CuratedModule.TransactOpts, nodeOperatorId, targetLimitMode, targetLimit)
}

// CuratedModuleDepositableSigningKeysCountChangedIterator is returned from FilterDepositableSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for DepositableSigningKeysCountChanged events raised by the CuratedModule contract.
type CuratedModuleDepositableSigningKeysCountChangedIterator struct {
	Event *CuratedModuleDepositableSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *CuratedModuleDepositableSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleDepositableSigningKeysCountChanged)
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
		it.Event = new(CuratedModuleDepositableSigningKeysCountChanged)
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
func (it *CuratedModuleDepositableSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleDepositableSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleDepositableSigningKeysCountChanged represents a DepositableSigningKeysCountChanged event raised by the CuratedModule contract.
type CuratedModuleDepositableSigningKeysCountChanged struct {
	NodeOperatorId       *big.Int
	DepositableKeysCount *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterDepositableSigningKeysCountChanged is a free log retrieval operation binding the contract event 0xf9109091b368cedad2edff45414eef892edd6b4fe80084bd590aa8f8def8ed33.
//
// Solidity: event DepositableSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 depositableKeysCount)
func (_CuratedModule *CuratedModuleFilterer) FilterDepositableSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CuratedModuleDepositableSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "DepositableSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleDepositableSigningKeysCountChangedIterator{contract: _CuratedModule.contract, event: "DepositableSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchDepositableSigningKeysCountChanged is a free log subscription operation binding the contract event 0xf9109091b368cedad2edff45414eef892edd6b4fe80084bd590aa8f8def8ed33.
//
// Solidity: event DepositableSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 depositableKeysCount)
func (_CuratedModule *CuratedModuleFilterer) WatchDepositableSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *CuratedModuleDepositableSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "DepositableSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleDepositableSigningKeysCountChanged)
				if err := _CuratedModule.contract.UnpackLog(event, "DepositableSigningKeysCountChanged", log); err != nil {
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

// ParseDepositableSigningKeysCountChanged is a log parse operation binding the contract event 0xf9109091b368cedad2edff45414eef892edd6b4fe80084bd590aa8f8def8ed33.
//
// Solidity: event DepositableSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 depositableKeysCount)
func (_CuratedModule *CuratedModuleFilterer) ParseDepositableSigningKeysCountChanged(log types.Log) (*CuratedModuleDepositableSigningKeysCountChanged, error) {
	event := new(CuratedModuleDepositableSigningKeysCountChanged)
	if err := _CuratedModule.contract.UnpackLog(event, "DepositableSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleDepositedSigningKeysCountChangedIterator is returned from FilterDepositedSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for DepositedSigningKeysCountChanged events raised by the CuratedModule contract.
type CuratedModuleDepositedSigningKeysCountChangedIterator struct {
	Event *CuratedModuleDepositedSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *CuratedModuleDepositedSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleDepositedSigningKeysCountChanged)
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
		it.Event = new(CuratedModuleDepositedSigningKeysCountChanged)
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
func (it *CuratedModuleDepositedSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleDepositedSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleDepositedSigningKeysCountChanged represents a DepositedSigningKeysCountChanged event raised by the CuratedModule contract.
type CuratedModuleDepositedSigningKeysCountChanged struct {
	NodeOperatorId     *big.Int
	DepositedKeysCount *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDepositedSigningKeysCountChanged is a free log retrieval operation binding the contract event 0x24eb1c9e765ba41accf9437300ea91ece5ed3f897ec3cdee0e9debd7fe309b78.
//
// Solidity: event DepositedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 depositedKeysCount)
func (_CuratedModule *CuratedModuleFilterer) FilterDepositedSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CuratedModuleDepositedSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "DepositedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleDepositedSigningKeysCountChangedIterator{contract: _CuratedModule.contract, event: "DepositedSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchDepositedSigningKeysCountChanged is a free log subscription operation binding the contract event 0x24eb1c9e765ba41accf9437300ea91ece5ed3f897ec3cdee0e9debd7fe309b78.
//
// Solidity: event DepositedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 depositedKeysCount)
func (_CuratedModule *CuratedModuleFilterer) WatchDepositedSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *CuratedModuleDepositedSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "DepositedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleDepositedSigningKeysCountChanged)
				if err := _CuratedModule.contract.UnpackLog(event, "DepositedSigningKeysCountChanged", log); err != nil {
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
// Solidity: event DepositedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 depositedKeysCount)
func (_CuratedModule *CuratedModuleFilterer) ParseDepositedSigningKeysCountChanged(log types.Log) (*CuratedModuleDepositedSigningKeysCountChanged, error) {
	event := new(CuratedModuleDepositedSigningKeysCountChanged)
	if err := _CuratedModule.contract.UnpackLog(event, "DepositedSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleERC1155RecoveredIterator is returned from FilterERC1155Recovered and is used to iterate over the raw logs and unpacked data for ERC1155Recovered events raised by the CuratedModule contract.
type CuratedModuleERC1155RecoveredIterator struct {
	Event *CuratedModuleERC1155Recovered // Event containing the contract specifics and raw log

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
func (it *CuratedModuleERC1155RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleERC1155Recovered)
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
		it.Event = new(CuratedModuleERC1155Recovered)
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
func (it *CuratedModuleERC1155RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleERC1155RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleERC1155Recovered represents a ERC1155Recovered event raised by the CuratedModule contract.
type CuratedModuleERC1155Recovered struct {
	Token     common.Address
	TokenId   *big.Int
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC1155Recovered is a free log retrieval operation binding the contract event 0x5cf02e753b3eb0f4bee4460a72817d8e5e3c75cd4d65c1d0b06dca88b8032936.
//
// Solidity: event ERC1155Recovered(address indexed token, uint256 tokenId, address indexed recipient, uint256 amount)
func (_CuratedModule *CuratedModuleFilterer) FilterERC1155Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CuratedModuleERC1155RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "ERC1155Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleERC1155RecoveredIterator{contract: _CuratedModule.contract, event: "ERC1155Recovered", logs: logs, sub: sub}, nil
}

// WatchERC1155Recovered is a free log subscription operation binding the contract event 0x5cf02e753b3eb0f4bee4460a72817d8e5e3c75cd4d65c1d0b06dca88b8032936.
//
// Solidity: event ERC1155Recovered(address indexed token, uint256 tokenId, address indexed recipient, uint256 amount)
func (_CuratedModule *CuratedModuleFilterer) WatchERC1155Recovered(opts *bind.WatchOpts, sink chan<- *CuratedModuleERC1155Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "ERC1155Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleERC1155Recovered)
				if err := _CuratedModule.contract.UnpackLog(event, "ERC1155Recovered", log); err != nil {
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

// ParseERC1155Recovered is a log parse operation binding the contract event 0x5cf02e753b3eb0f4bee4460a72817d8e5e3c75cd4d65c1d0b06dca88b8032936.
//
// Solidity: event ERC1155Recovered(address indexed token, uint256 tokenId, address indexed recipient, uint256 amount)
func (_CuratedModule *CuratedModuleFilterer) ParseERC1155Recovered(log types.Log) (*CuratedModuleERC1155Recovered, error) {
	event := new(CuratedModuleERC1155Recovered)
	if err := _CuratedModule.contract.UnpackLog(event, "ERC1155Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleERC20RecoveredIterator is returned from FilterERC20Recovered and is used to iterate over the raw logs and unpacked data for ERC20Recovered events raised by the CuratedModule contract.
type CuratedModuleERC20RecoveredIterator struct {
	Event *CuratedModuleERC20Recovered // Event containing the contract specifics and raw log

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
func (it *CuratedModuleERC20RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleERC20Recovered)
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
		it.Event = new(CuratedModuleERC20Recovered)
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
func (it *CuratedModuleERC20RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleERC20RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleERC20Recovered represents a ERC20Recovered event raised by the CuratedModule contract.
type CuratedModuleERC20Recovered struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC20Recovered is a free log retrieval operation binding the contract event 0xaca8fb252cde442184e5f10e0f2e6e4029e8cd7717cae63559079610702436aa.
//
// Solidity: event ERC20Recovered(address indexed token, address indexed recipient, uint256 amount)
func (_CuratedModule *CuratedModuleFilterer) FilterERC20Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CuratedModuleERC20RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "ERC20Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleERC20RecoveredIterator{contract: _CuratedModule.contract, event: "ERC20Recovered", logs: logs, sub: sub}, nil
}

// WatchERC20Recovered is a free log subscription operation binding the contract event 0xaca8fb252cde442184e5f10e0f2e6e4029e8cd7717cae63559079610702436aa.
//
// Solidity: event ERC20Recovered(address indexed token, address indexed recipient, uint256 amount)
func (_CuratedModule *CuratedModuleFilterer) WatchERC20Recovered(opts *bind.WatchOpts, sink chan<- *CuratedModuleERC20Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "ERC20Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleERC20Recovered)
				if err := _CuratedModule.contract.UnpackLog(event, "ERC20Recovered", log); err != nil {
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

// ParseERC20Recovered is a log parse operation binding the contract event 0xaca8fb252cde442184e5f10e0f2e6e4029e8cd7717cae63559079610702436aa.
//
// Solidity: event ERC20Recovered(address indexed token, address indexed recipient, uint256 amount)
func (_CuratedModule *CuratedModuleFilterer) ParseERC20Recovered(log types.Log) (*CuratedModuleERC20Recovered, error) {
	event := new(CuratedModuleERC20Recovered)
	if err := _CuratedModule.contract.UnpackLog(event, "ERC20Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleERC721RecoveredIterator is returned from FilterERC721Recovered and is used to iterate over the raw logs and unpacked data for ERC721Recovered events raised by the CuratedModule contract.
type CuratedModuleERC721RecoveredIterator struct {
	Event *CuratedModuleERC721Recovered // Event containing the contract specifics and raw log

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
func (it *CuratedModuleERC721RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleERC721Recovered)
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
		it.Event = new(CuratedModuleERC721Recovered)
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
func (it *CuratedModuleERC721RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleERC721RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleERC721Recovered represents a ERC721Recovered event raised by the CuratedModule contract.
type CuratedModuleERC721Recovered struct {
	Token     common.Address
	TokenId   *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC721Recovered is a free log retrieval operation binding the contract event 0x8166bf75d2ff2fa3c8f3c44410540bf42e9a5359b48409e8d660291dc9f788c8.
//
// Solidity: event ERC721Recovered(address indexed token, uint256 tokenId, address indexed recipient)
func (_CuratedModule *CuratedModuleFilterer) FilterERC721Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CuratedModuleERC721RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "ERC721Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleERC721RecoveredIterator{contract: _CuratedModule.contract, event: "ERC721Recovered", logs: logs, sub: sub}, nil
}

// WatchERC721Recovered is a free log subscription operation binding the contract event 0x8166bf75d2ff2fa3c8f3c44410540bf42e9a5359b48409e8d660291dc9f788c8.
//
// Solidity: event ERC721Recovered(address indexed token, uint256 tokenId, address indexed recipient)
func (_CuratedModule *CuratedModuleFilterer) WatchERC721Recovered(opts *bind.WatchOpts, sink chan<- *CuratedModuleERC721Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "ERC721Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleERC721Recovered)
				if err := _CuratedModule.contract.UnpackLog(event, "ERC721Recovered", log); err != nil {
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

// ParseERC721Recovered is a log parse operation binding the contract event 0x8166bf75d2ff2fa3c8f3c44410540bf42e9a5359b48409e8d660291dc9f788c8.
//
// Solidity: event ERC721Recovered(address indexed token, uint256 tokenId, address indexed recipient)
func (_CuratedModule *CuratedModuleFilterer) ParseERC721Recovered(log types.Log) (*CuratedModuleERC721Recovered, error) {
	event := new(CuratedModuleERC721Recovered)
	if err := _CuratedModule.contract.UnpackLog(event, "ERC721Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleEtherRecoveredIterator is returned from FilterEtherRecovered and is used to iterate over the raw logs and unpacked data for EtherRecovered events raised by the CuratedModule contract.
type CuratedModuleEtherRecoveredIterator struct {
	Event *CuratedModuleEtherRecovered // Event containing the contract specifics and raw log

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
func (it *CuratedModuleEtherRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleEtherRecovered)
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
		it.Event = new(CuratedModuleEtherRecovered)
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
func (it *CuratedModuleEtherRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleEtherRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleEtherRecovered represents a EtherRecovered event raised by the CuratedModule contract.
type CuratedModuleEtherRecovered struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEtherRecovered is a free log retrieval operation binding the contract event 0x8e274e42262a7f013b700b35c2b4629ccce1702f8fe83f8dfb7eacbb26a4382c.
//
// Solidity: event EtherRecovered(address indexed recipient, uint256 amount)
func (_CuratedModule *CuratedModuleFilterer) FilterEtherRecovered(opts *bind.FilterOpts, recipient []common.Address) (*CuratedModuleEtherRecoveredIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "EtherRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleEtherRecoveredIterator{contract: _CuratedModule.contract, event: "EtherRecovered", logs: logs, sub: sub}, nil
}

// WatchEtherRecovered is a free log subscription operation binding the contract event 0x8e274e42262a7f013b700b35c2b4629ccce1702f8fe83f8dfb7eacbb26a4382c.
//
// Solidity: event EtherRecovered(address indexed recipient, uint256 amount)
func (_CuratedModule *CuratedModuleFilterer) WatchEtherRecovered(opts *bind.WatchOpts, sink chan<- *CuratedModuleEtherRecovered, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "EtherRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleEtherRecovered)
				if err := _CuratedModule.contract.UnpackLog(event, "EtherRecovered", log); err != nil {
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

// ParseEtherRecovered is a log parse operation binding the contract event 0x8e274e42262a7f013b700b35c2b4629ccce1702f8fe83f8dfb7eacbb26a4382c.
//
// Solidity: event EtherRecovered(address indexed recipient, uint256 amount)
func (_CuratedModule *CuratedModuleFilterer) ParseEtherRecovered(log types.Log) (*CuratedModuleEtherRecovered, error) {
	event := new(CuratedModuleEtherRecovered)
	if err := _CuratedModule.contract.UnpackLog(event, "EtherRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleExitedSigningKeysCountChangedIterator is returned from FilterExitedSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for ExitedSigningKeysCountChanged events raised by the CuratedModule contract.
type CuratedModuleExitedSigningKeysCountChangedIterator struct {
	Event *CuratedModuleExitedSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *CuratedModuleExitedSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleExitedSigningKeysCountChanged)
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
		it.Event = new(CuratedModuleExitedSigningKeysCountChanged)
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
func (it *CuratedModuleExitedSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleExitedSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleExitedSigningKeysCountChanged represents a ExitedSigningKeysCountChanged event raised by the CuratedModule contract.
type CuratedModuleExitedSigningKeysCountChanged struct {
	NodeOperatorId  *big.Int
	ExitedKeysCount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterExitedSigningKeysCountChanged is a free log retrieval operation binding the contract event 0x0f67960648751434ae86bf350db61194f387fda387e7f568b0ccd0ae0c220166.
//
// Solidity: event ExitedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 exitedKeysCount)
func (_CuratedModule *CuratedModuleFilterer) FilterExitedSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CuratedModuleExitedSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "ExitedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleExitedSigningKeysCountChangedIterator{contract: _CuratedModule.contract, event: "ExitedSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchExitedSigningKeysCountChanged is a free log subscription operation binding the contract event 0x0f67960648751434ae86bf350db61194f387fda387e7f568b0ccd0ae0c220166.
//
// Solidity: event ExitedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 exitedKeysCount)
func (_CuratedModule *CuratedModuleFilterer) WatchExitedSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *CuratedModuleExitedSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "ExitedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleExitedSigningKeysCountChanged)
				if err := _CuratedModule.contract.UnpackLog(event, "ExitedSigningKeysCountChanged", log); err != nil {
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
// Solidity: event ExitedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 exitedKeysCount)
func (_CuratedModule *CuratedModuleFilterer) ParseExitedSigningKeysCountChanged(log types.Log) (*CuratedModuleExitedSigningKeysCountChanged, error) {
	event := new(CuratedModuleExitedSigningKeysCountChanged)
	if err := _CuratedModule.contract.UnpackLog(event, "ExitedSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleFullDepositInfoUpdateRequestedIterator is returned from FilterFullDepositInfoUpdateRequested and is used to iterate over the raw logs and unpacked data for FullDepositInfoUpdateRequested events raised by the CuratedModule contract.
type CuratedModuleFullDepositInfoUpdateRequestedIterator struct {
	Event *CuratedModuleFullDepositInfoUpdateRequested // Event containing the contract specifics and raw log

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
func (it *CuratedModuleFullDepositInfoUpdateRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleFullDepositInfoUpdateRequested)
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
		it.Event = new(CuratedModuleFullDepositInfoUpdateRequested)
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
func (it *CuratedModuleFullDepositInfoUpdateRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleFullDepositInfoUpdateRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleFullDepositInfoUpdateRequested represents a FullDepositInfoUpdateRequested event raised by the CuratedModule contract.
type CuratedModuleFullDepositInfoUpdateRequested struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFullDepositInfoUpdateRequested is a free log retrieval operation binding the contract event 0x5b98e65689fe2cae1e3cd4c7e47c05e552433fe6dbcf63c9d128f02887338bc7.
//
// Solidity: event FullDepositInfoUpdateRequested()
func (_CuratedModule *CuratedModuleFilterer) FilterFullDepositInfoUpdateRequested(opts *bind.FilterOpts) (*CuratedModuleFullDepositInfoUpdateRequestedIterator, error) {

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "FullDepositInfoUpdateRequested")
	if err != nil {
		return nil, err
	}
	return &CuratedModuleFullDepositInfoUpdateRequestedIterator{contract: _CuratedModule.contract, event: "FullDepositInfoUpdateRequested", logs: logs, sub: sub}, nil
}

// WatchFullDepositInfoUpdateRequested is a free log subscription operation binding the contract event 0x5b98e65689fe2cae1e3cd4c7e47c05e552433fe6dbcf63c9d128f02887338bc7.
//
// Solidity: event FullDepositInfoUpdateRequested()
func (_CuratedModule *CuratedModuleFilterer) WatchFullDepositInfoUpdateRequested(opts *bind.WatchOpts, sink chan<- *CuratedModuleFullDepositInfoUpdateRequested) (event.Subscription, error) {

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "FullDepositInfoUpdateRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleFullDepositInfoUpdateRequested)
				if err := _CuratedModule.contract.UnpackLog(event, "FullDepositInfoUpdateRequested", log); err != nil {
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

// ParseFullDepositInfoUpdateRequested is a log parse operation binding the contract event 0x5b98e65689fe2cae1e3cd4c7e47c05e552433fe6dbcf63c9d128f02887338bc7.
//
// Solidity: event FullDepositInfoUpdateRequested()
func (_CuratedModule *CuratedModuleFilterer) ParseFullDepositInfoUpdateRequested(log types.Log) (*CuratedModuleFullDepositInfoUpdateRequested, error) {
	event := new(CuratedModuleFullDepositInfoUpdateRequested)
	if err := _CuratedModule.contract.UnpackLog(event, "FullDepositInfoUpdateRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleGeneralDelayedPenaltyCancelledIterator is returned from FilterGeneralDelayedPenaltyCancelled and is used to iterate over the raw logs and unpacked data for GeneralDelayedPenaltyCancelled events raised by the CuratedModule contract.
type CuratedModuleGeneralDelayedPenaltyCancelledIterator struct {
	Event *CuratedModuleGeneralDelayedPenaltyCancelled // Event containing the contract specifics and raw log

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
func (it *CuratedModuleGeneralDelayedPenaltyCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleGeneralDelayedPenaltyCancelled)
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
		it.Event = new(CuratedModuleGeneralDelayedPenaltyCancelled)
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
func (it *CuratedModuleGeneralDelayedPenaltyCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleGeneralDelayedPenaltyCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleGeneralDelayedPenaltyCancelled represents a GeneralDelayedPenaltyCancelled event raised by the CuratedModule contract.
type CuratedModuleGeneralDelayedPenaltyCancelled struct {
	NodeOperatorId *big.Int
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGeneralDelayedPenaltyCancelled is a free log retrieval operation binding the contract event 0x8b4e728bbcfa37901be4c710fcd1e3fa980fa4dc43354610ada4174684a561dd.
//
// Solidity: event GeneralDelayedPenaltyCancelled(uint256 indexed nodeOperatorId, uint256 amount)
func (_CuratedModule *CuratedModuleFilterer) FilterGeneralDelayedPenaltyCancelled(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CuratedModuleGeneralDelayedPenaltyCancelledIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "GeneralDelayedPenaltyCancelled", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleGeneralDelayedPenaltyCancelledIterator{contract: _CuratedModule.contract, event: "GeneralDelayedPenaltyCancelled", logs: logs, sub: sub}, nil
}

// WatchGeneralDelayedPenaltyCancelled is a free log subscription operation binding the contract event 0x8b4e728bbcfa37901be4c710fcd1e3fa980fa4dc43354610ada4174684a561dd.
//
// Solidity: event GeneralDelayedPenaltyCancelled(uint256 indexed nodeOperatorId, uint256 amount)
func (_CuratedModule *CuratedModuleFilterer) WatchGeneralDelayedPenaltyCancelled(opts *bind.WatchOpts, sink chan<- *CuratedModuleGeneralDelayedPenaltyCancelled, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "GeneralDelayedPenaltyCancelled", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleGeneralDelayedPenaltyCancelled)
				if err := _CuratedModule.contract.UnpackLog(event, "GeneralDelayedPenaltyCancelled", log); err != nil {
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

// ParseGeneralDelayedPenaltyCancelled is a log parse operation binding the contract event 0x8b4e728bbcfa37901be4c710fcd1e3fa980fa4dc43354610ada4174684a561dd.
//
// Solidity: event GeneralDelayedPenaltyCancelled(uint256 indexed nodeOperatorId, uint256 amount)
func (_CuratedModule *CuratedModuleFilterer) ParseGeneralDelayedPenaltyCancelled(log types.Log) (*CuratedModuleGeneralDelayedPenaltyCancelled, error) {
	event := new(CuratedModuleGeneralDelayedPenaltyCancelled)
	if err := _CuratedModule.contract.UnpackLog(event, "GeneralDelayedPenaltyCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleGeneralDelayedPenaltyCompensatedIterator is returned from FilterGeneralDelayedPenaltyCompensated and is used to iterate over the raw logs and unpacked data for GeneralDelayedPenaltyCompensated events raised by the CuratedModule contract.
type CuratedModuleGeneralDelayedPenaltyCompensatedIterator struct {
	Event *CuratedModuleGeneralDelayedPenaltyCompensated // Event containing the contract specifics and raw log

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
func (it *CuratedModuleGeneralDelayedPenaltyCompensatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleGeneralDelayedPenaltyCompensated)
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
		it.Event = new(CuratedModuleGeneralDelayedPenaltyCompensated)
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
func (it *CuratedModuleGeneralDelayedPenaltyCompensatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleGeneralDelayedPenaltyCompensatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleGeneralDelayedPenaltyCompensated represents a GeneralDelayedPenaltyCompensated event raised by the CuratedModule contract.
type CuratedModuleGeneralDelayedPenaltyCompensated struct {
	NodeOperatorId *big.Int
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGeneralDelayedPenaltyCompensated is a free log retrieval operation binding the contract event 0x7d380a8a0ebe307944d48f30df0c4bfdeda1f9d1093ea28bde067811be3b9449.
//
// Solidity: event GeneralDelayedPenaltyCompensated(uint256 indexed nodeOperatorId, uint256 amount)
func (_CuratedModule *CuratedModuleFilterer) FilterGeneralDelayedPenaltyCompensated(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CuratedModuleGeneralDelayedPenaltyCompensatedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "GeneralDelayedPenaltyCompensated", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleGeneralDelayedPenaltyCompensatedIterator{contract: _CuratedModule.contract, event: "GeneralDelayedPenaltyCompensated", logs: logs, sub: sub}, nil
}

// WatchGeneralDelayedPenaltyCompensated is a free log subscription operation binding the contract event 0x7d380a8a0ebe307944d48f30df0c4bfdeda1f9d1093ea28bde067811be3b9449.
//
// Solidity: event GeneralDelayedPenaltyCompensated(uint256 indexed nodeOperatorId, uint256 amount)
func (_CuratedModule *CuratedModuleFilterer) WatchGeneralDelayedPenaltyCompensated(opts *bind.WatchOpts, sink chan<- *CuratedModuleGeneralDelayedPenaltyCompensated, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "GeneralDelayedPenaltyCompensated", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleGeneralDelayedPenaltyCompensated)
				if err := _CuratedModule.contract.UnpackLog(event, "GeneralDelayedPenaltyCompensated", log); err != nil {
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

// ParseGeneralDelayedPenaltyCompensated is a log parse operation binding the contract event 0x7d380a8a0ebe307944d48f30df0c4bfdeda1f9d1093ea28bde067811be3b9449.
//
// Solidity: event GeneralDelayedPenaltyCompensated(uint256 indexed nodeOperatorId, uint256 amount)
func (_CuratedModule *CuratedModuleFilterer) ParseGeneralDelayedPenaltyCompensated(log types.Log) (*CuratedModuleGeneralDelayedPenaltyCompensated, error) {
	event := new(CuratedModuleGeneralDelayedPenaltyCompensated)
	if err := _CuratedModule.contract.UnpackLog(event, "GeneralDelayedPenaltyCompensated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleGeneralDelayedPenaltyReportedIterator is returned from FilterGeneralDelayedPenaltyReported and is used to iterate over the raw logs and unpacked data for GeneralDelayedPenaltyReported events raised by the CuratedModule contract.
type CuratedModuleGeneralDelayedPenaltyReportedIterator struct {
	Event *CuratedModuleGeneralDelayedPenaltyReported // Event containing the contract specifics and raw log

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
func (it *CuratedModuleGeneralDelayedPenaltyReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleGeneralDelayedPenaltyReported)
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
		it.Event = new(CuratedModuleGeneralDelayedPenaltyReported)
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
func (it *CuratedModuleGeneralDelayedPenaltyReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleGeneralDelayedPenaltyReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleGeneralDelayedPenaltyReported represents a GeneralDelayedPenaltyReported event raised by the CuratedModule contract.
type CuratedModuleGeneralDelayedPenaltyReported struct {
	NodeOperatorId *big.Int
	PenaltyType    [32]byte
	Amount         *big.Int
	AdditionalFine *big.Int
	Details        string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGeneralDelayedPenaltyReported is a free log retrieval operation binding the contract event 0xe661a158813f91626d18f998442c79bcee679a1649aecc04b508f7fc761e668a.
//
// Solidity: event GeneralDelayedPenaltyReported(uint256 indexed nodeOperatorId, bytes32 indexed penaltyType, uint256 amount, uint256 additionalFine, string details)
func (_CuratedModule *CuratedModuleFilterer) FilterGeneralDelayedPenaltyReported(opts *bind.FilterOpts, nodeOperatorId []*big.Int, penaltyType [][32]byte) (*CuratedModuleGeneralDelayedPenaltyReportedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var penaltyTypeRule []interface{}
	for _, penaltyTypeItem := range penaltyType {
		penaltyTypeRule = append(penaltyTypeRule, penaltyTypeItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "GeneralDelayedPenaltyReported", nodeOperatorIdRule, penaltyTypeRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleGeneralDelayedPenaltyReportedIterator{contract: _CuratedModule.contract, event: "GeneralDelayedPenaltyReported", logs: logs, sub: sub}, nil
}

// WatchGeneralDelayedPenaltyReported is a free log subscription operation binding the contract event 0xe661a158813f91626d18f998442c79bcee679a1649aecc04b508f7fc761e668a.
//
// Solidity: event GeneralDelayedPenaltyReported(uint256 indexed nodeOperatorId, bytes32 indexed penaltyType, uint256 amount, uint256 additionalFine, string details)
func (_CuratedModule *CuratedModuleFilterer) WatchGeneralDelayedPenaltyReported(opts *bind.WatchOpts, sink chan<- *CuratedModuleGeneralDelayedPenaltyReported, nodeOperatorId []*big.Int, penaltyType [][32]byte) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var penaltyTypeRule []interface{}
	for _, penaltyTypeItem := range penaltyType {
		penaltyTypeRule = append(penaltyTypeRule, penaltyTypeItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "GeneralDelayedPenaltyReported", nodeOperatorIdRule, penaltyTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleGeneralDelayedPenaltyReported)
				if err := _CuratedModule.contract.UnpackLog(event, "GeneralDelayedPenaltyReported", log); err != nil {
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

// ParseGeneralDelayedPenaltyReported is a log parse operation binding the contract event 0xe661a158813f91626d18f998442c79bcee679a1649aecc04b508f7fc761e668a.
//
// Solidity: event GeneralDelayedPenaltyReported(uint256 indexed nodeOperatorId, bytes32 indexed penaltyType, uint256 amount, uint256 additionalFine, string details)
func (_CuratedModule *CuratedModuleFilterer) ParseGeneralDelayedPenaltyReported(log types.Log) (*CuratedModuleGeneralDelayedPenaltyReported, error) {
	event := new(CuratedModuleGeneralDelayedPenaltyReported)
	if err := _CuratedModule.contract.UnpackLog(event, "GeneralDelayedPenaltyReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleGeneralDelayedPenaltySettledIterator is returned from FilterGeneralDelayedPenaltySettled and is used to iterate over the raw logs and unpacked data for GeneralDelayedPenaltySettled events raised by the CuratedModule contract.
type CuratedModuleGeneralDelayedPenaltySettledIterator struct {
	Event *CuratedModuleGeneralDelayedPenaltySettled // Event containing the contract specifics and raw log

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
func (it *CuratedModuleGeneralDelayedPenaltySettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleGeneralDelayedPenaltySettled)
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
		it.Event = new(CuratedModuleGeneralDelayedPenaltySettled)
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
func (it *CuratedModuleGeneralDelayedPenaltySettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleGeneralDelayedPenaltySettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleGeneralDelayedPenaltySettled represents a GeneralDelayedPenaltySettled event raised by the CuratedModule contract.
type CuratedModuleGeneralDelayedPenaltySettled struct {
	NodeOperatorId *big.Int
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGeneralDelayedPenaltySettled is a free log retrieval operation binding the contract event 0x89c387b7e64bd581674f53a27c3f1662310cd166492dc734a32e91dc4888722b.
//
// Solidity: event GeneralDelayedPenaltySettled(uint256 indexed nodeOperatorId, uint256 amount)
func (_CuratedModule *CuratedModuleFilterer) FilterGeneralDelayedPenaltySettled(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CuratedModuleGeneralDelayedPenaltySettledIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "GeneralDelayedPenaltySettled", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleGeneralDelayedPenaltySettledIterator{contract: _CuratedModule.contract, event: "GeneralDelayedPenaltySettled", logs: logs, sub: sub}, nil
}

// WatchGeneralDelayedPenaltySettled is a free log subscription operation binding the contract event 0x89c387b7e64bd581674f53a27c3f1662310cd166492dc734a32e91dc4888722b.
//
// Solidity: event GeneralDelayedPenaltySettled(uint256 indexed nodeOperatorId, uint256 amount)
func (_CuratedModule *CuratedModuleFilterer) WatchGeneralDelayedPenaltySettled(opts *bind.WatchOpts, sink chan<- *CuratedModuleGeneralDelayedPenaltySettled, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "GeneralDelayedPenaltySettled", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleGeneralDelayedPenaltySettled)
				if err := _CuratedModule.contract.UnpackLog(event, "GeneralDelayedPenaltySettled", log); err != nil {
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

// ParseGeneralDelayedPenaltySettled is a log parse operation binding the contract event 0x89c387b7e64bd581674f53a27c3f1662310cd166492dc734a32e91dc4888722b.
//
// Solidity: event GeneralDelayedPenaltySettled(uint256 indexed nodeOperatorId, uint256 amount)
func (_CuratedModule *CuratedModuleFilterer) ParseGeneralDelayedPenaltySettled(log types.Log) (*CuratedModuleGeneralDelayedPenaltySettled, error) {
	event := new(CuratedModuleGeneralDelayedPenaltySettled)
	if err := _CuratedModule.contract.UnpackLog(event, "GeneralDelayedPenaltySettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CuratedModule contract.
type CuratedModuleInitializedIterator struct {
	Event *CuratedModuleInitialized // Event containing the contract specifics and raw log

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
func (it *CuratedModuleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleInitialized)
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
		it.Event = new(CuratedModuleInitialized)
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
func (it *CuratedModuleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleInitialized represents a Initialized event raised by the CuratedModule contract.
type CuratedModuleInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CuratedModule *CuratedModuleFilterer) FilterInitialized(opts *bind.FilterOpts) (*CuratedModuleInitializedIterator, error) {

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CuratedModuleInitializedIterator{contract: _CuratedModule.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CuratedModule *CuratedModuleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CuratedModuleInitialized) (event.Subscription, error) {

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleInitialized)
				if err := _CuratedModule.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CuratedModule *CuratedModuleFilterer) ParseInitialized(log types.Log) (*CuratedModuleInitialized, error) {
	event := new(CuratedModuleInitialized)
	if err := _CuratedModule.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleKeyAllocatedBalanceChangedIterator is returned from FilterKeyAllocatedBalanceChanged and is used to iterate over the raw logs and unpacked data for KeyAllocatedBalanceChanged events raised by the CuratedModule contract.
type CuratedModuleKeyAllocatedBalanceChangedIterator struct {
	Event *CuratedModuleKeyAllocatedBalanceChanged // Event containing the contract specifics and raw log

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
func (it *CuratedModuleKeyAllocatedBalanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleKeyAllocatedBalanceChanged)
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
		it.Event = new(CuratedModuleKeyAllocatedBalanceChanged)
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
func (it *CuratedModuleKeyAllocatedBalanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleKeyAllocatedBalanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleKeyAllocatedBalanceChanged represents a KeyAllocatedBalanceChanged event raised by the CuratedModule contract.
type CuratedModuleKeyAllocatedBalanceChanged struct {
	NodeOperatorId *big.Int
	KeyIndex       *big.Int
	NewTotal       *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterKeyAllocatedBalanceChanged is a free log retrieval operation binding the contract event 0x2d41fbb1b43f39a045c78d6cfea06db95d33a62f4df941c1f7e77a5a09c481d7.
//
// Solidity: event KeyAllocatedBalanceChanged(uint256 indexed nodeOperatorId, uint256 indexed keyIndex, uint256 newTotal)
func (_CuratedModule *CuratedModuleFilterer) FilterKeyAllocatedBalanceChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int, keyIndex []*big.Int) (*CuratedModuleKeyAllocatedBalanceChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var keyIndexRule []interface{}
	for _, keyIndexItem := range keyIndex {
		keyIndexRule = append(keyIndexRule, keyIndexItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "KeyAllocatedBalanceChanged", nodeOperatorIdRule, keyIndexRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleKeyAllocatedBalanceChangedIterator{contract: _CuratedModule.contract, event: "KeyAllocatedBalanceChanged", logs: logs, sub: sub}, nil
}

// WatchKeyAllocatedBalanceChanged is a free log subscription operation binding the contract event 0x2d41fbb1b43f39a045c78d6cfea06db95d33a62f4df941c1f7e77a5a09c481d7.
//
// Solidity: event KeyAllocatedBalanceChanged(uint256 indexed nodeOperatorId, uint256 indexed keyIndex, uint256 newTotal)
func (_CuratedModule *CuratedModuleFilterer) WatchKeyAllocatedBalanceChanged(opts *bind.WatchOpts, sink chan<- *CuratedModuleKeyAllocatedBalanceChanged, nodeOperatorId []*big.Int, keyIndex []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var keyIndexRule []interface{}
	for _, keyIndexItem := range keyIndex {
		keyIndexRule = append(keyIndexRule, keyIndexItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "KeyAllocatedBalanceChanged", nodeOperatorIdRule, keyIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleKeyAllocatedBalanceChanged)
				if err := _CuratedModule.contract.UnpackLog(event, "KeyAllocatedBalanceChanged", log); err != nil {
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

// ParseKeyAllocatedBalanceChanged is a log parse operation binding the contract event 0x2d41fbb1b43f39a045c78d6cfea06db95d33a62f4df941c1f7e77a5a09c481d7.
//
// Solidity: event KeyAllocatedBalanceChanged(uint256 indexed nodeOperatorId, uint256 indexed keyIndex, uint256 newTotal)
func (_CuratedModule *CuratedModuleFilterer) ParseKeyAllocatedBalanceChanged(log types.Log) (*CuratedModuleKeyAllocatedBalanceChanged, error) {
	event := new(CuratedModuleKeyAllocatedBalanceChanged)
	if err := _CuratedModule.contract.UnpackLog(event, "KeyAllocatedBalanceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleKeyConfirmedBalanceChangedIterator is returned from FilterKeyConfirmedBalanceChanged and is used to iterate over the raw logs and unpacked data for KeyConfirmedBalanceChanged events raised by the CuratedModule contract.
type CuratedModuleKeyConfirmedBalanceChangedIterator struct {
	Event *CuratedModuleKeyConfirmedBalanceChanged // Event containing the contract specifics and raw log

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
func (it *CuratedModuleKeyConfirmedBalanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleKeyConfirmedBalanceChanged)
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
		it.Event = new(CuratedModuleKeyConfirmedBalanceChanged)
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
func (it *CuratedModuleKeyConfirmedBalanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleKeyConfirmedBalanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleKeyConfirmedBalanceChanged represents a KeyConfirmedBalanceChanged event raised by the CuratedModule contract.
type CuratedModuleKeyConfirmedBalanceChanged struct {
	NodeOperatorId *big.Int
	KeyIndex       *big.Int
	NewBalance     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterKeyConfirmedBalanceChanged is a free log retrieval operation binding the contract event 0xbfcda08fc85741d88f016f7e6682b8ec791f26913bf655d9553f56f1c9a8bdc7.
//
// Solidity: event KeyConfirmedBalanceChanged(uint256 indexed nodeOperatorId, uint256 indexed keyIndex, uint256 newBalance)
func (_CuratedModule *CuratedModuleFilterer) FilterKeyConfirmedBalanceChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int, keyIndex []*big.Int) (*CuratedModuleKeyConfirmedBalanceChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var keyIndexRule []interface{}
	for _, keyIndexItem := range keyIndex {
		keyIndexRule = append(keyIndexRule, keyIndexItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "KeyConfirmedBalanceChanged", nodeOperatorIdRule, keyIndexRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleKeyConfirmedBalanceChangedIterator{contract: _CuratedModule.contract, event: "KeyConfirmedBalanceChanged", logs: logs, sub: sub}, nil
}

// WatchKeyConfirmedBalanceChanged is a free log subscription operation binding the contract event 0xbfcda08fc85741d88f016f7e6682b8ec791f26913bf655d9553f56f1c9a8bdc7.
//
// Solidity: event KeyConfirmedBalanceChanged(uint256 indexed nodeOperatorId, uint256 indexed keyIndex, uint256 newBalance)
func (_CuratedModule *CuratedModuleFilterer) WatchKeyConfirmedBalanceChanged(opts *bind.WatchOpts, sink chan<- *CuratedModuleKeyConfirmedBalanceChanged, nodeOperatorId []*big.Int, keyIndex []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var keyIndexRule []interface{}
	for _, keyIndexItem := range keyIndex {
		keyIndexRule = append(keyIndexRule, keyIndexItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "KeyConfirmedBalanceChanged", nodeOperatorIdRule, keyIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleKeyConfirmedBalanceChanged)
				if err := _CuratedModule.contract.UnpackLog(event, "KeyConfirmedBalanceChanged", log); err != nil {
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

// ParseKeyConfirmedBalanceChanged is a log parse operation binding the contract event 0xbfcda08fc85741d88f016f7e6682b8ec791f26913bf655d9553f56f1c9a8bdc7.
//
// Solidity: event KeyConfirmedBalanceChanged(uint256 indexed nodeOperatorId, uint256 indexed keyIndex, uint256 newBalance)
func (_CuratedModule *CuratedModuleFilterer) ParseKeyConfirmedBalanceChanged(log types.Log) (*CuratedModuleKeyConfirmedBalanceChanged, error) {
	event := new(CuratedModuleKeyConfirmedBalanceChanged)
	if err := _CuratedModule.contract.UnpackLog(event, "KeyConfirmedBalanceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleKeyRemovalChargeAppliedIterator is returned from FilterKeyRemovalChargeApplied and is used to iterate over the raw logs and unpacked data for KeyRemovalChargeApplied events raised by the CuratedModule contract.
type CuratedModuleKeyRemovalChargeAppliedIterator struct {
	Event *CuratedModuleKeyRemovalChargeApplied // Event containing the contract specifics and raw log

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
func (it *CuratedModuleKeyRemovalChargeAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleKeyRemovalChargeApplied)
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
		it.Event = new(CuratedModuleKeyRemovalChargeApplied)
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
func (it *CuratedModuleKeyRemovalChargeAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleKeyRemovalChargeAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleKeyRemovalChargeApplied represents a KeyRemovalChargeApplied event raised by the CuratedModule contract.
type CuratedModuleKeyRemovalChargeApplied struct {
	NodeOperatorId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterKeyRemovalChargeApplied is a free log retrieval operation binding the contract event 0x1cbb8dafbedbdf4f813a8ed1f50d871def63e1104f8729b677af57905eda90f6.
//
// Solidity: event KeyRemovalChargeApplied(uint256 indexed nodeOperatorId)
func (_CuratedModule *CuratedModuleFilterer) FilterKeyRemovalChargeApplied(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CuratedModuleKeyRemovalChargeAppliedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "KeyRemovalChargeApplied", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleKeyRemovalChargeAppliedIterator{contract: _CuratedModule.contract, event: "KeyRemovalChargeApplied", logs: logs, sub: sub}, nil
}

// WatchKeyRemovalChargeApplied is a free log subscription operation binding the contract event 0x1cbb8dafbedbdf4f813a8ed1f50d871def63e1104f8729b677af57905eda90f6.
//
// Solidity: event KeyRemovalChargeApplied(uint256 indexed nodeOperatorId)
func (_CuratedModule *CuratedModuleFilterer) WatchKeyRemovalChargeApplied(opts *bind.WatchOpts, sink chan<- *CuratedModuleKeyRemovalChargeApplied, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "KeyRemovalChargeApplied", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleKeyRemovalChargeApplied)
				if err := _CuratedModule.contract.UnpackLog(event, "KeyRemovalChargeApplied", log); err != nil {
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

// ParseKeyRemovalChargeApplied is a log parse operation binding the contract event 0x1cbb8dafbedbdf4f813a8ed1f50d871def63e1104f8729b677af57905eda90f6.
//
// Solidity: event KeyRemovalChargeApplied(uint256 indexed nodeOperatorId)
func (_CuratedModule *CuratedModuleFilterer) ParseKeyRemovalChargeApplied(log types.Log) (*CuratedModuleKeyRemovalChargeApplied, error) {
	event := new(CuratedModuleKeyRemovalChargeApplied)
	if err := _CuratedModule.contract.UnpackLog(event, "KeyRemovalChargeApplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleNodeOperatorAddedIterator is returned from FilterNodeOperatorAdded and is used to iterate over the raw logs and unpacked data for NodeOperatorAdded events raised by the CuratedModule contract.
type CuratedModuleNodeOperatorAddedIterator struct {
	Event *CuratedModuleNodeOperatorAdded // Event containing the contract specifics and raw log

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
func (it *CuratedModuleNodeOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleNodeOperatorAdded)
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
		it.Event = new(CuratedModuleNodeOperatorAdded)
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
func (it *CuratedModuleNodeOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleNodeOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleNodeOperatorAdded represents a NodeOperatorAdded event raised by the CuratedModule contract.
type CuratedModuleNodeOperatorAdded struct {
	NodeOperatorId             *big.Int
	ManagerAddress             common.Address
	RewardAddress              common.Address
	ExtendedManagerPermissions bool
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorAdded is a free log retrieval operation binding the contract event 0xf17baf73d46b0a80157c3ea3dda1bf081a702732d53ff1720f85e55d9f0997c0.
//
// Solidity: event NodeOperatorAdded(uint256 indexed nodeOperatorId, address indexed managerAddress, address indexed rewardAddress, bool extendedManagerPermissions)
func (_CuratedModule *CuratedModuleFilterer) FilterNodeOperatorAdded(opts *bind.FilterOpts, nodeOperatorId []*big.Int, managerAddress []common.Address, rewardAddress []common.Address) (*CuratedModuleNodeOperatorAddedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var managerAddressRule []interface{}
	for _, managerAddressItem := range managerAddress {
		managerAddressRule = append(managerAddressRule, managerAddressItem)
	}
	var rewardAddressRule []interface{}
	for _, rewardAddressItem := range rewardAddress {
		rewardAddressRule = append(rewardAddressRule, rewardAddressItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "NodeOperatorAdded", nodeOperatorIdRule, managerAddressRule, rewardAddressRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleNodeOperatorAddedIterator{contract: _CuratedModule.contract, event: "NodeOperatorAdded", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorAdded is a free log subscription operation binding the contract event 0xf17baf73d46b0a80157c3ea3dda1bf081a702732d53ff1720f85e55d9f0997c0.
//
// Solidity: event NodeOperatorAdded(uint256 indexed nodeOperatorId, address indexed managerAddress, address indexed rewardAddress, bool extendedManagerPermissions)
func (_CuratedModule *CuratedModuleFilterer) WatchNodeOperatorAdded(opts *bind.WatchOpts, sink chan<- *CuratedModuleNodeOperatorAdded, nodeOperatorId []*big.Int, managerAddress []common.Address, rewardAddress []common.Address) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var managerAddressRule []interface{}
	for _, managerAddressItem := range managerAddress {
		managerAddressRule = append(managerAddressRule, managerAddressItem)
	}
	var rewardAddressRule []interface{}
	for _, rewardAddressItem := range rewardAddress {
		rewardAddressRule = append(rewardAddressRule, rewardAddressItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "NodeOperatorAdded", nodeOperatorIdRule, managerAddressRule, rewardAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleNodeOperatorAdded)
				if err := _CuratedModule.contract.UnpackLog(event, "NodeOperatorAdded", log); err != nil {
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

// ParseNodeOperatorAdded is a log parse operation binding the contract event 0xf17baf73d46b0a80157c3ea3dda1bf081a702732d53ff1720f85e55d9f0997c0.
//
// Solidity: event NodeOperatorAdded(uint256 indexed nodeOperatorId, address indexed managerAddress, address indexed rewardAddress, bool extendedManagerPermissions)
func (_CuratedModule *CuratedModuleFilterer) ParseNodeOperatorAdded(log types.Log) (*CuratedModuleNodeOperatorAdded, error) {
	event := new(CuratedModuleNodeOperatorAdded)
	if err := _CuratedModule.contract.UnpackLog(event, "NodeOperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleNodeOperatorBalanceUpdatedIterator is returned from FilterNodeOperatorBalanceUpdated and is used to iterate over the raw logs and unpacked data for NodeOperatorBalanceUpdated events raised by the CuratedModule contract.
type CuratedModuleNodeOperatorBalanceUpdatedIterator struct {
	Event *CuratedModuleNodeOperatorBalanceUpdated // Event containing the contract specifics and raw log

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
func (it *CuratedModuleNodeOperatorBalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleNodeOperatorBalanceUpdated)
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
		it.Event = new(CuratedModuleNodeOperatorBalanceUpdated)
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
func (it *CuratedModuleNodeOperatorBalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleNodeOperatorBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleNodeOperatorBalanceUpdated represents a NodeOperatorBalanceUpdated event raised by the CuratedModule contract.
type CuratedModuleNodeOperatorBalanceUpdated struct {
	OperatorId *big.Int
	BalanceWei *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorBalanceUpdated is a free log retrieval operation binding the contract event 0x3d23c31684a7b8fe8c0d0e2dd1c25dedb14b8c3020c94de061cb72a4bb07f1b7.
//
// Solidity: event NodeOperatorBalanceUpdated(uint256 indexed operatorId, uint256 balanceWei)
func (_CuratedModule *CuratedModuleFilterer) FilterNodeOperatorBalanceUpdated(opts *bind.FilterOpts, operatorId []*big.Int) (*CuratedModuleNodeOperatorBalanceUpdatedIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "NodeOperatorBalanceUpdated", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleNodeOperatorBalanceUpdatedIterator{contract: _CuratedModule.contract, event: "NodeOperatorBalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorBalanceUpdated is a free log subscription operation binding the contract event 0x3d23c31684a7b8fe8c0d0e2dd1c25dedb14b8c3020c94de061cb72a4bb07f1b7.
//
// Solidity: event NodeOperatorBalanceUpdated(uint256 indexed operatorId, uint256 balanceWei)
func (_CuratedModule *CuratedModuleFilterer) WatchNodeOperatorBalanceUpdated(opts *bind.WatchOpts, sink chan<- *CuratedModuleNodeOperatorBalanceUpdated, operatorId []*big.Int) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "NodeOperatorBalanceUpdated", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleNodeOperatorBalanceUpdated)
				if err := _CuratedModule.contract.UnpackLog(event, "NodeOperatorBalanceUpdated", log); err != nil {
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

// ParseNodeOperatorBalanceUpdated is a log parse operation binding the contract event 0x3d23c31684a7b8fe8c0d0e2dd1c25dedb14b8c3020c94de061cb72a4bb07f1b7.
//
// Solidity: event NodeOperatorBalanceUpdated(uint256 indexed operatorId, uint256 balanceWei)
func (_CuratedModule *CuratedModuleFilterer) ParseNodeOperatorBalanceUpdated(log types.Log) (*CuratedModuleNodeOperatorBalanceUpdated, error) {
	event := new(CuratedModuleNodeOperatorBalanceUpdated)
	if err := _CuratedModule.contract.UnpackLog(event, "NodeOperatorBalanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleNodeOperatorDepositInfoFullyUpdatedIterator is returned from FilterNodeOperatorDepositInfoFullyUpdated and is used to iterate over the raw logs and unpacked data for NodeOperatorDepositInfoFullyUpdated events raised by the CuratedModule contract.
type CuratedModuleNodeOperatorDepositInfoFullyUpdatedIterator struct {
	Event *CuratedModuleNodeOperatorDepositInfoFullyUpdated // Event containing the contract specifics and raw log

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
func (it *CuratedModuleNodeOperatorDepositInfoFullyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleNodeOperatorDepositInfoFullyUpdated)
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
		it.Event = new(CuratedModuleNodeOperatorDepositInfoFullyUpdated)
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
func (it *CuratedModuleNodeOperatorDepositInfoFullyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleNodeOperatorDepositInfoFullyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleNodeOperatorDepositInfoFullyUpdated represents a NodeOperatorDepositInfoFullyUpdated event raised by the CuratedModule contract.
type CuratedModuleNodeOperatorDepositInfoFullyUpdated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorDepositInfoFullyUpdated is a free log retrieval operation binding the contract event 0x2326ef260a3c0b1ad432887941d7041ce2672d2b8a3b87d2f5a9ff9f674ee35a.
//
// Solidity: event NodeOperatorDepositInfoFullyUpdated()
func (_CuratedModule *CuratedModuleFilterer) FilterNodeOperatorDepositInfoFullyUpdated(opts *bind.FilterOpts) (*CuratedModuleNodeOperatorDepositInfoFullyUpdatedIterator, error) {

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "NodeOperatorDepositInfoFullyUpdated")
	if err != nil {
		return nil, err
	}
	return &CuratedModuleNodeOperatorDepositInfoFullyUpdatedIterator{contract: _CuratedModule.contract, event: "NodeOperatorDepositInfoFullyUpdated", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorDepositInfoFullyUpdated is a free log subscription operation binding the contract event 0x2326ef260a3c0b1ad432887941d7041ce2672d2b8a3b87d2f5a9ff9f674ee35a.
//
// Solidity: event NodeOperatorDepositInfoFullyUpdated()
func (_CuratedModule *CuratedModuleFilterer) WatchNodeOperatorDepositInfoFullyUpdated(opts *bind.WatchOpts, sink chan<- *CuratedModuleNodeOperatorDepositInfoFullyUpdated) (event.Subscription, error) {

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "NodeOperatorDepositInfoFullyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleNodeOperatorDepositInfoFullyUpdated)
				if err := _CuratedModule.contract.UnpackLog(event, "NodeOperatorDepositInfoFullyUpdated", log); err != nil {
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

// ParseNodeOperatorDepositInfoFullyUpdated is a log parse operation binding the contract event 0x2326ef260a3c0b1ad432887941d7041ce2672d2b8a3b87d2f5a9ff9f674ee35a.
//
// Solidity: event NodeOperatorDepositInfoFullyUpdated()
func (_CuratedModule *CuratedModuleFilterer) ParseNodeOperatorDepositInfoFullyUpdated(log types.Log) (*CuratedModuleNodeOperatorDepositInfoFullyUpdated, error) {
	event := new(CuratedModuleNodeOperatorDepositInfoFullyUpdated)
	if err := _CuratedModule.contract.UnpackLog(event, "NodeOperatorDepositInfoFullyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleNodeOperatorManagerAddressChangeProposedIterator is returned from FilterNodeOperatorManagerAddressChangeProposed and is used to iterate over the raw logs and unpacked data for NodeOperatorManagerAddressChangeProposed events raised by the CuratedModule contract.
type CuratedModuleNodeOperatorManagerAddressChangeProposedIterator struct {
	Event *CuratedModuleNodeOperatorManagerAddressChangeProposed // Event containing the contract specifics and raw log

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
func (it *CuratedModuleNodeOperatorManagerAddressChangeProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleNodeOperatorManagerAddressChangeProposed)
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
		it.Event = new(CuratedModuleNodeOperatorManagerAddressChangeProposed)
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
func (it *CuratedModuleNodeOperatorManagerAddressChangeProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleNodeOperatorManagerAddressChangeProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleNodeOperatorManagerAddressChangeProposed represents a NodeOperatorManagerAddressChangeProposed event raised by the CuratedModule contract.
type CuratedModuleNodeOperatorManagerAddressChangeProposed struct {
	NodeOperatorId     *big.Int
	OldProposedAddress common.Address
	NewProposedAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorManagerAddressChangeProposed is a free log retrieval operation binding the contract event 0x4048f15a706950765ca59f99d0fa6fe8edaaa3f3e3d0337417082e2131df82fb.
//
// Solidity: event NodeOperatorManagerAddressChangeProposed(uint256 indexed nodeOperatorId, address indexed oldProposedAddress, address indexed newProposedAddress)
func (_CuratedModule *CuratedModuleFilterer) FilterNodeOperatorManagerAddressChangeProposed(opts *bind.FilterOpts, nodeOperatorId []*big.Int, oldProposedAddress []common.Address, newProposedAddress []common.Address) (*CuratedModuleNodeOperatorManagerAddressChangeProposedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var oldProposedAddressRule []interface{}
	for _, oldProposedAddressItem := range oldProposedAddress {
		oldProposedAddressRule = append(oldProposedAddressRule, oldProposedAddressItem)
	}
	var newProposedAddressRule []interface{}
	for _, newProposedAddressItem := range newProposedAddress {
		newProposedAddressRule = append(newProposedAddressRule, newProposedAddressItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "NodeOperatorManagerAddressChangeProposed", nodeOperatorIdRule, oldProposedAddressRule, newProposedAddressRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleNodeOperatorManagerAddressChangeProposedIterator{contract: _CuratedModule.contract, event: "NodeOperatorManagerAddressChangeProposed", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorManagerAddressChangeProposed is a free log subscription operation binding the contract event 0x4048f15a706950765ca59f99d0fa6fe8edaaa3f3e3d0337417082e2131df82fb.
//
// Solidity: event NodeOperatorManagerAddressChangeProposed(uint256 indexed nodeOperatorId, address indexed oldProposedAddress, address indexed newProposedAddress)
func (_CuratedModule *CuratedModuleFilterer) WatchNodeOperatorManagerAddressChangeProposed(opts *bind.WatchOpts, sink chan<- *CuratedModuleNodeOperatorManagerAddressChangeProposed, nodeOperatorId []*big.Int, oldProposedAddress []common.Address, newProposedAddress []common.Address) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var oldProposedAddressRule []interface{}
	for _, oldProposedAddressItem := range oldProposedAddress {
		oldProposedAddressRule = append(oldProposedAddressRule, oldProposedAddressItem)
	}
	var newProposedAddressRule []interface{}
	for _, newProposedAddressItem := range newProposedAddress {
		newProposedAddressRule = append(newProposedAddressRule, newProposedAddressItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "NodeOperatorManagerAddressChangeProposed", nodeOperatorIdRule, oldProposedAddressRule, newProposedAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleNodeOperatorManagerAddressChangeProposed)
				if err := _CuratedModule.contract.UnpackLog(event, "NodeOperatorManagerAddressChangeProposed", log); err != nil {
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

// ParseNodeOperatorManagerAddressChangeProposed is a log parse operation binding the contract event 0x4048f15a706950765ca59f99d0fa6fe8edaaa3f3e3d0337417082e2131df82fb.
//
// Solidity: event NodeOperatorManagerAddressChangeProposed(uint256 indexed nodeOperatorId, address indexed oldProposedAddress, address indexed newProposedAddress)
func (_CuratedModule *CuratedModuleFilterer) ParseNodeOperatorManagerAddressChangeProposed(log types.Log) (*CuratedModuleNodeOperatorManagerAddressChangeProposed, error) {
	event := new(CuratedModuleNodeOperatorManagerAddressChangeProposed)
	if err := _CuratedModule.contract.UnpackLog(event, "NodeOperatorManagerAddressChangeProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleNodeOperatorManagerAddressChangedIterator is returned from FilterNodeOperatorManagerAddressChanged and is used to iterate over the raw logs and unpacked data for NodeOperatorManagerAddressChanged events raised by the CuratedModule contract.
type CuratedModuleNodeOperatorManagerAddressChangedIterator struct {
	Event *CuratedModuleNodeOperatorManagerAddressChanged // Event containing the contract specifics and raw log

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
func (it *CuratedModuleNodeOperatorManagerAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleNodeOperatorManagerAddressChanged)
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
		it.Event = new(CuratedModuleNodeOperatorManagerAddressChanged)
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
func (it *CuratedModuleNodeOperatorManagerAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleNodeOperatorManagerAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleNodeOperatorManagerAddressChanged represents a NodeOperatorManagerAddressChanged event raised by the CuratedModule contract.
type CuratedModuleNodeOperatorManagerAddressChanged struct {
	NodeOperatorId *big.Int
	OldAddress     common.Address
	NewAddress     common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorManagerAddressChanged is a free log retrieval operation binding the contract event 0x862021f23449d6e8516867bd839be15a3d8698a7561c5c2c35069074b7e91e61.
//
// Solidity: event NodeOperatorManagerAddressChanged(uint256 indexed nodeOperatorId, address indexed oldAddress, address indexed newAddress)
func (_CuratedModule *CuratedModuleFilterer) FilterNodeOperatorManagerAddressChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int, oldAddress []common.Address, newAddress []common.Address) (*CuratedModuleNodeOperatorManagerAddressChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "NodeOperatorManagerAddressChanged", nodeOperatorIdRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleNodeOperatorManagerAddressChangedIterator{contract: _CuratedModule.contract, event: "NodeOperatorManagerAddressChanged", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorManagerAddressChanged is a free log subscription operation binding the contract event 0x862021f23449d6e8516867bd839be15a3d8698a7561c5c2c35069074b7e91e61.
//
// Solidity: event NodeOperatorManagerAddressChanged(uint256 indexed nodeOperatorId, address indexed oldAddress, address indexed newAddress)
func (_CuratedModule *CuratedModuleFilterer) WatchNodeOperatorManagerAddressChanged(opts *bind.WatchOpts, sink chan<- *CuratedModuleNodeOperatorManagerAddressChanged, nodeOperatorId []*big.Int, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "NodeOperatorManagerAddressChanged", nodeOperatorIdRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleNodeOperatorManagerAddressChanged)
				if err := _CuratedModule.contract.UnpackLog(event, "NodeOperatorManagerAddressChanged", log); err != nil {
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

// ParseNodeOperatorManagerAddressChanged is a log parse operation binding the contract event 0x862021f23449d6e8516867bd839be15a3d8698a7561c5c2c35069074b7e91e61.
//
// Solidity: event NodeOperatorManagerAddressChanged(uint256 indexed nodeOperatorId, address indexed oldAddress, address indexed newAddress)
func (_CuratedModule *CuratedModuleFilterer) ParseNodeOperatorManagerAddressChanged(log types.Log) (*CuratedModuleNodeOperatorManagerAddressChanged, error) {
	event := new(CuratedModuleNodeOperatorManagerAddressChanged)
	if err := _CuratedModule.contract.UnpackLog(event, "NodeOperatorManagerAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleNodeOperatorRewardAddressChangeProposedIterator is returned from FilterNodeOperatorRewardAddressChangeProposed and is used to iterate over the raw logs and unpacked data for NodeOperatorRewardAddressChangeProposed events raised by the CuratedModule contract.
type CuratedModuleNodeOperatorRewardAddressChangeProposedIterator struct {
	Event *CuratedModuleNodeOperatorRewardAddressChangeProposed // Event containing the contract specifics and raw log

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
func (it *CuratedModuleNodeOperatorRewardAddressChangeProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleNodeOperatorRewardAddressChangeProposed)
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
		it.Event = new(CuratedModuleNodeOperatorRewardAddressChangeProposed)
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
func (it *CuratedModuleNodeOperatorRewardAddressChangeProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleNodeOperatorRewardAddressChangeProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleNodeOperatorRewardAddressChangeProposed represents a NodeOperatorRewardAddressChangeProposed event raised by the CuratedModule contract.
type CuratedModuleNodeOperatorRewardAddressChangeProposed struct {
	NodeOperatorId     *big.Int
	OldProposedAddress common.Address
	NewProposedAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorRewardAddressChangeProposed is a free log retrieval operation binding the contract event 0xb5878cdb1d66f971efe3b138a71c64bc5bc519314db2533e0e4cde954409ea5a.
//
// Solidity: event NodeOperatorRewardAddressChangeProposed(uint256 indexed nodeOperatorId, address indexed oldProposedAddress, address indexed newProposedAddress)
func (_CuratedModule *CuratedModuleFilterer) FilterNodeOperatorRewardAddressChangeProposed(opts *bind.FilterOpts, nodeOperatorId []*big.Int, oldProposedAddress []common.Address, newProposedAddress []common.Address) (*CuratedModuleNodeOperatorRewardAddressChangeProposedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var oldProposedAddressRule []interface{}
	for _, oldProposedAddressItem := range oldProposedAddress {
		oldProposedAddressRule = append(oldProposedAddressRule, oldProposedAddressItem)
	}
	var newProposedAddressRule []interface{}
	for _, newProposedAddressItem := range newProposedAddress {
		newProposedAddressRule = append(newProposedAddressRule, newProposedAddressItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "NodeOperatorRewardAddressChangeProposed", nodeOperatorIdRule, oldProposedAddressRule, newProposedAddressRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleNodeOperatorRewardAddressChangeProposedIterator{contract: _CuratedModule.contract, event: "NodeOperatorRewardAddressChangeProposed", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorRewardAddressChangeProposed is a free log subscription operation binding the contract event 0xb5878cdb1d66f971efe3b138a71c64bc5bc519314db2533e0e4cde954409ea5a.
//
// Solidity: event NodeOperatorRewardAddressChangeProposed(uint256 indexed nodeOperatorId, address indexed oldProposedAddress, address indexed newProposedAddress)
func (_CuratedModule *CuratedModuleFilterer) WatchNodeOperatorRewardAddressChangeProposed(opts *bind.WatchOpts, sink chan<- *CuratedModuleNodeOperatorRewardAddressChangeProposed, nodeOperatorId []*big.Int, oldProposedAddress []common.Address, newProposedAddress []common.Address) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var oldProposedAddressRule []interface{}
	for _, oldProposedAddressItem := range oldProposedAddress {
		oldProposedAddressRule = append(oldProposedAddressRule, oldProposedAddressItem)
	}
	var newProposedAddressRule []interface{}
	for _, newProposedAddressItem := range newProposedAddress {
		newProposedAddressRule = append(newProposedAddressRule, newProposedAddressItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "NodeOperatorRewardAddressChangeProposed", nodeOperatorIdRule, oldProposedAddressRule, newProposedAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleNodeOperatorRewardAddressChangeProposed)
				if err := _CuratedModule.contract.UnpackLog(event, "NodeOperatorRewardAddressChangeProposed", log); err != nil {
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

// ParseNodeOperatorRewardAddressChangeProposed is a log parse operation binding the contract event 0xb5878cdb1d66f971efe3b138a71c64bc5bc519314db2533e0e4cde954409ea5a.
//
// Solidity: event NodeOperatorRewardAddressChangeProposed(uint256 indexed nodeOperatorId, address indexed oldProposedAddress, address indexed newProposedAddress)
func (_CuratedModule *CuratedModuleFilterer) ParseNodeOperatorRewardAddressChangeProposed(log types.Log) (*CuratedModuleNodeOperatorRewardAddressChangeProposed, error) {
	event := new(CuratedModuleNodeOperatorRewardAddressChangeProposed)
	if err := _CuratedModule.contract.UnpackLog(event, "NodeOperatorRewardAddressChangeProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleNodeOperatorRewardAddressChangedIterator is returned from FilterNodeOperatorRewardAddressChanged and is used to iterate over the raw logs and unpacked data for NodeOperatorRewardAddressChanged events raised by the CuratedModule contract.
type CuratedModuleNodeOperatorRewardAddressChangedIterator struct {
	Event *CuratedModuleNodeOperatorRewardAddressChanged // Event containing the contract specifics and raw log

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
func (it *CuratedModuleNodeOperatorRewardAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleNodeOperatorRewardAddressChanged)
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
		it.Event = new(CuratedModuleNodeOperatorRewardAddressChanged)
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
func (it *CuratedModuleNodeOperatorRewardAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleNodeOperatorRewardAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleNodeOperatorRewardAddressChanged represents a NodeOperatorRewardAddressChanged event raised by the CuratedModule contract.
type CuratedModuleNodeOperatorRewardAddressChanged struct {
	NodeOperatorId *big.Int
	OldAddress     common.Address
	NewAddress     common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorRewardAddressChanged is a free log retrieval operation binding the contract event 0x069ac7cd8230db015b7250c8e5425149cf1a3e912d9569f497165e55b3b6b7b2.
//
// Solidity: event NodeOperatorRewardAddressChanged(uint256 indexed nodeOperatorId, address indexed oldAddress, address indexed newAddress)
func (_CuratedModule *CuratedModuleFilterer) FilterNodeOperatorRewardAddressChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int, oldAddress []common.Address, newAddress []common.Address) (*CuratedModuleNodeOperatorRewardAddressChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "NodeOperatorRewardAddressChanged", nodeOperatorIdRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleNodeOperatorRewardAddressChangedIterator{contract: _CuratedModule.contract, event: "NodeOperatorRewardAddressChanged", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorRewardAddressChanged is a free log subscription operation binding the contract event 0x069ac7cd8230db015b7250c8e5425149cf1a3e912d9569f497165e55b3b6b7b2.
//
// Solidity: event NodeOperatorRewardAddressChanged(uint256 indexed nodeOperatorId, address indexed oldAddress, address indexed newAddress)
func (_CuratedModule *CuratedModuleFilterer) WatchNodeOperatorRewardAddressChanged(opts *bind.WatchOpts, sink chan<- *CuratedModuleNodeOperatorRewardAddressChanged, nodeOperatorId []*big.Int, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "NodeOperatorRewardAddressChanged", nodeOperatorIdRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleNodeOperatorRewardAddressChanged)
				if err := _CuratedModule.contract.UnpackLog(event, "NodeOperatorRewardAddressChanged", log); err != nil {
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

// ParseNodeOperatorRewardAddressChanged is a log parse operation binding the contract event 0x069ac7cd8230db015b7250c8e5425149cf1a3e912d9569f497165e55b3b6b7b2.
//
// Solidity: event NodeOperatorRewardAddressChanged(uint256 indexed nodeOperatorId, address indexed oldAddress, address indexed newAddress)
func (_CuratedModule *CuratedModuleFilterer) ParseNodeOperatorRewardAddressChanged(log types.Log) (*CuratedModuleNodeOperatorRewardAddressChanged, error) {
	event := new(CuratedModuleNodeOperatorRewardAddressChanged)
	if err := _CuratedModule.contract.UnpackLog(event, "NodeOperatorRewardAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleNonceChangedIterator is returned from FilterNonceChanged and is used to iterate over the raw logs and unpacked data for NonceChanged events raised by the CuratedModule contract.
type CuratedModuleNonceChangedIterator struct {
	Event *CuratedModuleNonceChanged // Event containing the contract specifics and raw log

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
func (it *CuratedModuleNonceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleNonceChanged)
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
		it.Event = new(CuratedModuleNonceChanged)
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
func (it *CuratedModuleNonceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleNonceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleNonceChanged represents a NonceChanged event raised by the CuratedModule contract.
type CuratedModuleNonceChanged struct {
	Nonce *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNonceChanged is a free log retrieval operation binding the contract event 0x7220970e1f1f12864ecccd8942690a837c7a8dd45d158cb891eb45a8a69134aa.
//
// Solidity: event NonceChanged(uint256 nonce)
func (_CuratedModule *CuratedModuleFilterer) FilterNonceChanged(opts *bind.FilterOpts) (*CuratedModuleNonceChangedIterator, error) {

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "NonceChanged")
	if err != nil {
		return nil, err
	}
	return &CuratedModuleNonceChangedIterator{contract: _CuratedModule.contract, event: "NonceChanged", logs: logs, sub: sub}, nil
}

// WatchNonceChanged is a free log subscription operation binding the contract event 0x7220970e1f1f12864ecccd8942690a837c7a8dd45d158cb891eb45a8a69134aa.
//
// Solidity: event NonceChanged(uint256 nonce)
func (_CuratedModule *CuratedModuleFilterer) WatchNonceChanged(opts *bind.WatchOpts, sink chan<- *CuratedModuleNonceChanged) (event.Subscription, error) {

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "NonceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleNonceChanged)
				if err := _CuratedModule.contract.UnpackLog(event, "NonceChanged", log); err != nil {
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
func (_CuratedModule *CuratedModuleFilterer) ParseNonceChanged(log types.Log) (*CuratedModuleNonceChanged, error) {
	event := new(CuratedModuleNonceChanged)
	if err := _CuratedModule.contract.UnpackLog(event, "NonceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModulePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CuratedModule contract.
type CuratedModulePausedIterator struct {
	Event *CuratedModulePaused // Event containing the contract specifics and raw log

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
func (it *CuratedModulePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModulePaused)
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
		it.Event = new(CuratedModulePaused)
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
func (it *CuratedModulePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModulePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModulePaused represents a Paused event raised by the CuratedModule contract.
type CuratedModulePaused struct {
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_CuratedModule *CuratedModuleFilterer) FilterPaused(opts *bind.FilterOpts) (*CuratedModulePausedIterator, error) {

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CuratedModulePausedIterator{contract: _CuratedModule.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_CuratedModule *CuratedModuleFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CuratedModulePaused) (event.Subscription, error) {

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModulePaused)
				if err := _CuratedModule.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_CuratedModule *CuratedModuleFilterer) ParsePaused(log types.Log) (*CuratedModulePaused, error) {
	event := new(CuratedModulePaused)
	if err := _CuratedModule.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleReferrerSetIterator is returned from FilterReferrerSet and is used to iterate over the raw logs and unpacked data for ReferrerSet events raised by the CuratedModule contract.
type CuratedModuleReferrerSetIterator struct {
	Event *CuratedModuleReferrerSet // Event containing the contract specifics and raw log

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
func (it *CuratedModuleReferrerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleReferrerSet)
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
		it.Event = new(CuratedModuleReferrerSet)
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
func (it *CuratedModuleReferrerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleReferrerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleReferrerSet represents a ReferrerSet event raised by the CuratedModule contract.
type CuratedModuleReferrerSet struct {
	NodeOperatorId *big.Int
	Referrer       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterReferrerSet is a free log retrieval operation binding the contract event 0x67334334c388385e5f244703f8a8b28b7f4ffe52909130aca69bc62a8e27f09a.
//
// Solidity: event ReferrerSet(uint256 indexed nodeOperatorId, address indexed referrer)
func (_CuratedModule *CuratedModuleFilterer) FilterReferrerSet(opts *bind.FilterOpts, nodeOperatorId []*big.Int, referrer []common.Address) (*CuratedModuleReferrerSetIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var referrerRule []interface{}
	for _, referrerItem := range referrer {
		referrerRule = append(referrerRule, referrerItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "ReferrerSet", nodeOperatorIdRule, referrerRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleReferrerSetIterator{contract: _CuratedModule.contract, event: "ReferrerSet", logs: logs, sub: sub}, nil
}

// WatchReferrerSet is a free log subscription operation binding the contract event 0x67334334c388385e5f244703f8a8b28b7f4ffe52909130aca69bc62a8e27f09a.
//
// Solidity: event ReferrerSet(uint256 indexed nodeOperatorId, address indexed referrer)
func (_CuratedModule *CuratedModuleFilterer) WatchReferrerSet(opts *bind.WatchOpts, sink chan<- *CuratedModuleReferrerSet, nodeOperatorId []*big.Int, referrer []common.Address) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var referrerRule []interface{}
	for _, referrerItem := range referrer {
		referrerRule = append(referrerRule, referrerItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "ReferrerSet", nodeOperatorIdRule, referrerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleReferrerSet)
				if err := _CuratedModule.contract.UnpackLog(event, "ReferrerSet", log); err != nil {
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

// ParseReferrerSet is a log parse operation binding the contract event 0x67334334c388385e5f244703f8a8b28b7f4ffe52909130aca69bc62a8e27f09a.
//
// Solidity: event ReferrerSet(uint256 indexed nodeOperatorId, address indexed referrer)
func (_CuratedModule *CuratedModuleFilterer) ParseReferrerSet(log types.Log) (*CuratedModuleReferrerSet, error) {
	event := new(CuratedModuleReferrerSet)
	if err := _CuratedModule.contract.UnpackLog(event, "ReferrerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleResumedIterator is returned from FilterResumed and is used to iterate over the raw logs and unpacked data for Resumed events raised by the CuratedModule contract.
type CuratedModuleResumedIterator struct {
	Event *CuratedModuleResumed // Event containing the contract specifics and raw log

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
func (it *CuratedModuleResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleResumed)
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
		it.Event = new(CuratedModuleResumed)
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
func (it *CuratedModuleResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleResumed represents a Resumed event raised by the CuratedModule contract.
type CuratedModuleResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResumed is a free log retrieval operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_CuratedModule *CuratedModuleFilterer) FilterResumed(opts *bind.FilterOpts) (*CuratedModuleResumedIterator, error) {

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return &CuratedModuleResumedIterator{contract: _CuratedModule.contract, event: "Resumed", logs: logs, sub: sub}, nil
}

// WatchResumed is a free log subscription operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_CuratedModule *CuratedModuleFilterer) WatchResumed(opts *bind.WatchOpts, sink chan<- *CuratedModuleResumed) (event.Subscription, error) {

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleResumed)
				if err := _CuratedModule.contract.UnpackLog(event, "Resumed", log); err != nil {
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

// ParseResumed is a log parse operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_CuratedModule *CuratedModuleFilterer) ParseResumed(log types.Log) (*CuratedModuleResumed, error) {
	event := new(CuratedModuleResumed)
	if err := _CuratedModule.contract.UnpackLog(event, "Resumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the CuratedModule contract.
type CuratedModuleRoleAdminChangedIterator struct {
	Event *CuratedModuleRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CuratedModuleRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleRoleAdminChanged)
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
		it.Event = new(CuratedModuleRoleAdminChanged)
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
func (it *CuratedModuleRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleRoleAdminChanged represents a RoleAdminChanged event raised by the CuratedModule contract.
type CuratedModuleRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CuratedModule *CuratedModuleFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CuratedModuleRoleAdminChangedIterator, error) {

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

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleRoleAdminChangedIterator{contract: _CuratedModule.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CuratedModule *CuratedModuleFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CuratedModuleRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleRoleAdminChanged)
				if err := _CuratedModule.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_CuratedModule *CuratedModuleFilterer) ParseRoleAdminChanged(log types.Log) (*CuratedModuleRoleAdminChanged, error) {
	event := new(CuratedModuleRoleAdminChanged)
	if err := _CuratedModule.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the CuratedModule contract.
type CuratedModuleRoleGrantedIterator struct {
	Event *CuratedModuleRoleGranted // Event containing the contract specifics and raw log

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
func (it *CuratedModuleRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleRoleGranted)
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
		it.Event = new(CuratedModuleRoleGranted)
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
func (it *CuratedModuleRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleRoleGranted represents a RoleGranted event raised by the CuratedModule contract.
type CuratedModuleRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CuratedModule *CuratedModuleFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CuratedModuleRoleGrantedIterator, error) {

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

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleRoleGrantedIterator{contract: _CuratedModule.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CuratedModule *CuratedModuleFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CuratedModuleRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleRoleGranted)
				if err := _CuratedModule.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_CuratedModule *CuratedModuleFilterer) ParseRoleGranted(log types.Log) (*CuratedModuleRoleGranted, error) {
	event := new(CuratedModuleRoleGranted)
	if err := _CuratedModule.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the CuratedModule contract.
type CuratedModuleRoleRevokedIterator struct {
	Event *CuratedModuleRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CuratedModuleRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleRoleRevoked)
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
		it.Event = new(CuratedModuleRoleRevoked)
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
func (it *CuratedModuleRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleRoleRevoked represents a RoleRevoked event raised by the CuratedModule contract.
type CuratedModuleRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CuratedModule *CuratedModuleFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CuratedModuleRoleRevokedIterator, error) {

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

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleRoleRevokedIterator{contract: _CuratedModule.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CuratedModule *CuratedModuleFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CuratedModuleRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleRoleRevoked)
				if err := _CuratedModule.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_CuratedModule *CuratedModuleFilterer) ParseRoleRevoked(log types.Log) (*CuratedModuleRoleRevoked, error) {
	event := new(CuratedModuleRoleRevoked)
	if err := _CuratedModule.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleSigningKeyAddedIterator is returned from FilterSigningKeyAdded and is used to iterate over the raw logs and unpacked data for SigningKeyAdded events raised by the CuratedModule contract.
type CuratedModuleSigningKeyAddedIterator struct {
	Event *CuratedModuleSigningKeyAdded // Event containing the contract specifics and raw log

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
func (it *CuratedModuleSigningKeyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleSigningKeyAdded)
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
		it.Event = new(CuratedModuleSigningKeyAdded)
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
func (it *CuratedModuleSigningKeyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleSigningKeyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleSigningKeyAdded represents a SigningKeyAdded event raised by the CuratedModule contract.
type CuratedModuleSigningKeyAdded struct {
	NodeOperatorId *big.Int
	Pubkey         []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSigningKeyAdded is a free log retrieval operation binding the contract event 0xc77a17d6b857abe6d6e6c37301621bc72c4dd52fa8830fb54dfa715c04911a89.
//
// Solidity: event SigningKeyAdded(uint256 indexed nodeOperatorId, bytes pubkey)
func (_CuratedModule *CuratedModuleFilterer) FilterSigningKeyAdded(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CuratedModuleSigningKeyAddedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "SigningKeyAdded", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleSigningKeyAddedIterator{contract: _CuratedModule.contract, event: "SigningKeyAdded", logs: logs, sub: sub}, nil
}

// WatchSigningKeyAdded is a free log subscription operation binding the contract event 0xc77a17d6b857abe6d6e6c37301621bc72c4dd52fa8830fb54dfa715c04911a89.
//
// Solidity: event SigningKeyAdded(uint256 indexed nodeOperatorId, bytes pubkey)
func (_CuratedModule *CuratedModuleFilterer) WatchSigningKeyAdded(opts *bind.WatchOpts, sink chan<- *CuratedModuleSigningKeyAdded, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "SigningKeyAdded", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleSigningKeyAdded)
				if err := _CuratedModule.contract.UnpackLog(event, "SigningKeyAdded", log); err != nil {
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

// ParseSigningKeyAdded is a log parse operation binding the contract event 0xc77a17d6b857abe6d6e6c37301621bc72c4dd52fa8830fb54dfa715c04911a89.
//
// Solidity: event SigningKeyAdded(uint256 indexed nodeOperatorId, bytes pubkey)
func (_CuratedModule *CuratedModuleFilterer) ParseSigningKeyAdded(log types.Log) (*CuratedModuleSigningKeyAdded, error) {
	event := new(CuratedModuleSigningKeyAdded)
	if err := _CuratedModule.contract.UnpackLog(event, "SigningKeyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleSigningKeyRemovedIterator is returned from FilterSigningKeyRemoved and is used to iterate over the raw logs and unpacked data for SigningKeyRemoved events raised by the CuratedModule contract.
type CuratedModuleSigningKeyRemovedIterator struct {
	Event *CuratedModuleSigningKeyRemoved // Event containing the contract specifics and raw log

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
func (it *CuratedModuleSigningKeyRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleSigningKeyRemoved)
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
		it.Event = new(CuratedModuleSigningKeyRemoved)
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
func (it *CuratedModuleSigningKeyRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleSigningKeyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleSigningKeyRemoved represents a SigningKeyRemoved event raised by the CuratedModule contract.
type CuratedModuleSigningKeyRemoved struct {
	NodeOperatorId *big.Int
	Pubkey         []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSigningKeyRemoved is a free log retrieval operation binding the contract event 0xea4b75aaf57196f73d338cadf79ecd0a437902e2dd0d2c4c2cf3ea71b8ab27b9.
//
// Solidity: event SigningKeyRemoved(uint256 indexed nodeOperatorId, bytes pubkey)
func (_CuratedModule *CuratedModuleFilterer) FilterSigningKeyRemoved(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CuratedModuleSigningKeyRemovedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "SigningKeyRemoved", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleSigningKeyRemovedIterator{contract: _CuratedModule.contract, event: "SigningKeyRemoved", logs: logs, sub: sub}, nil
}

// WatchSigningKeyRemoved is a free log subscription operation binding the contract event 0xea4b75aaf57196f73d338cadf79ecd0a437902e2dd0d2c4c2cf3ea71b8ab27b9.
//
// Solidity: event SigningKeyRemoved(uint256 indexed nodeOperatorId, bytes pubkey)
func (_CuratedModule *CuratedModuleFilterer) WatchSigningKeyRemoved(opts *bind.WatchOpts, sink chan<- *CuratedModuleSigningKeyRemoved, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "SigningKeyRemoved", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleSigningKeyRemoved)
				if err := _CuratedModule.contract.UnpackLog(event, "SigningKeyRemoved", log); err != nil {
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

// ParseSigningKeyRemoved is a log parse operation binding the contract event 0xea4b75aaf57196f73d338cadf79ecd0a437902e2dd0d2c4c2cf3ea71b8ab27b9.
//
// Solidity: event SigningKeyRemoved(uint256 indexed nodeOperatorId, bytes pubkey)
func (_CuratedModule *CuratedModuleFilterer) ParseSigningKeyRemoved(log types.Log) (*CuratedModuleSigningKeyRemoved, error) {
	event := new(CuratedModuleSigningKeyRemoved)
	if err := _CuratedModule.contract.UnpackLog(event, "SigningKeyRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleStETHSharesRecoveredIterator is returned from FilterStETHSharesRecovered and is used to iterate over the raw logs and unpacked data for StETHSharesRecovered events raised by the CuratedModule contract.
type CuratedModuleStETHSharesRecoveredIterator struct {
	Event *CuratedModuleStETHSharesRecovered // Event containing the contract specifics and raw log

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
func (it *CuratedModuleStETHSharesRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleStETHSharesRecovered)
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
		it.Event = new(CuratedModuleStETHSharesRecovered)
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
func (it *CuratedModuleStETHSharesRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleStETHSharesRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleStETHSharesRecovered represents a StETHSharesRecovered event raised by the CuratedModule contract.
type CuratedModuleStETHSharesRecovered struct {
	Recipient common.Address
	Shares    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStETHSharesRecovered is a free log retrieval operation binding the contract event 0x426e7e0100db57255d4af4a46cd49552ef74f5f002bbdc8d4ebb6371c0070a02.
//
// Solidity: event StETHSharesRecovered(address indexed recipient, uint256 shares)
func (_CuratedModule *CuratedModuleFilterer) FilterStETHSharesRecovered(opts *bind.FilterOpts, recipient []common.Address) (*CuratedModuleStETHSharesRecoveredIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "StETHSharesRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleStETHSharesRecoveredIterator{contract: _CuratedModule.contract, event: "StETHSharesRecovered", logs: logs, sub: sub}, nil
}

// WatchStETHSharesRecovered is a free log subscription operation binding the contract event 0x426e7e0100db57255d4af4a46cd49552ef74f5f002bbdc8d4ebb6371c0070a02.
//
// Solidity: event StETHSharesRecovered(address indexed recipient, uint256 shares)
func (_CuratedModule *CuratedModuleFilterer) WatchStETHSharesRecovered(opts *bind.WatchOpts, sink chan<- *CuratedModuleStETHSharesRecovered, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "StETHSharesRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleStETHSharesRecovered)
				if err := _CuratedModule.contract.UnpackLog(event, "StETHSharesRecovered", log); err != nil {
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

// ParseStETHSharesRecovered is a log parse operation binding the contract event 0x426e7e0100db57255d4af4a46cd49552ef74f5f002bbdc8d4ebb6371c0070a02.
//
// Solidity: event StETHSharesRecovered(address indexed recipient, uint256 shares)
func (_CuratedModule *CuratedModuleFilterer) ParseStETHSharesRecovered(log types.Log) (*CuratedModuleStETHSharesRecovered, error) {
	event := new(CuratedModuleStETHSharesRecovered)
	if err := _CuratedModule.contract.UnpackLog(event, "StETHSharesRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleTargetValidatorsCountChangedIterator is returned from FilterTargetValidatorsCountChanged and is used to iterate over the raw logs and unpacked data for TargetValidatorsCountChanged events raised by the CuratedModule contract.
type CuratedModuleTargetValidatorsCountChangedIterator struct {
	Event *CuratedModuleTargetValidatorsCountChanged // Event containing the contract specifics and raw log

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
func (it *CuratedModuleTargetValidatorsCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleTargetValidatorsCountChanged)
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
		it.Event = new(CuratedModuleTargetValidatorsCountChanged)
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
func (it *CuratedModuleTargetValidatorsCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleTargetValidatorsCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleTargetValidatorsCountChanged represents a TargetValidatorsCountChanged event raised by the CuratedModule contract.
type CuratedModuleTargetValidatorsCountChanged struct {
	NodeOperatorId        *big.Int
	TargetLimitMode       *big.Int
	TargetValidatorsCount *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterTargetValidatorsCountChanged is a free log retrieval operation binding the contract event 0xf92eb109ce5b449e9b121c352c6aeb4319538a90738cb95d84f08e41274e92d2.
//
// Solidity: event TargetValidatorsCountChanged(uint256 indexed nodeOperatorId, uint256 targetLimitMode, uint256 targetValidatorsCount)
func (_CuratedModule *CuratedModuleFilterer) FilterTargetValidatorsCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CuratedModuleTargetValidatorsCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "TargetValidatorsCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleTargetValidatorsCountChangedIterator{contract: _CuratedModule.contract, event: "TargetValidatorsCountChanged", logs: logs, sub: sub}, nil
}

// WatchTargetValidatorsCountChanged is a free log subscription operation binding the contract event 0xf92eb109ce5b449e9b121c352c6aeb4319538a90738cb95d84f08e41274e92d2.
//
// Solidity: event TargetValidatorsCountChanged(uint256 indexed nodeOperatorId, uint256 targetLimitMode, uint256 targetValidatorsCount)
func (_CuratedModule *CuratedModuleFilterer) WatchTargetValidatorsCountChanged(opts *bind.WatchOpts, sink chan<- *CuratedModuleTargetValidatorsCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "TargetValidatorsCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleTargetValidatorsCountChanged)
				if err := _CuratedModule.contract.UnpackLog(event, "TargetValidatorsCountChanged", log); err != nil {
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

// ParseTargetValidatorsCountChanged is a log parse operation binding the contract event 0xf92eb109ce5b449e9b121c352c6aeb4319538a90738cb95d84f08e41274e92d2.
//
// Solidity: event TargetValidatorsCountChanged(uint256 indexed nodeOperatorId, uint256 targetLimitMode, uint256 targetValidatorsCount)
func (_CuratedModule *CuratedModuleFilterer) ParseTargetValidatorsCountChanged(log types.Log) (*CuratedModuleTargetValidatorsCountChanged, error) {
	event := new(CuratedModuleTargetValidatorsCountChanged)
	if err := _CuratedModule.contract.UnpackLog(event, "TargetValidatorsCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleTotalSigningKeysCountChangedIterator is returned from FilterTotalSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for TotalSigningKeysCountChanged events raised by the CuratedModule contract.
type CuratedModuleTotalSigningKeysCountChangedIterator struct {
	Event *CuratedModuleTotalSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *CuratedModuleTotalSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleTotalSigningKeysCountChanged)
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
		it.Event = new(CuratedModuleTotalSigningKeysCountChanged)
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
func (it *CuratedModuleTotalSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleTotalSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleTotalSigningKeysCountChanged represents a TotalSigningKeysCountChanged event raised by the CuratedModule contract.
type CuratedModuleTotalSigningKeysCountChanged struct {
	NodeOperatorId *big.Int
	TotalKeysCount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTotalSigningKeysCountChanged is a free log retrieval operation binding the contract event 0xdd01838a366ae4dc9a86e1922512c0716abebc9a440baae0e22d2dec578223f0.
//
// Solidity: event TotalSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 totalKeysCount)
func (_CuratedModule *CuratedModuleFilterer) FilterTotalSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CuratedModuleTotalSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "TotalSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleTotalSigningKeysCountChangedIterator{contract: _CuratedModule.contract, event: "TotalSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchTotalSigningKeysCountChanged is a free log subscription operation binding the contract event 0xdd01838a366ae4dc9a86e1922512c0716abebc9a440baae0e22d2dec578223f0.
//
// Solidity: event TotalSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 totalKeysCount)
func (_CuratedModule *CuratedModuleFilterer) WatchTotalSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *CuratedModuleTotalSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "TotalSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleTotalSigningKeysCountChanged)
				if err := _CuratedModule.contract.UnpackLog(event, "TotalSigningKeysCountChanged", log); err != nil {
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
// Solidity: event TotalSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 totalKeysCount)
func (_CuratedModule *CuratedModuleFilterer) ParseTotalSigningKeysCountChanged(log types.Log) (*CuratedModuleTotalSigningKeysCountChanged, error) {
	event := new(CuratedModuleTotalSigningKeysCountChanged)
	if err := _CuratedModule.contract.UnpackLog(event, "TotalSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleTotalWithdrawnValidatorsRebuiltIterator is returned from FilterTotalWithdrawnValidatorsRebuilt and is used to iterate over the raw logs and unpacked data for TotalWithdrawnValidatorsRebuilt events raised by the CuratedModule contract.
type CuratedModuleTotalWithdrawnValidatorsRebuiltIterator struct {
	Event *CuratedModuleTotalWithdrawnValidatorsRebuilt // Event containing the contract specifics and raw log

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
func (it *CuratedModuleTotalWithdrawnValidatorsRebuiltIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleTotalWithdrawnValidatorsRebuilt)
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
		it.Event = new(CuratedModuleTotalWithdrawnValidatorsRebuilt)
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
func (it *CuratedModuleTotalWithdrawnValidatorsRebuiltIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleTotalWithdrawnValidatorsRebuiltIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleTotalWithdrawnValidatorsRebuilt represents a TotalWithdrawnValidatorsRebuilt event raised by the CuratedModule contract.
type CuratedModuleTotalWithdrawnValidatorsRebuilt struct {
	TotalWithdrawnValidators *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterTotalWithdrawnValidatorsRebuilt is a free log retrieval operation binding the contract event 0x71e2ce26210a7e784b7d32a343866872e308e56cb66d4e6a017cc0b4571edf4c.
//
// Solidity: event TotalWithdrawnValidatorsRebuilt(uint256 totalWithdrawnValidators)
func (_CuratedModule *CuratedModuleFilterer) FilterTotalWithdrawnValidatorsRebuilt(opts *bind.FilterOpts) (*CuratedModuleTotalWithdrawnValidatorsRebuiltIterator, error) {

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "TotalWithdrawnValidatorsRebuilt")
	if err != nil {
		return nil, err
	}
	return &CuratedModuleTotalWithdrawnValidatorsRebuiltIterator{contract: _CuratedModule.contract, event: "TotalWithdrawnValidatorsRebuilt", logs: logs, sub: sub}, nil
}

// WatchTotalWithdrawnValidatorsRebuilt is a free log subscription operation binding the contract event 0x71e2ce26210a7e784b7d32a343866872e308e56cb66d4e6a017cc0b4571edf4c.
//
// Solidity: event TotalWithdrawnValidatorsRebuilt(uint256 totalWithdrawnValidators)
func (_CuratedModule *CuratedModuleFilterer) WatchTotalWithdrawnValidatorsRebuilt(opts *bind.WatchOpts, sink chan<- *CuratedModuleTotalWithdrawnValidatorsRebuilt) (event.Subscription, error) {

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "TotalWithdrawnValidatorsRebuilt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleTotalWithdrawnValidatorsRebuilt)
				if err := _CuratedModule.contract.UnpackLog(event, "TotalWithdrawnValidatorsRebuilt", log); err != nil {
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

// ParseTotalWithdrawnValidatorsRebuilt is a log parse operation binding the contract event 0x71e2ce26210a7e784b7d32a343866872e308e56cb66d4e6a017cc0b4571edf4c.
//
// Solidity: event TotalWithdrawnValidatorsRebuilt(uint256 totalWithdrawnValidators)
func (_CuratedModule *CuratedModuleFilterer) ParseTotalWithdrawnValidatorsRebuilt(log types.Log) (*CuratedModuleTotalWithdrawnValidatorsRebuilt, error) {
	event := new(CuratedModuleTotalWithdrawnValidatorsRebuilt)
	if err := _CuratedModule.contract.UnpackLog(event, "TotalWithdrawnValidatorsRebuilt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleValidatorSlashingReportedIterator is returned from FilterValidatorSlashingReported and is used to iterate over the raw logs and unpacked data for ValidatorSlashingReported events raised by the CuratedModule contract.
type CuratedModuleValidatorSlashingReportedIterator struct {
	Event *CuratedModuleValidatorSlashingReported // Event containing the contract specifics and raw log

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
func (it *CuratedModuleValidatorSlashingReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleValidatorSlashingReported)
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
		it.Event = new(CuratedModuleValidatorSlashingReported)
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
func (it *CuratedModuleValidatorSlashingReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleValidatorSlashingReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleValidatorSlashingReported represents a ValidatorSlashingReported event raised by the CuratedModule contract.
type CuratedModuleValidatorSlashingReported struct {
	NodeOperatorId *big.Int
	KeyIndex       *big.Int
	Pubkey         []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValidatorSlashingReported is a free log retrieval operation binding the contract event 0x4e37fce6f335bba671f891995433c6d300e4384b81a98ebb32eb5cfd5122109c.
//
// Solidity: event ValidatorSlashingReported(uint256 indexed nodeOperatorId, uint256 keyIndex, bytes pubkey)
func (_CuratedModule *CuratedModuleFilterer) FilterValidatorSlashingReported(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CuratedModuleValidatorSlashingReportedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "ValidatorSlashingReported", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleValidatorSlashingReportedIterator{contract: _CuratedModule.contract, event: "ValidatorSlashingReported", logs: logs, sub: sub}, nil
}

// WatchValidatorSlashingReported is a free log subscription operation binding the contract event 0x4e37fce6f335bba671f891995433c6d300e4384b81a98ebb32eb5cfd5122109c.
//
// Solidity: event ValidatorSlashingReported(uint256 indexed nodeOperatorId, uint256 keyIndex, bytes pubkey)
func (_CuratedModule *CuratedModuleFilterer) WatchValidatorSlashingReported(opts *bind.WatchOpts, sink chan<- *CuratedModuleValidatorSlashingReported, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "ValidatorSlashingReported", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleValidatorSlashingReported)
				if err := _CuratedModule.contract.UnpackLog(event, "ValidatorSlashingReported", log); err != nil {
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

// ParseValidatorSlashingReported is a log parse operation binding the contract event 0x4e37fce6f335bba671f891995433c6d300e4384b81a98ebb32eb5cfd5122109c.
//
// Solidity: event ValidatorSlashingReported(uint256 indexed nodeOperatorId, uint256 keyIndex, bytes pubkey)
func (_CuratedModule *CuratedModuleFilterer) ParseValidatorSlashingReported(log types.Log) (*CuratedModuleValidatorSlashingReported, error) {
	event := new(CuratedModuleValidatorSlashingReported)
	if err := _CuratedModule.contract.UnpackLog(event, "ValidatorSlashingReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleValidatorWithdrawnIterator is returned from FilterValidatorWithdrawn and is used to iterate over the raw logs and unpacked data for ValidatorWithdrawn events raised by the CuratedModule contract.
type CuratedModuleValidatorWithdrawnIterator struct {
	Event *CuratedModuleValidatorWithdrawn // Event containing the contract specifics and raw log

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
func (it *CuratedModuleValidatorWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleValidatorWithdrawn)
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
		it.Event = new(CuratedModuleValidatorWithdrawn)
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
func (it *CuratedModuleValidatorWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleValidatorWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleValidatorWithdrawn represents a ValidatorWithdrawn event raised by the CuratedModule contract.
type CuratedModuleValidatorWithdrawn struct {
	NodeOperatorId  *big.Int
	KeyIndex        *big.Int
	ExitBalance     *big.Int
	SlashingPenalty *big.Int
	Pubkey          []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorWithdrawn is a free log retrieval operation binding the contract event 0x7e465451d2ca8f7f7859d898fa02125d274d5aada615daaee964691e1fdeb222.
//
// Solidity: event ValidatorWithdrawn(uint256 indexed nodeOperatorId, uint256 keyIndex, uint256 exitBalance, uint256 slashingPenalty, bytes pubkey)
func (_CuratedModule *CuratedModuleFilterer) FilterValidatorWithdrawn(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CuratedModuleValidatorWithdrawnIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "ValidatorWithdrawn", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleValidatorWithdrawnIterator{contract: _CuratedModule.contract, event: "ValidatorWithdrawn", logs: logs, sub: sub}, nil
}

// WatchValidatorWithdrawn is a free log subscription operation binding the contract event 0x7e465451d2ca8f7f7859d898fa02125d274d5aada615daaee964691e1fdeb222.
//
// Solidity: event ValidatorWithdrawn(uint256 indexed nodeOperatorId, uint256 keyIndex, uint256 exitBalance, uint256 slashingPenalty, bytes pubkey)
func (_CuratedModule *CuratedModuleFilterer) WatchValidatorWithdrawn(opts *bind.WatchOpts, sink chan<- *CuratedModuleValidatorWithdrawn, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "ValidatorWithdrawn", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleValidatorWithdrawn)
				if err := _CuratedModule.contract.UnpackLog(event, "ValidatorWithdrawn", log); err != nil {
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

// ParseValidatorWithdrawn is a log parse operation binding the contract event 0x7e465451d2ca8f7f7859d898fa02125d274d5aada615daaee964691e1fdeb222.
//
// Solidity: event ValidatorWithdrawn(uint256 indexed nodeOperatorId, uint256 keyIndex, uint256 exitBalance, uint256 slashingPenalty, bytes pubkey)
func (_CuratedModule *CuratedModuleFilterer) ParseValidatorWithdrawn(log types.Log) (*CuratedModuleValidatorWithdrawn, error) {
	event := new(CuratedModuleValidatorWithdrawn)
	if err := _CuratedModule.contract.UnpackLog(event, "ValidatorWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleVettedSigningKeysCountChangedIterator is returned from FilterVettedSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for VettedSigningKeysCountChanged events raised by the CuratedModule contract.
type CuratedModuleVettedSigningKeysCountChangedIterator struct {
	Event *CuratedModuleVettedSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *CuratedModuleVettedSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleVettedSigningKeysCountChanged)
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
		it.Event = new(CuratedModuleVettedSigningKeysCountChanged)
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
func (it *CuratedModuleVettedSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleVettedSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleVettedSigningKeysCountChanged represents a VettedSigningKeysCountChanged event raised by the CuratedModule contract.
type CuratedModuleVettedSigningKeysCountChanged struct {
	NodeOperatorId  *big.Int
	VettedKeysCount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVettedSigningKeysCountChanged is a free log retrieval operation binding the contract event 0x947f955eec7e1f626bee3afd2aa47b5de04ddcdd3fe78dc8838213015ef58dfd.
//
// Solidity: event VettedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 vettedKeysCount)
func (_CuratedModule *CuratedModuleFilterer) FilterVettedSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CuratedModuleVettedSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "VettedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleVettedSigningKeysCountChangedIterator{contract: _CuratedModule.contract, event: "VettedSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchVettedSigningKeysCountChanged is a free log subscription operation binding the contract event 0x947f955eec7e1f626bee3afd2aa47b5de04ddcdd3fe78dc8838213015ef58dfd.
//
// Solidity: event VettedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 vettedKeysCount)
func (_CuratedModule *CuratedModuleFilterer) WatchVettedSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *CuratedModuleVettedSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "VettedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleVettedSigningKeysCountChanged)
				if err := _CuratedModule.contract.UnpackLog(event, "VettedSigningKeysCountChanged", log); err != nil {
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
// Solidity: event VettedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 vettedKeysCount)
func (_CuratedModule *CuratedModuleFilterer) ParseVettedSigningKeysCountChanged(log types.Log) (*CuratedModuleVettedSigningKeysCountChanged, error) {
	event := new(CuratedModuleVettedSigningKeysCountChanged)
	if err := _CuratedModule.contract.UnpackLog(event, "VettedSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CuratedModuleVettedSigningKeysCountDecreasedIterator is returned from FilterVettedSigningKeysCountDecreased and is used to iterate over the raw logs and unpacked data for VettedSigningKeysCountDecreased events raised by the CuratedModule contract.
type CuratedModuleVettedSigningKeysCountDecreasedIterator struct {
	Event *CuratedModuleVettedSigningKeysCountDecreased // Event containing the contract specifics and raw log

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
func (it *CuratedModuleVettedSigningKeysCountDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CuratedModuleVettedSigningKeysCountDecreased)
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
		it.Event = new(CuratedModuleVettedSigningKeysCountDecreased)
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
func (it *CuratedModuleVettedSigningKeysCountDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CuratedModuleVettedSigningKeysCountDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CuratedModuleVettedSigningKeysCountDecreased represents a VettedSigningKeysCountDecreased event raised by the CuratedModule contract.
type CuratedModuleVettedSigningKeysCountDecreased struct {
	NodeOperatorId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterVettedSigningKeysCountDecreased is a free log retrieval operation binding the contract event 0xe5725d045d5c47bd1483feba445e395dc8647486963e6d54aad9ed03ff7d6ce6.
//
// Solidity: event VettedSigningKeysCountDecreased(uint256 indexed nodeOperatorId)
func (_CuratedModule *CuratedModuleFilterer) FilterVettedSigningKeysCountDecreased(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CuratedModuleVettedSigningKeysCountDecreasedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.FilterLogs(opts, "VettedSigningKeysCountDecreased", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleVettedSigningKeysCountDecreasedIterator{contract: _CuratedModule.contract, event: "VettedSigningKeysCountDecreased", logs: logs, sub: sub}, nil
}

// WatchVettedSigningKeysCountDecreased is a free log subscription operation binding the contract event 0xe5725d045d5c47bd1483feba445e395dc8647486963e6d54aad9ed03ff7d6ce6.
//
// Solidity: event VettedSigningKeysCountDecreased(uint256 indexed nodeOperatorId)
func (_CuratedModule *CuratedModuleFilterer) WatchVettedSigningKeysCountDecreased(opts *bind.WatchOpts, sink chan<- *CuratedModuleVettedSigningKeysCountDecreased, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CuratedModule.contract.WatchLogs(opts, "VettedSigningKeysCountDecreased", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CuratedModuleVettedSigningKeysCountDecreased)
				if err := _CuratedModule.contract.UnpackLog(event, "VettedSigningKeysCountDecreased", log); err != nil {
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

// ParseVettedSigningKeysCountDecreased is a log parse operation binding the contract event 0xe5725d045d5c47bd1483feba445e395dc8647486963e6d54aad9ed03ff7d6ce6.
//
// Solidity: event VettedSigningKeysCountDecreased(uint256 indexed nodeOperatorId)
func (_CuratedModule *CuratedModuleFilterer) ParseVettedSigningKeysCountDecreased(log types.Log) (*CuratedModuleVettedSigningKeysCountDecreased, error) {
	event := new(CuratedModuleVettedSigningKeysCountDecreased)
	if err := _CuratedModule.contract.UnpackLog(event, "VettedSigningKeysCountDecreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package csm

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

// ICSAccountingPermitInput is an auto generated low-level Go binding around an user-defined struct.
type ICSAccountingPermitInput struct {
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
}

// NodeOperatorManagementProperties is an auto generated low-level Go binding around an user-defined struct.
type NodeOperatorManagementProperties struct {
	ManagerAddress             common.Address
	RewardAddress              common.Address
	ExtendedManagerPermissions bool
}

// CsmMetaData contains all meta data concerning the Csm contract.
var CsmMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"moduleType\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"minSlashingPenaltyQuotient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"elRewardsStealingFine\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxKeysPerOperatorEA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxKeyRemovalCharge\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lidoLocator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyActivated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyProposed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadySubmitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExitedKeysDecrease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExitedKeysHigherThanTotalDeposited\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToSendEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInput\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeysCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReportData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVetKeysPointer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxSigningKeysCountExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MethodCallIsNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeOperatorDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedToJoinYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedToRecover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughKeys\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PauseUntilMustBeInFuture\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PausedExpected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QueueIsEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QueueLookupNoLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ResumedExpected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotEligible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotManagerAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotProposedAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotRewardAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SigningKeysInvalidOffset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StuckKeysHigherThanNonExited\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAccountingAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAdminAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroLocatorAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroPauseDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroRewardAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"BatchEnqueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositableKeysCount\",\"type\":\"uint256\"}],\"name\":\"DepositableSigningKeysCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositedKeysCount\",\"type\":\"uint256\"}],\"name\":\"DepositedSigningKeysCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ELRewardsStealingPenaltyCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ELRewardsStealingPenaltyCompensated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"proposedBlockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stolenAmount\",\"type\":\"uint256\"}],\"name\":\"ELRewardsStealingPenaltyReported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"ELRewardsStealingPenaltySettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC1155Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ERC721Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EtherRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exitedKeysCount\",\"type\":\"uint256\"}],\"name\":\"ExitedSigningKeysCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"InitialSlashingSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"KeyRemovalChargeApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"KeyRemovalChargeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"managerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardAddress\",\"type\":\"address\"}],\"name\":\"NodeOperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldProposedAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newProposedAddress\",\"type\":\"address\"}],\"name\":\"NodeOperatorManagerAddressChangeProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"NodeOperatorManagerAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldProposedAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newProposedAddress\",\"type\":\"address\"}],\"name\":\"NodeOperatorRewardAddressChangeProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"NodeOperatorRewardAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"NonceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PublicRelease\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"ReferrerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Resumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"SigningKeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"SigningKeyRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"StETHSharesRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stuckKeysCount\",\"type\":\"uint256\"}],\"name\":\"StuckSigningKeysCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetLimitMode\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetValidatorsCount\",\"type\":\"uint256\"}],\"name\":\"TargetValidatorsCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalKeysCount\",\"type\":\"uint256\"}],\"name\":\"TotalSigningKeysCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vettedKeysCount\",\"type\":\"uint256\"}],\"name\":\"VettedSigningKeysCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"VettedSigningKeysCountDecreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"WithdrawalSubmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EL_REWARDS_STEALING_FINE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIAL_SLASHING_PENALTY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIDO_LOCATOR\",\"outputs\":[{\"internalType\":\"contractILidoLocator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_KEY_REMOVAL_CHARGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_SIGNING_KEYS_PER_OPERATOR_BEFORE_PUBLIC_RELEASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_INFINITELY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RECOVERER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REPORT_EL_REWARDS_STEALING_PENALTY_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESUME_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SETTLE_EL_REWARDS_STEALING_PENALTY_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_ROUTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STETH\",\"outputs\":[{\"internalType\":\"contractIStETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFIER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accounting\",\"outputs\":[{\"internalType\":\"contractICSAccounting\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activatePublicRelease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"publicKeys\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"managerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"extendedManagerPermissions\",\"type\":\"bool\"}],\"internalType\":\"structNodeOperatorManagementProperties\",\"name\":\"managementProperties\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"eaProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"addNodeOperatorETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"publicKeys\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"managerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"extendedManagerPermissions\",\"type\":\"bool\"}],\"internalType\":\"structNodeOperatorManagementProperties\",\"name\":\"managementProperties\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structICSAccounting.PermitInput\",\"name\":\"permit\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"eaProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"addNodeOperatorStETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"publicKeys\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"managerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"extendedManagerPermissions\",\"type\":\"bool\"}],\"internalType\":\"structNodeOperatorManagementProperties\",\"name\":\"managementProperties\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structICSAccounting.PermitInput\",\"name\":\"permit\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"eaProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"addNodeOperatorWstETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"publicKeys\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"addValidatorKeysETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"publicKeys\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structICSAccounting.PermitInput\",\"name\":\"permit\",\"type\":\"tuple\"}],\"name\":\"addValidatorKeysStETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"publicKeys\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structICSAccounting.PermitInput\",\"name\":\"permit\",\"type\":\"tuple\"}],\"name\":\"addValidatorKeysWstETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"cancelELRewardsStealingPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"changeNodeOperatorRewardAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stETHAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeFeeShares\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"rewardsProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimRewardsStETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stEthAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeFeeShares\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"rewardsProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimRewardsUnstETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wstETHAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeFeeShares\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"rewardsProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimRewardsWstETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxItems\",\"type\":\"uint256\"}],\"name\":\"cleanDepositQueue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"removed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRemovedAtDepth\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"compensateELRewardsStealingPenalty\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"confirmNodeOperatorManagerAddressChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"confirmNodeOperatorRewardAddressChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeOperatorIds\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vettedSigningKeysCounts\",\"type\":\"bytes\"}],\"name\":\"decreaseVettedSigningKeysCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositQueue\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"head\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"tail\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"index\",\"type\":\"uint128\"}],\"name\":\"depositQueueItem\",\"outputs\":[{\"internalType\":\"Batch\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stETHAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structICSAccounting.PermitInput\",\"name\":\"permit\",\"type\":\"tuple\"}],\"name\":\"depositStETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wstETHAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structICSAccounting.PermitInput\",\"name\":\"permit\",\"type\":\"tuple\"}],\"name\":\"depositWstETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"earlyAdoption\",\"outputs\":[{\"internalType\":\"contractICSEarlyAdoption\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveNodeOperatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getNodeOperator\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"totalAddedKeys\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"totalWithdrawnKeys\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"totalDepositedKeys\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"totalVettedKeys\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"stuckValidatorsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"depositableValidatorsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"targetLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"targetLimitMode\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"totalExitedKeys\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"enqueuedCount\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"managerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposedManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposedRewardAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"extendedManagerPermissions\",\"type\":\"bool\"}],\"internalType\":\"structNodeOperator\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getNodeOperatorIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nodeOperatorIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getNodeOperatorIsActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getNodeOperatorNonWithdrawnKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getNodeOperatorSummary\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"targetLimitMode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stuckValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stuckPenaltyEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalExitedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDepositedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositableValidatorsCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeOperatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getResumeSinceTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"}],\"name\":\"getSigningKeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"}],\"name\":\"getSigningKeysWithSignatures\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"keys\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingModuleSummary\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalExitedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDepositedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositableValidatorsCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_accounting\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_earlyAdoption\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_keyRemovalCharge\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"}],\"name\":\"isValidatorSlashed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"}],\"name\":\"isValidatorWithdrawn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keyRemovalCharge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"normalizeQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositsCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"obtainDepositData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKeys\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onExitedAndStuckValidatorsCountsUpdated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"}],\"name\":\"onRewardsMinted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onWithdrawalCredentialsChanged\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"pauseFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"}],\"name\":\"proposeNodeOperatorManagerAddressChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"}],\"name\":\"proposeNodeOperatorRewardAddressChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicRelease\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recoverERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recoverERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverEther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"}],\"name\":\"removeKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"reportELRewardsStealingPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"resetNodeOperatorManagerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setKeyRemovalCharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nodeOperatorIds\",\"type\":\"uint256[]\"}],\"name\":\"settleELRewardsStealingPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"}],\"name\":\"submitInitialSlashing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSlashed\",\"type\":\"bool\"}],\"name\":\"submitWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exitedValidatorsKeysCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stuckValidatorsKeysCount\",\"type\":\"uint256\"}],\"name\":\"unsafeUpdateValidatorsCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeOperatorIds\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"exitedValidatorsCounts\",\"type\":\"bytes\"}],\"name\":\"updateExitedValidatorsCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"updateRefundedValidatorsCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeOperatorIds\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"stuckValidatorsCounts\",\"type\":\"bytes\"}],\"name\":\"updateStuckValidatorsCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetLimitMode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetLimit\",\"type\":\"uint256\"}],\"name\":\"updateTargetValidatorsLimits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CsmABI is the input ABI used to generate the binding from.
// Deprecated: Use CsmMetaData.ABI instead.
var CsmABI = CsmMetaData.ABI

// Csm is an auto generated Go binding around an Ethereum contract.
type Csm struct {
	CsmCaller     // Read-only binding to the contract
	CsmTransactor // Write-only binding to the contract
	CsmFilterer   // Log filterer for contract events
}

// CsmCaller is an auto generated read-only Go binding around an Ethereum contract.
type CsmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CsmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CsmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CsmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CsmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CsmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CsmSession struct {
	Contract     *Csm              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CsmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CsmCallerSession struct {
	Contract *CsmCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CsmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CsmTransactorSession struct {
	Contract     *CsmTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CsmRaw is an auto generated low-level Go binding around an Ethereum contract.
type CsmRaw struct {
	Contract *Csm // Generic contract binding to access the raw methods on
}

// CsmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CsmCallerRaw struct {
	Contract *CsmCaller // Generic read-only contract binding to access the raw methods on
}

// CsmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CsmTransactorRaw struct {
	Contract *CsmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCsm creates a new instance of Csm, bound to a specific deployed contract.
func NewCsm(address common.Address, backend bind.ContractBackend) (*Csm, error) {
	contract, err := bindCsm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Csm{CsmCaller: CsmCaller{contract: contract}, CsmTransactor: CsmTransactor{contract: contract}, CsmFilterer: CsmFilterer{contract: contract}}, nil
}

// NewCsmCaller creates a new read-only instance of Csm, bound to a specific deployed contract.
func NewCsmCaller(address common.Address, caller bind.ContractCaller) (*CsmCaller, error) {
	contract, err := bindCsm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CsmCaller{contract: contract}, nil
}

// NewCsmTransactor creates a new write-only instance of Csm, bound to a specific deployed contract.
func NewCsmTransactor(address common.Address, transactor bind.ContractTransactor) (*CsmTransactor, error) {
	contract, err := bindCsm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CsmTransactor{contract: contract}, nil
}

// NewCsmFilterer creates a new log filterer instance of Csm, bound to a specific deployed contract.
func NewCsmFilterer(address common.Address, filterer bind.ContractFilterer) (*CsmFilterer, error) {
	contract, err := bindCsm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CsmFilterer{contract: contract}, nil
}

// bindCsm binds a generic wrapper to an already deployed contract.
func bindCsm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CsmMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Csm *CsmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Csm.Contract.CsmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Csm *CsmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csm.Contract.CsmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Csm *CsmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Csm.Contract.CsmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Csm *CsmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Csm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Csm *CsmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Csm *CsmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Csm.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Csm *CsmCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Csm *CsmSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Csm.Contract.DEFAULTADMINROLE(&_Csm.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Csm *CsmCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Csm.Contract.DEFAULTADMINROLE(&_Csm.CallOpts)
}

// ELREWARDSSTEALINGFINE is a free data retrieval call binding the contract method 0xbdac46a2.
//
// Solidity: function EL_REWARDS_STEALING_FINE() view returns(uint256)
func (_Csm *CsmCaller) ELREWARDSSTEALINGFINE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "EL_REWARDS_STEALING_FINE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ELREWARDSSTEALINGFINE is a free data retrieval call binding the contract method 0xbdac46a2.
//
// Solidity: function EL_REWARDS_STEALING_FINE() view returns(uint256)
func (_Csm *CsmSession) ELREWARDSSTEALINGFINE() (*big.Int, error) {
	return _Csm.Contract.ELREWARDSSTEALINGFINE(&_Csm.CallOpts)
}

// ELREWARDSSTEALINGFINE is a free data retrieval call binding the contract method 0xbdac46a2.
//
// Solidity: function EL_REWARDS_STEALING_FINE() view returns(uint256)
func (_Csm *CsmCallerSession) ELREWARDSSTEALINGFINE() (*big.Int, error) {
	return _Csm.Contract.ELREWARDSSTEALINGFINE(&_Csm.CallOpts)
}

// INITIALSLASHINGPENALTY is a free data retrieval call binding the contract method 0xd6477919.
//
// Solidity: function INITIAL_SLASHING_PENALTY() view returns(uint256)
func (_Csm *CsmCaller) INITIALSLASHINGPENALTY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "INITIAL_SLASHING_PENALTY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITIALSLASHINGPENALTY is a free data retrieval call binding the contract method 0xd6477919.
//
// Solidity: function INITIAL_SLASHING_PENALTY() view returns(uint256)
func (_Csm *CsmSession) INITIALSLASHINGPENALTY() (*big.Int, error) {
	return _Csm.Contract.INITIALSLASHINGPENALTY(&_Csm.CallOpts)
}

// INITIALSLASHINGPENALTY is a free data retrieval call binding the contract method 0xd6477919.
//
// Solidity: function INITIAL_SLASHING_PENALTY() view returns(uint256)
func (_Csm *CsmCallerSession) INITIALSLASHINGPENALTY() (*big.Int, error) {
	return _Csm.Contract.INITIALSLASHINGPENALTY(&_Csm.CallOpts)
}

// LIDOLOCATOR is a free data retrieval call binding the contract method 0xdbba4b48.
//
// Solidity: function LIDO_LOCATOR() view returns(address)
func (_Csm *CsmCaller) LIDOLOCATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "LIDO_LOCATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIDOLOCATOR is a free data retrieval call binding the contract method 0xdbba4b48.
//
// Solidity: function LIDO_LOCATOR() view returns(address)
func (_Csm *CsmSession) LIDOLOCATOR() (common.Address, error) {
	return _Csm.Contract.LIDOLOCATOR(&_Csm.CallOpts)
}

// LIDOLOCATOR is a free data retrieval call binding the contract method 0xdbba4b48.
//
// Solidity: function LIDO_LOCATOR() view returns(address)
func (_Csm *CsmCallerSession) LIDOLOCATOR() (common.Address, error) {
	return _Csm.Contract.LIDOLOCATOR(&_Csm.CallOpts)
}

// MAXKEYREMOVALCHARGE is a free data retrieval call binding the contract method 0x4baf13cc.
//
// Solidity: function MAX_KEY_REMOVAL_CHARGE() view returns(uint256)
func (_Csm *CsmCaller) MAXKEYREMOVALCHARGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "MAX_KEY_REMOVAL_CHARGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXKEYREMOVALCHARGE is a free data retrieval call binding the contract method 0x4baf13cc.
//
// Solidity: function MAX_KEY_REMOVAL_CHARGE() view returns(uint256)
func (_Csm *CsmSession) MAXKEYREMOVALCHARGE() (*big.Int, error) {
	return _Csm.Contract.MAXKEYREMOVALCHARGE(&_Csm.CallOpts)
}

// MAXKEYREMOVALCHARGE is a free data retrieval call binding the contract method 0x4baf13cc.
//
// Solidity: function MAX_KEY_REMOVAL_CHARGE() view returns(uint256)
func (_Csm *CsmCallerSession) MAXKEYREMOVALCHARGE() (*big.Int, error) {
	return _Csm.Contract.MAXKEYREMOVALCHARGE(&_Csm.CallOpts)
}

// MAXSIGNINGKEYSPEROPERATORBEFOREPUBLICRELEASE is a free data retrieval call binding the contract method 0x47faf311.
//
// Solidity: function MAX_SIGNING_KEYS_PER_OPERATOR_BEFORE_PUBLIC_RELEASE() view returns(uint256)
func (_Csm *CsmCaller) MAXSIGNINGKEYSPEROPERATORBEFOREPUBLICRELEASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "MAX_SIGNING_KEYS_PER_OPERATOR_BEFORE_PUBLIC_RELEASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSIGNINGKEYSPEROPERATORBEFOREPUBLICRELEASE is a free data retrieval call binding the contract method 0x47faf311.
//
// Solidity: function MAX_SIGNING_KEYS_PER_OPERATOR_BEFORE_PUBLIC_RELEASE() view returns(uint256)
func (_Csm *CsmSession) MAXSIGNINGKEYSPEROPERATORBEFOREPUBLICRELEASE() (*big.Int, error) {
	return _Csm.Contract.MAXSIGNINGKEYSPEROPERATORBEFOREPUBLICRELEASE(&_Csm.CallOpts)
}

// MAXSIGNINGKEYSPEROPERATORBEFOREPUBLICRELEASE is a free data retrieval call binding the contract method 0x47faf311.
//
// Solidity: function MAX_SIGNING_KEYS_PER_OPERATOR_BEFORE_PUBLIC_RELEASE() view returns(uint256)
func (_Csm *CsmCallerSession) MAXSIGNINGKEYSPEROPERATORBEFOREPUBLICRELEASE() (*big.Int, error) {
	return _Csm.Contract.MAXSIGNINGKEYSPEROPERATORBEFOREPUBLICRELEASE(&_Csm.CallOpts)
}

// MODULEMANAGERROLE is a free data retrieval call binding the contract method 0xcb17ed3e.
//
// Solidity: function MODULE_MANAGER_ROLE() view returns(bytes32)
func (_Csm *CsmCaller) MODULEMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "MODULE_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MODULEMANAGERROLE is a free data retrieval call binding the contract method 0xcb17ed3e.
//
// Solidity: function MODULE_MANAGER_ROLE() view returns(bytes32)
func (_Csm *CsmSession) MODULEMANAGERROLE() ([32]byte, error) {
	return _Csm.Contract.MODULEMANAGERROLE(&_Csm.CallOpts)
}

// MODULEMANAGERROLE is a free data retrieval call binding the contract method 0xcb17ed3e.
//
// Solidity: function MODULE_MANAGER_ROLE() view returns(bytes32)
func (_Csm *CsmCallerSession) MODULEMANAGERROLE() ([32]byte, error) {
	return _Csm.Contract.MODULEMANAGERROLE(&_Csm.CallOpts)
}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_Csm *CsmCaller) PAUSEINFINITELY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "PAUSE_INFINITELY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_Csm *CsmSession) PAUSEINFINITELY() (*big.Int, error) {
	return _Csm.Contract.PAUSEINFINITELY(&_Csm.CallOpts)
}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_Csm *CsmCallerSession) PAUSEINFINITELY() (*big.Int, error) {
	return _Csm.Contract.PAUSEINFINITELY(&_Csm.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Csm *CsmCaller) PAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Csm *CsmSession) PAUSEROLE() ([32]byte, error) {
	return _Csm.Contract.PAUSEROLE(&_Csm.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Csm *CsmCallerSession) PAUSEROLE() ([32]byte, error) {
	return _Csm.Contract.PAUSEROLE(&_Csm.CallOpts)
}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_Csm *CsmCaller) RECOVERERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "RECOVERER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_Csm *CsmSession) RECOVERERROLE() ([32]byte, error) {
	return _Csm.Contract.RECOVERERROLE(&_Csm.CallOpts)
}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_Csm *CsmCallerSession) RECOVERERROLE() ([32]byte, error) {
	return _Csm.Contract.RECOVERERROLE(&_Csm.CallOpts)
}

// REPORTELREWARDSSTEALINGPENALTYROLE is a free data retrieval call binding the contract method 0x8573e351.
//
// Solidity: function REPORT_EL_REWARDS_STEALING_PENALTY_ROLE() view returns(bytes32)
func (_Csm *CsmCaller) REPORTELREWARDSSTEALINGPENALTYROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "REPORT_EL_REWARDS_STEALING_PENALTY_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REPORTELREWARDSSTEALINGPENALTYROLE is a free data retrieval call binding the contract method 0x8573e351.
//
// Solidity: function REPORT_EL_REWARDS_STEALING_PENALTY_ROLE() view returns(bytes32)
func (_Csm *CsmSession) REPORTELREWARDSSTEALINGPENALTYROLE() ([32]byte, error) {
	return _Csm.Contract.REPORTELREWARDSSTEALINGPENALTYROLE(&_Csm.CallOpts)
}

// REPORTELREWARDSSTEALINGPENALTYROLE is a free data retrieval call binding the contract method 0x8573e351.
//
// Solidity: function REPORT_EL_REWARDS_STEALING_PENALTY_ROLE() view returns(bytes32)
func (_Csm *CsmCallerSession) REPORTELREWARDSSTEALINGPENALTYROLE() ([32]byte, error) {
	return _Csm.Contract.REPORTELREWARDSSTEALINGPENALTYROLE(&_Csm.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Csm *CsmCaller) RESUMEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "RESUME_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Csm *CsmSession) RESUMEROLE() ([32]byte, error) {
	return _Csm.Contract.RESUMEROLE(&_Csm.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Csm *CsmCallerSession) RESUMEROLE() ([32]byte, error) {
	return _Csm.Contract.RESUMEROLE(&_Csm.CallOpts)
}

// SETTLEELREWARDSSTEALINGPENALTYROLE is a free data retrieval call binding the contract method 0x3f04f0c8.
//
// Solidity: function SETTLE_EL_REWARDS_STEALING_PENALTY_ROLE() view returns(bytes32)
func (_Csm *CsmCaller) SETTLEELREWARDSSTEALINGPENALTYROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "SETTLE_EL_REWARDS_STEALING_PENALTY_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETTLEELREWARDSSTEALINGPENALTYROLE is a free data retrieval call binding the contract method 0x3f04f0c8.
//
// Solidity: function SETTLE_EL_REWARDS_STEALING_PENALTY_ROLE() view returns(bytes32)
func (_Csm *CsmSession) SETTLEELREWARDSSTEALINGPENALTYROLE() ([32]byte, error) {
	return _Csm.Contract.SETTLEELREWARDSSTEALINGPENALTYROLE(&_Csm.CallOpts)
}

// SETTLEELREWARDSSTEALINGPENALTYROLE is a free data retrieval call binding the contract method 0x3f04f0c8.
//
// Solidity: function SETTLE_EL_REWARDS_STEALING_PENALTY_ROLE() view returns(bytes32)
func (_Csm *CsmCallerSession) SETTLEELREWARDSSTEALINGPENALTYROLE() ([32]byte, error) {
	return _Csm.Contract.SETTLEELREWARDSSTEALINGPENALTYROLE(&_Csm.CallOpts)
}

// STAKINGROUTERROLE is a free data retrieval call binding the contract method 0x80231f15.
//
// Solidity: function STAKING_ROUTER_ROLE() view returns(bytes32)
func (_Csm *CsmCaller) STAKINGROUTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "STAKING_ROUTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGROUTERROLE is a free data retrieval call binding the contract method 0x80231f15.
//
// Solidity: function STAKING_ROUTER_ROLE() view returns(bytes32)
func (_Csm *CsmSession) STAKINGROUTERROLE() ([32]byte, error) {
	return _Csm.Contract.STAKINGROUTERROLE(&_Csm.CallOpts)
}

// STAKINGROUTERROLE is a free data retrieval call binding the contract method 0x80231f15.
//
// Solidity: function STAKING_ROUTER_ROLE() view returns(bytes32)
func (_Csm *CsmCallerSession) STAKINGROUTERROLE() ([32]byte, error) {
	return _Csm.Contract.STAKINGROUTERROLE(&_Csm.CallOpts)
}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_Csm *CsmCaller) STETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "STETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_Csm *CsmSession) STETH() (common.Address, error) {
	return _Csm.Contract.STETH(&_Csm.CallOpts)
}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_Csm *CsmCallerSession) STETH() (common.Address, error) {
	return _Csm.Contract.STETH(&_Csm.CallOpts)
}

// VERIFIERROLE is a free data retrieval call binding the contract method 0xe7705db6.
//
// Solidity: function VERIFIER_ROLE() view returns(bytes32)
func (_Csm *CsmCaller) VERIFIERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "VERIFIER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VERIFIERROLE is a free data retrieval call binding the contract method 0xe7705db6.
//
// Solidity: function VERIFIER_ROLE() view returns(bytes32)
func (_Csm *CsmSession) VERIFIERROLE() ([32]byte, error) {
	return _Csm.Contract.VERIFIERROLE(&_Csm.CallOpts)
}

// VERIFIERROLE is a free data retrieval call binding the contract method 0xe7705db6.
//
// Solidity: function VERIFIER_ROLE() view returns(bytes32)
func (_Csm *CsmCallerSession) VERIFIERROLE() ([32]byte, error) {
	return _Csm.Contract.VERIFIERROLE(&_Csm.CallOpts)
}

// Accounting is a free data retrieval call binding the contract method 0x9624e83e.
//
// Solidity: function accounting() view returns(address)
func (_Csm *CsmCaller) Accounting(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "accounting")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Accounting is a free data retrieval call binding the contract method 0x9624e83e.
//
// Solidity: function accounting() view returns(address)
func (_Csm *CsmSession) Accounting() (common.Address, error) {
	return _Csm.Contract.Accounting(&_Csm.CallOpts)
}

// Accounting is a free data retrieval call binding the contract method 0x9624e83e.
//
// Solidity: function accounting() view returns(address)
func (_Csm *CsmCallerSession) Accounting() (common.Address, error) {
	return _Csm.Contract.Accounting(&_Csm.CallOpts)
}

// DepositQueue is a free data retrieval call binding the contract method 0xf617eecc.
//
// Solidity: function depositQueue() view returns(uint128 head, uint128 tail)
func (_Csm *CsmCaller) DepositQueue(opts *bind.CallOpts) (struct {
	Head *big.Int
	Tail *big.Int
}, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "depositQueue")

	outstruct := new(struct {
		Head *big.Int
		Tail *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Head = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Tail = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DepositQueue is a free data retrieval call binding the contract method 0xf617eecc.
//
// Solidity: function depositQueue() view returns(uint128 head, uint128 tail)
func (_Csm *CsmSession) DepositQueue() (struct {
	Head *big.Int
	Tail *big.Int
}, error) {
	return _Csm.Contract.DepositQueue(&_Csm.CallOpts)
}

// DepositQueue is a free data retrieval call binding the contract method 0xf617eecc.
//
// Solidity: function depositQueue() view returns(uint128 head, uint128 tail)
func (_Csm *CsmCallerSession) DepositQueue() (struct {
	Head *big.Int
	Tail *big.Int
}, error) {
	return _Csm.Contract.DepositQueue(&_Csm.CallOpts)
}

// DepositQueueItem is a free data retrieval call binding the contract method 0x5e169fb8.
//
// Solidity: function depositQueueItem(uint128 index) view returns(uint256)
func (_Csm *CsmCaller) DepositQueueItem(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "depositQueueItem", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositQueueItem is a free data retrieval call binding the contract method 0x5e169fb8.
//
// Solidity: function depositQueueItem(uint128 index) view returns(uint256)
func (_Csm *CsmSession) DepositQueueItem(index *big.Int) (*big.Int, error) {
	return _Csm.Contract.DepositQueueItem(&_Csm.CallOpts, index)
}

// DepositQueueItem is a free data retrieval call binding the contract method 0x5e169fb8.
//
// Solidity: function depositQueueItem(uint128 index) view returns(uint256)
func (_Csm *CsmCallerSession) DepositQueueItem(index *big.Int) (*big.Int, error) {
	return _Csm.Contract.DepositQueueItem(&_Csm.CallOpts, index)
}

// EarlyAdoption is a free data retrieval call binding the contract method 0x26a666e4.
//
// Solidity: function earlyAdoption() view returns(address)
func (_Csm *CsmCaller) EarlyAdoption(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "earlyAdoption")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EarlyAdoption is a free data retrieval call binding the contract method 0x26a666e4.
//
// Solidity: function earlyAdoption() view returns(address)
func (_Csm *CsmSession) EarlyAdoption() (common.Address, error) {
	return _Csm.Contract.EarlyAdoption(&_Csm.CallOpts)
}

// EarlyAdoption is a free data retrieval call binding the contract method 0x26a666e4.
//
// Solidity: function earlyAdoption() view returns(address)
func (_Csm *CsmCallerSession) EarlyAdoption() (common.Address, error) {
	return _Csm.Contract.EarlyAdoption(&_Csm.CallOpts)
}

// GetActiveNodeOperatorsCount is a free data retrieval call binding the contract method 0x8469cbd3.
//
// Solidity: function getActiveNodeOperatorsCount() view returns(uint256)
func (_Csm *CsmCaller) GetActiveNodeOperatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "getActiveNodeOperatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveNodeOperatorsCount is a free data retrieval call binding the contract method 0x8469cbd3.
//
// Solidity: function getActiveNodeOperatorsCount() view returns(uint256)
func (_Csm *CsmSession) GetActiveNodeOperatorsCount() (*big.Int, error) {
	return _Csm.Contract.GetActiveNodeOperatorsCount(&_Csm.CallOpts)
}

// GetActiveNodeOperatorsCount is a free data retrieval call binding the contract method 0x8469cbd3.
//
// Solidity: function getActiveNodeOperatorsCount() view returns(uint256)
func (_Csm *CsmCallerSession) GetActiveNodeOperatorsCount() (*big.Int, error) {
	return _Csm.Contract.GetActiveNodeOperatorsCount(&_Csm.CallOpts)
}

// GetNodeOperator is a free data retrieval call binding the contract method 0x65c14dc7.
//
// Solidity: function getNodeOperator(uint256 nodeOperatorId) view returns((uint32,uint32,uint32,uint32,uint32,uint32,uint32,uint8,uint32,uint32,address,address,address,address,bool))
func (_Csm *CsmCaller) GetNodeOperator(opts *bind.CallOpts, nodeOperatorId *big.Int) (NodeOperator, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "getNodeOperator", nodeOperatorId)

	if err != nil {
		return *new(NodeOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(NodeOperator)).(*NodeOperator)

	return out0, err

}

// GetNodeOperator is a free data retrieval call binding the contract method 0x65c14dc7.
//
// Solidity: function getNodeOperator(uint256 nodeOperatorId) view returns((uint32,uint32,uint32,uint32,uint32,uint32,uint32,uint8,uint32,uint32,address,address,address,address,bool))
func (_Csm *CsmSession) GetNodeOperator(nodeOperatorId *big.Int) (NodeOperator, error) {
	return _Csm.Contract.GetNodeOperator(&_Csm.CallOpts, nodeOperatorId)
}

// GetNodeOperator is a free data retrieval call binding the contract method 0x65c14dc7.
//
// Solidity: function getNodeOperator(uint256 nodeOperatorId) view returns((uint32,uint32,uint32,uint32,uint32,uint32,uint32,uint8,uint32,uint32,address,address,address,address,bool))
func (_Csm *CsmCallerSession) GetNodeOperator(nodeOperatorId *big.Int) (NodeOperator, error) {
	return _Csm.Contract.GetNodeOperator(&_Csm.CallOpts, nodeOperatorId)
}

// GetNodeOperatorIds is a free data retrieval call binding the contract method 0x4febc81b.
//
// Solidity: function getNodeOperatorIds(uint256 offset, uint256 limit) view returns(uint256[] nodeOperatorIds)
func (_Csm *CsmCaller) GetNodeOperatorIds(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "getNodeOperatorIds", offset, limit)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetNodeOperatorIds is a free data retrieval call binding the contract method 0x4febc81b.
//
// Solidity: function getNodeOperatorIds(uint256 offset, uint256 limit) view returns(uint256[] nodeOperatorIds)
func (_Csm *CsmSession) GetNodeOperatorIds(offset *big.Int, limit *big.Int) ([]*big.Int, error) {
	return _Csm.Contract.GetNodeOperatorIds(&_Csm.CallOpts, offset, limit)
}

// GetNodeOperatorIds is a free data retrieval call binding the contract method 0x4febc81b.
//
// Solidity: function getNodeOperatorIds(uint256 offset, uint256 limit) view returns(uint256[] nodeOperatorIds)
func (_Csm *CsmCallerSession) GetNodeOperatorIds(offset *big.Int, limit *big.Int) ([]*big.Int, error) {
	return _Csm.Contract.GetNodeOperatorIds(&_Csm.CallOpts, offset, limit)
}

// GetNodeOperatorIsActive is a free data retrieval call binding the contract method 0x5e2fb908.
//
// Solidity: function getNodeOperatorIsActive(uint256 nodeOperatorId) view returns(bool)
func (_Csm *CsmCaller) GetNodeOperatorIsActive(opts *bind.CallOpts, nodeOperatorId *big.Int) (bool, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "getNodeOperatorIsActive", nodeOperatorId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetNodeOperatorIsActive is a free data retrieval call binding the contract method 0x5e2fb908.
//
// Solidity: function getNodeOperatorIsActive(uint256 nodeOperatorId) view returns(bool)
func (_Csm *CsmSession) GetNodeOperatorIsActive(nodeOperatorId *big.Int) (bool, error) {
	return _Csm.Contract.GetNodeOperatorIsActive(&_Csm.CallOpts, nodeOperatorId)
}

// GetNodeOperatorIsActive is a free data retrieval call binding the contract method 0x5e2fb908.
//
// Solidity: function getNodeOperatorIsActive(uint256 nodeOperatorId) view returns(bool)
func (_Csm *CsmCallerSession) GetNodeOperatorIsActive(nodeOperatorId *big.Int) (bool, error) {
	return _Csm.Contract.GetNodeOperatorIsActive(&_Csm.CallOpts, nodeOperatorId)
}

// GetNodeOperatorNonWithdrawnKeys is a free data retrieval call binding the contract method 0x8ec69028.
//
// Solidity: function getNodeOperatorNonWithdrawnKeys(uint256 nodeOperatorId) view returns(uint256)
func (_Csm *CsmCaller) GetNodeOperatorNonWithdrawnKeys(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "getNodeOperatorNonWithdrawnKeys", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeOperatorNonWithdrawnKeys is a free data retrieval call binding the contract method 0x8ec69028.
//
// Solidity: function getNodeOperatorNonWithdrawnKeys(uint256 nodeOperatorId) view returns(uint256)
func (_Csm *CsmSession) GetNodeOperatorNonWithdrawnKeys(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csm.Contract.GetNodeOperatorNonWithdrawnKeys(&_Csm.CallOpts, nodeOperatorId)
}

// GetNodeOperatorNonWithdrawnKeys is a free data retrieval call binding the contract method 0x8ec69028.
//
// Solidity: function getNodeOperatorNonWithdrawnKeys(uint256 nodeOperatorId) view returns(uint256)
func (_Csm *CsmCallerSession) GetNodeOperatorNonWithdrawnKeys(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csm.Contract.GetNodeOperatorNonWithdrawnKeys(&_Csm.CallOpts, nodeOperatorId)
}

// GetNodeOperatorSummary is a free data retrieval call binding the contract method 0xb3076c3c.
//
// Solidity: function getNodeOperatorSummary(uint256 nodeOperatorId) view returns(uint256 targetLimitMode, uint256 targetValidatorsCount, uint256 stuckValidatorsCount, uint256 refundedValidatorsCount, uint256 stuckPenaltyEndTimestamp, uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_Csm *CsmCaller) GetNodeOperatorSummary(opts *bind.CallOpts, nodeOperatorId *big.Int) (struct {
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
	err := _Csm.contract.Call(opts, &out, "getNodeOperatorSummary", nodeOperatorId)

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
func (_Csm *CsmSession) GetNodeOperatorSummary(nodeOperatorId *big.Int) (struct {
	TargetLimitMode            *big.Int
	TargetValidatorsCount      *big.Int
	StuckValidatorsCount       *big.Int
	RefundedValidatorsCount    *big.Int
	StuckPenaltyEndTimestamp   *big.Int
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	return _Csm.Contract.GetNodeOperatorSummary(&_Csm.CallOpts, nodeOperatorId)
}

// GetNodeOperatorSummary is a free data retrieval call binding the contract method 0xb3076c3c.
//
// Solidity: function getNodeOperatorSummary(uint256 nodeOperatorId) view returns(uint256 targetLimitMode, uint256 targetValidatorsCount, uint256 stuckValidatorsCount, uint256 refundedValidatorsCount, uint256 stuckPenaltyEndTimestamp, uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_Csm *CsmCallerSession) GetNodeOperatorSummary(nodeOperatorId *big.Int) (struct {
	TargetLimitMode            *big.Int
	TargetValidatorsCount      *big.Int
	StuckValidatorsCount       *big.Int
	RefundedValidatorsCount    *big.Int
	StuckPenaltyEndTimestamp   *big.Int
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	return _Csm.Contract.GetNodeOperatorSummary(&_Csm.CallOpts, nodeOperatorId)
}

// GetNodeOperatorsCount is a free data retrieval call binding the contract method 0xa70c70e4.
//
// Solidity: function getNodeOperatorsCount() view returns(uint256)
func (_Csm *CsmCaller) GetNodeOperatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "getNodeOperatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeOperatorsCount is a free data retrieval call binding the contract method 0xa70c70e4.
//
// Solidity: function getNodeOperatorsCount() view returns(uint256)
func (_Csm *CsmSession) GetNodeOperatorsCount() (*big.Int, error) {
	return _Csm.Contract.GetNodeOperatorsCount(&_Csm.CallOpts)
}

// GetNodeOperatorsCount is a free data retrieval call binding the contract method 0xa70c70e4.
//
// Solidity: function getNodeOperatorsCount() view returns(uint256)
func (_Csm *CsmCallerSession) GetNodeOperatorsCount() (*big.Int, error) {
	return _Csm.Contract.GetNodeOperatorsCount(&_Csm.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Csm *CsmCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Csm *CsmSession) GetNonce() (*big.Int, error) {
	return _Csm.Contract.GetNonce(&_Csm.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Csm *CsmCallerSession) GetNonce() (*big.Int, error) {
	return _Csm.Contract.GetNonce(&_Csm.CallOpts)
}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_Csm *CsmCaller) GetResumeSinceTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "getResumeSinceTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_Csm *CsmSession) GetResumeSinceTimestamp() (*big.Int, error) {
	return _Csm.Contract.GetResumeSinceTimestamp(&_Csm.CallOpts)
}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_Csm *CsmCallerSession) GetResumeSinceTimestamp() (*big.Int, error) {
	return _Csm.Contract.GetResumeSinceTimestamp(&_Csm.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Csm *CsmCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Csm *CsmSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Csm.Contract.GetRoleAdmin(&_Csm.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Csm *CsmCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Csm.Contract.GetRoleAdmin(&_Csm.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Csm *CsmCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Csm *CsmSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Csm.Contract.GetRoleMember(&_Csm.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Csm *CsmCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Csm.Contract.GetRoleMember(&_Csm.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Csm *CsmCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Csm *CsmSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Csm.Contract.GetRoleMemberCount(&_Csm.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Csm *CsmCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Csm.Contract.GetRoleMemberCount(&_Csm.CallOpts, role)
}

// GetSigningKeys is a free data retrieval call binding the contract method 0x59e25c12.
//
// Solidity: function getSigningKeys(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(bytes)
func (_Csm *CsmCaller) GetSigningKeys(opts *bind.CallOpts, nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "getSigningKeys", nodeOperatorId, startIndex, keysCount)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetSigningKeys is a free data retrieval call binding the contract method 0x59e25c12.
//
// Solidity: function getSigningKeys(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(bytes)
func (_Csm *CsmSession) GetSigningKeys(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) ([]byte, error) {
	return _Csm.Contract.GetSigningKeys(&_Csm.CallOpts, nodeOperatorId, startIndex, keysCount)
}

// GetSigningKeys is a free data retrieval call binding the contract method 0x59e25c12.
//
// Solidity: function getSigningKeys(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(bytes)
func (_Csm *CsmCallerSession) GetSigningKeys(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) ([]byte, error) {
	return _Csm.Contract.GetSigningKeys(&_Csm.CallOpts, nodeOperatorId, startIndex, keysCount)
}

// GetSigningKeysWithSignatures is a free data retrieval call binding the contract method 0x50388cb6.
//
// Solidity: function getSigningKeysWithSignatures(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(bytes keys, bytes signatures)
func (_Csm *CsmCaller) GetSigningKeysWithSignatures(opts *bind.CallOpts, nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) (struct {
	Keys       []byte
	Signatures []byte
}, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "getSigningKeysWithSignatures", nodeOperatorId, startIndex, keysCount)

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
func (_Csm *CsmSession) GetSigningKeysWithSignatures(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) (struct {
	Keys       []byte
	Signatures []byte
}, error) {
	return _Csm.Contract.GetSigningKeysWithSignatures(&_Csm.CallOpts, nodeOperatorId, startIndex, keysCount)
}

// GetSigningKeysWithSignatures is a free data retrieval call binding the contract method 0x50388cb6.
//
// Solidity: function getSigningKeysWithSignatures(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(bytes keys, bytes signatures)
func (_Csm *CsmCallerSession) GetSigningKeysWithSignatures(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) (struct {
	Keys       []byte
	Signatures []byte
}, error) {
	return _Csm.Contract.GetSigningKeysWithSignatures(&_Csm.CallOpts, nodeOperatorId, startIndex, keysCount)
}

// GetStakingModuleSummary is a free data retrieval call binding the contract method 0x9abddf09.
//
// Solidity: function getStakingModuleSummary() view returns(uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_Csm *CsmCaller) GetStakingModuleSummary(opts *bind.CallOpts) (struct {
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "getStakingModuleSummary")

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
func (_Csm *CsmSession) GetStakingModuleSummary() (struct {
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	return _Csm.Contract.GetStakingModuleSummary(&_Csm.CallOpts)
}

// GetStakingModuleSummary is a free data retrieval call binding the contract method 0x9abddf09.
//
// Solidity: function getStakingModuleSummary() view returns(uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_Csm *CsmCallerSession) GetStakingModuleSummary() (struct {
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	return _Csm.Contract.GetStakingModuleSummary(&_Csm.CallOpts)
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(bytes32)
func (_Csm *CsmCaller) GetType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "getType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(bytes32)
func (_Csm *CsmSession) GetType() ([32]byte, error) {
	return _Csm.Contract.GetType(&_Csm.CallOpts)
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(bytes32)
func (_Csm *CsmCallerSession) GetType() ([32]byte, error) {
	return _Csm.Contract.GetType(&_Csm.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Csm *CsmCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Csm *CsmSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Csm.Contract.HasRole(&_Csm.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Csm *CsmCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Csm.Contract.HasRole(&_Csm.CallOpts, role, account)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Csm *CsmCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Csm *CsmSession) IsPaused() (bool, error) {
	return _Csm.Contract.IsPaused(&_Csm.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Csm *CsmCallerSession) IsPaused() (bool, error) {
	return _Csm.Contract.IsPaused(&_Csm.CallOpts)
}

// IsValidatorSlashed is a free data retrieval call binding the contract method 0x3dbe8b5a.
//
// Solidity: function isValidatorSlashed(uint256 nodeOperatorId, uint256 keyIndex) view returns(bool)
func (_Csm *CsmCaller) IsValidatorSlashed(opts *bind.CallOpts, nodeOperatorId *big.Int, keyIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "isValidatorSlashed", nodeOperatorId, keyIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorSlashed is a free data retrieval call binding the contract method 0x3dbe8b5a.
//
// Solidity: function isValidatorSlashed(uint256 nodeOperatorId, uint256 keyIndex) view returns(bool)
func (_Csm *CsmSession) IsValidatorSlashed(nodeOperatorId *big.Int, keyIndex *big.Int) (bool, error) {
	return _Csm.Contract.IsValidatorSlashed(&_Csm.CallOpts, nodeOperatorId, keyIndex)
}

// IsValidatorSlashed is a free data retrieval call binding the contract method 0x3dbe8b5a.
//
// Solidity: function isValidatorSlashed(uint256 nodeOperatorId, uint256 keyIndex) view returns(bool)
func (_Csm *CsmCallerSession) IsValidatorSlashed(nodeOperatorId *big.Int, keyIndex *big.Int) (bool, error) {
	return _Csm.Contract.IsValidatorSlashed(&_Csm.CallOpts, nodeOperatorId, keyIndex)
}

// IsValidatorWithdrawn is a free data retrieval call binding the contract method 0x53433643.
//
// Solidity: function isValidatorWithdrawn(uint256 nodeOperatorId, uint256 keyIndex) view returns(bool)
func (_Csm *CsmCaller) IsValidatorWithdrawn(opts *bind.CallOpts, nodeOperatorId *big.Int, keyIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "isValidatorWithdrawn", nodeOperatorId, keyIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorWithdrawn is a free data retrieval call binding the contract method 0x53433643.
//
// Solidity: function isValidatorWithdrawn(uint256 nodeOperatorId, uint256 keyIndex) view returns(bool)
func (_Csm *CsmSession) IsValidatorWithdrawn(nodeOperatorId *big.Int, keyIndex *big.Int) (bool, error) {
	return _Csm.Contract.IsValidatorWithdrawn(&_Csm.CallOpts, nodeOperatorId, keyIndex)
}

// IsValidatorWithdrawn is a free data retrieval call binding the contract method 0x53433643.
//
// Solidity: function isValidatorWithdrawn(uint256 nodeOperatorId, uint256 keyIndex) view returns(bool)
func (_Csm *CsmCallerSession) IsValidatorWithdrawn(nodeOperatorId *big.Int, keyIndex *big.Int) (bool, error) {
	return _Csm.Contract.IsValidatorWithdrawn(&_Csm.CallOpts, nodeOperatorId, keyIndex)
}

// KeyRemovalCharge is a free data retrieval call binding the contract method 0xd9df8c92.
//
// Solidity: function keyRemovalCharge() view returns(uint256)
func (_Csm *CsmCaller) KeyRemovalCharge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "keyRemovalCharge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KeyRemovalCharge is a free data retrieval call binding the contract method 0xd9df8c92.
//
// Solidity: function keyRemovalCharge() view returns(uint256)
func (_Csm *CsmSession) KeyRemovalCharge() (*big.Int, error) {
	return _Csm.Contract.KeyRemovalCharge(&_Csm.CallOpts)
}

// KeyRemovalCharge is a free data retrieval call binding the contract method 0xd9df8c92.
//
// Solidity: function keyRemovalCharge() view returns(uint256)
func (_Csm *CsmCallerSession) KeyRemovalCharge() (*big.Int, error) {
	return _Csm.Contract.KeyRemovalCharge(&_Csm.CallOpts)
}

// PublicRelease is a free data retrieval call binding the contract method 0xe21a430b.
//
// Solidity: function publicRelease() view returns(bool)
func (_Csm *CsmCaller) PublicRelease(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "publicRelease")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PublicRelease is a free data retrieval call binding the contract method 0xe21a430b.
//
// Solidity: function publicRelease() view returns(bool)
func (_Csm *CsmSession) PublicRelease() (bool, error) {
	return _Csm.Contract.PublicRelease(&_Csm.CallOpts)
}

// PublicRelease is a free data retrieval call binding the contract method 0xe21a430b.
//
// Solidity: function publicRelease() view returns(bool)
func (_Csm *CsmCallerSession) PublicRelease() (bool, error) {
	return _Csm.Contract.PublicRelease(&_Csm.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Csm *CsmCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Csm.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Csm *CsmSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Csm.Contract.SupportsInterface(&_Csm.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Csm *CsmCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Csm.Contract.SupportsInterface(&_Csm.CallOpts, interfaceId)
}

// ActivatePublicRelease is a paid mutator transaction binding the contract method 0xd3931457.
//
// Solidity: function activatePublicRelease() returns()
func (_Csm *CsmTransactor) ActivatePublicRelease(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "activatePublicRelease")
}

// ActivatePublicRelease is a paid mutator transaction binding the contract method 0xd3931457.
//
// Solidity: function activatePublicRelease() returns()
func (_Csm *CsmSession) ActivatePublicRelease() (*types.Transaction, error) {
	return _Csm.Contract.ActivatePublicRelease(&_Csm.TransactOpts)
}

// ActivatePublicRelease is a paid mutator transaction binding the contract method 0xd3931457.
//
// Solidity: function activatePublicRelease() returns()
func (_Csm *CsmTransactorSession) ActivatePublicRelease() (*types.Transaction, error) {
	return _Csm.Contract.ActivatePublicRelease(&_Csm.TransactOpts)
}

// AddNodeOperatorETH is a paid mutator transaction binding the contract method 0x157a039b.
//
// Solidity: function addNodeOperatorETH(uint256 keysCount, bytes publicKeys, bytes signatures, (address,address,bool) managementProperties, bytes32[] eaProof, address referrer) payable returns()
func (_Csm *CsmTransactor) AddNodeOperatorETH(opts *bind.TransactOpts, keysCount *big.Int, publicKeys []byte, signatures []byte, managementProperties NodeOperatorManagementProperties, eaProof [][32]byte, referrer common.Address) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "addNodeOperatorETH", keysCount, publicKeys, signatures, managementProperties, eaProof, referrer)
}

// AddNodeOperatorETH is a paid mutator transaction binding the contract method 0x157a039b.
//
// Solidity: function addNodeOperatorETH(uint256 keysCount, bytes publicKeys, bytes signatures, (address,address,bool) managementProperties, bytes32[] eaProof, address referrer) payable returns()
func (_Csm *CsmSession) AddNodeOperatorETH(keysCount *big.Int, publicKeys []byte, signatures []byte, managementProperties NodeOperatorManagementProperties, eaProof [][32]byte, referrer common.Address) (*types.Transaction, error) {
	return _Csm.Contract.AddNodeOperatorETH(&_Csm.TransactOpts, keysCount, publicKeys, signatures, managementProperties, eaProof, referrer)
}

// AddNodeOperatorETH is a paid mutator transaction binding the contract method 0x157a039b.
//
// Solidity: function addNodeOperatorETH(uint256 keysCount, bytes publicKeys, bytes signatures, (address,address,bool) managementProperties, bytes32[] eaProof, address referrer) payable returns()
func (_Csm *CsmTransactorSession) AddNodeOperatorETH(keysCount *big.Int, publicKeys []byte, signatures []byte, managementProperties NodeOperatorManagementProperties, eaProof [][32]byte, referrer common.Address) (*types.Transaction, error) {
	return _Csm.Contract.AddNodeOperatorETH(&_Csm.TransactOpts, keysCount, publicKeys, signatures, managementProperties, eaProof, referrer)
}

// AddNodeOperatorStETH is a paid mutator transaction binding the contract method 0x6a5f2c4a.
//
// Solidity: function addNodeOperatorStETH(uint256 keysCount, bytes publicKeys, bytes signatures, (address,address,bool) managementProperties, (uint256,uint256,uint8,bytes32,bytes32) permit, bytes32[] eaProof, address referrer) returns()
func (_Csm *CsmTransactor) AddNodeOperatorStETH(opts *bind.TransactOpts, keysCount *big.Int, publicKeys []byte, signatures []byte, managementProperties NodeOperatorManagementProperties, permit ICSAccountingPermitInput, eaProof [][32]byte, referrer common.Address) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "addNodeOperatorStETH", keysCount, publicKeys, signatures, managementProperties, permit, eaProof, referrer)
}

// AddNodeOperatorStETH is a paid mutator transaction binding the contract method 0x6a5f2c4a.
//
// Solidity: function addNodeOperatorStETH(uint256 keysCount, bytes publicKeys, bytes signatures, (address,address,bool) managementProperties, (uint256,uint256,uint8,bytes32,bytes32) permit, bytes32[] eaProof, address referrer) returns()
func (_Csm *CsmSession) AddNodeOperatorStETH(keysCount *big.Int, publicKeys []byte, signatures []byte, managementProperties NodeOperatorManagementProperties, permit ICSAccountingPermitInput, eaProof [][32]byte, referrer common.Address) (*types.Transaction, error) {
	return _Csm.Contract.AddNodeOperatorStETH(&_Csm.TransactOpts, keysCount, publicKeys, signatures, managementProperties, permit, eaProof, referrer)
}

// AddNodeOperatorStETH is a paid mutator transaction binding the contract method 0x6a5f2c4a.
//
// Solidity: function addNodeOperatorStETH(uint256 keysCount, bytes publicKeys, bytes signatures, (address,address,bool) managementProperties, (uint256,uint256,uint8,bytes32,bytes32) permit, bytes32[] eaProof, address referrer) returns()
func (_Csm *CsmTransactorSession) AddNodeOperatorStETH(keysCount *big.Int, publicKeys []byte, signatures []byte, managementProperties NodeOperatorManagementProperties, permit ICSAccountingPermitInput, eaProof [][32]byte, referrer common.Address) (*types.Transaction, error) {
	return _Csm.Contract.AddNodeOperatorStETH(&_Csm.TransactOpts, keysCount, publicKeys, signatures, managementProperties, permit, eaProof, referrer)
}

// AddNodeOperatorWstETH is a paid mutator transaction binding the contract method 0xacc446eb.
//
// Solidity: function addNodeOperatorWstETH(uint256 keysCount, bytes publicKeys, bytes signatures, (address,address,bool) managementProperties, (uint256,uint256,uint8,bytes32,bytes32) permit, bytes32[] eaProof, address referrer) returns()
func (_Csm *CsmTransactor) AddNodeOperatorWstETH(opts *bind.TransactOpts, keysCount *big.Int, publicKeys []byte, signatures []byte, managementProperties NodeOperatorManagementProperties, permit ICSAccountingPermitInput, eaProof [][32]byte, referrer common.Address) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "addNodeOperatorWstETH", keysCount, publicKeys, signatures, managementProperties, permit, eaProof, referrer)
}

// AddNodeOperatorWstETH is a paid mutator transaction binding the contract method 0xacc446eb.
//
// Solidity: function addNodeOperatorWstETH(uint256 keysCount, bytes publicKeys, bytes signatures, (address,address,bool) managementProperties, (uint256,uint256,uint8,bytes32,bytes32) permit, bytes32[] eaProof, address referrer) returns()
func (_Csm *CsmSession) AddNodeOperatorWstETH(keysCount *big.Int, publicKeys []byte, signatures []byte, managementProperties NodeOperatorManagementProperties, permit ICSAccountingPermitInput, eaProof [][32]byte, referrer common.Address) (*types.Transaction, error) {
	return _Csm.Contract.AddNodeOperatorWstETH(&_Csm.TransactOpts, keysCount, publicKeys, signatures, managementProperties, permit, eaProof, referrer)
}

// AddNodeOperatorWstETH is a paid mutator transaction binding the contract method 0xacc446eb.
//
// Solidity: function addNodeOperatorWstETH(uint256 keysCount, bytes publicKeys, bytes signatures, (address,address,bool) managementProperties, (uint256,uint256,uint8,bytes32,bytes32) permit, bytes32[] eaProof, address referrer) returns()
func (_Csm *CsmTransactorSession) AddNodeOperatorWstETH(keysCount *big.Int, publicKeys []byte, signatures []byte, managementProperties NodeOperatorManagementProperties, permit ICSAccountingPermitInput, eaProof [][32]byte, referrer common.Address) (*types.Transaction, error) {
	return _Csm.Contract.AddNodeOperatorWstETH(&_Csm.TransactOpts, keysCount, publicKeys, signatures, managementProperties, permit, eaProof, referrer)
}

// AddValidatorKeysETH is a paid mutator transaction binding the contract method 0xfe7ed3cd.
//
// Solidity: function addValidatorKeysETH(uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures) payable returns()
func (_Csm *CsmTransactor) AddValidatorKeysETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "addValidatorKeysETH", nodeOperatorId, keysCount, publicKeys, signatures)
}

// AddValidatorKeysETH is a paid mutator transaction binding the contract method 0xfe7ed3cd.
//
// Solidity: function addValidatorKeysETH(uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures) payable returns()
func (_Csm *CsmSession) AddValidatorKeysETH(nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte) (*types.Transaction, error) {
	return _Csm.Contract.AddValidatorKeysETH(&_Csm.TransactOpts, nodeOperatorId, keysCount, publicKeys, signatures)
}

// AddValidatorKeysETH is a paid mutator transaction binding the contract method 0xfe7ed3cd.
//
// Solidity: function addValidatorKeysETH(uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures) payable returns()
func (_Csm *CsmTransactorSession) AddValidatorKeysETH(nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte) (*types.Transaction, error) {
	return _Csm.Contract.AddValidatorKeysETH(&_Csm.TransactOpts, nodeOperatorId, keysCount, publicKeys, signatures)
}

// AddValidatorKeysStETH is a paid mutator transaction binding the contract method 0x946654ad.
//
// Solidity: function addValidatorKeysStETH(uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csm *CsmTransactor) AddValidatorKeysStETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "addValidatorKeysStETH", nodeOperatorId, keysCount, publicKeys, signatures, permit)
}

// AddValidatorKeysStETH is a paid mutator transaction binding the contract method 0x946654ad.
//
// Solidity: function addValidatorKeysStETH(uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csm *CsmSession) AddValidatorKeysStETH(nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csm.Contract.AddValidatorKeysStETH(&_Csm.TransactOpts, nodeOperatorId, keysCount, publicKeys, signatures, permit)
}

// AddValidatorKeysStETH is a paid mutator transaction binding the contract method 0x946654ad.
//
// Solidity: function addValidatorKeysStETH(uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csm *CsmTransactorSession) AddValidatorKeysStETH(nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csm.Contract.AddValidatorKeysStETH(&_Csm.TransactOpts, nodeOperatorId, keysCount, publicKeys, signatures, permit)
}

// AddValidatorKeysWstETH is a paid mutator transaction binding the contract method 0x9ec3c24c.
//
// Solidity: function addValidatorKeysWstETH(uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csm *CsmTransactor) AddValidatorKeysWstETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "addValidatorKeysWstETH", nodeOperatorId, keysCount, publicKeys, signatures, permit)
}

// AddValidatorKeysWstETH is a paid mutator transaction binding the contract method 0x9ec3c24c.
//
// Solidity: function addValidatorKeysWstETH(uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csm *CsmSession) AddValidatorKeysWstETH(nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csm.Contract.AddValidatorKeysWstETH(&_Csm.TransactOpts, nodeOperatorId, keysCount, publicKeys, signatures, permit)
}

// AddValidatorKeysWstETH is a paid mutator transaction binding the contract method 0x9ec3c24c.
//
// Solidity: function addValidatorKeysWstETH(uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csm *CsmTransactorSession) AddValidatorKeysWstETH(nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csm.Contract.AddValidatorKeysWstETH(&_Csm.TransactOpts, nodeOperatorId, keysCount, publicKeys, signatures, permit)
}

// CancelELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x40044801.
//
// Solidity: function cancelELRewardsStealingPenalty(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csm *CsmTransactor) CancelELRewardsStealingPenalty(opts *bind.TransactOpts, nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "cancelELRewardsStealingPenalty", nodeOperatorId, amount)
}

// CancelELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x40044801.
//
// Solidity: function cancelELRewardsStealingPenalty(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csm *CsmSession) CancelELRewardsStealingPenalty(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.CancelELRewardsStealingPenalty(&_Csm.TransactOpts, nodeOperatorId, amount)
}

// CancelELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x40044801.
//
// Solidity: function cancelELRewardsStealingPenalty(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csm *CsmTransactorSession) CancelELRewardsStealingPenalty(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.CancelELRewardsStealingPenalty(&_Csm.TransactOpts, nodeOperatorId, amount)
}

// ChangeNodeOperatorRewardAddress is a paid mutator transaction binding the contract method 0x75a401da.
//
// Solidity: function changeNodeOperatorRewardAddress(uint256 nodeOperatorId, address newAddress) returns()
func (_Csm *CsmTransactor) ChangeNodeOperatorRewardAddress(opts *bind.TransactOpts, nodeOperatorId *big.Int, newAddress common.Address) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "changeNodeOperatorRewardAddress", nodeOperatorId, newAddress)
}

// ChangeNodeOperatorRewardAddress is a paid mutator transaction binding the contract method 0x75a401da.
//
// Solidity: function changeNodeOperatorRewardAddress(uint256 nodeOperatorId, address newAddress) returns()
func (_Csm *CsmSession) ChangeNodeOperatorRewardAddress(nodeOperatorId *big.Int, newAddress common.Address) (*types.Transaction, error) {
	return _Csm.Contract.ChangeNodeOperatorRewardAddress(&_Csm.TransactOpts, nodeOperatorId, newAddress)
}

// ChangeNodeOperatorRewardAddress is a paid mutator transaction binding the contract method 0x75a401da.
//
// Solidity: function changeNodeOperatorRewardAddress(uint256 nodeOperatorId, address newAddress) returns()
func (_Csm *CsmTransactorSession) ChangeNodeOperatorRewardAddress(nodeOperatorId *big.Int, newAddress common.Address) (*types.Transaction, error) {
	return _Csm.Contract.ChangeNodeOperatorRewardAddress(&_Csm.TransactOpts, nodeOperatorId, newAddress)
}

// ClaimRewardsStETH is a paid mutator transaction binding the contract method 0x8409d4fe.
//
// Solidity: function claimRewardsStETH(uint256 nodeOperatorId, uint256 stETHAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_Csm *CsmTransactor) ClaimRewardsStETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, stETHAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "claimRewardsStETH", nodeOperatorId, stETHAmount, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsStETH is a paid mutator transaction binding the contract method 0x8409d4fe.
//
// Solidity: function claimRewardsStETH(uint256 nodeOperatorId, uint256 stETHAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_Csm *CsmSession) ClaimRewardsStETH(nodeOperatorId *big.Int, stETHAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csm.Contract.ClaimRewardsStETH(&_Csm.TransactOpts, nodeOperatorId, stETHAmount, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsStETH is a paid mutator transaction binding the contract method 0x8409d4fe.
//
// Solidity: function claimRewardsStETH(uint256 nodeOperatorId, uint256 stETHAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_Csm *CsmTransactorSession) ClaimRewardsStETH(nodeOperatorId *big.Int, stETHAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csm.Contract.ClaimRewardsStETH(&_Csm.TransactOpts, nodeOperatorId, stETHAmount, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsUnstETH is a paid mutator transaction binding the contract method 0x3df6c438.
//
// Solidity: function claimRewardsUnstETH(uint256 nodeOperatorId, uint256 stEthAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_Csm *CsmTransactor) ClaimRewardsUnstETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, stEthAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "claimRewardsUnstETH", nodeOperatorId, stEthAmount, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsUnstETH is a paid mutator transaction binding the contract method 0x3df6c438.
//
// Solidity: function claimRewardsUnstETH(uint256 nodeOperatorId, uint256 stEthAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_Csm *CsmSession) ClaimRewardsUnstETH(nodeOperatorId *big.Int, stEthAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csm.Contract.ClaimRewardsUnstETH(&_Csm.TransactOpts, nodeOperatorId, stEthAmount, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsUnstETH is a paid mutator transaction binding the contract method 0x3df6c438.
//
// Solidity: function claimRewardsUnstETH(uint256 nodeOperatorId, uint256 stEthAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_Csm *CsmTransactorSession) ClaimRewardsUnstETH(nodeOperatorId *big.Int, stEthAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csm.Contract.ClaimRewardsUnstETH(&_Csm.TransactOpts, nodeOperatorId, stEthAmount, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsWstETH is a paid mutator transaction binding the contract method 0x5097ef59.
//
// Solidity: function claimRewardsWstETH(uint256 nodeOperatorId, uint256 wstETHAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_Csm *CsmTransactor) ClaimRewardsWstETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, wstETHAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "claimRewardsWstETH", nodeOperatorId, wstETHAmount, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsWstETH is a paid mutator transaction binding the contract method 0x5097ef59.
//
// Solidity: function claimRewardsWstETH(uint256 nodeOperatorId, uint256 wstETHAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_Csm *CsmSession) ClaimRewardsWstETH(nodeOperatorId *big.Int, wstETHAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csm.Contract.ClaimRewardsWstETH(&_Csm.TransactOpts, nodeOperatorId, wstETHAmount, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsWstETH is a paid mutator transaction binding the contract method 0x5097ef59.
//
// Solidity: function claimRewardsWstETH(uint256 nodeOperatorId, uint256 wstETHAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_Csm *CsmTransactorSession) ClaimRewardsWstETH(nodeOperatorId *big.Int, wstETHAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csm.Contract.ClaimRewardsWstETH(&_Csm.TransactOpts, nodeOperatorId, wstETHAmount, cumulativeFeeShares, rewardsProof)
}

// CleanDepositQueue is a paid mutator transaction binding the contract method 0x735dfa28.
//
// Solidity: function cleanDepositQueue(uint256 maxItems) returns(uint256 removed, uint256 lastRemovedAtDepth)
func (_Csm *CsmTransactor) CleanDepositQueue(opts *bind.TransactOpts, maxItems *big.Int) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "cleanDepositQueue", maxItems)
}

// CleanDepositQueue is a paid mutator transaction binding the contract method 0x735dfa28.
//
// Solidity: function cleanDepositQueue(uint256 maxItems) returns(uint256 removed, uint256 lastRemovedAtDepth)
func (_Csm *CsmSession) CleanDepositQueue(maxItems *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.CleanDepositQueue(&_Csm.TransactOpts, maxItems)
}

// CleanDepositQueue is a paid mutator transaction binding the contract method 0x735dfa28.
//
// Solidity: function cleanDepositQueue(uint256 maxItems) returns(uint256 removed, uint256 lastRemovedAtDepth)
func (_Csm *CsmTransactorSession) CleanDepositQueue(maxItems *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.CleanDepositQueue(&_Csm.TransactOpts, maxItems)
}

// CompensateELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x6efe37a2.
//
// Solidity: function compensateELRewardsStealingPenalty(uint256 nodeOperatorId) payable returns()
func (_Csm *CsmTransactor) CompensateELRewardsStealingPenalty(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "compensateELRewardsStealingPenalty", nodeOperatorId)
}

// CompensateELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x6efe37a2.
//
// Solidity: function compensateELRewardsStealingPenalty(uint256 nodeOperatorId) payable returns()
func (_Csm *CsmSession) CompensateELRewardsStealingPenalty(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.CompensateELRewardsStealingPenalty(&_Csm.TransactOpts, nodeOperatorId)
}

// CompensateELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x6efe37a2.
//
// Solidity: function compensateELRewardsStealingPenalty(uint256 nodeOperatorId) payable returns()
func (_Csm *CsmTransactorSession) CompensateELRewardsStealingPenalty(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.CompensateELRewardsStealingPenalty(&_Csm.TransactOpts, nodeOperatorId)
}

// ConfirmNodeOperatorManagerAddressChange is a paid mutator transaction binding the contract method 0x6bb1bfdf.
//
// Solidity: function confirmNodeOperatorManagerAddressChange(uint256 nodeOperatorId) returns()
func (_Csm *CsmTransactor) ConfirmNodeOperatorManagerAddressChange(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "confirmNodeOperatorManagerAddressChange", nodeOperatorId)
}

// ConfirmNodeOperatorManagerAddressChange is a paid mutator transaction binding the contract method 0x6bb1bfdf.
//
// Solidity: function confirmNodeOperatorManagerAddressChange(uint256 nodeOperatorId) returns()
func (_Csm *CsmSession) ConfirmNodeOperatorManagerAddressChange(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.ConfirmNodeOperatorManagerAddressChange(&_Csm.TransactOpts, nodeOperatorId)
}

// ConfirmNodeOperatorManagerAddressChange is a paid mutator transaction binding the contract method 0x6bb1bfdf.
//
// Solidity: function confirmNodeOperatorManagerAddressChange(uint256 nodeOperatorId) returns()
func (_Csm *CsmTransactorSession) ConfirmNodeOperatorManagerAddressChange(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.ConfirmNodeOperatorManagerAddressChange(&_Csm.TransactOpts, nodeOperatorId)
}

// ConfirmNodeOperatorRewardAddressChange is a paid mutator transaction binding the contract method 0x5204281c.
//
// Solidity: function confirmNodeOperatorRewardAddressChange(uint256 nodeOperatorId) returns()
func (_Csm *CsmTransactor) ConfirmNodeOperatorRewardAddressChange(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "confirmNodeOperatorRewardAddressChange", nodeOperatorId)
}

// ConfirmNodeOperatorRewardAddressChange is a paid mutator transaction binding the contract method 0x5204281c.
//
// Solidity: function confirmNodeOperatorRewardAddressChange(uint256 nodeOperatorId) returns()
func (_Csm *CsmSession) ConfirmNodeOperatorRewardAddressChange(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.ConfirmNodeOperatorRewardAddressChange(&_Csm.TransactOpts, nodeOperatorId)
}

// ConfirmNodeOperatorRewardAddressChange is a paid mutator transaction binding the contract method 0x5204281c.
//
// Solidity: function confirmNodeOperatorRewardAddressChange(uint256 nodeOperatorId) returns()
func (_Csm *CsmTransactorSession) ConfirmNodeOperatorRewardAddressChange(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.ConfirmNodeOperatorRewardAddressChange(&_Csm.TransactOpts, nodeOperatorId)
}

// DecreaseVettedSigningKeysCount is a paid mutator transaction binding the contract method 0xb643189b.
//
// Solidity: function decreaseVettedSigningKeysCount(bytes nodeOperatorIds, bytes vettedSigningKeysCounts) returns()
func (_Csm *CsmTransactor) DecreaseVettedSigningKeysCount(opts *bind.TransactOpts, nodeOperatorIds []byte, vettedSigningKeysCounts []byte) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "decreaseVettedSigningKeysCount", nodeOperatorIds, vettedSigningKeysCounts)
}

// DecreaseVettedSigningKeysCount is a paid mutator transaction binding the contract method 0xb643189b.
//
// Solidity: function decreaseVettedSigningKeysCount(bytes nodeOperatorIds, bytes vettedSigningKeysCounts) returns()
func (_Csm *CsmSession) DecreaseVettedSigningKeysCount(nodeOperatorIds []byte, vettedSigningKeysCounts []byte) (*types.Transaction, error) {
	return _Csm.Contract.DecreaseVettedSigningKeysCount(&_Csm.TransactOpts, nodeOperatorIds, vettedSigningKeysCounts)
}

// DecreaseVettedSigningKeysCount is a paid mutator transaction binding the contract method 0xb643189b.
//
// Solidity: function decreaseVettedSigningKeysCount(bytes nodeOperatorIds, bytes vettedSigningKeysCounts) returns()
func (_Csm *CsmTransactorSession) DecreaseVettedSigningKeysCount(nodeOperatorIds []byte, vettedSigningKeysCounts []byte) (*types.Transaction, error) {
	return _Csm.Contract.DecreaseVettedSigningKeysCount(&_Csm.TransactOpts, nodeOperatorIds, vettedSigningKeysCounts)
}

// DepositETH is a paid mutator transaction binding the contract method 0x5358fbda.
//
// Solidity: function depositETH(uint256 nodeOperatorId) payable returns()
func (_Csm *CsmTransactor) DepositETH(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "depositETH", nodeOperatorId)
}

// DepositETH is a paid mutator transaction binding the contract method 0x5358fbda.
//
// Solidity: function depositETH(uint256 nodeOperatorId) payable returns()
func (_Csm *CsmSession) DepositETH(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.DepositETH(&_Csm.TransactOpts, nodeOperatorId)
}

// DepositETH is a paid mutator transaction binding the contract method 0x5358fbda.
//
// Solidity: function depositETH(uint256 nodeOperatorId) payable returns()
func (_Csm *CsmTransactorSession) DepositETH(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.DepositETH(&_Csm.TransactOpts, nodeOperatorId)
}

// DepositStETH is a paid mutator transaction binding the contract method 0xe1aa105d.
//
// Solidity: function depositStETH(uint256 nodeOperatorId, uint256 stETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csm *CsmTransactor) DepositStETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, stETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "depositStETH", nodeOperatorId, stETHAmount, permit)
}

// DepositStETH is a paid mutator transaction binding the contract method 0xe1aa105d.
//
// Solidity: function depositStETH(uint256 nodeOperatorId, uint256 stETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csm *CsmSession) DepositStETH(nodeOperatorId *big.Int, stETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csm.Contract.DepositStETH(&_Csm.TransactOpts, nodeOperatorId, stETHAmount, permit)
}

// DepositStETH is a paid mutator transaction binding the contract method 0xe1aa105d.
//
// Solidity: function depositStETH(uint256 nodeOperatorId, uint256 stETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csm *CsmTransactorSession) DepositStETH(nodeOperatorId *big.Int, stETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csm.Contract.DepositStETH(&_Csm.TransactOpts, nodeOperatorId, stETHAmount, permit)
}

// DepositWstETH is a paid mutator transaction binding the contract method 0x3f214bb2.
//
// Solidity: function depositWstETH(uint256 nodeOperatorId, uint256 wstETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csm *CsmTransactor) DepositWstETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, wstETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "depositWstETH", nodeOperatorId, wstETHAmount, permit)
}

// DepositWstETH is a paid mutator transaction binding the contract method 0x3f214bb2.
//
// Solidity: function depositWstETH(uint256 nodeOperatorId, uint256 wstETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csm *CsmSession) DepositWstETH(nodeOperatorId *big.Int, wstETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csm.Contract.DepositWstETH(&_Csm.TransactOpts, nodeOperatorId, wstETHAmount, permit)
}

// DepositWstETH is a paid mutator transaction binding the contract method 0x3f214bb2.
//
// Solidity: function depositWstETH(uint256 nodeOperatorId, uint256 wstETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csm *CsmTransactorSession) DepositWstETH(nodeOperatorId *big.Int, wstETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csm.Contract.DepositWstETH(&_Csm.TransactOpts, nodeOperatorId, wstETHAmount, permit)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Csm *CsmTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Csm *CsmSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csm.Contract.GrantRole(&_Csm.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Csm *CsmTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csm.Contract.GrantRole(&_Csm.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe203094.
//
// Solidity: function initialize(address _accounting, address _earlyAdoption, uint256 _keyRemovalCharge, address admin) returns()
func (_Csm *CsmTransactor) Initialize(opts *bind.TransactOpts, _accounting common.Address, _earlyAdoption common.Address, _keyRemovalCharge *big.Int, admin common.Address) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "initialize", _accounting, _earlyAdoption, _keyRemovalCharge, admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe203094.
//
// Solidity: function initialize(address _accounting, address _earlyAdoption, uint256 _keyRemovalCharge, address admin) returns()
func (_Csm *CsmSession) Initialize(_accounting common.Address, _earlyAdoption common.Address, _keyRemovalCharge *big.Int, admin common.Address) (*types.Transaction, error) {
	return _Csm.Contract.Initialize(&_Csm.TransactOpts, _accounting, _earlyAdoption, _keyRemovalCharge, admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe203094.
//
// Solidity: function initialize(address _accounting, address _earlyAdoption, uint256 _keyRemovalCharge, address admin) returns()
func (_Csm *CsmTransactorSession) Initialize(_accounting common.Address, _earlyAdoption common.Address, _keyRemovalCharge *big.Int, admin common.Address) (*types.Transaction, error) {
	return _Csm.Contract.Initialize(&_Csm.TransactOpts, _accounting, _earlyAdoption, _keyRemovalCharge, admin)
}

// NormalizeQueue is a paid mutator transaction binding the contract method 0xb1520dc5.
//
// Solidity: function normalizeQueue(uint256 nodeOperatorId) returns()
func (_Csm *CsmTransactor) NormalizeQueue(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "normalizeQueue", nodeOperatorId)
}

// NormalizeQueue is a paid mutator transaction binding the contract method 0xb1520dc5.
//
// Solidity: function normalizeQueue(uint256 nodeOperatorId) returns()
func (_Csm *CsmSession) NormalizeQueue(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.NormalizeQueue(&_Csm.TransactOpts, nodeOperatorId)
}

// NormalizeQueue is a paid mutator transaction binding the contract method 0xb1520dc5.
//
// Solidity: function normalizeQueue(uint256 nodeOperatorId) returns()
func (_Csm *CsmTransactorSession) NormalizeQueue(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.NormalizeQueue(&_Csm.TransactOpts, nodeOperatorId)
}

// ObtainDepositData is a paid mutator transaction binding the contract method 0xbee41b58.
//
// Solidity: function obtainDepositData(uint256 depositsCount, bytes ) returns(bytes publicKeys, bytes signatures)
func (_Csm *CsmTransactor) ObtainDepositData(opts *bind.TransactOpts, depositsCount *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "obtainDepositData", depositsCount, arg1)
}

// ObtainDepositData is a paid mutator transaction binding the contract method 0xbee41b58.
//
// Solidity: function obtainDepositData(uint256 depositsCount, bytes ) returns(bytes publicKeys, bytes signatures)
func (_Csm *CsmSession) ObtainDepositData(depositsCount *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _Csm.Contract.ObtainDepositData(&_Csm.TransactOpts, depositsCount, arg1)
}

// ObtainDepositData is a paid mutator transaction binding the contract method 0xbee41b58.
//
// Solidity: function obtainDepositData(uint256 depositsCount, bytes ) returns(bytes publicKeys, bytes signatures)
func (_Csm *CsmTransactorSession) ObtainDepositData(depositsCount *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _Csm.Contract.ObtainDepositData(&_Csm.TransactOpts, depositsCount, arg1)
}

// OnExitedAndStuckValidatorsCountsUpdated is a paid mutator transaction binding the contract method 0xe864299e.
//
// Solidity: function onExitedAndStuckValidatorsCountsUpdated() returns()
func (_Csm *CsmTransactor) OnExitedAndStuckValidatorsCountsUpdated(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "onExitedAndStuckValidatorsCountsUpdated")
}

// OnExitedAndStuckValidatorsCountsUpdated is a paid mutator transaction binding the contract method 0xe864299e.
//
// Solidity: function onExitedAndStuckValidatorsCountsUpdated() returns()
func (_Csm *CsmSession) OnExitedAndStuckValidatorsCountsUpdated() (*types.Transaction, error) {
	return _Csm.Contract.OnExitedAndStuckValidatorsCountsUpdated(&_Csm.TransactOpts)
}

// OnExitedAndStuckValidatorsCountsUpdated is a paid mutator transaction binding the contract method 0xe864299e.
//
// Solidity: function onExitedAndStuckValidatorsCountsUpdated() returns()
func (_Csm *CsmTransactorSession) OnExitedAndStuckValidatorsCountsUpdated() (*types.Transaction, error) {
	return _Csm.Contract.OnExitedAndStuckValidatorsCountsUpdated(&_Csm.TransactOpts)
}

// OnRewardsMinted is a paid mutator transaction binding the contract method 0x8d7e4017.
//
// Solidity: function onRewardsMinted(uint256 totalShares) returns()
func (_Csm *CsmTransactor) OnRewardsMinted(opts *bind.TransactOpts, totalShares *big.Int) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "onRewardsMinted", totalShares)
}

// OnRewardsMinted is a paid mutator transaction binding the contract method 0x8d7e4017.
//
// Solidity: function onRewardsMinted(uint256 totalShares) returns()
func (_Csm *CsmSession) OnRewardsMinted(totalShares *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.OnRewardsMinted(&_Csm.TransactOpts, totalShares)
}

// OnRewardsMinted is a paid mutator transaction binding the contract method 0x8d7e4017.
//
// Solidity: function onRewardsMinted(uint256 totalShares) returns()
func (_Csm *CsmTransactorSession) OnRewardsMinted(totalShares *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.OnRewardsMinted(&_Csm.TransactOpts, totalShares)
}

// OnWithdrawalCredentialsChanged is a paid mutator transaction binding the contract method 0x90c09bdb.
//
// Solidity: function onWithdrawalCredentialsChanged() returns()
func (_Csm *CsmTransactor) OnWithdrawalCredentialsChanged(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "onWithdrawalCredentialsChanged")
}

// OnWithdrawalCredentialsChanged is a paid mutator transaction binding the contract method 0x90c09bdb.
//
// Solidity: function onWithdrawalCredentialsChanged() returns()
func (_Csm *CsmSession) OnWithdrawalCredentialsChanged() (*types.Transaction, error) {
	return _Csm.Contract.OnWithdrawalCredentialsChanged(&_Csm.TransactOpts)
}

// OnWithdrawalCredentialsChanged is a paid mutator transaction binding the contract method 0x90c09bdb.
//
// Solidity: function onWithdrawalCredentialsChanged() returns()
func (_Csm *CsmTransactorSession) OnWithdrawalCredentialsChanged() (*types.Transaction, error) {
	return _Csm.Contract.OnWithdrawalCredentialsChanged(&_Csm.TransactOpts)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 duration) returns()
func (_Csm *CsmTransactor) PauseFor(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "pauseFor", duration)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 duration) returns()
func (_Csm *CsmSession) PauseFor(duration *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.PauseFor(&_Csm.TransactOpts, duration)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 duration) returns()
func (_Csm *CsmTransactorSession) PauseFor(duration *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.PauseFor(&_Csm.TransactOpts, duration)
}

// ProposeNodeOperatorManagerAddressChange is a paid mutator transaction binding the contract method 0x8cabe959.
//
// Solidity: function proposeNodeOperatorManagerAddressChange(uint256 nodeOperatorId, address proposedAddress) returns()
func (_Csm *CsmTransactor) ProposeNodeOperatorManagerAddressChange(opts *bind.TransactOpts, nodeOperatorId *big.Int, proposedAddress common.Address) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "proposeNodeOperatorManagerAddressChange", nodeOperatorId, proposedAddress)
}

// ProposeNodeOperatorManagerAddressChange is a paid mutator transaction binding the contract method 0x8cabe959.
//
// Solidity: function proposeNodeOperatorManagerAddressChange(uint256 nodeOperatorId, address proposedAddress) returns()
func (_Csm *CsmSession) ProposeNodeOperatorManagerAddressChange(nodeOperatorId *big.Int, proposedAddress common.Address) (*types.Transaction, error) {
	return _Csm.Contract.ProposeNodeOperatorManagerAddressChange(&_Csm.TransactOpts, nodeOperatorId, proposedAddress)
}

// ProposeNodeOperatorManagerAddressChange is a paid mutator transaction binding the contract method 0x8cabe959.
//
// Solidity: function proposeNodeOperatorManagerAddressChange(uint256 nodeOperatorId, address proposedAddress) returns()
func (_Csm *CsmTransactorSession) ProposeNodeOperatorManagerAddressChange(nodeOperatorId *big.Int, proposedAddress common.Address) (*types.Transaction, error) {
	return _Csm.Contract.ProposeNodeOperatorManagerAddressChange(&_Csm.TransactOpts, nodeOperatorId, proposedAddress)
}

// ProposeNodeOperatorRewardAddressChange is a paid mutator transaction binding the contract method 0x1b40b231.
//
// Solidity: function proposeNodeOperatorRewardAddressChange(uint256 nodeOperatorId, address proposedAddress) returns()
func (_Csm *CsmTransactor) ProposeNodeOperatorRewardAddressChange(opts *bind.TransactOpts, nodeOperatorId *big.Int, proposedAddress common.Address) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "proposeNodeOperatorRewardAddressChange", nodeOperatorId, proposedAddress)
}

// ProposeNodeOperatorRewardAddressChange is a paid mutator transaction binding the contract method 0x1b40b231.
//
// Solidity: function proposeNodeOperatorRewardAddressChange(uint256 nodeOperatorId, address proposedAddress) returns()
func (_Csm *CsmSession) ProposeNodeOperatorRewardAddressChange(nodeOperatorId *big.Int, proposedAddress common.Address) (*types.Transaction, error) {
	return _Csm.Contract.ProposeNodeOperatorRewardAddressChange(&_Csm.TransactOpts, nodeOperatorId, proposedAddress)
}

// ProposeNodeOperatorRewardAddressChange is a paid mutator transaction binding the contract method 0x1b40b231.
//
// Solidity: function proposeNodeOperatorRewardAddressChange(uint256 nodeOperatorId, address proposedAddress) returns()
func (_Csm *CsmTransactorSession) ProposeNodeOperatorRewardAddressChange(nodeOperatorId *big.Int, proposedAddress common.Address) (*types.Transaction, error) {
	return _Csm.Contract.ProposeNodeOperatorRewardAddressChange(&_Csm.TransactOpts, nodeOperatorId, proposedAddress)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_Csm *CsmTransactor) RecoverERC1155(opts *bind.TransactOpts, token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "recoverERC1155", token, tokenId)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_Csm *CsmSession) RecoverERC1155(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.RecoverERC1155(&_Csm.TransactOpts, token, tokenId)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_Csm *CsmTransactorSession) RecoverERC1155(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.RecoverERC1155(&_Csm.TransactOpts, token, tokenId)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_Csm *CsmTransactor) RecoverERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "recoverERC20", token, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_Csm *CsmSession) RecoverERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.RecoverERC20(&_Csm.TransactOpts, token, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_Csm *CsmTransactorSession) RecoverERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.RecoverERC20(&_Csm.TransactOpts, token, amount)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_Csm *CsmTransactor) RecoverERC721(opts *bind.TransactOpts, token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "recoverERC721", token, tokenId)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_Csm *CsmSession) RecoverERC721(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.RecoverERC721(&_Csm.TransactOpts, token, tokenId)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_Csm *CsmTransactorSession) RecoverERC721(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.RecoverERC721(&_Csm.TransactOpts, token, tokenId)
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_Csm *CsmTransactor) RecoverEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "recoverEther")
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_Csm *CsmSession) RecoverEther() (*types.Transaction, error) {
	return _Csm.Contract.RecoverEther(&_Csm.TransactOpts)
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_Csm *CsmTransactorSession) RecoverEther() (*types.Transaction, error) {
	return _Csm.Contract.RecoverEther(&_Csm.TransactOpts)
}

// RemoveKeys is a paid mutator transaction binding the contract method 0x8b3ac71d.
//
// Solidity: function removeKeys(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) returns()
func (_Csm *CsmTransactor) RemoveKeys(opts *bind.TransactOpts, nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "removeKeys", nodeOperatorId, startIndex, keysCount)
}

// RemoveKeys is a paid mutator transaction binding the contract method 0x8b3ac71d.
//
// Solidity: function removeKeys(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) returns()
func (_Csm *CsmSession) RemoveKeys(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.RemoveKeys(&_Csm.TransactOpts, nodeOperatorId, startIndex, keysCount)
}

// RemoveKeys is a paid mutator transaction binding the contract method 0x8b3ac71d.
//
// Solidity: function removeKeys(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) returns()
func (_Csm *CsmTransactorSession) RemoveKeys(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.RemoveKeys(&_Csm.TransactOpts, nodeOperatorId, startIndex, keysCount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Csm *CsmTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Csm *CsmSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Csm.Contract.RenounceRole(&_Csm.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Csm *CsmTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Csm.Contract.RenounceRole(&_Csm.TransactOpts, role, callerConfirmation)
}

// ReportELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x388dd1d1.
//
// Solidity: function reportELRewardsStealingPenalty(uint256 nodeOperatorId, bytes32 blockHash, uint256 amount) returns()
func (_Csm *CsmTransactor) ReportELRewardsStealingPenalty(opts *bind.TransactOpts, nodeOperatorId *big.Int, blockHash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "reportELRewardsStealingPenalty", nodeOperatorId, blockHash, amount)
}

// ReportELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x388dd1d1.
//
// Solidity: function reportELRewardsStealingPenalty(uint256 nodeOperatorId, bytes32 blockHash, uint256 amount) returns()
func (_Csm *CsmSession) ReportELRewardsStealingPenalty(nodeOperatorId *big.Int, blockHash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.ReportELRewardsStealingPenalty(&_Csm.TransactOpts, nodeOperatorId, blockHash, amount)
}

// ReportELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x388dd1d1.
//
// Solidity: function reportELRewardsStealingPenalty(uint256 nodeOperatorId, bytes32 blockHash, uint256 amount) returns()
func (_Csm *CsmTransactorSession) ReportELRewardsStealingPenalty(nodeOperatorId *big.Int, blockHash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.ReportELRewardsStealingPenalty(&_Csm.TransactOpts, nodeOperatorId, blockHash, amount)
}

// ResetNodeOperatorManagerAddress is a paid mutator transaction binding the contract method 0x6a6304cc.
//
// Solidity: function resetNodeOperatorManagerAddress(uint256 nodeOperatorId) returns()
func (_Csm *CsmTransactor) ResetNodeOperatorManagerAddress(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "resetNodeOperatorManagerAddress", nodeOperatorId)
}

// ResetNodeOperatorManagerAddress is a paid mutator transaction binding the contract method 0x6a6304cc.
//
// Solidity: function resetNodeOperatorManagerAddress(uint256 nodeOperatorId) returns()
func (_Csm *CsmSession) ResetNodeOperatorManagerAddress(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.ResetNodeOperatorManagerAddress(&_Csm.TransactOpts, nodeOperatorId)
}

// ResetNodeOperatorManagerAddress is a paid mutator transaction binding the contract method 0x6a6304cc.
//
// Solidity: function resetNodeOperatorManagerAddress(uint256 nodeOperatorId) returns()
func (_Csm *CsmTransactorSession) ResetNodeOperatorManagerAddress(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.ResetNodeOperatorManagerAddress(&_Csm.TransactOpts, nodeOperatorId)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Csm *CsmTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Csm *CsmSession) Resume() (*types.Transaction, error) {
	return _Csm.Contract.Resume(&_Csm.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Csm *CsmTransactorSession) Resume() (*types.Transaction, error) {
	return _Csm.Contract.Resume(&_Csm.TransactOpts)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Csm *CsmTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Csm *CsmSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csm.Contract.RevokeRole(&_Csm.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Csm *CsmTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csm.Contract.RevokeRole(&_Csm.TransactOpts, role, account)
}

// SetKeyRemovalCharge is a paid mutator transaction binding the contract method 0xba1557ae.
//
// Solidity: function setKeyRemovalCharge(uint256 amount) returns()
func (_Csm *CsmTransactor) SetKeyRemovalCharge(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "setKeyRemovalCharge", amount)
}

// SetKeyRemovalCharge is a paid mutator transaction binding the contract method 0xba1557ae.
//
// Solidity: function setKeyRemovalCharge(uint256 amount) returns()
func (_Csm *CsmSession) SetKeyRemovalCharge(amount *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.SetKeyRemovalCharge(&_Csm.TransactOpts, amount)
}

// SetKeyRemovalCharge is a paid mutator transaction binding the contract method 0xba1557ae.
//
// Solidity: function setKeyRemovalCharge(uint256 amount) returns()
func (_Csm *CsmTransactorSession) SetKeyRemovalCharge(amount *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.SetKeyRemovalCharge(&_Csm.TransactOpts, amount)
}

// SettleELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x37b12b5f.
//
// Solidity: function settleELRewardsStealingPenalty(uint256[] nodeOperatorIds) returns()
func (_Csm *CsmTransactor) SettleELRewardsStealingPenalty(opts *bind.TransactOpts, nodeOperatorIds []*big.Int) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "settleELRewardsStealingPenalty", nodeOperatorIds)
}

// SettleELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x37b12b5f.
//
// Solidity: function settleELRewardsStealingPenalty(uint256[] nodeOperatorIds) returns()
func (_Csm *CsmSession) SettleELRewardsStealingPenalty(nodeOperatorIds []*big.Int) (*types.Transaction, error) {
	return _Csm.Contract.SettleELRewardsStealingPenalty(&_Csm.TransactOpts, nodeOperatorIds)
}

// SettleELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x37b12b5f.
//
// Solidity: function settleELRewardsStealingPenalty(uint256[] nodeOperatorIds) returns()
func (_Csm *CsmTransactorSession) SettleELRewardsStealingPenalty(nodeOperatorIds []*big.Int) (*types.Transaction, error) {
	return _Csm.Contract.SettleELRewardsStealingPenalty(&_Csm.TransactOpts, nodeOperatorIds)
}

// SubmitInitialSlashing is a paid mutator transaction binding the contract method 0xf96d3952.
//
// Solidity: function submitInitialSlashing(uint256 nodeOperatorId, uint256 keyIndex) returns()
func (_Csm *CsmTransactor) SubmitInitialSlashing(opts *bind.TransactOpts, nodeOperatorId *big.Int, keyIndex *big.Int) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "submitInitialSlashing", nodeOperatorId, keyIndex)
}

// SubmitInitialSlashing is a paid mutator transaction binding the contract method 0xf96d3952.
//
// Solidity: function submitInitialSlashing(uint256 nodeOperatorId, uint256 keyIndex) returns()
func (_Csm *CsmSession) SubmitInitialSlashing(nodeOperatorId *big.Int, keyIndex *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.SubmitInitialSlashing(&_Csm.TransactOpts, nodeOperatorId, keyIndex)
}

// SubmitInitialSlashing is a paid mutator transaction binding the contract method 0xf96d3952.
//
// Solidity: function submitInitialSlashing(uint256 nodeOperatorId, uint256 keyIndex) returns()
func (_Csm *CsmTransactorSession) SubmitInitialSlashing(nodeOperatorId *big.Int, keyIndex *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.SubmitInitialSlashing(&_Csm.TransactOpts, nodeOperatorId, keyIndex)
}

// SubmitWithdrawal is a paid mutator transaction binding the contract method 0xf408b551.
//
// Solidity: function submitWithdrawal(uint256 nodeOperatorId, uint256 keyIndex, uint256 amount, bool isSlashed) returns()
func (_Csm *CsmTransactor) SubmitWithdrawal(opts *bind.TransactOpts, nodeOperatorId *big.Int, keyIndex *big.Int, amount *big.Int, isSlashed bool) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "submitWithdrawal", nodeOperatorId, keyIndex, amount, isSlashed)
}

// SubmitWithdrawal is a paid mutator transaction binding the contract method 0xf408b551.
//
// Solidity: function submitWithdrawal(uint256 nodeOperatorId, uint256 keyIndex, uint256 amount, bool isSlashed) returns()
func (_Csm *CsmSession) SubmitWithdrawal(nodeOperatorId *big.Int, keyIndex *big.Int, amount *big.Int, isSlashed bool) (*types.Transaction, error) {
	return _Csm.Contract.SubmitWithdrawal(&_Csm.TransactOpts, nodeOperatorId, keyIndex, amount, isSlashed)
}

// SubmitWithdrawal is a paid mutator transaction binding the contract method 0xf408b551.
//
// Solidity: function submitWithdrawal(uint256 nodeOperatorId, uint256 keyIndex, uint256 amount, bool isSlashed) returns()
func (_Csm *CsmTransactorSession) SubmitWithdrawal(nodeOperatorId *big.Int, keyIndex *big.Int, amount *big.Int, isSlashed bool) (*types.Transaction, error) {
	return _Csm.Contract.SubmitWithdrawal(&_Csm.TransactOpts, nodeOperatorId, keyIndex, amount, isSlashed)
}

// UnsafeUpdateValidatorsCount is a paid mutator transaction binding the contract method 0xf2e2ca63.
//
// Solidity: function unsafeUpdateValidatorsCount(uint256 nodeOperatorId, uint256 exitedValidatorsKeysCount, uint256 stuckValidatorsKeysCount) returns()
func (_Csm *CsmTransactor) UnsafeUpdateValidatorsCount(opts *bind.TransactOpts, nodeOperatorId *big.Int, exitedValidatorsKeysCount *big.Int, stuckValidatorsKeysCount *big.Int) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "unsafeUpdateValidatorsCount", nodeOperatorId, exitedValidatorsKeysCount, stuckValidatorsKeysCount)
}

// UnsafeUpdateValidatorsCount is a paid mutator transaction binding the contract method 0xf2e2ca63.
//
// Solidity: function unsafeUpdateValidatorsCount(uint256 nodeOperatorId, uint256 exitedValidatorsKeysCount, uint256 stuckValidatorsKeysCount) returns()
func (_Csm *CsmSession) UnsafeUpdateValidatorsCount(nodeOperatorId *big.Int, exitedValidatorsKeysCount *big.Int, stuckValidatorsKeysCount *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.UnsafeUpdateValidatorsCount(&_Csm.TransactOpts, nodeOperatorId, exitedValidatorsKeysCount, stuckValidatorsKeysCount)
}

// UnsafeUpdateValidatorsCount is a paid mutator transaction binding the contract method 0xf2e2ca63.
//
// Solidity: function unsafeUpdateValidatorsCount(uint256 nodeOperatorId, uint256 exitedValidatorsKeysCount, uint256 stuckValidatorsKeysCount) returns()
func (_Csm *CsmTransactorSession) UnsafeUpdateValidatorsCount(nodeOperatorId *big.Int, exitedValidatorsKeysCount *big.Int, stuckValidatorsKeysCount *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.UnsafeUpdateValidatorsCount(&_Csm.TransactOpts, nodeOperatorId, exitedValidatorsKeysCount, stuckValidatorsKeysCount)
}

// UpdateExitedValidatorsCount is a paid mutator transaction binding the contract method 0x9b00c146.
//
// Solidity: function updateExitedValidatorsCount(bytes nodeOperatorIds, bytes exitedValidatorsCounts) returns()
func (_Csm *CsmTransactor) UpdateExitedValidatorsCount(opts *bind.TransactOpts, nodeOperatorIds []byte, exitedValidatorsCounts []byte) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "updateExitedValidatorsCount", nodeOperatorIds, exitedValidatorsCounts)
}

// UpdateExitedValidatorsCount is a paid mutator transaction binding the contract method 0x9b00c146.
//
// Solidity: function updateExitedValidatorsCount(bytes nodeOperatorIds, bytes exitedValidatorsCounts) returns()
func (_Csm *CsmSession) UpdateExitedValidatorsCount(nodeOperatorIds []byte, exitedValidatorsCounts []byte) (*types.Transaction, error) {
	return _Csm.Contract.UpdateExitedValidatorsCount(&_Csm.TransactOpts, nodeOperatorIds, exitedValidatorsCounts)
}

// UpdateExitedValidatorsCount is a paid mutator transaction binding the contract method 0x9b00c146.
//
// Solidity: function updateExitedValidatorsCount(bytes nodeOperatorIds, bytes exitedValidatorsCounts) returns()
func (_Csm *CsmTransactorSession) UpdateExitedValidatorsCount(nodeOperatorIds []byte, exitedValidatorsCounts []byte) (*types.Transaction, error) {
	return _Csm.Contract.UpdateExitedValidatorsCount(&_Csm.TransactOpts, nodeOperatorIds, exitedValidatorsCounts)
}

// UpdateRefundedValidatorsCount is a paid mutator transaction binding the contract method 0xa2e080f1.
//
// Solidity: function updateRefundedValidatorsCount(uint256 , uint256 ) returns()
func (_Csm *CsmTransactor) UpdateRefundedValidatorsCount(opts *bind.TransactOpts, arg0 *big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "updateRefundedValidatorsCount", arg0, arg1)
}

// UpdateRefundedValidatorsCount is a paid mutator transaction binding the contract method 0xa2e080f1.
//
// Solidity: function updateRefundedValidatorsCount(uint256 , uint256 ) returns()
func (_Csm *CsmSession) UpdateRefundedValidatorsCount(arg0 *big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.UpdateRefundedValidatorsCount(&_Csm.TransactOpts, arg0, arg1)
}

// UpdateRefundedValidatorsCount is a paid mutator transaction binding the contract method 0xa2e080f1.
//
// Solidity: function updateRefundedValidatorsCount(uint256 , uint256 ) returns()
func (_Csm *CsmTransactorSession) UpdateRefundedValidatorsCount(arg0 *big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.UpdateRefundedValidatorsCount(&_Csm.TransactOpts, arg0, arg1)
}

// UpdateStuckValidatorsCount is a paid mutator transaction binding the contract method 0x9b3d1900.
//
// Solidity: function updateStuckValidatorsCount(bytes nodeOperatorIds, bytes stuckValidatorsCounts) returns()
func (_Csm *CsmTransactor) UpdateStuckValidatorsCount(opts *bind.TransactOpts, nodeOperatorIds []byte, stuckValidatorsCounts []byte) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "updateStuckValidatorsCount", nodeOperatorIds, stuckValidatorsCounts)
}

// UpdateStuckValidatorsCount is a paid mutator transaction binding the contract method 0x9b3d1900.
//
// Solidity: function updateStuckValidatorsCount(bytes nodeOperatorIds, bytes stuckValidatorsCounts) returns()
func (_Csm *CsmSession) UpdateStuckValidatorsCount(nodeOperatorIds []byte, stuckValidatorsCounts []byte) (*types.Transaction, error) {
	return _Csm.Contract.UpdateStuckValidatorsCount(&_Csm.TransactOpts, nodeOperatorIds, stuckValidatorsCounts)
}

// UpdateStuckValidatorsCount is a paid mutator transaction binding the contract method 0x9b3d1900.
//
// Solidity: function updateStuckValidatorsCount(bytes nodeOperatorIds, bytes stuckValidatorsCounts) returns()
func (_Csm *CsmTransactorSession) UpdateStuckValidatorsCount(nodeOperatorIds []byte, stuckValidatorsCounts []byte) (*types.Transaction, error) {
	return _Csm.Contract.UpdateStuckValidatorsCount(&_Csm.TransactOpts, nodeOperatorIds, stuckValidatorsCounts)
}

// UpdateTargetValidatorsLimits is a paid mutator transaction binding the contract method 0x08a679ad.
//
// Solidity: function updateTargetValidatorsLimits(uint256 nodeOperatorId, uint256 targetLimitMode, uint256 targetLimit) returns()
func (_Csm *CsmTransactor) UpdateTargetValidatorsLimits(opts *bind.TransactOpts, nodeOperatorId *big.Int, targetLimitMode *big.Int, targetLimit *big.Int) (*types.Transaction, error) {
	return _Csm.contract.Transact(opts, "updateTargetValidatorsLimits", nodeOperatorId, targetLimitMode, targetLimit)
}

// UpdateTargetValidatorsLimits is a paid mutator transaction binding the contract method 0x08a679ad.
//
// Solidity: function updateTargetValidatorsLimits(uint256 nodeOperatorId, uint256 targetLimitMode, uint256 targetLimit) returns()
func (_Csm *CsmSession) UpdateTargetValidatorsLimits(nodeOperatorId *big.Int, targetLimitMode *big.Int, targetLimit *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.UpdateTargetValidatorsLimits(&_Csm.TransactOpts, nodeOperatorId, targetLimitMode, targetLimit)
}

// UpdateTargetValidatorsLimits is a paid mutator transaction binding the contract method 0x08a679ad.
//
// Solidity: function updateTargetValidatorsLimits(uint256 nodeOperatorId, uint256 targetLimitMode, uint256 targetLimit) returns()
func (_Csm *CsmTransactorSession) UpdateTargetValidatorsLimits(nodeOperatorId *big.Int, targetLimitMode *big.Int, targetLimit *big.Int) (*types.Transaction, error) {
	return _Csm.Contract.UpdateTargetValidatorsLimits(&_Csm.TransactOpts, nodeOperatorId, targetLimitMode, targetLimit)
}

// CsmBatchEnqueuedIterator is returned from FilterBatchEnqueued and is used to iterate over the raw logs and unpacked data for BatchEnqueued events raised by the Csm contract.
type CsmBatchEnqueuedIterator struct {
	Event *CsmBatchEnqueued // Event containing the contract specifics and raw log

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
func (it *CsmBatchEnqueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmBatchEnqueued)
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
		it.Event = new(CsmBatchEnqueued)
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
func (it *CsmBatchEnqueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmBatchEnqueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmBatchEnqueued represents a BatchEnqueued event raised by the Csm contract.
type CsmBatchEnqueued struct {
	NodeOperatorId *big.Int
	Count          *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBatchEnqueued is a free log retrieval operation binding the contract event 0x162b3db9d9ca7d0abe51ad8229dc058550a74b769457fd055579b5bdc5492536.
//
// Solidity: event BatchEnqueued(uint256 indexed nodeOperatorId, uint256 count)
func (_Csm *CsmFilterer) FilterBatchEnqueued(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmBatchEnqueuedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.FilterLogs(opts, "BatchEnqueued", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmBatchEnqueuedIterator{contract: _Csm.contract, event: "BatchEnqueued", logs: logs, sub: sub}, nil
}

// WatchBatchEnqueued is a free log subscription operation binding the contract event 0x162b3db9d9ca7d0abe51ad8229dc058550a74b769457fd055579b5bdc5492536.
//
// Solidity: event BatchEnqueued(uint256 indexed nodeOperatorId, uint256 count)
func (_Csm *CsmFilterer) WatchBatchEnqueued(opts *bind.WatchOpts, sink chan<- *CsmBatchEnqueued, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.WatchLogs(opts, "BatchEnqueued", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmBatchEnqueued)
				if err := _Csm.contract.UnpackLog(event, "BatchEnqueued", log); err != nil {
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

// ParseBatchEnqueued is a log parse operation binding the contract event 0x162b3db9d9ca7d0abe51ad8229dc058550a74b769457fd055579b5bdc5492536.
//
// Solidity: event BatchEnqueued(uint256 indexed nodeOperatorId, uint256 count)
func (_Csm *CsmFilterer) ParseBatchEnqueued(log types.Log) (*CsmBatchEnqueued, error) {
	event := new(CsmBatchEnqueued)
	if err := _Csm.contract.UnpackLog(event, "BatchEnqueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmDepositableSigningKeysCountChangedIterator is returned from FilterDepositableSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for DepositableSigningKeysCountChanged events raised by the Csm contract.
type CsmDepositableSigningKeysCountChangedIterator struct {
	Event *CsmDepositableSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *CsmDepositableSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmDepositableSigningKeysCountChanged)
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
		it.Event = new(CsmDepositableSigningKeysCountChanged)
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
func (it *CsmDepositableSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmDepositableSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmDepositableSigningKeysCountChanged represents a DepositableSigningKeysCountChanged event raised by the Csm contract.
type CsmDepositableSigningKeysCountChanged struct {
	NodeOperatorId       *big.Int
	DepositableKeysCount *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterDepositableSigningKeysCountChanged is a free log retrieval operation binding the contract event 0xf9109091b368cedad2edff45414eef892edd6b4fe80084bd590aa8f8def8ed33.
//
// Solidity: event DepositableSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 depositableKeysCount)
func (_Csm *CsmFilterer) FilterDepositableSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmDepositableSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.FilterLogs(opts, "DepositableSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmDepositableSigningKeysCountChangedIterator{contract: _Csm.contract, event: "DepositableSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchDepositableSigningKeysCountChanged is a free log subscription operation binding the contract event 0xf9109091b368cedad2edff45414eef892edd6b4fe80084bd590aa8f8def8ed33.
//
// Solidity: event DepositableSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 depositableKeysCount)
func (_Csm *CsmFilterer) WatchDepositableSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *CsmDepositableSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.WatchLogs(opts, "DepositableSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmDepositableSigningKeysCountChanged)
				if err := _Csm.contract.UnpackLog(event, "DepositableSigningKeysCountChanged", log); err != nil {
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
func (_Csm *CsmFilterer) ParseDepositableSigningKeysCountChanged(log types.Log) (*CsmDepositableSigningKeysCountChanged, error) {
	event := new(CsmDepositableSigningKeysCountChanged)
	if err := _Csm.contract.UnpackLog(event, "DepositableSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmDepositedSigningKeysCountChangedIterator is returned from FilterDepositedSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for DepositedSigningKeysCountChanged events raised by the Csm contract.
type CsmDepositedSigningKeysCountChangedIterator struct {
	Event *CsmDepositedSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *CsmDepositedSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmDepositedSigningKeysCountChanged)
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
		it.Event = new(CsmDepositedSigningKeysCountChanged)
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
func (it *CsmDepositedSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmDepositedSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmDepositedSigningKeysCountChanged represents a DepositedSigningKeysCountChanged event raised by the Csm contract.
type CsmDepositedSigningKeysCountChanged struct {
	NodeOperatorId     *big.Int
	DepositedKeysCount *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDepositedSigningKeysCountChanged is a free log retrieval operation binding the contract event 0x24eb1c9e765ba41accf9437300ea91ece5ed3f897ec3cdee0e9debd7fe309b78.
//
// Solidity: event DepositedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 depositedKeysCount)
func (_Csm *CsmFilterer) FilterDepositedSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmDepositedSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.FilterLogs(opts, "DepositedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmDepositedSigningKeysCountChangedIterator{contract: _Csm.contract, event: "DepositedSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchDepositedSigningKeysCountChanged is a free log subscription operation binding the contract event 0x24eb1c9e765ba41accf9437300ea91ece5ed3f897ec3cdee0e9debd7fe309b78.
//
// Solidity: event DepositedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 depositedKeysCount)
func (_Csm *CsmFilterer) WatchDepositedSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *CsmDepositedSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.WatchLogs(opts, "DepositedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmDepositedSigningKeysCountChanged)
				if err := _Csm.contract.UnpackLog(event, "DepositedSigningKeysCountChanged", log); err != nil {
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
func (_Csm *CsmFilterer) ParseDepositedSigningKeysCountChanged(log types.Log) (*CsmDepositedSigningKeysCountChanged, error) {
	event := new(CsmDepositedSigningKeysCountChanged)
	if err := _Csm.contract.UnpackLog(event, "DepositedSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmELRewardsStealingPenaltyCancelledIterator is returned from FilterELRewardsStealingPenaltyCancelled and is used to iterate over the raw logs and unpacked data for ELRewardsStealingPenaltyCancelled events raised by the Csm contract.
type CsmELRewardsStealingPenaltyCancelledIterator struct {
	Event *CsmELRewardsStealingPenaltyCancelled // Event containing the contract specifics and raw log

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
func (it *CsmELRewardsStealingPenaltyCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmELRewardsStealingPenaltyCancelled)
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
		it.Event = new(CsmELRewardsStealingPenaltyCancelled)
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
func (it *CsmELRewardsStealingPenaltyCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmELRewardsStealingPenaltyCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmELRewardsStealingPenaltyCancelled represents a ELRewardsStealingPenaltyCancelled event raised by the Csm contract.
type CsmELRewardsStealingPenaltyCancelled struct {
	NodeOperatorId *big.Int
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterELRewardsStealingPenaltyCancelled is a free log retrieval operation binding the contract event 0x1e7ebd3c5f4de9502000b6f7e6e7cf5d4ecb27d6fe1778e43fb9d1d0ca87d0e7.
//
// Solidity: event ELRewardsStealingPenaltyCancelled(uint256 indexed nodeOperatorId, uint256 amount)
func (_Csm *CsmFilterer) FilterELRewardsStealingPenaltyCancelled(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmELRewardsStealingPenaltyCancelledIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.FilterLogs(opts, "ELRewardsStealingPenaltyCancelled", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmELRewardsStealingPenaltyCancelledIterator{contract: _Csm.contract, event: "ELRewardsStealingPenaltyCancelled", logs: logs, sub: sub}, nil
}

// WatchELRewardsStealingPenaltyCancelled is a free log subscription operation binding the contract event 0x1e7ebd3c5f4de9502000b6f7e6e7cf5d4ecb27d6fe1778e43fb9d1d0ca87d0e7.
//
// Solidity: event ELRewardsStealingPenaltyCancelled(uint256 indexed nodeOperatorId, uint256 amount)
func (_Csm *CsmFilterer) WatchELRewardsStealingPenaltyCancelled(opts *bind.WatchOpts, sink chan<- *CsmELRewardsStealingPenaltyCancelled, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.WatchLogs(opts, "ELRewardsStealingPenaltyCancelled", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmELRewardsStealingPenaltyCancelled)
				if err := _Csm.contract.UnpackLog(event, "ELRewardsStealingPenaltyCancelled", log); err != nil {
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

// ParseELRewardsStealingPenaltyCancelled is a log parse operation binding the contract event 0x1e7ebd3c5f4de9502000b6f7e6e7cf5d4ecb27d6fe1778e43fb9d1d0ca87d0e7.
//
// Solidity: event ELRewardsStealingPenaltyCancelled(uint256 indexed nodeOperatorId, uint256 amount)
func (_Csm *CsmFilterer) ParseELRewardsStealingPenaltyCancelled(log types.Log) (*CsmELRewardsStealingPenaltyCancelled, error) {
	event := new(CsmELRewardsStealingPenaltyCancelled)
	if err := _Csm.contract.UnpackLog(event, "ELRewardsStealingPenaltyCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmELRewardsStealingPenaltyCompensatedIterator is returned from FilterELRewardsStealingPenaltyCompensated and is used to iterate over the raw logs and unpacked data for ELRewardsStealingPenaltyCompensated events raised by the Csm contract.
type CsmELRewardsStealingPenaltyCompensatedIterator struct {
	Event *CsmELRewardsStealingPenaltyCompensated // Event containing the contract specifics and raw log

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
func (it *CsmELRewardsStealingPenaltyCompensatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmELRewardsStealingPenaltyCompensated)
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
		it.Event = new(CsmELRewardsStealingPenaltyCompensated)
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
func (it *CsmELRewardsStealingPenaltyCompensatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmELRewardsStealingPenaltyCompensatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmELRewardsStealingPenaltyCompensated represents a ELRewardsStealingPenaltyCompensated event raised by the Csm contract.
type CsmELRewardsStealingPenaltyCompensated struct {
	NodeOperatorId *big.Int
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterELRewardsStealingPenaltyCompensated is a free log retrieval operation binding the contract event 0xb1858b4c2ab6242521725a8f7350a6cb22ad4ecae009c9b63ef114baffb054be.
//
// Solidity: event ELRewardsStealingPenaltyCompensated(uint256 indexed nodeOperatorId, uint256 amount)
func (_Csm *CsmFilterer) FilterELRewardsStealingPenaltyCompensated(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmELRewardsStealingPenaltyCompensatedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.FilterLogs(opts, "ELRewardsStealingPenaltyCompensated", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmELRewardsStealingPenaltyCompensatedIterator{contract: _Csm.contract, event: "ELRewardsStealingPenaltyCompensated", logs: logs, sub: sub}, nil
}

// WatchELRewardsStealingPenaltyCompensated is a free log subscription operation binding the contract event 0xb1858b4c2ab6242521725a8f7350a6cb22ad4ecae009c9b63ef114baffb054be.
//
// Solidity: event ELRewardsStealingPenaltyCompensated(uint256 indexed nodeOperatorId, uint256 amount)
func (_Csm *CsmFilterer) WatchELRewardsStealingPenaltyCompensated(opts *bind.WatchOpts, sink chan<- *CsmELRewardsStealingPenaltyCompensated, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.WatchLogs(opts, "ELRewardsStealingPenaltyCompensated", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmELRewardsStealingPenaltyCompensated)
				if err := _Csm.contract.UnpackLog(event, "ELRewardsStealingPenaltyCompensated", log); err != nil {
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

// ParseELRewardsStealingPenaltyCompensated is a log parse operation binding the contract event 0xb1858b4c2ab6242521725a8f7350a6cb22ad4ecae009c9b63ef114baffb054be.
//
// Solidity: event ELRewardsStealingPenaltyCompensated(uint256 indexed nodeOperatorId, uint256 amount)
func (_Csm *CsmFilterer) ParseELRewardsStealingPenaltyCompensated(log types.Log) (*CsmELRewardsStealingPenaltyCompensated, error) {
	event := new(CsmELRewardsStealingPenaltyCompensated)
	if err := _Csm.contract.UnpackLog(event, "ELRewardsStealingPenaltyCompensated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmELRewardsStealingPenaltyReportedIterator is returned from FilterELRewardsStealingPenaltyReported and is used to iterate over the raw logs and unpacked data for ELRewardsStealingPenaltyReported events raised by the Csm contract.
type CsmELRewardsStealingPenaltyReportedIterator struct {
	Event *CsmELRewardsStealingPenaltyReported // Event containing the contract specifics and raw log

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
func (it *CsmELRewardsStealingPenaltyReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmELRewardsStealingPenaltyReported)
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
		it.Event = new(CsmELRewardsStealingPenaltyReported)
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
func (it *CsmELRewardsStealingPenaltyReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmELRewardsStealingPenaltyReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmELRewardsStealingPenaltyReported represents a ELRewardsStealingPenaltyReported event raised by the Csm contract.
type CsmELRewardsStealingPenaltyReported struct {
	NodeOperatorId    *big.Int
	ProposedBlockHash [32]byte
	StolenAmount      *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterELRewardsStealingPenaltyReported is a free log retrieval operation binding the contract event 0xeec4d6dbe34149c6728a9638eca869d0e5a7fcd85c7a96178f7e9780b4b7fe4b.
//
// Solidity: event ELRewardsStealingPenaltyReported(uint256 indexed nodeOperatorId, bytes32 proposedBlockHash, uint256 stolenAmount)
func (_Csm *CsmFilterer) FilterELRewardsStealingPenaltyReported(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmELRewardsStealingPenaltyReportedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.FilterLogs(opts, "ELRewardsStealingPenaltyReported", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmELRewardsStealingPenaltyReportedIterator{contract: _Csm.contract, event: "ELRewardsStealingPenaltyReported", logs: logs, sub: sub}, nil
}

// WatchELRewardsStealingPenaltyReported is a free log subscription operation binding the contract event 0xeec4d6dbe34149c6728a9638eca869d0e5a7fcd85c7a96178f7e9780b4b7fe4b.
//
// Solidity: event ELRewardsStealingPenaltyReported(uint256 indexed nodeOperatorId, bytes32 proposedBlockHash, uint256 stolenAmount)
func (_Csm *CsmFilterer) WatchELRewardsStealingPenaltyReported(opts *bind.WatchOpts, sink chan<- *CsmELRewardsStealingPenaltyReported, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.WatchLogs(opts, "ELRewardsStealingPenaltyReported", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmELRewardsStealingPenaltyReported)
				if err := _Csm.contract.UnpackLog(event, "ELRewardsStealingPenaltyReported", log); err != nil {
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

// ParseELRewardsStealingPenaltyReported is a log parse operation binding the contract event 0xeec4d6dbe34149c6728a9638eca869d0e5a7fcd85c7a96178f7e9780b4b7fe4b.
//
// Solidity: event ELRewardsStealingPenaltyReported(uint256 indexed nodeOperatorId, bytes32 proposedBlockHash, uint256 stolenAmount)
func (_Csm *CsmFilterer) ParseELRewardsStealingPenaltyReported(log types.Log) (*CsmELRewardsStealingPenaltyReported, error) {
	event := new(CsmELRewardsStealingPenaltyReported)
	if err := _Csm.contract.UnpackLog(event, "ELRewardsStealingPenaltyReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmELRewardsStealingPenaltySettledIterator is returned from FilterELRewardsStealingPenaltySettled and is used to iterate over the raw logs and unpacked data for ELRewardsStealingPenaltySettled events raised by the Csm contract.
type CsmELRewardsStealingPenaltySettledIterator struct {
	Event *CsmELRewardsStealingPenaltySettled // Event containing the contract specifics and raw log

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
func (it *CsmELRewardsStealingPenaltySettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmELRewardsStealingPenaltySettled)
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
		it.Event = new(CsmELRewardsStealingPenaltySettled)
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
func (it *CsmELRewardsStealingPenaltySettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmELRewardsStealingPenaltySettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmELRewardsStealingPenaltySettled represents a ELRewardsStealingPenaltySettled event raised by the Csm contract.
type CsmELRewardsStealingPenaltySettled struct {
	NodeOperatorId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterELRewardsStealingPenaltySettled is a free log retrieval operation binding the contract event 0x00f4fe19c0404d2fbb58da6f646c0a3ee5a6994a034213bbd22b072ed1ca5c27.
//
// Solidity: event ELRewardsStealingPenaltySettled(uint256 indexed nodeOperatorId)
func (_Csm *CsmFilterer) FilterELRewardsStealingPenaltySettled(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmELRewardsStealingPenaltySettledIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.FilterLogs(opts, "ELRewardsStealingPenaltySettled", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmELRewardsStealingPenaltySettledIterator{contract: _Csm.contract, event: "ELRewardsStealingPenaltySettled", logs: logs, sub: sub}, nil
}

// WatchELRewardsStealingPenaltySettled is a free log subscription operation binding the contract event 0x00f4fe19c0404d2fbb58da6f646c0a3ee5a6994a034213bbd22b072ed1ca5c27.
//
// Solidity: event ELRewardsStealingPenaltySettled(uint256 indexed nodeOperatorId)
func (_Csm *CsmFilterer) WatchELRewardsStealingPenaltySettled(opts *bind.WatchOpts, sink chan<- *CsmELRewardsStealingPenaltySettled, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.WatchLogs(opts, "ELRewardsStealingPenaltySettled", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmELRewardsStealingPenaltySettled)
				if err := _Csm.contract.UnpackLog(event, "ELRewardsStealingPenaltySettled", log); err != nil {
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

// ParseELRewardsStealingPenaltySettled is a log parse operation binding the contract event 0x00f4fe19c0404d2fbb58da6f646c0a3ee5a6994a034213bbd22b072ed1ca5c27.
//
// Solidity: event ELRewardsStealingPenaltySettled(uint256 indexed nodeOperatorId)
func (_Csm *CsmFilterer) ParseELRewardsStealingPenaltySettled(log types.Log) (*CsmELRewardsStealingPenaltySettled, error) {
	event := new(CsmELRewardsStealingPenaltySettled)
	if err := _Csm.contract.UnpackLog(event, "ELRewardsStealingPenaltySettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmERC1155RecoveredIterator is returned from FilterERC1155Recovered and is used to iterate over the raw logs and unpacked data for ERC1155Recovered events raised by the Csm contract.
type CsmERC1155RecoveredIterator struct {
	Event *CsmERC1155Recovered // Event containing the contract specifics and raw log

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
func (it *CsmERC1155RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmERC1155Recovered)
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
		it.Event = new(CsmERC1155Recovered)
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
func (it *CsmERC1155RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmERC1155RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmERC1155Recovered represents a ERC1155Recovered event raised by the Csm contract.
type CsmERC1155Recovered struct {
	Token     common.Address
	TokenId   *big.Int
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC1155Recovered is a free log retrieval operation binding the contract event 0x5cf02e753b3eb0f4bee4460a72817d8e5e3c75cd4d65c1d0b06dca88b8032936.
//
// Solidity: event ERC1155Recovered(address indexed token, uint256 tokenId, address indexed recipient, uint256 amount)
func (_Csm *CsmFilterer) FilterERC1155Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CsmERC1155RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csm.contract.FilterLogs(opts, "ERC1155Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsmERC1155RecoveredIterator{contract: _Csm.contract, event: "ERC1155Recovered", logs: logs, sub: sub}, nil
}

// WatchERC1155Recovered is a free log subscription operation binding the contract event 0x5cf02e753b3eb0f4bee4460a72817d8e5e3c75cd4d65c1d0b06dca88b8032936.
//
// Solidity: event ERC1155Recovered(address indexed token, uint256 tokenId, address indexed recipient, uint256 amount)
func (_Csm *CsmFilterer) WatchERC1155Recovered(opts *bind.WatchOpts, sink chan<- *CsmERC1155Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csm.contract.WatchLogs(opts, "ERC1155Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmERC1155Recovered)
				if err := _Csm.contract.UnpackLog(event, "ERC1155Recovered", log); err != nil {
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
func (_Csm *CsmFilterer) ParseERC1155Recovered(log types.Log) (*CsmERC1155Recovered, error) {
	event := new(CsmERC1155Recovered)
	if err := _Csm.contract.UnpackLog(event, "ERC1155Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmERC20RecoveredIterator is returned from FilterERC20Recovered and is used to iterate over the raw logs and unpacked data for ERC20Recovered events raised by the Csm contract.
type CsmERC20RecoveredIterator struct {
	Event *CsmERC20Recovered // Event containing the contract specifics and raw log

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
func (it *CsmERC20RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmERC20Recovered)
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
		it.Event = new(CsmERC20Recovered)
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
func (it *CsmERC20RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmERC20RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmERC20Recovered represents a ERC20Recovered event raised by the Csm contract.
type CsmERC20Recovered struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC20Recovered is a free log retrieval operation binding the contract event 0xaca8fb252cde442184e5f10e0f2e6e4029e8cd7717cae63559079610702436aa.
//
// Solidity: event ERC20Recovered(address indexed token, address indexed recipient, uint256 amount)
func (_Csm *CsmFilterer) FilterERC20Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CsmERC20RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csm.contract.FilterLogs(opts, "ERC20Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsmERC20RecoveredIterator{contract: _Csm.contract, event: "ERC20Recovered", logs: logs, sub: sub}, nil
}

// WatchERC20Recovered is a free log subscription operation binding the contract event 0xaca8fb252cde442184e5f10e0f2e6e4029e8cd7717cae63559079610702436aa.
//
// Solidity: event ERC20Recovered(address indexed token, address indexed recipient, uint256 amount)
func (_Csm *CsmFilterer) WatchERC20Recovered(opts *bind.WatchOpts, sink chan<- *CsmERC20Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csm.contract.WatchLogs(opts, "ERC20Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmERC20Recovered)
				if err := _Csm.contract.UnpackLog(event, "ERC20Recovered", log); err != nil {
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
func (_Csm *CsmFilterer) ParseERC20Recovered(log types.Log) (*CsmERC20Recovered, error) {
	event := new(CsmERC20Recovered)
	if err := _Csm.contract.UnpackLog(event, "ERC20Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmERC721RecoveredIterator is returned from FilterERC721Recovered and is used to iterate over the raw logs and unpacked data for ERC721Recovered events raised by the Csm contract.
type CsmERC721RecoveredIterator struct {
	Event *CsmERC721Recovered // Event containing the contract specifics and raw log

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
func (it *CsmERC721RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmERC721Recovered)
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
		it.Event = new(CsmERC721Recovered)
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
func (it *CsmERC721RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmERC721RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmERC721Recovered represents a ERC721Recovered event raised by the Csm contract.
type CsmERC721Recovered struct {
	Token     common.Address
	TokenId   *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC721Recovered is a free log retrieval operation binding the contract event 0x8166bf75d2ff2fa3c8f3c44410540bf42e9a5359b48409e8d660291dc9f788c8.
//
// Solidity: event ERC721Recovered(address indexed token, uint256 tokenId, address indexed recipient)
func (_Csm *CsmFilterer) FilterERC721Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CsmERC721RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csm.contract.FilterLogs(opts, "ERC721Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsmERC721RecoveredIterator{contract: _Csm.contract, event: "ERC721Recovered", logs: logs, sub: sub}, nil
}

// WatchERC721Recovered is a free log subscription operation binding the contract event 0x8166bf75d2ff2fa3c8f3c44410540bf42e9a5359b48409e8d660291dc9f788c8.
//
// Solidity: event ERC721Recovered(address indexed token, uint256 tokenId, address indexed recipient)
func (_Csm *CsmFilterer) WatchERC721Recovered(opts *bind.WatchOpts, sink chan<- *CsmERC721Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csm.contract.WatchLogs(opts, "ERC721Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmERC721Recovered)
				if err := _Csm.contract.UnpackLog(event, "ERC721Recovered", log); err != nil {
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
func (_Csm *CsmFilterer) ParseERC721Recovered(log types.Log) (*CsmERC721Recovered, error) {
	event := new(CsmERC721Recovered)
	if err := _Csm.contract.UnpackLog(event, "ERC721Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmEtherRecoveredIterator is returned from FilterEtherRecovered and is used to iterate over the raw logs and unpacked data for EtherRecovered events raised by the Csm contract.
type CsmEtherRecoveredIterator struct {
	Event *CsmEtherRecovered // Event containing the contract specifics and raw log

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
func (it *CsmEtherRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmEtherRecovered)
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
		it.Event = new(CsmEtherRecovered)
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
func (it *CsmEtherRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmEtherRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmEtherRecovered represents a EtherRecovered event raised by the Csm contract.
type CsmEtherRecovered struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEtherRecovered is a free log retrieval operation binding the contract event 0x8e274e42262a7f013b700b35c2b4629ccce1702f8fe83f8dfb7eacbb26a4382c.
//
// Solidity: event EtherRecovered(address indexed recipient, uint256 amount)
func (_Csm *CsmFilterer) FilterEtherRecovered(opts *bind.FilterOpts, recipient []common.Address) (*CsmEtherRecoveredIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csm.contract.FilterLogs(opts, "EtherRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsmEtherRecoveredIterator{contract: _Csm.contract, event: "EtherRecovered", logs: logs, sub: sub}, nil
}

// WatchEtherRecovered is a free log subscription operation binding the contract event 0x8e274e42262a7f013b700b35c2b4629ccce1702f8fe83f8dfb7eacbb26a4382c.
//
// Solidity: event EtherRecovered(address indexed recipient, uint256 amount)
func (_Csm *CsmFilterer) WatchEtherRecovered(opts *bind.WatchOpts, sink chan<- *CsmEtherRecovered, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csm.contract.WatchLogs(opts, "EtherRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmEtherRecovered)
				if err := _Csm.contract.UnpackLog(event, "EtherRecovered", log); err != nil {
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
func (_Csm *CsmFilterer) ParseEtherRecovered(log types.Log) (*CsmEtherRecovered, error) {
	event := new(CsmEtherRecovered)
	if err := _Csm.contract.UnpackLog(event, "EtherRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmExitedSigningKeysCountChangedIterator is returned from FilterExitedSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for ExitedSigningKeysCountChanged events raised by the Csm contract.
type CsmExitedSigningKeysCountChangedIterator struct {
	Event *CsmExitedSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *CsmExitedSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmExitedSigningKeysCountChanged)
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
		it.Event = new(CsmExitedSigningKeysCountChanged)
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
func (it *CsmExitedSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmExitedSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmExitedSigningKeysCountChanged represents a ExitedSigningKeysCountChanged event raised by the Csm contract.
type CsmExitedSigningKeysCountChanged struct {
	NodeOperatorId  *big.Int
	ExitedKeysCount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterExitedSigningKeysCountChanged is a free log retrieval operation binding the contract event 0x0f67960648751434ae86bf350db61194f387fda387e7f568b0ccd0ae0c220166.
//
// Solidity: event ExitedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 exitedKeysCount)
func (_Csm *CsmFilterer) FilterExitedSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmExitedSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.FilterLogs(opts, "ExitedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmExitedSigningKeysCountChangedIterator{contract: _Csm.contract, event: "ExitedSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchExitedSigningKeysCountChanged is a free log subscription operation binding the contract event 0x0f67960648751434ae86bf350db61194f387fda387e7f568b0ccd0ae0c220166.
//
// Solidity: event ExitedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 exitedKeysCount)
func (_Csm *CsmFilterer) WatchExitedSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *CsmExitedSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.WatchLogs(opts, "ExitedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmExitedSigningKeysCountChanged)
				if err := _Csm.contract.UnpackLog(event, "ExitedSigningKeysCountChanged", log); err != nil {
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
func (_Csm *CsmFilterer) ParseExitedSigningKeysCountChanged(log types.Log) (*CsmExitedSigningKeysCountChanged, error) {
	event := new(CsmExitedSigningKeysCountChanged)
	if err := _Csm.contract.UnpackLog(event, "ExitedSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmInitialSlashingSubmittedIterator is returned from FilterInitialSlashingSubmitted and is used to iterate over the raw logs and unpacked data for InitialSlashingSubmitted events raised by the Csm contract.
type CsmInitialSlashingSubmittedIterator struct {
	Event *CsmInitialSlashingSubmitted // Event containing the contract specifics and raw log

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
func (it *CsmInitialSlashingSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmInitialSlashingSubmitted)
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
		it.Event = new(CsmInitialSlashingSubmitted)
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
func (it *CsmInitialSlashingSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmInitialSlashingSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmInitialSlashingSubmitted represents a InitialSlashingSubmitted event raised by the Csm contract.
type CsmInitialSlashingSubmitted struct {
	NodeOperatorId *big.Int
	KeyIndex       *big.Int
	Pubkey         []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterInitialSlashingSubmitted is a free log retrieval operation binding the contract event 0x0d541877c9d326d4c8ccfd72e6719f06dccb62a28292ae647e923441bcaad5c0.
//
// Solidity: event InitialSlashingSubmitted(uint256 indexed nodeOperatorId, uint256 keyIndex, bytes pubkey)
func (_Csm *CsmFilterer) FilterInitialSlashingSubmitted(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmInitialSlashingSubmittedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.FilterLogs(opts, "InitialSlashingSubmitted", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmInitialSlashingSubmittedIterator{contract: _Csm.contract, event: "InitialSlashingSubmitted", logs: logs, sub: sub}, nil
}

// WatchInitialSlashingSubmitted is a free log subscription operation binding the contract event 0x0d541877c9d326d4c8ccfd72e6719f06dccb62a28292ae647e923441bcaad5c0.
//
// Solidity: event InitialSlashingSubmitted(uint256 indexed nodeOperatorId, uint256 keyIndex, bytes pubkey)
func (_Csm *CsmFilterer) WatchInitialSlashingSubmitted(opts *bind.WatchOpts, sink chan<- *CsmInitialSlashingSubmitted, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.WatchLogs(opts, "InitialSlashingSubmitted", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmInitialSlashingSubmitted)
				if err := _Csm.contract.UnpackLog(event, "InitialSlashingSubmitted", log); err != nil {
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

// ParseInitialSlashingSubmitted is a log parse operation binding the contract event 0x0d541877c9d326d4c8ccfd72e6719f06dccb62a28292ae647e923441bcaad5c0.
//
// Solidity: event InitialSlashingSubmitted(uint256 indexed nodeOperatorId, uint256 keyIndex, bytes pubkey)
func (_Csm *CsmFilterer) ParseInitialSlashingSubmitted(log types.Log) (*CsmInitialSlashingSubmitted, error) {
	event := new(CsmInitialSlashingSubmitted)
	if err := _Csm.contract.UnpackLog(event, "InitialSlashingSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Csm contract.
type CsmInitializedIterator struct {
	Event *CsmInitialized // Event containing the contract specifics and raw log

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
func (it *CsmInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmInitialized)
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
		it.Event = new(CsmInitialized)
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
func (it *CsmInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmInitialized represents a Initialized event raised by the Csm contract.
type CsmInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Csm *CsmFilterer) FilterInitialized(opts *bind.FilterOpts) (*CsmInitializedIterator, error) {

	logs, sub, err := _Csm.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CsmInitializedIterator{contract: _Csm.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Csm *CsmFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CsmInitialized) (event.Subscription, error) {

	logs, sub, err := _Csm.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmInitialized)
				if err := _Csm.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Csm *CsmFilterer) ParseInitialized(log types.Log) (*CsmInitialized, error) {
	event := new(CsmInitialized)
	if err := _Csm.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmKeyRemovalChargeAppliedIterator is returned from FilterKeyRemovalChargeApplied and is used to iterate over the raw logs and unpacked data for KeyRemovalChargeApplied events raised by the Csm contract.
type CsmKeyRemovalChargeAppliedIterator struct {
	Event *CsmKeyRemovalChargeApplied // Event containing the contract specifics and raw log

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
func (it *CsmKeyRemovalChargeAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmKeyRemovalChargeApplied)
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
		it.Event = new(CsmKeyRemovalChargeApplied)
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
func (it *CsmKeyRemovalChargeAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmKeyRemovalChargeAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmKeyRemovalChargeApplied represents a KeyRemovalChargeApplied event raised by the Csm contract.
type CsmKeyRemovalChargeApplied struct {
	NodeOperatorId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterKeyRemovalChargeApplied is a free log retrieval operation binding the contract event 0x1cbb8dafbedbdf4f813a8ed1f50d871def63e1104f8729b677af57905eda90f6.
//
// Solidity: event KeyRemovalChargeApplied(uint256 indexed nodeOperatorId)
func (_Csm *CsmFilterer) FilterKeyRemovalChargeApplied(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmKeyRemovalChargeAppliedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.FilterLogs(opts, "KeyRemovalChargeApplied", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmKeyRemovalChargeAppliedIterator{contract: _Csm.contract, event: "KeyRemovalChargeApplied", logs: logs, sub: sub}, nil
}

// WatchKeyRemovalChargeApplied is a free log subscription operation binding the contract event 0x1cbb8dafbedbdf4f813a8ed1f50d871def63e1104f8729b677af57905eda90f6.
//
// Solidity: event KeyRemovalChargeApplied(uint256 indexed nodeOperatorId)
func (_Csm *CsmFilterer) WatchKeyRemovalChargeApplied(opts *bind.WatchOpts, sink chan<- *CsmKeyRemovalChargeApplied, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.WatchLogs(opts, "KeyRemovalChargeApplied", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmKeyRemovalChargeApplied)
				if err := _Csm.contract.UnpackLog(event, "KeyRemovalChargeApplied", log); err != nil {
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
func (_Csm *CsmFilterer) ParseKeyRemovalChargeApplied(log types.Log) (*CsmKeyRemovalChargeApplied, error) {
	event := new(CsmKeyRemovalChargeApplied)
	if err := _Csm.contract.UnpackLog(event, "KeyRemovalChargeApplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmKeyRemovalChargeSetIterator is returned from FilterKeyRemovalChargeSet and is used to iterate over the raw logs and unpacked data for KeyRemovalChargeSet events raised by the Csm contract.
type CsmKeyRemovalChargeSetIterator struct {
	Event *CsmKeyRemovalChargeSet // Event containing the contract specifics and raw log

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
func (it *CsmKeyRemovalChargeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmKeyRemovalChargeSet)
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
		it.Event = new(CsmKeyRemovalChargeSet)
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
func (it *CsmKeyRemovalChargeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmKeyRemovalChargeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmKeyRemovalChargeSet represents a KeyRemovalChargeSet event raised by the Csm contract.
type CsmKeyRemovalChargeSet struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterKeyRemovalChargeSet is a free log retrieval operation binding the contract event 0x699ec9c671aad1f3dcc15e571375584a1d6fb7176afd545d14467fd31477e98e.
//
// Solidity: event KeyRemovalChargeSet(uint256 amount)
func (_Csm *CsmFilterer) FilterKeyRemovalChargeSet(opts *bind.FilterOpts) (*CsmKeyRemovalChargeSetIterator, error) {

	logs, sub, err := _Csm.contract.FilterLogs(opts, "KeyRemovalChargeSet")
	if err != nil {
		return nil, err
	}
	return &CsmKeyRemovalChargeSetIterator{contract: _Csm.contract, event: "KeyRemovalChargeSet", logs: logs, sub: sub}, nil
}

// WatchKeyRemovalChargeSet is a free log subscription operation binding the contract event 0x699ec9c671aad1f3dcc15e571375584a1d6fb7176afd545d14467fd31477e98e.
//
// Solidity: event KeyRemovalChargeSet(uint256 amount)
func (_Csm *CsmFilterer) WatchKeyRemovalChargeSet(opts *bind.WatchOpts, sink chan<- *CsmKeyRemovalChargeSet) (event.Subscription, error) {

	logs, sub, err := _Csm.contract.WatchLogs(opts, "KeyRemovalChargeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmKeyRemovalChargeSet)
				if err := _Csm.contract.UnpackLog(event, "KeyRemovalChargeSet", log); err != nil {
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

// ParseKeyRemovalChargeSet is a log parse operation binding the contract event 0x699ec9c671aad1f3dcc15e571375584a1d6fb7176afd545d14467fd31477e98e.
//
// Solidity: event KeyRemovalChargeSet(uint256 amount)
func (_Csm *CsmFilterer) ParseKeyRemovalChargeSet(log types.Log) (*CsmKeyRemovalChargeSet, error) {
	event := new(CsmKeyRemovalChargeSet)
	if err := _Csm.contract.UnpackLog(event, "KeyRemovalChargeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmNodeOperatorAddedIterator is returned from FilterNodeOperatorAdded and is used to iterate over the raw logs and unpacked data for NodeOperatorAdded events raised by the Csm contract.
type CsmNodeOperatorAddedIterator struct {
	Event *CsmNodeOperatorAdded // Event containing the contract specifics and raw log

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
func (it *CsmNodeOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmNodeOperatorAdded)
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
		it.Event = new(CsmNodeOperatorAdded)
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
func (it *CsmNodeOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmNodeOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmNodeOperatorAdded represents a NodeOperatorAdded event raised by the Csm contract.
type CsmNodeOperatorAdded struct {
	NodeOperatorId *big.Int
	ManagerAddress common.Address
	RewardAddress  common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorAdded is a free log retrieval operation binding the contract event 0xf35982c84fdc94f58d48e901c54c615804cf7d7939b9b8f76ce4d459354e6363.
//
// Solidity: event NodeOperatorAdded(uint256 indexed nodeOperatorId, address indexed managerAddress, address indexed rewardAddress)
func (_Csm *CsmFilterer) FilterNodeOperatorAdded(opts *bind.FilterOpts, nodeOperatorId []*big.Int, managerAddress []common.Address, rewardAddress []common.Address) (*CsmNodeOperatorAddedIterator, error) {

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

	logs, sub, err := _Csm.contract.FilterLogs(opts, "NodeOperatorAdded", nodeOperatorIdRule, managerAddressRule, rewardAddressRule)
	if err != nil {
		return nil, err
	}
	return &CsmNodeOperatorAddedIterator{contract: _Csm.contract, event: "NodeOperatorAdded", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorAdded is a free log subscription operation binding the contract event 0xf35982c84fdc94f58d48e901c54c615804cf7d7939b9b8f76ce4d459354e6363.
//
// Solidity: event NodeOperatorAdded(uint256 indexed nodeOperatorId, address indexed managerAddress, address indexed rewardAddress)
func (_Csm *CsmFilterer) WatchNodeOperatorAdded(opts *bind.WatchOpts, sink chan<- *CsmNodeOperatorAdded, nodeOperatorId []*big.Int, managerAddress []common.Address, rewardAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Csm.contract.WatchLogs(opts, "NodeOperatorAdded", nodeOperatorIdRule, managerAddressRule, rewardAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmNodeOperatorAdded)
				if err := _Csm.contract.UnpackLog(event, "NodeOperatorAdded", log); err != nil {
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

// ParseNodeOperatorAdded is a log parse operation binding the contract event 0xf35982c84fdc94f58d48e901c54c615804cf7d7939b9b8f76ce4d459354e6363.
//
// Solidity: event NodeOperatorAdded(uint256 indexed nodeOperatorId, address indexed managerAddress, address indexed rewardAddress)
func (_Csm *CsmFilterer) ParseNodeOperatorAdded(log types.Log) (*CsmNodeOperatorAdded, error) {
	event := new(CsmNodeOperatorAdded)
	if err := _Csm.contract.UnpackLog(event, "NodeOperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmNodeOperatorManagerAddressChangeProposedIterator is returned from FilterNodeOperatorManagerAddressChangeProposed and is used to iterate over the raw logs and unpacked data for NodeOperatorManagerAddressChangeProposed events raised by the Csm contract.
type CsmNodeOperatorManagerAddressChangeProposedIterator struct {
	Event *CsmNodeOperatorManagerAddressChangeProposed // Event containing the contract specifics and raw log

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
func (it *CsmNodeOperatorManagerAddressChangeProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmNodeOperatorManagerAddressChangeProposed)
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
		it.Event = new(CsmNodeOperatorManagerAddressChangeProposed)
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
func (it *CsmNodeOperatorManagerAddressChangeProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmNodeOperatorManagerAddressChangeProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmNodeOperatorManagerAddressChangeProposed represents a NodeOperatorManagerAddressChangeProposed event raised by the Csm contract.
type CsmNodeOperatorManagerAddressChangeProposed struct {
	NodeOperatorId     *big.Int
	OldProposedAddress common.Address
	NewProposedAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorManagerAddressChangeProposed is a free log retrieval operation binding the contract event 0x4048f15a706950765ca59f99d0fa6fe8edaaa3f3e3d0337417082e2131df82fb.
//
// Solidity: event NodeOperatorManagerAddressChangeProposed(uint256 indexed nodeOperatorId, address indexed oldProposedAddress, address indexed newProposedAddress)
func (_Csm *CsmFilterer) FilterNodeOperatorManagerAddressChangeProposed(opts *bind.FilterOpts, nodeOperatorId []*big.Int, oldProposedAddress []common.Address, newProposedAddress []common.Address) (*CsmNodeOperatorManagerAddressChangeProposedIterator, error) {

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

	logs, sub, err := _Csm.contract.FilterLogs(opts, "NodeOperatorManagerAddressChangeProposed", nodeOperatorIdRule, oldProposedAddressRule, newProposedAddressRule)
	if err != nil {
		return nil, err
	}
	return &CsmNodeOperatorManagerAddressChangeProposedIterator{contract: _Csm.contract, event: "NodeOperatorManagerAddressChangeProposed", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorManagerAddressChangeProposed is a free log subscription operation binding the contract event 0x4048f15a706950765ca59f99d0fa6fe8edaaa3f3e3d0337417082e2131df82fb.
//
// Solidity: event NodeOperatorManagerAddressChangeProposed(uint256 indexed nodeOperatorId, address indexed oldProposedAddress, address indexed newProposedAddress)
func (_Csm *CsmFilterer) WatchNodeOperatorManagerAddressChangeProposed(opts *bind.WatchOpts, sink chan<- *CsmNodeOperatorManagerAddressChangeProposed, nodeOperatorId []*big.Int, oldProposedAddress []common.Address, newProposedAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Csm.contract.WatchLogs(opts, "NodeOperatorManagerAddressChangeProposed", nodeOperatorIdRule, oldProposedAddressRule, newProposedAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmNodeOperatorManagerAddressChangeProposed)
				if err := _Csm.contract.UnpackLog(event, "NodeOperatorManagerAddressChangeProposed", log); err != nil {
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
func (_Csm *CsmFilterer) ParseNodeOperatorManagerAddressChangeProposed(log types.Log) (*CsmNodeOperatorManagerAddressChangeProposed, error) {
	event := new(CsmNodeOperatorManagerAddressChangeProposed)
	if err := _Csm.contract.UnpackLog(event, "NodeOperatorManagerAddressChangeProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmNodeOperatorManagerAddressChangedIterator is returned from FilterNodeOperatorManagerAddressChanged and is used to iterate over the raw logs and unpacked data for NodeOperatorManagerAddressChanged events raised by the Csm contract.
type CsmNodeOperatorManagerAddressChangedIterator struct {
	Event *CsmNodeOperatorManagerAddressChanged // Event containing the contract specifics and raw log

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
func (it *CsmNodeOperatorManagerAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmNodeOperatorManagerAddressChanged)
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
		it.Event = new(CsmNodeOperatorManagerAddressChanged)
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
func (it *CsmNodeOperatorManagerAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmNodeOperatorManagerAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmNodeOperatorManagerAddressChanged represents a NodeOperatorManagerAddressChanged event raised by the Csm contract.
type CsmNodeOperatorManagerAddressChanged struct {
	NodeOperatorId *big.Int
	OldAddress     common.Address
	NewAddress     common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorManagerAddressChanged is a free log retrieval operation binding the contract event 0x862021f23449d6e8516867bd839be15a3d8698a7561c5c2c35069074b7e91e61.
//
// Solidity: event NodeOperatorManagerAddressChanged(uint256 indexed nodeOperatorId, address indexed oldAddress, address indexed newAddress)
func (_Csm *CsmFilterer) FilterNodeOperatorManagerAddressChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int, oldAddress []common.Address, newAddress []common.Address) (*CsmNodeOperatorManagerAddressChangedIterator, error) {

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

	logs, sub, err := _Csm.contract.FilterLogs(opts, "NodeOperatorManagerAddressChanged", nodeOperatorIdRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &CsmNodeOperatorManagerAddressChangedIterator{contract: _Csm.contract, event: "NodeOperatorManagerAddressChanged", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorManagerAddressChanged is a free log subscription operation binding the contract event 0x862021f23449d6e8516867bd839be15a3d8698a7561c5c2c35069074b7e91e61.
//
// Solidity: event NodeOperatorManagerAddressChanged(uint256 indexed nodeOperatorId, address indexed oldAddress, address indexed newAddress)
func (_Csm *CsmFilterer) WatchNodeOperatorManagerAddressChanged(opts *bind.WatchOpts, sink chan<- *CsmNodeOperatorManagerAddressChanged, nodeOperatorId []*big.Int, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Csm.contract.WatchLogs(opts, "NodeOperatorManagerAddressChanged", nodeOperatorIdRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmNodeOperatorManagerAddressChanged)
				if err := _Csm.contract.UnpackLog(event, "NodeOperatorManagerAddressChanged", log); err != nil {
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
func (_Csm *CsmFilterer) ParseNodeOperatorManagerAddressChanged(log types.Log) (*CsmNodeOperatorManagerAddressChanged, error) {
	event := new(CsmNodeOperatorManagerAddressChanged)
	if err := _Csm.contract.UnpackLog(event, "NodeOperatorManagerAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmNodeOperatorRewardAddressChangeProposedIterator is returned from FilterNodeOperatorRewardAddressChangeProposed and is used to iterate over the raw logs and unpacked data for NodeOperatorRewardAddressChangeProposed events raised by the Csm contract.
type CsmNodeOperatorRewardAddressChangeProposedIterator struct {
	Event *CsmNodeOperatorRewardAddressChangeProposed // Event containing the contract specifics and raw log

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
func (it *CsmNodeOperatorRewardAddressChangeProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmNodeOperatorRewardAddressChangeProposed)
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
		it.Event = new(CsmNodeOperatorRewardAddressChangeProposed)
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
func (it *CsmNodeOperatorRewardAddressChangeProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmNodeOperatorRewardAddressChangeProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmNodeOperatorRewardAddressChangeProposed represents a NodeOperatorRewardAddressChangeProposed event raised by the Csm contract.
type CsmNodeOperatorRewardAddressChangeProposed struct {
	NodeOperatorId     *big.Int
	OldProposedAddress common.Address
	NewProposedAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorRewardAddressChangeProposed is a free log retrieval operation binding the contract event 0xb5878cdb1d66f971efe3b138a71c64bc5bc519314db2533e0e4cde954409ea5a.
//
// Solidity: event NodeOperatorRewardAddressChangeProposed(uint256 indexed nodeOperatorId, address indexed oldProposedAddress, address indexed newProposedAddress)
func (_Csm *CsmFilterer) FilterNodeOperatorRewardAddressChangeProposed(opts *bind.FilterOpts, nodeOperatorId []*big.Int, oldProposedAddress []common.Address, newProposedAddress []common.Address) (*CsmNodeOperatorRewardAddressChangeProposedIterator, error) {

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

	logs, sub, err := _Csm.contract.FilterLogs(opts, "NodeOperatorRewardAddressChangeProposed", nodeOperatorIdRule, oldProposedAddressRule, newProposedAddressRule)
	if err != nil {
		return nil, err
	}
	return &CsmNodeOperatorRewardAddressChangeProposedIterator{contract: _Csm.contract, event: "NodeOperatorRewardAddressChangeProposed", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorRewardAddressChangeProposed is a free log subscription operation binding the contract event 0xb5878cdb1d66f971efe3b138a71c64bc5bc519314db2533e0e4cde954409ea5a.
//
// Solidity: event NodeOperatorRewardAddressChangeProposed(uint256 indexed nodeOperatorId, address indexed oldProposedAddress, address indexed newProposedAddress)
func (_Csm *CsmFilterer) WatchNodeOperatorRewardAddressChangeProposed(opts *bind.WatchOpts, sink chan<- *CsmNodeOperatorRewardAddressChangeProposed, nodeOperatorId []*big.Int, oldProposedAddress []common.Address, newProposedAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Csm.contract.WatchLogs(opts, "NodeOperatorRewardAddressChangeProposed", nodeOperatorIdRule, oldProposedAddressRule, newProposedAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmNodeOperatorRewardAddressChangeProposed)
				if err := _Csm.contract.UnpackLog(event, "NodeOperatorRewardAddressChangeProposed", log); err != nil {
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
func (_Csm *CsmFilterer) ParseNodeOperatorRewardAddressChangeProposed(log types.Log) (*CsmNodeOperatorRewardAddressChangeProposed, error) {
	event := new(CsmNodeOperatorRewardAddressChangeProposed)
	if err := _Csm.contract.UnpackLog(event, "NodeOperatorRewardAddressChangeProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmNodeOperatorRewardAddressChangedIterator is returned from FilterNodeOperatorRewardAddressChanged and is used to iterate over the raw logs and unpacked data for NodeOperatorRewardAddressChanged events raised by the Csm contract.
type CsmNodeOperatorRewardAddressChangedIterator struct {
	Event *CsmNodeOperatorRewardAddressChanged // Event containing the contract specifics and raw log

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
func (it *CsmNodeOperatorRewardAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmNodeOperatorRewardAddressChanged)
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
		it.Event = new(CsmNodeOperatorRewardAddressChanged)
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
func (it *CsmNodeOperatorRewardAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmNodeOperatorRewardAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmNodeOperatorRewardAddressChanged represents a NodeOperatorRewardAddressChanged event raised by the Csm contract.
type CsmNodeOperatorRewardAddressChanged struct {
	NodeOperatorId *big.Int
	OldAddress     common.Address
	NewAddress     common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorRewardAddressChanged is a free log retrieval operation binding the contract event 0x069ac7cd8230db015b7250c8e5425149cf1a3e912d9569f497165e55b3b6b7b2.
//
// Solidity: event NodeOperatorRewardAddressChanged(uint256 indexed nodeOperatorId, address indexed oldAddress, address indexed newAddress)
func (_Csm *CsmFilterer) FilterNodeOperatorRewardAddressChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int, oldAddress []common.Address, newAddress []common.Address) (*CsmNodeOperatorRewardAddressChangedIterator, error) {

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

	logs, sub, err := _Csm.contract.FilterLogs(opts, "NodeOperatorRewardAddressChanged", nodeOperatorIdRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &CsmNodeOperatorRewardAddressChangedIterator{contract: _Csm.contract, event: "NodeOperatorRewardAddressChanged", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorRewardAddressChanged is a free log subscription operation binding the contract event 0x069ac7cd8230db015b7250c8e5425149cf1a3e912d9569f497165e55b3b6b7b2.
//
// Solidity: event NodeOperatorRewardAddressChanged(uint256 indexed nodeOperatorId, address indexed oldAddress, address indexed newAddress)
func (_Csm *CsmFilterer) WatchNodeOperatorRewardAddressChanged(opts *bind.WatchOpts, sink chan<- *CsmNodeOperatorRewardAddressChanged, nodeOperatorId []*big.Int, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Csm.contract.WatchLogs(opts, "NodeOperatorRewardAddressChanged", nodeOperatorIdRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmNodeOperatorRewardAddressChanged)
				if err := _Csm.contract.UnpackLog(event, "NodeOperatorRewardAddressChanged", log); err != nil {
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
func (_Csm *CsmFilterer) ParseNodeOperatorRewardAddressChanged(log types.Log) (*CsmNodeOperatorRewardAddressChanged, error) {
	event := new(CsmNodeOperatorRewardAddressChanged)
	if err := _Csm.contract.UnpackLog(event, "NodeOperatorRewardAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmNonceChangedIterator is returned from FilterNonceChanged and is used to iterate over the raw logs and unpacked data for NonceChanged events raised by the Csm contract.
type CsmNonceChangedIterator struct {
	Event *CsmNonceChanged // Event containing the contract specifics and raw log

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
func (it *CsmNonceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmNonceChanged)
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
		it.Event = new(CsmNonceChanged)
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
func (it *CsmNonceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmNonceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmNonceChanged represents a NonceChanged event raised by the Csm contract.
type CsmNonceChanged struct {
	Nonce *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNonceChanged is a free log retrieval operation binding the contract event 0x7220970e1f1f12864ecccd8942690a837c7a8dd45d158cb891eb45a8a69134aa.
//
// Solidity: event NonceChanged(uint256 nonce)
func (_Csm *CsmFilterer) FilterNonceChanged(opts *bind.FilterOpts) (*CsmNonceChangedIterator, error) {

	logs, sub, err := _Csm.contract.FilterLogs(opts, "NonceChanged")
	if err != nil {
		return nil, err
	}
	return &CsmNonceChangedIterator{contract: _Csm.contract, event: "NonceChanged", logs: logs, sub: sub}, nil
}

// WatchNonceChanged is a free log subscription operation binding the contract event 0x7220970e1f1f12864ecccd8942690a837c7a8dd45d158cb891eb45a8a69134aa.
//
// Solidity: event NonceChanged(uint256 nonce)
func (_Csm *CsmFilterer) WatchNonceChanged(opts *bind.WatchOpts, sink chan<- *CsmNonceChanged) (event.Subscription, error) {

	logs, sub, err := _Csm.contract.WatchLogs(opts, "NonceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmNonceChanged)
				if err := _Csm.contract.UnpackLog(event, "NonceChanged", log); err != nil {
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
func (_Csm *CsmFilterer) ParseNonceChanged(log types.Log) (*CsmNonceChanged, error) {
	event := new(CsmNonceChanged)
	if err := _Csm.contract.UnpackLog(event, "NonceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Csm contract.
type CsmPausedIterator struct {
	Event *CsmPaused // Event containing the contract specifics and raw log

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
func (it *CsmPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmPaused)
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
		it.Event = new(CsmPaused)
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
func (it *CsmPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmPaused represents a Paused event raised by the Csm contract.
type CsmPaused struct {
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_Csm *CsmFilterer) FilterPaused(opts *bind.FilterOpts) (*CsmPausedIterator, error) {

	logs, sub, err := _Csm.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CsmPausedIterator{contract: _Csm.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_Csm *CsmFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CsmPaused) (event.Subscription, error) {

	logs, sub, err := _Csm.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmPaused)
				if err := _Csm.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Csm *CsmFilterer) ParsePaused(log types.Log) (*CsmPaused, error) {
	event := new(CsmPaused)
	if err := _Csm.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmPublicReleaseIterator is returned from FilterPublicRelease and is used to iterate over the raw logs and unpacked data for PublicRelease events raised by the Csm contract.
type CsmPublicReleaseIterator struct {
	Event *CsmPublicRelease // Event containing the contract specifics and raw log

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
func (it *CsmPublicReleaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmPublicRelease)
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
		it.Event = new(CsmPublicRelease)
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
func (it *CsmPublicReleaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmPublicReleaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmPublicRelease represents a PublicRelease event raised by the Csm contract.
type CsmPublicRelease struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPublicRelease is a free log retrieval operation binding the contract event 0xe5eb57aa4d841adeece4ac87bd294965df4a894f0aa24db4a4a55a39ab101d6e.
//
// Solidity: event PublicRelease()
func (_Csm *CsmFilterer) FilterPublicRelease(opts *bind.FilterOpts) (*CsmPublicReleaseIterator, error) {

	logs, sub, err := _Csm.contract.FilterLogs(opts, "PublicRelease")
	if err != nil {
		return nil, err
	}
	return &CsmPublicReleaseIterator{contract: _Csm.contract, event: "PublicRelease", logs: logs, sub: sub}, nil
}

// WatchPublicRelease is a free log subscription operation binding the contract event 0xe5eb57aa4d841adeece4ac87bd294965df4a894f0aa24db4a4a55a39ab101d6e.
//
// Solidity: event PublicRelease()
func (_Csm *CsmFilterer) WatchPublicRelease(opts *bind.WatchOpts, sink chan<- *CsmPublicRelease) (event.Subscription, error) {

	logs, sub, err := _Csm.contract.WatchLogs(opts, "PublicRelease")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmPublicRelease)
				if err := _Csm.contract.UnpackLog(event, "PublicRelease", log); err != nil {
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

// ParsePublicRelease is a log parse operation binding the contract event 0xe5eb57aa4d841adeece4ac87bd294965df4a894f0aa24db4a4a55a39ab101d6e.
//
// Solidity: event PublicRelease()
func (_Csm *CsmFilterer) ParsePublicRelease(log types.Log) (*CsmPublicRelease, error) {
	event := new(CsmPublicRelease)
	if err := _Csm.contract.UnpackLog(event, "PublicRelease", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmReferrerSetIterator is returned from FilterReferrerSet and is used to iterate over the raw logs and unpacked data for ReferrerSet events raised by the Csm contract.
type CsmReferrerSetIterator struct {
	Event *CsmReferrerSet // Event containing the contract specifics and raw log

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
func (it *CsmReferrerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmReferrerSet)
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
		it.Event = new(CsmReferrerSet)
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
func (it *CsmReferrerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmReferrerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmReferrerSet represents a ReferrerSet event raised by the Csm contract.
type CsmReferrerSet struct {
	NodeOperatorId *big.Int
	Referrer       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterReferrerSet is a free log retrieval operation binding the contract event 0x67334334c388385e5f244703f8a8b28b7f4ffe52909130aca69bc62a8e27f09a.
//
// Solidity: event ReferrerSet(uint256 indexed nodeOperatorId, address indexed referrer)
func (_Csm *CsmFilterer) FilterReferrerSet(opts *bind.FilterOpts, nodeOperatorId []*big.Int, referrer []common.Address) (*CsmReferrerSetIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var referrerRule []interface{}
	for _, referrerItem := range referrer {
		referrerRule = append(referrerRule, referrerItem)
	}

	logs, sub, err := _Csm.contract.FilterLogs(opts, "ReferrerSet", nodeOperatorIdRule, referrerRule)
	if err != nil {
		return nil, err
	}
	return &CsmReferrerSetIterator{contract: _Csm.contract, event: "ReferrerSet", logs: logs, sub: sub}, nil
}

// WatchReferrerSet is a free log subscription operation binding the contract event 0x67334334c388385e5f244703f8a8b28b7f4ffe52909130aca69bc62a8e27f09a.
//
// Solidity: event ReferrerSet(uint256 indexed nodeOperatorId, address indexed referrer)
func (_Csm *CsmFilterer) WatchReferrerSet(opts *bind.WatchOpts, sink chan<- *CsmReferrerSet, nodeOperatorId []*big.Int, referrer []common.Address) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var referrerRule []interface{}
	for _, referrerItem := range referrer {
		referrerRule = append(referrerRule, referrerItem)
	}

	logs, sub, err := _Csm.contract.WatchLogs(opts, "ReferrerSet", nodeOperatorIdRule, referrerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmReferrerSet)
				if err := _Csm.contract.UnpackLog(event, "ReferrerSet", log); err != nil {
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
func (_Csm *CsmFilterer) ParseReferrerSet(log types.Log) (*CsmReferrerSet, error) {
	event := new(CsmReferrerSet)
	if err := _Csm.contract.UnpackLog(event, "ReferrerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmResumedIterator is returned from FilterResumed and is used to iterate over the raw logs and unpacked data for Resumed events raised by the Csm contract.
type CsmResumedIterator struct {
	Event *CsmResumed // Event containing the contract specifics and raw log

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
func (it *CsmResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmResumed)
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
		it.Event = new(CsmResumed)
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
func (it *CsmResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmResumed represents a Resumed event raised by the Csm contract.
type CsmResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResumed is a free log retrieval operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Csm *CsmFilterer) FilterResumed(opts *bind.FilterOpts) (*CsmResumedIterator, error) {

	logs, sub, err := _Csm.contract.FilterLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return &CsmResumedIterator{contract: _Csm.contract, event: "Resumed", logs: logs, sub: sub}, nil
}

// WatchResumed is a free log subscription operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Csm *CsmFilterer) WatchResumed(opts *bind.WatchOpts, sink chan<- *CsmResumed) (event.Subscription, error) {

	logs, sub, err := _Csm.contract.WatchLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmResumed)
				if err := _Csm.contract.UnpackLog(event, "Resumed", log); err != nil {
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
func (_Csm *CsmFilterer) ParseResumed(log types.Log) (*CsmResumed, error) {
	event := new(CsmResumed)
	if err := _Csm.contract.UnpackLog(event, "Resumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Csm contract.
type CsmRoleAdminChangedIterator struct {
	Event *CsmRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CsmRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmRoleAdminChanged)
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
		it.Event = new(CsmRoleAdminChanged)
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
func (it *CsmRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmRoleAdminChanged represents a RoleAdminChanged event raised by the Csm contract.
type CsmRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Csm *CsmFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CsmRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Csm.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CsmRoleAdminChangedIterator{contract: _Csm.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Csm *CsmFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CsmRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Csm.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmRoleAdminChanged)
				if err := _Csm.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Csm *CsmFilterer) ParseRoleAdminChanged(log types.Log) (*CsmRoleAdminChanged, error) {
	event := new(CsmRoleAdminChanged)
	if err := _Csm.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Csm contract.
type CsmRoleGrantedIterator struct {
	Event *CsmRoleGranted // Event containing the contract specifics and raw log

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
func (it *CsmRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmRoleGranted)
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
		it.Event = new(CsmRoleGranted)
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
func (it *CsmRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmRoleGranted represents a RoleGranted event raised by the Csm contract.
type CsmRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Csm *CsmFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CsmRoleGrantedIterator, error) {

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

	logs, sub, err := _Csm.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CsmRoleGrantedIterator{contract: _Csm.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Csm *CsmFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CsmRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Csm.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmRoleGranted)
				if err := _Csm.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Csm *CsmFilterer) ParseRoleGranted(log types.Log) (*CsmRoleGranted, error) {
	event := new(CsmRoleGranted)
	if err := _Csm.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Csm contract.
type CsmRoleRevokedIterator struct {
	Event *CsmRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CsmRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmRoleRevoked)
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
		it.Event = new(CsmRoleRevoked)
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
func (it *CsmRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmRoleRevoked represents a RoleRevoked event raised by the Csm contract.
type CsmRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Csm *CsmFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CsmRoleRevokedIterator, error) {

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

	logs, sub, err := _Csm.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CsmRoleRevokedIterator{contract: _Csm.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Csm *CsmFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CsmRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Csm.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmRoleRevoked)
				if err := _Csm.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Csm *CsmFilterer) ParseRoleRevoked(log types.Log) (*CsmRoleRevoked, error) {
	event := new(CsmRoleRevoked)
	if err := _Csm.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmSigningKeyAddedIterator is returned from FilterSigningKeyAdded and is used to iterate over the raw logs and unpacked data for SigningKeyAdded events raised by the Csm contract.
type CsmSigningKeyAddedIterator struct {
	Event *CsmSigningKeyAdded // Event containing the contract specifics and raw log

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
func (it *CsmSigningKeyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmSigningKeyAdded)
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
		it.Event = new(CsmSigningKeyAdded)
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
func (it *CsmSigningKeyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmSigningKeyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmSigningKeyAdded represents a SigningKeyAdded event raised by the Csm contract.
type CsmSigningKeyAdded struct {
	NodeOperatorId *big.Int
	Pubkey         []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSigningKeyAdded is a free log retrieval operation binding the contract event 0xc77a17d6b857abe6d6e6c37301621bc72c4dd52fa8830fb54dfa715c04911a89.
//
// Solidity: event SigningKeyAdded(uint256 indexed nodeOperatorId, bytes pubkey)
func (_Csm *CsmFilterer) FilterSigningKeyAdded(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmSigningKeyAddedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.FilterLogs(opts, "SigningKeyAdded", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmSigningKeyAddedIterator{contract: _Csm.contract, event: "SigningKeyAdded", logs: logs, sub: sub}, nil
}

// WatchSigningKeyAdded is a free log subscription operation binding the contract event 0xc77a17d6b857abe6d6e6c37301621bc72c4dd52fa8830fb54dfa715c04911a89.
//
// Solidity: event SigningKeyAdded(uint256 indexed nodeOperatorId, bytes pubkey)
func (_Csm *CsmFilterer) WatchSigningKeyAdded(opts *bind.WatchOpts, sink chan<- *CsmSigningKeyAdded, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.WatchLogs(opts, "SigningKeyAdded", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmSigningKeyAdded)
				if err := _Csm.contract.UnpackLog(event, "SigningKeyAdded", log); err != nil {
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
func (_Csm *CsmFilterer) ParseSigningKeyAdded(log types.Log) (*CsmSigningKeyAdded, error) {
	event := new(CsmSigningKeyAdded)
	if err := _Csm.contract.UnpackLog(event, "SigningKeyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmSigningKeyRemovedIterator is returned from FilterSigningKeyRemoved and is used to iterate over the raw logs and unpacked data for SigningKeyRemoved events raised by the Csm contract.
type CsmSigningKeyRemovedIterator struct {
	Event *CsmSigningKeyRemoved // Event containing the contract specifics and raw log

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
func (it *CsmSigningKeyRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmSigningKeyRemoved)
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
		it.Event = new(CsmSigningKeyRemoved)
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
func (it *CsmSigningKeyRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmSigningKeyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmSigningKeyRemoved represents a SigningKeyRemoved event raised by the Csm contract.
type CsmSigningKeyRemoved struct {
	NodeOperatorId *big.Int
	Pubkey         []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSigningKeyRemoved is a free log retrieval operation binding the contract event 0xea4b75aaf57196f73d338cadf79ecd0a437902e2dd0d2c4c2cf3ea71b8ab27b9.
//
// Solidity: event SigningKeyRemoved(uint256 indexed nodeOperatorId, bytes pubkey)
func (_Csm *CsmFilterer) FilterSigningKeyRemoved(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmSigningKeyRemovedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.FilterLogs(opts, "SigningKeyRemoved", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmSigningKeyRemovedIterator{contract: _Csm.contract, event: "SigningKeyRemoved", logs: logs, sub: sub}, nil
}

// WatchSigningKeyRemoved is a free log subscription operation binding the contract event 0xea4b75aaf57196f73d338cadf79ecd0a437902e2dd0d2c4c2cf3ea71b8ab27b9.
//
// Solidity: event SigningKeyRemoved(uint256 indexed nodeOperatorId, bytes pubkey)
func (_Csm *CsmFilterer) WatchSigningKeyRemoved(opts *bind.WatchOpts, sink chan<- *CsmSigningKeyRemoved, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.WatchLogs(opts, "SigningKeyRemoved", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmSigningKeyRemoved)
				if err := _Csm.contract.UnpackLog(event, "SigningKeyRemoved", log); err != nil {
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
func (_Csm *CsmFilterer) ParseSigningKeyRemoved(log types.Log) (*CsmSigningKeyRemoved, error) {
	event := new(CsmSigningKeyRemoved)
	if err := _Csm.contract.UnpackLog(event, "SigningKeyRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmStETHSharesRecoveredIterator is returned from FilterStETHSharesRecovered and is used to iterate over the raw logs and unpacked data for StETHSharesRecovered events raised by the Csm contract.
type CsmStETHSharesRecoveredIterator struct {
	Event *CsmStETHSharesRecovered // Event containing the contract specifics and raw log

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
func (it *CsmStETHSharesRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmStETHSharesRecovered)
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
		it.Event = new(CsmStETHSharesRecovered)
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
func (it *CsmStETHSharesRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmStETHSharesRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmStETHSharesRecovered represents a StETHSharesRecovered event raised by the Csm contract.
type CsmStETHSharesRecovered struct {
	Recipient common.Address
	Shares    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStETHSharesRecovered is a free log retrieval operation binding the contract event 0x426e7e0100db57255d4af4a46cd49552ef74f5f002bbdc8d4ebb6371c0070a02.
//
// Solidity: event StETHSharesRecovered(address indexed recipient, uint256 shares)
func (_Csm *CsmFilterer) FilterStETHSharesRecovered(opts *bind.FilterOpts, recipient []common.Address) (*CsmStETHSharesRecoveredIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csm.contract.FilterLogs(opts, "StETHSharesRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsmStETHSharesRecoveredIterator{contract: _Csm.contract, event: "StETHSharesRecovered", logs: logs, sub: sub}, nil
}

// WatchStETHSharesRecovered is a free log subscription operation binding the contract event 0x426e7e0100db57255d4af4a46cd49552ef74f5f002bbdc8d4ebb6371c0070a02.
//
// Solidity: event StETHSharesRecovered(address indexed recipient, uint256 shares)
func (_Csm *CsmFilterer) WatchStETHSharesRecovered(opts *bind.WatchOpts, sink chan<- *CsmStETHSharesRecovered, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csm.contract.WatchLogs(opts, "StETHSharesRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmStETHSharesRecovered)
				if err := _Csm.contract.UnpackLog(event, "StETHSharesRecovered", log); err != nil {
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
func (_Csm *CsmFilterer) ParseStETHSharesRecovered(log types.Log) (*CsmStETHSharesRecovered, error) {
	event := new(CsmStETHSharesRecovered)
	if err := _Csm.contract.UnpackLog(event, "StETHSharesRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmStuckSigningKeysCountChangedIterator is returned from FilterStuckSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for StuckSigningKeysCountChanged events raised by the Csm contract.
type CsmStuckSigningKeysCountChangedIterator struct {
	Event *CsmStuckSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *CsmStuckSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmStuckSigningKeysCountChanged)
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
		it.Event = new(CsmStuckSigningKeysCountChanged)
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
func (it *CsmStuckSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmStuckSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmStuckSigningKeysCountChanged represents a StuckSigningKeysCountChanged event raised by the Csm contract.
type CsmStuckSigningKeysCountChanged struct {
	NodeOperatorId *big.Int
	StuckKeysCount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStuckSigningKeysCountChanged is a free log retrieval operation binding the contract event 0xb4f5879eca27b32881cec7907d1310378e9b4c79927062fb7d4a321434b5b04a.
//
// Solidity: event StuckSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 stuckKeysCount)
func (_Csm *CsmFilterer) FilterStuckSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmStuckSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.FilterLogs(opts, "StuckSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmStuckSigningKeysCountChangedIterator{contract: _Csm.contract, event: "StuckSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchStuckSigningKeysCountChanged is a free log subscription operation binding the contract event 0xb4f5879eca27b32881cec7907d1310378e9b4c79927062fb7d4a321434b5b04a.
//
// Solidity: event StuckSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 stuckKeysCount)
func (_Csm *CsmFilterer) WatchStuckSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *CsmStuckSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.WatchLogs(opts, "StuckSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmStuckSigningKeysCountChanged)
				if err := _Csm.contract.UnpackLog(event, "StuckSigningKeysCountChanged", log); err != nil {
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

// ParseStuckSigningKeysCountChanged is a log parse operation binding the contract event 0xb4f5879eca27b32881cec7907d1310378e9b4c79927062fb7d4a321434b5b04a.
//
// Solidity: event StuckSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 stuckKeysCount)
func (_Csm *CsmFilterer) ParseStuckSigningKeysCountChanged(log types.Log) (*CsmStuckSigningKeysCountChanged, error) {
	event := new(CsmStuckSigningKeysCountChanged)
	if err := _Csm.contract.UnpackLog(event, "StuckSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmTargetValidatorsCountChangedIterator is returned from FilterTargetValidatorsCountChanged and is used to iterate over the raw logs and unpacked data for TargetValidatorsCountChanged events raised by the Csm contract.
type CsmTargetValidatorsCountChangedIterator struct {
	Event *CsmTargetValidatorsCountChanged // Event containing the contract specifics and raw log

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
func (it *CsmTargetValidatorsCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmTargetValidatorsCountChanged)
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
		it.Event = new(CsmTargetValidatorsCountChanged)
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
func (it *CsmTargetValidatorsCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmTargetValidatorsCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmTargetValidatorsCountChanged represents a TargetValidatorsCountChanged event raised by the Csm contract.
type CsmTargetValidatorsCountChanged struct {
	NodeOperatorId        *big.Int
	TargetLimitMode       *big.Int
	TargetValidatorsCount *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterTargetValidatorsCountChanged is a free log retrieval operation binding the contract event 0xf92eb109ce5b449e9b121c352c6aeb4319538a90738cb95d84f08e41274e92d2.
//
// Solidity: event TargetValidatorsCountChanged(uint256 indexed nodeOperatorId, uint256 targetLimitMode, uint256 targetValidatorsCount)
func (_Csm *CsmFilterer) FilterTargetValidatorsCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmTargetValidatorsCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.FilterLogs(opts, "TargetValidatorsCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmTargetValidatorsCountChangedIterator{contract: _Csm.contract, event: "TargetValidatorsCountChanged", logs: logs, sub: sub}, nil
}

// WatchTargetValidatorsCountChanged is a free log subscription operation binding the contract event 0xf92eb109ce5b449e9b121c352c6aeb4319538a90738cb95d84f08e41274e92d2.
//
// Solidity: event TargetValidatorsCountChanged(uint256 indexed nodeOperatorId, uint256 targetLimitMode, uint256 targetValidatorsCount)
func (_Csm *CsmFilterer) WatchTargetValidatorsCountChanged(opts *bind.WatchOpts, sink chan<- *CsmTargetValidatorsCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.WatchLogs(opts, "TargetValidatorsCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmTargetValidatorsCountChanged)
				if err := _Csm.contract.UnpackLog(event, "TargetValidatorsCountChanged", log); err != nil {
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
func (_Csm *CsmFilterer) ParseTargetValidatorsCountChanged(log types.Log) (*CsmTargetValidatorsCountChanged, error) {
	event := new(CsmTargetValidatorsCountChanged)
	if err := _Csm.contract.UnpackLog(event, "TargetValidatorsCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmTotalSigningKeysCountChangedIterator is returned from FilterTotalSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for TotalSigningKeysCountChanged events raised by the Csm contract.
type CsmTotalSigningKeysCountChangedIterator struct {
	Event *CsmTotalSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *CsmTotalSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmTotalSigningKeysCountChanged)
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
		it.Event = new(CsmTotalSigningKeysCountChanged)
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
func (it *CsmTotalSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmTotalSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmTotalSigningKeysCountChanged represents a TotalSigningKeysCountChanged event raised by the Csm contract.
type CsmTotalSigningKeysCountChanged struct {
	NodeOperatorId *big.Int
	TotalKeysCount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTotalSigningKeysCountChanged is a free log retrieval operation binding the contract event 0xdd01838a366ae4dc9a86e1922512c0716abebc9a440baae0e22d2dec578223f0.
//
// Solidity: event TotalSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 totalKeysCount)
func (_Csm *CsmFilterer) FilterTotalSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmTotalSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.FilterLogs(opts, "TotalSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmTotalSigningKeysCountChangedIterator{contract: _Csm.contract, event: "TotalSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchTotalSigningKeysCountChanged is a free log subscription operation binding the contract event 0xdd01838a366ae4dc9a86e1922512c0716abebc9a440baae0e22d2dec578223f0.
//
// Solidity: event TotalSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 totalKeysCount)
func (_Csm *CsmFilterer) WatchTotalSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *CsmTotalSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.WatchLogs(opts, "TotalSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmTotalSigningKeysCountChanged)
				if err := _Csm.contract.UnpackLog(event, "TotalSigningKeysCountChanged", log); err != nil {
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
func (_Csm *CsmFilterer) ParseTotalSigningKeysCountChanged(log types.Log) (*CsmTotalSigningKeysCountChanged, error) {
	event := new(CsmTotalSigningKeysCountChanged)
	if err := _Csm.contract.UnpackLog(event, "TotalSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmVettedSigningKeysCountChangedIterator is returned from FilterVettedSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for VettedSigningKeysCountChanged events raised by the Csm contract.
type CsmVettedSigningKeysCountChangedIterator struct {
	Event *CsmVettedSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *CsmVettedSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmVettedSigningKeysCountChanged)
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
		it.Event = new(CsmVettedSigningKeysCountChanged)
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
func (it *CsmVettedSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmVettedSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmVettedSigningKeysCountChanged represents a VettedSigningKeysCountChanged event raised by the Csm contract.
type CsmVettedSigningKeysCountChanged struct {
	NodeOperatorId  *big.Int
	VettedKeysCount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVettedSigningKeysCountChanged is a free log retrieval operation binding the contract event 0x947f955eec7e1f626bee3afd2aa47b5de04ddcdd3fe78dc8838213015ef58dfd.
//
// Solidity: event VettedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 vettedKeysCount)
func (_Csm *CsmFilterer) FilterVettedSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmVettedSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.FilterLogs(opts, "VettedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmVettedSigningKeysCountChangedIterator{contract: _Csm.contract, event: "VettedSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchVettedSigningKeysCountChanged is a free log subscription operation binding the contract event 0x947f955eec7e1f626bee3afd2aa47b5de04ddcdd3fe78dc8838213015ef58dfd.
//
// Solidity: event VettedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 vettedKeysCount)
func (_Csm *CsmFilterer) WatchVettedSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *CsmVettedSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.WatchLogs(opts, "VettedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmVettedSigningKeysCountChanged)
				if err := _Csm.contract.UnpackLog(event, "VettedSigningKeysCountChanged", log); err != nil {
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
func (_Csm *CsmFilterer) ParseVettedSigningKeysCountChanged(log types.Log) (*CsmVettedSigningKeysCountChanged, error) {
	event := new(CsmVettedSigningKeysCountChanged)
	if err := _Csm.contract.UnpackLog(event, "VettedSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmVettedSigningKeysCountDecreasedIterator is returned from FilterVettedSigningKeysCountDecreased and is used to iterate over the raw logs and unpacked data for VettedSigningKeysCountDecreased events raised by the Csm contract.
type CsmVettedSigningKeysCountDecreasedIterator struct {
	Event *CsmVettedSigningKeysCountDecreased // Event containing the contract specifics and raw log

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
func (it *CsmVettedSigningKeysCountDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmVettedSigningKeysCountDecreased)
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
		it.Event = new(CsmVettedSigningKeysCountDecreased)
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
func (it *CsmVettedSigningKeysCountDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmVettedSigningKeysCountDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmVettedSigningKeysCountDecreased represents a VettedSigningKeysCountDecreased event raised by the Csm contract.
type CsmVettedSigningKeysCountDecreased struct {
	NodeOperatorId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterVettedSigningKeysCountDecreased is a free log retrieval operation binding the contract event 0xe5725d045d5c47bd1483feba445e395dc8647486963e6d54aad9ed03ff7d6ce6.
//
// Solidity: event VettedSigningKeysCountDecreased(uint256 indexed nodeOperatorId)
func (_Csm *CsmFilterer) FilterVettedSigningKeysCountDecreased(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmVettedSigningKeysCountDecreasedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.FilterLogs(opts, "VettedSigningKeysCountDecreased", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmVettedSigningKeysCountDecreasedIterator{contract: _Csm.contract, event: "VettedSigningKeysCountDecreased", logs: logs, sub: sub}, nil
}

// WatchVettedSigningKeysCountDecreased is a free log subscription operation binding the contract event 0xe5725d045d5c47bd1483feba445e395dc8647486963e6d54aad9ed03ff7d6ce6.
//
// Solidity: event VettedSigningKeysCountDecreased(uint256 indexed nodeOperatorId)
func (_Csm *CsmFilterer) WatchVettedSigningKeysCountDecreased(opts *bind.WatchOpts, sink chan<- *CsmVettedSigningKeysCountDecreased, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.WatchLogs(opts, "VettedSigningKeysCountDecreased", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmVettedSigningKeysCountDecreased)
				if err := _Csm.contract.UnpackLog(event, "VettedSigningKeysCountDecreased", log); err != nil {
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
func (_Csm *CsmFilterer) ParseVettedSigningKeysCountDecreased(log types.Log) (*CsmVettedSigningKeysCountDecreased, error) {
	event := new(CsmVettedSigningKeysCountDecreased)
	if err := _Csm.contract.UnpackLog(event, "VettedSigningKeysCountDecreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmWithdrawalSubmittedIterator is returned from FilterWithdrawalSubmitted and is used to iterate over the raw logs and unpacked data for WithdrawalSubmitted events raised by the Csm contract.
type CsmWithdrawalSubmittedIterator struct {
	Event *CsmWithdrawalSubmitted // Event containing the contract specifics and raw log

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
func (it *CsmWithdrawalSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmWithdrawalSubmitted)
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
		it.Event = new(CsmWithdrawalSubmitted)
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
func (it *CsmWithdrawalSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmWithdrawalSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmWithdrawalSubmitted represents a WithdrawalSubmitted event raised by the Csm contract.
type CsmWithdrawalSubmitted struct {
	NodeOperatorId *big.Int
	KeyIndex       *big.Int
	Amount         *big.Int
	Pubkey         []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalSubmitted is a free log retrieval operation binding the contract event 0x9bc54857932b6f10bb750fdad91f736b82dd4de202ed5c2f9f076773bb5b3fb7.
//
// Solidity: event WithdrawalSubmitted(uint256 indexed nodeOperatorId, uint256 keyIndex, uint256 amount, bytes pubkey)
func (_Csm *CsmFilterer) FilterWithdrawalSubmitted(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmWithdrawalSubmittedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.FilterLogs(opts, "WithdrawalSubmitted", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmWithdrawalSubmittedIterator{contract: _Csm.contract, event: "WithdrawalSubmitted", logs: logs, sub: sub}, nil
}

// WatchWithdrawalSubmitted is a free log subscription operation binding the contract event 0x9bc54857932b6f10bb750fdad91f736b82dd4de202ed5c2f9f076773bb5b3fb7.
//
// Solidity: event WithdrawalSubmitted(uint256 indexed nodeOperatorId, uint256 keyIndex, uint256 amount, bytes pubkey)
func (_Csm *CsmFilterer) WatchWithdrawalSubmitted(opts *bind.WatchOpts, sink chan<- *CsmWithdrawalSubmitted, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csm.contract.WatchLogs(opts, "WithdrawalSubmitted", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmWithdrawalSubmitted)
				if err := _Csm.contract.UnpackLog(event, "WithdrawalSubmitted", log); err != nil {
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

// ParseWithdrawalSubmitted is a log parse operation binding the contract event 0x9bc54857932b6f10bb750fdad91f736b82dd4de202ed5c2f9f076773bb5b3fb7.
//
// Solidity: event WithdrawalSubmitted(uint256 indexed nodeOperatorId, uint256 keyIndex, uint256 amount, bytes pubkey)
func (_Csm *CsmFilterer) ParseWithdrawalSubmitted(log types.Log) (*CsmWithdrawalSubmitted, error) {
	event := new(CsmWithdrawalSubmitted)
	if err := _Csm.contract.UnpackLog(event, "WithdrawalSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

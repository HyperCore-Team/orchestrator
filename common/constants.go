package common

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"orchestrator/common/bridge"
	"strings"
	"time"
)

const (
	// Orchestrator
	WrapCeremonyState = iota
	UnwrapCeremonyState

	DefaultEventsDir            = "events"
	DefaultQueuesDirs           = "queues"
	DefaultTssDir               = "tss"
	DefaultLogsDir              = "logs"
	DefaultAdministratorLogFile = "administrator.log"

	// UnredeemedStatus for unwrap events, unredeemed also means unsent
	UnredeemedStatus uint32 = 0
	// PendingRedeemStatus For wrap events it is the first redeem call, for unwrap events it is when the event is added
	PendingRedeemStatus uint32 = 1
	// RedeemedStatus Redeemed status occurs after the funds have been released on both directions
	RedeemedStatus uint32 = 2
	// RevokedStatus this occurs when and event was revoked and the funds cannot be released anymore
	RevokedStatus uint32 = 3

	// ZNN
	ZenonNetworkName         = "znn"
	EstimatedBlockTimeZnn    = 10 * time.Second
	ConfirmationsRequiredZnn = 2
	WrapRequestsProcessSize  = 30 // sign 30 events at most

	// EVM
	haltedSignatureString   = "Halted()"
	unhaltedSignatureString = "Unhalted()"
	// This event occurs when an users tries to unwrap tokens back to zenon network
	unwrappedSignatureString = "Unwrapped(address,address,string,uint256)"
	// We need this event in case we need to revoke it
	registeredRedeemSignatureString = "RegisteredRedeem(uint256,address,address,uint256)"
	// An entry was redeemed and we can set it as so in the orchestrator local storage
	redeemedSignatureString = "Redeemed(uint256,address,address,uint256)"
	//
	revokedRedeemSignatureString = "RevokedRedeem(uint256)"

	pendingTokenInfoSignatureString           = "PendingTokenInfo(address)"
	setTokenInfoSignatureString               = "SetTokenInfo(address)"
	pendingAdministratorSignatureString       = "PendingAdministrator(address)"
	setAdministratorSignatureString           = "SetAdministrator(address,address)"
	pendingTssSignatureString                 = "PendingTss(address)"
	setTssSignatureString                     = "SetTss(address,address)"
	pendingGuardiansSignatureString           = "PendingGuardians()"
	setGuardiansSignatureString               = "SetGuardians()"
	setAdministratorDelaySignatureString      = "SetAdministratorDelay(uint256)"
	setSoftDelaySignatureString               = "SetSoftDelay(uint256)"
	SetUnhaltDurationSignatureString          = "SetUnhaltDuration(uint256)"
	setEstimatedBlockTimeSignatureString      = "SetEstimatedBlockTime(uint64)"
	setAllowKeyGenSignatureString             = "SetAllowKeyGen(bool)"
	setConfirmationsToFinalitySignatureString = "SetConfirmationsToFinality(uint64)"
)

var (
	// EVM
	EvmContractAbi, _                 = abi.JSON(strings.NewReader(bridge.BridgeMetaData.ABI))
	UnwrapSigHash                     = crypto.Keccak256Hash([]byte(unwrappedSignatureString))
	RegisteredRedeemSigHash           = crypto.Keccak256Hash([]byte(registeredRedeemSignatureString))
	RedeemedSigHash                   = crypto.Keccak256Hash([]byte(redeemedSignatureString))
	RevokedRedeemSigHash              = crypto.Keccak256Hash([]byte(revokedRedeemSignatureString))
	HaltedSigHash                     = crypto.Keccak256Hash([]byte(haltedSignatureString))
	UnhaltedSigHash                   = crypto.Keccak256Hash([]byte(unhaltedSignatureString))
	PendingTokenInfoSigHash           = crypto.Keccak256Hash([]byte(pendingTokenInfoSignatureString))
	SetTokenInfoSigHash               = crypto.Keccak256Hash([]byte(setTokenInfoSignatureString))
	PendingAdministratorSigHash       = crypto.Keccak256Hash([]byte(pendingAdministratorSignatureString))
	SetAdministratorSigHash           = crypto.Keccak256Hash([]byte(setAdministratorSignatureString))
	PendingTssSigHash                 = crypto.Keccak256Hash([]byte(pendingTssSignatureString))
	SetTssSigHash                     = crypto.Keccak256Hash([]byte(setTssSignatureString))
	PendingGuardiansSigHash           = crypto.Keccak256Hash([]byte(pendingGuardiansSignatureString))
	SetGuardiansSigHash               = crypto.Keccak256Hash([]byte(setGuardiansSignatureString))
	SetAdministratorDelaySigHash      = crypto.Keccak256Hash([]byte(setAdministratorDelaySignatureString))
	SetSoftDelaySigHash               = crypto.Keccak256Hash([]byte(setSoftDelaySignatureString))
	SetUnhaltDurationSigHash          = crypto.Keccak256Hash([]byte(SetUnhaltDurationSignatureString))
	SetEstimatedBlockTimeSigHash      = crypto.Keccak256Hash([]byte(setEstimatedBlockTimeSignatureString))
	SetAllowKeyGenSigHash             = crypto.Keccak256Hash([]byte(setAllowKeyGenSignatureString))
	SetConfirmationsToFinalitySigHash = crypto.Keccak256Hash([]byte(setConfirmationsToFinalitySignatureString))

	Topics = [][]common.Hash{{UnwrapSigHash, RegisteredRedeemSigHash, RedeemedSigHash, RevokedRedeemSigHash, HaltedSigHash, UnhaltedSigHash, PendingTokenInfoSigHash, SetTokenInfoSigHash,
		PendingAdministratorSigHash, SetAdministratorSigHash, PendingTssSigHash, SetTssSigHash, PendingGuardiansSigHash, SetGuardiansSigHash}}
)

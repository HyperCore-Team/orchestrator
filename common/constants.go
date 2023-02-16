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

	DefaultEventsDir  = "events"
	DefaultQueuesDirs = "queues"
	DefaultTssDir     = "tss"

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
	haltedSignatureString = "Halted()"
	// This event occurs when an users tries to unwrap tokens back to zenon network
	unwrappedSignatureString = "Unwrapped(address,address,string,uint256)"
	// We need this event in case we need to revoke it
	registeredRedeemSignatureString = "RegisteredRedeem(uint256,address,address,uint256)"
	// An entry was redeemed and we can set it as so in the orchestrator local storage
	redeemedSignatureString = "Redeemed(uint256,address,address,uint256)"
	//
	revokedRedeemSignatureString = "RevokedRedeem(uint256)"
)

var (
	// EVM
	EvmContractAbi, _       = abi.JSON(strings.NewReader(bridge.BridgeMetaData.ABI))
	UnwrapSigHash           = crypto.Keccak256Hash([]byte(unwrappedSignatureString))
	RegisteredRedeemSigHash = crypto.Keccak256Hash([]byte(registeredRedeemSignatureString))
	RedeemedSigHash         = crypto.Keccak256Hash([]byte(redeemedSignatureString))
	RevokedRedeemSigHash    = crypto.Keccak256Hash([]byte(revokedRedeemSignatureString))
	HaltedSigHash           = crypto.Keccak256Hash([]byte(haltedSignatureString))
	Topics                  = [][]common.Hash{{UnwrapSigHash, RegisteredRedeemSigHash, RedeemedSigHash, RevokedRedeemSigHash}}
)

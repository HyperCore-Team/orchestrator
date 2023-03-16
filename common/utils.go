package common

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/pkg/errors"
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/vm/constants"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
	"math/big"
	"orchestrator/common/events"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

func ParseAddressString(address string, networkClass uint32) (interface{}, error) {
	switch networkClass {
	case definition.EvmClass:
		if common.IsHexAddress(address) {
			return common.HexToAddress(address), nil
		}
		return nil, errors.New("not a valid evm address")
	case definition.NoMClass:
		return types.ParseAddress(address)
	}

	return nil, errors.New("config not supported")
}

func homeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	return ""
}

// DefaultDataDir is the default data directory to use for the databases and other persistence requirements.
func DefaultDataDir() string {
	// Try to place the data folder in the user's home dir
	home := homeDir()
	if home != "" {
		switch runtime.GOOS {
		case "darwin":
			panic("darwin not supported")
		case "windows":
			panic("windows not supported")
		default:
			return filepath.Join(home, ".orchestrator")
		}
	}
	// As we cannot guess a stable location, return empty and handle later
	return ""
}

func ZnnWrapToOrchestratorWrap(rpcEvent *definition.WrapTokenRequest) events.WrapRequestZnn {
	return events.WrapRequestZnn{
		NetworkClass:  rpcEvent.NetworkClass,
		ChainId:       rpcEvent.ChainId,
		Id:            rpcEvent.Id,
		ToAddress:     rpcEvent.ToAddress,
		TokenAddress:  rpcEvent.TokenAddress,
		Amount:        big.NewInt(0).Set(rpcEvent.Amount),
		Fee:           big.NewInt(0).Set(rpcEvent.Fee),
		Signature:     rpcEvent.Signature,
		RedeemStatus:  UnredeemedStatus,
		SentSignature: false,
	}
}

func ZnnUnwrapToOrchestatorUnwrap(event *definition.UnwrapTokenRequest) events.UnwrapRequestEvm {
	var status uint32
	if event.Revoked == 1 {
		status = RevokedStatus
	} else if event.Redeemed == 1 {
		status = RedeemedStatus
	} else {
		if len(event.Signature) > 0 {
			status = PendingRedeemStatus
		} else {
			status = UnredeemedStatus
		}
	}
	return events.UnwrapRequestEvm{
		NetworkClass:    event.NetworkClass,
		ChainId:         event.ChainId,
		BlockNumber:     0,
		TransactionHash: common.Hash(event.TransactionHash),
		LogIndex:        event.LogIndex,
		From:            common.Address{},
		To:              event.ToAddress.String(),
		Token:           common.HexToAddress(event.TokenAddress),
		Amount:          event.Amount,
		Signature:       event.Signature,
		RedeemStatus:    status,
	}
}

func GetPublicKeyFilePath(dataDir, pubKey string) (string, error) {
	pubKeyBytes, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return "", err
	}
	pubKeyHex := hex.EncodeToString(pubKeyBytes)
	fileName := fmt.Sprintf("localstate-%s.json", pubKeyHex)
	path := filepath.Join(dataDir, fileName)
	return path, nil
}

func DeletePubKeyFile(dataDir, pubKey string) error {
	filePath, err := GetPublicKeyFilePath(dataDir, pubKey)
	if err != nil {
		return err
	}
	if err := os.Remove(filePath); err != nil {
		return err
	}
	return nil
}

func StateToText(state uint8) string {
	stateStr := ""
	switch state {
	case LiveState:
		stateStr = "LiveState"
	case KeyGenState:
		stateStr = "KeyGenState"
	case HaltedState:
		stateStr = "HaltedState"
	case EmergencyState:
		stateStr = "EmergencyState"
	}
	return stateStr
}

func DecompressPubKey(pubKey string) (string, error) {
	pubKeyBytes, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return "", nil
	} else if len(pubKeyBytes) != constants.CompressedECDSAPubKeyLength {
		return "", nil
	}
	X, Y := secp256k1.DecompressPubkey(pubKeyBytes)
	dPubKeyBytes := make([]byte, 1)
	dPubKeyBytes[0] = 4
	dPubKeyBytes = append(dPubKeyBytes, X.Bytes()...)
	dPubKeyBytes = append(dPubKeyBytes, Y.Bytes()...)
	return base64.StdEncoding.EncodeToString(dPubKeyBytes), nil
}

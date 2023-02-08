package config

import (
	ecommon "github.com/ethereum/go-ethereum/common"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
	"orchestrator/common"
	"time"
)

type EvmParams struct {
	networkName           string
	networkClass, chainId uint32
	contractAddress       ecommon.Address
	// This tells us when should we start looking for events
	contractDeploymentHeight uint64
	estimatedBlockTime       time.Duration
	confirmationsToFinality  uint64
}

func NewEvmParams(network *definition.NetworkInfo) (EvmParams, error) {
	return EvmParams{
		networkName:              network.Name,
		networkClass:             network.Class,
		chainId:                  network.Id,
		contractAddress:          ecommon.HexToAddress(network.ContractAddress),
		contractDeploymentHeight: 0,
		estimatedBlockTime:       0,
		confirmationsToFinality:  0,
	}, nil
}

func (ec *EvmParams) Log() {
	common.GlobalLogger.Infof("ec.networkName: %s", ec.networkName)
	common.GlobalLogger.Infof("ec.networkClass: %d", ec.networkClass)
	common.GlobalLogger.Infof("ec.chainId: %d", ec.chainId)
	common.GlobalLogger.Infof("ec.contractAddress: %s", ec.contractAddress.String())
	common.GlobalLogger.Infof("ec.contractDeploymentHeight: %d", ec.contractDeploymentHeight)
	common.GlobalLogger.Infof("ec.estimatedBlockTime: %d", ec.estimatedBlockTime)
	common.GlobalLogger.Infof("ec.confirmationsToFinality: %d", ec.confirmationsToFinality)
}

func (ec *EvmParams) NetworkName() string {
	return ec.networkName
}

func (ec *EvmParams) NetworkClass() uint32 {
	return ec.networkClass
}

func (ec *EvmParams) ChainId() uint32 {
	return ec.chainId
}

func (ec *EvmParams) ContractAddress() *ecommon.Address {
	return &ec.contractAddress
}

func (ec *EvmParams) ContractDeploymentHeight() uint64 {
	return ec.contractDeploymentHeight
}

func (ec *EvmParams) SetContractDeploymentHeight(height uint64) {
	ec.contractDeploymentHeight = height
}

func (ec *EvmParams) EstimatedBlockTime() time.Duration {
	return ec.estimatedBlockTime
}

func (ec *EvmParams) SetEstimatedBlockTime(blockTime uint64) {
	ec.estimatedBlockTime = time.Second * time.Duration(int64(blockTime))
}

func (ec *EvmParams) ConfirmationsToFinality() uint64 {
	return ec.confirmationsToFinality
}

func (ec *EvmParams) SetConfirmationsToFinality(confirmations uint64) {
	ec.confirmationsToFinality = confirmations
	common.GlobalLogger.Infof("set ec.confirmationsToFinality %d", ec.confirmationsToFinality)
}

type ZnnParams struct {
	windowSize                                                                        uint64
	keyGenThreshold, confirmationsToFinality, estimatedMomentumTime, keySignThreshold uint32
}

func NewZnnParams(orchestratorInfo *definition.OrchestratorInfo) (*ZnnParams, error) {
	return &ZnnParams{
		windowSize:              orchestratorInfo.WindowSize,
		keyGenThreshold:         orchestratorInfo.KeyGenThreshold,
		confirmationsToFinality: orchestratorInfo.ConfirmationsToFinality,
		estimatedMomentumTime:   orchestratorInfo.EstimatedMomentumTime,
		// todo discuss about removing it in orchestrator and go-zenon
		keySignThreshold: orchestratorInfo.KeySignThreshold,
	}, nil
}

func (zP *ZnnParams) SetWindowSize(windowSize uint64) {
	zP.windowSize = windowSize
}
func (zP *ZnnParams) SetKeyGenThreshold(keyGenThreshold uint32) {
	zP.keyGenThreshold = keyGenThreshold
}
func (zP *ZnnParams) SetConfirmationsToFinality(confirmationsToFinality uint32) {
	zP.confirmationsToFinality = confirmationsToFinality
}
func (zP *ZnnParams) SetEstimatedMomentumTime(estimatedMomentumTime uint32) {
	zP.estimatedMomentumTime = estimatedMomentumTime
}
func (zP *ZnnParams) SetKeySignThreshold(keySignThreshold uint32) {
	zP.keySignThreshold = keySignThreshold
}
func (zP *ZnnParams) KeyGenThreshold() uint32 {
	return zP.keyGenThreshold
}
func (zP *ZnnParams) KeySignThreshold() uint32 {
	return zP.keySignThreshold
}
func (zP *ZnnParams) ConfirmationsToFinality() uint32 {
	return zP.confirmationsToFinality
}
func (zP *ZnnParams) EstimatedMomentumTime() uint32 {
	return zP.estimatedMomentumTime
}
func (zP *ZnnParams) WindowSize() uint64 {
	return zP.windowSize
}

package rpc

import (
	ecommon "github.com/ethereum/go-ethereum/common"
	"orchestrator/common"
	"orchestrator/common/config"
	"os"
	"syscall"
)

type Manager struct {
	znnClient  *ZnnRpc
	evmClients map[uint32]*EvmRpc
	stopChan   chan os.Signal
}

func NewRpcManager(networkConfig config.BaseNetworkConfig, stop chan os.Signal) (*Manager, error) {
	newZnnClient, err := NewZnnRpcClient(networkConfig)
	if err != nil {
		return nil, err
	}

	return &Manager{
		znnClient:  newZnnClient,
		evmClients: make(map[uint32]*EvmRpc, 0),
		stopChan:   stop,
	}, nil
}

func (m *Manager) AddEvmClient(networkConfig config.BaseNetworkConfig, chainId uint32, networkName string, address ecommon.Address) error {
	if newClient, err := NewEvmRpcClient(networkConfig, networkName, address); err != nil {
		return err
	} else {
		m.evmClients[chainId] = newClient
	}
	return nil
}

func (m *Manager) RemoveEvmClient(chainId uint32) {
	m.Evm(chainId).Stop()
	m.Evm(chainId).DeleteDirectories()
	delete(m.evmClients, chainId)
}

func (m *Manager) HasEvmNetwork(chainId uint32) bool {
	_, ok := m.evmClients[chainId]
	if !ok {
		return false
	}
	return true
}

func (m *Manager) Znn() *ZnnRpc {
	return m.znnClient
}

func (m *Manager) Evm(chainId uint32) *EvmRpc {
	client, ok := m.evmClients[chainId]
	if !ok || client == nil {
		common.GlobalLogger.Debugf("evm rpc client non existent for chainId %d", chainId)
		m.stopChan <- syscall.SIGINT
	}
	return client
}

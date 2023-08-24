package manager

import (
	zdb "github.com/zenon-network/go-zenon/common/db"
	"go.uber.org/zap"
	"orchestrator/common"
	"orchestrator/common/storage"
	"orchestrator/db"
	"orchestrator/db/unwrap"
	"orchestrator/db/wrap"
	"os"
	"syscall"
)

type Manager struct {
	znnStorage db.ZnnStorage
	// key is the chainId of the network
	evmStorage map[uint32]db.EvmStorage
	stopChan   chan os.Signal
	logger     *zap.SugaredLogger
}

func NewDbManager(stop chan os.Signal) (*Manager, error) {
	newZnnLdb, err := storage.CreateOrOpenLevelDb(common.ZenonNetworkName)
	if err != nil {
		return nil, err
	}
	newLogger, errLog := common.CreateSugarLogger()
	if errLog != nil {
		return nil, errLog
	}

	newDbManager := &Manager{
		znnStorage: wrap.NewZnnStorage(zdb.NewLevelDBWrapper(newZnnLdb), stop),
		evmStorage: make(map[uint32]db.EvmStorage),
		stopChan:   stop,
		logger:     newLogger,
	}
	return newDbManager, nil
}

func (m *Manager) AddEvmEventStore(chainId uint32, name string, contractDeploymentHeight uint64) {
	newLdb, err := storage.CreateOrOpenLevelDb(name)
	if err != nil {
		m.logger.Info("sent SIGINT from here 1")
		m.stopChan <- syscall.SIGINT
		m.logger.Error(err)
		return
	}
	m.evmStorage[chainId] = unwrap.NewEvmStorage(zdb.NewLevelDBWrapper(newLdb), m.stopChan, contractDeploymentHeight)
}

func (m *Manager) ZnnStorage() db.ZnnStorage {
	return m.znnStorage
}

func (m *Manager) EvmStorage(chainId uint32) db.EvmStorage {
	store, ok := m.evmStorage[chainId]
	if !ok {
		m.stopChan <- syscall.SIGINT
		m.logger.Error("evm network storage non existent")
	}
	return store
}

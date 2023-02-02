package unwrap

import (
	zdb "github.com/zenon-network/go-zenon/common/db"
	"orchestrator/db"
	"os"
	"syscall"
)

func getStorageIterator() []byte {
	return unwrapRequestPrefix
}

type evmStorage struct {
	zdb.DB
	stopChan                 chan os.Signal
	contractDeploymentHeight uint64
}

func (es *evmStorage) Storage() zdb.DB {
	return zdb.DisableNotFound(es.DB.Subset(getStorageIterator()))
}
func (es *evmStorage) Snapshot() db.EvmStorage {
	return NewEvmStorage(es.DB.Snapshot(), es.stopChan, es.contractDeploymentHeight)
}

func (es *evmStorage) SendKillSignal() {
	es.stopChan <- syscall.SIGKILL
}

func NewEvmStorage(db zdb.DB, stopChan chan os.Signal, contractDeploymentHeight uint64) db.EvmStorage {
	if db == nil {
		panic("account store can't operate with nil db")
	}
	return &evmStorage{
		DB:                       db,
		stopChan:                 stopChan,
		contractDeploymentHeight: contractDeploymentHeight,
	}
}

package wrap

import (
	zdb "github.com/zenon-network/go-zenon/common/db"
	"orchestrator/db"
	"os"
	"syscall"
)

func getStorageIterator() []byte {
	return wrapEventPrefix
}

// eventStore address is the contract address for the events
type eventStore struct {
	zdb.DB
	stopChan chan os.Signal
}

func (es *eventStore) Storage() zdb.DB {
	return zdb.DisableNotFound(es.DB.Subset(getStorageIterator()))
}
func (es *eventStore) Snapshot() db.ZnnStorage {
	return NewZnnStorage(es.DB.Snapshot(), es.stopChan)
}

func (es *eventStore) SendKillSignal() {
	es.stopChan <- syscall.SIGKILL
}

func NewZnnStorage(db zdb.DB, stopChan chan os.Signal) db.ZnnStorage {
	if db == nil {
		panic("account store can't operate with nil db")
	}
	return &eventStore{
		DB:       db,
		stopChan: stopChan,
	}
}

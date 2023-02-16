package storage

import (
	"github.com/joncrlsn/dque"
	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb"
	lerrors "github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
	"math"
	"orchestrator/common"
	"orchestrator/common/events"
	"os"
	"path/filepath"
)

func CreateOrOpenQueue(networkClass uint32, networkName string) (*dque.DQue, error) {
	qDir := filepath.Join(common.DefaultDataDir(), common.DefaultQueuesDirs)
	if _, err := os.Stat(qDir); os.IsNotExist(err) {
		err := os.MkdirAll(qDir, 0700)
		if err != nil {
			return nil, err
		}
	}

	segmentSize := math.MaxInt // todo how big?

	var queue *dque.DQue
	var err error
	switch networkClass {
	case definition.EvmClass:
		queue, err = dque.NewOrOpen(networkName, qDir, segmentSize, events.UnwrapRequestEvmBuilder)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("config not supported")
	}
	return queue, nil
}

func DeleteQueue(name string) error {
	queuePath := filepath.Join(common.DefaultDataDir(), common.DefaultQueuesDirs, name)
	return os.RemoveAll(queuePath)
}

func CreateOrOpenLevelDb(name string) (*leveldb.DB, error) {
	opts := &opt.Options{OpenFilesCacheCapacity: 200}
	evDir := filepath.Join(common.DefaultDataDir(), common.DefaultEventsDir)
	if _, err := os.Stat(evDir); os.IsNotExist(err) {
		if err = os.MkdirAll(evDir, 0700); err != nil {
			return nil, err
		}
	}
	dbDir := filepath.Join(evDir, name)
	ldb, err := leveldb.OpenFile(dbDir, opts)
	if _, isCorrupted := err.(*lerrors.ErrCorrupted); isCorrupted {
		ldb, err = leveldb.RecoverFile(dbDir, nil)
		if err != nil {
			return nil, err
		}
	}

	return ldb, nil
}

func DeleteLvlDb(name string) error {
	dbPath := filepath.Join(common.DefaultDataDir(), common.DefaultEventsDir, name)
	return os.RemoveAll(dbPath)
}

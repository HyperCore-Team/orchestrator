package unwrap

import (
	"encoding/binary"
	ecommon "github.com/ethereum/go-ethereum/common"
	"github.com/syndtr/goleveldb/leveldb"
	zcommon "github.com/zenon-network/go-zenon/common"
	"orchestrator/common"
	"orchestrator/common/events"
)

func getUnwrapRequestPrefix() []byte {
	return zcommon.JoinBytes(unwrapRequestPrefix)
}

func getUnwrapRequestKey(txHash ecommon.Hash, logIndex uint32) []byte {
	logNumberBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(logNumberBytes, logIndex)
	return zcommon.JoinBytes(getUnwrapRequestPrefix(), txHash.Bytes(), logNumberBytes)
}

func (es *evmStorage) AddUnwrapRequest(event events.UnwrapRequestEvm) error {
	if eventBytes, err := event.Serialize(); err != nil {
		es.SendKillSignal()
		return err
	} else {
		if err := es.DB.Put(getUnwrapRequestKey(event.TransactionHash, event.LogIndex), eventBytes); err != nil {
			es.SendKillSignal()
			return err
		}
	}
	return nil
}

func (es *evmStorage) UpdateUnwrapRequestBlockNumber(event events.UnwrapRequestEvm) error {
	localEvent, err := es.GetUnwrapRequestByHashAndLog(event.TransactionHash, event.LogIndex)
	if err != nil {
		es.SendKillSignal()
		return err
	}
	if localEvent == nil {
		return es.AddUnwrapRequest(event)
	}
	localEvent.BlockNumber = event.BlockNumber
	localEventBytes, err := localEvent.Serialize()
	if err != nil {
		es.SendKillSignal()
		return err
	}
	if err := es.DB.Put(getUnwrapRequestKey(event.TransactionHash, event.LogIndex), localEventBytes); err != nil {
		es.SendKillSignal()
		return err
	}
	return nil
}

func (es *evmStorage) GetUnwrapRequestByHashAndLog(txHash ecommon.Hash, logIndex uint32) (*events.UnwrapRequestEvm, error) {
	data, err := es.DB.Get(getUnwrapRequestKey(txHash, logIndex))
	if err == leveldb.ErrNotFound {
		return nil, nil
	}

	if err != nil {
		es.SendKillSignal()
		return nil, err
	}

	event, err := events.DeserializeEvmUnwrapRequest(data)
	if err != nil {
		es.SendKillSignal()
		return nil, err
	}
	return event, nil
}

func (es *evmStorage) SetUnwrapRequestStatus(txHash ecommon.Hash, logIndex, status uint32) error {
	if event, err := es.GetUnwrapRequestByHashAndLog(txHash, logIndex); err != nil {
		es.SendKillSignal()
		return err
	} else if event == nil {
		return leveldb.ErrNotFound
	} else {
		event.RedeemStatus = status
		if eventBytes, err := event.Serialize(); err != nil {
			es.SendKillSignal()
			return err
		} else {
			if err := es.DB.Put(getUnwrapRequestKey(event.TransactionHash, event.LogIndex), eventBytes); err != nil {
				es.SendKillSignal()
				return err
			}
		}
	}
	return nil
}

func (es *evmStorage) SetUnwrapRequestSignature(txHash ecommon.Hash, logIndex uint32, signature string) error {
	if event, err := es.GetUnwrapRequestByHashAndLog(txHash, logIndex); err != nil {
		es.SendKillSignal()
		return err
	} else if event == nil {
		return leveldb.ErrNotFound
	} else {
		event.Signature = signature
		if eventBytes, err := event.Serialize(); err != nil {
			es.SendKillSignal()
			return err
		} else {
			if err := es.DB.Put(getUnwrapRequestKey(event.TransactionHash, event.LogIndex), eventBytes); err != nil {
				es.SendKillSignal()
				return err
			}
		}
	}
	return nil
}

func (es *evmStorage) SetUnsentUnwrapRequestAsUnsigned(txHash ecommon.Hash, logIndex uint32) error {
	if event, err := es.GetUnwrapRequestByHashAndLog(txHash, logIndex); err != nil {
		es.SendKillSignal()
		return err
	} else if event == nil {
		return leveldb.ErrNotFound
	} else {
		event.Signature = ""
		event.RedeemStatus = common.UnredeemedStatus
		if eventBytes, err := event.Serialize(); err != nil {
			es.SendKillSignal()
			return err
		} else {
			if err := es.DB.Put(getUnwrapRequestKey(event.TransactionHash, event.LogIndex), eventBytes); err != nil {
				es.SendKillSignal()
				return err
			}
		}
	}
	return nil
}

func (es *evmStorage) GetUnwrapRequestsByStatus(status uint32) ([]*events.UnwrapRequestEvm, error) {
	iterator := es.DB.NewIterator(getUnwrapRequestPrefix())
	defer iterator.Release()
	result := make([]*events.UnwrapRequestEvm, 0)

	for {
		if !iterator.Next() {
			if iterator.Error() != nil {
				es.SendKillSignal()
				return nil, iterator.Error()
			}
			break
		}
		if iterator.Value() == nil {
			continue
		}

		event, err := events.DeserializeEvmUnwrapRequest(iterator.Value())
		if err != nil {
			es.SendKillSignal()
			return nil, err
		}
		if event.RedeemStatus != status {
			continue
		}

		result = append(result, event)
	}
	return result, nil
}

func (es *evmStorage) GetUnsignedUnwrapRequests() ([]*events.UnwrapRequestEvm, error) {
	iterator := es.DB.NewIterator(getUnwrapRequestPrefix())
	defer iterator.Release()
	result := make([]*events.UnwrapRequestEvm, 0)

	for {
		if !iterator.Next() {
			if iterator.Error() != nil {
				es.SendKillSignal()
				return nil, iterator.Error()
			}
			break
		}
		if iterator.Value() == nil {
			continue
		}

		event, err := events.DeserializeEvmUnwrapRequest(iterator.Value())
		if err != nil {
			es.SendKillSignal()
			return nil, err
		}
		if len(event.Signature) > 0 {
			continue
		}

		result = append(result, event)
	}
	return result, nil
}

func (es *evmStorage) GetUnsentSignedUnwrapRequests() ([]*events.UnwrapRequestEvm, error) {
	iterator := es.DB.NewIterator(getUnwrapRequestPrefix())
	defer iterator.Release()
	result := make([]*events.UnwrapRequestEvm, 0)

	for {
		if !iterator.Next() {
			if iterator.Error() != nil {
				es.SendKillSignal()
				return nil, iterator.Error()
			}
			break
		}
		if iterator.Value() == nil {
			continue
		}

		event, err := events.DeserializeEvmUnwrapRequest(iterator.Value())
		if err != nil {
			es.SendKillSignal()
			return nil, err
		}
		// unredeemed means unsent
		if len(event.Signature) > 0 && event.RedeemStatus == common.UnredeemedStatus {
			result = append(result, event)
		}
	}
	return result, nil
}

func getLastUpdateKey() []byte {
	return zcommon.JoinBytes(lastUpdatePrefix)
}

func (es *evmStorage) GetLastUpdateHeight() (uint64, error) {
	data, err := es.DB.Get(getLastUpdateKey())
	if err == leveldb.ErrNotFound {
		return es.contractDeploymentHeight, nil
	}

	if err != nil {
		es.SendKillSignal()
		return 0, err
	}

	return binary.LittleEndian.Uint64(data), nil
}

func (es *evmStorage) SetLastUpdateHeight(blockHeight uint64) error {
	if _, err := es.GetLastUpdateHeight(); err != nil {
		es.SendKillSignal()
		return err
	} else {
		bytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(bytes, blockHeight)
		if err := es.DB.Put(getLastUpdateKey(), bytes); err != nil {
			es.SendKillSignal()
			return err
		}
	}
	return nil
}

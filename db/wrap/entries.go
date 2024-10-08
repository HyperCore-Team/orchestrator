package wrap

import (
	"encoding/binary"
	"errors"
	"github.com/syndtr/goleveldb/leveldb"
	zcommon "github.com/zenon-network/go-zenon/common"
	"github.com/zenon-network/go-zenon/common/types"
	"orchestrator/common"
	"orchestrator/common/events"
)

func getWrapEventPrefix() []byte {
	return zcommon.JoinBytes(wrapEventPrefix)
}

func getWrapEventKey(id types.Hash) []byte {
	return zcommon.JoinBytes(getWrapEventPrefix(), id.Bytes())
}

func (es *eventStore) AddWrapRequest(event events.WrapRequestZnn) error {
	if eventBytes, err := event.Serialize(); err != nil {
		es.SendSigInt()
		return err
	} else {
		if err := es.DB.Put(getWrapEventKey(event.Id), eventBytes); err != nil {
			es.SendSigInt()
			return err
		}
	}
	return nil
}

func (es *eventStore) GetWrapRequestById(id types.Hash) (*events.WrapRequestZnn, error) {
	data, err := es.DB.Get(getWrapEventKey(id))
	if err == leveldb.ErrNotFound {
		return nil, nil
	}
	if err != nil {
		es.SendSigInt()
		return nil, err
	}

	event, err := events.DeserializeWrapEventZnn(data)
	if err != nil {
		es.SendSigInt()
		return nil, err
	}
	return event, nil
}

func (es *eventStore) SetWrapRequestStatus(id types.Hash, status uint32) error {
	if event, err := es.GetWrapRequestById(id); err != nil {
		es.SendSigInt()
		return err
	} else {
		if event == nil {
			return leveldb.ErrNotFound
		}
		event.RedeemStatus = status
		if eventBytes, err := event.Serialize(); err != nil {
			es.SendSigInt()
			return err
		} else {
			if err := es.DB.Put(getWrapEventKey(event.Id), eventBytes); err != nil {
				es.SendSigInt()
				return err
			}
		}
	}
	return nil
}

func (es *eventStore) SetWrapRequestSignature(id types.Hash, signature string) error {
	if event, err := es.GetWrapRequestById(id); err != nil {
		es.SendSigInt()
		return err
	} else {
		if event == nil {
			return leveldb.ErrNotFound
		}
		event.Signature = signature
		if eventBytes, err := event.Serialize(); err != nil {
			es.SendSigInt()
			return err
		} else {
			if err := es.DB.Put(getWrapEventKey(event.Id), eventBytes); err != nil {
				es.SendSigInt()
				return err
			}
		}
	}
	return nil
}

func (es *eventStore) SetWrapRequestSentSignature(id types.Hash, value bool) error {
	if event, err := es.GetWrapRequestById(id); err != nil {
		es.SendSigInt()
		return err
	} else {
		if event == nil {
			return leveldb.ErrNotFound
		}
		event.SentSignature = value
		if eventBytes, err := event.Serialize(); err != nil {
			es.SendSigInt()
			return err
		} else {
			if err := es.DB.Put(getWrapEventKey(event.Id), eventBytes); err != nil {
				es.SendSigInt()
				return err
			}
		}
	}
	return nil
}

func (es *eventStore) GetUnsentSignedWrapRequests() ([]*events.WrapRequestZnn, error) {
	iterator := es.DB.NewIterator(getWrapEventPrefix())
	defer iterator.Release()
	result := make([]*events.WrapRequestZnn, 0)

	for {
		if !iterator.Next() {
			if iterator.Error() != nil {
				es.SendSigInt()
				return nil, iterator.Error()
			}
			break
		}
		if iterator.Value() == nil {
			continue
		}

		event, err := events.DeserializeWrapEventZnn(iterator.Value())
		if err != nil {
			es.SendSigInt()
			return nil, err
		}
		if event.SentSignature || len(event.Signature) == 0 {
			continue
		}

		result = append(result, event)
	}
	return result, nil
}

func (es *eventStore) GetResignableWrapRequests() ([]*events.WrapRequestZnn, error) {
	iterator := es.DB.NewIterator(getWrapEventPrefix())
	defer iterator.Release()
	result := make([]*events.WrapRequestZnn, 0)

	for {
		if !iterator.Next() {
			if iterator.Error() != nil {
				es.SendSigInt()
				return nil, iterator.Error()
			}
			break
		}
		if iterator.Value() == nil {
			continue
		}

		event, err := events.DeserializeWrapEventZnn(iterator.Value())
		if err != nil {
			es.SendSigInt()
			return nil, err
		}
		resignable, err := es.GetResignStatus(event.Id)
		if err != nil {
			es.SendSigInt()
			return nil, err
		}
		if !resignable {
			continue
		}

		result = append(result, event)
	}
	return result, nil
}

func (es *eventStore) GetUnredeemedWrapRequests() ([]*events.WrapRequestZnn, error) {
	iterator := es.DB.NewIterator(getWrapEventPrefix())
	defer iterator.Release()
	result := make([]*events.WrapRequestZnn, 0)

	for {
		if !iterator.Next() {
			if iterator.Error() != nil {
				es.SendSigInt()
				return nil, iterator.Error()
			}
			break
		}
		if iterator.Value() == nil {
			continue
		}

		event, err := events.DeserializeWrapEventZnn(iterator.Value())
		if err != nil {
			es.SendSigInt()
			return nil, err
		}
		if event.RedeemStatus != common.UnredeemedStatus {
			continue
		}

		result = append(result, event)
	}
	return result, nil
}

func (es *eventStore) GetUnsignedWrapRequests() ([]*events.WrapRequestZnn, error) {
	iterator := es.DB.NewIterator(getWrapEventPrefix())
	defer iterator.Release()
	result := make([]*events.WrapRequestZnn, 0)

	for {
		if !iterator.Next() {
			if iterator.Error() != nil {
				es.SendSigInt()
				return nil, iterator.Error()
			}
			break
		}
		if iterator.Value() == nil {
			continue
		}

		event, err := events.DeserializeWrapEventZnn(iterator.Value())
		if err != nil {
			es.SendSigInt()
			return nil, err
		}
		if len(event.Signature) > 0 {
			continue
		}

		result = append(result, event)
	}
	return result, nil
}

func getLastUpdateKey() []byte {
	return zcommon.JoinBytes(lastUpdatePrefix)
}

func (es *eventStore) GetLastUpdateHeight() (uint64, error) {
	data, err := es.DB.Get(getLastUpdateKey())
	if err == leveldb.ErrNotFound {
		return 1, nil
	}

	if err != nil {
		es.SendSigInt()
		return 0, err
	}

	return binary.LittleEndian.Uint64(data), nil
}

func (es *eventStore) SetLastUpdateHeight(accBlHeight uint64) error {
	if _, err := es.GetLastUpdateHeight(); err != nil {
		return err
	} else {
		bytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(bytes, accBlHeight)
		if err := es.DB.Put(getLastUpdateKey(), bytes); err != nil {
			es.SendSigInt()
			return err
		}
	}
	return nil
}

func getResignedStatusKey(id types.Hash) []byte {
	return zcommon.JoinBytes(resignedStatusPrefix, id.Bytes())
}

func (es *eventStore) GetResignStatus(id types.Hash) (bool, error) {
	data, err := es.DB.Get(getResignedStatusKey(id))
	if errors.Is(err, leveldb.ErrNotFound) {
		return false, nil
	}

	if err != nil {
		es.SendSigInt()
		return false, err
	}
	value := binary.LittleEndian.Uint32(data)
	ans := false
	if value > 0 {
		ans = true
	}
	return ans, nil
}

func (es *eventStore) SetResignStatus(id types.Hash, status bool) error {
	bytes := make([]byte, 4)
	value := uint32(0)
	if status {
		value = 1
	}
	binary.LittleEndian.PutUint32(bytes, value)
	if err := es.DB.Put(getResignedStatusKey(id), bytes); err != nil {
		es.SendSigInt()
		return err
	}
	return nil
}

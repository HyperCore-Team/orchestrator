package events

import (
	ecommon "github.com/ethereum/go-ethereum/common"
	"github.com/golang/protobuf/proto"
	zcommon "github.com/zenon-network/go-zenon/common"
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
	"github.com/zenon-network/go-zenon/vm/embedded/implementation"
	"math/big"
)

// WrapRequestZnn znn -> wZnn networkType and ChainId belong to the destination network
type WrapRequestZnn struct {
	NetworkType   uint32
	ChainId       uint32
	Id            types.Hash // hash of the send block
	ToAddress     string
	TokenAddress  string
	Amount        *big.Int
	Fee           *big.Int
	Signature     string
	RedeemStatus  uint32
	SentSignature bool
}

func (e *WrapRequestZnn) GetMessage(address *ecommon.Address) ([]byte, error) {
	request := &definition.WrapTokenRequest{
		NetworkType:  e.NetworkType,
		ChainId:      e.ChainId,
		Id:           e.Id,
		ToAddress:    e.ToAddress,
		TokenAddress: e.TokenAddress,
		Amount:       e.Amount,
		Fee:          e.Fee,
	}
	return implementation.GetWrapTokenRequestMessage(request, address)
}

func (e *WrapRequestZnn) Proto() *WrapRequestZnnProto {
	return &WrapRequestZnnProto{
		NetworkType:   e.NetworkType,
		ChainId:       e.ChainId,
		Id:            e.Id.Bytes(),
		ToAddress:     e.ToAddress,
		TokenAddress:  e.TokenAddress,
		Amount:        e.Amount.Bytes(),
		Fee:           e.Fee.Bytes(),
		Signature:     e.Signature,
		Status:        e.RedeemStatus,
		SentSignature: e.SentSignature,
	}
}

func DeProtoWrapEventZnn(e *WrapRequestZnnProto) *WrapRequestZnn {
	ev := &WrapRequestZnn{
		NetworkType:   e.NetworkType,
		ChainId:       e.ChainId,
		Id:            types.BytesToHashPanic(e.Id),
		ToAddress:     e.ToAddress,
		TokenAddress:  e.TokenAddress,
		Amount:        zcommon.BytesToBigInt(e.Amount),
		Fee:           zcommon.BytesToBigInt(e.Fee),
		Signature:     e.Signature,
		RedeemStatus:  e.Status,
		SentSignature: e.SentSignature,
	}
	return ev
}

func (e *WrapRequestZnn) Serialize() ([]byte, error) {
	return proto.Marshal(e.Proto())
}
func DeserializeWrapEventZnn(data []byte) (*WrapRequestZnn, error) {
	ev := &WrapRequestZnnProto{}
	if err := proto.Unmarshal(data, ev); err != nil {
		return nil, err
	}
	return DeProtoWrapEventZnn(ev), nil
}

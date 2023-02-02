package events

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	ecommon "github.com/ethereum/go-ethereum/common"
	"github.com/golang/protobuf/proto"
	zcommon "github.com/zenon-network/go-zenon/common"
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
	"github.com/zenon-network/go-zenon/vm/embedded/implementation"
	"math/big"
)

// UnwrapRequestEvm This event occurs when users try to go from wZNN to ZNN and they call the smart contract method for unwrapping
// txHash - transaction hash
// blockHeight - block height in which the transaction was included
// networkId - id representing the evm network
// fromAddress - evm tx sender
// ToAddress - Zenon Address
// token - evm token address
// Amount - Amount to be unwrapped
type UnwrapRequestEvm struct {
	NetworkType     uint32          `json:"networkType"`
	ChainId         uint32          `json:"networkId"`
	BlockNumber     uint64          `json:"blockNumber"`
	BlockHash       ecommon.Hash    `json:"blockHash"`
	TransactionHash ecommon.Hash    `json:"transactionHash"`
	LogIndex        uint32          `json:"logIndex"`
	From            ecommon.Address `json:"from"`
	To              string          `json:"to"`
	Token           ecommon.Address `json:"token"`
	Amount          *big.Int        `json:"amount"`
	Signature       string          `json:"signature"`
	RedeemStatus    uint32          `json:"redeemStatus"`
}

// UnwrapRequestEvmBuilder creates a new event and returns a pointer to it.
// This is used when we load a segment of the queue from disk.
func UnwrapRequestEvmBuilder() interface{} {
	return &UnwrapRequestEvm{}
}

func (e *UnwrapRequestEvm) Hash() (string, error) {
	args := abi.Arguments{{Type: definition.Uint256Ty}, {Type: definition.Uint256Ty}, {Type: definition.AddressTy}, {Type: definition.AddressTy}, {Type: definition.Uint256Ty}}
	values := make([]interface{}, 0)
	values = append(values,
		big.NewInt(0).SetBytes(e.TransactionHash.Bytes()),
		big.NewInt(0).SetUint64(uint64(e.LogIndex)),
		e.From,
		ecommon.HexToAddress(e.To),
		big.NewInt(0).SetBytes(e.Amount.Bytes()),
	)

	messageBytes, err := args.PackValues(values)
	if err != nil {
		return "", err
	}
	hash := types.NewHash(messageBytes)
	return hash.String(), nil
}

func (e *UnwrapRequestEvm) GetMessage() ([]byte, error) {
	toAddress, err := types.ParseAddress(e.To)
	if err != nil {
		return nil, err
	}
	param := &definition.UnwrapTokenParam{
		NetworkType:     e.NetworkType,
		ChainId:         e.ChainId,
		TransactionHash: types.Hash(e.TransactionHash),
		LogIndex:        e.LogIndex,
		ToAddress:       toAddress,
		TokenAddress:    e.Token.String(),
		Amount:          e.Amount,
	}
	return implementation.GetUnwrapTokenRequestMessage(param)
}

func (e *UnwrapRequestEvm) Proto() *UnwrapRequestEvmProto {
	return &UnwrapRequestEvmProto{
		NetworkType:     e.NetworkType,
		ChainId:         e.ChainId,
		BlockNumber:     e.BlockNumber,
		BlockHash:       e.BlockHash.Bytes(),
		TransactionHash: e.TransactionHash.Bytes(),
		LogIndex:        e.LogIndex,
		From:            e.From.Bytes(),
		To:              e.To,
		Token:           e.Token.Bytes(),
		Amount:          zcommon.BigIntToBytes(e.Amount),
		Signature:       e.Signature,
		RedeemStatus:    e.RedeemStatus,
	}
}

func DeProtoEvmUnwrapRequest(e *UnwrapRequestEvmProto) *UnwrapRequestEvm {
	ev := &UnwrapRequestEvm{
		NetworkType:     e.NetworkType,
		ChainId:         e.ChainId,
		BlockNumber:     e.BlockNumber,
		BlockHash:       ecommon.BytesToHash(e.BlockHash),
		TransactionHash: ecommon.BytesToHash(e.TransactionHash),
		LogIndex:        e.LogIndex,
		From:            ecommon.BytesToAddress(e.From),
		To:              e.To,
		Token:           ecommon.BytesToAddress(e.Token),
		Amount:          zcommon.BytesToBigInt(e.Amount),
		Signature:       e.Signature,
		RedeemStatus:    e.RedeemStatus,
	}
	return ev
}

func (e *UnwrapRequestEvm) Serialize() ([]byte, error) {
	return proto.Marshal(e.Proto())
}
func DeserializeEvmUnwrapRequest(data []byte) (*UnwrapRequestEvm, error) {
	ev := &UnwrapRequestEvmProto{}
	if err := proto.Unmarshal(data, ev); err != nil {
		return nil, err
	}
	return DeProtoEvmUnwrapRequest(ev), nil
}

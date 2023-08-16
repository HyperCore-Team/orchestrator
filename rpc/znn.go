package rpc

import (
	"encoding/base64"
	sdk_rpc_client "github.com/MoonBaZZe/znn-sdk-go/rpc_client"
	"github.com/MoonBaZZe/znn-sdk-go/zenon"
	ic "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/pkg/errors"
	"github.com/zenon-network/go-zenon/chain/nom"
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/protocol"
	"github.com/zenon-network/go-zenon/rpc/api"
	"github.com/zenon-network/go-zenon/rpc/api/embedded"
	"github.com/zenon-network/go-zenon/rpc/api/subscribe"
	"github.com/zenon-network/go-zenon/rpc/server"
	"github.com/zenon-network/go-zenon/vm/constants"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
	"github.com/zenon-network/go-zenon/wallet"
	"math"
	"orchestrator/common"
	"orchestrator/common/config"
	"orchestrator/common/events"
)

type ZnnRpc struct {
	rpcClient *sdk_rpc_client.RpcClient
	Urls      config.UrlsInfo

	momentumsSub  *server.ClientSubscription
	momentumsChan chan []subscribe.Momentum

	accountBlocksSub  *server.ClientSubscription
	accountBlocksChan chan []subscribe.AccountBlock
}

func NewZnnRpcClient(networkConfig config.BaseNetworkConfig) (*ZnnRpc, error) {
	newUrls, err := config.NewUrlsInfo(networkConfig)
	if err != nil {
		return nil, err
	}
	var newZnnClient *sdk_rpc_client.RpcClient
	currentUrl := newUrls.GetCurrentUrl()
	for {
		newZnnClient, err = sdk_rpc_client.NewRpcClient(currentUrl)
		if err != nil {
			common.GlobalLogger.Infof("Error when dialing %s, got: %s\n", currentUrl, err)
		} else {
			break
		}
		currentUrl = newUrls.NextUrl()
		if len(currentUrl) == 0 {
			return nil, errors.New("cannot connect to any urls to a znn node")
		}
	}

	newUrls.Clear()
	return &ZnnRpc{
		rpcClient: newZnnClient,
		Urls:      *newUrls,
	}, nil
}

/// Utils

func (r *ZnnRpc) Stop() error {
	close(r.momentumsChan)
	close(r.accountBlocksChan)
	if r.rpcClient == nil {
		return errors.New("znn rpc client is nil")
	}
	r.rpcClient.Stop()
	return nil
}

func (r *ZnnRpc) Reconnect() error {
	// todo?this should have a semaphore
	return nil
}

/// Subscriptions

func (r *ZnnRpc) SubscribeToMomentums() (*server.ClientSubscription, chan []subscribe.Momentum, error) {
	var err error
	r.momentumsSub, r.momentumsChan, err = r.rpcClient.SubscriberApi.ToMomentums()
	if err != nil {
		return nil, nil, err
	}
	return r.momentumsSub, r.momentumsChan, nil
}

func (r *ZnnRpc) SubscribeToAccountBlocks(address types.Address) (*server.ClientSubscription, chan []subscribe.AccountBlock, error) {
	var err error
	r.accountBlocksSub, r.accountBlocksChan, err = r.rpcClient.SubscriberApi.ToAccountBlocksByAddress(address)
	if err != nil {
		return nil, nil, err
	}
	return r.accountBlocksSub, r.accountBlocksChan, nil
}

/// Transactions

func (r *ZnnRpc) BroadcastTransaction(tx *nom.AccountBlock, keyPair *wallet.KeyPair) error {
	if err := zenon.CheckAndSetFields(r.rpcClient, tx, keyPair.Address, keyPair.Public); err != nil {
		return err
	}
	if err := zenon.SetDifficulty(r.rpcClient, tx); err != nil {
		return err
	}

	tx.Hash = tx.ComputeHash()
	tx.Signature = keyPair.Sign(tx.Hash.Bytes())

	return r.rpcClient.LedgerApi.PublishRawTransaction(tx)
}

func (r *ZnnRpc) UpdateWrapRequest(id types.Hash, signature string, keyPair *wallet.KeyPair) error {
	tx := r.rpcClient.BridgeApi.UpdateWrapRequest(id, signature)
	return r.BroadcastTransaction(tx, keyPair)
}

func (r *ZnnRpc) SendUnwrapRequest(event *events.UnwrapRequestEvm, keyPair *wallet.KeyPair) error {
	hash := types.BytesToHashPanic(event.TransactionHash.Bytes())
	toAddress := types.ParseAddressPanic(event.To)
	tx := r.rpcClient.BridgeApi.UnwrapToken(event.NetworkClass, event.ChainId, event.Token.String(), hash, event.LogIndex, event.Amount, toAddress, event.Signature)
	return r.BroadcastTransaction(tx, keyPair)
}

func (r *ZnnRpc) ChangeTssEcdsaPubKey(pubKey, signature, newSignature string, keyPair *wallet.KeyPair) error {
	tx := r.rpcClient.BridgeApi.ChangeTssECDSAPubKey(pubKey, signature, newSignature)
	return r.BroadcastTransaction(tx, keyPair)
}

func (r *ZnnRpc) Halt(signature string, keyPair *wallet.KeyPair) error {
	tx := r.rpcClient.BridgeApi.Halt(signature)
	return r.BroadcastTransaction(tx, keyPair)
}

/// RPC Calls

func (r *ZnnRpc) GetBridgeInfo() (*definition.BridgeInfoVariable, error) {
	return r.rpcClient.BridgeApi.GetBridgeInfo()
}

func (r *ZnnRpc) GetSyncInfo() (*protocol.SyncInfo, error) {
	return r.rpcClient.StatsApi.SyncInfo()
}

func (r *ZnnRpc) GetOrchestratorInfo() (*definition.OrchestratorInfo, error) {
	orchestratorInfo, err := r.rpcClient.BridgeApi.GetOrchestratorInfo()
	if err != nil {
		return nil, err
	}
	if orchestratorInfo.WindowSize == 0 || orchestratorInfo.KeyGenThreshold == 0 || orchestratorInfo.ConfirmationsToFinality == 0 || orchestratorInfo.EstimatedMomentumTime == 0 {
		return nil, constants.ErrOrchestratorNotInitialized
	}
	return orchestratorInfo, nil
}

func (r *ZnnRpc) GetSecurityInfo() (*definition.SecurityInfoVariable, error) {
	return r.rpcClient.BridgeApi.GetSecurityInfo()
}

func (r *ZnnRpc) GetAllUnsignedWrapTokenRequests(pageIndex, pageSize uint32) (*embedded.WrapTokenRequestList, error) {
	return r.rpcClient.BridgeApi.GetAllUnsignedWrapTokenRequests(pageIndex, pageSize)
}

func (r *ZnnRpc) GetWrapTokenRequestById(id types.Hash) (*definition.WrapTokenRequest, error) {
	return r.rpcClient.BridgeApi.GetWrapTokenRequestById(id)
}

func (r *ZnnRpc) GetUnwrapTokenRequestByHashAndLog(txHash types.Hash, logIndex uint32) (*definition.UnwrapTokenRequest, error) {
	return r.rpcClient.BridgeApi.GetUnwrapTokenRequestByHashAndLog(txHash, logIndex)
}

func (r *ZnnRpc) GetAllWrapTokenRequests(pageIndex, pageSize uint32) (*embedded.WrapTokenRequestList, error) {
	return r.rpcClient.BridgeApi.GetAllWrapTokenRequests(pageIndex, pageSize)
}

func (r *ZnnRpc) GetAccountBlockByHash(hash types.Hash) (*api.AccountBlock, error) {
	return r.rpcClient.LedgerApi.GetAccountBlockByHash(hash)
}

func (r *ZnnRpc) GetAccountBlocksByHeight(address types.Address, height, count uint64) (*api.AccountBlockList, error) {
	return r.rpcClient.LedgerApi.GetAccountBlocksByHeight(address, height, count)
}

func (r *ZnnRpc) GetMomentumsByHeight(height, count uint64) (*api.MomentumList, error) {
	return r.rpcClient.LedgerApi.GetMomentumsByHeight(height, count)
}

func (r *ZnnRpc) GetFrontierMomentum() (*api.Momentum, error) {
	return r.rpcClient.LedgerApi.GetFrontierMomentum()
}

func (r *ZnnRpc) GetAllNetworks() ([]*definition.NetworkInfo, error) {
	page := uint32(0)
	ans := make([]*definition.NetworkInfo, 0)

	for {
		networks, err := r.rpcClient.BridgeApi.GetAllNetworks(page, api.RpcMaxCountSize)
		if err != nil {
			return nil, err
		}
		if len(networks.List) == 0 {
			break
		}
		page += 1
		for _, network := range networks.List {
			ans = append(ans, network)
		}
	}
	return ans, nil
}

func (r *ZnnRpc) GetNetworkByClassAndId(networkClass, id uint32) (*definition.NetworkInfo, error) {
	network, err := r.rpcClient.BridgeApi.GetNetworkInfo(networkClass, id)
	if err != nil {
		return nil, err
	} else if network.NetworkClass == 0 {
		return nil, nil
	}

	return network, nil
}

func (r *ZnnRpc) GetPillarPublicKeys() (map[string]string, error) {
	orchestratorInfo, err := r.GetOrchestratorInfo()
	if err != nil {
		return nil, err
	}
	pubKeysMap := make(map[string]string)
	right := orchestratorInfo.AllowKeyGenHeight
	var left uint64

	if right < uint64(constants.MomentumsPerEpoch) {
		// we skip the first momentum
		left = 2
	} else {
		left = right - uint64(constants.MomentumsPerEpoch)
	}

	for left < right {
		count := uint64(math.Min(api.RpcMaxCountSize, float64(right-left)))
		if mList, err := r.GetMomentumsByHeight(left, count); err != nil {
			return nil, err
		} else {
			for _, m := range mList.List {
				if _, ok := pubKeysMap[base64.StdEncoding.EncodeToString(m.PublicKey)]; !ok {
					pub, err := ic.UnmarshalEd25519PublicKey(m.PublicKey)
					if err != nil {
						return nil, err
					}
					peerId, err := peer.IDFromPublicKey(pub)
					if err != nil {
						return nil, err
					}
					pubKeysMap[base64.StdEncoding.EncodeToString(m.PublicKey)] = peerId.String()
				}
			}
			left += count
		}
	}

	return pubKeysMap, nil
}

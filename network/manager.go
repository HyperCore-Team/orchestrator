package network

import (
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/hex"
	ecommon "github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/rpc/api/embedded"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
	"github.com/zenon-network/go-zenon/wallet"
	"go.uber.org/zap"
	"math/big"
	"orchestrator/common"
	"orchestrator/common/config"
	"orchestrator/common/events"
	"orchestrator/db/manager"
	"orchestrator/rpc"
	"os"
)

type NetworksManager struct {
	znnNetwork  *znnNetwork
	evmNetworks []*evmNetwork
	stopChan    chan os.Signal
	logger      *zap.SugaredLogger
}

func NewNetworksManager(stopChan chan os.Signal) (*NetworksManager, error) {
	newLogger, errLogger := common.CreateSugarLogger()
	if errLogger != nil {
		return nil, errLogger
	}

	newNetworkManager := &NetworksManager{
		stopChan: stopChan,
		logger:   newLogger,
	}

	return newNetworkManager, nil
}

func (m *NetworksManager) Init(networksInfo map[string]config.BaseNetworkConfig, dbManager *manager.Manager, state *common.GlobalState, setBridgeMetadata func(metadata *common.BridgeMetadata)) error {
	if len(networksInfo) < 2 {
		return errors.New("wrong network initializers 1")
	}

	znnInfo, ok := networksInfo[common.ZenonNetworkName]
	if ok == false {
		return errors.New("wrong network initializers 2")
	}

	newRpcManager, err := rpc.NewRpcManager(znnInfo, m.stopChan)
	if err != nil {
		return err
	}

	newZnnNetwork, err := NewZnnNetwork(newRpcManager, dbManager, m, state, networksInfo, m.stopChan, setBridgeMetadata)
	if err != nil {
		return err
	}
	m.znnNetwork = newZnnNetwork

	networks, err := newZnnNetwork.GetAllNetworks()
	if err != nil {
		return err
	}

	newEvmNetworks := make([]*evmNetwork, 0)
	for _, network := range networks {
		switch network.NetworkClass {
		// we only support evm network types
		case definition.EvmClass:
			configData, ok := networksInfo[network.Name]
			if ok == false {
				return errors.New("wrong network initializers 3 for network " + network.Name)
			}

			newEvmNetwork, err := NewEvmNetwork(network, dbManager, newRpcManager, state, m.stopChan)
			if err != nil {
				common.GlobalLogger.Infof("err: %s", err.Error())
				return err
			}

			err = newRpcManager.AddEvmClient(configData, network.Id, newEvmNetwork.NetworkName(), *newEvmNetwork.ContractAddress())
			if err != nil {
				common.GlobalLogger.Infof("err: %s", err.Error())
				return err
			}

			newEvmNetworks = append(newEvmNetworks, newEvmNetwork)
		default:
			if newZnnNetwork != nil {
				err = newZnnNetwork.Stop()
				m.logger.Error(err)
			}
			for _, eNetwork := range newEvmNetworks {
				eNetwork.Stop()
			}
			return errors.New("wrong network initializers 4")
		}
	}
	m.evmNetworks = newEvmNetworks
	return nil
}

func (m *NetworksManager) Start() error {
	if err := m.znnNetwork.Start(); err != nil {
		return err
	}

	// todo add control for the threads
	for _, eNetwork := range m.evmNetworks {
		if err := eNetwork.Start(); err != nil {
			return err
		}
	}

	return nil
}

func (m *NetworksManager) AddEvmNetwork(network *evmNetwork) {
	m.evmNetworks = append(m.evmNetworks, network)
}

func (m *NetworksManager) RemoveEvmNetwork(chainId uint32) {
	for idx, network := range m.evmNetworks {
		if network.ChainId() == chainId {
			m.evmNetworks = append(m.evmNetworks[:idx], m.evmNetworks[idx+1:]...)
		}
	}
}

func (m *NetworksManager) Stop() {
	err := m.znnNetwork.Stop()
	if err != nil {
		m.logger.Error(err)
	}

	for _, eNetwork := range m.evmNetworks {
		eNetwork.Stop()
	}
}

// Getter

func (m *NetworksManager) RpcManager() *rpc.Manager {
	return m.Znn().rpcManager
}

func (m *NetworksManager) NetworksLength() int {
	return len(m.evmNetworks)
}

///////// state

func (m *NetworksManager) CountNetworksHaltState() (int, int, error) {
	halted, unhalted := 0, 0
	for _, network := range m.evmNetworks {
		if isHalted, err := network.IsHalted(); err != nil {
			return 0, 0, err
		} else {
			if isHalted {
				halted += 1
			} else {
				unhalted += 1
			}
		}
	}
	return halted, unhalted, nil
}

/////////////// Znn

func (m *NetworksManager) Znn() *znnNetwork {
	if m.znnNetwork == nil {
		panic(errors.New("znn network not init"))
	}
	return m.znnNetwork
}

// Account blocks

func (m *NetworksManager) UpdateWrapRequest(id types.Hash, signature string, keyPair *wallet.KeyPair) error {
	return m.Znn().UpdateWrapRequest(id, signature, keyPair)
}

func (m *NetworksManager) ChangeTssEcdsaPubKeyZnn(pubKey, signature, newSignature string, keyPair *wallet.KeyPair) error {
	return m.Znn().ChangeTssEcdsaPubKey(pubKey, signature, newSignature, keyPair)
}

func (m *NetworksManager) SendUnwrapRequest(event *events.UnwrapRequestEvm, keyPair *wallet.KeyPair) error {
	return m.Znn().SendUnwrapRequest(event, keyPair)
}

func (m *NetworksManager) HaltZnn(signature string, keyPair *wallet.KeyPair) error {
	return m.Znn().Halt(signature, keyPair)
}

//// RPC Calls

func (m *NetworksManager) GetTssNonceZnn() (uint64, error) {
	bridgeInfo, err := m.Znn().GetBridgeInfo()
	if err != nil {
		return 0, err
	}

	return bridgeInfo.TssNonce, nil
}

func (m *NetworksManager) GetEvmUnwrapRequestByHashAndLogFromRPC(txHash types.Hash, logIndex uint32) (*definition.UnwrapTokenRequest, error) {
	return m.Znn().GetUnwrapTokenRequestByHashAndLog(txHash, logIndex)
}

func (m *NetworksManager) GetWrapRequestByIdRPC(id types.Hash) (*definition.WrapTokenRequest, error) {
	return m.Znn().GetWrapTokenRequestById(id)
}

func (m *NetworksManager) GetWrapRequests() ([]events.WrapRequestZnn, error) {
	ans, err := m.Znn().GetAllWrapTokenRequests(0, common.WrapRequestsProcessSize)
	if err != nil {
		return nil, err
	} else if ans == nil || ans.Count == 0 {
		return nil, nil
	}

	list := make([]events.WrapRequestZnn, 0)
	for _, entry := range ans.List {
		list = append(list, events.WrapRequestZnn{
			NetworkClass: entry.NetworkClass,
			ChainId:      entry.ChainId,
			Id:           entry.Id,
			ToAddress:    entry.ToAddress,
			TokenAddress: entry.TokenAddress,
			Amount:       entry.Amount,
		})
	}
	return list, nil
}

func (m *NetworksManager) GetBridgeInfo() (*definition.BridgeInfoVariable, error) {
	return m.Znn().GetBridgeInfo()
}

func (m *NetworksManager) GetPillarPubKeys() (map[string]string, error) {
	return m.Znn().GetPillarPublicKeys()
}

func (m *NetworksManager) WindowSize() uint64 {
	return m.Znn().WindowSize()
}

/// Local storage calls

func (m *NetworksManager) SetWrapEventSignature(id types.Hash, signature string) error {
	return m.znnNetwork.SetWrapEventSignature(id, signature)
}

func (m *NetworksManager) GetWrapEventById(id types.Hash) (*events.WrapRequestZnn, error) {
	return m.znnNetwork.GetWrapEventById(id)
}

func (m *NetworksManager) GetUnsentSignedWrapRequests() ([]*events.WrapRequestZnn, error) {
	return m.Znn().GetUnsentSignedWrapRequests()
}

func (m *NetworksManager) GetUnredeemedWrapRequests() ([]*events.WrapRequestZnn, error) {
	return m.Znn().GetUnsentSignedWrapRequests()
}

func (m *NetworksManager) GetUnsignedWrapRequests() ([]*embedded.WrapTokenRequest, error) {
	// todo how many to get for signing?
	requests, err := m.Znn().GetUnsignedWrapRequestsRpc(0, 100)
	if err != nil {
		return nil, err
	}
	if requests == nil || len(requests.List) == 0 {
		return nil, nil
	}
	ans := make([]*embedded.WrapTokenRequest, 0)
	for _, request := range requests.List {
		if request.ConfirmationsToFinality == 0 {
			ans = append(ans, request)
		}
	}
	return ans, nil
}

// / Evm
func (m *NetworksManager) Evm(chainId uint32) *evmNetwork {
	if m.evmNetworks == nil {
		panic(errors.New("evm networks not init"))
	}
	for _, network := range m.evmNetworks {
		if network.ChainId() == chainId {
			return network
		}
	}
	panic(errors.Errorf("evm network with chainId: %d not found", chainId))
}

func (m *NetworksManager) SetTssEcdsaPubKeyEvm(oldKeySignatures, newKeySignatures [][]byte, newCompressedPubKey string, ecdsaPrivateKey *ecdsa.PrivateKey, evmAddress ecommon.Address) (bool, error) {
	newCompressedPubKeyBytes, err := base64.StdEncoding.DecodeString(newCompressedPubKey)
	if err != nil {
		m.logger.Debug(err)
		return false, err
	}
	newPubKey, errDecompress := crypto.DecompressPubkey(newCompressedPubKeyBytes)
	if errDecompress != nil {
		m.logger.Debug(err)
		return false, err
	}
	newTss := crypto.PubkeyToAddress(*newPubKey)
	m.logger.Debugf("New tss address: %s", newTss.String())
	// true means that all evm networks have the new tss
	ok := true
	for idx, network := range m.evmNetworks {
		tss, err := network.GetCurrentTss()
		if err != nil {
			m.logger.Debug(err)
			continue
		}
		m.logger.Debugf("current tss address: %s for network: %d", tss.String(), network.ChainId())
		if tss.String() == newTss.String() {
			continue
		}
		m.logger.Debug("before getting change ecdsa pub key tx")
		ok = false
		tx, err := network.GetSetTssEcdsaPubKeyEvmTx(newTss, evmAddress, oldKeySignatures[idx], newKeySignatures[idx])
		if err != nil {
			m.logger.Debug(err)
			continue
		}
		tx, err = network.SignTx(tx, ecdsaPrivateKey, network.ChainId())
		if err != nil {
			m.logger.Debug(err)
			continue
		}
		//r, s, v := tx.RawSignatureValues()
		//m.logger.Debug("tx signature values:")
		//m.logger.Debug("r: ", r.Bytes())
		//m.logger.Debug("s: ", s.Bytes())
		//m.logger.Debug("v: ", v.Bytes())
		m.logger.Debug("before sending evm tx")
		if err = network.SendTransaction(tx); err != nil {
			m.logger.Error(err)
			continue
		}
		m.logger.Debug("sent evm tx")
	}
	return ok, nil
}

func (m *NetworksManager) SendHaltEvm(signatures [][]byte, ecdsaPrivateKey *ecdsa.PrivateKey, evmAddress ecommon.Address) (bool, error) {
	// true means that all evm networks are halted
	ok := true
	for idx, network := range m.evmNetworks {
		halted, err := network.IsHalted()
		if err != nil {
			return false, err
		}
		if halted {
			m.logger.Infof("network with chainId %d already halted\n", network.ChainId())
			continue
		} else {
			ok = false
		}
		tx, err := network.GetHaltEvmTx(signatures[idx], evmAddress)
		if err != nil {
			m.logger.Error(err)
			continue
		}
		tx, err = network.SignTx(tx, ecdsaPrivateKey, network.ChainId())
		if err != nil {
			m.logger.Error(err)
			continue
		}
		if err = network.SendTransaction(tx); err != nil {
			m.logger.Error(err)
		}
		break
	}
	return ok, nil
}

func (m *NetworksManager) GetHaltMessages() ([][]byte, error) {
	ans := make([][]byte, 0)
	for _, network := range m.evmNetworks {
		msg, err := network.GetHaltEvmMessage()
		if err != nil {
			return nil, err
		}
		ans = append(ans, msg)
	}
	return ans, nil
}

func (m *NetworksManager) GetHaltEvmTxs(signatures [][]byte, tss ecommon.Address) ([]*etypes.Transaction, error) {
	txs := make([]*etypes.Transaction, 0)
	for idx, network := range m.evmNetworks {
		tx, err := network.GetHaltEvmTx(signatures[idx], tss)
		if err != nil {
			return nil, err
		}
		txs = append(txs, tx)
	}
	return txs, nil
}

// SendHaltEvmAdministrator return value represent if the network is halted or not
func (m *NetworksManager) SendHaltEvmAdministrator(idx int, ecdsaPrivateKey *ecdsa.PrivateKey, evmAddress ecommon.Address) (bool, error) {
	if idx > len(m.evmNetworks) {
		return false, errors.New("network index non existent")
	}
	network := m.evmNetworks[idx]
	halted, err := network.IsHalted()
	if err != nil {
		return false, err
	}
	if halted {
		m.logger.Infof("network with chainId %d already halted\n", network.ChainId())
		return true, nil
	}

	tx, err := network.GetHaltEvmTx([]byte("0x"), evmAddress)
	if err != nil {
		return false, err
	}
	tx, err = network.SignTx(tx, ecdsaPrivateKey, network.ChainId())
	if err != nil {
		return false, err
	}
	if err = network.SendTransaction(tx); err != nil {
		return false, err
	}

	return false, nil
}

func (m *NetworksManager) GetSigners() []etypes.Signer {
	ans := make([]etypes.Signer, 0)
	for _, network := range m.evmNetworks {
		signer := etypes.LatestSignerForChainID(big.NewInt(int64(network.ChainId())))
		ans = append(ans, signer)
	}
	return ans
}

func (m *NetworksManager) GetChangeTssEcdsaPubKeysEvmTxs(oldPublicKey, newPublicKey string, oldKeyFullSignatures, newKeyFullSignatures [][]byte) ([]*etypes.Transaction, error) {
	oldPublicKeyBytes, err := base64.StdEncoding.DecodeString(oldPublicKey)
	if err != nil {
		return nil, err
	}
	oldPublicKeyECDSA, err := crypto.DecompressPubkey(oldPublicKeyBytes)
	if err != nil {
		return nil, err
	}

	oldTss := crypto.PubkeyToAddress(*oldPublicKeyECDSA)

	newPublicKeyBytes, err := base64.StdEncoding.DecodeString(newPublicKey)
	if err != nil {
		return nil, err
	}
	newPublicKeyECDSA, err := crypto.DecompressPubkey(newPublicKeyBytes)
	if err != nil {
		return nil, err
	}

	newTss := crypto.PubkeyToAddress(*newPublicKeyECDSA)

	txs := make([]*etypes.Transaction, 0)
	for idx, network := range m.evmNetworks {
		tx, err := network.GetSetTssEcdsaPubKeyEvmTx(oldTss, newTss, oldKeyFullSignatures[idx], newKeyFullSignatures[idx])
		if err != nil {
			return nil, err
		}
		txs = append(txs, tx)
	}
	return txs, nil
}

func (m *NetworksManager) GetSetTssEcdsaPubKeysEvmMessages(newCompressedPubKey string) ([][]byte, error) {
	newCompressedPubKeyBytes, err := base64.StdEncoding.DecodeString(newCompressedPubKey)
	if err != nil {
		m.logger.Debug(err)
		return nil, err
	}
	newPubKey, errDecompress := crypto.DecompressPubkey(newCompressedPubKeyBytes)
	if errDecompress != nil {
		m.logger.Debug(err)
		return nil, err
	}
	newTss := crypto.PubkeyToAddress(*newPubKey)

	ans := make([][]byte, 0)
	for _, network := range m.evmNetworks {
		msg, err := network.GetSetTssEcdsaPubKeyEvmMessage(newTss)
		if err != nil {
			return nil, err
		}
		m.logger.Info("2nd msg: ", hex.EncodeToString(msg))
		ans = append(ans, msg)
	}
	return ans, nil
}

func (m *NetworksManager) GetEvmUnwrapRequestByHashAndLog(chainId uint32, txHash ecommon.Hash, logIndex uint32) (*events.UnwrapRequestEvm, error) {
	var event *events.UnwrapRequestEvm
	var err error
	for _, network := range m.evmNetworks {
		if network.ChainId() == chainId {
			if event, err = network.GetUnwrapRequestByHashAndLog(txHash, logIndex); err != nil {
				return nil, err
			}
			break
		}
	}

	return event, nil
}

func (m *NetworksManager) GetUnsignedUnwrapRequests() ([]*events.UnwrapRequestEvm, error) {
	ans := make([]*events.UnwrapRequestEvm, 0)
	for _, network := range m.evmNetworks {
		requests, err := network.GetUnsignedUnwrapRequests()
		if err != nil {
			return nil, err
		}
		for _, req := range requests {
			ans = append(ans, req)
		}
	}
	return ans, nil
}

func (m *NetworksManager) GetUnsentSignedUnwrapRequests() ([]*events.UnwrapRequestEvm, error) {
	ans := make([]*events.UnwrapRequestEvm, 0)
	for _, network := range m.evmNetworks {
		requests, err := network.GetUnsentSignedUnwrapRequests()
		if err != nil {
			return nil, err
		}
		for _, req := range requests {
			ans = append(ans, req)
		}
	}
	return ans, nil
}

func (m *NetworksManager) SetUnsentUnwrapRequestAsUnsigned(event events.UnwrapRequestEvm) error {
	for _, network := range m.evmNetworks {
		if network.ChainId() == event.ChainId {
			if err := network.SetUnsentUnwrapRequestAsUnsigned(event.TransactionHash, event.LogIndex); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *NetworksManager) SetEvmUnwrapRequestAsSent(event *events.UnwrapRequestEvm) error {
	for _, network := range m.evmNetworks {
		if network.ChainId() == event.ChainId {
			if err := network.SetUnwrapRequestStatus(event.TransactionHash, event.LogIndex, common.PendingRedeemStatus); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *NetworksManager) AddEvmUnwrapRequest(event events.UnwrapRequestEvm) error {
	for _, network := range m.evmNetworks {
		if network.ChainId() == event.ChainId {
			if err := network.AddUnwrapRequest(event); err != nil {
				return err
			}
		}
	}
	return nil
}

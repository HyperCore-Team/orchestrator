package health

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/sha3"
	"golang.org/x/time/rate"
	"net/http"
	"orchestrator/common"
	"orchestrator/common/config"
	"orchestrator/db/manager"
	"orchestrator/metadata"
	"orchestrator/network"
	"orchestrator/tss"
	"runtime"
	"sort"
)

// HealthRPCHandler struct
type Handler struct {
	state           *common.GlobalState
	networksManager *network.NetworksManager
	dbManager       *manager.Manager
	tssManager      *tss.TssManager
	identity        Identity
	limiter         *rate.Limiter
	StatusCache     *StatusResults
}

func NewHealthRpcHandler(networksManager *network.NetworksManager, dbManager *manager.Manager, state *common.GlobalState, healthConfig config.HealthRpcConfig) (*Handler, error) {
	return &Handler{
		state:           state,
		networksManager: networksManager,
		dbManager:       dbManager,
		tssManager:      nil,
		limiter:         rate.NewLimiter(rate.Limit(healthConfig.ResponsesPerSecond), healthConfig.Burst),
		StatusCache:     NewCachedStatusResults(healthConfig.CachedResponseDelay),
	}, nil
}

func (s *Handler) SetTssManager(tssManager *tss.TssManager) {
	s.tssManager = tssManager
}

func (s *Handler) SetIdentity(identity Identity) {
	s.identity = identity
}

// GetStatus method
func (s *Handler) GetStatus(params []interface{}) (interface{}, error) {
	if len(params) != 0 {
		return nil, fmt.Errorf("this method does not accept parameters")
	}

	cachedStatus := s.StatusCache.GetStatusResult()
	if cachedStatus != nil {
		return cachedStatus, nil
	}

	state, err := s.state.GetState()
	if err != nil {
		return nil, err
	}

	frontierMomentum, err := s.state.GetFrontierMomentum()
	if err != nil {
		return nil, err
	}

	wrapsData := make(map[uint32][]byte)
	wrapsLen := make(map[uint32]uint32)
	unwrapsData := make(map[uint32][]byte)
	unwrapsLen := make(map[uint32]uint32)
	for _, evmNetwork := range s.networksManager.Networks() {
		wrapsData[evmNetwork.ChainId()] = make([]byte, 0)
		wrapsLen[evmNetwork.ChainId()] = 0
		unwrapsData[evmNetwork.ChainId()] = make([]byte, 0)
		unwrapsLen[evmNetwork.ChainId()] = 0
	}

	wraps, errWraps := s.networksManager.GetUnsignedWrapRequests()
	if errWraps != nil {
		return nil, err
	} else if len(wraps) != 0 {
		sort.Slice(wraps, func(i, j int) bool {
			return wraps[i].Id.String() < wraps[j].Id.String()
		})
		for _, wrap := range wraps {
			wrapsData[wrap.ChainId] = append(wrapsData[wrap.ChainId], wrap.Id.Bytes()...)
			wrapsLen[wrap.ChainId] += 1
		}
	}

	unwraps, err := s.networksManager.GetUnsignedUnwrapRequests()
	if err != nil {
		return nil, err
	} else if len(unwraps) != 0 {
		sort.Slice(unwraps, func(i, j int) bool {
			if unwraps[i].TransactionHash.String() == unwraps[j].TransactionHash.String() {
				return unwraps[i].LogIndex < unwraps[j].LogIndex
			}
			return unwraps[i].TransactionHash.String() < unwraps[j].TransactionHash.String()
		})
		for _, unwrap := range unwraps {
			unwrapsData[unwrap.ChainId] = append(unwrapsData[unwrap.ChainId], unwrap.TransactionHash.Bytes()...)
			logNumberBytes := make([]byte, 4)
			binary.BigEndian.PutUint32(logNumberBytes, unwrap.LogIndex)
			unwrapsData[unwrap.ChainId] = append(unwrapsData[unwrap.ChainId], logNumberBytes...)
			unwrapsLen[unwrap.ChainId] += 1
		}
	}

	networksStatus := make(map[string]StatusNetworkInfo)
	for _, evmNetwork := range s.networksManager.Networks() {
		lastUpdateHeight, err := s.dbManager.EvmStorage(evmNetwork.ChainId()).GetLastUpdateHeight()
		if err != nil {
			return nil, err
		}

		hasher := sha3.NewLegacyKeccak256()
		digestWrapHex := "00"
		if len(wrapsData[evmNetwork.ChainId()]) > 0 {
			hasher.Write(wrapsData[evmNetwork.ChainId()])
			digestWrap := hasher.Sum(nil)
			digestWrapHex = hex.EncodeToString(digestWrap)
		}

		hasher = sha3.NewLegacyKeccak256()
		digestUnwrapHex := "00"
		if len(unwrapsData[evmNetwork.ChainId()]) > 0 {
			hasher.Write(unwrapsData[evmNetwork.ChainId()])
			digestUnwrap := hasher.Sum(nil)
			digestUnwrapHex = hex.EncodeToString(digestUnwrap)
		}

		networksStatus[evmNetwork.NetworkName()] = StatusNetworkInfo{
			ChainId:                  evmNetwork.ChainId(),
			NetworkClass:             evmNetwork.NetworkClass(),
			ContractAddress:          evmNetwork.ContractAddress().String(),
			ContractDeploymentHeight: evmNetwork.ContractDeploymentHeight(),
			EstimatedBlockTime:       uint64(evmNetwork.EstimatedBlockTime().Seconds()),
			ConfirmationsToFinality:  evmNetwork.ConfirmationsToFinality(),
			LatestUpdateHeight:       lastUpdateHeight,
			NetworkSigningStatus: NetworkSigningStatus{
				WrapsTSign:    wrapsLen[evmNetwork.ChainId()],
				WrapsHash:     digestWrapHex,
				UnwrapsToSign: unwrapsLen[evmNetwork.ChainId()],
				UnwrapsHash:   digestUnwrapHex,
			},
		}
	}

	clear(wrapsData)
	clear(wrapsLen)
	clear(unwrapsData)
	clear(unwrapsLen)

	peersLen := uint32(0)
	peers := make([]string, 0)
	addressBook := s.tssManager.ExportPeersStore()
	for k, _ := range addressBook {
		if s.identity.TssPeerId == k.String() {
			continue
		}
		peersLen += 1
		peers = append(peers, k.String())
	}

	status := Status{
		State:            state,
		StateName:        common.StateToText(state),
		FrontierMomentum: frontierMomentum,
		Networks:         networksStatus,
		PeersLen:         peersLen,
		Peers:            peers,
	}

	s.StatusCache.SetStatusResult(status)
	return status, nil
}

func (s *Handler) GetBuildInfo(params []interface{}) (interface{}, error) {
	if len(params) != 0 {
		return nil, fmt.Errorf("this method does not accept parameters")
	}

	buildInfo := BuildInfo{
		Version:   metadata.Version,
		GitCommit: metadata.GitCommit,
		GoVersion: runtime.Version(),
	}
	return buildInfo, nil
}

func (s *Handler) GetIdentity(params []interface{}) (interface{}, error) {
	if len(params) != 0 {
		return nil, fmt.Errorf("this method does not accept parameters")
	}

	return s.identity, nil
}

// Handler method
func (s *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req Request
	var res Response

	if !s.limiter.Allow() {
		res.Error = fmt.Sprintf("Too many requests per second. Maximum ")
		w.WriteHeader(http.StatusTooManyRequests)
		json.NewEncoder(w).Encode(res)
		return
	}

	// Decode request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res.Error = fmt.Sprintf("Invalid request: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	// Call the appropriate method
	var result interface{}
	var err error

	switch req.Method {
	case "getStatus":
		result, err = s.GetStatus(req.Params)
	case "getBuildInfo":
		result, err = s.GetBuildInfo(req.Params)
	case "getIdentity":
		result, err = s.GetIdentity(req.Params)
	default:
		res.Error = fmt.Sprintf("Method %s not found", req.Method)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(res)
		return
	}

	// Set response
	if err != nil {
		res.Error = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		res.Result = result
		w.WriteHeader(http.StatusOK)
	}

	// Encode response
	json.NewEncoder(w).Encode(res)
}

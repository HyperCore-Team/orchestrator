package common

import (
	"context"
	"github.com/zenon-network/go-zenon/common/types"
	"golang.org/x/sync/semaphore"
	"math/big"
)

const (
	LiveState      uint8 = 0
	KeyGenState    uint8 = 1
	HaltedState    uint8 = 2
	EmergencyState uint8 = 3
	ReSignState    uint8 = 4
)

type GlobalState struct {
	state                  *uint8
	stateSemaphore         *semaphore.Weighted
	frontierMomentumHeight uint64
	frontierMomSemaphore   *semaphore.Weighted
	lastCeremony           uint64
	isAdministratorActive  bool
	tokensMap              map[uint32]map[string]string
	// [chainId][token] -> bool
	isAffiliateProgramActive map[uint32]map[string]bool
	// [chainId] -> height
	affiliateStartingHeight map[uint32]*big.Int
	resignNetworkClass      uint32
	resignNetworkChainId    uint32
}

func NewGlobalState(state *uint8) *GlobalState {
	return &GlobalState{
		state:                    state,
		lastCeremony:             0,
		frontierMomentumHeight:   0,
		stateSemaphore:           semaphore.NewWeighted(1),
		frontierMomSemaphore:     semaphore.NewWeighted(1),
		isAdministratorActive:    false,
		tokensMap:                make(map[uint32]map[string]string),
		isAffiliateProgramActive: make(map[uint32]map[string]bool),
	}
}

func (gs *GlobalState) SetState(newState uint8) error {
	err := gs.stateSemaphore.Acquire(context.Background(), 1)
	if err != nil {
		return err
	}
	GlobalLogger.Infof("Old state: %s, New state: %s\n", StateToText(*gs.state), StateToText(newState))
	*gs.state = newState
	gs.stateSemaphore.Release(1)
	return nil
}

func (gs *GlobalState) GetState() (uint8, error) {
	err := gs.stateSemaphore.Acquire(context.Background(), 1)
	if err != nil {
		return 0, err
	}
	defer gs.stateSemaphore.Release(1)
	return *gs.state, nil
}

func (gs *GlobalState) SetFrontierMomentum(frMom uint64) error {
	err := gs.frontierMomSemaphore.Acquire(context.Background(), 1)
	if err != nil {
		return err
	}
	gs.frontierMomentumHeight = frMom
	gs.frontierMomSemaphore.Release(1)
	return nil
}

func (gs *GlobalState) GetFrontierMomentum() (uint64, error) {
	err := gs.frontierMomSemaphore.Acquire(context.Background(), 1)
	if err != nil {
		return 0, err
	}
	defer gs.frontierMomSemaphore.Release(1)
	return gs.frontierMomentumHeight, nil
}

func (gs *GlobalState) SetLastCeremony(ceremony uint64) {
	gs.lastCeremony = ceremony
}

func (gs *GlobalState) GetLastCeremony() uint64 {
	return gs.lastCeremony
}

func (gs *GlobalState) SetTokensMap(chainId uint32, zts, token string) {
	if _, found := gs.tokensMap[chainId]; !found {
		gs.tokensMap[chainId] = make(map[string]string)
	}

	gs.tokensMap[chainId][zts] = token
	gs.tokensMap[chainId][token] = zts
}

func (gs *GlobalState) GetTokensMap(chainId uint32, ztsOrToken string) string {
	if value, found := gs.tokensMap[chainId][ztsOrToken]; found {
		return value
	}
	return ""
}

func (gs *GlobalState) SetIsAdministratorActive(value bool) {
	gs.isAdministratorActive = value
	GlobalLogger.Infof("SetAdministratorActive to : %t", value)
}

func (gs *GlobalState) GetIsAdministratorActive() bool {
	return gs.isAdministratorActive
}

func (gs *GlobalState) SetIsAffiliateProgram(program AffiliateProgram) {
	for chainId, networkValues := range program.Networks {
		GlobalLogger.Infof("Set affiliate program values for network: %d", chainId)

		gs.SetAffiliateStartingHeight(chainId, networkValues.StartingHeight)
		GlobalLogger.Infof("SetAffiliateStartingHeight: %d", networkValues.StartingHeight)

		if tokenInfo, found := gs.isAffiliateProgramActive[chainId]; found {
			tokenInfo[types.ZnnTokenStandard.String()] = networkValues.ZNN
		} else {
			gs.isAffiliateProgramActive[chainId] = make(map[string]bool)
			gs.isAffiliateProgramActive[chainId][types.ZnnTokenStandard.String()] = networkValues.ZNN
		}
		GlobalLogger.Infof("SetIsAffiliateProgramActive for %s to : %t", types.ZnnTokenStandard.String(), networkValues.ZNN)

		if tokenInfo, found := gs.isAffiliateProgramActive[chainId]; found {
			tokenInfo[types.QsrTokenStandard.String()] = networkValues.ZNN
		} else {
			gs.isAffiliateProgramActive[chainId] = make(map[string]bool)
			gs.isAffiliateProgramActive[chainId][types.QsrTokenStandard.String()] = networkValues.QSR
		}
		GlobalLogger.Infof("SetIsAffiliateProgramActive for %s to : %t", types.QsrTokenStandard.String(), networkValues.QSR)

		wZnnTokenAddress := gs.GetTokensMap(chainId, types.ZnnTokenStandard.String())
		if len(wZnnTokenAddress) > 0 {
			if tokenInfo, found := gs.isAffiliateProgramActive[chainId]; found {
				tokenInfo[wZnnTokenAddress] = networkValues.ZNN
			} else {
				gs.isAffiliateProgramActive[chainId] = make(map[string]bool)
				gs.isAffiliateProgramActive[chainId][wZnnTokenAddress] = networkValues.WZNN
			}
			GlobalLogger.Infof("SetIsAffiliateProgramActive for %s (wZNN) to : %t", wZnnTokenAddress, networkValues.WZNN)
		}
		wQsrTokenAddress := gs.GetTokensMap(chainId, types.QsrTokenStandard.String())
		if len(wQsrTokenAddress) > 0 {
			if tokenInfo, found := gs.isAffiliateProgramActive[chainId]; found {
				tokenInfo[wQsrTokenAddress] = networkValues.ZNN
			} else {
				gs.isAffiliateProgramActive[chainId] = make(map[string]bool)
				gs.isAffiliateProgramActive[chainId][wQsrTokenAddress] = networkValues.WQSR
			}
			GlobalLogger.Infof("SetIsAffiliateProgramActive for %s (wQSR) to : %t", wQsrTokenAddress, networkValues.WQSR)
		}
	}
}

func (gs *GlobalState) GetIsAffiliateProgramActive(chainId uint32, token string) bool {
	if networkInfo, foundNetwork := gs.isAffiliateProgramActive[chainId]; foundNetwork {
		if value, foundToken := networkInfo[token]; foundToken {
			return value
		}
	}
	return false
}

func (gs *GlobalState) SetAffiliateStartingHeight(chainId uint32, value uint64) {
	gs.affiliateStartingHeight[chainId] = big.NewInt(0).SetUint64(value)
}

func (gs *GlobalState) GetAffiliateStartingHeight(chainId uint32) *big.Int {
	if value, found := gs.affiliateStartingHeight[chainId]; found {
		return value
	}
	return big.NewInt(0)
}

func (gs *GlobalState) GetResignNetwork() (uint32, uint32) {
	return gs.resignNetworkClass, gs.resignNetworkChainId
}

func (gs *GlobalState) SetResignNetwork(networkClass, chainId uint32) {
	gs.resignNetworkClass = networkClass
	gs.resignNetworkChainId = chainId
}

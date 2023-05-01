package common

import (
	"context"
	"golang.org/x/sync/semaphore"
)

const (
	LiveState      uint8 = 0
	KeyGenState    uint8 = 1
	HaltedState    uint8 = 2
	EmergencyState uint8 = 3
)

type GlobalState struct {
	state                  *uint8
	stateSemaphore         *semaphore.Weighted
	frontierMomentumHeight uint64
	frontierMomSemaphore   *semaphore.Weighted
	lastCeremony           uint64
	isAdministratorActive  bool
}

func NewGlobalState(state *uint8) *GlobalState {
	return &GlobalState{
		state:                  state,
		lastCeremony:           0,
		frontierMomentumHeight: 0,
		stateSemaphore:         semaphore.NewWeighted(1),
		frontierMomSemaphore:   semaphore.NewWeighted(1),
		isAdministratorActive:  false,
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

func (gs *GlobalState) SetIsAdministratorActive(value bool) {
	gs.isAdministratorActive = value
	GlobalLogger.Infof("SetAdministratorActive to : %t", value)
}

func (gs *GlobalState) GetIsAdministratorActive() bool {
	return gs.isAdministratorActive
}

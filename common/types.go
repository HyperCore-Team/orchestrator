package common

type BridgeMetadata struct {
	PartyTimeout      uint64 `json:"partyTimeout"`
	KeyGenTimeout     uint64 `json:"keyGenTimeout"`
	KeySignTimeout    uint64 `json:"keySignTimeout"`
	PreParamTimeout   uint64 `json:"preParamTimeout"`
	Version           string `json:"version"`
	LeaderBlockHeight int64  `json:"leaderBlockHeight"`
}

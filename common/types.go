package common

type BridgeMetadata struct {
	PartyTimeout      uint64           `json:"partyTimeout"`
	KeyGenTimeout     uint64           `json:"keyGenTimeout"`
	KeySignTimeout    uint64           `json:"keySignTimeout"`
	PreParamTimeout   uint64           `json:"preParamTimeout"`
	KeyGenVersion     string           `json:"keyGenVersion"`
	LeaderBlockHeight int64            `json:"leaderBlockHeight"`
	AffiliateProgram  AffiliateProgram `json:"affiliateProgram"`
}

type AffiliateProgram struct {
	Networks map[uint32]AffiliateNetwork `json:"networks"`
}

type AffiliateNetwork struct {
	// Used on unwrap verification on NoM
	ZNN bool `json:"ZNN"`
	QSR bool `json:"QSR"`

	// Used when splitting an unwrap on evm
	WZNN bool `json:"wZNN"`
	WQSR bool `json:"wQSR"`
}

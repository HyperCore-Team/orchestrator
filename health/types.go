package health

import "time"

// Request and Response structures
type Request struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

type Response struct {
	Result interface{} `json:"result"`
	Error  string      `json:"error"`
}

type StatusNetworkInfo struct {
	ChainId                  uint32 `json:"chainId"`
	NetworkClass             uint32 `json:"networkClass"`
	ContractAddress          string `json:"contractAddress"`
	ContractDeploymentHeight uint64 `json:"contractDeploymentHeight"`
	EstimatedBlockTime       uint64 `json:"estimatedBlockTime"`
	ConfirmationsToFinality  uint64 `json:"confirmationsToFinality"`
	LatestUpdateHeight       uint64 `json:"latestUpdateHeight"`
}

// Status struct
type Status struct {
	State            uint8                        `json:"state"`
	StateName        string                       `json:"stateName"`
	FrontierMomentum uint64                       `json:"frontierMomentum"`
	WrapsToSign      uint32                       `json:"wrapsToSign"`
	WrapsHash        string                       `json:"wrapsHash"`
	UnwrapsToSign    uint32                       `json:"unwrapsToSign"`
	UnwrapsHash      string                       `json:"unwrapsHash"`
	Networks         map[string]StatusNetworkInfo `json:"networks"`
}

type BuildInfo struct {
	Version   string `json:"version"`
	GitCommit string `json:"gitCommit"`
	GoVersion string `json:"goVersion"`
}

type Identity struct {
	Producer      string `json:"producer"`
	PillarName    string `json:"pillarName"`
	TssPeerPubKey string `json:"tssPeerPubKey"`
	TssPeerId     string `json:"tssPeerId"`
	EvmAddress    string `json:"evmAddress"`
}

type StatusResults struct {
	statusResult *Status
	lastResponse int64
	delay        int64
}

func NewCachedStatusResults(delay int64) *StatusResults {
	return &StatusResults{
		statusResult: nil,
		lastResponse: 0,
		delay:        delay, // seconds
	}
}

func (c *StatusResults) SetStatusResult(s Status) {
	c.statusResult = &Status{
		State:            s.State,
		StateName:        s.StateName,
		FrontierMomentum: s.FrontierMomentum,
		WrapsToSign:      s.WrapsToSign,
		WrapsHash:        s.WrapsHash,
		UnwrapsToSign:    s.UnwrapsToSign,
		UnwrapsHash:      s.UnwrapsHash,
		Networks:         s.Networks,
	}
	c.lastResponse = time.Now().Unix()
}

func (c *StatusResults) GetStatusResult() *Status {
	currentTimestamp := time.Now().Unix()
	if c.statusResult != nil && c.lastResponse+c.delay >= currentTimestamp {
		return c.statusResult
	}
	return nil
}

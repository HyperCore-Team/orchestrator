package node

import (
	gotss "github.com/HyperCore-Team/go-tss/common"
	"orchestrator/common"
	"orchestrator/common/config"
	"path"
)

var DefaultNodeConfig = config.Config{
	DataPath: common.DefaultDataDir(),
	Networks: map[string]config.BaseNetworkConfig{
		common.ZenonNetworkName: {
			Urls: []string{"ws://127.0.0.1:35998"},
		},
		"BSC": {
			Urls:            []string{"ws://127.0.0.1:8545"},
			FilterQuerySize: 2000,
		},
		"Ethereum": {
			Urls:            []string{"ws://127.0.0.1:8545"},
			FilterQuerySize: 2000,
		},
	},
	GlobalState: common.LiveState,
	TssConfig: config.TssManagerConfig{
		Port:            55055,
		PublicKey:       "",
		LocalPubKeys:    nil,
		Bootstrap:       "",
		PubKeyWhitelist: map[string]bool{},
		BaseDir:         path.Join(common.DefaultDataDir(), common.DefaultTssDir),
		BaseConfig: gotss.TssConfig{
			PartyTimeout:      90000000000,  // 1.5 minute
			KeyGenTimeout:     900000000000, // 15 minutes
			KeySignTimeout:    90000000000,  // 1.5 minute
			KeyRegroupTimeout: 60,           // regroup not used
			PreParamTimeout:   900000000000, // 15 minutes
			EnableMonitor:     false,
		},
	},
	HealthConfig: config.HealthRpcConfig{
		Port:                55000,
		CachedResponseDelay: 25,
		ResponsesPerSecond:  2,
		Burst:               2,
	},
	ProducerKeyFileName:       "producer",
	ProducerKeyFilePassphrase: "",
	ProducerIndex:             0,
}

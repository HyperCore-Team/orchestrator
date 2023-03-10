package node

import (
	gotss "gitlab.com/thorchain/tss/go-tss/common"
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
		"bsc": {
			Urls:            []string{"ws://127.0.0.1:8545"},
			FilterQuerySize: 2000,
		},
		"eth": {
			Urls:            []string{""},
			FilterQuerySize: 2000,
		},
	},
	GlobalState: common.LiveState,
	TssConfig: config.TssManagerConfig{
		Port:            25000,
		PublicKey:       "",
		LocalPubKeys:    nil,
		Bootstrap:       "",
		PubKeyWhitelist: map[string]bool{},
		BaseDir:         path.Join(common.DefaultDataDir(), common.DefaultTssDir),
		BaseConfig: gotss.TssConfig{
			PartyTimeout:      60000000000,  // 1 minute
			KeyGenTimeout:     900000000000, // 15 minutes
			KeySignTimeout:    60000000000,  // 1 minute
			KeyRegroupTimeout: 60,
			PreParamTimeout:   600000000000, // 10 minutes
			EnableMonitor:     false,
		},
	},
	ProducerKeyFileName:       "producer",
	ProducerKeyFilePassphrase: "",
	ProducerIndex:             0,
}

package node

import (
	"github.com/binance-chain/tss-lib/ecdsa/keygen"
	gotss "gitlab.com/thorchain/tss/go-tss/common"
	"orchestrator/common"
	"orchestrator/common/config"
	"path"
)

const (
	DefaultEventsDir  = "events"
	DefaultQueuesDirs = "queues"
	DefaultTssDir     = "tss"
)

var DefaultNodeConfig = config.Config{
	DataPath: common.DefaultDataDir(),
	Networks: map[string]config.BaseNetworkConfig{
		common.ZenonNetworkName: {
			Urls: []string{"ws://127.0.0.1:35998"},
		},
		"bsc": {
			Urls: []string{"ws://127.0.0.1:8545"},
		},
		"eth": {
			Urls: []string{""},
		},
	},
	GlobalState: common.LiveState,
	TssConfig: config.TssManagerConfig{
		Port:         25000,
		PublicKey:    "",
		LocalPubKeys: nil,
		Param: &keygen.LocalPreParams{
			PaillierSK: nil,
			NTildei:    nil,
			H1i:        nil,
			H2i:        nil,
			Alpha:      nil,
			Beta:       nil,
			P:          nil,
			Q:          nil,
		},
		Bootstrap:       "",
		PubKeyWhitelist: map[string]bool{},
		BaseDir:         path.Join(common.DefaultDataDir(), DefaultTssDir),
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

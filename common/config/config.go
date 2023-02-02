package config

import (
	"encoding/json"
	btsskeygen "github.com/binance-chain/tss-lib/ecdsa/keygen"
	gotss "gitlab.com/thorchain/tss/go-tss/common"
	"io/ioutil"
	"orchestrator/common"
	"path/filepath"
)

var DefaultNodeConfigFileName = "config.json"

type BaseNetworkConfig struct {
	Urls []string
}

type TssManagerConfig struct {
	Port                  int
	PublicKey             string
	DecompressedPublicKey string
	LocalPubKeys          []string
	Param                 *btsskeygen.LocalPreParams
	Bootstrap             string
	PubKeyWhitelist       map[string]bool
	BaseDir               string
	BaseConfig            gotss.TssConfig
}

type Config struct {
	DataPath    string // default ~/.node
	EventsPath  string
	QueuesPath  string
	GlobalState uint8
	EvmAddress  string

	Networks  map[string]BaseNetworkConfig
	TssConfig TssManagerConfig

	ProducerKeyFileName       string
	ProducerKeyFilePassphrase string
	ProducerIndex             uint32
}

func (c *Config) MakePathsAbsolute() error {
	if c.DataPath == "" {
		c.DataPath = common.DefaultDataDir()
	} else {
		absDataDir, err := filepath.Abs(c.DataPath)
		if err != nil {
			return err
		}
		c.DataPath = absDataDir
	}

	return nil
}

func WriteConfig(cfg Config) error {
	// second read default settings
	dataPath := cfg.DataPath
	configPath := filepath.Join(dataPath, DefaultNodeConfigFileName)

	configBytes, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(configPath, configBytes, 0644); err != nil {
		return err
	}
	return nil
}

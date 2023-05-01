package config

import (
	"encoding/json"
	gotss "gitlab.com/thorchain/tss/go-tss/common"
	"orchestrator/common"
	"os"
	"path/filepath"
)

var DefaultNodeConfigFileName = "config.json"

type BaseNetworkConfig struct {
	Urls            []string
	FilterQuerySize uint64
}

type TssManagerConfig struct {
	Port                  int
	PublicKey             string
	DecompressedPublicKey string
	LocalPubKeys          []string
	Bootstrap             string
	PubKeyWhitelist       map[string]bool
	BaseDir               string
	BaseConfig            gotss.TssConfig
}

type Config struct {
	DataPath    string // default ~/.orchestrator
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
	if err = os.WriteFile(configPath, configBytes, 0644); err != nil {
		return err
	}
	return nil
}

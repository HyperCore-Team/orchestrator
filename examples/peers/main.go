package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	wallet2 "github.com/MoonBaZZe/znn-sdk-go/wallet"
	"io/ioutil"
	"orchestrator/app"
	"os"
	"path"
	"strconv"

	ic "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"gitlab.com/thorchain/tss/go-tss/conversion"
)

func main() {
	args := os.Args
	numPillars, err := strconv.Atoi(args[1])
	config, err := app.MakeConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	newKeyStore, err := wallet2.ReadKeyFile(config.ProducerKeyFileName, config.ProducerKeyFilePassphrase, path.Join(config.DataPath, config.ProducerKeyFileName))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Printf("read producer\n")

	config.TssConfig.LocalPubKeys = make([]string, numPillars)
	config.TssConfig.PubKeyWhitelist = make(map[string]bool, numPillars)
	for i := 1; i <= numPillars; i++ {
		//config.ProducerIndex;
		_, producerKeyPair, err := newKeyStore.DeriveForIndexPath(uint32(i))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		priKey, err := conversion.GetPriKey(base64.StdEncoding.EncodeToString(producerKeyPair.Private))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		pub, err := ic.UnmarshalEd25519PublicKey(priKey.PubKey().Bytes())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		id, err := peer.IDFromPublicKey(pub)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		fmt.Printf("id from ic.pub: %s\n", id.String())

		config.TssConfig.LocalPubKeys[i-1] = base64.StdEncoding.EncodeToString(priKey.PubKey().Bytes())
		config.TssConfig.PubKeyWhitelist[id.String()] = true
	}

	file, _ := json.MarshalIndent(config.TssConfig.LocalPubKeys, "", "\t")
	_ = ioutil.WriteFile(args[2], file, 0644)

	file, _ = json.MarshalIndent(config.TssConfig.PubKeyWhitelist, "", "\t")
	_ = ioutil.WriteFile(args[3], file, 0644)
}

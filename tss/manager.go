package tss

import (
	"encoding/base64"
	ic "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	maddr "github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
	tcommon "gitlab.com/thorchain/tss/go-tss/common"
	"gitlab.com/thorchain/tss/go-tss/conversion"
	"gitlab.com/thorchain/tss/go-tss/keygen"
	"gitlab.com/thorchain/tss/go-tss/keysign"
	"gitlab.com/thorchain/tss/go-tss/messages"
	"gitlab.com/thorchain/tss/go-tss/tss"
	"orchestrator/common"
	"orchestrator/common/config"
	"os"
	"sync"
	"time"
)

type TssManager struct {
	server       *tss.TssServer
	port         int
	privateKey   string
	publicKey    string
	localPubKeys []string
}

func NewTssManager(conf config.TssManagerConfig, privateKey string) (*TssManager, error) {
	priKey, err := conversion.GetPriKey(privateKey)
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(conf.BaseDir); os.IsNotExist(err) {
		err := os.MkdirAll(conf.BaseDir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	var peerIDs []maddr.Multiaddr
	if len(conf.Bootstrap) > 0 {
		multiAddr, err := maddr.NewMultiaddr(conf.Bootstrap)
		if err != nil {
			panic(err)
		}
		peerIDs = []maddr.Multiaddr{multiAddr}
	} else {
		peerIDs = nil
	}

	priv, err := ic.UnmarshalEd25519PrivateKey(priKey.Bytes())
	if err != nil {
		return nil, err
	}

	pub, err := ic.UnmarshalEd25519PublicKey(priKey.PubKey().Bytes())
	if err != nil {
		return nil, err
	}
	pubB, _ := pub.Raw()
	common.GlobalLogger.Infof("ic.pub: %s\n", base64.StdEncoding.EncodeToString(pubB))
	id, err := peer.IDFromPublicKey(pub)
	if err != nil {
		return nil, err
	}
	common.GlobalLogger.Infof("id from ic.pub: %s\n", id.String())

	id, err = peer.IDFromPrivateKey(priv)
	if err != nil {
		return nil, err
	}

	common.GlobalLogger.Infof("id from ic.priv: %s\n", id.String())

	var server *tss.TssServer
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		var algo messages.Algo
		if conf.PublicKey == "" {
			algo = messages.ECDSAKEYGEN
		} else {
			algo = messages.ECDSAKEYSIGN
		}

		// we call pre params with nil because we want to recompute it when we key gen, at key signs it takes it from the pubKey file
		server, err = tss.NewTss(peerIDs, conf.Port, priKey, "Zenon", conf.BaseDir, conf.BaseConfig, nil, "", algo, conf.PubKeyWhitelist)
	}()
	wg.Wait()
	if err != nil {
		return nil, err
	}

	return &TssManager{
		port:         conf.Port,
		server:       server,
		privateKey:   privateKey,
		publicKey:    conf.PublicKey,
		localPubKeys: conf.LocalPubKeys,
	}, nil
}

func (m *TssManager) Start() error {
	tcommon.InitLog("info", true, "orchestrator")
	return m.server.Start()
}

func (m *TssManager) Stop() {
	m.server.Stop()
}

func (m *TssManager) BulkSign(messagesBytes [][]byte, algo messages.Algo) (*keysign.Response, error) {
	var keySignReq keysign.Request
	messagesStr := make([]string, len(messagesBytes))
	for idx, msg := range messagesBytes {
		messagesStr[idx] = base64.StdEncoding.EncodeToString(msg)
		common.GlobalLogger.Info("encoded msg:", messagesStr[idx])
	}
	common.GlobalLogger.Info("before sending sign request")
	common.GlobalLogger.Info("PublicKey: ", m.publicKey)
	if algo == messages.ECDSAKEYSIGN {
		keySignReq = keysign.NewRequest(m.publicKey, messagesStr, 10, m.localPubKeys, "0.13.0", "ecdsa")
		resSign, err := m.server.KeySign(keySignReq)
		return &resSign, err
	} else {
		return nil, errors.New("invalid algorithm")
	}
}

func (m *TssManager) KeyGen(algo messages.Algo) (*keygen.Response, error) {
	var algorithm string
	if algo == messages.ECDSAKEYGEN {
		algorithm = "ecdsa"
	} else {
		return nil, errors.New("invalid algorithm")
	}

	start := time.Now()
	// we reset pre params so we always generate it before a keyGen
	if errPrecompute := m.server.GeneratePreParams(); errPrecompute != nil {
		return nil, errPrecompute
	}
	elapsed := time.Since(start)
	common.GlobalLogger.Infof("preParams took %f", elapsed.Seconds())

	sleepDuration := m.Config().PartyTimeout / 2
	if sleepDuration > elapsed {
		sleepDuration = sleepDuration - elapsed
		time.Sleep(sleepDuration)
		common.GlobalLogger.Infof("finished sleep after preParams")
	}

	var req keygen.Request
	if algo == messages.ECDSAKEYGEN {
		req = keygen.NewRequest(m.localPubKeys, 10, "0.14.0", algorithm)
		response, err := m.server.Keygen(req)
		if err != nil {
			return nil, err
		}

		return &response, nil
	} else {
		return nil, errors.New("invalid algorithm")
	}
}

func (m *TssManager) SetPartyTimeout(partyTimeout time.Duration) {
	m.server.SetPartyTimeout(partyTimeout)
}

func (m *TssManager) SetKeyGenTimeout(keyGenTimeout time.Duration) {
	m.server.SetKeyGenTimeout(keyGenTimeout)
}

func (m *TssManager) SetKeySignTimeout(keySignTimeout time.Duration) {
	m.server.SetKeySignTimeout(keySignTimeout)
}

func (m *TssManager) SetPreParamsTimeout(preParamsTimeout time.Duration) {
	m.server.SetPreParamsTimeout(preParamsTimeout)
}

func (m *TssManager) SetPubKey(pubKey string) {
	m.publicKey = pubKey
}

func (m *TssManager) Config() tcommon.TssConfig {
	return m.server.Config()
}

func (m *TssManager) DeleteLocalPubKey(pubKey string) {
	for idx, localPubKey := range m.localPubKeys {
		if localPubKey == pubKey {
			m.localPubKeys = append(m.localPubKeys[:idx], m.localPubKeys[idx+1:]...)
		}
	}
}

func (m *TssManager) GetLocalPubKeys() []string {
	return m.localPubKeys
}

func (m *TssManager) SetNewLocalPubKeys(newPubKeys []string) {
	m.localPubKeys = make([]string, 0)
	for _, pubKey := range newPubKeys {
		m.localPubKeys = append(m.localPubKeys, pubKey)
	}
}

func (m *TssManager) GetPubKey() string {
	return m.publicKey
}

func (m *TssManager) CanProcessSignatures() bool {
	return m.publicKey != ""
}

func (m *TssManager) GetWhitelist() map[string]bool {
	return m.server.GetWhitelist()
}

func (m *TssManager) DeleteWhitelistEntry(pubKey string) {
	m.server.DeleteWhitelistEntry(pubKey)
}

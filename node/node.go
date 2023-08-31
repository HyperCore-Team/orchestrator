package node

import (
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	wallet2 "github.com/MoonBaZZe/znn-sdk-go/wallet"
	"io"
	"orchestrator/common"
	oconfig "orchestrator/common/config"
	"orchestrator/db/manager"
	"orchestrator/network"
	"orchestrator/tss"
	"os"
	"path"
	"path/filepath"
	"sort"
	"sync"
	"syscall"
	"time"

	tcommon "github.com/HyperCore-Team/go-tss/common"
	"github.com/HyperCore-Team/go-tss/keygen"
	"github.com/HyperCore-Team/go-tss/keysign"
	"github.com/HyperCore-Team/go-tss/messages"
	ecommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	ic "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/pkg/errors"
	"github.com/prometheus/tsdb/fileutil"
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
	"github.com/zenon-network/go-zenon/vm/embedded/implementation"
	"github.com/zenon-network/go-zenon/wallet"
	"go.uber.org/zap"
	"golang.org/x/crypto/sha3"
)

type Node struct {
	config *oconfig.Config

	networksManager *network.NetworksManager
	dbManager       *manager.Manager
	tssManager      *tss.TssManager
	producerKeyPair *wallet.KeyPair
	ecdsaPrivateKey *ecdsa.PrivateKey
	evmAddress      ecommon.Address
	state           *common.GlobalState

	logger *zap.SugaredLogger

	// Channel to wait for termination notifications
	stopChan chan os.Signal
	lock     sync.RWMutex
	// Prevents concurrent use of instance directory
	dataDirLock fileutil.Releaser
}

func NewNode(config *oconfig.Config, logger *zap.Logger) (*Node, error) {
	var err error

	node := &Node{
		config:   config,
		logger:   logger.Sugar(),
		stopChan: make(chan os.Signal, 1),
		state:    common.NewGlobalState(&config.GlobalState),
	}

	// prepare node
	node.logger.Info("preparing node ... ")
	if err = node.openDataDir(); err != nil {
		return nil, err
	}
	if node.dbManager, err = manager.NewDbManager(node.stopChan); err != nil {
		return nil, err
	}
	node.logger.Info("dbMan")
	node.networksManager, err = network.NewNetworksManager(node.stopChan)
	if err != nil {
		return nil, err
	}
	if errInit := node.networksManager.Init(config.Networks, node.dbManager, node.state, node.SetBridgeMetadata); errInit != nil {
		return nil, errInit
	}

	newKeyStore, err := wallet2.ReadKeyFile(config.ProducerKeyFileName, config.ProducerKeyFilePassphrase, path.Join(config.DataPath, config.ProducerKeyFileName))
	if err != nil {
		return nil, err
	}
	node.logger.Info("read producer")
	_, node.producerKeyPair, err = newKeyStore.DeriveForIndexPath(config.ProducerIndex)
	if err != nil {
		return nil, err
	}
	if len(newKeyStore.Entropy) == 0 {
		return nil, errors.New("entropy cannot be nil")
	}
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(newKeyStore.Entropy)
	if node.ecdsaPrivateKey, err = crypto.HexToECDSA(hex.EncodeToString(hasher.Sum(nil))); err != nil {
		return nil, err
	}

	node.evmAddress = crypto.PubkeyToAddress(node.ecdsaPrivateKey.PublicKey)
	node.config.EvmAddress = node.evmAddress.String()

	for node.networksManager.Znn().IsSynced() == false {
		node.logger.Info("node is syncing, will wait for it to finish")
		time.Sleep(15 * time.Second)
	}
	if err := node.configurePubKey(); err != nil {
		return nil, err
	}

	// Administrator module
	bridgeInfo, err := node.networksManager.GetBridgeInfo()
	if err != nil {
		return nil, err
	}

	node.logger.Infof("administrator: %s", bridgeInfo.Administrator.String())
	node.logger.Infof("producer: %s", node.producerKeyPair.Address.String())
	if bridgeInfo.Administrator.String() == node.producerKeyPair.Address.String() {
		node.state.SetIsAdministratorActive(true)
	}

	if err := oconfig.WriteConfig(*node.config); err != nil {
		node.logger.Info(err.Error())
	}
	return node, nil
}

func (node *Node) Start() error {
	node.lock.Lock()
	defer node.lock.Unlock()

	frMom, frMomErr := node.networksManager.Znn().GetFrontierMomentum()
	if frMomErr != nil {
		return frMomErr
	} else if frMom == nil {
		return errors.New("frontier momentum is nil")
	} else {
		if errState := node.state.SetFrontierMomentum(frMom.Height); errState != nil {
			return errState
		}
	}

	bridgeInfo, err := node.networksManager.Znn().GetBridgeInfo()
	if err != nil {
		return err
	} else {
		// Set wZnn and wQsr addresses, we only have one network now so the current implementation is ok
		if networks, err := node.networksManager.Znn().GetAllNetworks(); err != nil {
			return err
		} else {
			for _, n := range networks {
				for _, pair := range n.TokenPairs {
					if pair.TokenStandard.String() == types.ZnnTokenStandard.String() {
						node.state.SetTokensMap(n.Id, types.ZnnTokenStandard.String(), pair.TokenAddress)
						node.logger.Infof("Set chainId: %d wZNN address to %s", n.Id, pair.TokenAddress)
					} else if pair.TokenStandard.String() == types.QsrTokenStandard.String() {
						node.state.SetTokensMap(n.Id, types.QsrTokenStandard.String(), pair.TokenAddress)
						node.logger.Infof("Set chainId: %d wQSR address to %s", n.Id, pair.TokenAddress)
					}
				}
			}
		}

		metadata := &common.BridgeMetadata{}
		if err := json.Unmarshal([]byte(bridgeInfo.Metadata), metadata); err != nil {
			return err
		} else {
			node.state.SetIsAffiliateProgram(metadata.AffiliateProgram)
		}
	}

	if err := node.networksManager.Start(); err != nil {
		return err
	}

	//  Don't start tss if the administrator runs the orchestrator
	if !node.state.GetIsAdministratorActive() {
		var tssErr error
		node.tssManager, tssErr = tss.NewTssManager(node.config.TssConfig, base64.StdEncoding.EncodeToString(node.producerKeyPair.Private))
		if tssErr != nil {
			return tssErr
		}
	}

	bridgeInfo, err = node.networksManager.Znn().GetBridgeInfo()
	if err != nil {
		return err
	} else {
		metadata := &common.BridgeMetadata{}
		if err := json.Unmarshal([]byte(bridgeInfo.Metadata), metadata); err != nil {
			return err
		} else {
			node.SetBridgeMetadata(metadata)
		}
	}

	currentState, stateErr := node.state.GetState()
	if stateErr != nil {
		return stateErr
	}
	node.logger.Infof("Current state is: %s", common.StateToText(currentState))

	// If we are in emergency we should stay so we can halt the network
	// Otherwise we should check
	if currentState == common.LiveState {
		if halted, _, err := node.networksManager.CountNetworksHaltState(); err != nil {
			return err
		} else {
			if node.networksManager.Znn().IsHalted() || halted > 0 {
				if err := node.state.SetState(common.HaltedState); err != nil {
					return err
				}
			}
		}
	}

	go node.processSignatures()
	go node.sendSignatures()
	return nil
}

func (node *Node) Stop() error {
	node.lock.Lock()
	defer node.lock.Unlock()
	defer close(node.stopChan)
	node.logger.Info("stopping node ...")

	node.tssManager.Stop()
	node.networksManager.Stop()

	// Release instance directory lock.
	node.closeDataDir()
	node.logger.Info("closed stopChan")
	return nil
}

func (node *Node) Wait() {
	signalReceived := <-node.stopChan
	node.logger.Info("signal from wait: ", signalReceived)
}

// Sign messages methods

func (node *Node) processSignatures() {
	node.logger.Info("in node processSignatures")
	for {
		time.Sleep(time.Second)
		currentState, err := node.state.GetState()
		if err != nil {
			node.logger.Debug("currentState error in processSig")
			node.logger.Debug(err.Error())
			time.Sleep(5 * time.Second)
			continue
		}

		if node.state.GetIsAdministratorActive() {
			if currentState != common.EmergencyState {
				continue
			}
			node.haltNetworksAdministrator()
			time.Sleep(5 * time.Second)
			continue
		}
		switch currentState {
		case common.KeyGenState:
			mom, err := node.networksManager.Znn().GetFrontierMomentum()
			if err != nil {
				node.logger.Debug(err)
				continue
			}
			// We start a key gen every 216 momentums ( 36 minutes )
			for mom.Height%270 > 3 {
				time.Sleep(10 * time.Second)
				mom, err = node.networksManager.Znn().GetFrontierMomentum()
				if err != nil {
					node.logger.Debug(err)
					continue
				}

				if mom.Height%24 == 0 {
					node.logger.Infof("%d momentums left before starting keygen", 270-mom.Height%270)
				}

				// we query the state every 1 minute
				if mom.Height%6 == 0 {
					state, err := node.state.GetState()
					if err != nil {
						node.logger.Debug("currentState error while sleeping for key gen")
						node.logger.Debug(err.Error())
						continue
					}
					if state != common.KeyGenState {
						node.logger.Info("State no longer keyGen so will exit key generation")
						break
					}
				}
			}

			node.logger.Info("Starting keyGen")

			// We start looking 24h of momentums behind the momentum height specified by the bridgeInfo.shouldKeyGenAt
			// and select every pubKey that produced a momentum
			participatingPubKeys, err := node.networksManager.GetPillarPubKeys()
			if err != nil {
				node.logger.Debug(err)
				continue
			}

			// we create a join of the old pub keys list and the new ones
			node.allowNewParticipants(participatingPubKeys)

			// each round of keyGen will remove nodes that do not participate
			// we try to keygen until we have min parties threshold
			var keyGenResponse *keygen.Response

			initialParticipantsLength := node.getParticipantsLength()
			node.logger.Debug("node.getParticipantsLength(): ", initialParticipantsLength)

			for node.networksManager.Znn().KeyGenThreshold() <= node.getParticipantsLength() {
				state, err := node.state.GetState()
				if err != nil {
					node.logger.Debug("currentState error in processSig")
					node.logger.Debug(err.Error())
					time.Sleep(5 * time.Second)
					continue
				}
				if state != common.KeyGenState {
					node.logger.Info("State no longer keyGen so will exit key generation")
					keyGenResponse = nil
					break
				}
				keyGenThreshold := node.networksManager.Znn().KeyGenThreshold()
				node.logger.Debug("keyGenThreshold: ", keyGenThreshold)
				node.logger.Info("Started ECDSA Keygen")
				node.logger.Debug("len(node.participatingPubKeys): ", node.getParticipantsLength())

				duration := time.Duration(5 * 60 * 1e9) // 5 minutes
				node.logger.Infof("Old party timeout value: %f minutes", node.tssManager.Config().PartyTimeout.Minutes())
				node.tssManager.SetPartyTimeout(duration)
				node.logger.Infof("Set party timeout to value: %f minutes", duration.Minutes())
				node.logger.Infof("New party timeout value: %f minutes", node.tssManager.Config().PartyTimeout.Minutes())

				// start the key gen
				start := time.Now()
				keyGenResponse, err = node.tssManager.KeyGen(messages.ECDSAKEYGEN)
				elapsed := time.Since(start)
				node.logger.Infof("keyGen took %f", elapsed.Seconds())
				if keyGenResponse != nil {
					node.logger.Infof("keyGen threshold %d", keyGenResponse.Threshold)
					node.logger.Infof("KeyGen status: %d (0 NA, 1 Success, 2 Fail)", keyGenResponse.Status)
				}

				// Set the old party timeout
				node.tssManager.SetPartyTimeout(node.config.TssConfig.BaseConfig.PartyTimeout)
				node.logger.Infof("Set party timeout to old value: %f minutes", node.config.TssConfig.BaseConfig.PartyTimeout.Minutes())

				if err != nil {
					node.logger.Debug(err)
					continue
				}

				node.logger.Debug("len(node.participatingPubKeys): ", node.getParticipantsLength())
				p2pWhitelist, partyWhitelist, sigWhitelist := node.tssManager.GetWhitelists()
				for _, blamedNode := range keyGenResponse.Blame.BlameNodes {
					node.removeParticipant(blamedNode.Pubkey)
					node.logger.Debugf("Blamed node pubKey: %s", blamedNode.Pubkey)

					// remove peerId from whitelist
					pubKeyBytes, err := base64.StdEncoding.DecodeString(blamedNode.Pubkey)
					if err != nil {
						node.logger.Debugf("PubKey from blamed node could not be decoded: %s", blamedNode.Pubkey)
						continue
					}
					pub, err := ic.UnmarshalEd25519PublicKey(pubKeyBytes)
					if err != nil {
						node.logger.Debugf("PubKey from blamed node could not be unmarshaled: %s", blamedNode.Pubkey)
						continue
					}
					id, err := peer.IDFromPublicKey(pub)
					if err != nil {
						node.logger.Debugf("PubKey from blamed node could not be transformed into peerId: %s", blamedNode.Pubkey)
						continue
					}
					p2pWhitelist[id.String()] = false
					partyWhitelist[id.String()] = false
					sigWhitelist[id.String()] = false
				}
				node.logger.Debug("len(node.participatingPubKeys) after removing blamed nodes: ", node.getParticipantsLength())

				if keyGenThreshold > node.getParticipantsLength() {
					node.logger.Info("KeyGen threshold was not met")
					if keyGenResponse.Status == tcommon.Success {
						if err := common.DeletePubKeyFile(node.config.TssConfig.BaseDir, keyGenResponse.PubKey); err != nil {
							node.logger.Error(err)
						}
					}
					keyGenResponse = nil
					break
				}

				// key gen was generated
				if keyGenResponse.Status == tcommon.Success {
					node.logger.Infof("Generated key: %s", keyGenResponse.PubKey)
					break
				} else {
					// Sleep so we can close current connections
					time.Sleep(15 * time.Second)
				}
			}
			// there was en error while trying to keyGen, we retry
			if keyGenResponse == nil {
				continue
			} else if keyGenResponse.Status != tcommon.Success {
				continue
			}

			// 33 bytes in base64 will always be 44 chars
			if len(keyGenResponse.PubKey) != 44 {
				node.logger.Debug("Generated pubKey base64 encoded form length is not 44 characters long")
				continue
			}

			decompressedKeyGenPubKey, errDecompress := common.DecompressPubKey(keyGenResponse.PubKey)
			if errDecompress != nil {
				node.logger.Debug(errDecompress)
				continue
			}

			time.Sleep(5 * time.Second)
			// we save the old pubKey because we will need it
			oldKey := node.tssManager.GetPubKey()
			node.logger.Debugf("Old pubKey: %s", oldKey)

			// if len is 0, it means that we have not generated a key before and the administrator needs to set it manually
			// otherwise, orchestrator will generate a signature using the new key and validate it
			// it will then create a signature with the old key to send an account block that changes the key
			if len(oldKey) != 0 {
				decompressedOldKey, errDecompress := common.DecompressPubKey(oldKey)
				if errDecompress != nil {
					node.logger.Debug(errDecompress)
					continue
				}

				node.logger.Debug("Old key found")
				znnTssNonce, err := node.networksManager.GetTssNonceZnn()
				if err != nil {
					node.logger.Error(err)
					continue
				}

				// ZNN
				msgsIndexes := make(map[string]int)
				znnMessage, err := implementation.GetChangePubKeyMessage(definition.ChangeTssECDSAPubKeyMethodName, definition.NoMClass, 1, znnTssNonce, keyGenResponse.PubKey)
				if err != nil {
					node.logger.Debug(err)
					continue
				}
				msgsIndexes[base64.StdEncoding.EncodeToString(znnMessage)] = 0
				messagesToSign := make([][]byte, 0)
				messagesToSign = append(messagesToSign, znnMessage)

				// EVM messages
				toSignMessagesEvm, err := node.networksManager.GetSetTssEcdsaPubKeysEvmMessages(keyGenResponse.PubKey)
				if err != nil {
					node.logger.Debug(err)
					continue
				}
				for idx, msg := range toSignMessagesEvm {
					messagesToSign = append(messagesToSign, msg)
					msgsIndexes[base64.StdEncoding.EncodeToString(msg)] = idx + 1
				}
				node.logger.Debug("msg to sign after evm append: ", messagesToSign)
				// New key sign of messages
				// we will try to generate a signature using the new key secret shares and validate it
				node.tssManager.SetPubKey(keyGenResponse.PubKey)

				newKeySignResponse, err := node.signMessages(messagesToSign, msgsIndexes)
				if err != nil {
					// we set back the old key
					node.tssManager.SetPubKey(oldKey)
					node.logger.Debug(err)
					node.logger.Debug(" Error after signing with the new generated key 1")
					continue
				} else if newKeySignResponse.Status != tcommon.Success {
					node.tssManager.SetPubKey(oldKey)
					node.logger.Debugf("keySignStatus: %d", newKeySignResponse.Status)
					node.logger.Debug(" Error after signing with the new generated key 2")
					continue
				}

				// Verify all signatures
				znnNewKeySignature, newKeyFullSignatures, errValidate := node.validateSignatures(newKeySignResponse, decompressedKeyGenPubKey, messagesToSign)
				if errValidate != nil {
					node.logger.Info(errValidate.Error())
					continue
				}
				node.logger.Debug("Signatures generated from the new key are valid!")

				// Old key sign of messages
				node.tssManager.SetPubKey(oldKey)

				oldKeySignResponse, err := node.signMessages(messagesToSign, msgsIndexes)
				if err != nil {
					node.logger.Debug(err)
					node.logger.Debug("Error when trying to sign evm msgs with the old key")
					continue
				} else if oldKeySignResponse.Status != tcommon.Success {
					node.logger.Debugf("keySignStatus: %d", oldKeySignResponse.Status)
					node.logger.Debug(" Error after signing with the old key 2")
					continue
				}

				// Verify all signatures
				znnOldKeySignature, oldKeyFullSignatures, errValidate := node.validateSignatures(oldKeySignResponse, decompressedOldKey, messagesToSign)
				if errValidate != nil {
					node.logger.Debug(errValidate)
					continue
				}
				node.logger.Debug("Signatures generated from the old key are valid!")

				var senders sync.WaitGroup
				senders.Add(2)
				//we send the account block on the znn network
				go func() {
					defer senders.Done()
					// we take the first base64 char of the signature and transform it to its ascii value
					index := uint32(int(znnOldKeySignature[0])) % node.getParticipantsLength()
					for {
						index = (index + 1) % node.getParticipantsLength()
						producerPubKey := base64.StdEncoding.EncodeToString(node.producerKeyPair.Public)
						bridgeInfo, err := node.networksManager.GetBridgeInfo()
						if err != nil {
							node.logger.Debug(err)
							continue
						}
						// The pubKey was changed
						if bridgeInfo.CompressedTssECDSAPubKey == keyGenResponse.PubKey {
							break
						}
						if producerPubKey == node.getParticipant(index) {
							node.logger.Debug("[sendZnnTx PubKey] this is me")
							err = node.networksManager.ChangeTssEcdsaPubKeyZnn(keyGenResponse.PubKey, znnOldKeySignature, znnNewKeySignature, node.producerKeyPair)
							if err != nil {
								node.logger.Debug(err)
								continue
							}
							node.logger.Debug("[sendZnnTx PubKey] sent tx")
							// we wait 2 momentums so that the send and receive block are inserted and the pubKey changes
							// todo use constants
							time.Sleep(20 * time.Second)

						}
						time.Sleep(20 * time.Second)
					}
				}()

				go func() {
					defer senders.Done()
					index := uint32(int(znnOldKeySignature[0])) % node.getParticipantsLength()
					for {
						index = (index + 1) % node.getParticipantsLength()
						producerPubKey := base64.StdEncoding.EncodeToString(node.producerKeyPair.Public)
						if producerPubKey == node.getParticipant(index) {
							node.logger.Debug("[send set tss ecdsa pub key evm tx] this is me")
							if changed, err := node.networksManager.SetTssEcdsaPubKeyEvm(oldKeyFullSignatures, newKeyFullSignatures, keyGenResponse.PubKey, node.ecdsaPrivateKey, node.evmAddress); err != nil {
								node.logger.Debug(err)
								continue
							} else if changed {
								break
							}
						}
						// todo use estimatedBlockTime / TimeToFinality
						time.Sleep(20 * time.Second)
					}
					node.logger.Info("Successfully set pubKey on all networks")
				}()

				senders.Wait()
			} else {
				znnTssNonce, err := node.networksManager.GetTssNonceZnn()
				if err != nil {
					continue
				}

				// ZNN
				msgsIndexes := make(map[string]int)
				znnMessage, err := implementation.GetChangePubKeyMessage(definition.ChangeTssECDSAPubKeyMethodName, definition.NoMClass, 1, znnTssNonce, keyGenResponse.PubKey)
				if err != nil {
					continue
				}
				msgsIndexes[base64.StdEncoding.EncodeToString(znnMessage)] = 0
				messagesToSign := make([][]byte, 0)
				messagesToSign = append(messagesToSign, znnMessage)

				// EVM messages
				toSignMessagesEvm, err := node.networksManager.GetSetTssEcdsaPubKeysEvmMessages(keyGenResponse.PubKey)
				if err != nil {
					continue
				}
				for idx, msg := range toSignMessagesEvm {
					messagesToSign = append(messagesToSign, msg)
					msgsIndexes[base64.StdEncoding.EncodeToString(msg)] = idx + 1
				}
				// New key sign of messages
				// we will try to generate a signature using the new key secret shares and validate it
				node.tssManager.SetPubKey(keyGenResponse.PubKey)

				newKeySignResponse, err := node.signMessages(messagesToSign, msgsIndexes)
				if err != nil {
					continue
				}

				// Verify all signatures
				znnNewKeySignature, newKeyFullSignatures, errValidate := node.validateSignatures(newKeySignResponse, decompressedKeyGenPubKey, messagesToSign)
				if errValidate != nil {
					node.logger.Info(errValidate.Error())
					continue
				}

				node.logger.Infof("znnNewKeySignature: %s\n", znnNewKeySignature)
				node.logger.Infof("newKeyFullSignatures: %s\n", hex.EncodeToString(newKeyFullSignatures[0]))
			}

			node.resetSignatures()

			node.logger.Infof("ECDSA KeyGen Response here: %v", keyGenResponse)

			node.config.TssConfig.PublicKey = keyGenResponse.PubKey
			node.config.TssConfig.DecompressedPublicKey = decompressedKeyGenPubKey
			node.tssManager.SetPubKey(keyGenResponse.PubKey)
			if stateErr := node.state.SetState(common.LiveState); stateErr != nil {
				node.logger.Info("Could not set state to live after key gen")
				node.logger.Debug(stateErr)
				node.stopChan <- syscall.SIGINT
				return
			}
			// todo delete it?
			//if len(oldKey) > 0 {
			//	if delErr := common.DeletePubKeyFile(node.config.TssConfig.BaseDir, oldKey); delErr != nil {
			//		node.logger.Debug(delErr)
			//	}
			//}

			if err := oconfig.WriteConfig(*node.config); err != nil {
				node.logger.Info(err.Error())
			}
		case common.LiveState:
			time.Sleep(10 * time.Second)
			bridgeInfo, err := node.networksManager.GetBridgeInfo()
			if err != nil {
				node.logger.Debug(err)
				continue
			} else if bridgeInfo == nil {
				node.logger.Debug("processSignatures bridgeInfo == nil")
				continue
			}
			if bridgeInfo.AllowKeyGen == true {
				// this means we generated a key and we wait for the administrator to change it
				if len(bridgeInfo.CompressedTssECDSAPubKey) == 0 && len(node.tssManager.GetPubKey()) != 0 {
					time.Sleep(10 * time.Second)
					continue
				}
				if err := node.state.SetState(common.KeyGenState); err != nil {
					node.logger.Debug(err.Error())
				}
			} else if node.tssManager.CanProcessSignatures() && bridgeInfo.CompressedTssECDSAPubKey == node.tssManager.GetPubKey() {
				frontierMomentum, getFrErr := node.state.GetFrontierMomentum()
				if getFrErr != nil {
					node.logger.Debug(getFrErr)
					continue
				}

				currentCeremony := frontierMomentum / node.networksManager.WindowSize()
				state := currentCeremony % 2
				lastCeremony := node.state.GetLastCeremony()
				if state == common.WrapCeremonyState && currentCeremony != lastCeremony {
					node.logger.Debugf("state wrap - lastCeremony: %d", lastCeremony)
					node.state.SetLastCeremony(currentCeremony)
					if err, ok := node.processSignaturesWrap(); err != nil {
						node.logger.Debug(err)
					} else if ok == false {
						node.logger.Debug("Nothing to sign, try unwrap")
						// If we had nothing to sign we try the other direction
						if err, _ := node.processSignaturesUnwrap(); err != nil {
							node.logger.Debug(err)
						}
					}
				} else if state == common.UnwrapCeremonyState && currentCeremony != lastCeremony {
					node.logger.Debugf("state unwrap - lastCeremony: %d", lastCeremony)
					node.state.SetLastCeremony(currentCeremony)
					if err, ok := node.processSignaturesUnwrap(); err != nil {
						node.logger.Debug(err)
					} else if ok == false {
						node.logger.Debug("Nothing to sign, try wrap")
						// If we had nothing to sign we try the other direction
						if err, _ := node.processSignaturesWrap(); err != nil {
							node.logger.Debug(err)
						}
					}
				}
			}
		case common.EmergencyState:
			if node.state.GetIsAdministratorActive() {
				node.haltNetworksAdministrator()
				continue
			}

			// todo when to exit this state in case we don't halt the network after some time
			node.logger.Debug("Process in emergency state")
			time.Sleep(5 * time.Second)
			znnTssNonce, err := node.networksManager.GetTssNonceZnn()
			if err != nil {
				node.logger.Debug(err)
				continue
			}
			decompressedPubKey, err := common.DecompressPubKey(node.tssManager.GetPubKey())
			if err != nil {
				node.logger.Debug(err)
				continue
			}

			msgsIndexes := make(map[string]int)
			znnMessage, err := implementation.GetBasicMethodMessage("Halt", znnTssNonce, definition.NoMClass, 1)
			if err != nil {
				node.logger.Debug(err)
				continue
			}
			msgsIndexes[base64.StdEncoding.EncodeToString(znnMessage)] = 0
			messagesToSign := make([][]byte, 0)
			messagesToSign = append(messagesToSign, znnMessage)

			evmMessages, err := node.networksManager.GetHaltMessages()
			if err != nil {
				node.logger.Debug(err)
				continue
			}
			for idx, msg := range evmMessages {
				messagesToSign = append(messagesToSign, msg)
				msgsIndexes[base64.StdEncoding.EncodeToString(msg)] = idx + 1
			}
			node.logger.Debug("msg to sign after evm append: ", messagesToSign)

			keySignResponse, err := node.signMessages(messagesToSign, msgsIndexes)
			if err != nil {
				node.logger.Debug(err)
				node.logger.Debug(" error after signing halt 1")
				continue
			} else if keySignResponse.Status != tcommon.Success {
				node.logger.Debug(keySignResponse.Status)
				node.logger.Debug("error after signing halt 2")
				continue
			}

			znnSignature, evmFullSignatures, err := node.validateSignatures(keySignResponse, decompressedPubKey, messagesToSign)
			if err != nil {
				node.logger.Debug(err)
				continue
			}

			node.logger.Debug("Halt signatures are valid!")
			var senders sync.WaitGroup
			senders.Add(2)
			go func() {
				defer senders.Done()
				index := uint32(int(znnSignature[0])) % node.getParticipantsLength()
				if !node.networksManager.Znn().IsHalted() {
					for {
						index = (index + 1) % node.getParticipantsLength()
						if node.networksManager.Znn().IsHalted() {
							break
						}
						producerPubKey := base64.StdEncoding.EncodeToString(node.producerKeyPair.Public)
						if producerPubKey == node.getParticipant(index) {
							node.logger.Debug("[send Halt Znn Tx] this is me")
							err = node.networksManager.HaltZnn(znnSignature, node.producerKeyPair)
							if err != nil {
								node.logger.Info(err.Error())
								continue
							}
							node.logger.Debug("[send Halt Znn Tx] sent request")
							// we wait 2 momentums so that the send and receive block are inserted and the pubKey changes
							// todo wait use constant
							time.Sleep(20 * time.Second)
						}
					}
					time.Sleep(time.Second)
				}
			}()

			go func() {
				defer senders.Done()
				index := uint32(int(znnSignature[0])) % node.getParticipantsLength()
				for {
					index = (index + 1) % node.getParticipantsLength()
					producerPubKey := base64.StdEncoding.EncodeToString(node.producerKeyPair.Public)
					if producerPubKey == node.getParticipant(index) {
						node.logger.Debug("[send halt evm tx] this is me")
						// todo discuss using concurrent send for different networks with different senders
						if changed, err := node.networksManager.SendHaltEvm(evmFullSignatures, node.ecdsaPrivateKey, node.evmAddress); err != nil {
							node.logger.Debug(err)
							continue
						} else if changed {
							break
						} else {
							// todo use estimatedBlockTime / TimeToFinality
							time.Sleep(15 * time.Second)
						}
					}
					time.Sleep(time.Second)
				}
			}()
			senders.Wait()

			if err := node.state.SetState(common.HaltedState); err != nil {
				node.logger.Error(err)
				node.stopChan <- syscall.SIGINT
				return
			}
		case common.HaltedState:
			if halted, unhalted, err := node.networksManager.CountNetworksHaltState(); err != nil {
				node.logger.Debug(err)
			} else {
				// the state should be given by the znn network
				// when unhalting, start with the znn network
				if node.networksManager.Znn().IsHalted() {
					// if some of the networks are not halted, either some error occurred while sending halt tx or the tx is not accepted yet
					// in any case we want to try to halt the networks again
					if unhalted > 0 {
						if err := node.state.SetState(common.EmergencyState); err != nil {
							node.logger.Debug(err)
							continue
						}
					}
				} else {
					// we wait so the administrator unhalts all the networks if znn is unhalted
					if halted == 0 {
						if err := node.state.SetState(common.LiveState); err != nil {
							node.logger.Debug(err)
							continue
						}
					}
				}
			}
		}
	}
}

// ok == true we had a signing ceremony
// ok == false means that we had nothing to sign
func (node *Node) processSignaturesWrap() (error, bool) {
	wrapRequestsIds, err := node.networksManager.GetUnsignedWrapRequests()
	if err != nil {
		return err, false
	} else if wrapRequestsIds == nil {
		return nil, false
	}

	messagesToSign := make([][]byte, 0)
	msgsIndexes := make(map[string]int)
	node.logger.Debugf("WrapRequests len: %d", len(wrapRequestsIds))
	for idx, request := range wrapRequestsIds {
		event, err := node.networksManager.GetWrapEventById(request.Id)
		if err != nil || event == nil {
			// if the rpc returns it but we don't have it locally there was a problem
			// if we don't have it we just add it now
			if errStorage := node.networksManager.Znn().AddWrapEvent(request.WrapTokenRequest); errStorage != nil {
				node.logger.Error(err)
				node.stopChan <- syscall.SIGINT
				return err, false
			}
			event, err = node.networksManager.GetWrapEventById(request.Id)
			if err != nil || event == nil {
				node.logger.Error(err)
				node.stopChan <- syscall.SIGINT
				return err, false
			}
		} else if len(event.Signature) > 0 {
			continue
		}
		node.logger.Debugf("param.Id: %s", request.Id.String())
		node.logger.Debugf("param.ToAddress: %s", request.ToAddress)
		node.logger.Debugf("param.TokenAddress: %s", request.TokenAddress)
		node.logger.Debugf("param.Amount: %s", request.Amount.String())

		if localEvent, errStorage := node.dbManager.ZnnStorage().GetWrapRequestById(event.Id); err != nil {
			node.logger.Debug(errStorage)
		} else {
			// if we have a signature but it was not sent yet, do not sign again
			// on a new key sign, this will be set as unsigned if the signature set will not be succeeded
			if len(localEvent.Signature) > 0 {
				continue
			}
		}

		msg, err := event.GetMessage(node.networksManager.Evm(event.ChainId).ContractAddress())
		if err != nil {
			node.logger.Debug(err)
			// todo set status to skipped
			continue
		}
		messagesToSign = append(messagesToSign, msg)
		msgsIndexes[base64.StdEncoding.EncodeToString(msg)] = idx
	}

	if len(messagesToSign) == 0 {
		return nil, false
	}

	response, err := node.signMessages(messagesToSign, msgsIndexes)
	if err != nil {
		return err, true
	} else if response.Status != tcommon.Success {
		node.logger.Debug(response.Status)
		node.logger.Debug(" error signing wrap")
		return nil, true
	}

	// we apply the signatures that don't return error
	for idx, sig := range response.Signatures {
		signature, err := base64.StdEncoding.DecodeString(sig.Signature)
		if err != nil {
			node.logger.Debug(err)
			continue
		}
		recoverID, err := base64.StdEncoding.DecodeString(sig.RecoveryID)
		fullSignature := append(signature, recoverID...)
		fullSignatureStr := base64.StdEncoding.EncodeToString(fullSignature)

		ok, err := implementation.CheckECDSASignature(messagesToSign[idx], node.config.TssConfig.DecompressedPublicKey, fullSignatureStr)
		if err != nil {
			node.logger.Debug("Error checking ecdsa signature for wrap: %s", err.Error())
			continue
		} else if ok == false {
			node.logger.Debugf("invalid signature when checking ecdsa signature for wrap msg: %s", messagesToSign[idx])
			continue
		}

		if err = node.networksManager.SetWrapEventSignature(wrapRequestsIds[msgsIndexes[sig.Msg]].Id, fullSignatureStr); err != nil {
			node.logger.Debug(err)
			continue
		}
		node.logger.Infof("%d. msg: %s sig: %s\n", msgsIndexes[sig.Msg], base64.StdEncoding.EncodeToString(messagesToSign[idx]), sig.Signature)
	}
	return nil, true
}

// ok == true means that we had a signing ceremony or at least we tried
// ok == false means that we had nothing to sign
func (node *Node) processSignaturesUnwrap() (error, bool) {
	// todo discuss what requests should the orchestrator sign in a ceremony
	requests, err := node.networksManager.GetUnsignedUnwrapRequests()
	if err != nil {
		return err, true
	}
	if len(requests) == 0 {
		return nil, false
	}

	msgsIndexes := make(map[string]int)
	// arr for signing
	messagesToSign := make([][]byte, 0)
	for idx, req := range requests {
		node.logger.Debugf("req.To: %s", req.To)
		node.logger.Debugf("req.Amount: %d", req.Amount.Uint64())
		node.logger.Debugf("req.Token: %s", req.Token.String())
		node.logger.Debugf("req.TransactionHash: %s", req.TransactionHash.String())

		msg, err := req.GetMessage()
		if err != nil {
			node.logger.Debug(err)
			// todo set status to skipped
			continue
		}
		messagesToSign = append(messagesToSign, msg)
		msgsIndexes[base64.StdEncoding.EncodeToString(msg)] = idx
	}

	if len(messagesToSign) == 0 {
		return nil, false
	}
	node.logger.Debug("MessagesToSign: ", messagesToSign)
	response, err := node.signMessages(messagesToSign, msgsIndexes)
	if err != nil {
		return err, true
	} else if response.Status != tcommon.Success {
		node.logger.Debug(response.Status)
		node.logger.Debug(" error signing unwrap")
		return nil, true
	}

	for idx, sig := range response.Signatures {
		signature, err := base64.StdEncoding.DecodeString(sig.Signature)
		if err != nil {
			node.logger.Debug(err)
			continue
		}
		recoverID, err := base64.StdEncoding.DecodeString(sig.RecoveryID)
		fullSignature := append(signature, recoverID...)
		requests[msgsIndexes[sig.Msg]].Signature = base64.StdEncoding.EncodeToString(fullSignature)

		ok, err := implementation.CheckECDSASignature(messagesToSign[idx], node.config.TssConfig.DecompressedPublicKey, requests[msgsIndexes[sig.Msg]].Signature)
		if err != nil {
			node.logger.Debug("Error checking ecdsa signature for unwrap: %s", err.Error())
			continue
		} else if ok == false {
			node.logger.Debugf("invalid signature when checking ecdsa signature for unwrap msg: %s", messagesToSign[idx])
			continue
		}

		if err = node.networksManager.AddEvmUnwrapRequest(*requests[msgsIndexes[sig.Msg]]); err != nil {
			node.logger.Debug(err.Error())
			continue
		}
		node.logger.Debugf("\n%d. msg: %s\n\n", idx, base64.StdEncoding.EncodeToString(messagesToSign[idx]))
		node.logger.Debugf("\n%d. sig: %s, recoveryID: %s, FinalSig: %s \n", msgsIndexes[sig.Msg], sig.Signature, sig.RecoveryID, requests[msgsIndexes[sig.Msg]].Signature)
	}
	return nil, true
}

// Send signatures methods

func (node *Node) sendSignatures() {
	// Early return
	if node.state.GetIsAdministratorActive() {
		return
	}

	seenEventsCount := make(map[string]uint32)
	for {
		currentState, err := node.state.GetState()
		if err != nil {
			node.logger.Debug("currentState error in processSig")
			node.logger.Debug(err.Error())
			time.Sleep(5 * time.Second)
			continue
		}
		switch currentState {
		case common.LiveState:
			if len(node.tssManager.GetPubKey()) != 0 {
				var senders sync.WaitGroup
				senders.Add(2)
				go func() {
					defer senders.Done()
					node.sendSignaturesWrap(seenEventsCount)
				}()
				go func() {
					defer senders.Done()
					node.sendUnwrapRequests(seenEventsCount)
				}()
				senders.Wait()
			}
		}
		// todo use constant for momentum duration
		time.Sleep(10 * time.Second)
	}
}

func (node *Node) sendSignaturesWrap(seenEventsCount map[string]uint32) {
	requests, err := node.networksManager.GetUnsentSignedWrapRequests()
	if err != nil {
		node.logger.Debug(err)
		return
	}

	for _, req := range requests {
		rpcRequest, err := node.networksManager.GetWrapRequestByIdRPC(req.Id)
		if err != nil {
			node.logger.Debug(err)
			continue
		} else if len(rpcRequest.Signature) != 0 {
			delete(seenEventsCount, req.Id.String())
			continue
		}

		index := uint32(req.Id.Bytes()[31]) % node.getParticipantsLength()
		if seenEventsCount[req.Id.String()] > 2 {
			index = (index + seenEventsCount[req.Id.String()]) % node.getParticipantsLength()
		}
		seenEventsCount[req.Id.String()] += 1
		producerPubKey := base64.StdEncoding.EncodeToString(node.producerKeyPair.Public)
		if producerPubKey == node.getParticipant(index) {
			node.logger.Info("[sendSignaturesWrap] this is me")
			err = node.networksManager.UpdateWrapRequest(req.Id, req.Signature, node.producerKeyPair)
			if err != nil {
				node.logger.Debug(err)
				continue
			}
			delete(seenEventsCount, req.Id.String())
			node.logger.Info("[sendSignaturesWrap] sent request")
		}
		// todo how much to wait between sends?
		time.Sleep(25 * time.Second)
	}
}

func (node *Node) sendUnwrapRequests(seenEventsCount map[string]uint32) {
	requests, err := node.networksManager.GetUnsentSignedUnwrapRequests()
	if err != nil {
		node.logger.Debug(err)
		return
	}
	for _, req := range requests {
		rpcReq, err := node.networksManager.GetEvmUnwrapRequestByHashAndLogFromRPC(types.Hash(req.TransactionHash), req.LogIndex)
		if rpcReq != nil {
			node.logger.Debug("event exists, set it as sent locally")
			if err := node.networksManager.SetEvmUnwrapRequestAsSent(req); err != nil {
				node.logger.Debug(err)
			}
			delete(seenEventsCount, req.TransactionHash.String())
			continue
		}

		index := uint32(req.TransactionHash.Bytes()[31]) % node.getParticipantsLength()
		if seenEventsCount[req.TransactionHash.String()] > 2 {
			index = (index + seenEventsCount[req.TransactionHash.String()]) % node.getParticipantsLength()
		}
		seenEventsCount[req.TransactionHash.String()] += 1
		producerPubKey := base64.StdEncoding.EncodeToString(node.producerKeyPair.Public)
		if producerPubKey == node.getParticipant(index) {
			node.logger.Info("[sendUnwrapRequests] this is me")
			if len(req.Signature) == 0 {
				node.logger.Debug("signature missing")
				continue
			}
			err = node.networksManager.SendUnwrapRequest(req, node.producerKeyPair)
			if err != nil {
				node.logger.Debug(err)
			}
		}
		// todo how much to wait between sends?
		time.Sleep(25 * time.Second)
	}
}

// Administrator methods

func (node *Node) haltNetworksAdministrator() {
	var senders sync.WaitGroup
	senders.Add(1)

	go func() {
		defer senders.Done()
		for {
			if node.networksManager.Znn().IsHalted() {
				node.logger.Info("ZNN Network is halted")
				break
			}

			errZnn := node.networksManager.HaltZnn("", node.producerKeyPair)
			if errZnn != nil {
				node.logger.Info(errZnn.Error())
			}
			time.Sleep(20 * time.Second)
		}
	}()

	l := node.networksManager.NetworksLength()
	senders.Add(l)

	for idx := 0; idx < l; idx++ {
		go func(index int) {
			defer senders.Done()
			for {
				halted, err := node.networksManager.SendHaltEvmAdministrator(index, node.ecdsaPrivateKey, node.evmAddress)
				if err != nil {
					node.logger.Error(err)
				} else if halted {
					break
				}
				// todo use network block time
				time.Sleep(12 * time.Second)
			}
		}(idx)
	}
	senders.Wait()
	if err := node.state.SetState(common.HaltedState); err != nil {
		node.logger.Debug(err)
	}
}

// Getters

func (node *Node) GetConfig() *oconfig.Config {
	return node.config
}

func (node *Node) getParticipantsLength() uint32 {
	// we know for sure that these pubKeys were in our keyGen, or we don't have the keyGen at all
	return uint32(len(node.config.TssConfig.LocalPubKeys))
}

func (node *Node) getParticipant(index uint32) string {
	return node.config.TssConfig.LocalPubKeys[index]
}

// Setters

func (node *Node) SetBridgeMetadata(metadata *common.BridgeMetadata) {
	node.logger.Info("In SetBridgeMetadata:")
	if metadata != nil {
		if node.tssManager != nil {
			if metadata.PartyTimeout != 0 {
				duration := time.Duration(metadata.PartyTimeout) * time.Second
				node.logger.Infof("PartyTimeout in seconds - old: %f, new: %d", node.config.TssConfig.BaseConfig.PartyTimeout.Seconds(), metadata.PartyTimeout)
				node.tssManager.SetPartyTimeout(duration)
				node.config.TssConfig.BaseConfig.PartyTimeout = duration
			}

			if metadata.KeyGenTimeout != 0 {
				duration := time.Duration(metadata.KeyGenTimeout) * time.Second
				node.logger.Infof("KeyGenTimeout in seconds - old: %f, new: %d", node.config.TssConfig.BaseConfig.KeyGenTimeout.Seconds(), metadata.KeyGenTimeout)
				node.tssManager.SetKeyGenTimeout(duration)
				node.config.TssConfig.BaseConfig.KeyGenTimeout = duration
			}

			if metadata.KeySignTimeout != 0 {
				duration := time.Duration(metadata.KeySignTimeout) * time.Second
				node.logger.Infof("KeySignTimeout in seconds - old: %f, new: %d", node.config.TssConfig.BaseConfig.KeySignTimeout.Seconds(), metadata.KeySignTimeout)
				node.tssManager.SetKeySignTimeout(duration)
				node.config.TssConfig.BaseConfig.KeySignTimeout = duration
			}

			if metadata.PreParamTimeout != 0 {
				duration := time.Duration(metadata.PreParamTimeout) * time.Second
				node.logger.Infof("PreParamsTimeout in seconds - old: %f, new: %d", node.config.TssConfig.BaseConfig.PreParamTimeout.Seconds(), metadata.PreParamTimeout)
				node.tssManager.SetPreParamsTimeout(duration)
				node.config.TssConfig.BaseConfig.PreParamTimeout = duration
			}

			if len(metadata.JoinPartyVersion) > 0 {
				node.logger.Infof("joinPartyVersion - old: %s, new: %s", node.tssManager.GetJoinPartyVersion(), metadata.JoinPartyVersion)
				node.tssManager.SetKeyGenVersion(metadata.JoinPartyVersion)
			}
		}

		node.state.SetIsAffiliateProgram(metadata.AffiliateProgram)
	}
}

// Utils

// Some signed events will not work anymore so we need to resign them
func (node *Node) resetSignatures() {
	// 1. Unwrap events that were signed but not send
	if requests, err := node.networksManager.GetUnsentSignedUnwrapRequests(); err != nil {
		node.logger.Debug(err)
	} else {
		for _, request := range requests {
			if err := node.networksManager.SetUnsentUnwrapRequestAsUnsigned(*request); err != nil {
				node.logger.Debugf("Error: %s for event: %s", err.Error(), request.TransactionHash.String())
				continue
			}
		}
	}

	// 2. Wrap requests that were signed but the signature was not sent
	if requests, err := node.networksManager.GetUnsentSignedWrapRequests(); err != nil {
		node.logger.Debug(err)
	} else {
		for _, request := range requests {
			if err := node.networksManager.SetWrapEventSignature(request.Id, ""); err != nil {
				node.logger.Debugf("Error: %s for event: %s", err.Error(), request.Id.String())
				continue
			}
		}
	}

	// 3. Signed wrap request that were not redeemed even once
	if requests, err := node.networksManager.Znn().GetUnredeemedWrapRequests(); err != nil {
		node.logger.Debug(err)
	} else {
		for _, request := range requests {
			if err := node.networksManager.SetWrapEventSignature(request.Id, ""); err != nil {
				node.logger.Debugf("Error: %s for event: %s", err.Error(), request.Id.String())
				continue
			}
		}
	}
}

func (node *Node) signMessages(messagesToSign [][]byte, msgsIndexes map[string]int) (*keysign.Response, error) {
	start := time.Now()
	response, err := node.tssManager.BulkSign(messagesToSign, messages.ECDSAKEYSIGN)
	if err != nil {
		return nil, err
	}
	elapsed := time.Since(start)
	node.logger.Infof("bulkSign took %f", elapsed.Seconds())
	sort.Slice(response.Signatures, func(i, j int) bool {
		return msgsIndexes[response.Signatures[i].Msg] < msgsIndexes[response.Signatures[j].Msg]
	})
	return response, err
}

func (node *Node) validateSignatures(response *keysign.Response, pubKey string, messagesToSign [][]byte) (string, [][]byte, error) {
	var znnSignature string
	evmFullSignatures := make([][]byte, 0)
	for idx, sig := range response.Signatures {
		signature, err := base64.StdEncoding.DecodeString(sig.Signature)
		if err != nil {
			return "", nil, err
		}

		recoverID, err := base64.StdEncoding.DecodeString(sig.RecoveryID)
		if err != nil {
			return "", nil, err
		}
		fullSignatureBytes := append(signature, recoverID...)
		fullSignature := base64.StdEncoding.EncodeToString(fullSignatureBytes)

		ok, err := implementation.CheckECDSASignature(messagesToSign[idx], pubKey, fullSignature)
		if err != nil {
			return "", nil, err
		} else if ok == false {
			return "", nil, errors.New("invalid signature when checking ecdsa signature")
		}

		// znn sig
		if idx == 0 {
			znnSignature = base64.StdEncoding.EncodeToString(fullSignatureBytes)
			continue
		}

		// add 27 for evm
		fullSignatureBytes[len(fullSignatureBytes)-1] += 27
		evmFullSignatures = append(evmFullSignatures, fullSignatureBytes)
	}
	return znnSignature, evmFullSignatures, nil
}

func (node *Node) removeParticipant(pubKey string) {
	for idx, localPubKey := range node.config.TssConfig.LocalPubKeys {
		if localPubKey == pubKey {
			node.config.TssConfig.LocalPubKeys = append(node.config.TssConfig.LocalPubKeys[:idx], node.config.TssConfig.LocalPubKeys[idx+1:]...)
		}
	}
	delete(node.config.TssConfig.PubKeyWhitelist, pubKey)
	node.tssManager.DeleteLocalPubKey(pubKey)
	node.tssManager.DeleteWhitelistEntry(pubKey)
}

func (node *Node) allowNewParticipants(pubKeysMap map[string]string) {
	localPubKeysMap := make(map[string]bool)
	for _, pubKey := range node.config.TssConfig.LocalPubKeys {
		localPubKeysMap[pubKey] = true
	}

	// we add the new producing pub keys to the whitelist
	p2pWhitelist, partyWhitelist, sigWhitelist := node.tssManager.GetWhitelists()
	p2pWhitelist = make(map[string]bool)
	partyWhitelist = make(map[string]bool)
	sigWhitelist = make(map[string]bool)
	for k, v := range pubKeysMap {
		// eddsa pubKey
		localPubKeysMap[k] = true
		// peer id
		p2pWhitelist[v] = true
		partyWhitelist[v] = true
		sigWhitelist[v] = true
		// peer id
		node.config.TssConfig.PubKeyWhitelist[v] = true
	}

	// reset them
	node.config.TssConfig.LocalPubKeys = make([]string, 0)
	for k, _ := range localPubKeysMap {
		node.config.TssConfig.LocalPubKeys = append(node.config.TssConfig.LocalPubKeys, k)
	}
	sort.Strings(node.config.TssConfig.LocalPubKeys)
	node.tssManager.SetNewLocalPubKeys(node.config.TssConfig.LocalPubKeys)
	node.logger.Info("node.config.TssConfig.LocalPubKeys: ", node.config.TssConfig.LocalPubKeys)
	node.logger.Info("node.config.TssConfig.WhiteList: ", node.config.TssConfig.PubKeyWhitelist)
	node.logger.Info("node.config.TssConfig.LocalPubKeys len: ", len(node.config.TssConfig.LocalPubKeys))
	node.logger.Info("node.tssManager.GetLocalPubKeys(): ", node.tssManager.GetLocalPubKeys())
}

func (node *Node) openDataDir() error {
	if node.config.DataPath == "" {
		return nil
	}

	if err := os.MkdirAll(node.config.DataPath, 0700); err != nil {
		return err
	}
	node.logger.Info("successfully ensured DataPath exists", zap.String("data-path", node.config.DataPath))

	// Lock the instance directory to prevent concurrent use by another instance as well as
	// accidental use of the instance directory as a database.
	if fileLock, _, err := fileutil.Flock(filepath.Join(node.config.DataPath, ".lock")); err != nil {
		node.logger.Info("unable to acquire file-lock", zap.String("reason", err.Error()))
		return convertFileLockError(err)
	} else {
		node.dataDirLock = fileLock
	}

	node.logger.Info("successfully locked dataDir")
	return nil
}

func (node *Node) closeDataDir() {
	node.logger.Info("releasing dataDir lock ... ")
	// Release instance directory lock.
	if node.dataDirLock != nil {
		if err := node.dataDirLock.Release(); err != nil {
			node.logger.Error("can't release dataDir lock", zap.String("reason", err.Error()))
		}
		node.dataDirLock = nil
	}
}

func (node *Node) configurePubKey() error {
	node.config.TssConfig.LocalPubKeys = make([]string, 0)
	node.config.TssConfig.PubKeyWhitelist = make(map[string]bool)
	bridgeInfo, err := node.networksManager.GetBridgeInfo()
	if err != nil {
		return err
	}

	// just try to load the path of the pub key coming from the rpc
	pubKeyPath, err := common.GetPublicKeyFilePath(node.config.TssConfig.BaseDir, bridgeInfo.CompressedTssECDSAPubKey)
	if err != nil {
		return err
	}
	if _, err := os.Stat(pubKeyPath); os.IsNotExist(err) {
		// 1. we had a hard reset, in this case we can no longer sign be cause we don't have the pubKey files
		// 2. we have not participated yet in a keyGen ceremony
		// 3. there was no key gen (bootstrap), in this case participatingPubKeys will be empty
		// pub key file does not exist, we can only wait for a new key gen
		node.logger.Info(err.Error())
		node.config.TssConfig.PublicKey = ""
		node.config.TssConfig.DecompressedPublicKey = ""
	} else {
		node.config.TssConfig.PublicKey = bridgeInfo.CompressedTssECDSAPubKey
		node.config.TssConfig.DecompressedPublicKey = bridgeInfo.DecompressedTssECDSAPubKey
		jsonFile, err := os.Open(pubKeyPath)
		if err != nil {
			return err
		}

		bytesValue, _ := io.ReadAll(jsonFile)
		if err := jsonFile.Close(); err != nil {
			return err
		}

		var result map[string]interface{}
		if err := json.Unmarshal(bytesValue, &result); err != nil {
			return err
		}

		for _, publicKey := range result["participant_keys"].([]interface{}) {
			node.logger.Info(publicKey.(string))
			publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKey.(string))
			if err != nil {
				return err
			}
			pub, err := ic.UnmarshalEd25519PublicKey(publicKeyBytes)
			if err != nil {
				return err
			}
			peerId, err := peer.IDFromPublicKey(pub)
			if err != nil {
				return err
			}

			node.config.TssConfig.PubKeyWhitelist[peerId.String()] = true
			node.config.TssConfig.LocalPubKeys = append(node.config.TssConfig.LocalPubKeys, publicKey.(string))
		}
	}
	sort.Strings(node.config.TssConfig.LocalPubKeys)

	return nil
}

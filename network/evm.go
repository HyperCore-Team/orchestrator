package network

import (
	"crypto/ecdsa"
	ecommon "github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/joncrlsn/dque"
	"github.com/pkg/errors"
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/vm/constants"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
	"go.uber.org/zap"
	"orchestrator/common/config"
	"orchestrator/common/storage"
	"orchestrator/db"
	"orchestrator/db/manager"
	"orchestrator/rpc"
	"os"
	"reflect"
	"strings"
	"syscall"

	"math/big"
	"orchestrator/common"
	"orchestrator/common/events"
	"time"
)

type evmNetwork struct {
	config.EvmParams
	unconfirmedQueue *dque.DQue
	dbManager        *manager.Manager
	rpcManager       *rpc.Manager
	UrlsInfo         config.UrlsInfo
	state            *common.GlobalState
	stopChan         chan os.Signal
	logger           *zap.SugaredLogger
}

func NewEvmNetwork(network *definition.NetworkInfo, dbManager *manager.Manager, rpcManager *rpc.Manager, state *common.GlobalState, stop chan os.Signal) (*evmNetwork, error) {
	newConfig, err := config.NewEvmParams(network)
	if err != nil {
		return nil, err
	}

	newQueue, err := storage.CreateOrOpenQueue(network.NetworkClass, network.Name)
	if err != nil {
		return nil, err
	}
	dbManager.AddEvmEventStore(network.Id, network.Name, newConfig.ContractDeploymentHeight())

	newLogger, errLog := common.CreateSugarLogger()
	if errLog != nil {
		return nil, errLog
	}

	newEvmNetwork := &evmNetwork{
		EvmParams:        newConfig,
		dbManager:        dbManager,
		rpcManager:       rpcManager,
		unconfirmedQueue: newQueue,
		state:            state,
		stopChan:         stop,
		logger:           newLogger,
	}

	return newEvmNetwork, nil
}

func (eN *evmNetwork) Start() error {
	synced := eN.rpcManager.Evm(eN.ChainId()).IsSynced()
	for synced == false {
		eN.logger.Infof("evm node for network %d is not synced, will wait for it to sync\n", eN.ChainId())
		time.Sleep(10 * time.Second)
		synced = eN.rpcManager.Evm(eN.ChainId()).IsSynced()
	}
	if err := eN.FillEvmParamsRpc(); err != nil {
		eN.logger.Debug(err)
		return err
	}
	lastUpdateHeight, err := eN.dbManager.EvmStorage(eN.ChainId()).GetLastUpdateHeight()
	if err != nil {
		eN.logger.Debug(err)
		return err
	}
	if lastUpdateHeight == 0 {
		if err := eN.dbManager.EvmStorage(eN.ChainId()).SetLastUpdateHeight(eN.ContractDeploymentHeight()); err != nil {
			eN.logger.Debug(err)
			return err
		}
	}
	if err := eN.Sync(); err != nil {
		eN.logger.Debug(err)
		return err
	}

	go eN.SubscribeToEvents()
	// todo defer unsub?

	go eN.ProcessEvents()
	return nil
}

func (eN *evmNetwork) eventsStore() db.EvmStorage {
	return eN.dbManager.EvmStorage(eN.ChainId())
}

func (eN *evmNetwork) EvmRpc() *rpc.EvmRpc {
	return eN.rpcManager.Evm(eN.ChainId())
}

func (eN *evmNetwork) Sync() error {
	eN.logger.Info("In sync evm")
	if updateHeight, err := eN.eventsStore().GetLastUpdateHeight(); err != nil {
		return err
	} else {
		eN.logger.Info("updateHeight: ", updateHeight)
		eN.logger.Debugf("confToFin: %d", eN.ConfirmationsToFinality())
		for {
			latestBlock, rpcErr := eN.EvmRpc().BlockNumber()
			if rpcErr != nil {
				eN.logger.Debug(rpcErr)
				return rpcErr
			}

			if updateHeight > latestBlock {
				// todo this can happen if we have a fork
				return errors.Errorf("sync evm problem for network: %s, chainId: %d", eN.NetworkName(), eN.ChainId())
			}

			end := false
			filterQuerySize := eN.rpcManager.Evm(eN.ChainId()).FilterQuerySize()
			eN.logger.Debugf("filterQuerySize: %d", filterQuerySize)
			distance := latestBlock - updateHeight
			if distance < eN.ConfirmationsToFinality() {
				filterQuerySize = distance
				end = true
			} else if distance < filterQuerySize {
				filterQuerySize = distance
			}
			//eN.logger.Infof("distange: %d, left: %d, right: %d, filterQuerySize: %d\n", distance, updateHeight, updateHeight+filterQuerySize, filterQuerySize)

			if logs, err := eN.EvmRpc().FilterLogs(updateHeight, updateHeight+filterQuerySize); err != nil {
				return err
			} else {
				for _, log := range logs {
					eN.logger.Info("log: ", log)
					// if we have confirmations then we are live, otherwise we are not
					if err := eN.InterpretLog(log, latestBlock-log.BlockNumber < eN.ConfirmationsToFinality()); err != nil {
						eN.logger.Error(err)
						return err
					}
				}
			}

			updateHeight += filterQuerySize
			if err := eN.eventsStore().SetLastUpdateHeight(updateHeight); err != nil {
				return err
			}
			if end {
				break
			}
		}
	}
	return nil
}

func (eN *evmNetwork) InterpretLog(log etypes.Log, live bool) error {
	eN.logger.Infof("InterpretLog - tx: %s and log topic: %s - live: %v", log.TxHash.String(), log.Topics[0].Hex(), live)

	switch log.Topics[0].Hex() {
	case common.UnwrapSigHash.Hex():
		unwrapped, errParse := eN.EvmRpc().Bridge().ParseUnwrapped(log)
		if errParse != nil {
			eN.logger.Debug(errParse)
			return nil
		}
		eN.logger.Debug("Found unwrap event evm")
		// only process events that have a valid address
		if _, errParse = common.ParseAddressString(unwrapped.To, definition.NoMClass); errParse != nil {
			eN.logger.Debugf("Could not parse zenon address: %s with error: %s", unwrapped.To, errParse.Error())
			return nil
		}

		event := events.UnwrapRequestEvm{
			NetworkClass:    eN.NetworkClass(),
			ChainId:         eN.ChainId(),
			BlockNumber:     log.BlockNumber,
			BlockHash:       log.BlockHash,
			TransactionHash: log.TxHash,
			LogIndex:        uint32(log.Index),
			From:            unwrapped.From,
			To:              unwrapped.To,
			Token:           unwrapped.Token,
			Amount:          big.NewInt(0).Set(unwrapped.Amount),
			Signature:       "",
			RedeemStatus:    common.UnredeemedStatus,
		}
		// todo refactor all logs on a single line
		eN.logger.Infof("chainId: %d", event.ChainId)
		eN.logger.Infof("txHash: %s", event.TransactionHash.String())
		eN.logger.Infof("logIndex: %d", event.LogIndex)
		eN.logger.Infof("To: %s", event.To)
		eN.logger.Infof("From: %s", event.From.String())
		eN.logger.Infof("Token: %s", event.Token.String())
		eN.logger.Infof("Amount: %d", event.Amount.Uint64())

		// we enqueue the event only if we are live or we don't have confirmations
		if live {
			err := eN.unconfirmedQueue.Enqueue(&event)
			if err != nil {
				eN.logger.Error(err)
				return err
			}
		} else {
			if localEvent, dbErr := eN.eventsStore().GetUnwrapRequestByHashAndLog(event.TransactionHash, event.LogIndex); dbErr != nil {
				return dbErr
			} else if localEvent == nil {
				if err := eN.eventsStore().AddUnwrapRequest(event); err != nil {
					return err
				}
			} else {
				// the event was added by the znn sync, we update the block number
				if err := eN.eventsStore().UpdateUnwrapRequestBlockNumber(event); err != nil {
					return err
				}
			}
		}
	case common.RegisteredRedeemSigHash.Hex():
		id := types.Hash(log.Topics[1])
		eN.logger.Infof("found RegisteredRedeemSigHash: %s", id.String())

		if rpcEvent, rpcErr := eN.rpcManager.Znn().GetWrapTokenRequestById(id); rpcErr != nil {
			eN.logger.Debugf("call: eN.rpcManager.Znn().GetWrapTokenRequestById(id) error: %s", rpcErr.Error())
			if rpcErr == constants.ErrDataNonExistent {
				if live {
					if stateErr := eN.state.SetState(common.EmergencyState); stateErr != nil {
						eN.logger.Info("sent sigkill from here 4")
						eN.stopChan <- syscall.SIGKILL
						return stateErr
					}
				}
			}
			return rpcErr
		} else if rpcEvent == nil {
			eN.logger.Info("wrap event not found for register redeem id: ", id.String())
			if live {
				if stateErr := eN.state.SetState(common.EmergencyState); stateErr != nil {
					eN.logger.Info("sent sigkill from here 5")
					eN.stopChan <- syscall.SIGKILL
					return stateErr
				}
			}
		} else {
			if registeredRedeem, parseErr := eN.rpcManager.Evm(eN.ChainId()).Bridge().ParseRegisteredRedeem(log); parseErr != nil {
				eN.logger.Debug(parseErr)
			} else {
				deductedFeeAmount := big.NewInt(0).Set(rpcEvent.Amount)
				deductedFeeAmount.Sub(deductedFeeAmount, rpcEvent.Fee)

				// We have to check every field here because one that has control over tss can create a redeem and spoof the id, token, amount or destination
				if deductedFeeAmount.Cmp(registeredRedeem.Amount) != 0 || rpcEvent.ToAddress != strings.ToLower(registeredRedeem.To.String()) ||
					rpcEvent.TokenAddress != strings.ToLower(registeredRedeem.Token.String()) {
					if live {
						if stateErr := eN.state.SetState(common.EmergencyState); stateErr != nil {
							eN.logger.Info("sent sigkill from here 8")
							eN.stopChan <- syscall.SIGKILL
							return stateErr
						}
					}
				} else {
					if event, storageErr := eN.dbManager.ZnnStorage().GetWrapRequestById(id); storageErr != nil {
						return storageErr
					} else if event == nil {
						if storageErr = eN.dbManager.ZnnStorage().AddWrapRequest(common.ZnnWrapToOrchestratorWrap(rpcEvent)); storageErr != nil {
							return storageErr
						}
					}
					if err := eN.dbManager.ZnnStorage().SetWrapRequestStatus(id, common.PendingRedeemStatus); err != nil {
						return err
					}
				}
			}
		}
	case common.RedeemedSigHash.Hex():
		redeem, errParse := eN.EvmRpc().Bridge().ParseRedeemed(log)
		if errParse != nil {
			return errParse
		}
		id, err := types.BytesToHash(redeem.Nonce.Bytes())
		if err != nil {
			return err
		}
		if event, storageErr := eN.dbManager.ZnnStorage().GetWrapRequestById(id); storageErr != nil {
			return storageErr
		} else if event == nil {
			// todo check if event exists on znn otherwise in this case there is not much we can do
		} else {
			if err = eN.dbManager.ZnnStorage().SetWrapRequestStatus(id, common.RedeemedStatus); err != nil {
				return err
			}
		}
	case common.RevokedRedeemSigHash.Hex():
		redeem, errParse := eN.EvmRpc().Bridge().ParseRedeemed(log)
		if errParse != nil {
			return errParse
		}
		id, err := types.BytesToHash(redeem.Nonce.Bytes())
		if err != nil {
			return err
		}
		if event, storageErr := eN.dbManager.ZnnStorage().GetWrapRequestById(id); storageErr != nil {
			return storageErr
		} else if event == nil {
			// todo check if event exists on znn otherwise in this case there is not much we can do
		} else {
			if err = eN.dbManager.ZnnStorage().SetWrapRequestStatus(id, common.RevokedStatus); err != nil {
				return err
			}
		}
	case common.HaltedSigHash.Hex():
		if live {
			if err := eN.state.SetState(common.HaltedState); err != nil {
				eN.logger.Debug(err)
				eN.stopChan <- syscall.SIGKILL
				return err
			}
		}
	}
	if err := eN.eventsStore().SetLastUpdateHeight(log.BlockNumber); err != nil {
		return err
	}
	eN.logger.Infof("Set last blockNumber as: %d", log.BlockNumber)

	return nil
}

// FillEvmParamsRpc this should be called after we create the network
func (eN *evmNetwork) FillEvmParamsRpc() error {
	if confirmations, err := eN.ConfirmationsToFinalityRpc(); err != nil {
		return err
	} else {
		eN.SetConfirmationsToFinality(confirmations)
	}
	if height, err := eN.ContractDeploymentHeightRpc(); err != nil {
		return err
	} else {
		eN.SetContractDeploymentHeight(height)
	}
	if blockTime, err := eN.EstimatedBlockTimeRpc(); err != nil {
		return err
	} else {
		eN.SetEstimatedBlockTime(blockTime)
	}

	eN.Log()
	return nil
}

func (eN *evmNetwork) SubscribeToEvents() {
	eN.logger.Infof("SubscribeToEvents for network with chainId: %d", eN.ChainId())
	logSub, logChan, err := eN.EvmRpc().SubscribeToLogs()
	if err != nil {
		eN.logger.Error(err)
		eN.stopChan <- syscall.SIGKILL
		return
	}

	go func() {
		for {
			time.Sleep(3 * time.Minute)
			if logSub != nil {
				logSub.Unsubscribe()
			}
			errSync := eN.Sync()
			if errSync != nil {
				eN.logger.Debug(errSync)
				continue
			}
			if logSub, logChan, err = eN.EvmRpc().SubscribeToLogs(); err != nil {
				eN.logger.Error(err)
				eN.stopChan <- syscall.SIGKILL
			}
		}
	}()

	for {
		select {
		case subErr := <-logSub.Err():
			if subErr != nil {
				eN.logger.Error(subErr)
				eN.stopChan <- syscall.SIGKILL
			}
		case newLog := <-logChan:
			if errInterpret := eN.InterpretLog(newLog, true); errInterpret != nil {
				eN.logger.Debug(errInterpret)
			}
		}
	}
}

func (eN *evmNetwork) ProcessEvents() {
	// this means we should dequeue
	dequeue := false
	// this means we dequeued an item
	dequeued := false
	for {
		var peekedEvent *events.UnwrapRequestEvm
		var peekedInterface interface{}
		var errQueue error
		if dequeue {
			peekedInterface, errQueue = eN.unconfirmedQueue.DequeueBlock()
			if errQueue != nil {
				eN.logger.Error(errQueue)
				eN.stopChan <- syscall.SIGKILL
				return
			}
			dequeue = false
			dequeued = true
		} else {
			peekedInterface, errQueue = eN.unconfirmedQueue.PeekBlock()
			if errQueue != nil {
				eN.logger.Error(errQueue)
				eN.stopChan <- syscall.SIGKILL
				return
			}
			dequeued = false
		}
		var y bool
		peekedEvent, y = peekedInterface.(*events.UnwrapRequestEvm)
		if !y {
			eN.logger.Info("Dequeued object is not an Item pointer")
			if dequeued == false {
				dequeue = true
			}
			continue
		}

		var txReceipt *etypes.Receipt
		eN.logger.Debugf("processing evm event with tx hash %s and logIndex: %d", peekedEvent.TransactionHash.String(), peekedEvent.LogIndex)
		txReceipt, err := eN.EvmRpc().TransactionReceipt(peekedEvent.TransactionHash)
		if err != nil {
			eN.logger.Debug(err)
			if dequeued == false {
				dequeue = true
			}
			continue
		} else if txReceipt.Status != etypes.ReceiptStatusSuccessful {
			eN.logger.Infof("txReceipt for tx hash %s not successful\n", txReceipt.TxHash.String())
			if dequeued == false {
				dequeue = true
			}
			continue
		}

		for {
			time.Sleep(2 * time.Second)
			currentBlockHeight, err := eN.EvmRpc().BlockNumber()
			if err != nil {
				eN.logger.Debug(err)
				continue
			}
			if currentBlockHeight < txReceipt.BlockNumber.Uint64() {
				eN.logger.Errorf("blockNumber on evm with chain id: %d is less than the transaction block number, we are probably still syncing", eN.ChainId())
				// we stop the binary so it restarts and wait for the node to sync
				eN.stopChan <- syscall.SIGKILL
				return
			}

			confirmations := currentBlockHeight - txReceipt.BlockNumber.Uint64()
			if confirmations < eN.EvmParams.ConfirmationsToFinality() {
				// we need to wait confirmationsRequired blocks * estimated time per block
				confirmationsRequired := eN.EvmParams.ConfirmationsToFinality() - confirmations
				timeToWait := time.Duration(confirmationsRequired) * eN.EvmParams.EstimatedBlockTime()
				time.Sleep(timeToWait)
				continue
			}
			break
		}

		txReceipt, err = eN.EvmRpc().TransactionReceipt(peekedEvent.TransactionHash)
		if err != nil {
			eN.logger.Debug(err)
			if dequeued == false {
				dequeue = true
			}
			continue
		} else if txReceipt.Status != etypes.ReceiptStatusSuccessful {
			eN.logger.Infof("txReceipt for tx hash %s not successful\n", txReceipt.TxHash.String())
			if dequeued == false {
				dequeue = true
			}
			continue
		}

		if !reflect.DeepEqual(txReceipt.BlockHash.Bytes(), peekedEvent.BlockHash.Bytes()) {
			eN.logger.Info("Transaction %s has a different block hash %s, expected %s", peekedEvent.TransactionHash.String(), txReceipt.BlockHash.String(), peekedEvent.BlockHash.String())
			if dequeued == false {
				dequeue = true
			}
			continue
		}

		// We double-check that this transaction exists
		tx, _, err := eN.EvmRpc().TransactionByHash(peekedEvent.TransactionHash)
		if err != nil {
			eN.logger.Debug(err)
			if dequeued == false {
				dequeue = true
			}
			continue
		} else if tx == nil {
			eN.logger.Infof("Transaction %s does not exist or has not executed successfully", peekedEvent.TransactionHash.String())
			if dequeued == false {
				dequeue = true
			}
			continue
		}

		eN.logger.Infof("Event with hash: %s and logIndex: %d is confirmed", peekedEvent.TransactionHash, peekedEvent.LogIndex)

		if err := eN.eventsStore().AddUnwrapRequest(*peekedEvent); err != nil {
			return
		}
		eN.logger.Infof("Added event hash: %s logIndex: %d to persistent storage", peekedEvent.TransactionHash.String(), peekedEvent.LogIndex)

		if !dequeued {
			_, errQueue = eN.unconfirmedQueue.DequeueBlock()
			if errQueue != nil {
				eN.logger.Error(errQueue)
				eN.stopChan <- syscall.SIGKILL
				return
			}
		}
	}
}

func (eN *evmNetwork) Stop() {
	eN.EvmRpc().Stop()
	_ = eN.unconfirmedQueue.Close()
}

func (eN *evmNetwork) SignTx(tx *etypes.Transaction, ecdsaPrivateKey *ecdsa.PrivateKey, chainId uint32) (*etypes.Transaction, error) {
	signer := etypes.LatestSignerForChainID(big.NewInt(int64(chainId)))
	return etypes.SignTx(tx, signer, ecdsaPrivateKey)
}

// Local storage

func (eN *evmNetwork) GetUnsignedUnwrapRequests() ([]*events.UnwrapRequestEvm, error) {
	return eN.eventsStore().GetUnsignedUnwrapRequests()
}

func (eN *evmNetwork) AddUnwrapRequest(event events.UnwrapRequestEvm) error {
	return eN.eventsStore().AddUnwrapRequest(event)
}

func (eN *evmNetwork) GetUnwrapRequestByHashAndLog(txHash ecommon.Hash, logIndex uint32) (*events.UnwrapRequestEvm, error) {
	return eN.eventsStore().GetUnwrapRequestByHashAndLog(txHash, logIndex)
}

func (eN *evmNetwork) SetUnwrapRequestStatus(txHash ecommon.Hash, logIndex, status uint32) error {
	return eN.eventsStore().SetUnwrapRequestStatus(txHash, logIndex, status)
}

func (eN *evmNetwork) SetUnwrapRequestSignature(txHash ecommon.Hash, logIndex uint32, signature string) error {
	return eN.eventsStore().SetUnwrapRequestSignature(txHash, logIndex, signature)
}

func (eN *evmNetwork) SetUnsentUnwrapRequestAsUnsigned(txHash ecommon.Hash, logIndex uint32) error {
	return eN.eventsStore().SetUnsentUnwrapRequestAsUnsigned(txHash, logIndex)
}

func (eN *evmNetwork) GetUnwrapRequestsByStatus(status uint32) ([]*events.UnwrapRequestEvm, error) {
	return eN.eventsStore().GetUnwrapRequestsByStatus(status)
}

func (eN *evmNetwork) GetUnsentSignedUnwrapRequests() ([]*events.UnwrapRequestEvm, error) {
	return eN.eventsStore().GetUnsentSignedUnwrapRequests()
}

///////// RPC

func (eN *evmNetwork) SendTransaction(tx *etypes.Transaction) error {
	return eN.EvmRpc().SendTransaction(tx)
}

func (eN *evmNetwork) GetCurrentTss() (ecommon.Address, error) {
	return eN.EvmRpc().GetCurrentTss()
}

func (eN *evmNetwork) GetSetTssEcdsaPubKeyEvmMessage(newAddress ecommon.Address) ([]byte, error) {
	return eN.EvmRpc().GetSetTssEcdsaPubKeyEvmMessage(newAddress, eN.NetworkClass(), eN.ChainId(), eN.ContractAddress())
}

func (eN *evmNetwork) GetSetTssEcdsaPubKeyEvmTx(newTss, sender ecommon.Address, oldFullSignature, newFullSignature []byte) (*etypes.Transaction, error) {
	return eN.EvmRpc().GetSetTssEcdsaPubKeyEvmTx(newTss, sender, oldFullSignature, newFullSignature, eN.ContractAddress())
}

func (eN *evmNetwork) GetHaltEvmMessage() ([]byte, error) {
	return eN.EvmRpc().GetHaltEvmMessage(definition.EvmClass, eN.ChainId(), eN.ContractAddress())
}

func (eN *evmNetwork) GetHaltEvmTx(signature []byte, sender ecommon.Address) (*etypes.Transaction, error) {
	return eN.EvmRpc().GetHaltTxEvm(sender, signature, eN.ContractAddress())
}

func (eN *evmNetwork) IsHalted() (bool, error) {
	return eN.EvmRpc().IsHalted()
}

func (eN *evmNetwork) ContractDeploymentHeightRpc() (uint64, error) {
	return eN.EvmRpc().ContractDeploymentHeight()
}

func (eN *evmNetwork) EstimatedBlockTimeRpc() (uint64, error) {
	return eN.EvmRpc().EstimatedBlockTime()
}

func (eN *evmNetwork) ConfirmationsToFinalityRpc() (uint64, error) {
	return eN.EvmRpc().ConfirmationsToFinality()
}

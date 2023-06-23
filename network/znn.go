package network

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	ecommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/zenon-network/go-zenon/chain/nom"
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/protocol"
	"github.com/zenon-network/go-zenon/rpc/api"
	"github.com/zenon-network/go-zenon/rpc/api/embedded"
	"github.com/zenon-network/go-zenon/vm/constants"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
	"github.com/zenon-network/go-zenon/wallet"
	"go.uber.org/zap"
	"math/big"
	"orchestrator/common"
	"orchestrator/common/config"
	"orchestrator/common/events"
	"orchestrator/db"
	"orchestrator/db/manager"
	"orchestrator/rpc"
	"os"
	"strings"
	"syscall"
	"time"
)

type znnNetwork struct {
	config.ZnnParams
	dbManager         *manager.Manager
	rpcManager        *rpc.Manager
	networkManager    *NetworksManager
	networksInfo      map[string]config.BaseNetworkConfig
	state             *common.GlobalState
	stopChan          chan os.Signal
	setBridgeMetadata func(metadata *common.BridgeMetadata)
	logger            *zap.SugaredLogger
	administrator     string
}

// CheckOrchestratorInfoInitialized this method should have the same checks as in go-zenon
func CheckOrchestratorInfoInitialized(orchestratorInfo *definition.OrchestratorInfo) error {
	if orchestratorInfo == nil {
		return errors.New("OrchestratorInfo not initialised 1")
	} else if orchestratorInfo.WindowSize == 0 || orchestratorInfo.KeyGenThreshold == 0 || orchestratorInfo.ConfirmationsToFinality == 0 || orchestratorInfo.EstimatedMomentumTime == 0 {
		return errors.New("OrchestratorInfo not initialised 2")
	}
	return nil
}

// CheckSecurityInfoInitialized this method should have the same checks as in go-zenon
func CheckSecurityInfoInitialized(securityInfo *definition.SecurityInfoVariable) error {
	if len(securityInfo.Guardians) < constants.MinGuardians {
		return errors.New("SecurityInfo not initialised")
	}
	return nil
}

func NewZnnNetwork(rpcManager *rpc.Manager, dbManager *manager.Manager, networkManager *NetworksManager, state *common.GlobalState, networksInfo map[string]config.BaseNetworkConfig, stopChan chan os.Signal, setBridgeMetadata func(metadata *common.BridgeMetadata)) (*znnNetwork, error) {
	bridgeInfo, err := rpcManager.Znn().GetBridgeInfo()
	if err != nil {
		return nil, err
	} else if bridgeInfo == nil {
		return nil, errors.New("BridgeInfo not initialised")
	}

	orchestratorInfo, err := rpcManager.Znn().GetOrchestratorInfo()
	if err != nil {
		return nil, err
	} else if bridgeErr := CheckOrchestratorInfoInitialized(orchestratorInfo); bridgeErr != nil {
		return nil, bridgeErr
	}

	securityInfo, err := rpcManager.Znn().GetSecurityInfo()
	if err != nil {
		return nil, err
	} else if bridgeErr := CheckSecurityInfoInitialized(securityInfo); bridgeErr != nil {
		return nil, bridgeErr
	}

	newZnnParams, err := config.NewZnnParams(orchestratorInfo)
	if err != nil {
		return nil, err
	}

	newLogger, errLog := common.CreateSugarLogger()
	if errLog != nil {
		return nil, errLog
	}

	newZnnNetwork := &znnNetwork{
		ZnnParams:         *newZnnParams,
		rpcManager:        rpcManager,
		dbManager:         dbManager,
		networkManager:    networkManager,
		networksInfo:      networksInfo,
		state:             state,
		setBridgeMetadata: setBridgeMetadata,
		stopChan:          stopChan,
		logger:            newLogger,
		administrator:     bridgeInfo.Administrator.String(),
	}
	return newZnnNetwork, nil
}

/// Utils

func (rC *znnNetwork) Start() error {
	go rC.ListenForMomentumHeight()

	if err := rC.Sync(); err != nil {
		rC.logger.Debugf("error: %s", err.Error())
		return err
	}

	go rC.ListenForEmbeddedBridgeAccountBlocks()
	return nil
}

func (rC *znnNetwork) Stop() error {
	return rC.ZnnRpc().Stop()
}
func (rC *znnNetwork) eventsStore() db.ZnnStorage {
	return rC.dbManager.ZnnStorage()
}
func (rC *znnNetwork) ZnnRpc() *rpc.ZnnRpc {
	return rC.rpcManager.Znn()
}

func (rC *znnNetwork) Sync() error {
	rC.logger.Debug("In sync znn")
	if accountBlockHeight, err := rC.eventsStore().GetLastUpdateHeight(); err != nil {
		return err
	} else {
		rC.logger.Debugf("last account block update height: %d", accountBlockHeight)
		accountBlockList, errRpc := rC.ZnnRpc().GetAccountBlocksByHeight(types.BridgeContract, accountBlockHeight+1, 30)
		if errRpc != nil {
			return errRpc
		}
		for len(accountBlockList.List) > 0 {
			for _, accBlock := range accountBlockList.List {
				if accBlock.BlockType == nom.BlockTypeContractReceive {
					rC.logger.Debug("found receive block")
					hash := accBlock.Hash
					for {
						rC.logger.Debugf("confDetail is nil: %v for %s", accBlock.ConfirmationDetail == nil, hash.String())
						accBlock, errRpc = rC.ZnnRpc().GetAccountBlockByHash(hash)
						if errRpc != nil {
							rC.logger.Debug(err)
						} else if accBlock == nil {

						} else if accBlock.ConfirmationDetail != nil {
							break
						}
						time.Sleep(5 * time.Second)
						continue
					}

					if sendBlock, errRpc := rC.ZnnRpc().GetAccountBlockByHash(accBlock.FromBlockHash); err != nil {
						return errRpc
					} else if sendBlock == nil {
						return errors.Errorf("Send block %s for associated receive %s is non existent", accBlock.Hash.String(), accBlock.FromBlockHash.String())
					} else {
						var live bool
						frMomHeight, errFrMom := rC.state.GetFrontierMomentum()
						if errFrMom != nil {
							return errFrMom
						}
						if frMomHeight < accBlock.ConfirmationDetail.MomentumHeight {
							return errors.New(fmt.Sprintf("frMomHeight %d cannot be less than the height of the momentum %d in which was included the acc block we process", frMomHeight, accBlock.ConfirmationDetail.MomentumHeight))
						}
						live = (frMomHeight - accBlock.ConfirmationDetail.MomentumHeight) < uint64(rC.ConfirmationsToFinality())
						live = live && rC.IsSynced()
						if newErr := rC.InterpretSendBlockData(sendBlock, live, accBlock.Height); newErr != nil {
							return newErr
						}
					}
				}
			}
			accountBlockHeight += uint64(len(accountBlockList.List))
			accountBlockList, err = rC.ZnnRpc().GetAccountBlocksByHeight(types.BridgeContract, accountBlockHeight+1, 30)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// InterpretSendBlockData We assume that if the receive acc block was created then it was no error
func (rC *znnNetwork) InterpretSendBlockData(sendBlock *api.AccountBlock, live bool, receiveBlockHeight uint64) error {
	rC.logger.Debugf("InterpretSendBlockData, live: %v", live)
	methodSig := base64.StdEncoding.EncodeToString(sendBlock.Data[0:4])
	switch methodSig {
	case base64.StdEncoding.EncodeToString(definition.ABIBridge.Methods[definition.WrapTokenMethodName].Id()):
		rC.logger.Debug("found WrapTokenMethodName")
		param := new(definition.WrapTokenParam)
		if err := definition.ABIBridge.UnpackMethod(param, definition.WrapTokenMethodName, sendBlock.Data); err != nil {
			// todo ignore some of these errors
			return constants.ErrUnpackError
		}

		if request, err := rC.ZnnRpc().GetWrapTokenRequestById(sendBlock.Hash); err != nil {
			if err.Error() == constants.ErrDataNonExistent.Error() {
				rC.logger.Debug(constants.ErrDataNonExistent)
				return nil
			}
			return err
		} else if request == nil {
			rC.logger.Info("request non existent")
			return nil
		} else {
			if err = rC.AddWrapEvent(request); err != nil {
				return err
			}
		}
	case base64.StdEncoding.EncodeToString(definition.ABIBridge.Methods[definition.UpdateWrapRequestMethodName].Id()):
		rC.logger.Debug("found UpdateWrapRequestMethodName")
		param := new(definition.UpdateWrapRequestParam)
		if err := definition.ABIBridge.UnpackMethod(param, definition.UpdateWrapRequestMethodName, sendBlock.Data); err != nil {
			return constants.ErrUnpackError
		}

		if request, err := rC.ZnnRpc().GetWrapTokenRequestById(param.Id); err != nil {
			if err.Error() == constants.ErrDataNonExistent.Error() {
				rC.logger.Debug(constants.ErrDataNonExistent)
				return nil
			}
			return err
		} else if request == nil {
			rC.logger.Info("update request non existent")
			return nil
		} else {
			if localRequest, err := rC.eventsStore().GetWrapRequestById(param.Id); err != nil {
				return err
			} else if localRequest == nil {
				rC.logger.Info("request does not exist locally, will add it")
				if err := rC.AddWrapEvent(request); err != nil {
					return err
				}
			}
			if len(request.Signature) > 0 {
				if err = rC.eventsStore().SetWrapRequestSentSignature(request.Id); err != nil {
					return err
				}
			}
		}
	case base64.StdEncoding.EncodeToString(definition.ABIBridge.Methods[definition.RedeemUnwrapMethodName].Id()):
		rC.logger.Info("found RedeemUnwrapMethodName")
		param := new(definition.RedeemParam)
		if err := definition.ABIBridge.UnpackMethod(param, definition.RedeemUnwrapMethodName, sendBlock.Data); err != nil {
			return constants.ErrUnpackError
		}

		rC.logger.Debugf("redeem for tx: %s and log index: %d", param.TransactionHash.String(), param.LogIndex)
		if rpcEvent, rpcErr := rC.GetUnwrapTokenRequestByHashAndLog(param.TransactionHash, param.LogIndex); rpcErr != nil {
			if rpcErr.Error() == constants.ErrDataNonExistent.Error() {
				rC.logger.Info("there is a redeem attempt for a non existing unwrap event")
				rC.logger.Debug(rpcErr)
				return nil
			}
			return rpcErr
		} else if rpcEvent == nil {
			// someone is trying to redeem a non existent event
			rC.logger.Info("there is a redeem attempt for a non existing unwrap event")
			return nil
		} else {
			if localEvent, err := rC.dbManager.EvmStorage(rpcEvent.ChainId).GetUnwrapRequestByHashAndLog(ecommon.Hash(rpcEvent.TransactionHash), rpcEvent.LogIndex); err != nil {
				return err
			} else if localEvent == nil {
				if storageErr := rC.dbManager.EvmStorage(rpcEvent.ChainId).AddUnwrapRequest(common.ZnnUnwrapToOrchestatorUnwrap(rpcEvent)); storageErr != nil {
					return storageErr
				}
			}
			// if the event was redeemed we also set it locally
			if rpcEvent.Redeemed == 1 {
				if storageErr := rC.dbManager.EvmStorage(rpcEvent.ChainId).SetUnwrapRequestStatus(ecommon.Hash(rpcEvent.TransactionHash), rpcEvent.LogIndex, common.RedeemedStatus); storageErr != nil {
					return storageErr
				}
			} else {
				if storageErr := rC.dbManager.EvmStorage(rpcEvent.ChainId).SetUnwrapRequestStatus(ecommon.Hash(rpcEvent.TransactionHash), rpcEvent.LogIndex, common.PendingRedeemStatus); storageErr != nil {
					return storageErr
				}
			}
		}
	case base64.StdEncoding.EncodeToString(definition.ABIBridge.Methods[definition.UnwrapTokenMethodName].Id()):
		rC.logger.Debug("found UnwrapTokenMethodName")
		param := new(definition.UnwrapTokenParam)
		if err := definition.ABIBridge.UnpackMethod(param, definition.UnwrapTokenMethodName, sendBlock.Data); err != nil {
			return constants.ErrUnpackError
		}

		rC.logger.Infof("NetworkClass: %d, ChainId: %d, TransactionHash: %s, LogIndex: %d, ToAddress: %s, TokenAddress: %s, Amount: %s, Signature: %s",
			param.NetworkClass, param.ChainId, param.TransactionHash, param.LogIndex, param.ToAddress.String(), param.TokenAddress, param.Amount.String(), param.Signature)

		if rpcZnnEvent, rpcZnnErr := rC.GetUnwrapTokenRequestByHashAndLog(param.TransactionHash, param.LogIndex); rpcZnnErr != nil {
			if rpcZnnErr.Error() == constants.ErrDataNonExistent.Error() {
				rC.logger.Debug(constants.ErrDataNonExistent)
				return nil
			}
			rC.logger.Debugf("get for tx %s and log :%d rpc error: %s", param.TransactionHash.String(), param.LogIndex, rpcZnnErr.Error())
			return rpcZnnErr
		} else if rpcZnnEvent == nil {
			// We don't care if it the rpc does not return it, it means the tx returned an error
			rC.logger.Infof("unwrap event non existent: %s", param.TransactionHash.String())
			return nil
		} else {
			if tx, rpcEvmErr := rC.rpcManager.Evm(param.ChainId).TransactionReceipt(ecommon.Hash(param.TransactionHash)); rpcEvmErr != nil {
				rC.logger.Debugf("error: %s", rpcEvmErr.Error())
				// todo filter errors, maybe just an rpc error
				if rpcEvmErr.Error() == ethereum.NotFound.Error() || rpcEvmErr.Error() == "server returned transaction without signature" {
					if live {
						if stateErr := rC.state.SetState(common.EmergencyState); stateErr != nil {
							rC.logger.Info("error setting emergency state")
							rC.logger.Info(stateErr)
							rC.stopChan <- syscall.SIGKILL
						}
						return nil
					}
				} else {
					return rpcEvmErr
				}
			} else {
				found := false
				if logs, err := rC.rpcManager.Evm(param.ChainId).FilterBlockLogs(tx.BlockHash); err != nil {
					rC.logger.Debug(err)
					return err
				} else {
					logIndexToCheck := param.LogIndex
					if logIndexToCheck >= common.AffiliateLogIndexAddition {
						logIndexToCheck -= common.AffiliateLogIndexAddition
					}
					if uint32(len(logs)) <= logIndexToCheck {
						rC.logger.Debugf("Logs length %d is less than log index %d", len(logs), param.LogIndex)
					} else {
						log := logs[logIndexToCheck]
						if strings.ToLower(log.Address.String()) != strings.ToLower(rC.rpcManager.Evm(param.ChainId).BridgeAddress().String()) {
							rC.logger.Debugf("Address that generated this log %s is different than contract address %s",
								strings.ToLower(log.Address.String()), strings.ToLower(rC.rpcManager.Evm(param.ChainId).BridgeAddress().String()))
						} else if log.Topics[0].Hex() != common.UnwrapSigHash.Hex() {
							rC.logger.Debugf("log Topic hash %s is different than unwrapSigHash %s", log.Topics[0].Hex(), common.UnwrapSigHash.Hex())
						} else {
							if unwrapRequest, parseErr := rC.rpcManager.Evm(param.ChainId).Bridge().ParseUnwrapped(log); parseErr != nil {
								rC.logger.Debugf("Could not parse unwrapped event with error: %s", parseErr.Error())
							} else {
								addresses := strings.Split(unwrapRequest.To, common.AffiliateProgramAddressSeparator)
								if _, errParse := common.ParseAddressString(addresses[0], definition.NoMClass); errParse != nil {
									rC.logger.Debugf("Could not parse zenon address: %s with error: %s", addresses[0], errParse.Error())
								} else {
									// We have to check the proper amounts for the initiator or the affiliate amount
									var errParse error
									if len(addresses) > 1 {
										if _, errParse = common.ParseAddressString(addresses[1], definition.NoMClass); errParse != nil {
											rC.logger.Debugf("Could not parse affiliate zenon address: %s with error: %s", addresses[1], errParse.Error())
										}
									}

									// No matter if affiliate program is active, this transaction should be treated as normal
									zts := rC.state.GetTokensMap(param.ChainId, strings.ToLower(unwrapRequest.Token.String()))
									// Here we check the zts because the checks on evm only might not be active
									if len(addresses) == 1 || errParse != nil || !rC.state.GetIsAffiliateProgramActive(zts) {
										if param.LogIndex >= common.AffiliateLogIndexAddition {
											rC.logger.Debugf("Found affiliate logIndex but this is a non affiliate unwrap - len(addresses): %d, errParse != nil: %t,"+
												"!affiliateActive(): %t", len(addresses), errParse != nil, !rC.state.GetIsAffiliateProgramActive(zts))
										} else if addresses[0] != rpcZnnEvent.ToAddress.String() {
											rC.logger.Debugf("Normal unwrap event address %s different than znn unwrap toAddress %s", addresses[0], rpcZnnEvent.ToAddress.String())
										} else if unwrapRequest.Amount.Cmp(rpcZnnEvent.Amount) != 0 {
											rC.logger.Debugf("Normal unwrap event amount %s different than znn unwrap amount %s", unwrapRequest.Amount.String(), rpcZnnEvent.Amount.String())
										} else if strings.ToLower(unwrapRequest.Token.String()) != rpcZnnEvent.TokenAddress {
											rC.logger.Debugf("Normal unwrap event Token %s different than znn unwrap TokenAddress %s", unwrapRequest.Token.String(), rpcZnnEvent.TokenAddress)
										} else {
											found = true
										}
									} else {
										addressToCheck := addresses[0]
										amountToCheck := big.NewInt(0).Set(unwrapRequest.Amount)
										fee := big.NewInt(0).Div(amountToCheck, big.NewInt(100)) // 1%
										amountToCheck.Add(amountToCheck, fee)                    // 101%

										// we check the affiliate tx
										if param.LogIndex >= common.AffiliateLogIndexAddition {
											affiliateAmount := big.NewInt(0).Set(unwrapRequest.Amount)
											affiliateAmount.Mul(affiliateAmount, big.NewInt(2))
											affiliateAmount.Div(affiliateAmount, big.NewInt(100)) // 2%

											amountToCheck.Set(affiliateAmount)
											addressToCheck = addresses[1] // we checked that we have this address
										}

										rC.logger.Infof("addressToCheck: %s", addressToCheck)
										rC.logger.Infof("amountToCheck: %s", amountToCheck.String())
										if addressToCheck != rpcZnnEvent.ToAddress.String() {
											rC.logger.Debugf("Affiliate unwrap event address %s different than znn unwrap toAddress %s", addressToCheck, rpcZnnEvent.ToAddress.String())
										} else if amountToCheck.Cmp(rpcZnnEvent.Amount) != 0 {
											rC.logger.Debugf("Affiliate unwrap event amount %s different than znn unwrap amount %s", amountToCheck.String(), rpcZnnEvent.Amount.String())
										} else if strings.ToLower(unwrapRequest.Token.String()) != rpcZnnEvent.TokenAddress {
											rC.logger.Debugf("Affiliate unwrap event Token %s different than znn unwrap TokenAddress %s", unwrapRequest.Token.String(), rpcZnnEvent.TokenAddress)
										} else {
											found = true
										}
									}
								}
							}
						}
					}
				}
				rC.logger.Infof("found: %t", found)
				if !found && live {
					if stateErr := rC.state.SetState(common.EmergencyState); stateErr != nil {
						rC.logger.Info(stateErr)
						rC.stopChan <- syscall.SIGKILL
						return stateErr
					}
					break
				}
				rC.logger.Debug("Trying to add unwrap request")
				// we just overwrite the event and set it's status as pending redeem, we don't care about block number anymore
				if storageErr := rC.dbManager.EvmStorage(param.ChainId).AddUnwrapRequest(common.ZnnUnwrapToOrchestatorUnwrap(rpcZnnEvent)); storageErr != nil {
					return storageErr
				}
				rC.logger.Debug("Successfully added unwrap request")
				if err := rC.dbManager.EvmStorage(param.ChainId).SetUnwrapRequestStatus(ecommon.Hash(param.TransactionHash), param.LogIndex, common.PendingRedeemStatus); err != nil {
					return err
				}
			}
		}
	case base64.StdEncoding.EncodeToString(definition.ABIBridge.Methods[definition.RevokeUnwrapRequestMethodName].Id()):
		if sendBlock.Address.String() != rC.administrator {
			break
		}
		param := new(definition.RevokeUnwrapParam)
		if err := definition.ABIBridge.UnpackMethod(param, definition.RevokeUnwrapRequestMethodName, sendBlock.Data); err != nil {
			return constants.ErrUnpackError
		}

		common.AdministratorLogger.Infof("RevokeUnwrapRequestMethodName TxHash: %s, LogIndex: %d", param.TransactionHash.String(), param.LogIndex)
		if rpcEvent, rpcErr := rC.GetUnwrapTokenRequestByHashAndLog(param.TransactionHash, param.LogIndex); rpcErr != nil {
			if rpcErr.Error() == constants.ErrDataNonExistent.Error() {
				rC.logger.Debug(rpcErr)
				return nil
			}
			return rpcErr
		} else if rpcEvent == nil {
			// someone is trying to redeem a non existent event
			rC.logger.Info("event non existent")
			return nil
		} else {
			// if the event was revoked we also set it locally
			if localEvent, err := rC.dbManager.EvmStorage(rpcEvent.ChainId).GetUnwrapRequestByHashAndLog(ecommon.Hash(rpcEvent.TransactionHash), rpcEvent.LogIndex); err != nil {
				return err
			} else if localEvent == nil {
				if storageErr := rC.dbManager.EvmStorage(rpcEvent.ChainId).AddUnwrapRequest(common.ZnnUnwrapToOrchestatorUnwrap(rpcEvent)); storageErr != nil {
					return storageErr
				}
			}
			// if the event was redeemed we also set it locally
			if rpcEvent.Revoked == 1 {
				if storageErr := rC.dbManager.EvmStorage(rpcEvent.ChainId).SetUnwrapRequestStatus(ecommon.Hash(rpcEvent.TransactionHash), rpcEvent.LogIndex, common.RevokedStatus); storageErr != nil {
					return storageErr
				}
			}
		}
	case base64.StdEncoding.EncodeToString(definition.ABIBridge.Methods[definition.SetNetworkMethodName].Id()):
		if !live || sendBlock.Address.String() != rC.administrator {
			break
		}
		param := new(definition.NetworkInfoParam)
		if err := definition.ABIBridge.UnpackMethod(param, definition.SetNetworkMethodName, sendBlock.Data); err != nil {
			return constants.ErrUnpackError
		}
		common.AdministratorLogger.Info("SetNetworkMethodName NetworkClass: %d, ChainId: %d, Name: %s, ContractAddress: %s, Metadata: %s",
			param.NetworkClass, param.ChainId, param.Name, param.ContractAddress, param.Metadata)
		network, err := rC.ZnnRpc().GetNetworkByClassAndId(param.NetworkClass, param.ChainId)
		if err != nil {
			return err
		} else if network == nil {
			// we don't do anything
			rC.logger.Info("network not added")
			return nil
		}
		rC.logger.Debugf("network found in go-zeonon: %s, %d, %d", network.Name, network.NetworkClass, network.Id)
		// check locally that the network is added
		switch param.NetworkClass {
		case definition.EvmClass:
			existent := rC.rpcManager.HasEvmNetwork(param.ChainId)
			if existent {
				rC.logger.Info("network already existent")
				break
			}
			rC.logger.Debug("network non existent")
			configData, ok := rC.networksInfo[network.Name]
			if ok == false {
				rC.logger.Infof("network url non existent for network: %s chainId: %d", network.Name, network.Id)
				rC.stopChan <- syscall.SIGKILL
				return errors.New("network url non existent")
			}
			rC.logger.Info("configData: ", configData)
			newEvmNetwork, err := NewEvmNetwork(network, rC.dbManager, rC.rpcManager, rC.state, rC.stopChan)
			if err != nil {
				rC.logger.Error(err)
				rC.stopChan <- syscall.SIGKILL
				return err
			}
			err = rC.rpcManager.AddEvmClient(configData, network.Id, newEvmNetwork.NetworkName(), *newEvmNetwork.ContractAddress())
			if err != nil {
				rC.logger.Error(err)
				rC.stopChan <- syscall.SIGKILL
				return err
			}
			rC.logger.Debug("add evm client ok")
			if err := newEvmNetwork.Start(); err != nil {
				rC.logger.Error(err)
				rC.stopChan <- syscall.SIGKILL
				return err
			}
			rC.logger.Debug("network start ok")
			rC.networkManager.AddEvmNetwork(newEvmNetwork)
		}
	case base64.StdEncoding.EncodeToString(definition.ABIBridge.Methods[definition.RemoveNetworkMethodName].Id()):
		if !live || sendBlock.Address.String() != rC.administrator {
			break
		}
		param := new(definition.NetworkInfoParam)
		if err := definition.ABIBridge.UnpackMethod(param, definition.RemoveNetworkMethodName, sendBlock.Data); err != nil {
			return constants.ErrUnpackError
		}
		common.AdministratorLogger.Info("RemoveNetworkMethodName NetworkClass: %d, ChainId: %d, Name: %s, ContractAddress: %s, Metadata: %s",
			param.NetworkClass, param.ChainId, param.Name, param.ContractAddress, param.Metadata)
		network, err := rC.ZnnRpc().GetNetworkByClassAndId(param.NetworkClass, param.ChainId)
		if err != nil {
			return err
		} else if network == nil {
			switch param.NetworkClass {
			case definition.EvmClass:
				existent := rC.rpcManager.HasEvmNetwork(param.ChainId)
				if !existent {
					rC.logger.Info("network already deleted")
					break
				}
				// todo integrate these two
				rC.rpcManager.RemoveEvmClient(param.ChainId)
				rC.networkManager.RemoveEvmNetwork(param.ChainId)
			}
		}
	case base64.StdEncoding.EncodeToString(definition.ABIBridge.Methods[definition.SetTokenPairMethod].Id()):
		if !live || sendBlock.Address.String() != rC.administrator {
			break
		}
		param := new(definition.TokenPairParam)
		if err := definition.ABIBridge.UnpackMethod(param, definition.SetTokenPairMethod, sendBlock.Data); err != nil {
			return constants.ErrUnpackError
		}
		common.AdministratorLogger.Infof("SetTokenPairMethod NetworkClass: %d, ChainId: %d, TokenStandard: %s, TokenAddress: %s, Bridgeable: %t, "+
			"Redeemable: %t, Owned: %t, MinAmount: %s, FeePercentage: %d, RedeemDelay: %d, Metadata: %s",
			param.NetworkClass, param.ChainId, param.TokenStandard.String(), param.TokenAddress, param.Bridgeable, param.Redeemable, param.Owned, param.MinAmount.String(),
			param.FeePercentage, param.RedeemDelay, param.Metadata)
	case base64.StdEncoding.EncodeToString(definition.ABIBridge.Methods[definition.RemoveTokenPairMethodName].Id()):
		if !live || sendBlock.Address.String() != rC.administrator {
			break
		}
		param := new(definition.TokenPairParam)
		if err := definition.ABIBridge.UnpackMethod(param, definition.RemoveTokenPairMethodName, sendBlock.Data); err != nil {
			return constants.ErrUnpackError
		}
		common.AdministratorLogger.Infof("RemoveTokenPairMethodName NetworkClass: %d, ChainId: %d, TokenStandard: %s, TokenAddress: %s",
			param.NetworkClass, param.ChainId, param.TokenStandard.String(), param.TokenAddress)
	case base64.StdEncoding.EncodeToString(definition.ABIBridge.Methods[definition.HaltMethodName].Id()):
		if !live {
			break
		}
		common.AdministratorLogger.Info("HaltMethodName")
		currentState, err := rC.state.GetState()
		if err != nil {
			rC.logger.Debug(err)
			rC.stopChan <- syscall.SIGKILL
			return err
		}
		// if the node is in emergency, it will set the state to halted after all txs, we don't need to do it after we see one
		if currentState != common.EmergencyState {
			halted := rC.IsHalted()
			if halted {
				if err := rC.state.SetState(common.HaltedState); err != nil {
					rC.logger.Error(err)
					rC.stopChan <- syscall.SIGKILL
					return err
				}
			}
		}
	case base64.StdEncoding.EncodeToString(definition.ABIBridge.Methods[definition.NominateGuardiansMethodName].Id()):
		if !live || sendBlock.Address.String() != rC.administrator {
			break
		}
		guardians := new([]types.Address)
		if err := definition.ABIBridge.UnpackMethod(guardians, definition.NominateGuardiansMethodName, sendBlock.Data); err != nil {
			return constants.ErrUnpackError
		}
		for idx, guardian := range *guardians {
			common.AdministratorLogger.Infof("NominateGuardiansMethodName Idx: %d, Guardian: %s", idx, guardian.String())
		}
	case base64.StdEncoding.EncodeToString(definition.ABIBridge.Methods[definition.UnhaltMethodName].Id()):
		if !live || sendBlock.Address.String() != rC.administrator {
			break
		}
		common.AdministratorLogger.Info("UnhaltMethodName")
	case base64.StdEncoding.EncodeToString(definition.ABIBridge.Methods[definition.ProposeAdministratorMethodName].Id()):
		if !live {
			break
		}
		address := new(types.Address)
		if err := definition.ABIBridge.UnpackMethod(address, definition.ProposeAdministratorMethodName, sendBlock.Data); err != nil {
			return constants.ErrUnpackError
		}
		common.AdministratorLogger.Infof("ProposeAdministratorMethodName %s", address.String())
	case base64.StdEncoding.EncodeToString(definition.ABIBridge.Methods[definition.EmergencyMethodName].Id()):
		if !live || sendBlock.Address.String() != rC.administrator {
			break
		}
		common.AdministratorLogger.Info("EmergencyMethodName")
		halted := rC.IsHalted()
		if halted {
			if err := rC.state.SetState(common.HaltedState); err != nil {
				rC.logger.Error(err)
				rC.stopChan <- syscall.SIGKILL
				return err
			}
		}
		bridgeInfo, err := rC.GetBridgeInfo()
		if err != nil {
			return err
		}
		rC.administrator = bridgeInfo.Administrator.String()
	case base64.StdEncoding.EncodeToString(definition.ABIBridge.Methods[definition.ChangeTssECDSAPubKeyMethodName].Id()):
		if !live {
			break
		}
		param := new(definition.ChangeECDSAPubKeyParam)
		if err := definition.ABIBridge.UnpackMethod(param, definition.ChangeTssECDSAPubKeyMethodName, sendBlock.Data); err != nil {
			return constants.ErrUnpackError
		}
		common.AdministratorLogger.Infof("ChangeTssECDSAPubKeyMethodName PubKey: %s, OldSig: %s, NewSig: %s", param.PubKey, param.OldPubKeySignature, param.NewPubKeySignature)
		bridgeInfo, err := rC.GetBridgeInfo()
		if err != nil {
			return err
		}
		// If the key was changed to the param, we are no longer in keyGen
		if bridgeInfo.CompressedTssECDSAPubKey == param.PubKey {
			currentState, err := rC.state.GetState()
			if err != nil {
				rC.logger.Debug(err)
				rC.stopChan <- syscall.SIGKILL
				return err
			}
			if currentState == common.KeyGenState {
				if err := rC.state.SetState(common.LiveState); err != nil {
					rC.logger.Error(err)
					rC.stopChan <- syscall.SIGKILL
					return err
				}
			}
		}
	case base64.StdEncoding.EncodeToString(definition.ABIBridge.Methods[definition.ChangeAdministratorMethodName].Id()):
		if !live || sendBlock.Address.String() != rC.administrator {
			break
		}
		address := new(types.Address)
		if err := definition.ABIBridge.UnpackMethod(address, definition.ChangeAdministratorMethodName, sendBlock.Data); err != nil {
			return constants.ErrUnpackError
		}
		common.AdministratorLogger.Infof("ChangeAdministratorMethodName %s", address.String())
		bridgeInfo, err := rC.GetBridgeInfo()
		if err != nil {
			return err
		}
		rC.administrator = bridgeInfo.Administrator.String()
	case base64.StdEncoding.EncodeToString(definition.ABIBridge.Methods[definition.SetAllowKeygenMethodName].Id()):
		if !live || sendBlock.Address.String() != rC.administrator {
			break
		}
		var value bool
		if err := definition.ABIBridge.UnpackMethod(&value, definition.SetAllowKeygenMethodName, sendBlock.Data); err != nil {
			return constants.ErrUnpackError
		}

		common.AdministratorLogger.Infof("SetAllowKeygenMethodName %t", value)
		if value == false {
			currentState, err := rC.state.GetState()
			if err != nil {
				rC.logger.Debug(err)
				rC.stopChan <- syscall.SIGKILL
				return err
			}
			if currentState == common.KeyGenState {
				if err := rC.state.SetState(common.LiveState); err != nil {
					rC.logger.Error(err)
					rC.stopChan <- syscall.SIGKILL
					return err
				}
			}
		}
	case base64.StdEncoding.EncodeToString(definition.ABIBridge.Methods[definition.SetOrchestratorInfoMethodName].Id()):
		if !live || sendBlock.Address.String() != rC.administrator {
			break
		}
		common.AdministratorLogger.Info("SetOrchestratorInfoMethodName")
		rC.logger.Debug("found SetOrchestratorInfoMethodName")
		orchestratorInfo, err := rC.GetOrchestratorInfo()
		if err != nil {
			rC.logger.Error(err)
			return err
		}
		rC.SetWindowSize(orchestratorInfo.WindowSize)
		rC.SetKeyGenThreshold(orchestratorInfo.KeyGenThreshold)
		rC.SetConfirmationsToFinality(orchestratorInfo.ConfirmationsToFinality)
		rC.SetEstimatedMomentumTime(orchestratorInfo.EstimatedMomentumTime)
	case base64.StdEncoding.EncodeToString(definition.ABIBridge.Methods[definition.SetBridgeMetadataMethodName].Id()):
		if !live || sendBlock.Address.String() != rC.administrator {
			break
		}
		var param string
		if err := definition.ABIBridge.UnpackMethod(&param, definition.SetBridgeMetadataMethodName, sendBlock.Data); err != nil {
			return constants.ErrUnpackError
		}
		common.AdministratorLogger.Infof("SetBridgeMetadataMethodName %s", param)

		metadata := &common.BridgeMetadata{}

		if err := json.Unmarshal([]byte(param), metadata); err != nil {
			return err
		} else {
			rC.setBridgeMetadata(metadata)
		}
	case base64.StdEncoding.EncodeToString(definition.ABIBridge.Methods[definition.SetNetworkMetadataMethodName].Id()):
		if !live || sendBlock.Address.String() != rC.administrator {
			break
		}
		param := new(definition.SetNetworkMetadataParam)
		if err := definition.ABIBridge.UnpackMethod(param, definition.SetNetworkMetadataMethodName, sendBlock.Data); err != nil {
			return constants.ErrUnpackError
		}
		common.AdministratorLogger.Infof("SetNetworkMetadataMethodName NetworkClass: %d, ChainId: %d, Metadata: %s",
			param.NetworkClass, param.ChainId, param.Metadata)
	}

	return rC.eventsStore().SetLastUpdateHeight(receiveBlockHeight)
}

// Subscriptions

func (rC *znnNetwork) ListenForMomentumHeight() {
	rC.logger.Debug("func (rC *znnNetwork) ListenForMomentumHeight() {")
	momSub, momChan, err := rC.ZnnRpc().SubscribeToMomentums()
	if err != nil {
		rC.logger.Error(err)
		rC.stopChan <- syscall.SIGKILL
		return
	}
	rC.logger.Debug("Successfully started to listen for momentums")
	for {
		select {
		case errSub := <-momSub.Err():
			if errSub != nil {
				rC.logger.Debugf("listen mom err: %s", errSub.Error())
				rC.stopChan <- syscall.SIGKILL
				return
			}
		case momentums := <-momChan:
			for _, mom := range momentums {
				if frMom, errState := rC.state.GetFrontierMomentum(); errState != nil {
					rC.logger.Info("error when trying to get frontierMomentum from state")
					rC.logger.Error(errState)
				} else {
					if mom.Height > frMom {
						if errState = rC.state.SetFrontierMomentum(mom.Height); errState != nil {
							rC.logger.Error(errState)
							rC.logger.Info("error when trying to set frontier momentum")
						}
					}
				}
			}
		}
	}
}

func (rC *znnNetwork) ListenForEmbeddedBridgeAccountBlocks() {
	rC.logger.Debug("ListenForEmbeddedBridgeAccountBlocks")
	accBlSub, accBlCh, err := rC.ZnnRpc().SubscribeToAccountBlocks(types.BridgeContract)
	if err != nil {
		rC.logger.Info("sub accBerr: ", err)
		rC.stopChan <- syscall.SIGKILL
		return
	}
	rC.logger.Debug("Successfully started to listen for account blocks")
	for {
		select {
		case errSub := <-accBlSub.Err():
			if errSub != nil {
				rC.logger.Debugf("listen accB err: %s", errSub.Error())
				rC.stopChan <- syscall.SIGKILL
				return
			}
		case accBlocks := <-accBlCh:
			// these accountBlocks are seen before being inserted in a momentum
			for _, accBlock := range accBlocks {
				if accBlock.BlockType != nom.BlockTypeContractReceive {
					continue
				}
				// we wait for the acc block to be inserted in a momentum
				for {
					time.Sleep(4 * time.Second)
					if receiveBlock, err := rC.ZnnRpc().GetAccountBlockByHash(accBlock.Hash); err != nil {
						rC.logger.Info("receive block non existent")
						continue
					} else if receiveBlock == nil {
						rC.logger.Info("receive block non existent")
						continue
					}
					break
				}

				rC.logger.Info("detected block type receive")
				if sendBlock, err := rC.ZnnRpc().GetAccountBlockByHash(accBlock.FromHash); err != nil {
					rC.logger.Error(err)
				} else if sendBlock == nil {
					rC.logger.Info("send block non existent")
				} else {
					rC.logger.Info("found send block")
					rC.logger.Infof("confirmationDetail is nil: %v", sendBlock.ConfirmationDetail == nil)
					if newErr := rC.InterpretSendBlockData(sendBlock, true, accBlock.Height); newErr != nil {
						rC.logger.Info(newErr)
						// Try one more time
						time.Sleep(3 * time.Second)
						if newErr = rC.InterpretSendBlockData(sendBlock, true, accBlock.Height); newErr != nil {
							rC.logger.Info(newErr)
						}
					}
				}
			}
		}
	}
}

// Transactions

func (rC *znnNetwork) UpdateWrapRequest(id types.Hash, signature string, keyPair *wallet.KeyPair) error {
	return rC.ZnnRpc().UpdateWrapRequest(id, signature, keyPair)
}

func (rC *znnNetwork) SendUnwrapRequest(event *events.UnwrapRequestEvm, keyPair *wallet.KeyPair) error {
	return rC.ZnnRpc().SendUnwrapRequest(event, keyPair)
}

func (rC *znnNetwork) Halt(signature string, keyPair *wallet.KeyPair) error {
	return rC.ZnnRpc().Halt(signature, keyPair)
}

/// Rpc Calls

func (rC *znnNetwork) GetFrontierMomentum() (*api.Momentum, error) {
	return rC.ZnnRpc().GetFrontierMomentum()
}

func (rC *znnNetwork) GetUnsignedWrapRequestsRpc(pageIndex, pageSize uint32) (*embedded.WrapTokenRequestList, error) {
	return rC.ZnnRpc().GetAllUnsignedWrapTokenRequests(pageIndex, pageSize)
}

func (rC *znnNetwork) GetAllWrapTokenRequests(pageIndex, pageSize uint32) (*embedded.WrapTokenRequestList, error) {
	return rC.ZnnRpc().GetAllWrapTokenRequests(pageIndex, pageSize)
}

func (rC *znnNetwork) GetUnwrapTokenRequestByHashAndLog(txHash types.Hash, logIndex uint32) (*definition.UnwrapTokenRequest, error) {
	return rC.ZnnRpc().GetUnwrapTokenRequestByHashAndLog(txHash, logIndex)
}

func (rC *znnNetwork) GetWrapTokenRequestById(id types.Hash) (*definition.WrapTokenRequest, error) {
	return rC.ZnnRpc().GetWrapTokenRequestById(id)
}

func (rC *znnNetwork) ChangeTssEcdsaPubKey(pubKey, signature, newSignature string, keyPair *wallet.KeyPair) error {
	return rC.ZnnRpc().ChangeTssEcdsaPubKey(pubKey, signature, newSignature, keyPair)
}

func (rC *znnNetwork) GetPillarPublicKeys() (map[string]string, error) {
	return rC.ZnnRpc().GetPillarPublicKeys()
}

func (rC *znnNetwork) GetAllNetworks() ([]*definition.NetworkInfo, error) {
	return rC.ZnnRpc().GetAllNetworks()
}

func (rC *znnNetwork) GetBridgeInfo() (*definition.BridgeInfoVariable, error) {
	return rC.ZnnRpc().GetBridgeInfo()
}

func (rC *znnNetwork) GetOrchestratorInfo() (*definition.OrchestratorInfo, error) {
	return rC.ZnnRpc().GetOrchestratorInfo()
}

func (rC *znnNetwork) IsSynced() bool {
	syncInfo, err := rC.ZnnRpc().GetSyncInfo()
	if err != nil {
		rC.logger.Error(err)
	}
	return syncInfo.State == protocol.SyncDone
}

func (rC *znnNetwork) IsHalted() bool {
	bridgeInfo, err := rC.GetBridgeInfo()
	if err != nil {
		rC.logger.Error(err)
		rC.stopChan <- syscall.SIGKILL
		return true
	}
	frMom, frMomErr := rC.state.GetFrontierMomentum()
	if frMomErr != nil {
		rC.logger.Error(err)
		rC.stopChan <- syscall.SIGKILL
		return true
	}

	return bridgeInfo.Halted || bridgeInfo.UnhaltedAt+bridgeInfo.UnhaltDurationInMomentums >= frMom
}

/// Local storage

func (rC *znnNetwork) GetUnsentSignedWrapRequests() ([]*events.WrapRequestZnn, error) {
	return rC.eventsStore().GetUnsentSignedWrapRequests()
}

func (rC *znnNetwork) GetUnredeemedWrapRequests() ([]*events.WrapRequestZnn, error) {
	return rC.eventsStore().GetUnredeemedWrapRequests()
}

func (rC *znnNetwork) SetWrapEventSignature(id types.Hash, signature string) error {
	return rC.eventsStore().SetWrapRequestSignature(id, signature)
}

func (rC *znnNetwork) GetWrapEventById(id types.Hash) (*events.WrapRequestZnn, error) {
	return rC.eventsStore().GetWrapRequestById(id)
}

func (rC *znnNetwork) AddWrapEvent(rpcEvent *definition.WrapTokenRequest) error {
	event := events.WrapRequestZnn{
		NetworkClass:  rpcEvent.NetworkClass,
		ChainId:       rpcEvent.ChainId,
		Id:            rpcEvent.Id,
		ToAddress:     rpcEvent.ToAddress,
		TokenAddress:  rpcEvent.TokenAddress,
		Amount:        big.NewInt(0).Set(rpcEvent.Amount),
		Fee:           big.NewInt(0).Set(rpcEvent.Fee),
		Signature:     rpcEvent.Signature,
		RedeemStatus:  common.UnredeemedStatus,
		SentSignature: false,
	}

	return rC.eventsStore().AddWrapRequest(event)
}

func (rC *znnNetwork) SetWrapEventStatus(id types.Hash, status uint32) error {
	return rC.eventsStore().SetWrapRequestStatus(id, status)
}

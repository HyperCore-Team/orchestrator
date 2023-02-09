package rpc

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ecommon "github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
	"github.com/zenon-network/go-zenon/vm/embedded/implementation"
	"go.uber.org/zap"
	"golang.org/x/crypto/sha3"
	"math/big"
	"orchestrator/common"
	"orchestrator/common/bridge"
	"orchestrator/common/config"
)

type EvmRpc struct {
	rpcClient *ethclient.Client
	Urls      config.UrlsInfo

	bridgeContract  *bridge.Bridge
	bridgeAddress   ecommon.Address
	filterQuery     ethereum.FilterQuery
	filterQuerySize uint64

	logSub  ethereum.Subscription
	logChan chan etypes.Log
	logger  *zap.SugaredLogger
}

func NewEvmRpcClient(networkConfig config.BaseNetworkConfig, address ecommon.Address) (*EvmRpc, error) {
	logger, errLog := common.CreateSugarLogger()
	if errLog != nil {
		return nil, errLog
	}

	newUrls, err := config.NewUrlsInfo(networkConfig)
	if err != nil {
		return nil, err
	}
	var newRpcClient *ethclient.Client
	currentUrl := newUrls.GetCurrentUrl()
	for {
		newRpcClient, err = ethclient.Dial(currentUrl)
		if err != nil {
			logger.Infof("Error when dialing %s, got: %s\n", currentUrl, err)
		} else {
			break
		}
		currentUrl = newUrls.NextUrl()
		if len(currentUrl) == 0 {
			return nil, errors.New("cannot connect to any url on evm")
		}
	}
	newUrls.Clear()

	newBridgeContract, err := bridge.NewBridge(address, newRpcClient)
	if err != nil {
		return nil, err
	}

	newFilterQuery := ethereum.FilterQuery{
		Addresses: []ecommon.Address{address},
		Topics:    common.Topics,
	}

	return &EvmRpc{
		rpcClient:       newRpcClient,
		Urls:            *newUrls,
		bridgeContract:  newBridgeContract,
		bridgeAddress:   address,
		filterQuery:     newFilterQuery,
		filterQuerySize: networkConfig.FilterQuerySize,
		logChan:         make(chan etypes.Log, 20000),
		logger:          logger,
	}, nil
}

/// Utils

func (r *EvmRpc) Bridge() *bridge.Bridge {
	return r.bridgeContract
}

func (r *EvmRpc) BridgeAddress() ecommon.Address {
	return r.bridgeAddress
}

func (r *EvmRpc) FilterQuerySize() uint64 {
	return r.filterQuerySize
}

func (r *EvmRpc) Stop() {
	r.rpcClient.Close()
	r.logSub.Unsubscribe()
}

func (r *EvmRpc) IsSynced() bool {
	syncProgress, err := r.rpcClient.SyncProgress(context.Background())
	if err != nil {
		r.logger.Debug(err)
		return false
	}
	if syncProgress == nil {
		return true
	}

	// we consider that we are synced if we are at most 4 blocks behind
	return syncProgress.HighestBlock-syncProgress.CurrentBlock < 4
}

/// Subscribe

func (r *EvmRpc) SubscribeToLogs() (ethereum.Subscription, chan etypes.Log, error) {
	var err error
	r.logSub, err = r.rpcClient.SubscribeFilterLogs(context.Background(), r.filterQuery, r.logChan)
	if err != nil {
		r.logger.Error("Error after subscribe: %s", err.Error())
		return nil, nil, err
	}
	return r.logSub, r.logChan, nil
}

/// Transactions

func (r *EvmRpc) SendTransaction(tx *etypes.Transaction) error {
	return r.rpcClient.SendTransaction(context.Background(), tx)
}

func (r *EvmRpc) GetChangeTssEcdsaPubKeyEvmMessage(newAddress ecommon.Address, networkClass, chainId uint32, contractAddress *ecommon.Address) ([]byte, error) {
	actionsNonce, err := r.GetActionNonce()
	if err != nil {
		return nil, err
	}

	args := abi.Arguments{{Type: definition.StringTy}, {Type: definition.Uint256Ty}, {Type: definition.Uint256Ty}, {Type: definition.AddressTy}, {Type: definition.Uint256Ty}, {Type: definition.AddressTy}}
	values := make([]interface{}, 0)
	values = append(values, "changeTssAddress",
		big.NewInt(int64(networkClass)),
		big.NewInt(int64(chainId)),
		contractAddress,
		actionsNonce,
		newAddress)

	packedData, err := args.PackValues(values)
	if err != nil {
		return nil, err
	}

	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(packedData)
	data := hasher.Sum(nil)
	return implementation.GetMessageToSignEvm(data)
}

func (r *EvmRpc) GetChangeTssEcdsaPubKeyEvmTx(newTssAddress, sender ecommon.Address, oldFullSignature, newFullSignature []byte, contractAddress *ecommon.Address) (*etypes.Transaction, error) {
	if blockHeight, err := r.BlockNumber(); err != nil {
		return nil, err
	} else {
		if balance, err := r.BalanceAt(sender, blockHeight); err != nil {
			return nil, err
		} else {
			if balance.Cmp(big.NewInt(0)) == 0 {
				return nil, errors.New("Balance is 0, not enough to send change ecdsa pub key tx")
			}
			r.logger.Debugf("Balance: %d", balance.Uint64())
			if nonce, err := r.NonceAt(sender, blockHeight); err != nil {
				return nil, err
			} else {
				gasPrice, err := r.SuggestGasPrice()
				if err != nil {
					return nil, err
				}

				encodedData, err := common.EvmContractAbi.Pack("changeTssAddress", newTssAddress, oldFullSignature, newFullSignature)
				if err != nil {
					return nil, err
				}
				msg := ethereum.CallMsg{
					From:     sender,
					To:       contractAddress,
					Gas:      0,
					GasPrice: gasPrice,
					Value:    big.NewInt(0),
					Data:     encodedData,
				}
				estimatedGas, err := r.EstimateGas(msg)
				if err != nil {
					return nil, err
				} else {
					r.logger.Debug("estimatedGas: ", estimatedGas)
				}

				fees := big.NewInt(0).Mul(gasPrice, big.NewInt(0).SetUint64(estimatedGas))
				// We subtract the fees and send the difference to the contract
				// We subtract the fees and send the difference to the contract
				if balance.Cmp(fees) < 0 {
					return nil, errors.New("not enough balance to send change ecdsa pub key tx")
				}
				balance.Sub(balance, fees)
				tx := etypes.NewTx(&etypes.LegacyTx{
					Nonce:    nonce,
					To:       contractAddress,
					Value:    big.NewInt(0),
					Gas:      estimatedGas,
					GasPrice: gasPrice,
					Data:     encodedData,
				})

				return tx, nil
			}
		}
	}
}

func (r *EvmRpc) GetHaltEvmMessage(networkClass, chainId uint32, contractAddress *ecommon.Address) ([]byte, error) {
	actionsNonce, err := r.GetActionNonce()
	if err != nil {
		return nil, err
	}

	args := abi.Arguments{{Type: definition.StringTy}, {Type: definition.Uint256Ty}, {Type: definition.Uint256Ty}, {Type: definition.AddressTy}, {Type: definition.Uint256Ty}}
	values := make([]interface{}, 0)
	values = append(values, "halt", big.NewInt(int64(networkClass)), big.NewInt(int64(chainId)), contractAddress, actionsNonce)

	packedData, err := args.PackValues(values)
	if err != nil {
		return nil, err
	}

	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(packedData)
	data := hasher.Sum(nil)
	return implementation.GetMessageToSignEvm(data)
}

func (r *EvmRpc) GetHaltTxEvm(sender ecommon.Address, signature []byte, contractAddress *ecommon.Address) (*etypes.Transaction, error) {
	if blockHeight, err := r.BlockNumber(); err != nil {
		return nil, err
	} else {
		if balance, err := r.BalanceAt(sender, blockHeight); err != nil {
			return nil, err
		} else {
			if balance.Cmp(big.NewInt(0)) == 0 {
				return nil, errors.New("Balance is 0, not enough to send halt evm tx")
			}
			r.logger.Debugf("Balance: %d", balance.Uint64())
			if nonce, err := r.NonceAt(sender, blockHeight); err != nil {
				return nil, err
			} else {
				gasPrice, err := r.SuggestGasPrice()
				if err != nil {
					return nil, err
				}

				encodedData, err := common.EvmContractAbi.Pack("halt", signature)
				if err != nil {
					return nil, err
				}
				msg := ethereum.CallMsg{
					From:     sender,
					To:       contractAddress,
					Gas:      0,
					GasPrice: gasPrice,
					Value:    big.NewInt(0),
					Data:     encodedData,
				}
				estimatedGas, err := r.EstimateGas(msg)
				if err != nil {
					return nil, err
				}

				fees := big.NewInt(0).Mul(gasPrice, big.NewInt(0).SetUint64(estimatedGas))
				// We subtract the fees and send the difference to the contract
				if balance.Cmp(fees) < 0 {
					return nil, errors.New("not enough balance to send halt evm tx")
				}
				balance.Sub(balance, fees)
				tx := etypes.NewTx(&etypes.LegacyTx{
					Nonce:    nonce,
					To:       contractAddress,
					Value:    big.NewInt(0),
					Gas:      estimatedGas,
					GasPrice: gasPrice,
					Data:     encodedData,
				})

				return tx, nil
			}
		}
	}
}

/// Rpc Calls

func (r *EvmRpc) FilterLogs(left, right uint64) ([]etypes.Log, error) {
	r.filterQuery.FromBlock = big.NewInt(0).SetUint64(left)
	r.filterQuery.ToBlock = big.NewInt(0).SetUint64(right)
	defer func() {
		r.filterQuery.FromBlock = nil
		r.filterQuery.ToBlock = nil
	}()
	return r.rpcClient.FilterLogs(context.Background(), r.filterQuery)
}

func (r *EvmRpc) TransactionReceipt(txHash ecommon.Hash) (*etypes.Receipt, error) {
	return r.rpcClient.TransactionReceipt(context.Background(), txHash)
}

func (r *EvmRpc) EstimateGas(msg ethereum.CallMsg) (uint64, error) {
	return r.rpcClient.EstimateGas(context.Background(), msg)
}

func (r *EvmRpc) SuggestGasPrice() (*big.Int, error) {
	return r.rpcClient.SuggestGasPrice(context.Background())
}

func (r *EvmRpc) NonceAt(address ecommon.Address, blockNumber uint64) (uint64, error) {
	return r.rpcClient.NonceAt(context.Background(), address, big.NewInt(0).SetUint64(blockNumber))
}

func (r *EvmRpc) BalanceAt(address ecommon.Address, blockNumber uint64) (*big.Int, error) {
	return r.rpcClient.BalanceAt(context.Background(), address, big.NewInt(0).SetUint64(blockNumber))
}

func (r *EvmRpc) BlockNumber() (uint64, error) {
	return r.rpcClient.BlockNumber(context.Background())
}

func (r *EvmRpc) GetCurrentTssAddress() (ecommon.Address, error) {
	return r.bridgeContract.TssAddress(nil)
}

func (r *EvmRpc) IsHalted() (bool, error) {
	return r.bridgeContract.IsHalted(nil)
}

func (r *EvmRpc) GetActionNonce() (*big.Int, error) {
	return r.bridgeContract.ActionsNonce(nil)
}

func (r *EvmRpc) EstimatedBlockTime() (uint64, error) {
	return r.bridgeContract.EstimatedBlockTime(nil)
}

func (r *EvmRpc) ConfirmationsToFinality() (uint64, error) {
	return r.bridgeContract.ConfirmationsToFinality(nil)
}

func (r *EvmRpc) ContractDeploymentHeight() (uint64, error) {
	ans, err := r.bridgeContract.ContractDeploymentHeight(nil)
	if err != nil {
		return 0, err
	}
	return ans.Uint64(), nil
}

func (r *EvmRpc) TransactionByHash(hash ecommon.Hash) (*etypes.Transaction, bool, error) {
	return r.rpcClient.TransactionByHash(context.Background(), hash)
}

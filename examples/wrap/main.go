package main

import (
	"fmt"
	"github.com/MoonBaZZe/znn-sdk-go/zenon"
	"math/big"
	"time"

	"github.com/zenon-network/go-zenon/common/types"
)

func main() {
	// Initialize zenon client with keyFile
	z, err := zenon.NewZenon("producer.json")

	if err != nil {
		zenon.CommonLogger.Error("Error while creating Zenon SDK instance", "error", err)
		return
	}

	if err := z.Start("Pass-123456", "ws://127.0.0.1:35998", 1); err != nil {
		fmt.Println(err)
		zenon.CommonLogger.Error("Error while trying to connect to node", "error", err)
		return
	}

	// Test wrap
	fmt.Println("Test wrap")
	if err := z.Send(z.Client.BridgeApi.WrapToken(2, 31337, "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266", big.NewInt(10*1e8), types.ZnnTokenStandard)); err != nil {
		fmt.Println(err)
	}
	if err := z.Send(z.Client.BridgeApi.WrapToken(2, 31337, "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266", big.NewInt(8*1e8), types.QsrTokenStandard)); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Done")

	// Wait at least 1 Momentum
	time.Sleep(31 * time.Second)

	if momentum, err := z.Client.LedgerApi.GetFrontierMomentum(); err == nil {
		fmt.Println("Current Momentum Height: ", momentum.Momentum.Height)
	}

	fmt.Println("--------------------------------\nBridge Info")
	bridgeInfo, err := z.Client.BridgeApi.GetBridgeInfo()
	if bridgeInfo == nil {
		fmt.Println("bridge info nil")
	} else {
		fmt.Println("CompressedTssECDSAPubKey: ", bridgeInfo.CompressedTssECDSAPubKey)
		fmt.Println("DecompressedTssECDSAPubKey: ", bridgeInfo.DecompressedTssECDSAPubKey)
		fmt.Println("AllowKeygen: ", bridgeInfo.AllowKeyGen)
		fmt.Println("HaltActivated: ", bridgeInfo.Halted)
		fmt.Println("UnhaltHeight: ", bridgeInfo.UnhaltedAt)
		fmt.Println("UnhaltDurationInMomentums: ", bridgeInfo.UnhaltDurationInMomentums)
		fmt.Println("Metadata: ", bridgeInfo.Metadata)
	}

	fmt.Println("--------------------------------\nOrchestrator Info")
	orchestratorInfo, err := z.Client.BridgeApi.GetOrchestratorInfo()
	if orchestratorInfo == nil {
		fmt.Println("orchestrator info nil")
	} else {
		fmt.Println("WindowSize: ", orchestratorInfo.WindowSize)
		fmt.Println("KeyGenThreshold: ", orchestratorInfo.KeyGenThreshold)
		fmt.Println("EstimatedMomentumTime: ", orchestratorInfo.EstimatedMomentumTime)
		fmt.Println("ConfirmationsToFinality: ", orchestratorInfo.ConfirmationsToFinality)
		fmt.Println("AllowKeyGenHeight: ", orchestratorInfo.AllowKeyGenHeight)
	}

	fmt.Println("--------------------------------\nSecurity Info")
	securityInfo, err := z.Client.BridgeApi.GetSecurityInfo()
	if securityInfo == nil {
		fmt.Println("securityInfo nil")
	} else {
		fmt.Println("AdministratorDelay: ", securityInfo.AdministratorDelay)
		fmt.Println("SoftDelay: ", securityInfo.SoftDelay)
		fmt.Println("Guardians: ", securityInfo.Guardians)
		fmt.Println("GuardiansVotes: ", securityInfo.GuardiansVotes)
	}

	fmt.Println("--------------------------------\nNetworks")
	networksInfo, err := z.Client.BridgeApi.GetAllNetworks(0, 5)
	if err != nil {
		fmt.Println(err)
	}
	for _, networkInfo := range networksInfo.List {
		if len(networkInfo.Name) == 0 {
			fmt.Println("network info is nil")
		} else {
			fmt.Println("Name: ", networkInfo.Name)
			fmt.Println("Id: ", networkInfo.Id)
			fmt.Println("ContractAddress: ", networkInfo.ContractAddress)

			fmt.Println("---------------------------------------\nToken Pairs for " + networkInfo.Name)
			for _, token := range networkInfo.TokenPairs {
				fmt.Println("TokenStandard: ", token.TokenStandard)
				fmt.Println("TokenAddress: ", token.TokenAddress)
				fmt.Println("Metadata: ", token.Metadata)
				fmt.Println("FeePercentage: ", token.FeePercentage)
				fmt.Println("MinAmount: ", token.MinAmount)
				fmt.Println("Owned: ", token.Owned)
				fmt.Println("RedeemDelay: ", token.RedeemDelay)
				fmt.Println("Redeemable: ", token.Redeemable)
			}
		}
	}

	fmt.Println("---------------------------------------\nWraps")
	requests, err := z.Client.BridgeApi.GetAllWrapTokenRequests(0, 5)
	if err != nil {
		fmt.Println(err)
	} else if requests == nil || requests.List == nil {
		fmt.Println("requests nil")
	} else {
		fmt.Println(requests.Count)
		for _, request := range requests.List {
			fmt.Println("TokenStandard:", request.TokenStandard.String())
			fmt.Println("Id: ", request.Id.String())
			fmt.Println("ChainId: ", request.ChainId)
			fmt.Println("Amount: ", request.Amount.Uint64())
			fmt.Println("Fee: ", request.Fee.Uint64())
			fmt.Println("ToAddress: ", request.ToAddress)
			fmt.Println("TokenAddress: ", request.TokenAddress)
			fmt.Println("ConfirmationsToFinality: ", request.ConfirmationsToFinality)
			fmt.Println("Signature: ", request.Signature)
		}
	}

	fmt.Println("--------------------------------------\nUnwraps")
	requests2, err := z.Client.BridgeApi.GetAllUnwrapTokenRequests(0, 250)
	if err != nil {
		fmt.Println(err)
	} else if requests2 == nil || requests2.List == nil {
		fmt.Println("requests nil")
	} else {
		for _, requestF := range requests2.List {
			fmt.Println("TransactionHash: ", requestF.TransactionHash.String())
			fmt.Println("LogIndex: ", requestF.LogIndex)
			fmt.Println("RegistrationMomentumHeight: ", requestF.RegistrationMomentumHeight)
			fmt.Println("NetworkClass: ", requestF.NetworkClass)
			fmt.Println("ChainId: ", requestF.ChainId)
			fmt.Println("Amount: ", requestF.Amount.String())
			fmt.Println("Redeemed: ", requestF.Redeemed)
			fmt.Println("Revoked: ", requestF.Revoked)
			fmt.Println("ToAddress: ", requestF.ToAddress)
		}
	}

	if err := z.Stop(); err != nil {
		zenon.CommonLogger.Error("Error while stopping Zenon SDK instance", "error", err)
	}
}

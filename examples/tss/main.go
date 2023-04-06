package main

import (
	"fmt"
	"math/big"
	"orchestrator/app"
	"os"
	"strconv"
	"time"
	"znn-sdk-go/zenon"

	"github.com/zenon-network/go-zenon/common/types"
)

func main() {
	args := os.Args
	numPillars, err := strconv.Atoi(args[1])

	// Initialize zenon client with keyFile
	z, err := zenon.NewZenon("producer.json")

	if err != nil {
		zenon.CommonLogger.Error("Error while creating Zenon SDK instance", "error", err)
		return
	}

	if err := z.Start("Pass-123456", "ws://127.0.0.1:35998", 0); err != nil {
		fmt.Println(err)
		zenon.CommonLogger.Error("Error while trying to connect to node", "error", err)
		return
	}

	for {
		// Checking if the guardians are already set
		fmt.Printf("Checking if the guardians are already set\n")
		guardiansCheck, err := z.Client.BridgeApi.GetSecurityInfo()
		if err != nil {
			fmt.Println(err)
			return
		} else if len(guardiansCheck.Guardians) > 0 {
			fmt.Printf("Guardians already set, skipping\n")
			break
		} else {
			// Wait 5 momentums for Bridge spork to be activated
			fmt.Printf("Wait 5 momentums for Bridge spork to be activated\n")
			for {
				if momentum, err := z.Client.LedgerApi.GetFrontierMomentum(); err == nil {
					fmt.Printf("Current Momentum Height %d of %d ", momentum.Momentum.Height, 5)

					if momentum.Momentum.Height > 5 {
						break
					}
				}

				// Wait at least 1 Momentum
				time.Sleep(11 * time.Second)
			}
			fmt.Println("Done")

			// Nominate Guardians (Step 1/2)
			// We won't be able to perform any other action on the bridge if the guardians are not set
			fmt.Println("Nominate Guardians (Step 1/2)")
			if err := z.Send(z.Client.BridgeApi.NominateGuardians([]string{"qwTYX+dQSceeIM6WYAJjlbe5S0HHIdOnoBqlcaXB8YM=", "PGXPrhy88Mt9FoMHbJMyD/ECxT5o/T843MbnYzBV234=", "BEDqbIoqicO03NPOsvBIaVzkH4QAaE+3dsDALB03IQg=", "1PWYLNn2Xvw8t/IP/Q9msGAg72GH/7B5mLJRjJ80P3U=", "xhiXsj84MWLSYlDhppkgSdec5ptKd97yJ4iHICJ5oJU="})); err != nil {
				fmt.Println(err)
			}
			fmt.Println("Done")

			// Wait at least 1 Momentum
			time.Sleep(31 * time.Second)

			// Commit Guardians (Step 2/2)
			fmt.Println("Commit Guardians (Step 2/2)")
			if err := z.Send(z.Client.BridgeApi.NominateGuardians([]string{"qwTYX+dQSceeIM6WYAJjlbe5S0HHIdOnoBqlcaXB8YM=", "PGXPrhy88Mt9FoMHbJMyD/ECxT5o/T843MbnYzBV234=", "BEDqbIoqicO03NPOsvBIaVzkH4QAaE+3dsDALB03IQg=", "1PWYLNn2Xvw8t/IP/Q9msGAg72GH/7B5mLJRjJ80P3U=", "xhiXsj84MWLSYlDhppkgSdec5ptKd97yJ4iHICJ5oJU="})); err != nil {
				fmt.Println(err)
			}
			fmt.Println("Done")

			// Wait at least 1 Momentum
			time.Sleep(31 * time.Second)

			// Set orchestrator info
			fmt.Println("Set orchestrator info")
			if err := z.Send(z.Client.BridgeApi.SetOrchestratorInfo(6, uint32(numPillars), 1, 10)); err != nil {
				fmt.Println(err)
			}
			fmt.Println("Done")

			// Wait at least 1 Momentum
			time.Sleep(31 * time.Second)
		}
	}

	for {
		// Checking if the TSS key has been set
		fmt.Printf("Checking if the TSS key has been set\n")
		tssCheck, err := z.Client.BridgeApi.GetBridgeInfo()
		if err != nil {
			fmt.Println(err)
			return
		} else if len(tssCheck.CompressedTssECDSAPubKey) > 0 {
			fmt.Printf("TSS key set, skipping\n")
			break
		} else {
			// Wait for pillars to produce enough momentus so the orchestrator can collect their public keys
			fmt.Printf("Wait %d momentums for pillars to produce enough momentus so the orchestrator can collect their public keys\n", (4 * numPillars))
			for {
				if momentum, err := z.Client.LedgerApi.GetFrontierMomentum(); err == nil {
					fmt.Printf("Current Momentum Height %d of %d ", momentum.Momentum.Height, (4 * numPillars))

					if momentum.Momentum.Height > uint64(4*numPillars) {
						break
					}
				}

				// Wait at least 1 Momentum
				time.Sleep(11 * time.Second)
			}
			fmt.Println("Done")

			// Start keygen
			fmt.Println("Start keygen")
			if err := z.Send(z.Client.BridgeApi.AllowKeygen()); err != nil {
				fmt.Println(err)
			}
			fmt.Println("Done")

			// Takes at least 1 minute for the keygen to finish
			time.Sleep(61 * time.Second)

			// Wait for the keygen to be done
			fmt.Println("Wait for the keygen to be done")
			tssPubKey := ""
			for {
				config, err := app.MakeConfig()

				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					return
				}

				fmt.Println("config.TssConfig.PublicKey: ", config.TssConfig.PublicKey)

				if config.TssConfig.PublicKey != "" {
					tssPubKey = config.TssConfig.PublicKey
					break
				}

				fmt.Println("Not found, trying again in 21s")

				// Try again later
				time.Sleep(31 * time.Second)
			}
			fmt.Println("Done")

			// Keygen is done, setting TSS Public Key
			fmt.Println("Keygen is done, setting TSS Public Key", tssPubKey)
			if err := z.Send(z.Client.BridgeApi.ChangeTssECDSAPubKey(tssPubKey, "", "", uint32(numPillars))); err != nil {
				fmt.Println("ChangeTssECDSAPubKey err:")
				fmt.Println(err)
			}
			fmt.Println("Done")

			// Wait at least 1 Momentum
			time.Sleep(31 * time.Second)
		}
	}

	for {
		// Checking if the bridge has already been configured
		fmt.Printf("Checking if the bridge has already been configured\n")
		networksCount, err := z.Client.BridgeApi.GetAllNetworks(0, 5)
		if err != nil {
			fmt.Println(err)
		} else if networksCount.Count > 0 {
			fmt.Printf("Bridge already configured, skipping\n")
			break
		} else {
			// Add Hardhat network (ignore the name "bsc")
			fmt.Println("Add Hardhat network")
			if err := z.Send(z.Client.BridgeApi.AddNetwork(2, 31337, "bsc", "0x8464135c8F25Da09e49BC8782676a84730C318bC", `{}`)); err != nil {
				fmt.Println(err)
			}
			fmt.Println("Done")

			// Wait at least 1 Momentum
			time.Sleep(31 * time.Second)

			// Add ZNN/wZNN token pair
			fmt.Println("Add ZNN/wZNN token pair")
			if err := z.Send(z.Client.BridgeApi.SetTokenPair(2, 31337, types.ZnnTokenStandard, "0x5FbDB2315678afecb367f032d93F642f64180aa3", true, true, false, big.NewInt(1e7), 100, 20, `{}`)); err != nil {
				fmt.Println(err)
			}
			fmt.Println("Done")

			// Add QSR/wQSR token pair
			fmt.Println("Add QSR/wQSR token pair")
			if err := z.Send(z.Client.BridgeApi.SetTokenPair(2, 31337, types.QsrTokenStandard, "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512", true, true, false, big.NewInt(1e8), 300, 35, `{}`)); err != nil {
				fmt.Println(err)
			}
			fmt.Println("Done")

			// Wait at least 1 Momentum
			time.Sleep(31 * time.Second)
		}
	}

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
		fmt.Println("KeySignThreshold: ", orchestratorInfo.KeySignThreshold)
	}

	fmt.Println("--------------------------------\nSecurity Info")
	securityInfo, err := z.Client.BridgeApi.GetSecurityInfo()
	if securityInfo == nil {
		fmt.Println("securityInfo nil")
	} else {
		fmt.Println("AdministratorDelay: ", securityInfo.AdministratorDelay)
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
	if false {
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
	}

	fmt.Println("--------------------------------------\nUnwraps")
	if true {
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
	}

	if err := z.Stop(); err != nil {
		zenon.CommonLogger.Error("Error while stopping Zenon SDK instance", "error", err)
	}
}

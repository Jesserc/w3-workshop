package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/w3types"
	"github.com/lmittmann/w3/w3vm"
)

func main() {
	// Connect to an Ethereum node
	client := w3.MustDial("https://rpc.ankr.com/eth")
	defer client.Close()

	// Create a VM instance
	vmInstance, err := w3vm.New(
		w3vm.WithFork(client, nil),
		w3vm.WithNoBaseFee(),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create a message
	msg := &w3types.Message{
		From:  common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e"),
		To:    w3.APtr("0x742d35Cc6634C0532925a3b844Bc454e4438f44f"),
		Value: w3.I("1 ether"),
	}

	// Create hooks
	hooks := &tracing.Hooks{
		OnEnter: func(depth int, typ byte, from common.Address, to common.Address, input []byte, gas uint64, value *big.Int) {
			fmt.Printf("Entering depth %d, from %s to %s, gas: %d, value: %s\n", depth, from, to, gas, value)
			fmt.Println()
		},
		OnExit: func(depth int, output []byte, gasUsed uint64, err error, reverted bool) {
			fmt.Printf("Exiting depth %d, gas used: %d, reverted: %v\n", depth, gasUsed, reverted)
			fmt.Println()
			if err != nil {
				fmt.Printf("Error during execution: %v\n", err)
				fmt.Println()
			}
		},
		OnOpcode: func(pc uint64, op byte, gas, cost uint64, scope tracing.OpContext, rData []byte, depth int, err error) {
			fmt.Printf("Executing opcode %s at PC %d, gas: %d, cost: %d\n", vm.OpCode(op), pc, gas, cost)
			fmt.Println()
		},
		OnGasChange: func(oldGas, newGas uint64, reason tracing.GasChangeReason) {
			fmt.Printf("Gas changed from %d to %d, reason: %v\n", oldGas, newGas, reason)
			fmt.Println()
		},
		OnBalanceChange: func(addr common.Address, prev, new *big.Int, reason tracing.BalanceChangeReason) {
			fmt.Printf("Balance changed for %s from %s to %s, reason: %v\n", addr, prev, new, reason)
			fmt.Println()
		},
	}

	// Apply the message with hooks
	receipt, err := vmInstance.Apply(msg, hooks)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error applying message: %v", err))
	}

	// Print the receipt
	fmt.Printf("Transaction receipt: %+v\n", receipt.GasUsed)
}

package main

import (
	"fmt"
	"log"

	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/debug"
	"github.com/lmittmann/w3/w3types"
)

var (
	fromAddr = w3.A("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")

	// ERC20 token contract address (using USDC on Ethereum mainnet as an example)
	tokenAddress = w3.A("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")

	// ERC20 transfer function signature
	balanceOf = w3.MustNewFunc("balanceOf(address)", "uint256")

	selector = balanceOf.Selector
)

func main() {
	// Connect to an Ethereum node
	client := w3.MustDial("https://docs-demo.quiknode.pro") // quick node demo rpc url
	defer client.Close()

	// Create a message for the ERC20 transfer
	msg := &w3types.Message{
		From: fromAddr, // Replace with your address
		To:   &tokenAddress,
		Func: balanceOf,
		Args: []any{
			fromAddr, // Pass fromAddr as argument for balanceOf(address)
		},
	}

	// Configure the tracer
	config := &debug.TraceConfig{
		EnableStack:   true,
		EnableMemory:  true,
		EnableStorage: true,
	}

	// Perform the trace call
	var trace *debug.Trace
	err := client.Call(
		debug.TraceCall(msg, nil, config).Returns(&trace),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Print trace results
	for _, structLog := range trace.StructLogs {
		opcode := structLog.Op
		fmt.Printf("Opcode: %s, Gas: %d, Depth: %d\n", opcode, structLog.Gas, structLog.Depth)
		stack := make([]string, 0)
		for _, s := range structLog.Stack {
			stack = append(stack, s.Hex())
		}
		fmt.Printf("  Stack: %v\n", stack)
		fmt.Printf("  Memory: %x\n", structLog.Memory)
		if len(structLog.Storage) > 0 {
			fmt.Printf("  Storage:\n")
			for k, v := range structLog.Storage {
				fmt.Printf("    %x: %x\n", k, v)
			}
		}
		fmt.Println()
	}
}

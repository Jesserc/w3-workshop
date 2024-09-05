package main

import (
	"fmt"
	"log"

	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/w3types"
	"github.com/lmittmann/w3/w3vm"
)

func main() {
	// Connect to RPC endpoint
	client := w3.MustDial("https://rpc.ankr.com/eth")
	defer client.Close()

	var (
		// Define addresses
		addrFrom = w3.A("0x742d35Cc6634C0532925a3b844Bc454e4438f44e")
		addrTo   = w3.A("0x1a2d3FE26634C0532925a3b844bc454e4438F44F")

		value = w3.I("10 ether")
	)

	// Create a VM that forks the Mainnet state from the latest block
	vm, err := w3vm.New(
		w3vm.WithFork(client, nil),
		w3vm.WithNoBaseFee(),
	)
	if err != nil {
		log.Fatalf("Failed to create VM: %v", err)
	}

	// Get balance of recipient before transfer
	balance, err := vm.Balance(addrTo)
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}
	fmt.Printf("Recipient balance before transfer: %.6s ETH\n", w3.FromWei(balance, 18))

	// Simulate a simple ETH transfer
	// We can even simulate a UniSwap v3 swap, see: w3/w3vm/vm_test.go
	receipt, err := vm.Apply(&w3types.Message{
		From:  addrFrom,
		To:    &addrTo,
		Value: value,
	})
	if err != nil {
		log.Fatalf("Failed to apply transaction: %v", err)
	}

	// Print transaction details
	fmt.Printf("Transaction successful:\n")
	fmt.Printf("  From: %s\n", addrFrom)
	fmt.Printf("  To: %s\n", addrTo)
	fmt.Printf("  Value: %s ETH\n", value)
	fmt.Printf("  Gas Used: %d\n", receipt.GasUsed)

	// Get balance of recipient after transfer
	balance, err = vm.Balance(addrTo)
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}

	fmt.Printf("Recipient balance after transfer: %.6s ETH\n", w3.FromWei(balance, 18))
}

package main

import (
	"fmt"
	"math/big"
	"os"

	"github.com/lmittmann/go-solc"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/w3types"
	"github.com/lmittmann/w3/w3vm"
)

// Your contract ABI
var (
	funcStore    = w3.MustNewFunc("store(uint256)", "")
	funcRetrieve = w3.MustNewFunc("retrieve()", "uint256")

	// Solidity compiler
	c = solc.New("0.8.25")
)

func main() {
	// 0. Get contract bytecode
	contract := c.MustCompile(".", "Store")
	contractBytecode := contract.DeployCode

	// 1. Set up a local VM environment
	vm, err := w3vm.New(
		w3vm.WithChainConfig(nil), // Use default chain config
		w3vm.WithNoBaseFee(),      // Disable base fee for simplicity
	)
	if err != nil {
		fmt.Printf("Failed to create vm instance, err: %v\n", err)
		os.Exit(1)
	}

	// 2. Deploy the contract
	deployerAddr := w3.A("0x742d35Cc6634C0532925a3b844Bc454e4438f44e")
	deployMsg := &w3types.Message{
		From:  deployerAddr,
		Input: contractBytecode,
	}

	receipt, err := vm.Apply(deployMsg)
	if err != nil {
		fmt.Printf("Failed to deploy contract, err: %v\n", err)
		os.Exit(1)
	}

	contractAddr := *receipt.ContractAddress
	fmt.Printf("Contract deployed at: %s\n", contractAddr)

	// 3. Interact with the deployed contract

	// Store a value
	storeMsg := &w3types.Message{
		From: deployerAddr,
		To:   &contractAddr,
		Func: funcStore,
		Args: []any{big.NewInt(42)},
	}

	_, err = vm.Apply(storeMsg)
	if err != nil {
		fmt.Printf("Failed to send store msg, err: %v\n", err)
		os.Exit(1)
	}

	// Retrieve the value
	retrieveMsg := &w3types.Message{
		From: deployerAddr,
		To:   &contractAddr,
		Func: funcRetrieve,
	}

	receipt, err = vm.Call(retrieveMsg)
	if err != nil {
		fmt.Printf("Failed to send retrieve msg, err: %v\n", err)
		os.Exit(1)
	}

	// 4. Verify the results
	var storedValue *big.Int
	if err := receipt.DecodeReturns(&storedValue); err != nil {
		fmt.Printf("Failed to decode returned value, err: %v\n", err)
	}

	fmt.Printf("Stored value: %s\n", storedValue)

	// Assert the value is correct
	if storedValue.Cmp(big.NewInt(42)) != 0 {
		fmt.Printf("Returned value is not equal to expected value, err: %v\n", err)
	}

	fmt.Println("Test passed successfully!")
}

package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/w3types"
	"github.com/lmittmann/w3/w3vm"
)

// Your contract ABI
var (
	funcStore    = w3.MustNewFunc("store(uint256)", "")
	funcRetrieve = w3.MustNewFunc("retrieve()", "uint256")
)

// Your contract bytecode (replace this with your actual bytecode)
var contractBytecode = common.FromHex("608060405234801561001057600080fd5b50610150806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632e64cec11461003b5780636057361d14610059575b600080fd5b610043610075565b60405161005091906100d9565b60405180910390f35b610073600480360381019061006e919061009d565b61007e565b005b60008054905090565b8060008190555050565b60008135905061009781610103565b92915050565b6000602082840312156100b3576100b26100fe565b5b60006100c184828501610088565b91505092915050565b6100d3816100f4565b82525050565b60006020820190506100ee60008301846100ca565b92915050565b6000819050919050565b600080fd5b61010c816100f4565b811461011757600080fd5b5056fea2646970667358221220404e37f487a89a932dca5e77faaf6ca2de3b991f93d230604b1b8daaef64766264736f6c63430008090033")

func main() {
	// 1. Set up a local VM environment
	vm, err := w3vm.New(
		w3vm.WithChainConfig(nil), // Use default chain config
		w3vm.WithNoBaseFee(),      // Disable base fee for simplicity
	)
	if err != nil {
		log.Fatal(err)
	}

	// 2. Deploy the contract
	deployerAddr := w3.A("0x742d35Cc6634C0532925a3b844Bc454e4438f44e")
	deployMsg := &w3types.Message{
		From:  deployerAddr,
		Input: contractBytecode,
	}

	receipt, err := vm.Apply(deployMsg)
	if err != nil {
		log.Fatal(err)
	}

	contractAddr := *receipt.ContractAddress
	fmt.Printf("Contract deployed at: %s\n", contractAddr.Hex())

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
		log.Fatal(err)
	}

	// Retrieve the value
	retrieveMsg := &w3types.Message{
		From: deployerAddr,
		To:   &contractAddr,
		Func: funcRetrieve,
	}

	receipt, err = vm.Call(retrieveMsg)
	if err != nil {
		log.Fatal(err)
	}

	// 4. Verify the results
	var storedValue *big.Int
	if err := receipt.DecodeReturns(&storedValue); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Stored value: %s\n", storedValue.String())

	// Assert the value is correct
	if storedValue.Cmp(big.NewInt(42)) != 0 {
		log.Fatal("Unexpected stored value")
	}

	fmt.Println("Test passed successfully!")
}

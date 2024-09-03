package main

import (
	"fmt"
	"math/big"

	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
)

func main() {
	// Connect to RPC endpoint
	client := w3.MustDial("https://rpc.ankr.com/eth")
	defer client.Close()

	var (
		addr = w3.A("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
		usdc = w3.A("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")

		// Declare a Smart Contract function using Solidity syntax,
		// no "abigen" and ABI JSON file needed.
		balanceOf = w3.MustNewFunc("balanceOf(address)", "uint256")

		// Declare variables for the RPC responses.
		ethBalance  *big.Int
		usdcBalance *big.Int
	)

	// Do batch request (both RPC requests are send in the same
	// HTTP request).
	if err := client.Call(
		eth.Balance(addr, nil).Returns(&ethBalance),
		eth.CallFunc(usdc, balanceOf, addr).Returns(&usdcBalance),
	); err != nil {
		fmt.Printf("Request failed: %v\n", err)
		return
	}

	fmt.Printf("Combined balance: %v wei",
		new(big.Int).Add(ethBalance, usdcBalance),
	)
}

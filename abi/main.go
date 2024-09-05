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
		addr  = w3.A("0xb1b2d032AA2F52347fbcfd08E5C3Cc55216E8404")
		weth9 = w3.A("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")

		// Declare a Smart Contract function using Solidity syntax,
		// no "abigen" and ABI JSON file needed.
		balanceOf = w3.MustNewFunc("balanceOf(address)", "uint256")

		// Declare variables for the RPC responses.
		ethBalance   *big.Int
		weth9Balance *big.Int
	)

	// Do batch request (both RPC requests are send in the same
	// HTTP request).
	if err := client.Call(
		eth.Balance(addr, nil).Returns(&ethBalance),
		eth.CallFunc(weth9, balanceOf, addr).Returns(&weth9Balance),
	); err != nil {
		fmt.Printf("Request failed: %v\n", err)
		return
	}

	fmt.Printf("Eth balance: %v\n", w3.FromWei(ethBalance, 18))
	fmt.Printf("Weth balance: %v\n", w3.FromWei(weth9Balance, 18))
}

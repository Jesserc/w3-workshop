package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
)

var (
	rpcUrl    = "https://rpc.ankr.com/eth"
	blockHash = common.HexToHash("0x3ec47b634bc7b7bf5f5e1dfb0ab4c5f1b0b3ada3d11a761b80ed5f462019639f")
)

func main() {
	// connect to RPC endpoint
	client := w3.MustDial(rpcUrl)
	defer client.Close()

	// fetch blocks in bulk
	var block *types.Block
	call := eth.BlockByHash(blockHash).Returns(&block)

	err := client.Call(call)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Printf("Block Number: %v\n", block.Number())
	fmt.Printf("Block transaction count: %v\n", len(block.Transactions()))
	fmt.Printf("Block Coinbase: %v\n", block.Coinbase())

	fmt.Println()
	fmt.Println("Processing first 5 transactions...")
	processBlock(block)
}

func processBlock(b *types.Block) {
	// Process the first 5 blocks
	for i, transaction := range b.Transactions() {
		if i == 5 {
			return
		}
		sender, err := types.Sender(types.LatestSignerForChainID(big.NewInt(1)), transaction)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("-%v:\n", i+1)
		fmt.Printf("Hash: %v\nFrom: %v\nTo: %v\nValue: %v ETH\n\n",
			transaction.Hash(),
			sender,
			transaction.To(),
			w3.FromWei(transaction.Value(), 18))
	}
}

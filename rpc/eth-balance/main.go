package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	w3eth "github.com/lmittmann/w3/module/eth"
)

func main() {
	// connect to RPC endpoint
	client := w3.MustDial("https://rpc.ankr.com/eth")
	defer client.Close()

	var balance *big.Int
	err := client.Call(w3eth.Balance(w3.A("0x6058A1cDdeC5873c0b116b8F0A528bCb6aBc05dA"), nil).Returns(&balance))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Balance for %v: %.4v Ether\n", common.HexToAddress("0x6058A1cDdeC5873c0b116b8F0A528bCb6aBc05dA"), w3.FromWei(balance, 18))
}

package main

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/debug"
	"github.com/lmittmann/w3/w3types"
)

func main() {
	client := w3.MustDial("https://rpc.ankr.com/eth")
	defer client.Close()

	msg := &w3types.Message{
		From:  common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e"),
		To:    w3.APtr("0x742d35Cc6634C0532925a3b844Bc454e4438f44f"),
		Value: w3.I("1 ether"),
	}

	var trace *debug.CallTrace
	err := client.Call(
		debug.CallTraceCall(msg, nil, nil).Returns(&trace),
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("From: %s\n", trace.From.Hex())
	fmt.Printf("To: %s\n", trace.To.Hex())
	fmt.Printf("Value: %s\n", trace.Value.String())
	fmt.Printf("Gas Used: %d\n", trace.GasUsed)
	fmt.Printf("Error: %s\n", trace.Error)
	fmt.Printf("Revert Reason: %s\n", trace.RevertReason)

	printNestedCalls(trace.Calls, 1)
}

func printNestedCalls(calls []*debug.CallTrace, depth int) {
	for _, call := range calls {
		fmt.Printf("%sCall to %s, Type: %s, Gas Used: %d\n",
			strings.Repeat("  ", depth), call.To.Hex(), call.Type, call.GasUsed)
		printNestedCalls(call.Calls, depth+1)
	}
}

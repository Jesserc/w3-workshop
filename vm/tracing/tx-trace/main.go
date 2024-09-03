package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/goccy/go-json"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/debug"
	"github.com/lmittmann/w3/w3types"
)

func main() {
	client := w3.MustDial("https://docs-demo.quiknode.pro/")

	defer client.Close()

	msg := &w3types.Message{
		From:  common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e"),
		To:    w3.APtr("0x742d35Cc6634C0532925a3b844Bc454e4438f44f"),
		Value: w3.I("1 ether"),
	}

	config := &debug.TraceConfig{
		EnableStack:   true,
		EnableMemory:  true,
		EnableStorage: true,
	}

	var trace *debug.Trace
	err := client.Call(
		debug.TraceCall(msg, nil, config).Returns(&trace),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Gas used: %d\n", trace.Gas)
	fmt.Printf("Execution failed: %v\n", trace.Failed)
	fmt.Printf("Output: %x\n", trace.Output)
	// for i, structLog := range trace.StructLogs {
	// 	fmt.Printf("Step %d: Opcode %s, Gas: %d, Depth: %d\n", i, structLog.Op, structLog.Gas, structLog.Depth)
	// }

	indent, err := json.MarshalIndent(trace.StructLogs, "", "\t")
	if err != nil {
		return
	}

	fmt.Println(string(indent))
}

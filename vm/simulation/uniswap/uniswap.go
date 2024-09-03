package main

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/debug"
	"github.com/lmittmann/w3/w3types"
	"github.com/lmittmann/w3/w3vm"
)

func main() {
	var (
		addrEOA    = w3.A("0x000000000000000000000000000000000000c0Fe")
		addrWETH   = w3.A("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
		addrUNI    = w3.A("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984")
		addrRouter = w3.A("0xE592427A0AEce92De3Edee1F18E0157C05861564")

		funcExactInput = w3.MustNewFunc(`exactInput(
			(
			bytes path,
			address recipient,
			uint256 deadline,
			uint256 amountIn,
			uint256 amountOutMinimum
			) params
		)`, "uint256 amountOut")
	)

	type ExactInputParams struct {
		Path             []byte
		Recipient        common.Address
		Deadline         *big.Int
		AmountIn         *big.Int
		AmountOutMinimum *big.Int
	}

	encodePath := func(tokenA common.Address, fee uint32, tokenB common.Address) []byte {
		path := make([]byte, 43)
		copy(path, tokenA[:])
		path[20], path[21], path[22] = byte(fee>>16), byte(fee>>8), byte(fee)
		copy(path[23:], tokenB[:])
		return path
	}

	client := w3.MustDial("https://docs-demo.quiknode.pro/")
	defer client.Close()

	// 1. Create a VM that forks the Mainnet state from the latest block,
	// disables the base fee, and has a fake WETH balance and approval for the router
	vm, err := w3vm.New(
		w3vm.WithFork(client, nil),
		w3vm.WithNoBaseFee(),
		w3vm.WithState(w3types.State{ // create state
			addrWETH: {Storage: map[common.Hash]common.Hash{
				w3vm.WETHBalanceSlot(addrEOA):               common.BigToHash(w3.I("1 ether")),
				w3vm.WETHAllowanceSlot(addrEOA, addrRouter): common.BigToHash(w3.I("1 ether")),
			}},
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	// 2. Simulate a UniSwap v3 swap
	msg := &w3types.Message{
		From: addrEOA,
		To:   &addrRouter,
		Func: funcExactInput,
		Args: []any{&ExactInputParams{
			Path:             encodePath(addrWETH, 500, addrUNI),
			Recipient:        addrEOA,
			Deadline:         big.NewInt(time.Now().Unix()),
			AmountIn:         w3.I("1 ether"),
			AmountOutMinimum: w3.Big0,
		}}}
	receipt, err := vm.Apply(msg)
	if err != nil {
		log.Fatal(err)
	}

	// 3. Decode output amount
	var amountOut *big.Int
	if err := receipt.DecodeReturns(&amountOut); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("amount out: %s UNI\n", w3.FromWei(amountOut, 18))

	// Print event from executing the tx
	/*for _, lg := range receipt.Logs {
		jsonBytes, err := lg.MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}

		b := bytes.Buffer{}
		err = json.Indent(&b, jsonBytes, "", "\t")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(&b)
	}*/

	// Trace
	config := &debug.TraceConfig{
		EnableStack:   true,
		EnableMemory:  true,
		EnableStorage: true,
	}

	var trace *debug.Trace
	err = client.Call(
		debug.TraceCall(msg, nil, config).Returns(&trace),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Gas used: %d\n", trace.Gas)
	fmt.Printf("Execution failed: %v\n", trace.Failed)
	fmt.Printf("Output: %x\n", trace.Output)
	for i, structLog := range trace.StructLogs {
		fmt.Printf("Step %d: Opcode %s, Gas: %d, Depth: %d\n", i, structLog.Op, structLog.Gas, structLog.Depth)
	}
}

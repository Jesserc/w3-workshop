package main

import (
	"fmt"

	"github.com/lmittmann/w3"
)

func main() {
	addr := w3.A("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")
	hash := w3.H("0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3")
	bytes := w3.B("0x27c5342c")
	amount := w3.I("12.34 ether")

	fmt.Printf(
		"\n"+
			"Address: %v\n"+
			"Hash: %v\n"+
			"Bytes: %v\n"+
			"Amount in Wei: %v\n", addr, hash, bytes, amount)
}

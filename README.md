# w3-workshop

This repository contains a collection of Go examples and utilities for interacting with Ethereum using the w3 library.

## Directory Structure

```
w3-workshop/
│
├── abi/
│   └── main.go
│
├── rpc/
│   ├── eth-balance/
│   │   └── main.go
│   └── eth-block-by-hash/
│       └── main.go
│
├── utilities/
│   └── main.go
│
├── vm/
│   ├── simulation/
│   │   └── uniswap/
│   │       ├── uniswap.go
│   │       └── main.go
│   │
│   ├── testing/
│   │   └── main.go
│   │
│   └── tracing/
│       ├── trace-hooks/
│       │   └── main.go
│       └── tx-trace/
│           ├── _main.go
│           └── main.go
│
└── .gitignore
```

## Project Overview

This workshop is designed to showcase various aspects of Ethereum development using the w3 library in Go. Here's a brief overview of each directory:

### abi/
Contains examples of working with Ethereum ABI (Application Binary Interface).

### rpc/
Demonstrates making RPC calls to Ethereum nodes:
- `eth-balance/`: Example of checking ETH balance
- `eth-block-by-hash/`: Example of fetching block information by its hash

### utilities/
Utility functions and helpers for Ethereum development.

### vm/
Ethereum Virtual Machine related examples and simulations:
- `simulation/`: Contains simulations of Ethereum transactions and contract interactions
    - `uniswap/`: Specific example simulating Uniswap interactions
- `testing/`: Examples of testing smart contracts using w3
- `tracing/`: Tools and examples for tracing Ethereum transactions
    - `trace-hooks/`: Examples of using trace hooks
    - `tx-trace/`: Transaction tracing utilities

## Getting Started

1. Clone this repository
2. Ensure you have Go installed (version 1.16 or later recommended)
3. Install dependencies:
   ```
   go mod tidy
   ```
4. Navigate to any of the example directories and run the Go files, for example:
   ```
   cd rpc/eth-balance
   go run main.go
   ```

<div align="center">
  <h1>Coderouter</h1>
  <p><strong>High-Performance Uniswap V4 Optimal Route Finder</strong></p>
</div>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.26+-00ADD8?style=for-the-badge&logo=go" alt="Go Version">
  <img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" alt="License">
</p>

## Overview

Coderouter is an open-source, high-performance Go backend service designed to find optimal token swap routes exclusively on **Uniswap V4**. By leveraging Go's robust concurrency model (worker pools and goroutines), it concurrently analyzes multiple Uniswap V4 pool hops to guarantee the best possible execution path with minimal latency.

## Key Features

- **Lightning Fast:** Highly optimized route finding engine using advanced Go concurrency patterns.
- **Uniswap V4 Native:** Purpose-built for the intricacies of Uniswap V4 (Hooks, Singleton Architecture, Flash Accounting).
- **Memory Efficient:** Minimal allocations and zero-allocation critical paths.
- **REST API:** Clean and simple HTTP interface for seamless integration into any frontend or trading bot.

## Quick Start

### Prerequisites
- Go 1.26 or higher
- Access to an Ethereum RPC endpoint

### Installation

```bash
git clone https://github.com/coderouter/coderouter.git
cd coderouter
```

### Usage

1. Set your Ethereum RPC URL:
```bash
export RPC_URL="https://eth.llamarpc.com"
```

2. Run the server:
```bash
go run ./cmd/coderouter
```

3. Query a route:
```bash
curl -X GET "http://localhost:8080/quote?tokenIn=0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48&tokenOut=0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2&amount=1000000000"
```

## Architecture

Coderouter follows strict clean architecture and idiomatic Go principles:

- `cmd/`: Application entry points.
- `internal/api/`: Fast REST HTTP handlers.
- `internal/aggregator/`: The core concurrency engine and route comparison logic.
- `internal/dex/uniswap/`: Uniswap V4 specific RPC and state management.
- `internal/domain/`: Core business models.

## License

MIT License. See `LICENSE` for details.

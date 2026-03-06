package domain

import (
	"math/big"
)

type Token struct {
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	Decimals uint8  `json:"decimals"`
}

type Pool struct {
	Address   string
	Token0    Token
	Token1    Token
	Fee       *big.Int
	Tick      int
	Liquidity *big.Int
}

type Route struct {
	Pools        []Pool
	Tokens       []Token
	AmountIn     *big.Int
	AmountOut    *big.Int
	EstimatedGas *big.Int
}

type QuoteRequest struct {
	TokenIn  string   `json:"tokenIn"`
	TokenOut string   `json:"tokenOut"`
	Amount   *big.Int `json:"amount"`
}

type QuoteResponse struct {
	AmountOut   *big.Int `json:"amountOut"`
	Path        []string `json:"path"`
	GasEstimate uint64   `json:"gasEstimate"`
}

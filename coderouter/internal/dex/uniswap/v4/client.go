package v4

import (
	"context"
	"math/big"

	"github.com/coderouter/coderouter/internal/domain"
)

type Client struct {
	rpcURL string
}

func NewClient(rpcURL string) *Client {
	return &Client{
		rpcURL: rpcURL,
	}
}

func (c *Client) GetQuote(ctx context.Context, req *domain.QuoteRequest) (*domain.QuoteResponse, error) {
	return &domain.QuoteResponse{
		AmountOut:   big.NewInt(0),
		Path:        []string{},
		GasEstimate: 0,
	}, nil
}

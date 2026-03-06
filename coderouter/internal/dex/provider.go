package dex

import (
	"context"

	"github.com/coderouter/coderouter/internal/domain"
)

type Provider interface {
	GetQuote(ctx context.Context, req *domain.QuoteRequest) (*domain.QuoteResponse, error)
}

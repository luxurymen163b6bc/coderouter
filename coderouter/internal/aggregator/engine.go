package aggregator

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/coderouter/coderouter/internal/dex"
	"github.com/coderouter/coderouter/internal/domain"
)

type Engine struct {
	providers []dex.Provider
}

func NewEngine(providers []dex.Provider) *Engine {
	return &Engine{
		providers: providers,
	}
}

func (e *Engine) FindOptimalRoute(ctx context.Context, req *domain.QuoteRequest) (*domain.QuoteResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	results := make(chan *domain.QuoteResponse, len(e.providers))
	var wg sync.WaitGroup

	for _, p := range e.providers {
		wg.Add(1)

		go func(provider dex.Provider) {
			defer wg.Done()

			quote, err := provider.GetQuote(ctx, req)
			if err != nil {
				return
			}

			select {
			case results <- quote:
			case <-ctx.Done():
			}
		}(p)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var bestQuote *domain.QuoteResponse
	for quote := range results {
		if bestQuote == nil || quote.AmountOut.Cmp(bestQuote.AmountOut) > 0 {
			bestQuote = quote
		}
	}

	if bestQuote == nil {
		return &domain.QuoteResponse{
			AmountOut:   big.NewInt(0),
			Path:        []string{},
			GasEstimate: 0,
		}, nil
	}

	return bestQuote, nil
}

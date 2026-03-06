package api

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/coderouter/coderouter/internal/aggregator"
	"github.com/coderouter/coderouter/internal/domain"
)

type Server struct {
	router *http.ServeMux
	engine *aggregator.Engine
}

func NewServer(engine *aggregator.Engine) *Server {
	s := &Server{
		router: http.NewServeMux(),
		engine: engine,
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.router.HandleFunc("/quote", s.handleQuote())
}

func (s *Server) handleQuote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		query := r.URL.Query()
		tokenIn := query.Get("tokenIn")
		tokenOut := query.Get("tokenOut")
		amountStr := query.Get("amount")

		if tokenIn == "" || tokenOut == "" || amountStr == "" {
			http.Error(w, "Missing required query parameters", http.StatusBadRequest)
			return
		}

		amount, ok := new(big.Int).SetString(amountStr, 10)
		if !ok {
			http.Error(w, "Invalid amount", http.StatusBadRequest)
			return
		}

		req := &domain.QuoteRequest{
			TokenIn:  tokenIn,
			TokenOut: tokenOut,
			Amount:   amount,
		}

		quote, err := s.engine.FindOptimalRoute(context.Background(), req)
		if err != nil {
			http.Error(w, "Error fetching quote", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(quote); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

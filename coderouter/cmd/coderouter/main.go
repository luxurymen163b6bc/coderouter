package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/coderouter/coderouter/internal/aggregator"
	"github.com/coderouter/coderouter/internal/api"
	"github.com/coderouter/coderouter/internal/config"
	"github.com/coderouter/coderouter/internal/dex"
	v4 "github.com/coderouter/coderouter/internal/dex/uniswap/v4"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	uniswapV4Client := v4.NewClient(cfg.RPCURL)

	providers := []dex.Provider{
		uniswapV4Client,
	}

	engine := aggregator.NewEngine(providers)
	server := api.NewServer(engine)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: server,
	}

	go func() {
		log.Printf("Starting Coderouter on port %s...", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}

package config

import (
	"os"
)

type Config struct {
	RPCURL   string
	Port     string
	LogLevel string
}

func Load() (*Config, error) {
	rpcUrl := os.Getenv("RPC_URL")
	if rpcUrl == "" {
		rpcUrl = "https://eth.llamarpc.com"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	return &Config{
		RPCURL:   rpcUrl,
		Port:     port,
		LogLevel: logLevel,
	}, nil
}

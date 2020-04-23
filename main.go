package main

import (
	"github.com/syrilster/go-currency-exchange-service/internal"
	"github.com/syrilster/go-currency-exchange-service/internal/config"
)

func main() {
	cfg := config.NewApplicationConfig()
	server := internal.SetupServer(cfg)
	server.Start("", cfg.ServerPort())
}

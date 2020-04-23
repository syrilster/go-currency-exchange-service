package internal

import (
	"fmt"
	"github.com/syrilster/go-currency-exchange-service/internal/config"
	"github.com/syrilster/go-currency-exchange-service/internal/currencyexchange"
	"github.com/syrilster/go-currency-exchange-service/internal/exchange"
	"github.com/syrilster/go-currency-exchange-service/internal/middlewares"
	"net/http"
)

func StatusRoute() (route config.Route) {
	route = config.Route{
		Path:    "/status",
		Method:  http.MethodGet,
		Handler: middlewares.RuntimeHealthCheck(),
	}
	return route
}

type ServerConfig interface {
	Version() string
	BaseURL() string
	CurrencyExchangeClient() exchange.ClientInterface
}

func SetupServer(cfg ServerConfig) *config.Server {
	basePath := fmt.Sprintf("/%v", cfg.Version())
	currencyExchangeService := currencyexchange.NewService(cfg.CurrencyExchangeClient())
	server := config.NewServer().
		WithRoutes(
			"", StatusRoute(),
		).
		WithRoutes(
			basePath,
			currencyexchange.Route(currencyExchangeService),
		)
	return server
}

package config

import (
	"github.com/syrilster/go-currency-exchange-service/internal/customhttp"
	"github.com/syrilster/go-currency-exchange-service/internal/exchange"
	"net/http"
	"time"
)

type ApplicationConfig struct {
	envValues              *envConfig
	currencyExchangeClient exchange.ClientInterface
}

func (cfg *ApplicationConfig) Version() string {
	return cfg.envValues.Version
}

func (cfg *ApplicationConfig) ServerPort() int {
	return cfg.envValues.ServerPort
}

func (cfg *ApplicationConfig) BaseURL() string {
	return cfg.envValues.BaseUrl
}

func (cfg *ApplicationConfig) CurrencyExchangeClient() exchange.ClientInterface {
	return cfg.currencyExchangeClient
}

func NewApplicationConfig() *ApplicationConfig {
	envValues := newEnvironmentConfig()
	httpCommand := NewHttpCommand()
	ceClient := exchange.NewClient(envValues.CurrencyExchangeEndpoint, httpCommand, envValues.AppID)

	return &ApplicationConfig{
		envValues:              envValues,
		currencyExchangeClient: ceClient,
	}
}

func NewHttpCommand() customhttp.HTTPCommand {
	httpCommand := customhttp.New(
		customhttp.WithHTTPClient(&http.Client{Timeout: 5 * time.Second}),
	).Build()

	return httpCommand
}

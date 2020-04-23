package currencyexchange

import (
	"context"
	"github.com/syrilster/go-currency-exchange-service/internal/config"
	"net/http"
)

type ExchangeRateFetcher interface {
	FetchExchangeRate(ctx context.Context, req Request) (*Response, error)
}

func Route(fetcher ExchangeRateFetcher) (route config.Route) {
	route = config.Route{
		Path:    "/currency-exchange/from/{from}/to/{to}",
		Method:  http.MethodGet,
		Handler: handler(fetcher),
	}

	return route
}

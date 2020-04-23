package currencyexchange

import (
	"github.com/syrilster/go-currency-exchange-service/internal/util"
	"net/http"
)

func handler(fetcher ExchangeRateFetcher) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		util.WithBodyAndStatus("OK", http.StatusOK, w)
	}
}

package currencyexchange

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/syrilster/go-currency-exchange-service/internal/util"
	"net/http"
)

func handler(fetcher ExchangeRateFetcher) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		contextLogger := log.WithContext(ctx)

		request := parseFetchRequest(r)
		resp, err := fetcher.FetchExchangeRate(ctx, request)
		if err != nil {
			contextLogger.WithError(err).Error("Failed to fetch currency exchange rates")
			util.WithBodyAndStatus(nil, http.StatusInternalServerError, w)
			return
		}
		util.WithBodyAndStatus(resp, http.StatusOK, w)
	}
}

func parseFetchRequest(r *http.Request) Request {
	fromCurrency := mux.Vars(r)["from"]
	toCurrency := mux.Vars(r)["to"]

	return Request{
		FromCurrency: fromCurrency,
		ToCurrency:   toCurrency,
	}
}

package middlewares

import (
	"github.com/syrilster/go-currency-exchange-service/internal/util"
	"net/http"
)

func RuntimeHealthCheck() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		util.WithBodyAndStatus("All OK", http.StatusOK, w)
	}
}

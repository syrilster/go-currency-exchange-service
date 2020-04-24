package currencyexchange

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/syrilster/go-currency-exchange-service/internal/exchange"
	"strings"
)

type Service struct {
	client exchange.ClientInterface
}

func NewService(c exchange.ClientInterface) *Service {
	return &Service{client: c}
}

func (s Service) FetchExchangeRate(ctx context.Context, req Request) (*Response, error) {
	ctxLogger := log.WithContext(ctx)
	ctxLogger.Infof("Calling the exchange (openexchangerates) fetch API")

	request := toExchangeRequest(req)
	resp, err := s.client.GetExchangeRate(ctx, request)
	if err != nil {
		ctxLogger.Infof("Failed to fetch currency exchange rate: %v", err)
		return nil, err
	}
	response := unMarshallExchangeRate(resp, req)
	return response, nil
}

func unMarshallExchangeRate(resp *exchange.Response, req Request) *Response {
	var fromCurrency = req.FromCurrency
	var toCurrency = req.ToCurrency
	var conversionMultiple string
	var exchangeRate float64
	if strings.EqualFold(fromCurrency, "USD") {
		exchangeRate = getRateForCurrency(resp.Rates, toCurrency)
		conversionMultiple = fmt.Sprintf("%f", exchangeRate)
	} else if strings.EqualFold(toCurrency, "USD") {
		exchangeRate = getRateForCurrency(resp.Rates, fromCurrency)
		conversionMultiple = fmt.Sprintf("%f", float64(1)/exchangeRate)
	} else {
		// FromCurrency to USD and then USD to toCurrency
		exchangeRate = getRateForCurrency(resp.Rates, toCurrency)
		usdToFromCurrency := getRateForCurrency(resp.Rates, fromCurrency)
		toCurrencyToUSD := float64(1) / exchangeRate
		foreignCurrencyFactor := float64(1) / usdToFromCurrency
		conversionMultiple = fmt.Sprintf("%f", foreignCurrencyFactor/toCurrencyToUSD)
	}

	return &Response{
		FromCurrency: req.FromCurrency,
		ToCurrency:   req.ToCurrency,
		ConvMultiple: conversionMultiple,
	}
}

func getRateForCurrency(rates map[string]interface{}, currency string) float64 {
	var exchangeRate float64
	for key, rate := range rates {
		if strings.EqualFold(key, currency) {
			exchangeRate = rate.(float64)
			break
		}
	}
	return exchangeRate
}

func toExchangeRequest(request Request) exchange.Request {
	return exchange.Request{
		FromCurrency: request.FromCurrency,
		ToCurrency:   request.ToCurrency,
	}
}

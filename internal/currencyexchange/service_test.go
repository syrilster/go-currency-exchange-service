package currencyexchange

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/syrilster/go-currency-exchange-service/internal/exchange"
	"testing"
)

type MockExchangeRateClient struct {
	mock.Mock
}

func (client MockExchangeRateClient) GetExchangeRate(ctx context.Context, request exchange.Request) (*exchange.Response, error) {
	args := client.Called(ctx, request)
	return args.Get(0).(*exchange.Response), args.Error(1)
}

func TestCurrencyExchange(t *testing.T) {
	var rates = make(map[string]interface{})
	rates["AED"] = 3.6732
	rates["AUD"] = 1.539603
	rates["INR"] = 76.245
	rates["JPY"] = 107.5285

	var rateTests = []struct {
		currencyExchangeRequest exchange.Request
		expected                *Response
	}{
		{exchange.Request{
			FromCurrency: "AUD",
			ToCurrency:   "INR",
		}, &Response{
			FromCurrency: "AUD",
			ToCurrency:   "INR",
			ConvMultiple: "49.522507",
		}},
		{exchange.Request{
			FromCurrency: "USD",
			ToCurrency:   "AED",
		}, &Response{
			FromCurrency: "USD",
			ToCurrency:   "AED",
			ConvMultiple: "3.673200",
		}},
		{exchange.Request{
			FromCurrency: "JPY",
			ToCurrency:   "USD",
		}, &Response{
			FromCurrency: "JPY",
			ToCurrency:   "USD",
			ConvMultiple: "0.009300",
		}},
	}

	var mockResponse = &exchange.Response{
		Base:  "USD",
		Rates: rates,
	}

	t.Run("success", func(t *testing.T) {

		for _, tt := range rateTests {
			mockClient := new(MockExchangeRateClient)
			mockClient.On("GetExchangeRate", context.Background(), tt.currencyExchangeRequest).Return(mockResponse, nil)
			service := NewService(mockClient)
			req := Request{
				FromCurrency: tt.currencyExchangeRequest.FromCurrency,
				ToCurrency:   tt.currencyExchangeRequest.ToCurrency,
			}
			actual, err := service.FetchExchangeRate(context.Background(), req)
			if actual.ConvMultiple != tt.expected.ConvMultiple {
				t.Errorf("TestCurrencyExchange(%v): expected %v, actual %v", tt.currencyExchangeRequest, tt.expected, actual)
			}
			assert.Nil(t, err)
			mockClient.AssertExpectations(t)
		}
	})
}

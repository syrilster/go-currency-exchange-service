package currencyexchange

import (
	"context"
	"github.com/syrilster/go-currency-exchange-service/internal/exchange"
)

type Service struct {
	client exchange.ClientInterface
}

func (s Service) FetchExchangeRate(ctx context.Context, req Request) (*Response, error) {
	panic("implement me")
}

func NewService(c exchange.ClientInterface) *Service {
	return &Service{client: c}
}

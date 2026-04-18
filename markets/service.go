package markets

import (
	"net/http"
	"time"
)

type Service interface {
	GetAllMarkets() (response []MarketsResponse, err error)
}

type service struct {
	client *http.Client
}

func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *service) GetAllMarkets() (response []MarketsResponse, err error) {
	return
}

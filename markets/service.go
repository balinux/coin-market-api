package markets

import (
	"encoding/json"
	"net/http"
	"time"

	"go-mrt-schedule/common/client"
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
	url := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return nil, err
	}

	var markets []Markets
	err = json.Unmarshal(byteResponse, &markets)

	for _, item := range markets {
		response = append(response, MarketsResponse{
			Name:         item.Name,
			Image:        item.Image,
			CurrentPrice: item.CurrentPrice,
		})
	}

	return
}

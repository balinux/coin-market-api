package markets

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"coin-market-api/common/client"
)

type Service interface {
	GetAllMarkets() (response []MarketsResponse, err error)
	GetCoinDetail(id string) (response CoinDetailResponse, err error)
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

func (s *service) GetCoinDetail(id string) (response CoinDetailResponse, err error) {
	url := "https://api.coingecko.com/api/v3/coins/" + id
	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return CoinDetailResponse{}, err
	}

	var result CoinDetail

	err = json.Unmarshal(byteResponse, &result)
	if err != nil {
		return CoinDetailResponse{}, err
	}

	price, ok := result.MarketData.CurrentPrice["usd"]
	if !ok {
		return CoinDetailResponse{}, fmt.Errorf("usd price not found")
	}
	return CoinDetailResponse{
		Name:         result.Name,
		Image:        result.Image.Large,
		CurrentPrice: price,
		Description:  result.Description.EN,
		Categories:   result.Categories,
	}, nil
}

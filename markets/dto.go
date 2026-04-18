package markets

type Markets struct {
	Name         string  `json:"name"`
	Image        string  `json:"image"`
	CurrentPrice float64 `json:"current_price"`
}

type MarketsResponse struct {
	Name         string  `json:"name"`
	Image        string  `json:"image"`
	CurrentPrice float64 `json:"current_price"`
}

//	type CoinDetail struct {
//		Name string `json:"name"`
//
//		Image struct {
//			Large string `json:"large"`
//		} `json:"image"`
//
//		MarketData struct {
//			CurrentPrice map[string]float64 `json:"current_price"`
//		} `json:"market_data"`
//	}
type CoinDetailResponse struct {
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	CurrentPrice float64  `json:"current_price"`
	Description  string   `json:"description,omitempty"`
	Categories   []string `json:"categories,omitempty"`
}

type CoinDetail struct {
	Name        string         `json:"name"`
	Image       CoinImage      `json:"image"`
	MarketData  CoinMarketData `json:"market_data"`
	Categories  []string       `json:"categories"`
	Description struct {
		EN string `json:"en"`
	} `json:"description"`
}

type CoinImage struct {
	Large string `json:"large"`
}

type CoinMarketData struct {
	CurrentPrice map[string]float64 `json:"current_price"`
}

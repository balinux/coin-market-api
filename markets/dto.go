package markets

type Markets struct {
	Name         string `json:"name"`
	Image        string `json:"image"`
	CurrentPrice int    `json:"current_price"`
}

type MarketsResponse struct {
	Name         string `json:"name"`
	Image        string `json:"image"`
	CurrentPrice int    `json:"current_price"`
}

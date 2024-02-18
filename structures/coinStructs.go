package structures

import "time"

type Coins struct {
	LastUpdate time.Time `json:"lastUpdate"`
	AllCoins   []Coin    `json:"coins"`
}

type Coin struct {
	ID                           string    `json:"id"`
	Symbol                       string    `json:"symbol"`
	Name                         string    `json:"name"`
	Image                        string    `json:"image"`
	CurrentPrice                 float64   `json:"current_price"`
	MarketCap                    float64   `json:"market_cap"`
	MarketCapRank                float64   `json:"market_cap_rank"`
	FullyDilutedValuation        float64   `json:"fully_diluted_valuation"`
	TotalVolume                  float64   `json:"total_volume"`
	High24h                      float64   `json:"high_24h"`
	Low24h                       float64   `json:"low_24h"`
	PriceChange24h               float64   `json:"price_change_24h"`
	PriceChangePercentage24h     float64   `json:"price_change_percentage_24h"`
	MarketCapChange24h           float64   `json:"market_cap_change_24h"`
	MarketCapChangePercentage24h float64   `json:"market_cap_change_percentage_24h"`
	CirculatingSupply            float64   `json:"circulating_supply"`
	TotalSupply                  float64   `json:"total_supply"`
	MaxSupply                    float64   `json:"max_supply"`
	Ath                          float64   `json:"ath"`
	AthChangePercentage          float64   `json:"ath_change_percentage"`
	AthDate                      time.Time `json:"ath_date"`
	Atl                          float64   `json:"atl"`
	AtlChangePercentage          float64   `json:"atl_change_percentage"`
	AtlDate                      time.Time `json:"atl_date"`
	ROI                          ROI       `json:"roi"`
	LastUpdated                  time.Time `json:"last_updated"`
}

type ROI struct {
	Times      float64 `json:"times"`
	Currency   string  `json:"currency"`
	Percentage float64 `json:"percentage"`
}

package coincap

import "fmt"

type AssertsResponse struct {
	Assets    []AssetData `json:"data"`
	Timestamp int64       `json: "timestamp"`
}

type AssertResponse struct {
	Asset     AssetData `json:"data"`
	Timestamp int64     `json: "timestamp"`
}

type AssetData struct {
	Id                string `json:"id"`
	Rank              string `json:"rank"`
	Symbol            string `json: "symbol"`
	Name              string `json: "name"`
	Supply            string `json: "supply"`
	MaxSupply         string `json: "maxSupply"`
	MarketCapUsd      string `json: "marketCapUsd"`
	VolumeUsd24Hr     string `json: "volumeUsd24Hr"`
	PriceUsd          string `json: "priceUsd"`
	ChangePercent24Hr string `json: "changePercent24Hr"`
	Vwap24Hr          string `json: "vwap24Hr"`
}

func (d *AssetData) Info() string {
	return fmt.Sprintf("[ID] %s | [RANK] %s | [SYMBOL] %s | [PRICE] %s", d.Id, d.Rank, d.Symbol, d.PriceUsd)
}

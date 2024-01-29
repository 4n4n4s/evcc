package smartenergy

import (
	"encoding/json"
	"time"
)

const RegionURI = "https://apis.smartenergy.at/market/v1/price"

type Prices struct {
	Data []PriceInfo
}

type PriceInfo struct {
	StartTimestamp time.Time `json:"start_timestamp"`
	EndTimestamp   time.Time `json:"end_timestamp"`
	Marketprice    float64   `json:"marketprice"`
	Unit           string    `json:"unit"`
}

func (p *PriceInfo) UnmarshalJSON(data []byte) error {
	var s struct {
		StartTimestamp string  `json:"date"`
		Marketprice    float64 `json:"value"`
		Unit           string  `json:"unit"`
	}

	err := json.Unmarshal(data, &s)
	format := "2006-01-02T15:04:05-07:00"
	startDate, _ := time.Parse(format, s.StartTimestamp)
	endDate := startDate.Add(time.Minute*14 + time.Second*59) // prices are in 15 minute intervals
	if err == nil {
		p.StartTimestamp = startDate
		p.EndTimestamp = endDate
		p.Marketprice = s.Marketprice / 100 // price in ct/kWh
		p.Unit = "ct/kWh"
	}

	return err
}

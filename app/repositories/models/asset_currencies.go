package models

import "time"

type AssetCurrencies struct {
	Base
	ApiID      int64
	Exchange   string
	Currency   string
	Type       int64
	Amount     float64
	Price      float64
	Worth      float64
	TimeWindow time.Time
	CreatedAt  time.Time
}

func (t *AssetCurrencies) TableName() string {
	return "asset_currencies"
}



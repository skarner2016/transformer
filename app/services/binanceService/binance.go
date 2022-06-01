package binanceService

import "transformer/app/repositories/models"

type Binance struct {
	ApiID          int64
	Exchange       string
	SymbolPriceMap map[string]float64
	Data           *BinanceData
}

type BinanceData struct {
	SpotAssetCurrency *[]models.AssetCurrencies
}

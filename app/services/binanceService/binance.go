package binanceService

import (
	"sync"
	"time"
	"transformer/app/enums/exchange"
	"transformer/app/repositories/models"
)

type BinanceService struct {
	*Binance
	ApiKey         string
	ApiSecret      string
	SymbolPriceMap map[string]float64
	Data           *BinanceData
	Error      error
	Start          time.Time
	End            time.Time
}

type Binance struct {
	ApiID    int64
	Exchange string
}

type BinanceData struct {
	SpotAssetCurrency     []*models.AssetCurrencies
	FutureAssetCurrency   []*models.AssetCurrencies
	FuturePosition        []*models.Positions
	DeliveryAssetCurrency []*models.AssetCurrencies
	DeliveryPosition      []*models.Positions
}

func NewBinanceService(apiID int64, apiKey, apiSecret string, start, end time.Time) *BinanceService {
	return &BinanceService{
		Binance: &Binance{
			ApiID:    apiID,
			Exchange: string(exchange.Binance),
		},
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
	}
}

func (b *BinanceService) Do() {
	spotService := NewSpotService(b.ApiID, b.ApiKey, b.ApiSecret)
	symbolPriceMap, err := spotService.GetSetSymbolPrice()
	if err != nil {
		b.Error = err
		return
	}
	b.SymbolPriceMap = symbolPriceMap

	wg := sync.WaitGroup{}

	// 现货
	go func() {
		wg.Add(1)
		defer wg.Done()
		spotAssetCurrency, err := spotService.GetBalance()
		if err != nil {
			b.Error = err
			return
		}
		b.Data.SpotAssetCurrency = spotAssetCurrency
	}()

	// 金本位合约
	futureService := NewFutureService(b.ApiID, b.ApiKey, b.ApiSecret)
	go func() {
		wg.Add(1)
		defer wg.Done()

		assetCurrency, positions, err:= futureService.GetAccount()
		if err != nil {
			b.Error = err
			return
		}

		b.Data.FutureAssetCurrency = assetCurrency
		b.Data.FuturePosition = positions
	}()

	// 币本位合约
	deliveryService := NewDeliveryService(b.ApiID, b.ApiKey,b.ApiSecret)
	go func() {
		wg.Add(1)
		defer wg.Done()

		assetCurrency, positions, err:= deliveryService.GetAccount()
		if err != nil {
			b.Error = err
			return
		}

		b.Data.DeliveryAssetCurrency = assetCurrency
		b.Data.DeliveryPosition = positions
	}()

	wg.Wait()

	if b.Error != nil {
		// TODO: 写入进度, models.AssetCurrency models.Positions 添加字段做区分
	}



}

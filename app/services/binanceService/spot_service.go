package binanceService

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
	"transformer/app/repositories/symbolPriceRepository"

	"transformer/app/enums/exchange"
	"transformer/app/repositories/models"
	"transformer/library/cache"
	"transformer/library/cache/redis"
	"transformer/library/helper"
	"transformer/library/log"

	"github.com/adshao/go-binance/v2"
)

type SpotService struct {
	*Binance
	Client *binance.Client
	SymbolPriceMap map[string]float64
}

func NewSpotService(apiID int64, apiKey, apiSecret string) *SpotService {
	return &SpotService{
		Binance: &Binance{
			ApiID:    apiID,
			Exchange: string(exchange.Binance),
		},
		Client: binance.NewClient(apiKey, apiSecret),
	}
}

func (s *SpotService) GetSetSymbolPrice() (map[string]float64, error) {
	symbolPrices, err := s.Client.NewListPricesService().Do(context.Background())
	if err != nil {
		log.NewLogger().Error(err)
		return nil,err
	}

	sp := make(map[string]float64, 0)
	spm := make(map[string]interface{}, 0)
	for _, symbolPrice := range symbolPrices {
		price, err := strconv.ParseFloat(symbolPrice.Price, 64)
		if err != nil {
			log.NewLogger().Error(err)
			return nil,err
		}
		spm[symbolPrice.Symbol] = price
		sp[symbolPrice.Symbol] = price
	}

	err = symbolPriceRepository.NewSymbolPriceRepository().SetSymbolPrice(spm)
	if err != nil {
		return nil, err
	}

	return sp, nil
}

func (s *SpotService) GetSymbolInfo() error {
	exchangeInfo, err := s.Client.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		log.NewLogger().Error(err)
		return err
	}

	symbolInfo := make([]*models.SymbolInfos, 0)
	sim := make(map[string]interface{})
	for _, symbol := range exchangeInfo.Symbols {
		tmp := &models.SymbolInfos{
			Exchange:       s.Exchange,
			Symbol:         symbol.Symbol,
			BaseCurrency:   symbol.BaseAsset,
			BasePrecision:  int64(symbol.BaseAssetPrecision),
			QuoteCurrency:  symbol.QuoteAsset,
			QuotePrecision: int64(symbol.QuotePrecision),
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}
		symbolInfo = append(symbolInfo, tmp)

		jsonTmp, err := json.Marshal(tmp)
		if err != nil {
			log.NewLogger().Error(err)
			return err
		}
		sim[symbol.Symbol] = string(jsonTmp)
	}
	rdc, err := redis.NewRedis(redis.InstantDefault)
	if err != nil {
		log.NewLogger().Error(err)
		return err
	}

	cacheKey := cache.NewKey().GetSymbolInfo()
	rdc.HSet(context.Background(), cacheKey, sim)

	return nil
}

func (s *SpotService) GetBalance() ([]*models.AssetCurrencies, error) {
	account, err := s.Client.NewGetAccountService().Do(context.Background())
	if err != nil {
		log.NewLogger().Error(err)
		return nil, err
	}

	sac := make([]*models.AssetCurrencies, 0)
	for _, b := range account.Balances {
		amt := helper.AddString(b.Free, b.Locked)
		price, ok := s.SymbolPriceMap[b.Asset]
		if !ok {
			log.NewLogger().Error(errors.New(fmt.Sprintf("[api_id:%d] [exchange:%s] get price fail:%s", s.ApiID, s.Exchange, b.Asset)))
			return nil, err
		}

		tmp := &models.AssetCurrencies{
			ApiID:      s.ApiID,
			Exchange:   s.Exchange,
			Currency:   b.Asset,
			Amount:     amt,
			Price:      price,
			TimeWindow: time.Time{},
		}
		sac = append(sac, tmp)
	}

	return sac, nil
}

func (s *SpotService) GetOrder(start, end time.Time) error {
	rdc, err := redis.NewRedis(redis.InstantDefault)
	if err != nil {
		log.NewLogger().Error(err)
		return err
	}
	cacheKey := cache.NewKey().GetSymbolInfo()
	symbols := rdc.HGetAll(context.Background(), cacheKey)

	// TODO:

	fmt.Println("symbols:", symbols, start, end)

	//orders, err := s.Client.NewListOrdersService().
	//	StartTime(start.UnixMilli()).
	//	EndTime(end.UnixMilli()).
	//	Do(context.Background())
	//if err != nil {
	//	log.NewLogger().Error(err)
	//	return err
	//}
	//
	//for _, o := range orders {
	//	fmt.Println(o)
	//}

	return nil
}

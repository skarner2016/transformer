package symbolPriceRepository

import (
	"context"
	redisV8 "github.com/go-redis/redis/v8"
	"transformer/library/cache"
	"transformer/library/cache/redis"
	"strconv"
)

type SymbolPriceRepository struct {
	redisKey string
	rdc      *redisV8.Client
}

func NewSymbolPriceRepository() *SymbolPriceRepository {
	rdc, _ := redis.NewRedis(redis.InstantDefault)
	return &SymbolPriceRepository{
		redisKey: cache.NewKey().GetSymbolPrice(),
		rdc:      rdc,
	}
}

func (s *SymbolPriceRepository) SetSymbolPrice(spotPriceMap map[string]interface{}) error {
	_, err := s.rdc.HSet(context.Background(), s.redisKey, spotPriceMap).Result()

	return err
}

func (s *SymbolPriceRepository) GetSymbolPrice(symbol string) (float64, error) {
	priceString, err := s.rdc.HGet(context.Background(), s.redisKey, symbol).Result()
	if err != nil{
		return 0, err
	}

	price, err := strconv.ParseFloat(priceString, 64)
	if err != nil {
		return 0, err
	}

	return price, nil
}

func (s *SymbolPriceRepository) GetCurrencyPrice(currency string) (float64, error)  {
	if currency == "USDT" {
		return 1, nil
	}
	// 如:BTCUSDT
	symbol := currency + "USDT"

	return s.GetSymbolPrice(symbol)
}

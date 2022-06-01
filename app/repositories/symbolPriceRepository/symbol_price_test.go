package symbolPriceRepository

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"transformer/library/config"
	"testing"
)

func TestSymbolPriceRepository_GetSpotPrice(t *testing.T) {
	config.InitConfig()
	spr := NewSymbolPriceRepository()

	currency := "BTCUSDT"
	price, err := spr.GetSymbolPrice(currency)
	if err != nil {
		if err != redis.Nil {
			t.Error(err)
		}
		fmt.Println(fmt.Sprintf("%s don't have price", currency))
	}

	fmt.Println(currency, price)
}

func TestSymbolPriceRepository_SetSymbolPrice(t *testing.T) {
	config.InitConfig()
	spr := NewSymbolPriceRepository()

	spotPriceMap := map[string]interface{}{
		"AAA": 123.456,
		"BBB": 1.2345,
	}
	err := spr.SetSymbolPrice(spotPriceMap)
	if err != nil {
		fmt.Println(err)
	}
}

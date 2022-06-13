package binanceService

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2"
	"testing"
	"time"
	"transformer/library/config"
)

func GetApi() (apiID int64, apiKey, apiSecret string) {
	config.InitConfig()
	apiID = int64(123)
	//apiKey = "Gu4R5T2KtFxpdsDgQFoNzkea8h5xRdpfJZqX0dzJpt8RAEGkIzXyyk3GLZ90chgV"
	//apiSecret = "LUlk0bWv0riBFZq0aVePODjyR6OrylYHYGJRbnlPWiO1DsZYQJyyp9iCRoHbRwyc"

	apiKey = "5aiYx5WJh0LvkPbRhdKtJdtfvt0ngrKcVKgJyDyWoxEnqtZpONeqXFj2v0Uomn4g"
	apiSecret ="ayMPjg52K8VAwMZ5T2CrkXsMLtdNLofxQU3BlM2UEYE3k3re16LCo778D7vUD8iy"

	return apiID, apiKey, apiSecret
}

func TestSpotService_GetAssetPrice(t *testing.T) {
	spot := NewSpotService(GetApi())

	spm, err := spot.GetSetSymbolPrice()
	if err != nil {
		t.Error(err)
	}

	for s, f := range spm {
		fmt.Println(s, f)
	}
}

func TestSpotService_GetSymbolInfo(t *testing.T) {
	spot := NewSpotService(GetApi())

	err := spot.GetSymbolInfo()
	if err != nil {
		t.Error(err)
	}
}

func TestSpotService_GetBalance(t *testing.T) {
	spot := NewSpotService(GetApi())

	AssetCurrencies, err := spot.GetBalance()
	if err != nil {
		t.Error(err)
	}

	for _, currency := range AssetCurrencies {
		fmt.Println(currency)
	}
}

func TestSpotService_GetOrder(t *testing.T) {
	spot := NewSpotService(GetApi())

	subTime, err := time.ParseDuration("-24h")
	if err != nil {
		t.Error(err)
	}
	start := time.Now().Add(subTime)
	end := time.Now()

	err = spot.GetOrder(start, end)
	if err != nil {
		t.Error(err)
	}
}

func TestSpotService_Handle(t *testing.T) {
	_, apiKey, apiSecret := GetApi()

	//c := binance.NewClient(apiKey, apiSecret)

	c:=binance.NewFuturesClient(apiKey, apiSecret)

	startTime := time.Now().AddDate(0, 0, -29).UnixMilli()
	endTime := time.Now().UnixMilli()

	res, err := c.NewGetIncomeHistoryService().
		StartTime(startTime).
		EndTime(endTime).
		Do(context.Background())

	if err != nil {
		t.Error(err)
	}

	for _, i := range res {
		it := time.UnixMilli(i.Time)

		fmt.Println(it, i.Asset, i.Income, i.TradeID, i.TradeID, i.Symbol)
	}

}




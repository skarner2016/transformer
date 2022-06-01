package binanceService

import (
	"context"
	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
	"transformer/app/enums/exchange"
	"transformer/app/repositories/models"
	"transformer/app/repositories/symbolPriceRepository"
	"transformer/library/log"
	"strconv"
	"time"
)

type FutureService struct {
	*Binance
	Client *futures.Client
}

func NewFutureService(apiID int64, apiKey, apiSecret string) *FutureService {
	return &FutureService{
		Binance: &Binance{
			ApiID:    apiID,
			Exchange: string(exchange.Binance),
		},
		Client: binance.NewFuturesClient(apiKey, apiSecret),
	}
}

func (f *FutureService) GetAccount() ([]*models.AssetCurrencies, []*models.TmpPositions, error) {
	accountInfo, err := f.Client.NewGetAccountService().Do(context.Background())
	if err != nil {
		return nil, nil, err
	}

	now := time.Now()

	cs := make([]*models.AssetCurrencies, 0)
	for _, a := range accountInfo.Assets {
		amt, err := strconv.ParseFloat(a.WalletBalance, 64)
		if err != nil {
			return nil, nil, err
		}

		price, err := symbolPriceRepository.NewSymbolPriceRepository().GetCurrencyPrice(a.Asset)
		if err != nil {
			log.NewLogger().Error(err)
			return nil, nil, err
		}

		// 价值低于10U，略过
		worth := price * amt
		if worth < 10 {
			continue
		}

		tmp := &models.AssetCurrencies{
			ApiID:      f.ApiID,
			Exchange:   f.Exchange,
			Currency:   a.Asset,
			Type:       exchange.AssetTypeFuture,
			Amount:     amt,
			Price:      price,
			Worth:      worth,
			TimeWindow: now,
			CreatedAt:  now,
		}

		cs = append(cs, tmp)
	}

	ps := make([]*models.TmpPositions, 0)
	for _, p := range accountInfo.Positions {
		amt, err := strconv.ParseFloat(p.PositionAmt, 64)
		if err != nil {
			log.NewLogger().Error(err)
			return nil, nil, err
		}

		if amt == 0 {
			continue
		}

		leverage, err := strconv.ParseInt(p.Leverage, 10, 64)
		if err != nil {
			log.NewLogger().Error(err)
			return nil, nil, err
		}

		entryPrice, err := strconv.ParseFloat(p.EntryPrice, 64)
		if err != nil {
			log.NewLogger().Error(err)
			return nil, nil, err
		}

		side := exchange.PositionSideLong
		if p.PositionSide == "BOTH" {
			if amt < 0 {
				side = exchange.PositionSideShort
			}
		}

		if p.PositionSide == "SHORT" {
			side = exchange.PositionSideShort
		}

		tmp := &models.TmpPositions{
			ApiID:      f.ApiID,
			Exchange:   f.Exchange,
			Symbol:     p.Symbol,
			Leverage:   leverage,
			EntryPrice: entryPrice,
			Amount:     amt,
			Side:       int64(side),
			TimeWindow: now,
			CreatedAt:  now,
		}

		ps = append(ps, tmp)
	}

	return cs, ps, nil
}

func (f *FutureService) GetOrders()  {
	// TODO:
}

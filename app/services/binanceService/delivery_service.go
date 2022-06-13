package binanceService

import (
	"context"
	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/delivery"
	"strconv"
	"time"
	"transformer/app/enums/exchange"
	"transformer/app/repositories/models"
	"transformer/library/helper"
)

type DeliveryService struct {
	*Binance
	Client *delivery.Client
}

func NewDeliveryService(apiID int64, apiKey, apiSecret string) *DeliveryService {
	return &DeliveryService{
		Binance: &Binance{
			ApiID:    apiID,
			Exchange: string(exchange.Binance),
		},
		Client: binance.NewDeliveryClient(apiKey, apiSecret),
	}
}

func (d *DeliveryService) GetAccount() ([]*models.AssetCurrencies, []*models.Positions, error) {
	account, err := d.Client.NewGetAccountService().Do(context.Background())
	if err != nil {
		return nil, nil, err
	}

	now  := time.Now()
	acs := make([]*models.AssetCurrencies, 0)
	for _, asset := range account.Assets {
		amt := helper.AddString(asset.WalletBalance, asset.UnrealizedProfit)
		if amt == 0 {
			continue
		}

		tmp := &models.AssetCurrencies{
			ApiID:      d.ApiID,
			Exchange:   d.Exchange,
			Currency:   asset.Asset,
			Type:       0,
			Amount:     amt,
			Price:      0,
			Worth:      0,
			TimeWindow: now,
			CreatedAt:  now,
		}
		acs = append(acs, tmp)
	}

	ps := make([]*models.Positions, 0)
	for _, position := range account.Positions {
		amt, err := strconv.ParseFloat(position.PositionAmt, 64)
		if err != nil {
			return nil, nil, err
		}

		if amt == 0 {
			 continue
		}

		leverage, err := strconv.ParseInt(position.Leverage, 10, 64)
		if err != nil {
			return nil, nil, err
		}

		entryPrice, err := strconv.ParseFloat(position.EntryPrice, 64)
		if err != nil {
			return nil, nil, err
		}

		side := exchange.PositionSideLong
		if position.PositionSide == "SHORT" {
			side = exchange.PositionSideShort
		}
		if position.PositionSide == "BOTH" {
			if amt < 0 {
				 side = exchange.PositionSideShort
			}
		}

		tmp := &models.Positions{
			ApiID:      d.ApiID,
			Exchange:   d.Exchange,
			Symbol:     position.Symbol,
			Leverage:   leverage,
			EntryPrice: entryPrice,
			Amount:     amt,
			Side:       int64(side),
			TimeWindow: now,
			CreatedAt:  now,
		}
		ps = append(ps, tmp)
	}

	return acs, ps, err
}

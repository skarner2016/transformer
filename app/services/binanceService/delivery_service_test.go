package binanceService

import (
	"fmt"
	"testing"
)

func TestDeliveryService_GetAccount(t *testing.T) {
	d := NewDeliveryService(GetApi())

	acs, ps, err := d.GetAccount()
	if err != nil {
		t.Error(err)
	}

	for _, ac := range acs {
		fmt.Println(ac)
	}

	for _, p := range ps {
		fmt.Println(p.Symbol, p.Amount, p.Leverage, p.Side)
	}
}

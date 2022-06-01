package binanceService

import (
	"fmt"
	"testing"
)

func TestFutureService_GetAccount(t *testing.T) {
	f := NewFutureService(GetApi())

	//cs, ps, err := f.GetAccount()
	//cs, _, err := f.GetAccount()
	_, ps, err := f.GetAccount()
	if err != nil {
		t.Error(err)
	}

	//for _, a := range cs {
	//	fmt.Println(a.Currency, a.Amount, a.Price, a.Worth)
	//}

	for _, p := range ps {
		fmt.Println(p.Symbol,p.Amount, p.Leverage, p.EntryPrice)
	}
}

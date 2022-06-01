package models

import "time"

type SymbolInfos struct {
	Base
	Exchange       string
	Symbol         string
	BaseCurrency   string
	BasePrecision  int64
	QuoteCurrency  string
	QuotePrecision int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

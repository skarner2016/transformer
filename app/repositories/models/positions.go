package models

import (
	"time"
)

type Positions struct {
	Base
	ApiID      int64
	Exchange   string
	Symbol     string
	Leverage   int64
	EntryPrice float64
	Amount     float64
	Side       int64

	TimeWindow time.Time
	CreatedAt  time.Time
}

func (t *Positions) TableName() string {
	return "tmp_positions"
}

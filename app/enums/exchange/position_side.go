package exchange

type PositionSide int64

const (
	PositionSideLong  PositionSide = 1 // 做多
	PositionSideShort PositionSide = 2 // 做空
)

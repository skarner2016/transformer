package exchange

type ApiType int64

const (
	ApiTypeSpot     ApiType = 1 // 现货
	ApiTypeContract ApiType = 2 // 合约
)

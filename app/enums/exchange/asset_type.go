package exchange

type AssetType int64

const (
	AssetTypeSpot     = 1 // 现货
	AssetTypeFuture   = 2 // U本位合约
	AssetTypeDelivery = 3 // 币本位合约
)

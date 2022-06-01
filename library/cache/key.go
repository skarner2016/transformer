package cache

type Key struct {

}

func NewKey() *Key {
	return &Key{}
}

func (c *Key) GetSymbolPrice() string {
	return "symbol:price:binance"
}

func (c *Key) GetSymbolInfo() string {
	return "symbol:binance"
}

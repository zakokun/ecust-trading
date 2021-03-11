package strategy

// 策略模块，传入价格，返回买卖动作
type St interface {
	GetName() string
	SendPrice(f float64) *TradeMsg
}

type TradeMsg struct {
	// 交易动作,包括buy-market, sell-market, buy-limit, sell-limit
	Tp string
	// 交易价格
	Price float64
	// 交易数量
	Num float64
	// 交易对象
	Symbol string
}

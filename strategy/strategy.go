package strategy

// 策略模块，传入价格，返回买卖动作
type St interface {
	GetName() string
	SendPrice(f float64)
	Close()
}

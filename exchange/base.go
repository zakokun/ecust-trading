package exchange

var ex Ex

// Ex 所有交易所要实现的接口
type Ex interface {
	Start() error
	Close() error
	// 实时行情价格的方法 返回消息指针的channel
	TickListener(symbol string) chan *Trade
}

// 交易所返回的价格消息
type Trade struct {
	From   string  // 交易所名称
	Symbol string  // 交易对名称
	Price  float32 // 价格
	TS     int64   // 时间戳
}

func Register(sy string) {
	ex = new(Huobi)
	if err := ex.Start(); err != nil {
		panic(err)
	}
}
package exchange

var ex Ex

var (
	tickChan = make(chan *TickData, 1024)
)

// Ex 所有交易所要实现的接口
type Ex interface {
	Start() error
	Close() error
	// 实时行情价格的方法 返回消息指针的channel
	TickListener() chan *TickData
}

// 交易所返回的实时价格消息
type TickData struct {
	From   string  // 交易所名称
	Symbol string  // 交易对名称
	Price  float64 // 价格
	TS     int64   // 时间戳
}

func New(sy string) Ex {
	ex = new(Huobi)
	//if err := ex.Start(); err != nil {
	//	panic(err)
	//}
	return ex
}

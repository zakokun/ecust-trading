package strategy

import (
	"time"
)

// 网格交易
type Grid struct {
	Name        string
	buyStep     float64
	sellStep    float64
	lastTdPrice float64
}

// 根据5%初始化每次网格的容量
// 记录上次交易的价格
func (g *Grid) Init(f float64) {
	// init step
	g.buyStep = f * 0.05
	g.sellStep = f * 0.05
	g.lastTdPrice = f
	go func() {
		for {
			time.Sleep(time.Hour * 24)
			// 根据数据，重新计算买入和卖出的step
		}
	}()
}

// 获取实时价格，接收，做判断
func (g *Grid) GetName() string {
	return g.Name
}

// 获取实时价格，接收，做判断
func (g *Grid) SendPrice(f float64) (msg *TradeMsg) {
	//log.Info("get price is:%.2f\n", f)
	//if g.buyStep == 0 || g.sellStep == 0 {
	//	g.Init(f)
	//}
	//if f > g.lastTdPrice && f-g.lastTdPrice > g.sellStep { //上涨超过step
	//	msg = &TradeMsg{
	//		Tp:     "sell_market",
	//		Price:  f,
	//		Num:    (f - g.lastTdPrice) / g.sellStep,
	//		Symbol: "btsusdt",
	//	}
	//
	//	if err := DB.GetDB().SaveTradeData(msg); err != nil {
	//		log.Error("db.SaveTradeData(%v) err(%v)", msg, err)
	//		return
	//	}
	//	g.lastTdPrice = f
	//} else if f < g.lastTdPrice && g.lastTdPrice-f > g.buyStep { // 反之买入 判断buyStep
	//	msg = &TradeMsg{
	//		Tp:     "buy_market",
	//		Price:  f,
	//		Num:    (g.lastTdPrice - f) / g.buyStep,
	//		Symbol: "btsusdt",
	//	}
	//	if err := DB.GetDB().SaveTradeData(msg); err != nil {
	//		log.Error("db.SaveTradeData(%v) err(%v)", msg, err)
	//		return
	//	}
	//	g.lastTdPrice = f
	//}
	return
}

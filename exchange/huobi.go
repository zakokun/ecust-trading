package exchange

import (
	"math/rand"
	"time"

	"ecust-trading/conf"
	"ecust-trading/utils/log"

	"github.com/davecgh/go-spew/spew"
	"github.com/huobirdcenter/huobi_golang/pkg/client/marketwebsocketclient"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
)

type Huobi struct {
	Name   string
	AppKey string
	Secret string
	Wallet float64 // 余额
	Stock  float64 // 持仓
	Client *marketwebsocketclient.Last24hCandlestickWebSocketClient
}

// 连接 监听数据，把各种数据写到对应的chan里面
func (h *Huobi) Start() (err error) {
	cf := conf.Get().Ex.Huobi
	h.Client = new(marketwebsocketclient.Last24hCandlestickWebSocketClient).Init(cf.Host)
	h.Client.SetHandler(
		h.subscribe,
		h.handler,
	)
	return
}

func (h *Huobi) subscribe() {
	cf := conf.Get().Ex.Huobi
	symbol := conf.Get().Trade.Symbol
	//client.Request("btcusdt", "1608")
	h.Client.Subscribe(symbol, cf.ClientId)
}

func (h *Huobi) handler(resp interface{}) {
	candlestickResponse, ok := resp.(market.SubscribeLast24hCandlestickResponse)
	if ok {
		if &candlestickResponse != nil && candlestickResponse.Tick != nil {
			t := candlestickResponse.Tick
			log.Info("get data %s", spew.Sdump(t))
		}
	} else {
		log.Info("get unexpect data from client: %s", spew.Sdump(resp))
	}
}

func (h *Huobi) Close() (err error) {
	return
}

// TickListener 返回实时价格的channel
// 持续获取价格数据
func (h *Huobi) TickListener() chan *TickData {
	go func() {
		for {
			xx := rand.Float64()
			td := &TickData{
				From:   "huobi",
				Symbol: "BTC/USDT",
				Price:  xx,
				TS:     time.Now().Unix(),
			}
			tickChan <- td
			time.Sleep(time.Second)
		}
	}()
	return tickChan
}

func (h *Huobi) Kindle1DayListener() chan *CandleData {
	return Candle1DayChan
}

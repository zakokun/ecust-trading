package exchange

import (
	"ecust-trading/utils/DB"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/order"
	"time"

	"ecust-trading/conf"
	"ecust-trading/utils/log"

	"github.com/huobirdcenter/huobi_golang/pkg/client/marketwebsocketclient"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
)

type Huobi struct {
	Name   string
	AppKey string
	Secret string
	Wallet float64 // 余额
	Stock  float64 // 持仓
	client *marketwebsocketclient.CandlestickWebSocketClient
}

// SubMsg websocket 订阅资产的消息
type SubMsg struct {
	Sub string
	Id  string
}

// UnsubMsg websocket 取消订阅资产
type UnsubMsg struct {
	Unsub string
	Id    string
}

// 连接 监听数据，把各种数据写到对应的chan里面
func (h *Huobi) Start() (err error) {
	cf := conf.Get().Ex.Huobi
	h.client = new(marketwebsocketclient.CandlestickWebSocketClient).Init(cf.Host)
	h.client.SetHandler(
		h.subscribe,
		h.handler,
	)
	go h.startListener()
	return
}

func reqAndSubscribeCandlestick() {

	client := new(marketwebsocketclient.CandlestickWebSocketClient).Init(config.Host)

	client.SetHandler(
		func() {
			client.Request("btcusdt", "1min", 1569361140, 1569366420, "2305")

			client.Subscribe("btcusdt", "1min", "2118")
		},
		func(response interface{}) {
			resp, ok := response.(market.SubscribeCandlestickResponse)
			if ok {
				if &resp != nil {
					if resp.Tick != nil {
						t := resp.Tick
						applogger.Info("Candlestick update, id: %d, count: %d, vol: %v [%v-%v-%v-%v]",
							t.Id, t.Count, t.Vol, t.Open, t.Close, t.Low, t.High)
					}

					if resp.Data != nil {
						applogger.Info("WebSocket returned data, count=%d", len(resp.Data))
						for _, t := range resp.Data {
							applogger.Info("Candlestick data, id: %d, count: %d, vol: %v [%v-%v-%v-%v]",
								t.Id, t.Count, t.Vol, t.Open, t.Count, t.Low, t.High)
						}
					}
				}
			} else {
				applogger.Warn("Unknown response: %v", resp)
			}

		})

	client.Connect(true)

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	client.UnSubscribe("btcusdt", "1min", "2118")

	client.Close()
	applogger.Info("Client closed")
}

func (h *Huobi) getAllFinance() {
	//db := DB.GetDB().Table("")

}

func (h *Huobi) startListener() {
	for {
		cf := conf.Get().Ex.Huobi
		ct := new(client.MarketClient).Init(cf.APIHost)
		optionalRequest := market.GetCandlestickOptionalRequest{Period: market.DAY1, Size: 10}
		resp, err := ct.GetCandlestick("btcusdt", optionalRequest)
		if err != nil {
			log.Warn("ct.GetCandlestick(%s) err(%v)", "btcusdt", err)
		}
		log.Warn("v.Open.Float64(%v) err(%v)", resp, err)
		for _, v := range resp {
			op, ok := v.Open.Float64()
			if !ok {
				log.Warn("v.Open.Float64(%v) err(%v)", v.Open, op)
				continue
			}
			cl, ok := v.Close.Float64()
			if !ok {
				log.Warn("v.Close.Float64(%v) err(%v)", v.Open, err)
				continue
			}
			lo, ok := v.Low.Float64()
			if !ok {
				log.Warn("v.Low.Float64(%v) err(%v)", v.Open, err)
				continue
			}
			hi, ok := v.High.Float64()
			if !ok {
				log.Warn("v.High.Float64(%v) err(%v)", v.Open, err)
				continue
			}
			td := &CandleData{
				From:   "huobi",
				Symbol: conf.Get().Trade.Symbol,
				Open:   op,
				Close:  cl,
				Low:    lo,
				High:   hi,
				TS:     time.Now().Unix(),
			}
			Candle1DayChan <- td
		}
		time.Sleep(time.Second)
	}
}

func (h *Huobi) subscribe() {
	cf := conf.Get().Ex.Huobi
	//client.Request("btcusdt", "1608")

	flist := DB.GetDB().Table("stocks").Find()

	h.client.Subscribe(h.makeSymbol())
}
func (h *Huobi) makeSymbol(name, ts string) string {
	return fmt.Sprintf("market.%s.kline.%s", name, ts)
}

func (h *Huobi) handler(resp interface{}) {
	candlestickResponse, ok := resp.(market.SubscribeLast24hCandlestickResponse)
	if ok {
		if &candlestickResponse != nil && candlestickResponse.Tick != nil {
			t := candlestickResponse.Tick
			log.Info("get data %s", spew.Sdump(t))
			fp, ok := t.Close.Float64()
			if !ok {
				log.Warn("format decimal to float64(%v) err!", t.Close)
				return
			}
			td := &TickData{
				From:   "huobi",
				Symbol: conf.Get().Trade.Symbol,
				Price:  fp,
				TS:     time.Now().Unix(),
			}
			tickChan <- td
		}
	} else {
		log.Info("get unexpect data from client: %s", spew.Sdump(resp))
	}
}

func (h *Huobi) Close() (err error) {
	h.tickClient.Close()
	close(tickChan)
	close(Candle1DayChan)
	return
}

func (h *Huobi) Trade(td *TradeMsg) (err error) {
	cf := conf.Get().Ex.Huobi
	ct := new(client.OrderClient).Init(cf.AppKey, cf.Secret, cf.APIHost)
	od := &order.PlaceOrderRequest{
		AccountId: cf.ClientId,
		Symbol:    td.Symbol,
		Type:      td.Tp,
		Amount:    fmt.Sprintf("%.2f", td.Num),
		Price:     fmt.Sprintf("%.2f", td.Price),
	}
	_, err = ct.PlaceOrder(od)
	if err != nil {
		log.Info("PlaceOrder error!:%v", err)
		return
	}
	return
}

// TickListener 返回实时价格的channel
// 持续获取价格数据
func (h *Huobi) TickListener() chan *TickData {
	return tickChan
}

func (h *Huobi) Kindle1DayListener() chan *CandleData {
	return Candle1DayChan
}

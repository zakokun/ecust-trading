package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
	"github.com/jinzhu/gorm"
	"strings"
)

type Stocks struct {
	Id     int64
	TS     int64
	Symbol string
	Open   float64
	Close  float64
	Low    float64
	High   float64
	Volume int64
}

func main() {
	ls := []string{
		"usdt",
		"btc",
		"bch",
		"eth",
		"xrp",
		"ltc",
		"ht",
		"ada",
		"eos",
		"iota",
		"xem",
		"xmr",
		"dash",
		"neo",
		"trx",
		"icx",
		"lsk",
		"qtum",
		"etc",
	}
	path := strings.Join([]string{"root", ":", "root", "@tcp(", "127.0.0.1", ":", "3306", ")/", "trading", "?charset=utf8"}, "")
	conn, err := gorm.Open("mysql", path)
	if err != nil {
		panic(err)
	}
	//conn.Table("stocks").Create()
	ct := new(client.MarketClient).Init(config.Host)
	optionalRequest := market.GetCandlestickOptionalRequest{Period: market.DAY1, Size: 2}

	for _, v := range ls {
		resp, err := ct.GetCandlestick(v+"usdt", optionalRequest)
		if err != nil {
			applogger.Error("get err(%v)", err)
			continue
		}
		for _, vv := range resp {
			o, _ := vv.Open.Float64()
			c, _ := vv.Close.Float64()
			l, _ := vv.Low.Float64()
			h, _ := vv.High.Float64()
			vo := vv.Vol.IntPart()
			ts := vv.Id

			ss := &Stocks{
				Symbol: v + "usdt",
				Open:   o,
				Close:  c,
				Low:    l,
				High:   h,
				Volume: vo,
				TS:     ts,
			}
			if err = conn.Table("stocks").Create(ss).Error; err != nil {
				applogger.Error("save db get err(%v) val(%+v)", err, ss)
			}
		}

		applogger.Info("get price %v", resp)
	}

	//
	//
	//client := new(marketwebsocketclient.Last24hCandlestickWebSocketClient).Init(config.Host)
	//
	//// Set the callback handlers
	//client.SetHandler(
	//	// Connected handler
	//	func() {
	//		//client.Request("btcusdt", "1608")
	//		client.Subscribe("btcusdt", "1608")
	//	},
	//	// Response handler
	//	func(resp interface{}) {
	//		candlestickResponse, ok := resp.(market.SubscribeLast24hCandlestickResponse)
	//		if ok {
	//			if &candlestickResponse != nil {
	//				if candlestickResponse.Tick != nil {
	//					t := candlestickResponse.Tick
	//					applogger.Info("Candlestick update, id: %d, count: %v, volume: %v [%v-%v-%v-%v]",
	//						t.Id, t.Count, t.Vol, t.Open, t.Close, t.Low, t.High)
	//				}
	//
	//				if candlestickResponse.Data != nil {
	//					t := candlestickResponse.Data
	//					applogger.Info("Candlestick data, id: %d, count: %v, volume: %v [%v-%v-%v-%v]",
	//						t.Id, t.Count, t.Vol, t.Open, t.Close, t.Low, t.High)
	//				}
	//			}
	//		} else {
	//			applogger.Warn("Unknown response: %v", resp)
	//		}
	//	})
	//// Connect to the server and wait for the handler to handle the response
	//client.Connect(true)
	//time.Sleep(time.Minute)
}

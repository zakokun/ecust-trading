package exchange

import (
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/url"
	"time"
)

type Huobi struct {
	Name string
	Ws   *websocket.Conn
	Ch   chan *Trade
}

func (h *Huobi) Start() (err error) {
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8000", Path: "/echo"}
	log.Printf("connecting to %s", u.String())
	h.Ws, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer h.Ws.Close()
	return
}

func (h *Huobi) Close() (err error) {
	return
	return h.Ws.Close()
}

// TickListener 返回实时价格的channel
// 持续获取价格数据
func (h *Huobi) TickListener() chan *Trade {
	h.Ch = make(chan *Trade, 1024)
	go func() {
		for {
			xx := rand.Float32()
			td := &Trade{
				From:   "huobi",
				Symbol: "BTC/USDT",
				Price:  xx,
				TS:     time.Now().Unix(),
			}
			h.Ch <- td
			time.Sleep(time.Second)
		}
	}()
	return h.Ch
}

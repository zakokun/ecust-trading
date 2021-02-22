package finance

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

type Huobi struct {
	Name string
	Ws   *websocket.Conn
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
	return h.Ws.Close()
}

// TickListener 返回实时价格的channel
// 持续获取价格数据
func (h *Huobi) TickListener() (ch chan *Trade) {
	return
}

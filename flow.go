package main

import (
	"ecust-trading/exchange"
	"ecust-trading/strategy"
)

type Service struct {
	AA        int
	Ex        exchange.Ex // 对接的交易所接口
	CloseChan chan bool
}

func New(e exchange.Ex,st strategy.St) *Service {
	svr := new(Service)
	svr.Ex = e
	return svr
}

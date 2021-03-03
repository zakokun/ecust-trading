package main

import (
	"ecust-trading/exchange"
	"ecust-trading/strategy"
	"github.com/davecgh/go-spew/spew"
)

type Service struct {
	AA int
	Ex exchange.Ex // 对接的交易所接口
	St map[string]strategy.St
}

func New(e exchange.Ex, st ...strategy.St) *Service {
	svr := new(Service)
	svr.Ex = e
	svr.St = make(map[string]strategy.St)
	for _, v := range st {
		svr.St[v.GetName()] = v
	}
	return svr
}

func (s *Service) Close() {
	spew.Dump("close all service")
	s.Ex.Close()
}
func (s *Service) ListenTick() {
	ch := s.Ex.TickListener()
	go func() {
		for {
			select {
			case td, ok := <-ch:
				if !ok {
					s.Close()
					return
				}
				for _, v := range s.St {
					v.SendPrice(td.Price)
				}
			}
		}
	}()
}

func (s *Service) SaveDB() {
}

package main

import (
	"ecust-trading/exchange"
	"ecust-trading/strategy"
	"github.com/davecgh/go-spew/spew"
)

type Service struct {
	AA int
	Ex exchange.Ex // 对接的交易所接口
	St strategy.St
}

func New(e exchange.Ex, st strategy.St) *Service {
	svr := new(Service)
	svr.Ex = e
	svr.St = st
	return svr
}

func (s *Service) Close() {
	spew.Dump("close all service")
	s.Ex.Close()
	s.St.Close()
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
				s.St.GetPrice(td.Price)
			}
		}
	}()
}

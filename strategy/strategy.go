package strategy

import "ecust-trading/exchange"

type Service struct {
	AA        int
	Ex        exchange.Ex // 对接的交易所接口
	CloseChan chan bool
}

func New(e exchange.Ex) *Service {
	svr := new(Service)
	svr.Ex = e
	return svr
}

func (s *Service) ListenTick() {
	ch := s.Ex.TickListener()
	go func() {
		for {
			select {
			case td := <-ch:
				s.SaveTrade(td)
			case <-s.CloseChan:
				close(ch)
			}
		}
	}()
}

// 保存到db，交易处理
func (s *Service) SaveTrade(td *exchange.Trade) {

}

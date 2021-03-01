package main

import (
	"ecust-trading/conf"
	"ecust-trading/exchange"
	"ecust-trading/strategy"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ex := exchange.New(conf.Get().Symbol)
	st := strategy.New()
	svr := New(ex, st)


	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

package main

import (
	"ecust-trading/exchange"
	"ecust-trading/strategy"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ex := exchange.New()
	st := new(strategy.Grid)
	svr := New(ex, st)
	svr.ListenTick()

	http.ListenAndServe(":8000", nil)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			svr.Close()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

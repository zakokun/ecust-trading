package main

import (
	"fmt"
	"strconv"
	"time"
)

func okex(symbol string) chan int {
	ch := make(chan int, 10)
	go func() {
		i := 0
		for {
			ch <- i
			i += 2
			time.Sleep(time.Second)
		}
	}()
	return ch
}

func huobi(symbol string) chan int {
	ch := make(chan int, 10)
	go func() {
		i := 1
		for {
			ch <- i
			i += 2
			time.Sleep(time.Second)
		}
	}()
	return ch
}

func main() {
	mp := make(map[string]chan int)
	mp["okex"] = okex("b")
	mp["huobi"] = huobi("eos")
	for name, ch := range mp {
		name := name
		ch := ch
		go func() {
			for {
				select {
				case val := <-ch:
					fmt.Println("get " + name + " tick val:" + strconv.FormatInt(int64(val), 10))
				}
			}
		}()
	}
}

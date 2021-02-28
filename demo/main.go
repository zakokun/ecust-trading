package main

import (
	"fmt"
	"math"
)

var (
	stock          = int64(0)
	rmb            = 15000.0
	nowPrice       = 20.0
	lastTradePrice = nowPrice
	diff           = 1.0
	unit           = int64(0)
)

func initStock() {
	half := rmb / 2.0
	buy := int64(math.Abs(half / nowPrice))
	fmt.Println(buy)
	stock = buy
	rmb -= float64(buy) * nowPrice
	diff = 20 * 0.1
	unit = countUnit(nowPrice, diff, rmb)
}

func countUnit(p, d, r float64) int64 {
	avg := 0.0
	for i := 1; i <= 5; i++ {
		avg += p - d*float64(i)
	}
	return int64(r / avg)
}

func main() {
	rd := 0
	initStock()
	fmt.Printf("diff:%.2f unit:%d stock:%d rmb:%.2f\n", diff, unit, stock, rmb)
	for {
		if rd > 100 {
			break
		}
		rd++
		fmt.Printf("round %d\n", rd)

		old := unit
		unit = countUnit(nowPrice, diff, rmb)
		fmt.Printf("unit reset from %d to %d\n", old, unit)

		for i := 0; i <= 13; i++ {
			fmt.Printf("rmb:%.2f stock:%d nowPrice:%.2f total:%.2f unit:%d\n", rmb, stock, nowPrice, rmb+float64(stock)*nowPrice, unit)
			nowPrice += 1.0
			trade(nowPrice)
		}
		for i := 10; i >= 0; i-- {
			fmt.Printf("rmb:%.2f stock:%d nowPrice:%.2f total:%.2f unit:%d\n", rmb, stock, nowPrice, rmb+float64(stock)*nowPrice, unit)
			nowPrice -= 1.0
			trade(nowPrice)
		}
	}
	//fmt.Printf("rmb:%.2f stock:%d\n nowPrice:%.2f", rmb, stock, nowPrice)
}

func trade(newPrice float64) {
	rg := math.Abs(newPrice - lastTradePrice)
	if rg >= diff { // 触发操作
		num := int64(rg/diff) * unit   // 操作股票数量
		if newPrice > lastTradePrice { // 卖出
			if num > stock { //不够卖
				num = stock
			}
			stock -= num
			rmb += float64(num) * newPrice
		} else {                             // 买入
			if float64(num)*newPrice > rmb { //不够买
				num = int64(rmb / newPrice)
			}
			stock += num
			rmb -= float64(num) * newPrice
		}
		lastTradePrice = newPrice
	}
}

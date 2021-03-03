package strategy

import (
	"fmt"
)

// 网格交易
type Grid struct {
	Name string
	Step int64
}

// 获取实时价格，接收，做判断
func (g *Grid) GetName() string {
	return g.Name
}

// 获取实时价格，接收，做判断
func (g *Grid) SendPrice(f float32) {
	fmt.Printf("get price is:%.2f\n", f)
}

// 获取实时价格，接收，做判断
func (g *Grid) Close() {
	// 策略模块看看要不要关闭
}

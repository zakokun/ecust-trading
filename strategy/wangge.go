package strategy

// 网格交易
type Grid struct {
	Step int64
}

// 获取实时价格，接收，做判断
func (g *Grid) GetPrice(f float32) {
}

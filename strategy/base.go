package strategy

type St interface {
	GetPrice(f float32)
}

func New() St {
	return new(Grid)
}

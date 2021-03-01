package strategy

type St interface {
	GetPrice(f float32)
	Close()
}

func New() St {
	return new(Grid)
}

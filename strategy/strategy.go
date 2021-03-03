package strategy

type St interface {
	GetName() string
	SendPrice(f float32)
	Close()
}
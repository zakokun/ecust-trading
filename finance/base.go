package finance

type Finance interface {
	Start() error
	Close() error
	Tick() float32
}

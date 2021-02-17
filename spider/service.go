package spider

type Service struct {
	List financeGroup
}

type TickPrice struct {

}

type Spider interface {
	Start() error
	Close() error
	Tick() float32
}

type financeGroup map[string]*Spider

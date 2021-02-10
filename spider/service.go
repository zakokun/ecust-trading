package spider

type Service struct {
	List financeGroup
}

type Spider interface {
	Start()
	Close()
	Tick()
}

type financeGroup map[string]*Spider

package observer_szl

type Paper struct {
	Observers []Observer
}

func NewPaper() *Paper {
	return &Paper{Observers: make([]Observer, 0)}
}

func (p *Paper) Attach(observer Observer) {
	p.Observers = append(p.Observers, observer)
}

func (p *Paper) Notify() {
	for _, observer := range p.Observers {
		observer.Update(p)
	}
}

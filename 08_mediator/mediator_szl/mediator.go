package mediator_szl

import (
	"mediator_szl/domain"
	"sync"
)

type Mediator struct {
	CDDriver  domain.CDDriver
	CPUDriver domain.CPUDriver
	VideoCard domain.Player
	SoundCard domain.Player
}

var (
	instance *Mediator
	once     sync.Once
)

func GetMediatorInstance() *Mediator {
	once.Do(func() {
		instance = &Mediator{}
	})
	return instance
}

func (m *Mediator) Change(i interface{}) {
	switch i.(type) {
	case domain.CDDriver:
		data := m.CDDriver.GetData()
		m.CPUDriver.Process(data)
	case domain.CPUDriver:
		m.SoundCard.Play(i.(domain.CPUDriver).GetSound())
		m.VideoCard.Play(i.(domain.CPUDriver).GetVideo())
	}
}

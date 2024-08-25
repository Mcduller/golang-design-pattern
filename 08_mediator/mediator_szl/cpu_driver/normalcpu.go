package cpu_driver

import (
	"mediator_szl"
	"mediator_szl/domain"
	"mediator_szl/types"
)

type normalCpu struct {
	Video string
	Sound string
}

func (n *normalCpu) GetSound() string {
	return n.Sound
}

func (n *normalCpu) GetVideo() string {
	return n.Video
}

func (n *normalCpu) Process(data types.Data) {
	n.Sound = data.Sound
	n.Video = data.Video
	mediator_szl.GetMediatorInstance().Change(n)
}

func NewCpuDriver() domain.CPUDriver {
	return &normalCpu{}
}

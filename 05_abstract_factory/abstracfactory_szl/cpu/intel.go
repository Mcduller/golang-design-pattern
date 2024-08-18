package cpu

type IntelCPU struct {
	cpuHole int
}

func (i *IntelCPU) InstallCPU() string {
	return "Install Intel"
}

func NewIntelCPU(cpuHole int) CPUApi {
	return &IntelCPU{cpuHole: cpuHole}
}

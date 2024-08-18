package cpu

type AMDCpu struct {
	cpuHole int
}

func (A *AMDCpu) InstallCPU() string {
	return "Install AMD"
}

func NewAMDCPU(cpuHole int) *AMDCpu {
	return &AMDCpu{cpuHole: cpuHole}
}

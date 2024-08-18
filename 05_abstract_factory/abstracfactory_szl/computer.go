package abstracfactory_szl

import (
	"abstracfactory_szl/board"
	"abstracfactory_szl/cpu"
)

type InstallComputer interface {
	CreateCPU() cpu.CPUApi
	CreateBoard() board.BoardApi
}

type AMDMsiSchema struct {
}

func (s *AMDMsiSchema) CreateCPU() cpu.CPUApi {
	return cpu.NewAMDCPU(990)
}

func (s *AMDMsiSchema) CreateBoard() board.BoardApi {
	return board.NewMsiBoard(990)
}

type IntelGAMSchema struct {
}

func (i *IntelGAMSchema) CreateCPU() cpu.CPUApi {
	return cpu.NewIntelCPU(130)
}

func (i *IntelGAMSchema) CreateBoard() board.BoardApi {
	return board.NewGamBoard(130)
}

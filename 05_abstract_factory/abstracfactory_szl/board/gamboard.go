package board

type GamBoard struct {
	cpuHole int
}

func (g GamBoard) InstallBoard() string {
	return "Install GAM"
}

func NewGamBoard(cpuHole int) *GamBoard {
	return &GamBoard{cpuHole: cpuHole}
}

func GAMBoardFactory(cpuHole int) BoardApi {
	return NewGamBoard(cpuHole)
}

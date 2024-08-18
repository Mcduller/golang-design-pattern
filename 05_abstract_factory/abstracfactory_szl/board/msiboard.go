package board

type MsiBoard struct {
	cpuHole int
}

func (m MsiBoard) InstallBoard() string {
	return "Install Msi"
}

func NewMsiBoard(cpuHole int) *MsiBoard {
	return &MsiBoard{
		cpuHole: cpuHole,
	}
}

func MisBoardFactory() {

}

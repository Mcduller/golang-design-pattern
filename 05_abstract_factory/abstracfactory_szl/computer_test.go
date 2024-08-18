package abstracfactory_szl

import "testing"

func TestInstall(t *testing.T) {
	t.Run("AMDMsi", func(t *testing.T) {
		var schema InstallComputer
		schema = &AMDMsiSchema{}

		wantCpu := "Install AMD"
		if gotCPU := schema.CreateCPU().InstallCPU(); gotCPU != wantCpu {
			t.Errorf("wantCpu；%s,gotCPU:%s", wantCpu, gotCPU)
		}

		wantBoard := "Install Msi"
		if gotBoard := schema.CreateBoard().InstallBoard(); gotBoard != wantBoard {
			t.Errorf("wantBoard:%s.gotBoard；%s", wantBoard, gotBoard)
		}
	})

	t.Run("IntelGAM", func(t *testing.T) {
		var schma InstallComputer
		schma = &IntelGAMSchema{}

		wantCPU := "Install Intel"
		if gotCpu := schma.CreateCPU().InstallCPU(); gotCpu != wantCPU {
			t.Errorf("wantCpu:%s,gotCpu:%s", wantCPU, gotCpu)
		}

		wantBoard := "Install GAM"
		if gotBoard := schma.CreateBoard().InstallBoard(); gotBoard != wantBoard {
			t.Errorf("wantBoard:%s,gotBoard；%s", wantBoard, gotBoard)
		}
	})

}

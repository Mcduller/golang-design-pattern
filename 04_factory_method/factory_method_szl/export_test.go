package factory_method_szl

import "testing"

func TestBaseExporter_SetName(t *testing.T) {
	t.Run("DB export", func(t *testing.T) {
		dbExporterFactory := NewDBExporterFactory()
		dbExporter := dbExporterFactory.CreateExporter()
		want := "DB Export"
		if got := dbExporter.Export(); got != want {
			t.Errorf("want:%s,got:%s", want, got)
		}
	})

	t.Run("Txt Export", func(t *testing.T) {
		txtExporterFactory := NewTxtExporterFactory()
		txtExporter := txtExporterFactory.CreateExporter()
		want := "TxT Export"
		if got := txtExporter.Export(); got != want {
			t.Errorf("want:%s,got：%s", want, got)
		}
	})
}

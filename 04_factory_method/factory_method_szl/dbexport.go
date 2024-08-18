package factory_method_szl

import "fmt"

type DBExporterFactory struct {
}

type dBExporter struct {
	*BaseExporter
}

func (d *dBExporter) Export() string {
	return fmt.Sprintf("DB Export")
}

func (d *DBExporterFactory) CreateExporter() ExportApi {
	return &dBExporter{}
}

func NewDBExporterFactory() ExporterFactory {
	return &DBExporterFactory{}
}

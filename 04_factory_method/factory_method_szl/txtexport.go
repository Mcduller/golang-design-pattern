package factory_method_szl

import "fmt"

type TxtExporterFactory struct {
}

type txtExporter struct {
	*BaseExporter
}

func (t *txtExporter) Export() string {
	return fmt.Sprintf("TxT Export")
}

func (t *TxtExporterFactory) CreateExporter() ExportApi {
	return &txtExporter{}
}

func NewTxtExporterFactory() ExporterFactory {
	return &TxtExporterFactory{}
}

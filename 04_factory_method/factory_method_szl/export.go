package factory_method_szl

type ExportApi interface {
	SetName(name string)
	Export() string
}

type ExporterFactory interface {
	CreateExporter() ExportApi
}

type BaseExport interface {
}

type BaseExporter struct {
	name string
}

func (b BaseExporter) SetName(name string) {
	b.name = name
}

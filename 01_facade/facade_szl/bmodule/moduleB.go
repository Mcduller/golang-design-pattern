package bmodule

type ModuleBApi interface {
	BRun() string
}

type moduleBImpl struct {
}

func NewModuleBApi() ModuleBApi {
	return moduleBImpl{}
}

func (m moduleBImpl) BRun() string {
	return "BRun"
}

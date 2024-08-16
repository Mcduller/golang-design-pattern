package amodule

type ModuleAApi interface {
	ARun() string
}

type moduleAImpl struct {
}

func NewModuleAApi() ModuleAApi {
	return &moduleAImpl{}
}

func (m moduleAImpl) ARun() string {
	return "ARun"
}

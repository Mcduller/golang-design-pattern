package facade_szl

import (
	"facade_szl/amodule"
	"facade_szl/bmodule"
	"fmt"
)

type FacadeApi interface {
	Run() string
}

type facadeImpl struct {
	a amodule.ModuleAApi
	b bmodule.ModuleBApi
}

func NewFacade() FacadeApi {
	return facadeImpl{
		a: amodule.NewModuleAApi(),
		b: bmodule.NewModuleBApi(),
	}
}

func (f facadeImpl) Run() string {
	aRun := f.a.ARun()
	bRun := f.b.BRun()

	return fmt.Sprintf("%s,%s", aRun, bRun)
}

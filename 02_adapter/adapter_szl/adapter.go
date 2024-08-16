package adapter_szl

import "adapter_szl/logfile"

type Adaptee interface {
	GetLog() string
}

type Adapter struct {
	adaptee Adaptee
}

func NewAdapter() Adapter {
	return Adapter{
		adaptee: logfile.NewFileLog(),
	}
}

func (a Adapter) GetLogFromDB() string {
	return a.adaptee.GetLog()
}

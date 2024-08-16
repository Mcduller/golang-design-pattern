package adapter_szl

type DBLog interface {
	GetLogFromDB() string
}

type DBLogImpl struct {
	adapter Adapter
}

func NewDBLog() DBLog {
	return DBLogImpl{adapter: NewAdapter()}
}

func (D DBLogImpl) GetLogFromDB() string {
	return D.adapter.GetLogFromDB()
}

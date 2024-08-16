package logfile

type FileLog interface {
	GetLog() string
}

type fileLogImpl struct {
}

func NewFileLog() FileLog {
	return &fileLogImpl{}
}

func (f fileLogImpl) GetLog() string {
	return "From File"
}

package log

type loggerMock struct {
}

// NewMock creates new instance of dhtlog.Logger that does nothing.
// It can be used for tests.
func NewMock() Logger {
	mock := &loggerMock{}

	return NewLogger(mock)
}

func (logger *loggerMock) Write(level string, message string) {
}

func (logger *loggerMock) Writef(level string, message string, args ...interface{}) {
}

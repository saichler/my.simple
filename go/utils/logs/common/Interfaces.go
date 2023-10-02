package common

type Logger interface {
	Trace(interface{}, ...interface{})
	Debug(interface{}, ...interface{})
	Info(interface{}, ...interface{})
	Warning(interface{}, ...interface{})
	Error(interface{}, ...interface{}) error
}

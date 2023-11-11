package logs

type Logger interface {
	Trace(interface{}, ...interface{})
	Debug(interface{}, ...interface{})
	Info(interface{}, ...interface{})
	Warning(interface{}, ...interface{})
	Error(interface{}, ...interface{}) error
	Empty() bool
}

type LoggerImpl interface {
	Print(string)
}

var Log Logger = NewLoggerQueue(NewFmtLogger())

func Trace(any interface{}, anys ...interface{}) {
	Log.Trace(any, anys...)
}

func Debug(any interface{}, anys ...interface{}) {
	Log.Debug(any, anys...)
}

func Info(any interface{}, anys ...interface{}) {
	Log.Info(any, anys...)
}

func Warning(any interface{}, anys ...interface{}) {
	Log.Warning(any, anys...)
}

func Error(any interface{}, anys ...interface{}) error {
	return Log.Error(any, anys...)
}

func Empty() bool {
	return Log.Empty()
}

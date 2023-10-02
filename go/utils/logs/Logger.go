package logs

import "github.com/saichler/my.simple/go/utils/logs/common"

var Log common.Logger = NewFmtLogger()

func Trace(any interface{}, anys ...interface{}) {
	Log.Trace(any, anys)
}

func Debug(any interface{}, anys ...interface{}) {
	Log.Debug(any, anys)
}

func Info(any interface{}, anys ...interface{}) {
	Log.Info(any, anys)
}

func Warning(any interface{}, anys ...interface{}) {
	Log.Warning(any, anys)
}

func Error(any interface{}, anys ...interface{}) error {
	return Log.Error(any, anys)
}

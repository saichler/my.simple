package logs

import (
	"errors"
	"fmt"
)

type FmtLogger struct {
}

func NewFmtLogger() *FmtLogger {
	fl := &FmtLogger{}
	return fl
}

func (logger FmtLogger) Trace(any interface{}, anys ...interface{}) {
	fmt.Println(TraceToString(any, anys))
}

func (logger FmtLogger) Debug(any interface{}, anys ...interface{}) {
	fmt.Println(DebugToString(any, anys))
}

func (logger FmtLogger) Info(any interface{}, anys ...interface{}) {
	fmt.Println(InfoToString(any, anys))
}

func (logger FmtLogger) Warning(any interface{}, anys ...interface{}) {
	fmt.Println(WarningToString(any, anys))
}

func (logger FmtLogger) Error(any interface{}, anys ...interface{}) error {
	return errors.New(ErrorToString(any, anys))
}

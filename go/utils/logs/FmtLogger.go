package logs

import "fmt"

type FmtLogger struct {
}

func NewFmtLogger() *FmtLogger {
	fl := &FmtLogger{}
	return fl
}

func (logger FmtLogger) Print(str string) {
	fmt.Println(str)
}

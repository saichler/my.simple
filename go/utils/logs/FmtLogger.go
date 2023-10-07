package logs

import (
	"errors"
	"fmt"
	"github.com/saichler/my.simple/go/utils/strng"
)

type FmtLogger struct {
}

func NewFmtLogger() *FmtLogger {
	fl := &FmtLogger{}
	return fl
}

func (logger FmtLogger) Trace(any interface{}, anys ...interface{}) {
	str := strng.New("<Trace   > ")
	str.TypesPrefix = false
	str.AddSpaceWhenAdding = true
	s, _ := str.StringOf(any)
	str.Add(s)
	if anys != nil {
		for _, a := range anys {
			s, _ = str.StringOf(a)
			str.Add(s)
		}
	}
	fmt.Println(str.String())
}

func (logger FmtLogger) Debug(any interface{}, anys ...interface{}) {
	str := strng.New("< Debug   > ")
	str.TypesPrefix = false
	str.AddSpaceWhenAdding = true
	s, _ := str.StringOf(any)
	str.Add(s)
	if anys != nil {
		for _, a := range anys {
			s, _ = str.StringOf(a)
			str.Add(s)
		}
	}
	fmt.Println(str.String())
}

func (logger FmtLogger) Info(any interface{}, anys ...interface{}) {
	str := strng.New("<  Info    > ")
	str.TypesPrefix = false
	str.AddSpaceWhenAdding = true
	s, _ := str.StringOf(any)
	str.Add(s)
	if anys != nil {
		for _, a := range anys {
			s, _ = str.StringOf(a)
			str.Add(s)
		}
	}
	fmt.Println(str.String())
}

func (logger FmtLogger) Warning(any interface{}, anys ...interface{}) {
	str := strng.New("<   Warning > ")
	str.TypesPrefix = false
	str.AddSpaceWhenAdding = true
	s, _ := str.StringOf(any)
	str.Add(s)
	if anys != nil {
		for _, a := range anys {
			s, _ = str.StringOf(a)
			str.Add(s)
		}
	}
	fmt.Println(str.String())
}

func (logger FmtLogger) Error(any interface{}, anys ...interface{}) error {
	str := strng.New("<      Error > ")
	str.TypesPrefix = false
	str.AddSpaceWhenAdding = true
	s, _ := str.StringOf(any)
	str.Add(s)
	if anys != nil {
		for _, a := range anys {
			s, _ = str.StringOf(a)
			str.Add(s)
		}
	}
	fmt.Println(str.String())
	return errors.New(str.String())
}

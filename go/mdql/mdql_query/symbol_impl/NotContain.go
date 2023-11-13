package symbol_impl

import (
	"reflect"
)

type NotContain struct {
	contain *Contain
}

func NewNotContain() *NotContain {
	notContain := &NotContain{}
	notContain.contain = NewContain()
	return notContain
}

func (notContain *NotContain) Exec(aSide, zSide []reflect.Value) bool {
	return !Exec(aSide, zSide, notContain.contain.kind2Contain, "Not Contain")
}

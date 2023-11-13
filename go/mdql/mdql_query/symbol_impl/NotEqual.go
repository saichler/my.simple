package symbol_impl

import (
	"reflect"
)

type NotEqual struct {
	eq *Equal
}

func NewNotEqual() *NotEqual {
	notEq := &NotEqual{}
	notEq.eq = NewEqual()
	return notEq
}

func (notequal *NotEqual) Exec(aSide, zSide []reflect.Value) bool {
	return !Exec(aSide, zSide, notequal.eq.kind2EqualFunc, "Not Equal")
}

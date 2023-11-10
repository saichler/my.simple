package symbol_impl

import (
	"reflect"
	"strings"
)

type LessThanOrEqual struct {
	kind2LTEQ map[reflect.Kind]func(reflect.Value, reflect.Value) bool
}

func NewLessThanOrEqual() *LessThanOrEqual {
	c := &LessThanOrEqual{}
	c.kind2LTEQ = make(map[reflect.Kind]func(reflect.Value, reflect.Value) bool)
	c.kind2LTEQ[reflect.String] = lteqString
	c.kind2LTEQ[reflect.Int] = lteqInt
	c.kind2LTEQ[reflect.Int8] = lteqInt
	c.kind2LTEQ[reflect.Int16] = lteqInt
	c.kind2LTEQ[reflect.Int32] = lteqInt
	c.kind2LTEQ[reflect.Int64] = lteqInt
	c.kind2LTEQ[reflect.Uint] = lteqUint
	c.kind2LTEQ[reflect.Uint8] = lteqUint
	c.kind2LTEQ[reflect.Uint16] = lteqUint
	c.kind2LTEQ[reflect.Uint32] = lteqUint
	c.kind2LTEQ[reflect.Uint64] = lteqUint
	return c
}

func (lteq *LessThanOrEqual) Exec(aSide, zSide []reflect.Value) bool {
	return Exec(aSide, zSide, lteq.kind2LTEQ, "Less Than Or Equal")
}

func lteqString(aSide, zSide reflect.Value) bool {
	aside := removeSingleQuote(strings.ToLower(aSide.String()))
	zside := removeSingleQuote(strings.ToLower(zSide.String()))
	return aside <= zside
}

func lteqInt(aSide, zSide reflect.Value) bool {
	aside, ok := getInt64(aSide)
	if !ok {
		return false
	}
	zside, ok := getInt64(zSide)
	if !ok {
		return false
	}
	return aside <= zside
}

func lteqUint(aSide, zSide reflect.Value) bool {
	aside, ok := getUint64(aSide)
	if !ok {
		return false
	}
	zside, ok := getUint64(zSide)
	if !ok {
		return false
	}
	return aside <= zside
}

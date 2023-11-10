package symbol_impl

import (
	"reflect"
	"strings"
)

type LessThan struct {
	kind2LT map[reflect.Kind]func(reflect.Value, reflect.Value) bool
}

func NewLessThan() *LessThan {
	c := &LessThan{}
	c.kind2LT = make(map[reflect.Kind]func(reflect.Value, reflect.Value) bool)
	c.kind2LT[reflect.String] = ltString
	c.kind2LT[reflect.Int] = ltInt
	c.kind2LT[reflect.Int8] = ltInt
	c.kind2LT[reflect.Int16] = ltInt
	c.kind2LT[reflect.Int32] = ltInt
	c.kind2LT[reflect.Int64] = ltInt
	c.kind2LT[reflect.Uint] = ltUint
	c.kind2LT[reflect.Uint8] = ltUint
	c.kind2LT[reflect.Uint16] = ltUint
	c.kind2LT[reflect.Uint32] = ltUint
	c.kind2LT[reflect.Uint64] = ltUint
	return c
}

func (lt *LessThan) Exec(aSide, zSide []reflect.Value) bool {
	return Exec(aSide, zSide, lt.kind2LT, "Less Than")
}

func ltString(aSide, zSide reflect.Value) bool {
	aside := removeSingleQuote(strings.ToLower(aSide.String()))
	zside := removeSingleQuote(strings.ToLower(zSide.String()))
	return aside < zside
}

func ltInt(aSide, zSide reflect.Value) bool {
	aside, ok := getInt64(aSide)
	if !ok {
		return false
	}
	zside, ok := getInt64(zSide)
	if !ok {
		return false
	}
	return aside < zside
}

func ltUint(aSide, zSide reflect.Value) bool {
	aside, ok := getUint64(aSide)
	if !ok {
		return false
	}
	zside, ok := getUint64(zSide)
	if !ok {
		return false
	}
	return aside < zside
}

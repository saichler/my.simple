package symbol_impl

import (
	"reflect"
	"strings"
)

type GreaterThanOrEqual struct {
	kind2GTEQ map[reflect.Kind]func(reflect.Value, reflect.Value) bool
}

func NewGreaterThanOrEqual() *GreaterThanOrEqual {
	c := &GreaterThanOrEqual{}
	c.kind2GTEQ = make(map[reflect.Kind]func(reflect.Value, reflect.Value) bool)
	c.kind2GTEQ[reflect.String] = gteqString
	c.kind2GTEQ[reflect.Int] = gteqInt
	c.kind2GTEQ[reflect.Int8] = gteqInt
	c.kind2GTEQ[reflect.Int16] = gteqInt
	c.kind2GTEQ[reflect.Int32] = gteqInt
	c.kind2GTEQ[reflect.Int64] = gteqInt
	c.kind2GTEQ[reflect.Uint] = gteqUint
	c.kind2GTEQ[reflect.Uint8] = gteqUint
	c.kind2GTEQ[reflect.Uint16] = gteqUint
	c.kind2GTEQ[reflect.Uint32] = gteqUint
	c.kind2GTEQ[reflect.Uint64] = gteqUint
	return c
}

func (gteq *GreaterThanOrEqual) Exec(aSide, zSide []reflect.Value) bool {
	return Exec(aSide, zSide, gteq.kind2GTEQ, "Greater Than Or Equal")
}

func gteqString(aSide, zSide reflect.Value) bool {
	aside := removeSingleQuote(strings.ToLower(aSide.String()))
	zside := removeSingleQuote(strings.ToLower(zSide.String()))
	return aside >= zside
}

func gteqInt(aSide, zSide reflect.Value) bool {
	aside, ok := getInt64(aSide)
	if !ok {
		return false
	}
	zside, ok := getInt64(zSide)
	if !ok {
		return false
	}
	return aside >= zside
}

func gteqUint(aSide, zSide reflect.Value) bool {
	aside, ok := getUint64(aSide)
	if !ok {
		return false
	}
	zside, ok := getUint64(zSide)
	if !ok {
		return false
	}
	return aside >= zside
}

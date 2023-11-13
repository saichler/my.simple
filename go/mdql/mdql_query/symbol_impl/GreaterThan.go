package symbol_impl

import (
	"reflect"
	"strings"
)

type GreaterThan struct {
	kind2GT map[reflect.Kind]func(reflect.Value, reflect.Value) bool
}

func NewGreaterThan() *GreaterThan {
	c := &GreaterThan{}
	c.kind2GT = make(map[reflect.Kind]func(reflect.Value, reflect.Value) bool)
	c.kind2GT[reflect.String] = gtString
	c.kind2GT[reflect.Int] = gtInt
	c.kind2GT[reflect.Int8] = gtInt
	c.kind2GT[reflect.Int16] = gtInt
	c.kind2GT[reflect.Int32] = gtInt
	c.kind2GT[reflect.Int64] = gtInt
	c.kind2GT[reflect.Uint] = gtUint
	c.kind2GT[reflect.Uint8] = gtUint
	c.kind2GT[reflect.Uint16] = gtUint
	c.kind2GT[reflect.Uint32] = gtUint
	c.kind2GT[reflect.Uint64] = gtUint
	return c
}

func (gt *GreaterThan) Exec(aSide, zSide []reflect.Value) bool {
	return Exec(aSide, zSide, gt.kind2GT, "Greater Than")
}

func gtString(aSide, zSide reflect.Value) bool {
	aside := removeSingleQuote(strings.ToLower(aSide.String()))
	zside := removeSingleQuote(strings.ToLower(zSide.String()))
	return aside > zside
}

func gtInt(aSide, zSide reflect.Value) bool {
	aside, ok := getInt64(aSide)
	if !ok {
		return false
	}
	zside, ok := getInt64(zSide)
	if !ok {
		return false
	}
	return aside > zside
}

func gtUint(aSide, zSide reflect.Value) bool {
	aside, ok := getUint64(aSide)
	if !ok {
		return false
	}
	zside, ok := getUint64(zSide)
	if !ok {
		return false
	}
	return aside > zside
}

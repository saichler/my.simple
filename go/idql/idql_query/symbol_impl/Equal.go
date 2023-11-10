package symbol_impl

import (
	"reflect"
	"strings"
)

type Equal struct {
	kind2EqualFunc map[reflect.Kind]func(reflect.Value, reflect.Value) bool
}

func NewEqual() *Equal {
	eq := &Equal{}
	eq.kind2EqualFunc = make(map[reflect.Kind]func(reflect.Value, reflect.Value) bool)
	eq.kind2EqualFunc[reflect.String] = eqString
	eq.kind2EqualFunc[reflect.Int] = eqInt
	eq.kind2EqualFunc[reflect.Int8] = eqInt
	eq.kind2EqualFunc[reflect.Int16] = eqInt
	eq.kind2EqualFunc[reflect.Int32] = eqInt
	eq.kind2EqualFunc[reflect.Int64] = eqInt
	eq.kind2EqualFunc[reflect.Uint] = eqUint
	eq.kind2EqualFunc[reflect.Uint8] = eqUint
	eq.kind2EqualFunc[reflect.Uint16] = eqUint
	eq.kind2EqualFunc[reflect.Uint32] = eqUint
	eq.kind2EqualFunc[reflect.Uint64] = eqUint
	eq.kind2EqualFunc[reflect.Ptr] = eqPtr
	eq.kind2EqualFunc[reflect.Bool] = eqBool
	return eq
}

func (equal *Equal) Exec(aSide, zSide []reflect.Value) bool {
	return Exec(aSide, zSide, equal.kind2EqualFunc, "Equal")
}

func eqString(aSide, zSide reflect.Value) bool {
	aside := removeSingleQuote(strings.ToLower(aSide.String()))
	zside := removeSingleQuote(strings.ToLower(zSide.String()))
	if aside == "nil" && zside == "" {
		return true
	}
	if zside == "nil" && aside == "" {
		return true
	}
	splits := GetWildCardSubstrings(zside)
	if splits == nil {
		return aside == zside
	}
	for _, substr := range splits {
		if substr != "" && strings.Contains(aside, substr) {
			return true
		}
	}
	return false
}

func eqPtr(aSide, zSide reflect.Value) bool {
	if aSide.Kind() == reflect.Ptr && zSide.IsValid() && zSide.String() == "nil" && aSide.IsNil() {
		return true
	}
	if zSide.Kind() == reflect.Ptr && aSide.IsValid() && aSide.String() == "nil" && zSide.IsNil() {
		return true
	}
	return false
}

func eqInt(aSide, zSide reflect.Value) bool {
	aside, aok := getInt64(aSide)
	zside, zok := getInt64(zSide)
	if zSide.String() == "nil" && aok && aside == 0 {
		return true
	}
	if aSide.String() == "nil" && zok && zside == 0 {
		return true
	}
	if !aok || !zok {
		return false
	}
	return aside == zside
}

func eqBool(aSide, zSide reflect.Value) bool {
	aside, aok := getBool(aSide)
	zside, zok := getBool(zSide)
	if !aok || !zok {
		return false
	}
	return aside == zside
}

func eqUint(aSide, zSide reflect.Value) bool {
	aside, ok := getUint64(aSide)
	if !ok {
		return false
	}
	zside, ok := getUint64(zSide)
	if !ok {
		return false
	}
	return aside == zside
}

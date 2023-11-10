package symbol_impl

import (
	"reflect"
	"strconv"
	"strings"
)

func Exec(aSide, zSide []reflect.Value, symbolImpl map[reflect.Kind]func(reflect.Value, reflect.Value) bool, name string) bool {
	kind := getKind(aSide, zSide)
	fnc := symbolImpl[kind]
	if fnc == nil {
		panic("Cannot find symbol impl func for:" + name + " Kind:" + kind.String())
	}
	for _, aside := range aSide {
		for _, zside := range zSide {
			if fnc(aside, zside) {
				return true
			}
		}
	}
	return false
}

func getKind(aside, zside []reflect.Value) reflect.Kind {
	aSideKind := reflect.String
	zSideKind := reflect.String
	if len(aside) > 0 {
		aSideKind = aside[0].Kind()
	}
	if len(zside) > 0 {
		zSideKind = zside[0].Kind()
	}
	if aSideKind != reflect.String {
		return aSideKind
	} else if zSideKind != reflect.String {
		return zSideKind
	}
	return aSideKind
}

func removeSingleQuote(value string) string {
	if strings.Contains(value, "'") {
		return value[1 : len(value)-1]
	}
	return value
}

func getInt64(value reflect.Value) (int64, bool) {
	if value.Kind() != reflect.String {
		return value.Int(), true
	} else {
		i, e := strconv.Atoi(value.String())
		if e != nil {
			return 0, false
		}
		return int64(i), true
	}
}

func getUint64(value reflect.Value) (uint64, bool) {
	if value.Kind() != reflect.String {
		return value.Uint(), true
	} else {
		i, e := strconv.Atoi(value.String())
		if e != nil {
			return 0, false
		}
		return uint64(i), true
	}
}

func GetWildCardSubstrings(str string) []string {
	if !strings.Contains(str, "*") {
		return nil
	}
	return strings.Split(str, "*")
}

func getBool(value reflect.Value) (bool, bool) {
	if value.Kind() != reflect.String {
		return value.Bool(), true
	} else {
		b, e := strconv.ParseBool(value.String())
		if e != nil {
			return false, false
		}
		return b, true
	}
}

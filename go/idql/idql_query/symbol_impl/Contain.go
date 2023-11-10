package symbol_impl

import (
	"reflect"
	"strconv"
	"strings"
)

type Contain struct {
	kind2Contain map[reflect.Kind]func(reflect.Value, reflect.Value) bool
}

func NewContain() *Contain {
	contain := &Contain{}
	contain.kind2Contain = make(map[reflect.Kind]func(reflect.Value, reflect.Value) bool)
	contain.kind2Contain[reflect.String] = containString
	contain.kind2Contain[reflect.Int] = containInt
	contain.kind2Contain[reflect.Int8] = containInt
	contain.kind2Contain[reflect.Int16] = containInt
	contain.kind2Contain[reflect.Int32] = containInt
	contain.kind2Contain[reflect.Int64] = containInt
	contain.kind2Contain[reflect.Uint] = containUint
	contain.kind2Contain[reflect.Uint8] = containUint
	contain.kind2Contain[reflect.Uint16] = containUint
	contain.kind2Contain[reflect.Uint32] = containUint
	contain.kind2Contain[reflect.Uint64] = containUint
	return contain
}

func (contain *Contain) Exec(aSide, zSide []reflect.Value) bool {
	return Exec(aSide, zSide, contain.kind2Contain, "Contain")
}

func containString(aSide, zSide reflect.Value) bool {
	aside := removeSingleQuote(strings.ToLower(aSide.String()))
	zsideList := strings.ToLower(zSide.String())
	values := getContainStringList(zsideList)
	for _, v := range values {
		if aside == v {
			return true
		}
	}
	return false
}

func containInt(aSide, zSide reflect.Value) bool {
	aside, ok := getInt64(aSide)
	if !ok {
		return false
	}

	zsideList := strings.ToLower(zSide.String())

	values := getContainStringList(zsideList)
	for _, v := range values {
		intV, e := strconv.Atoi(v)
		if e != nil {
			return false
		}
		if aside == int64(intV) {
			return true
		}
	}
	return false
}

func containUint(aSide, zSide reflect.Value) bool {
	aside, ok := getUint64(aSide)
	if !ok {
		return false
	}

	zsideList := strings.ToLower(zSide.String())

	values := getContainStringList(zsideList)
	for _, v := range values {
		intV, e := strconv.Atoi(v)
		if e != nil {
			return false
		}
		if aside == uint64(intV) {
			return true
		}
	}
	return false
}

func getContainStringList(str string) []string {
	index := strings.Index(str, "{")
	index2 := strings.Index(str, "}")
	lst := str[index+1 : index2]
	values := strings.Split(lst, ",")
	result := make([]string, 0)
	for _, v := range values {
		result = append(result, removeSingleQuote(v))
	}
	return result
}

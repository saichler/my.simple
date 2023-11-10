package idql_query

import (
	"errors"
	"github.com/saichler/my.simple/go/idql/idql_parser"
	"github.com/saichler/my.simple/go/idql/idql_query/symbol_impl"
	"github.com/saichler/my.simple/go/utils/strng"
	"reflect"
	"strings"
)

type VarSymbol struct {
	aSide         string
	aSideInstance *introspect2.Instance
	symbol        idql_parser.Symbol
	zSide         string
	zSideInstance *introspect2.Instance
}

type VarSymbolImpl interface {
	Exec([]reflect.Value, []reflect.Value) bool
}

var symbolImpls = make(map[idql_parser.Symbol]VarSymbolImpl)

func init() {
	symbolImpls[idql_parser.Equal] = symbol_impl.NewEqual()
	symbolImpls[idql_parser.NotEqual] = symbol_impl.NewNotEqual()
	symbolImpls[idql_parser.Contain] = symbol_impl.NewContain()
	symbolImpls[idql_parser.NotContain] = symbol_impl.NewNotContain()
	symbolImpls[idql_parser.GreaterThan] = symbol_impl.NewGreaterThan()
	symbolImpls[idql_parser.GreaterOrEqual] = symbol_impl.NewGreaterThanOrEqual()
	symbolImpls[idql_parser.LessThan] = symbol_impl.NewLessThan()
	symbolImpls[idql_parser.LessOrEqual] = symbol_impl.NewLessThanOrEqual()
}

func (varSymbol *VarSymbol) String() string {
	s := &strng.String{}
	if varSymbol.aSideInstance != nil {
		s.Add(varSymbol.aSideInstance.InstanceId())
	} else {
		s.Add(varSymbol.aSide)
	}
	s.Add(varSymbol.symbol.String())
	if varSymbol.zSideInstance != nil {
		s.Add(varSymbol.zSideInstance.InstanceId())
	} else {
		s.Add(varSymbol.zSide)
	}
	return s.String()
}

func (varSymbol *VarSymbol) Simple() string {
	s := &strng.String{}
	isArrayOrMap := false
	if varSymbol.aSideInstance != nil {
		s.Add(varSymbol.aSideInstance.Node().FieldName)
		isArrayOrMap = varSymbol.aSideInstance.Node().IsSlice || varSymbol.aSideInstance.Node().IsMap
	} else {
		s.Add(varSymbol.aSide)
	}
	if isArrayOrMap {
		s.Add(" like ")
	} else {
		s.Add(varSymbol.symbol.String())
	}
	if varSymbol.zSideInstance != nil {
		s.Add(varSymbol.zSideInstance.Node().FieldName)
	} else {
		if isArrayOrMap {
			v := varSymbol.zSide
			if v[0] != '\'' {
				s.Add("'%").Add(v).Add("%'")
			} else {
				s.Add("'%").Add(v[1 : len(v)-1]).Add("%'")
			}
		} else {
			s.Add(varSymbol.zSide)
		}
	}
	return s.String()
}

func newVarSymbol(pVarSymbol *idql_parser.VarSymbol, elementType string) (*VarSymbol, error) {
	varSymbol := &VarSymbol{}
	varSymbol.symbol = pVarSymbol.Symbol()
	varSymbol.aSide = pVarSymbol.ASide()
	varSymbol.zSide = pVarSymbol.ZSide()
	aSide := varSymbol.aSide
	if !strings.HasPrefix(aSide, elementType) {
		s := strng.String{}
		s.Add(elementType)
		s.Add(".")
		s.Add(aSide)
		aSide = s.String()
	}
	zSide := varSymbol.zSide
	if !strings.HasPrefix(zSide, elementType) {
		s := strng.String{}
		s.Add(elementType)
		s.Add(".")
		s.Add(zSide)
		zSide = s.String()
	}
	aSideInstance, aSideErr := introspect2.Introspector.InstanceOf(aSide)
	varSymbol.aSideInstance = aSideInstance
	zSideInstance, zSideErr := introspect2.Introspector.InstanceOf(zSide)
	varSymbol.zSideInstance = zSideInstance

	if aSideErr != nil && zSideErr != nil {
		return nil, errors.New(aSideErr.Error() + " or " + zSideErr.Error())
	}

	return varSymbol, nil
}

func (varSymbol *VarSymbol) Match(any reflect.Value) (bool, error) {
	var aSideValues []reflect.Value
	var zSideValues []reflect.Value
	if varSymbol.aSideInstance != nil {
		aSideValues = varSymbol.aSideInstance.GetAsValues(any.Interface())
	} else {
		aSideValues = []reflect.Value{reflect.ValueOf(varSymbol.aSide)}
	}

	if varSymbol.zSideInstance != nil {
		zSideValues = varSymbol.zSideInstance.GetAsValues(any.Interface())
	} else {
		zSideValues = []reflect.Value{reflect.ValueOf(varSymbol.zSide)}
	}

	impl := symbolImpls[varSymbol.symbol]
	if impl == nil {
		panic("No impl found for: " + varSymbol.symbol + " symbol.")
	}
	return impl.Exec(aSideValues, zSideValues), nil
}

func (varSymbol *VarSymbol) isForType(typeName string) bool {
	if varSymbol.aSideInstance != nil && varSymbol.aSideInstance.Node().Parent.TypeName == typeName {
		return true
	}
	if varSymbol.zSideInstance != nil && varSymbol.zSideInstance.Node().Parent.TypeName == typeName {
		return true
	}
	return false
}

func (varSymbol *VarSymbol) ASideInstance() *introspect2.Instance {
	return varSymbol.aSideInstance
}

func (varSymbol *VarSymbol) ZSideInstance() *introspect2.Instance {
	return varSymbol.zSideInstance
}

func (varSymbol *VarSymbol) Symbol() idql_parser.Symbol {
	return varSymbol.symbol
}

func (varSymbol *VarSymbol) Aside() string {
	return varSymbol.aSide
}

func (varSymbol *VarSymbol) Zside() string {
	return varSymbol.zSide
}

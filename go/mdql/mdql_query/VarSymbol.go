package mdql_query

import (
	"errors"
	"fmt"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/instance"
	"github.com/saichler/my.simple/go/mdql/mdql_parser"
	"github.com/saichler/my.simple/go/mdql/mdql_query/symbol_impl"
	"github.com/saichler/my.simple/go/utils/strng"
	"reflect"
	"strings"
)

type VarSymbol struct {
	aSide         string
	aSideInstance *instance.Instance
	symbol        mdql_parser.Symbol
	zSide         string
	zSideInstance *instance.Instance
}

type VarSymbolImpl interface {
	Exec([]reflect.Value, []reflect.Value) bool
}

var symbolImpls = make(map[mdql_parser.Symbol]VarSymbolImpl)

func init() {
	symbolImpls[mdql_parser.Equal] = symbol_impl.NewEqual()
	symbolImpls[mdql_parser.NotEqual] = symbol_impl.NewNotEqual()
	symbolImpls[mdql_parser.Contain] = symbol_impl.NewContain()
	symbolImpls[mdql_parser.NotContain] = symbol_impl.NewNotContain()
	symbolImpls[mdql_parser.GreaterThan] = symbol_impl.NewGreaterThan()
	symbolImpls[mdql_parser.GreaterOrEqual] = symbol_impl.NewGreaterThanOrEqual()
	symbolImpls[mdql_parser.LessThan] = symbol_impl.NewLessThan()
	symbolImpls[mdql_parser.LessOrEqual] = symbol_impl.NewLessThanOrEqual()
}

func (varSymbol *VarSymbol) String() string {
	s := &strng.String{}
	if varSymbol.aSideInstance != nil {
		id, err := varSymbol.aSideInstance.InstanceId()
		if err != nil {
			fmt.Println("Error Aside ToString instanceId")
		}
		s.Add(id)
	} else {
		s.Add(varSymbol.aSide)
	}
	s.Add(varSymbol.symbol.String())
	if varSymbol.zSideInstance != nil {
		id, err := varSymbol.zSideInstance.InstanceId()
		if err != nil {
			fmt.Println("Error Zside ToString instanceId")
		}
		s.Add(id)
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

func newVarSymbol(pVarSymbol *mdql_parser.VarSymbol, elementType string, introspect common.IIntrospect) (*VarSymbol, error) {
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
	aSideInstance, aSideErr := instance.InstanceOf(aSide, introspect)
	varSymbol.aSideInstance = aSideInstance
	zSideInstance, zSideErr := instance.InstanceOf(zSide, introspect)
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

func (varSymbol *VarSymbol) ASideInstance() *instance.Instance {
	return varSymbol.aSideInstance
}

func (varSymbol *VarSymbol) ZSideInstance() *instance.Instance {
	return varSymbol.zSideInstance
}

func (varSymbol *VarSymbol) Symbol() mdql_parser.Symbol {
	return varSymbol.symbol
}

func (varSymbol *VarSymbol) Aside() string {
	return varSymbol.aSide
}

func (varSymbol *VarSymbol) Zside() string {
	return varSymbol.zSide
}

package idql_query

import (
	"bytes"
	"errors"
	"github.com/saichler/my.simple/go/idql/idql_parser"
	"github.com/saichler/my.simple/go/utils/strng"
	"reflect"
)

type CriteriaSymbol struct {
	varSymbol          *VarSymbol
	symbol             idql_parser.Symbol
	nextCriteriaSymbol *CriteriaSymbol
}

func newCriteriaSymbol(pCriteriaSymbol *idql_parser.CriteriaSymbol, elementType string) (*CriteriaSymbol, error) {
	criteriaSymbol := &CriteriaSymbol{}
	criteriaSymbol.symbol = pCriteriaSymbol.Symbol()
	varSymbol, e := newVarSymbol(pCriteriaSymbol.VarSymbol(), elementType)
	if e != nil {
		return nil, e
	}
	criteriaSymbol.varSymbol = varSymbol

	if pCriteriaSymbol.NextCriteriaSymbol() != nil {
		n, e := newCriteriaSymbol(pCriteriaSymbol.NextCriteriaSymbol(), elementType)
		if e != nil {
			return nil, e
		}
		criteriaSymbol.nextCriteriaSymbol = n
	}
	return criteriaSymbol, nil
}

func (criteriaSymbol *CriteriaSymbol) String() string {
	s := &strng.String{}
	s.Add("(")
	criteriaSymbol.toString(s)
	s.Add(")")
	return s.String()
}

func (criteriaSymbol *CriteriaSymbol) toString(s *strng.String) {
	if criteriaSymbol.varSymbol != nil {
		s.Add(criteriaSymbol.varSymbol.String())
	}
	if criteriaSymbol.nextCriteriaSymbol != nil {
		s.Add(criteriaSymbol.symbol.String())
		criteriaSymbol.nextCriteriaSymbol.toString(s)
	}
}

func (criteriaSymbol *CriteriaSymbol) Match(any reflect.Value) (bool, error) {
	m, e := criteriaSymbol.varSymbol.Match(any)
	if e != nil {
		return false, e
	}
	nm := true
	if criteriaSymbol.symbol == idql_parser.Or {
		nm = false
	}
	if criteriaSymbol.nextCriteriaSymbol != nil {
		nm, e = criteriaSymbol.nextCriteriaSymbol.Match(any)
		if e != nil {
			return false, e
		}
	}
	if criteriaSymbol.symbol == "" {
		return nm && m, nil
	}
	if criteriaSymbol.symbol == idql_parser.And {
		return m && nm, nil
	}
	if criteriaSymbol.symbol == idql_parser.Or {
		return m || nm, nil
	}
	return false, errors.New("Unsupported symbol :" + criteriaSymbol.symbol.String())
}

func (criteriaSymbol *CriteriaSymbol) VarSymbolsForType(typeName string) string {
	buff := bytes.Buffer{}
	if criteriaSymbol.varSymbol != nil && criteriaSymbol.varSymbol.isForType(typeName) {
		buff.WriteString(criteriaSymbol.varSymbol.Simple())
	}
	if criteriaSymbol.nextCriteriaSymbol != nil {
		buff.WriteString(string(criteriaSymbol.symbol))
		buff.WriteString(criteriaSymbol.nextCriteriaSymbol.VarSymbolsForType(typeName))
	}
	return buff.String()
}

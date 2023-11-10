package idql_query

import (
	"bytes"
	"errors"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/idql/idql_parser"
	"github.com/saichler/my.simple/go/utils/strng"
	"reflect"
)

type Criteria struct {
	criteriaSymbol *CriteriaSymbol
	symbol         idql_parser.Symbol
	nextCriteria   *Criteria
	subCriteria    *Criteria
}

func (criteria *Criteria) String() string {
	s := &strng.String{}
	if criteria.criteriaSymbol != nil {
		s.Add(criteria.criteriaSymbol.String())
	} else {
		s.Add("(")
	}
	if criteria.subCriteria != nil {
		s.Add(criteria.subCriteria.String())
	}
	if criteria.criteriaSymbol == nil {
		s.Add(")")
	}
	if criteria.nextCriteria != nil {
		s.Add(criteria.symbol.String())
		s.Add(criteria.nextCriteria.String())
	}
	return s.String()
}

func newCriteria(pCriteria *idql_parser.Criteria, elementType string, introspect common.IIntrospect) (*Criteria, error) {
	if pCriteria == nil {
		return nil, nil
	}
	criteria := &Criteria{}
	criteria.symbol = pCriteria.Symbol()

	if pCriteria.CriteriaSymbol() != nil {
		criteriaSymbol, e := newCriteriaSymbol(pCriteria.CriteriaSymbol(), elementType, introspect)
		if e != nil {
			return nil, e
		}
		criteria.criteriaSymbol = criteriaSymbol
	}

	if criteria.subCriteria != nil {
		subCriteria, e := newCriteria(pCriteria.SubCriteria(), elementType, introspect)
		if e != nil {
			return nil, e
		}
		criteria.subCriteria = subCriteria
	}

	if pCriteria.NextCriteria() != nil {
		nextCriteria, e := newCriteria(pCriteria.NextCriteria(), elementType, introspect)
		if e != nil {
			return nil, e
		}
		criteria.nextCriteria = nextCriteria
	}

	return criteria, nil
}

func (criteria *Criteria) Match(any reflect.Value) (bool, error) {
	cSymbol := true
	cSub := true
	cNext := true
	var e error
	if criteria.symbol == idql_parser.Or {
		cSymbol = false
		cSub = false
		cNext = false
	}
	if criteria.criteriaSymbol != nil {
		cSymbol, e = criteria.criteriaSymbol.Match(any)
		if e != nil {
			return false, e
		}
	}
	if criteria.subCriteria != nil {
		cSub, e = criteria.subCriteria.Match(any)
		if e != nil {
			return false, e
		}
	}
	if criteria.nextCriteria != nil {
		cNext, e = criteria.nextCriteria.Match(any)
		if e != nil {
			return false, e
		}
	}
	if criteria.symbol == "" {
		return cSub && cNext && cSymbol, nil
	}
	if criteria.symbol == idql_parser.And {
		return cSub && cNext && cSymbol, nil
	}
	if criteria.symbol == idql_parser.Or {
		return cSub || cNext || cSymbol, nil
	}

	return false, errors.New("Unsupported Criteria symbol " + string(criteria.symbol))
}

func (criteria *Criteria) VarSymbolsForType(typeName string) string {
	buff := bytes.Buffer{}
	if criteria.criteriaSymbol != nil {
		buff.WriteString(criteria.criteriaSymbol.VarSymbolsForType(typeName))
	}
	if criteria.subCriteria != nil {
		buff.WriteString("(")
		buff.WriteString(criteria.subCriteria.VarSymbolsForType(typeName))
		buff.WriteString(")")
	}
	if criteria.nextCriteria != nil {
		buff.WriteString(string(criteria.symbol))
		buff.WriteString(criteria.nextCriteria.VarSymbolsForType(typeName))
	}
	return buff.String()
}

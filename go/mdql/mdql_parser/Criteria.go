package mdql_parser

import (
	"github.com/saichler/my.simple/go/utils/strng"
	"strings"
)

type Criteria struct {
	criteriaSymbol *CriteriaSymbol
	symbol         Symbol
	nextCriteria   *Criteria
	subCriteria    *Criteria
}

func (criteria *Criteria) CriteriaSymbol() *CriteriaSymbol {
	return criteria.criteriaSymbol
}

func (criteria *Criteria) Symbol() Symbol {
	return criteria.symbol
}

func (criteria *Criteria) NextCriteria() *Criteria {
	return criteria.nextCriteria
}

func (criteria *Criteria) SubCriteria() *Criteria {
	return criteria.subCriteria
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

func (criteria *Criteria) Print(ind int) string {
	s := &strng.String{}
	s.Add(printIndent(ind))
	s.Add("Expression\n")
	if criteria.criteriaSymbol != nil {
		s.Add(criteria.criteriaSymbol.Print(ind + 1))
	}
	if criteria.subCriteria != nil {
		s.Add(criteria.subCriteria.Print(ind + 1))
	}
	if criteria.nextCriteria != nil {
		s.Add(printIndent(ind))
		s.Add(strings.TrimSpace(string(criteria.symbol.String())))
		s.Add("\n")
		s.Add(criteria.nextCriteria.Print(ind))
	}
	return s.String()
}

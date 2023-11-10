package idql_parser

import (
	"errors"
	"github.com/saichler/my.simple/go/utils/strng"
	"strings"
)

type CriteriaSymbol struct {
	varSymbol          *VarSymbol
	symbol             Symbol
	nextCriteriaSymbol *CriteriaSymbol
}

func (criteriaSymbol *CriteriaSymbol) VarSymbol() *VarSymbol {
	return criteriaSymbol.varSymbol
}

func (criteriaSymbol *CriteriaSymbol) Symbol() Symbol {
	return criteriaSymbol.symbol
}

func (criteriaSymbol *CriteriaSymbol) NextCriteriaSymbol() *CriteriaSymbol {
	return criteriaSymbol.nextCriteriaSymbol
}

func (criteriaSymbol *CriteriaSymbol) String() string {
	s := &strng.String{}
	s.Add("(")
	criteriaSymbol.toString(s)
	s.Add(")")
	return s.String()
}

func (criteriaSymbol *CriteriaSymbol) Print(lvl int) string {
	s := &strng.String{}
	s.Add(printIndent(lvl))
	s.Add("Criteria Symbol\n")
	if criteriaSymbol.varSymbol != nil {
		s.Add(criteriaSymbol.varSymbol.Print(lvl + 1))
	}
	if criteriaSymbol.nextCriteriaSymbol != nil {
		s.Add(printIndent(lvl))
		s.Add(strings.TrimSpace(criteriaSymbol.symbol.String()))
		s.Add("\n")
		s.Add(criteriaSymbol.nextCriteriaSymbol.Print(lvl))
	}
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

func newCriteriaSymbol(expression string) (*CriteriaSymbol, error) {
	if expression == "" {
		return nil, nil
	}
	index := MAX
	var symbol Symbol
	and := strings.Index(expression, And.String())
	if and != -1 {
		index = and
		symbol = And
	}
	or := strings.Index(expression, Or.String())
	if or != -1 && or < index {
		index = or
		symbol = Or
	}

	criteriaSymbol := &CriteriaSymbol{}
	if index == MAX {
		varSymbol, e := newVarSymbol(expression)
		if e != nil {
			return nil, e
		}
		criteriaSymbol.varSymbol = varSymbol
		return criteriaSymbol, nil
	}

	varSymbol, e := newVarSymbol(expression[0:index])
	if e != nil {
		return nil, e
	}

	criteriaSymbol.varSymbol = varSymbol
	criteriaSymbol.symbol = symbol

	expression = expression[index+len(symbol):]
	n, e := newCriteriaSymbol(expression)
	if e != nil {
		return nil, e
	}

	criteriaSymbol.nextCriteriaSymbol = n
	return criteriaSymbol, nil
}

func lastSymbol(expression string) (Symbol, int, error) {
	index := -1
	var symbol Symbol

	and := strings.LastIndex(expression, And.String())
	if and > index {
		symbol = And
		index = and
	}

	or := strings.LastIndex(expression, Or.String())
	if or > index {
		symbol = Or
		index = or
	}

	if index == -1 {
		return "", 0, errors.New("no last Criteria Symbol was found")
	}
	return symbol, index, nil
}

func firstSymbol(expression string) (Symbol, int, error) {
	index := MAX
	var symbol Symbol
	and := strings.Index(expression, And.String())
	if and != -1 {
		index = and
		symbol = And
	}
	or := strings.Index(expression, Or.String())
	if or != -1 && or < index {
		index = or
		symbol = Or
	}

	if index == MAX {
		return "", 0, errors.New("no first Criteria Symbol was found")
	}

	return symbol, index, nil
}

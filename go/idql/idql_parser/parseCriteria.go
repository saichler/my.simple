package idql_parser

import (
	"errors"
	"strings"
)

func parseCriteria(expression string) (*Criteria, error) {
	if expression == "" {
		return nil, nil
	}
	expression = strings.TrimSpace(expression)
	begin := bracketBegin(expression)
	if begin == -1 {
		return parseNoBrackets(expression)
	}
	if begin > 0 {
		return parseBeforeBrackets(expression, begin)
	}
	return parseWithBrackets(expression, begin)
}

func parseWithBrackets(expression string, begin int) (*Criteria, error) {
	end, e := bracketEnd(expression, begin)
	if e != nil {
		return nil, e
	}
	criteria := &Criteria{}
	subCriteria, e := parseCriteria(expression[1:end])
	if e != nil {
		return nil, e
	}

	criteria.subCriteria = subCriteria

	if end < len(expression)-1 {
		symbol, index, e := firstSymbol(expression[end+1:])
		if e != nil {
			return nil, e
		}
		criteria.symbol = symbol
		nextCriteria, e := parseCriteria(expression[end+1+index+len(symbol):])
		if e != nil {
			return nil, e
		}
		criteria.nextCriteria = nextCriteria
	}
	return criteria, nil
}

func parseBeforeBrackets(expression string, begin int) (*Criteria, error) {
	prefix := expression[0:begin]
	symbol, index, e := lastSymbol(prefix)
	if e != nil {
		return nil, e
	}
	criteria, e := parseNoBrackets(prefix[0:index])
	if e != nil {
		return nil, e
	}
	criteria.symbol = symbol
	nextCriteria, e := parseCriteria(expression[begin:])
	if e != nil {
		return nil, e
	}
	criteria.nextCriteria = nextCriteria
	return criteria, nil
}

func parseNoBrackets(ws string) (*Criteria, error) {
	if ws == "" {
		return nil, nil
	}
	criteria := &Criteria{}
	criteriaSymbol, e := newCriteriaSymbol(ws)
	if e != nil {
		return nil, e
	}
	criteria.criteriaSymbol = criteriaSymbol
	return criteria, nil
}

func bracketBegin(expression string) int {
	return strings.Index(expression, "(")
}

func bracketEnd(expression string, begin int) (int, error) {
	count := 0
	for i := begin; i < len(expression); i++ {
		if byte(expression[i]) == byte('(') {
			count++
		} else if byte(expression[i]) == byte(')') {
			count--
		}
		if count == 0 {
			return i, nil
		}
	}
	return -1, errors.New("missing close bracket in expression " + expression)
}

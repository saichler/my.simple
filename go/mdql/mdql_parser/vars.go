package mdql_parser

import (
	"bytes"
	"strings"
)

type Symbol string

const (
	Equal          Symbol = "="
	NotEqual       Symbol = "!="
	GreaterThan    Symbol = ">"
	LessThan       Symbol = "<"
	GreaterOrEqual Symbol = ">="
	LessOrEqual    Symbol = "<="
	Contain        Symbol = " contain "
	NotContain     Symbol = " not contain "
	And            Symbol = " and "
	Or             Symbol = " or "
)

const (
	FETCH    = "fetch "
	ONLY     = " only "
	CRITERIA = " criteria "
	SORT     = " sort "
	PAGE     = " page "
	LIMIT    = " limit "
	MAX      = 99999
)

var reserved = map[string]string{FETCH: FETCH, ONLY: ONLY, CRITERIA: CRITERIA, SORT: SORT, PAGE: PAGE, LIMIT: LIMIT}

var varsSymbols = make([]Symbol, 0)
var initParser = initParserVars()

func initParserVars() bool {
	varsSymbols = append(varsSymbols, NotContain)
	varsSymbols = append(varsSymbols, Contain)
	varsSymbols = append(varsSymbols, LessOrEqual)
	varsSymbols = append(varsSymbols, GreaterOrEqual)
	varsSymbols = append(varsSymbols, GreaterThan)
	varsSymbols = append(varsSymbols, LessThan)
	varsSymbols = append(varsSymbols, GreaterThan)
	varsSymbols = append(varsSymbols, NotEqual)
	varsSymbols = append(varsSymbols, Equal)
	return true
}

func (symb Symbol) String() string {
	return string(symb)
}

func printIndent(ind int) string {
	buff := bytes.Buffer{}
	buff.WriteString("|")
	for i := 0; i < ind; i++ {
		buff.WriteString("--")
	}
	return buff.String()
}

func validateVar(azValue string) string {
	open := strings.Index(azValue, "(")
	close := strings.Index(azValue, ")")
	if open != -1 || close != -1 {
		return azValue + " has either only '(' or ')' symbols"
	}
	return ""
}

package idql_parser

import (
	"bytes"
	"errors"
	"github.com/saichler/my.simple/go/utils/strng"
	"strings"
)

type VarSymbol struct {
	aSide  string
	symbol Symbol
	zSide  string
}

func (varSymbol *VarSymbol) ASide() string {
	return varSymbol.aSide
}

func (varSymbol *VarSymbol) ZSide() string {
	return varSymbol.zSide
}

func (varSymbol *VarSymbol) Symbol() Symbol {
	return varSymbol.symbol
}

func (varSymbol *VarSymbol) String() string {
	s := &strng.String{}
	s.Add(varSymbol.aSide)
	s.Add(varSymbol.symbol.String())
	s.Add(varSymbol.zSide)
	return s.String()
}

func (varSymbol *VarSymbol) Print(lvl int) string {
	s := &strng.String{}
	s.Add(printIndent(lvl))
	s.Add("Var Symbol (")
	s.Add(varSymbol.aSide)
	s.Add(varSymbol.symbol.String())
	s.Add(varSymbol.zSide)
	s.Add(")\n")
	return s.String()
}

func newVarSymbol(expression string) (*VarSymbol, error) {
	for _, symb := range varsSymbols {
		index := strings.Index(expression, symb.String())
		if index != -1 {
			sym := &VarSymbol{}
			sym.aSide = strings.TrimSpace(expression[0:index])
			sym.zSide = strings.TrimSpace(expression[index+len(symb):])
			sym.symbol = symb
			if validateVar(sym.aSide) != "" {
				return nil, errors.New(validateVar(sym.aSide))
			}
			if validateVar(sym.zSide) != "" {
				return nil, errors.New(validateVar(sym.zSide))
			}
			return sym, nil
		}
	}
	return nil, errors.New("Cannot find a Symbol in: " + expression)
}

func ToLower(str string) string {
	buff := bytes.Buffer{}
	open := false
	for i, c := range str {
		if c == '\'' {
			open = !open
		}
		if !open {
			buff.WriteString(strings.ToLower(str[i : i+1]))
		} else {
			buff.WriteString(str[i : i+1])
		}
	}
	return buff.String()
}

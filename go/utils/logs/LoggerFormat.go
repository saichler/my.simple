package logs

import "github.com/saichler/my.simple/go/utils/strng"

func toString(tag string, any interface{}, anys ...interface{}) string {
	str := strng.New(tag)
	str.TypesPrefix = false
	str.AddSpaceWhenAdding = true
	s, _ := str.StringOf(any)
	str.Add(s)
	if anys != nil {
		for _, a := range anys {
			s, _ = str.StringOf(a)
			str.Add(s)
		}
	}
	return str.String()
}

func TraceToString(any interface{}, anys ...interface{}) string {
	return toString("Tr ->", any, anys...)
}

func DebugToString(any interface{}, anys ...interface{}) string {
	return toString(" Dg ->", any, anys...)
}

func InfoToString(any interface{}, anys ...interface{}) string {
	return toString("  In ->", any, anys...)
}

func WarningToString(any interface{}, anys ...interface{}) string {
	return toString("   Wr ->", any, anys...)
}

func ErrorToString(any interface{}, anys ...interface{}) string {
	return toString("    Er ->", any, anys...)
}

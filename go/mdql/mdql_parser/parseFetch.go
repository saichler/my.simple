package mdql_parser

import (
	"github.com/saichler/my.simple/go/utils/strng"
	"strings"
)

func parseFetch(fetch string) map[string]string {
	result := make(map[string]string)
	fetch = ToLower(strings.TrimSpace(fetch))
	for reservedWord, _ := range reserved {
		openB := false
		openQ := false
		openV := false
		word := strng.New("")
		value := strng.New("")

		for _, c := range fetch {
			if c == '{' {
				openB = true
			} else if c == '}' {
				openB = false
			} else if c == '\'' {
				openQ = !openQ
			}

			if openV {
				value.Add(string(c))
				strValue := value.String()
				found := false
				for _, res := range reserved {
					if len(strValue) >= len(res) && strValue[len(strValue)-len(res):] == res && !openQ && !openB {
						result[reservedWord] = strValue[0 : len(strValue)-len(res)]
						openV = false
						found = true
						break
					}
				}
				if found {
					break
				}
			}

			if !openB && !openQ && !openV {
				word.Add(string(c))
				if word.Len() > len(reservedWord) {
					temp := strng.New("")
					temp.AddBytes(word.Bytes()[1:])
					word = temp
				}
				if word.String() == reservedWord {
					openV = true
				}
			}
		}
		if openV {
			result[reservedWord] = value.String()
		}
	}
	return result
}

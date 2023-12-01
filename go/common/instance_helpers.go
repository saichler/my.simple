package common

import "github.com/saichler/my.simple/go/utils/strng"

func NodeKey(instanceId string) string {
	buff := &strng.String{}
	open := false
	for _, c := range instanceId {
		if c == '<' {
			open = true
		} else if c == '>' {
			open = false
		} else if !open {
			buff.Add(string(c))
		}
	}
	return buff.String()
}

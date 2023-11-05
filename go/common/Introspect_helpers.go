package common

import (
	"github.com/saichler/my.simple/go/introspect/model"
	"reflect"
	"strings"
)

func ValueTypeLower(any interface{}) (reflect.Value, reflect.Type, string) {
	v := reflect.ValueOf(any)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	ts := strings.ToLower(t.Name())
	return v, t, ts
}

func IsLeaf(node *model.Node) bool {
	if node.Attributes == nil || len(node.Attributes) == 0 {
		return true
	}
	return false
}

func IsRoot(node *model.Node) bool {
	if node.Parent == nil {
		return true
	}
	return false
}

func IgnoreName(fieldName string) bool {
	if fieldName == "DoNotCompare" {
		return true
	}
	if fieldName == "DoNotCopy" {
		return true
	}
	if len(fieldName) > 3 && fieldName[0:3] == "XXX" {
		return true
	}
	if fieldName[0:1] == strings.ToLower(fieldName[0:1]) {
		return true
	}
	return false
}

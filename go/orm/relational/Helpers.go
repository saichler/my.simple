package relational

import (
	"github.com/google/uuid"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/introspect/model"
	"github.com/saichler/my.simple/go/utils/strng"
	"reflect"
)

func newRow() *Row {
	return &Row{colValues: make(map[string]reflect.Value)}
}

func keyOf(key, value reflect.Value, node *model.Node, path, attr string, inspect common.IIntrospect) string {
	primary := inspect.StringOfPrimaryDecorator(node, value)
	//This is a root key
	if path == "" {
		if primary == "" && key.IsValid() {
			return strng.New(node.TypeName, "<", toString(key), ">").String()
		}
		if primary != "" {
			return strng.New(node.TypeName, "<", primary, ">").String()
		}
		//No key for the item was found and this is a root element
		if primary == "" && path == "" {
			return strng.New(node.TypeName, "<", uuid.New().String(), ">").String()
		}
	}
	if primary != "" {
		return strng.New(path, "<", primary, ">").String()
	}
	if key.IsValid() {
		str := strng.New()
		str.TypesPrefix = true
		keyString := str.ToString(key)
		return strng.New(path, ".", attr, "<", keyString, ">").String()
	}
	return strng.New(path, ".", attr).String()
}

func tableName(value reflect.Value) string {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return value.Type().Name()
}

func removePtr(value reflect.Value) reflect.Value {
	if value.Kind() == reflect.Ptr {
		return value.Elem()
	}
	return value
}

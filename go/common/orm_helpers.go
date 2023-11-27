package common

import (
	"github.com/saichler/my.simple/go/introspect/model"
	"github.com/saichler/my.simple/go/utils/strng"
	"reflect"
)

func PrimaryDecorator(node *model.Node, value reflect.Value) string {
	fields := PrimaryDecoratorFields(node)
	if fields == nil {
		return ""
	}
	str := strng.New()
	str.TypesPrefix = true
	for _, field := range fields {
		v := value.FieldByName(field).Interface()
		str.Add(str.StringOf(v))
	}
	return str.String()
}

func PrimaryDecoratorFields(node *model.Node) []string {
	decValue := node.Decorators[int32(model.DecoratorType_Primary)]
	fields, ok := strng.InstanceOf(decValue).([]string)
	if !ok {
		return nil
	}
	return fields
}

func DeepDecorator(node *model.Node) bool {
	decValue := node.Decorators[int32(model.DecoratorType_Deep)]
	if decValue == "" {
		return false
	}
	return true
}

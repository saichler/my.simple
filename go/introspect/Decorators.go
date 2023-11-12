package introspect

import (
	"github.com/saichler/my.simple/go/introspect/model"
	"github.com/saichler/my.simple/go/utils/strng"
	"reflect"
)

func (i *Introspect) AddDecorator(decoratorType model.DecoratorType, any interface{}, node *model.Node) {
	s := &strng.String{TypesPrefix: true}
	str := s.StringOf(any)
	if node.Decorators == nil {
		node.Decorators = make(map[int32]string)
	}
	node.Decorators[int32(decoratorType)] = str
}

func (i *Introspect) DecoratorOf(decoratorType model.DecoratorType, node *model.Node) interface{} {
	decValue := node.Decorators[int32(decoratorType)]
	v := strng.InstanceOf(decValue)
	return v
}

func (i *Introspect) StringOfPrimaryDecorator(node *model.Node, value reflect.Value) string {
	decValue := node.Decorators[int32(model.DecoratorType_Primary)]
	fields, ok := strng.InstanceOf(decValue).([]string)
	if !ok {
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

func (i *Introspect) DeepDecorator(node *model.Node) bool {
	decValue := node.Decorators[int32(model.DecoratorType_Deep)]
	if decValue == "" {
		return false
	}
	return true
}

package introspect

import (
	"github.com/saichler/my.simple/go/introspect/model"
	"github.com/saichler/my.simple/go/utils/strng"
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

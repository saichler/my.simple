package introspect

import (
	"github.com/saichler/my.simple/go/introspect/model"
	"github.com/saichler/my.simple/go/utils/strng"
)

func AddDecorator(decoratorType model.DecoratorType, any interface{}, node *model.Node) error {
	s := &strng.String{TypesPrefix: true}
	str, err := s.StringOf(any)
	if node.Decorators == nil {
		node.Decorators = make(map[int32]string)
	}
	node.Decorators[int32(decoratorType)] = str
	return err
}

func DecoratorOf(decoratorType model.DecoratorType, node *model.Node) (interface{}, error) {
	decValue := node.Decorators[int32(decoratorType)]
	v, err := strng.InstanceOf(decValue)
	return v, err
}

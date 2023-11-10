package common

import (
	"github.com/saichler/my.simple/go/introspect/model"
	"reflect"
)

type IIntrospect interface {
	Inspect(interface{}) (*model.Node, error)
	Node(string) (*model.Node, bool)
	NodeByType(interface{}) (*model.Node, bool)
	Nodes(bool, bool) []*model.Node
	Print()
	Registry() IRegistry
	Kind(*model.Node) reflect.Kind
	Clone(interface{}) interface{}
	AddDecorator(model.DecoratorType, interface{}, *model.Node) error
	DecoratorOf(model.DecoratorType, *model.Node) (interface{}, error)
}

var Introspect IIntrospect

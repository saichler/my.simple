package introspect

import (
	"github.com/saichler/my.simple/go/introspect/model"
	"github.com/saichler/my.simple/go/utils/registry"
	"github.com/saichler/my.simple/go/utils/strng"
	"reflect"
	"strings"
)

func NodeKey(node *model.Node) string {
	if node.CachedKey != "" {
		return node.CachedKey
	}
	if node.Parent == nil {
		return strings.ToLower(node.TypeName)
	}
	buff := &strng.String{}
	buff.Add(NodeKey(node.Parent))
	buff.Add(".")
	buff.Add(strings.ToLower(node.FieldName))
	node.CachedKey = buff.String()
	return node.CachedKey
}

func addAttribute(node *model.Node, _type reflect.Type, _fieldName string) *model.Node {
	registry.RegisterStruct(_type)
	if node != nil && node.Attributes == nil {
		node.Attributes = make(map[string]*model.Node)
	}

	subNode := &model.Node{}
	subNode.TypeName = _type.Name()
	subNode.Parent = node
	subNode.FieldName = _fieldName

	if node != nil {
		node.Attributes[subNode.FieldName] = subNode
	}
	return subNode
}

func Kind(node *model.Node) reflect.Kind {
	t, err := registry.TypeByName(node.TypeName)
	if err != nil {
		panic(err.Error())
	}
	return t.Kind()
}

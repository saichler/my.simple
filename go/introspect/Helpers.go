package introspect

import (
	"fmt"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/introspect/model"
	"reflect"
)

func (i *Introspect) addAttribute(node *model.Node, _type reflect.Type, _fieldName string) *model.Node {
	i.registry.RegisterStruct(_type)
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

func (i *Introspect) addNode(_type reflect.Type, _parent *model.Node, _fieldName string) (*model.Node, bool) {
	exist, ok := i.typeToNode.Get(_type.Name())
	if ok && !common.IsLeaf(exist) {
		clone := i.cloner.Clone(exist).(*model.Node)
		clone.Parent = _parent
		clone.FieldName = _fieldName
		clone.CachedKey = ""
		nodePath := NodeKey(clone)
		i.pathToNode.Put(nodePath, clone)
		return clone, true
	}

	node := i.addAttribute(_parent, _type, _fieldName)
	nodePath := NodeKey(node)
	_, ok = i.pathToNode.Get(nodePath)
	if ok {
		return nil, false
	}
	i.pathToNode.Put(nodePath, node)
	if _type.Kind() == reflect.Struct {
		i.typeToNode.Put(node.TypeName, node)
	}
	return node, false
}

func (i *Introspect) inspectStruct(_type reflect.Type, _parent *model.Node, _fieldName string) *model.Node {
	node, isClone := i.addNode(_type, _parent, _fieldName)
	if isClone {
		return node
	}
	i.registry.RegisterStructType(_type)
	for index := 0; index < _type.NumField(); index++ {
		field := _type.Field(index)
		if common.IgnoreName(field.Name) {
			continue
		}
		if field.Type.Kind() == reflect.Slice {
			i.inspectSlice(field.Type, node, field.Name)
		} else if field.Type.Kind() == reflect.Map {
			i.inspectMap(field.Type, node, field.Name)
		} else if field.Type.Kind() == reflect.Ptr {
			subnode := i.inspectPtr(field.Type.Elem(), node, field.Name)
			i.typeToNode.Put(subnode.TypeName, subnode)
		} else {
			i.addNode(field.Type, node, field.Name)
		}
	}
	return node
}

func (i *Introspect) inspectPtr(_type reflect.Type, _parent *model.Node, _fieldName string) *model.Node {
	switch _type.Kind() {
	case reflect.Struct:
		return i.inspectStruct(_type, _parent, _fieldName)
	}
	panic("unknown ptr kind " + _type.Kind().String())
}

func (i *Introspect) inspectMap(_type reflect.Type, _parent *model.Node, _fieldName string) *model.Node {
	if _type.Elem().Kind() == reflect.Ptr && _type.Elem().Elem().Kind() == reflect.Struct {
		subNode := i.inspectStruct(_type.Elem().Elem(), _parent, _fieldName)
		subNode.IsMap = true
		_parent.Attributes[_fieldName] = subNode
		return subNode
	} else {
		subNode, _ := i.addNode(_type.Elem(), _parent, _fieldName)
		subNode.IsMap = true
		return node
	}
}

func (i *Introspect) inspectSlice(_type reflect.Type, _parent *model.Node, _fieldName string) *model.Node {
	if _type.Elem().Kind() == reflect.Ptr && _type.Elem().Elem().Kind() == reflect.Struct {
		subNode := i.inspectStruct(_type.Elem().Elem(), _parent, _fieldName)
		subNode.IsSlice = true
		_parent.Attributes[_fieldName] = subNode
		return subNode
	} else {
		subNode, _ := i.addNode(_type.Elem(), _parent, _fieldName)
		subNode.IsSlice = true
		return subNode
	}
}

func printDo(key, val interface{}) {
	node := val.(*model.Node)
	fmt.Println(key, "-", node.TypeName, ", map=", node.IsMap, ", slice=", node.IsSlice, ", leaf=", common.IsLeaf(node))
}

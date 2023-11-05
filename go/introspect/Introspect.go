package introspect

import (
	"fmt"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/introspect/model"
	"github.com/saichler/my.simple/go/utils/logs"
	"github.com/saichler/my.simple/go/utils/maps"
	"github.com/saichler/my.simple/go/utils/registry"
	"reflect"
)

type Introspect struct {
	pathToNode *maps.IntrospectNodeMap
	typeToNode *maps.IntrospectNodeMap
}

var DefaultIntrospect = NewIntrospect()

func NewIntrospect() *Introspect {
	i := &Introspect{}
	i.pathToNode = maps.NewIntrospectNodeMap()
	i.typeToNode = maps.NewIntrospectNodeMap()
	return i
}

func Inspect(any interface{}) (*model.Node, error) {
	return DefaultIntrospect.Inspect(any)
}

func (i *Introspect) Inspect(any interface{}) (*model.Node, error) {
	if any == nil {
		return nil, logs.Error("Cannot introspect a nil value")
	}
	registry.RegisterStruct(any)
	_, t, ts := common.ValueTypeLower(any)
	if t.Kind() != reflect.Struct {
		return nil, logs.Error("Cannot introspect a value that is not a struct")
	}
	node, ok := i.pathToNode.Get(ts)
	if ok {
		return node, nil
	}
	return i.inspectStruct(t, nil, ""), nil
}

func (i *Introspect) Node(path string) (*model.Node, bool) {
	return i.pathToNode.Get(path)
}

func (i *Introspect) NodeByType(any interface{}) (*model.Node, bool) {
	typ := reflect.ValueOf(any)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return i.typeToNode.Get(typ.Type().Name())
}

func (i *Introspect) Nodes(onlyLeafs, onlyRoots bool) []*model.Node {

	filter := func(any interface{}) bool {
		node := any.(*model.Node)
		if onlyLeafs && !common.IsLeaf(node) {
			return false
		}
		if onlyRoots && !common.IsRoot(node) {
			return false
		}
		return true
	}

	return i.pathToNode.NodesList(filter)
}

func (i *Introspect) addNode(_type reflect.Type, _parent *model.Node, _fieldName string) (*model.Node, bool) {
	exist, ok := i.typeToNode.Get(_type.Name())
	if ok && !common.IsLeaf(exist) {
		clone := Clone(exist).(*model.Node)
		clone.Parent = _parent
		clone.FieldName = _fieldName
		nodePath := NodeKey(clone)
		i.pathToNode.Put(nodePath, clone)
		return clone, true
	}

	node := addAttribute(_parent, _type, _fieldName)
	nodePath := NodeKey(node)
	_, ok = i.pathToNode.Get(nodePath)
	if ok {
		return nil, false
	}
	i.pathToNode.Put(nodePath, node)
	i.typeToNode.Put(node.TypeName, node)
	return node, false
}

func (i *Introspect) inspectStruct(_type reflect.Type, _parent *model.Node, _fieldName string) *model.Node {
	node, isClone := i.addNode(_type, _parent, _fieldName)
	if isClone {
		return node
	}
	for index := 0; index < _type.NumField(); index++ {
		field := _type.Field(index)
		if common.IgnoreName(field.Name) {
			continue
		}
		if field.Type.Kind() == reflect.Slice {
			subNode := i.inspectSlice(field.Type, node, field.Name)
			subNode.IsSlice = true
		} else if field.Type.Kind() == reflect.Map {
			subNode := i.inspectMap(field.Type, node, field.Name)
			subNode.IsMap = true
		} else if field.Type.Kind() == reflect.Ptr {
			i.inspectPtr(field.Type.Elem(), node, field.Name)
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
		return i.inspectStruct(_type.Elem().Elem(), _parent, _fieldName)
	} else {
		node, _ := i.addNode(_type.Elem(), _parent, _fieldName)
		return node
	}
}

func (i *Introspect) inspectSlice(_type reflect.Type, _parent *model.Node, _fieldName string) *model.Node {
	if _type.Elem().Kind() == reflect.Ptr && _type.Elem().Elem().Kind() == reflect.Struct {
		return i.inspectStruct(_type.Elem().Elem(), _parent, _fieldName)
	} else {
		node, _ := i.addNode(_type.Elem(), _parent, _fieldName)
		return node
	}
}

func printDo(key, val interface{}) {
	node := val.(*model.Node)
	fmt.Println(key, "-", node.TypeName, ", map=", node.IsMap, ", slice=", node.IsSlice, ", leaf=", common.IsLeaf(node))
}

func (i *Introspect) Print() {
	i.pathToNode.Iterate(printDo)
}

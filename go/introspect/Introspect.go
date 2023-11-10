package introspect

import (
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/introspect/model"
	"github.com/saichler/my.simple/go/utils/logs"
	"github.com/saichler/my.simple/go/utils/maps"
	"github.com/saichler/my.simple/go/utils/strng"
	"reflect"
	"strings"
)

type Introspect struct {
	pathToNode *maps.IntrospectNodeMap
	typeToNode *maps.IntrospectNodeMap
	registry   common.IRegistry
	cloner     *Cloner
}

func NewIntrospect(registry common.IRegistry) *Introspect {
	i := &Introspect{}
	i.registry = registry
	i.cloner = newCloner()
	i.pathToNode = maps.NewIntrospectNodeMap()
	i.typeToNode = maps.NewIntrospectNodeMap()
	return i
}

func (i *Introspect) Registry() common.IRegistry {
	return i.registry
}

func (i *Introspect) Inspect(any interface{}) (*model.Node, error) {
	if any == nil {
		return nil, logs.Error("Cannot introspect a nil value")
	}
	i.registry.RegisterStruct(any)
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

func (i *Introspect) Print() {
	i.pathToNode.Iterate(printDo)
}

func (i *Introspect) Kind(node *model.Node) reflect.Kind {
	t, err := i.registry.TypeByName(node.TypeName)
	if err != nil {
		panic(err.Error())
	}
	return t.Kind()
}

func (i *Introspect) Clone(any interface{}) interface{} {
	return i.cloner.Clone(any)
}

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

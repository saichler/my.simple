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
	pathToNode *NodeMap
	typeToNode *NodeMap
	registry   common.IRegistry
	cloner     *Cloner
	tableViews *maps.SyncMap
}

func NewIntrospect(registry common.IRegistry) *Introspect {
	i := &Introspect{}
	i.registry = registry
	i.cloner = newCloner()
	i.pathToNode = NewIntrospectNodeMap()
	i.typeToNode = NewIntrospectNodeMap()
	i.tableViews = maps.NewSyncMap()
	return i
}

func (i *Introspect) Registry() common.IRegistry {
	return i.registry
}

func (i *Introspect) Inspect(any interface{}) (*model.Node, error) {
	if any == nil {
		return nil, logs.Error("Cannot introspect a nil value")
	}

	_, t := common.ValueAndType(any)
	if t.Kind() == reflect.Slice && t.Kind() == reflect.Map {
		t = t.Elem().Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, logs.Error("Cannot introspect a value that is not a struct")
	}
	node, ok := i.pathToNode.Get(strings.ToLower(t.Name()))
	if ok {
		return node, nil
	}
	return i.inspectStruct(t, nil, ""), nil
}

func (i *Introspect) Node(path string) (*model.Node, bool) {
	return i.pathToNode.Get(strings.ToLower(path))
}

func (i *Introspect) NodeByValue(any interface{}) (*model.Node, bool) {
	val := reflect.ValueOf(any)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return i.NodeByType(val.Type())
}

func (i *Introspect) NodeByType(typ reflect.Type) (*model.Node, bool) {
	return i.NodeByTypeName(typ.Name())
}

func (i *Introspect) NodeByTypeName(name string) (*model.Node, bool) {
	return i.typeToNode.Get(name)
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

func (i *Introspect) addTableView(node *model.Node) {
	tv := &model.TableView{Table: node, Columns: make([]*model.Node, 0)}
	for _, attr := range node.Attributes {
		if common.IsLeaf(attr) {
			tv.Columns = append(tv.Columns, attr)
		}
	}
	i.tableViews.Put(node.TypeName, tv)
}

func (i *Introspect) TableView(name string) (*model.TableView, bool) {
	tv, ok := i.tableViews.Get(name)
	if !ok {
		return nil, ok
	}
	return tv.(*model.TableView), ok
}

func (i *Introspect) TableViews() []*model.TableView {
	list := i.tableViews.ValuesAsList(reflect.TypeOf(&model.TableView{}), nil)
	return list.([]*model.TableView)
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

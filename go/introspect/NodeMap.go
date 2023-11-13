package introspect

import (
	"github.com/saichler/my.simple/go/introspect/model"
	"github.com/saichler/my.simple/go/utils/maps"
	"reflect"
)

var node *model.Node
var nodeType = reflect.TypeOf(node)

type NodeMap struct {
	impl *maps.SyncMap
}

func NewIntrospectNodeMap() *NodeMap {
	m := &NodeMap{}
	m.impl = maps.NewSyncMap()
	return m
}

func (m *NodeMap) Put(key string, value *model.Node) bool {
	return m.impl.Put(key, value)
}

func (m *NodeMap) Get(key string) (*model.Node, bool) {
	value, ok := m.impl.Get(key)
	if value != nil {
		return value.(*model.Node), ok
	}
	return nil, ok
}

func (m *NodeMap) Contains(key string) bool {
	return m.impl.Contains(key)
}

func (m *NodeMap) NodesList(filter func(v interface{}) bool) []*model.Node {
	return m.impl.ValuesAsList(nodeType, filter).([]*model.Node)
}

func (m *NodeMap) Iterate(do func(k, v interface{})) {
	m.impl.Iterate(do)
}

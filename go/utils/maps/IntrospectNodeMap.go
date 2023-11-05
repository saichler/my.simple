package maps

import (
	"github.com/saichler/my.simple/go/introspect/model"
	"reflect"
)

var node *model.Node
var nodeType = reflect.TypeOf(node)

type IntrospectNodeMap struct {
	impl *SyncMap
}

func NewIntrospectNodeMap() *IntrospectNodeMap {
	m := &IntrospectNodeMap{}
	m.impl = NewSyncMap()
	return m
}

func (m *IntrospectNodeMap) Put(key string, value *model.Node) bool {
	return m.impl.Put(key, value)
}

func (m *IntrospectNodeMap) Get(key string) (*model.Node, bool) {
	value, ok := m.impl.Get(key)
	if value != nil {
		return value.(*model.Node), ok
	}
	return nil, ok
}

func (m *IntrospectNodeMap) Contains(key string) bool {
	return m.impl.Contains(key)
}

func (m *IntrospectNodeMap) NodesList(filter func(v interface{}) bool) []*model.Node {
	return m.impl.valuesAsList(nodeType, filter).([]*model.Node)
}

func (m *IntrospectNodeMap) Iterate(do func(k, v interface{})) {
	m.impl.Iterate(do)
}

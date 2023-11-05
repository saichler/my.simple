package maps

import (
	"github.com/saichler/my.simple/go/common"
	"reflect"
)

type PortMap struct {
	impl *SyncMap
}

var port common.Port
var portType = reflect.TypeOf(port)

func NewPortMap() *PortMap {
	m := &PortMap{}
	m.impl = NewSyncMap()
	return m
}

func (pm *PortMap) Put(key string, value common.Port) bool {
	return pm.impl.Put(key, value)
}

func (pm *PortMap) Get(key string) (common.Port, bool) {
	value, ok := pm.impl.Get(key)
	if value != nil {
		return value.(common.Port), ok
	}
	return nil, ok
}

func (pm *PortMap) Contains(key string) bool {
	return pm.impl.Contains(key)
}

func (pm *PortMap) PortList() []common.Port {
	return pm.impl.valuesAsList(portType, nil).([]common.Port)
}

func (pm *PortMap) Iterate(do func(k, v interface{})) {
	pm.impl.Iterate(do)
}

package service_point

import (
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/utils/maps"
)

type String2ServicePointMap struct {
	impl *maps.SyncMap
}

func NewString2ServicePointMap() *String2ServicePointMap {
	newMap := &String2ServicePointMap{}
	newMap.impl = maps.NewSyncMap()
	return newMap
}

func (mp *String2ServicePointMap) Put(key string, value common.ServicePointHandler) bool {
	return mp.impl.Put(key, value)
}

func (mp *String2ServicePointMap) Get(key string) (common.ServicePointHandler, bool) {
	value, ok := mp.impl.Get(key)
	if value != nil {
		return value.(common.ServicePointHandler), ok
	}
	return nil, ok
}

func (mp *String2ServicePointMap) Contains(key string) bool {
	return mp.impl.Contains(key)
}

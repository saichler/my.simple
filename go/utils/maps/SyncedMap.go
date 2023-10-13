package maps

import (
	"reflect"
	"sync"
)

type SyncMap struct {
	m map[interface{}]interface{}
	s *sync.RWMutex
}

func NewSyncMap() *SyncMap {
	mm := &SyncMap{}
	mm.m = make(map[interface{}]interface{})
	mm.s = &sync.RWMutex{}
	return mm
}

func (syncMap *SyncMap) Put(key, value interface{}) bool {
	syncMap.s.Lock()
	defer syncMap.s.Unlock()
	_, ok := syncMap.m[key]
	syncMap.m[key] = value
	return ok
}

func (syncMap *SyncMap) Get(key interface{}) (interface{}, bool) {
	syncMap.s.RLock()
	defer syncMap.s.RUnlock()
	v, ok := syncMap.m[key]
	return v, ok
}

func (syncMap *SyncMap) Contains(key interface{}) bool {
	syncMap.s.RLock()
	defer syncMap.s.RUnlock()
	_, ok := syncMap.m[key]
	return ok
}

func (syncMap *SyncMap) Delete(key interface{}) (interface{}, bool) {
	syncMap.s.Lock()
	defer syncMap.s.Unlock()
	v, ok := syncMap.m[key]
	delete(syncMap.m, key)
	return v, ok
}

func (syncMap *SyncMap) Size() int {
	syncMap.s.RLock()
	defer syncMap.s.RUnlock()
	return len(syncMap.m)
}

func (syncMap *SyncMap) Clean() map[interface{}]interface{} {
	syncMap.s.Lock()
	defer syncMap.s.Unlock()
	result := syncMap.m
	syncMap.m = make(map[interface{}]interface{})
	return result
}

func (syncMap *SyncMap) valuesAsList(typ reflect.Type) interface{} {
	syncMap.s.RLock()
	defer syncMap.s.RUnlock()
	newSlice := reflect.MakeSlice(typ, len(syncMap.m), len(syncMap.m))
	index := 0
	for _, v := range syncMap.m {
		newSlice.Index(index).Set(reflect.ValueOf(v))
		index++
	}
	return newSlice.Interface()
}

func (syncMap *SyncMap) Iterate(do func(k, v interface{})) {
	syncMap.s.RLock()
	defer syncMap.s.RUnlock()
	for k, v := range syncMap.m {
		do(k, v)
	}
}

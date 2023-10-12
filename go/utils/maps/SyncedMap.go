package maps

import "sync"

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

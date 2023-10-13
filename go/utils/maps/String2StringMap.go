package maps

type String2StringMap struct {
	impl *SyncMap
}

func NewString2StringMap() *String2StringMap {
	s2s := &String2StringMap{}
	s2s.impl = NewSyncMap()
	return s2s
}

func (s2s *String2StringMap) Put(key string, value string) bool {
	return s2s.impl.Put(key, value)
}

func (s2s *String2StringMap) Get(key string) (string, bool) {
	value, ok := s2s.impl.Get(key)
	if value != nil {
		return value.(string), ok
	}
	return "", ok
}

func (s2s *String2StringMap) Contains(key string) bool {
	return s2s.impl.Contains(key)
}

func (s2s *String2StringMap) Iterate(do func(k, v interface{})) {
	s2s.Iterate(do)
}

package maps

type String2BoolMap struct {
	impl *SyncMap
}

func NewString2BoolMap() *String2BoolMap {
	s2s := &String2BoolMap{}
	s2s.impl = NewSyncMap()
	return s2s
}

func (s2b *String2BoolMap) Put(key string, value bool) bool {
	return s2b.impl.Put(key, value)
}

func (s2b *String2BoolMap) Get(key string) (string, bool) {
	value, ok := s2b.impl.Get(key)
	if value != nil {
		return value.(string), ok
	}
	return "", ok
}

func (s2b *String2BoolMap) Contains(key string) bool {
	return s2b.impl.Contains(key)
}

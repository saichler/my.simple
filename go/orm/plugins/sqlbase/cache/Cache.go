package cache

import (
	"github.com/saichler/my.simple/go/utils/maps"
	"github.com/saichler/my.simple/go/utils/strng"
	"reflect"
)

type EntryType string

const (
	Insert EntryType = "I"
	Select EntryType = "S"
	Update EntryType = "U"
	Delete EntryType = "D"
	Table  EntryType = "T"
	Field  EntryType = "F"
)

type Cache struct {
	cache *maps.SyncMap
}

func NewCache() *Cache {
	return &Cache{cache: maps.NewSyncMap()}
}

func (c *Cache) TableName(name string) bool {
	return c.cache.Contains(keyOf(Table, name))
}

func (c *Cache) Tables() []string {
	ct := c.cache.KeysAsList(reflect.TypeOf(""), func(i interface{}) bool {
		key := i.(string)
		if key[0:2] == "T_" {
			return true
		}
		return false
	}).([]string)
	result := make([]string, len(ct))
	for i, v := range ct {
		result[i] = v[2:]
	}
	return result
}

func (c *Cache) AddTable(name string) {
	c.cache.Put(keyOf(Table, name), true)
}

func (c *Cache) PutIfNotExist(entryType EntryType, name string, v interface{}) {
	key := keyOf(entryType, name)
	if !c.cache.Contains(key) {
		c.cache.Put(key, v)
	}
}

func (c *Cache) Get(entryType EntryType, name string) (interface{}, bool) {
	return c.cache.Get(keyOf(entryType, name))
}

func keyOf(entryType EntryType, name string) string {
	return strng.New(entryType, "_", name).String()
}

package relational

import (
	"errors"
	"github.com/saichler/my.simple/go/common"
	"reflect"
)

func (data *RelationalData) addSliceRoot(value reflect.Value, inspect common.IIntrospect) error {
	var table *Table
	for i := 0; i < value.Len(); i++ {
		if table == nil {
			tName := tableName(value.Index(i))
			data.rootTableName = tName
			table = data.getOrCreateTable(tName)
		}
		err := table.addRoot(reflect.ValueOf(i), removePtr(value.Index(i)), inspect, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (data *RelationalData) addMapRoot(value reflect.Value, inspect common.IIntrospect) error {
	mapKeys := value.MapKeys()
	var table *Table
	for _, k := range mapKeys {
		if table == nil {
			tName := tableName(value.MapIndex(k))
			data.rootTableName = tName
			table = data.getOrCreateTable(tName)
		}
		err := table.addRoot(k, removePtr(value.MapIndex(k)), inspect, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (data *RelationalData) addValueRoot(value reflect.Value, inspect common.IIntrospect) error {
	tName := tableName(value)
	data.rootTableName = tName
	table := data.getOrCreateTable(tName)
	return table.addRoot(reflect.ValueOf(nil), removePtr(value), inspect, data)
}

func (data *RelationalData) AddInstances(any interface{}, inspect common.IIntrospect) error {
	value := reflect.ValueOf(any)

	if !value.IsValid() {
		return errors.New("added element is not valid")
	}

	if value.Kind() == reflect.Ptr && value.IsNil() {
		return errors.New("added element is nil")
	}

	if value.Kind() == reflect.Slice {
		return data.addSliceRoot(value, inspect)
	} else if value.Kind() == reflect.Map {
		return data.addMapRoot(value, inspect)
	} else if value.Kind() == reflect.Ptr && value.Elem().Kind() == reflect.Struct {
		return data.addValueRoot(value, inspect)
	}

	return errors.New("input element is not a slice of struct,map of struct or struct")
}

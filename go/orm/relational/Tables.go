package relational

import (
	"errors"
	"fmt"
	"github.com/saichler/my.simple/go/common"
	"reflect"
)

type Tables struct {
	trId       string
	name2Table map[string]*Table
}

func NewTables(trId string) *Tables {
	return &Tables{trId: trId, name2Table: make(map[string]*Table)}
}

func (tables *Tables) getOrCreateTable(typ string) *Table {
	table, ok := tables.name2Table[typ]
	if !ok {
		table = newTable()
		tables.name2Table[typ] = table
	}
	return table
}

func (tables *Tables) addSliceRoot(value reflect.Value, inspect common.IIntrospect) error {
	var table *Table
	for i := 0; i < value.Len(); i++ {
		if table == nil {
			table = tables.getOrCreateTable(tableName(value.Index(i)))
		}
		err := table.addRoot(reflect.ValueOf(i), removePtr(value.Index(i)), inspect, tables)
		if err != nil {
			return err
		}
	}
	return nil
}

func (tables *Tables) addMapRoot(value reflect.Value, inspect common.IIntrospect) error {
	mapKeys := value.MapKeys()
	var table *Table
	for _, k := range mapKeys {
		if table == nil {
			table = tables.getOrCreateTable(tableName(value.MapIndex(k)))
		}
		err := table.addRoot(k, removePtr(value.MapIndex(k)), inspect, tables)
		if err != nil {
			return err
		}
	}
	return nil
}

func (tables *Tables) addValueRoot(value reflect.Value, inspect common.IIntrospect) error {
	table := tables.getOrCreateTable(tableName(value))
	return table.addRoot(reflect.ValueOf(nil), removePtr(value), inspect, tables)
}

func (tables *Tables) Add(any interface{}, inspect common.IIntrospect) error {
	value := reflect.ValueOf(any)

	if !value.IsValid() {
		return errors.New("added element is not valid")
	}

	if value.Kind() == reflect.Ptr && value.IsNil() {
		return errors.New("added element is nil")
	}

	if value.Kind() == reflect.Slice {
		return tables.addSliceRoot(value, inspect)
	} else if value.Kind() == reflect.Map {
		return tables.addMapRoot(value, inspect)
	} else if value.Kind() == reflect.Ptr && value.Elem().Kind() == reflect.Struct {
		return tables.addValueRoot(value, inspect)
	}

	return errors.New("input element is not a slice of struct,map of struct or struct")
}

func (tables *Tables) Print() {
	for name, table := range tables.name2Table {
		fmt.Println(name)
		table.Print()
	}
}

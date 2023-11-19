package relational

import (
	"errors"
	"fmt"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/introspect/model"
	"github.com/saichler/my.simple/go/utils/strng"
	"reflect"
)

type Table struct {
	rows map[string]map[string]map[string]*Row
}

func newTable() *Table {
	return &Table{rows: make(map[string]map[string]map[string]*Row)}
}

func toString(value reflect.Value) string {
	s := strng.New()
	s.TypesPrefix = true
	return s.ToString(value)
}

func (table *Table) Print() {
	for k, _ := range table.rows {
		fmt.Println("   ", k)
	}
}

func (table *Table) addRow(key, value reflect.Value, node *model.Node, path, attr string, inspect common.IIntrospect, tables *RelationalData) error {
	recKey := keyOf(key, value, node, path, attr, inspect)
	row := newRow()
	row.addValues(value, recKey, node, inspect, tables)
	table.addRowToMap(path, attr, recKey, row)
	return nil
}

func (table *Table) addRowToMap(path, colName, recKey string, row *Row) {
	pathMap, ok := table.rows[path]
	if !ok {
		pathMap = make(map[string]map[string]*Row)
		table.rows[path] = pathMap
	}
	colNameMap, ok := pathMap[colName]
	if !ok {
		colNameMap = make(map[string]*Row)
		pathMap[colName] = colNameMap
	}
	colNameMap[recKey] = row
}

func (table *Table) addRoot(key, value reflect.Value, inspect common.IIntrospect, tables *RelationalData) error {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	rootNode, ok := inspect.Node(value.Type().Name())
	if !ok {
		return errors.New("Cannot find inspected data for " + value.Type().Name())
	}
	return table.addRow(key, value, rootNode, "", "", inspect, tables)
}

func (table *Table) add(value reflect.Value, node *model.Node, path, attr string, inspect common.IIntrospect, tables *RelationalData) {

	if !value.IsValid() {
		return
	}

	if value.Kind() == reflect.Ptr && value.IsNil() {
		return
	}

	if value.Kind() == reflect.Slice {
		for i := 0; i < value.Len(); i++ {
			table.addRow(reflect.ValueOf(i), removePtr(value.Index(i)), node, path, attr, inspect, tables)
		}
	} else if value.Kind() == reflect.Map {
		mapKeys := value.MapKeys()
		for _, k := range mapKeys {
			table.addRow(k, removePtr(value.MapIndex(k)), node, path, attr, inspect, tables)
		}
	} else if value.Kind() == reflect.Struct {
		table.addRow(reflect.ValueOf(nil), removePtr(value), node, path, attr, inspect, tables)
	}
}

func (table *Table) rowsOf(colName, key string) map[string]*Row {
	colNameMap, ok := table.rows[key]
	if ok {
		return colNameMap[colName]
	}
	return nil
}

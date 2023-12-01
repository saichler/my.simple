package relational

import (
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/introspect/model"
	"github.com/saichler/my.simple/go/utils/strng"
	"reflect"
)

type Row struct {
	colValues map[string]reflect.Value
}

func newRow() *Row {
	return &Row{colValues: make(map[string]reflect.Value)}
}

func (row *Row) ValueOf(name string) (reflect.Value, bool) {
	v, ok := row.colValues[name]
	return v, ok
}

func (row *Row) addValue(colName string, parent reflect.Value) {
	row.colValues[colName] = parent.FieldByName(colName)
}

func (row *Row) addRecValue(attr *model.Node, value interface{}) {
	if attr.IsSlice || attr.IsMap {
		value = strng.FromString(value.(string)).Interface()
	}
	row.colValues[attr.FieldName] = reflect.ValueOf(value)
}

func (row *Row) addValues(parent reflect.Value, path string, node *model.Node, inspect common.IIntrospect, relationalData *RelationalData) {
	for colName, n := range node.Attributes {
		if common.IsLeaf(n) {
			row.addValue(colName, parent)
		} else {
			row.colValues[strng.New("_", colName).String()] = reflect.ValueOf(n.TypeName)
			table := relationalData.getOrCreateTable(n.TypeName)
			table.add(removePtr(parent.FieldByName(colName)), n, path, n.FieldName, inspect, relationalData)
		}
	}
}

func FromRec(args []interface{}, view *model.TableView) *Row {
	row := newRow()
	row.colValues[common.RECKEY] = reflect.ValueOf(args[0].(string))
	for i, attr := range view.Columns {
		row.addRecValue(attr, args[i+1])
	}
	for _, subT := range view.SubTables {
		row.colValues[strng.New("_", subT.FieldName).String()] = reflect.ValueOf(subT.TypeName)
	}
	return row
}

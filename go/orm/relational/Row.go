package relational

import (
	"fmt"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/introspect/model"
	"reflect"
)

type Row struct {
	colValues map[string]reflect.Value
}

func (row *Row) addValue(colName string, parent reflect.Value) {
	row.colValues[colName] = parent.FieldByName(colName)
}

func (row *Row) addValues(parent reflect.Value, path string, node *model.Node, inspect common.IIntrospect, tables *Tables) {
	for colName, n := range node.Attributes {
		if common.IsLeaf(n) {
			fmt.Println("Leaf:", node.TypeName, ":", n.FieldName)
			row.addValue(colName, parent)
		} else {
			fmt.Println("None Leaf:", n.FieldName, ":", path)
			table := tables.getOrCreateTable(n.TypeName)
			table.add(removePtr(parent.FieldByName(colName)), n, path, n.FieldName, inspect, tables)
		}
	}
}

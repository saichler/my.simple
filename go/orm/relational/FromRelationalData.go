package relational

import (
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/utils/logs"
	"github.com/saichler/my.simple/go/utils/strng"
	"reflect"
	"strings"
)

func (data *RelationalData) ToIstances(inspect common.IIntrospect) (map[string]interface{}, error) {
	rootRows := make(map[string]*Row, 0)
	table := data.name2Table[data.rootTableName]
	for key, row := range table.rows[""][""] {
		rootRows[key] = row
	}
	return data.toIstances(data.rootTableName, rootRows, inspect)
}

func (data *RelationalData) toIstances(tName string, rows map[string]*Row, inspect common.IIntrospect) (map[string]interface{}, error) {
	instances := make(map[string]interface{})
	for key, row := range rows {
		instance, err := data.rowToInstance(tName, key, row, inspect)
		if err != nil {
			return instances, err
		}
		instances[key] = instance
	}
	return instances, nil
}

func (data *RelationalData) rowToInstance(tName, key string, row *Row, inspect common.IIntrospect) (interface{}, error) {
	rowInstance, err := inspect.Registry().NewInstance(tName)
	if err != nil {
		return nil, err
	}
	rowValue := reflect.ValueOf(rowInstance).Elem()
	for fieldName, value := range row.colValues {
		if fieldName[0] == '_' {
			rows := data.rowsOf(value.String(), fieldName[1:], key)
			if rows != nil {
				data.setValues(key, rowValue, rows, fieldName[1:], inspect)
			}
			continue
		}
		setValue(rowValue, value, fieldName)
	}
	return rowInstance, nil
}

func (data *RelationalData) rowsOf(tName, colName, key string) map[string]*Row {
	table, ok := data.name2Table[tName]
	if ok {
		return table.rowsOf(colName, key)
	}
	return nil
}

func setValue(parent, value reflect.Value, fieldName string) {
	parent.FieldByName(fieldName).Set(value)
}

func (data *RelationalData) setValues(key string, parent reflect.Value, rows map[string]*Row, fieldName string, inspect common.IIntrospect) {
	field := parent.FieldByName(fieldName)
	tName := ""
	var newSlice reflect.Value
	var newMap reflect.Value
	if field.Kind() == reflect.Slice {
		tName = field.Type().Elem().Elem().Name()
		newSlice = reflect.MakeSlice(field.Type(), len(rows), len(rows))
	} else if field.Kind() == reflect.Map {
		tName = field.Type().Elem().Elem().Name()
		newMap = reflect.MakeMap(field.Type())
	} else if field.Kind() == reflect.Ptr {
		tName = field.Type().Elem().Name()
	}

	for rowKey, row := range rows {
		instance, err := data.rowToInstance(tName, key, row, inspect)
		if err != nil {
			logs.Error("Failed From RelationalData", err)
			continue
		}
		if field.Kind() == reflect.Ptr {
			field.Set(reflect.ValueOf(instance))
		} else if field.Kind() == reflect.Slice {
			index := strng.FromString(extractKeyValue(rowKey)).Interface().(int)
			newSlice.Index(index).Set(reflect.ValueOf(instance))
		} else if field.Kind() == reflect.Map {
			mapKey := strng.FromString(extractKeyValue(rowKey))
			newMap.SetMapIndex(mapKey, reflect.ValueOf(instance))
		}
	}
	if field.Kind() == reflect.Slice {
		field.Set(newSlice)
	} else if field.Kind() == reflect.Map {
		field.Set(newMap)
	}
}

func extractKeyValue(key string) string {
	index1 := strings.LastIndex(key, "<")
	index2 := strings.LastIndex(key, ">")
	return key[index1+1 : index2]
}

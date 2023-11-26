package postgres

import (
	"errors"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/introspect/model"
	"github.com/saichler/my.simple/go/orm/relational"
	"github.com/saichler/my.simple/go/utils/strng"
	"reflect"
	"strings"
)

func (plugin *OrmPostgresPlugin) validateTables(rdata *relational.RelationalData, o common.IORM) error {
	tables := rdata.TablesMap()
	for tname, _ := range tables {
		err := plugin.validateOrCreateTable(tname, o.Introspect())
		if err != nil {
			return err
		}
	}
	return nil
}

func (plugin *OrmPostgresPlugin) validateOrCreateTable(tname string, inspect common.IIntrospect) error {
	node, ok := inspect.NodeByTypeName(tname)
	if !ok {
		return errors.New("Cannot find inspect data for: " + tname)
	}

	if plugin.names.Contains(tname) {
		return plugin.validateOrCreateFields(node)
	}

	sq := strng.New("select count(*) from ", tname).String()
	_, err := plugin.db.Exec(sq)
	if err != nil && strings.Contains(err.Error(), "relation") &&
		strings.Contains(err.Error(), "does not exist") {
		err = plugin.createTable(node, inspect)
		if err != nil {
			return err
		}
		plugin.names.Put(tname, true)
	}
	return nil
}

func (plugin *OrmPostgresPlugin) createTable(node *model.Node, inspect common.IIntrospect) error {
	//if we need to ignore this table and not persist it
	if inspect.DecoratorOf(model.DecoratorType_Ignore, node) != nil {
		return nil
	}
	ignoredAttr, _ := inspect.DecoratorOf(model.DecoratorType_IgnoreAttr, node).(map[string]bool)
	sq := strng.New("CREATE TABLE IF NOT EXISTS ")
	sq.Add(plugin.schema).Add(".").Add(node.TypeName).Add(" (\n")
	sq.Add("    ").Add(common.RECKEY).Add("    ").Add("VARCHAR,\n")
	for _, attr := range node.Attributes {
		//This attribute was marked as none persist, hence ignore it
		if ignoredAttr != nil && ignoredAttr[attr.FieldName] {
			continue
		} else if common.IsLeaf(attr) {
			k := inspect.Kind(attr)
			if attr.IsSlice || attr.IsMap {
				k = reflect.Slice
			}
			fldSql, err := fieldDef(attr.FieldName, ",\n", k)
			if err != nil {
				return err
			}
			sq.Add(fldSql)
		}
	}

	//The primary key is always RECKEY, with the value of the primary key decorator
	sq.Add("PRIMARY KEY (").Add(common.RECKEY).Add(")\n")

	//@TODO add primary, unique & none unique indexes from decorators

	sq.Add(");")
	sqlStr := sq.String()
	_, err := plugin.db.Exec(sqlStr)
	if err != nil {
		return errors.New(err.Error() + "\n" + sq.String())
	}
	return nil
}

func fieldDef(fieldName, delimiter string, kind reflect.Kind) (string, error) {
	sq := strng.New()
	//User is a saved word, hence add _ to it
	if fieldName == "User" {
		fieldName = "_" + fieldName
	}
	sq.Add("    ").Add(fieldName).Add("    ")

	switch kind {
	case reflect.Slice:
		fallthrough
	case reflect.Map:
		fallthrough
	case reflect.Float64:
		fallthrough
	case reflect.Float32:
		fallthrough
	case reflect.String:
		sq.Add("VARCHAR").Add(delimiter)
	case reflect.Int32:
		fallthrough
	case reflect.Uint32:
		sq.Add("integer DEFAULT 0").Add(delimiter)
	case reflect.Int:
		fallthrough
	case reflect.Uint64:
		fallthrough
	case reflect.Int64:
		sq.Add("bigint DEFAULT 0").Add(delimiter)
	case reflect.Bool:
		sq.Add("boolean DEFAULT FALSE").Add(delimiter)
	default:
		return "", errors.New("unsupported kind " + kind.String())
	}
	return sq.String(), nil
}

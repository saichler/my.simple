package stmt

import (
	"errors"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/introspect/model"
	"github.com/saichler/my.simple/go/utils/strng"
	"reflect"
)

func (sb *SqlStatementBuilder) CreateSchema() error {
	if sb.schema == "" {
		return nil
	}
	st := strng.New("CREATE SCHEMA IF NOT EXISTS ")
	st.Add(sb.schema).Add(";")
	_, err := sb.db.Exec(st.String())
	if err != nil {
		return errors.New(err.Error() + "\n" + st.String())
	}
	return nil
}

func (sb *SqlStatementBuilder) CreateTable() error {
	//if we need to ignore this table and not persist it
	if sb.o.Introspect().DecoratorOf(model.DecoratorType_Ignore, sb.node) != nil {
		return nil
	}
	ignoredAttr, _ := sb.o.Introspect().DecoratorOf(model.DecoratorType_IgnoreAttr, sb.node).(map[string]bool)
	sq := strng.New("CREATE TABLE IF NOT EXISTS ")
	sq.Add(sb.tableName()).Add(" (\n")
	sq.Add("    ").Add(common.RECKEY).Add("    ").Add("VARCHAR,\n")
	for _, attr := range sb.node.Attributes {
		//This attribute was marked as none persist, hence ignore it
		if ignoredAttr != nil && ignoredAttr[attr.FieldName] {
			continue
		} else if common.IsLeaf(attr) {
			k := sb.o.Introspect().Kind(attr)
			if attr.IsSlice || attr.IsMap {
				k = reflect.Slice
			}
			fldSql, err := FieldDef(attr.FieldName, ",\n", k)
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
	_, err := sb.db.Exec(sqlStr)
	if err != nil {
		return errors.New(err.Error() + "\n" + sq.String())
	}
	return nil
}

func FieldDef(fieldName, delimiter string, kind reflect.Kind) (string, error) {
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

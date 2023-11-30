package sqlbase

import (
	"database/sql"
	"errors"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/introspect/model"
	"github.com/saichler/my.simple/go/orm/plugins/sqlbase/cache"
	"github.com/saichler/my.simple/go/utils/strng"
	"reflect"
	"strings"
)

func CreateSchema(schema string, db *sql.DB, o common.IORM, c *cache.Cache) error {
	if schema == "" {
		return nil
	}
	st := strng.New("CREATE SCHEMA IF NOT EXISTS ")
	st.Add(schema).Add(";")
	_, err := db.Exec(st.String())
	if err != nil {
		return errors.New(err.Error() + "\n" + st.String())
	}
	return CreateSchemaTables(db, o, c)
}

func CreateSchemaTables(db *sql.DB, o common.IORM, c *cache.Cache) error {
	views := o.Introspect().TableViews()
	for _, view := range views {
		err := CheckSchemaTable(view, db, o, c)
		if err != nil {
			return err
		}
	}
	return nil
}

func CheckSchemaTable(view *model.TableView, db *sql.DB, o common.IORM, c *cache.Cache) error {
	if c.TableName(view.Table.TypeName) {
		return CheckFields(view)
	}

	sq := strng.New("select count(*) from ", view.Table.TypeName).String()

	_, err := db.Exec(sq)
	if err != nil && (strings.Contains(err.Error(), "relation") &&
		strings.Contains(err.Error(), "does not exist") ||
		strings.Contains(err.Error(), "no such table")) {
		return CreateSchemaTable(view, db, o, c)
	} else if err != nil {
		return err
	}

	return nil
}

// @TODO implement this method
func CheckFields(view *model.TableView) error {
	return nil
}

func CreateSchemaTable(view *model.TableView, db *sql.DB, o common.IORM, c *cache.Cache) error {
	//Was table already created
	if c.TableName(view.Table.TypeName) {
		return nil
	}

	//if we need to ignore this table and not persist it
	if o.Introspect().DecoratorOf(model.DecoratorType_Ignore, view.Table) != nil {
		return nil
	}

	ignoredAttr, _ := o.Introspect().DecoratorOf(model.DecoratorType_IgnoreAttr, view.Table).(map[string]bool)
	sq := strng.New("CREATE TABLE IF NOT EXISTS ")
	sq.Add(view.Table.TypeName).Add(" (\n")
	sq.Add("    ").Add(common.RECKEY).Add("    ").Add("VARCHAR,\n")
	for _, attr := range view.Columns {
		//This attribute was marked as none persist, hence ignore it
		if ignoredAttr != nil && ignoredAttr[attr.FieldName] {
			continue
		} else {
			k := o.Introspect().Kind(attr)
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
	_, err := db.Exec(sqlStr)
	if err != nil {
		return errors.New(err.Error() + "\n" + sq.String())
	}
	c.AddTable(view.Table.TypeName)
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
		return "", errors.New("unsupported kind " + kind.String() + " " + fieldName)
	}
	return sq.String(), nil
}

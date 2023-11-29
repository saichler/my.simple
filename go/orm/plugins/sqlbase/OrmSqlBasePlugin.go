package sqlbase

import (
	"database/sql"
	"errors"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/orm/stmt"
	"github.com/saichler/my.simple/go/utils/maps"
)

type OrmSqlBasePlugin struct {
	decorator common.DataStoreDecorator
	stmt      *stmt.SqlStatementBuilder
	names     *maps.String2BoolMap
}

func NewOrmSqlBasePlugin(decorator common.DataStoreDecorator) common.IOrmPlugin {
	plugin := &OrmSqlBasePlugin{}
	plugin.names = maps.NewString2BoolMap()
	plugin.decorator = decorator
	return plugin
}

func (plugin *OrmSqlBasePlugin) Init(o common.IORM, args ...interface{}) error {
	if args == nil || len(args) != 2 {
		return errors.New("Sql base plugin requires 2 arguments of *sql.DB & schema name")
	}
	db, ok := args[0].(*sql.DB)
	if !ok || db == nil {
		return errors.New("Sql base plugin requires first argument to be a valid *sql.DB instance")
	}
	schema, ok := args[1].(string)
	if !ok {
		return errors.New("Sql base plugin requires second argument to be a schema name")
	}
	return plugin.init(db, schema, o)
}

func (plugin *OrmSqlBasePlugin) init(db *sql.DB, schema string, o common.IORM) error {
	plugin.stmt, _ = stmt.NewSqlStatementBuilder(schema, "", o, db, plugin.names, plugin.decorator)
	return plugin.stmt.CreateSchema()
}

func (plugin *OrmSqlBasePlugin) RelationalData() bool {
	return true
}

func (plugin *OrmSqlBasePlugin) Decorator() common.DataStoreDecorator {
	return plugin.decorator
}

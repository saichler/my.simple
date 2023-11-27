package sqlbase

import (
	"database/sql"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/orm/stmt"
	"github.com/saichler/my.simple/go/utils/maps"
)

type OrmSqlBasePlugin struct {
	sqlDatabaseDecorator common.SqlDatabaseDecorator
	stmt                 *stmt.SqlStatementBuilder
	names                *maps.String2BoolMap
}

func NewOrmSqlBasePlugin(decorator common.SqlDatabaseDecorator) *OrmSqlBasePlugin {
	plugin := &OrmSqlBasePlugin{}
	plugin.names = maps.NewString2BoolMap()
	plugin.sqlDatabaseDecorator = decorator
	return plugin
}

func (plugin *OrmSqlBasePlugin) Init(db *sql.DB, schema string, o common.IORM) error {
	plugin.stmt, _ = stmt.NewSqlStatementBuilder(schema, "", o, db, plugin.names, plugin.sqlDatabaseDecorator)
	return plugin.stmt.CreateSchema()
}

func (plugin *OrmSqlBasePlugin) SQL() bool {
	return true
}

func (plugin *OrmSqlBasePlugin) Decorator() common.SqlDatabaseDecorator {
	return plugin.sqlDatabaseDecorator
}

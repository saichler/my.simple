package sqlite

import (
	"database/sql"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/orm/stmt"
	"github.com/saichler/my.simple/go/utils/maps"
)

type OrmSqlitePlugin struct {
	stmt  *stmt.SqlStatementBuilder
	names *maps.String2BoolMap
}

func NewOrmSqlitePlugin() *OrmSqlitePlugin {
	plugin := &OrmSqlitePlugin{}
	plugin.names = maps.NewString2BoolMap()
	return plugin
}

func (plugin *OrmSqlitePlugin) Init(db *sql.DB, schema string, o common.IORM) error {
	plugin.stmt, _ = stmt.NewSqlStatementBuilder(schema, "", o, db, plugin.names)
	return plugin.stmt.CreateSchema()
}

func (plugin *OrmSqlitePlugin) SQL() bool {
	return true
}

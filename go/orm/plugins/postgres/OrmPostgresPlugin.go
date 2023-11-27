package postgres

import (
	"database/sql"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/orm/stmt"
	"github.com/saichler/my.simple/go/utils/maps"
)

type OrmPostgresPlugin struct {
	stmt  *stmt.SqlStatementBuilder
	names *maps.String2BoolMap
}

func NewOrmPostgresPlugin() *OrmPostgresPlugin {
	plugin := &OrmPostgresPlugin{}
	plugin.names = maps.NewString2BoolMap()
	return plugin
}

func (plugin *OrmPostgresPlugin) Init(db *sql.DB, schema string, o common.IORM) error {
	plugin.stmt, _ = stmt.NewSqlStatementBuilder(schema, "", o, db, plugin.names)
	return plugin.stmt.CreateSchema()
}

func (plugin *OrmPostgresPlugin) SQL() bool {
	return true
}

package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/orm/plugins/sqlbase"
)

type OrmSqlitePluginDecorator struct {
}

func NewOrmSqlitePlugin() common.IOrmPlugin {
	sDecorator := &OrmSqlitePluginDecorator{}
	plugin := sqlbase.NewOrmSqlBasePlugin(sDecorator)
	return plugin
}

func (plugin *OrmSqlitePluginDecorator) DataStoreTypeName() string {
	return "SQLite"
}

func (plugin *OrmSqlitePluginDecorator) Connect(args ...string) interface{} {
	db, err := sql.Open("sqlite3", args[0])
	if err != nil {
		panic(err)
	}
	return db
}

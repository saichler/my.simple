package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/saichler/my.simple/go/orm/plugins/sqlbase"
)

type OrmSqlitePluginDecorator struct {
}

func NewOrmSqlitePlugin() *sqlbase.OrmSqlBasePlugin {
	sDecorator := &OrmSqlitePluginDecorator{}
	plugin := sqlbase.NewOrmSqlBasePlugin(sDecorator)
	return plugin
}

func (plugin *OrmSqlitePluginDecorator) DbType() string {
	return "SQLite"
}

func (plugin *OrmSqlitePluginDecorator) Connect(args ...string) *sql.DB {
	db, err := sql.Open("sqlite3", args[0])
	if err != nil {
		panic(err)
	}
	return db
}

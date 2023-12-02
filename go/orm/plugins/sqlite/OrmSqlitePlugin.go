package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/orm/plugins/sqlbase"
	"strings"
)

type OrmSqlitePluginDecorator struct {
}

func NewOrmSqlitePlugin() common.IOrmPlugin {
	sDecorator := &OrmSqlitePluginDecorator{}
	plugin := sqlbase.NewOrmSqlBasePlugin(sDecorator)
	return plugin
}

func (decorator *OrmSqlitePluginDecorator) DataStoreTypeName() string {
	return "SQLite"
}

func (decorator *OrmSqlitePluginDecorator) Connect(args ...string) interface{} {
	db, err := sql.Open("sqlite3", args[0])
	if err != nil {
		panic(err)
	}
	return db
}

func (decorator *OrmSqlitePluginDecorator) DoesNotExistError(err error) bool {
	if err != nil && strings.Contains(err.Error(), "no such table") {
		return true
	}
	return false
}

package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/orm/plugins/sqlbase"
	"github.com/saichler/my.simple/go/utils/strng"
)

type OrmPostgresPluginDecorator struct {
}

func NewOrmPostgresPlugin() common.IOrmPlugin {
	pDecorator := &OrmPostgresPluginDecorator{}
	plugin := sqlbase.NewOrmSqlBasePlugin(pDecorator)
	return plugin
}

func (decorator *OrmPostgresPluginDecorator) DataStoreTypeName() string {
	return "Postgres"
}

func (decorator *OrmPostgresPluginDecorator) Connect(args ...string) interface{} {
	def := strng.New("host=", args[0], " port=", args[1], " user=", args[2], " password=", args[3], " dbname=", args[4], " sslmode=", args[5])
	db, err := sql.Open("postgres", def.String())
	if err != nil {
		panic(err)
	}
	return db
}

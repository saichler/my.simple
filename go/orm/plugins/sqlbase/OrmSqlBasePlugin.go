package sqlbase

import (
	"database/sql"
	"errors"
	"github.com/saichler/my.simple/go/common"
	cache2 "github.com/saichler/my.simple/go/orm/plugins/sqlbase/cache"
	"github.com/saichler/my.simple/go/utils/strng"
)

type OrmSqlBasePlugin struct {
	decorator common.DataStoreDecorator
	schema    string
	db        *sql.DB
	o         common.IORM
	cache     *cache2.Cache
}

func NewOrmSqlBasePlugin(decorator common.DataStoreDecorator) common.IOrmPlugin {
	plugin := &OrmSqlBasePlugin{}
	plugin.decorator = decorator
	plugin.cache = cache2.NewCache()
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
	plugin.schema = schema
	plugin.o = o
	plugin.db = db
	return plugin.init()
}

func (plugin *OrmSqlBasePlugin) init() error {
	return CreateSchema(plugin.schema, plugin.db, plugin.o, plugin.cache, plugin.decorator)
}

func (plugin *OrmSqlBasePlugin) RelationalData() bool {
	return true
}

func (plugin *OrmSqlBasePlugin) Decorator() common.DataStoreDecorator {
	return plugin.decorator
}

func TableName(schema, name string) string {
	if schema == "" {
		return name
	}
	return strng.New(schema, ".", name).String()
}

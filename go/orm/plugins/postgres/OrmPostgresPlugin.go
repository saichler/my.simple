package postgres

import (
	"database/sql"
	"errors"
	"github.com/saichler/my.simple/go/utils/maps"
	"github.com/saichler/my.simple/go/utils/strng"
)

type OrmPostgresPlugin struct {
	db     *sql.DB
	schema string
	names  *maps.String2BoolMap
}

func NewOrmPostgresPlugin(db *sql.DB, schema string) (*OrmPostgresPlugin, error) {
	plugin := &OrmPostgresPlugin{db: db, schema: schema}
	plugin.names = maps.NewString2BoolMap()
	err := plugin.createSchema()
	return plugin, err
}

func (plugin *OrmPostgresPlugin) SQL() bool {
	return true
}

func (plugin *OrmPostgresPlugin) createSchema() error {
	st := strng.New("CREATE SCHEMA IF NOT EXISTS ")
	st.Add(plugin.schema).Add(";")
	_, err := plugin.db.Exec(st.String())
	if err != nil {
		return errors.New(err.Error() + "\n" + st.String())
	}
	return nil
}

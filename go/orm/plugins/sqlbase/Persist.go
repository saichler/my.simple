package sqlbase

import (
	"errors"
	"github.com/saichler/my.simple/go/orm/plugins/sqlbase/cache"
	"github.com/saichler/my.simple/go/orm/plugins/sqlbase/stmt"
	"github.com/saichler/my.simple/go/orm/relational"
	"github.com/saichler/my.simple/go/utils/logs"
)

func (plugin *OrmSqlBasePlugin) Persist(data interface{}) error {
	rdata := data.(*relational.RelationalData)
	err := plugin.prepareStmt(rdata)
	if err != nil {
		return err
	}
	err = plugin.write(rdata)
	return err
}

func (plugin *OrmSqlBasePlugin) write(rdata *relational.RelationalData) error {
	tables := rdata.TablesMap()
	tx, _ := plugin.db.Begin()
	defer tx.Rollback()
	for tname, table := range tables {
		rows := table.Rows()
		isb, ok := plugin.cache.Get(cache.Insert, tname)
		if !ok {
			return errors.New("No insert statement was found for " + tname)
		}
		sb := isb.(*stmt.StmtBuilder)

		for rk, row := range rows {
			err := sb.Insert(rk, row, tx, plugin.o, plugin.cache)
			if err != nil {
				return logs.Error("Failed to insert record:", err)
			}
		}
	}
	return tx.Commit()
}

func (plugin *OrmSqlBasePlugin) prepareStmt(rdata *relational.RelationalData) error {
	tables := rdata.TablesMap()
	for tableName, _ := range tables {
		view, ok := plugin.o.Introspect().TableView(tableName)
		if !ok {
			return errors.New("Cannot find introspect view data for: " + tableName)
		}
		err := CheckSchemaTable(view, plugin.db, plugin.o, plugin.cache, plugin.decorator)
		if err != nil {
			return err
		}
		plugin.cache.PutIfNotExist(cache.Insert, view.Table.TypeName, stmt.NewStmtBuilder(plugin.schema, view))
	}
	return nil
}

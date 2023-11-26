package postgres

import (
	"fmt"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/orm/relational"
	"github.com/saichler/my.simple/go/orm/stmt"
)

func (plugin *OrmPostgresPlugin) Write(data interface{}, o common.IORM) error {
	rdata := data.(*relational.RelationalData)
	err := plugin.validateTables(rdata, o)
	plugin.write(rdata, o)
	return err
}

func (plugin *OrmPostgresPlugin) write(rdata *relational.RelationalData, o common.IORM) error {
	tables := rdata.TablesMap()
	tx, _ := plugin.db.Begin()
	defer tx.Rollback()
	for tname, table := range tables {
		rows := table.Rows()
		sb := &stmt.SqlStatementBuilder{}
		node, _ := o.Introspect().NodeByTypeName(tname)

		for rk, row := range rows {
			err := sb.Insert(plugin.schema, rk, node, o.Introspect(), row, tx)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	tx.Commit()
	return nil
}

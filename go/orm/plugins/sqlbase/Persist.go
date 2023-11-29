package sqlbase

import (
	"errors"
	"fmt"
	"github.com/saichler/my.simple/go/orm/relational"
	"github.com/saichler/my.simple/go/orm/stmt"
)

func (plugin *OrmSqlBasePlugin) Persist(data interface{}) error {
	rdata := data.(*relational.RelationalData)
	err := plugin.stmt.ValidateTables(rdata, plugin.names)
	if err != nil {
		return err
	}
	err = plugin.write(rdata)
	return err
}

func (plugin *OrmSqlBasePlugin) write(rdata *relational.RelationalData) error {
	tables := rdata.TablesMap()
	tx, _ := plugin.stmt.Db().Begin()
	defer tx.Rollback()
	for tname, table := range tables {
		rows := table.Rows()
		sb, ok := plugin.stmt.BuilderOf(stmt.BInsert, tname)
		if !ok {
			return errors.New("No Statement Builder found for " + tname)
		}

		for rk, row := range rows {
			err := sb.Insert(rk, row, tx)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	tx.Commit()
	return nil
}

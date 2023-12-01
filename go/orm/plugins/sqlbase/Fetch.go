package sqlbase

import (
	"fmt"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/orm/plugins/sqlbase/stmt"
	"github.com/saichler/my.simple/go/orm/relational"
)

func (plugin *OrmSqlBasePlugin) Fetch(fetch common.IFetch) (interface{}, error) {
	tx, err := plugin.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	rdata := relational.NewRelationalData("")
	for _, tname := range plugin.cache.Tables() {
		view, _ := plugin.o.Introspect().TableView(tname)
		sb := stmt.NewStmtBuilder(plugin.schema, view)
		fmt.Println(tname)
		recs, err := sb.Fetch(fetch, tx, plugin.o)
		if err != nil {
			return nil, err
		}
		rdata.ImportData(view, recs)
	}
	return rdata, nil
}

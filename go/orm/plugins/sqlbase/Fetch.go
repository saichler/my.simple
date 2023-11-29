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
		node, _ := plugin.o.Introspect().NodeByTypeName(tname)
		sb := stmt.NewStmtBuilder(plugin.schema, node)
		fmt.Println(tname)
		sb.Fetch(fetch, tx, plugin.o, plugin.cache)
	}
	return rdata, nil
}

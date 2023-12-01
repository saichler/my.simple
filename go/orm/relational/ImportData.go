package relational

import (
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/introspect/model"
	"strings"
)

func (data *RelationalData) ImportData(view *model.TableView, recs map[string][]interface{}) {
	table := newTable()
	data.name2Table[view.Table.TypeName] = table
	for key, rec := range recs {
		path := ""
		typ := ""
		lastIndex := strings.LastIndex(key, ".")
		if lastIndex != -1 {
			path = key[0:lastIndex]
			typ = common.NodeKey(key[lastIndex+1:])
		}
		if table.rows[path] == nil {
			table.rows[path] = make(map[string]map[string]*Row)
		}
		if table.rows[path][typ] == nil {
			table.rows[path][typ] = make(map[string]*Row)
		}
		table.rows[path][typ][key] = FromRec(rec, view)
	}
}

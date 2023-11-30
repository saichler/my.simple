package stmt

import (
	"database/sql"
	"fmt"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/orm/plugins/sqlbase/cache"
	"github.com/saichler/my.simple/go/utils/strng"
	"reflect"
)

func (sb *StmtBuilder) Fetch(fetch common.IFetch, tx *sql.Tx, o common.IORM, c *cache.Cache) error {
	if sb.stmt == nil {
		err := sb.createSelectStatement(tx)
		if err != nil {
			return err
		}
	}

	rows, err := sb.stmt.Query()
	args := make([]interface{}, len(sb.view.Columns)+1)
	rk := ""
	args[0] = &rk
	for i, attr := range sb.view.Columns {
		typ := reflect.TypeOf("")
		if !attr.IsSlice && !attr.IsMap {
			t, e := o.Introspect().Registry().TypeByName(attr.TypeName)
			if e != nil {
				panic(e.Error())
			}
			typ = t
		}
		args[i+1] = reflect.New(typ).Interface()
	}
	for rows.Next() {
		err := rows.Scan(args...)
		if err == nil {
			for _, arg := range args {
				val := reflect.ValueOf(arg).Elem().Interface()
				fmt.Print(val, " ")
			}
			fmt.Println()
		}
		fmt.Println(err)
	}
	return err
}

func (sb *StmtBuilder) createSelectStatement(tx *sql.Tx) error {

	selectSql := strng.New("Select ", common.RECKEY)

	for _, attr := range sb.view.Columns {
		selectSql.Add(",").Add(attr.FieldName)
	}
	selectSql.Add(" from ").Add(sb.view.Table.TypeName).Add(";")

	sb.stmtString = selectSql.String()

	stmt, err := tx.Prepare(sb.stmtString)
	if err != nil {
		return err
	}
	sb.stmt = stmt
	return err
}

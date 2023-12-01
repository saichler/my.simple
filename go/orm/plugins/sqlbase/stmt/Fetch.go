package stmt

import (
	"database/sql"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/utils/strng"
	"reflect"
)

func (sb *StmtBuilder) Fetch(fetch common.IFetch, tx *sql.Tx, o common.IORM) (map[string][]interface{}, error) {
	if sb.stmt == nil {
		err := sb.createSelectStatement(tx)
		if err != nil {
			return nil, err
		}
	}

	rows, err := sb.stmt.Query()
	recs := make(map[string][]interface{})

	for rows.Next() {
		args := sb.newArgs(o)
		vals := make([]interface{}, len(args))
		err = rows.Scan(args...)
		if err != nil {
			return nil, err
		}
		if err == nil {
			for i, arg := range args {
				vals[i] = reflect.ValueOf(arg).Elem().Interface()
			}
			recs[vals[0].(string)] = vals
		} else {
			return nil, err
		}
	}
	return recs, err
}

func (sb *StmtBuilder) newArgs(o common.IORM) []interface{} {
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
	return args
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

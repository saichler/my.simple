package stmt

import (
	"database/sql"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/orm/plugins/sqlbase/cache"
	"github.com/saichler/my.simple/go/orm/relational"
	"github.com/saichler/my.simple/go/utils/strng"
	"reflect"
	"strconv"
)

func (sb *StmtBuilder) Insert(rk string, row *relational.Row, tx *sql.Tx, o common.IORM, c *cache.Cache) error {
	if sb.stmt == nil {
		err := sb.createInsertStatement(tx, o, c)
		if err != nil {
			return err
		}
	}

	args := make([]interface{}, len(sb.view.Columns)+1)
	toString := strng.New()
	toString.TypesPrefix = true
	args[0] = rk
	for i, attr := range sb.view.Columns {
		val, ok := row.ValueOf(attr.FieldName)
		if ok {
			if val.Kind() == reflect.Map || val.Kind() == reflect.Slice {
				val = reflect.ValueOf(toString.ToString(val))
			}
			args[i+1] = val.Interface()
		} else {
			panic("unsupported insert type (yet) for " + attr.FieldName + " in " + sb.view.Table.TypeName)
		}
	}
	_, err := sb.stmt.Exec(args...)
	return err
}

func (sb *StmtBuilder) createInsertStatement(tx *sql.Tx, o common.IORM, c *cache.Cache) error {

	insertSql := strng.New("insert into ", sb.TableName(), " ")
	attrs := strng.New(" (", common.RECKEY, ",")
	values := strng.New(" values ($1,")

	first := true
	for i, attr := range sb.view.Columns {
		if !first {
			attrs.Add(",")
			values.Add(",")
		}
		first = false
		attrs.Add(attr.FieldName)
		values.Add("$")
		values.Add(strconv.Itoa(i + 2))
	}

	attrs.Add(") ")
	values.Add(")")
	insertSql.Join(attrs)
	insertSql.Join(values)
	onConflict := sb.createOnConflict()
	insertSql.Add(onConflict)
	sb.stmtString = insertSql.String()

	stmt, err := tx.Prepare(sb.stmtString)
	if err != nil {
		return err
	}
	sb.stmt = stmt
	return err
}

func (sb *StmtBuilder) createOnConflict() string {
	conflict := strng.New(" on conflict (").Add(common.RECKEY).Add(") do update set ")
	firstAttr := true
	for i, attr := range sb.view.Columns {
		if !firstAttr {
			conflict.Add(",")
		}
		firstAttr = false
		conflict.Add(attr.FieldName).Add("=").Add("$").Add(strconv.Itoa(i + 2))
	}
	return conflict.String()
}

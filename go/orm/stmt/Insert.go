package stmt

import (
	"database/sql"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/orm/relational"
	"github.com/saichler/my.simple/go/utils/strng"
	"reflect"
	"strconv"
)

func (sb *SqlStatementBuilder) createInsertStatement(tx *sql.Tx) error {

	sqls := strng.New("insert into ", sb.tableName(), " ")

	attrs := strng.New(" (", common.RECKEY, ",")
	values := strng.New(" values ($1,")
	first := true
	sb.attrNames = sb.o.Introspect().AttributesNames(sb.node)
	for i, attr := range sb.attrNames {
		if !first {
			attrs.Add(",")
			values.Add(",")
		}
		first = false
		attrs.Add(attr)
		values.Add("$")
		values.Add(strconv.Itoa(i + 2))
	}
	attrs.Add(") ")
	values.Add(")")
	sqls.Join(attrs)
	sqls.Join(values)
	onConflict := sb.conflict()
	sqls.Add(onConflict)
	sb.stmtString = sqls.String()
	stmt, err := tx.Prepare(sb.stmtString)
	if err != nil {
		return err
	}
	sb.stmt = stmt
	return err
}

func (sb *SqlStatementBuilder) Insert(rk string, row *relational.Row, tx *sql.Tx) error {
	if sb.stmt == nil {
		err := sb.createInsertStatement(tx)
		if err != nil {
			return err
		}
	}
	args := make([]interface{}, len(sb.attrNames)+1)
	toString := strng.New()
	toString.TypesPrefix = true
	args[0] = rk
	for i, key := range sb.attrNames {
		val, ok := row.ValueOf(key)
		if ok {
			if val.Kind() == reflect.Map || val.Kind() == reflect.Slice {
				val = reflect.ValueOf(toString.ToString(val))
			}
			args[i+1] = val.Interface()
		} else {
			panic("unsupported insert type (yet) for " + key + " in " + sb.node.TypeName)
		}
	}
	_, err := sb.stmt.Exec(args...)
	return err
}

func (sb *SqlStatementBuilder) conflict() string {
	conflict := strng.New(" on conflict (").Add(common.RECKEY).Add(") do update set ")
	firstAttr := true
	for i, key := range sb.attrNames {
		if !firstAttr {
			conflict.Add(",")
		}
		firstAttr = false
		conflict.Add(key).Add("=").Add("$").Add(strconv.Itoa(i + 2))
	}
	return conflict.String()
}

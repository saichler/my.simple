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
	attrNames := c.AttrNames(sb.node.TypeName)
	args := make([]interface{}, len(attrNames)+1)
	toString := strng.New()
	toString.TypesPrefix = true
	args[0] = rk
	for i, key := range attrNames {
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

func (sb *StmtBuilder) createInsertStatement(tx *sql.Tx, o common.IORM, c *cache.Cache) error {

	insertSql := strng.New("insert into ", sb.TableName(), " ")
	attrs := strng.New(" (", common.RECKEY, ",")
	values := strng.New(" values ($1,")
	attrNames := c.AttrNames(sb.node.TypeName)
	if attrNames == nil {
		attrNames = o.Introspect().AttributesNames(sb.node)
		c.PutAttrNames(sb.node.TypeName, attrNames)
	}

	first := true
	for i, attr := range attrNames {
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
	insertSql.Join(attrs)
	insertSql.Join(values)
	onConflict := sb.createOnConflict(attrNames)
	insertSql.Add(onConflict)
	sb.stmtString = insertSql.String()

	stmt, err := tx.Prepare(sb.stmtString)
	if err != nil {
		return err
	}
	sb.stmt = stmt
	return err
}

func (sb *StmtBuilder) createOnConflict(attrNames []string) string {
	conflict := strng.New(" on conflict (").Add(common.RECKEY).Add(") do update set ")
	firstAttr := true
	for i, key := range attrNames {
		if !firstAttr {
			conflict.Add(",")
		}
		firstAttr = false
		conflict.Add(key).Add("=").Add("$").Add(strconv.Itoa(i + 2))
	}
	return conflict.String()
}

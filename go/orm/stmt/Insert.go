package stmt

import (
	"database/sql"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/introspect/model"
	"github.com/saichler/my.simple/go/orm/relational"
	"github.com/saichler/my.simple/go/utils/strng"
	"reflect"
	"strconv"
)

func (st *SqlStatementBuilder) createInsertStatement(schema string, node *model.Node, inspect common.IIntrospect, tx *sql.Tx) error {

	//conflict := st.conflict(node)

	sqls := strng.New("insert into ", schema, ".", node.TypeName, " ")

	attrs := strng.New(" (", common.RECKEY, ",")
	values := strng.New(" values ($1,")
	first := true
	st.attrNames = inspect.AttributesNames(node)
	for i, attr := range st.attrNames {
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
	//sqls.Add(conflict)
	st.stmtString = sqls.String()
	stmt, err := tx.Prepare(st.stmtString)
	if err != nil {
		return err
	}
	st.stmt = stmt
	return err
}

func (st *SqlStatementBuilder) Insert(schema, rk string, node *model.Node, inspect common.IIntrospect, row *relational.Row, tx *sql.Tx) error {
	if st.stmt == nil {
		err := st.createInsertStatement(schema, node, inspect, tx)
		if err != nil {
			return err
		}
	}
	args := make([]interface{}, len(st.attrNames)+1)
	toString := strng.New()
	toString.TypesPrefix = true
	args[0] = rk
	for i, key := range st.attrNames {
		val, ok := row.ValueOf(key)
		if ok {
			if val.Kind() == reflect.Map || val.Kind() == reflect.Slice {
				val = reflect.ValueOf(toString.ToString(val))
			}
			args[i+1] = val.Interface()
		} else {
			panic("unsupported insert type (yet) for " + key + " in " + node.TypeName)
		}
	}
	_, err := st.stmt.Exec(args...)
	return err
}

/*
func (st *SqlStatementBuilder) conflict(node *model.Node) string {

	conflict := utils.NewString("")
	indexFields := make(map[string]string)
	conflict.Add(" on conflict (")

	primary := introspect.Primary(node)

	if primary != nil {
		firstAttr := true
		for _, attr := range primary.Attributes {
			if !firstAttr {
				conflict.Add(",")
			}
			firstAttr = false
			conflict.Add(attr)
			indexFields[attr] = attr
		}
	} else {
		conflict.Add(table_model.REC_KEY)
	}
	conflict.Add(") do update set ")
	firstAttr := true
	for i, key := range st.keys {
		if !firstAttr {
			conflict.Add(",")
		}
		firstAttr = false
		conflict.Add(key).Add("=").Add("$").Add(strconv.Itoa(i + 1))
	}
	return conflict.String()
}*/

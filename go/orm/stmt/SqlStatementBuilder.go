package stmt

import (
	"database/sql"
	"errors"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/introspect/model"
	"github.com/saichler/my.simple/go/utils/maps"
	"github.com/saichler/my.simple/go/utils/strng"
)

type SqlStatementBuilder struct {
	o      common.IORM
	node   *model.Node
	schema string
	db     *sql.DB

	stmt       *sql.Stmt
	stmtString string
	attrNames  []string
	nameCache  *maps.String2BoolMap
	children   *maps.SyncMap
}

func NewSqlStatementBuilder(schema, tableName string, o common.IORM, db *sql.DB, namecache *maps.String2BoolMap) (*SqlStatementBuilder, error) {
	sb := &SqlStatementBuilder{schema: schema, o: o, nameCache: namecache, db: db}
	if tableName != "" {
		node, ok := sb.o.Introspect().NodeByTypeName(tableName)
		if !ok {
			return nil, errors.New("Cannot find introspect data for: " + node.TypeName)
		}
		sb.node = node
	}
	return sb, nil
}

func (sb *SqlStatementBuilder) Db() *sql.DB {
	return sb.db
}

func (sb *SqlStatementBuilder) tableName() string {
	if sb.schema == "" {
		return sb.node.TypeName
	}
	return strng.New(sb.schema, ".", sb.node.TypeName).String()
}

func (sb *SqlStatementBuilder) BuilderOf(tname string) (*SqlStatementBuilder, bool) {
	bldr, ok := sb.children.Get(tname)
	if !ok {
		return nil, ok
	}
	return bldr.(*SqlStatementBuilder), ok
}

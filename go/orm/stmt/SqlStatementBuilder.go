package stmt

import (
	"database/sql"
	"errors"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/introspect/model"
	"github.com/saichler/my.simple/go/utils/maps"
	"github.com/saichler/my.simple/go/utils/strng"
)

type BuilderType int

const (
	BInsert BuilderType = 1
	BFetch  BuilderType = 2
)

type SqlStatementBuilder struct {
	o         common.IORM
	node      *model.Node
	schema    string
	db        *sql.DB
	decorator common.DataStoreDecorator

	stmt       *sql.Stmt
	stmtString string
	attrNames  []string
	nameCache  *maps.String2BoolMap
	builders   map[BuilderType]*maps.SyncMap
}

func NewSqlStatementBuilder(schema, tableName string, o common.IORM, db *sql.DB, namecache *maps.String2BoolMap, decorator common.DataStoreDecorator) (*SqlStatementBuilder, error) {
	sb := &SqlStatementBuilder{schema: schema, o: o, nameCache: namecache, db: db, decorator: decorator}
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

func (sb *SqlStatementBuilder) BuilderOf(bt BuilderType, tname string) (*SqlStatementBuilder, bool) {
	bldr, ok := sb.builders[bt].Get(tname)
	if !ok {
		return nil, ok
	}
	return bldr.(*SqlStatementBuilder), ok
}

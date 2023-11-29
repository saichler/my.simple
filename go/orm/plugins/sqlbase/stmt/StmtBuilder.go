package stmt

import (
	"database/sql"
	"github.com/saichler/my.simple/go/introspect/model"
	"github.com/saichler/my.simple/go/utils/strng"
)

type StmtBuilder struct {
	stmt       *sql.Stmt
	stmtString string
	node       *model.Node
	schema     string
}

func NewStmtBuilder(schema string, node *model.Node) *StmtBuilder {
	return &StmtBuilder{schema: schema, node: node}
}

func (sb *StmtBuilder) TableName() string {
	if sb.schema == "" {
		return sb.node.TypeName
	} else {
		return strng.New(sb.schema, ".", sb.node.TypeName).String()
	}
}

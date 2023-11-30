package stmt

import (
	"database/sql"
	"github.com/saichler/my.simple/go/introspect/model"
	"github.com/saichler/my.simple/go/utils/strng"
)

type StmtBuilder struct {
	stmt       *sql.Stmt
	stmtString string
	view       *model.TableView
	schema     string
}

func NewStmtBuilder(schema string, view *model.TableView) *StmtBuilder {
	return &StmtBuilder{schema: schema, view: view}
}

func (sb *StmtBuilder) TableName() string {
	if sb.schema == "" {
		return sb.view.Table.TypeName
	} else {
		return strng.New(sb.schema, ".", sb.view.Table.TypeName).String()
	}
}

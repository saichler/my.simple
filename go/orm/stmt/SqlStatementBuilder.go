package stmt

import "database/sql"

type SqlStatementBuilder struct {
	stmt       *sql.Stmt
	stmtString string
	attrNames  []string
}

package stmt

import (
	"database/sql"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/utils/strng"
)

func (sb *SqlStatementBuilder) Fetch(fetch common.IFetch, tx *sql.Tx) error {
	if sb.stmt == nil {
		err := sb.createSelectStatement(tx)
		if err != nil {
			return err
		}
	}
	rows, err := sb.stmt.Query()
	return err
}

func (sb *SqlStatementBuilder) createSelectStatement(tx *sql.Tx) error {

	selectSql := strng.New("Select ", common.RECKEY)
	sb.attrNames = sb.o.Introspect().AttributesNames(sb.node)
	for _, attr := range sb.attrNames {
		selectSql.Add(",").Add(attr)
	}
	selectSql.Add(" from ").Add(sb.node.TypeName).Add(";")

	sb.stmtString = selectSql.String()

	stmt, err := tx.Prepare(sb.stmtString)
	if err != nil {
		return err
	}
	sb.stmt = stmt
	return err
}

package stmt

import (
	"database/sql"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/orm/plugins/sqlbase/cache"
	"github.com/saichler/my.simple/go/utils/strng"
)

func (sb *StmtBuilder) Fetch(fetch common.IFetch, tx *sql.Tx, o common.IORM, c *cache.Cache) error {
	if sb.stmt == nil {
		err := sb.createSelectStatement(tx, o, c)
		if err != nil {
			return err
		}
	}
	_, err := sb.stmt.Query()
	return err
}

func (sb *StmtBuilder) createSelectStatement(tx *sql.Tx, o common.IORM, c *cache.Cache) error {

	selectSql := strng.New("Select ", common.RECKEY)
	attrNames := c.AttrNames(sb.node.TypeName)
	if attrNames == nil {
		attrNames = o.Introspect().AttributesNames(sb.node)
		c.PutAttrNames(sb.node.TypeName, attrNames)
	}

	for _, attr := range attrNames {
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

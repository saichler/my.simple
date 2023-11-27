package stmt

import (
	"github.com/saichler/my.simple/go/orm/relational"
	"github.com/saichler/my.simple/go/utils/maps"
	"github.com/saichler/my.simple/go/utils/strng"
	"strings"
)

func (sb *SqlStatementBuilder) ValidateTables(rdata *relational.RelationalData, nameCache *maps.String2BoolMap) error {
	tables := rdata.TablesMap()
	sb.children = maps.NewSyncMap()

	for tname, _ := range tables {
		sbt, err := NewSqlStatementBuilder(sb.schema, tname, sb.o, sb.db, nameCache)
		if err != nil {
			return err
		}
		sb.children.Put(tname, sbt)
		err = sbt.ValidateTable()
		if err != nil {
			return err
		}
	}
	return nil
}

func (sb *SqlStatementBuilder) ValidateTable() error {

	if sb.nameCache != nil && sb.nameCache.Contains(sb.node.TypeName) {
		return sb.ValidateFields()
	}

	sq := strng.New("select count(*) from ", sb.node.TypeName).String()

	_, err := sb.db.Exec(sq)
	if err != nil && (strings.Contains(err.Error(), "relation") &&
		strings.Contains(err.Error(), "does not exist") ||
		strings.Contains(err.Error(), "no such table")) {
		err = sb.CreateTable()
		if err != nil {
			return err
		}
		if sb.nameCache != nil {
			sb.nameCache.Put(sb.node.TypeName, true)
		}
	} else if err != nil {
		return err
	}

	return nil
}

// @TODO implement this method
func (stmt *SqlStatementBuilder) ValidateFields() error {
	return nil
}

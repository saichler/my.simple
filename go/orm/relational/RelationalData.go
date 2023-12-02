package relational

import (
	"github.com/saichler/my.simple/go/utils/strng"
)

type RelationalData struct {
	trId          string
	rootTableName string
	name2Table    map[string]*Table
}

func NewRelationalData(trId string) *RelationalData {
	return &RelationalData{trId: trId, name2Table: make(map[string]*Table)}
}

func (data *RelationalData) getOrCreateTable(typ string) *Table {
	table, ok := data.name2Table[typ]
	if !ok {
		table = newTable()
		data.name2Table[typ] = table
	}
	return table
}

func (data *RelationalData) TablesMap() map[string]*Table {
	return data.name2Table
}

func (data *RelationalData) String() string {
	str := strng.New("Relational Data:\n")
	for name, table := range data.name2Table {
		str.Add("|- ", name, "\n")
		table.String(str)
	}
	return str.String()
}

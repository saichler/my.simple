package relational

import (
	"fmt"
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

func (data *RelationalData) Print() {
	for name, table := range data.name2Table {
		fmt.Println(name)
		table.Print()
	}
}

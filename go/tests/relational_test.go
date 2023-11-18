package tests

import (
	"fmt"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/orm/relational"
	"github.com/saichler/my.simple/go/tests/model"
	"testing"
)

func TestRelationalData(t *testing.T) {
	common.Introspect.Inspect(&model.MyTestModel{})
	data := createTestModelInstance(1)
	tbls := relational.NewTables("1")
	err := tbls.Add(data, common.Introspect)
	fmt.Println(err)
	tbls.Print()
}

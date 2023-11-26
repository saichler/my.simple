package tests

import (
	"fmt"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/orm"
	"github.com/saichler/my.simple/go/orm/plugins/postgres"
	"testing"
)

func TestPostgresPlugin(t *testing.T) {
	db := newPostgresConnection()
	pp, err := postgres.NewOrmPostgresPlugin(db, "test")
	if err != nil {
		t.Fail()
		fmt.Println("Unable to open database:", err)
		return
	}

	o := orm.NewOrm(pp, common.Introspect)
	sample := createTestModelInstance(5)
	common.Introspect.Inspect(sample)
	err = o.Write(sample)
	if err != nil {
		t.Fail()
		fmt.Println(err)
		return
	}
}

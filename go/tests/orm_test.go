package tests

import (
	"fmt"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/orm"
	"github.com/saichler/my.simple/go/orm/plugins/postgres"
	"github.com/saichler/my.simple/go/orm/plugins/sqlite"
	"testing"
)

func TestPostgresPlugin(t *testing.T) {
	pp := postgres.NewOrmPostgresPlugin()
	db := newPostgresConnection(pp.Decorator())
	o := orm.NewOrm(pp, common.Introspect)
	err := pp.Init(db, "test", o)

	if err != nil {
		t.Fail()
		fmt.Println("Unable to open database:", err)
		return
	}

	sample := createTestModelInstance(5)
	common.Introspect.Inspect(sample)
	err = o.Write(sample)
	if err != nil {
		t.Fail()
		fmt.Println(err)
		return
	}
}

func TestSqlitePlugin(t *testing.T) {
	sp := sqlite.NewOrmSqlitePlugin()
	db := newSqliteConnection(sp.Decorator())
	o := orm.NewOrm(sp, common.Introspect)
	err := sp.Init(db, "", o)

	if err != nil {
		t.Fail()
		fmt.Println("Unable to open database:", err)
		return
	}

	sample := createTestModelInstance(5)
	common.Introspect.Inspect(sample)
	err = o.Write(sample)
	if err != nil {
		t.Fail()
		fmt.Println(err)
		return
	}
	count := 0
	rows, _ := db.Query("select count(*) from mytestmodel;")
	for rows.Next() {
		rows.Scan(&count)
	}
	fmt.Println(count)
}

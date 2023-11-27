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
	db := newPostgresConnection()
	pp := postgres.NewOrmPostgresPlugin()
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
	db := newSqliteConnection()
	pp := sqlite.NewOrmSqlitePlugin()
	o := orm.NewOrm(pp, common.Introspect)
	err := pp.Init(db, "", o)

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

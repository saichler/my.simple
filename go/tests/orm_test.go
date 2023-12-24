package tests

import (
	"fmt"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/orm"
	"github.com/saichler/my.simple/go/orm/plugins/postgres"
	"github.com/saichler/my.simple/go/orm/plugins/sqlite"
	"github.com/saichler/my.simple/go/tests/model"
	"reflect"
	"testing"
)

func TestPostgresPlugin(t *testing.T) {
	pp := postgres.NewOrmPostgresPlugin()
	db := newPostgresConnection(pp.Decorator())
	o := orm.NewOrm(pp, common.Introspect)
	err := pp.Init(o, db, "test")

	if err != nil {
		t.Fail()
		fmt.Println("Unable to open database:", err)
		return
	}

	sample := createTestModelInstance(5)
	common.Introspect.Inspect(sample)
	err = o.Persist(sample)
	if err != nil {
		t.Fail()
		fmt.Println(err)
		return
	}
	count := 0
	rows, _ := db.Query("select count(*) from test.mytestmodel;")
	for rows.Next() {
		rows.Scan(&count)
	}
	fmt.Println(count)

	d, err := o.Fetch(nil)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
	data := d.(map[string]interface{})

	for _, rec := range data {
		sample2 := rec.(*model.MyTestModel)
		/*
			a := fmt.Sprintf("%s", sample)
			b := fmt.Sprintf("%s", sample2)
			fmt.Println(a)
			fmt.Println(b)

		*/
		if !reflect.DeepEqual(sample, sample2) {
			t.Fail()
			fmt.Println("Not Equale")
			return
		}
	}
}

func testSqlitePlugin(t *testing.T) {
	sp := sqlite.NewOrmSqlitePlugin()
	db := newSqliteConnection(sp.Decorator())
	o := orm.NewOrm(sp, common.Introspect)
	err := sp.Init(o, db, "")

	if err != nil {
		t.Fail()
		fmt.Println("Unable to open database:", err)
		return
	}

	sample := createTestModelInstance(5)
	common.Introspect.Inspect(sample)
	err = o.Persist(sample)
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

	d, err := o.Fetch(nil)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
	data := d.(map[string]interface{})

	for _, rec := range data {
		sample2 := rec.(*model.MyTestModel)
		/*
			a := fmt.Sprintf("%s", sample)
			b := fmt.Sprintf("%s", sample2)
			fmt.Println(a)
			fmt.Println(b)

		*/
		if !reflect.DeepEqual(sample, sample2) {
			t.Fail()
			fmt.Println("Not Equale")
			return
		}
	}
}

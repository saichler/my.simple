package tests

import (
	"database/sql"
	"fmt"
	"github.com/saichler/my.security/go/sec"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/defaults"
	model2 "github.com/saichler/my.simple/go/introspect/model"
	"github.com/saichler/my.simple/go/orm"
	"github.com/saichler/my.simple/go/orm/plugins/postgres"
	"github.com/saichler/my.simple/go/security"
	"github.com/saichler/my.simple/go/tests/model"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func init() {
	sec.SetProvider(security.NewShallowSecurityProvider("v7mdWmN7YtkK9o9RXVlezRd7j5Qntohg", "Top Secret"))
	defaults.ApplyDefaults()
	decorateModel()
}

func createTestModelInstance(index int) *model.MyTestModel {
	tag := strconv.Itoa(index)
	sub := &model.MyTestSubModelSingle{
		MyString: "string-sub-" + tag,
		MyInt64:  time.Now().Unix(),
	}
	sub1 := &model.MyTestSubModelSingle{
		MyString: "string-sub-1-" + tag,
		MyInt64:  time.Now().Unix(),
	}
	sub2 := &model.MyTestSubModelSingle{
		MyString: "string-sub-2-" + tag,
		MyInt64:  time.Now().Unix(),
	}
	i := &model.MyTestModel{
		MyString:           "string-" + tag,
		MyFloat64:          123456.123456,
		MyBool:             true,
		MyFloat32:          123.123,
		MyInt32:            int32(index),
		MyInt64:            int64(index * 10),
		MyInt32Slice:       []int32{1, 2, 3, int32(index)},
		MyStringSlice:      []string{"a", "b", "c", "d", tag},
		MyInt32ToInt64Map:  map[int32]int64{1: 11, 2: 22, 3: 33, 4: 44, int32(index): int64(index * 10)},
		MyString2StringMap: map[string]string{"a": "aa", "b": "bb", "c": "cc", tag: tag + tag},
		MySingle:           sub,
		MyModelSlice:       []*model.MyTestSubModelSingle{sub1, sub2},
		MyString2ModelMap:  map[string]*model.MyTestSubModelSingle{sub1.MyString: sub1, sub2.MyString: sub2},
	}
	return i
}

func extractKeyValue(key string) string {
	index1 := strings.LastIndex(key, "<")
	index2 := strings.LastIndex(key, ">")
	return key[index1+1 : index2]
}

func newPostgresConnection(decorator common.DataStoreDecorator) *sql.DB {
	db := decorator.Connect("127.0.0.1", "5432", "postgres", "admin", "test", "disable")
	return db.(*sql.DB)
}

func newSqliteConnection(decorator common.DataStoreDecorator) *sql.DB {
	os.Remove("/tmp/sqlite.db")
	file, _ := os.Create("/tmp/sqlite.db")
	file.Close()
	db := decorator.Connect("/tmp/sqlite.db")
	return db.(*sql.DB)
}

func newPostgresOrm(t *testing.T) (common.IORM, *sql.DB) {
	pp := postgres.NewOrmPostgresPlugin()
	db := newPostgresConnection(pp.Decorator())
	o := orm.NewOrm(pp, common.Introspect)
	err := pp.Init(o, db, "test")

	if err != nil {
		t.Fail()
		fmt.Println("Unable to open database:", err)
		return nil, nil
	}
	sample := createTestModelInstance(0)
	_, err = common.Introspect.Inspect(sample)
	if err != nil {
		t.Fail()
		fmt.Println(err)
		return nil, nil
	}
	return o, db
}

func decorateModel() {
	m := &model.MyTestModel{}
	node, _ := common.Introspect.Inspect(m)
	common.Introspect.AddDecorator(model2.DecoratorType_Primary, []string{"MyString"}, node)
}

package tests

import (
	"database/sql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/saichler/my.security/go/sec"
	"github.com/saichler/my.simple/go/defaults"
	"github.com/saichler/my.simple/go/security"
	"github.com/saichler/my.simple/go/tests/model"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	sec.SetProvider(security.NewShallowSecurityProvider("v7mdWmN7YtkK9o9RXVlezRd7j5Qntohg", "Top Secret"))
	defaults.ApplyDefaults()
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

func newPostgresConnection() *sql.DB {
	def := "host=127.0.0.1 port=5432 user=postgres password=admin dbname=test sslmode=disable"

	// open database
	db, err := sql.Open("postgres", def)
	if err != nil {
		panic(err)
	}
	return db
}

func newSqliteConnection() *sql.DB {
	os.Remove("/tmp/sqlite.db")
	file, _ := os.Create("/tmp/sqlite.db")
	file.Close()

	db, _ := sql.Open("sqlite3", "/tmp/sqlite.db")
	return db
}

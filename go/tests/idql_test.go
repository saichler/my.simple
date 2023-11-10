package tests

import (
	"fmt"
	"github.com/saichler/my.simple/go/idql/idql_parser"
	"github.com/saichler/my.simple/go/idql/idql_query"
	"github.com/saichler/my.simple/go/introspect"
	"testing"
)

func TestFetchParse(t *testing.T) {
	fetch := "fetch elem only 1,2,3 criteria 1=4 and 6=7"
	f, e := idql_parser.NewFetch(fetch)
	if e != nil {
		t.Fail()
		fmt.Println(e)
		return
	}
	if f.String() != "fetch elem only 1,2,3 criteria (1=4 and 6=7)" {
		t.Fail()
		fmt.Println(f.String(), "not eq")
	}
}

func TestFetchQuery(t *testing.T) {
	test := createTestModelInstance(1)
	introspect.Inspect(test)

	request := "fetch mytestmodel only mysingle.mystring, mysingle.myint64 criteria mysingle.mystring='string-sub-1'"
	fetch, err := idql_query.NewFetch(request, introspect.DefaultIntrospect)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	m, e := fetch.Match(test)
	if e != nil {
		fmt.Println(e)
		t.Fail()
		return
	}
	if !m {
		fmt.Println("No Match at ", fetch.String())
		t.Fail()
		return
	}
}

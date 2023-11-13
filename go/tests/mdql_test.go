package tests

import (
	"fmt"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/mdql/mdql_parser"
	"github.com/saichler/my.simple/go/mdql/mdql_query"
	"testing"
)

func TestFetchParse(t *testing.T) {
	fetch := "fetch elem only 1,2,3 criteria 1=4 and 6=7"
	f, e := mdql_parser.NewFetch(fetch)
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
	common.Introspect.Inspect(test)

	request := "fetch mytestmodel only mysingle.mystring, mysingle.myint64 criteria mysingle.mystring='string-sub-1'"
	fetch, err := mdql_query.NewFetch(request, common.Introspect)
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

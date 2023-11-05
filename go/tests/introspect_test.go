package tests

import (
	"fmt"
	"github.com/saichler/my.simple/go/introspect"
	"github.com/saichler/my.simple/go/tests/model"
	"testing"
)

func TestIntrospect(t *testing.T) {
	m := &model.MyTestModel{}
	_, err := introspect.Inspect(m)
	if err != nil {
		t.Fail()
		fmt.Println(err)
		return
	}
	introspect.DefaultIntrospect.Print()
	nodes := introspect.DefaultIntrospect.Nodes(false, false)
	if len(nodes) != 13 {
		fmt.Println("Expected length to be 13 but got ", len(nodes))
		t.Fail()
		return
	}

	nodes = introspect.DefaultIntrospect.Nodes(false, true)
	if len(nodes) != 2 {
		fmt.Println("Expected length to be 2 roots but got ", len(nodes))
		t.Fail()
		return
	}

	nodes = introspect.DefaultIntrospect.Nodes(true, false)
	if len(nodes) != 12 {
		fmt.Println("Expected length to be 12 leafs but got ", len(nodes))
		t.Fail()
		return
	}

	_, ok := introspect.DefaultIntrospect.Node("mytestmodel.myint32toint64map")
	if !ok {
		fmt.Println("Could not fetch node")
		t.Fail()
		return
	}

	_, ok = introspect.DefaultIntrospect.NodeByType(&model.MyTestSubModelSingle{})
	if !ok {
		fmt.Println("Could not fetch node by type")
		t.Fail()
		return
	}
}

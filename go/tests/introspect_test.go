package tests

import (
	"fmt"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/tests/model"
	"testing"
)

func TestIntrospect(t *testing.T) {
	m := &model.MyTestModel{}
	_, err := common.Introspect.Inspect(m)
	if err != nil {
		t.Fail()
		fmt.Println(err)
		return
	}
	common.Introspect.Print()
	nodes := common.Introspect.Nodes(false, false)
	expectedNodes := 16
	if len(nodes) != expectedNodes {
		fmt.Println("Expected length to be ", expectedNodes, " but got ", len(nodes))
		t.Fail()
		return
	}

	nodes = common.Introspect.Nodes(false, true)
	if len(nodes) != 2 {
		fmt.Println("Expected length to be 2 roots but got ", len(nodes))
		t.Fail()
		return
	}

	nodes = common.Introspect.Nodes(true, false)
	if len(nodes) != 13 {
		fmt.Println("Expected length to be 13 leafs but got ", len(nodes))
		t.Fail()
		return
	}

	_, ok := common.Introspect.Node("mytestmodel.myint32toint64map")
	if !ok {
		fmt.Println("Could not fetch node")
		t.Fail()
		return
	}

	_, ok = common.Introspect.NodeByType(&model.MyTestSubModelSingle{})
	if !ok {
		fmt.Println("Could not fetch node by type")
		t.Fail()
		return
	}
}

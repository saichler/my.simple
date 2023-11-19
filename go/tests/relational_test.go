package tests

import (
	"fmt"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/orm/relational"
	"github.com/saichler/my.simple/go/tests/model"
	"github.com/saichler/my.simple/go/utils/strng"
	"reflect"
	"testing"
)

func TestRelationalData(t *testing.T) {
	common.Introspect.Inspect(&model.MyTestModel{})
	data := []*model.MyTestModel{createTestModelInstance(1),
		createTestModelInstance(2)}
	rdata := relational.NewRelationalData("1")

	err := rdata.AddInstances(data, common.Introspect)
	if err != nil {
		t.Fail()
		fmt.Println(err)
		return
	}
	//rdata.Print()

	instances, err := rdata.ToIstances(common.Introspect)
	if len(instances) != 2 {
		t.Fail()
		fmt.Println("Expected 2 instances but got ", len(instances))
		return
	}
	for key, val := range instances {
		index := strng.FromString(extractKeyValue(key)).Interface().(int)
		if !reflect.DeepEqual(data[index], val) {
			t.Fail()
			fmt.Println("Index", index, " not equal")
			return
		}
	}
}

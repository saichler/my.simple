package tests

import (
	"fmt"
	"github.com/saichler/my.simple/go/introspect"
	"github.com/saichler/my.simple/go/orm/relational"
	registry2 "github.com/saichler/my.simple/go/registry"
	"github.com/saichler/my.simple/go/tests/model"
	"github.com/saichler/my.simple/go/utils/strng"
	"reflect"
	"testing"
)

func TestRelationalData(t *testing.T) {
	registry := registry2.NewStructRegistry()
	inspect := introspect.NewIntrospect(registry)

	inspect.Inspect(&model.MyTestModel{})
	data := []*model.MyTestModel{createTestModelInstance(1),
		createTestModelInstance(2)}
	relationalData := relational.NewRelationalData("<transaction ref>")

	err := relationalData.AddInstances(data, inspect)
	if err != nil {
		t.Fail()
		fmt.Println(err)
		return
	}
	//relationalData.Print()

	instances, err := relationalData.ToIstances(inspect)
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

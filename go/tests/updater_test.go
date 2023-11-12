package tests

import (
	"fmt"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/tests/model"
	"github.com/saichler/my.simple/go/updater"
	"testing"
)

func TestUpdater(t *testing.T) {
	common.Introspect.Inspect(&model.MyTestModel{})
	updater := updater.NewUpdater(common.Introspect, false)
	aside := createTestModelInstance(0)
	zside := &model.MyTestModel{MyString: "updated"}
	uside := common.Introspect.Clone(aside).(*model.MyTestModel)
	err := updater.Update(aside, zside, common.Introspect)
	if err != nil {
		t.Fail()
		fmt.Println(err)
		return
	}

	changes := updater.Changes()

	if len(changes) != 1 {
		t.Fail()
		fmt.Println("Expected 1 change but got ", len(updater.Changes()))
		for _, c := range changes {
			fmt.Println(c.String())
		}
		return
	}

	if aside.MyString != zside.MyString {
		t.Fail()
		fmt.Println("1 Expected ", zside.MyString, " got ", aside.MyString)
		return
	}

	for _, change := range changes {
		change.Apply(uside)
	}

	if uside.MyString != aside.MyString {
		fmt.Println("2 Expected ", aside.MyString, " got ", uside.MyString)
		t.Fail()
		return
	}
}

package tests

import (
	"github.com/saichler/my.simple/go/common"
	model2 "github.com/saichler/my.simple/go/net/model"
	"github.com/saichler/my.simple/go/net/port"
	"github.com/saichler/my.simple/go/net/switching"
	security2 "github.com/saichler/my.simple/go/security"
	"github.com/saichler/my.simple/go/tests/model"
	"github.com/saichler/my.simple/go/utils/logs"
	"testing"
	"time"
)

var securityProvider = security2.NewShallowSecurityProvider(common.GenerateAES256Key(), "testing 1..2..3")

func TestPortsSwitch(t *testing.T) {
	mh := &model.MyTestModelHandler{}
	mh.Expected = "Hello Test"
	common.ServicePoints.RegisterServicePoint(&model.MyTestModel{}, mh, common.Registry)

	sw := switching.NewSwitchService(common.NetConfig.DefaultSwitchPort, common.Registry, common.HealthCenter, common.ServicePoints)
	go sw.Start()
	time.Sleep(time.Millisecond * 100)

	p1, err := port.ConnectTo("127.0.0.1", common.NetConfig.DefaultSwitchPort, nil, common.Registry, common.ServicePoints)
	if err != nil {
		t.Fail()
		logs.Error("err:", err.Error())
		return
	}

	p2, err := port.ConnectTo("127.0.0.1", common.NetConfig.DefaultSwitchPort, nil, common.Registry, common.ServicePoints)
	if err != nil {
		t.Fail()
		logs.Error(err)
		return
	}

	m := &model.MyTestModel{}
	m.MyString = mh.Expected

	time.Sleep(time.Millisecond * 100)
	err = p1.Do(model2.Action_Action_Get, p2.Uuid(), m)

	time.Sleep(time.Second * 10)

	if !mh.Passed {
		t.Fail()
	}

	p1.Shutdown()
	p2.Shutdown()
	sw.Shutdown()
	time.Sleep(time.Second)
}

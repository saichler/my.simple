package tests

import (
	"github.com/saichler/my.simple/go/common"
	model2 "github.com/saichler/my.simple/go/net/model"
	"github.com/saichler/my.simple/go/net/port"
	"github.com/saichler/my.simple/go/net/switching"
	"github.com/saichler/my.simple/go/tests/model"
	"github.com/saichler/my.simple/go/utils/security"
	"testing"
	"time"
)

var key = security.GenerateAES256Key()
var secret = "testing..1..2..3"

func TestPortsSwitch(t *testing.T) {
	sw := switching.NewSwitchService(key, secret, common.NetConfig.DefaultSwitchPort)
	go sw.Start()
	time.Sleep(time.Millisecond * 100)
	p1, err := port.ConnectTo("127.0.0.1", key, secret, common.NetConfig.DefaultSwitchPort, nil)
	if err != nil {
		t.Fail()
		return
	}
	p2, err := port.ConnectTo("127.0.0.1", key, secret, common.NetConfig.DefaultSwitchPort, nil)
	if err != nil {
		t.Fail()
		return
	}

	m := &model.MyTestModel{}
	m.MyString = "Hello World"

	err = p1.Do(model2.Action_Action_Get, p2.Uuid(), m)

	time.Sleep(time.Second * 10)

	p1.Shutdown()
	p2.Shutdown()
	sw.Shutdown()
	time.Sleep(time.Second)
}

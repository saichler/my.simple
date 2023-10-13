package tests

import (
	"fmt"
	"github.com/saichler/my.simple/go/common"
	model2 "github.com/saichler/my.simple/go/net/model"
	"github.com/saichler/my.simple/go/net/port"
	"github.com/saichler/my.simple/go/net/switching"
	"github.com/saichler/my.simple/go/services/service_point"
	"github.com/saichler/my.simple/go/tests/model"
	"github.com/saichler/my.simple/go/utils/logs"
	"github.com/saichler/my.simple/go/utils/security"
	"google.golang.org/protobuf/proto"
	"testing"
	"time"
)

var key = security.GenerateAES256Key()
var secret = "testing..1..2..3"

type MyTestModelHandler struct {
}

func (my *MyTestModelHandler) Post(pb proto.Message, port common.Port) (proto.Message, error) {
	return nil, nil
}
func (my *MyTestModelHandler) Put(pb proto.Message, port common.Port) (proto.Message, error) {
	return nil, nil
}
func (my *MyTestModelHandler) Patch(pb proto.Message, port common.Port) (proto.Message, error) {
	return nil, nil
}
func (my *MyTestModelHandler) Delete(pb proto.Message, port common.Port) (proto.Message, error) {
	return nil, nil
}
func (my *MyTestModelHandler) Get(pb proto.Message, port common.Port) (proto.Message, error) {
	mmm := pb.(*model.MyTestModel)
	fmt.Println(mmm.MyString)
	return nil, nil
}

func TestPortsSwitch(t *testing.T) {

	service_point.RegisterServicePoint(&model.MyTestModel{}, &MyTestModelHandler{})

	sw := switching.NewSwitchService(key, secret, common.NetConfig.DefaultSwitchPort)
	go sw.Start()
	time.Sleep(time.Millisecond * 100)

	p1, err := port.ConnectTo("127.0.0.1", key, secret, common.NetConfig.DefaultSwitchPort, nil)
	if err != nil {
		t.Fail()
		logs.Error(err)
		return
	}

	p2, err := port.ConnectTo("127.0.0.1", key, secret, common.NetConfig.DefaultSwitchPort, nil)
	if err != nil {
		t.Fail()
		logs.Error(err)
		return
	}

	m := &model.MyTestModel{}
	m.MyString = "Hello World"

	time.Sleep(time.Millisecond * 100)
	err = p1.Do(model2.Action_Action_Get, p2.Uuid(), m)

	time.Sleep(time.Second * 10)

	p1.Shutdown()
	p2.Shutdown()
	sw.Shutdown()
	time.Sleep(time.Second)
}

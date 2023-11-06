package model

import (
	"fmt"
	"github.com/saichler/my.simple/go/common"
	"google.golang.org/protobuf/proto"
)

type MyTestModelHandler struct {
	Passed   bool
	Expected string
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
	mdl := pb.(*MyTestModel)
	fmt.Println("Got message with ", mdl.MyString)
	if mdl.MyString == my.Expected {
		my.Passed = true
	}
	return nil, nil
}
func (my *MyTestModelHandler) EndPoint() string {
	return "/test"
}

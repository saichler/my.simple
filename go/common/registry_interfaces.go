package common

import (
	"google.golang.org/protobuf/proto"
	"reflect"
)

type IRegistry interface {
	RegisterStruct(interface{}) bool
	RegisterStructType(reflect.Type) bool
	NewProtobufInstance(string) (proto.Message, error)
	NewInstance(string) (interface{}, error)
	TypeByName(string) (reflect.Type, error)
}

var Registry IRegistry

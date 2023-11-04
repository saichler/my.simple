package registry

import (
	"errors"
	"github.com/saichler/my.simple/go/utils/logs"
	"github.com/saichler/my.simple/go/utils/maps"
	"google.golang.org/protobuf/proto"
	"reflect"
)

type StructRegistryImpl struct {
	structName2Type *maps.String2TypeMap
}

var structRegistry = newStructRegistry()

func newStructRegistry() *StructRegistryImpl {
	sr := &StructRegistryImpl{}
	sr.structName2Type = maps.NewString2TypeMao()
	return sr
}

func RegisterStruct(any interface{}) bool {
	v := reflect.ValueOf(any)
	if v.Kind() == reflect.Ptr {
		return RegisterStructType(v.Elem().Type())
	}
	return RegisterStructType(v.Type())
}

func RegisterStructType(t reflect.Type) bool {
	return structRegistry.structName2Type.Put(t.Name(), t)
}

func TypeByName(name string) (reflect.Type, error) {
	value, ok := structRegistry.structName2Type.Get(name)
	if !ok {
		return nil, errors.New("Unknown Struct Type: " + name)
	}
	return value, nil
}

func NewProtobufInstance(t string) (proto.Message, error) {
	if t == "" {
		return nil, logs.Error("cannot create a new protobuf instance from blank type name")
	}
	typ, ok := structRegistry.structName2Type.Get(t)
	if !ok {
		return nil, logs.Error("Struct Type ", t, " is not registered")
	}
	n := reflect.New(typ)
	if !n.IsValid() {
		return nil, logs.Error("Was not able to create new instance of type ", t)
	}
	pb, ok := n.Interface().(proto.Message)
	if !ok {
		return nil, logs.Error("Type ", t, " is not a protobuf")
	}
	return pb, nil
}

func NewInstance(name string) (interface{}, error) {
	typ, ok := structRegistry.structName2Type.Get(name)
	if !ok {
		return nil, errors.New("Unknown Struct Type: " + name)
	}
	return reflect.New(typ).Interface(), nil
}

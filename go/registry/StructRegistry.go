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

func NewStructRegistry() *StructRegistryImpl {
	sr := &StructRegistryImpl{}
	sr.structName2Type = maps.NewString2TypeMao()
	return sr
}

func (r *StructRegistryImpl) RegisterStruct(any interface{}) bool {
	v := reflect.ValueOf(any)
	if v.Kind() == reflect.Ptr {
		return r.RegisterStructType(v.Elem().Type())
	}
	return r.RegisterStructType(v.Type())
}

func (r *StructRegistryImpl) RegisterStructType(t reflect.Type) bool {
	return r.structName2Type.Put(t.Name(), t)
}

func (r *StructRegistryImpl) TypeByName(name string) (reflect.Type, error) {
	value, ok := r.structName2Type.Get(name)
	if !ok {
		return nil, errors.New("Unknown Struct Type: " + name)
	}
	return value, nil
}

func (r *StructRegistryImpl) NewProtobufInstance(t string) (proto.Message, error) {
	if t == "" {
		return nil, logs.Error("cannot create a new protobuf instance from blank type name")
	}
	typ, ok := r.structName2Type.Get(t)
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

func (r *StructRegistryImpl) NewInstance(name string) (interface{}, error) {
	typ, ok := r.structName2Type.Get(name)
	if !ok {
		return nil, errors.New("Unknown Struct Type: " + name)
	}
	return reflect.New(typ).Interface(), nil
}

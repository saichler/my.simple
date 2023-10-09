package types

import (
	"errors"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/utils/logs"
	"google.golang.org/protobuf/proto"
	"reflect"
	"sync"
)

type TypesImpl struct {
	typeName2Type        map[string]reflect.Type
	typeName2TypeHandler map[string]common.TypeHandler
	mtx                  *sync.RWMutex
}

var types = newTypes()

func newTypes() *TypesImpl {
	types := &TypesImpl{}
	types.typeName2Type = make(map[string]reflect.Type)
	types.typeName2TypeHandler = make(map[string]common.TypeHandler)
	types.mtx = &sync.RWMutex{}
	return types
}

func (types *TypesImpl) registerType(t reflect.Type) error {
	types.mtx.Lock()
	defer types.mtx.Unlock()
	_, ok := types.typeName2Type[t.Name()]
	if ok {
		return nil
	}
	types.typeName2Type[t.Name()] = t
	return nil
}

func Type(name string) (reflect.Type, error) {
	return types.typ(name)
}

func (types *TypesImpl) typ(name string) (reflect.Type, error) {
	types.mtx.RLock()
	defer types.mtx.RUnlock()
	typ, ok := types.typeName2Type[name]
	if !ok {
		return nil, errors.New("Unknown Struct Type: " + name)
	}
	return typ, nil
}

func NewProto(name string) (proto.Message, error) {
	return types.newProto(name)
}

func NewInterface(name string) (interface{}, error) {
	return types.newInterface(name)
}

func (types *TypesImpl) newInterface(name string) (interface{}, error) {
	types.mtx.RLock()
	defer types.mtx.RUnlock()
	typ, ok := types.typeName2Type[name]
	if !ok {
		return nil, errors.New("Unknown Struct Type: " + name)
	}
	return reflect.New(typ).Interface(), nil
}

func (types *TypesImpl) newProto(t string) (proto.Message, error) {
	if t == "" {
		return nil, logs.Error("Cannot New with blank type")
	}
	types.mtx.Lock()
	typ, ok := types.typeName2Type[t]
	types.mtx.Unlock()
	if !ok {
		return nil, logs.Error("Type ", t, " is not registered")
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

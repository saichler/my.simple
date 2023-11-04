package protobuf_object

import (
	"github.com/saichler/my.simple/go/utils/logs"
	"reflect"
	"sync"
)

var protobufObjectTypes = make(map[reflect.Kind]ProtobufObjectType)
var sizeObjectType = &Int32{}
var stringObjectType = &String{}
var mtx = &sync.Mutex{}

type ProtobufObject struct {
	data     []byte
	location int
}

type ProtobufObjectType interface {
	add(interface{}) ([]byte, int)
	get([]byte, int) (interface{}, int)
}

func init() {
	mtx.Lock()
	defer mtx.Unlock()
	if len(protobufObjectTypes) == 0 {
		protobufObjectTypes[reflect.Int] = &Int{}
		protobufObjectTypes[reflect.Uint32] = &UInt32{}
		protobufObjectTypes[reflect.Uint64] = &UInt64{}
		protobufObjectTypes[reflect.Int32] = &Int32{}
		protobufObjectTypes[reflect.Int64] = &Int64{}
		protobufObjectTypes[reflect.Float32] = &Float32{}
		protobufObjectTypes[reflect.Float64] = &Float64{}
		protobufObjectTypes[reflect.String] = &String{}
		protobufObjectTypes[reflect.Ptr] = &Proto{}
		protobufObjectTypes[reflect.Slice] = &Slice{}
		protobufObjectTypes[reflect.Map] = &Map{}
		protobufObjectTypes[reflect.Bool] = &Bool{}
	}
}

func NewProtobufObject(data []byte, location int) *ProtobufObject {
	obj := &ProtobufObject{}
	obj.data = data
	obj.location = location
	return obj
}

func (obj *ProtobufObject) Data() []byte {
	return obj.data
}

func (obj *ProtobufObject) Location() int {
	return obj.location
}

func (obj *ProtobufObject) Add(any interface{}) error {
	kind := reflect.ValueOf(any).Kind()
	mtx.Lock()
	et, ok := protobufObjectTypes[kind]
	mtx.Unlock()
	if !ok {
		return logs.Error("Did not find any Object for kind", kind.String())
	}
	obj.addKind(kind)
	b, l := et.add(any)
	obj.location += l
	obj.data = append(obj.data, b...)
	return nil
}

func (obj *ProtobufObject) Get() (interface{}, error) {
	kind := obj.getKind()
	mtx.Lock()
	et, ok := protobufObjectTypes[kind]
	mtx.Unlock()
	if !ok {
		return nil, logs.Error("Did not find any Object for kind", kind.String())
	}
	d, l := et.get(obj.data, obj.location)
	obj.location += l
	return d, nil
}

func (obj *ProtobufObject) addKind(kind reflect.Kind) {
	b, l := sizeObjectType.add(int32(kind))
	obj.location += l
	obj.data = append(obj.data, b...)
}

func (obj *ProtobufObject) getKind() reflect.Kind {
	i, l := sizeObjectType.get(obj.data, obj.location)
	obj.location += l
	return reflect.Kind(i.(int32))
}

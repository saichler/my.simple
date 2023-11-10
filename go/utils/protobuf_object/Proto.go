package protobuf_object

import (
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/utils/logs"
	"google.golang.org/protobuf/proto"
	"reflect"
)

type Proto struct {
}

func (_type *Proto) Add(any interface{}) ([]byte, int) {
	if any == nil || reflect.ValueOf(any).IsNil() {
		sizeBytes, _ := sizeObjectType.Add(int32(-1))
		return sizeBytes, 4
	}

	typ := reflect.ValueOf(any).Elem().Type()
	typeName := typ.Name()
	_, err := common.Registry.TypeByName(typeName)
	if err != nil {
		common.Registry.RegisterStructType(typ)
	}
	pb := any.(proto.Message)
	pbData, err := proto.Marshal(pb)
	if err != nil {
		logs.Error("Failed To marshal proto ", typeName, " in protobuf object:", err)
		return []byte{}, 0
	}

	data, _ := stringObjectType.Add(typeName)
	sizeData, _ := sizeObjectType.Add(int32(len(pbData)))
	data = append(data, sizeData...)
	data = append(data, pbData...)

	return data, len(data)
}

func (_type *Proto) Get(data []byte, location int) (interface{}, int) {
	l, _ := sizeObjectType.Get(data, location)
	size := l.(int32)
	if size == -1 || size == 0 {
		return nil, 4
	}

	typeN, typeSize := stringObjectType.Get(data, location)
	typeName := typeN.(string)
	pb, err := common.Registry.NewProtobufInstance(typeName)
	if err != nil {
		logs.Error("Unknown proto name ", typeName, " in registry, please register it.")
		return []byte{}, 0
	}
	location += typeSize
	s, _ := sizeObjectType.Get(data, location)
	size = s.(int32)
	location += 4
	protoData := data[location : location+int(size)]

	err = proto.Unmarshal(protoData, pb)
	if err != nil {
		logs.Error("Failed To unmarshal proto ", typeName, ":", err)
		return []byte{}, 0
	}
	return pb, typeSize + 4 + int(size)
}

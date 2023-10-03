package types

import (
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/p2p/model"
	"github.com/saichler/my.simple/go/utils/logs"
	"google.golang.org/protobuf/proto"
	"reflect"
)

func (types *TypesImpl) RegisterTypeHandler(pb proto.Message, handler common.TypeHandler) error {
	if pb == nil {
		return logs.Error("Cannot register handler with nil proto")
	}
	t := reflect.ValueOf(pb).Elem().Type()
	if handler == nil {
		return logs.Error("Cannot register nil handler for type ", t.Name())
	}

	types.registerType(common.TypeOf(pb))

	types.mtx.Lock()
	defer types.mtx.Unlock()
	types.typeName2TypeHandler[t.Name()] = handler
	return nil
}

func (types *TypesImpl) Hanlde(packet *model.Packet, port common.Port) (proto.Message, error) {
	pbInstance, err := Types.new(packet.ProtoTypeName)
	if err != nil {
		return nil, err
	}
	err = proto.Unmarshal(packet.Data, pbInstance)
	if err != nil {
		return nil, logs.Error("Unmarshal failed with:", err)
	}
	return Types.handle(pbInstance, packet.Action, port)
}

func (types *TypesImpl) handle(pb proto.Message, action model.Action, port common.Port) (proto.Message, error) {
	tName := reflect.ValueOf(pb).Elem().Type().Name()
	types.mtx.Lock()
	h, ok := types.typeName2TypeHandler[tName]
	types.mtx.Unlock()
	if !ok {
		return nil, logs.Error("Cannot find handler for type ", tName)
	}
	switch action {
	case model.Action_Action_Post:
		return h.Post(pb, port)
	case model.Action_Action_Put:
		return h.Put(pb, port)
	case model.Action_Action_Patch:
		return h.Patch(pb, port)
	case model.Action_Action_Delete:
		return h.Delete(pb, port)
	case model.Action_Action_Get:
		return h.Get(pb, port)
	case model.Action_Action_Invalid:
		return nil, logs.Error("Invalid Crud Operation, ignoring")
	}
	panic("Unknown Action:" + action.String())
}

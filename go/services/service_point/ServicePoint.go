package service_point

import (
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/net/model"
	"github.com/saichler/my.simple/go/utils/logs"
	"github.com/saichler/my.simple/go/utils/maps"
	"google.golang.org/protobuf/proto"
	"reflect"
)

type ServicePoints struct {
	structName2ServicePoint *maps.String2ServicePointMap
}

func NewServicePoints() *ServicePoints {
	sp := &ServicePoints{}
	sp.structName2ServicePoint = maps.NewString2ServicePointMap()
	return sp
}

func (servicePoints *ServicePoints) RegisterServicePoint(pb proto.Message, handler common.ServicePointHandler, registry common.IRegistry) error {
	if pb == nil {
		return logs.Error("cannot register handler with nil proto")
	}
	typ := reflect.ValueOf(pb).Elem().Type()
	if handler == nil {
		return logs.Error("cannot register nil handler for type ", typ.Name())
	}
	registry.RegisterStructType(typ)
	servicePoints.structName2ServicePoint.Put(typ.Name(), handler)
	return nil
}

func (servicePoints *ServicePoints) Handle(pb proto.Message, action model.Action, port common.Port) (proto.Message, error) {
	tName := reflect.ValueOf(pb).Elem().Type().Name()
	h, ok := servicePoints.structName2ServicePoint.Get(tName)
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
		return nil, logs.Error("Invalid Action, ignoring")
	}
	panic("Unknown Action:" + action.String())
}

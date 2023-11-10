package common

import (
	"github.com/saichler/my.simple/go/net/model"
	"google.golang.org/protobuf/proto"
)

type ServicePointHandler interface {
	Post(proto.Message, Port) (proto.Message, error)
	Put(proto.Message, Port) (proto.Message, error)
	Patch(proto.Message, Port) (proto.Message, error)
	Delete(proto.Message, Port) (proto.Message, error)
	Get(proto.Message, Port) (proto.Message, error)
	EndPoint() string
}

type IServicePoints interface {
	RegisterServicePoint(proto.Message, ServicePointHandler, IRegistry) error
	Handle(proto.Message, model.Action, Port) (proto.Message, error)
}

var ServicePoints IServicePoints

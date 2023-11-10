package common

import (
	"github.com/saichler/my.simple/go/net/model"
	"google.golang.org/protobuf/proto"
)

type IServicePoints interface {
	RegisterServicePoint(proto.Message, ServicePointHandler, IRegistry) error
	Handle(proto.Message, model.Action, Port) (proto.Message, error)
}

var ServicePoints IServicePoints

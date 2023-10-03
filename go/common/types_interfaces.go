package common

import "google.golang.org/protobuf/proto"

type TypeHandler interface {
	Post(proto.Message, Port) (proto.Message, error)
	Put(proto.Message, Port) (proto.Message, error)
	Patch(proto.Message, Port) (proto.Message, error)
	Delete(proto.Message, Port) (proto.Message, error)
	Get(proto.Message, Port) (proto.Message, error)
}

package common

import (
	"github.com/saichler/my.simple/go/net/model"
	"google.golang.org/protobuf/proto"
)

var NetConfig = &model.NetConfig{
	MaxDataSize:        1024 * 1024,
	DefaultTxQueueSize: 1000,
	DefaultRxQueueSize: 1000,
	DefaultSwitchPort:  50000,
}

type Port interface {
	Start()
	Addr() string
	Uuid() string
	Send([]byte) error
	Name() string
	Do(model.Action, string, proto.Message) error
	Shutdown()
	CreatedAt() int64
}

type DatatListener interface {
	PortShutdown(port Port)
	HandleData([]byte, Port)
}

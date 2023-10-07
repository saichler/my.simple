package common

import "github.com/saichler/my.simple/go/net/model"

var NetConfig = &model.NetConfig{
	MaxDataSize:        1024 * 1024,
	DefaultTxQueueSize: 1000,
	DefaultRxQueueSize: 1000,
}

type Port interface {
	Start()
	Addr() string
	Uuid() string
	Send([]byte) error
	Shutdown()
}

type RawDataListener interface {
	DataReceived([]byte, Port)
	PortShutdown(port Port)
}

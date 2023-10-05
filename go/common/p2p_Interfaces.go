package common

import "google.golang.org/protobuf/proto"

var MAX_DATA_SIZE = 1024 * 1024 * 50
var LARGE_PACKET = 1024 * 1024 * 5

type Port interface {
	SendRawData([]byte) error
	SendProtobuf(proto.Message) error
}

type Listener interface {
	DataReceived([]byte, Port)
	PortShutdown(port Port)
}

package common

var MAX_DATA_SIZE = int64(1024 * 1024 * 50)
var LARGE_PACKET = 1024 * 1024 * 5

type Port interface {
	Send([]byte) error
}

type Listener interface {
	DataReceived([]byte, Port)
	PortShutdown(port Port)
}

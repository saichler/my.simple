package port

import "github.com/saichler/my.simple/go/utils/logs"

// loop decoded data ([]byte) from the NX queue and notify the data listener
func (port *PortImpl) incomingDataNotifier() {
	// As long as the port is active
	for port.active {
		// Get the next data in the queue, if the queue is empty this is a blocking call
		data := port.nx.Next()
		// If there is data and listener
		if data != nil && port.listener != nil {
			// notify the data listener that there is a incoming data and providing the port as reference
			// If a reply is needed
			//@TODO - should we call below as go routing?
			port.listener.DataReceived(data, port)
		} else if data != nil {
			logs.Info("No Data Listener for packet:", string(data))
		}
	}
}

package port

import (
	"fmt"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/net/protocol"
	"github.com/saichler/my.simple/go/utils/logs"
)

// loop and Read incoming data from the socket
func (port *PortImpl) readFromSocket() {
	// While the port is active
	for port.active {
		// read data ([]byte) from socket
		data, err := common.Read(port.conn)

		//If therer is an error
		if err != nil {
			// If the secret is not blank, it means that this port is
			// initating the connection, so try to reconnect
			if port.secret != "" {
				// Attempt to reconnect
				port.attemptToReconnect()
				// And try to read the data again
				data, err = common.Read(port.conn)
			} else {
				// If this is the receiving port, break and clean resources.
				logs.Error(err)
				break
			}
		}
		if data != nil {
			// If still active, write the data to the RX queue
			if port.active {
				port.rx.Add(data)
			}
		} else {
			// If data is nil, it means the port was shutdown
			// so break and cleanup
			break
		}
	}
	logs.Info("Connection Read for ", port.Name(), " ended.")
	//Just in case, mark the port as shutdown so other thread will stop as well.
	port.Shutdown()
}

// Notify the RawDataListener on new data
func (port *PortImpl) notifyRawDataListener() {
	// While the port is active
	for port.active {
		// Read next data ([]byte) block
		data := port.rx.Next()
		// If data is not nil
		if data != nil {
			if port.listener != nil {
				port.listener.DataReceived(data, port)
			} else {
				s, _, _ := protocol.HeaderOf(data)
				fmt.Println(port.Name(), "Received message from", s)
			}
		}
	}
	logs.Info("notify data listener for ", port.Name(), " Ended")
	port.Shutdown()
}

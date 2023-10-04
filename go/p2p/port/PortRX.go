package port

import (
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/utils/logs"
	"github.com/saichler/my.simple/go/utils/security"
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

// loop and decode incoming data from the RX queue and place it in the NX queue
func (port *PortImpl) decodeIncomingData() {
	// While the port is active
	for port.active {
		// Read next data ([]byte) block
		data := port.rx.Next()
		// If data is not nil
		if data != nil {
			encString := string(data)
			// Decode and unencrypt the data
			decodedData, err := security.Decode(encString, port.key)
			if err != nil {
				// On any error, break and cleanup
				break
			}
			// Add the decoded data to the NX queue
			port.nx.Add(decodedData)
		}
	}
	logs.Info("Message Processing for ", port.Name(), " Ended")
	port.Shutdown()
}

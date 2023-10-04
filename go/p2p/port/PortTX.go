package port

import (
	"errors"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/utils/logs"
	"github.com/saichler/my.simple/go/utils/security"
	"google.golang.org/protobuf/proto"
)

// loop of Writing data to socket
func (port *PortImpl) write() {
	// As long ad the port is active
	for port.active {
		// Get next data to write to the socket from the TX queue, if no data, this is a blocking call
		data := port.tx.Next()
		// if the data is not nil
		if data != nil && port.active {
			port.writeMutex.L.Lock()
			//Write the data to the socket
			err := common.Write(data, port.conn)
			port.writeMutex.L.Unlock()
			// If there is an error
			if err != nil {
				// If the port has a secret, it means it is the initiating port, so try to reconnect
				// and send the data.
				if port.secret != "" {
					port.attemptToReconnect()
					err = common.Write(data, port.conn)
				} else {
					break
				}
			}
		} else {
			// if the data is nil, break and cleanup
			break
		}
	}
	logs.Info("Connection Write for ", port.Name(), " ended.")
	port.Shutdown()
}

// Encrypt and send raw ([]byte) data
func (port *PortImpl) SendRawData(data []byte) error {
	// if the port is still active
	if port.active {
		// Encrypt the data
		encData, err := security.Encode(data, port.key)
		if err != nil {
			return err
		}
		// Add the encrypted data to the TX queue
		port.tx.Add([]byte(encData))
	} else {
		return errors.New("Port is not active")
	}
	return nil
}

// marshal and send protobuf
func (port *PortImpl) SendProtobuf(pb proto.Message) error {
	// marshal the protobuf
	data, err := proto.Marshal(pb)
	if err != nil {
		return err
	}
	// Send the data over the wire
	return port.SendRawData(data)
}

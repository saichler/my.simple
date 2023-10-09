package port

import (
	"errors"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/net/model"
	"github.com/saichler/my.simple/go/net/protocol"
	"github.com/saichler/my.simple/go/utils/logs"
	"google.golang.org/protobuf/proto"
)

// loop of Writing data to socket
func (port *PortImpl) writeToSocket() {
	// As long ad the port is active
	for port.active {
		// Get next data to write to the socket from the TX queue, if no data, this is a blocking call
		data := port.tx.Next()
		// if the data is not nil
		if data != nil && port.active {
			//Write the data to the socket
			err := common.Write(data, port.conn)
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
	logs.Debug("Connection Write for ", port.Name(), " ended.")
	port.Shutdown()
}

// Send Add the raw data to the tx queue to be written to the socket
func (port *PortImpl) Send(data []byte) error {
	// if the port is still active
	if port.active {
		// Add the data to the TX queue
		port.tx.Add(data)
	} else {
		return errors.New("Port is not active")
	}
	return nil
}

// Do is wrapping a protobuf with a secure message and send it to the switch
func (port *PortImpl) Do(action model.Action, dest string, pb proto.Message) error {
	// Create message payload
	data, err := protocol.CreateMessageFor(model.Priority_P0, action, port.key, port.uuid, dest, pb)
	if err != nil {
		logs.Error("Failed to create message:", err)
		return err
	}
	//Send the secure message to the switch
	port.Send(data)
	return nil
}

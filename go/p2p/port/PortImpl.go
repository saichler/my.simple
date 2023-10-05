package port

import (
	"github.com/google/uuid"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/utils/logs"
	"github.com/saichler/my.simple/go/utils/queues"
	"github.com/saichler/my.simple/go/utils/strng"
	"net"
	"sync"
	"time"
)

type PortImpl struct {
	// The encryption key
	key string
	// The secret of this connection
	secret string
	// generated UUID for the port
	uuid string
	// Once connected, what is the other side uuid
	zside string
	// The incoming data queue
	rx *queues.ByteSliceQueue
	// The outgoing data queue
	tx *queues.ByteSliceQueue
	// The incoming data queue already decoded
	nx *queues.ByteSliceQueue
	// The connection
	conn net.Conn
	// is the port active
	active bool
	// The incoming data listener
	listener common.Listener
	// The ip & port of this port.
	ipAndPort string

	portType  string
	serviceId string

	host          string
	destPort      int
	reconnectMtx  *sync.Mutex
	doneReconnect bool
}

// Instantiate a new port with a connection
func NewPortImpl(local bool, con net.Conn, key string, listener common.Listener, maxInputQueueSize, maxOutputQueueSize int) *PortImpl {
	port := &PortImpl{}
	port.uuid = uuid.New().String()
	port.conn = con
	port.active = true
	port.key = key
	port.listener = listener

	if local {
		port.portType = "Switch"
		port.ipAndPort = con.RemoteAddr().String()
	} else {
		port.ipAndPort = con.LocalAddr().String()
	}

	port.rx = queues.NewByteSliceQueue("RX", maxInputQueueSize)
	port.tx = queues.NewByteSliceQueue("TX", maxOutputQueueSize)
	port.nx = queues.NewByteSliceQueue("NX", maxInputQueueSize)

	return port
}

// This is the method that the service port is using to connect to the switch for the VM/machine
func ConnectTo(host, key, secret string, destPort int, listener common.Listener, maxIn, maxOut, notifiers int) (common.Port, error) {

	// Dial the destination and validate the secret and key
	conn, err := common.ConnectAndValidateSecretAndKey(host, secret, key, destPort)
	if err != nil {
		return nil, err
	}

	// Instantiate the port
	port := NewPortImpl(false, conn, key, listener, maxIn, maxOut)

	//Below attributes are only for the port initiating the connection
	port.secret = secret
	port.host = host
	port.destPort = destPort
	port.reconnectMtx = &sync.Mutex{}

	// Request the connecting port to send over its uuid
	zuuid, err := common.ExchangeUuid(port.uuid, port.key, conn)
	if err != nil {
		return nil, err
	}
	port.zside = zuuid

	//We have only one go routing per each because we want to keep the order of incoming and outgoing messages

	// Start loop reading from the socket
	go port.readFromSocket()
	// Start loop reading from the TX queue and writing to the socket
	go port.writeToSocket()
	// Start loop decoding from RX queue
	go port.decodeIncomingData()
	// Start loop of reading decoded data from the NX queue and notify the listener
	go port.incomingDataNotifier()

	return port, nil
}

func (port *PortImpl) Uuid() string {
	return port.uuid
}

func (port *PortImpl) ZSide() string {
	return port.zside
}

func (port *PortImpl) Shutdown() {
	port.active = false
	if port.conn != nil {
		port.conn.Close()
	}
	port.rx.Shutdown()
	port.tx.Shutdown()
	port.nx.Shutdown()
	if port.listener != nil {
		port.listener.PortShutdown(port)
	}
}

func (port *PortImpl) attemptToReconnect() {
	port.reconnectMtx.Lock()
	defer port.reconnectMtx.Unlock()
	if port.doneReconnect {
		port.doneReconnect = false
		return
	}
	port.doneReconnect = true
	for {
		time.Sleep(time.Second * 5)
		logs.Info("Connection issues, trying to reconnect to switch")

		err := port.reconnect()
		if err == nil {
			break
		}

	}
	logs.Info("Reconnected!")
}

func (port *PortImpl) reconnect() error {
	// Dial the destination and validate the secret and key
	conn, err := common.ConnectAndValidateSecretAndKey(port.host, port.secret, port.key, port.destPort)
	if err != nil {
		return err
	}

	// Request the connecting port to send over its uuid
	zuuid, err := common.ExchangeUuid(port.uuid, port.key, conn)
	if err != nil {
		return err
	}

	port.zside = zuuid
	port.conn = conn

	return nil
}

func (port *PortImpl) Name() string {
	/*
		if port.serviceId == "" {
			port.serviceId = handlers.Directory.ServiceId(port.uuid)
		}*/
	name := strng.New("")
	name.Add(port.portType)
	name.Add("(")
	name.Add(port.serviceId)
	name.Add(" ")
	name.Add(port.uuid)
	name.Add("[")
	name.Add(port.ipAndPort)
	name.Add("]")
	name.Add(")")
	return name.String()
}

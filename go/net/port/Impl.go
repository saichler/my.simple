package port

import (
	"github.com/google/uuid"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/net/protocol"
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
	// The connection
	conn net.Conn
	// is the port active
	active bool
	// The incoming data listener
	dataListener common.DatatListener
	// The local/remote address, depending on if this port is the initiator of the connection
	addr string
	//port reconnect info, only valid if the port is the initiating side
	reconnectInfo *ReconnectInfo
	// Is this port belongs to the switch
	isSwitch bool
	// created at
	createdAt int64
}

type ReconnectInfo struct {
	//The host
	host string
	//The port
	port int32
	// Mutex as multiple go routines might call reconnect
	reconnectMtx *sync.Mutex
	// Indicates if the port was already reconnected
	alreadyReconnected bool
}

// Instantiate a new port with a connection
func NewPortImpl(incomingConnection bool, con net.Conn, key, secret, _uuid string, dataListener common.DatatListener) *PortImpl {
	port := &PortImpl{}
	port.uuid = _uuid
	port.createdAt = time.Now().Unix()
	if port.uuid == "" {
		port.uuid = uuid.New().String()
	}
	port.conn = con
	port.active = true
	port.key = key
	port.secret = secret
	port.dataListener = dataListener

	if incomingConnection {
		port.addr = con.RemoteAddr().String()
		port.isSwitch = true
	} else {
		port.addr = con.LocalAddr().String()
	}

	port.rx = queues.NewByteSliceQueue("RX", int(common.NetConfig.DefaultRxQueueSize))
	port.tx = queues.NewByteSliceQueue("TX", int(common.NetConfig.DefaultTxQueueSize))

	return port
}

// This is the method that the service port is using to connect to the switch for the VM/machine
func ConnectTo(host, key, secret string, destPort int32, datalistener common.DatatListener) (common.Port, error) {

	// Dial the destination and validate the secret and key
	conn, err := protocol.ConnectToAndValidateSecretAndKey(host, secret, key, destPort)
	if err != nil {
		return nil, err
	}

	// Instantiate the port
	port := NewPortImpl(false, conn, key, secret, "", datalistener)

	//Below attributes are only for the port initiating the connection
	port.secret = secret
	port.reconnectInfo = &ReconnectInfo{
		host:         host,
		port:         destPort,
		reconnectMtx: &sync.Mutex{},
	}

	// Request the connecting port to send over its uuid
	zuuid, err := protocol.ExchangeUuid(port.uuid, port.key, conn)
	if err != nil {
		return nil, err
	}
	port.zside = zuuid

	//We have only one go routing per each because we want to keep the order of incoming and outgoing messages

	port.Start()

	return port, nil
}

func (port *PortImpl) Start() {
	// Start loop reading from the socket
	go port.readFromSocket()
	// Start loop reading from the TX queue and writing to the socket
	go port.writeToSocket()
	// Start loop notifying the raw data listener on new incoming data
	go port.notifyRawDataListener()
	logs.Info(port.Name(), "Started!")
}

// Addr The address of this port, either remote addr or local
// depending on if this is the initiator of the connection
func (port *PortImpl) Addr() string {
	return port.addr
}

// The port uuid
func (port *PortImpl) Uuid() string {
	return port.uuid
}

func (port *PortImpl) ZSide() string {
	return port.zside
}

func (port *PortImpl) Shutdown() {
	logs.Info(port.Name(), "Shutdown called...")
	port.active = false
	if port.conn != nil {
		port.conn.Close()
	}
	port.rx.Shutdown()
	port.tx.Shutdown()

	if port.dataListener != nil {
		port.dataListener.PortShutdown(port)
	}
}

func (port *PortImpl) attemptToReconnect() {
	// Should not be a valid scenario, however bugs do happen
	if port.reconnectInfo == nil {
		return
	}

	port.reconnectInfo.reconnectMtx.Lock()
	defer port.reconnectInfo.reconnectMtx.Unlock()
	if port.reconnectInfo.alreadyReconnected {
		return
	}
	for {
		time.Sleep(time.Second * 5)
		logs.Info("Connection issues, trying to reconnect to switch")

		err := port.reconnect()
		if err == nil {
			port.reconnectInfo.alreadyReconnected = true
			go func() {
				time.Sleep(time.Second)
				port.reconnectInfo.alreadyReconnected = false
			}()
			break
		}

	}
	logs.Info("Reconnected!")
}

func (port *PortImpl) reconnect() error {
	// Dial the destination and validate the secret and key
	conn, err := protocol.ConnectToAndValidateSecretAndKey(port.reconnectInfo.host, port.secret, port.key, port.reconnectInfo.port)
	if err != nil {
		return err
	}

	// Request the connecting port to send over its uuid
	zuuid, err := protocol.ExchangeUuid(port.uuid, port.key, conn)
	if err != nil {
		return err
	}

	port.zside = zuuid
	port.conn = conn

	return nil
}

func (port *PortImpl) Name() string {
	name := strng.New("")
	if port.isSwitch {
		name.Add("Switch Port ")
	} else {
		name.Add("Node Port ")
	}
	name.Add(port.uuid)
	name.Add("[")
	name.Add(port.addr)
	name.Add("]")
	return name.String()
}

func (port *PortImpl) CreatedAt() int64 {
	return port.createdAt
}

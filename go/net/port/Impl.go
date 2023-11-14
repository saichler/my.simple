package port

import (
	"github.com/google/uuid"
	"github.com/saichler/my.security/go/sec_common"
	"github.com/saichler/my.simple/go/common"
	model2 "github.com/saichler/my.simple/go/net/model"
	"github.com/saichler/my.simple/go/services/health"
	"github.com/saichler/my.simple/go/services/health/model"
	"github.com/saichler/my.simple/go/utils/logs"
	"github.com/saichler/my.simple/go/utils/queues"
	"github.com/saichler/my.simple/go/utils/strng"
	"net"
	"sync"
	"time"
)

type PortImpl struct {
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
	//The used registry
	registry common.IRegistry
	//Service Points
	servicePoints common.IServicePoints
}

type ReconnectInfo struct {
	//The host
	host string
	//The port
	port uint32
	// Mutex as multiple go routines might call reconnect
	reconnectMtx *sync.Mutex
	// Indicates if the port was already reconnected
	alreadyReconnected bool
}

// Instantiate a new port with a connection
func NewPortImpl(incomingConnection bool, con net.Conn, uid string, dataListener common.DatatListener, registry common.IRegistry, servicePoints common.IServicePoints) *PortImpl {
	port := &PortImpl{}
	port.uuid = uid
	port.registry = registry
	port.servicePoints = servicePoints
	port.createdAt = time.Now().Unix()
	port.conn = con
	port.active = true
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
func ConnectTo(host string, destPort uint32, datalistener common.DatatListener, registry common.IRegistry, servicePoints common.IServicePoints) (common.Port, error) {

	// Dial the destination and validate the secret and key
	conn, err := sec_common.MySecurityProvider.CanDial(host, destPort)
	if err != nil {
		return nil, err
	}

	auuid := uuid.New().String()
	zuuid, err := sec_common.MySecurityProvider.ValidateConnection(conn, auuid)

	port := NewPortImpl(false, conn, auuid, datalistener, registry, servicePoints)

	//Below attributes are only for the port initiating the connection
	port.reconnectInfo = &ReconnectInfo{
		host:         host,
		port:         destPort,
		reconnectMtx: &sync.Mutex{},
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

	go port.reportHealthStatus()
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
		logs.Warning("Connection issues, trying to reconnect to switch")

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
	conn, err := sec_common.MySecurityProvider.CanDial(port.reconnectInfo.host, port.reconnectInfo.port)
	if err != nil {
		return err
	}
	zuuid, err := sec_common.MySecurityProvider.ValidateConnection(conn, port.uuid)
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

func (port *PortImpl) reportHealthStatus() {
	for port.active {
		time.Sleep(time.Second * 5)
		report := health.CreateReport(port.uuid, model.HealthStatus_Health_Live)
		port.Do(model2.Action_Action_Post, health.Health_Center_Topic, report)
	}
}

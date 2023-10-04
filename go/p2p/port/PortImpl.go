package port

import (
	"errors"
	"github.com/google/uuid"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/utils/logs"
	"github.com/saichler/my.simple/go/utils/queues"
	"github.com/saichler/my.simple/go/utils/security"
	"github.com/saichler/my.simple/go/utils/strng"
	"net"
	"strconv"
	"sync"
	"time"
)

type PortImpl struct {
	key        string
	secret     string
	uuid       string
	zside      string
	rx         *queues.ByteSliceQueue
	tx         *queues.ByteSliceQueue
	nx         *queues.ByteSliceQueue
	writeMutex *sync.Cond
	conn       net.Conn
	active     bool
	listener   common.Listener
	ipAndPort  string
	portType   string
	serviceId  string

	host          string
	destPort      int
	reconnectMtx  *sync.Mutex
	doneReconnect bool
}

func NewPortImpl(local bool, con net.Conn, key string, listener common.Listener, maxInputQueueSize, maxOutputQueueSize int) *PortImpl {
	port := &PortImpl{}
	port.uuid = uuid.New().String()
	port.conn = con
	if local {
		port.portType = "Switch"
		port.ipAndPort = con.RemoteAddr().String()
	} else {
		port.ipAndPort = con.LocalAddr().String()
	}
	port.rx = queues.NewByteSliceQueue("RX", maxInputQueueSize)
	port.tx = queues.NewByteSliceQueue("TX", maxOutputQueueSize)
	port.nx = queues.NewByteSliceQueue("NX", maxInputQueueSize)
	port.active = true
	port.key = key
	port.listener = listener
	port.writeMutex = sync.NewCond(&sync.Mutex{})
	return port
}

func ConnectTo(host, key, secret string, destPort int, listener common.Listener, maxIn, maxOut, notifiers int) (common.Port, error) {
	conn, err := net.Dial("tcp", host+":"+strconv.Itoa(destPort))
	if err != nil {
		return nil, err
	}

	data, err := security.Encode([]byte(secret), key)
	if err != nil {
		return nil, err
	}

	err = common.Write([]byte(data), conn)
	if err != nil {
		return nil, err
	}

	inData, err := common.Read(conn)
	if string(inData) != "OK" {
		return nil, errors.New("Failed to connect, incorrect Key/Secret")
	}

	port := NewPortImpl(false, conn, key, listener, maxIn, maxOut)
	port.secret = secret
	port.host = host
	port.destPort = destPort
	port.reconnectMtx = &sync.Mutex{}

	data, err = security.Encode([]byte(port.uuid), port.key)
	if err != nil {
		return nil, err
	}

	err = common.Write([]byte(data), conn)
	if err != nil {
		return nil, err
	}

	inData, err = common.Read(conn)
	port.zside = string(inData)

	// Start loop reading from the socker
	go port.readFromSocket()
	go port.write()
	// Start loop decoding from RX queue
	go port.decodeIncomingData()
	if notifiers <= 0 {
		go port.notifier()
	} else {
		for i := 0; i < notifiers; i++ {
			go port.notifier()
		}
	}
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
	port.writeMutex.Broadcast()
	if port.listener != nil {
		port.listener.PortShutdown(port)
	}
}

func (port *PortImpl) notifier() {
	for port.active {
		data := port.nx.Next()
		if data != nil && port.listener != nil {
			port.listener.DataReceived(data, port)
		} else if data != nil {
			logs.Info("No Data Listener for packet:", string(data))
		}
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
	conn, err := net.Dial("tcp", port.host+":"+strconv.Itoa(port.destPort))
	if err != nil {
		return logs.Error("Unable to reconnect to switch... ", err.Error())
	}

	data, err := security.Encode([]byte(port.secret), port.key)
	if err != nil {
		return logs.Error("Unable to encode when reconnecting to switch...", err.Error())
	}

	err = common.Write([]byte(data), conn)
	if err != nil {
		return logs.Error("Unable to write when reconnecting to switch...", err.Error())
	}

	inData, err := common.Read(conn)
	if string(inData) != "OK" {
		return logs.Error("Failed to reconnect, incorrect Key/Secret")
	}

	data, err = security.Encode([]byte(port.uuid), port.key)
	if err != nil {
		return err
	}

	err = common.Write([]byte(data), conn)
	if err != nil {
		return err
	}

	inData, err = common.Read(conn)
	port.zside = string(inData)

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

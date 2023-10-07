package switching

import (
	"errors"
	"github.com/google/uuid"
	"github.com/saichler/my.simple/go/common"
	port2 "github.com/saichler/my.simple/go/net/port"
	"github.com/saichler/my.simple/go/net/protocol"
	"github.com/saichler/my.simple/go/utils/logs"
	"net"
	"strconv"
)

type SwitchService struct {
	uuid        string
	port        int
	key         string
	secret      string
	socket      net.Listener
	active      bool
	switchTable *SwitchTable
}

func newSwitchService(key, secret string, port int) *SwitchService {
	switchService := &SwitchService{}
	switchService.uuid = uuid.New().String()
	switchService.key = key
	switchService.secret = secret
	switchService.port = port
	switchService.switchTable = newSwitchTable()
	return switchService
}

func (switchService *SwitchService) start() error {
	if switchService.port == 0 {
		return errors.New("Switch Port does not have a port defined")
	}
	if switchService.secret == "" {
		return errors.New("Switch Port does not have a secret")
	}
	if switchService.key == "" {
		return errors.New("Switch Port does not have a key")
	}

	err := switchService.bind()
	if err != nil {
		return err
	}

	for switchService.active {
		logs.Info("Waiting for connections...")
		conn, e := switchService.socket.Accept()
		if e != nil {
			logs.Error("Failed to accept socket connection:", err)
			continue
		}
		logs.Info("Accepted socket connection...")
		go switchService.connect(conn)
	}
	logs.Info("Switch Service has ended")
	return nil
}

func (switchService *SwitchService) bind() error {
	socket, e := net.Listen("tcp", ":"+strconv.Itoa(switchService.port))
	if e != nil {
		return logs.Error("Unable to bind to port ", switchService.port, e.Error())
	}
	logs.Info("Bind Successfully to port ", switchService.port)
	switchService.socket = socket
	return nil
}

func (switchService *SwitchService) connect(conn net.Conn) {
	uuid, err := protocol.Incoming(conn, switchService.key, switchService.secret, switchService.uuid)
	if err != nil {
		logs.Error("Failed to connect:", err.Error())
		return
	}
	port := port2.NewPortImpl(true, conn, switchService.key, switchService.secret, uuid, switchService)
	port.Start()
	switchService.switchTable.addPort(port)
}

func (switchService *SwitchService) DataReceived(data []byte, port common.Port) {

}

func (switchService *SwitchService) PortShutdown(port common.Port) {

}

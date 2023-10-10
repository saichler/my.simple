package switching

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/saichler/my.simple/go/common"
	port2 "github.com/saichler/my.simple/go/net/port"
	"github.com/saichler/my.simple/go/net/protocol"
	"github.com/saichler/my.simple/go/services/health"
	"github.com/saichler/my.simple/go/types"
	"github.com/saichler/my.simple/go/utils/logs"
	"google.golang.org/protobuf/proto"
	"net"
	"strconv"
)

type SwitchService struct {
	uuid        string
	port        int32
	key         string
	secret      string
	socket      net.Listener
	active      bool
	switchTable *SwitchTable
}

func NewSwitchService(key, secret string, port int32) *SwitchService {
	switchService := &SwitchService{}
	switchService.uuid = uuid.New().String()
	switchService.key = key
	switchService.secret = secret
	switchService.port = port
	switchService.switchTable = newSwitchTable()
	switchService.active = true
	return switchService
}

func (switchService *SwitchService) Start() error {
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
		if e != nil && switchService.active {
			logs.Error("Failed to accept socket connection:", err)
			continue
		}
		if switchService.active {
			logs.Info("Accepted socket connection...")
			go switchService.connect(conn)
		}
	}
	logs.Warning("Switch Service has ended")
	return nil
}

func (switchService *SwitchService) bind() error {
	socket, e := net.Listen("tcp", ":"+strconv.Itoa(int(switchService.port)))
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
	switchService.notifyNewPort(port)
}

func (switchService *SwitchService) notifyNewPort(port common.Port) {
	go switchService.switchTable.addPort(port, switchService.key, switchService.uuid)
}

func (switchService *SwitchService) Shutdown() {
	switchService.active = false
	switchService.socket.Close()
}

func (switchService *SwitchService) HandleData(data []byte, port common.Port) {
	source, destination, pri := protocol.HeaderOf(data)
	fmt.Println(source, destination, pri.String())
	//The destination is the switch
	if destination == switchService.uuid {
		switchService.switchDataReceived(data, port)
		return
	}

	uuidList := health.ServiceUuids(destination)
	if uuidList != nil {
		switchService.sendToPorts(uuidList, data)
		return
	}

	//The destination is a single port
	p := switchService.switchTable.fetchPortByUuid(destination)
	if p == nil {
		logs.Error("Cannot find destination port for ", destination)
		return
	}
	p.Send(data)
}

func (switchService *SwitchService) sendToPorts(uuids []string, data []byte) {
	for _, uuid := range uuids {
		port := switchService.switchTable.fetchPortByUuid(uuid)
		if port != nil {
			port.Send(data)
		}
	}
}

func (switchService *SwitchService) publish(pb proto.Message) {

}

func (switchService *SwitchService) PortShutdown(port common.Port) {
}

func (switchService *SwitchService) switchDataReceived(data []byte, port common.Port) {
	msg, err := protocol.MessageOf(data)
	if err != nil {
		logs.Error(err)
		return
	}
	pb, err := protocol.ProtoOf(msg, switchService.key)
	if err != nil {
		logs.Error(err)
		return
	}
	// Otherwise call the handler per the action & the type
	types.Handle(pb, msg.Action, port)

}

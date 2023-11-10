package switching

import (
	"fmt"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/net/model"
	"github.com/saichler/my.simple/go/net/protocol"
	"github.com/saichler/my.simple/go/services/health"
	"github.com/saichler/my.simple/go/utils/logs"
	"github.com/saichler/my.simple/go/utils/maps"
	"google.golang.org/protobuf/proto"
)

type SwitchTable struct {
	internalPorts *maps.PortMap
	externalPorts *maps.PortMap
	health        common.IHealthCeter
}

func newSwitchTable(health common.IHealthCeter) *SwitchTable {
	switchTable := &SwitchTable{}
	switchTable.internalPorts = maps.NewPortMap()
	switchTable.externalPorts = maps.NewPortMap()
	switchTable.health = health
	return switchTable
}

func (switchTable *SwitchTable) allPortsList() []common.Port {
	ports := make([]common.Port, 0)
	switchTable.internalPorts.Iterate(func(k, v interface{}) {
		ports = append(ports, v.(common.Port))
	})
	switchTable.externalPorts.Iterate(func(k, v interface{}) {
		ports = append(ports, v.(common.Port))
	})
	return ports
}

func (switchTable *SwitchTable) broadcast(topic string, action model.Action, switchUuid string, pb proto.Message) {
	fmt.Println("Broadcast")
	ports := switchTable.allPortsList()
	data, err := protocol.CreateMessageFor(model.Priority_P0, action, switchUuid, topic, pb)
	if err != nil {
		logs.Error("Failed to send broadcast:", err)
		return
	}
	for _, port := range ports {
		port.Send(data)
	}
}

func (switchTable *SwitchTable) addPort(port common.Port, switchUuid string) {
	//check if this port is local to the machine, e.g. not belong to public subnet
	isLocal := ipSegment.isLocal(port.Addr())
	// If it is local, add it to the internal map
	if isLocal {
		//check if the port already exist
		ep, ok := switchTable.internalPorts.Get(port.Uuid())
		if ok {
			//If it exists, then shutdown the existing instance as we want the new one to be used.
			ep.Shutdown()
		}
		switchTable.internalPorts.Put(port.Uuid(), port)
	} else {
		// If it is public, add it to the external map
		// but first check if it already exists
		ep, ok := switchTable.externalPorts.Get(port.Uuid())
		if ok {
			//if it already exists, shut it down.
			ep.Shutdown()
		}
		switchTable.externalPorts.Put(port.Uuid(), port)
	}
	switchTable.health.AddPort(port)
	switchTable.health.AddService(health.Health_Center_Topic, port.Uuid())
	go switchTable.broadcast(health.Health_Center_Topic, model.Action_Action_Post, switchUuid, switchTable.health.Clone())
}

func (switchTable *SwitchTable) fetchPortByUuid(id string) common.Port {
	p, ok := switchTable.internalPorts.Get(id)
	if !ok {
		p, ok = switchTable.externalPorts.Get(id)
	}
	return p
}

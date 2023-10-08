package switching

import (
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/net/model"
	"github.com/saichler/my.simple/go/net/protocol"
	"github.com/saichler/my.simple/go/utils/logs"
	"google.golang.org/protobuf/proto"
	"sync"
)

type SwitchTable struct {
	internalPorts map[string]common.Port
	externalPorts map[string]common.Port
	mtx           *sync.RWMutex
}

func newSwitchTable() *SwitchTable {
	switchTable := &SwitchTable{}
	switchTable.mtx = &sync.RWMutex{}
	switchTable.internalPorts = make(map[string]common.Port)
	switchTable.externalPorts = make(map[string]common.Port)
	return switchTable
}

func (switchTable *SwitchTable) allPortsList() []common.Port {
	switchTable.mtx.RLock()
	defer switchTable.mtx.RUnlock()
	ports := make([]common.Port, 0)
	for _, port := range switchTable.internalPorts {
		ports = append(ports, port)
	}
	for _, port := range switchTable.externalPorts {
		ports = append(ports, port)
	}
	return ports
}

func (switchTable *SwitchTable) broadcast(action model.Action, key string, pb proto.Message) {
	ports := switchTable.allPortsList()
	data, err := protocol.CreateMessageFor(model.Priority_P0, action, key, "_b", "_b", pb)
	if err != nil {
		logs.Error("Failed to send broadcast:", err)
		return
	}
	for _, port := range ports {
		port.Send(data)
	}
}

func (switchTable *SwitchTable) addPort(port common.Port) {
	switchTable.mtx.Lock()
	defer switchTable.mtx.Unlock()
	//check if this port is local to the machine, e.g. not belong to public subnet
	isLocal := ipSegment.isLocal(port.Addr())
	// If it is local, add it to the internal map
	if isLocal {
		//check if the port already exist
		ep, ok := switchTable.internalPorts[port.Uuid()]
		if ok {
			//If it exists, then shutdown the existing instance as we want the new one to be used.
			ep.Shutdown()
		}
		switchTable.internalPorts[port.Uuid()] = port
	} else {
		// If it is public, add it to the external map
		// but first check if it already exists
		ep, ok := switchTable.externalPorts[port.Uuid()]
		if ok {
			//if it already exists, shut it down.
			ep.Shutdown()
		}
		switchTable.externalPorts[port.Uuid()] = port
	}
}

func (switchTable *SwitchTable) fetchPortByUuid(id string) common.Port {
	switchTable.mtx.RLock()
	defer switchTable.mtx.RUnlock()
	p, ok := switchTable.internalPorts[id]
	if !ok {
		p = switchTable.externalPorts[id]
	}
	return p
}

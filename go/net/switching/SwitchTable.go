package switching

import (
	"github.com/saichler/my.simple/go/common"
	"sync"
)

type SwitchTable struct {
	internalPorts map[string]common.Port
	externalPorts map[string]common.Port
	mtx           *sync.Mutex
}

func newSwitchTable() *SwitchTable {
	switchTable := &SwitchTable{}
	switchTable.mtx = &sync.Mutex{}
	switchTable.internalPorts = make(map[string]common.Port)
	switchTable.externalPorts = make(map[string]common.Port)
	return switchTable
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

package health

import (
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/services/health/model"
	"github.com/saichler/my.simple/go/types"
	"google.golang.org/protobuf/proto"
)

var serviceCenter = newServiceCenter()

const (
	Service_Center_Topic = "Service Center"
)

type ServiceCenter struct {
}

func newServiceCenter() *ServiceCenter {
	sc := &ServiceCenter{}
	types.RegisterTypeHandler(&model.Report{}, sc)
	return sc
}

func (h *ServiceCenter) Post(pb proto.Message, port common.Port) (proto.Message, error) {
	return nil, nil
}

func (h *ServiceCenter) Put(pb proto.Message, port common.Port) (proto.Message, error) {
	return nil, nil
}

func (h *ServiceCenter) Patch(pb proto.Message, port common.Port) (proto.Message, error) {
	return nil, nil
}

func (h *ServiceCenter) Delete(pb proto.Message, port common.Port) (proto.Message, error) {
	return nil, nil
}

func (h *ServiceCenter) Get(pb proto.Message, port common.Port) (proto.Message, error) {
	return nil, nil
}

func AssignServiceTopicToProvider(uuid, topic string) {
	healthCenter.mtx.L.Lock()
	defer healthCenter.mtx.L.Unlock()
	var port *model.Port
	for uid, p := range healthCenter.health.Ports {
		if uid == uuid {
			port = p
			break
		}
	}
	if port != nil {
		service, ok := healthCenter.health.Services[topic]
		if !ok {
			service = &model.Service{}
			service.PortUuids = make([]string, 0)
			healthCenter.health.Services[topic] = service
		}
		service.PortUuids = append(service.PortUuids, uuid)
	}
}

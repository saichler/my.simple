package health

import (
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/services/health/model"
	"github.com/saichler/my.simple/go/types"
	"google.golang.org/protobuf/proto"
	"time"
)

var serviceCenter = newServiceCenter()

const (
	Service_Center_Topic = "Service Center"
)

type ServiceCenter struct {
}

func newServiceCenter() *ServiceCenter {
	sc := &ServiceCenter{}
	types.RegisterTypeHandler(&model.ServiceProvider{}, sc)
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

	var nodeHealth *model.NodesHealth
	for uid, nh := range healthCenter.health.Nodes {
		if uid == uuid {
			nodeHealth = nh
			break
		}
	}
	if nodeHealth == nil {
		nodeHealth = &model.NodesHealth{}
		nodeHealth.CreatedAt = time.Now().Unix()
		nodeHealth.PortUuid = uuid
		nodeHealth.Report = &model.HealthReport{}
		nodeHealth.Services = make(map[string]bool)
		healthCenter.health.Nodes[uuid] = nodeHealth
	}

	nodeHealth.Services[topic] = true
	nodeHealth.Status = model.HealthStatus_Health_Live

	providers, ok := healthCenter.health.Providers[topic]
	if !ok {
		providers = &model.ServiceProviders{}
		providers.ProvidersUuids = make([]string, 0)
		healthCenter.health.Providers[topic] = providers
	}
	providers.ProvidersUuids = append(providers.ProvidersUuids, uuid)
}

/*
func (h *ServiceCenter) service(serviceTopic string) string {
	ServiceCenter.mtx.L.Lock()
	defer ServiceCenter.mtx.L.Unlock()
	providers := h.health.Services[serviceTopic]
	if providers == nil {
		return ""
	}
	if providers.ProvidersUuids == nil || len(providers.ProvidersUuids) == 0 {
		return ""
	}
	return providers.ProvidersUuids[0]
}*/

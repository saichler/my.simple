package health

import (
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/services/health/model"
	"github.com/saichler/my.simple/go/types"
	"github.com/saichler/my.simple/go/utils/logs"
	"google.golang.org/protobuf/proto"
	"sync"
	"time"
)

type HealthCenter struct {
	health *model.HealthCenter
	mtx    *sync.Cond
}

var healthCenter = newHealthCenter()

const (
	Health_Center_Topic = "Healh Center"
)

func newHealthCenter() *HealthCenter {
	hc := &HealthCenter{}
	hc.mtx = sync.NewCond(&sync.Mutex{})
	hc.health = &model.HealthCenter{}
	hc.health.Nodes = make(map[string]*model.NodesHealth)
	hc.health.Services = make(map[string]*model.NodeServices)
	types.Types.RegisterTypeHandler(hc.health, hc)
	return hc
}

func (h *HealthCenter) Post(pb proto.Message, port common.Port) (proto.Message, error) {
	logs.Info("Health Center Report ", port.Name())
	other := pb.(*model.HealthCenter)
	h.mtx.L.Lock()
	defer h.mtx.L.Unlock()

	h.health.LastReportTime = other.LastReportTime
	for k, v := range other.Nodes {
		h.health.Nodes[k] = v
	}
	return nil, nil
}

func (h *HealthCenter) Put(pb proto.Message, port common.Port) (proto.Message, error) {
	return nil, nil
}

func (h *HealthCenter) Patch(pb proto.Message, port common.Port) (proto.Message, error) {
	return nil, nil
}

func (h *HealthCenter) Delete(pb proto.Message, port common.Port) (proto.Message, error) {
	return nil, nil
}

func (h *HealthCenter) Get(pb proto.Message, port common.Port) (proto.Message, error) {
	return nil, nil
}

func CloneHealth() *model.HealthCenter {
	healthCenter.mtx.L.Lock()
	defer healthCenter.mtx.L.Unlock()
	clone := &model.HealthCenter{}
	clone.LastReportTime = time.Now().Unix()
	clone.Nodes = cloneNodesHealthMap()
	clone.Services = cloneServicesMap()
	return clone
}

func cloneNodesHealthMap() map[string]*model.NodesHealth {
	cloneMap := make(map[string]*model.NodesHealth)
	for k, v := range healthCenter.health.Nodes {
		cloneMap[k] = cloneNodesHealth(v)
	}
	return cloneMap
}

func cloneNodesHealth(nh *model.NodesHealth) *model.NodesHealth {
	clone := &model.NodesHealth{}
	clone.CreatedAt = nh.CreatedAt
	clone.PortUuid = nh.PortUuid
	clone.Services = make(map[string]bool)
	if nh.Services != nil {
		for k, v := range nh.Services {
			clone.Services[k] = v
		}
	}
	return clone
}

func cloneServicesMap() map[string]*model.NodeServices {
	cloneMap := make(map[string]*model.NodeServices)
	if healthCenter.health.Services != nil {
		for k, v := range healthCenter.health.Services {
			cloneMap[k] = cloneNodeServices(v)
		}
	}
	return cloneMap
}

func cloneNodeServices(ns *model.NodeServices) *model.NodeServices {
	clone := &model.NodeServices{}
	clone.ProvidersUuids = make([]string, 0)
	for _, v := range ns.ProvidersUuids {
		clone.ProvidersUuids = append(clone.ProvidersUuids, v)
	}
	return clone
}

func (h *HealthCenter) service(serviceTopic string) string {
	healthCenter.mtx.L.Lock()
	defer healthCenter.mtx.L.Unlock()
	providers := h.health.Services[serviceTopic]
	if providers == nil {
		return ""
	}
	if providers.ProvidersUuids == nil || len(providers.ProvidersUuids) == 0 {
		return ""
	}

	return providers.ProvidersUuids[0]
}

package health

import (
	"github.com/saichler/my.simple/go/common"
	model2 "github.com/saichler/my.simple/go/net/model"
	"github.com/saichler/my.simple/go/services/health/model"
	"github.com/saichler/my.simple/go/types"
	"github.com/saichler/my.simple/go/utils/logs"
	"google.golang.org/protobuf/proto"
	"sync"
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
	hc.health.Providers = make(map[string]*model.ServiceProviders)
	types.RegisterTypeHandler(hc.health, hc)
	return hc
}

func (h *HealthCenter) Post(pb proto.Message, port common.Port) (proto.Message, error) {
	logs.Debug("Health Center Report port:", port.Name())
	other := pb.(*model.HealthCenter)
	h.mtx.L.Lock()
	defer h.mtx.L.Unlock()

	h.health.LastReportTime = other.LastReportTime
	for k, v := range other.Nodes {
		logs.Debug("    ", port.Name(), " -> ", v.PortUuid)
		h.health.Nodes[k] = v
	}
	if h.health.Providers == nil {
		h.health.Providers = make(map[string]*model.ServiceProviders)
	}
	for k, v := range other.Providers {
		logs.Debug("    ", k, " -> ")
		for _, uuid := range v.ProvidersUuids {
			logs.Debug("       ", uuid)
		}
		h.health.Providers[k] = v
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
	health := CloneHealth()
	port.Do(model2.Action_Action_Post, port.Uuid(), health)
	return nil, nil
}

package health

import (
	"github.com/saichler/my.simple/go/common"
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

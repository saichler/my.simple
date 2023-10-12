package health

import (
	"github.com/saichler/my.simple/go/common"
	model2 "github.com/saichler/my.simple/go/net/model"
	"github.com/saichler/my.simple/go/services/health/model"
	"github.com/saichler/my.simple/go/services/service_point"
	"github.com/saichler/my.simple/go/utils/logs"
	"google.golang.org/protobuf/proto"
	"runtime"
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
	hc.health.Ports = make(map[string]*model.Port)
	hc.health.Services = make(map[string]*model.Service)
	hc.health.Reports = make(map[string]*model.Report)
	service_point.RegisterServicePoint(hc.health, hc)
	return hc
}

func (h *HealthCenter) Post(pb proto.Message, port common.Port) (proto.Message, error) {
	logs.Debug("Health Center Report port:", port.Name())
	other := pb.(*model.HealthCenter)
	h.mtx.L.Lock()
	defer h.mtx.L.Unlock()

	for k, v := range other.Ports {
		logs.Debug("     Port:", k, " -> ", v.PortUuid)
		h.health.Ports[k] = v
	}
	for k, v := range other.Services {
		logs.Debug("     Service:", k, " -> ")
		for _, uuid := range v.PortUuids {
			logs.Debug("       ", uuid)
		}
		h.health.Services[k] = v
	}
	for k, v := range other.Reports {
		logs.Debug("      Reports:", k, " ->", v.PortUuid)
		h.health.Reports[k] = v
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

func AddPort(port common.Port) {
	p, ok := healthCenter.health.Ports[port.Uuid()]
	if !ok {
		p = &model.Port{}
		p.PortUuid = port.Uuid()
		p.CreatedAt = time.Now().Unix()
		healthCenter.health.Ports[port.Uuid()] = p
	}
}

func CreateReport(portUuid string, stat model.HealthStatus) *model.Report {
	report := &model.Report{}
	report.PortUuid = portUuid
	report.ReportTime = time.Now().Unix()
	report.Status = stat
	mem := &runtime.MemStats{}
	report.MemoryUsage = mem.TotalAlloc - mem.Frees
	return report
}

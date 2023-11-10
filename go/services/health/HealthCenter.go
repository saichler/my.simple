package health

import (
	"github.com/saichler/my.simple/go/common"
	model2 "github.com/saichler/my.simple/go/net/model"
	"github.com/saichler/my.simple/go/services/health/model"
	"github.com/saichler/my.simple/go/utils/logs"
	"google.golang.org/protobuf/proto"
	"runtime"
	"sync"
	"time"
)

type HealthCenter struct {
	health       *model.HealthCenter
	mtx          *sync.Cond
	introspect   common.IIntrospect
	servicePoint common.IServicePoints
}

const (
	Health_Center_Topic = "Healh Center"
)

func NewHealthCenter(introspect common.IIntrospect, servicePoints common.IServicePoints) *HealthCenter {
	hc := &HealthCenter{}
	hc.mtx = sync.NewCond(&sync.Mutex{})
	hc.introspect = introspect
	hc.servicePoint = servicePoints
	hc.health = &model.HealthCenter{}
	hc.health.Ports = make(map[string]*model.Port)
	hc.health.Services = make(map[string]*model.Service)
	hc.health.Reports = make(map[string]*model.Report)
	hc.servicePoint.RegisterServicePoint(hc.health, hc, hc.introspect.Registry())
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
	health := h.Clone()
	port.Do(model2.Action_Action_Post, port.Uuid(), health)
	return nil, nil
}

func (h *HealthCenter) EndPoint() string {
	return "/health"
}

func (h *HealthCenter) AddPort(port common.Port) {
	h.mtx.L.Lock()
	defer h.mtx.L.Unlock()
	p, ok := h.health.Ports[port.Uuid()]
	if !ok {
		p = &model.Port{}
		p.PortUuid = port.Uuid()
		p.CreatedAt = time.Now().Unix()
		h.health.Ports[port.Uuid()] = p
	}
}

func (h *HealthCenter) ApplyReport(report *model.Report) {
	h.mtx.L.Lock()
	defer h.mtx.L.Unlock()
	logs.Debug("Received report from ", report.PortUuid)
	h.health.Reports[report.PortUuid] = report
}

func (h *HealthCenter) AddService(topic, uuid string) {
	h.mtx.L.Lock()
	defer h.mtx.L.Unlock()
	var port *model.Port
	for uid, p := range h.health.Ports {
		if uid == uuid {
			port = p
			break
		}
	}
	if port != nil {
		service, ok := h.health.Services[topic]
		if !ok {
			service = &model.Service{}
			service.PortUuids = make([]string, 0)
			h.health.Services[topic] = service
		}
		service.PortUuids = append(service.PortUuids, uuid)
	}
}

func (h *HealthCenter) ServiceUuids(topic string) []string {
	h.mtx.L.Lock()
	defer h.mtx.L.Unlock()
	uuids := h.health.Services[topic]
	if uuids == nil {
		return nil
	}
	result := make([]string, len(uuids.PortUuids))
	for i, v := range uuids.PortUuids {
		result[i] = v
	}
	return result
}

func (h *HealthCenter) Clone() *model.HealthCenter {
	h.mtx.L.Lock()
	defer h.mtx.L.Unlock()
	clone := h.introspect.Clone(h.health)
	return clone.(*model.HealthCenter)
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

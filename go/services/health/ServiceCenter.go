package health

import (
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/services/health/model"
	"google.golang.org/protobuf/proto"
)

const (
	Service_Center_Topic = "Service Center"
)

type ServiceCenter struct {
	registry      common.IRegistry
	health        common.IHealthCeter
	servicePoints common.IServicePoints
}

func NewServiceCenter(registry common.IRegistry, health common.IHealthCeter, servicePoints common.IServicePoints) *ServiceCenter {
	sc := &ServiceCenter{}
	sc.registry = registry
	sc.health = health
	sc.servicePoints = servicePoints
	sc.servicePoints.RegisterServicePoint(&model.Report{}, sc, sc.registry)
	return sc
}

func (h *ServiceCenter) Post(pb proto.Message, port common.Port) (proto.Message, error) {
	report := pb.(*model.Report)
	h.health.ApplyReport(report)
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

func (h *ServiceCenter) EndPoint() string {
	return "/service-center"
}

func (h *ServiceCenter) AddService(topic, uuid string) {
	h.health.AddService(topic, uuid)
}

func (h *ServiceCenter) ServiceUuids(topic string) []string {
	return h.health.ServiceUuids(topic)
}

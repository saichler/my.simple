package health

import (
	"github.com/saichler/my.simple/go/services/health/model"
)

func CloneHealth() *model.HealthCenter {
	healthCenter.mtx.L.Lock()
	defer healthCenter.mtx.L.Unlock()
	clone := &model.HealthCenter{}
	clone.Ports = clonePortsMap()
	clone.Services = cloneServicesMap()
	clone.Reports = cloneReportsMap()
	return clone
}

func clonePortsMap() map[string]*model.Port {
	cloneMap := make(map[string]*model.Port)
	for k, v := range healthCenter.health.Ports {
		cloneMap[k] = clonePort(v)
	}
	return cloneMap
}

func clonePort(nh *model.Port) *model.Port {
	clone := &model.Port{}
	clone.CreatedAt = nh.CreatedAt
	clone.PortUuid = nh.PortUuid
	return clone
}

func cloneServicesMap() map[string]*model.Service {
	cloneMap := make(map[string]*model.Service)
	for k, v := range healthCenter.health.Services {
		cloneMap[k] = cloneService(v)
	}
	return cloneMap
}

func cloneService(ns *model.Service) *model.Service {
	clone := &model.Service{}
	clone.PortUuids = make([]string, 0)
	for _, v := range ns.PortUuids {
		clone.PortUuids = append(clone.PortUuids, v)
	}
	return clone
}

func cloneReportsMap() map[string]*model.Report {
	cloneMap := make(map[string]*model.Report)
	for k, v := range healthCenter.health.Reports {
		cloneMap[k] = cloneReport(v)
	}
	return cloneMap
}

func cloneReport(report *model.Report) *model.Report {
	clone := &model.Report{}
	clone.ReportTime = report.ReportTime
	clone.MemoryUsage = report.MemoryUsage
	clone.PortUuid = report.PortUuid
	clone.Status = report.Status
	return clone
}

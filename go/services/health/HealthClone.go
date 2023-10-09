package health

import (
	"github.com/saichler/my.simple/go/services/health/model"
	"time"
)

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

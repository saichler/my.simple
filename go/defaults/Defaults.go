package defaults

import (
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/introspect"
	"github.com/saichler/my.simple/go/services/health"
	"github.com/saichler/my.simple/go/services/service_point"
	"github.com/saichler/my.simple/go/utils/registry"
)

func ApplyDefaults() {
	common.Registry = registry.NewStructRegistry()
	common.ServicePoints = service_point.NewServicePoints()
	common.Introspect = introspect.NewIntrospect(common.Registry)
	common.HealthCenter = health.NewHealthCenter(common.Introspect, common.ServicePoints)
	common.ServiceCenter = health.NewServiceCenter(common.Registry, common.HealthCenter, common.ServicePoints)
}

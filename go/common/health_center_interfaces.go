package common

import "github.com/saichler/my.simple/go/services/health/model"

type IHealthCeter interface {
	AddPort(Port)
	ApplyReport(*model.Report)
	AddService(string, string)
	ServiceUuids(string) []string
	Clone() *model.HealthCenter
}

var HealthCenter IHealthCeter

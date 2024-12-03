package adapters

import (
	"bedatabase/dgraphdb"

	"go.uber.org/fx"
)

type healthCheckAdapter struct {
	*dgraphdb.BeDgraphDB
}

type HealthCheckAdapter interface {
	GetDatabaseConnectionStatus(id string)
}

type HealthCheckAdapterParam struct {
	fx.In

	DgraphDB *dgraphdb.BeDgraphDB `name:"dgraphdb"`
}

func NewHealthCheckAdapter(p HealthCheckAdapterParam) HealthCheckAdapter {
	return healthCheckAdapter{p.DgraphDB}
}

func (e healthCheckAdapter) GetDatabaseConnectionStatus(id string) {

	e.Query(id)
}

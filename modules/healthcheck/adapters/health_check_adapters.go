package adapters

import (
	"bedatabase"
)

type HealthCheckAdapter interface {
	GetDatabaseConnectionStatus(id string)
}

func NewHealthCheckAdapter(db bedatabase.BeDatabase) HealthCheckAdapter {
	return healthCheckAdapter{db}
}

type healthCheckAdapter struct {
	bedatabase.BeDatabase
}

func (e healthCheckAdapter) GetDatabaseConnectionStatus(id string) {

	e.Query(id)
}

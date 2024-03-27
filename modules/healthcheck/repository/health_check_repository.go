package repository

import (
	"fmt"
	"healthcheck/adapters"
)

type HealthCheckRepository interface {
	GetCPU(id string)
}

type healthCheckRepository struct {
	adapters.HealthCheckAdapter
}

func NewHealthCheckRepository(adapter adapters.HealthCheckAdapter) HealthCheckRepository {
	return &healthCheckRepository{adapter}
}

// GetCPU implements HealtCheckRepository.
func (h *healthCheckRepository) GetCPU(id string) {
	fmt.Println("GetCPU : ", id)
}

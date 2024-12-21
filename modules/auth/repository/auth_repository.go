package repository

import (
	"auth/adapters"
	"fmt"
)

type AuthRepository interface {
	GetCPU(id string)
}

type authRepository struct {
	adapters.AuthAdapter
}

func NewHealthCheckRepository(adapter adapters.AuthAdapter) AuthRepository {
	return &authRepository{adapter}
}

// GetCPU implements HealtCheckRepository.
func (h *authRepository) GetCPU(id string) {
	fmt.Println("GetCPU : ", id)
}

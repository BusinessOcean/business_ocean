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

func NewAuthRepository(adapter adapters.AuthAdapter) AuthRepository {
	return &authRepository{adapter}
}

// GetCPU implements AuthRepository.
func (h *authRepository) GetCPU(id string) {
	fmt.Println("GetCPU : ", id)
}

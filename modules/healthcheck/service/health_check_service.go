package service

import (
	"beservice/healthcheck/apis"
	"context"
	"healthcheck/repository"
)

type HealthCheckService struct {
	repository repository.HealthCheckRepository
	apis.UnimplementedHealthCheckServiceServer
}

func NewHealthCheckService(repository repository.HealthCheckRepository) *HealthCheckService {

	return &HealthCheckService{repository: repository}
}

func (hc HealthCheckService) CheckHealthService(ctx context.Context, req *apis.CheckHealthServiceRequest) (res *apis.CheckHealthServiceResponse, error error) {
	// status.Errorf(codes.Unimplemented, "method CheckHealthService not implemented")

	return &apis.CheckHealthServiceResponse{Status: apis.CheckHealthServiceResponse_STATUS_HEALTHY, Message: "Health check Service is running"}, nil
}

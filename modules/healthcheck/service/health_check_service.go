package service

import (
	"beservice/healthcheck/apis"
	"context"
	"healthcheck/repository"

	firebase "firebase.google.com/go/v4"

	"github.com/kataras/iris/v12"
)

type HealthCheckService struct {
	repository repository.HealthCheckRepository
	apis.UnimplementedHealthCheckServiceServer
}

func NewHealthCheckService(repository repository.HealthCheckRepository, firebseApp *firebase.App) *HealthCheckService {

	return &HealthCheckService{repository: repository}
}

func (hc HealthCheckService) CheckHealthService(context.Context, *apis.CheckHealthServiceRequest) (*apis.CheckHealthServiceResponse, error) {

	return &apis.CheckHealthServiceResponse{Status: apis.CheckHealthServiceResponse_STATUS_HEALTHY, Message: "Health check Service is running"}, nil
}

func (hc HealthCheckService) HealthCheckApiRoute(ctx iris.Context) {
	ctx.JSON(iris.Map{"status": "HealthCheckApiRoute"})
}

func (hc HealthCheckService) HealthCheckApiRoute2(ctx iris.Context) {

	ctx.JSON(iris.Map{"status": "HealthCheckApiRoute2"})
}

package service

import (
	"auth/repository"
	"beservice/auth/apis"
	"context"

	firebase "firebase.google.com/go/v4"

	"github.com/kataras/iris/v12"
)

type AuthService struct {
	repository repository.AuthRepository
	apis.UnimplementedAuthServiceServer
}

func NewAuthService(repository repository.AuthRepository, firebseApp *firebase.App) *AuthService {

	return &AuthService{repository: repository}
}

func (hc AuthService) AuthStart(context.Context, *apis.AuthStartRequest) (*apis.AuthStartResponse, error) {

	return &apis.AuthStartResponse{Message: "Auth working"}, nil
}

func (hc AuthService) AuthHealthCheck(ctx iris.Context) {
	ctx.JSON(iris.Map{"status": "Auth service is up and running"})
}

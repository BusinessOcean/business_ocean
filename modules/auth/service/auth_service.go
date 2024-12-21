package service

import (
	"auth/repository"
	auth "beservice/auth/apis"
	"context"

	firebase "firebase.google.com/go/v4"

	"github.com/kataras/iris/v12"
)

type AuthService struct {
	repository repository.AuthRepository
	auth.UnimplementedAuthServiceServer
}

func NewAuthService(repository repository.AuthRepository, firebseApp *firebase.App) *AuthService {

	return &AuthService{repository: repository}
}

// AuthStart(context.Context, *AuthStartRequest) (*AuthStartResponse, error)
func (as AuthService) AuthStart(context.Context, *auth.AuthStartRequest) (*auth.AuthStartResponse, error) {

	return &auth.AuthStartResponse{Status: auth.AuthStartResponse_STATUS_UNKNOWN, Message: "Auth check Service is running"}, nil
}

func (as AuthService) AuthApiRoute(ctx iris.Context) {
	ctx.JSON(iris.Map{"status": "AuthApiRoute"})
}

func (as AuthService) AuthApiRoute2(ctx iris.Context) {

	ctx.JSON(iris.Map{"status": "AuthApiRoute2"})
}

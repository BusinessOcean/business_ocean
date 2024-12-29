package service

import (
	"auth/repository"
	"becommon/beevent"
	"beservice/auth/apis"
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"github.com/google/uuid"

	"github.com/kataras/iris/v12"
)

type AuthService struct {
	repository repository.AuthRepository
	eventbus   *beevent.BeEventBus
	apis.UnimplementedAuthServiceServer
}

func NewAuthService(repository repository.AuthRepository, firebseApp *firebase.App, eventBus *beevent.BeEventBus) *AuthService {

	return &AuthService{repository: repository, eventbus: eventBus}
}

func (hc AuthService) AuthStart(context.Context, *apis.AuthStartRequest) (*apis.AuthStartResponse, error) {

	return &apis.AuthStartResponse{Message: "Auth working"}, nil
}

func (hc AuthService) AuthHealthCheck(ctx iris.Context) {
	ctx.JSON(iris.Map{"status": "Auth service is up and running"})
}

func (hc AuthService) CreateUserEvent(ctx iris.Context) {
	err := hc.eventbus.Publish(beevent.UserCreatedEvent{
		Username: uuid.New().String(),
	})
	fmt.Println("User created event published:", err)
	ctx.JSON(iris.Map{"status": "Event test log"})
}

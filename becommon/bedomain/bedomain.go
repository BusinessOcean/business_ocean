package bedomain

import (
	"becore/beserver"

	"google.golang.org/grpc"
)

type BeDomainName string

type IBeDomain interface {
	Setup() error
	Register(grpc.ServiceDesc, interface{}) error
	GetServer() *beserver.BegoServer
	Run() error
	OnTerminate() error
}

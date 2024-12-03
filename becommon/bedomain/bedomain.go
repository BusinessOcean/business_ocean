package bedomain

import (
	"becore/beroutes"

	"google.golang.org/grpc"
)

type BeDomainName string

type IBeDomain interface {
	Setup() error
	Register(grpc.ServiceDesc, interface{}) error
	RegisterRoutes(grpc.ServiceDesc, []*beroutes.Route) error
	Run() error
	OnTerminate() error
}

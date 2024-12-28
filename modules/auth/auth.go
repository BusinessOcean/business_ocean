package auth

import (
	"auth/service"
	"becommon/bedomain"
	"becore/beconfig"
	"becore/belogger"
	"becore/beroutes"
	"becore/beserver"

	"go.uber.org/fx"
)

var _ bedomain.IBeDomain = (*BeAuthkDomain)(nil)

type DomainParams struct {
	fx.In
	Config      *beconfig.AppConfig
	Logger      *belogger.BeLogger
	BaseModules *bedomain.BaseDomain `name:"domain"`
	Service     *service.AuthService
	Routes      []*beroutes.Route `name:"authroutes"`
}

type BeAuthkDomain struct {
	*bedomain.BaseDomain

	service *service.AuthService
	routes  []*beroutes.Route
}

func NewBeAuthDomain(params DomainParams) *BeAuthkDomain {
	// TODO: Do it in better way. Should not create new server here using uber fx
	grpc := beserver.NewBeGRPCServer()
	http := beserver.NewBeHttpServer(params.Config)
	server := beserver.NewBegoServer(params.Logger, http, grpc)
	base := bedomain.NewBaseDomain(params.Config, params.Logger, server)

	// ISSUE: if injected server it was using same server instance which cause run one module at a time

	return &BeAuthkDomain{
		BaseDomain: base,
		service:    params.Service,
		routes:     params.Routes,
	}
}

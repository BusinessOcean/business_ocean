package auth

import (
	"auth/service"
	"becommon/bedomain"
	"becore/beconfig"
	"becore/belogger"
	"becore/beroutes"

	"go.uber.org/fx"
)

var _ bedomain.IBeDomain = (*BeAuthkDomain)(nil)

type DomainParams struct {
	fx.In
	Config      *beconfig.AppConfig
	Logger      *belogger.BeLogger
	BaseModules *bedomain.BaseDomain `name:"authdomain"`
	Service     *service.AuthService
	Routes      []*beroutes.Route `name:"authroutes"`
}

type BeAuthkDomain struct {
	*bedomain.BaseDomain

	service *service.AuthService
	routes  []*beroutes.Route
}

func NewBeAuthDomain(params DomainParams) *BeAuthkDomain {

	return &BeAuthkDomain{
		BaseDomain: params.BaseModules,
		service:    params.Service,
		routes:     params.Routes,
	}
}

package healthcheck

import (
	"becommon/bedomain"
	"becore/beroutes"
	"healthcheck/service"

	"go.uber.org/fx"
)

var _ bedomain.IBeDomain = (*BeHealthCheckDomain)(nil)

type DomainParams struct {
	fx.In

	BaseModules *bedomain.BaseDomain `name:"domain"`
	Service     *service.HealthCheckService
	Routes      []beroutes.IRegisterRouteAPI
}

type BeHealthCheckDomain struct {
	*bedomain.BaseDomain
	service *service.HealthCheckService
	routes  []beroutes.IRegisterRouteAPI
}

func NewBeHealthCheckDomain(params DomainParams) *BeHealthCheckDomain {
	return &BeHealthCheckDomain{
		BaseDomain: params.BaseModules,
		service:    params.Service,
		routes:     params.Routes,
	}
}

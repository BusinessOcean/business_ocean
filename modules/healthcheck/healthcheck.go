package healthcheck

import (
	"becommon/bedomain"
	"healthcheck/service"

	"go.uber.org/fx"
)

var _ bedomain.IBeDomain = (*BeHealthCheckDomain)(nil)

type DomainParams struct {
	fx.In

	BaseModules *bedomain.BaseDomain `name:"domain"`
	Service     *service.HealthCheckService
}

type BeHealthCheckDomain struct {
	*bedomain.BaseDomain

	service *service.HealthCheckService
}

func NewBeHealthCheckDomain(params DomainParams) *BeHealthCheckDomain {
	return &BeHealthCheckDomain{
		BaseDomain: params.BaseModules,
		service:    params.Service,
	}
}

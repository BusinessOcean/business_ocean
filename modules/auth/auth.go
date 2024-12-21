package auth

import (
	"auth/service"
	"becommon/bedomain"
	"becore/beroutes"

	"go.uber.org/fx"
)

var _ bedomain.IBeDomain = (*BeAuthDomain)(nil)

type DomainParams struct {
	fx.In

	BaseModules *bedomain.BaseDomain `name:"domain"`
	Service     *service.AuthService
	// Routes      []*beroutes.Route
}

type BeAuthDomain struct {
	*bedomain.BaseDomain

	service *service.AuthService
	routes  []*beroutes.Route
}

func NewBeAuthDomain(params DomainParams) *BeAuthDomain {
	return &BeAuthDomain{
		BaseDomain: params.BaseModules,
		service:    params.Service,
		// routes:     params.Routes,
	}
}

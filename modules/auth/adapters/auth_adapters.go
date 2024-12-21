package adapters

import (
	"bedatabase/dgraphdb"

	"go.uber.org/fx"
)

type authAdapter struct {
	*dgraphdb.BeDgraphDB
}

type AuthAdapter interface {
	GetDatabaseConnectionStatus(id string)
}

type AuthAdapterParam struct {
	fx.In

	DgraphDB *dgraphdb.BeDgraphDB `name:"dgraphdb"`
}

func NewAuthAdapter(p AuthAdapterParam) AuthAdapter {
	return authAdapter{p.DgraphDB}
}

func (e authAdapter) GetDatabaseConnectionStatus(id string) {

	e.Query(id)
}

package adapters

import (
	"becore/belogger"
	"bedatabase/dgraphdb"

	"go.uber.org/fx"
)

type authAdapter struct {
	*dgraphdb.BeDgraphDB
	logger *belogger.BeLogger
}

type AuthAdapter interface {
	GetAuthServiceStatus() error
}

type AuthAdapterParam struct {
	fx.In

	DgraphDB *dgraphdb.BeDgraphDB `name:"dgraphdb"`
	Logger   *belogger.BeLogger
}

func NewAuthAdapter(p AuthAdapterParam) AuthAdapter {
	return authAdapter{BeDgraphDB: p.DgraphDB, logger: p.Logger}
}

func (e authAdapter) GetAuthServiceStatus() error {
	e.logger.Info("GetAuthServiceStatus : Ping Firebase Auth Service")

	// e.Query()
	return nil
}

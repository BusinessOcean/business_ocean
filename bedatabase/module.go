package bedatabase

import (
	"becommon/fxutil"
	"bedatabase/dgraphdb"
	"bedatabase/memorydb"
	"context"
	"fmt"

	"go.uber.org/fx"
)

type DatabaseModuleParams struct {
	fx.In
	Lifecycle  fx.Lifecycle
	BeDgraphDB *dgraphdb.BeDgraphDB `name:"dgraphdb"`
	InMemoryDB *memorydb.InMemoryDB `name:"memorydb"`
}

var DatabaseModules = fx.Options(
	fxutil.AnnotatedProvide(dgraphdb.NewBeDgraphDB, `name:"dgraphdb"`),
	fxutil.AnnotatedProvide(memorydb.NewInMemoryDB, `name:"memorydb"`),
	fx.Invoke(registerGraphDBHooks),
)

func registerGraphDBHooks(param DatabaseModuleParams) {
	param.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Connecting to BeDgraphDB...")
			param.BeDgraphDB.Connect()
			param.InMemoryDB.Connect()
			// graphDB.Ping()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Disconnecting from BeDgraphDB...")
			param.BeDgraphDB.Disconnect()
			return nil
		},
	})
}

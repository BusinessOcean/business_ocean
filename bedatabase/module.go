package bedatabase

import (
	"bedatabase/dgraphdb"
	"bedatabase/memorydb"
	"context"
	"fmt"

	"go.uber.org/fx"
)

var DatabaseModule = fx.Options(
	fx.Provide(memorydb.NewInMemoryDB),
	fx.Provide(dgraphdb.NewBeDgraphDB),
	fx.Provide(fx.Annotate(dgraphdb.NewBeDgraphDB, fx.As(new(BeDatabase)))),
	fx.Invoke(registerGraphDBHooks),
)

func registerGraphDBHooks(lc fx.Lifecycle, graphDB *dgraphdb.BeDgraphDB) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Connecting to BeDgraphDB...")
			graphDB.Connect()
			// graphDB.Ping()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Disconnecting from BeDgraphDB...")
			graphDB.Disconnect()
			return nil
		},
	})
}

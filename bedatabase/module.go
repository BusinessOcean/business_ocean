package bedatabase

import (
	"bedatabase/dgraphdb"
	"bedatabase/memorydb"
	"fmt"

	"go.uber.org/fx"
)

var DatabaseModule = fx.Options(
	fx.Provide(memorydb.NewInMemoryDB),
	fx.Provide(dgraphdb.NewBeDgraphDB),
	fx.Provide(fx.Annotate(dgraphdb.NewBeDgraphDB, fx.As(new(BeDatabase)))),
	fx.Invoke(func(graphDB *dgraphdb.BeDgraphDB) {
		graphDB.Connect()
		fmt.Println("DatabaseModule #########################################################")
	}),
)

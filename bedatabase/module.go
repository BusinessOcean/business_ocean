package bedatabase

import (
	"go.uber.org/fx"
)

var DatabaseModule = fx.Options(
	fx.Provide(NewInMemoryDatabase),
)

package becommon

import (
	"becommon/bedomain"

	"go.uber.org/fx"
)

var BeCommonModule = fx.Options(
	fx.Provide(bedomain.NewBaseDomainModule),
)

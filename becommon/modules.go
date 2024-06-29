package becommon

import (
	"becommon/bedomain"
	"becommon/fxutil"

	"go.uber.org/fx"
)

var BeCommonModule = fx.Options(
	fxutil.AnnotatedProvide(bedomain.NewBaseDomain, `name:"domain"`),
)

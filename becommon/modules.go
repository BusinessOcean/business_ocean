package becommon

import (
	"becommon/bedomain"
	"becore/beconfig"
	"becore/belogger"
	"becore/beserver"

	"go.uber.org/fx"
)

var BeCommonModule = fx.Options(
	// fxutil.AnnotatedProvide(bedomain.NewBaseDomain, `name:"domain"`),
	fx.Provide(
		fx.Annotate(
			func(
				config *beconfig.AppConfig,
				logger *belogger.BeLogger,
				server *beserver.BegoServer,
			) *bedomain.BaseDomain {
				logger.Warn("----------------------------------------------------")
				logger.Warn("Createing instance of authdomain BaseDomain")
				logger.Warn("----------------------------------------------------")
				grpc := beserver.NewBeGRPCServer()
				http := beserver.NewBeHttpServer(config)
				ns := beserver.NewBegoServer(logger, http, grpc)
				return bedomain.NewBaseDomain(config, logger, ns)
			},
			fx.ResultTags(`name:"authdomain"`),
		),
	),
	fx.Provide(
		fx.Annotate(
			func(
				config *beconfig.AppConfig,
				logger *belogger.BeLogger,
				server *beserver.BegoServer,
			) *bedomain.BaseDomain {
				logger.Warn("----------------------------------------------------")
				logger.Warn("Createing instance of healthcheckdomain BaseDomain")
				logger.Warn("----------------------------------------------------")
				grpc := beserver.NewBeGRPCServer()
				http := beserver.NewBeHttpServer(config)
				ns := beserver.NewBegoServer(logger, http, grpc)
				return bedomain.NewBaseDomain(config, logger, ns)
			},
			fx.ResultTags(`name:"healthcheckdomain"`),
		),
	),
)

package bego

import (
	"beconsole/commands"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type bego struct{}

// NewBego returns a new instance of Bego
func NewBegoServer() *bego {
	return &bego{}
}

func (b *bego) Short() string {
	return " Run Bego Server "
}
func (b *bego) Setup(cmd *cobra.Command) error {
	return nil
}
func (b *bego) Run() (commands.CommandRunner, error) {
	// return func() {}, nil

	return func(
	// middleware middlewares.Middlewares,
	// env *lib.Env,
	// router infrastructure.Router,
	// route routes.Routes,
	// logger lib.Logger,
	// database infrastructure.Database,
	// seeds seeds.Seeds,

	) {
		// logger.Info(`+-----------------------+`)
		// logger.Info(`| GO CLEAN ARCHITECTURE |`)
		// logger.Info(`+-----------------------+`)

		// Using time zone as specified in env file
		// loc, _ := time.LoadLocation(env.TimeZone)
		// time.Local = loc

		// middleware.Setup()
		// route.Setup()
		// seeds.Setup()

		// if env.Environment != "local" && env.SentryDSN != "" {
		// 	err := sentry.Init(sentry.ClientOptions{
		// 		Dsn:              env.SentryDSN,
		// 		AttachStacktrace: true,
		// 	})
		// 	if err != nil {
		// 		logger.Error("sentry initialization failed")
		// 		logger.Error(err.Error())
		// 	}
		// }
		// logger.Info("Running server")
		// if env.ServerPort == "" {
		// 	if err := router.Run(); err != nil {
		// 		logger.Fatal(err)
		// 		return
		// 	}
		// } else {
		// 	if err := router.Run(":" + env.ServerPort); err != nil {
		// 		logger.Fatal(err)
		// 		return
		// 	}
		// }
		fmt.Println("I am working")

		http.ListenAndServe(":8080", nil)
	}, nil
}

package beserver

import (
	"becore/beconfig"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

type BeServer struct {
	*iris.Application
	AppName string
}

type BeContext = iris.Context

// type BeContext  iris.Context

func NewBeServer(config beconfig.ServerConfig) *BeServer {
	_app := iris.New()
	_app.SetName(config.AppName)
	_app.Use(iris.Compression)
	_app.Favicon("./static/favicons/businessocean.ico")

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	_app.Use(crs)

	return &BeServer{iris.New(), config.AppName}
}

func (b *BeServer) RegisterRoutes() {
	// routes := ro.NewRoutes()
	// _ = routes.Setup()
}

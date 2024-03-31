package beserver

import (
	"becore/beconfig"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

type BeHTTPServer struct {
	*iris.Application
	AppName string
}

type BeContext = iris.Context

// type BeContext  iris.Context
func NewBeHttpServer(config *beconfig.Config) *BeHTTPServer {
	_app := iris.New()
	_app.SetName(config.AppName)
	_app.Use(iris.Compression)
	_app.Favicon("./static/favicons/businessocean.ico")

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	_app.Use(crs)

	return &BeHTTPServer{iris.New(), config.AppName}
}

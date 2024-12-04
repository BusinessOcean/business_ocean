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

// type BeContext  iris.Context
func NewBeHttpServer(config *beconfig.AppConfig) *BeHTTPServer {

	server := &BeHTTPServer{iris.New(), config.Info.AppName}
	server.Logger().SetLevel("debug")
	server.SetName(config.Info.AppName)
	server.Use(iris.Compression)
	server.Favicon("./static/favicons/businessocean.ico")
	server.HandleDir("/", iris.Dir("./public"), iris.DirOptions{
		IndexName: "index.html",
		SPA:       true,
	})
	corsOpt := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}

	crs := cors.New(corsOpt)
	server.Use(crs)

	return server

}

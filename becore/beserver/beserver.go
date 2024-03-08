package beserver

import (
	beconfig "becore/config"

	"github.com/kataras/iris/v12"
)

type BeServer struct {
	*iris.Application
	AppName string
}

type BeContext iris.Context

// type BeContext  iris.Context

func NewBeServer(config beconfig.ServerConfig) *BeServer {
	_app := iris.New()
	_app.SetName(config.AppName)
	_app.Use(iris.Compression)

	return &BeServer{iris.New(), config.AppName}
}

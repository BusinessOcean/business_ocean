package beconfig

type ServerConfig struct {
	Port         string
	LogLevel     string
	AppName      string
	IsProduction bool
	FavIconPath  string
}

// NewServerConfig returns a new instance of ServerConfig
func NewServerConfig() ServerConfig {
	// TODO: load from env
	return ServerConfig{Port: "8080",
		LogLevel:     "debug",
		AppName:      "Bego",
		IsProduction: false,
		FavIconPath:  "./static/favicons/businessocean.ico",
	}
}

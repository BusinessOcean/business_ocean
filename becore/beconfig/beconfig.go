package beconfig

import (
	"becore/belogger"

	"github.com/spf13/viper"
)

type Config struct {
	LogLevel           string `mapstructure:"LOG_LEVEL"`
	ServerPort         string `mapstructure:"SERVER_PORT"`
	AppName            string `mapstructure:"APP_NAME"`
	FavIconPath        string `mapstructure:"FAV_ICON_PATH"`
	TimeZone           string `mapstructure:"TIMEZONE"`
	MaxMultipartMemory int64  `mapstructure:"MAX_MULTIPART_MEMORY"`
}

// NewConfig returns a new instance of ServerConfig
func NewConfig(logger belogger.BeLogger) *Config {

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal("cannot read configuration from .env file with viper ", err)
	}

	viper.SetDefault("TIMEZONE", "UTC")

	err = viper.Unmarshal(&globalConfig)
	if err != nil {
		logger.Fatal("environment cant be loaded: ", err)
	}

	return &globalConfig
}

var globalConfig = Config{
	MaxMultipartMemory: 10 << 20, // 10 MB
}

func GetConfig() Config {
	return globalConfig
}

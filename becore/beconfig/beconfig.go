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

// Config holds the configuration values
type AppConfig struct {
	Dgraph   DgraphConfig
	Postgres PostgresConfig
	JWT      JWTConfig
	Cloud    CloudConfig
	Common   CommonConfig
}

type DgraphConfig struct {
	Host string
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type JWTConfig struct {
	Secret string
}

type CloudConfig struct {
	Secret string
}

type CommonConfig struct {
	SomeOtherSecret string
}

// NewConfig returns a new instance of ServerConfig
func NewConfig(logger belogger.BeLogger) (*Config, error) {

	// viper.SetConfigFile(".env")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var config AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal("cannot read configuration from .env file with viper ", err)
	}

	viper.SetDefault("TIMEZONE", "UTC")

	err = viper.Unmarshal(&globalConfig)
	if err != nil {
		logger.Fatal("environment cant be loaded: ", err)
	}

	return &globalConfig, nil
}

var globalConfig = Config{
	MaxMultipartMemory: 10 << 20, // 10 MB
}

func GetConfig() Config {
	return globalConfig
}

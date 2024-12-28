package beconfig

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config struct to hold configuration values
type AppConfig struct {
	Info        Info           `mapstructure:"info"`
	Dgraph      DgraphConfig   `mapstructure:"dgraph"`
	Postgres    PostgresConfig `mapstructure:"postgres"`
	Jwt         JwtConfig      `mapstructure:"jwt"`
	Cloud       CloudConfig    `mapstructure:"cloud"`
	Auth        ModuleConfig   `mapstructure:"auth"`
	HealthCheck ModuleConfig   `mapstructure:"healthcheck"`
}

// DgraphConfig holds Dgraph-related settings
type DgraphConfig struct {
	Host string `mapstructure:"host"`
}

type ModuleConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

// PostgresConfig holds Postgres-related settings
type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
}

// JwtConfig holds JWT-related settings
type JwtConfig struct {
	Secret string `mapstructure:"secret"`
}

// CloudConfig holds Cloud-related settings
type CloudConfig struct {
	Secret string `mapstructure:"secret"`
}

// Info holds other common secrets
type Info struct {
	AppName      string `mapstructure:"app_name"`
	AppVersion   string `mapstructure:"app_version"`
	BuildVersion string `mapstructure:"build_version"`
	Mode         string `mapstructure:"mode"`
}

// NewAppConfig loads the configuration from a file and returns a Config struct
func NewAppConfig() (*AppConfig, error) {
	// Initialize Viper
	v := viper.New()

	// Set the config type to YAML
	v.SetConfigType("yaml")

	// Set paths to look for the config file
	v.AddConfigPath(".")      // Current directory
	v.SetConfigName("config") // Name of the config file (without extension)

	// Read in the config file
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Log the raw config data (for debugging purposes)
	// fmt.Println("Raw config data:", v.AllSettings())

	// Create an instance of the Config struct
	var config AppConfig

	// Unmarshal the config into the struct
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	// Return the config struct
	return &config, nil
}

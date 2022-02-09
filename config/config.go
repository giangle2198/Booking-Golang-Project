package config

import (
	"bytes"
	"strings"

	"github.com/spf13/viper"
)

var defaultConfig = []byte(`
grpc_address: 6000
http_address: 6001
environment: D
app: booking-center-service
sensitive_fields: 
mongo:
  url: mongodb://172.27.42.11:27017
  database: giang
bim_client:
  host: 172.27.42.11
  port: 20000
excluded_paths:
  - GetProductDetail
  - GetListProduct
  - GetServices
  - GetCountries
enabled_cors: true
`)

type (
	//
	Config struct {
		Base `mapstructure:",squash"`
	}

	Base struct {
		Environment     string            `mapstructure:"enviroment"`
		App             string            `mapstructure:"app"`
		HTTPAddress     int               `mapstructure:"http_address"`
		GRPCAddress     int               `mapstructure:"grpc_address"`
		SensitiveFields map[string]string `mapstructure:"sensitive_fields"`
		Mongo           mongo             `mapstructure:"mongo"`
		BIMClient       bimClient         `mapstructure:"bim_client"`
		ExcludedPaths   []string          `mapstructure:"excluded_paths"`
		EnabledCors     bool              `mapstructure:"enabled_cors"`
	}

	mongo struct {
		URL      string `mapstructure:"url"`
		Database string `mapstructure:"database"`
	}

	bimClient struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	}
)

func Load() (*Config, error) {
	var cfg = &Config{}
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(defaultConfig))
	if err != nil {
		return nil, err
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

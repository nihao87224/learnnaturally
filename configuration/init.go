package configuration

import (
	"fmt"
	"github.com/spf13/viper"
)

type ConfigData struct {
	Common Common `yaml:"common"`
}
type Common struct {
	ProxyUrl string `yaml:"proxy_url"`
}

func LoadConfig(configPath string) (*ConfigData, error) {
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %s", err)
	}

	var cfg ConfigData
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %s", err)
	}
	return &cfg, nil
}

package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBname        string `mapstructure:"DB_NAME"`
	CollName      string `mapstructure:"COLL_NAME"`
	DBuri         string `mapstructure:"DB_URI"`
	DBuser        string `mapstructure:"DB_USER"`
	HttpAddrSite  string `mapstructure:"HTTP_ADDR_SITE"`
	HttpAddrAdmin string `mapstructure:"HTTP_ADDR_ADMIN"`
	Env           string `mapstructure:"Env"`
	AdminPass     string `mapstructure:"ADMIN_PASSWORD"`
	AdminName     string `mapstructure:"ADMIN_NAME"`
	AdminCookie   string `mapstructure:"ADMIN_COOKIE"`
}

func InitConfig(path string) (*Config, error) {
	var config *Config
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return config, nil
}

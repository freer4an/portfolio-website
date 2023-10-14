package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App      *App      `mapstructure:"app"`
	Database *Database `mapstructure:"db"`
	Env      string    `mapstructure:"env"`
}

type App struct {
	Addr     string `mapstructure:"addr"`
	MaxAge   string `mapstructure:"maxage"`
	Duration string `mapstructure:"duration"`
}

type Database struct {
	Name        string `mapstructure:"name"`
	Uri         string `mapstructure:"uri"`
	User        string `mapstructure:"user"`
	Passowrd    string `mapstructure:"password"`
	CollProject string `mapstructure:"coll-projects"`
}

func InitConfig(path string) (*Config, error) {
	var config *Config

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return config, nil
}

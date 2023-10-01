package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Cookie   *Cookie   `mapstructure:"cookie"`
	App      *App      `mapstructure:"app"`
	Admin    *Admin    `mapstructure:"admin"`
	Database *Database `mapstructure:"db"`
	Env      string    `mapstructure:"env"`
}

type Cookie struct {
	Admin    string `mapstructure:"admin"`
	MaxAge   string `mapstructure:"maxage"`
	Duration string `mapstructure:"duration"`
}

type App struct {
	Addr     string `mapstructure:"addr"`
	MaxAge   string `mapstructure:"maxage"`
	Duration string `mapstructure:"duration"`
}

type Admin struct {
	Login    string `mapstructure:"login"`
	Password string `mapstructure:"password"`
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

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return config, nil
}

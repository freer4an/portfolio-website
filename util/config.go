package util

import (
	"github.com/rs/zerolog/log"

	"github.com/spf13/viper"
)

type Config struct {
	DBname   string `mapstructure:"DB_NAME"`
	CollName string `mapstructure:"COLL_NAME"`
	DBuri    string `mapstructure:"DB_URI"`
	DBuser   string `mapstructure:"DB_USER"`
	HttpAddr string `mapstructure:"HTTP_ADDR"`
}

func InitConfig(path string) (config Config) {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal().Err(err)
	}

	return
}

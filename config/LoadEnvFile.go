package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadEnvFile(config AppConfig) (err error) {
	// find file in current working directory
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("unmarshal", err)
	}
	return
}

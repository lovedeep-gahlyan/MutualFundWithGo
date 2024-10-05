package config

import (
	"log"
	"github.com/spf13/viper"
)

// *viper.Viper is a struct in Go that comes from the Viper Library.
// It is used as a configuration management tool

func InitConfig(fileName string) *viper.Viper {
	config := viper.New()

	// automatically search for extention - .yml, .toml
	config.SetConfigName(fileName)

	// search where go program starts - where main.go lies
	config.AddConfigPath(".")
	// search config file in home directory also
	config.AddConfigPath("$HOME")

	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("Error while parsing configuration file", err)
	}

	return config
}

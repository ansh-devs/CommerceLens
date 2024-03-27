package config

import (
	"log"
)
import "github.com/spf13/viper"

type AppConfig struct {
	HttpAddr         string `mapstructure:"HTTPPORT"`
	DatabaseName     string `mapstructure:"DBNAME"`
	DatabaseHost     string `mapstructure:"DBHOST"`
	DatabaseUser     string `mapstructure:"DBUSER"`
	DatabasePassword string `mapstructure:"DBPASSWORD"`
}

var AppConfigs *AppConfig

func LoadEnvVariables() (config *AppConfig) {
	viper.SetDefault("HTTPPORT", ":8080")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	AppConfigs = config
	return
}

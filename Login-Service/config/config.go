package config

import "log"
import "github.com/spf13/viper"

var AppConfigs *AppConfig

func InitEnvConfigs() {
	AppConfigs = loadEnvVariables()
}

type AppConfig struct {
	HttpAddr         string `mapstructure:"HTTPPORT"`
	DatabaseName     string `mapstructure:"DBNAME"`
	DatabaseHost     string `mapstructure:"DBHOST"`
	DatabaseUser     string `mapstructure:"DBUSER"`
	DatabasePassword string `mapstructure:"DBPASSWORD"`
}

func loadEnvVariables() (config *AppConfig) {
	viper.AddConfigPath(".")
	viper.AddConfigPath("/app")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}

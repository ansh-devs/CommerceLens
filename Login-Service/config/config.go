package config

import "log"
import "github.com/spf13/viper"

var AppConfigs *AppConfig

func InitEnvConfigs() {
	AppConfigs = loadEnvVariables()
}

type AppConfig struct {
	GrpcAddr    string `mapstructure:"GRPC_ADDRESS"`
	HttpAddr    string `mapstructure:"HTTP_ADDRESS"`
	DatabaseUrl string `mapstructure:"DATABASE_URL"`
}

func loadEnvVariables() (config *AppConfig) {
	// Tell viper the path/location of your env file. If it is root just add "."
	viper.AddConfigPath(".")

	// Tell viper the name of your file
	viper.SetConfigName(".env")

	// Tell viper the type of your file
	viper.SetConfigType("env")

	// Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}

func NewServerConfig(grpcAddr string, httpAddr string, databaseUrl string) *AppConfig {
	return &AppConfig{
		GrpcAddr:    grpcAddr,
		HttpAddr:    httpAddr,
		DatabaseUrl: databaseUrl,
	}
}

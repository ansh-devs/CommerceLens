package main

import (
	"github.com/ansh-devs/microservices_project/login-service/config"
	"github.com/ansh-devs/microservices_project/login-service/server"
)

func main() {
	config.InitEnvConfigs()
	srvc := server.Init()
	srvc.Boot()
}

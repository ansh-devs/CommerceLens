package main

import (
	"github.com/ansh-devs/microservices_project/product-service/config"
	"github.com/ansh-devs/microservices_project/product-service/server"
)

func main() {

	config.InitEnvConfigs()
	srvc := server.Init()
	srvc.Boot()
}

package main

import (
	"github.com/ArthurHlt/go-eureka-client/eureka"
	"github.com/ansh-devs/microservices_project/product-service/config"
	"github.com/ansh-devs/microservices_project/product-service/db"
	"github.com/ansh-devs/microservices_project/product-service/server"
)

func registerService() {

	client := eureka.NewClient([]string{
		"http://192.168.1.5:50000/eureka", //From a spring boot based eureka server
		// add others servers here
	})
	instance := eureka.NewInstanceInfo("192.168.1.5", "Product-Service", "192.168.1.5", 50005, 5, false) //Create a new instance to register
	instance.Metadata = &eureka.MetaData{
		Map: make(map[string]string),
	}
	instance.Metadata.Map["grpcuri"] = "192.168.1.5:50004"
	err := client.RegisterInstance("product-service", instance)
	if err != nil {
		return
	} // Register new instance in your eureka(s)
	_, _ = client.GetApplications() // Retrieves all applications from eureka server(s)
	_, err = client.GetApplication(instance.App)
	if err != nil {
		return
	} // retrieve the application "test"
	_, err = client.GetInstance(instance.App, instance.HostName)
	if err != nil {
		return
	} // retrieve the instance from "test.com" inside "test"" app
	err = client.SendHeartbeat(instance.App, instance.HostName)
	if err != nil {
		return
	} // say to eureka that your app is alive (here you must send heartbeat before 30 sec)
}

func main() {
	config.InitEnvConfigs()
	pgConn := db.MustConnectToPostgress(config.AppConfigs.DatabaseUrl)
	go server.MustStartGrpcServer(pgConn, config.AppConfigs.GrpcAddr)
	server.MustStartHttpServer(pgConn, config.AppConfigs.HttpAddr)
}

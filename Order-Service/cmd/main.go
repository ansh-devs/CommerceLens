package main

import (
	"fmt"
	"github.com/ArthurHlt/go-eureka-client/eureka"
	"github.com/ansh-devs/microservices_project/order-service/config"
	"github.com/ansh-devs/microservices_project/order-service/db"
	"github.com/ansh-devs/microservices_project/order-service/server"
)

func registerService() *eureka.Applications {

	client := eureka.NewClient([]string{
		"http://192.168.1.5:50000/eureka", //From a spring boot based eureka server
		// add others servers here
	})
	instance := eureka.NewInstanceInfo("192.168.1.5", "Order-Service", "192.168.1.5", 50007, 5, false) //Create a new instance to register
	instance.Metadata = &eureka.MetaData{
		Map: make(map[string]string),
	}
	instance.Metadata.Map["foo"] = "bar" //add metadata for example
	err := client.RegisterInstance("order-service", instance)
	if err != nil {
		fmt.Println(err)
	} // Register new instance in your eureka(s)
	apps, err := client.GetApplications() // Retrieves all applications from eureka server(s)
	if err != nil {
		fmt.Println(err)
	}
	_, err = client.GetApplication(instance.App)
	if err != nil {
		fmt.Println(err)
	} // retrieve the application "test"
	_, err = client.GetInstance(instance.App, instance.HostName)
	if err != nil {
		fmt.Println(err)
	} // retrieve the instance from "test.com" inside "test"" app
	go func() {
		err = client.SendHeartbeat(instance.App, instance.HostName)
		if err != nil {
			fmt.Println(err)
		}
	}()
	return apps
}

func main() {
	config.InitEnvConfigs()
	pgConn := db.MustConnectToPostgress(config.AppConfigs.DatabaseUrl)
	go server.MustStartGrpcServer(pgConn, config.AppConfigs.GrpcAddr)
	server.MustStartHttpServer(pgConn, config.AppConfigs.HttpAddr)
}

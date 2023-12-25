package server

import (
	"github.com/ansh-devs/microservices_project/product-service/config"
	"github.com/ansh-devs/microservices_project/product-service/db"
	db_gen "github.com/ansh-devs/microservices_project/product-service/db/generated"
	"github.com/hashicorp/consul/api"
)

type ProductService struct {
	*db_gen.Queries
	//httpAddr string
	//grpcAddr string
	consulClient *api.Client
}

func Init() *ProductService {
	client, _ := api.NewClient(&api.Config{})
	return &ProductService{
		Queries:      nil,
		consulClient: client,
	}
}
func (srv *ProductService) Boot() {
	srv.registerService(config.AppConfigs.HttpAddr)
	go srv.updateHealthStatus()
	_, loginsrvcport := srv.getloginservice()
	pgConn := db.MustConnectToPostgress(config.AppConfigs.DatabaseUrl)
	go MustStartGrpcServer(pgConn, config.AppConfigs.GrpcAddr, loginsrvcport)
	MustStartHttpServer(pgConn, config.AppConfigs.HttpAddr, loginsrvcport)
}

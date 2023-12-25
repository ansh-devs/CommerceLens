package server

import (
	"github.com/ansh-devs/microservices_project/login-service/config"
	"github.com/ansh-devs/microservices_project/login-service/db"
	db_gen "github.com/ansh-devs/microservices_project/login-service/db/generated"
	"github.com/hashicorp/consul/api"
)

type LoginService struct {
	*db_gen.Queries
	//httpAddr string
	//grpcAddr string
	consulClient *api.Client
}

func Init() *LoginService {
	client, _ := api.NewClient(&api.Config{})
	return &LoginService{
		Queries:      nil,
		consulClient: client,
		//httpAddr: httpAddress,
		//grpcAddr: grpcAddress,
	}
}
func (srv *LoginService) Boot() {
	srv.registerService(config.AppConfigs.HttpAddr)
	go srv.updateHealthStatus()
	pgConn := db.MustConnectToPostgress(config.AppConfigs.DatabaseUrl)
	go MustStartGrpcServer(pgConn, config.AppConfigs.GrpcAddr)
	MustStartHttpServer(pgConn, config.AppConfigs.HttpAddr)
}

package main

import (
	"github.com/ansh-devs/microservices_project/order-service/db"
	"github.com/ansh-devs/microservices_project/order-service/repo"
	"github.com/ansh-devs/microservices_project/order-service/service"
	"testing"
)

func TestServer(m *testing.T) {
	m.Run("testing server", func(t *testing.T) {
		var srv *service.OrderService
		dbConn := db.MustConnectToPostgress(dbSource)
		repository := repo.NewRepo(dbConn, nil, nil)
		srv = service.NewService(repository, nil, nil)
		if srv == nil {
			m.Error("Server is empty")
		}
	})
}

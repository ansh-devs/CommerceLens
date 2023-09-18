package gapi

import (
	db "github.com/ansh-devs/microservices_project/order-service/db/generated"
	baseproto "github.com/ansh-devs/microservices_project/order-service/proto"
)

type OrderRpcServer struct {
	*db.Queries
	baseproto.UnimplementedOrderServiceServer
	ProductServicePORT string
	LoginServicePORT   string
}

func NewGrpcServer(pgConn *db.Queries) *OrderRpcServer {
	return &OrderRpcServer{
		Queries:                         pgConn,
		UnimplementedOrderServiceServer: baseproto.UnimplementedOrderServiceServer{},
		ProductServicePORT:              ":50004",
		LoginServicePORT:                ":50002",
	}
}

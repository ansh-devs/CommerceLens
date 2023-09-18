package gapi

import (
	db "github.com/ansh-devs/microservices_project/product-service/db/generated"
	baseproto "github.com/ansh-devs/microservices_project/product-service/proto"
)

type ProductRpcServer struct {
	*db.Queries
	baseproto.UnimplementedProductServiceServer
}

func NewGrpcServer(pgConn *db.Queries) *ProductRpcServer {
	return &ProductRpcServer{
		Queries: pgConn,
	}
}

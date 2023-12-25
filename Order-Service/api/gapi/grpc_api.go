package gapi

import (
	db "github.com/ansh-devs/microservices_project/order-service/db/generated"
	baseproto "github.com/ansh-devs/microservices_project/order-service/proto"
	"strconv"
)

type OrderRpcServer struct {
	*db.Queries
	baseproto.UnimplementedOrderServiceServer
	ProductServicePORT string
	LoginServicePORT   string
}

func NewGrpcServer(pgConn *db.Queries, loginsrvcaddr string, loginsrvcport int) *OrderRpcServer {
	return &OrderRpcServer{
		Queries:                         pgConn,
		UnimplementedOrderServiceServer: baseproto.UnimplementedOrderServiceServer{},
		ProductServicePORT:              ":50004",
		LoginServicePORT:                ":" + strconv.Itoa(loginsrvcport),
	}
}

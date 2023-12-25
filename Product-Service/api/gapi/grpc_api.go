package gapi

import (
	db "github.com/ansh-devs/microservices_project/product-service/db/generated"
	baseproto "github.com/ansh-devs/microservices_project/product-service/proto"
	"strconv"
)

type ProductRpcServer struct {
	*db.Queries
	baseproto.UnimplementedProductServiceServer
	ProductServicePORT string
	LoginServicePORT   string
}

func NewGrpcServer(pgConn *db.Queries, loginsrvcport int) *ProductRpcServer {
	return &ProductRpcServer{
		Queries:            pgConn,
		ProductServicePORT: ":50004",
		LoginServicePORT:   ":" + strconv.Itoa(loginsrvcport),
	}
}

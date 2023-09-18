package server

import (
	"fmt"
	baserepo "github.com/ansh-devs/microservices_project/order-service/api/gapi"
	db "github.com/ansh-devs/microservices_project/order-service/db/generated"
	baseproto "github.com/ansh-devs/microservices_project/order-service/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func MustStartGrpcServer(dbConn *db.Queries, rpcAddr string) {

	lis, err := net.Listen("tcp", rpcAddr)
	if err != nil {
		log.Fatal(err)
	}
	srvr := baserepo.NewGrpcServer(dbConn)
	grpcInterceptor := grpc.UnaryInterceptor(baserepo.GrpcInterceptor)
	grpcServer := grpc.NewServer(grpcInterceptor)
	baseproto.RegisterOrderServiceServer(grpcServer, srvr)
	fmt.Printf("[INFO]: gRPC server started at : %s \n", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

package gapi

import (
	db "github.com/ansh-devs/microservices_project/login-service/db/generated"
	baseproto "github.com/ansh-devs/microservices_project/login-service/proto"
)

type LoginRpcServer struct {
	*db.Queries
	baseproto.UnimplementedLoginServiceServer
}

func NewGrpcServer(pgConn *db.Queries) *LoginRpcServer {
	return &LoginRpcServer{
		Queries: pgConn,
	}
}

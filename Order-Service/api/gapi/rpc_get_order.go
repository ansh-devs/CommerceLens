package gapi

import (
	"context"
	baseproto "github.com/ansh-devs/microservices_project/order-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (srv *OrderRpcServer) GetOrder(ctx context.Context, req *baseproto.GetIDRequest) (*baseproto.GetOrderResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}

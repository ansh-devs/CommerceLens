package gapi

import (
	"context"
	baseproto "github.com/ansh-devs/microservices_project/order-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (srv *OrderRpcServer) CancelOrder(ctx context.Context, req *baseproto.GetIDRequest) (*baseproto.CancelOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}

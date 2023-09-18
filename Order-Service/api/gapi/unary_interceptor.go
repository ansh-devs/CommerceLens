package gapi

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func GrpcInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Println("received Grpc Request")
	return handler(ctx, req)
}

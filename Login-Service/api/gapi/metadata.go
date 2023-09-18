package gapi

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
)

const (
	UserAgent = "grpcgateway-user-agent"
	ClientIp  = "x-forwarded-for"
)

func (l *LoginRpcServer) extractmetadata(ctx context.Context) {
	if mtdt, ok := metadata.FromIncomingContext(ctx); ok {
		fmt.Printf("[INFO]: UserAgent : %s\n", mtdt[UserAgent])
		fmt.Printf("[INFO]: ClientIp : %s\n", mtdt[ClientIp])
	}

}

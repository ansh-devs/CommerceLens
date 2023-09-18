package gapi

import (
	"context"
	"fmt"
	baseproto "github.com/ansh-devs/microservices_project/login-service/proto"
	"github.com/ansh-devs/microservices_project/login-service/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (l *LoginRpcServer) GetUserDetails(ctx context.Context, req *baseproto.GetUserDetailsRequest) (*baseproto.GetUserDetailsResponse, error) {
	//l.extractmetadata(ctx)
	fmt.Println(req.GetAccessToken())
	decrypter, err := token.TokenDecrypter(req.GetAccessToken())
	fmt.Printf("%+v", decrypter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't find the user : %s", err)
	}
	model, err := l.CheckUserById(ctx, decrypter.UserId)
	return &baseproto.GetUserDetailsResponse{
		UserData: &baseproto.User{
			Id:        model.ID,
			Email:     model.Email,
			Fullname:  model.Email,
			Address:   model.Address,
			CreatedAt: nil,
		},
		Msg: "message",
	}, nil
	//return nil, status.Errorf(codes.Unimplemented, "method GetUserDetails not implemented")
}

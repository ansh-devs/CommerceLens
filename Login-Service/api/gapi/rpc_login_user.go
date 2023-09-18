package gapi

import (
	"context"
	"fmt"
	proto "github.com/ansh-devs/microservices_project/login-service/proto"
	"github.com/ansh-devs/microservices_project/login-service/token"
	"github.com/golang/protobuf/ptypes/timestamp"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (l *LoginRpcServer) LoginUser(ctx context.Context, req *proto.LoginUserRequest) (*proto.LoginUserResponse, error) {
	// checking if the user with given email exists on database...
	model, err := l.CheckUserByEmail(ctx, req.Email)
	if err != nil {
		fmt.Println(err)
	}
	// hashing the password from the request...
	hashedpwderr := bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(req.Password))
	if err != nil {
		fmt.Println(err)
	}
	// comparing hashed password with that which is stored in database...
	if hashedpwderr == nil {
		// Authentication Successful...
		generatedtoken, err := token.GenerateToken(model.ID)
		if err != nil {
			fmt.Println(err)
		}

		decrypter, _ := token.TokenDecrypter(generatedtoken)
		fmt.Printf(" Decrypted Text : %+v", decrypter)

		user := proto.User{
			Id:       model.ID,
			Fullname: model.Fullname,
			Email:    model.Email,
			Address:  model.Address,
			CreatedAt: &timestamp.Timestamp{
				Seconds: 23,
				Nanos:   42,
			},
		}
		return &proto.LoginUserResponse{
			User:        &user,
			AccessToken: generatedtoken,
			Msg:         "Welcome World",
		}, nil
	} else {
		return nil, status.Errorf(codes.Unauthenticated, "Password Does Not Match")
	}
}

package gapi

import (
	"context"
	"fmt"
	db "github.com/ansh-devs/microservices_project/login-service/db/generated"
	baseproto "github.com/ansh-devs/microservices_project/login-service/proto"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (l *LoginRpcServer) RegisterUser(rpcctx context.Context, req *baseproto.RegisterUserRequest) (*baseproto.RegisterUserResponse, error) {
	tm := time.Now()
	id := uuid.NewString()

	count, err := l.CheckUserIsRegistered(rpcctx, req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't verify if this user is already registered. : %s ", err)
	}
	if count == 0 {
		hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.MinCost)
		user, err := l.CreateUser(rpcctx, db.CreateUserParams{
			ID:        id,
			Fullname:  req.GetName(),
			Email:     req.GetEmail(),
			Password:  string(hashedPwd),
			Address:   req.GetAddress(),
			CreatedAt: time.Now(),
		})
		if err != nil {
			return nil, err
		}
		fmt.Printf("user_id : %s \n", id)
		fmt.Printf("time taken : %s \n", time.Since(tm))
		return &baseproto.RegisterUserResponse{
			User: &baseproto.User{
				Id:        user.ID,
				Fullname:  user.Fullname,
				Email:     user.Email,
				Password:  user.Password,
				Address:   user.Address,
				CreatedAt: nil,
			},
			Msg: "user created successfully.",
		}, nil
	} else {
		return nil, status.Errorf(codes.AlreadyExists, "User with this email is already registered.")
	}
}

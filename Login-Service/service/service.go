package service

import (
	"context"
	"github.com/ansh-devs/commercelens/login-service/dto"
)

type Service interface {
	RegisterUser(ctx context.Context, payload dto.RegisterUserRequest) (dto.RegisterUserResponse, error)
	LoginUser(ctx context.Context, email, password string) (dto.LoginUserResponse, error)
	GetUserDetails(ctx context.Context, accessToken string) (dto.GetUserDetailsResponse, error)
	GetUserWithNats(ctx context.Context)
	RegisterService(addr *string)
	UpdateHealthStatus()
}

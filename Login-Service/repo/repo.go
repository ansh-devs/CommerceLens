package repo

import (
	"context"
	"github.com/ansh-devs/commercelens/login-service/dto"
)

type Repository interface {
	CreateUser(ctx context.Context, usr dto.RegisterUserRequest) (dto.User, error)
	GetUser(ctx context.Context, id string) (dto.User, error)
	CheckUserByEmail(ctx context.Context, email string) (dto.User, error)
}

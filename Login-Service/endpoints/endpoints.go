package endpoints

import (
	"context"
	"github.com/ansh-devs/ecomm-poc/login-service/dto"
	"github.com/ansh-devs/ecomm-poc/login-service/service"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	RegisterUser   endpoint.Endpoint
	LoginUser      endpoint.Endpoint
	GetUserDetails endpoint.Endpoint
}

func NewEndpoints(s service.Service) *Endpoints {
	return &Endpoints{
		RegisterUser:   makeRegisterUserEndpoint(s),
		LoginUser:      makeLoginUserEndpoint(s),
		GetUserDetails: makeGetUserDetailsEndpoint(s),
	}
}

func makeRegisterUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.RegisterUserRequest)
		ok, err := s.RegisterUser(ctx, req)
		return ok, err
	}
}

func makeLoginUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.LoginUserRequest)
		ok, err := s.LoginUser(ctx, req.Email, req.Password)
		return ok, err
	}
}

func makeGetUserDetailsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.GetUserDetailsRequest)
		ok, err := s.GetUserDetails(ctx, req.AccessToken)
		return ok, err
	}
}

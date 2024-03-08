package transport

import (
	"context"
	"github.com/ansh-devs/microservices_project/login-service/endpoints"
	transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHttpServer(ctx context.Context, endpoints *endpoints.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(JsonTypeReWrittermiddleware)

	r.Methods("POST").Path("/user/v1/register").Handler(transport.NewServer(
		endpoints.RegisterUser,
		JsonRegisterUserRequestDecoder,
		JsonResponseEncoder,
	))

	r.Methods("POST").Path("/user/v1/login").Handler(transport.NewServer(
		endpoints.LoginUser,
		JsonLoginUserRequestDecoder,
		JsonResponseEncoder,
	))

	r.Methods("GET").Path("/user/v1/get-user/{id}").Handler(transport.NewServer(
		endpoints.GetUserDetails,
		JsonGetUserDetailsRequestDecoder,
		JsonResponseEncoder,
	))

	return r
}

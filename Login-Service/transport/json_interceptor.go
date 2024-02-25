package transport

import (
	"context"
	"encoding/json"
	"github.com/ansh-devs/microservices_project/login-service/dto"
	"github.com/gorilla/mux"
	"net/http"
)

func JsonResponseEncoder(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	return json.NewEncoder(w).Encode(resp)
}

func RegisterJsonResponseEncoder(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(resp)
}

func JsonRegisterUserRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func JsonRegisterUserResponseDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.RegisterUserResponse
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func JsonLoginUserRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.LoginUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func JsonLoginUserResponseDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.LoginUserResponse
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func JsonGetUserDetailsRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.GetUserDetailsRequest
	vars := mux.Vars(r)
	req.AccessToken = vars["id"]
	return req, nil
}

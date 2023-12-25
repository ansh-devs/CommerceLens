package Order_Service

import (
	"context"
	"encoding/json"
	"github.com/ansh-devs/microservices_project/order-service/dto"
	"net/http"
)

func JsonResponseEncoder(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	return json.NewEncoder(w).Encode(resp)
}

func JsonPlaceOrderResponseDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.PlaceOrderReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func JsonGetOrderResponseDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.GetOrderReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func JsonCancelOrderResponseDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.CancelOrderReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func JsonGetAllUserOrdersResponseDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.GetAllUserOrdersReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

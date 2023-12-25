package order

import (
	"context"
	"encoding/json"
	"net/http"
)

func JsonResponseEncoder(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	return json.NewEncoder(w).Encode(resp)
}

func JsonPlaceOrderResponseDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var req PlaceOrderReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func JsonGetOrderResponseDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetOrderReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func JsonCancelOrderResponseDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CancelOrderReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func JsonGetAllUserOrdersResponseDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetAllUserOrdersReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

package transport

import (
	"context"
	"encoding/json"
	"github.com/ansh-devs/ecomm-poc/order-service/dto"
	"github.com/gorilla/mux"
	"net/http"
)

func JsonResponseEncoder(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	return json.NewEncoder(w).Encode(resp)
}

func JsonPlaceOrderResponseDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.PlaceOrderReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func JsonGetOrderResponseDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.GetOrderReq
	vars := mux.Vars(r)
	//	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	//	return nil, err
	//  }
	req.OrderID = vars["id"]
	return req, nil
}

func JsonCancelOrderResponseDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.CancelOrderReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func JsonGetAllUserOrdersResponseDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.GetAllUserOrdersReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

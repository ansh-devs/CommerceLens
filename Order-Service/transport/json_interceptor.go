package transport

import (
	"context"
	"encoding/json"
	"github.com/ansh-devs/commercelens/order-service/dto"
	"github.com/gorilla/mux"
	"net/http"
)

func JsonResponseEncoder(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	return json.NewEncoder(w).Encode(resp)
}

func JsonGetOrderResponseDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.GetOrderReq
	vars := mux.Vars(r)
	req.OrderID = vars["id"]
	return req, nil
}

func JsonCancelOrderResponseDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.CancelOrderReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func JsonGetAllUserOrdersResponseDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.GetAllUserOrdersReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

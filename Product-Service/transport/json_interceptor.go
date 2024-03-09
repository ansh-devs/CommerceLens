package transport

import (
	"context"
	"encoding/json"
	"github.com/ansh-devs/ecomm-poc/product-service/dto"
	"github.com/gorilla/mux"
	"net/http"
)

func JsonResponseEncoder(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	return json.NewEncoder(w).Encode(resp)
}

func JsonGetProductByIDRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.GetProductReq
	vars := mux.Vars(r)
	req.ProductID = vars["id"]
	return req, nil
}

func JsonGetAllProductsResponseDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}

func JsonPurchaseProductResponseDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.GetAllProductsResp
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

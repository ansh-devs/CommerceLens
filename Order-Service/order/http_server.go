package order

import (
	"context"
	transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHttpServer(ctx context.Context, endpoints *Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(JsonTypeReWrittermiddleware)

	r.Methods("GET").Path("/v1/orders/place-order").Handler(transport.NewServer(
		endpoints.PlaceOrder,
		JsonPlaceOrderResponseDecoder,
		JsonResponseEncoder,
	))

	r.Methods("POST").Path("/v1/orders/get-order").Handler(transport.NewServer(
		endpoints.GetOrder,
		JsonGetOrderResponseDecoder,
		JsonResponseEncoder,
	))

	r.Methods("GET").Path("/v1/orders/cancel-order").Handler(transport.NewServer(
		endpoints.PlaceOrder,
		JsonCancelOrderResponseDecoder,
		JsonResponseEncoder,
	))

	r.Methods("GET").Path("/v1/orders/get-user-all-orders").Handler(transport.NewServer(
		endpoints.PlaceOrder,
		JsonGetAllUserOrdersResponseDecoder,
		JsonResponseEncoder,
	))
	return r
}

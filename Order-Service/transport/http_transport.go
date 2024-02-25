package transport

import (
	"context"
	"github.com/ansh-devs/microservices_project/order-service/endpoints"
	transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHttpServer(_ context.Context, endpoints *endpoints.HttpEndpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(JsonTypeReWrittermiddleware)

	/*	r.Methods("POST").Path("/v1/orders/place-order").Handler(transport.NewServer(
		endpoints.PlaceOrder,
		JsonPlaceOrderResponseDecoder,
		JsonResponseEncoder,
	))*/

	r.Methods("GET").Path("/v1/orders/get-order/{id}").Handler(transport.NewServer(
		endpoints.GetOrder,
		JsonGetOrderResponseDecoder,
		JsonResponseEncoder,
	))

	r.Methods("GET").Path("/v1/orders/cancel-order").Handler(transport.NewServer(
		endpoints.CancelOrder,
		JsonCancelOrderResponseDecoder,
		JsonResponseEncoder,
	))

	r.Methods("GET").Path("/v1/orders/get-user-all-orders").Handler(transport.NewServer(
		endpoints.GetAllUserOrders,
		JsonGetAllUserOrdersResponseDecoder,
		JsonResponseEncoder,
	))
	return r
}

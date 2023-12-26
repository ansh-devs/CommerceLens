package transport

import (
	"context"
	"github.com/ansh-devs/microservices_project/order-service"
	"github.com/ansh-devs/microservices_project/order-service/endpoints"
	transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHttpServer(ctx context.Context, endpoints *endpoints.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(Order_Service.JsonTypeReWrittermiddleware)

	r.Methods("GET").Path("/v1/orders/place-order").Handler(transport.NewServer(
		endpoints.PlaceOrder,
		Order_Service.JsonPlaceOrderResponseDecoder,
		Order_Service.JsonResponseEncoder,
	))

	r.Methods("POST").Path("/v1/orders/get-order").Handler(transport.NewServer(
		endpoints.GetOrder,
		Order_Service.JsonGetOrderResponseDecoder,
		Order_Service.JsonResponseEncoder,
	))

	r.Methods("GET").Path("/v1/orders/cancel-order").Handler(transport.NewServer(
		endpoints.PlaceOrder,
		Order_Service.JsonCancelOrderResponseDecoder,
		Order_Service.JsonResponseEncoder,
	))

	r.Methods("GET").Path("/v1/orders/get-user-all-orders").Handler(transport.NewServer(
		endpoints.PlaceOrder,
		Order_Service.JsonGetAllUserOrdersResponseDecoder,
		Order_Service.JsonResponseEncoder,
	))
	return r
}

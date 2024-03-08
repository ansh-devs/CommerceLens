package transport

import (
	"context"
	"github.com/ansh-devs/microservices_project/product-service/endpoints"
	transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHttpServer(ctx context.Context, endpoints *endpoints.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(JsonTypeReWrittermiddleware)

	r.Methods("POST").Path("/products/v1/purchase").Handler(transport.NewServer(
		endpoints.PurchaseProduct,
		JsonPurchaseProductResponseDecoder,
		JsonResponseEncoder,
	))

	r.Methods("GET").Path("/products/v1/get-product/{id}").Handler(transport.NewServer(
		endpoints.GetProductById,
		JsonGetProductByIDRequestDecoder,
		JsonResponseEncoder,
	))

	r.Methods("GET").Path("/products/v1/get-all").Handler(transport.NewServer(
		endpoints.GetAllProducts,
		JsonGetAllProductsResponseDecoder,
		JsonResponseEncoder,
	))
	return r
}

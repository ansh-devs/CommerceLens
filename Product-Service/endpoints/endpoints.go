package endpoints

import (
	"context"
	"fmt"
	"github.com/ansh-devs/microservices_project/product-service/dto"
	"github.com/ansh-devs/microservices_project/product-service/service"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints - endpoint layer for http
type Endpoints struct {
	GetProductById  endpoint.Endpoint
	GetAllProducts  endpoint.Endpoint
	PurchaseProduct endpoint.Endpoint
}

func NewEndpoints(s service.Service) *Endpoints {
	return &Endpoints{
		GetProductById:  makeGetProductByIdEndpoint(s),
		GetAllProducts:  makeGetAllProductsEndpoint(s),
		PurchaseProduct: makePurchaseProductEndpoint(s),
	}
}

func makeGetProductByIdEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.GetProductReq)
		ok, err := s.GetProductById(ctx, req.ProductID)
		return ok, err
	}
}

func makeGetAllProductsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		ok, err := s.GetAllProducts(ctx)
		if err != nil {

			fmt.Println("Error makeGetAll Ke Andar" + err.Error())
		}
		return ok, err
	}
}

func makePurchaseProductEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.PurchaseOrderReq)
		err = s.PurchaseProduct(ctx, req.UserAccessToken, req.ProductID)
		return nil, err
	}
}

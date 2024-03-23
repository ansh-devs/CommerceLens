package endpoints

import (
	"context"
	"github.com/ansh-devs/commercelens/product-service/dto"
	"github.com/ansh-devs/commercelens/product-service/service"
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
		resp := new(dto.GetProductResp)
		if err == nil {
			resp.Product = dto.Product{
				ID:          ok.ID,
				ProductName: ok.ProductName,
				Description: ok.Description,
				Price:       ok.Price,
			}
			resp.Status = "successful"
			resp.Message = "Item fetched Successfully"
			return resp, err
		} else {
			resp.Status = "failed"
			resp.Message = err.Error()
			return resp, err
		}
	}
}

func makeGetAllProductsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		resp := new(dto.GetAllProductsResp)
		prod, err := s.GetAllProducts(ctx)
		if err == nil {
			for _, v := range prod {
				resp.Product = append(resp.Product, dto.Product{
					ID:          v.ID,
					ProductName: v.ProductName,
					Description: v.Description,
					Price:       v.Price,
				})
			}
			resp.Status = "successful"
			resp.Message = "Items fetched Successfully"
			return resp, err
		} else {
			resp.Status = "failed"
			resp.Message = err.Error()
			return resp, err
		}
	}
}

func makePurchaseProductEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.PurchaseOrderReq)
		prod, err := s.PurchaseProduct(ctx, req.UserAccessToken, req.ProductID)
		return prod, err
	}
}

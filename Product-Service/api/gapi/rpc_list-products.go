package gapi

import (
	"context"
	baseproto "github.com/ansh-devs/microservices_project/product-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *ProductRpcServer) GetAllProducts(ctx context.Context, req *baseproto.GetAllProductsRequest) (*baseproto.GetAllProductsResponse, error) {
	resp, err := p.Queries.GetAllProducts(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "can't fetch products. %s", err)
	}

	payload := make([]*baseproto.AddProductResponse, 0)
	for _, v := range resp {
		payload = append(payload, &baseproto.AddProductResponse{
			Id:          v.ID,
			ProductName: v.ProductName,
			Description: v.Description,
			Price:       v.Price,
		})
	}

	return &baseproto.GetAllProductsResponse{
		Status:   "ok",
		Products: payload,
	}, nil
}

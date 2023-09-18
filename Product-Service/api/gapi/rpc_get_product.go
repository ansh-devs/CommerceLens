package gapi

import (
	"context"
	baseproto "github.com/ansh-devs/microservices_project/product-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *ProductRpcServer) GetProduct(ctx context.Context, req *baseproto.GetProductRequest) (*baseproto.GetProductResponse, error) {

	product, err := p.Queries.GetProductById(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "can't fetch the product.")
	}
	return &baseproto.GetProductResponse{
		Id:          product.ID,
		ProductName: product.ProductName,
		Description: product.Description,
		Price:       product.Price,
	}, nil

}

package gapi

import (
	"context"
	db "github.com/ansh-devs/microservices_project/product-service/db/generated"
	baseproto "github.com/ansh-devs/microservices_project/product-service/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (p *ProductRpcServer) AddProduct(ctx context.Context, req *baseproto.AddProductRequest) (*baseproto.AddProductResponse, error) {
	tm := time.Now()
	id := uuid.NewString()
	resp, err := p.Queries.CreateProduct(ctx, db.CreateProductParams{
		ID:          id,
		ProductName: req.ProductName,
		Description: req.Description,
		Price:       req.Price,
		CreatedAt:   tm,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't add the product. %s", err)
	}
	return &baseproto.AddProductResponse{
		Id:          resp.ID,
		ProductName: resp.ProductName,
		Description: resp.Description,
		Price:       resp.Price,
	}, nil
	//return nil, status.Errorf(codes.NotFound, "can't fetch products.")
}

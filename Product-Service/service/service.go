package service

import (
	"context"
	db "github.com/ansh-devs/microservices_project/product-service/db/generated"
)

type Service interface {
	PurchaseProduct(ctx context.Context, userId, productId string) error
	GetProductById(ctx context.Context, productId string) (db.Product, error)
	GetAllProducts(ctx context.Context) ([]db.Product, error)
}

package service

import (
	"context"
	db "github.com/ansh-devs/ecomm-poc/product-service/db/generated"
)

type Service interface {
	PurchaseProduct(ctx context.Context, userId, productId string) (db.Product, error)
	GetProductById(ctx context.Context, productId string) (db.Product, error)
	GetAllProducts(ctx context.Context) ([]db.Product, error)
}

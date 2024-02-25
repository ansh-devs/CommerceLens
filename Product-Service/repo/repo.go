package repo

import (
	"context"
	db "github.com/ansh-devs/microservices_project/product-service/db/generated"
	"github.com/opentracing/opentracing-go"
)

type Repository interface {
	GetProductByID(ctx context.Context, prodID string, span opentracing.Span) (db.Product, error)
	GetAllProducts(ctx context.Context, span opentracing.Span) ([]db.Product, error)
}

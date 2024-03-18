package repo

import (
	"context"
	db "github.com/ansh-devs/ecomm-poc/product-service/db/generated"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/opentracing/opentracing-go"
)

type PGRepo struct {
	db     *db.Queries
	logger log.Logger
	trace  opentracing.Tracer
}

func NewRepo(db *db.Queries, logger log.Logger, tracer opentracing.Tracer) *PGRepo {
	return &PGRepo{
		db:     db,
		logger: log.With(logger, "layer", "repository"),
		trace:  tracer,
	}
}

func (r *PGRepo) GetProductByID(ctx context.Context, prodID string, span opentracing.Span) (db.Product, error) {
	_ = level.Info(r.logger).Log("method-invoked", "GetProductByID")
	sp := r.trace.StartSpan("get_product_by_id_db_layer_call", opentracing.ChildOf(span.Context()))
	defer sp.Finish()
	prods, err := r.db.GetProductById(ctx, prodID)
	if err != nil {
		return db.Product{}, err
	}
	return prods, nil
}
func (r *PGRepo) GetAllProducts(ctx context.Context, span opentracing.Span) ([]db.Product, error) {
	_ = level.Info(r.logger).Log("method-invoked", "GetAllProducts")
	sp := r.trace.StartSpan("get_all_products_db_layer_call", opentracing.ChildOf(span.Context()))
	defer sp.Finish()
	prods, err := r.db.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}
	return prods, nil
}

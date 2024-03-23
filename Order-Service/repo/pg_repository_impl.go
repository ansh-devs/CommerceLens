package repo

import (
	"context"
	"errors"
	"time"

	db "github.com/ansh-devs/commercelens/order-service/db/generated"
	"github.com/ansh-devs/commercelens/order-service/dto"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
)

var (
	ErrRepo = errors.New("unable to handle repo request")
)

type PGRepo struct {
	db.Queries
	log.Logger
	opentracing.Tracer
}

func NewPostgresRepository(db *db.Queries, logger log.Logger, tracer opentracing.Tracer) Repository {
	return &PGRepo{*db, log.With(logger, "layer", "repository"), tracer}
}

func (r *PGRepo) PlaceOrder(ctx context.Context, product dto.Product, user dto.NatsUser, span opentracing.Span) (dto.Order, error) {
	tm := time.Now()
	id := uuid.NewString()
	createdOrder, err := r.Queries.CreateOrder(ctx, db.CreateOrderParams{
		ID:          id,
		ProductID:   product.ID,
		UserID:      user.ID,
		TotalCost:   product.Price,
		Status:      "placed",
		Fullname:    user.FullName,
		Address:     user.FullName,
		ProductName: product.ProductName,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   tm,
	})

	if err != nil {
		return dto.Order{}, err
	}
	_ = level.Info(r.Logger).Log("method-invoked", "PlaceOrder")
	return dto.Order{
		ID:              createdOrder.ID,
		ProductID:       createdOrder.ProductID,
		UserID:          createdOrder.UserID,
		TotalCost:       createdOrder.TotalCost,
		Username:        createdOrder.Fullname,
		ProductName:     createdOrder.ProductName,
		Description:     createdOrder.Description,
		Price:           createdOrder.Price,
		ShippingAddress: createdOrder.Address,
	}, nil

}

func (r *PGRepo) CancelOrder(ctx context.Context, orderId string, span opentracing.Span) (string, error) {
	sp := r.Tracer.StartSpan(
		"cancel-order-db-call",
		opentracing.ChildOf(span.Context()))
	defer sp.Finish()
	_ = level.Info(r.Logger).Log("method-invoked", "CancelOrder")
	err := r.Queries.ChangeOrderStatusById(ctx, db.ChangeOrderStatusByIdParams{
		ID:     orderId,
		Status: "canceled",
	})
	if err != nil {
		return "failed", err
	}
	return "success", nil
}

func (r *PGRepo) GetUserAllOrders(ctx context.Context, userId string, span opentracing.Span) ([]dto.Order, error) {
	sp := r.Tracer.StartSpan(
		"get-all-user-orders-db-call",
		opentracing.ChildOf(span.Context()))
	defer sp.Finish()
	_ = level.Info(r.Logger).Log("method-invoked", "GetUserAllOrders")
	result, err := r.Queries.GetAllOrdersByUserId(ctx, userId)
	if err != nil {
		return []dto.Order{}, err
	}

	var resp []dto.Order

	for _, v := range result {
		model := dto.Order{
			ID:              v.ID,
			ProductID:       v.ProductID,
			UserID:          v.UserID,
			TotalCost:       v.TotalCost,
			Username:        v.Fullname,
			ProductName:     v.ProductName,
			Description:     v.Description,
			Price:           v.Price,
			Status:          v.Status,
			ShippingAddress: v.Address,
		}
		resp = append(resp, model)
	}

	return resp, nil
}

func (r *PGRepo) GetOrder(ctx context.Context, orderId string, span opentracing.Span) (dto.Order, error) {
	if orderId == "" {
		return dto.Order{}, errors.New("empty order id passed")
	}
	sp := r.Tracer.StartSpan(
		"get-order-db-call",
		opentracing.ChildOf(span.Context()))
	defer sp.Finish()
	_ = level.Info(r.Logger).Log("method-invoked", "GetOrder for "+orderId)
	result, err := r.Queries.GetOrderById(ctx, orderId)
	if err != nil {
		return dto.Order{}, err
	}

	return dto.Order{
		ID:              result.ID,
		ProductID:       result.ProductID,
		UserID:          result.UserID,
		TotalCost:       result.TotalCost,
		Username:        result.Fullname,
		ProductName:     result.ProductName,
		Description:     result.Description,
		Price:           result.Price,
		ShippingAddress: result.Address,
	}, nil
}

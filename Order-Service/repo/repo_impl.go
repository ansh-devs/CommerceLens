package repo

import (
	"context"
	"errors"
	db "github.com/ansh-devs/microservices_project/order-service/db/generated"
	"github.com/ansh-devs/microservices_project/order-service/dto"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"time"
)

var (
	RepoErr = errors.New("unable to handle repo request")
)

type Repo struct {
	db     *db.Queries
	logger log.Logger
	trace  opentracing.Tracer
}

func NewRepo(db *db.Queries, logger log.Logger, tracer opentracing.Tracer) *Repo {
	return &Repo{db: db, logger: log.With(logger, "layer", "repository"), trace: tracer}
}

func (r *Repo) PlaceOrder(ctx context.Context, product dto.Product, user dto.NatsUser, span opentracing.Span) (dto.Order, error) {
	tm := time.Now()
	id := uuid.NewString()
	createdOrder, err := r.db.CreateOrder(ctx, db.CreateOrderParams{
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
	_ = level.Info(r.logger).Log("method-invoked", "PlaceOrder")
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

func (r *Repo) CancelOrder(ctx context.Context, orderId string, span opentracing.Span) (string, error) {
	sp := r.trace.StartSpan(
		"cancel-order-db-call",
		opentracing.ChildOf(span.Context()))
	defer sp.Finish()
	_ = level.Info(r.logger).Log("method-invoked", "CancelOrder")
	err := r.db.ChangeOrderStatusById(ctx, db.ChangeOrderStatusByIdParams{
		ID:     orderId,
		Status: "canceled",
	})
	if err != nil {
		return "failed", err
	}
	return "success", nil
}

func (r *Repo) GetUserAllOrders(ctx context.Context, userId string, span opentracing.Span) ([]dto.Order, error) {
	sp := r.trace.StartSpan(
		"get-all-user-orders-db-call",
		opentracing.ChildOf(span.Context()))
	defer sp.Finish()
	_ = level.Info(r.logger).Log("method-invoked", "GetUserAllOrders")
	result, err := r.db.GetAllOrdersByUserId(ctx, userId)
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

func (r *Repo) GetOrder(ctx context.Context, orderId string, span opentracing.Span) (dto.Order, error) {
	sp := r.trace.StartSpan(
		"get-order-db-call",
		opentracing.ChildOf(span.Context()))
	defer sp.Finish()
	_ = level.Info(r.logger).Log("method-invoked", "GetOrder for "+orderId)
	result, err := r.db.GetOrderById(ctx, orderId)
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

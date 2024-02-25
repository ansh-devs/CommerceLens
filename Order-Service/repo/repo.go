package repo

import (
	"context"
	"github.com/ansh-devs/microservices_project/order-service/dto"
	"github.com/opentracing/opentracing-go"
)

type Repository interface {
	PlaceOrder(ctx context.Context, product dto.Product, userId dto.NatsUser, span opentracing.Span) (dto.Order, error)
	CancelOrder(ctx context.Context, orderId string, span opentracing.Span) (string, error)
	GetUserAllOrders(ctx context.Context, userId string, span opentracing.Span) ([]dto.Order, error)
	GetOrder(ctx context.Context, orderId string, span opentracing.Span) (dto.Order, error)
}

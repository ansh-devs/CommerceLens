package repo

import (
	"context"

	"github.com/ansh-devs/commercelens/order-service/dto"
	"github.com/opentracing/opentracing-go"
)

// Repository is the BaseClass for Repository related operations.
type Repository interface {
	// PlaceOrder Places an order for the given user with the given product.
	PlaceOrder(ctx context.Context, product dto.Product, userId dto.NatsUser, span opentracing.Span) (dto.Order, error)
	// CancelOrder cancels the order.
	CancelOrder(ctx context.Context, orderId string, span opentracing.Span) (string, error)
	//GetUserAllOrders fetches all the orders for the given user.
	GetUserAllOrders(ctx context.Context, userId string, span opentracing.Span) ([]dto.Order, error)
	// GetOrder fetches the single order with the provided orderId.
	GetOrder(ctx context.Context, orderId string, span opentracing.Span) (dto.Order, error)
}

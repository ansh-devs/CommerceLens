package order

import "context"

type Repository interface {
	PlaceOrder(ctx context.Context, orderId string) (Order, error)
	CancelOrder(ctx context.Context, orderId string) (string, error)
	GetUserAllOrders(ctx context.Context, userId string) ([]Order, error)
	GetOrder(ctx context.Context, orderId string) (Order, error)
}

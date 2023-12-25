package repo

import (
	"context"
	"github.com/ansh-devs/microservices_project/order-service/dto"
)

type Repository interface {
	PlaceOrder(ctx context.Context, orderId string) (dto.Order, error)
	CancelOrder(ctx context.Context, orderId string) (string, error)
	GetUserAllOrders(ctx context.Context, userId string) ([]dto.Order, error)
	GetOrder(ctx context.Context, orderId string) (dto.Order, error)
}

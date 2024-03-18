package service

import (
	"context"
	"github.com/ansh-devs/ecomm-poc/order-service/dto"
)

// Service is the Base Class for the service related operations.
type Service interface {
	// PlaceOrder Places an order
	PlaceOrder(ctx context.Context)
	// GetOrder fetches an order for the given productId.
	GetOrder(ctx context.Context, productId string) (dto.GetOrderResp, error)
	// CancelOrder cancels the order with the given orderId.
	CancelOrder(ctx context.Context, orderId string) (dto.CancelOrderResp, error)
	// GetAllUserOrders Fetches all the users orders with the provided userId.
	GetAllUserOrders(ctx context.Context, userId string) (dto.GetAllUserOrdersResp, error)
	// RegisterService Registers the service in consul cluster with the port provided.
	RegisterService(addr *string)
	// UpdateHealthStatus updates the service status in consul's registry.
	UpdateHealthStatus()
}

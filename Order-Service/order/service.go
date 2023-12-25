package order

import (
	"context"
)

type Service interface {
	PlaceOrder(ctx context.Context, productId string) (PlaceOrderResp, error)
	GetOrder(ctx context.Context, productId string) (GetOrderResp, error)
	CancelOrder(ctx context.Context, productId string) (CancelOrderResp, error)
	GetAllUserOrders(ctx context.Context, userId string) (GetAllUserOrdersResp, error)
	RegisterService(addr *string)
	UpdateHealthStatus()
}

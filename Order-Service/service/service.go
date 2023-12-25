package service

import (
	"context"
	"github.com/ansh-devs/microservices_project/order-service/dto"
)

type Service interface {
	PlaceOrder(ctx context.Context, productId string) (dto.PlaceOrderResp, error)
	GetOrder(ctx context.Context, productId string) (dto.GetOrderResp, error)
	CancelOrder(ctx context.Context, productId string) (dto.CancelOrderResp, error)
	GetAllUserOrders(ctx context.Context, userId string) (dto.GetAllUserOrdersResp, error)
	RegisterService(addr *string)
	UpdateHealthStatus()
}

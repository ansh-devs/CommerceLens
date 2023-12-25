package order

import (
	"context"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/hashicorp/consul/api"
)

type OrderService struct {
	repository   Repository
	logger       log.Logger
	consulClient *api.Client
}

func NewService(rep Repository, logger log.Logger) *OrderService {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Println(err)
	}
	return &OrderService{
		repository:   rep,
		logger:       log.With(logger, "layer", "service"),
		consulClient: client,
	}
}

// PlaceOrder - place order wrapper function around the method that makes calls to the repo...
func (s *OrderService) PlaceOrder(ctx context.Context, productId string) (PlaceOrderResp, error) {
	_ = level.Info(s.logger).Log("method-invoked", "PlaceOrder")

	_, err := s.repository.PlaceOrder(ctx, productId)
	if err != nil {
		return PlaceOrderResp{}, err
	}
	return PlaceOrderResp{}, nil
}
func (s *OrderService) GetOrder(ctx context.Context, productId string) (GetOrderResp, error) {
	_ = level.Info(s.logger).Log("method-invoked", "GetOrder")
	_, err := s.repository.GetOrder(ctx, productId)
	if err != nil {
		return GetOrderResp{}, err
	}
	return GetOrderResp{}, nil
}
func (s *OrderService) CancelOrder(ctx context.Context, productId string) (CancelOrderResp, error) {
	_ = level.Info(s.logger).Log("method-invoked", "CancelOrder")
	_, err := s.repository.CancelOrder(ctx, productId)
	if err != nil {
		return CancelOrderResp{}, err
	}
	return CancelOrderResp{
		Status:  "",
		Message: "",
	}, nil
}
func (s *OrderService) GetAllUserOrders(ctx context.Context, userId string) (GetAllUserOrdersResp, error) {
	_ = level.Info(s.logger).Log("method-invoked", "GetAllUserOrders")
	_, err := s.repository.GetUserAllOrders(ctx, userId)
	if err != nil {
		return GetAllUserOrdersResp{}, err
	}
	return GetAllUserOrdersResp{}, nil
}

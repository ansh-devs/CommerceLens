package service

import (
	"context"
	"fmt"
	"github.com/ansh-devs/microservices_project/order-service/dto"
	"github.com/ansh-devs/microservices_project/order-service/repo"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/hashicorp/consul/api"
	"net"
)

type OrderService struct {
	repository   repo.Repository
	logger       log.Logger
	consulClient *api.Client
}

func NewService(rep repo.Repository, logger log.Logger) *OrderService {
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

// PlaceOrder - place dto wrapper function around the method that makes calls to the repo...
func (s *OrderService) PlaceOrder(ctx context.Context, productId string) (dto.PlaceOrderResp, error) {
	_ = level.Info(s.logger).Log("method-invoked", "PlaceOrder")

	_, err := s.repository.PlaceOrder(ctx, productId)
	if err != nil {
		return dto.PlaceOrderResp{}, err
	}
	return dto.PlaceOrderResp{}, nil
}
func (s *OrderService) GetOrder(ctx context.Context, productId string) (dto.GetOrderResp, error) {
	_ = level.Info(s.logger).Log("method-invoked", "GetOrder")
	_, err := s.repository.GetOrder(ctx, productId)
	if err != nil {
		return dto.GetOrderResp{}, err
	}
	return dto.GetOrderResp{}, nil
}
func (s *OrderService) CancelOrder(ctx context.Context, productId string) (dto.CancelOrderResp, error) {
	_ = level.Info(s.logger).Log("method-invoked", "CancelOrder")
	_, err := s.repository.CancelOrder(ctx, productId)
	if err != nil {
		return dto.CancelOrderResp{}, err
	}
	return dto.CancelOrderResp{
		Status:  "",
		Message: "",
	}, nil
}
func (s *OrderService) GetAllUserOrders(ctx context.Context, userId string) (dto.GetAllUserOrdersResp, error) {
	_ = level.Info(s.logger).Log("method-invoked", "GetAllUserOrders")
	_, err := s.repository.GetUserAllOrders(ctx, userId)
	if err != nil {
		return dto.GetAllUserOrdersResp{}, err
	}
	return dto.GetAllUserOrdersResp{}, nil
}

func (s *OrderService) getLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		_ = level.Error(s.logger).Log("err", "can't get local ip")
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

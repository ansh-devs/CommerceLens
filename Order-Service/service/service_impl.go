package service

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/ansh-devs/ecomm-poc/order-service/dto"
	"github.com/ansh-devs/ecomm-poc/order-service/natsutil"
	"github.com/ansh-devs/ecomm-poc/order-service/repo"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/hashicorp/consul/api"
	"github.com/nats-io/nats.go"
	"github.com/opentracing/opentracing-go"
)

// OrderService - Implementation of service interface for business logic...
type OrderService struct {
	SrvID        string
	repository   repo.Repository
	logger       log.Logger
	ConsulClient *api.Client
	trace        opentracing.Tracer
	nats         *natsutil.NATSComponent
}

// NewOrderService - constructor for the OrderService...
func NewOrderService(rep repo.Repository, logger log.Logger, tracer opentracing.Tracer) *OrderService {
	client, err := api.NewClient(&api.Config{
		Address: "service-discovery:8500",
	})
	if err != nil {
		fmt.Println(err)
	}
	srvID := "instance_" + strconv.Itoa(rand.Intn(99))

	return &OrderService{
		repository:   rep,
		logger:       log.With(logger, "layer", "service"),
		ConsulClient: client,
		trace:        tracer,
		SrvID:        srvID,
	}
}

func (r *OrderService) InitNATS() {
	nc := natsutil.NewNatsComponent(r.SrvID)
	err := nc.ConnectToNATS("nats://nats-srvr:4222", nil)
	if err != nil {
		fmt.Println(err)
	}

	r.nats = nc
}

// PlaceOrder - dto wrapper function around the method that makes calls to the repository.
func (s *OrderService) PlaceOrder(ctx context.Context) {
	msgHandler := func(msg *nats.Msg) {
		spn := s.trace.StartSpan("place-order")
		defer spn.Finish()
		_ = level.Info(s.logger).Log("method-invoked", "PlaceOrder")
		order, _ := s.nats.DecryptMsgToOrder(msg.Data)
		buffUserID, err := s.nats.UserIdEncoder(order.UserId)
		if err != nil {
			fmt.Println(err)
		}
		request, err := s.nats.NATS().Request("user.getdetails", buffUserID.Bytes(), time.Second*2)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second * 2)
		user, err := s.nats.DecryptMsgToUser(request.Data)
		if err != nil {
			fmt.Println(err)
		}
		newCtx := opentracing.ContextWithSpan(ctx, spn)
		_, err = s.repository.PlaceOrder(newCtx, order.Product, user, spn)
		if err != nil {
			fmt.Println(err)
		}
	}
	for {
		_, _ = s.nats.NATS().Subscribe("products.purchase", msgHandler)
	}

}

// GetOrder - place dto wrapper function around the method that makes calls to the repo...
func (s *OrderService) GetOrder(ctx context.Context, productId string) (dto.GetOrderResp, error) {
	spn := s.trace.StartSpan("get-order")
	defer spn.Finish()
	_ = level.Info(s.logger).Log("method-invoked", "GetOrder for "+productId)
	newCtx := opentracing.ContextWithSpan(ctx, spn)
	result, err := s.repository.GetOrder(newCtx, productId, spn)
	if err != nil {
		return dto.GetOrderResp{}, err
	}
	return dto.GetOrderResp{
		Status:  "successful",
		Message: "ok",
		Order: dto.Order{
			ID:              result.ID,
			ProductID:       result.ProductID,
			UserID:          result.UserID,
			TotalCost:       result.TotalCost,
			Username:        result.Username,
			ProductName:     result.ProductName,
			Description:     result.Description,
			Price:           result.Price,
			Status:          result.Status,
			ShippingAddress: result.ShippingAddress,
		},
	}, nil
}

// CancelOrder - place dto wrapper function around the method that makes calls to the repo...
func (s *OrderService) CancelOrder(ctx context.Context, orderId string) (dto.CancelOrderResp, error) {
	_ = level.Info(s.logger).Log("method-invoked", "CancelOrder")
	spn := s.trace.StartSpan("cancel-order")
	defer spn.Finish()
	newCtx := opentracing.ContextWithSpan(ctx, spn)
	result, err := s.repository.CancelOrder(newCtx, orderId, spn)
	if err != nil {
		return dto.CancelOrderResp{
			Status:  "failed",
			Message: "can't cancel the order",
		}, nil
	}
	if result == "success" {
		return dto.CancelOrderResp{
			Status:  "successful",
			Message: "Order Cancelled successfully.",
		}, nil
	} else {
		return dto.CancelOrderResp{
			Status:  "failed",
			Message: "can't cancel the order",
		}, nil
	}
}

// GetAllUserOrders - place dto wrapper function around the method that makes calls to the repo...
func (s *OrderService) GetAllUserOrders(ctx context.Context, userId string) (dto.GetAllUserOrdersResp, error) {
	_ = level.Info(s.logger).Log("method-invoked", "GetAllUserOrders")
	spn := s.trace.StartSpan("get-all-user-orders")
	defer spn.Finish()
	newCtx := opentracing.ContextWithSpan(ctx, spn)
	orders, err := s.repository.GetUserAllOrders(newCtx, userId, spn)
	if err != nil {
		return dto.GetAllUserOrdersResp{}, err
	}
	return dto.GetAllUserOrdersResp{
		Status:  "successful",
		Message: "list of all orders for the user",
		Orders:  orders,
	}, nil
}

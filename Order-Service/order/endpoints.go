package order

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints - endpoint layer for http
type Endpoints struct {
	PlaceOrder       endpoint.Endpoint
	GetOrder         endpoint.Endpoint
	CancelOrder      endpoint.Endpoint
	GetAllUserOrders endpoint.Endpoint
}

func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		PlaceOrder:       makePlaceOrderEndpoint(s),
		GetOrder:         makeGetOrderEndpoint(s),
		CancelOrder:      makeCancelOrderEndpoint(s),
		GetAllUserOrders: makeGetAllUserOrdersEndpoint(s),
	}
}

func makeGetAllUserOrdersEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetAllUserOrdersReq)
		ok, err := s.GetAllUserOrders(ctx, req.UserID)
		return ok, err
	}
}

func makeCancelOrderEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CancelOrderReq)
		ok, err := s.PlaceOrder(ctx, req.OrderID)
		return ok, err
	}
}

func makeGetOrderEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetOrderReq)
		ok, err := s.GetOrder(ctx, req.OrderID)
		return ok, err
	}
}

func makePlaceOrderEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(PlaceOrderReq)
		ok, err := s.PlaceOrder(ctx, req.ProductID)
		return ok, err
	}
}

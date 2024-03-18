package endpoints

import (
	"context"
	"github.com/ansh-devs/ecomm-poc/order-service/dto"
	"github.com/ansh-devs/ecomm-poc/order-service/service"
	"github.com/go-kit/kit/endpoint"
)

// HttpEndpoints - endpoint layer for http
type HttpEndpoints struct {
	//PlaceOrder       endpoint.Endpoint
	GetOrder         endpoint.Endpoint
	CancelOrder      endpoint.Endpoint
	GetAllUserOrders endpoint.Endpoint
}

// NewEndpoints - constructor for the HttpEndpoints.
func NewEndpoints(s service.Service) *HttpEndpoints {
	return &HttpEndpoints{
		//PlaceOrder:       makePlaceOrderEndpoint(s),
		GetOrder:         makeGetOrderEndpoint(s),
		CancelOrder:      makeCancelOrderEndpoint(s),
		GetAllUserOrders: makeGetAllUserOrdersEndpoint(s),
	}
}

func makeGetAllUserOrdersEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.GetAllUserOrdersReq)
		ok, err := s.GetAllUserOrders(ctx, req.UserID)
		return ok, err
	}
}

func makeCancelOrderEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.CancelOrderReq)
		ok, err := s.CancelOrder(ctx, req.OrderID)
		return ok, err
	}
}

func makeGetOrderEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.GetOrderReq)
		ok, err := s.GetOrder(ctx, req.OrderID)
		return ok, err
	}
}

/*func makePlaceOrderEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.PlaceOrderReq)
		ok, err := s.PlaceOrder(ctx, req.ProductID, req.UserID)
		return ok, err
	}
}*/

package gapi

import (
	"context"
	"fmt"
	baseproto "github.com/ansh-devs/microservices_project/order-service/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func (srv *OrderRpcServer) GetAllUserOrders(ctx context.Context, req *baseproto.GetIDRequest) (*baseproto.GetAllUserOrdersResponse, error) {
	conn2, err := grpc.Dial("localhost"+srv.LoginServicePORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn2)

	loginsrvc := baseproto.NewLoginServiceClient(conn2)
	details, err := loginsrvc.GetUserDetails(ctx, &baseproto.GetUserDetailsRequest{AccessToken: req.GetId()})
	if err != nil {
		fmt.Println(err)
	}

	orders, err := srv.GetAllOrdersByUserId(ctx, details.UserData.Id)

	if err != nil {
		fmt.Println(err.Error())
	}
	var neworders []*baseproto.Order
	for _, v := range orders {
		tm := v.CreatedAt.(timestamp.Timestamp)
		neworders = append(neworders, &baseproto.Order{
			Id:              v.ID,
			ProductId:       v.ProductID,
			UserId:          v.UserID,
			TotalCost:       v.TotalCost,
			Username:        v.Fullname,
			ProductName:     v.ProductName,
			Description:     v.Description,
			Price:           v.Price,
			ShippingAddress: v.Address,
			CreatedAt:       &tm,
		})
	}

	return &baseproto.GetAllUserOrdersResponse{
		Status: "ok",
		Orders: neworders,
	}, nil

	return nil, status.Errorf(codes.Unimplemented, "method GetAllUserOrders not implemented")
}

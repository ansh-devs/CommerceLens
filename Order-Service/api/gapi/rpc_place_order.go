package gapi

import (
	"context"
	"fmt"
	db "github.com/ansh-devs/microservices_project/order-service/db/generated"
	baseproto "github.com/ansh-devs/microservices_project/order-service/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"time"
)

func (srv *OrderRpcServer) PlaceOrder(ctx context.Context, req *baseproto.PlaceOrderRequest) (*baseproto.PlaceOrderResponse, error) {
	id := uuid.NewString()
	conn, err := grpc.Dial("localhost"+srv.ProductServicePORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)

	productsrvc := baseproto.NewProductServiceClient(conn)

	product, err := productsrvc.GetProduct(ctx, &baseproto.GetProductRequest{Id: req.ProductId})
	if err != nil {
		fmt.Println(err)
	}
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

	user, err := loginsrvc.GetUserDetails(ctx, &baseproto.GetUserDetailsRequest{AccessToken: req.AccessToken})
	if err != nil {
		fmt.Println(err)
	}

	order, err := srv.Queries.CreateOrder(ctx, db.CreateOrderParams{
		ID:          id,
		ProductID:   req.ProductId,
		UserID:      user.UserData.Id,
		TotalCost:   product.Price,
		Status:      "placed",
		Fullname:    user.UserData.Fullname,
		Address:     user.UserData.Address,
		ProductName: product.ProductName,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   time.Now(),
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "error occured %s", err)
	}

	return &baseproto.PlaceOrderResponse{
		Status: "ok",
		Order: &baseproto.Order{
			Id:              order.ID,
			ProductId:       order.ProductID,
			UserId:          order.UserID,
			TotalCost:       order.TotalCost,
			Username:        user.UserData.Fullname,
			ProductName:     product.ProductName,
			Description:     product.Description,
			Price:           product.Price,
			ShippingAddress: user.UserData.Address,
			CreatedAt: &timestamp.Timestamp{
				Seconds: time.Now().Unix(),
				Nanos:   0,
			},
		},
	}, nil
}

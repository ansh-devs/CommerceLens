package repo

import (
	"context"
	"errors"
	db "github.com/ansh-devs/microservices_project/order-service/db/generated"
	"github.com/ansh-devs/microservices_project/order-service/dto"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/google/uuid"
	"time"
)

var (
	RepoErr = errors.New("unable to handle repo request")
)

type Repo struct {
	db     *db.Queries
	logger log.Logger
}

func NewRepo(db *db.Queries, logger log.Logger) *Repo {
	return &Repo{db: db, logger: log.With(logger, "layer", "repository")}
}

func (r *Repo) PlaceOrder(ctx context.Context, orderId, userId string) (dto.Order, error) {
	tm := time.Now()
	id := uuid.NewString()
	createdOrder, err := r.db.CreateOrder(ctx, db.CreateOrderParams{
		ID:          id,
		ProductID:   id,
		UserID:      userId,
		TotalCost:   "$9999",
		Status:      "placed",
		Fullname:    "Username",
		Address:     "mock_address",
		ProductName: "mocked_product_name",
		Description: "description",
		Price:       "$ 9999",
		CreatedAt:   tm,
	})

	if err != nil {
		return dto.Order{}, err
	}
	_ = level.Info(r.logger).Log("method-invoked", "PlaceOrder")
	return dto.Order{
		ID:              createdOrder.ID,
		ProductID:       createdOrder.ProductID,
		UserID:          createdOrder.UserID,
		TotalCost:       createdOrder.TotalCost,
		Username:        createdOrder.Fullname,
		ProductName:     createdOrder.ProductName,
		Description:     createdOrder.Description,
		Price:           createdOrder.Price,
		ShippingAddress: createdOrder.Address,
	}, nil

}

func (r *Repo) CancelOrder(ctx context.Context, orderId string) (string, error) {
	_ = level.Info(r.logger).Log("method-invoked", "CancelOrder")
	// result,err:=r.db.
	return "", nil
}

func (r *Repo) GetUserAllOrders(ctx context.Context, userId string) ([]dto.Order, error) {
	_ = level.Info(r.logger).Log("method-invoked", "GetUserAllOrders")

	result, err := r.db.GetAllOrdersByUserId(ctx, userId)
	if err != nil {
		return []dto.Order{}, err
	}

	var resp []dto.Order

	for _, v := range result {
		model := dto.Order{
			ID:              v.ID,
			ProductID:       v.ProductID,
			UserID:          v.UserID,
			TotalCost:       v.TotalCost,
			Username:        v.Fullname,
			ProductName:     v.ProductName,
			Description:     v.Description,
			Price:           v.Price,
			ShippingAddress: v.Address,
		}
		resp = append(resp, model)
	}

	return resp, nil
}

func (r *Repo) GetOrder(ctx context.Context, orderId string) (dto.Order, error) {
	_ = level.Info(r.logger).Log("method-invoked", "GetOrder")
	result, err := r.db.GetOrderById(ctx, orderId)
	if err != nil {
		return dto.Order{}, err
	}

	return dto.Order{
		ID:              result.ID,
		ProductID:       result.ProductID,
		UserID:          result.UserID,
		TotalCost:       result.TotalCost,
		Username:        result.Fullname,
		ProductName:     result.ProductName,
		Description:     result.Description,
		Price:           result.Price,
		ShippingAddress: result.Address,
	}, nil
}

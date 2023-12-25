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

func (r *Repo) PlaceOrder(ctx context.Context, orderId string) (dto.Order, error) {
	tm := time.Now()
	id := uuid.NewString()
	createdOrder, err := r.db.CreateOrder(ctx, db.CreateOrderParams{
		ID:          id,
		ProductID:   "",
		UserID:      "",
		TotalCost:   "",
		Status:      "",
		Fullname:    "",
		Address:     "",
		ProductName: "",
		Description: "",
		Price:       "",
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

	return "", nil
}

func (r *Repo) GetUserAllOrders(ctx context.Context, userId string) ([]dto.Order, error) {
	_ = level.Info(r.logger).Log("method-invoked", "GetUserAllOrders")
	return []dto.Order{}, nil
}

func (r *Repo) GetOrder(ctx context.Context, orderId string) (dto.Order, error) {
	_ = level.Info(r.logger).Log("method-invoked", "GetOrder")

	return dto.Order{}, nil
}

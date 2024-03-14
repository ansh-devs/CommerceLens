package service

import (
	"context"
	"github.com/ansh-devs/ecomm-poc/order-service/dto"
	"github.com/ansh-devs/ecomm-poc/order-service/mocks/github.com/ansh-devs/ecomm-poc/order-service/repo"
	"github.com/go-kit/log"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
	"github.com/uber/jaeger-client-go/config"
	"testing"
)

func TestGetAllUserOrders(t *testing.T) {
	cfg := &config.Configuration{Disabled: true}
	tr, closer, _ := cfg.NewTracer()
	defer closer.Close()
	mockedRepo := repo.NewMockRequesterVariadic(t)
	userId := uuid.NewString()
	order1Id := uuid.NewString()
	order2Id := uuid.NewString()
	orderResponse := []dto.Order{{ID: order1Id, UserID: userId}, {ID: order2Id, UserID: userId}}
	ctx := context.Background()
	spn := tr.StartSpan("test_span")
	newCtx := opentracing.ContextWithSpan(ctx, spn)
	logger := log.NewNopLogger()
	mockedRepo.On("GetUserAllOrders", newCtx, userId, spn).
		Return(orderResponse, nil)
	svc := NewOrderService(mockedRepo, logger, tr)
	resp, _ := svc.GetAllUserOrders(ctx, userId)
	assert.Equal(t, "successful", resp.Status)
	assert.Equal(t, order1Id, resp.Orders[0].ID)
	assert.Equal(t, order2Id, resp.Orders[1].ID)

}

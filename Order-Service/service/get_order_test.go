package service

import (
	"context"
	"testing"

	"github.com/ansh-devs/commercelens/order-service/dto"
	"github.com/ansh-devs/commercelens/order-service/mocks/github.com/ansh-devs/commercelens/order-service/repo"
	"github.com/go-kit/log"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
	"github.com/uber/jaeger-client-go/config"
)

const GetOrderMethod = "GetOrder"

func TestGetOrder(t *testing.T) {
	cfg := &config.Configuration{Disabled: true}
	tr, closer, _ := cfg.NewTracer()
	defer closer.Close()
	mockedRepo := repo.NewMockRequesterVariadic(t)
	id := uuid.NewString()
	ctx := context.Background()
	spn := tr.StartSpan("test_span")
	newCtx := opentracing.ContextWithSpan(ctx, spn)
	logger := log.NewNopLogger()
	mockedRepo.On(GetOrderMethod, newCtx, id, spn).Return(dto.Order{ID: id}, nil)
	svc := NewOrderService(mockedRepo, logger, tr)
	resp, err := svc.GetOrder(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, id, resp.Order.ID)
	assert.Equal(t, "successful", resp.Status)
}

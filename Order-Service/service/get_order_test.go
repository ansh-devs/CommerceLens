package service

import (
	"context"
	"testing"

	"github.com/ansh-devs/ecomm-poc/order-service/dto"
	"github.com/ansh-devs/ecomm-poc/order-service/mocks/github.com/ansh-devs/ecomm-poc/order-service/repo"
	"github.com/go-kit/log"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
	"github.com/uber/jaeger-client-go/config"
)

func TestGetOrder(t *testing.T) {
	tr := opentracing.GlobalTracer()
	cfg := &config.Configuration{
		Disabled: true,
		// Sampler:  &config.SamplerConfig{Type: "0"},
	}
	tr, closer, _ := cfg.NewTracer()
	defer closer.Close()
	mockedRepo := repo.NewMockRequesterVariadic(t)
	id := uuid.NewString()
	ctx := context.Background()
	spn := tr.StartSpan("test_span")
	newCtx := opentracing.ContextWithSpan(ctx, spn)
	logger := log.NewNopLogger()
	mockedRepo.On("GetOrder", newCtx, id, spn).Return(
		dto.Order{
			ID:        id,
			ProductID: uuid.NewString(),
			UserID:    uuid.NewString(),
			Username:  "username",
		}, nil)
	svc := NewOrderService(mockedRepo, logger, tr)
	resp, err := svc.GetOrder(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, id, resp.Order.ID)
	assert.Equal(t, "successful", resp.Status)

}

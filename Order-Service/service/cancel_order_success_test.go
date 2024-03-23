package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/ansh-devs/commercelens/order-service/mocks/github.com/ansh-devs/commercelens/order-service/repo"
	"github.com/go-kit/log"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
	"github.com/uber/jaeger-client-go/config"
)

const CancelOrderMethod = "CancelOrder"

func TestCancelOrderSuccess(t *testing.T) {
	cfg := &config.Configuration{Disabled: true}
	tr, closer, _ := cfg.NewTracer()
	defer closer.Close()
	mockedRepo := repo.NewMockRequesterVariadic(t)
	id := uuid.NewString()
	ctx := context.Background()
	spn := tr.StartSpan("test_span")
	newCtx := opentracing.ContextWithSpan(ctx, spn)
	logger := log.NewNopLogger()
	mockedRepo.On(CancelOrderMethod, newCtx, id, spn).Return(fmt.Sprint("failed"), nil)
	svc2 := NewOrderService(mockedRepo, logger, tr)
	resp2, _ := svc2.CancelOrder(ctx, id)
	assert.Equal(t, "failed", resp2.Status)

}

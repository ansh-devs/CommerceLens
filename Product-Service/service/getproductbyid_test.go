package service

import (
	"context"
	db "github.com/ansh-devs/ecomm-poc/product-service/db/generated"
	"github.com/ansh-devs/ecomm-poc/product-service/mocks/github.com/ansh-devs/ecomm-poc/product-service/repo"
	"github.com/go-kit/log"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
	"github.com/uber/jaeger-client-go/config"
	"testing"
)

const GetProductByIdName = "GetProductById"

func TestProductService_GetProductById(t *testing.T) {
	cfg := &config.Configuration{Disabled: true}
	tr, closer, _ := cfg.NewTracer()
	defer closer.Close()
	mockedRepo := repo.NewMockRequesterVariadic(t)
	id := uuid.New().String()
	ctx := context.Background()
	spn := tr.StartSpan("test_span")
	newCtx := opentracing.ContextWithSpan(ctx, spn)
	//mockedRepo.On(GetProductByIdName, newCtx, id, spn).Return(db.Product{ID: id}, nil).Times(1)
	mockedRepo.EXPECT().GetProductByID(newCtx, id, spn).Return(db.Product{ID: id}, nil).Times(1)
	logger := log.NewNopLogger()
	svc := NewService(mockedRepo, logger, tr)
	resp, err := svc.GetProductById(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, id, resp.ID)
}

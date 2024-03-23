package service

import (
	"context"
	db "github.com/ansh-devs/commercelens/product-service/db/generated"
	"github.com/ansh-devs/commercelens/product-service/mocks/github.com/ansh-devs/commercelens/product-service/repo"
	"github.com/go-kit/log"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
	"github.com/uber/jaeger-client-go/config"
	"testing"
)

func TestProductService_GetAllProducts(t *testing.T) {
	cfg := &config.Configuration{Disabled: true}
	tr, closer, _ := cfg.NewTracer()
	var products []db.Product
	products = append(products, db.Product{
		ID:          uuid.New().String(),
		ProductName: "Product 1",
		Description: "product one description",
		Price:       "$473.81",
		CreatedAt:   nil,
	})
	products = append(products, db.Product{
		ID:          uuid.New().String(),
		ProductName: "Product 2",
		Description: "product two description",
		Price:       "$253.62",
		CreatedAt:   nil,
	})
	defer closer.Close()
	mockedRepo := repo.NewMockRequesterVariadic(t)
	ctx := context.Background()
	spn := tr.StartSpan("test_span")
	newCtx := opentracing.ContextWithSpan(ctx, spn)
	mockedRepo.EXPECT().GetAllProducts(newCtx, spn).Return(products, nil).Times(1)
	logger := log.NewNopLogger()
	svc := NewService(mockedRepo, logger, tr)
	resp, err := svc.GetAllProducts(newCtx)
	assert.NoError(t, err)
	assert.Equal(t, len(products), len(resp))
}

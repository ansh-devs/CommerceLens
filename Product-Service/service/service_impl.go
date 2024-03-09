package service

import (
	"context"
	"fmt"
	db "github.com/ansh-devs/ecomm-poc/product-service/db/generated"
	"github.com/ansh-devs/ecomm-poc/product-service/dto"
	"github.com/ansh-devs/ecomm-poc/product-service/natsutil"
	"github.com/ansh-devs/ecomm-poc/product-service/repo"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/hashicorp/consul/api"
	"github.com/opentracing/opentracing-go"
	"math/rand"
	"strconv"
)

type ProductService struct {
	repository   repo.Repository
	logger       log.Logger
	ConsulClient *api.Client
	trace        opentracing.Tracer
	SrvID        string
	nats         *natsutil.NATSComponent
}

func NewService(rep repo.Repository, logger log.Logger, tracer opentracing.Tracer) *ProductService {
	client, err := api.NewClient(&api.Config{
		Address: "service-discovery:8500",
	})
	if err != nil {
		fmt.Println(err)
	}
	srvID := "instance_" + strconv.Itoa(rand.Intn(99))
	nc := natsutil.NewNatsComponent(srvID)
	err = nc.ConnectToNATS("nats://nats-srvr:4222", nil)
	if err != nil {
		fmt.Println(err)
	}
	return &ProductService{
		repository:   rep,
		logger:       log.With(logger, "layer", "service"),
		ConsulClient: client,
		trace:        tracer,
		SrvID:        srvID,
		nats:         nc,
	}
}

func (s *ProductService) PurchaseProduct(ctx context.Context, userId, productId string) error {
	_ = level.Info(s.logger).Log("method-invoked", "PurchaseOrder")
	spn := s.trace.StartSpan("purchase-product-product-srv")
	defer spn.Finish()
	resp, err := s.repository.GetProductByID(ctx, productId, spn)
	order := dto.NatsPurchaseOrder{
		UserId: userId,
		Product: dto.Product{
			ID:          resp.ID,
			ProductName: resp.ProductName,
			Description: resp.Description,
			Price:       resp.Price,
		},
	}
	if err != nil {
		return err
	}
	err = s.nats.Publish("products.purchase", order)
	if err != nil {
		return err
	} else {
		return nil
	}
}
func (s *ProductService) GetProductById(ctx context.Context, productId string) (db.Product, error) {
	_ = level.Info(s.logger).Log("method-invoked", "GetProductById")
	spn := s.trace.StartSpan("get_product_by_id_service_layer_call")
	defer spn.Finish()
	resp, err := s.repository.GetProductByID(ctx, productId, spn)
	if err != nil {
		return db.Product{}, err
	}
	return resp, nil
}

func (s *ProductService) GetAllProducts(ctx context.Context) ([]db.Product, error) {
	_ = level.Info(s.logger).Log("method-invoked", "GetAllProducts")
	spn := s.trace.StartSpan("get_all_products_service_layer_call")
	defer spn.Finish()
	resp, err := s.repository.GetAllProducts(ctx, spn)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

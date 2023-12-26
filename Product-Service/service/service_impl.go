package service

import (
	"fmt"
	"github.com/ansh-devs/microservices_project/product-service/repo"
	"github.com/go-kit/log"
	"github.com/hashicorp/consul/api"
)

type ProductService struct {
	repository   repo.Repository
	logger       log.Logger
	consulClient *api.Client
}

func NewService(rep repo.Repository, logger log.Logger) *ProductService {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Println(err)
	}
	return &ProductService{
		repository:   rep,
		logger:       log.With(logger, "layer", "service"),
		consulClient: client,
	}
}

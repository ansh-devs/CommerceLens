package repo

import (
	db "github.com/ansh-devs/microservices_project/product-service/db/generated"
	"github.com/go-kit/log"
)

type Repo struct {
	db     *db.Queries
	logger log.Logger
}

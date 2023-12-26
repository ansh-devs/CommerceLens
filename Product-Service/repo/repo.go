package repo

import "context"

type Repository interface {
	PlaceOrder(context.Context)
}

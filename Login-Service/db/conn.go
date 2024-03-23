package db

import (
	"context"
	"fmt"

	db "github.com/ansh-devs/commercelens/login-service/db/generated"
	"github.com/jackc/pgx/v5"
)

// MustConnectToPostgress urlExample := "postgres://username:password@localhost:5432/database_name"
func MustConnectToPostgress(uri string) db.Querier {
	ctx := context.Background()
	pgconn, err := pgx.Connect(ctx, uri)

	if err != nil {
		panic(err)
	}
	q := db.New(pgconn)
	if pgconn == nil {
		defer func(pgconn *pgx.Conn, ctx context.Context) {
			err := pgconn.Close(ctx)
			if err != nil {
				fmt.Println(err)
			}
		}(pgconn, ctx)
	}
	return q
}

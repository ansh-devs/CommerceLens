package db

import (
	"context"
	"fmt"
	db "github.com/ansh-devs/commercelens/product-service/db/generated"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

// MustConnectToPostgress urlExample := "postgres://username:password@localhost:5432/database_name"

func dbConfWithPool(url string) *pgxpool.Config {
	const defaultMaxConns = int32(3)
	const defaultMinConns = int32(2)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 10
	config, err := pgxpool.ParseConfig(url)

	config.MaxConns = defaultMaxConns
	config.MinConns = defaultMinConns
	config.MaxConnLifetime = defaultMaxConnLifetime
	config.MaxConnIdleTime = defaultMaxConnIdleTime
	config.HealthCheckPeriod = defaultHealthCheckPeriod
	config.ConnConfig.ConnectTimeout = defaultConnectTimeout

	if err != nil {
		return nil
	}
	return config
}

func MustConnectToPostgress(uri string) *db.Queries {
	ctx := context.Background()
	poolConf := dbConfWithPool(uri)
	config, err := pgxpool.NewWithConfig(ctx, poolConf)
	if err != nil {
		panic(fmt.Sprintf("can't establish conection pooling %v", err.Error()))
	}
	q := db.New(config)
	return q
}

package order

import (
	"context"
	"flag"
	"fmt"
	"github.com/ansh-devs/microservices_project/order-service/db"
	"github.com/ansh-devs/microservices_project/order-service/order"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const dbSource = "pgsql_url"

func main() {
	var httpAddr = flag.String("http", ":8080", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "order_service",
			"time", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	_ = level.Info(logger).Log("service started")

	flag.Parse()
	ctx := context.Background()
	var srv *order.OrderService
	{
		dbConn := db.MustConnectToPostgress(dbSource)
		repository := order.NewRepo(dbConn, logger)
		srv = order.NewService(repository, logger)

	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := order.NewEndpoints(srv)

	go func() {
		fmt.Println("listening on port :", *httpAddr)
		handler := order.NewHttpServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()
	_ = level.Error(logger).Log("exit", <-errs)
}

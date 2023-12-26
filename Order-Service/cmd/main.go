package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ansh-devs/microservices_project/order-service/db"
	"github.com/ansh-devs/microservices_project/order-service/endpoints"
	"github.com/ansh-devs/microservices_project/order-service/repo"
	"github.com/ansh-devs/microservices_project/order-service/service"
	"github.com/ansh-devs/microservices_project/order-service/transport"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	_ "github.com/lib/pq"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const dbSource = "pgsql_url"

func main() {
	var httpAddr = flag.String("http", ":8080", "http listen address")

	cfg := &config.Configuration{
		ServiceName: "order-service",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
	}

	errs := make(chan error)

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		errs <- fmt.Errorf("%s", err)
	}
	defer func(closer io.Closer) {
		err := closer.Close()
		if err != nil {
			errs <- fmt.Errorf("%s", err)
		}
	}(closer)

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
	var srv *service.OrderService
	{
		dbConn := db.MustConnectToPostgress(dbSource)
		repository := repo.NewRepo(dbConn, logger)
		srv = service.NewService(repository, logger, tracer)

	}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoint := endpoints.NewEndpoints(srv)

	go func() {
		fmt.Println("listening on port :", *httpAddr)
		handler := transport.NewHttpServer(ctx, endpoint)

		errs <- http.ListenAndServe(*httpAddr, handler)
	}()
	_ = level.Error(logger).Log("exit", <-errs)
}

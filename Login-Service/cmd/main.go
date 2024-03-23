package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	config2 "github.com/ansh-devs/commercelens/login-service/config"
	"github.com/ansh-devs/commercelens/login-service/db"
	"github.com/ansh-devs/commercelens/login-service/endpoints"
	"github.com/ansh-devs/commercelens/login-service/repo"
	"github.com/ansh-devs/commercelens/login-service/service"
	"github.com/ansh-devs/commercelens/login-service/transport"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	_ "github.com/lib/pq"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func main() {
	config2.InitEnvConfigs()
	var httpAddr = &config2.AppConfigs.HttpAddr
	//var httpAddr = flag.String("http", ":8081", "http listen address")
	tracer := opentracing.GlobalTracer()
	cfg := &config.Configuration{
		ServiceName: "login-service",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
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
			"service", "login_service",
			"time", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	_ = level.Info(logger).Log("msg", "service started")

	flag.Parse()
	ctx := context.Background()
	var srv *service.LoginService
	{

		var dbSource = fmt.Sprintf("postgres://%s:%s@%s/%s",
			config2.AppConfigs.DatabaseUser,
			config2.AppConfigs.DatabasePassword,
			config2.AppConfigs.DatabaseHost,
			config2.AppConfigs.DatabaseName,
		)

		dbConn := db.MustConnectToPostgress(dbSource)
		repository := repo.NewPostgresRepo(dbConn, logger, tracer)
		srv = service.NewService(repository, logger, tracer)
	}
	go srv.GetUserWithNats(context.Background())
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoint := endpoints.NewEndpoints(srv)
	srv.RegisterService(httpAddr)
	go srv.UpdateHealthStatus()

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := transport.NewHttpServer(ctx, endpoint)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	for sig := range errs {
		_ = level.Error(logger).Log("status", sig, "GRACEFULLY SHUTTING DOWN !")
		_ = srv.ConsulClient.Agent().ServiceDeregister(srv.SrvID)
		os.Exit(0)
	}

}

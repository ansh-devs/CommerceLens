package main

import (
	"github.com/ansh-devs/microservices_project/order-service/db"
	"github.com/ansh-devs/microservices_project/order-service/endpoints"
	repo2 "github.com/ansh-devs/microservices_project/order-service/repo"
	"github.com/ansh-devs/microservices_project/order-service/service"
	"github.com/ansh-devs/microservices_project/order-service/transport"
	"github.com/opentracing/opentracing-go"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHttp(t *testing.T) {
	tracer := opentracing.GlobalTracer()
	dbConf := db.MustConnectToPostgress("")
	repo := repo2.NewRepo(dbConf, nil, tracer)
	svc := service.NewService(repo, nil, tracer)
	eps := endpoints.NewEndpoints(svc)
	mux := transport.NewHttpServer(eps)
	srv := httptest.NewServer(mux)
	defer srv.Close()

	for _, testcase := range []struct {
		method, url, body, want string
	}{
		{"GET", srv.URL + "/concat", `{"a":"1","b":"2"}`, `{"v":"12"}`},
		{"GET", srv.URL + "/sum", `{"a":1,"b":2}`, `{"v":3}`},
	} {
		req, _ := http.NewRequest(testcase.method, testcase.url, strings.NewReader(testcase.body))
		resp, _ := http.DefaultClient.Do(req)
		body, _ := io.ReadAll(resp.Body)
		if want, have := testcase.want, strings.TrimSpace(string(body)); want != have {
			t.Errorf("%s %s %s: want %q, have %q", testcase.method, testcase.url, testcase.body, want, have)
		}
	}

}

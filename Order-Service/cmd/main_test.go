package main

// type ServerTestSuite struct {
// 	suite.Suite
// }

// func TestHttp(t *testing.T) {
// 	tracer := opentracing.GlobalTracer()
// 	dbConf := db.MustConnectToPostgress("")
// 	repo := repo2.NewPostgresRepository(dbConf, nil, tracer)
// 	svc := service.NewOrderService(repo, nil, tracer)
// 	eps := endpoints.NewEndpoints(svc)
// 	mux := transport.NewHttpServer(eps)
// 	srv := httptest.NewServer(mux)
// 	defer srv.Close()

// 	for _, testcase := range []struct {
// 		method, url, body, want string
// 	}{
// 		{"GET", srv.URL + "/orders/v1/get-order/id", "", `{"v":"12"}`},
// 		{"GET", srv.URL + "/orders/v1/", `{"a":1,"b":2}`, `{"v":3}`},
// 	} {
// 		req, _ := http.NewRequest(testcase.method, testcase.url, strings.NewReader(testcase.body))
// 		resp, _ := http.DefaultClient.Do(req)
// 		body, _ := io.ReadAll(resp.Body)
// 		if want, have := testcase.want, strings.TrimSpace(string(body)); want != have {
// 			t.Errorf("%s %s %s: want %q, have %q", testcase.method, testcase.url, testcase.body, want, have)
// 		}
// 	}

// }

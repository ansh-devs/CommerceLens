package transport

import (
	"context"
	"encoding/json"
	"fmt"
	transport "github.com/go-kit/kit/transport/http"
	"net/http"
)

var (
	errorDecorator = []transport.ServerOption{
		transport.ServerErrorEncoder(func(ctx context.Context, e error, w http.ResponseWriter) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			errorModel := struct {
				Err     string `json:"err"`
				Code    int    `json:"code"`
				Message string `json:"message"`
			}{
				Err:     e.Error(),
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			}
			marshal, err := json.Marshal(errorModel)
			if err != nil {
				return
			}
			_, err = w.Write(marshal)
			if err != nil {
				fmt.Println(err)
			}
		}),
	}

	notFoundHandler = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		errorModel := struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}
		marshal, err := json.Marshal(errorModel)
		if err != nil {
			return
		}
		_, err = writer.Write(marshal)
		if err != nil {
			return
		}
	})
)

func JsonTypeReWrittermiddleware(nextRoute http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		nextRoute.ServeHTTP(writer, request)
	})
}

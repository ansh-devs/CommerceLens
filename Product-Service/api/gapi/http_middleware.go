package gapi

import (
	"log"
	"net/http"
	"time"
)

func HttpMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		startTime := time.Now()
		handler.ServeHTTP(writer, request)
		timeTaken := time.Since(startTime)
		log.Printf("[HTTP]: Method : %s | Path : %s |  Duration : %s\n", request.Method, request.RequestURI, timeTaken)
	})
}

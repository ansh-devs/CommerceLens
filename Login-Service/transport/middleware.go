package transport

import "net/http"

func JsonTypeReWrittermiddleware(nextRoute http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		nextRoute.ServeHTTP(writer, request)
	})
}

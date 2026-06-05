package core_http_middleware

import (
	"fmt"
	"net/http"

	core_logger "github.com/Mommsent/todoapp-Studying.git/internal/core/logger"
)

func Dummy(s string) Middleware {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(rwriter http.ResponseWriter, request *http.Request) {
			ctx := request.Context()
			log := core_logger.FromContext(ctx)

			log.Debug(fmt.Sprintf("---> before: %s", s))
			handler.ServeHTTP(rwriter, request)
			log.Debug(fmt.Sprintf("<--- after: %s", s))
		})
	}
}

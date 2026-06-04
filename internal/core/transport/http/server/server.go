package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	core_logger "github.com/Mommsent/todoapp-Studying.git/internal/core/logger"
	core_http_middleware "github.com/Mommsent/todoapp-Studying.git/internal/core/transport/http/middleware"
	"go.uber.org/zap"
)

type HTTPServer struct {
	mux        *http.ServeMux
	config     Config
	log        *core_logger.Logger
	middleware []core_http_middleware.Middleware
}

func NewHTTPServer(
	conf Config,
	logger *core_logger.Logger,
	middleware ...core_http_middleware.Middleware) *HTTPServer {
	return &HTTPServer{
		mux:        http.NewServeMux(),
		config:     conf,
		log:        logger,
		middleware: middleware,
	}
}

func (handler *HTTPServer) RegisterAPIRouters(routers ...*APIVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.apiVersion)

		handler.mux.Handle(
			prefix+"/",
			http.StripPrefix(prefix, router),
		)
	}
}

func (httpServer *HTTPServer) Run(ctx context.Context) error {

	mux := core_http_middleware.ChainMiddleware(httpServer.mux, httpServer.middleware...)

	server := &http.Server{
		Addr:    httpServer.config.Addr,
		Handler: mux,
	}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)

		httpServer.log.Warn("Start HTTP server", zap.String("adrr", httpServer.config.Addr))

		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("listen and server HTTP: %w", err)
		}
	case <-ctx.Done():
		httpServer.log.Warn("Shutting down http server...")

		shutdonwCtx, canel := context.WithTimeout(
			context.Background(),
			httpServer.config.ShotdownTimeout,
		)

		defer canel()

		if err := server.Shutdown(shutdonwCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("shutting down HTTP server: %w", err)
		}

		httpServer.log.Warn("Http server stopped")
	}

	return nil
}

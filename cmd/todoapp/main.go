package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/Mommsent/todoapp-Studying.git/internal/core/logger"
	core_posgres_pool "github.com/Mommsent/todoapp-Studying.git/internal/core/repository/postgres/pool"
	core_http_middleware "github.com/Mommsent/todoapp-Studying.git/internal/core/transport/http/middleware"
	core_http_server "github.com/Mommsent/todoapp-Studying.git/internal/core/transport/http/server"
	users_postgres_repository "github.com/Mommsent/todoapp-Studying.git/internal/features/users/repository/postgres"
	users_service "github.com/Mommsent/todoapp-Studying.git/internal/features/users/service"
	users_transport_http "github.com/Mommsent/todoapp-Studying.git/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	loggerConfig := core_logger.NewConfigMust()
	logger, err := core_logger.NewLogger(loggerConfig)
	if err != nil {
		fmt.Println("failed to init application logger: ", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("Initializing postgres connection pool")
	pool, err := core_posgres_pool.NewConnectionPool(ctx, core_posgres_pool.NewConfigMust())
	if err != nil {
		logger.Fatal("Failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("initializing features", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransportHTTP := users_transport_http.NewUserHTTPHandler(usersService)

	logger.Debug("Initialize HTTP server")
	servConfig := core_http_server.NewConfigMust()
	httpServer := core_http_server.NewHTTPServer(
		servConfig,
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)
	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.APIVersion1)
	apiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes()...)
	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	core_config "github.com/Mommsent/todoapp-Studying.git/internal/core/config"
	core_logger "github.com/Mommsent/todoapp-Studying.git/internal/core/logger"
	core_pgx_pool "github.com/Mommsent/todoapp-Studying.git/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/Mommsent/todoapp-Studying.git/internal/core/transport/http/middleware"
	core_http_server "github.com/Mommsent/todoapp-Studying.git/internal/core/transport/http/server"
	statistics_postgres_repository "github.com/Mommsent/todoapp-Studying.git/internal/features/statistics/repository/postgres"
	statistics_service "github.com/Mommsent/todoapp-Studying.git/internal/features/statistics/service"
	statistics_transport_http "github.com/Mommsent/todoapp-Studying.git/internal/features/statistics/transport/http"
	tasks_postgres_repository "github.com/Mommsent/todoapp-Studying.git/internal/features/tasks/repository/postgres"
	tasks_service "github.com/Mommsent/todoapp-Studying.git/internal/features/tasks/service"
	tasks_transport_http "github.com/Mommsent/todoapp-Studying.git/internal/features/tasks/transport/http"
	users_postgres_repository "github.com/Mommsent/todoapp-Studying.git/internal/features/users/repository/postgres"
	users_service "github.com/Mommsent/todoapp-Studying.git/internal/features/users/service"
	users_transport_http "github.com/Mommsent/todoapp-Studying.git/internal/features/users/transport/http"
	"go.uber.org/zap"

	_ "github.com/Mommsent/todoapp-Studying.git/docs"
)

// @title           Todo App API
// @version         1.0
// @description     API для управления задачами
// @host            127.0.0.1:5050
// @BasePath        /api/v1
func main() {

	cfg := core_config.NewConfigMust()
	time.Local = cfg.TimeZone

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

	logger.Debug("application time zone", zap.Any("zone", time.Local))
	logger.Debug("Initializing postgres connection pool")

	pool, err := core_pgx_pool.NewPool(
		ctx,
		core_pgx_pool.NewConfigMust(),
	)
	if err != nil {
		logger.Fatal("Failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("initializing features", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransportHTTP := users_transport_http.NewUserHTTPHandler(usersService)

	logger.Debug("initializing features", zap.String("feature", "tasks"))
	tasksRepository := tasks_postgres_repository.NewTasksRepository(pool)
	tasksService := tasks_service.NewTasksService(tasksRepository)
	tasksTransportHTTP := tasks_transport_http.NewTasksHTTPHandler(tasksService)

	logger.Debug("initializing feature", zap.String("feature", "statistics"))
	statisticsRepository := statistics_postgres_repository.NewStatisticsRepository(pool)
	statisticsService := statistics_service.NewStatisticsService(statisticsRepository)
	statisticsTransportHTTP := statistics_transport_http.NewStatisticsHTTPHandler(statisticsService)

	logger.Debug("Initialize HTTP server")
	servConfig := core_http_server.NewConfigMust()
	httpServer := core_http_server.NewHTTPServer(
		servConfig,
		logger,
		core_http_middleware.CORS(),
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)
	apiVersionRouterV1 := core_http_server.NewAPIVersionRouter(core_http_server.APIVersion1)
	apiVersionRouterV1.RegisterRoutes(usersTransportHTTP.Routes()...)
	apiVersionRouterV1.RegisterRoutes(tasksTransportHTTP.Routes()...)
	apiVersionRouterV1.RegisterRoutes(statisticsTransportHTTP.Routes()...)

	apiVersionRouterV2 := core_http_server.NewAPIVersionRouter(
		core_http_server.APIVersion2,
		core_http_middleware.Dummy("api v2 middleware"),
	)
	apiVersionRouterV2.RegisterRoutes(usersTransportHTTP.Routes()...)
	apiVersionRouterV2.RegisterRoutes(tasksTransportHTTP.Routes()...)
	apiVersionRouterV2.RegisterRoutes(statisticsTransportHTTP.Routes()...)

	httpServer.RegisterAPIRouters(
		apiVersionRouterV1,
		apiVersionRouterV2,
	)

	httpServer.RegisterSwagger()

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}

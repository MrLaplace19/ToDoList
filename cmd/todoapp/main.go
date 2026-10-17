package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/MrLaplace19/ToDoList/internal/core/logger"
	core_pgx_pool "github.com/MrLaplace19/ToDoList/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/MrLaplace19/ToDoList/internal/core/transport/http/middelware"
	core_http_server "github.com/MrLaplace19/ToDoList/internal/core/transport/http/server"
	statistics_postgres_repository "github.com/MrLaplace19/ToDoList/internal/features/statistics/repository/postgres"
	statistics_service "github.com/MrLaplace19/ToDoList/internal/features/statistics/service"
	statistics_transport_http "github.com/MrLaplace19/ToDoList/internal/features/statistics/transport/http"
	tasks_postgres_repository "github.com/MrLaplace19/ToDoList/internal/features/tasks/repository/postgres"
	tasks_service "github.com/MrLaplace19/ToDoList/internal/features/tasks/service"
	tasks_transport_http "github.com/MrLaplace19/ToDoList/internal/features/tasks/transport/http"
	users_postgres_repository "github.com/MrLaplace19/ToDoList/internal/features/users/repository/postgres"
	users_service "github.com/MrLaplace19/ToDoList/internal/features/users/service"
	users_transport_http "github.com/MrLaplace19/ToDoList/internal/features/users/transport/http"
	"go.uber.org/zap"
	_ "github.com/MrLaplace19/ToDoList/docs"
)


// @title golang-todoapp
// @version 1.0
// @description restAPI-schema
// @BasePath /api/v1
func main() {

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)

	defer cancel()

	fmt.Println("HELLO TODOAPP")

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())

	if err != nil {
		fmt.Println("FAIL LOGGER", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("initializing to init connection pool")
	pool, err := core_pgx_pool.NewPgxConnectionPool(core_pgx_pool.NewConfigMust(), ctx)
	if err != nil {
		logger.Fatal("failed to init connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("initializing feature", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersServices := users_service.NewUserService(usersRepository)
	usersTransportHTTP := users_transport_http.NewUserHttpHandler(usersServices)

	logger.Debug("initialized feature", zap.String("feature", "tasks"))
	taskRepository := tasks_postgres_repository.NewTaskRepository(pool)
	taskServices := tasks_service.NewTasksService(taskRepository)
	taskTransportHTTP := tasks_transport_http.NewTasksHTTPHandler(taskServices)

	logger.Debug("initialized feature", zap.String("feature", "statistics"))
	statisticsRepository := statistics_postgres_repository.NewStatisticsRepository(pool)
	statisticsService := statistics_service.NewStatisticsService(statisticsRepository)
	statisticsTransportHTTP := statistics_transport_http.NewStatisticsHTTPHandler(statisticsService)

	logger.Debug("initialized HTTP server")

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.CORS(),
		core_http_middleware.RequestId(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)
	apiVersionRouterV1 := core_http_server.NewAPIVersion(core_http_server.ApiVersion1)
	apiVersionRouterV1.RegisterRoutes(usersTransportHTTP.Routes()...)
	apiVersionRouterV1.RegisterRoutes(taskTransportHTTP.Routes()...)
	apiVersionRouterV1.RegisterRoutes(statisticsTransportHTTP.Routes()...)

	httpServer.RegisterAPIRouters(
		apiVersionRouterV1,
	)
	httpServer.RegisterSwaggerRoute()
	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}

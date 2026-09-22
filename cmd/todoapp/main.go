package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/MrLaplace19/ToDoList/internal/core/logger"
	core_postgres_pool "github.com/MrLaplace19/ToDoList/internal/core/repository/postgres/pool"
	core_http_middleware "github.com/MrLaplace19/ToDoList/internal/core/transport/http/middelware"
	core_http_server "github.com/MrLaplace19/ToDoList/internal/core/transport/http/server"
	users_postgres_repository "github.com/MrLaplace19/ToDoList/internal/features/users/repository/postgres"
	users_service "github.com/MrLaplace19/ToDoList/internal/features/users/service"
	users_transport_http "github.com/MrLaplace19/ToDoList/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)

	defer cancel()

	fmt.Println("HELLO TODOAPP")

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	
	if err != nil{
		fmt.Println("FAIL LOGGER", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("initializing to init connection pool")
	pool, err := core_postgres_pool.NewConnectionPool(core_postgres_pool.NewConfigMust(), ctx)
	if err != nil{
		logger.Fatal("failed to init connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("initializing feature", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersServices := users_service.NewUserService(usersRepository)
	usersTransportHTTP := users_transport_http.NewUserHttpHandler(usersServices)

	logger.Debug("initialized HTTP server")

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestId(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)
	apiVersionRouter := core_http_server.NewAPIVersion(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes()...)
	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err !=nil{
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
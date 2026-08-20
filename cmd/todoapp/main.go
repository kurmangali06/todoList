package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	core_logger "todolist/internal/core/logger"
	"todolist/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "todolist/internal/core/transport/http/middleware"
	corre_http_server "todolist/internal/core/transport/http/server"
	tasks_postgres_repository "todolist/internal/features/tasks/repository/postgres"
	tasks_service "todolist/internal/features/tasks/service"
	tasks_transport_http "todolist/internal/features/tasks/transport/http"
	users_postgres_repository "todolist/internal/features/users/repository/postgres"
	users_service "todolist/internal/features/users/service"
	users_transport_http "todolist/internal/features/users/transport/http"

	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())

	if err != nil {
		fmt.Println("Logger init error:", err)
		os.Exit(1)
	}
	defer logger.Close()
	logger.Debug("init postgres connection pool")
	postgresPool, err := core_pgx_pool.NewPool(
		ctx,
		core_pgx_pool.NewConfigMust(),
	)
	if err != nil {
		logger.Fatal("Postgres connection pool init error", zap.Error(err))
	}

	defer postgresPool.Close()

	logger.Debug("initializing feature", zap.String("feature", "users"))
	userRepository := users_postgres_repository.NewUsersRepository(postgresPool)
	userService := users_service.NewUsersService(userRepository)
	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(userService)

	logger.Debug("initializing feature", zap.String("feature", "tasks"))
	tasksRepository := tasks_postgres_repository.NewTasksRepository(postgresPool)
	tasksService := tasks_service.NewTasksService(tasksRepository)
	tasksTransportHTTP := tasks_transport_http.NewTasksHTTPHandler(tasksService)

	logger.Debug("init users HTTP server")

	usersRoutes := usersTransportHTTP.Routes()
	tasksRoutes := tasksTransportHTTP.Routes()

	apiVersionRouter := corre_http_server.NewAPIVersionRouter(corre_http_server.ApiVersionV1)
	apiVersionRouter.RegisterRoutes(usersRoutes...)
	apiVersionRouter.RegisterRoutes(tasksRoutes...)

	httpServer := corre_http_server.NewHTTPServer(
		corre_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestId(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	httpServer.RegisterAPIRoutes(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}

}

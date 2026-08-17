package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	core_logger "todolist/internal/core/logger"
	core_http_middleware "todolist/internal/core/transport/http/middleware"
	corre_http_server "todolist/internal/core/transport/http/server"
	users_transport_http "todolist/internal/features/users/transport/http"

	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	fmt.Println("Hello, World!")
	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())

	if err != nil {
		fmt.Println("Logger init error:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("Starting TodoApp")

	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(nil)
	usersRoutes := usersTransportHTTP.Routes()

	apiVersionRouter := corre_http_server.NewAPIVersionRouter(corre_http_server.ApiVersionV1)
	apiVersionRouter.RegisterRoutes(usersRoutes...)

	httpServer := corre_http_server.NewHTTPServer(
		corre_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestId(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)

	httpServer.RegisterAPIRoutes(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}

}

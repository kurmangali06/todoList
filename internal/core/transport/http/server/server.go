package corre_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	core_logger "todolist/internal/core/logger"
	core_http_middleware "todolist/internal/core/transport/http/middleware"

	"go.uber.org/zap"
)

type HTTPServer struct {
	mux        *http.ServeMux
	config     Config
	log        *core_logger.Logger
	middleware []core_http_middleware.Middleware
}

func NewHTTPServer(
	config Config,
	log *core_logger.Logger,
	middleware ...core_http_middleware.Middleware,
) *HTTPServer {
	return &HTTPServer{
		mux:        http.NewServeMux(),
		config:     config,
		log:        log,
		middleware: middleware,
	}
}

func (h *HTTPServer) RegisterAPIRoutes(routes ...*APIVersionRouter) {
	for _, route := range routes {
		prefix := "/api/" + string(route.apiVersion)

		h.mux.Handle(
			prefix+"/",
			http.StripPrefix(prefix, route),
		)
	}
}

func (h *HTTPServer) Run(ctx context.Context) error {
	mux := core_http_middleware.ChainMiddlewares(h.mux, h.middleware...)
	server := &http.Server{
		Addr:    h.config.Addr,
		Handler: mux,
	}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)

		h.log.Warn("start HTTP server", zap.String("addr", h.config.Addr))

		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("listern and  run HTTP server: %w", err)
		}
	case <-ctx.Done():
		h.log.Warn("shutdown HTTP server")

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			h.config.ShutdownTimeout,
		)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("shutdown HTTP server: %w", err)
		}

		h.log.Info("HTTP server stopped")
	}

	return nil

}

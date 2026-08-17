package core_http_middleware

import (
	"context"
	"net/http"
	"time"
	core_logger "todolist/internal/core/logger"
	core_http_response "todolist/internal/core/transport/http/response"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

const requestIdHeader = "X-Request-Id"

func RequestId() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestId := r.Header.Get(requestIdHeader)
			if requestId == "" {
				requestId = uuid.NewString()
			}

			r.Header.Set(requestIdHeader, requestId)

			next.ServeHTTP(w, r)
		})
	}
}

func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestId := r.Header.Get(requestIdHeader)

			l := log.With(
				zap.String("requestId", requestId),
				zap.String("url", r.URL.String()))

			ctx := context.WithValue(r.Context(), "log", l)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			responseHadler := core_http_response.NewHTTPResponseHandler(log, w)
			defer func() {
				if p := recover(); p != nil {
					responseHadler.PanicResponse(
						p,
						"during handler HTTP request got unexpected error ",
					)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			rw := core_http_response.NewHTTPResponseWriter(w)

			before := time.Now()

			log.Debug(
				">>> incoming HTTP request",
				zap.Time("time", before.UTC()),
			)

			next.ServeHTTP(rw, r)

			log.Debug(
				"<<< outgoing HTTP response",
				zap.Int("status", rw.GetStatusCodeOrPanic()),
				zap.Duration("latency", time.Now().Sub(before)),
			)
		})
	}
}

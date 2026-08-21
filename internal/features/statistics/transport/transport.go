package statistics_transport_http

import (
	"context"
	"net/http"
	"time"
	"todolist/internal/core/domain"
	corre_http_server "todolist/internal/core/transport/http/server"

	"github.com/google/uuid"
)

type StatisticsHTTPHandler struct {
	statisticsService StatisticsService
}

type StatisticsService interface {
	GetStatistics(
		ctx context.Context,
		userID *uuid.UUID,
		from *time.Time,
		to *time.Time,
	) (domain.Statistics, error)
}

func NewStatisticsHTTPHandler(statisticsService StatisticsService) *StatisticsHTTPHandler {
	return &StatisticsHTTPHandler{statisticsService: statisticsService}
}

func (h *StatisticsHTTPHandler) Routes() []corre_http_server.Route {
	return []corre_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/statistics",
			Handler: h.GetStatistics,
		},
	}
}

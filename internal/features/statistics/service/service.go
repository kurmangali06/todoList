package statistics_service

import (
	"context"
	"time"
	"todolist/internal/core/domain"

	"github.com/google/uuid"
)

type StatisticsService struct {
	statisticsRepository StatisticsRepository
}

type StatisticsRepository interface {
	GetTasks(
		ctx context.Context,
		userID *uuid.UUID,
		from *time.Time,
		to *time.Time,
	) ([]domain.Task, error)
}

func NewStatisticsService(statisticsRepository StatisticsRepository) *StatisticsService {
	return &StatisticsService{statisticsRepository: statisticsRepository}
}

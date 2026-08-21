package statistics_service

import (
	"context"
	"fmt"
	"time"
	"todolist/internal/core/domain"
	core_error "todolist/internal/core/errors"

	"github.com/google/uuid"
)

func (s *StatisticsService) GetStatistics(
	ctx context.Context,
	userID *uuid.UUID,
	from *time.Time,
	to *time.Time,
) (domain.Statistics, error) {
	if from != nil && to != nil {
		if to.Before(*from) || to.Equal(*from) {
			return domain.Statistics{}, fmt.Errorf(
				"`to` must be after `from`: %w",
				core_error.ErrInvalidArgument,
			)
		}
	}

	tasks, err := s.statisticsRepository.GetTasks(ctx, userID, from, to)

	if err != nil {
		return domain.Statistics{}, fmt.Errorf("get statistics from repository: %w", err)
	}

	statistics := domain.CreateStatistics(tasks)
	return statistics, nil
}

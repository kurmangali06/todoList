package tasks_service

import (
	"context"
	"fmt"
	"todolist/internal/core/domain"
	core_error "todolist/internal/core/errors"

	"github.com/google/uuid"
)

func (s *TasksService) GetTasks(
	ctx context.Context,
	userID *uuid.UUID,
	limit *int,
	offset *int) ([]domain.Task, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf("limit must be positive : %w", core_error.ErrInvalidArgument)
	}
	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf("offset must be positive : %w", core_error.ErrInvalidArgument)
	}

	tasks, err := s.tasksRepository.GetTasks(ctx, userID, limit, offset)

	if err != nil {
		return nil, fmt.Errorf("get tasks from repository: %w", err)
	}
	return tasks, nil

}

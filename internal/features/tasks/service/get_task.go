package tasks_service

import (
	"context"
	"fmt"
	"todolist/internal/core/domain"

	"github.com/google/uuid"
)

func (s *TasksService) GetTask(
	ctx context.Context,
	userID uuid.UUID,
) (domain.Task, error) {
	task, err := s.tasksRepository.GetTask(ctx, userID)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get task from repository: %w", err)
	}
	return task, nil
}

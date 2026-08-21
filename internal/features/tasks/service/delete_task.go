package tasks_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *TasksService) DeleteTask(ctx context.Context, id uuid.UUID) error {
	if err := s.tasksRepository.DeleteTask(ctx, id); err != nil {
		return fmt.Errorf("delete task from repository: %w", err)
	}
	return nil
}

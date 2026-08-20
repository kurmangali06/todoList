package tasks_service

import (
	"context"
	"fmt"
	"todolist/internal/core/domain"

	"github.com/google/uuid"
)

func (s *TasksService) CreateTask(
	ctx context.Context,
	title string,
	description *string,
	authorUserID uuid.UUID,
) (domain.Task, error) {
	task := domain.CreateTask(title, description, authorUserID)

	if err := task.Validate(); err != nil {
		return domain.Task{}, fmt.Errorf("validate task domain: %w", err)
	}
	task, err := s.tasksRepository.CreateTask(ctx, task)

	if err != nil {
		return domain.Task{}, fmt.Errorf("create task in repository: %w", err)
	}
	return task, nil
}

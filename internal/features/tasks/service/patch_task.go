package tasks_service

import (
	"context"
	"fmt"
	"todolist/internal/core/domain"

	"github.com/google/uuid"
)

func (s *TasksService) PatchTask(
	ctx context.Context,
	ID uuid.UUID,
	patch domain.TaskPatch,
) (domain.Task, error) {
	task, err := s.tasksRepository.GetTask(ctx, ID)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get task from repository: %w", err)
	}

	if err := task.ApplyPatch(patch); err != nil {
		return domain.Task{}, fmt.Errorf("apply patch to task: %w", err)
	}

	taskDomain, err := s.tasksRepository.PatchTask(ctx, ID, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("update task in repository: %w", err)
	}
	return taskDomain, nil
}

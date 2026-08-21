package tasks_service

import (
	"context"
	"todolist/internal/core/domain"

	"github.com/google/uuid"
)

type TasksService struct {
	tasksRepository TasksRepository
}

type TasksRepository interface {
	CreateTask(
		ctx context.Context,
		task domain.Task,
	) (domain.Task, error)
	GetTasks(
		ctx context.Context,
		userID *uuid.UUID,
		limit *int,
		offset *int,
	) ([]domain.Task, error)
	GetTask(
		ctx context.Context,
		userID uuid.UUID,
	) (domain.Task, error)
	DeleteTask(
		ctx context.Context,
		ID uuid.UUID,
	) error
	PatchTask(
		ctx context.Context,
		ID uuid.UUID,
		patch domain.Task,
	) (domain.Task, error)
}

func NewTasksService(
	tasksRepository TasksRepository,
) *TasksService {
	return &TasksService{
		tasksRepository: tasksRepository,
	}
}

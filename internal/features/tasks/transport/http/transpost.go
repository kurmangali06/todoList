package tasks_transport_http

import (
	"context"
	"todolist/internal/core/domain"
)

type TaskHTTPHandler struct {
	tasksService TaskService
}

type TaskService interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
}

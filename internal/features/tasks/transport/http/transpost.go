package tasks_transport_http

import (
	"context"
	"net/http"
	"todolist/internal/core/domain"
	corre_http_server "todolist/internal/core/transport/http/server"

	"github.com/google/uuid"
)

type TaskHTTPHandler struct {
	tasksService TasksService
}

type TasksService interface {
	CreateTask(
		ctx context.Context,
		title string,
		description *string,
		authorUserID uuid.UUID,
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
		patch domain.TaskPatch,
	) (domain.Task, error)
}

func NewTasksHTTPHandler(
	tasksService TasksService,
) *TaskHTTPHandler {
	return &TaskHTTPHandler{
		tasksService: tasksService,
	}
}

func (h *TaskHTTPHandler) Routes() []corre_http_server.Route {
	return []corre_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/tasks",
			Handler: h.CreateTask,
		},
		{
			Method:  http.MethodGet,
			Path:    "/tasks",
			Handler: h.GetTasks,
		},
		{
			Method:  http.MethodGet,
			Path:    "/tasks/{id}",
			Handler: h.GetTask,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/tasks/{id}",
			Handler: h.DeleteTask,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/tasks/{id}",
			Handler: h.PatchTask,
		},
	}
}

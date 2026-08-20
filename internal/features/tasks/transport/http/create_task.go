package tasks_transport_http

import (
	"net/http"
	core_logger "todolist/internal/core/logger"
	core_http_request "todolist/internal/core/transport/http/request"
	core_http_response "todolist/internal/core/transport/http/response"

	"github.com/google/uuid"
)

type CreateTaskRequest struct {
	Title        string    `json:"title" validate:"required,min=1,max=100"         example:"Домашнее задание"`
	Description  *string   `json:"description" validate:"omitempty,min=1,max=1000" example:"Сделать до четверга домашнее задание по математике"`
	AuthorUserID uuid.UUID `json:"author_user_id" validate:"required"              example:"550e8400-e29b-41d4-a716-446655440000"`
}

type CreateTaskResponse TaskDTOResponse

func (h *TaskHTTPHandler) CreateTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)
	log.Debug("create tasks")
	var req CreateTaskRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
	}

	taskDomain, err := h.tasksService.CreateTask(
		ctx,
		req.Title,
		req.Description,
		req.AuthorUserID,
	)

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create task")
	}
	response := CreateTaskResponse(taskDTOFromDomain(taskDomain))
	responseHandler.JsonResponse(response, http.StatusCreated)

}

package tasks_transport_http

import (
	"net/http"
	core_logger "todolist/internal/core/logger"
	core_http_request "todolist/internal/core/transport/http/request"
	core_http_response "todolist/internal/core/transport/http/response"
)

type GetTaskResponse TaskDTOResponse

func (h *TaskHTTPHandler) GetTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)
	taskID, err := core_http_request.GetUUIDPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get task id")
		return
	}

	taskDomain, err := h.tasksService.GetTask(ctx, taskID)

	response := GetTaskResponse(taskDTOFromDomain(taskDomain))

	responseHandler.JsonResponse(response, http.StatusOK)
}

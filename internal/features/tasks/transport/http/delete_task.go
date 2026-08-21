package tasks_transport_http

import (
	"net/http"
	core_logger "todolist/internal/core/logger"
	core_http_request "todolist/internal/core/transport/http/request"
	core_http_response "todolist/internal/core/transport/http/response"
)

func (h *TaskHTTPHandler) DeleteTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	taskID, err := core_http_request.GetUUIDPathValue(r, "id")

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get task id")
		return
	}

	err = h.tasksService.DeleteTask(ctx, taskID)

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to delete task")
		return
	}
	responseHandler.NoContentResponse()

}

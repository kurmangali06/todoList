package tasks_transport_http

import (
	"net/http"
	core_logger "todolist/internal/core/logger"
	core_http_response "todolist/internal/core/transport/http/response"
)

func (h *TaskHTTPHandler) CreateTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)
	log.Debug("create tasks")
	responseHandler.NoContentResponse()

}

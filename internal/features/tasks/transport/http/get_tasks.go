package tasks_transport_http

import (
	"fmt"
	"net/http"
	core_logger "todolist/internal/core/logger"
	core_http_request "todolist/internal/core/transport/http/request"
	core_http_response "todolist/internal/core/transport/http/response"

	"github.com/google/uuid"
)

type GetTasksResponse []TaskDTOResponse

func (h *TaskHTTPHandler) GetTasks(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, limit, offset, err := getUserIDLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get limit and offset query params")
		return
	}

	tasksDomin, err := h.tasksService.GetTasks(ctx, userID, limit, offset)

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get tasks")
		return
	}

	res := taskDTOsFromDomains(tasksDomin)
	responseHandler.JsonResponse(
		GetTasksResponse(res),
		http.StatusOK,
	)

}

func getUserIDLimitOffsetQueryParams(r *http.Request) (*uuid.UUID, *int, *int, error) {
	const (
		userIDQueryParamKey = "user_id"
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
	)
	userID, err := core_http_request.GetUUIDQueryParam(r, userIDQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'user_id' query param: %w", err)
	}
	limit, err := core_http_request.GetIntQueryParam(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}
	offset, err := core_http_request.GetIntQueryParam(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'offset' query param: %w", err)
	}
	return userID, limit, offset, nil
}

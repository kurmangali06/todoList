package users_transport_http

import (
	"fmt"
	"net/http"
	core_logger "todolist/internal/core/logger"
	core_http_response "todolist/internal/core/transport/http/response"
	core_http_utils "todolist/internal/core/transport/http/utils"
)

type GetUsersResponse []UserDTOResponse

func (h *UserHTTPHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get limit and offset query params")
		return
	}
	userDomains, err := h.usersService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get users")
		return
	}
	response := GetUsersResponse(usersDTOFromDomains(userDomains))
	responseHandler.JsonResponse(
		response,
		http.StatusOK,
	)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	limit, err := core_http_utils.GetIntQueryParams(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}
	offset, err := core_http_utils.GetIntQueryParams(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("get 'offset' query param: %w", err)
	}
	return limit, offset, nil
}

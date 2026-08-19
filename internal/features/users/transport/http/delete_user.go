package users_transport_http

import (
	"net/http"
	core_logger "todolist/internal/core/logger"
	core_http_response "todolist/internal/core/transport/http/response"
	core_http_utils "todolist/internal/core/transport/http/utils"
)

func (h *UserHTTPHandler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_utils.GetUUIDPathValue(r, "id")

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userID path value")
		return
	}
	err = h.usersService.DeleteUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to delete user")
		return
	}
	responseHandler.NoContentResponse()
}

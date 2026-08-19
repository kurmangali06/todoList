package users_transport_http

import (
	"fmt"
	"net/http"
	"strings"
	"todolist/internal/core/domain"
	core_logger "todolist/internal/core/logger"
	core_http_request "todolist/internal/core/transport/http/request"
	core_http_response "todolist/internal/core/transport/http/response"
	core_http_types "todolist/internal/core/transport/http/types"
	core_http_utils "todolist/internal/core/transport/http/utils"
)

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number"`
}

func (r *PatchUserRequest) Validate() error {
	if r.FullName.Set {
		if r.FullName.Value == nil {
			return fmt.Errorf("`FullName` can't be NULL")
		}

		fullNameLen := len([]rune(*r.FullName.Value))
		if fullNameLen < 3 || fullNameLen > 100 {
			return fmt.Errorf("`FullName` must be between 3 and 100 symbols")
		}
	}

	if r.PhoneNumber.Set {
		if r.PhoneNumber.Value != nil {
			phoneNumberLen := len([]rune(*r.PhoneNumber.Value))
			if phoneNumberLen < 10 || phoneNumberLen > 15 {
				return fmt.Errorf("`PhoneNumber` must be between 10 and 15 symbols")
			}

			if !strings.HasPrefix(*r.PhoneNumber.Value, "+") {
				return fmt.Errorf("`PhoneNumber` must startswith '+' symbol")
			}
		}
	}

	return nil
}

type PatchUserResponse UserDTOResponse

func (h *UserHTTPHandler) PatchUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)
	userID, err := core_http_utils.GetUUIDPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID path value",
		)
		return
	}
	var req PatchUserRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	userPatch := userPatchFromRequest(req)

	userDomain, err := h.usersService.PatchUser(ctx, userID, userPatch)

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch user")
		return
	}
	response := PatchUserResponse(userDTOFromDomain(userDomain))
	responseHandler.JsonResponse(response, http.StatusOK)
}

func userPatchFromRequest(req PatchUserRequest) domain.UserPatch {
	return domain.UserPatch{
		FullName:    req.FullName.ToDomain(),
		PhoneNumber: req.PhoneNumber.ToDomain(),
	}
}

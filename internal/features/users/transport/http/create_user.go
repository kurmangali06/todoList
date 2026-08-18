package users_transport_http

import (
	"net/http"
	"todolist/internal/core/domain"
	core_logger "todolist/internal/core/logger"
	core_http_request "todolist/internal/core/transport/http/request"
	core_http_response "todolist/internal/core/transport/http/response"

	"github.com/google/uuid"
)

type CreateUserRequest struct {
	FullName    string `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNumber string `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+"`
}

type CreateUserResponse struct {
	ID          uuid.UUID `json:"id"`
	FullName    string    `json:"full_name"`
	PhoneNumber *string   `json:"phone_number"`
	Version     int       `json:"version"`
}

func (h *UserHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)
	log.Debug("create user")

	var req CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")

		return
	}
	userDomain := domainFromDTO(req)
	user, err := h.usersService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return
	}
	response := dtoFromDomain(user)
	responseHandler.JsonResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	var phoneNumber *string
	if dto.PhoneNumber != "" {
		phoneNumber = &dto.PhoneNumber
	}

	return domain.NewUser(
		dto.FullName,
		phoneNumber,
		uuid.New(),
		-1)
}

func dtoFromDomain(user domain.User) CreateUserResponse {
	return CreateUserResponse{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}

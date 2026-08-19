package users_transport_http

import (
	"context"
	"net/http"
	"todolist/internal/core/domain"
	corre_http_server "todolist/internal/core/transport/http/server"

	"github.com/google/uuid"
)

type UserHTTPHandler struct {
	usersService UsersService
}

type UsersService interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)
	GetUser(
		ctx context.Context,
		ID uuid.UUID,
	) (domain.User, error)
	DeleteUser(
		ctx context.Context,
		ID uuid.UUID,
	) error
	PatchUser(
		ctx context.Context,
		ID uuid.UUID,
		user domain.UserPatch,
	) (domain.User, error)
}

func NewUsersHTTPHandler(usersService UsersService) *UserHTTPHandler {
	return &UserHTTPHandler{usersService: usersService}
}

func (h *UserHTTPHandler) Routes() []corre_http_server.Route {
	return []corre_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: h.GetUsers,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: h.GetUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/{id}",
			Handler: h.DeleteUser,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users/{id}",
			Handler: h.PatchUser,
		},
	}
}
